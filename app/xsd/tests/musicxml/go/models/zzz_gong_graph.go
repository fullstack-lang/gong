// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (a_directive *A_directive) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_directives[a_directive]
	return ok
}

func (a_measure *A_measure) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_measures[a_measure]
	return ok
}

func (a_measure_1 *A_measure_1) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_measure_1s[a_measure_1]
	return ok
}

func (a_part *A_part) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_parts[a_part]
	return ok
}

func (a_part_1 *A_part_1) GongIsStaged(stage *Stage) bool {
	_, ok := stage.A_part_1s[a_part_1]
	return ok
}

func (accidental *Accidental) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Accidentals[accidental]
	return ok
}

func (accidental_mark *Accidental_mark) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Accidental_marks[accidental_mark]
	return ok
}

func (accidental_text *Accidental_text) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Accidental_texts[accidental_text]
	return ok
}

func (accord *Accord) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Accords[accord]
	return ok
}

func (accordion_registration *Accordion_registration) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Accordion_registrations[accordion_registration]
	return ok
}

func (appearance *Appearance) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Appearances[appearance]
	return ok
}

func (arpeggiate *Arpeggiate) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Arpeggiates[arpeggiate]
	return ok
}

func (arrow *Arrow) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Arrows[arrow]
	return ok
}

func (articulations *Articulations) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Articulationss[articulations]
	return ok
}

func (assess *Assess) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Assesss[assess]
	return ok
}

func (attributes *Attributes) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Attributess[attributes]
	return ok
}

func (backup *Backup) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Backups[backup]
	return ok
}

func (bar_style_color *Bar_style_color) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Bar_style_colors[bar_style_color]
	return ok
}

func (barline *Barline) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Barlines[barline]
	return ok
}

func (barre *Barre) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Barres[barre]
	return ok
}

func (bass *Bass) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Basss[bass]
	return ok
}

func (bass_step *Bass_step) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Bass_steps[bass_step]
	return ok
}

func (beam *Beam) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Beams[beam]
	return ok
}

func (beat_repeat *Beat_repeat) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Beat_repeats[beat_repeat]
	return ok
}

func (beat_unit_tied *Beat_unit_tied) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Beat_unit_tieds[beat_unit_tied]
	return ok
}

func (beater *Beater) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Beaters[beater]
	return ok
}

func (bend *Bend) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Bends[bend]
	return ok
}

func (bookmark *Bookmark) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Bookmarks[bookmark]
	return ok
}

func (bracket *Bracket) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Brackets[bracket]
	return ok
}

func (breath_mark *Breath_mark) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Breath_marks[breath_mark]
	return ok
}

func (caesura *Caesura) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Caesuras[caesura]
	return ok
}

func (cancel *Cancel) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Cancels[cancel]
	return ok
}

func (clef *Clef) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Clefs[clef]
	return ok
}

func (coda *Coda) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Codas[coda]
	return ok
}

func (credit *Credit) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Credits[credit]
	return ok
}

func (dashes *Dashes) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Dashess[dashes]
	return ok
}

func (defaults *Defaults) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Defaultss[defaults]
	return ok
}

func (degree *Degree) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Degrees[degree]
	return ok
}

func (degree_alter *Degree_alter) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Degree_alters[degree_alter]
	return ok
}

func (degree_type *Degree_type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Degree_types[degree_type]
	return ok
}

func (degree_value *Degree_value) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Degree_values[degree_value]
	return ok
}

func (direction *Direction) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Directions[direction]
	return ok
}

func (direction_type *Direction_type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Direction_types[direction_type]
	return ok
}

func (distance *Distance) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Distances[distance]
	return ok
}

func (double *Double) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Doubles[double]
	return ok
}

func (dynamics *Dynamics) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Dynamicss[dynamics]
	return ok
}

func (effect *Effect) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Effects[effect]
	return ok
}

func (elision *Elision) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Elisions[elision]
	return ok
}

func (empty *Empty) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Emptys[empty]
	return ok
}

func (empty_font *Empty_font) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Empty_fonts[empty_font]
	return ok
}

func (empty_line *Empty_line) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Empty_lines[empty_line]
	return ok
}

func (empty_placement *Empty_placement) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Empty_placements[empty_placement]
	return ok
}

func (empty_placement_smufl *Empty_placement_smufl) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Empty_placement_smufls[empty_placement_smufl]
	return ok
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Empty_print_object_style_aligns[empty_print_object_style_align]
	return ok
}

func (empty_print_style *Empty_print_style) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Empty_print_styles[empty_print_style]
	return ok
}

func (empty_print_style_align *Empty_print_style_align) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Empty_print_style_aligns[empty_print_style_align]
	return ok
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Empty_print_style_align_ids[empty_print_style_align_id]
	return ok
}

func (empty_trill_sound *Empty_trill_sound) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Empty_trill_sounds[empty_trill_sound]
	return ok
}

func (encoding *Encoding) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Encodings[encoding]
	return ok
}

func (ending *Ending) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Endings[ending]
	return ok
}

func (extend *Extend) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Extends[extend]
	return ok
}

func (feature *Feature) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Features[feature]
	return ok
}

func (fermata *Fermata) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Fermatas[fermata]
	return ok
}

func (figure *Figure) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Figures[figure]
	return ok
}

func (figured_bass *Figured_bass) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Figured_basss[figured_bass]
	return ok
}

func (fingering *Fingering) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Fingerings[fingering]
	return ok
}

func (first_fret *First_fret) GongIsStaged(stage *Stage) bool {
	_, ok := stage.First_frets[first_fret]
	return ok
}

func (for_part *For_part) GongIsStaged(stage *Stage) bool {
	_, ok := stage.For_parts[for_part]
	return ok
}

func (formatted_symbol *Formatted_symbol) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Formatted_symbols[formatted_symbol]
	return ok
}

func (formatted_symbol_id *Formatted_symbol_id) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Formatted_symbol_ids[formatted_symbol_id]
	return ok
}

func (formatted_text *Formatted_text) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Formatted_texts[formatted_text]
	return ok
}

func (formatted_text_id *Formatted_text_id) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Formatted_text_ids[formatted_text_id]
	return ok
}

func (forward *Forward) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Forwards[forward]
	return ok
}

func (frame *Frame) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Frames[frame]
	return ok
}

func (frame_note *Frame_note) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Frame_notes[frame_note]
	return ok
}

func (fret *Fret) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Frets[fret]
	return ok
}

func (glass *Glass) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Glasss[glass]
	return ok
}

func (glissando *Glissando) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Glissandos[glissando]
	return ok
}

func (glyph *Glyph) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Glyphs[glyph]
	return ok
}

func (grace *Grace) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Graces[grace]
	return ok
}

func (group_barline *Group_barline) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Group_barlines[group_barline]
	return ok
}

func (group_name *Group_name) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Group_names[group_name]
	return ok
}

func (group_symbol *Group_symbol) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Group_symbols[group_symbol]
	return ok
}

func (grouping *Grouping) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Groupings[grouping]
	return ok
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Hammer_on_pull_offs[hammer_on_pull_off]
	return ok
}

func (handbell *Handbell) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Handbells[handbell]
	return ok
}

func (harmon_closed *Harmon_closed) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Harmon_closeds[harmon_closed]
	return ok
}

func (harmon_mute *Harmon_mute) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Harmon_mutes[harmon_mute]
	return ok
}

func (harmonic *Harmonic) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Harmonics[harmonic]
	return ok
}

func (harmony *Harmony) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Harmonys[harmony]
	return ok
}

func (harmony_alter *Harmony_alter) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Harmony_alters[harmony_alter]
	return ok
}

func (harp_pedals *Harp_pedals) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Harp_pedalss[harp_pedals]
	return ok
}

func (heel_toe *Heel_toe) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Heel_toes[heel_toe]
	return ok
}

func (hole *Hole) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Holes[hole]
	return ok
}

func (hole_closed *Hole_closed) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Hole_closeds[hole_closed]
	return ok
}

func (horizontal_turn *Horizontal_turn) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Horizontal_turns[horizontal_turn]
	return ok
}

func (identification *Identification) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Identifications[identification]
	return ok
}

func (image *Image) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Images[image]
	return ok
}

func (instrument *Instrument) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Instruments[instrument]
	return ok
}

func (instrument_change *Instrument_change) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Instrument_changes[instrument_change]
	return ok
}

func (instrument_link *Instrument_link) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Instrument_links[instrument_link]
	return ok
}

func (interchangeable *Interchangeable) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Interchangeables[interchangeable]
	return ok
}

func (inversion *Inversion) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Inversions[inversion]
	return ok
}

func (key *Key) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Keys[key]
	return ok
}

func (key_accidental *Key_accidental) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Key_accidentals[key_accidental]
	return ok
}

func (key_octave *Key_octave) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Key_octaves[key_octave]
	return ok
}

func (kind *Kind) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Kinds[kind]
	return ok
}

func (level *Level) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Levels[level]
	return ok
}

func (line_detail *Line_detail) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Line_details[line_detail]
	return ok
}

func (line_width *Line_width) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Line_widths[line_width]
	return ok
}

func (link *Link) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Links[link]
	return ok
}

func (listen *Listen) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Listens[listen]
	return ok
}

func (listening *Listening) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Listenings[listening]
	return ok
}

func (lyric *Lyric) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Lyrics[lyric]
	return ok
}

func (lyric_font *Lyric_font) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Lyric_fonts[lyric_font]
	return ok
}

func (lyric_language *Lyric_language) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Lyric_languages[lyric_language]
	return ok
}

func (measure_layout *Measure_layout) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Measure_layouts[measure_layout]
	return ok
}

func (measure_numbering *Measure_numbering) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Measure_numberings[measure_numbering]
	return ok
}

func (measure_repeat *Measure_repeat) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Measure_repeats[measure_repeat]
	return ok
}

func (measure_style *Measure_style) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Measure_styles[measure_style]
	return ok
}

func (membrane *Membrane) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Membranes[membrane]
	return ok
}

func (metal *Metal) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Metals[metal]
	return ok
}

func (metronome *Metronome) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Metronomes[metronome]
	return ok
}

func (metronome_beam *Metronome_beam) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Metronome_beams[metronome_beam]
	return ok
}

func (metronome_note *Metronome_note) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Metronome_notes[metronome_note]
	return ok
}

func (metronome_tied *Metronome_tied) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Metronome_tieds[metronome_tied]
	return ok
}

func (metronome_tuplet *Metronome_tuplet) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Metronome_tuplets[metronome_tuplet]
	return ok
}

func (midi_device *Midi_device) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Midi_devices[midi_device]
	return ok
}

func (midi_instrument *Midi_instrument) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Midi_instruments[midi_instrument]
	return ok
}

func (miscellaneous *Miscellaneous) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Miscellaneouss[miscellaneous]
	return ok
}

func (miscellaneous_field *Miscellaneous_field) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Miscellaneous_fields[miscellaneous_field]
	return ok
}

func (mordent *Mordent) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Mordents[mordent]
	return ok
}

func (multiple_rest *Multiple_rest) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Multiple_rests[multiple_rest]
	return ok
}

func (name_display *Name_display) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Name_displays[name_display]
	return ok
}

func (non_arpeggiate *Non_arpeggiate) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Non_arpeggiates[non_arpeggiate]
	return ok
}

func (notations *Notations) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Notationss[notations]
	return ok
}

func (note *Note) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Notes[note]
	return ok
}

func (note_size *Note_size) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Note_sizes[note_size]
	return ok
}

func (note_type *Note_type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Note_types[note_type]
	return ok
}

func (notehead *Notehead) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Noteheads[notehead]
	return ok
}

func (notehead_text *Notehead_text) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Notehead_texts[notehead_text]
	return ok
}

func (numeral *Numeral) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Numerals[numeral]
	return ok
}

func (numeral_key *Numeral_key) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Numeral_keys[numeral_key]
	return ok
}

func (numeral_root *Numeral_root) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Numeral_roots[numeral_root]
	return ok
}

func (octave_shift *Octave_shift) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Octave_shifts[octave_shift]
	return ok
}

func (offset *Offset) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Offsets[offset]
	return ok
}

func (opus *Opus) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Opuss[opus]
	return ok
}

func (ornaments *Ornaments) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Ornamentss[ornaments]
	return ok
}

func (other_appearance *Other_appearance) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Other_appearances[other_appearance]
	return ok
}

func (other_direction *Other_direction) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Other_directions[other_direction]
	return ok
}

func (other_listening *Other_listening) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Other_listenings[other_listening]
	return ok
}

func (other_notation *Other_notation) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Other_notations[other_notation]
	return ok
}

func (other_placement_text *Other_placement_text) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Other_placement_texts[other_placement_text]
	return ok
}

func (other_play *Other_play) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Other_plays[other_play]
	return ok
}

func (other_text *Other_text) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Other_texts[other_text]
	return ok
}

func (page_layout *Page_layout) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Page_layouts[page_layout]
	return ok
}

func (page_margins *Page_margins) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Page_marginss[page_margins]
	return ok
}

func (part_clef *Part_clef) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Part_clefs[part_clef]
	return ok
}

func (part_group *Part_group) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Part_groups[part_group]
	return ok
}

func (part_link *Part_link) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Part_links[part_link]
	return ok
}

func (part_list *Part_list) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Part_lists[part_list]
	return ok
}

func (part_name *Part_name) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Part_names[part_name]
	return ok
}

func (part_symbol *Part_symbol) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Part_symbols[part_symbol]
	return ok
}

func (part_transpose *Part_transpose) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Part_transposes[part_transpose]
	return ok
}

func (pedal *Pedal) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Pedals[pedal]
	return ok
}

func (pedal_tuning *Pedal_tuning) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Pedal_tunings[pedal_tuning]
	return ok
}

func (per_minute *Per_minute) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Per_minutes[per_minute]
	return ok
}

func (percussion *Percussion) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Percussions[percussion]
	return ok
}

func (pitch *Pitch) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Pitchs[pitch]
	return ok
}

func (pitched *Pitched) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Pitcheds[pitched]
	return ok
}

func (placement_text *Placement_text) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Placement_texts[placement_text]
	return ok
}

func (play *Play) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Plays[play]
	return ok
}

func (player *Player) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Players[player]
	return ok
}

func (principal_voice *Principal_voice) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Principal_voices[principal_voice]
	return ok
}

func (print *Print) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Prints[print]
	return ok
}

func (release *Release) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Releases[release]
	return ok
}

func (repeat *Repeat) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Repeats[repeat]
	return ok
}

func (rest *Rest) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Rests[rest]
	return ok
}

func (root *Root) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Roots[root]
	return ok
}

func (root_step *Root_step) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Root_steps[root_step]
	return ok
}

func (scaling *Scaling) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Scalings[scaling]
	return ok
}

func (scordatura *Scordatura) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Scordaturas[scordatura]
	return ok
}

func (score_instrument *Score_instrument) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Score_instruments[score_instrument]
	return ok
}

func (score_part *Score_part) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Score_parts[score_part]
	return ok
}

func (score_partwise *Score_partwise) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Score_partwises[score_partwise]
	return ok
}

func (score_timewise *Score_timewise) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Score_timewises[score_timewise]
	return ok
}

func (segno *Segno) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Segnos[segno]
	return ok
}

func (slash *Slash) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Slashs[slash]
	return ok
}

func (slide *Slide) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Slides[slide]
	return ok
}

func (slur *Slur) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Slurs[slur]
	return ok
}

func (sound *Sound) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Sounds[sound]
	return ok
}

func (staff_details *Staff_details) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Staff_detailss[staff_details]
	return ok
}

func (staff_divide *Staff_divide) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Staff_divides[staff_divide]
	return ok
}

func (staff_layout *Staff_layout) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Staff_layouts[staff_layout]
	return ok
}

func (staff_size *Staff_size) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Staff_sizes[staff_size]
	return ok
}

func (staff_tuning *Staff_tuning) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Staff_tunings[staff_tuning]
	return ok
}

func (stem *Stem) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Stems[stem]
	return ok
}

func (stick *Stick) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Sticks[stick]
	return ok
}

func (string_mute *String_mute) GongIsStaged(stage *Stage) bool {
	_, ok := stage.String_mutes[string_mute]
	return ok
}

func (string_type *String_type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.String_types[string_type]
	return ok
}

func (strong_accent *Strong_accent) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Strong_accents[strong_accent]
	return ok
}

func (style_text *Style_text) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Style_texts[style_text]
	return ok
}

func (supports *Supports) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Supportss[supports]
	return ok
}

func (swing *Swing) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Swings[swing]
	return ok
}

func (sync *Sync) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Syncs[sync]
	return ok
}

func (system_dividers *System_dividers) GongIsStaged(stage *Stage) bool {
	_, ok := stage.System_dividerss[system_dividers]
	return ok
}

func (system_layout *System_layout) GongIsStaged(stage *Stage) bool {
	_, ok := stage.System_layouts[system_layout]
	return ok
}

func (system_margins *System_margins) GongIsStaged(stage *Stage) bool {
	_, ok := stage.System_marginss[system_margins]
	return ok
}

func (tap *Tap) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Taps[tap]
	return ok
}

func (technical *Technical) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Technicals[technical]
	return ok
}

func (text_element_data *Text_element_data) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Text_element_datas[text_element_data]
	return ok
}

func (tie *Tie) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Ties[tie]
	return ok
}

func (tied *Tied) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tieds[tied]
	return ok
}

func (time *Time) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Times[time]
	return ok
}

func (time_modification *Time_modification) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Time_modifications[time_modification]
	return ok
}

func (timpani *Timpani) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Timpanis[timpani]
	return ok
}

func (transpose *Transpose) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Transposes[transpose]
	return ok
}

func (tremolo *Tremolo) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tremolos[tremolo]
	return ok
}

func (tuplet *Tuplet) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tuplets[tuplet]
	return ok
}

func (tuplet_dot *Tuplet_dot) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tuplet_dots[tuplet_dot]
	return ok
}

func (tuplet_number *Tuplet_number) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tuplet_numbers[tuplet_number]
	return ok
}

func (tuplet_portion *Tuplet_portion) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tuplet_portions[tuplet_portion]
	return ok
}

func (tuplet_type *Tuplet_type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tuplet_types[tuplet_type]
	return ok
}

func (typed_text *Typed_text) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Typed_texts[typed_text]
	return ok
}

func (unpitched *Unpitched) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Unpitcheds[unpitched]
	return ok
}

func (virtual_instrument *Virtual_instrument) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Virtual_instruments[virtual_instrument]
	return ok
}

func (wait *Wait) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Waits[wait]
	return ok
}

func (wavy_line *Wavy_line) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Wavy_lines[wavy_line]
	return ok
}

func (wedge *Wedge) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Wedges[wedge]
	return ok
}

func (wood *Wood) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Woods[wood]
	return ok
}

func (work *Work) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Works[work]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (a_directive *A_directive) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_directive) {
		return
	}

	a_directive.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_measure *A_measure) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_measure) {
		return
	}

	a_measure.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _note := range a_measure.Note {
		stage.StageBranch(_note)
	}
	for _, _backup := range a_measure.Backup {
		stage.StageBranch(_backup)
	}
	for _, _forward := range a_measure.Forward {
		stage.StageBranch(_forward)
	}
	for _, _direction := range a_measure.Direction {
		stage.StageBranch(_direction)
	}
	for _, _attributes := range a_measure.Attributes {
		stage.StageBranch(_attributes)
	}
	for _, _harmony := range a_measure.Harmony {
		stage.StageBranch(_harmony)
	}
	for _, _figured_bass := range a_measure.Figured_bass {
		stage.StageBranch(_figured_bass)
	}
	for _, _print := range a_measure.Print {
		stage.StageBranch(_print)
	}
	for _, _sound := range a_measure.Sound {
		stage.StageBranch(_sound)
	}
	for _, _listening := range a_measure.Listening {
		stage.StageBranch(_listening)
	}
	for _, _barline := range a_measure.Barline {
		stage.StageBranch(_barline)
	}
	for _, _grouping := range a_measure.Grouping {
		stage.StageBranch(_grouping)
	}
	for _, _link := range a_measure.Link {
		stage.StageBranch(_link)
	}
	for _, _bookmark := range a_measure.Bookmark {
		stage.StageBranch(_bookmark)
	}

}

func (a_measure_1 *A_measure_1) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_measure_1) {
		return
	}

	a_measure_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_part_1 := range a_measure_1.Part {
		stage.StageBranch(_a_part_1)
	}

}

func (a_part *A_part) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_part) {
		return
	}

	a_part.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_measure := range a_part.Measure {
		stage.StageBranch(_a_measure)
	}

}

func (a_part_1 *A_part_1) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(a_part_1) {
		return
	}

	a_part_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _note := range a_part_1.Note {
		stage.StageBranch(_note)
	}
	for _, _backup := range a_part_1.Backup {
		stage.StageBranch(_backup)
	}
	for _, _forward := range a_part_1.Forward {
		stage.StageBranch(_forward)
	}
	for _, _direction := range a_part_1.Direction {
		stage.StageBranch(_direction)
	}
	for _, _attributes := range a_part_1.Attributes {
		stage.StageBranch(_attributes)
	}
	for _, _harmony := range a_part_1.Harmony {
		stage.StageBranch(_harmony)
	}
	for _, _figured_bass := range a_part_1.Figured_bass {
		stage.StageBranch(_figured_bass)
	}
	for _, _print := range a_part_1.Print {
		stage.StageBranch(_print)
	}
	for _, _sound := range a_part_1.Sound {
		stage.StageBranch(_sound)
	}
	for _, _listening := range a_part_1.Listening {
		stage.StageBranch(_listening)
	}
	for _, _barline := range a_part_1.Barline {
		stage.StageBranch(_barline)
	}
	for _, _grouping := range a_part_1.Grouping {
		stage.StageBranch(_grouping)
	}
	for _, _link := range a_part_1.Link {
		stage.StageBranch(_link)
	}
	for _, _bookmark := range a_part_1.Bookmark {
		stage.StageBranch(_bookmark)
	}

}

func (accidental *Accidental) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(accidental) {
		return
	}

	accidental.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (accidental_mark *Accidental_mark) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(accidental_mark) {
		return
	}

	accidental_mark.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (accidental_text *Accidental_text) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(accidental_text) {
		return
	}

	accidental_text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (accord *Accord) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(accord) {
		return
	}

	accord.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (accordion_registration *Accordion_registration) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(accordion_registration) {
		return
	}

	accordion_registration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (appearance *Appearance) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(appearance) {
		return
	}

	appearance.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _line_width := range appearance.Line_width {
		stage.StageBranch(_line_width)
	}
	for _, _note_size := range appearance.Note_size {
		stage.StageBranch(_note_size)
	}
	for _, _distance := range appearance.Distance {
		stage.StageBranch(_distance)
	}
	for _, _glyph := range appearance.Glyph {
		stage.StageBranch(_glyph)
	}
	for _, _other_appearance := range appearance.Other_appearance {
		stage.StageBranch(_other_appearance)
	}

}

func (arpeggiate *Arpeggiate) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(arpeggiate) {
		return
	}

	arpeggiate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arrow *Arrow) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(arrow) {
		return
	}

	arrow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (articulations *Articulations) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(articulations) {
		return
	}

	articulations.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty_placement := range articulations.Accent {
		stage.StageBranch(_empty_placement)
	}
	for _, _strong_accent := range articulations.Strong_accent {
		stage.StageBranch(_strong_accent)
	}
	for _, _empty_placement := range articulations.Staccato {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_placement := range articulations.Tenuto {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_placement := range articulations.Detached_legato {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_placement := range articulations.Staccatissimo {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_placement := range articulations.Spiccato {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_line := range articulations.Scoop {
		stage.StageBranch(_empty_line)
	}
	for _, _empty_line := range articulations.Plop {
		stage.StageBranch(_empty_line)
	}
	for _, _empty_line := range articulations.Doit {
		stage.StageBranch(_empty_line)
	}
	for _, _empty_line := range articulations.Falloff {
		stage.StageBranch(_empty_line)
	}
	for _, _breath_mark := range articulations.Breath_mark {
		stage.StageBranch(_breath_mark)
	}
	for _, _caesura := range articulations.Caesura {
		stage.StageBranch(_caesura)
	}
	for _, _empty_placement := range articulations.Stress {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_placement := range articulations.Unstress {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_placement := range articulations.Soft_accent {
		stage.StageBranch(_empty_placement)
	}
	for _, _other_placement_text := range articulations.Other_articulation {
		stage.StageBranch(_other_placement_text)
	}

}

func (assess *Assess) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(assess) {
		return
	}

	assess.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attributes *Attributes) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(attributes) {
		return
	}

	attributes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributes.Footnote != nil {
		stage.StageBranch(attributes.Footnote)
	}
	if attributes.Level != nil {
		stage.StageBranch(attributes.Level)
	}
	if attributes.Part_symbol != nil {
		stage.StageBranch(attributes.Part_symbol)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key := range attributes.Key {
		stage.StageBranch(_key)
	}
	for _, _time := range attributes.Time {
		stage.StageBranch(_time)
	}
	for _, _clef := range attributes.Clef {
		stage.StageBranch(_clef)
	}
	for _, _staff_details := range attributes.Staff_details {
		stage.StageBranch(_staff_details)
	}
	for _, _transpose := range attributes.Transpose {
		stage.StageBranch(_transpose)
	}
	for _, _for_part := range attributes.For_part {
		stage.StageBranch(_for_part)
	}
	for _, _a_directive := range attributes.Directive {
		stage.StageBranch(_a_directive)
	}
	for _, _measure_style := range attributes.Measure_style {
		stage.StageBranch(_measure_style)
	}

}

func (backup *Backup) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(backup) {
		return
	}

	backup.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if backup.Footnote != nil {
		stage.StageBranch(backup.Footnote)
	}
	if backup.Level != nil {
		stage.StageBranch(backup.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bar_style_color *Bar_style_color) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bar_style_color) {
		return
	}

	bar_style_color.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (barline *Barline) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(barline) {
		return
	}

	barline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if barline.Bar_style != nil {
		stage.StageBranch(barline.Bar_style)
	}
	if barline.Footnote != nil {
		stage.StageBranch(barline.Footnote)
	}
	if barline.Level != nil {
		stage.StageBranch(barline.Level)
	}
	if barline.Wavy_line != nil {
		stage.StageBranch(barline.Wavy_line)
	}
	if barline.Segno_1 != nil {
		stage.StageBranch(barline.Segno_1)
	}
	if barline.Coda_1 != nil {
		stage.StageBranch(barline.Coda_1)
	}
	if barline.Fermata != nil {
		stage.StageBranch(barline.Fermata)
	}
	if barline.Ending != nil {
		stage.StageBranch(barline.Ending)
	}
	if barline.Repeat != nil {
		stage.StageBranch(barline.Repeat)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (barre *Barre) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(barre) {
		return
	}

	barre.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bass *Bass) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bass) {
		return
	}

	bass.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bass.Bass_separator != nil {
		stage.StageBranch(bass.Bass_separator)
	}
	if bass.Bass_step != nil {
		stage.StageBranch(bass.Bass_step)
	}
	if bass.Bass_alter != nil {
		stage.StageBranch(bass.Bass_alter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bass_step *Bass_step) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bass_step) {
		return
	}

	bass_step.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beam *Beam) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(beam) {
		return
	}

	beam.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beat_repeat *Beat_repeat) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(beat_repeat) {
		return
	}

	beat_repeat.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beat_unit_tied *Beat_unit_tied) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(beat_unit_tied) {
		return
	}

	beat_unit_tied.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beater *Beater) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(beater) {
		return
	}

	beater.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bend *Bend) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bend) {
		return
	}

	bend.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bend.Release != nil {
		stage.StageBranch(bend.Release)
	}
	if bend.With_bar != nil {
		stage.StageBranch(bend.With_bar)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bookmark *Bookmark) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bookmark) {
		return
	}

	bookmark.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bracket *Bracket) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bracket) {
		return
	}

	bracket.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (breath_mark *Breath_mark) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(breath_mark) {
		return
	}

	breath_mark.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (caesura *Caesura) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(caesura) {
		return
	}

	caesura.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cancel *Cancel) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(cancel) {
		return
	}

	cancel.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clef *Clef) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(clef) {
		return
	}

	clef.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coda *Coda) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coda) {
		return
	}

	coda.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (credit *Credit) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(credit) {
		return
	}

	credit.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if credit.Credit_image != nil {
		stage.StageBranch(credit.Credit_image)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range credit.Link {
		stage.StageBranch(_link)
	}
	for _, _bookmark := range credit.Bookmark {
		stage.StageBranch(_bookmark)
	}
	for _, _formatted_text_id := range credit.Credit_words {
		stage.StageBranch(_formatted_text_id)
	}
	for _, _formatted_symbol_id := range credit.Credit_symbol {
		stage.StageBranch(_formatted_symbol_id)
	}

}

func (dashes *Dashes) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(dashes) {
		return
	}

	dashes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (defaults *Defaults) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(defaults) {
		return
	}

	defaults.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if defaults.Scaling != nil {
		stage.StageBranch(defaults.Scaling)
	}
	if defaults.Page_layout != nil {
		stage.StageBranch(defaults.Page_layout)
	}
	if defaults.System_layout != nil {
		stage.StageBranch(defaults.System_layout)
	}
	if defaults.Appearance != nil {
		stage.StageBranch(defaults.Appearance)
	}
	if defaults.Music_font != nil {
		stage.StageBranch(defaults.Music_font)
	}
	if defaults.Word_font != nil {
		stage.StageBranch(defaults.Word_font)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _staff_layout := range defaults.Staff_layout {
		stage.StageBranch(_staff_layout)
	}
	for _, _lyric_font := range defaults.Lyric_font {
		stage.StageBranch(_lyric_font)
	}
	for _, _lyric_language := range defaults.Lyric_language {
		stage.StageBranch(_lyric_language)
	}

}

func (degree *Degree) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(degree) {
		return
	}

	degree.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if degree.Degree_value != nil {
		stage.StageBranch(degree.Degree_value)
	}
	if degree.Degree_alter != nil {
		stage.StageBranch(degree.Degree_alter)
	}
	if degree.Degree_type != nil {
		stage.StageBranch(degree.Degree_type)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degree_alter *Degree_alter) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(degree_alter) {
		return
	}

	degree_alter.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degree_type *Degree_type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(degree_type) {
		return
	}

	degree_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degree_value *Degree_value) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(degree_value) {
		return
	}

	degree_value.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (direction *Direction) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(direction) {
		return
	}

	direction.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if direction.Offset != nil {
		stage.StageBranch(direction.Offset)
	}
	if direction.Footnote != nil {
		stage.StageBranch(direction.Footnote)
	}
	if direction.Level != nil {
		stage.StageBranch(direction.Level)
	}
	if direction.Sound != nil {
		stage.StageBranch(direction.Sound)
	}
	if direction.Listening != nil {
		stage.StageBranch(direction.Listening)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _direction_type := range direction.Direction_type {
		stage.StageBranch(_direction_type)
	}

}

func (direction_type *Direction_type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(direction_type) {
		return
	}

	direction_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if direction_type.Wedge != nil {
		stage.StageBranch(direction_type.Wedge)
	}
	if direction_type.Dashes != nil {
		stage.StageBranch(direction_type.Dashes)
	}
	if direction_type.Bracket != nil {
		stage.StageBranch(direction_type.Bracket)
	}
	if direction_type.Pedal != nil {
		stage.StageBranch(direction_type.Pedal)
	}
	if direction_type.Metronome != nil {
		stage.StageBranch(direction_type.Metronome)
	}
	if direction_type.Octave_shift != nil {
		stage.StageBranch(direction_type.Octave_shift)
	}
	if direction_type.Harp_pedals != nil {
		stage.StageBranch(direction_type.Harp_pedals)
	}
	if direction_type.Damp != nil {
		stage.StageBranch(direction_type.Damp)
	}
	if direction_type.Damp_all != nil {
		stage.StageBranch(direction_type.Damp_all)
	}
	if direction_type.Eyeglasses != nil {
		stage.StageBranch(direction_type.Eyeglasses)
	}
	if direction_type.String_mute != nil {
		stage.StageBranch(direction_type.String_mute)
	}
	if direction_type.Scordatura != nil {
		stage.StageBranch(direction_type.Scordatura)
	}
	if direction_type.Image != nil {
		stage.StageBranch(direction_type.Image)
	}
	if direction_type.Principal_voice != nil {
		stage.StageBranch(direction_type.Principal_voice)
	}
	if direction_type.Accordion_registration != nil {
		stage.StageBranch(direction_type.Accordion_registration)
	}
	if direction_type.Staff_divide != nil {
		stage.StageBranch(direction_type.Staff_divide)
	}
	if direction_type.Other_direction != nil {
		stage.StageBranch(direction_type.Other_direction)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formatted_text_id := range direction_type.Rehearsal {
		stage.StageBranch(_formatted_text_id)
	}
	for _, _segno := range direction_type.Segno {
		stage.StageBranch(_segno)
	}
	for _, _coda := range direction_type.Coda {
		stage.StageBranch(_coda)
	}
	for _, _formatted_text_id := range direction_type.Words {
		stage.StageBranch(_formatted_text_id)
	}
	for _, _formatted_symbol_id := range direction_type.Symbol {
		stage.StageBranch(_formatted_symbol_id)
	}
	for _, _dynamics := range direction_type.Dynamics {
		stage.StageBranch(_dynamics)
	}
	for _, _percussion := range direction_type.Percussion {
		stage.StageBranch(_percussion)
	}

}

func (distance *Distance) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(distance) {
		return
	}

	distance.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (double *Double) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(double) {
		return
	}

	double.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dynamics *Dynamics) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(dynamics) {
		return
	}

	dynamics.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _other_text := range dynamics.Other_dynamics {
		stage.StageBranch(_other_text)
	}

}

func (effect *Effect) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(effect) {
		return
	}

	effect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (elision *Elision) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(elision) {
		return
	}

	elision.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty *Empty) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(empty) {
		return
	}

	empty.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_font *Empty_font) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(empty_font) {
		return
	}

	empty_font.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_line *Empty_line) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(empty_line) {
		return
	}

	empty_line.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_placement *Empty_placement) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(empty_placement) {
		return
	}

	empty_placement.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_placement_smufl *Empty_placement_smufl) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(empty_placement_smufl) {
		return
	}

	empty_placement_smufl.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_print_object_style_align *Empty_print_object_style_align) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(empty_print_object_style_align) {
		return
	}

	empty_print_object_style_align.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_print_style *Empty_print_style) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(empty_print_style) {
		return
	}

	empty_print_style.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_print_style_align *Empty_print_style_align) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(empty_print_style_align) {
		return
	}

	empty_print_style_align.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_print_style_align_id *Empty_print_style_align_id) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(empty_print_style_align_id) {
		return
	}

	empty_print_style_align_id.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_trill_sound *Empty_trill_sound) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(empty_trill_sound) {
		return
	}

	empty_trill_sound.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (encoding *Encoding) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(encoding) {
		return
	}

	encoding.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typed_text := range encoding.Encoder {
		stage.StageBranch(_typed_text)
	}
	for _, _supports := range encoding.Supports {
		stage.StageBranch(_supports)
	}

}

func (ending *Ending) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ending) {
		return
	}

	ending.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (extend *Extend) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(extend) {
		return
	}

	extend.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (feature *Feature) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(feature) {
		return
	}

	feature.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (fermata *Fermata) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(fermata) {
		return
	}

	fermata.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (figure *Figure) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(figure) {
		return
	}

	figure.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if figure.Prefix != nil {
		stage.StageBranch(figure.Prefix)
	}
	if figure.Figure_number != nil {
		stage.StageBranch(figure.Figure_number)
	}
	if figure.Suffix != nil {
		stage.StageBranch(figure.Suffix)
	}
	if figure.Extend != nil {
		stage.StageBranch(figure.Extend)
	}
	if figure.Footnote != nil {
		stage.StageBranch(figure.Footnote)
	}
	if figure.Level != nil {
		stage.StageBranch(figure.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (figured_bass *Figured_bass) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(figured_bass) {
		return
	}

	figured_bass.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if figured_bass.Footnote != nil {
		stage.StageBranch(figured_bass.Footnote)
	}
	if figured_bass.Level != nil {
		stage.StageBranch(figured_bass.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _figure := range figured_bass.Figure {
		stage.StageBranch(_figure)
	}

}

func (fingering *Fingering) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(fingering) {
		return
	}

	fingering.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (first_fret *First_fret) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(first_fret) {
		return
	}

	first_fret.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (for_part *For_part) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(for_part) {
		return
	}

	for_part.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if for_part.Part_clef != nil {
		stage.StageBranch(for_part.Part_clef)
	}
	if for_part.Part_transpose != nil {
		stage.StageBranch(for_part.Part_transpose)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formatted_symbol *Formatted_symbol) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(formatted_symbol) {
		return
	}

	formatted_symbol.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formatted_symbol_id *Formatted_symbol_id) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(formatted_symbol_id) {
		return
	}

	formatted_symbol_id.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formatted_text *Formatted_text) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(formatted_text) {
		return
	}

	formatted_text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formatted_text_id *Formatted_text_id) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(formatted_text_id) {
		return
	}

	formatted_text_id.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (forward *Forward) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(forward) {
		return
	}

	forward.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if forward.Footnote != nil {
		stage.StageBranch(forward.Footnote)
	}
	if forward.Level != nil {
		stage.StageBranch(forward.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (frame *Frame) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(frame) {
		return
	}

	frame.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if frame.First_fret != nil {
		stage.StageBranch(frame.First_fret)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _frame_note := range frame.Frame_note {
		stage.StageBranch(_frame_note)
	}

}

func (frame_note *Frame_note) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(frame_note) {
		return
	}

	frame_note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if frame_note.String != nil {
		stage.StageBranch(frame_note.String)
	}
	if frame_note.Fret != nil {
		stage.StageBranch(frame_note.Fret)
	}
	if frame_note.Fingering != nil {
		stage.StageBranch(frame_note.Fingering)
	}
	if frame_note.Barre != nil {
		stage.StageBranch(frame_note.Barre)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (fret *Fret) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(fret) {
		return
	}

	fret.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (glass *Glass) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(glass) {
		return
	}

	glass.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (glissando *Glissando) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(glissando) {
		return
	}

	glissando.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (glyph *Glyph) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(glyph) {
		return
	}

	glyph.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (grace *Grace) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(grace) {
		return
	}

	grace.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (group_barline *Group_barline) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(group_barline) {
		return
	}

	group_barline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (group_name *Group_name) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(group_name) {
		return
	}

	group_name.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (group_symbol *Group_symbol) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(group_symbol) {
		return
	}

	group_symbol.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (grouping *Grouping) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(grouping) {
		return
	}

	grouping.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _feature := range grouping.Feature {
		stage.StageBranch(_feature)
	}

}

func (hammer_on_pull_off *Hammer_on_pull_off) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(hammer_on_pull_off) {
		return
	}

	hammer_on_pull_off.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (handbell *Handbell) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(handbell) {
		return
	}

	handbell.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmon_closed *Harmon_closed) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(harmon_closed) {
		return
	}

	harmon_closed.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmon_mute *Harmon_mute) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(harmon_mute) {
		return
	}

	harmon_mute.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmon_mute.Harmon_closed != nil {
		stage.StageBranch(harmon_mute.Harmon_closed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmonic *Harmonic) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(harmonic) {
		return
	}

	harmonic.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmony *Harmony) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(harmony) {
		return
	}

	harmony.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmony.Root != nil {
		stage.StageBranch(harmony.Root)
	}
	if harmony.Numeral != nil {
		stage.StageBranch(harmony.Numeral)
	}
	if harmony.Function != nil {
		stage.StageBranch(harmony.Function)
	}
	if harmony.Kind != nil {
		stage.StageBranch(harmony.Kind)
	}
	if harmony.Inversion != nil {
		stage.StageBranch(harmony.Inversion)
	}
	if harmony.Bass != nil {
		stage.StageBranch(harmony.Bass)
	}
	if harmony.Frame != nil {
		stage.StageBranch(harmony.Frame)
	}
	if harmony.Offset != nil {
		stage.StageBranch(harmony.Offset)
	}
	if harmony.Footnote != nil {
		stage.StageBranch(harmony.Footnote)
	}
	if harmony.Level != nil {
		stage.StageBranch(harmony.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _degree := range harmony.Degree {
		stage.StageBranch(_degree)
	}

}

func (harmony_alter *Harmony_alter) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(harmony_alter) {
		return
	}

	harmony_alter.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harp_pedals *Harp_pedals) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(harp_pedals) {
		return
	}

	harp_pedals.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _pedal_tuning := range harp_pedals.Pedal_tuning {
		stage.StageBranch(_pedal_tuning)
	}

}

func (heel_toe *Heel_toe) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(heel_toe) {
		return
	}

	heel_toe.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (hole *Hole) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(hole) {
		return
	}

	hole.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if hole.Hole_closed != nil {
		stage.StageBranch(hole.Hole_closed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (hole_closed *Hole_closed) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(hole_closed) {
		return
	}

	hole_closed.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (horizontal_turn *Horizontal_turn) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(horizontal_turn) {
		return
	}

	horizontal_turn.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (identification *Identification) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(identification) {
		return
	}

	identification.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if identification.Encoding != nil {
		stage.StageBranch(identification.Encoding)
	}
	if identification.Miscellaneous != nil {
		stage.StageBranch(identification.Miscellaneous)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typed_text := range identification.Creator {
		stage.StageBranch(_typed_text)
	}
	for _, _typed_text := range identification.Rights {
		stage.StageBranch(_typed_text)
	}
	for _, _typed_text := range identification.Relation {
		stage.StageBranch(_typed_text)
	}

}

func (image *Image) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(image) {
		return
	}

	image.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (instrument *Instrument) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(instrument) {
		return
	}

	instrument.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (instrument_change *Instrument_change) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(instrument_change) {
		return
	}

	instrument_change.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if instrument_change.Virtual_instrument != nil {
		stage.StageBranch(instrument_change.Virtual_instrument)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (instrument_link *Instrument_link) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(instrument_link) {
		return
	}

	instrument_link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (interchangeable *Interchangeable) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(interchangeable) {
		return
	}

	interchangeable.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (inversion *Inversion) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(inversion) {
		return
	}

	inversion.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (key *Key) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(key) {
		return
	}

	key.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if key.Cancel != nil {
		stage.StageBranch(key.Cancel)
	}
	if key.Key_accidental != nil {
		stage.StageBranch(key.Key_accidental)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key_octave := range key.Key_octave {
		stage.StageBranch(_key_octave)
	}

}

func (key_accidental *Key_accidental) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(key_accidental) {
		return
	}

	key_accidental.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (key_octave *Key_octave) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(key_octave) {
		return
	}

	key_octave.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (kind *Kind) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(kind) {
		return
	}

	kind.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (level *Level) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(level) {
		return
	}

	level.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (line_detail *Line_detail) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(line_detail) {
		return
	}

	line_detail.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (line_width *Line_width) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(line_width) {
		return
	}

	line_width.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (link *Link) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(link) {
		return
	}

	link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (listen *Listen) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(listen) {
		return
	}

	listen.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assess := range listen.Assess {
		stage.StageBranch(_assess)
	}
	for _, _wait := range listen.Wait {
		stage.StageBranch(_wait)
	}
	for _, _other_listening := range listen.Other_listen {
		stage.StageBranch(_other_listening)
	}

}

func (listening *Listening) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(listening) {
		return
	}

	listening.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if listening.Offset != nil {
		stage.StageBranch(listening.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sync := range listening.Sync {
		stage.StageBranch(_sync)
	}
	for _, _other_listening := range listening.Other_listening {
		stage.StageBranch(_other_listening)
	}

}

func (lyric *Lyric) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(lyric) {
		return
	}

	lyric.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lyric.Extend != nil {
		stage.StageBranch(lyric.Extend)
	}
	if lyric.Footnote != nil {
		stage.StageBranch(lyric.Footnote)
	}
	if lyric.Level != nil {
		stage.StageBranch(lyric.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _elision := range lyric.Elision {
		stage.StageBranch(_elision)
	}
	for _, _text_element_data := range lyric.Text {
		stage.StageBranch(_text_element_data)
	}

}

func (lyric_font *Lyric_font) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(lyric_font) {
		return
	}

	lyric_font.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lyric_language *Lyric_language) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(lyric_language) {
		return
	}

	lyric_language.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measure_layout *Measure_layout) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measure_layout) {
		return
	}

	measure_layout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measure_numbering *Measure_numbering) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measure_numbering) {
		return
	}

	measure_numbering.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measure_repeat *Measure_repeat) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measure_repeat) {
		return
	}

	measure_repeat.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measure_style *Measure_style) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measure_style) {
		return
	}

	measure_style.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if measure_style.Multiple_rest != nil {
		stage.StageBranch(measure_style.Multiple_rest)
	}
	if measure_style.Measure_repeat != nil {
		stage.StageBranch(measure_style.Measure_repeat)
	}
	if measure_style.Beat_repeat != nil {
		stage.StageBranch(measure_style.Beat_repeat)
	}
	if measure_style.Slash != nil {
		stage.StageBranch(measure_style.Slash)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (membrane *Membrane) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(membrane) {
		return
	}

	membrane.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metal *Metal) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(metal) {
		return
	}

	metal.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metronome *Metronome) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(metronome) {
		return
	}

	metronome.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if metronome.Per_minute != nil {
		stage.StageBranch(metronome.Per_minute)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _beat_unit_tied := range metronome.Beat_unit_tied {
		stage.StageBranch(_beat_unit_tied)
	}
	for _, _metronome_note := range metronome.Metronome_note {
		stage.StageBranch(_metronome_note)
	}

}

func (metronome_beam *Metronome_beam) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(metronome_beam) {
		return
	}

	metronome_beam.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metronome_note *Metronome_note) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(metronome_note) {
		return
	}

	metronome_note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if metronome_note.Metronome_tied != nil {
		stage.StageBranch(metronome_note.Metronome_tied)
	}
	if metronome_note.Metronome_tuplet != nil {
		stage.StageBranch(metronome_note.Metronome_tuplet)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _metronome_beam := range metronome_note.Metronome_beam {
		stage.StageBranch(_metronome_beam)
	}

}

func (metronome_tied *Metronome_tied) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(metronome_tied) {
		return
	}

	metronome_tied.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metronome_tuplet *Metronome_tuplet) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(metronome_tuplet) {
		return
	}

	metronome_tuplet.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (midi_device *Midi_device) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(midi_device) {
		return
	}

	midi_device.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (midi_instrument *Midi_instrument) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(midi_instrument) {
		return
	}

	midi_instrument.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (miscellaneous *Miscellaneous) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(miscellaneous) {
		return
	}

	miscellaneous.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _miscellaneous_field := range miscellaneous.Miscellaneous_field {
		stage.StageBranch(_miscellaneous_field)
	}

}

func (miscellaneous_field *Miscellaneous_field) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(miscellaneous_field) {
		return
	}

	miscellaneous_field.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (mordent *Mordent) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(mordent) {
		return
	}

	mordent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (multiple_rest *Multiple_rest) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(multiple_rest) {
		return
	}

	multiple_rest.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (name_display *Name_display) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(name_display) {
		return
	}

	name_display.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formatted_text := range name_display.Display_text {
		stage.StageBranch(_formatted_text)
	}
	for _, _accidental_text := range name_display.Accidental_text {
		stage.StageBranch(_accidental_text)
	}

}

func (non_arpeggiate *Non_arpeggiate) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(non_arpeggiate) {
		return
	}

	non_arpeggiate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notations *Notations) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(notations) {
		return
	}

	notations.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notations.Footnote != nil {
		stage.StageBranch(notations.Footnote)
	}
	if notations.Level != nil {
		stage.StageBranch(notations.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tied := range notations.Tied {
		stage.StageBranch(_tied)
	}
	for _, _slur := range notations.Slur {
		stage.StageBranch(_slur)
	}
	for _, _tuplet := range notations.Tuplet {
		stage.StageBranch(_tuplet)
	}
	for _, _glissando := range notations.Glissando {
		stage.StageBranch(_glissando)
	}
	for _, _slide := range notations.Slide {
		stage.StageBranch(_slide)
	}
	for _, _ornaments := range notations.Ornaments {
		stage.StageBranch(_ornaments)
	}
	for _, _technical := range notations.Technical {
		stage.StageBranch(_technical)
	}
	for _, _articulations := range notations.Articulations {
		stage.StageBranch(_articulations)
	}
	for _, _dynamics := range notations.Dynamics {
		stage.StageBranch(_dynamics)
	}
	for _, _fermata := range notations.Fermata {
		stage.StageBranch(_fermata)
	}
	for _, _arpeggiate := range notations.Arpeggiate {
		stage.StageBranch(_arpeggiate)
	}
	for _, _non_arpeggiate := range notations.Non_arpeggiate {
		stage.StageBranch(_non_arpeggiate)
	}
	for _, _accidental_mark := range notations.Accidental_mark {
		stage.StageBranch(_accidental_mark)
	}
	for _, _other_notation := range notations.Other_notation {
		stage.StageBranch(_other_notation)
	}

}

func (note *Note) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if note.Grace != nil {
		stage.StageBranch(note.Grace)
	}
	if note.Pitch != nil {
		stage.StageBranch(note.Pitch)
	}
	if note.Unpitched != nil {
		stage.StageBranch(note.Unpitched)
	}
	if note.Rest != nil {
		stage.StageBranch(note.Rest)
	}
	if note.Tie != nil {
		stage.StageBranch(note.Tie)
	}
	if note.Footnote != nil {
		stage.StageBranch(note.Footnote)
	}
	if note.Level != nil {
		stage.StageBranch(note.Level)
	}
	if note.Type != nil {
		stage.StageBranch(note.Type)
	}
	if note.Accidental != nil {
		stage.StageBranch(note.Accidental)
	}
	if note.Time_modification != nil {
		stage.StageBranch(note.Time_modification)
	}
	if note.Stem != nil {
		stage.StageBranch(note.Stem)
	}
	if note.Notehead != nil {
		stage.StageBranch(note.Notehead)
	}
	if note.Notehead_text != nil {
		stage.StageBranch(note.Notehead_text)
	}
	if note.Beam != nil {
		stage.StageBranch(note.Beam)
	}
	if note.Play != nil {
		stage.StageBranch(note.Play)
	}
	if note.Listen != nil {
		stage.StageBranch(note.Listen)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument := range note.Instrument {
		stage.StageBranch(_instrument)
	}
	for _, _empty_placement := range note.Dot {
		stage.StageBranch(_empty_placement)
	}
	for _, _notations := range note.Notations {
		stage.StageBranch(_notations)
	}
	for _, _lyric := range note.Lyric {
		stage.StageBranch(_lyric)
	}

}

func (note_size *Note_size) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(note_size) {
		return
	}

	note_size.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (note_type *Note_type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(note_type) {
		return
	}

	note_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notehead *Notehead) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(notehead) {
		return
	}

	notehead.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notehead_text *Notehead_text) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(notehead_text) {
		return
	}

	notehead_text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formatted_text := range notehead_text.Display_text {
		stage.StageBranch(_formatted_text)
	}
	for _, _accidental_text := range notehead_text.Accidental_text {
		stage.StageBranch(_accidental_text)
	}

}

func (numeral *Numeral) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(numeral) {
		return
	}

	numeral.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if numeral.Numeral_root != nil {
		stage.StageBranch(numeral.Numeral_root)
	}
	if numeral.Numeral_alter != nil {
		stage.StageBranch(numeral.Numeral_alter)
	}
	if numeral.Numeral_key != nil {
		stage.StageBranch(numeral.Numeral_key)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (numeral_key *Numeral_key) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(numeral_key) {
		return
	}

	numeral_key.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (numeral_root *Numeral_root) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(numeral_root) {
		return
	}

	numeral_root.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (octave_shift *Octave_shift) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(octave_shift) {
		return
	}

	octave_shift.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (offset *Offset) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(offset) {
		return
	}

	offset.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (opus *Opus) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(opus) {
		return
	}

	opus.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ornaments *Ornaments) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ornaments) {
		return
	}

	ornaments.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty_trill_sound := range ornaments.Trill_mark {
		stage.StageBranch(_empty_trill_sound)
	}
	for _, _horizontal_turn := range ornaments.Turn {
		stage.StageBranch(_horizontal_turn)
	}
	for _, _horizontal_turn := range ornaments.Delayed_turn {
		stage.StageBranch(_horizontal_turn)
	}
	for _, _horizontal_turn := range ornaments.Inverted_turn {
		stage.StageBranch(_horizontal_turn)
	}
	for _, _horizontal_turn := range ornaments.Delayed_inverted_turn {
		stage.StageBranch(_horizontal_turn)
	}
	for _, _empty_trill_sound := range ornaments.Vertical_turn {
		stage.StageBranch(_empty_trill_sound)
	}
	for _, _empty_trill_sound := range ornaments.Inverted_vertical_turn {
		stage.StageBranch(_empty_trill_sound)
	}
	for _, _empty_trill_sound := range ornaments.Shake {
		stage.StageBranch(_empty_trill_sound)
	}
	for _, _wavy_line := range ornaments.Wavy_line {
		stage.StageBranch(_wavy_line)
	}
	for _, _mordent := range ornaments.Mordent {
		stage.StageBranch(_mordent)
	}
	for _, _mordent := range ornaments.Inverted_mordent {
		stage.StageBranch(_mordent)
	}
	for _, _empty_placement := range ornaments.Schleifer {
		stage.StageBranch(_empty_placement)
	}
	for _, _tremolo := range ornaments.Tremolo {
		stage.StageBranch(_tremolo)
	}
	for _, _empty_trill_sound := range ornaments.Haydn {
		stage.StageBranch(_empty_trill_sound)
	}
	for _, _other_placement_text := range ornaments.Other_ornament {
		stage.StageBranch(_other_placement_text)
	}
	for _, _accidental_mark := range ornaments.Accidental_mark {
		stage.StageBranch(_accidental_mark)
	}

}

func (other_appearance *Other_appearance) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(other_appearance) {
		return
	}

	other_appearance.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (other_direction *Other_direction) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(other_direction) {
		return
	}

	other_direction.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (other_listening *Other_listening) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(other_listening) {
		return
	}

	other_listening.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (other_notation *Other_notation) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(other_notation) {
		return
	}

	other_notation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (other_placement_text *Other_placement_text) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(other_placement_text) {
		return
	}

	other_placement_text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (other_play *Other_play) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(other_play) {
		return
	}

	other_play.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (other_text *Other_text) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(other_text) {
		return
	}

	other_text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (page_layout *Page_layout) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(page_layout) {
		return
	}

	page_layout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if page_layout.Page_margins != nil {
		stage.StageBranch(page_layout.Page_margins)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (page_margins *Page_margins) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(page_margins) {
		return
	}

	page_margins.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part_clef *Part_clef) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(part_clef) {
		return
	}

	part_clef.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part_group *Part_group) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(part_group) {
		return
	}

	part_group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if part_group.Group_name != nil {
		stage.StageBranch(part_group.Group_name)
	}
	if part_group.Group_name_display != nil {
		stage.StageBranch(part_group.Group_name_display)
	}
	if part_group.Group_abbreviation != nil {
		stage.StageBranch(part_group.Group_abbreviation)
	}
	if part_group.Group_abbreviation_display != nil {
		stage.StageBranch(part_group.Group_abbreviation_display)
	}
	if part_group.Group_symbol != nil {
		stage.StageBranch(part_group.Group_symbol)
	}
	if part_group.Group_barline != nil {
		stage.StageBranch(part_group.Group_barline)
	}
	if part_group.Footnote != nil {
		stage.StageBranch(part_group.Footnote)
	}
	if part_group.Level != nil {
		stage.StageBranch(part_group.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part_link *Part_link) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(part_link) {
		return
	}

	part_link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument_link := range part_link.Instrument_link {
		stage.StageBranch(_instrument_link)
	}

}

func (part_list *Part_list) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(part_list) {
		return
	}

	part_list.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if part_list.Part_group != nil {
		stage.StageBranch(part_list.Part_group)
	}
	if part_list.Score_part != nil {
		stage.StageBranch(part_list.Score_part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part_name *Part_name) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(part_name) {
		return
	}

	part_name.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part_symbol *Part_symbol) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(part_symbol) {
		return
	}

	part_symbol.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part_transpose *Part_transpose) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(part_transpose) {
		return
	}

	part_transpose.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pedal *Pedal) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pedal) {
		return
	}

	pedal.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pedal_tuning *Pedal_tuning) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pedal_tuning) {
		return
	}

	pedal_tuning.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (per_minute *Per_minute) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(per_minute) {
		return
	}

	per_minute.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (percussion *Percussion) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(percussion) {
		return
	}

	percussion.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if percussion.Glass != nil {
		stage.StageBranch(percussion.Glass)
	}
	if percussion.Metal != nil {
		stage.StageBranch(percussion.Metal)
	}
	if percussion.Wood != nil {
		stage.StageBranch(percussion.Wood)
	}
	if percussion.Pitched != nil {
		stage.StageBranch(percussion.Pitched)
	}
	if percussion.Membrane != nil {
		stage.StageBranch(percussion.Membrane)
	}
	if percussion.Effect != nil {
		stage.StageBranch(percussion.Effect)
	}
	if percussion.Timpani != nil {
		stage.StageBranch(percussion.Timpani)
	}
	if percussion.Beater != nil {
		stage.StageBranch(percussion.Beater)
	}
	if percussion.Stick != nil {
		stage.StageBranch(percussion.Stick)
	}
	if percussion.Other_percussion != nil {
		stage.StageBranch(percussion.Other_percussion)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pitch *Pitch) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pitch) {
		return
	}

	pitch.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pitched *Pitched) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pitched) {
		return
	}

	pitched.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (placement_text *Placement_text) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(placement_text) {
		return
	}

	placement_text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (play *Play) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(play) {
		return
	}

	play.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _other_play := range play.Other_play {
		stage.StageBranch(_other_play)
	}

}

func (player *Player) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(player) {
		return
	}

	player.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (principal_voice *Principal_voice) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(principal_voice) {
		return
	}

	principal_voice.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (print *Print) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(print) {
		return
	}

	print.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if print.Page_layout != nil {
		stage.StageBranch(print.Page_layout)
	}
	if print.System_layout != nil {
		stage.StageBranch(print.System_layout)
	}
	if print.Measure_layout != nil {
		stage.StageBranch(print.Measure_layout)
	}
	if print.Measure_numbering != nil {
		stage.StageBranch(print.Measure_numbering)
	}
	if print.Part_name_display != nil {
		stage.StageBranch(print.Part_name_display)
	}
	if print.Part_abbreviation_display != nil {
		stage.StageBranch(print.Part_abbreviation_display)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _staff_layout := range print.Staff_layout {
		stage.StageBranch(_staff_layout)
	}

}

func (release *Release) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(release) {
		return
	}

	release.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (repeat *Repeat) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(repeat) {
		return
	}

	repeat.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rest *Rest) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rest) {
		return
	}

	rest.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (root *Root) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(root) {
		return
	}

	root.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if root.Root_step != nil {
		stage.StageBranch(root.Root_step)
	}
	if root.Root_alter != nil {
		stage.StageBranch(root.Root_alter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (root_step *Root_step) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(root_step) {
		return
	}

	root_step.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (scaling *Scaling) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(scaling) {
		return
	}

	scaling.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (scordatura *Scordatura) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(scordatura) {
		return
	}

	scordatura.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _accord := range scordatura.Accord {
		stage.StageBranch(_accord)
	}

}

func (score_instrument *Score_instrument) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(score_instrument) {
		return
	}

	score_instrument.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if score_instrument.Virtual_instrument != nil {
		stage.StageBranch(score_instrument.Virtual_instrument)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (score_part *Score_part) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(score_part) {
		return
	}

	score_part.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if score_part.Identification != nil {
		stage.StageBranch(score_part.Identification)
	}
	if score_part.Part_name != nil {
		stage.StageBranch(score_part.Part_name)
	}
	if score_part.Part_name_display != nil {
		stage.StageBranch(score_part.Part_name_display)
	}
	if score_part.Part_abbreviation != nil {
		stage.StageBranch(score_part.Part_abbreviation)
	}
	if score_part.Part_abbreviation_display != nil {
		stage.StageBranch(score_part.Part_abbreviation_display)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part_link := range score_part.Part_link {
		stage.StageBranch(_part_link)
	}
	for _, _score_instrument := range score_part.Score_instrument {
		stage.StageBranch(_score_instrument)
	}
	for _, _player := range score_part.Player {
		stage.StageBranch(_player)
	}
	for _, _midi_device := range score_part.Midi_device {
		stage.StageBranch(_midi_device)
	}
	for _, _midi_instrument := range score_part.Midi_instrument {
		stage.StageBranch(_midi_instrument)
	}

}

func (score_partwise *Score_partwise) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(score_partwise) {
		return
	}

	score_partwise.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if score_partwise.Work != nil {
		stage.StageBranch(score_partwise.Work)
	}
	if score_partwise.Identification != nil {
		stage.StageBranch(score_partwise.Identification)
	}
	if score_partwise.Defaults != nil {
		stage.StageBranch(score_partwise.Defaults)
	}
	if score_partwise.Part_list != nil {
		stage.StageBranch(score_partwise.Part_list)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range score_partwise.Credit {
		stage.StageBranch(_credit)
	}
	for _, _a_part := range score_partwise.Part {
		stage.StageBranch(_a_part)
	}

}

func (score_timewise *Score_timewise) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(score_timewise) {
		return
	}

	score_timewise.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if score_timewise.Work != nil {
		stage.StageBranch(score_timewise.Work)
	}
	if score_timewise.Identification != nil {
		stage.StageBranch(score_timewise.Identification)
	}
	if score_timewise.Defaults != nil {
		stage.StageBranch(score_timewise.Defaults)
	}
	if score_timewise.Part_list != nil {
		stage.StageBranch(score_timewise.Part_list)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range score_timewise.Credit {
		stage.StageBranch(_credit)
	}
	for _, _a_measure_1 := range score_timewise.Measure {
		stage.StageBranch(_a_measure_1)
	}

}

func (segno *Segno) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(segno) {
		return
	}

	segno.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (slash *Slash) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(slash) {
		return
	}

	slash.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (slide *Slide) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(slide) {
		return
	}

	slide.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (slur *Slur) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(slur) {
		return
	}

	slur.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (sound *Sound) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(sound) {
		return
	}

	sound.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sound.Swing != nil {
		stage.StageBranch(sound.Swing)
	}
	if sound.Offset != nil {
		stage.StageBranch(sound.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument_change := range sound.Instrument_change {
		stage.StageBranch(_instrument_change)
	}
	for _, _midi_device := range sound.Midi_device {
		stage.StageBranch(_midi_device)
	}
	for _, _midi_instrument := range sound.Midi_instrument {
		stage.StageBranch(_midi_instrument)
	}
	for _, _play := range sound.Play {
		stage.StageBranch(_play)
	}

}

func (staff_details *Staff_details) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(staff_details) {
		return
	}

	staff_details.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if staff_details.Staff_size != nil {
		stage.StageBranch(staff_details.Staff_size)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _line_detail := range staff_details.Line_detail {
		stage.StageBranch(_line_detail)
	}
	for _, _staff_tuning := range staff_details.Staff_tuning {
		stage.StageBranch(_staff_tuning)
	}

}

func (staff_divide *Staff_divide) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(staff_divide) {
		return
	}

	staff_divide.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (staff_layout *Staff_layout) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(staff_layout) {
		return
	}

	staff_layout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (staff_size *Staff_size) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(staff_size) {
		return
	}

	staff_size.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (staff_tuning *Staff_tuning) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(staff_tuning) {
		return
	}

	staff_tuning.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stem *Stem) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stem) {
		return
	}

	stem.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stick *Stick) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(stick) {
		return
	}

	stick.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (string_mute *String_mute) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(string_mute) {
		return
	}

	string_mute.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (string_type *String_type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(string_type) {
		return
	}

	string_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (strong_accent *Strong_accent) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(strong_accent) {
		return
	}

	strong_accent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (style_text *Style_text) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(style_text) {
		return
	}

	style_text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (supports *Supports) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(supports) {
		return
	}

	supports.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (swing *Swing) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(swing) {
		return
	}

	swing.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (sync *Sync) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(sync) {
		return
	}

	sync.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (system_dividers *System_dividers) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(system_dividers) {
		return
	}

	system_dividers.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if system_dividers.Left_divider != nil {
		stage.StageBranch(system_dividers.Left_divider)
	}
	if system_dividers.Right_divider != nil {
		stage.StageBranch(system_dividers.Right_divider)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (system_layout *System_layout) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(system_layout) {
		return
	}

	system_layout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if system_layout.System_margins != nil {
		stage.StageBranch(system_layout.System_margins)
	}
	if system_layout.System_dividers != nil {
		stage.StageBranch(system_layout.System_dividers)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (system_margins *System_margins) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(system_margins) {
		return
	}

	system_margins.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tap *Tap) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tap) {
		return
	}

	tap.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (technical *Technical) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(technical) {
		return
	}

	technical.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty_placement := range technical.Up_bow {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_placement := range technical.Down_bow {
		stage.StageBranch(_empty_placement)
	}
	for _, _harmonic := range technical.Harmonic {
		stage.StageBranch(_harmonic)
	}
	for _, _empty_placement := range technical.Open_string {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_placement := range technical.Thumb_position {
		stage.StageBranch(_empty_placement)
	}
	for _, _fingering := range technical.Fingering {
		stage.StageBranch(_fingering)
	}
	for _, _placement_text := range technical.Pluck {
		stage.StageBranch(_placement_text)
	}
	for _, _empty_placement := range technical.Double_tongue {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_placement := range technical.Triple_tongue {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_placement_smufl := range technical.Stopped {
		stage.StageBranch(_empty_placement_smufl)
	}
	for _, _empty_placement := range technical.Snap_pizzicato {
		stage.StageBranch(_empty_placement)
	}
	for _, _fret := range technical.Fret {
		stage.StageBranch(_fret)
	}
	for _, _string_type := range technical.String {
		stage.StageBranch(_string_type)
	}
	for _, _hammer_on_pull_off := range technical.Hammer_on {
		stage.StageBranch(_hammer_on_pull_off)
	}
	for _, _hammer_on_pull_off := range technical.Pull_off {
		stage.StageBranch(_hammer_on_pull_off)
	}
	for _, _bend := range technical.Bend {
		stage.StageBranch(_bend)
	}
	for _, _tap := range technical.Tap {
		stage.StageBranch(_tap)
	}
	for _, _heel_toe := range technical.Heel {
		stage.StageBranch(_heel_toe)
	}
	for _, _heel_toe := range technical.Toe {
		stage.StageBranch(_heel_toe)
	}
	for _, _empty_placement := range technical.Fingernails {
		stage.StageBranch(_empty_placement)
	}
	for _, _hole := range technical.Hole {
		stage.StageBranch(_hole)
	}
	for _, _arrow := range technical.Arrow {
		stage.StageBranch(_arrow)
	}
	for _, _handbell := range technical.Handbell {
		stage.StageBranch(_handbell)
	}
	for _, _empty_placement := range technical.Brass_bend {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_placement := range technical.Flip {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_placement := range technical.Smear {
		stage.StageBranch(_empty_placement)
	}
	for _, _empty_placement_smufl := range technical.Open {
		stage.StageBranch(_empty_placement_smufl)
	}
	for _, _empty_placement_smufl := range technical.Half_muted {
		stage.StageBranch(_empty_placement_smufl)
	}
	for _, _harmon_mute := range technical.Harmon_mute {
		stage.StageBranch(_harmon_mute)
	}
	for _, _empty_placement := range technical.Golpe {
		stage.StageBranch(_empty_placement)
	}
	for _, _other_placement_text := range technical.Other_technical {
		stage.StageBranch(_other_placement_text)
	}

}

func (text_element_data *Text_element_data) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(text_element_data) {
		return
	}

	text_element_data.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tie *Tie) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tie) {
		return
	}

	tie.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tied *Tied) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tied) {
		return
	}

	tied.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (time *Time) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(time) {
		return
	}

	time.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if time.Interchangeable != nil {
		stage.StageBranch(time.Interchangeable)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (time_modification *Time_modification) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(time_modification) {
		return
	}

	time_modification.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timpani *Timpani) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timpani) {
		return
	}

	timpani.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (transpose *Transpose) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(transpose) {
		return
	}

	transpose.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tremolo *Tremolo) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tremolo) {
		return
	}

	tremolo.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tuplet *Tuplet) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tuplet) {
		return
	}

	tuplet.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tuplet.Tuplet_actual != nil {
		stage.StageBranch(tuplet.Tuplet_actual)
	}
	if tuplet.Tuplet_normal != nil {
		stage.StageBranch(tuplet.Tuplet_normal)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tuplet_dot *Tuplet_dot) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tuplet_dot) {
		return
	}

	tuplet_dot.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tuplet_number *Tuplet_number) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tuplet_number) {
		return
	}

	tuplet_number.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tuplet_portion *Tuplet_portion) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tuplet_portion) {
		return
	}

	tuplet_portion.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tuplet_portion.Tuplet_number != nil {
		stage.StageBranch(tuplet_portion.Tuplet_number)
	}
	if tuplet_portion.Tuplet_type != nil {
		stage.StageBranch(tuplet_portion.Tuplet_type)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tuplet_dot := range tuplet_portion.Tuplet_dot {
		stage.StageBranch(_tuplet_dot)
	}

}

func (tuplet_type *Tuplet_type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tuplet_type) {
		return
	}

	tuplet_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (typed_text *Typed_text) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(typed_text) {
		return
	}

	typed_text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (unpitched *Unpitched) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(unpitched) {
		return
	}

	unpitched.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (virtual_instrument *Virtual_instrument) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(virtual_instrument) {
		return
	}

	virtual_instrument.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wait *Wait) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(wait) {
		return
	}

	wait.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wavy_line *Wavy_line) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(wavy_line) {
		return
	}

	wavy_line.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wedge *Wedge) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(wedge) {
		return
	}

	wedge.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wood *Wood) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(wood) {
		return
	}

	wood.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (work *Work) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(work) {
		return
	}

	work.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if work.Opus != nil {
		stage.StageBranch(work.Opus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *A_directive:
		toT := GongCopyBranchA_directive(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_measure:
		toT := GongCopyBranchA_measure(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_measure_1:
		toT := GongCopyBranchA_measure_1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_part:
		toT := GongCopyBranchA_part(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_part_1:
		toT := GongCopyBranchA_part_1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Accidental:
		toT := GongCopyBranchAccidental(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Accidental_mark:
		toT := GongCopyBranchAccidental_mark(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Accidental_text:
		toT := GongCopyBranchAccidental_text(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Accord:
		toT := GongCopyBranchAccord(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Accordion_registration:
		toT := GongCopyBranchAccordion_registration(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Appearance:
		toT := GongCopyBranchAppearance(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Arpeggiate:
		toT := GongCopyBranchArpeggiate(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Arrow:
		toT := GongCopyBranchArrow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Articulations:
		toT := GongCopyBranchArticulations(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Assess:
		toT := GongCopyBranchAssess(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Attributes:
		toT := GongCopyBranchAttributes(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Backup:
		toT := GongCopyBranchBackup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bar_style_color:
		toT := GongCopyBranchBar_style_color(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Barline:
		toT := GongCopyBranchBarline(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Barre:
		toT := GongCopyBranchBarre(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bass:
		toT := GongCopyBranchBass(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bass_step:
		toT := GongCopyBranchBass_step(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Beam:
		toT := GongCopyBranchBeam(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Beat_repeat:
		toT := GongCopyBranchBeat_repeat(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Beat_unit_tied:
		toT := GongCopyBranchBeat_unit_tied(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Beater:
		toT := GongCopyBranchBeater(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bend:
		toT := GongCopyBranchBend(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bookmark:
		toT := GongCopyBranchBookmark(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bracket:
		toT := GongCopyBranchBracket(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Breath_mark:
		toT := GongCopyBranchBreath_mark(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Caesura:
		toT := GongCopyBranchCaesura(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Cancel:
		toT := GongCopyBranchCancel(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Clef:
		toT := GongCopyBranchClef(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Coda:
		toT := GongCopyBranchCoda(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Credit:
		toT := GongCopyBranchCredit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Dashes:
		toT := GongCopyBranchDashes(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Defaults:
		toT := GongCopyBranchDefaults(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Degree:
		toT := GongCopyBranchDegree(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Degree_alter:
		toT := GongCopyBranchDegree_alter(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Degree_type:
		toT := GongCopyBranchDegree_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Degree_value:
		toT := GongCopyBranchDegree_value(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Direction:
		toT := GongCopyBranchDirection(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Direction_type:
		toT := GongCopyBranchDirection_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Distance:
		toT := GongCopyBranchDistance(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Double:
		toT := GongCopyBranchDouble(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Dynamics:
		toT := GongCopyBranchDynamics(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Effect:
		toT := GongCopyBranchEffect(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Elision:
		toT := GongCopyBranchElision(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty:
		toT := GongCopyBranchEmpty(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_font:
		toT := GongCopyBranchEmpty_font(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_line:
		toT := GongCopyBranchEmpty_line(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_placement:
		toT := GongCopyBranchEmpty_placement(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_placement_smufl:
		toT := GongCopyBranchEmpty_placement_smufl(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_print_object_style_align:
		toT := GongCopyBranchEmpty_print_object_style_align(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_print_style:
		toT := GongCopyBranchEmpty_print_style(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_print_style_align:
		toT := GongCopyBranchEmpty_print_style_align(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_print_style_align_id:
		toT := GongCopyBranchEmpty_print_style_align_id(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_trill_sound:
		toT := GongCopyBranchEmpty_trill_sound(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Encoding:
		toT := GongCopyBranchEncoding(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Ending:
		toT := GongCopyBranchEnding(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Extend:
		toT := GongCopyBranchExtend(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Feature:
		toT := GongCopyBranchFeature(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Fermata:
		toT := GongCopyBranchFermata(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Figure:
		toT := GongCopyBranchFigure(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Figured_bass:
		toT := GongCopyBranchFigured_bass(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Fingering:
		toT := GongCopyBranchFingering(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *First_fret:
		toT := GongCopyBranchFirst_fret(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *For_part:
		toT := GongCopyBranchFor_part(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Formatted_symbol:
		toT := GongCopyBranchFormatted_symbol(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Formatted_symbol_id:
		toT := GongCopyBranchFormatted_symbol_id(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Formatted_text:
		toT := GongCopyBranchFormatted_text(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Formatted_text_id:
		toT := GongCopyBranchFormatted_text_id(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Forward:
		toT := GongCopyBranchForward(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Frame:
		toT := GongCopyBranchFrame(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Frame_note:
		toT := GongCopyBranchFrame_note(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Fret:
		toT := GongCopyBranchFret(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Glass:
		toT := GongCopyBranchGlass(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Glissando:
		toT := GongCopyBranchGlissando(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Glyph:
		toT := GongCopyBranchGlyph(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Grace:
		toT := GongCopyBranchGrace(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group_barline:
		toT := GongCopyBranchGroup_barline(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group_name:
		toT := GongCopyBranchGroup_name(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group_symbol:
		toT := GongCopyBranchGroup_symbol(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Grouping:
		toT := GongCopyBranchGrouping(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Hammer_on_pull_off:
		toT := GongCopyBranchHammer_on_pull_off(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Handbell:
		toT := GongCopyBranchHandbell(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harmon_closed:
		toT := GongCopyBranchHarmon_closed(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harmon_mute:
		toT := GongCopyBranchHarmon_mute(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harmonic:
		toT := GongCopyBranchHarmonic(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harmony:
		toT := GongCopyBranchHarmony(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harmony_alter:
		toT := GongCopyBranchHarmony_alter(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harp_pedals:
		toT := GongCopyBranchHarp_pedals(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Heel_toe:
		toT := GongCopyBranchHeel_toe(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Hole:
		toT := GongCopyBranchHole(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Hole_closed:
		toT := GongCopyBranchHole_closed(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Horizontal_turn:
		toT := GongCopyBranchHorizontal_turn(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Identification:
		toT := GongCopyBranchIdentification(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Image:
		toT := GongCopyBranchImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Instrument:
		toT := GongCopyBranchInstrument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Instrument_change:
		toT := GongCopyBranchInstrument_change(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Instrument_link:
		toT := GongCopyBranchInstrument_link(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Interchangeable:
		toT := GongCopyBranchInterchangeable(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Inversion:
		toT := GongCopyBranchInversion(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Key:
		toT := GongCopyBranchKey(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Key_accidental:
		toT := GongCopyBranchKey_accidental(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Key_octave:
		toT := GongCopyBranchKey_octave(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Kind:
		toT := GongCopyBranchKind(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Level:
		toT := GongCopyBranchLevel(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Line_detail:
		toT := GongCopyBranchLine_detail(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Line_width:
		toT := GongCopyBranchLine_width(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Link:
		toT := GongCopyBranchLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Listen:
		toT := GongCopyBranchListen(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Listening:
		toT := GongCopyBranchListening(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Lyric:
		toT := GongCopyBranchLyric(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Lyric_font:
		toT := GongCopyBranchLyric_font(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Lyric_language:
		toT := GongCopyBranchLyric_language(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Measure_layout:
		toT := GongCopyBranchMeasure_layout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Measure_numbering:
		toT := GongCopyBranchMeasure_numbering(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Measure_repeat:
		toT := GongCopyBranchMeasure_repeat(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Measure_style:
		toT := GongCopyBranchMeasure_style(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Membrane:
		toT := GongCopyBranchMembrane(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metal:
		toT := GongCopyBranchMetal(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metronome:
		toT := GongCopyBranchMetronome(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metronome_beam:
		toT := GongCopyBranchMetronome_beam(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metronome_note:
		toT := GongCopyBranchMetronome_note(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metronome_tied:
		toT := GongCopyBranchMetronome_tied(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metronome_tuplet:
		toT := GongCopyBranchMetronome_tuplet(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Midi_device:
		toT := GongCopyBranchMidi_device(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Midi_instrument:
		toT := GongCopyBranchMidi_instrument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Miscellaneous:
		toT := GongCopyBranchMiscellaneous(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Miscellaneous_field:
		toT := GongCopyBranchMiscellaneous_field(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Mordent:
		toT := GongCopyBranchMordent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Multiple_rest:
		toT := GongCopyBranchMultiple_rest(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Name_display:
		toT := GongCopyBranchName_display(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Non_arpeggiate:
		toT := GongCopyBranchNon_arpeggiate(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Notations:
		toT := GongCopyBranchNotations(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := GongCopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note_size:
		toT := GongCopyBranchNote_size(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note_type:
		toT := GongCopyBranchNote_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Notehead:
		toT := GongCopyBranchNotehead(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Notehead_text:
		toT := GongCopyBranchNotehead_text(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Numeral:
		toT := GongCopyBranchNumeral(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Numeral_key:
		toT := GongCopyBranchNumeral_key(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Numeral_root:
		toT := GongCopyBranchNumeral_root(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Octave_shift:
		toT := GongCopyBranchOctave_shift(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Offset:
		toT := GongCopyBranchOffset(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Opus:
		toT := GongCopyBranchOpus(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Ornaments:
		toT := GongCopyBranchOrnaments(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Other_appearance:
		toT := GongCopyBranchOther_appearance(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Other_direction:
		toT := GongCopyBranchOther_direction(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Other_listening:
		toT := GongCopyBranchOther_listening(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Other_notation:
		toT := GongCopyBranchOther_notation(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Other_placement_text:
		toT := GongCopyBranchOther_placement_text(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Other_play:
		toT := GongCopyBranchOther_play(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Other_text:
		toT := GongCopyBranchOther_text(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Page_layout:
		toT := GongCopyBranchPage_layout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Page_margins:
		toT := GongCopyBranchPage_margins(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part_clef:
		toT := GongCopyBranchPart_clef(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part_group:
		toT := GongCopyBranchPart_group(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part_link:
		toT := GongCopyBranchPart_link(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part_list:
		toT := GongCopyBranchPart_list(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part_name:
		toT := GongCopyBranchPart_name(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part_symbol:
		toT := GongCopyBranchPart_symbol(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part_transpose:
		toT := GongCopyBranchPart_transpose(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pedal:
		toT := GongCopyBranchPedal(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pedal_tuning:
		toT := GongCopyBranchPedal_tuning(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Per_minute:
		toT := GongCopyBranchPer_minute(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Percussion:
		toT := GongCopyBranchPercussion(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pitch:
		toT := GongCopyBranchPitch(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pitched:
		toT := GongCopyBranchPitched(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Placement_text:
		toT := GongCopyBranchPlacement_text(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Play:
		toT := GongCopyBranchPlay(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Player:
		toT := GongCopyBranchPlayer(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Principal_voice:
		toT := GongCopyBranchPrincipal_voice(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Print:
		toT := GongCopyBranchPrint(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Release:
		toT := GongCopyBranchRelease(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Repeat:
		toT := GongCopyBranchRepeat(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Rest:
		toT := GongCopyBranchRest(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Root:
		toT := GongCopyBranchRoot(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Root_step:
		toT := GongCopyBranchRoot_step(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Scaling:
		toT := GongCopyBranchScaling(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Scordatura:
		toT := GongCopyBranchScordatura(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Score_instrument:
		toT := GongCopyBranchScore_instrument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Score_part:
		toT := GongCopyBranchScore_part(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Score_partwise:
		toT := GongCopyBranchScore_partwise(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Score_timewise:
		toT := GongCopyBranchScore_timewise(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Segno:
		toT := GongCopyBranchSegno(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Slash:
		toT := GongCopyBranchSlash(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Slide:
		toT := GongCopyBranchSlide(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Slur:
		toT := GongCopyBranchSlur(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Sound:
		toT := GongCopyBranchSound(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Staff_details:
		toT := GongCopyBranchStaff_details(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Staff_divide:
		toT := GongCopyBranchStaff_divide(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Staff_layout:
		toT := GongCopyBranchStaff_layout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Staff_size:
		toT := GongCopyBranchStaff_size(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Staff_tuning:
		toT := GongCopyBranchStaff_tuning(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Stem:
		toT := GongCopyBranchStem(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Stick:
		toT := GongCopyBranchStick(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *String_mute:
		toT := GongCopyBranchString_mute(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *String_type:
		toT := GongCopyBranchString_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Strong_accent:
		toT := GongCopyBranchStrong_accent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Style_text:
		toT := GongCopyBranchStyle_text(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Supports:
		toT := GongCopyBranchSupports(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Swing:
		toT := GongCopyBranchSwing(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Sync:
		toT := GongCopyBranchSync(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *System_dividers:
		toT := GongCopyBranchSystem_dividers(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *System_layout:
		toT := GongCopyBranchSystem_layout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *System_margins:
		toT := GongCopyBranchSystem_margins(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tap:
		toT := GongCopyBranchTap(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Technical:
		toT := GongCopyBranchTechnical(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Text_element_data:
		toT := GongCopyBranchText_element_data(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tie:
		toT := GongCopyBranchTie(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tied:
		toT := GongCopyBranchTied(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Time:
		toT := GongCopyBranchTime(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Time_modification:
		toT := GongCopyBranchTime_modification(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Timpani:
		toT := GongCopyBranchTimpani(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Transpose:
		toT := GongCopyBranchTranspose(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tremolo:
		toT := GongCopyBranchTremolo(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tuplet:
		toT := GongCopyBranchTuplet(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tuplet_dot:
		toT := GongCopyBranchTuplet_dot(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tuplet_number:
		toT := GongCopyBranchTuplet_number(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tuplet_portion:
		toT := GongCopyBranchTuplet_portion(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tuplet_type:
		toT := GongCopyBranchTuplet_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Typed_text:
		toT := GongCopyBranchTyped_text(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Unpitched:
		toT := GongCopyBranchUnpitched(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Virtual_instrument:
		toT := GongCopyBranchVirtual_instrument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Wait:
		toT := GongCopyBranchWait(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Wavy_line:
		toT := GongCopyBranchWavy_line(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Wedge:
		toT := GongCopyBranchWedge(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Wood:
		toT := GongCopyBranchWood(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Work:
		toT := GongCopyBranchWork(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchA_directive(mapOrigCopy map[any]any, a_directiveFrom *A_directive) (a_directiveTo *A_directive) {
	var alreadyCopied bool
	a_directiveTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_directiveFrom)
	if alreadyCopied {
		return
	}
	a_directiveFrom.GongCopyBasicFields(a_directiveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_measure(mapOrigCopy map[any]any, a_measureFrom *A_measure) (a_measureTo *A_measure) {
	var alreadyCopied bool
	a_measureTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_measureFrom)
	if alreadyCopied {
		return
	}
	a_measureFrom.GongCopyBasicFields(a_measureTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _note := range a_measureFrom.Note {
		a_measureTo.Note = append(a_measureTo.Note, GongCopyBranchNote(mapOrigCopy, _note))
	}
	for _, _backup := range a_measureFrom.Backup {
		a_measureTo.Backup = append(a_measureTo.Backup, GongCopyBranchBackup(mapOrigCopy, _backup))
	}
	for _, _forward := range a_measureFrom.Forward {
		a_measureTo.Forward = append(a_measureTo.Forward, GongCopyBranchForward(mapOrigCopy, _forward))
	}
	for _, _direction := range a_measureFrom.Direction {
		a_measureTo.Direction = append(a_measureTo.Direction, GongCopyBranchDirection(mapOrigCopy, _direction))
	}
	for _, _attributes := range a_measureFrom.Attributes {
		a_measureTo.Attributes = append(a_measureTo.Attributes, GongCopyBranchAttributes(mapOrigCopy, _attributes))
	}
	for _, _harmony := range a_measureFrom.Harmony {
		a_measureTo.Harmony = append(a_measureTo.Harmony, GongCopyBranchHarmony(mapOrigCopy, _harmony))
	}
	for _, _figured_bass := range a_measureFrom.Figured_bass {
		a_measureTo.Figured_bass = append(a_measureTo.Figured_bass, GongCopyBranchFigured_bass(mapOrigCopy, _figured_bass))
	}
	for _, _print := range a_measureFrom.Print {
		a_measureTo.Print = append(a_measureTo.Print, GongCopyBranchPrint(mapOrigCopy, _print))
	}
	for _, _sound := range a_measureFrom.Sound {
		a_measureTo.Sound = append(a_measureTo.Sound, GongCopyBranchSound(mapOrigCopy, _sound))
	}
	for _, _listening := range a_measureFrom.Listening {
		a_measureTo.Listening = append(a_measureTo.Listening, GongCopyBranchListening(mapOrigCopy, _listening))
	}
	for _, _barline := range a_measureFrom.Barline {
		a_measureTo.Barline = append(a_measureTo.Barline, GongCopyBranchBarline(mapOrigCopy, _barline))
	}
	for _, _grouping := range a_measureFrom.Grouping {
		a_measureTo.Grouping = append(a_measureTo.Grouping, GongCopyBranchGrouping(mapOrigCopy, _grouping))
	}
	for _, _link := range a_measureFrom.Link {
		a_measureTo.Link = append(a_measureTo.Link, GongCopyBranchLink(mapOrigCopy, _link))
	}
	for _, _bookmark := range a_measureFrom.Bookmark {
		a_measureTo.Bookmark = append(a_measureTo.Bookmark, GongCopyBranchBookmark(mapOrigCopy, _bookmark))
	}

	return
}

func GongCopyBranchA_measure_1(mapOrigCopy map[any]any, a_measure_1From *A_measure_1) (a_measure_1To *A_measure_1) {
	var alreadyCopied bool
	a_measure_1To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_measure_1From)
	if alreadyCopied {
		return
	}
	a_measure_1From.GongCopyBasicFields(a_measure_1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_part_1 := range a_measure_1From.Part {
		a_measure_1To.Part = append(a_measure_1To.Part, GongCopyBranchA_part_1(mapOrigCopy, _a_part_1))
	}

	return
}

func GongCopyBranchA_part(mapOrigCopy map[any]any, a_partFrom *A_part) (a_partTo *A_part) {
	var alreadyCopied bool
	a_partTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_partFrom)
	if alreadyCopied {
		return
	}
	a_partFrom.GongCopyBasicFields(a_partTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_measure := range a_partFrom.Measure {
		a_partTo.Measure = append(a_partTo.Measure, GongCopyBranchA_measure(mapOrigCopy, _a_measure))
	}

	return
}

func GongCopyBranchA_part_1(mapOrigCopy map[any]any, a_part_1From *A_part_1) (a_part_1To *A_part_1) {
	var alreadyCopied bool
	a_part_1To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, a_part_1From)
	if alreadyCopied {
		return
	}
	a_part_1From.GongCopyBasicFields(a_part_1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _note := range a_part_1From.Note {
		a_part_1To.Note = append(a_part_1To.Note, GongCopyBranchNote(mapOrigCopy, _note))
	}
	for _, _backup := range a_part_1From.Backup {
		a_part_1To.Backup = append(a_part_1To.Backup, GongCopyBranchBackup(mapOrigCopy, _backup))
	}
	for _, _forward := range a_part_1From.Forward {
		a_part_1To.Forward = append(a_part_1To.Forward, GongCopyBranchForward(mapOrigCopy, _forward))
	}
	for _, _direction := range a_part_1From.Direction {
		a_part_1To.Direction = append(a_part_1To.Direction, GongCopyBranchDirection(mapOrigCopy, _direction))
	}
	for _, _attributes := range a_part_1From.Attributes {
		a_part_1To.Attributes = append(a_part_1To.Attributes, GongCopyBranchAttributes(mapOrigCopy, _attributes))
	}
	for _, _harmony := range a_part_1From.Harmony {
		a_part_1To.Harmony = append(a_part_1To.Harmony, GongCopyBranchHarmony(mapOrigCopy, _harmony))
	}
	for _, _figured_bass := range a_part_1From.Figured_bass {
		a_part_1To.Figured_bass = append(a_part_1To.Figured_bass, GongCopyBranchFigured_bass(mapOrigCopy, _figured_bass))
	}
	for _, _print := range a_part_1From.Print {
		a_part_1To.Print = append(a_part_1To.Print, GongCopyBranchPrint(mapOrigCopy, _print))
	}
	for _, _sound := range a_part_1From.Sound {
		a_part_1To.Sound = append(a_part_1To.Sound, GongCopyBranchSound(mapOrigCopy, _sound))
	}
	for _, _listening := range a_part_1From.Listening {
		a_part_1To.Listening = append(a_part_1To.Listening, GongCopyBranchListening(mapOrigCopy, _listening))
	}
	for _, _barline := range a_part_1From.Barline {
		a_part_1To.Barline = append(a_part_1To.Barline, GongCopyBranchBarline(mapOrigCopy, _barline))
	}
	for _, _grouping := range a_part_1From.Grouping {
		a_part_1To.Grouping = append(a_part_1To.Grouping, GongCopyBranchGrouping(mapOrigCopy, _grouping))
	}
	for _, _link := range a_part_1From.Link {
		a_part_1To.Link = append(a_part_1To.Link, GongCopyBranchLink(mapOrigCopy, _link))
	}
	for _, _bookmark := range a_part_1From.Bookmark {
		a_part_1To.Bookmark = append(a_part_1To.Bookmark, GongCopyBranchBookmark(mapOrigCopy, _bookmark))
	}

	return
}

func GongCopyBranchAccidental(mapOrigCopy map[any]any, accidentalFrom *Accidental) (accidentalTo *Accidental) {
	var alreadyCopied bool
	accidentalTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, accidentalFrom)
	if alreadyCopied {
		return
	}
	accidentalFrom.GongCopyBasicFields(accidentalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAccidental_mark(mapOrigCopy map[any]any, accidental_markFrom *Accidental_mark) (accidental_markTo *Accidental_mark) {
	var alreadyCopied bool
	accidental_markTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, accidental_markFrom)
	if alreadyCopied {
		return
	}
	accidental_markFrom.GongCopyBasicFields(accidental_markTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAccidental_text(mapOrigCopy map[any]any, accidental_textFrom *Accidental_text) (accidental_textTo *Accidental_text) {
	var alreadyCopied bool
	accidental_textTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, accidental_textFrom)
	if alreadyCopied {
		return
	}
	accidental_textFrom.GongCopyBasicFields(accidental_textTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAccord(mapOrigCopy map[any]any, accordFrom *Accord) (accordTo *Accord) {
	var alreadyCopied bool
	accordTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, accordFrom)
	if alreadyCopied {
		return
	}
	accordFrom.GongCopyBasicFields(accordTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAccordion_registration(mapOrigCopy map[any]any, accordion_registrationFrom *Accordion_registration) (accordion_registrationTo *Accordion_registration) {
	var alreadyCopied bool
	accordion_registrationTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, accordion_registrationFrom)
	if alreadyCopied {
		return
	}
	accordion_registrationFrom.GongCopyBasicFields(accordion_registrationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAppearance(mapOrigCopy map[any]any, appearanceFrom *Appearance) (appearanceTo *Appearance) {
	var alreadyCopied bool
	appearanceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, appearanceFrom)
	if alreadyCopied {
		return
	}
	appearanceFrom.GongCopyBasicFields(appearanceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _line_width := range appearanceFrom.Line_width {
		appearanceTo.Line_width = append(appearanceTo.Line_width, GongCopyBranchLine_width(mapOrigCopy, _line_width))
	}
	for _, _note_size := range appearanceFrom.Note_size {
		appearanceTo.Note_size = append(appearanceTo.Note_size, GongCopyBranchNote_size(mapOrigCopy, _note_size))
	}
	for _, _distance := range appearanceFrom.Distance {
		appearanceTo.Distance = append(appearanceTo.Distance, GongCopyBranchDistance(mapOrigCopy, _distance))
	}
	for _, _glyph := range appearanceFrom.Glyph {
		appearanceTo.Glyph = append(appearanceTo.Glyph, GongCopyBranchGlyph(mapOrigCopy, _glyph))
	}
	for _, _other_appearance := range appearanceFrom.Other_appearance {
		appearanceTo.Other_appearance = append(appearanceTo.Other_appearance, GongCopyBranchOther_appearance(mapOrigCopy, _other_appearance))
	}

	return
}

func GongCopyBranchArpeggiate(mapOrigCopy map[any]any, arpeggiateFrom *Arpeggiate) (arpeggiateTo *Arpeggiate) {
	var alreadyCopied bool
	arpeggiateTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, arpeggiateFrom)
	if alreadyCopied {
		return
	}
	arpeggiateFrom.GongCopyBasicFields(arpeggiateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArrow(mapOrigCopy map[any]any, arrowFrom *Arrow) (arrowTo *Arrow) {
	var alreadyCopied bool
	arrowTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, arrowFrom)
	if alreadyCopied {
		return
	}
	arrowFrom.GongCopyBasicFields(arrowTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArticulations(mapOrigCopy map[any]any, articulationsFrom *Articulations) (articulationsTo *Articulations) {
	var alreadyCopied bool
	articulationsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, articulationsFrom)
	if alreadyCopied {
		return
	}
	articulationsFrom.GongCopyBasicFields(articulationsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty_placement := range articulationsFrom.Accent {
		articulationsTo.Accent = append(articulationsTo.Accent, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _strong_accent := range articulationsFrom.Strong_accent {
		articulationsTo.Strong_accent = append(articulationsTo.Strong_accent, GongCopyBranchStrong_accent(mapOrigCopy, _strong_accent))
	}
	for _, _empty_placement := range articulationsFrom.Staccato {
		articulationsTo.Staccato = append(articulationsTo.Staccato, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_placement := range articulationsFrom.Tenuto {
		articulationsTo.Tenuto = append(articulationsTo.Tenuto, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_placement := range articulationsFrom.Detached_legato {
		articulationsTo.Detached_legato = append(articulationsTo.Detached_legato, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_placement := range articulationsFrom.Staccatissimo {
		articulationsTo.Staccatissimo = append(articulationsTo.Staccatissimo, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_placement := range articulationsFrom.Spiccato {
		articulationsTo.Spiccato = append(articulationsTo.Spiccato, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_line := range articulationsFrom.Scoop {
		articulationsTo.Scoop = append(articulationsTo.Scoop, GongCopyBranchEmpty_line(mapOrigCopy, _empty_line))
	}
	for _, _empty_line := range articulationsFrom.Plop {
		articulationsTo.Plop = append(articulationsTo.Plop, GongCopyBranchEmpty_line(mapOrigCopy, _empty_line))
	}
	for _, _empty_line := range articulationsFrom.Doit {
		articulationsTo.Doit = append(articulationsTo.Doit, GongCopyBranchEmpty_line(mapOrigCopy, _empty_line))
	}
	for _, _empty_line := range articulationsFrom.Falloff {
		articulationsTo.Falloff = append(articulationsTo.Falloff, GongCopyBranchEmpty_line(mapOrigCopy, _empty_line))
	}
	for _, _breath_mark := range articulationsFrom.Breath_mark {
		articulationsTo.Breath_mark = append(articulationsTo.Breath_mark, GongCopyBranchBreath_mark(mapOrigCopy, _breath_mark))
	}
	for _, _caesura := range articulationsFrom.Caesura {
		articulationsTo.Caesura = append(articulationsTo.Caesura, GongCopyBranchCaesura(mapOrigCopy, _caesura))
	}
	for _, _empty_placement := range articulationsFrom.Stress {
		articulationsTo.Stress = append(articulationsTo.Stress, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_placement := range articulationsFrom.Unstress {
		articulationsTo.Unstress = append(articulationsTo.Unstress, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_placement := range articulationsFrom.Soft_accent {
		articulationsTo.Soft_accent = append(articulationsTo.Soft_accent, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _other_placement_text := range articulationsFrom.Other_articulation {
		articulationsTo.Other_articulation = append(articulationsTo.Other_articulation, GongCopyBranchOther_placement_text(mapOrigCopy, _other_placement_text))
	}

	return
}

func GongCopyBranchAssess(mapOrigCopy map[any]any, assessFrom *Assess) (assessTo *Assess) {
	var alreadyCopied bool
	assessTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, assessFrom)
	if alreadyCopied {
		return
	}
	assessFrom.GongCopyBasicFields(assessTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAttributes(mapOrigCopy map[any]any, attributesFrom *Attributes) (attributesTo *Attributes) {
	var alreadyCopied bool
	attributesTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, attributesFrom)
	if alreadyCopied {
		return
	}
	attributesFrom.GongCopyBasicFields(attributesTo)

	//insertion point for the staging of instances referenced by pointers
	if attributesFrom.Footnote != nil {
		attributesTo.Footnote = GongCopyBranchFormatted_text(mapOrigCopy, attributesFrom.Footnote)
	}
	if attributesFrom.Level != nil {
		attributesTo.Level = GongCopyBranchLevel(mapOrigCopy, attributesFrom.Level)
	}
	if attributesFrom.Part_symbol != nil {
		attributesTo.Part_symbol = GongCopyBranchPart_symbol(mapOrigCopy, attributesFrom.Part_symbol)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key := range attributesFrom.Key {
		attributesTo.Key = append(attributesTo.Key, GongCopyBranchKey(mapOrigCopy, _key))
	}
	for _, _time := range attributesFrom.Time {
		attributesTo.Time = append(attributesTo.Time, GongCopyBranchTime(mapOrigCopy, _time))
	}
	for _, _clef := range attributesFrom.Clef {
		attributesTo.Clef = append(attributesTo.Clef, GongCopyBranchClef(mapOrigCopy, _clef))
	}
	for _, _staff_details := range attributesFrom.Staff_details {
		attributesTo.Staff_details = append(attributesTo.Staff_details, GongCopyBranchStaff_details(mapOrigCopy, _staff_details))
	}
	for _, _transpose := range attributesFrom.Transpose {
		attributesTo.Transpose = append(attributesTo.Transpose, GongCopyBranchTranspose(mapOrigCopy, _transpose))
	}
	for _, _for_part := range attributesFrom.For_part {
		attributesTo.For_part = append(attributesTo.For_part, GongCopyBranchFor_part(mapOrigCopy, _for_part))
	}
	for _, _a_directive := range attributesFrom.Directive {
		attributesTo.Directive = append(attributesTo.Directive, GongCopyBranchA_directive(mapOrigCopy, _a_directive))
	}
	for _, _measure_style := range attributesFrom.Measure_style {
		attributesTo.Measure_style = append(attributesTo.Measure_style, GongCopyBranchMeasure_style(mapOrigCopy, _measure_style))
	}

	return
}

func GongCopyBranchBackup(mapOrigCopy map[any]any, backupFrom *Backup) (backupTo *Backup) {
	var alreadyCopied bool
	backupTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, backupFrom)
	if alreadyCopied {
		return
	}
	backupFrom.GongCopyBasicFields(backupTo)

	//insertion point for the staging of instances referenced by pointers
	if backupFrom.Footnote != nil {
		backupTo.Footnote = GongCopyBranchFormatted_text(mapOrigCopy, backupFrom.Footnote)
	}
	if backupFrom.Level != nil {
		backupTo.Level = GongCopyBranchLevel(mapOrigCopy, backupFrom.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBar_style_color(mapOrigCopy map[any]any, bar_style_colorFrom *Bar_style_color) (bar_style_colorTo *Bar_style_color) {
	var alreadyCopied bool
	bar_style_colorTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bar_style_colorFrom)
	if alreadyCopied {
		return
	}
	bar_style_colorFrom.GongCopyBasicFields(bar_style_colorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBarline(mapOrigCopy map[any]any, barlineFrom *Barline) (barlineTo *Barline) {
	var alreadyCopied bool
	barlineTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, barlineFrom)
	if alreadyCopied {
		return
	}
	barlineFrom.GongCopyBasicFields(barlineTo)

	//insertion point for the staging of instances referenced by pointers
	if barlineFrom.Bar_style != nil {
		barlineTo.Bar_style = GongCopyBranchBar_style_color(mapOrigCopy, barlineFrom.Bar_style)
	}
	if barlineFrom.Footnote != nil {
		barlineTo.Footnote = GongCopyBranchFormatted_text(mapOrigCopy, barlineFrom.Footnote)
	}
	if barlineFrom.Level != nil {
		barlineTo.Level = GongCopyBranchLevel(mapOrigCopy, barlineFrom.Level)
	}
	if barlineFrom.Wavy_line != nil {
		barlineTo.Wavy_line = GongCopyBranchWavy_line(mapOrigCopy, barlineFrom.Wavy_line)
	}
	if barlineFrom.Segno_1 != nil {
		barlineTo.Segno_1 = GongCopyBranchSegno(mapOrigCopy, barlineFrom.Segno_1)
	}
	if barlineFrom.Coda_1 != nil {
		barlineTo.Coda_1 = GongCopyBranchCoda(mapOrigCopy, barlineFrom.Coda_1)
	}
	if barlineFrom.Fermata != nil {
		barlineTo.Fermata = GongCopyBranchFermata(mapOrigCopy, barlineFrom.Fermata)
	}
	if barlineFrom.Ending != nil {
		barlineTo.Ending = GongCopyBranchEnding(mapOrigCopy, barlineFrom.Ending)
	}
	if barlineFrom.Repeat != nil {
		barlineTo.Repeat = GongCopyBranchRepeat(mapOrigCopy, barlineFrom.Repeat)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBarre(mapOrigCopy map[any]any, barreFrom *Barre) (barreTo *Barre) {
	var alreadyCopied bool
	barreTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, barreFrom)
	if alreadyCopied {
		return
	}
	barreFrom.GongCopyBasicFields(barreTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBass(mapOrigCopy map[any]any, bassFrom *Bass) (bassTo *Bass) {
	var alreadyCopied bool
	bassTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bassFrom)
	if alreadyCopied {
		return
	}
	bassFrom.GongCopyBasicFields(bassTo)

	//insertion point for the staging of instances referenced by pointers
	if bassFrom.Bass_separator != nil {
		bassTo.Bass_separator = GongCopyBranchStyle_text(mapOrigCopy, bassFrom.Bass_separator)
	}
	if bassFrom.Bass_step != nil {
		bassTo.Bass_step = GongCopyBranchBass_step(mapOrigCopy, bassFrom.Bass_step)
	}
	if bassFrom.Bass_alter != nil {
		bassTo.Bass_alter = GongCopyBranchHarmony_alter(mapOrigCopy, bassFrom.Bass_alter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBass_step(mapOrigCopy map[any]any, bass_stepFrom *Bass_step) (bass_stepTo *Bass_step) {
	var alreadyCopied bool
	bass_stepTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bass_stepFrom)
	if alreadyCopied {
		return
	}
	bass_stepFrom.GongCopyBasicFields(bass_stepTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBeam(mapOrigCopy map[any]any, beamFrom *Beam) (beamTo *Beam) {
	var alreadyCopied bool
	beamTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, beamFrom)
	if alreadyCopied {
		return
	}
	beamFrom.GongCopyBasicFields(beamTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBeat_repeat(mapOrigCopy map[any]any, beat_repeatFrom *Beat_repeat) (beat_repeatTo *Beat_repeat) {
	var alreadyCopied bool
	beat_repeatTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, beat_repeatFrom)
	if alreadyCopied {
		return
	}
	beat_repeatFrom.GongCopyBasicFields(beat_repeatTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBeat_unit_tied(mapOrigCopy map[any]any, beat_unit_tiedFrom *Beat_unit_tied) (beat_unit_tiedTo *Beat_unit_tied) {
	var alreadyCopied bool
	beat_unit_tiedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, beat_unit_tiedFrom)
	if alreadyCopied {
		return
	}
	beat_unit_tiedFrom.GongCopyBasicFields(beat_unit_tiedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBeater(mapOrigCopy map[any]any, beaterFrom *Beater) (beaterTo *Beater) {
	var alreadyCopied bool
	beaterTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, beaterFrom)
	if alreadyCopied {
		return
	}
	beaterFrom.GongCopyBasicFields(beaterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBend(mapOrigCopy map[any]any, bendFrom *Bend) (bendTo *Bend) {
	var alreadyCopied bool
	bendTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bendFrom)
	if alreadyCopied {
		return
	}
	bendFrom.GongCopyBasicFields(bendTo)

	//insertion point for the staging of instances referenced by pointers
	if bendFrom.Release != nil {
		bendTo.Release = GongCopyBranchRelease(mapOrigCopy, bendFrom.Release)
	}
	if bendFrom.With_bar != nil {
		bendTo.With_bar = GongCopyBranchPlacement_text(mapOrigCopy, bendFrom.With_bar)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBookmark(mapOrigCopy map[any]any, bookmarkFrom *Bookmark) (bookmarkTo *Bookmark) {
	var alreadyCopied bool
	bookmarkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bookmarkFrom)
	if alreadyCopied {
		return
	}
	bookmarkFrom.GongCopyBasicFields(bookmarkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBracket(mapOrigCopy map[any]any, bracketFrom *Bracket) (bracketTo *Bracket) {
	var alreadyCopied bool
	bracketTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bracketFrom)
	if alreadyCopied {
		return
	}
	bracketFrom.GongCopyBasicFields(bracketTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBreath_mark(mapOrigCopy map[any]any, breath_markFrom *Breath_mark) (breath_markTo *Breath_mark) {
	var alreadyCopied bool
	breath_markTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, breath_markFrom)
	if alreadyCopied {
		return
	}
	breath_markFrom.GongCopyBasicFields(breath_markTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCaesura(mapOrigCopy map[any]any, caesuraFrom *Caesura) (caesuraTo *Caesura) {
	var alreadyCopied bool
	caesuraTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, caesuraFrom)
	if alreadyCopied {
		return
	}
	caesuraFrom.GongCopyBasicFields(caesuraTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCancel(mapOrigCopy map[any]any, cancelFrom *Cancel) (cancelTo *Cancel) {
	var alreadyCopied bool
	cancelTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, cancelFrom)
	if alreadyCopied {
		return
	}
	cancelFrom.GongCopyBasicFields(cancelTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchClef(mapOrigCopy map[any]any, clefFrom *Clef) (clefTo *Clef) {
	var alreadyCopied bool
	clefTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, clefFrom)
	if alreadyCopied {
		return
	}
	clefFrom.GongCopyBasicFields(clefTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCoda(mapOrigCopy map[any]any, codaFrom *Coda) (codaTo *Coda) {
	var alreadyCopied bool
	codaTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codaFrom)
	if alreadyCopied {
		return
	}
	codaFrom.GongCopyBasicFields(codaTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCredit(mapOrigCopy map[any]any, creditFrom *Credit) (creditTo *Credit) {
	var alreadyCopied bool
	creditTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, creditFrom)
	if alreadyCopied {
		return
	}
	creditFrom.GongCopyBasicFields(creditTo)

	//insertion point for the staging of instances referenced by pointers
	if creditFrom.Credit_image != nil {
		creditTo.Credit_image = GongCopyBranchImage(mapOrigCopy, creditFrom.Credit_image)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range creditFrom.Link {
		creditTo.Link = append(creditTo.Link, GongCopyBranchLink(mapOrigCopy, _link))
	}
	for _, _bookmark := range creditFrom.Bookmark {
		creditTo.Bookmark = append(creditTo.Bookmark, GongCopyBranchBookmark(mapOrigCopy, _bookmark))
	}
	for _, _formatted_text_id := range creditFrom.Credit_words {
		creditTo.Credit_words = append(creditTo.Credit_words, GongCopyBranchFormatted_text_id(mapOrigCopy, _formatted_text_id))
	}
	for _, _formatted_symbol_id := range creditFrom.Credit_symbol {
		creditTo.Credit_symbol = append(creditTo.Credit_symbol, GongCopyBranchFormatted_symbol_id(mapOrigCopy, _formatted_symbol_id))
	}

	return
}

func GongCopyBranchDashes(mapOrigCopy map[any]any, dashesFrom *Dashes) (dashesTo *Dashes) {
	var alreadyCopied bool
	dashesTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dashesFrom)
	if alreadyCopied {
		return
	}
	dashesFrom.GongCopyBasicFields(dashesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDefaults(mapOrigCopy map[any]any, defaultsFrom *Defaults) (defaultsTo *Defaults) {
	var alreadyCopied bool
	defaultsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, defaultsFrom)
	if alreadyCopied {
		return
	}
	defaultsFrom.GongCopyBasicFields(defaultsTo)

	//insertion point for the staging of instances referenced by pointers
	if defaultsFrom.Scaling != nil {
		defaultsTo.Scaling = GongCopyBranchScaling(mapOrigCopy, defaultsFrom.Scaling)
	}
	if defaultsFrom.Page_layout != nil {
		defaultsTo.Page_layout = GongCopyBranchPage_layout(mapOrigCopy, defaultsFrom.Page_layout)
	}
	if defaultsFrom.System_layout != nil {
		defaultsTo.System_layout = GongCopyBranchSystem_layout(mapOrigCopy, defaultsFrom.System_layout)
	}
	if defaultsFrom.Appearance != nil {
		defaultsTo.Appearance = GongCopyBranchAppearance(mapOrigCopy, defaultsFrom.Appearance)
	}
	if defaultsFrom.Music_font != nil {
		defaultsTo.Music_font = GongCopyBranchEmpty_font(mapOrigCopy, defaultsFrom.Music_font)
	}
	if defaultsFrom.Word_font != nil {
		defaultsTo.Word_font = GongCopyBranchEmpty_font(mapOrigCopy, defaultsFrom.Word_font)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _staff_layout := range defaultsFrom.Staff_layout {
		defaultsTo.Staff_layout = append(defaultsTo.Staff_layout, GongCopyBranchStaff_layout(mapOrigCopy, _staff_layout))
	}
	for _, _lyric_font := range defaultsFrom.Lyric_font {
		defaultsTo.Lyric_font = append(defaultsTo.Lyric_font, GongCopyBranchLyric_font(mapOrigCopy, _lyric_font))
	}
	for _, _lyric_language := range defaultsFrom.Lyric_language {
		defaultsTo.Lyric_language = append(defaultsTo.Lyric_language, GongCopyBranchLyric_language(mapOrigCopy, _lyric_language))
	}

	return
}

func GongCopyBranchDegree(mapOrigCopy map[any]any, degreeFrom *Degree) (degreeTo *Degree) {
	var alreadyCopied bool
	degreeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, degreeFrom)
	if alreadyCopied {
		return
	}
	degreeFrom.GongCopyBasicFields(degreeTo)

	//insertion point for the staging of instances referenced by pointers
	if degreeFrom.Degree_value != nil {
		degreeTo.Degree_value = GongCopyBranchDegree_value(mapOrigCopy, degreeFrom.Degree_value)
	}
	if degreeFrom.Degree_alter != nil {
		degreeTo.Degree_alter = GongCopyBranchDegree_alter(mapOrigCopy, degreeFrom.Degree_alter)
	}
	if degreeFrom.Degree_type != nil {
		degreeTo.Degree_type = GongCopyBranchDegree_type(mapOrigCopy, degreeFrom.Degree_type)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDegree_alter(mapOrigCopy map[any]any, degree_alterFrom *Degree_alter) (degree_alterTo *Degree_alter) {
	var alreadyCopied bool
	degree_alterTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, degree_alterFrom)
	if alreadyCopied {
		return
	}
	degree_alterFrom.GongCopyBasicFields(degree_alterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDegree_type(mapOrigCopy map[any]any, degree_typeFrom *Degree_type) (degree_typeTo *Degree_type) {
	var alreadyCopied bool
	degree_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, degree_typeFrom)
	if alreadyCopied {
		return
	}
	degree_typeFrom.GongCopyBasicFields(degree_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDegree_value(mapOrigCopy map[any]any, degree_valueFrom *Degree_value) (degree_valueTo *Degree_value) {
	var alreadyCopied bool
	degree_valueTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, degree_valueFrom)
	if alreadyCopied {
		return
	}
	degree_valueFrom.GongCopyBasicFields(degree_valueTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDirection(mapOrigCopy map[any]any, directionFrom *Direction) (directionTo *Direction) {
	var alreadyCopied bool
	directionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, directionFrom)
	if alreadyCopied {
		return
	}
	directionFrom.GongCopyBasicFields(directionTo)

	//insertion point for the staging of instances referenced by pointers
	if directionFrom.Offset != nil {
		directionTo.Offset = GongCopyBranchOffset(mapOrigCopy, directionFrom.Offset)
	}
	if directionFrom.Footnote != nil {
		directionTo.Footnote = GongCopyBranchFormatted_text(mapOrigCopy, directionFrom.Footnote)
	}
	if directionFrom.Level != nil {
		directionTo.Level = GongCopyBranchLevel(mapOrigCopy, directionFrom.Level)
	}
	if directionFrom.Sound != nil {
		directionTo.Sound = GongCopyBranchSound(mapOrigCopy, directionFrom.Sound)
	}
	if directionFrom.Listening != nil {
		directionTo.Listening = GongCopyBranchListening(mapOrigCopy, directionFrom.Listening)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _direction_type := range directionFrom.Direction_type {
		directionTo.Direction_type = append(directionTo.Direction_type, GongCopyBranchDirection_type(mapOrigCopy, _direction_type))
	}

	return
}

func GongCopyBranchDirection_type(mapOrigCopy map[any]any, direction_typeFrom *Direction_type) (direction_typeTo *Direction_type) {
	var alreadyCopied bool
	direction_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, direction_typeFrom)
	if alreadyCopied {
		return
	}
	direction_typeFrom.GongCopyBasicFields(direction_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if direction_typeFrom.Wedge != nil {
		direction_typeTo.Wedge = GongCopyBranchWedge(mapOrigCopy, direction_typeFrom.Wedge)
	}
	if direction_typeFrom.Dashes != nil {
		direction_typeTo.Dashes = GongCopyBranchDashes(mapOrigCopy, direction_typeFrom.Dashes)
	}
	if direction_typeFrom.Bracket != nil {
		direction_typeTo.Bracket = GongCopyBranchBracket(mapOrigCopy, direction_typeFrom.Bracket)
	}
	if direction_typeFrom.Pedal != nil {
		direction_typeTo.Pedal = GongCopyBranchPedal(mapOrigCopy, direction_typeFrom.Pedal)
	}
	if direction_typeFrom.Metronome != nil {
		direction_typeTo.Metronome = GongCopyBranchMetronome(mapOrigCopy, direction_typeFrom.Metronome)
	}
	if direction_typeFrom.Octave_shift != nil {
		direction_typeTo.Octave_shift = GongCopyBranchOctave_shift(mapOrigCopy, direction_typeFrom.Octave_shift)
	}
	if direction_typeFrom.Harp_pedals != nil {
		direction_typeTo.Harp_pedals = GongCopyBranchHarp_pedals(mapOrigCopy, direction_typeFrom.Harp_pedals)
	}
	if direction_typeFrom.Damp != nil {
		direction_typeTo.Damp = GongCopyBranchEmpty_print_style_align_id(mapOrigCopy, direction_typeFrom.Damp)
	}
	if direction_typeFrom.Damp_all != nil {
		direction_typeTo.Damp_all = GongCopyBranchEmpty_print_style_align_id(mapOrigCopy, direction_typeFrom.Damp_all)
	}
	if direction_typeFrom.Eyeglasses != nil {
		direction_typeTo.Eyeglasses = GongCopyBranchEmpty_print_style_align_id(mapOrigCopy, direction_typeFrom.Eyeglasses)
	}
	if direction_typeFrom.String_mute != nil {
		direction_typeTo.String_mute = GongCopyBranchString_mute(mapOrigCopy, direction_typeFrom.String_mute)
	}
	if direction_typeFrom.Scordatura != nil {
		direction_typeTo.Scordatura = GongCopyBranchScordatura(mapOrigCopy, direction_typeFrom.Scordatura)
	}
	if direction_typeFrom.Image != nil {
		direction_typeTo.Image = GongCopyBranchImage(mapOrigCopy, direction_typeFrom.Image)
	}
	if direction_typeFrom.Principal_voice != nil {
		direction_typeTo.Principal_voice = GongCopyBranchPrincipal_voice(mapOrigCopy, direction_typeFrom.Principal_voice)
	}
	if direction_typeFrom.Accordion_registration != nil {
		direction_typeTo.Accordion_registration = GongCopyBranchAccordion_registration(mapOrigCopy, direction_typeFrom.Accordion_registration)
	}
	if direction_typeFrom.Staff_divide != nil {
		direction_typeTo.Staff_divide = GongCopyBranchStaff_divide(mapOrigCopy, direction_typeFrom.Staff_divide)
	}
	if direction_typeFrom.Other_direction != nil {
		direction_typeTo.Other_direction = GongCopyBranchOther_direction(mapOrigCopy, direction_typeFrom.Other_direction)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formatted_text_id := range direction_typeFrom.Rehearsal {
		direction_typeTo.Rehearsal = append(direction_typeTo.Rehearsal, GongCopyBranchFormatted_text_id(mapOrigCopy, _formatted_text_id))
	}
	for _, _segno := range direction_typeFrom.Segno {
		direction_typeTo.Segno = append(direction_typeTo.Segno, GongCopyBranchSegno(mapOrigCopy, _segno))
	}
	for _, _coda := range direction_typeFrom.Coda {
		direction_typeTo.Coda = append(direction_typeTo.Coda, GongCopyBranchCoda(mapOrigCopy, _coda))
	}
	for _, _formatted_text_id := range direction_typeFrom.Words {
		direction_typeTo.Words = append(direction_typeTo.Words, GongCopyBranchFormatted_text_id(mapOrigCopy, _formatted_text_id))
	}
	for _, _formatted_symbol_id := range direction_typeFrom.Symbol {
		direction_typeTo.Symbol = append(direction_typeTo.Symbol, GongCopyBranchFormatted_symbol_id(mapOrigCopy, _formatted_symbol_id))
	}
	for _, _dynamics := range direction_typeFrom.Dynamics {
		direction_typeTo.Dynamics = append(direction_typeTo.Dynamics, GongCopyBranchDynamics(mapOrigCopy, _dynamics))
	}
	for _, _percussion := range direction_typeFrom.Percussion {
		direction_typeTo.Percussion = append(direction_typeTo.Percussion, GongCopyBranchPercussion(mapOrigCopy, _percussion))
	}

	return
}

func GongCopyBranchDistance(mapOrigCopy map[any]any, distanceFrom *Distance) (distanceTo *Distance) {
	var alreadyCopied bool
	distanceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, distanceFrom)
	if alreadyCopied {
		return
	}
	distanceFrom.GongCopyBasicFields(distanceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDouble(mapOrigCopy map[any]any, doubleFrom *Double) (doubleTo *Double) {
	var alreadyCopied bool
	doubleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, doubleFrom)
	if alreadyCopied {
		return
	}
	doubleFrom.GongCopyBasicFields(doubleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDynamics(mapOrigCopy map[any]any, dynamicsFrom *Dynamics) (dynamicsTo *Dynamics) {
	var alreadyCopied bool
	dynamicsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dynamicsFrom)
	if alreadyCopied {
		return
	}
	dynamicsFrom.GongCopyBasicFields(dynamicsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _other_text := range dynamicsFrom.Other_dynamics {
		dynamicsTo.Other_dynamics = append(dynamicsTo.Other_dynamics, GongCopyBranchOther_text(mapOrigCopy, _other_text))
	}

	return
}

func GongCopyBranchEffect(mapOrigCopy map[any]any, effectFrom *Effect) (effectTo *Effect) {
	var alreadyCopied bool
	effectTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, effectFrom)
	if alreadyCopied {
		return
	}
	effectFrom.GongCopyBasicFields(effectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchElision(mapOrigCopy map[any]any, elisionFrom *Elision) (elisionTo *Elision) {
	var alreadyCopied bool
	elisionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, elisionFrom)
	if alreadyCopied {
		return
	}
	elisionFrom.GongCopyBasicFields(elisionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmpty(mapOrigCopy map[any]any, emptyFrom *Empty) (emptyTo *Empty) {
	var alreadyCopied bool
	emptyTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, emptyFrom)
	if alreadyCopied {
		return
	}
	emptyFrom.GongCopyBasicFields(emptyTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmpty_font(mapOrigCopy map[any]any, empty_fontFrom *Empty_font) (empty_fontTo *Empty_font) {
	var alreadyCopied bool
	empty_fontTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, empty_fontFrom)
	if alreadyCopied {
		return
	}
	empty_fontFrom.GongCopyBasicFields(empty_fontTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmpty_line(mapOrigCopy map[any]any, empty_lineFrom *Empty_line) (empty_lineTo *Empty_line) {
	var alreadyCopied bool
	empty_lineTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, empty_lineFrom)
	if alreadyCopied {
		return
	}
	empty_lineFrom.GongCopyBasicFields(empty_lineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmpty_placement(mapOrigCopy map[any]any, empty_placementFrom *Empty_placement) (empty_placementTo *Empty_placement) {
	var alreadyCopied bool
	empty_placementTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, empty_placementFrom)
	if alreadyCopied {
		return
	}
	empty_placementFrom.GongCopyBasicFields(empty_placementTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmpty_placement_smufl(mapOrigCopy map[any]any, empty_placement_smuflFrom *Empty_placement_smufl) (empty_placement_smuflTo *Empty_placement_smufl) {
	var alreadyCopied bool
	empty_placement_smuflTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, empty_placement_smuflFrom)
	if alreadyCopied {
		return
	}
	empty_placement_smuflFrom.GongCopyBasicFields(empty_placement_smuflTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmpty_print_object_style_align(mapOrigCopy map[any]any, empty_print_object_style_alignFrom *Empty_print_object_style_align) (empty_print_object_style_alignTo *Empty_print_object_style_align) {
	var alreadyCopied bool
	empty_print_object_style_alignTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, empty_print_object_style_alignFrom)
	if alreadyCopied {
		return
	}
	empty_print_object_style_alignFrom.GongCopyBasicFields(empty_print_object_style_alignTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmpty_print_style(mapOrigCopy map[any]any, empty_print_styleFrom *Empty_print_style) (empty_print_styleTo *Empty_print_style) {
	var alreadyCopied bool
	empty_print_styleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, empty_print_styleFrom)
	if alreadyCopied {
		return
	}
	empty_print_styleFrom.GongCopyBasicFields(empty_print_styleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmpty_print_style_align(mapOrigCopy map[any]any, empty_print_style_alignFrom *Empty_print_style_align) (empty_print_style_alignTo *Empty_print_style_align) {
	var alreadyCopied bool
	empty_print_style_alignTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, empty_print_style_alignFrom)
	if alreadyCopied {
		return
	}
	empty_print_style_alignFrom.GongCopyBasicFields(empty_print_style_alignTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmpty_print_style_align_id(mapOrigCopy map[any]any, empty_print_style_align_idFrom *Empty_print_style_align_id) (empty_print_style_align_idTo *Empty_print_style_align_id) {
	var alreadyCopied bool
	empty_print_style_align_idTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, empty_print_style_align_idFrom)
	if alreadyCopied {
		return
	}
	empty_print_style_align_idFrom.GongCopyBasicFields(empty_print_style_align_idTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEmpty_trill_sound(mapOrigCopy map[any]any, empty_trill_soundFrom *Empty_trill_sound) (empty_trill_soundTo *Empty_trill_sound) {
	var alreadyCopied bool
	empty_trill_soundTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, empty_trill_soundFrom)
	if alreadyCopied {
		return
	}
	empty_trill_soundFrom.GongCopyBasicFields(empty_trill_soundTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEncoding(mapOrigCopy map[any]any, encodingFrom *Encoding) (encodingTo *Encoding) {
	var alreadyCopied bool
	encodingTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, encodingFrom)
	if alreadyCopied {
		return
	}
	encodingFrom.GongCopyBasicFields(encodingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typed_text := range encodingFrom.Encoder {
		encodingTo.Encoder = append(encodingTo.Encoder, GongCopyBranchTyped_text(mapOrigCopy, _typed_text))
	}
	for _, _supports := range encodingFrom.Supports {
		encodingTo.Supports = append(encodingTo.Supports, GongCopyBranchSupports(mapOrigCopy, _supports))
	}

	return
}

func GongCopyBranchEnding(mapOrigCopy map[any]any, endingFrom *Ending) (endingTo *Ending) {
	var alreadyCopied bool
	endingTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, endingFrom)
	if alreadyCopied {
		return
	}
	endingFrom.GongCopyBasicFields(endingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchExtend(mapOrigCopy map[any]any, extendFrom *Extend) (extendTo *Extend) {
	var alreadyCopied bool
	extendTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, extendFrom)
	if alreadyCopied {
		return
	}
	extendFrom.GongCopyBasicFields(extendTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFeature(mapOrigCopy map[any]any, featureFrom *Feature) (featureTo *Feature) {
	var alreadyCopied bool
	featureTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, featureFrom)
	if alreadyCopied {
		return
	}
	featureFrom.GongCopyBasicFields(featureTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFermata(mapOrigCopy map[any]any, fermataFrom *Fermata) (fermataTo *Fermata) {
	var alreadyCopied bool
	fermataTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, fermataFrom)
	if alreadyCopied {
		return
	}
	fermataFrom.GongCopyBasicFields(fermataTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFigure(mapOrigCopy map[any]any, figureFrom *Figure) (figureTo *Figure) {
	var alreadyCopied bool
	figureTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, figureFrom)
	if alreadyCopied {
		return
	}
	figureFrom.GongCopyBasicFields(figureTo)

	//insertion point for the staging of instances referenced by pointers
	if figureFrom.Prefix != nil {
		figureTo.Prefix = GongCopyBranchStyle_text(mapOrigCopy, figureFrom.Prefix)
	}
	if figureFrom.Figure_number != nil {
		figureTo.Figure_number = GongCopyBranchStyle_text(mapOrigCopy, figureFrom.Figure_number)
	}
	if figureFrom.Suffix != nil {
		figureTo.Suffix = GongCopyBranchStyle_text(mapOrigCopy, figureFrom.Suffix)
	}
	if figureFrom.Extend != nil {
		figureTo.Extend = GongCopyBranchExtend(mapOrigCopy, figureFrom.Extend)
	}
	if figureFrom.Footnote != nil {
		figureTo.Footnote = GongCopyBranchFormatted_text(mapOrigCopy, figureFrom.Footnote)
	}
	if figureFrom.Level != nil {
		figureTo.Level = GongCopyBranchLevel(mapOrigCopy, figureFrom.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFigured_bass(mapOrigCopy map[any]any, figured_bassFrom *Figured_bass) (figured_bassTo *Figured_bass) {
	var alreadyCopied bool
	figured_bassTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, figured_bassFrom)
	if alreadyCopied {
		return
	}
	figured_bassFrom.GongCopyBasicFields(figured_bassTo)

	//insertion point for the staging of instances referenced by pointers
	if figured_bassFrom.Footnote != nil {
		figured_bassTo.Footnote = GongCopyBranchFormatted_text(mapOrigCopy, figured_bassFrom.Footnote)
	}
	if figured_bassFrom.Level != nil {
		figured_bassTo.Level = GongCopyBranchLevel(mapOrigCopy, figured_bassFrom.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _figure := range figured_bassFrom.Figure {
		figured_bassTo.Figure = append(figured_bassTo.Figure, GongCopyBranchFigure(mapOrigCopy, _figure))
	}

	return
}

func GongCopyBranchFingering(mapOrigCopy map[any]any, fingeringFrom *Fingering) (fingeringTo *Fingering) {
	var alreadyCopied bool
	fingeringTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, fingeringFrom)
	if alreadyCopied {
		return
	}
	fingeringFrom.GongCopyBasicFields(fingeringTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFirst_fret(mapOrigCopy map[any]any, first_fretFrom *First_fret) (first_fretTo *First_fret) {
	var alreadyCopied bool
	first_fretTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, first_fretFrom)
	if alreadyCopied {
		return
	}
	first_fretFrom.GongCopyBasicFields(first_fretTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFor_part(mapOrigCopy map[any]any, for_partFrom *For_part) (for_partTo *For_part) {
	var alreadyCopied bool
	for_partTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, for_partFrom)
	if alreadyCopied {
		return
	}
	for_partFrom.GongCopyBasicFields(for_partTo)

	//insertion point for the staging of instances referenced by pointers
	if for_partFrom.Part_clef != nil {
		for_partTo.Part_clef = GongCopyBranchPart_clef(mapOrigCopy, for_partFrom.Part_clef)
	}
	if for_partFrom.Part_transpose != nil {
		for_partTo.Part_transpose = GongCopyBranchPart_transpose(mapOrigCopy, for_partFrom.Part_transpose)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormatted_symbol(mapOrigCopy map[any]any, formatted_symbolFrom *Formatted_symbol) (formatted_symbolTo *Formatted_symbol) {
	var alreadyCopied bool
	formatted_symbolTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formatted_symbolFrom)
	if alreadyCopied {
		return
	}
	formatted_symbolFrom.GongCopyBasicFields(formatted_symbolTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormatted_symbol_id(mapOrigCopy map[any]any, formatted_symbol_idFrom *Formatted_symbol_id) (formatted_symbol_idTo *Formatted_symbol_id) {
	var alreadyCopied bool
	formatted_symbol_idTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formatted_symbol_idFrom)
	if alreadyCopied {
		return
	}
	formatted_symbol_idFrom.GongCopyBasicFields(formatted_symbol_idTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormatted_text(mapOrigCopy map[any]any, formatted_textFrom *Formatted_text) (formatted_textTo *Formatted_text) {
	var alreadyCopied bool
	formatted_textTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formatted_textFrom)
	if alreadyCopied {
		return
	}
	formatted_textFrom.GongCopyBasicFields(formatted_textTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormatted_text_id(mapOrigCopy map[any]any, formatted_text_idFrom *Formatted_text_id) (formatted_text_idTo *Formatted_text_id) {
	var alreadyCopied bool
	formatted_text_idTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formatted_text_idFrom)
	if alreadyCopied {
		return
	}
	formatted_text_idFrom.GongCopyBasicFields(formatted_text_idTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchForward(mapOrigCopy map[any]any, forwardFrom *Forward) (forwardTo *Forward) {
	var alreadyCopied bool
	forwardTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, forwardFrom)
	if alreadyCopied {
		return
	}
	forwardFrom.GongCopyBasicFields(forwardTo)

	//insertion point for the staging of instances referenced by pointers
	if forwardFrom.Footnote != nil {
		forwardTo.Footnote = GongCopyBranchFormatted_text(mapOrigCopy, forwardFrom.Footnote)
	}
	if forwardFrom.Level != nil {
		forwardTo.Level = GongCopyBranchLevel(mapOrigCopy, forwardFrom.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFrame(mapOrigCopy map[any]any, frameFrom *Frame) (frameTo *Frame) {
	var alreadyCopied bool
	frameTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, frameFrom)
	if alreadyCopied {
		return
	}
	frameFrom.GongCopyBasicFields(frameTo)

	//insertion point for the staging of instances referenced by pointers
	if frameFrom.First_fret != nil {
		frameTo.First_fret = GongCopyBranchFirst_fret(mapOrigCopy, frameFrom.First_fret)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _frame_note := range frameFrom.Frame_note {
		frameTo.Frame_note = append(frameTo.Frame_note, GongCopyBranchFrame_note(mapOrigCopy, _frame_note))
	}

	return
}

func GongCopyBranchFrame_note(mapOrigCopy map[any]any, frame_noteFrom *Frame_note) (frame_noteTo *Frame_note) {
	var alreadyCopied bool
	frame_noteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, frame_noteFrom)
	if alreadyCopied {
		return
	}
	frame_noteFrom.GongCopyBasicFields(frame_noteTo)

	//insertion point for the staging of instances referenced by pointers
	if frame_noteFrom.String != nil {
		frame_noteTo.String = GongCopyBranchString_type(mapOrigCopy, frame_noteFrom.String)
	}
	if frame_noteFrom.Fret != nil {
		frame_noteTo.Fret = GongCopyBranchFret(mapOrigCopy, frame_noteFrom.Fret)
	}
	if frame_noteFrom.Fingering != nil {
		frame_noteTo.Fingering = GongCopyBranchFingering(mapOrigCopy, frame_noteFrom.Fingering)
	}
	if frame_noteFrom.Barre != nil {
		frame_noteTo.Barre = GongCopyBranchBarre(mapOrigCopy, frame_noteFrom.Barre)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFret(mapOrigCopy map[any]any, fretFrom *Fret) (fretTo *Fret) {
	var alreadyCopied bool
	fretTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, fretFrom)
	if alreadyCopied {
		return
	}
	fretFrom.GongCopyBasicFields(fretTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGlass(mapOrigCopy map[any]any, glassFrom *Glass) (glassTo *Glass) {
	var alreadyCopied bool
	glassTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, glassFrom)
	if alreadyCopied {
		return
	}
	glassFrom.GongCopyBasicFields(glassTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGlissando(mapOrigCopy map[any]any, glissandoFrom *Glissando) (glissandoTo *Glissando) {
	var alreadyCopied bool
	glissandoTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, glissandoFrom)
	if alreadyCopied {
		return
	}
	glissandoFrom.GongCopyBasicFields(glissandoTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGlyph(mapOrigCopy map[any]any, glyphFrom *Glyph) (glyphTo *Glyph) {
	var alreadyCopied bool
	glyphTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, glyphFrom)
	if alreadyCopied {
		return
	}
	glyphFrom.GongCopyBasicFields(glyphTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrace(mapOrigCopy map[any]any, graceFrom *Grace) (graceTo *Grace) {
	var alreadyCopied bool
	graceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, graceFrom)
	if alreadyCopied {
		return
	}
	graceFrom.GongCopyBasicFields(graceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGroup_barline(mapOrigCopy map[any]any, group_barlineFrom *Group_barline) (group_barlineTo *Group_barline) {
	var alreadyCopied bool
	group_barlineTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, group_barlineFrom)
	if alreadyCopied {
		return
	}
	group_barlineFrom.GongCopyBasicFields(group_barlineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGroup_name(mapOrigCopy map[any]any, group_nameFrom *Group_name) (group_nameTo *Group_name) {
	var alreadyCopied bool
	group_nameTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, group_nameFrom)
	if alreadyCopied {
		return
	}
	group_nameFrom.GongCopyBasicFields(group_nameTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGroup_symbol(mapOrigCopy map[any]any, group_symbolFrom *Group_symbol) (group_symbolTo *Group_symbol) {
	var alreadyCopied bool
	group_symbolTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, group_symbolFrom)
	if alreadyCopied {
		return
	}
	group_symbolFrom.GongCopyBasicFields(group_symbolTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGrouping(mapOrigCopy map[any]any, groupingFrom *Grouping) (groupingTo *Grouping) {
	var alreadyCopied bool
	groupingTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, groupingFrom)
	if alreadyCopied {
		return
	}
	groupingFrom.GongCopyBasicFields(groupingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _feature := range groupingFrom.Feature {
		groupingTo.Feature = append(groupingTo.Feature, GongCopyBranchFeature(mapOrigCopy, _feature))
	}

	return
}

func GongCopyBranchHammer_on_pull_off(mapOrigCopy map[any]any, hammer_on_pull_offFrom *Hammer_on_pull_off) (hammer_on_pull_offTo *Hammer_on_pull_off) {
	var alreadyCopied bool
	hammer_on_pull_offTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, hammer_on_pull_offFrom)
	if alreadyCopied {
		return
	}
	hammer_on_pull_offFrom.GongCopyBasicFields(hammer_on_pull_offTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHandbell(mapOrigCopy map[any]any, handbellFrom *Handbell) (handbellTo *Handbell) {
	var alreadyCopied bool
	handbellTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, handbellFrom)
	if alreadyCopied {
		return
	}
	handbellFrom.GongCopyBasicFields(handbellTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHarmon_closed(mapOrigCopy map[any]any, harmon_closedFrom *Harmon_closed) (harmon_closedTo *Harmon_closed) {
	var alreadyCopied bool
	harmon_closedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, harmon_closedFrom)
	if alreadyCopied {
		return
	}
	harmon_closedFrom.GongCopyBasicFields(harmon_closedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHarmon_mute(mapOrigCopy map[any]any, harmon_muteFrom *Harmon_mute) (harmon_muteTo *Harmon_mute) {
	var alreadyCopied bool
	harmon_muteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, harmon_muteFrom)
	if alreadyCopied {
		return
	}
	harmon_muteFrom.GongCopyBasicFields(harmon_muteTo)

	//insertion point for the staging of instances referenced by pointers
	if harmon_muteFrom.Harmon_closed != nil {
		harmon_muteTo.Harmon_closed = GongCopyBranchHarmon_closed(mapOrigCopy, harmon_muteFrom.Harmon_closed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHarmonic(mapOrigCopy map[any]any, harmonicFrom *Harmonic) (harmonicTo *Harmonic) {
	var alreadyCopied bool
	harmonicTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, harmonicFrom)
	if alreadyCopied {
		return
	}
	harmonicFrom.GongCopyBasicFields(harmonicTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHarmony(mapOrigCopy map[any]any, harmonyFrom *Harmony) (harmonyTo *Harmony) {
	var alreadyCopied bool
	harmonyTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, harmonyFrom)
	if alreadyCopied {
		return
	}
	harmonyFrom.GongCopyBasicFields(harmonyTo)

	//insertion point for the staging of instances referenced by pointers
	if harmonyFrom.Root != nil {
		harmonyTo.Root = GongCopyBranchRoot(mapOrigCopy, harmonyFrom.Root)
	}
	if harmonyFrom.Numeral != nil {
		harmonyTo.Numeral = GongCopyBranchNumeral(mapOrigCopy, harmonyFrom.Numeral)
	}
	if harmonyFrom.Function != nil {
		harmonyTo.Function = GongCopyBranchStyle_text(mapOrigCopy, harmonyFrom.Function)
	}
	if harmonyFrom.Kind != nil {
		harmonyTo.Kind = GongCopyBranchKind(mapOrigCopy, harmonyFrom.Kind)
	}
	if harmonyFrom.Inversion != nil {
		harmonyTo.Inversion = GongCopyBranchInversion(mapOrigCopy, harmonyFrom.Inversion)
	}
	if harmonyFrom.Bass != nil {
		harmonyTo.Bass = GongCopyBranchBass(mapOrigCopy, harmonyFrom.Bass)
	}
	if harmonyFrom.Frame != nil {
		harmonyTo.Frame = GongCopyBranchFrame(mapOrigCopy, harmonyFrom.Frame)
	}
	if harmonyFrom.Offset != nil {
		harmonyTo.Offset = GongCopyBranchOffset(mapOrigCopy, harmonyFrom.Offset)
	}
	if harmonyFrom.Footnote != nil {
		harmonyTo.Footnote = GongCopyBranchFormatted_text(mapOrigCopy, harmonyFrom.Footnote)
	}
	if harmonyFrom.Level != nil {
		harmonyTo.Level = GongCopyBranchLevel(mapOrigCopy, harmonyFrom.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _degree := range harmonyFrom.Degree {
		harmonyTo.Degree = append(harmonyTo.Degree, GongCopyBranchDegree(mapOrigCopy, _degree))
	}

	return
}

func GongCopyBranchHarmony_alter(mapOrigCopy map[any]any, harmony_alterFrom *Harmony_alter) (harmony_alterTo *Harmony_alter) {
	var alreadyCopied bool
	harmony_alterTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, harmony_alterFrom)
	if alreadyCopied {
		return
	}
	harmony_alterFrom.GongCopyBasicFields(harmony_alterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHarp_pedals(mapOrigCopy map[any]any, harp_pedalsFrom *Harp_pedals) (harp_pedalsTo *Harp_pedals) {
	var alreadyCopied bool
	harp_pedalsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, harp_pedalsFrom)
	if alreadyCopied {
		return
	}
	harp_pedalsFrom.GongCopyBasicFields(harp_pedalsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _pedal_tuning := range harp_pedalsFrom.Pedal_tuning {
		harp_pedalsTo.Pedal_tuning = append(harp_pedalsTo.Pedal_tuning, GongCopyBranchPedal_tuning(mapOrigCopy, _pedal_tuning))
	}

	return
}

func GongCopyBranchHeel_toe(mapOrigCopy map[any]any, heel_toeFrom *Heel_toe) (heel_toeTo *Heel_toe) {
	var alreadyCopied bool
	heel_toeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, heel_toeFrom)
	if alreadyCopied {
		return
	}
	heel_toeFrom.GongCopyBasicFields(heel_toeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHole(mapOrigCopy map[any]any, holeFrom *Hole) (holeTo *Hole) {
	var alreadyCopied bool
	holeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, holeFrom)
	if alreadyCopied {
		return
	}
	holeFrom.GongCopyBasicFields(holeTo)

	//insertion point for the staging of instances referenced by pointers
	if holeFrom.Hole_closed != nil {
		holeTo.Hole_closed = GongCopyBranchHole_closed(mapOrigCopy, holeFrom.Hole_closed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHole_closed(mapOrigCopy map[any]any, hole_closedFrom *Hole_closed) (hole_closedTo *Hole_closed) {
	var alreadyCopied bool
	hole_closedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, hole_closedFrom)
	if alreadyCopied {
		return
	}
	hole_closedFrom.GongCopyBasicFields(hole_closedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchHorizontal_turn(mapOrigCopy map[any]any, horizontal_turnFrom *Horizontal_turn) (horizontal_turnTo *Horizontal_turn) {
	var alreadyCopied bool
	horizontal_turnTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, horizontal_turnFrom)
	if alreadyCopied {
		return
	}
	horizontal_turnFrom.GongCopyBasicFields(horizontal_turnTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchIdentification(mapOrigCopy map[any]any, identificationFrom *Identification) (identificationTo *Identification) {
	var alreadyCopied bool
	identificationTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, identificationFrom)
	if alreadyCopied {
		return
	}
	identificationFrom.GongCopyBasicFields(identificationTo)

	//insertion point for the staging of instances referenced by pointers
	if identificationFrom.Encoding != nil {
		identificationTo.Encoding = GongCopyBranchEncoding(mapOrigCopy, identificationFrom.Encoding)
	}
	if identificationFrom.Miscellaneous != nil {
		identificationTo.Miscellaneous = GongCopyBranchMiscellaneous(mapOrigCopy, identificationFrom.Miscellaneous)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typed_text := range identificationFrom.Creator {
		identificationTo.Creator = append(identificationTo.Creator, GongCopyBranchTyped_text(mapOrigCopy, _typed_text))
	}
	for _, _typed_text := range identificationFrom.Rights {
		identificationTo.Rights = append(identificationTo.Rights, GongCopyBranchTyped_text(mapOrigCopy, _typed_text))
	}
	for _, _typed_text := range identificationFrom.Relation {
		identificationTo.Relation = append(identificationTo.Relation, GongCopyBranchTyped_text(mapOrigCopy, _typed_text))
	}

	return
}

func GongCopyBranchImage(mapOrigCopy map[any]any, imageFrom *Image) (imageTo *Image) {
	var alreadyCopied bool
	imageTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, imageFrom)
	if alreadyCopied {
		return
	}
	imageFrom.GongCopyBasicFields(imageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInstrument(mapOrigCopy map[any]any, instrumentFrom *Instrument) (instrumentTo *Instrument) {
	var alreadyCopied bool
	instrumentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, instrumentFrom)
	if alreadyCopied {
		return
	}
	instrumentFrom.GongCopyBasicFields(instrumentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInstrument_change(mapOrigCopy map[any]any, instrument_changeFrom *Instrument_change) (instrument_changeTo *Instrument_change) {
	var alreadyCopied bool
	instrument_changeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, instrument_changeFrom)
	if alreadyCopied {
		return
	}
	instrument_changeFrom.GongCopyBasicFields(instrument_changeTo)

	//insertion point for the staging of instances referenced by pointers
	if instrument_changeFrom.Virtual_instrument != nil {
		instrument_changeTo.Virtual_instrument = GongCopyBranchVirtual_instrument(mapOrigCopy, instrument_changeFrom.Virtual_instrument)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInstrument_link(mapOrigCopy map[any]any, instrument_linkFrom *Instrument_link) (instrument_linkTo *Instrument_link) {
	var alreadyCopied bool
	instrument_linkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, instrument_linkFrom)
	if alreadyCopied {
		return
	}
	instrument_linkFrom.GongCopyBasicFields(instrument_linkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInterchangeable(mapOrigCopy map[any]any, interchangeableFrom *Interchangeable) (interchangeableTo *Interchangeable) {
	var alreadyCopied bool
	interchangeableTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, interchangeableFrom)
	if alreadyCopied {
		return
	}
	interchangeableFrom.GongCopyBasicFields(interchangeableTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInversion(mapOrigCopy map[any]any, inversionFrom *Inversion) (inversionTo *Inversion) {
	var alreadyCopied bool
	inversionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, inversionFrom)
	if alreadyCopied {
		return
	}
	inversionFrom.GongCopyBasicFields(inversionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKey(mapOrigCopy map[any]any, keyFrom *Key) (keyTo *Key) {
	var alreadyCopied bool
	keyTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, keyFrom)
	if alreadyCopied {
		return
	}
	keyFrom.GongCopyBasicFields(keyTo)

	//insertion point for the staging of instances referenced by pointers
	if keyFrom.Cancel != nil {
		keyTo.Cancel = GongCopyBranchCancel(mapOrigCopy, keyFrom.Cancel)
	}
	if keyFrom.Key_accidental != nil {
		keyTo.Key_accidental = GongCopyBranchKey_accidental(mapOrigCopy, keyFrom.Key_accidental)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key_octave := range keyFrom.Key_octave {
		keyTo.Key_octave = append(keyTo.Key_octave, GongCopyBranchKey_octave(mapOrigCopy, _key_octave))
	}

	return
}

func GongCopyBranchKey_accidental(mapOrigCopy map[any]any, key_accidentalFrom *Key_accidental) (key_accidentalTo *Key_accidental) {
	var alreadyCopied bool
	key_accidentalTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, key_accidentalFrom)
	if alreadyCopied {
		return
	}
	key_accidentalFrom.GongCopyBasicFields(key_accidentalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKey_octave(mapOrigCopy map[any]any, key_octaveFrom *Key_octave) (key_octaveTo *Key_octave) {
	var alreadyCopied bool
	key_octaveTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, key_octaveFrom)
	if alreadyCopied {
		return
	}
	key_octaveFrom.GongCopyBasicFields(key_octaveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKind(mapOrigCopy map[any]any, kindFrom *Kind) (kindTo *Kind) {
	var alreadyCopied bool
	kindTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, kindFrom)
	if alreadyCopied {
		return
	}
	kindFrom.GongCopyBasicFields(kindTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLevel(mapOrigCopy map[any]any, levelFrom *Level) (levelTo *Level) {
	var alreadyCopied bool
	levelTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, levelFrom)
	if alreadyCopied {
		return
	}
	levelFrom.GongCopyBasicFields(levelTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLine_detail(mapOrigCopy map[any]any, line_detailFrom *Line_detail) (line_detailTo *Line_detail) {
	var alreadyCopied bool
	line_detailTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, line_detailFrom)
	if alreadyCopied {
		return
	}
	line_detailFrom.GongCopyBasicFields(line_detailTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLine_width(mapOrigCopy map[any]any, line_widthFrom *Line_width) (line_widthTo *Line_width) {
	var alreadyCopied bool
	line_widthTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, line_widthFrom)
	if alreadyCopied {
		return
	}
	line_widthFrom.GongCopyBasicFields(line_widthTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLink(mapOrigCopy map[any]any, linkFrom *Link) (linkTo *Link) {
	var alreadyCopied bool
	linkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, linkFrom)
	if alreadyCopied {
		return
	}
	linkFrom.GongCopyBasicFields(linkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchListen(mapOrigCopy map[any]any, listenFrom *Listen) (listenTo *Listen) {
	var alreadyCopied bool
	listenTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, listenFrom)
	if alreadyCopied {
		return
	}
	listenFrom.GongCopyBasicFields(listenTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assess := range listenFrom.Assess {
		listenTo.Assess = append(listenTo.Assess, GongCopyBranchAssess(mapOrigCopy, _assess))
	}
	for _, _wait := range listenFrom.Wait {
		listenTo.Wait = append(listenTo.Wait, GongCopyBranchWait(mapOrigCopy, _wait))
	}
	for _, _other_listening := range listenFrom.Other_listen {
		listenTo.Other_listen = append(listenTo.Other_listen, GongCopyBranchOther_listening(mapOrigCopy, _other_listening))
	}

	return
}

func GongCopyBranchListening(mapOrigCopy map[any]any, listeningFrom *Listening) (listeningTo *Listening) {
	var alreadyCopied bool
	listeningTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, listeningFrom)
	if alreadyCopied {
		return
	}
	listeningFrom.GongCopyBasicFields(listeningTo)

	//insertion point for the staging of instances referenced by pointers
	if listeningFrom.Offset != nil {
		listeningTo.Offset = GongCopyBranchOffset(mapOrigCopy, listeningFrom.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sync := range listeningFrom.Sync {
		listeningTo.Sync = append(listeningTo.Sync, GongCopyBranchSync(mapOrigCopy, _sync))
	}
	for _, _other_listening := range listeningFrom.Other_listening {
		listeningTo.Other_listening = append(listeningTo.Other_listening, GongCopyBranchOther_listening(mapOrigCopy, _other_listening))
	}

	return
}

func GongCopyBranchLyric(mapOrigCopy map[any]any, lyricFrom *Lyric) (lyricTo *Lyric) {
	var alreadyCopied bool
	lyricTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lyricFrom)
	if alreadyCopied {
		return
	}
	lyricFrom.GongCopyBasicFields(lyricTo)

	//insertion point for the staging of instances referenced by pointers
	if lyricFrom.Extend != nil {
		lyricTo.Extend = GongCopyBranchExtend(mapOrigCopy, lyricFrom.Extend)
	}
	if lyricFrom.Footnote != nil {
		lyricTo.Footnote = GongCopyBranchFormatted_text(mapOrigCopy, lyricFrom.Footnote)
	}
	if lyricFrom.Level != nil {
		lyricTo.Level = GongCopyBranchLevel(mapOrigCopy, lyricFrom.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _elision := range lyricFrom.Elision {
		lyricTo.Elision = append(lyricTo.Elision, GongCopyBranchElision(mapOrigCopy, _elision))
	}
	for _, _text_element_data := range lyricFrom.Text {
		lyricTo.Text = append(lyricTo.Text, GongCopyBranchText_element_data(mapOrigCopy, _text_element_data))
	}

	return
}

func GongCopyBranchLyric_font(mapOrigCopy map[any]any, lyric_fontFrom *Lyric_font) (lyric_fontTo *Lyric_font) {
	var alreadyCopied bool
	lyric_fontTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lyric_fontFrom)
	if alreadyCopied {
		return
	}
	lyric_fontFrom.GongCopyBasicFields(lyric_fontTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLyric_language(mapOrigCopy map[any]any, lyric_languageFrom *Lyric_language) (lyric_languageTo *Lyric_language) {
	var alreadyCopied bool
	lyric_languageTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lyric_languageFrom)
	if alreadyCopied {
		return
	}
	lyric_languageFrom.GongCopyBasicFields(lyric_languageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMeasure_layout(mapOrigCopy map[any]any, measure_layoutFrom *Measure_layout) (measure_layoutTo *Measure_layout) {
	var alreadyCopied bool
	measure_layoutTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measure_layoutFrom)
	if alreadyCopied {
		return
	}
	measure_layoutFrom.GongCopyBasicFields(measure_layoutTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMeasure_numbering(mapOrigCopy map[any]any, measure_numberingFrom *Measure_numbering) (measure_numberingTo *Measure_numbering) {
	var alreadyCopied bool
	measure_numberingTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measure_numberingFrom)
	if alreadyCopied {
		return
	}
	measure_numberingFrom.GongCopyBasicFields(measure_numberingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMeasure_repeat(mapOrigCopy map[any]any, measure_repeatFrom *Measure_repeat) (measure_repeatTo *Measure_repeat) {
	var alreadyCopied bool
	measure_repeatTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measure_repeatFrom)
	if alreadyCopied {
		return
	}
	measure_repeatFrom.GongCopyBasicFields(measure_repeatTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMeasure_style(mapOrigCopy map[any]any, measure_styleFrom *Measure_style) (measure_styleTo *Measure_style) {
	var alreadyCopied bool
	measure_styleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measure_styleFrom)
	if alreadyCopied {
		return
	}
	measure_styleFrom.GongCopyBasicFields(measure_styleTo)

	//insertion point for the staging of instances referenced by pointers
	if measure_styleFrom.Multiple_rest != nil {
		measure_styleTo.Multiple_rest = GongCopyBranchMultiple_rest(mapOrigCopy, measure_styleFrom.Multiple_rest)
	}
	if measure_styleFrom.Measure_repeat != nil {
		measure_styleTo.Measure_repeat = GongCopyBranchMeasure_repeat(mapOrigCopy, measure_styleFrom.Measure_repeat)
	}
	if measure_styleFrom.Beat_repeat != nil {
		measure_styleTo.Beat_repeat = GongCopyBranchBeat_repeat(mapOrigCopy, measure_styleFrom.Beat_repeat)
	}
	if measure_styleFrom.Slash != nil {
		measure_styleTo.Slash = GongCopyBranchSlash(mapOrigCopy, measure_styleFrom.Slash)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMembrane(mapOrigCopy map[any]any, membraneFrom *Membrane) (membraneTo *Membrane) {
	var alreadyCopied bool
	membraneTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, membraneFrom)
	if alreadyCopied {
		return
	}
	membraneFrom.GongCopyBasicFields(membraneTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMetal(mapOrigCopy map[any]any, metalFrom *Metal) (metalTo *Metal) {
	var alreadyCopied bool
	metalTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, metalFrom)
	if alreadyCopied {
		return
	}
	metalFrom.GongCopyBasicFields(metalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMetronome(mapOrigCopy map[any]any, metronomeFrom *Metronome) (metronomeTo *Metronome) {
	var alreadyCopied bool
	metronomeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, metronomeFrom)
	if alreadyCopied {
		return
	}
	metronomeFrom.GongCopyBasicFields(metronomeTo)

	//insertion point for the staging of instances referenced by pointers
	if metronomeFrom.Per_minute != nil {
		metronomeTo.Per_minute = GongCopyBranchPer_minute(mapOrigCopy, metronomeFrom.Per_minute)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _beat_unit_tied := range metronomeFrom.Beat_unit_tied {
		metronomeTo.Beat_unit_tied = append(metronomeTo.Beat_unit_tied, GongCopyBranchBeat_unit_tied(mapOrigCopy, _beat_unit_tied))
	}
	for _, _metronome_note := range metronomeFrom.Metronome_note {
		metronomeTo.Metronome_note = append(metronomeTo.Metronome_note, GongCopyBranchMetronome_note(mapOrigCopy, _metronome_note))
	}

	return
}

func GongCopyBranchMetronome_beam(mapOrigCopy map[any]any, metronome_beamFrom *Metronome_beam) (metronome_beamTo *Metronome_beam) {
	var alreadyCopied bool
	metronome_beamTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, metronome_beamFrom)
	if alreadyCopied {
		return
	}
	metronome_beamFrom.GongCopyBasicFields(metronome_beamTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMetronome_note(mapOrigCopy map[any]any, metronome_noteFrom *Metronome_note) (metronome_noteTo *Metronome_note) {
	var alreadyCopied bool
	metronome_noteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, metronome_noteFrom)
	if alreadyCopied {
		return
	}
	metronome_noteFrom.GongCopyBasicFields(metronome_noteTo)

	//insertion point for the staging of instances referenced by pointers
	if metronome_noteFrom.Metronome_tied != nil {
		metronome_noteTo.Metronome_tied = GongCopyBranchMetronome_tied(mapOrigCopy, metronome_noteFrom.Metronome_tied)
	}
	if metronome_noteFrom.Metronome_tuplet != nil {
		metronome_noteTo.Metronome_tuplet = GongCopyBranchMetronome_tuplet(mapOrigCopy, metronome_noteFrom.Metronome_tuplet)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _metronome_beam := range metronome_noteFrom.Metronome_beam {
		metronome_noteTo.Metronome_beam = append(metronome_noteTo.Metronome_beam, GongCopyBranchMetronome_beam(mapOrigCopy, _metronome_beam))
	}

	return
}

func GongCopyBranchMetronome_tied(mapOrigCopy map[any]any, metronome_tiedFrom *Metronome_tied) (metronome_tiedTo *Metronome_tied) {
	var alreadyCopied bool
	metronome_tiedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, metronome_tiedFrom)
	if alreadyCopied {
		return
	}
	metronome_tiedFrom.GongCopyBasicFields(metronome_tiedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMetronome_tuplet(mapOrigCopy map[any]any, metronome_tupletFrom *Metronome_tuplet) (metronome_tupletTo *Metronome_tuplet) {
	var alreadyCopied bool
	metronome_tupletTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, metronome_tupletFrom)
	if alreadyCopied {
		return
	}
	metronome_tupletFrom.GongCopyBasicFields(metronome_tupletTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMidi_device(mapOrigCopy map[any]any, midi_deviceFrom *Midi_device) (midi_deviceTo *Midi_device) {
	var alreadyCopied bool
	midi_deviceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, midi_deviceFrom)
	if alreadyCopied {
		return
	}
	midi_deviceFrom.GongCopyBasicFields(midi_deviceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMidi_instrument(mapOrigCopy map[any]any, midi_instrumentFrom *Midi_instrument) (midi_instrumentTo *Midi_instrument) {
	var alreadyCopied bool
	midi_instrumentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, midi_instrumentFrom)
	if alreadyCopied {
		return
	}
	midi_instrumentFrom.GongCopyBasicFields(midi_instrumentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMiscellaneous(mapOrigCopy map[any]any, miscellaneousFrom *Miscellaneous) (miscellaneousTo *Miscellaneous) {
	var alreadyCopied bool
	miscellaneousTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, miscellaneousFrom)
	if alreadyCopied {
		return
	}
	miscellaneousFrom.GongCopyBasicFields(miscellaneousTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _miscellaneous_field := range miscellaneousFrom.Miscellaneous_field {
		miscellaneousTo.Miscellaneous_field = append(miscellaneousTo.Miscellaneous_field, GongCopyBranchMiscellaneous_field(mapOrigCopy, _miscellaneous_field))
	}

	return
}

func GongCopyBranchMiscellaneous_field(mapOrigCopy map[any]any, miscellaneous_fieldFrom *Miscellaneous_field) (miscellaneous_fieldTo *Miscellaneous_field) {
	var alreadyCopied bool
	miscellaneous_fieldTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, miscellaneous_fieldFrom)
	if alreadyCopied {
		return
	}
	miscellaneous_fieldFrom.GongCopyBasicFields(miscellaneous_fieldTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMordent(mapOrigCopy map[any]any, mordentFrom *Mordent) (mordentTo *Mordent) {
	var alreadyCopied bool
	mordentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, mordentFrom)
	if alreadyCopied {
		return
	}
	mordentFrom.GongCopyBasicFields(mordentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMultiple_rest(mapOrigCopy map[any]any, multiple_restFrom *Multiple_rest) (multiple_restTo *Multiple_rest) {
	var alreadyCopied bool
	multiple_restTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, multiple_restFrom)
	if alreadyCopied {
		return
	}
	multiple_restFrom.GongCopyBasicFields(multiple_restTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchName_display(mapOrigCopy map[any]any, name_displayFrom *Name_display) (name_displayTo *Name_display) {
	var alreadyCopied bool
	name_displayTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, name_displayFrom)
	if alreadyCopied {
		return
	}
	name_displayFrom.GongCopyBasicFields(name_displayTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formatted_text := range name_displayFrom.Display_text {
		name_displayTo.Display_text = append(name_displayTo.Display_text, GongCopyBranchFormatted_text(mapOrigCopy, _formatted_text))
	}
	for _, _accidental_text := range name_displayFrom.Accidental_text {
		name_displayTo.Accidental_text = append(name_displayTo.Accidental_text, GongCopyBranchAccidental_text(mapOrigCopy, _accidental_text))
	}

	return
}

func GongCopyBranchNon_arpeggiate(mapOrigCopy map[any]any, non_arpeggiateFrom *Non_arpeggiate) (non_arpeggiateTo *Non_arpeggiate) {
	var alreadyCopied bool
	non_arpeggiateTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, non_arpeggiateFrom)
	if alreadyCopied {
		return
	}
	non_arpeggiateFrom.GongCopyBasicFields(non_arpeggiateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNotations(mapOrigCopy map[any]any, notationsFrom *Notations) (notationsTo *Notations) {
	var alreadyCopied bool
	notationsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notationsFrom)
	if alreadyCopied {
		return
	}
	notationsFrom.GongCopyBasicFields(notationsTo)

	//insertion point for the staging of instances referenced by pointers
	if notationsFrom.Footnote != nil {
		notationsTo.Footnote = GongCopyBranchFormatted_text(mapOrigCopy, notationsFrom.Footnote)
	}
	if notationsFrom.Level != nil {
		notationsTo.Level = GongCopyBranchLevel(mapOrigCopy, notationsFrom.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tied := range notationsFrom.Tied {
		notationsTo.Tied = append(notationsTo.Tied, GongCopyBranchTied(mapOrigCopy, _tied))
	}
	for _, _slur := range notationsFrom.Slur {
		notationsTo.Slur = append(notationsTo.Slur, GongCopyBranchSlur(mapOrigCopy, _slur))
	}
	for _, _tuplet := range notationsFrom.Tuplet {
		notationsTo.Tuplet = append(notationsTo.Tuplet, GongCopyBranchTuplet(mapOrigCopy, _tuplet))
	}
	for _, _glissando := range notationsFrom.Glissando {
		notationsTo.Glissando = append(notationsTo.Glissando, GongCopyBranchGlissando(mapOrigCopy, _glissando))
	}
	for _, _slide := range notationsFrom.Slide {
		notationsTo.Slide = append(notationsTo.Slide, GongCopyBranchSlide(mapOrigCopy, _slide))
	}
	for _, _ornaments := range notationsFrom.Ornaments {
		notationsTo.Ornaments = append(notationsTo.Ornaments, GongCopyBranchOrnaments(mapOrigCopy, _ornaments))
	}
	for _, _technical := range notationsFrom.Technical {
		notationsTo.Technical = append(notationsTo.Technical, GongCopyBranchTechnical(mapOrigCopy, _technical))
	}
	for _, _articulations := range notationsFrom.Articulations {
		notationsTo.Articulations = append(notationsTo.Articulations, GongCopyBranchArticulations(mapOrigCopy, _articulations))
	}
	for _, _dynamics := range notationsFrom.Dynamics {
		notationsTo.Dynamics = append(notationsTo.Dynamics, GongCopyBranchDynamics(mapOrigCopy, _dynamics))
	}
	for _, _fermata := range notationsFrom.Fermata {
		notationsTo.Fermata = append(notationsTo.Fermata, GongCopyBranchFermata(mapOrigCopy, _fermata))
	}
	for _, _arpeggiate := range notationsFrom.Arpeggiate {
		notationsTo.Arpeggiate = append(notationsTo.Arpeggiate, GongCopyBranchArpeggiate(mapOrigCopy, _arpeggiate))
	}
	for _, _non_arpeggiate := range notationsFrom.Non_arpeggiate {
		notationsTo.Non_arpeggiate = append(notationsTo.Non_arpeggiate, GongCopyBranchNon_arpeggiate(mapOrigCopy, _non_arpeggiate))
	}
	for _, _accidental_mark := range notationsFrom.Accidental_mark {
		notationsTo.Accidental_mark = append(notationsTo.Accidental_mark, GongCopyBranchAccidental_mark(mapOrigCopy, _accidental_mark))
	}
	for _, _other_notation := range notationsFrom.Other_notation {
		notationsTo.Other_notation = append(notationsTo.Other_notation, GongCopyBranchOther_notation(mapOrigCopy, _other_notation))
	}

	return
}

func GongCopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {
	var alreadyCopied bool
	noteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteFrom)
	if alreadyCopied {
		return
	}
	noteFrom.GongCopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers
	if noteFrom.Grace != nil {
		noteTo.Grace = GongCopyBranchGrace(mapOrigCopy, noteFrom.Grace)
	}
	if noteFrom.Pitch != nil {
		noteTo.Pitch = GongCopyBranchPitch(mapOrigCopy, noteFrom.Pitch)
	}
	if noteFrom.Unpitched != nil {
		noteTo.Unpitched = GongCopyBranchUnpitched(mapOrigCopy, noteFrom.Unpitched)
	}
	if noteFrom.Rest != nil {
		noteTo.Rest = GongCopyBranchRest(mapOrigCopy, noteFrom.Rest)
	}
	if noteFrom.Tie != nil {
		noteTo.Tie = GongCopyBranchTie(mapOrigCopy, noteFrom.Tie)
	}
	if noteFrom.Footnote != nil {
		noteTo.Footnote = GongCopyBranchFormatted_text(mapOrigCopy, noteFrom.Footnote)
	}
	if noteFrom.Level != nil {
		noteTo.Level = GongCopyBranchLevel(mapOrigCopy, noteFrom.Level)
	}
	if noteFrom.Type != nil {
		noteTo.Type = GongCopyBranchNote_type(mapOrigCopy, noteFrom.Type)
	}
	if noteFrom.Accidental != nil {
		noteTo.Accidental = GongCopyBranchAccidental(mapOrigCopy, noteFrom.Accidental)
	}
	if noteFrom.Time_modification != nil {
		noteTo.Time_modification = GongCopyBranchTime_modification(mapOrigCopy, noteFrom.Time_modification)
	}
	if noteFrom.Stem != nil {
		noteTo.Stem = GongCopyBranchStem(mapOrigCopy, noteFrom.Stem)
	}
	if noteFrom.Notehead != nil {
		noteTo.Notehead = GongCopyBranchNotehead(mapOrigCopy, noteFrom.Notehead)
	}
	if noteFrom.Notehead_text != nil {
		noteTo.Notehead_text = GongCopyBranchNotehead_text(mapOrigCopy, noteFrom.Notehead_text)
	}
	if noteFrom.Beam != nil {
		noteTo.Beam = GongCopyBranchBeam(mapOrigCopy, noteFrom.Beam)
	}
	if noteFrom.Play != nil {
		noteTo.Play = GongCopyBranchPlay(mapOrigCopy, noteFrom.Play)
	}
	if noteFrom.Listen != nil {
		noteTo.Listen = GongCopyBranchListen(mapOrigCopy, noteFrom.Listen)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument := range noteFrom.Instrument {
		noteTo.Instrument = append(noteTo.Instrument, GongCopyBranchInstrument(mapOrigCopy, _instrument))
	}
	for _, _empty_placement := range noteFrom.Dot {
		noteTo.Dot = append(noteTo.Dot, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _notations := range noteFrom.Notations {
		noteTo.Notations = append(noteTo.Notations, GongCopyBranchNotations(mapOrigCopy, _notations))
	}
	for _, _lyric := range noteFrom.Lyric {
		noteTo.Lyric = append(noteTo.Lyric, GongCopyBranchLyric(mapOrigCopy, _lyric))
	}

	return
}

func GongCopyBranchNote_size(mapOrigCopy map[any]any, note_sizeFrom *Note_size) (note_sizeTo *Note_size) {
	var alreadyCopied bool
	note_sizeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, note_sizeFrom)
	if alreadyCopied {
		return
	}
	note_sizeFrom.GongCopyBasicFields(note_sizeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNote_type(mapOrigCopy map[any]any, note_typeFrom *Note_type) (note_typeTo *Note_type) {
	var alreadyCopied bool
	note_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, note_typeFrom)
	if alreadyCopied {
		return
	}
	note_typeFrom.GongCopyBasicFields(note_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNotehead(mapOrigCopy map[any]any, noteheadFrom *Notehead) (noteheadTo *Notehead) {
	var alreadyCopied bool
	noteheadTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteheadFrom)
	if alreadyCopied {
		return
	}
	noteheadFrom.GongCopyBasicFields(noteheadTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNotehead_text(mapOrigCopy map[any]any, notehead_textFrom *Notehead_text) (notehead_textTo *Notehead_text) {
	var alreadyCopied bool
	notehead_textTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notehead_textFrom)
	if alreadyCopied {
		return
	}
	notehead_textFrom.GongCopyBasicFields(notehead_textTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formatted_text := range notehead_textFrom.Display_text {
		notehead_textTo.Display_text = append(notehead_textTo.Display_text, GongCopyBranchFormatted_text(mapOrigCopy, _formatted_text))
	}
	for _, _accidental_text := range notehead_textFrom.Accidental_text {
		notehead_textTo.Accidental_text = append(notehead_textTo.Accidental_text, GongCopyBranchAccidental_text(mapOrigCopy, _accidental_text))
	}

	return
}

func GongCopyBranchNumeral(mapOrigCopy map[any]any, numeralFrom *Numeral) (numeralTo *Numeral) {
	var alreadyCopied bool
	numeralTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, numeralFrom)
	if alreadyCopied {
		return
	}
	numeralFrom.GongCopyBasicFields(numeralTo)

	//insertion point for the staging of instances referenced by pointers
	if numeralFrom.Numeral_root != nil {
		numeralTo.Numeral_root = GongCopyBranchNumeral_root(mapOrigCopy, numeralFrom.Numeral_root)
	}
	if numeralFrom.Numeral_alter != nil {
		numeralTo.Numeral_alter = GongCopyBranchHarmony_alter(mapOrigCopy, numeralFrom.Numeral_alter)
	}
	if numeralFrom.Numeral_key != nil {
		numeralTo.Numeral_key = GongCopyBranchNumeral_key(mapOrigCopy, numeralFrom.Numeral_key)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNumeral_key(mapOrigCopy map[any]any, numeral_keyFrom *Numeral_key) (numeral_keyTo *Numeral_key) {
	var alreadyCopied bool
	numeral_keyTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, numeral_keyFrom)
	if alreadyCopied {
		return
	}
	numeral_keyFrom.GongCopyBasicFields(numeral_keyTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNumeral_root(mapOrigCopy map[any]any, numeral_rootFrom *Numeral_root) (numeral_rootTo *Numeral_root) {
	var alreadyCopied bool
	numeral_rootTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, numeral_rootFrom)
	if alreadyCopied {
		return
	}
	numeral_rootFrom.GongCopyBasicFields(numeral_rootTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOctave_shift(mapOrigCopy map[any]any, octave_shiftFrom *Octave_shift) (octave_shiftTo *Octave_shift) {
	var alreadyCopied bool
	octave_shiftTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, octave_shiftFrom)
	if alreadyCopied {
		return
	}
	octave_shiftFrom.GongCopyBasicFields(octave_shiftTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOffset(mapOrigCopy map[any]any, offsetFrom *Offset) (offsetTo *Offset) {
	var alreadyCopied bool
	offsetTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, offsetFrom)
	if alreadyCopied {
		return
	}
	offsetFrom.GongCopyBasicFields(offsetTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOpus(mapOrigCopy map[any]any, opusFrom *Opus) (opusTo *Opus) {
	var alreadyCopied bool
	opusTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, opusFrom)
	if alreadyCopied {
		return
	}
	opusFrom.GongCopyBasicFields(opusTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOrnaments(mapOrigCopy map[any]any, ornamentsFrom *Ornaments) (ornamentsTo *Ornaments) {
	var alreadyCopied bool
	ornamentsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ornamentsFrom)
	if alreadyCopied {
		return
	}
	ornamentsFrom.GongCopyBasicFields(ornamentsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty_trill_sound := range ornamentsFrom.Trill_mark {
		ornamentsTo.Trill_mark = append(ornamentsTo.Trill_mark, GongCopyBranchEmpty_trill_sound(mapOrigCopy, _empty_trill_sound))
	}
	for _, _horizontal_turn := range ornamentsFrom.Turn {
		ornamentsTo.Turn = append(ornamentsTo.Turn, GongCopyBranchHorizontal_turn(mapOrigCopy, _horizontal_turn))
	}
	for _, _horizontal_turn := range ornamentsFrom.Delayed_turn {
		ornamentsTo.Delayed_turn = append(ornamentsTo.Delayed_turn, GongCopyBranchHorizontal_turn(mapOrigCopy, _horizontal_turn))
	}
	for _, _horizontal_turn := range ornamentsFrom.Inverted_turn {
		ornamentsTo.Inverted_turn = append(ornamentsTo.Inverted_turn, GongCopyBranchHorizontal_turn(mapOrigCopy, _horizontal_turn))
	}
	for _, _horizontal_turn := range ornamentsFrom.Delayed_inverted_turn {
		ornamentsTo.Delayed_inverted_turn = append(ornamentsTo.Delayed_inverted_turn, GongCopyBranchHorizontal_turn(mapOrigCopy, _horizontal_turn))
	}
	for _, _empty_trill_sound := range ornamentsFrom.Vertical_turn {
		ornamentsTo.Vertical_turn = append(ornamentsTo.Vertical_turn, GongCopyBranchEmpty_trill_sound(mapOrigCopy, _empty_trill_sound))
	}
	for _, _empty_trill_sound := range ornamentsFrom.Inverted_vertical_turn {
		ornamentsTo.Inverted_vertical_turn = append(ornamentsTo.Inverted_vertical_turn, GongCopyBranchEmpty_trill_sound(mapOrigCopy, _empty_trill_sound))
	}
	for _, _empty_trill_sound := range ornamentsFrom.Shake {
		ornamentsTo.Shake = append(ornamentsTo.Shake, GongCopyBranchEmpty_trill_sound(mapOrigCopy, _empty_trill_sound))
	}
	for _, _wavy_line := range ornamentsFrom.Wavy_line {
		ornamentsTo.Wavy_line = append(ornamentsTo.Wavy_line, GongCopyBranchWavy_line(mapOrigCopy, _wavy_line))
	}
	for _, _mordent := range ornamentsFrom.Mordent {
		ornamentsTo.Mordent = append(ornamentsTo.Mordent, GongCopyBranchMordent(mapOrigCopy, _mordent))
	}
	for _, _mordent := range ornamentsFrom.Inverted_mordent {
		ornamentsTo.Inverted_mordent = append(ornamentsTo.Inverted_mordent, GongCopyBranchMordent(mapOrigCopy, _mordent))
	}
	for _, _empty_placement := range ornamentsFrom.Schleifer {
		ornamentsTo.Schleifer = append(ornamentsTo.Schleifer, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _tremolo := range ornamentsFrom.Tremolo {
		ornamentsTo.Tremolo = append(ornamentsTo.Tremolo, GongCopyBranchTremolo(mapOrigCopy, _tremolo))
	}
	for _, _empty_trill_sound := range ornamentsFrom.Haydn {
		ornamentsTo.Haydn = append(ornamentsTo.Haydn, GongCopyBranchEmpty_trill_sound(mapOrigCopy, _empty_trill_sound))
	}
	for _, _other_placement_text := range ornamentsFrom.Other_ornament {
		ornamentsTo.Other_ornament = append(ornamentsTo.Other_ornament, GongCopyBranchOther_placement_text(mapOrigCopy, _other_placement_text))
	}
	for _, _accidental_mark := range ornamentsFrom.Accidental_mark {
		ornamentsTo.Accidental_mark = append(ornamentsTo.Accidental_mark, GongCopyBranchAccidental_mark(mapOrigCopy, _accidental_mark))
	}

	return
}

func GongCopyBranchOther_appearance(mapOrigCopy map[any]any, other_appearanceFrom *Other_appearance) (other_appearanceTo *Other_appearance) {
	var alreadyCopied bool
	other_appearanceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, other_appearanceFrom)
	if alreadyCopied {
		return
	}
	other_appearanceFrom.GongCopyBasicFields(other_appearanceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOther_direction(mapOrigCopy map[any]any, other_directionFrom *Other_direction) (other_directionTo *Other_direction) {
	var alreadyCopied bool
	other_directionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, other_directionFrom)
	if alreadyCopied {
		return
	}
	other_directionFrom.GongCopyBasicFields(other_directionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOther_listening(mapOrigCopy map[any]any, other_listeningFrom *Other_listening) (other_listeningTo *Other_listening) {
	var alreadyCopied bool
	other_listeningTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, other_listeningFrom)
	if alreadyCopied {
		return
	}
	other_listeningFrom.GongCopyBasicFields(other_listeningTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOther_notation(mapOrigCopy map[any]any, other_notationFrom *Other_notation) (other_notationTo *Other_notation) {
	var alreadyCopied bool
	other_notationTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, other_notationFrom)
	if alreadyCopied {
		return
	}
	other_notationFrom.GongCopyBasicFields(other_notationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOther_placement_text(mapOrigCopy map[any]any, other_placement_textFrom *Other_placement_text) (other_placement_textTo *Other_placement_text) {
	var alreadyCopied bool
	other_placement_textTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, other_placement_textFrom)
	if alreadyCopied {
		return
	}
	other_placement_textFrom.GongCopyBasicFields(other_placement_textTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOther_play(mapOrigCopy map[any]any, other_playFrom *Other_play) (other_playTo *Other_play) {
	var alreadyCopied bool
	other_playTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, other_playFrom)
	if alreadyCopied {
		return
	}
	other_playFrom.GongCopyBasicFields(other_playTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOther_text(mapOrigCopy map[any]any, other_textFrom *Other_text) (other_textTo *Other_text) {
	var alreadyCopied bool
	other_textTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, other_textFrom)
	if alreadyCopied {
		return
	}
	other_textFrom.GongCopyBasicFields(other_textTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPage_layout(mapOrigCopy map[any]any, page_layoutFrom *Page_layout) (page_layoutTo *Page_layout) {
	var alreadyCopied bool
	page_layoutTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, page_layoutFrom)
	if alreadyCopied {
		return
	}
	page_layoutFrom.GongCopyBasicFields(page_layoutTo)

	//insertion point for the staging of instances referenced by pointers
	if page_layoutFrom.Page_margins != nil {
		page_layoutTo.Page_margins = GongCopyBranchPage_margins(mapOrigCopy, page_layoutFrom.Page_margins)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPage_margins(mapOrigCopy map[any]any, page_marginsFrom *Page_margins) (page_marginsTo *Page_margins) {
	var alreadyCopied bool
	page_marginsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, page_marginsFrom)
	if alreadyCopied {
		return
	}
	page_marginsFrom.GongCopyBasicFields(page_marginsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPart_clef(mapOrigCopy map[any]any, part_clefFrom *Part_clef) (part_clefTo *Part_clef) {
	var alreadyCopied bool
	part_clefTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, part_clefFrom)
	if alreadyCopied {
		return
	}
	part_clefFrom.GongCopyBasicFields(part_clefTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPart_group(mapOrigCopy map[any]any, part_groupFrom *Part_group) (part_groupTo *Part_group) {
	var alreadyCopied bool
	part_groupTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, part_groupFrom)
	if alreadyCopied {
		return
	}
	part_groupFrom.GongCopyBasicFields(part_groupTo)

	//insertion point for the staging of instances referenced by pointers
	if part_groupFrom.Group_name != nil {
		part_groupTo.Group_name = GongCopyBranchGroup_name(mapOrigCopy, part_groupFrom.Group_name)
	}
	if part_groupFrom.Group_name_display != nil {
		part_groupTo.Group_name_display = GongCopyBranchName_display(mapOrigCopy, part_groupFrom.Group_name_display)
	}
	if part_groupFrom.Group_abbreviation != nil {
		part_groupTo.Group_abbreviation = GongCopyBranchGroup_name(mapOrigCopy, part_groupFrom.Group_abbreviation)
	}
	if part_groupFrom.Group_abbreviation_display != nil {
		part_groupTo.Group_abbreviation_display = GongCopyBranchName_display(mapOrigCopy, part_groupFrom.Group_abbreviation_display)
	}
	if part_groupFrom.Group_symbol != nil {
		part_groupTo.Group_symbol = GongCopyBranchGroup_symbol(mapOrigCopy, part_groupFrom.Group_symbol)
	}
	if part_groupFrom.Group_barline != nil {
		part_groupTo.Group_barline = GongCopyBranchGroup_barline(mapOrigCopy, part_groupFrom.Group_barline)
	}
	if part_groupFrom.Footnote != nil {
		part_groupTo.Footnote = GongCopyBranchFormatted_text(mapOrigCopy, part_groupFrom.Footnote)
	}
	if part_groupFrom.Level != nil {
		part_groupTo.Level = GongCopyBranchLevel(mapOrigCopy, part_groupFrom.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPart_link(mapOrigCopy map[any]any, part_linkFrom *Part_link) (part_linkTo *Part_link) {
	var alreadyCopied bool
	part_linkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, part_linkFrom)
	if alreadyCopied {
		return
	}
	part_linkFrom.GongCopyBasicFields(part_linkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument_link := range part_linkFrom.Instrument_link {
		part_linkTo.Instrument_link = append(part_linkTo.Instrument_link, GongCopyBranchInstrument_link(mapOrigCopy, _instrument_link))
	}

	return
}

func GongCopyBranchPart_list(mapOrigCopy map[any]any, part_listFrom *Part_list) (part_listTo *Part_list) {
	var alreadyCopied bool
	part_listTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, part_listFrom)
	if alreadyCopied {
		return
	}
	part_listFrom.GongCopyBasicFields(part_listTo)

	//insertion point for the staging of instances referenced by pointers
	if part_listFrom.Part_group != nil {
		part_listTo.Part_group = GongCopyBranchPart_group(mapOrigCopy, part_listFrom.Part_group)
	}
	if part_listFrom.Score_part != nil {
		part_listTo.Score_part = GongCopyBranchScore_part(mapOrigCopy, part_listFrom.Score_part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPart_name(mapOrigCopy map[any]any, part_nameFrom *Part_name) (part_nameTo *Part_name) {
	var alreadyCopied bool
	part_nameTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, part_nameFrom)
	if alreadyCopied {
		return
	}
	part_nameFrom.GongCopyBasicFields(part_nameTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPart_symbol(mapOrigCopy map[any]any, part_symbolFrom *Part_symbol) (part_symbolTo *Part_symbol) {
	var alreadyCopied bool
	part_symbolTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, part_symbolFrom)
	if alreadyCopied {
		return
	}
	part_symbolFrom.GongCopyBasicFields(part_symbolTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPart_transpose(mapOrigCopy map[any]any, part_transposeFrom *Part_transpose) (part_transposeTo *Part_transpose) {
	var alreadyCopied bool
	part_transposeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, part_transposeFrom)
	if alreadyCopied {
		return
	}
	part_transposeFrom.GongCopyBasicFields(part_transposeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPedal(mapOrigCopy map[any]any, pedalFrom *Pedal) (pedalTo *Pedal) {
	var alreadyCopied bool
	pedalTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pedalFrom)
	if alreadyCopied {
		return
	}
	pedalFrom.GongCopyBasicFields(pedalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPedal_tuning(mapOrigCopy map[any]any, pedal_tuningFrom *Pedal_tuning) (pedal_tuningTo *Pedal_tuning) {
	var alreadyCopied bool
	pedal_tuningTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pedal_tuningFrom)
	if alreadyCopied {
		return
	}
	pedal_tuningFrom.GongCopyBasicFields(pedal_tuningTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPer_minute(mapOrigCopy map[any]any, per_minuteFrom *Per_minute) (per_minuteTo *Per_minute) {
	var alreadyCopied bool
	per_minuteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, per_minuteFrom)
	if alreadyCopied {
		return
	}
	per_minuteFrom.GongCopyBasicFields(per_minuteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPercussion(mapOrigCopy map[any]any, percussionFrom *Percussion) (percussionTo *Percussion) {
	var alreadyCopied bool
	percussionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, percussionFrom)
	if alreadyCopied {
		return
	}
	percussionFrom.GongCopyBasicFields(percussionTo)

	//insertion point for the staging of instances referenced by pointers
	if percussionFrom.Glass != nil {
		percussionTo.Glass = GongCopyBranchGlass(mapOrigCopy, percussionFrom.Glass)
	}
	if percussionFrom.Metal != nil {
		percussionTo.Metal = GongCopyBranchMetal(mapOrigCopy, percussionFrom.Metal)
	}
	if percussionFrom.Wood != nil {
		percussionTo.Wood = GongCopyBranchWood(mapOrigCopy, percussionFrom.Wood)
	}
	if percussionFrom.Pitched != nil {
		percussionTo.Pitched = GongCopyBranchPitched(mapOrigCopy, percussionFrom.Pitched)
	}
	if percussionFrom.Membrane != nil {
		percussionTo.Membrane = GongCopyBranchMembrane(mapOrigCopy, percussionFrom.Membrane)
	}
	if percussionFrom.Effect != nil {
		percussionTo.Effect = GongCopyBranchEffect(mapOrigCopy, percussionFrom.Effect)
	}
	if percussionFrom.Timpani != nil {
		percussionTo.Timpani = GongCopyBranchTimpani(mapOrigCopy, percussionFrom.Timpani)
	}
	if percussionFrom.Beater != nil {
		percussionTo.Beater = GongCopyBranchBeater(mapOrigCopy, percussionFrom.Beater)
	}
	if percussionFrom.Stick != nil {
		percussionTo.Stick = GongCopyBranchStick(mapOrigCopy, percussionFrom.Stick)
	}
	if percussionFrom.Other_percussion != nil {
		percussionTo.Other_percussion = GongCopyBranchOther_text(mapOrigCopy, percussionFrom.Other_percussion)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPitch(mapOrigCopy map[any]any, pitchFrom *Pitch) (pitchTo *Pitch) {
	var alreadyCopied bool
	pitchTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pitchFrom)
	if alreadyCopied {
		return
	}
	pitchFrom.GongCopyBasicFields(pitchTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPitched(mapOrigCopy map[any]any, pitchedFrom *Pitched) (pitchedTo *Pitched) {
	var alreadyCopied bool
	pitchedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pitchedFrom)
	if alreadyCopied {
		return
	}
	pitchedFrom.GongCopyBasicFields(pitchedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPlacement_text(mapOrigCopy map[any]any, placement_textFrom *Placement_text) (placement_textTo *Placement_text) {
	var alreadyCopied bool
	placement_textTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, placement_textFrom)
	if alreadyCopied {
		return
	}
	placement_textFrom.GongCopyBasicFields(placement_textTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPlay(mapOrigCopy map[any]any, playFrom *Play) (playTo *Play) {
	var alreadyCopied bool
	playTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, playFrom)
	if alreadyCopied {
		return
	}
	playFrom.GongCopyBasicFields(playTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _other_play := range playFrom.Other_play {
		playTo.Other_play = append(playTo.Other_play, GongCopyBranchOther_play(mapOrigCopy, _other_play))
	}

	return
}

func GongCopyBranchPlayer(mapOrigCopy map[any]any, playerFrom *Player) (playerTo *Player) {
	var alreadyCopied bool
	playerTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, playerFrom)
	if alreadyCopied {
		return
	}
	playerFrom.GongCopyBasicFields(playerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPrincipal_voice(mapOrigCopy map[any]any, principal_voiceFrom *Principal_voice) (principal_voiceTo *Principal_voice) {
	var alreadyCopied bool
	principal_voiceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, principal_voiceFrom)
	if alreadyCopied {
		return
	}
	principal_voiceFrom.GongCopyBasicFields(principal_voiceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPrint(mapOrigCopy map[any]any, printFrom *Print) (printTo *Print) {
	var alreadyCopied bool
	printTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, printFrom)
	if alreadyCopied {
		return
	}
	printFrom.GongCopyBasicFields(printTo)

	//insertion point for the staging of instances referenced by pointers
	if printFrom.Page_layout != nil {
		printTo.Page_layout = GongCopyBranchPage_layout(mapOrigCopy, printFrom.Page_layout)
	}
	if printFrom.System_layout != nil {
		printTo.System_layout = GongCopyBranchSystem_layout(mapOrigCopy, printFrom.System_layout)
	}
	if printFrom.Measure_layout != nil {
		printTo.Measure_layout = GongCopyBranchMeasure_layout(mapOrigCopy, printFrom.Measure_layout)
	}
	if printFrom.Measure_numbering != nil {
		printTo.Measure_numbering = GongCopyBranchMeasure_numbering(mapOrigCopy, printFrom.Measure_numbering)
	}
	if printFrom.Part_name_display != nil {
		printTo.Part_name_display = GongCopyBranchName_display(mapOrigCopy, printFrom.Part_name_display)
	}
	if printFrom.Part_abbreviation_display != nil {
		printTo.Part_abbreviation_display = GongCopyBranchName_display(mapOrigCopy, printFrom.Part_abbreviation_display)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _staff_layout := range printFrom.Staff_layout {
		printTo.Staff_layout = append(printTo.Staff_layout, GongCopyBranchStaff_layout(mapOrigCopy, _staff_layout))
	}

	return
}

func GongCopyBranchRelease(mapOrigCopy map[any]any, releaseFrom *Release) (releaseTo *Release) {
	var alreadyCopied bool
	releaseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, releaseFrom)
	if alreadyCopied {
		return
	}
	releaseFrom.GongCopyBasicFields(releaseTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRepeat(mapOrigCopy map[any]any, repeatFrom *Repeat) (repeatTo *Repeat) {
	var alreadyCopied bool
	repeatTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, repeatFrom)
	if alreadyCopied {
		return
	}
	repeatFrom.GongCopyBasicFields(repeatTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRest(mapOrigCopy map[any]any, restFrom *Rest) (restTo *Rest) {
	var alreadyCopied bool
	restTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, restFrom)
	if alreadyCopied {
		return
	}
	restFrom.GongCopyBasicFields(restTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRoot(mapOrigCopy map[any]any, rootFrom *Root) (rootTo *Root) {
	var alreadyCopied bool
	rootTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rootFrom)
	if alreadyCopied {
		return
	}
	rootFrom.GongCopyBasicFields(rootTo)

	//insertion point for the staging of instances referenced by pointers
	if rootFrom.Root_step != nil {
		rootTo.Root_step = GongCopyBranchRoot_step(mapOrigCopy, rootFrom.Root_step)
	}
	if rootFrom.Root_alter != nil {
		rootTo.Root_alter = GongCopyBranchHarmony_alter(mapOrigCopy, rootFrom.Root_alter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRoot_step(mapOrigCopy map[any]any, root_stepFrom *Root_step) (root_stepTo *Root_step) {
	var alreadyCopied bool
	root_stepTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, root_stepFrom)
	if alreadyCopied {
		return
	}
	root_stepFrom.GongCopyBasicFields(root_stepTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchScaling(mapOrigCopy map[any]any, scalingFrom *Scaling) (scalingTo *Scaling) {
	var alreadyCopied bool
	scalingTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, scalingFrom)
	if alreadyCopied {
		return
	}
	scalingFrom.GongCopyBasicFields(scalingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchScordatura(mapOrigCopy map[any]any, scordaturaFrom *Scordatura) (scordaturaTo *Scordatura) {
	var alreadyCopied bool
	scordaturaTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, scordaturaFrom)
	if alreadyCopied {
		return
	}
	scordaturaFrom.GongCopyBasicFields(scordaturaTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _accord := range scordaturaFrom.Accord {
		scordaturaTo.Accord = append(scordaturaTo.Accord, GongCopyBranchAccord(mapOrigCopy, _accord))
	}

	return
}

func GongCopyBranchScore_instrument(mapOrigCopy map[any]any, score_instrumentFrom *Score_instrument) (score_instrumentTo *Score_instrument) {
	var alreadyCopied bool
	score_instrumentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, score_instrumentFrom)
	if alreadyCopied {
		return
	}
	score_instrumentFrom.GongCopyBasicFields(score_instrumentTo)

	//insertion point for the staging of instances referenced by pointers
	if score_instrumentFrom.Virtual_instrument != nil {
		score_instrumentTo.Virtual_instrument = GongCopyBranchVirtual_instrument(mapOrigCopy, score_instrumentFrom.Virtual_instrument)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchScore_part(mapOrigCopy map[any]any, score_partFrom *Score_part) (score_partTo *Score_part) {
	var alreadyCopied bool
	score_partTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, score_partFrom)
	if alreadyCopied {
		return
	}
	score_partFrom.GongCopyBasicFields(score_partTo)

	//insertion point for the staging of instances referenced by pointers
	if score_partFrom.Identification != nil {
		score_partTo.Identification = GongCopyBranchIdentification(mapOrigCopy, score_partFrom.Identification)
	}
	if score_partFrom.Part_name != nil {
		score_partTo.Part_name = GongCopyBranchPart_name(mapOrigCopy, score_partFrom.Part_name)
	}
	if score_partFrom.Part_name_display != nil {
		score_partTo.Part_name_display = GongCopyBranchName_display(mapOrigCopy, score_partFrom.Part_name_display)
	}
	if score_partFrom.Part_abbreviation != nil {
		score_partTo.Part_abbreviation = GongCopyBranchPart_name(mapOrigCopy, score_partFrom.Part_abbreviation)
	}
	if score_partFrom.Part_abbreviation_display != nil {
		score_partTo.Part_abbreviation_display = GongCopyBranchName_display(mapOrigCopy, score_partFrom.Part_abbreviation_display)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part_link := range score_partFrom.Part_link {
		score_partTo.Part_link = append(score_partTo.Part_link, GongCopyBranchPart_link(mapOrigCopy, _part_link))
	}
	for _, _score_instrument := range score_partFrom.Score_instrument {
		score_partTo.Score_instrument = append(score_partTo.Score_instrument, GongCopyBranchScore_instrument(mapOrigCopy, _score_instrument))
	}
	for _, _player := range score_partFrom.Player {
		score_partTo.Player = append(score_partTo.Player, GongCopyBranchPlayer(mapOrigCopy, _player))
	}
	for _, _midi_device := range score_partFrom.Midi_device {
		score_partTo.Midi_device = append(score_partTo.Midi_device, GongCopyBranchMidi_device(mapOrigCopy, _midi_device))
	}
	for _, _midi_instrument := range score_partFrom.Midi_instrument {
		score_partTo.Midi_instrument = append(score_partTo.Midi_instrument, GongCopyBranchMidi_instrument(mapOrigCopy, _midi_instrument))
	}

	return
}

func GongCopyBranchScore_partwise(mapOrigCopy map[any]any, score_partwiseFrom *Score_partwise) (score_partwiseTo *Score_partwise) {
	var alreadyCopied bool
	score_partwiseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, score_partwiseFrom)
	if alreadyCopied {
		return
	}
	score_partwiseFrom.GongCopyBasicFields(score_partwiseTo)

	//insertion point for the staging of instances referenced by pointers
	if score_partwiseFrom.Work != nil {
		score_partwiseTo.Work = GongCopyBranchWork(mapOrigCopy, score_partwiseFrom.Work)
	}
	if score_partwiseFrom.Identification != nil {
		score_partwiseTo.Identification = GongCopyBranchIdentification(mapOrigCopy, score_partwiseFrom.Identification)
	}
	if score_partwiseFrom.Defaults != nil {
		score_partwiseTo.Defaults = GongCopyBranchDefaults(mapOrigCopy, score_partwiseFrom.Defaults)
	}
	if score_partwiseFrom.Part_list != nil {
		score_partwiseTo.Part_list = GongCopyBranchPart_list(mapOrigCopy, score_partwiseFrom.Part_list)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range score_partwiseFrom.Credit {
		score_partwiseTo.Credit = append(score_partwiseTo.Credit, GongCopyBranchCredit(mapOrigCopy, _credit))
	}
	for _, _a_part := range score_partwiseFrom.Part {
		score_partwiseTo.Part = append(score_partwiseTo.Part, GongCopyBranchA_part(mapOrigCopy, _a_part))
	}

	return
}

func GongCopyBranchScore_timewise(mapOrigCopy map[any]any, score_timewiseFrom *Score_timewise) (score_timewiseTo *Score_timewise) {
	var alreadyCopied bool
	score_timewiseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, score_timewiseFrom)
	if alreadyCopied {
		return
	}
	score_timewiseFrom.GongCopyBasicFields(score_timewiseTo)

	//insertion point for the staging of instances referenced by pointers
	if score_timewiseFrom.Work != nil {
		score_timewiseTo.Work = GongCopyBranchWork(mapOrigCopy, score_timewiseFrom.Work)
	}
	if score_timewiseFrom.Identification != nil {
		score_timewiseTo.Identification = GongCopyBranchIdentification(mapOrigCopy, score_timewiseFrom.Identification)
	}
	if score_timewiseFrom.Defaults != nil {
		score_timewiseTo.Defaults = GongCopyBranchDefaults(mapOrigCopy, score_timewiseFrom.Defaults)
	}
	if score_timewiseFrom.Part_list != nil {
		score_timewiseTo.Part_list = GongCopyBranchPart_list(mapOrigCopy, score_timewiseFrom.Part_list)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range score_timewiseFrom.Credit {
		score_timewiseTo.Credit = append(score_timewiseTo.Credit, GongCopyBranchCredit(mapOrigCopy, _credit))
	}
	for _, _a_measure_1 := range score_timewiseFrom.Measure {
		score_timewiseTo.Measure = append(score_timewiseTo.Measure, GongCopyBranchA_measure_1(mapOrigCopy, _a_measure_1))
	}

	return
}

func GongCopyBranchSegno(mapOrigCopy map[any]any, segnoFrom *Segno) (segnoTo *Segno) {
	var alreadyCopied bool
	segnoTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, segnoFrom)
	if alreadyCopied {
		return
	}
	segnoFrom.GongCopyBasicFields(segnoTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSlash(mapOrigCopy map[any]any, slashFrom *Slash) (slashTo *Slash) {
	var alreadyCopied bool
	slashTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, slashFrom)
	if alreadyCopied {
		return
	}
	slashFrom.GongCopyBasicFields(slashTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSlide(mapOrigCopy map[any]any, slideFrom *Slide) (slideTo *Slide) {
	var alreadyCopied bool
	slideTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, slideFrom)
	if alreadyCopied {
		return
	}
	slideFrom.GongCopyBasicFields(slideTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSlur(mapOrigCopy map[any]any, slurFrom *Slur) (slurTo *Slur) {
	var alreadyCopied bool
	slurTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, slurFrom)
	if alreadyCopied {
		return
	}
	slurFrom.GongCopyBasicFields(slurTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSound(mapOrigCopy map[any]any, soundFrom *Sound) (soundTo *Sound) {
	var alreadyCopied bool
	soundTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, soundFrom)
	if alreadyCopied {
		return
	}
	soundFrom.GongCopyBasicFields(soundTo)

	//insertion point for the staging of instances referenced by pointers
	if soundFrom.Swing != nil {
		soundTo.Swing = GongCopyBranchSwing(mapOrigCopy, soundFrom.Swing)
	}
	if soundFrom.Offset != nil {
		soundTo.Offset = GongCopyBranchOffset(mapOrigCopy, soundFrom.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument_change := range soundFrom.Instrument_change {
		soundTo.Instrument_change = append(soundTo.Instrument_change, GongCopyBranchInstrument_change(mapOrigCopy, _instrument_change))
	}
	for _, _midi_device := range soundFrom.Midi_device {
		soundTo.Midi_device = append(soundTo.Midi_device, GongCopyBranchMidi_device(mapOrigCopy, _midi_device))
	}
	for _, _midi_instrument := range soundFrom.Midi_instrument {
		soundTo.Midi_instrument = append(soundTo.Midi_instrument, GongCopyBranchMidi_instrument(mapOrigCopy, _midi_instrument))
	}
	for _, _play := range soundFrom.Play {
		soundTo.Play = append(soundTo.Play, GongCopyBranchPlay(mapOrigCopy, _play))
	}

	return
}

func GongCopyBranchStaff_details(mapOrigCopy map[any]any, staff_detailsFrom *Staff_details) (staff_detailsTo *Staff_details) {
	var alreadyCopied bool
	staff_detailsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, staff_detailsFrom)
	if alreadyCopied {
		return
	}
	staff_detailsFrom.GongCopyBasicFields(staff_detailsTo)

	//insertion point for the staging of instances referenced by pointers
	if staff_detailsFrom.Staff_size != nil {
		staff_detailsTo.Staff_size = GongCopyBranchStaff_size(mapOrigCopy, staff_detailsFrom.Staff_size)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _line_detail := range staff_detailsFrom.Line_detail {
		staff_detailsTo.Line_detail = append(staff_detailsTo.Line_detail, GongCopyBranchLine_detail(mapOrigCopy, _line_detail))
	}
	for _, _staff_tuning := range staff_detailsFrom.Staff_tuning {
		staff_detailsTo.Staff_tuning = append(staff_detailsTo.Staff_tuning, GongCopyBranchStaff_tuning(mapOrigCopy, _staff_tuning))
	}

	return
}

func GongCopyBranchStaff_divide(mapOrigCopy map[any]any, staff_divideFrom *Staff_divide) (staff_divideTo *Staff_divide) {
	var alreadyCopied bool
	staff_divideTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, staff_divideFrom)
	if alreadyCopied {
		return
	}
	staff_divideFrom.GongCopyBasicFields(staff_divideTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStaff_layout(mapOrigCopy map[any]any, staff_layoutFrom *Staff_layout) (staff_layoutTo *Staff_layout) {
	var alreadyCopied bool
	staff_layoutTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, staff_layoutFrom)
	if alreadyCopied {
		return
	}
	staff_layoutFrom.GongCopyBasicFields(staff_layoutTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStaff_size(mapOrigCopy map[any]any, staff_sizeFrom *Staff_size) (staff_sizeTo *Staff_size) {
	var alreadyCopied bool
	staff_sizeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, staff_sizeFrom)
	if alreadyCopied {
		return
	}
	staff_sizeFrom.GongCopyBasicFields(staff_sizeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStaff_tuning(mapOrigCopy map[any]any, staff_tuningFrom *Staff_tuning) (staff_tuningTo *Staff_tuning) {
	var alreadyCopied bool
	staff_tuningTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, staff_tuningFrom)
	if alreadyCopied {
		return
	}
	staff_tuningFrom.GongCopyBasicFields(staff_tuningTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStem(mapOrigCopy map[any]any, stemFrom *Stem) (stemTo *Stem) {
	var alreadyCopied bool
	stemTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stemFrom)
	if alreadyCopied {
		return
	}
	stemFrom.GongCopyBasicFields(stemTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStick(mapOrigCopy map[any]any, stickFrom *Stick) (stickTo *Stick) {
	var alreadyCopied bool
	stickTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stickFrom)
	if alreadyCopied {
		return
	}
	stickFrom.GongCopyBasicFields(stickTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchString_mute(mapOrigCopy map[any]any, string_muteFrom *String_mute) (string_muteTo *String_mute) {
	var alreadyCopied bool
	string_muteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, string_muteFrom)
	if alreadyCopied {
		return
	}
	string_muteFrom.GongCopyBasicFields(string_muteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchString_type(mapOrigCopy map[any]any, string_typeFrom *String_type) (string_typeTo *String_type) {
	var alreadyCopied bool
	string_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, string_typeFrom)
	if alreadyCopied {
		return
	}
	string_typeFrom.GongCopyBasicFields(string_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStrong_accent(mapOrigCopy map[any]any, strong_accentFrom *Strong_accent) (strong_accentTo *Strong_accent) {
	var alreadyCopied bool
	strong_accentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, strong_accentFrom)
	if alreadyCopied {
		return
	}
	strong_accentFrom.GongCopyBasicFields(strong_accentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStyle_text(mapOrigCopy map[any]any, style_textFrom *Style_text) (style_textTo *Style_text) {
	var alreadyCopied bool
	style_textTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, style_textFrom)
	if alreadyCopied {
		return
	}
	style_textFrom.GongCopyBasicFields(style_textTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSupports(mapOrigCopy map[any]any, supportsFrom *Supports) (supportsTo *Supports) {
	var alreadyCopied bool
	supportsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, supportsFrom)
	if alreadyCopied {
		return
	}
	supportsFrom.GongCopyBasicFields(supportsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSwing(mapOrigCopy map[any]any, swingFrom *Swing) (swingTo *Swing) {
	var alreadyCopied bool
	swingTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, swingFrom)
	if alreadyCopied {
		return
	}
	swingFrom.GongCopyBasicFields(swingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSync(mapOrigCopy map[any]any, syncFrom *Sync) (syncTo *Sync) {
	var alreadyCopied bool
	syncTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, syncFrom)
	if alreadyCopied {
		return
	}
	syncFrom.GongCopyBasicFields(syncTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSystem_dividers(mapOrigCopy map[any]any, system_dividersFrom *System_dividers) (system_dividersTo *System_dividers) {
	var alreadyCopied bool
	system_dividersTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, system_dividersFrom)
	if alreadyCopied {
		return
	}
	system_dividersFrom.GongCopyBasicFields(system_dividersTo)

	//insertion point for the staging of instances referenced by pointers
	if system_dividersFrom.Left_divider != nil {
		system_dividersTo.Left_divider = GongCopyBranchEmpty_print_object_style_align(mapOrigCopy, system_dividersFrom.Left_divider)
	}
	if system_dividersFrom.Right_divider != nil {
		system_dividersTo.Right_divider = GongCopyBranchEmpty_print_object_style_align(mapOrigCopy, system_dividersFrom.Right_divider)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSystem_layout(mapOrigCopy map[any]any, system_layoutFrom *System_layout) (system_layoutTo *System_layout) {
	var alreadyCopied bool
	system_layoutTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, system_layoutFrom)
	if alreadyCopied {
		return
	}
	system_layoutFrom.GongCopyBasicFields(system_layoutTo)

	//insertion point for the staging of instances referenced by pointers
	if system_layoutFrom.System_margins != nil {
		system_layoutTo.System_margins = GongCopyBranchSystem_margins(mapOrigCopy, system_layoutFrom.System_margins)
	}
	if system_layoutFrom.System_dividers != nil {
		system_layoutTo.System_dividers = GongCopyBranchSystem_dividers(mapOrigCopy, system_layoutFrom.System_dividers)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSystem_margins(mapOrigCopy map[any]any, system_marginsFrom *System_margins) (system_marginsTo *System_margins) {
	var alreadyCopied bool
	system_marginsTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, system_marginsFrom)
	if alreadyCopied {
		return
	}
	system_marginsFrom.GongCopyBasicFields(system_marginsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTap(mapOrigCopy map[any]any, tapFrom *Tap) (tapTo *Tap) {
	var alreadyCopied bool
	tapTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tapFrom)
	if alreadyCopied {
		return
	}
	tapFrom.GongCopyBasicFields(tapTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTechnical(mapOrigCopy map[any]any, technicalFrom *Technical) (technicalTo *Technical) {
	var alreadyCopied bool
	technicalTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, technicalFrom)
	if alreadyCopied {
		return
	}
	technicalFrom.GongCopyBasicFields(technicalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty_placement := range technicalFrom.Up_bow {
		technicalTo.Up_bow = append(technicalTo.Up_bow, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_placement := range technicalFrom.Down_bow {
		technicalTo.Down_bow = append(technicalTo.Down_bow, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _harmonic := range technicalFrom.Harmonic {
		technicalTo.Harmonic = append(technicalTo.Harmonic, GongCopyBranchHarmonic(mapOrigCopy, _harmonic))
	}
	for _, _empty_placement := range technicalFrom.Open_string {
		technicalTo.Open_string = append(technicalTo.Open_string, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_placement := range technicalFrom.Thumb_position {
		technicalTo.Thumb_position = append(technicalTo.Thumb_position, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _fingering := range technicalFrom.Fingering {
		technicalTo.Fingering = append(technicalTo.Fingering, GongCopyBranchFingering(mapOrigCopy, _fingering))
	}
	for _, _placement_text := range technicalFrom.Pluck {
		technicalTo.Pluck = append(technicalTo.Pluck, GongCopyBranchPlacement_text(mapOrigCopy, _placement_text))
	}
	for _, _empty_placement := range technicalFrom.Double_tongue {
		technicalTo.Double_tongue = append(technicalTo.Double_tongue, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_placement := range technicalFrom.Triple_tongue {
		technicalTo.Triple_tongue = append(technicalTo.Triple_tongue, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_placement_smufl := range technicalFrom.Stopped {
		technicalTo.Stopped = append(technicalTo.Stopped, GongCopyBranchEmpty_placement_smufl(mapOrigCopy, _empty_placement_smufl))
	}
	for _, _empty_placement := range technicalFrom.Snap_pizzicato {
		technicalTo.Snap_pizzicato = append(technicalTo.Snap_pizzicato, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _fret := range technicalFrom.Fret {
		technicalTo.Fret = append(technicalTo.Fret, GongCopyBranchFret(mapOrigCopy, _fret))
	}
	for _, _string_type := range technicalFrom.String {
		technicalTo.String = append(technicalTo.String, GongCopyBranchString_type(mapOrigCopy, _string_type))
	}
	for _, _hammer_on_pull_off := range technicalFrom.Hammer_on {
		technicalTo.Hammer_on = append(technicalTo.Hammer_on, GongCopyBranchHammer_on_pull_off(mapOrigCopy, _hammer_on_pull_off))
	}
	for _, _hammer_on_pull_off := range technicalFrom.Pull_off {
		technicalTo.Pull_off = append(technicalTo.Pull_off, GongCopyBranchHammer_on_pull_off(mapOrigCopy, _hammer_on_pull_off))
	}
	for _, _bend := range technicalFrom.Bend {
		technicalTo.Bend = append(technicalTo.Bend, GongCopyBranchBend(mapOrigCopy, _bend))
	}
	for _, _tap := range technicalFrom.Tap {
		technicalTo.Tap = append(technicalTo.Tap, GongCopyBranchTap(mapOrigCopy, _tap))
	}
	for _, _heel_toe := range technicalFrom.Heel {
		technicalTo.Heel = append(technicalTo.Heel, GongCopyBranchHeel_toe(mapOrigCopy, _heel_toe))
	}
	for _, _heel_toe := range technicalFrom.Toe {
		technicalTo.Toe = append(technicalTo.Toe, GongCopyBranchHeel_toe(mapOrigCopy, _heel_toe))
	}
	for _, _empty_placement := range technicalFrom.Fingernails {
		technicalTo.Fingernails = append(technicalTo.Fingernails, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _hole := range technicalFrom.Hole {
		technicalTo.Hole = append(technicalTo.Hole, GongCopyBranchHole(mapOrigCopy, _hole))
	}
	for _, _arrow := range technicalFrom.Arrow {
		technicalTo.Arrow = append(technicalTo.Arrow, GongCopyBranchArrow(mapOrigCopy, _arrow))
	}
	for _, _handbell := range technicalFrom.Handbell {
		technicalTo.Handbell = append(technicalTo.Handbell, GongCopyBranchHandbell(mapOrigCopy, _handbell))
	}
	for _, _empty_placement := range technicalFrom.Brass_bend {
		technicalTo.Brass_bend = append(technicalTo.Brass_bend, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_placement := range technicalFrom.Flip {
		technicalTo.Flip = append(technicalTo.Flip, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_placement := range technicalFrom.Smear {
		technicalTo.Smear = append(technicalTo.Smear, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _empty_placement_smufl := range technicalFrom.Open {
		technicalTo.Open = append(technicalTo.Open, GongCopyBranchEmpty_placement_smufl(mapOrigCopy, _empty_placement_smufl))
	}
	for _, _empty_placement_smufl := range technicalFrom.Half_muted {
		technicalTo.Half_muted = append(technicalTo.Half_muted, GongCopyBranchEmpty_placement_smufl(mapOrigCopy, _empty_placement_smufl))
	}
	for _, _harmon_mute := range technicalFrom.Harmon_mute {
		technicalTo.Harmon_mute = append(technicalTo.Harmon_mute, GongCopyBranchHarmon_mute(mapOrigCopy, _harmon_mute))
	}
	for _, _empty_placement := range technicalFrom.Golpe {
		technicalTo.Golpe = append(technicalTo.Golpe, GongCopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _other_placement_text := range technicalFrom.Other_technical {
		technicalTo.Other_technical = append(technicalTo.Other_technical, GongCopyBranchOther_placement_text(mapOrigCopy, _other_placement_text))
	}

	return
}

func GongCopyBranchText_element_data(mapOrigCopy map[any]any, text_element_dataFrom *Text_element_data) (text_element_dataTo *Text_element_data) {
	var alreadyCopied bool
	text_element_dataTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, text_element_dataFrom)
	if alreadyCopied {
		return
	}
	text_element_dataFrom.GongCopyBasicFields(text_element_dataTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTie(mapOrigCopy map[any]any, tieFrom *Tie) (tieTo *Tie) {
	var alreadyCopied bool
	tieTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tieFrom)
	if alreadyCopied {
		return
	}
	tieFrom.GongCopyBasicFields(tieTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTied(mapOrigCopy map[any]any, tiedFrom *Tied) (tiedTo *Tied) {
	var alreadyCopied bool
	tiedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tiedFrom)
	if alreadyCopied {
		return
	}
	tiedFrom.GongCopyBasicFields(tiedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTime(mapOrigCopy map[any]any, timeFrom *Time) (timeTo *Time) {
	var alreadyCopied bool
	timeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timeFrom)
	if alreadyCopied {
		return
	}
	timeFrom.GongCopyBasicFields(timeTo)

	//insertion point for the staging of instances referenced by pointers
	if timeFrom.Interchangeable != nil {
		timeTo.Interchangeable = GongCopyBranchInterchangeable(mapOrigCopy, timeFrom.Interchangeable)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTime_modification(mapOrigCopy map[any]any, time_modificationFrom *Time_modification) (time_modificationTo *Time_modification) {
	var alreadyCopied bool
	time_modificationTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, time_modificationFrom)
	if alreadyCopied {
		return
	}
	time_modificationFrom.GongCopyBasicFields(time_modificationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTimpani(mapOrigCopy map[any]any, timpaniFrom *Timpani) (timpaniTo *Timpani) {
	var alreadyCopied bool
	timpaniTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timpaniFrom)
	if alreadyCopied {
		return
	}
	timpaniFrom.GongCopyBasicFields(timpaniTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTranspose(mapOrigCopy map[any]any, transposeFrom *Transpose) (transposeTo *Transpose) {
	var alreadyCopied bool
	transposeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, transposeFrom)
	if alreadyCopied {
		return
	}
	transposeFrom.GongCopyBasicFields(transposeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTremolo(mapOrigCopy map[any]any, tremoloFrom *Tremolo) (tremoloTo *Tremolo) {
	var alreadyCopied bool
	tremoloTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tremoloFrom)
	if alreadyCopied {
		return
	}
	tremoloFrom.GongCopyBasicFields(tremoloTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTuplet(mapOrigCopy map[any]any, tupletFrom *Tuplet) (tupletTo *Tuplet) {
	var alreadyCopied bool
	tupletTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tupletFrom)
	if alreadyCopied {
		return
	}
	tupletFrom.GongCopyBasicFields(tupletTo)

	//insertion point for the staging of instances referenced by pointers
	if tupletFrom.Tuplet_actual != nil {
		tupletTo.Tuplet_actual = GongCopyBranchTuplet_portion(mapOrigCopy, tupletFrom.Tuplet_actual)
	}
	if tupletFrom.Tuplet_normal != nil {
		tupletTo.Tuplet_normal = GongCopyBranchTuplet_portion(mapOrigCopy, tupletFrom.Tuplet_normal)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTuplet_dot(mapOrigCopy map[any]any, tuplet_dotFrom *Tuplet_dot) (tuplet_dotTo *Tuplet_dot) {
	var alreadyCopied bool
	tuplet_dotTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tuplet_dotFrom)
	if alreadyCopied {
		return
	}
	tuplet_dotFrom.GongCopyBasicFields(tuplet_dotTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTuplet_number(mapOrigCopy map[any]any, tuplet_numberFrom *Tuplet_number) (tuplet_numberTo *Tuplet_number) {
	var alreadyCopied bool
	tuplet_numberTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tuplet_numberFrom)
	if alreadyCopied {
		return
	}
	tuplet_numberFrom.GongCopyBasicFields(tuplet_numberTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTuplet_portion(mapOrigCopy map[any]any, tuplet_portionFrom *Tuplet_portion) (tuplet_portionTo *Tuplet_portion) {
	var alreadyCopied bool
	tuplet_portionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tuplet_portionFrom)
	if alreadyCopied {
		return
	}
	tuplet_portionFrom.GongCopyBasicFields(tuplet_portionTo)

	//insertion point for the staging of instances referenced by pointers
	if tuplet_portionFrom.Tuplet_number != nil {
		tuplet_portionTo.Tuplet_number = GongCopyBranchTuplet_number(mapOrigCopy, tuplet_portionFrom.Tuplet_number)
	}
	if tuplet_portionFrom.Tuplet_type != nil {
		tuplet_portionTo.Tuplet_type = GongCopyBranchTuplet_type(mapOrigCopy, tuplet_portionFrom.Tuplet_type)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tuplet_dot := range tuplet_portionFrom.Tuplet_dot {
		tuplet_portionTo.Tuplet_dot = append(tuplet_portionTo.Tuplet_dot, GongCopyBranchTuplet_dot(mapOrigCopy, _tuplet_dot))
	}

	return
}

func GongCopyBranchTuplet_type(mapOrigCopy map[any]any, tuplet_typeFrom *Tuplet_type) (tuplet_typeTo *Tuplet_type) {
	var alreadyCopied bool
	tuplet_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tuplet_typeFrom)
	if alreadyCopied {
		return
	}
	tuplet_typeFrom.GongCopyBasicFields(tuplet_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTyped_text(mapOrigCopy map[any]any, typed_textFrom *Typed_text) (typed_textTo *Typed_text) {
	var alreadyCopied bool
	typed_textTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, typed_textFrom)
	if alreadyCopied {
		return
	}
	typed_textFrom.GongCopyBasicFields(typed_textTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchUnpitched(mapOrigCopy map[any]any, unpitchedFrom *Unpitched) (unpitchedTo *Unpitched) {
	var alreadyCopied bool
	unpitchedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, unpitchedFrom)
	if alreadyCopied {
		return
	}
	unpitchedFrom.GongCopyBasicFields(unpitchedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVirtual_instrument(mapOrigCopy map[any]any, virtual_instrumentFrom *Virtual_instrument) (virtual_instrumentTo *Virtual_instrument) {
	var alreadyCopied bool
	virtual_instrumentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, virtual_instrumentFrom)
	if alreadyCopied {
		return
	}
	virtual_instrumentFrom.GongCopyBasicFields(virtual_instrumentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchWait(mapOrigCopy map[any]any, waitFrom *Wait) (waitTo *Wait) {
	var alreadyCopied bool
	waitTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, waitFrom)
	if alreadyCopied {
		return
	}
	waitFrom.GongCopyBasicFields(waitTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchWavy_line(mapOrigCopy map[any]any, wavy_lineFrom *Wavy_line) (wavy_lineTo *Wavy_line) {
	var alreadyCopied bool
	wavy_lineTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, wavy_lineFrom)
	if alreadyCopied {
		return
	}
	wavy_lineFrom.GongCopyBasicFields(wavy_lineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchWedge(mapOrigCopy map[any]any, wedgeFrom *Wedge) (wedgeTo *Wedge) {
	var alreadyCopied bool
	wedgeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, wedgeFrom)
	if alreadyCopied {
		return
	}
	wedgeFrom.GongCopyBasicFields(wedgeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchWood(mapOrigCopy map[any]any, woodFrom *Wood) (woodTo *Wood) {
	var alreadyCopied bool
	woodTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, woodFrom)
	if alreadyCopied {
		return
	}
	woodFrom.GongCopyBasicFields(woodTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchWork(mapOrigCopy map[any]any, workFrom *Work) (workTo *Work) {
	var alreadyCopied bool
	workTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, workFrom)
	if alreadyCopied {
		return
	}
	workFrom.GongCopyBasicFields(workTo)

	//insertion point for the staging of instances referenced by pointers
	if workFrom.Opus != nil {
		workTo.Opus = GongCopyBranchOpus(mapOrigCopy, workFrom.Opus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (a_directive *A_directive) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_directive) {
		return
	}

	a_directive.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_measure *A_measure) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_measure) {
		return
	}

	a_measure.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _note := range a_measure.Note {
		stage.UnstageBranch(_note)
	}
	for _, _backup := range a_measure.Backup {
		stage.UnstageBranch(_backup)
	}
	for _, _forward := range a_measure.Forward {
		stage.UnstageBranch(_forward)
	}
	for _, _direction := range a_measure.Direction {
		stage.UnstageBranch(_direction)
	}
	for _, _attributes := range a_measure.Attributes {
		stage.UnstageBranch(_attributes)
	}
	for _, _harmony := range a_measure.Harmony {
		stage.UnstageBranch(_harmony)
	}
	for _, _figured_bass := range a_measure.Figured_bass {
		stage.UnstageBranch(_figured_bass)
	}
	for _, _print := range a_measure.Print {
		stage.UnstageBranch(_print)
	}
	for _, _sound := range a_measure.Sound {
		stage.UnstageBranch(_sound)
	}
	for _, _listening := range a_measure.Listening {
		stage.UnstageBranch(_listening)
	}
	for _, _barline := range a_measure.Barline {
		stage.UnstageBranch(_barline)
	}
	for _, _grouping := range a_measure.Grouping {
		stage.UnstageBranch(_grouping)
	}
	for _, _link := range a_measure.Link {
		stage.UnstageBranch(_link)
	}
	for _, _bookmark := range a_measure.Bookmark {
		stage.UnstageBranch(_bookmark)
	}

}

func (a_measure_1 *A_measure_1) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_measure_1) {
		return
	}

	a_measure_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_part_1 := range a_measure_1.Part {
		stage.UnstageBranch(_a_part_1)
	}

}

func (a_part *A_part) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_part) {
		return
	}

	a_part.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _a_measure := range a_part.Measure {
		stage.UnstageBranch(_a_measure)
	}

}

func (a_part_1 *A_part_1) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(a_part_1) {
		return
	}

	a_part_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _note := range a_part_1.Note {
		stage.UnstageBranch(_note)
	}
	for _, _backup := range a_part_1.Backup {
		stage.UnstageBranch(_backup)
	}
	for _, _forward := range a_part_1.Forward {
		stage.UnstageBranch(_forward)
	}
	for _, _direction := range a_part_1.Direction {
		stage.UnstageBranch(_direction)
	}
	for _, _attributes := range a_part_1.Attributes {
		stage.UnstageBranch(_attributes)
	}
	for _, _harmony := range a_part_1.Harmony {
		stage.UnstageBranch(_harmony)
	}
	for _, _figured_bass := range a_part_1.Figured_bass {
		stage.UnstageBranch(_figured_bass)
	}
	for _, _print := range a_part_1.Print {
		stage.UnstageBranch(_print)
	}
	for _, _sound := range a_part_1.Sound {
		stage.UnstageBranch(_sound)
	}
	for _, _listening := range a_part_1.Listening {
		stage.UnstageBranch(_listening)
	}
	for _, _barline := range a_part_1.Barline {
		stage.UnstageBranch(_barline)
	}
	for _, _grouping := range a_part_1.Grouping {
		stage.UnstageBranch(_grouping)
	}
	for _, _link := range a_part_1.Link {
		stage.UnstageBranch(_link)
	}
	for _, _bookmark := range a_part_1.Bookmark {
		stage.UnstageBranch(_bookmark)
	}

}

func (accidental *Accidental) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(accidental) {
		return
	}

	accidental.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (accidental_mark *Accidental_mark) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(accidental_mark) {
		return
	}

	accidental_mark.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (accidental_text *Accidental_text) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(accidental_text) {
		return
	}

	accidental_text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (accord *Accord) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(accord) {
		return
	}

	accord.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (accordion_registration *Accordion_registration) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(accordion_registration) {
		return
	}

	accordion_registration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (appearance *Appearance) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(appearance) {
		return
	}

	appearance.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _line_width := range appearance.Line_width {
		stage.UnstageBranch(_line_width)
	}
	for _, _note_size := range appearance.Note_size {
		stage.UnstageBranch(_note_size)
	}
	for _, _distance := range appearance.Distance {
		stage.UnstageBranch(_distance)
	}
	for _, _glyph := range appearance.Glyph {
		stage.UnstageBranch(_glyph)
	}
	for _, _other_appearance := range appearance.Other_appearance {
		stage.UnstageBranch(_other_appearance)
	}

}

func (arpeggiate *Arpeggiate) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(arpeggiate) {
		return
	}

	arpeggiate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (arrow *Arrow) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(arrow) {
		return
	}

	arrow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (articulations *Articulations) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(articulations) {
		return
	}

	articulations.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty_placement := range articulations.Accent {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _strong_accent := range articulations.Strong_accent {
		stage.UnstageBranch(_strong_accent)
	}
	for _, _empty_placement := range articulations.Staccato {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_placement := range articulations.Tenuto {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_placement := range articulations.Detached_legato {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_placement := range articulations.Staccatissimo {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_placement := range articulations.Spiccato {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_line := range articulations.Scoop {
		stage.UnstageBranch(_empty_line)
	}
	for _, _empty_line := range articulations.Plop {
		stage.UnstageBranch(_empty_line)
	}
	for _, _empty_line := range articulations.Doit {
		stage.UnstageBranch(_empty_line)
	}
	for _, _empty_line := range articulations.Falloff {
		stage.UnstageBranch(_empty_line)
	}
	for _, _breath_mark := range articulations.Breath_mark {
		stage.UnstageBranch(_breath_mark)
	}
	for _, _caesura := range articulations.Caesura {
		stage.UnstageBranch(_caesura)
	}
	for _, _empty_placement := range articulations.Stress {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_placement := range articulations.Unstress {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_placement := range articulations.Soft_accent {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _other_placement_text := range articulations.Other_articulation {
		stage.UnstageBranch(_other_placement_text)
	}

}

func (assess *Assess) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(assess) {
		return
	}

	assess.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attributes *Attributes) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(attributes) {
		return
	}

	attributes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributes.Footnote != nil {
		stage.UnstageBranch(attributes.Footnote)
	}
	if attributes.Level != nil {
		stage.UnstageBranch(attributes.Level)
	}
	if attributes.Part_symbol != nil {
		stage.UnstageBranch(attributes.Part_symbol)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key := range attributes.Key {
		stage.UnstageBranch(_key)
	}
	for _, _time := range attributes.Time {
		stage.UnstageBranch(_time)
	}
	for _, _clef := range attributes.Clef {
		stage.UnstageBranch(_clef)
	}
	for _, _staff_details := range attributes.Staff_details {
		stage.UnstageBranch(_staff_details)
	}
	for _, _transpose := range attributes.Transpose {
		stage.UnstageBranch(_transpose)
	}
	for _, _for_part := range attributes.For_part {
		stage.UnstageBranch(_for_part)
	}
	for _, _a_directive := range attributes.Directive {
		stage.UnstageBranch(_a_directive)
	}
	for _, _measure_style := range attributes.Measure_style {
		stage.UnstageBranch(_measure_style)
	}

}

func (backup *Backup) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(backup) {
		return
	}

	backup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if backup.Footnote != nil {
		stage.UnstageBranch(backup.Footnote)
	}
	if backup.Level != nil {
		stage.UnstageBranch(backup.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bar_style_color *Bar_style_color) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bar_style_color) {
		return
	}

	bar_style_color.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (barline *Barline) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(barline) {
		return
	}

	barline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if barline.Bar_style != nil {
		stage.UnstageBranch(barline.Bar_style)
	}
	if barline.Footnote != nil {
		stage.UnstageBranch(barline.Footnote)
	}
	if barline.Level != nil {
		stage.UnstageBranch(barline.Level)
	}
	if barline.Wavy_line != nil {
		stage.UnstageBranch(barline.Wavy_line)
	}
	if barline.Segno_1 != nil {
		stage.UnstageBranch(barline.Segno_1)
	}
	if barline.Coda_1 != nil {
		stage.UnstageBranch(barline.Coda_1)
	}
	if barline.Fermata != nil {
		stage.UnstageBranch(barline.Fermata)
	}
	if barline.Ending != nil {
		stage.UnstageBranch(barline.Ending)
	}
	if barline.Repeat != nil {
		stage.UnstageBranch(barline.Repeat)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (barre *Barre) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(barre) {
		return
	}

	barre.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bass *Bass) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bass) {
		return
	}

	bass.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bass.Bass_separator != nil {
		stage.UnstageBranch(bass.Bass_separator)
	}
	if bass.Bass_step != nil {
		stage.UnstageBranch(bass.Bass_step)
	}
	if bass.Bass_alter != nil {
		stage.UnstageBranch(bass.Bass_alter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bass_step *Bass_step) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bass_step) {
		return
	}

	bass_step.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beam *Beam) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(beam) {
		return
	}

	beam.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beat_repeat *Beat_repeat) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(beat_repeat) {
		return
	}

	beat_repeat.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beat_unit_tied *Beat_unit_tied) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(beat_unit_tied) {
		return
	}

	beat_unit_tied.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (beater *Beater) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(beater) {
		return
	}

	beater.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bend *Bend) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bend) {
		return
	}

	bend.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bend.Release != nil {
		stage.UnstageBranch(bend.Release)
	}
	if bend.With_bar != nil {
		stage.UnstageBranch(bend.With_bar)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bookmark *Bookmark) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bookmark) {
		return
	}

	bookmark.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bracket *Bracket) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bracket) {
		return
	}

	bracket.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (breath_mark *Breath_mark) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(breath_mark) {
		return
	}

	breath_mark.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (caesura *Caesura) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(caesura) {
		return
	}

	caesura.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cancel *Cancel) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(cancel) {
		return
	}

	cancel.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (clef *Clef) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(clef) {
		return
	}

	clef.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coda *Coda) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coda) {
		return
	}

	coda.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (credit *Credit) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(credit) {
		return
	}

	credit.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if credit.Credit_image != nil {
		stage.UnstageBranch(credit.Credit_image)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range credit.Link {
		stage.UnstageBranch(_link)
	}
	for _, _bookmark := range credit.Bookmark {
		stage.UnstageBranch(_bookmark)
	}
	for _, _formatted_text_id := range credit.Credit_words {
		stage.UnstageBranch(_formatted_text_id)
	}
	for _, _formatted_symbol_id := range credit.Credit_symbol {
		stage.UnstageBranch(_formatted_symbol_id)
	}

}

func (dashes *Dashes) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(dashes) {
		return
	}

	dashes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (defaults *Defaults) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(defaults) {
		return
	}

	defaults.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if defaults.Scaling != nil {
		stage.UnstageBranch(defaults.Scaling)
	}
	if defaults.Page_layout != nil {
		stage.UnstageBranch(defaults.Page_layout)
	}
	if defaults.System_layout != nil {
		stage.UnstageBranch(defaults.System_layout)
	}
	if defaults.Appearance != nil {
		stage.UnstageBranch(defaults.Appearance)
	}
	if defaults.Music_font != nil {
		stage.UnstageBranch(defaults.Music_font)
	}
	if defaults.Word_font != nil {
		stage.UnstageBranch(defaults.Word_font)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _staff_layout := range defaults.Staff_layout {
		stage.UnstageBranch(_staff_layout)
	}
	for _, _lyric_font := range defaults.Lyric_font {
		stage.UnstageBranch(_lyric_font)
	}
	for _, _lyric_language := range defaults.Lyric_language {
		stage.UnstageBranch(_lyric_language)
	}

}

func (degree *Degree) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(degree) {
		return
	}

	degree.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if degree.Degree_value != nil {
		stage.UnstageBranch(degree.Degree_value)
	}
	if degree.Degree_alter != nil {
		stage.UnstageBranch(degree.Degree_alter)
	}
	if degree.Degree_type != nil {
		stage.UnstageBranch(degree.Degree_type)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degree_alter *Degree_alter) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(degree_alter) {
		return
	}

	degree_alter.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degree_type *Degree_type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(degree_type) {
		return
	}

	degree_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (degree_value *Degree_value) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(degree_value) {
		return
	}

	degree_value.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (direction *Direction) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(direction) {
		return
	}

	direction.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if direction.Offset != nil {
		stage.UnstageBranch(direction.Offset)
	}
	if direction.Footnote != nil {
		stage.UnstageBranch(direction.Footnote)
	}
	if direction.Level != nil {
		stage.UnstageBranch(direction.Level)
	}
	if direction.Sound != nil {
		stage.UnstageBranch(direction.Sound)
	}
	if direction.Listening != nil {
		stage.UnstageBranch(direction.Listening)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _direction_type := range direction.Direction_type {
		stage.UnstageBranch(_direction_type)
	}

}

func (direction_type *Direction_type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(direction_type) {
		return
	}

	direction_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if direction_type.Wedge != nil {
		stage.UnstageBranch(direction_type.Wedge)
	}
	if direction_type.Dashes != nil {
		stage.UnstageBranch(direction_type.Dashes)
	}
	if direction_type.Bracket != nil {
		stage.UnstageBranch(direction_type.Bracket)
	}
	if direction_type.Pedal != nil {
		stage.UnstageBranch(direction_type.Pedal)
	}
	if direction_type.Metronome != nil {
		stage.UnstageBranch(direction_type.Metronome)
	}
	if direction_type.Octave_shift != nil {
		stage.UnstageBranch(direction_type.Octave_shift)
	}
	if direction_type.Harp_pedals != nil {
		stage.UnstageBranch(direction_type.Harp_pedals)
	}
	if direction_type.Damp != nil {
		stage.UnstageBranch(direction_type.Damp)
	}
	if direction_type.Damp_all != nil {
		stage.UnstageBranch(direction_type.Damp_all)
	}
	if direction_type.Eyeglasses != nil {
		stage.UnstageBranch(direction_type.Eyeglasses)
	}
	if direction_type.String_mute != nil {
		stage.UnstageBranch(direction_type.String_mute)
	}
	if direction_type.Scordatura != nil {
		stage.UnstageBranch(direction_type.Scordatura)
	}
	if direction_type.Image != nil {
		stage.UnstageBranch(direction_type.Image)
	}
	if direction_type.Principal_voice != nil {
		stage.UnstageBranch(direction_type.Principal_voice)
	}
	if direction_type.Accordion_registration != nil {
		stage.UnstageBranch(direction_type.Accordion_registration)
	}
	if direction_type.Staff_divide != nil {
		stage.UnstageBranch(direction_type.Staff_divide)
	}
	if direction_type.Other_direction != nil {
		stage.UnstageBranch(direction_type.Other_direction)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formatted_text_id := range direction_type.Rehearsal {
		stage.UnstageBranch(_formatted_text_id)
	}
	for _, _segno := range direction_type.Segno {
		stage.UnstageBranch(_segno)
	}
	for _, _coda := range direction_type.Coda {
		stage.UnstageBranch(_coda)
	}
	for _, _formatted_text_id := range direction_type.Words {
		stage.UnstageBranch(_formatted_text_id)
	}
	for _, _formatted_symbol_id := range direction_type.Symbol {
		stage.UnstageBranch(_formatted_symbol_id)
	}
	for _, _dynamics := range direction_type.Dynamics {
		stage.UnstageBranch(_dynamics)
	}
	for _, _percussion := range direction_type.Percussion {
		stage.UnstageBranch(_percussion)
	}

}

func (distance *Distance) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(distance) {
		return
	}

	distance.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (double *Double) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(double) {
		return
	}

	double.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dynamics *Dynamics) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(dynamics) {
		return
	}

	dynamics.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _other_text := range dynamics.Other_dynamics {
		stage.UnstageBranch(_other_text)
	}

}

func (effect *Effect) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(effect) {
		return
	}

	effect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (elision *Elision) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(elision) {
		return
	}

	elision.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty *Empty) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(empty) {
		return
	}

	empty.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_font *Empty_font) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(empty_font) {
		return
	}

	empty_font.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_line *Empty_line) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(empty_line) {
		return
	}

	empty_line.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_placement *Empty_placement) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(empty_placement) {
		return
	}

	empty_placement.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_placement_smufl *Empty_placement_smufl) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(empty_placement_smufl) {
		return
	}

	empty_placement_smufl.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_print_object_style_align *Empty_print_object_style_align) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(empty_print_object_style_align) {
		return
	}

	empty_print_object_style_align.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_print_style *Empty_print_style) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(empty_print_style) {
		return
	}

	empty_print_style.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_print_style_align *Empty_print_style_align) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(empty_print_style_align) {
		return
	}

	empty_print_style_align.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_print_style_align_id *Empty_print_style_align_id) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(empty_print_style_align_id) {
		return
	}

	empty_print_style_align_id.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (empty_trill_sound *Empty_trill_sound) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(empty_trill_sound) {
		return
	}

	empty_trill_sound.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (encoding *Encoding) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(encoding) {
		return
	}

	encoding.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typed_text := range encoding.Encoder {
		stage.UnstageBranch(_typed_text)
	}
	for _, _supports := range encoding.Supports {
		stage.UnstageBranch(_supports)
	}

}

func (ending *Ending) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ending) {
		return
	}

	ending.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (extend *Extend) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(extend) {
		return
	}

	extend.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (feature *Feature) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(feature) {
		return
	}

	feature.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (fermata *Fermata) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(fermata) {
		return
	}

	fermata.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (figure *Figure) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(figure) {
		return
	}

	figure.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if figure.Prefix != nil {
		stage.UnstageBranch(figure.Prefix)
	}
	if figure.Figure_number != nil {
		stage.UnstageBranch(figure.Figure_number)
	}
	if figure.Suffix != nil {
		stage.UnstageBranch(figure.Suffix)
	}
	if figure.Extend != nil {
		stage.UnstageBranch(figure.Extend)
	}
	if figure.Footnote != nil {
		stage.UnstageBranch(figure.Footnote)
	}
	if figure.Level != nil {
		stage.UnstageBranch(figure.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (figured_bass *Figured_bass) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(figured_bass) {
		return
	}

	figured_bass.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if figured_bass.Footnote != nil {
		stage.UnstageBranch(figured_bass.Footnote)
	}
	if figured_bass.Level != nil {
		stage.UnstageBranch(figured_bass.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _figure := range figured_bass.Figure {
		stage.UnstageBranch(_figure)
	}

}

func (fingering *Fingering) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(fingering) {
		return
	}

	fingering.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (first_fret *First_fret) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(first_fret) {
		return
	}

	first_fret.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (for_part *For_part) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(for_part) {
		return
	}

	for_part.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if for_part.Part_clef != nil {
		stage.UnstageBranch(for_part.Part_clef)
	}
	if for_part.Part_transpose != nil {
		stage.UnstageBranch(for_part.Part_transpose)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formatted_symbol *Formatted_symbol) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(formatted_symbol) {
		return
	}

	formatted_symbol.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formatted_symbol_id *Formatted_symbol_id) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(formatted_symbol_id) {
		return
	}

	formatted_symbol_id.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formatted_text *Formatted_text) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(formatted_text) {
		return
	}

	formatted_text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formatted_text_id *Formatted_text_id) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(formatted_text_id) {
		return
	}

	formatted_text_id.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (forward *Forward) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(forward) {
		return
	}

	forward.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if forward.Footnote != nil {
		stage.UnstageBranch(forward.Footnote)
	}
	if forward.Level != nil {
		stage.UnstageBranch(forward.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (frame *Frame) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(frame) {
		return
	}

	frame.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if frame.First_fret != nil {
		stage.UnstageBranch(frame.First_fret)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _frame_note := range frame.Frame_note {
		stage.UnstageBranch(_frame_note)
	}

}

func (frame_note *Frame_note) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(frame_note) {
		return
	}

	frame_note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if frame_note.String != nil {
		stage.UnstageBranch(frame_note.String)
	}
	if frame_note.Fret != nil {
		stage.UnstageBranch(frame_note.Fret)
	}
	if frame_note.Fingering != nil {
		stage.UnstageBranch(frame_note.Fingering)
	}
	if frame_note.Barre != nil {
		stage.UnstageBranch(frame_note.Barre)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (fret *Fret) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(fret) {
		return
	}

	fret.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (glass *Glass) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(glass) {
		return
	}

	glass.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (glissando *Glissando) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(glissando) {
		return
	}

	glissando.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (glyph *Glyph) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(glyph) {
		return
	}

	glyph.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (grace *Grace) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(grace) {
		return
	}

	grace.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (group_barline *Group_barline) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(group_barline) {
		return
	}

	group_barline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (group_name *Group_name) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(group_name) {
		return
	}

	group_name.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (group_symbol *Group_symbol) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(group_symbol) {
		return
	}

	group_symbol.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (grouping *Grouping) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(grouping) {
		return
	}

	grouping.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _feature := range grouping.Feature {
		stage.UnstageBranch(_feature)
	}

}

func (hammer_on_pull_off *Hammer_on_pull_off) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(hammer_on_pull_off) {
		return
	}

	hammer_on_pull_off.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (handbell *Handbell) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(handbell) {
		return
	}

	handbell.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmon_closed *Harmon_closed) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(harmon_closed) {
		return
	}

	harmon_closed.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmon_mute *Harmon_mute) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(harmon_mute) {
		return
	}

	harmon_mute.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmon_mute.Harmon_closed != nil {
		stage.UnstageBranch(harmon_mute.Harmon_closed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmonic *Harmonic) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(harmonic) {
		return
	}

	harmonic.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harmony *Harmony) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(harmony) {
		return
	}

	harmony.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmony.Root != nil {
		stage.UnstageBranch(harmony.Root)
	}
	if harmony.Numeral != nil {
		stage.UnstageBranch(harmony.Numeral)
	}
	if harmony.Function != nil {
		stage.UnstageBranch(harmony.Function)
	}
	if harmony.Kind != nil {
		stage.UnstageBranch(harmony.Kind)
	}
	if harmony.Inversion != nil {
		stage.UnstageBranch(harmony.Inversion)
	}
	if harmony.Bass != nil {
		stage.UnstageBranch(harmony.Bass)
	}
	if harmony.Frame != nil {
		stage.UnstageBranch(harmony.Frame)
	}
	if harmony.Offset != nil {
		stage.UnstageBranch(harmony.Offset)
	}
	if harmony.Footnote != nil {
		stage.UnstageBranch(harmony.Footnote)
	}
	if harmony.Level != nil {
		stage.UnstageBranch(harmony.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _degree := range harmony.Degree {
		stage.UnstageBranch(_degree)
	}

}

func (harmony_alter *Harmony_alter) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(harmony_alter) {
		return
	}

	harmony_alter.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (harp_pedals *Harp_pedals) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(harp_pedals) {
		return
	}

	harp_pedals.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _pedal_tuning := range harp_pedals.Pedal_tuning {
		stage.UnstageBranch(_pedal_tuning)
	}

}

func (heel_toe *Heel_toe) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(heel_toe) {
		return
	}

	heel_toe.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (hole *Hole) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(hole) {
		return
	}

	hole.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if hole.Hole_closed != nil {
		stage.UnstageBranch(hole.Hole_closed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (hole_closed *Hole_closed) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(hole_closed) {
		return
	}

	hole_closed.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (horizontal_turn *Horizontal_turn) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(horizontal_turn) {
		return
	}

	horizontal_turn.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (identification *Identification) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(identification) {
		return
	}

	identification.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if identification.Encoding != nil {
		stage.UnstageBranch(identification.Encoding)
	}
	if identification.Miscellaneous != nil {
		stage.UnstageBranch(identification.Miscellaneous)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typed_text := range identification.Creator {
		stage.UnstageBranch(_typed_text)
	}
	for _, _typed_text := range identification.Rights {
		stage.UnstageBranch(_typed_text)
	}
	for _, _typed_text := range identification.Relation {
		stage.UnstageBranch(_typed_text)
	}

}

func (image *Image) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(image) {
		return
	}

	image.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (instrument *Instrument) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(instrument) {
		return
	}

	instrument.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (instrument_change *Instrument_change) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(instrument_change) {
		return
	}

	instrument_change.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if instrument_change.Virtual_instrument != nil {
		stage.UnstageBranch(instrument_change.Virtual_instrument)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (instrument_link *Instrument_link) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(instrument_link) {
		return
	}

	instrument_link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (interchangeable *Interchangeable) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(interchangeable) {
		return
	}

	interchangeable.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (inversion *Inversion) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(inversion) {
		return
	}

	inversion.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (key *Key) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(key) {
		return
	}

	key.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if key.Cancel != nil {
		stage.UnstageBranch(key.Cancel)
	}
	if key.Key_accidental != nil {
		stage.UnstageBranch(key.Key_accidental)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key_octave := range key.Key_octave {
		stage.UnstageBranch(_key_octave)
	}

}

func (key_accidental *Key_accidental) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(key_accidental) {
		return
	}

	key_accidental.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (key_octave *Key_octave) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(key_octave) {
		return
	}

	key_octave.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (kind *Kind) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(kind) {
		return
	}

	kind.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (level *Level) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(level) {
		return
	}

	level.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (line_detail *Line_detail) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(line_detail) {
		return
	}

	line_detail.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (line_width *Line_width) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(line_width) {
		return
	}

	line_width.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (link *Link) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(link) {
		return
	}

	link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (listen *Listen) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(listen) {
		return
	}

	listen.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assess := range listen.Assess {
		stage.UnstageBranch(_assess)
	}
	for _, _wait := range listen.Wait {
		stage.UnstageBranch(_wait)
	}
	for _, _other_listening := range listen.Other_listen {
		stage.UnstageBranch(_other_listening)
	}

}

func (listening *Listening) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(listening) {
		return
	}

	listening.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if listening.Offset != nil {
		stage.UnstageBranch(listening.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sync := range listening.Sync {
		stage.UnstageBranch(_sync)
	}
	for _, _other_listening := range listening.Other_listening {
		stage.UnstageBranch(_other_listening)
	}

}

func (lyric *Lyric) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(lyric) {
		return
	}

	lyric.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lyric.Extend != nil {
		stage.UnstageBranch(lyric.Extend)
	}
	if lyric.Footnote != nil {
		stage.UnstageBranch(lyric.Footnote)
	}
	if lyric.Level != nil {
		stage.UnstageBranch(lyric.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _elision := range lyric.Elision {
		stage.UnstageBranch(_elision)
	}
	for _, _text_element_data := range lyric.Text {
		stage.UnstageBranch(_text_element_data)
	}

}

func (lyric_font *Lyric_font) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(lyric_font) {
		return
	}

	lyric_font.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (lyric_language *Lyric_language) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(lyric_language) {
		return
	}

	lyric_language.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measure_layout *Measure_layout) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measure_layout) {
		return
	}

	measure_layout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measure_numbering *Measure_numbering) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measure_numbering) {
		return
	}

	measure_numbering.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measure_repeat *Measure_repeat) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measure_repeat) {
		return
	}

	measure_repeat.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measure_style *Measure_style) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measure_style) {
		return
	}

	measure_style.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if measure_style.Multiple_rest != nil {
		stage.UnstageBranch(measure_style.Multiple_rest)
	}
	if measure_style.Measure_repeat != nil {
		stage.UnstageBranch(measure_style.Measure_repeat)
	}
	if measure_style.Beat_repeat != nil {
		stage.UnstageBranch(measure_style.Beat_repeat)
	}
	if measure_style.Slash != nil {
		stage.UnstageBranch(measure_style.Slash)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (membrane *Membrane) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(membrane) {
		return
	}

	membrane.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metal *Metal) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(metal) {
		return
	}

	metal.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metronome *Metronome) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(metronome) {
		return
	}

	metronome.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if metronome.Per_minute != nil {
		stage.UnstageBranch(metronome.Per_minute)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _beat_unit_tied := range metronome.Beat_unit_tied {
		stage.UnstageBranch(_beat_unit_tied)
	}
	for _, _metronome_note := range metronome.Metronome_note {
		stage.UnstageBranch(_metronome_note)
	}

}

func (metronome_beam *Metronome_beam) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(metronome_beam) {
		return
	}

	metronome_beam.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metronome_note *Metronome_note) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(metronome_note) {
		return
	}

	metronome_note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if metronome_note.Metronome_tied != nil {
		stage.UnstageBranch(metronome_note.Metronome_tied)
	}
	if metronome_note.Metronome_tuplet != nil {
		stage.UnstageBranch(metronome_note.Metronome_tuplet)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _metronome_beam := range metronome_note.Metronome_beam {
		stage.UnstageBranch(_metronome_beam)
	}

}

func (metronome_tied *Metronome_tied) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(metronome_tied) {
		return
	}

	metronome_tied.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metronome_tuplet *Metronome_tuplet) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(metronome_tuplet) {
		return
	}

	metronome_tuplet.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (midi_device *Midi_device) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(midi_device) {
		return
	}

	midi_device.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (midi_instrument *Midi_instrument) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(midi_instrument) {
		return
	}

	midi_instrument.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (miscellaneous *Miscellaneous) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(miscellaneous) {
		return
	}

	miscellaneous.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _miscellaneous_field := range miscellaneous.Miscellaneous_field {
		stage.UnstageBranch(_miscellaneous_field)
	}

}

func (miscellaneous_field *Miscellaneous_field) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(miscellaneous_field) {
		return
	}

	miscellaneous_field.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (mordent *Mordent) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(mordent) {
		return
	}

	mordent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (multiple_rest *Multiple_rest) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(multiple_rest) {
		return
	}

	multiple_rest.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (name_display *Name_display) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(name_display) {
		return
	}

	name_display.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formatted_text := range name_display.Display_text {
		stage.UnstageBranch(_formatted_text)
	}
	for _, _accidental_text := range name_display.Accidental_text {
		stage.UnstageBranch(_accidental_text)
	}

}

func (non_arpeggiate *Non_arpeggiate) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(non_arpeggiate) {
		return
	}

	non_arpeggiate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notations *Notations) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(notations) {
		return
	}

	notations.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notations.Footnote != nil {
		stage.UnstageBranch(notations.Footnote)
	}
	if notations.Level != nil {
		stage.UnstageBranch(notations.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tied := range notations.Tied {
		stage.UnstageBranch(_tied)
	}
	for _, _slur := range notations.Slur {
		stage.UnstageBranch(_slur)
	}
	for _, _tuplet := range notations.Tuplet {
		stage.UnstageBranch(_tuplet)
	}
	for _, _glissando := range notations.Glissando {
		stage.UnstageBranch(_glissando)
	}
	for _, _slide := range notations.Slide {
		stage.UnstageBranch(_slide)
	}
	for _, _ornaments := range notations.Ornaments {
		stage.UnstageBranch(_ornaments)
	}
	for _, _technical := range notations.Technical {
		stage.UnstageBranch(_technical)
	}
	for _, _articulations := range notations.Articulations {
		stage.UnstageBranch(_articulations)
	}
	for _, _dynamics := range notations.Dynamics {
		stage.UnstageBranch(_dynamics)
	}
	for _, _fermata := range notations.Fermata {
		stage.UnstageBranch(_fermata)
	}
	for _, _arpeggiate := range notations.Arpeggiate {
		stage.UnstageBranch(_arpeggiate)
	}
	for _, _non_arpeggiate := range notations.Non_arpeggiate {
		stage.UnstageBranch(_non_arpeggiate)
	}
	for _, _accidental_mark := range notations.Accidental_mark {
		stage.UnstageBranch(_accidental_mark)
	}
	for _, _other_notation := range notations.Other_notation {
		stage.UnstageBranch(_other_notation)
	}

}

func (note *Note) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if note.Grace != nil {
		stage.UnstageBranch(note.Grace)
	}
	if note.Pitch != nil {
		stage.UnstageBranch(note.Pitch)
	}
	if note.Unpitched != nil {
		stage.UnstageBranch(note.Unpitched)
	}
	if note.Rest != nil {
		stage.UnstageBranch(note.Rest)
	}
	if note.Tie != nil {
		stage.UnstageBranch(note.Tie)
	}
	if note.Footnote != nil {
		stage.UnstageBranch(note.Footnote)
	}
	if note.Level != nil {
		stage.UnstageBranch(note.Level)
	}
	if note.Type != nil {
		stage.UnstageBranch(note.Type)
	}
	if note.Accidental != nil {
		stage.UnstageBranch(note.Accidental)
	}
	if note.Time_modification != nil {
		stage.UnstageBranch(note.Time_modification)
	}
	if note.Stem != nil {
		stage.UnstageBranch(note.Stem)
	}
	if note.Notehead != nil {
		stage.UnstageBranch(note.Notehead)
	}
	if note.Notehead_text != nil {
		stage.UnstageBranch(note.Notehead_text)
	}
	if note.Beam != nil {
		stage.UnstageBranch(note.Beam)
	}
	if note.Play != nil {
		stage.UnstageBranch(note.Play)
	}
	if note.Listen != nil {
		stage.UnstageBranch(note.Listen)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument := range note.Instrument {
		stage.UnstageBranch(_instrument)
	}
	for _, _empty_placement := range note.Dot {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _notations := range note.Notations {
		stage.UnstageBranch(_notations)
	}
	for _, _lyric := range note.Lyric {
		stage.UnstageBranch(_lyric)
	}

}

func (note_size *Note_size) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(note_size) {
		return
	}

	note_size.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (note_type *Note_type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(note_type) {
		return
	}

	note_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notehead *Notehead) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(notehead) {
		return
	}

	notehead.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notehead_text *Notehead_text) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(notehead_text) {
		return
	}

	notehead_text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formatted_text := range notehead_text.Display_text {
		stage.UnstageBranch(_formatted_text)
	}
	for _, _accidental_text := range notehead_text.Accidental_text {
		stage.UnstageBranch(_accidental_text)
	}

}

func (numeral *Numeral) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(numeral) {
		return
	}

	numeral.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if numeral.Numeral_root != nil {
		stage.UnstageBranch(numeral.Numeral_root)
	}
	if numeral.Numeral_alter != nil {
		stage.UnstageBranch(numeral.Numeral_alter)
	}
	if numeral.Numeral_key != nil {
		stage.UnstageBranch(numeral.Numeral_key)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (numeral_key *Numeral_key) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(numeral_key) {
		return
	}

	numeral_key.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (numeral_root *Numeral_root) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(numeral_root) {
		return
	}

	numeral_root.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (octave_shift *Octave_shift) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(octave_shift) {
		return
	}

	octave_shift.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (offset *Offset) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(offset) {
		return
	}

	offset.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (opus *Opus) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(opus) {
		return
	}

	opus.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ornaments *Ornaments) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ornaments) {
		return
	}

	ornaments.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty_trill_sound := range ornaments.Trill_mark {
		stage.UnstageBranch(_empty_trill_sound)
	}
	for _, _horizontal_turn := range ornaments.Turn {
		stage.UnstageBranch(_horizontal_turn)
	}
	for _, _horizontal_turn := range ornaments.Delayed_turn {
		stage.UnstageBranch(_horizontal_turn)
	}
	for _, _horizontal_turn := range ornaments.Inverted_turn {
		stage.UnstageBranch(_horizontal_turn)
	}
	for _, _horizontal_turn := range ornaments.Delayed_inverted_turn {
		stage.UnstageBranch(_horizontal_turn)
	}
	for _, _empty_trill_sound := range ornaments.Vertical_turn {
		stage.UnstageBranch(_empty_trill_sound)
	}
	for _, _empty_trill_sound := range ornaments.Inverted_vertical_turn {
		stage.UnstageBranch(_empty_trill_sound)
	}
	for _, _empty_trill_sound := range ornaments.Shake {
		stage.UnstageBranch(_empty_trill_sound)
	}
	for _, _wavy_line := range ornaments.Wavy_line {
		stage.UnstageBranch(_wavy_line)
	}
	for _, _mordent := range ornaments.Mordent {
		stage.UnstageBranch(_mordent)
	}
	for _, _mordent := range ornaments.Inverted_mordent {
		stage.UnstageBranch(_mordent)
	}
	for _, _empty_placement := range ornaments.Schleifer {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _tremolo := range ornaments.Tremolo {
		stage.UnstageBranch(_tremolo)
	}
	for _, _empty_trill_sound := range ornaments.Haydn {
		stage.UnstageBranch(_empty_trill_sound)
	}
	for _, _other_placement_text := range ornaments.Other_ornament {
		stage.UnstageBranch(_other_placement_text)
	}
	for _, _accidental_mark := range ornaments.Accidental_mark {
		stage.UnstageBranch(_accidental_mark)
	}

}

func (other_appearance *Other_appearance) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(other_appearance) {
		return
	}

	other_appearance.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (other_direction *Other_direction) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(other_direction) {
		return
	}

	other_direction.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (other_listening *Other_listening) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(other_listening) {
		return
	}

	other_listening.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (other_notation *Other_notation) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(other_notation) {
		return
	}

	other_notation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (other_placement_text *Other_placement_text) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(other_placement_text) {
		return
	}

	other_placement_text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (other_play *Other_play) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(other_play) {
		return
	}

	other_play.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (other_text *Other_text) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(other_text) {
		return
	}

	other_text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (page_layout *Page_layout) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(page_layout) {
		return
	}

	page_layout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if page_layout.Page_margins != nil {
		stage.UnstageBranch(page_layout.Page_margins)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (page_margins *Page_margins) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(page_margins) {
		return
	}

	page_margins.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part_clef *Part_clef) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(part_clef) {
		return
	}

	part_clef.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part_group *Part_group) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(part_group) {
		return
	}

	part_group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if part_group.Group_name != nil {
		stage.UnstageBranch(part_group.Group_name)
	}
	if part_group.Group_name_display != nil {
		stage.UnstageBranch(part_group.Group_name_display)
	}
	if part_group.Group_abbreviation != nil {
		stage.UnstageBranch(part_group.Group_abbreviation)
	}
	if part_group.Group_abbreviation_display != nil {
		stage.UnstageBranch(part_group.Group_abbreviation_display)
	}
	if part_group.Group_symbol != nil {
		stage.UnstageBranch(part_group.Group_symbol)
	}
	if part_group.Group_barline != nil {
		stage.UnstageBranch(part_group.Group_barline)
	}
	if part_group.Footnote != nil {
		stage.UnstageBranch(part_group.Footnote)
	}
	if part_group.Level != nil {
		stage.UnstageBranch(part_group.Level)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part_link *Part_link) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(part_link) {
		return
	}

	part_link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument_link := range part_link.Instrument_link {
		stage.UnstageBranch(_instrument_link)
	}

}

func (part_list *Part_list) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(part_list) {
		return
	}

	part_list.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if part_list.Part_group != nil {
		stage.UnstageBranch(part_list.Part_group)
	}
	if part_list.Score_part != nil {
		stage.UnstageBranch(part_list.Score_part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part_name *Part_name) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(part_name) {
		return
	}

	part_name.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part_symbol *Part_symbol) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(part_symbol) {
		return
	}

	part_symbol.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part_transpose *Part_transpose) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(part_transpose) {
		return
	}

	part_transpose.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pedal *Pedal) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pedal) {
		return
	}

	pedal.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pedal_tuning *Pedal_tuning) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pedal_tuning) {
		return
	}

	pedal_tuning.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (per_minute *Per_minute) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(per_minute) {
		return
	}

	per_minute.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (percussion *Percussion) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(percussion) {
		return
	}

	percussion.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if percussion.Glass != nil {
		stage.UnstageBranch(percussion.Glass)
	}
	if percussion.Metal != nil {
		stage.UnstageBranch(percussion.Metal)
	}
	if percussion.Wood != nil {
		stage.UnstageBranch(percussion.Wood)
	}
	if percussion.Pitched != nil {
		stage.UnstageBranch(percussion.Pitched)
	}
	if percussion.Membrane != nil {
		stage.UnstageBranch(percussion.Membrane)
	}
	if percussion.Effect != nil {
		stage.UnstageBranch(percussion.Effect)
	}
	if percussion.Timpani != nil {
		stage.UnstageBranch(percussion.Timpani)
	}
	if percussion.Beater != nil {
		stage.UnstageBranch(percussion.Beater)
	}
	if percussion.Stick != nil {
		stage.UnstageBranch(percussion.Stick)
	}
	if percussion.Other_percussion != nil {
		stage.UnstageBranch(percussion.Other_percussion)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pitch *Pitch) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pitch) {
		return
	}

	pitch.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pitched *Pitched) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pitched) {
		return
	}

	pitched.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (placement_text *Placement_text) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(placement_text) {
		return
	}

	placement_text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (play *Play) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(play) {
		return
	}

	play.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _other_play := range play.Other_play {
		stage.UnstageBranch(_other_play)
	}

}

func (player *Player) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(player) {
		return
	}

	player.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (principal_voice *Principal_voice) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(principal_voice) {
		return
	}

	principal_voice.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (print *Print) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(print) {
		return
	}

	print.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if print.Page_layout != nil {
		stage.UnstageBranch(print.Page_layout)
	}
	if print.System_layout != nil {
		stage.UnstageBranch(print.System_layout)
	}
	if print.Measure_layout != nil {
		stage.UnstageBranch(print.Measure_layout)
	}
	if print.Measure_numbering != nil {
		stage.UnstageBranch(print.Measure_numbering)
	}
	if print.Part_name_display != nil {
		stage.UnstageBranch(print.Part_name_display)
	}
	if print.Part_abbreviation_display != nil {
		stage.UnstageBranch(print.Part_abbreviation_display)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _staff_layout := range print.Staff_layout {
		stage.UnstageBranch(_staff_layout)
	}

}

func (release *Release) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(release) {
		return
	}

	release.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (repeat *Repeat) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(repeat) {
		return
	}

	repeat.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rest *Rest) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rest) {
		return
	}

	rest.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (root *Root) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(root) {
		return
	}

	root.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if root.Root_step != nil {
		stage.UnstageBranch(root.Root_step)
	}
	if root.Root_alter != nil {
		stage.UnstageBranch(root.Root_alter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (root_step *Root_step) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(root_step) {
		return
	}

	root_step.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (scaling *Scaling) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(scaling) {
		return
	}

	scaling.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (scordatura *Scordatura) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(scordatura) {
		return
	}

	scordatura.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _accord := range scordatura.Accord {
		stage.UnstageBranch(_accord)
	}

}

func (score_instrument *Score_instrument) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(score_instrument) {
		return
	}

	score_instrument.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if score_instrument.Virtual_instrument != nil {
		stage.UnstageBranch(score_instrument.Virtual_instrument)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (score_part *Score_part) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(score_part) {
		return
	}

	score_part.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if score_part.Identification != nil {
		stage.UnstageBranch(score_part.Identification)
	}
	if score_part.Part_name != nil {
		stage.UnstageBranch(score_part.Part_name)
	}
	if score_part.Part_name_display != nil {
		stage.UnstageBranch(score_part.Part_name_display)
	}
	if score_part.Part_abbreviation != nil {
		stage.UnstageBranch(score_part.Part_abbreviation)
	}
	if score_part.Part_abbreviation_display != nil {
		stage.UnstageBranch(score_part.Part_abbreviation_display)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part_link := range score_part.Part_link {
		stage.UnstageBranch(_part_link)
	}
	for _, _score_instrument := range score_part.Score_instrument {
		stage.UnstageBranch(_score_instrument)
	}
	for _, _player := range score_part.Player {
		stage.UnstageBranch(_player)
	}
	for _, _midi_device := range score_part.Midi_device {
		stage.UnstageBranch(_midi_device)
	}
	for _, _midi_instrument := range score_part.Midi_instrument {
		stage.UnstageBranch(_midi_instrument)
	}

}

func (score_partwise *Score_partwise) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(score_partwise) {
		return
	}

	score_partwise.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if score_partwise.Work != nil {
		stage.UnstageBranch(score_partwise.Work)
	}
	if score_partwise.Identification != nil {
		stage.UnstageBranch(score_partwise.Identification)
	}
	if score_partwise.Defaults != nil {
		stage.UnstageBranch(score_partwise.Defaults)
	}
	if score_partwise.Part_list != nil {
		stage.UnstageBranch(score_partwise.Part_list)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range score_partwise.Credit {
		stage.UnstageBranch(_credit)
	}
	for _, _a_part := range score_partwise.Part {
		stage.UnstageBranch(_a_part)
	}

}

func (score_timewise *Score_timewise) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(score_timewise) {
		return
	}

	score_timewise.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if score_timewise.Work != nil {
		stage.UnstageBranch(score_timewise.Work)
	}
	if score_timewise.Identification != nil {
		stage.UnstageBranch(score_timewise.Identification)
	}
	if score_timewise.Defaults != nil {
		stage.UnstageBranch(score_timewise.Defaults)
	}
	if score_timewise.Part_list != nil {
		stage.UnstageBranch(score_timewise.Part_list)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range score_timewise.Credit {
		stage.UnstageBranch(_credit)
	}
	for _, _a_measure_1 := range score_timewise.Measure {
		stage.UnstageBranch(_a_measure_1)
	}

}

func (segno *Segno) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(segno) {
		return
	}

	segno.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (slash *Slash) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(slash) {
		return
	}

	slash.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (slide *Slide) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(slide) {
		return
	}

	slide.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (slur *Slur) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(slur) {
		return
	}

	slur.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (sound *Sound) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(sound) {
		return
	}

	sound.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sound.Swing != nil {
		stage.UnstageBranch(sound.Swing)
	}
	if sound.Offset != nil {
		stage.UnstageBranch(sound.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument_change := range sound.Instrument_change {
		stage.UnstageBranch(_instrument_change)
	}
	for _, _midi_device := range sound.Midi_device {
		stage.UnstageBranch(_midi_device)
	}
	for _, _midi_instrument := range sound.Midi_instrument {
		stage.UnstageBranch(_midi_instrument)
	}
	for _, _play := range sound.Play {
		stage.UnstageBranch(_play)
	}

}

func (staff_details *Staff_details) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(staff_details) {
		return
	}

	staff_details.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if staff_details.Staff_size != nil {
		stage.UnstageBranch(staff_details.Staff_size)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _line_detail := range staff_details.Line_detail {
		stage.UnstageBranch(_line_detail)
	}
	for _, _staff_tuning := range staff_details.Staff_tuning {
		stage.UnstageBranch(_staff_tuning)
	}

}

func (staff_divide *Staff_divide) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(staff_divide) {
		return
	}

	staff_divide.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (staff_layout *Staff_layout) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(staff_layout) {
		return
	}

	staff_layout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (staff_size *Staff_size) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(staff_size) {
		return
	}

	staff_size.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (staff_tuning *Staff_tuning) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(staff_tuning) {
		return
	}

	staff_tuning.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stem *Stem) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stem) {
		return
	}

	stem.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stick *Stick) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(stick) {
		return
	}

	stick.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (string_mute *String_mute) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(string_mute) {
		return
	}

	string_mute.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (string_type *String_type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(string_type) {
		return
	}

	string_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (strong_accent *Strong_accent) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(strong_accent) {
		return
	}

	strong_accent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (style_text *Style_text) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(style_text) {
		return
	}

	style_text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (supports *Supports) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(supports) {
		return
	}

	supports.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (swing *Swing) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(swing) {
		return
	}

	swing.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (sync *Sync) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(sync) {
		return
	}

	sync.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (system_dividers *System_dividers) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(system_dividers) {
		return
	}

	system_dividers.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if system_dividers.Left_divider != nil {
		stage.UnstageBranch(system_dividers.Left_divider)
	}
	if system_dividers.Right_divider != nil {
		stage.UnstageBranch(system_dividers.Right_divider)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (system_layout *System_layout) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(system_layout) {
		return
	}

	system_layout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if system_layout.System_margins != nil {
		stage.UnstageBranch(system_layout.System_margins)
	}
	if system_layout.System_dividers != nil {
		stage.UnstageBranch(system_layout.System_dividers)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (system_margins *System_margins) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(system_margins) {
		return
	}

	system_margins.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tap *Tap) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tap) {
		return
	}

	tap.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (technical *Technical) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(technical) {
		return
	}

	technical.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty_placement := range technical.Up_bow {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_placement := range technical.Down_bow {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _harmonic := range technical.Harmonic {
		stage.UnstageBranch(_harmonic)
	}
	for _, _empty_placement := range technical.Open_string {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_placement := range technical.Thumb_position {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _fingering := range technical.Fingering {
		stage.UnstageBranch(_fingering)
	}
	for _, _placement_text := range technical.Pluck {
		stage.UnstageBranch(_placement_text)
	}
	for _, _empty_placement := range technical.Double_tongue {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_placement := range technical.Triple_tongue {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_placement_smufl := range technical.Stopped {
		stage.UnstageBranch(_empty_placement_smufl)
	}
	for _, _empty_placement := range technical.Snap_pizzicato {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _fret := range technical.Fret {
		stage.UnstageBranch(_fret)
	}
	for _, _string_type := range technical.String {
		stage.UnstageBranch(_string_type)
	}
	for _, _hammer_on_pull_off := range technical.Hammer_on {
		stage.UnstageBranch(_hammer_on_pull_off)
	}
	for _, _hammer_on_pull_off := range technical.Pull_off {
		stage.UnstageBranch(_hammer_on_pull_off)
	}
	for _, _bend := range technical.Bend {
		stage.UnstageBranch(_bend)
	}
	for _, _tap := range technical.Tap {
		stage.UnstageBranch(_tap)
	}
	for _, _heel_toe := range technical.Heel {
		stage.UnstageBranch(_heel_toe)
	}
	for _, _heel_toe := range technical.Toe {
		stage.UnstageBranch(_heel_toe)
	}
	for _, _empty_placement := range technical.Fingernails {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _hole := range technical.Hole {
		stage.UnstageBranch(_hole)
	}
	for _, _arrow := range technical.Arrow {
		stage.UnstageBranch(_arrow)
	}
	for _, _handbell := range technical.Handbell {
		stage.UnstageBranch(_handbell)
	}
	for _, _empty_placement := range technical.Brass_bend {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_placement := range technical.Flip {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_placement := range technical.Smear {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _empty_placement_smufl := range technical.Open {
		stage.UnstageBranch(_empty_placement_smufl)
	}
	for _, _empty_placement_smufl := range technical.Half_muted {
		stage.UnstageBranch(_empty_placement_smufl)
	}
	for _, _harmon_mute := range technical.Harmon_mute {
		stage.UnstageBranch(_harmon_mute)
	}
	for _, _empty_placement := range technical.Golpe {
		stage.UnstageBranch(_empty_placement)
	}
	for _, _other_placement_text := range technical.Other_technical {
		stage.UnstageBranch(_other_placement_text)
	}

}

func (text_element_data *Text_element_data) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(text_element_data) {
		return
	}

	text_element_data.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tie *Tie) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tie) {
		return
	}

	tie.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tied *Tied) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tied) {
		return
	}

	tied.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (time *Time) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(time) {
		return
	}

	time.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if time.Interchangeable != nil {
		stage.UnstageBranch(time.Interchangeable)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (time_modification *Time_modification) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(time_modification) {
		return
	}

	time_modification.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timpani *Timpani) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timpani) {
		return
	}

	timpani.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (transpose *Transpose) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(transpose) {
		return
	}

	transpose.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tremolo *Tremolo) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tremolo) {
		return
	}

	tremolo.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tuplet *Tuplet) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tuplet) {
		return
	}

	tuplet.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tuplet.Tuplet_actual != nil {
		stage.UnstageBranch(tuplet.Tuplet_actual)
	}
	if tuplet.Tuplet_normal != nil {
		stage.UnstageBranch(tuplet.Tuplet_normal)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tuplet_dot *Tuplet_dot) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tuplet_dot) {
		return
	}

	tuplet_dot.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tuplet_number *Tuplet_number) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tuplet_number) {
		return
	}

	tuplet_number.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tuplet_portion *Tuplet_portion) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tuplet_portion) {
		return
	}

	tuplet_portion.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tuplet_portion.Tuplet_number != nil {
		stage.UnstageBranch(tuplet_portion.Tuplet_number)
	}
	if tuplet_portion.Tuplet_type != nil {
		stage.UnstageBranch(tuplet_portion.Tuplet_type)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tuplet_dot := range tuplet_portion.Tuplet_dot {
		stage.UnstageBranch(_tuplet_dot)
	}

}

func (tuplet_type *Tuplet_type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tuplet_type) {
		return
	}

	tuplet_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (typed_text *Typed_text) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(typed_text) {
		return
	}

	typed_text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (unpitched *Unpitched) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(unpitched) {
		return
	}

	unpitched.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (virtual_instrument *Virtual_instrument) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(virtual_instrument) {
		return
	}

	virtual_instrument.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wait *Wait) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(wait) {
		return
	}

	wait.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wavy_line *Wavy_line) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(wavy_line) {
		return
	}

	wavy_line.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wedge *Wedge) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(wedge) {
		return
	}

	wedge.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (wood *Wood) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(wood) {
		return
	}

	wood.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (work *Work) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(work) {
		return
	}

	work.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if work.Opus != nil {
		stage.UnstageBranch(work.Opus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *A_directive) GongReconstructPointersFromReferences(stage *Stage, instance *A_directive) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_measure) GongReconstructPointersFromReferences(stage *Stage, instance *A_measure) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Backup, stage.Backups_reference, instance.Backup)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Forward, stage.Forwards_reference, instance.Forward)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Direction, stage.Directions_reference, instance.Direction)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Attributes, stage.Attributess_reference, instance.Attributes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Harmony, stage.Harmonys_reference, instance.Harmony)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Figured_bass, stage.Figured_basss_reference, instance.Figured_bass)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Print, stage.Prints_reference, instance.Print)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sound, stage.Sounds_reference, instance.Sound)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Listening, stage.Listenings_reference, instance.Listening)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Barline, stage.Barlines_reference, instance.Barline)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Grouping, stage.Groupings_reference, instance.Grouping)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Link, stage.Links_reference, instance.Link)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Bookmark, stage.Bookmarks_reference, instance.Bookmark)
}

func (reference *A_measure_1) GongReconstructPointersFromReferences(stage *Stage, instance *A_measure_1) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Part, stage.A_part_1s_reference, instance.Part)
}

func (reference *A_part) GongReconstructPointersFromReferences(stage *Stage, instance *A_part) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Measure, stage.A_measures_reference, instance.Measure)
}

func (reference *A_part_1) GongReconstructPointersFromReferences(stage *Stage, instance *A_part_1) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Backup, stage.Backups_reference, instance.Backup)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Forward, stage.Forwards_reference, instance.Forward)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Direction, stage.Directions_reference, instance.Direction)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Attributes, stage.Attributess_reference, instance.Attributes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Harmony, stage.Harmonys_reference, instance.Harmony)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Figured_bass, stage.Figured_basss_reference, instance.Figured_bass)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Print, stage.Prints_reference, instance.Print)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sound, stage.Sounds_reference, instance.Sound)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Listening, stage.Listenings_reference, instance.Listening)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Barline, stage.Barlines_reference, instance.Barline)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Grouping, stage.Groupings_reference, instance.Grouping)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Link, stage.Links_reference, instance.Link)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Bookmark, stage.Bookmarks_reference, instance.Bookmark)
}

func (reference *Accidental) GongReconstructPointersFromReferences(stage *Stage, instance *Accidental) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Accidental_mark) GongReconstructPointersFromReferences(stage *Stage, instance *Accidental_mark) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Accidental_text) GongReconstructPointersFromReferences(stage *Stage, instance *Accidental_text) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Accord) GongReconstructPointersFromReferences(stage *Stage, instance *Accord) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Accordion_registration) GongReconstructPointersFromReferences(stage *Stage, instance *Accordion_registration) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Appearance) GongReconstructPointersFromReferences(stage *Stage, instance *Appearance) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Line_width, stage.Line_widths_reference, instance.Line_width)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Note_size, stage.Note_sizes_reference, instance.Note_size)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Distance, stage.Distances_reference, instance.Distance)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Glyph, stage.Glyphs_reference, instance.Glyph)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Other_appearance, stage.Other_appearances_reference, instance.Other_appearance)
}

func (reference *Arpeggiate) GongReconstructPointersFromReferences(stage *Stage, instance *Arpeggiate) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Arrow) GongReconstructPointersFromReferences(stage *Stage, instance *Arrow) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Articulations) GongReconstructPointersFromReferences(stage *Stage, instance *Articulations) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Accent, stage.Empty_placements_reference, instance.Accent)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Strong_accent, stage.Strong_accents_reference, instance.Strong_accent)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Staccato, stage.Empty_placements_reference, instance.Staccato)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tenuto, stage.Empty_placements_reference, instance.Tenuto)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Detached_legato, stage.Empty_placements_reference, instance.Detached_legato)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Staccatissimo, stage.Empty_placements_reference, instance.Staccatissimo)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Spiccato, stage.Empty_placements_reference, instance.Spiccato)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Scoop, stage.Empty_lines_reference, instance.Scoop)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Plop, stage.Empty_lines_reference, instance.Plop)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Doit, stage.Empty_lines_reference, instance.Doit)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Falloff, stage.Empty_lines_reference, instance.Falloff)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Breath_mark, stage.Breath_marks_reference, instance.Breath_mark)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Caesura, stage.Caesuras_reference, instance.Caesura)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Stress, stage.Empty_placements_reference, instance.Stress)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Unstress, stage.Empty_placements_reference, instance.Unstress)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Soft_accent, stage.Empty_placements_reference, instance.Soft_accent)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Other_articulation, stage.Other_placement_texts_reference, instance.Other_articulation)
}

func (reference *Assess) GongReconstructPointersFromReferences(stage *Stage, instance *Assess) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Attributes) GongReconstructPointersFromReferences(stage *Stage, instance *Attributes) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Footnote, stage.Formatted_texts_reference, instance.Footnote)
	__gong__reconstructPointer(&reference.Level, stage.Levels_reference, instance.Level)
	__gong__reconstructPointer(&reference.Part_symbol, stage.Part_symbols_reference, instance.Part_symbol)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Key, stage.Keys_reference, instance.Key)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Time, stage.Times_reference, instance.Time)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Clef, stage.Clefs_reference, instance.Clef)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Staff_details, stage.Staff_detailss_reference, instance.Staff_details)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Transpose, stage.Transposes_reference, instance.Transpose)
	__gong__reconstructSliceOfPointersFromReferences(&reference.For_part, stage.For_parts_reference, instance.For_part)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Directive, stage.A_directives_reference, instance.Directive)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Measure_style, stage.Measure_styles_reference, instance.Measure_style)
}

func (reference *Backup) GongReconstructPointersFromReferences(stage *Stage, instance *Backup) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Footnote, stage.Formatted_texts_reference, instance.Footnote)
	__gong__reconstructPointer(&reference.Level, stage.Levels_reference, instance.Level)
	// insertion point for slice of pointers field
}

func (reference *Bar_style_color) GongReconstructPointersFromReferences(stage *Stage, instance *Bar_style_color) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Barline) GongReconstructPointersFromReferences(stage *Stage, instance *Barline) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Bar_style, stage.Bar_style_colors_reference, instance.Bar_style)
	__gong__reconstructPointer(&reference.Footnote, stage.Formatted_texts_reference, instance.Footnote)
	__gong__reconstructPointer(&reference.Level, stage.Levels_reference, instance.Level)
	__gong__reconstructPointer(&reference.Wavy_line, stage.Wavy_lines_reference, instance.Wavy_line)
	__gong__reconstructPointer(&reference.Segno_1, stage.Segnos_reference, instance.Segno_1)
	__gong__reconstructPointer(&reference.Coda_1, stage.Codas_reference, instance.Coda_1)
	__gong__reconstructPointer(&reference.Fermata, stage.Fermatas_reference, instance.Fermata)
	__gong__reconstructPointer(&reference.Ending, stage.Endings_reference, instance.Ending)
	__gong__reconstructPointer(&reference.Repeat, stage.Repeats_reference, instance.Repeat)
	// insertion point for slice of pointers field
}

func (reference *Barre) GongReconstructPointersFromReferences(stage *Stage, instance *Barre) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Bass) GongReconstructPointersFromReferences(stage *Stage, instance *Bass) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Bass_separator, stage.Style_texts_reference, instance.Bass_separator)
	__gong__reconstructPointer(&reference.Bass_step, stage.Bass_steps_reference, instance.Bass_step)
	__gong__reconstructPointer(&reference.Bass_alter, stage.Harmony_alters_reference, instance.Bass_alter)
	// insertion point for slice of pointers field
}

func (reference *Bass_step) GongReconstructPointersFromReferences(stage *Stage, instance *Bass_step) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Beam) GongReconstructPointersFromReferences(stage *Stage, instance *Beam) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Beat_repeat) GongReconstructPointersFromReferences(stage *Stage, instance *Beat_repeat) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Beat_unit_tied) GongReconstructPointersFromReferences(stage *Stage, instance *Beat_unit_tied) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Beater) GongReconstructPointersFromReferences(stage *Stage, instance *Beater) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Bend) GongReconstructPointersFromReferences(stage *Stage, instance *Bend) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Release, stage.Releases_reference, instance.Release)
	__gong__reconstructPointer(&reference.With_bar, stage.Placement_texts_reference, instance.With_bar)
	// insertion point for slice of pointers field
}

func (reference *Bookmark) GongReconstructPointersFromReferences(stage *Stage, instance *Bookmark) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Bracket) GongReconstructPointersFromReferences(stage *Stage, instance *Bracket) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Breath_mark) GongReconstructPointersFromReferences(stage *Stage, instance *Breath_mark) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Caesura) GongReconstructPointersFromReferences(stage *Stage, instance *Caesura) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Cancel) GongReconstructPointersFromReferences(stage *Stage, instance *Cancel) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Clef) GongReconstructPointersFromReferences(stage *Stage, instance *Clef) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Coda) GongReconstructPointersFromReferences(stage *Stage, instance *Coda) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Credit) GongReconstructPointersFromReferences(stage *Stage, instance *Credit) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Credit_image, stage.Images_reference, instance.Credit_image)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Link, stage.Links_reference, instance.Link)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Bookmark, stage.Bookmarks_reference, instance.Bookmark)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Credit_words, stage.Formatted_text_ids_reference, instance.Credit_words)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Credit_symbol, stage.Formatted_symbol_ids_reference, instance.Credit_symbol)
}

func (reference *Dashes) GongReconstructPointersFromReferences(stage *Stage, instance *Dashes) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Defaults) GongReconstructPointersFromReferences(stage *Stage, instance *Defaults) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Scaling, stage.Scalings_reference, instance.Scaling)
	__gong__reconstructPointer(&reference.Page_layout, stage.Page_layouts_reference, instance.Page_layout)
	__gong__reconstructPointer(&reference.System_layout, stage.System_layouts_reference, instance.System_layout)
	__gong__reconstructPointer(&reference.Appearance, stage.Appearances_reference, instance.Appearance)
	__gong__reconstructPointer(&reference.Music_font, stage.Empty_fonts_reference, instance.Music_font)
	__gong__reconstructPointer(&reference.Word_font, stage.Empty_fonts_reference, instance.Word_font)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Staff_layout, stage.Staff_layouts_reference, instance.Staff_layout)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Lyric_font, stage.Lyric_fonts_reference, instance.Lyric_font)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Lyric_language, stage.Lyric_languages_reference, instance.Lyric_language)
}

func (reference *Degree) GongReconstructPointersFromReferences(stage *Stage, instance *Degree) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Degree_value, stage.Degree_values_reference, instance.Degree_value)
	__gong__reconstructPointer(&reference.Degree_alter, stage.Degree_alters_reference, instance.Degree_alter)
	__gong__reconstructPointer(&reference.Degree_type, stage.Degree_types_reference, instance.Degree_type)
	// insertion point for slice of pointers field
}

func (reference *Degree_alter) GongReconstructPointersFromReferences(stage *Stage, instance *Degree_alter) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Degree_type) GongReconstructPointersFromReferences(stage *Stage, instance *Degree_type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Degree_value) GongReconstructPointersFromReferences(stage *Stage, instance *Degree_value) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Direction) GongReconstructPointersFromReferences(stage *Stage, instance *Direction) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Offset, stage.Offsets_reference, instance.Offset)
	__gong__reconstructPointer(&reference.Footnote, stage.Formatted_texts_reference, instance.Footnote)
	__gong__reconstructPointer(&reference.Level, stage.Levels_reference, instance.Level)
	__gong__reconstructPointer(&reference.Sound, stage.Sounds_reference, instance.Sound)
	__gong__reconstructPointer(&reference.Listening, stage.Listenings_reference, instance.Listening)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Direction_type, stage.Direction_types_reference, instance.Direction_type)
}

func (reference *Direction_type) GongReconstructPointersFromReferences(stage *Stage, instance *Direction_type) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Wedge, stage.Wedges_reference, instance.Wedge)
	__gong__reconstructPointer(&reference.Dashes, stage.Dashess_reference, instance.Dashes)
	__gong__reconstructPointer(&reference.Bracket, stage.Brackets_reference, instance.Bracket)
	__gong__reconstructPointer(&reference.Pedal, stage.Pedals_reference, instance.Pedal)
	__gong__reconstructPointer(&reference.Metronome, stage.Metronomes_reference, instance.Metronome)
	__gong__reconstructPointer(&reference.Octave_shift, stage.Octave_shifts_reference, instance.Octave_shift)
	__gong__reconstructPointer(&reference.Harp_pedals, stage.Harp_pedalss_reference, instance.Harp_pedals)
	__gong__reconstructPointer(&reference.Damp, stage.Empty_print_style_align_ids_reference, instance.Damp)
	__gong__reconstructPointer(&reference.Damp_all, stage.Empty_print_style_align_ids_reference, instance.Damp_all)
	__gong__reconstructPointer(&reference.Eyeglasses, stage.Empty_print_style_align_ids_reference, instance.Eyeglasses)
	__gong__reconstructPointer(&reference.String_mute, stage.String_mutes_reference, instance.String_mute)
	__gong__reconstructPointer(&reference.Scordatura, stage.Scordaturas_reference, instance.Scordatura)
	__gong__reconstructPointer(&reference.Image, stage.Images_reference, instance.Image)
	__gong__reconstructPointer(&reference.Principal_voice, stage.Principal_voices_reference, instance.Principal_voice)
	__gong__reconstructPointer(&reference.Accordion_registration, stage.Accordion_registrations_reference, instance.Accordion_registration)
	__gong__reconstructPointer(&reference.Staff_divide, stage.Staff_divides_reference, instance.Staff_divide)
	__gong__reconstructPointer(&reference.Other_direction, stage.Other_directions_reference, instance.Other_direction)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Rehearsal, stage.Formatted_text_ids_reference, instance.Rehearsal)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Segno, stage.Segnos_reference, instance.Segno)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Coda, stage.Codas_reference, instance.Coda)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Words, stage.Formatted_text_ids_reference, instance.Words)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Symbol, stage.Formatted_symbol_ids_reference, instance.Symbol)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Dynamics, stage.Dynamicss_reference, instance.Dynamics)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Percussion, stage.Percussions_reference, instance.Percussion)
}

func (reference *Distance) GongReconstructPointersFromReferences(stage *Stage, instance *Distance) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Double) GongReconstructPointersFromReferences(stage *Stage, instance *Double) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Dynamics) GongReconstructPointersFromReferences(stage *Stage, instance *Dynamics) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Other_dynamics, stage.Other_texts_reference, instance.Other_dynamics)
}

func (reference *Effect) GongReconstructPointersFromReferences(stage *Stage, instance *Effect) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Elision) GongReconstructPointersFromReferences(stage *Stage, instance *Elision) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Empty) GongReconstructPointersFromReferences(stage *Stage, instance *Empty) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Empty_font) GongReconstructPointersFromReferences(stage *Stage, instance *Empty_font) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Empty_line) GongReconstructPointersFromReferences(stage *Stage, instance *Empty_line) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Empty_placement) GongReconstructPointersFromReferences(stage *Stage, instance *Empty_placement) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Empty_placement_smufl) GongReconstructPointersFromReferences(stage *Stage, instance *Empty_placement_smufl) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Empty_print_object_style_align) GongReconstructPointersFromReferences(stage *Stage, instance *Empty_print_object_style_align) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Empty_print_style) GongReconstructPointersFromReferences(stage *Stage, instance *Empty_print_style) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Empty_print_style_align) GongReconstructPointersFromReferences(stage *Stage, instance *Empty_print_style_align) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Empty_print_style_align_id) GongReconstructPointersFromReferences(stage *Stage, instance *Empty_print_style_align_id) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Empty_trill_sound) GongReconstructPointersFromReferences(stage *Stage, instance *Empty_trill_sound) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Encoding) GongReconstructPointersFromReferences(stage *Stage, instance *Encoding) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Encoder, stage.Typed_texts_reference, instance.Encoder)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Supports, stage.Supportss_reference, instance.Supports)
}

func (reference *Ending) GongReconstructPointersFromReferences(stage *Stage, instance *Ending) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Extend) GongReconstructPointersFromReferences(stage *Stage, instance *Extend) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Feature) GongReconstructPointersFromReferences(stage *Stage, instance *Feature) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Fermata) GongReconstructPointersFromReferences(stage *Stage, instance *Fermata) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Figure) GongReconstructPointersFromReferences(stage *Stage, instance *Figure) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Prefix, stage.Style_texts_reference, instance.Prefix)
	__gong__reconstructPointer(&reference.Figure_number, stage.Style_texts_reference, instance.Figure_number)
	__gong__reconstructPointer(&reference.Suffix, stage.Style_texts_reference, instance.Suffix)
	__gong__reconstructPointer(&reference.Extend, stage.Extends_reference, instance.Extend)
	__gong__reconstructPointer(&reference.Footnote, stage.Formatted_texts_reference, instance.Footnote)
	__gong__reconstructPointer(&reference.Level, stage.Levels_reference, instance.Level)
	// insertion point for slice of pointers field
}

func (reference *Figured_bass) GongReconstructPointersFromReferences(stage *Stage, instance *Figured_bass) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Footnote, stage.Formatted_texts_reference, instance.Footnote)
	__gong__reconstructPointer(&reference.Level, stage.Levels_reference, instance.Level)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Figure, stage.Figures_reference, instance.Figure)
}

func (reference *Fingering) GongReconstructPointersFromReferences(stage *Stage, instance *Fingering) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *First_fret) GongReconstructPointersFromReferences(stage *Stage, instance *First_fret) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *For_part) GongReconstructPointersFromReferences(stage *Stage, instance *For_part) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Part_clef, stage.Part_clefs_reference, instance.Part_clef)
	__gong__reconstructPointer(&reference.Part_transpose, stage.Part_transposes_reference, instance.Part_transpose)
	// insertion point for slice of pointers field
}

func (reference *Formatted_symbol) GongReconstructPointersFromReferences(stage *Stage, instance *Formatted_symbol) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Formatted_symbol_id) GongReconstructPointersFromReferences(stage *Stage, instance *Formatted_symbol_id) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Formatted_text) GongReconstructPointersFromReferences(stage *Stage, instance *Formatted_text) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Formatted_text_id) GongReconstructPointersFromReferences(stage *Stage, instance *Formatted_text_id) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Forward) GongReconstructPointersFromReferences(stage *Stage, instance *Forward) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Footnote, stage.Formatted_texts_reference, instance.Footnote)
	__gong__reconstructPointer(&reference.Level, stage.Levels_reference, instance.Level)
	// insertion point for slice of pointers field
}

func (reference *Frame) GongReconstructPointersFromReferences(stage *Stage, instance *Frame) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.First_fret, stage.First_frets_reference, instance.First_fret)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Frame_note, stage.Frame_notes_reference, instance.Frame_note)
}

func (reference *Frame_note) GongReconstructPointersFromReferences(stage *Stage, instance *Frame_note) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.String, stage.String_types_reference, instance.String)
	__gong__reconstructPointer(&reference.Fret, stage.Frets_reference, instance.Fret)
	__gong__reconstructPointer(&reference.Fingering, stage.Fingerings_reference, instance.Fingering)
	__gong__reconstructPointer(&reference.Barre, stage.Barres_reference, instance.Barre)
	// insertion point for slice of pointers field
}

func (reference *Fret) GongReconstructPointersFromReferences(stage *Stage, instance *Fret) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Glass) GongReconstructPointersFromReferences(stage *Stage, instance *Glass) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Glissando) GongReconstructPointersFromReferences(stage *Stage, instance *Glissando) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Glyph) GongReconstructPointersFromReferences(stage *Stage, instance *Glyph) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Grace) GongReconstructPointersFromReferences(stage *Stage, instance *Grace) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Group_barline) GongReconstructPointersFromReferences(stage *Stage, instance *Group_barline) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Group_name) GongReconstructPointersFromReferences(stage *Stage, instance *Group_name) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Group_symbol) GongReconstructPointersFromReferences(stage *Stage, instance *Group_symbol) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Grouping) GongReconstructPointersFromReferences(stage *Stage, instance *Grouping) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Feature, stage.Features_reference, instance.Feature)
}

func (reference *Hammer_on_pull_off) GongReconstructPointersFromReferences(stage *Stage, instance *Hammer_on_pull_off) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Handbell) GongReconstructPointersFromReferences(stage *Stage, instance *Handbell) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Harmon_closed) GongReconstructPointersFromReferences(stage *Stage, instance *Harmon_closed) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Harmon_mute) GongReconstructPointersFromReferences(stage *Stage, instance *Harmon_mute) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Harmon_closed, stage.Harmon_closeds_reference, instance.Harmon_closed)
	// insertion point for slice of pointers field
}

func (reference *Harmonic) GongReconstructPointersFromReferences(stage *Stage, instance *Harmonic) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Harmony) GongReconstructPointersFromReferences(stage *Stage, instance *Harmony) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Root, stage.Roots_reference, instance.Root)
	__gong__reconstructPointer(&reference.Numeral, stage.Numerals_reference, instance.Numeral)
	__gong__reconstructPointer(&reference.Function, stage.Style_texts_reference, instance.Function)
	__gong__reconstructPointer(&reference.Kind, stage.Kinds_reference, instance.Kind)
	__gong__reconstructPointer(&reference.Inversion, stage.Inversions_reference, instance.Inversion)
	__gong__reconstructPointer(&reference.Bass, stage.Basss_reference, instance.Bass)
	__gong__reconstructPointer(&reference.Frame, stage.Frames_reference, instance.Frame)
	__gong__reconstructPointer(&reference.Offset, stage.Offsets_reference, instance.Offset)
	__gong__reconstructPointer(&reference.Footnote, stage.Formatted_texts_reference, instance.Footnote)
	__gong__reconstructPointer(&reference.Level, stage.Levels_reference, instance.Level)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Degree, stage.Degrees_reference, instance.Degree)
}

func (reference *Harmony_alter) GongReconstructPointersFromReferences(stage *Stage, instance *Harmony_alter) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Harp_pedals) GongReconstructPointersFromReferences(stage *Stage, instance *Harp_pedals) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Pedal_tuning, stage.Pedal_tunings_reference, instance.Pedal_tuning)
}

func (reference *Heel_toe) GongReconstructPointersFromReferences(stage *Stage, instance *Heel_toe) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Hole) GongReconstructPointersFromReferences(stage *Stage, instance *Hole) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Hole_closed, stage.Hole_closeds_reference, instance.Hole_closed)
	// insertion point for slice of pointers field
}

func (reference *Hole_closed) GongReconstructPointersFromReferences(stage *Stage, instance *Hole_closed) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Horizontal_turn) GongReconstructPointersFromReferences(stage *Stage, instance *Horizontal_turn) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Identification) GongReconstructPointersFromReferences(stage *Stage, instance *Identification) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Encoding, stage.Encodings_reference, instance.Encoding)
	__gong__reconstructPointer(&reference.Miscellaneous, stage.Miscellaneouss_reference, instance.Miscellaneous)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Creator, stage.Typed_texts_reference, instance.Creator)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Rights, stage.Typed_texts_reference, instance.Rights)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Relation, stage.Typed_texts_reference, instance.Relation)
}

func (reference *Image) GongReconstructPointersFromReferences(stage *Stage, instance *Image) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Instrument) GongReconstructPointersFromReferences(stage *Stage, instance *Instrument) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Instrument_change) GongReconstructPointersFromReferences(stage *Stage, instance *Instrument_change) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Virtual_instrument, stage.Virtual_instruments_reference, instance.Virtual_instrument)
	// insertion point for slice of pointers field
}

func (reference *Instrument_link) GongReconstructPointersFromReferences(stage *Stage, instance *Instrument_link) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Interchangeable) GongReconstructPointersFromReferences(stage *Stage, instance *Interchangeable) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Inversion) GongReconstructPointersFromReferences(stage *Stage, instance *Inversion) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Key) GongReconstructPointersFromReferences(stage *Stage, instance *Key) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Cancel, stage.Cancels_reference, instance.Cancel)
	__gong__reconstructPointer(&reference.Key_accidental, stage.Key_accidentals_reference, instance.Key_accidental)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Key_octave, stage.Key_octaves_reference, instance.Key_octave)
}

func (reference *Key_accidental) GongReconstructPointersFromReferences(stage *Stage, instance *Key_accidental) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Key_octave) GongReconstructPointersFromReferences(stage *Stage, instance *Key_octave) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Kind) GongReconstructPointersFromReferences(stage *Stage, instance *Kind) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Level) GongReconstructPointersFromReferences(stage *Stage, instance *Level) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Line_detail) GongReconstructPointersFromReferences(stage *Stage, instance *Line_detail) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Line_width) GongReconstructPointersFromReferences(stage *Stage, instance *Line_width) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Link) GongReconstructPointersFromReferences(stage *Stage, instance *Link) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Listen) GongReconstructPointersFromReferences(stage *Stage, instance *Listen) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Assess, stage.Assesss_reference, instance.Assess)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Wait, stage.Waits_reference, instance.Wait)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Other_listen, stage.Other_listenings_reference, instance.Other_listen)
}

func (reference *Listening) GongReconstructPointersFromReferences(stage *Stage, instance *Listening) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Offset, stage.Offsets_reference, instance.Offset)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sync, stage.Syncs_reference, instance.Sync)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Other_listening, stage.Other_listenings_reference, instance.Other_listening)
}

func (reference *Lyric) GongReconstructPointersFromReferences(stage *Stage, instance *Lyric) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Extend, stage.Extends_reference, instance.Extend)
	__gong__reconstructPointer(&reference.Footnote, stage.Formatted_texts_reference, instance.Footnote)
	__gong__reconstructPointer(&reference.Level, stage.Levels_reference, instance.Level)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Elision, stage.Elisions_reference, instance.Elision)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Text, stage.Text_element_datas_reference, instance.Text)
}

func (reference *Lyric_font) GongReconstructPointersFromReferences(stage *Stage, instance *Lyric_font) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Lyric_language) GongReconstructPointersFromReferences(stage *Stage, instance *Lyric_language) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Measure_layout) GongReconstructPointersFromReferences(stage *Stage, instance *Measure_layout) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Measure_numbering) GongReconstructPointersFromReferences(stage *Stage, instance *Measure_numbering) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Measure_repeat) GongReconstructPointersFromReferences(stage *Stage, instance *Measure_repeat) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Measure_style) GongReconstructPointersFromReferences(stage *Stage, instance *Measure_style) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Multiple_rest, stage.Multiple_rests_reference, instance.Multiple_rest)
	__gong__reconstructPointer(&reference.Measure_repeat, stage.Measure_repeats_reference, instance.Measure_repeat)
	__gong__reconstructPointer(&reference.Beat_repeat, stage.Beat_repeats_reference, instance.Beat_repeat)
	__gong__reconstructPointer(&reference.Slash, stage.Slashs_reference, instance.Slash)
	// insertion point for slice of pointers field
}

func (reference *Membrane) GongReconstructPointersFromReferences(stage *Stage, instance *Membrane) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Metal) GongReconstructPointersFromReferences(stage *Stage, instance *Metal) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Metronome) GongReconstructPointersFromReferences(stage *Stage, instance *Metronome) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Per_minute, stage.Per_minutes_reference, instance.Per_minute)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Beat_unit_tied, stage.Beat_unit_tieds_reference, instance.Beat_unit_tied)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Metronome_note, stage.Metronome_notes_reference, instance.Metronome_note)
}

func (reference *Metronome_beam) GongReconstructPointersFromReferences(stage *Stage, instance *Metronome_beam) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Metronome_note) GongReconstructPointersFromReferences(stage *Stage, instance *Metronome_note) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Metronome_tied, stage.Metronome_tieds_reference, instance.Metronome_tied)
	__gong__reconstructPointer(&reference.Metronome_tuplet, stage.Metronome_tuplets_reference, instance.Metronome_tuplet)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Metronome_beam, stage.Metronome_beams_reference, instance.Metronome_beam)
}

func (reference *Metronome_tied) GongReconstructPointersFromReferences(stage *Stage, instance *Metronome_tied) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Metronome_tuplet) GongReconstructPointersFromReferences(stage *Stage, instance *Metronome_tuplet) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Midi_device) GongReconstructPointersFromReferences(stage *Stage, instance *Midi_device) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Midi_instrument) GongReconstructPointersFromReferences(stage *Stage, instance *Midi_instrument) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Miscellaneous) GongReconstructPointersFromReferences(stage *Stage, instance *Miscellaneous) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Miscellaneous_field, stage.Miscellaneous_fields_reference, instance.Miscellaneous_field)
}

func (reference *Miscellaneous_field) GongReconstructPointersFromReferences(stage *Stage, instance *Miscellaneous_field) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Mordent) GongReconstructPointersFromReferences(stage *Stage, instance *Mordent) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Multiple_rest) GongReconstructPointersFromReferences(stage *Stage, instance *Multiple_rest) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Name_display) GongReconstructPointersFromReferences(stage *Stage, instance *Name_display) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Display_text, stage.Formatted_texts_reference, instance.Display_text)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Accidental_text, stage.Accidental_texts_reference, instance.Accidental_text)
}

func (reference *Non_arpeggiate) GongReconstructPointersFromReferences(stage *Stage, instance *Non_arpeggiate) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Notations) GongReconstructPointersFromReferences(stage *Stage, instance *Notations) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Footnote, stage.Formatted_texts_reference, instance.Footnote)
	__gong__reconstructPointer(&reference.Level, stage.Levels_reference, instance.Level)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tied, stage.Tieds_reference, instance.Tied)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Slur, stage.Slurs_reference, instance.Slur)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tuplet, stage.Tuplets_reference, instance.Tuplet)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Glissando, stage.Glissandos_reference, instance.Glissando)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Slide, stage.Slides_reference, instance.Slide)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Ornaments, stage.Ornamentss_reference, instance.Ornaments)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Technical, stage.Technicals_reference, instance.Technical)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Articulations, stage.Articulationss_reference, instance.Articulations)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Dynamics, stage.Dynamicss_reference, instance.Dynamics)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Fermata, stage.Fermatas_reference, instance.Fermata)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Arpeggiate, stage.Arpeggiates_reference, instance.Arpeggiate)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Non_arpeggiate, stage.Non_arpeggiates_reference, instance.Non_arpeggiate)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Accidental_mark, stage.Accidental_marks_reference, instance.Accidental_mark)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Other_notation, stage.Other_notations_reference, instance.Other_notation)
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Grace, stage.Graces_reference, instance.Grace)
	__gong__reconstructPointer(&reference.Pitch, stage.Pitchs_reference, instance.Pitch)
	__gong__reconstructPointer(&reference.Unpitched, stage.Unpitcheds_reference, instance.Unpitched)
	__gong__reconstructPointer(&reference.Rest, stage.Rests_reference, instance.Rest)
	__gong__reconstructPointer(&reference.Tie, stage.Ties_reference, instance.Tie)
	__gong__reconstructPointer(&reference.Footnote, stage.Formatted_texts_reference, instance.Footnote)
	__gong__reconstructPointer(&reference.Level, stage.Levels_reference, instance.Level)
	__gong__reconstructPointer(&reference.Type, stage.Note_types_reference, instance.Type)
	__gong__reconstructPointer(&reference.Accidental, stage.Accidentals_reference, instance.Accidental)
	__gong__reconstructPointer(&reference.Time_modification, stage.Time_modifications_reference, instance.Time_modification)
	__gong__reconstructPointer(&reference.Stem, stage.Stems_reference, instance.Stem)
	__gong__reconstructPointer(&reference.Notehead, stage.Noteheads_reference, instance.Notehead)
	__gong__reconstructPointer(&reference.Notehead_text, stage.Notehead_texts_reference, instance.Notehead_text)
	__gong__reconstructPointer(&reference.Beam, stage.Beams_reference, instance.Beam)
	__gong__reconstructPointer(&reference.Play, stage.Plays_reference, instance.Play)
	__gong__reconstructPointer(&reference.Listen, stage.Listens_reference, instance.Listen)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Instrument, stage.Instruments_reference, instance.Instrument)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Dot, stage.Empty_placements_reference, instance.Dot)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Notations, stage.Notationss_reference, instance.Notations)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Lyric, stage.Lyrics_reference, instance.Lyric)
}

func (reference *Note_size) GongReconstructPointersFromReferences(stage *Stage, instance *Note_size) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Note_type) GongReconstructPointersFromReferences(stage *Stage, instance *Note_type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Notehead) GongReconstructPointersFromReferences(stage *Stage, instance *Notehead) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Notehead_text) GongReconstructPointersFromReferences(stage *Stage, instance *Notehead_text) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Display_text, stage.Formatted_texts_reference, instance.Display_text)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Accidental_text, stage.Accidental_texts_reference, instance.Accidental_text)
}

func (reference *Numeral) GongReconstructPointersFromReferences(stage *Stage, instance *Numeral) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Numeral_root, stage.Numeral_roots_reference, instance.Numeral_root)
	__gong__reconstructPointer(&reference.Numeral_alter, stage.Harmony_alters_reference, instance.Numeral_alter)
	__gong__reconstructPointer(&reference.Numeral_key, stage.Numeral_keys_reference, instance.Numeral_key)
	// insertion point for slice of pointers field
}

func (reference *Numeral_key) GongReconstructPointersFromReferences(stage *Stage, instance *Numeral_key) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Numeral_root) GongReconstructPointersFromReferences(stage *Stage, instance *Numeral_root) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Octave_shift) GongReconstructPointersFromReferences(stage *Stage, instance *Octave_shift) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Offset) GongReconstructPointersFromReferences(stage *Stage, instance *Offset) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Opus) GongReconstructPointersFromReferences(stage *Stage, instance *Opus) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Ornaments) GongReconstructPointersFromReferences(stage *Stage, instance *Ornaments) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Trill_mark, stage.Empty_trill_sounds_reference, instance.Trill_mark)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Turn, stage.Horizontal_turns_reference, instance.Turn)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Delayed_turn, stage.Horizontal_turns_reference, instance.Delayed_turn)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Inverted_turn, stage.Horizontal_turns_reference, instance.Inverted_turn)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Delayed_inverted_turn, stage.Horizontal_turns_reference, instance.Delayed_inverted_turn)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Vertical_turn, stage.Empty_trill_sounds_reference, instance.Vertical_turn)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Inverted_vertical_turn, stage.Empty_trill_sounds_reference, instance.Inverted_vertical_turn)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Shake, stage.Empty_trill_sounds_reference, instance.Shake)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Wavy_line, stage.Wavy_lines_reference, instance.Wavy_line)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Mordent, stage.Mordents_reference, instance.Mordent)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Inverted_mordent, stage.Mordents_reference, instance.Inverted_mordent)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Schleifer, stage.Empty_placements_reference, instance.Schleifer)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tremolo, stage.Tremolos_reference, instance.Tremolo)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Haydn, stage.Empty_trill_sounds_reference, instance.Haydn)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Other_ornament, stage.Other_placement_texts_reference, instance.Other_ornament)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Accidental_mark, stage.Accidental_marks_reference, instance.Accidental_mark)
}

func (reference *Other_appearance) GongReconstructPointersFromReferences(stage *Stage, instance *Other_appearance) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Other_direction) GongReconstructPointersFromReferences(stage *Stage, instance *Other_direction) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Other_listening) GongReconstructPointersFromReferences(stage *Stage, instance *Other_listening) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Other_notation) GongReconstructPointersFromReferences(stage *Stage, instance *Other_notation) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Other_placement_text) GongReconstructPointersFromReferences(stage *Stage, instance *Other_placement_text) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Other_play) GongReconstructPointersFromReferences(stage *Stage, instance *Other_play) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Other_text) GongReconstructPointersFromReferences(stage *Stage, instance *Other_text) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Page_layout) GongReconstructPointersFromReferences(stage *Stage, instance *Page_layout) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Page_margins, stage.Page_marginss_reference, instance.Page_margins)
	// insertion point for slice of pointers field
}

func (reference *Page_margins) GongReconstructPointersFromReferences(stage *Stage, instance *Page_margins) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Part_clef) GongReconstructPointersFromReferences(stage *Stage, instance *Part_clef) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Part_group) GongReconstructPointersFromReferences(stage *Stage, instance *Part_group) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Group_name, stage.Group_names_reference, instance.Group_name)
	__gong__reconstructPointer(&reference.Group_name_display, stage.Name_displays_reference, instance.Group_name_display)
	__gong__reconstructPointer(&reference.Group_abbreviation, stage.Group_names_reference, instance.Group_abbreviation)
	__gong__reconstructPointer(&reference.Group_abbreviation_display, stage.Name_displays_reference, instance.Group_abbreviation_display)
	__gong__reconstructPointer(&reference.Group_symbol, stage.Group_symbols_reference, instance.Group_symbol)
	__gong__reconstructPointer(&reference.Group_barline, stage.Group_barlines_reference, instance.Group_barline)
	__gong__reconstructPointer(&reference.Footnote, stage.Formatted_texts_reference, instance.Footnote)
	__gong__reconstructPointer(&reference.Level, stage.Levels_reference, instance.Level)
	// insertion point for slice of pointers field
}

func (reference *Part_link) GongReconstructPointersFromReferences(stage *Stage, instance *Part_link) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Instrument_link, stage.Instrument_links_reference, instance.Instrument_link)
}

func (reference *Part_list) GongReconstructPointersFromReferences(stage *Stage, instance *Part_list) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Part_group, stage.Part_groups_reference, instance.Part_group)
	__gong__reconstructPointer(&reference.Score_part, stage.Score_parts_reference, instance.Score_part)
	// insertion point for slice of pointers field
}

func (reference *Part_name) GongReconstructPointersFromReferences(stage *Stage, instance *Part_name) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Part_symbol) GongReconstructPointersFromReferences(stage *Stage, instance *Part_symbol) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Part_transpose) GongReconstructPointersFromReferences(stage *Stage, instance *Part_transpose) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Pedal) GongReconstructPointersFromReferences(stage *Stage, instance *Pedal) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Pedal_tuning) GongReconstructPointersFromReferences(stage *Stage, instance *Pedal_tuning) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Per_minute) GongReconstructPointersFromReferences(stage *Stage, instance *Per_minute) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Percussion) GongReconstructPointersFromReferences(stage *Stage, instance *Percussion) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Glass, stage.Glasss_reference, instance.Glass)
	__gong__reconstructPointer(&reference.Metal, stage.Metals_reference, instance.Metal)
	__gong__reconstructPointer(&reference.Wood, stage.Woods_reference, instance.Wood)
	__gong__reconstructPointer(&reference.Pitched, stage.Pitcheds_reference, instance.Pitched)
	__gong__reconstructPointer(&reference.Membrane, stage.Membranes_reference, instance.Membrane)
	__gong__reconstructPointer(&reference.Effect, stage.Effects_reference, instance.Effect)
	__gong__reconstructPointer(&reference.Timpani, stage.Timpanis_reference, instance.Timpani)
	__gong__reconstructPointer(&reference.Beater, stage.Beaters_reference, instance.Beater)
	__gong__reconstructPointer(&reference.Stick, stage.Sticks_reference, instance.Stick)
	__gong__reconstructPointer(&reference.Other_percussion, stage.Other_texts_reference, instance.Other_percussion)
	// insertion point for slice of pointers field
}

func (reference *Pitch) GongReconstructPointersFromReferences(stage *Stage, instance *Pitch) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Pitched) GongReconstructPointersFromReferences(stage *Stage, instance *Pitched) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Placement_text) GongReconstructPointersFromReferences(stage *Stage, instance *Placement_text) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Play) GongReconstructPointersFromReferences(stage *Stage, instance *Play) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Other_play, stage.Other_plays_reference, instance.Other_play)
}

func (reference *Player) GongReconstructPointersFromReferences(stage *Stage, instance *Player) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Principal_voice) GongReconstructPointersFromReferences(stage *Stage, instance *Principal_voice) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Print) GongReconstructPointersFromReferences(stage *Stage, instance *Print) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Page_layout, stage.Page_layouts_reference, instance.Page_layout)
	__gong__reconstructPointer(&reference.System_layout, stage.System_layouts_reference, instance.System_layout)
	__gong__reconstructPointer(&reference.Measure_layout, stage.Measure_layouts_reference, instance.Measure_layout)
	__gong__reconstructPointer(&reference.Measure_numbering, stage.Measure_numberings_reference, instance.Measure_numbering)
	__gong__reconstructPointer(&reference.Part_name_display, stage.Name_displays_reference, instance.Part_name_display)
	__gong__reconstructPointer(&reference.Part_abbreviation_display, stage.Name_displays_reference, instance.Part_abbreviation_display)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Staff_layout, stage.Staff_layouts_reference, instance.Staff_layout)
}

func (reference *Release) GongReconstructPointersFromReferences(stage *Stage, instance *Release) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Repeat) GongReconstructPointersFromReferences(stage *Stage, instance *Repeat) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Rest) GongReconstructPointersFromReferences(stage *Stage, instance *Rest) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Root) GongReconstructPointersFromReferences(stage *Stage, instance *Root) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Root_step, stage.Root_steps_reference, instance.Root_step)
	__gong__reconstructPointer(&reference.Root_alter, stage.Harmony_alters_reference, instance.Root_alter)
	// insertion point for slice of pointers field
}

func (reference *Root_step) GongReconstructPointersFromReferences(stage *Stage, instance *Root_step) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Scaling) GongReconstructPointersFromReferences(stage *Stage, instance *Scaling) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Scordatura) GongReconstructPointersFromReferences(stage *Stage, instance *Scordatura) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Accord, stage.Accords_reference, instance.Accord)
}

func (reference *Score_instrument) GongReconstructPointersFromReferences(stage *Stage, instance *Score_instrument) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Virtual_instrument, stage.Virtual_instruments_reference, instance.Virtual_instrument)
	// insertion point for slice of pointers field
}

func (reference *Score_part) GongReconstructPointersFromReferences(stage *Stage, instance *Score_part) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identification, stage.Identifications_reference, instance.Identification)
	__gong__reconstructPointer(&reference.Part_name, stage.Part_names_reference, instance.Part_name)
	__gong__reconstructPointer(&reference.Part_name_display, stage.Name_displays_reference, instance.Part_name_display)
	__gong__reconstructPointer(&reference.Part_abbreviation, stage.Part_names_reference, instance.Part_abbreviation)
	__gong__reconstructPointer(&reference.Part_abbreviation_display, stage.Name_displays_reference, instance.Part_abbreviation_display)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Part_link, stage.Part_links_reference, instance.Part_link)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Score_instrument, stage.Score_instruments_reference, instance.Score_instrument)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Player, stage.Players_reference, instance.Player)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Midi_device, stage.Midi_devices_reference, instance.Midi_device)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Midi_instrument, stage.Midi_instruments_reference, instance.Midi_instrument)
}

func (reference *Score_partwise) GongReconstructPointersFromReferences(stage *Stage, instance *Score_partwise) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Work, stage.Works_reference, instance.Work)
	__gong__reconstructPointer(&reference.Identification, stage.Identifications_reference, instance.Identification)
	__gong__reconstructPointer(&reference.Defaults, stage.Defaultss_reference, instance.Defaults)
	__gong__reconstructPointer(&reference.Part_list, stage.Part_lists_reference, instance.Part_list)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Credit, stage.Credits_reference, instance.Credit)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Part, stage.A_parts_reference, instance.Part)
}

func (reference *Score_timewise) GongReconstructPointersFromReferences(stage *Stage, instance *Score_timewise) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Work, stage.Works_reference, instance.Work)
	__gong__reconstructPointer(&reference.Identification, stage.Identifications_reference, instance.Identification)
	__gong__reconstructPointer(&reference.Defaults, stage.Defaultss_reference, instance.Defaults)
	__gong__reconstructPointer(&reference.Part_list, stage.Part_lists_reference, instance.Part_list)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Credit, stage.Credits_reference, instance.Credit)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Measure, stage.A_measure_1s_reference, instance.Measure)
}

func (reference *Segno) GongReconstructPointersFromReferences(stage *Stage, instance *Segno) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Slash) GongReconstructPointersFromReferences(stage *Stage, instance *Slash) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Slide) GongReconstructPointersFromReferences(stage *Stage, instance *Slide) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Slur) GongReconstructPointersFromReferences(stage *Stage, instance *Slur) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Sound) GongReconstructPointersFromReferences(stage *Stage, instance *Sound) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Swing, stage.Swings_reference, instance.Swing)
	__gong__reconstructPointer(&reference.Offset, stage.Offsets_reference, instance.Offset)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Instrument_change, stage.Instrument_changes_reference, instance.Instrument_change)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Midi_device, stage.Midi_devices_reference, instance.Midi_device)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Midi_instrument, stage.Midi_instruments_reference, instance.Midi_instrument)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Play, stage.Plays_reference, instance.Play)
}

func (reference *Staff_details) GongReconstructPointersFromReferences(stage *Stage, instance *Staff_details) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Staff_size, stage.Staff_sizes_reference, instance.Staff_size)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Line_detail, stage.Line_details_reference, instance.Line_detail)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Staff_tuning, stage.Staff_tunings_reference, instance.Staff_tuning)
}

func (reference *Staff_divide) GongReconstructPointersFromReferences(stage *Stage, instance *Staff_divide) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Staff_layout) GongReconstructPointersFromReferences(stage *Stage, instance *Staff_layout) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Staff_size) GongReconstructPointersFromReferences(stage *Stage, instance *Staff_size) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Staff_tuning) GongReconstructPointersFromReferences(stage *Stage, instance *Staff_tuning) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Stem) GongReconstructPointersFromReferences(stage *Stage, instance *Stem) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Stick) GongReconstructPointersFromReferences(stage *Stage, instance *Stick) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *String_mute) GongReconstructPointersFromReferences(stage *Stage, instance *String_mute) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *String_type) GongReconstructPointersFromReferences(stage *Stage, instance *String_type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Strong_accent) GongReconstructPointersFromReferences(stage *Stage, instance *Strong_accent) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Style_text) GongReconstructPointersFromReferences(stage *Stage, instance *Style_text) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Supports) GongReconstructPointersFromReferences(stage *Stage, instance *Supports) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Swing) GongReconstructPointersFromReferences(stage *Stage, instance *Swing) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Sync) GongReconstructPointersFromReferences(stage *Stage, instance *Sync) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *System_dividers) GongReconstructPointersFromReferences(stage *Stage, instance *System_dividers) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Left_divider, stage.Empty_print_object_style_aligns_reference, instance.Left_divider)
	__gong__reconstructPointer(&reference.Right_divider, stage.Empty_print_object_style_aligns_reference, instance.Right_divider)
	// insertion point for slice of pointers field
}

func (reference *System_layout) GongReconstructPointersFromReferences(stage *Stage, instance *System_layout) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.System_margins, stage.System_marginss_reference, instance.System_margins)
	__gong__reconstructPointer(&reference.System_dividers, stage.System_dividerss_reference, instance.System_dividers)
	// insertion point for slice of pointers field
}

func (reference *System_margins) GongReconstructPointersFromReferences(stage *Stage, instance *System_margins) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tap) GongReconstructPointersFromReferences(stage *Stage, instance *Tap) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Technical) GongReconstructPointersFromReferences(stage *Stage, instance *Technical) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Up_bow, stage.Empty_placements_reference, instance.Up_bow)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Down_bow, stage.Empty_placements_reference, instance.Down_bow)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Harmonic, stage.Harmonics_reference, instance.Harmonic)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Open_string, stage.Empty_placements_reference, instance.Open_string)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Thumb_position, stage.Empty_placements_reference, instance.Thumb_position)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Fingering, stage.Fingerings_reference, instance.Fingering)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Pluck, stage.Placement_texts_reference, instance.Pluck)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Double_tongue, stage.Empty_placements_reference, instance.Double_tongue)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Triple_tongue, stage.Empty_placements_reference, instance.Triple_tongue)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Stopped, stage.Empty_placement_smufls_reference, instance.Stopped)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Snap_pizzicato, stage.Empty_placements_reference, instance.Snap_pizzicato)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Fret, stage.Frets_reference, instance.Fret)
	__gong__reconstructSliceOfPointersFromReferences(&reference.String, stage.String_types_reference, instance.String)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Hammer_on, stage.Hammer_on_pull_offs_reference, instance.Hammer_on)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Pull_off, stage.Hammer_on_pull_offs_reference, instance.Pull_off)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Bend, stage.Bends_reference, instance.Bend)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tap, stage.Taps_reference, instance.Tap)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Heel, stage.Heel_toes_reference, instance.Heel)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Toe, stage.Heel_toes_reference, instance.Toe)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Fingernails, stage.Empty_placements_reference, instance.Fingernails)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Hole, stage.Holes_reference, instance.Hole)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Arrow, stage.Arrows_reference, instance.Arrow)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Handbell, stage.Handbells_reference, instance.Handbell)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Brass_bend, stage.Empty_placements_reference, instance.Brass_bend)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Flip, stage.Empty_placements_reference, instance.Flip)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Smear, stage.Empty_placements_reference, instance.Smear)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Open, stage.Empty_placement_smufls_reference, instance.Open)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Half_muted, stage.Empty_placement_smufls_reference, instance.Half_muted)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Harmon_mute, stage.Harmon_mutes_reference, instance.Harmon_mute)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Golpe, stage.Empty_placements_reference, instance.Golpe)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Other_technical, stage.Other_placement_texts_reference, instance.Other_technical)
}

func (reference *Text_element_data) GongReconstructPointersFromReferences(stage *Stage, instance *Text_element_data) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tie) GongReconstructPointersFromReferences(stage *Stage, instance *Tie) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tied) GongReconstructPointersFromReferences(stage *Stage, instance *Tied) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Time) GongReconstructPointersFromReferences(stage *Stage, instance *Time) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Interchangeable, stage.Interchangeables_reference, instance.Interchangeable)
	// insertion point for slice of pointers field
}

func (reference *Time_modification) GongReconstructPointersFromReferences(stage *Stage, instance *Time_modification) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Timpani) GongReconstructPointersFromReferences(stage *Stage, instance *Timpani) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Transpose) GongReconstructPointersFromReferences(stage *Stage, instance *Transpose) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tremolo) GongReconstructPointersFromReferences(stage *Stage, instance *Tremolo) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tuplet) GongReconstructPointersFromReferences(stage *Stage, instance *Tuplet) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Tuplet_actual, stage.Tuplet_portions_reference, instance.Tuplet_actual)
	__gong__reconstructPointer(&reference.Tuplet_normal, stage.Tuplet_portions_reference, instance.Tuplet_normal)
	// insertion point for slice of pointers field
}

func (reference *Tuplet_dot) GongReconstructPointersFromReferences(stage *Stage, instance *Tuplet_dot) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tuplet_number) GongReconstructPointersFromReferences(stage *Stage, instance *Tuplet_number) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Tuplet_portion) GongReconstructPointersFromReferences(stage *Stage, instance *Tuplet_portion) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Tuplet_number, stage.Tuplet_numbers_reference, instance.Tuplet_number)
	__gong__reconstructPointer(&reference.Tuplet_type, stage.Tuplet_types_reference, instance.Tuplet_type)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tuplet_dot, stage.Tuplet_dots_reference, instance.Tuplet_dot)
}

func (reference *Tuplet_type) GongReconstructPointersFromReferences(stage *Stage, instance *Tuplet_type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Typed_text) GongReconstructPointersFromReferences(stage *Stage, instance *Typed_text) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Unpitched) GongReconstructPointersFromReferences(stage *Stage, instance *Unpitched) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Virtual_instrument) GongReconstructPointersFromReferences(stage *Stage, instance *Virtual_instrument) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Wait) GongReconstructPointersFromReferences(stage *Stage, instance *Wait) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Wavy_line) GongReconstructPointersFromReferences(stage *Stage, instance *Wavy_line) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Wedge) GongReconstructPointersFromReferences(stage *Stage, instance *Wedge) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Wood) GongReconstructPointersFromReferences(stage *Stage, instance *Wood) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Work) GongReconstructPointersFromReferences(stage *Stage, instance *Work) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Opus, stage.Opuss_reference, instance.Opus)
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *A_directive) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_measure) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Note, stage.Notes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Backup, stage.Backups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Forward, stage.Forwards_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Direction, stage.Directions_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Attributes, stage.Attributess_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Harmony, stage.Harmonys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Figured_bass, stage.Figured_basss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Print, stage.Prints_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sound, stage.Sounds_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Listening, stage.Listenings_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Barline, stage.Barlines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Grouping, stage.Groupings_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Link, stage.Links_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Bookmark, stage.Bookmarks_instance)
}

func (reference *A_measure_1) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Part, stage.A_part_1s_instance)
}

func (reference *A_part) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Measure, stage.A_measures_instance)
}

func (reference *A_part_1) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Note, stage.Notes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Backup, stage.Backups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Forward, stage.Forwards_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Direction, stage.Directions_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Attributes, stage.Attributess_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Harmony, stage.Harmonys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Figured_bass, stage.Figured_basss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Print, stage.Prints_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sound, stage.Sounds_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Listening, stage.Listenings_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Barline, stage.Barlines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Grouping, stage.Groupings_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Link, stage.Links_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Bookmark, stage.Bookmarks_instance)
}

func (reference *Accidental) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Accidental_mark) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Accidental_text) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Accord) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Accordion_registration) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Appearance) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Line_width, stage.Line_widths_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Note_size, stage.Note_sizes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Distance, stage.Distances_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Glyph, stage.Glyphs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Other_appearance, stage.Other_appearances_instance)
}

func (reference *Arpeggiate) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Arrow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Articulations) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Accent, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Strong_accent, stage.Strong_accents_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Staccato, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tenuto, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Detached_legato, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Staccatissimo, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Spiccato, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Scoop, stage.Empty_lines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Plop, stage.Empty_lines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Doit, stage.Empty_lines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Falloff, stage.Empty_lines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Breath_mark, stage.Breath_marks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Caesura, stage.Caesuras_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Stress, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Unstress, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Soft_accent, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Other_articulation, stage.Other_placement_texts_instance)
}

func (reference *Assess) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Attributes) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Footnote, stage.Formatted_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Level, stage.Levels_instance)
	__gong__reconstructPointerFromInstance(&reference.Part_symbol, stage.Part_symbols_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Key, stage.Keys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Time, stage.Times_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Clef, stage.Clefs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Staff_details, stage.Staff_detailss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Transpose, stage.Transposes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.For_part, stage.For_parts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Directive, stage.A_directives_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Measure_style, stage.Measure_styles_instance)
}

func (reference *Backup) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Footnote, stage.Formatted_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Level, stage.Levels_instance)
	// insertion point for slice of pointers fields
}

func (reference *Bar_style_color) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Barline) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Bar_style, stage.Bar_style_colors_instance)
	__gong__reconstructPointerFromInstance(&reference.Footnote, stage.Formatted_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Level, stage.Levels_instance)
	__gong__reconstructPointerFromInstance(&reference.Wavy_line, stage.Wavy_lines_instance)
	__gong__reconstructPointerFromInstance(&reference.Segno_1, stage.Segnos_instance)
	__gong__reconstructPointerFromInstance(&reference.Coda_1, stage.Codas_instance)
	__gong__reconstructPointerFromInstance(&reference.Fermata, stage.Fermatas_instance)
	__gong__reconstructPointerFromInstance(&reference.Ending, stage.Endings_instance)
	__gong__reconstructPointerFromInstance(&reference.Repeat, stage.Repeats_instance)
	// insertion point for slice of pointers fields
}

func (reference *Barre) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Bass) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Bass_separator, stage.Style_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Bass_step, stage.Bass_steps_instance)
	__gong__reconstructPointerFromInstance(&reference.Bass_alter, stage.Harmony_alters_instance)
	// insertion point for slice of pointers fields
}

func (reference *Bass_step) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Beam) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Beat_repeat) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Beat_unit_tied) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Beater) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Bend) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Release, stage.Releases_instance)
	__gong__reconstructPointerFromInstance(&reference.With_bar, stage.Placement_texts_instance)
	// insertion point for slice of pointers fields
}

func (reference *Bookmark) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Bracket) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Breath_mark) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Caesura) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Cancel) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Clef) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Coda) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Credit) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Credit_image, stage.Images_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Link, stage.Links_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Bookmark, stage.Bookmarks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Credit_words, stage.Formatted_text_ids_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Credit_symbol, stage.Formatted_symbol_ids_instance)
}

func (reference *Dashes) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Defaults) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Scaling, stage.Scalings_instance)
	__gong__reconstructPointerFromInstance(&reference.Page_layout, stage.Page_layouts_instance)
	__gong__reconstructPointerFromInstance(&reference.System_layout, stage.System_layouts_instance)
	__gong__reconstructPointerFromInstance(&reference.Appearance, stage.Appearances_instance)
	__gong__reconstructPointerFromInstance(&reference.Music_font, stage.Empty_fonts_instance)
	__gong__reconstructPointerFromInstance(&reference.Word_font, stage.Empty_fonts_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Staff_layout, stage.Staff_layouts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Lyric_font, stage.Lyric_fonts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Lyric_language, stage.Lyric_languages_instance)
}

func (reference *Degree) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Degree_value, stage.Degree_values_instance)
	__gong__reconstructPointerFromInstance(&reference.Degree_alter, stage.Degree_alters_instance)
	__gong__reconstructPointerFromInstance(&reference.Degree_type, stage.Degree_types_instance)
	// insertion point for slice of pointers fields
}

func (reference *Degree_alter) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Degree_type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Degree_value) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Direction) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Offset, stage.Offsets_instance)
	__gong__reconstructPointerFromInstance(&reference.Footnote, stage.Formatted_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Level, stage.Levels_instance)
	__gong__reconstructPointerFromInstance(&reference.Sound, stage.Sounds_instance)
	__gong__reconstructPointerFromInstance(&reference.Listening, stage.Listenings_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Direction_type, stage.Direction_types_instance)
}

func (reference *Direction_type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Wedge, stage.Wedges_instance)
	__gong__reconstructPointerFromInstance(&reference.Dashes, stage.Dashess_instance)
	__gong__reconstructPointerFromInstance(&reference.Bracket, stage.Brackets_instance)
	__gong__reconstructPointerFromInstance(&reference.Pedal, stage.Pedals_instance)
	__gong__reconstructPointerFromInstance(&reference.Metronome, stage.Metronomes_instance)
	__gong__reconstructPointerFromInstance(&reference.Octave_shift, stage.Octave_shifts_instance)
	__gong__reconstructPointerFromInstance(&reference.Harp_pedals, stage.Harp_pedalss_instance)
	__gong__reconstructPointerFromInstance(&reference.Damp, stage.Empty_print_style_align_ids_instance)
	__gong__reconstructPointerFromInstance(&reference.Damp_all, stage.Empty_print_style_align_ids_instance)
	__gong__reconstructPointerFromInstance(&reference.Eyeglasses, stage.Empty_print_style_align_ids_instance)
	__gong__reconstructPointerFromInstance(&reference.String_mute, stage.String_mutes_instance)
	__gong__reconstructPointerFromInstance(&reference.Scordatura, stage.Scordaturas_instance)
	__gong__reconstructPointerFromInstance(&reference.Image, stage.Images_instance)
	__gong__reconstructPointerFromInstance(&reference.Principal_voice, stage.Principal_voices_instance)
	__gong__reconstructPointerFromInstance(&reference.Accordion_registration, stage.Accordion_registrations_instance)
	__gong__reconstructPointerFromInstance(&reference.Staff_divide, stage.Staff_divides_instance)
	__gong__reconstructPointerFromInstance(&reference.Other_direction, stage.Other_directions_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Rehearsal, stage.Formatted_text_ids_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Segno, stage.Segnos_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Coda, stage.Codas_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Words, stage.Formatted_text_ids_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Symbol, stage.Formatted_symbol_ids_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Dynamics, stage.Dynamicss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Percussion, stage.Percussions_instance)
}

func (reference *Distance) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Double) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Dynamics) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Other_dynamics, stage.Other_texts_instance)
}

func (reference *Effect) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Elision) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Empty) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Empty_font) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Empty_line) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Empty_placement) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Empty_placement_smufl) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Empty_print_object_style_align) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Empty_print_style) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Empty_print_style_align) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Empty_print_style_align_id) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Empty_trill_sound) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Encoding) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Encoder, stage.Typed_texts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Supports, stage.Supportss_instance)
}

func (reference *Ending) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Extend) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Feature) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Fermata) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Figure) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Prefix, stage.Style_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Figure_number, stage.Style_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Suffix, stage.Style_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Extend, stage.Extends_instance)
	__gong__reconstructPointerFromInstance(&reference.Footnote, stage.Formatted_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Level, stage.Levels_instance)
	// insertion point for slice of pointers fields
}

func (reference *Figured_bass) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Footnote, stage.Formatted_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Level, stage.Levels_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Figure, stage.Figures_instance)
}

func (reference *Fingering) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *First_fret) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *For_part) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Part_clef, stage.Part_clefs_instance)
	__gong__reconstructPointerFromInstance(&reference.Part_transpose, stage.Part_transposes_instance)
	// insertion point for slice of pointers fields
}

func (reference *Formatted_symbol) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Formatted_symbol_id) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Formatted_text) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Formatted_text_id) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Forward) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Footnote, stage.Formatted_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Level, stage.Levels_instance)
	// insertion point for slice of pointers fields
}

func (reference *Frame) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.First_fret, stage.First_frets_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Frame_note, stage.Frame_notes_instance)
}

func (reference *Frame_note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.String, stage.String_types_instance)
	__gong__reconstructPointerFromInstance(&reference.Fret, stage.Frets_instance)
	__gong__reconstructPointerFromInstance(&reference.Fingering, stage.Fingerings_instance)
	__gong__reconstructPointerFromInstance(&reference.Barre, stage.Barres_instance)
	// insertion point for slice of pointers fields
}

func (reference *Fret) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Glass) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Glissando) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Glyph) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Grace) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Group_barline) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Group_name) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Group_symbol) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Grouping) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Feature, stage.Features_instance)
}

func (reference *Hammer_on_pull_off) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Handbell) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Harmon_closed) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Harmon_mute) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Harmon_closed, stage.Harmon_closeds_instance)
	// insertion point for slice of pointers fields
}

func (reference *Harmonic) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Harmony) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Root, stage.Roots_instance)
	__gong__reconstructPointerFromInstance(&reference.Numeral, stage.Numerals_instance)
	__gong__reconstructPointerFromInstance(&reference.Function, stage.Style_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Kind, stage.Kinds_instance)
	__gong__reconstructPointerFromInstance(&reference.Inversion, stage.Inversions_instance)
	__gong__reconstructPointerFromInstance(&reference.Bass, stage.Basss_instance)
	__gong__reconstructPointerFromInstance(&reference.Frame, stage.Frames_instance)
	__gong__reconstructPointerFromInstance(&reference.Offset, stage.Offsets_instance)
	__gong__reconstructPointerFromInstance(&reference.Footnote, stage.Formatted_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Level, stage.Levels_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Degree, stage.Degrees_instance)
}

func (reference *Harmony_alter) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Harp_pedals) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Pedal_tuning, stage.Pedal_tunings_instance)
}

func (reference *Heel_toe) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Hole) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Hole_closed, stage.Hole_closeds_instance)
	// insertion point for slice of pointers fields
}

func (reference *Hole_closed) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Horizontal_turn) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Identification) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Encoding, stage.Encodings_instance)
	__gong__reconstructPointerFromInstance(&reference.Miscellaneous, stage.Miscellaneouss_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Creator, stage.Typed_texts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Rights, stage.Typed_texts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Relation, stage.Typed_texts_instance)
}

func (reference *Image) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Instrument) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Instrument_change) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Virtual_instrument, stage.Virtual_instruments_instance)
	// insertion point for slice of pointers fields
}

func (reference *Instrument_link) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Interchangeable) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Inversion) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Key) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Cancel, stage.Cancels_instance)
	__gong__reconstructPointerFromInstance(&reference.Key_accidental, stage.Key_accidentals_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Key_octave, stage.Key_octaves_instance)
}

func (reference *Key_accidental) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Key_octave) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Kind) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Level) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Line_detail) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Line_width) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Link) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Listen) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Assess, stage.Assesss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Wait, stage.Waits_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Other_listen, stage.Other_listenings_instance)
}

func (reference *Listening) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Offset, stage.Offsets_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sync, stage.Syncs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Other_listening, stage.Other_listenings_instance)
}

func (reference *Lyric) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Extend, stage.Extends_instance)
	__gong__reconstructPointerFromInstance(&reference.Footnote, stage.Formatted_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Level, stage.Levels_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Elision, stage.Elisions_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Text, stage.Text_element_datas_instance)
}

func (reference *Lyric_font) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Lyric_language) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Measure_layout) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Measure_numbering) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Measure_repeat) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Measure_style) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Multiple_rest, stage.Multiple_rests_instance)
	__gong__reconstructPointerFromInstance(&reference.Measure_repeat, stage.Measure_repeats_instance)
	__gong__reconstructPointerFromInstance(&reference.Beat_repeat, stage.Beat_repeats_instance)
	__gong__reconstructPointerFromInstance(&reference.Slash, stage.Slashs_instance)
	// insertion point for slice of pointers fields
}

func (reference *Membrane) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Metal) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Metronome) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Per_minute, stage.Per_minutes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Beat_unit_tied, stage.Beat_unit_tieds_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Metronome_note, stage.Metronome_notes_instance)
}

func (reference *Metronome_beam) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Metronome_note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Metronome_tied, stage.Metronome_tieds_instance)
	__gong__reconstructPointerFromInstance(&reference.Metronome_tuplet, stage.Metronome_tuplets_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Metronome_beam, stage.Metronome_beams_instance)
}

func (reference *Metronome_tied) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Metronome_tuplet) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Midi_device) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Midi_instrument) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Miscellaneous) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Miscellaneous_field, stage.Miscellaneous_fields_instance)
}

func (reference *Miscellaneous_field) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Mordent) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Multiple_rest) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Name_display) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Display_text, stage.Formatted_texts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Accidental_text, stage.Accidental_texts_instance)
}

func (reference *Non_arpeggiate) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Notations) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Footnote, stage.Formatted_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Level, stage.Levels_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tied, stage.Tieds_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Slur, stage.Slurs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tuplet, stage.Tuplets_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Glissando, stage.Glissandos_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Slide, stage.Slides_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Ornaments, stage.Ornamentss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Technical, stage.Technicals_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Articulations, stage.Articulationss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Dynamics, stage.Dynamicss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Fermata, stage.Fermatas_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Arpeggiate, stage.Arpeggiates_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Non_arpeggiate, stage.Non_arpeggiates_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Accidental_mark, stage.Accidental_marks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Other_notation, stage.Other_notations_instance)
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Grace, stage.Graces_instance)
	__gong__reconstructPointerFromInstance(&reference.Pitch, stage.Pitchs_instance)
	__gong__reconstructPointerFromInstance(&reference.Unpitched, stage.Unpitcheds_instance)
	__gong__reconstructPointerFromInstance(&reference.Rest, stage.Rests_instance)
	__gong__reconstructPointerFromInstance(&reference.Tie, stage.Ties_instance)
	__gong__reconstructPointerFromInstance(&reference.Footnote, stage.Formatted_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Level, stage.Levels_instance)
	__gong__reconstructPointerFromInstance(&reference.Type, stage.Note_types_instance)
	__gong__reconstructPointerFromInstance(&reference.Accidental, stage.Accidentals_instance)
	__gong__reconstructPointerFromInstance(&reference.Time_modification, stage.Time_modifications_instance)
	__gong__reconstructPointerFromInstance(&reference.Stem, stage.Stems_instance)
	__gong__reconstructPointerFromInstance(&reference.Notehead, stage.Noteheads_instance)
	__gong__reconstructPointerFromInstance(&reference.Notehead_text, stage.Notehead_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Beam, stage.Beams_instance)
	__gong__reconstructPointerFromInstance(&reference.Play, stage.Plays_instance)
	__gong__reconstructPointerFromInstance(&reference.Listen, stage.Listens_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Instrument, stage.Instruments_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Dot, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Notations, stage.Notationss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Lyric, stage.Lyrics_instance)
}

func (reference *Note_size) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Note_type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Notehead) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Notehead_text) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Display_text, stage.Formatted_texts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Accidental_text, stage.Accidental_texts_instance)
}

func (reference *Numeral) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Numeral_root, stage.Numeral_roots_instance)
	__gong__reconstructPointerFromInstance(&reference.Numeral_alter, stage.Harmony_alters_instance)
	__gong__reconstructPointerFromInstance(&reference.Numeral_key, stage.Numeral_keys_instance)
	// insertion point for slice of pointers fields
}

func (reference *Numeral_key) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Numeral_root) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Octave_shift) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Offset) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Opus) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Ornaments) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Trill_mark, stage.Empty_trill_sounds_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Turn, stage.Horizontal_turns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Delayed_turn, stage.Horizontal_turns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Inverted_turn, stage.Horizontal_turns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Delayed_inverted_turn, stage.Horizontal_turns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Vertical_turn, stage.Empty_trill_sounds_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Inverted_vertical_turn, stage.Empty_trill_sounds_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Shake, stage.Empty_trill_sounds_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Wavy_line, stage.Wavy_lines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Mordent, stage.Mordents_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Inverted_mordent, stage.Mordents_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Schleifer, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tremolo, stage.Tremolos_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Haydn, stage.Empty_trill_sounds_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Other_ornament, stage.Other_placement_texts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Accidental_mark, stage.Accidental_marks_instance)
}

func (reference *Other_appearance) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Other_direction) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Other_listening) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Other_notation) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Other_placement_text) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Other_play) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Other_text) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Page_layout) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Page_margins, stage.Page_marginss_instance)
	// insertion point for slice of pointers fields
}

func (reference *Page_margins) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Part_clef) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Part_group) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Group_name, stage.Group_names_instance)
	__gong__reconstructPointerFromInstance(&reference.Group_name_display, stage.Name_displays_instance)
	__gong__reconstructPointerFromInstance(&reference.Group_abbreviation, stage.Group_names_instance)
	__gong__reconstructPointerFromInstance(&reference.Group_abbreviation_display, stage.Name_displays_instance)
	__gong__reconstructPointerFromInstance(&reference.Group_symbol, stage.Group_symbols_instance)
	__gong__reconstructPointerFromInstance(&reference.Group_barline, stage.Group_barlines_instance)
	__gong__reconstructPointerFromInstance(&reference.Footnote, stage.Formatted_texts_instance)
	__gong__reconstructPointerFromInstance(&reference.Level, stage.Levels_instance)
	// insertion point for slice of pointers fields
}

func (reference *Part_link) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Instrument_link, stage.Instrument_links_instance)
}

func (reference *Part_list) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Part_group, stage.Part_groups_instance)
	__gong__reconstructPointerFromInstance(&reference.Score_part, stage.Score_parts_instance)
	// insertion point for slice of pointers fields
}

func (reference *Part_name) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Part_symbol) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Part_transpose) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Pedal) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Pedal_tuning) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Per_minute) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Percussion) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Glass, stage.Glasss_instance)
	__gong__reconstructPointerFromInstance(&reference.Metal, stage.Metals_instance)
	__gong__reconstructPointerFromInstance(&reference.Wood, stage.Woods_instance)
	__gong__reconstructPointerFromInstance(&reference.Pitched, stage.Pitcheds_instance)
	__gong__reconstructPointerFromInstance(&reference.Membrane, stage.Membranes_instance)
	__gong__reconstructPointerFromInstance(&reference.Effect, stage.Effects_instance)
	__gong__reconstructPointerFromInstance(&reference.Timpani, stage.Timpanis_instance)
	__gong__reconstructPointerFromInstance(&reference.Beater, stage.Beaters_instance)
	__gong__reconstructPointerFromInstance(&reference.Stick, stage.Sticks_instance)
	__gong__reconstructPointerFromInstance(&reference.Other_percussion, stage.Other_texts_instance)
	// insertion point for slice of pointers fields
}

func (reference *Pitch) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Pitched) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Placement_text) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Play) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Other_play, stage.Other_plays_instance)
}

func (reference *Player) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Principal_voice) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Print) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Page_layout, stage.Page_layouts_instance)
	__gong__reconstructPointerFromInstance(&reference.System_layout, stage.System_layouts_instance)
	__gong__reconstructPointerFromInstance(&reference.Measure_layout, stage.Measure_layouts_instance)
	__gong__reconstructPointerFromInstance(&reference.Measure_numbering, stage.Measure_numberings_instance)
	__gong__reconstructPointerFromInstance(&reference.Part_name_display, stage.Name_displays_instance)
	__gong__reconstructPointerFromInstance(&reference.Part_abbreviation_display, stage.Name_displays_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Staff_layout, stage.Staff_layouts_instance)
}

func (reference *Release) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Repeat) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Rest) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Root) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Root_step, stage.Root_steps_instance)
	__gong__reconstructPointerFromInstance(&reference.Root_alter, stage.Harmony_alters_instance)
	// insertion point for slice of pointers fields
}

func (reference *Root_step) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Scaling) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Scordatura) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Accord, stage.Accords_instance)
}

func (reference *Score_instrument) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Virtual_instrument, stage.Virtual_instruments_instance)
	// insertion point for slice of pointers fields
}

func (reference *Score_part) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identification, stage.Identifications_instance)
	__gong__reconstructPointerFromInstance(&reference.Part_name, stage.Part_names_instance)
	__gong__reconstructPointerFromInstance(&reference.Part_name_display, stage.Name_displays_instance)
	__gong__reconstructPointerFromInstance(&reference.Part_abbreviation, stage.Part_names_instance)
	__gong__reconstructPointerFromInstance(&reference.Part_abbreviation_display, stage.Name_displays_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Part_link, stage.Part_links_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Score_instrument, stage.Score_instruments_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Player, stage.Players_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Midi_device, stage.Midi_devices_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Midi_instrument, stage.Midi_instruments_instance)
}

func (reference *Score_partwise) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Work, stage.Works_instance)
	__gong__reconstructPointerFromInstance(&reference.Identification, stage.Identifications_instance)
	__gong__reconstructPointerFromInstance(&reference.Defaults, stage.Defaultss_instance)
	__gong__reconstructPointerFromInstance(&reference.Part_list, stage.Part_lists_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Credit, stage.Credits_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Part, stage.A_parts_instance)
}

func (reference *Score_timewise) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Work, stage.Works_instance)
	__gong__reconstructPointerFromInstance(&reference.Identification, stage.Identifications_instance)
	__gong__reconstructPointerFromInstance(&reference.Defaults, stage.Defaultss_instance)
	__gong__reconstructPointerFromInstance(&reference.Part_list, stage.Part_lists_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Credit, stage.Credits_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Measure, stage.A_measure_1s_instance)
}

func (reference *Segno) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Slash) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Slide) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Slur) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Sound) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Swing, stage.Swings_instance)
	__gong__reconstructPointerFromInstance(&reference.Offset, stage.Offsets_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Instrument_change, stage.Instrument_changes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Midi_device, stage.Midi_devices_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Midi_instrument, stage.Midi_instruments_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Play, stage.Plays_instance)
}

func (reference *Staff_details) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Staff_size, stage.Staff_sizes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Line_detail, stage.Line_details_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Staff_tuning, stage.Staff_tunings_instance)
}

func (reference *Staff_divide) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Staff_layout) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Staff_size) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Staff_tuning) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Stem) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Stick) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *String_mute) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *String_type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Strong_accent) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Style_text) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Supports) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Swing) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Sync) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *System_dividers) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Left_divider, stage.Empty_print_object_style_aligns_instance)
	__gong__reconstructPointerFromInstance(&reference.Right_divider, stage.Empty_print_object_style_aligns_instance)
	// insertion point for slice of pointers fields
}

func (reference *System_layout) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.System_margins, stage.System_marginss_instance)
	__gong__reconstructPointerFromInstance(&reference.System_dividers, stage.System_dividerss_instance)
	// insertion point for slice of pointers fields
}

func (reference *System_margins) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tap) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Technical) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Up_bow, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Down_bow, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Harmonic, stage.Harmonics_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Open_string, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Thumb_position, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Fingering, stage.Fingerings_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Pluck, stage.Placement_texts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Double_tongue, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Triple_tongue, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Stopped, stage.Empty_placement_smufls_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Snap_pizzicato, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Fret, stage.Frets_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.String, stage.String_types_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Hammer_on, stage.Hammer_on_pull_offs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Pull_off, stage.Hammer_on_pull_offs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Bend, stage.Bends_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tap, stage.Taps_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Heel, stage.Heel_toes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Toe, stage.Heel_toes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Fingernails, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Hole, stage.Holes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Arrow, stage.Arrows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Handbell, stage.Handbells_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Brass_bend, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Flip, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Smear, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Open, stage.Empty_placement_smufls_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Half_muted, stage.Empty_placement_smufls_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Harmon_mute, stage.Harmon_mutes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Golpe, stage.Empty_placements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Other_technical, stage.Other_placement_texts_instance)
}

func (reference *Text_element_data) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tie) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tied) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Time) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Interchangeable, stage.Interchangeables_instance)
	// insertion point for slice of pointers fields
}

func (reference *Time_modification) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Timpani) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Transpose) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tremolo) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tuplet) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Tuplet_actual, stage.Tuplet_portions_instance)
	__gong__reconstructPointerFromInstance(&reference.Tuplet_normal, stage.Tuplet_portions_instance)
	// insertion point for slice of pointers fields
}

func (reference *Tuplet_dot) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tuplet_number) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Tuplet_portion) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Tuplet_number, stage.Tuplet_numbers_instance)
	__gong__reconstructPointerFromInstance(&reference.Tuplet_type, stage.Tuplet_types_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tuplet_dot, stage.Tuplet_dots_instance)
}

func (reference *Tuplet_type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Typed_text) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Unpitched) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Virtual_instrument) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Wait) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Wavy_line) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Wedge) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Wood) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Work) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Opus, stage.Opuss_instance)
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_directive *A_directive) GongDiff(stage *Stage, a_directiveOther *A_directive) (diffs []string) {
	// insertion point for field diffs
	if a_directive.Name != a_directiveOther.Name {
		diffs = append(diffs, a_directive.GongMarshallField(stage, "Name"))
	}
	if a_directive.Lang != a_directiveOther.Lang {
		diffs = append(diffs, a_directive.GongMarshallField(stage, "Lang"))
	}
	if a_directive.Default_x != a_directiveOther.Default_x {
		diffs = append(diffs, a_directive.GongMarshallField(stage, "Default_x"))
	}
	if a_directive.Default_y != a_directiveOther.Default_y {
		diffs = append(diffs, a_directive.GongMarshallField(stage, "Default_y"))
	}
	if a_directive.Relative_x != a_directiveOther.Relative_x {
		diffs = append(diffs, a_directive.GongMarshallField(stage, "Relative_x"))
	}
	if a_directive.Relative_y != a_directiveOther.Relative_y {
		diffs = append(diffs, a_directive.GongMarshallField(stage, "Relative_y"))
	}
	if a_directive.Font_family != a_directiveOther.Font_family {
		diffs = append(diffs, a_directive.GongMarshallField(stage, "Font_family"))
	}
	if a_directive.Font_style != a_directiveOther.Font_style {
		diffs = append(diffs, a_directive.GongMarshallField(stage, "Font_style"))
	}
	if a_directive.Font_size != a_directiveOther.Font_size {
		diffs = append(diffs, a_directive.GongMarshallField(stage, "Font_size"))
	}
	if a_directive.Font_weight != a_directiveOther.Font_weight {
		diffs = append(diffs, a_directive.GongMarshallField(stage, "Font_weight"))
	}
	if a_directive.Color != a_directiveOther.Color {
		diffs = append(diffs, a_directive.GongMarshallField(stage, "Color"))
	}
	if a_directive.EnclosedText != a_directiveOther.EnclosedText {
		diffs = append(diffs, a_directive.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_measure *A_measure) GongDiff(stage *Stage, a_measureOther *A_measure) (diffs []string) {
	// insertion point for field diffs
	if a_measure.Name != a_measureOther.Name {
		diffs = append(diffs, a_measure.GongMarshallField(stage, "Name"))
	}
	if a_measure.Number != a_measureOther.Number {
		diffs = append(diffs, a_measure.GongMarshallField(stage, "Number"))
	}
	if a_measure.Text != a_measureOther.Text {
		diffs = append(diffs, a_measure.GongMarshallField(stage, "Text"))
	}
	if a_measure.Implicit != a_measureOther.Implicit {
		diffs = append(diffs, a_measure.GongMarshallField(stage, "Implicit"))
	}
	if a_measure.Non_controlling != a_measureOther.Non_controlling {
		diffs = append(diffs, a_measure.GongMarshallField(stage, "Non_controlling"))
	}
	if a_measure.Width != a_measureOther.Width {
		diffs = append(diffs, a_measure.GongMarshallField(stage, "Width"))
	}
	if a_measure.Id != a_measureOther.Id {
		diffs = append(diffs, a_measure.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Note", a_measureOther.Note, a_measure.Note); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Backup", a_measureOther.Backup, a_measure.Backup); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Forward", a_measureOther.Forward, a_measure.Forward); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Direction", a_measureOther.Direction, a_measure.Direction); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Attributes", a_measureOther.Attributes, a_measure.Attributes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Harmony", a_measureOther.Harmony, a_measure.Harmony); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Figured_bass", a_measureOther.Figured_bass, a_measure.Figured_bass); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Print", a_measureOther.Print, a_measure.Print); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Sound", a_measureOther.Sound, a_measure.Sound); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Listening", a_measureOther.Listening, a_measure.Listening); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Barline", a_measureOther.Barline, a_measure.Barline); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Grouping", a_measureOther.Grouping, a_measure.Grouping); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Link", a_measureOther.Link, a_measure.Link); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure, "Bookmark", a_measureOther.Bookmark, a_measure.Bookmark); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_measure_1 *A_measure_1) GongDiff(stage *Stage, a_measure_1Other *A_measure_1) (diffs []string) {
	// insertion point for field diffs
	if a_measure_1.Name != a_measure_1Other.Name {
		diffs = append(diffs, a_measure_1.GongMarshallField(stage, "Name"))
	}
	if a_measure_1.Number != a_measure_1Other.Number {
		diffs = append(diffs, a_measure_1.GongMarshallField(stage, "Number"))
	}
	if a_measure_1.Text != a_measure_1Other.Text {
		diffs = append(diffs, a_measure_1.GongMarshallField(stage, "Text"))
	}
	if a_measure_1.Implicit != a_measure_1Other.Implicit {
		diffs = append(diffs, a_measure_1.GongMarshallField(stage, "Implicit"))
	}
	if a_measure_1.Non_controlling != a_measure_1Other.Non_controlling {
		diffs = append(diffs, a_measure_1.GongMarshallField(stage, "Non_controlling"))
	}
	if a_measure_1.Width != a_measure_1Other.Width {
		diffs = append(diffs, a_measure_1.GongMarshallField(stage, "Width"))
	}
	if a_measure_1.Id != a_measure_1Other.Id {
		diffs = append(diffs, a_measure_1.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, a_measure_1, "Part", a_measure_1Other.Part, a_measure_1.Part); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_part *A_part) GongDiff(stage *Stage, a_partOther *A_part) (diffs []string) {
	// insertion point for field diffs
	if a_part.Name != a_partOther.Name {
		diffs = append(diffs, a_part.GongMarshallField(stage, "Name"))
	}
	if a_part.Id != a_partOther.Id {
		diffs = append(diffs, a_part.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part, "Measure", a_partOther.Measure, a_part.Measure); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_part_1 *A_part_1) GongDiff(stage *Stage, a_part_1Other *A_part_1) (diffs []string) {
	// insertion point for field diffs
	if a_part_1.Name != a_part_1Other.Name {
		diffs = append(diffs, a_part_1.GongMarshallField(stage, "Name"))
	}
	if a_part_1.Id != a_part_1Other.Id {
		diffs = append(diffs, a_part_1.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Note", a_part_1Other.Note, a_part_1.Note); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Backup", a_part_1Other.Backup, a_part_1.Backup); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Forward", a_part_1Other.Forward, a_part_1.Forward); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Direction", a_part_1Other.Direction, a_part_1.Direction); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Attributes", a_part_1Other.Attributes, a_part_1.Attributes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Harmony", a_part_1Other.Harmony, a_part_1.Harmony); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Figured_bass", a_part_1Other.Figured_bass, a_part_1.Figured_bass); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Print", a_part_1Other.Print, a_part_1.Print); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Sound", a_part_1Other.Sound, a_part_1.Sound); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Listening", a_part_1Other.Listening, a_part_1.Listening); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Barline", a_part_1Other.Barline, a_part_1.Barline); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Grouping", a_part_1Other.Grouping, a_part_1.Grouping); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Link", a_part_1Other.Link, a_part_1.Link); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, a_part_1, "Bookmark", a_part_1Other.Bookmark, a_part_1.Bookmark); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (accidental *Accidental) GongDiff(stage *Stage, accidentalOther *Accidental) (diffs []string) {
	// insertion point for field diffs
	if accidental.Name != accidentalOther.Name {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Name"))
	}
	if accidental.Cautionary != accidentalOther.Cautionary {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Cautionary"))
	}
	if accidental.Editorial != accidentalOther.Editorial {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Editorial"))
	}
	if accidental.Smufl != accidentalOther.Smufl {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Smufl"))
	}
	if accidental.Parentheses != accidentalOther.Parentheses {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Parentheses"))
	}
	if accidental.Bracket != accidentalOther.Bracket {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Bracket"))
	}
	if accidental.Size != accidentalOther.Size {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Size"))
	}
	if accidental.Default_x != accidentalOther.Default_x {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Default_x"))
	}
	if accidental.Default_y != accidentalOther.Default_y {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Default_y"))
	}
	if accidental.Relative_x != accidentalOther.Relative_x {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Relative_x"))
	}
	if accidental.Relative_y != accidentalOther.Relative_y {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Relative_y"))
	}
	if accidental.Font_family != accidentalOther.Font_family {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Font_family"))
	}
	if accidental.Font_style != accidentalOther.Font_style {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Font_style"))
	}
	if accidental.Font_size != accidentalOther.Font_size {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Font_size"))
	}
	if accidental.Font_weight != accidentalOther.Font_weight {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Font_weight"))
	}
	if accidental.Color != accidentalOther.Color {
		diffs = append(diffs, accidental.GongMarshallField(stage, "Color"))
	}
	if accidental.EnclosedText != accidentalOther.EnclosedText {
		diffs = append(diffs, accidental.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (accidental_mark *Accidental_mark) GongDiff(stage *Stage, accidental_markOther *Accidental_mark) (diffs []string) {
	// insertion point for field diffs
	if accidental_mark.Name != accidental_markOther.Name {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Name"))
	}
	if accidental_mark.Smufl != accidental_markOther.Smufl {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Smufl"))
	}
	if accidental_mark.Parentheses != accidental_markOther.Parentheses {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Parentheses"))
	}
	if accidental_mark.Bracket != accidental_markOther.Bracket {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Bracket"))
	}
	if accidental_mark.Size != accidental_markOther.Size {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Size"))
	}
	if accidental_mark.Default_x != accidental_markOther.Default_x {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Default_x"))
	}
	if accidental_mark.Default_y != accidental_markOther.Default_y {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Default_y"))
	}
	if accidental_mark.Relative_x != accidental_markOther.Relative_x {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Relative_x"))
	}
	if accidental_mark.Relative_y != accidental_markOther.Relative_y {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Relative_y"))
	}
	if accidental_mark.Font_family != accidental_markOther.Font_family {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Font_family"))
	}
	if accidental_mark.Font_style != accidental_markOther.Font_style {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Font_style"))
	}
	if accidental_mark.Font_size != accidental_markOther.Font_size {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Font_size"))
	}
	if accidental_mark.Font_weight != accidental_markOther.Font_weight {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Font_weight"))
	}
	if accidental_mark.Color != accidental_markOther.Color {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Color"))
	}
	if accidental_mark.Placement != accidental_markOther.Placement {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Placement"))
	}
	if accidental_mark.Id != accidental_markOther.Id {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "Id"))
	}
	if accidental_mark.EnclosedText != accidental_markOther.EnclosedText {
		diffs = append(diffs, accidental_mark.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (accidental_text *Accidental_text) GongDiff(stage *Stage, accidental_textOther *Accidental_text) (diffs []string) {
	// insertion point for field diffs
	if accidental_text.Name != accidental_textOther.Name {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Name"))
	}
	if accidental_text.Smufl != accidental_textOther.Smufl {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Smufl"))
	}
	if accidental_text.Lang != accidental_textOther.Lang {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Lang"))
	}
	if accidental_text.Space != accidental_textOther.Space {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Space"))
	}
	if accidental_text.Justify != accidental_textOther.Justify {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Justify"))
	}
	if accidental_text.Default_x != accidental_textOther.Default_x {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Default_x"))
	}
	if accidental_text.Default_y != accidental_textOther.Default_y {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Default_y"))
	}
	if accidental_text.Relative_x != accidental_textOther.Relative_x {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Relative_x"))
	}
	if accidental_text.Relative_y != accidental_textOther.Relative_y {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Relative_y"))
	}
	if accidental_text.Font_family != accidental_textOther.Font_family {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Font_family"))
	}
	if accidental_text.Font_style != accidental_textOther.Font_style {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Font_style"))
	}
	if accidental_text.Font_size != accidental_textOther.Font_size {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Font_size"))
	}
	if accidental_text.Font_weight != accidental_textOther.Font_weight {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Font_weight"))
	}
	if accidental_text.Color != accidental_textOther.Color {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Color"))
	}
	if accidental_text.Halign != accidental_textOther.Halign {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Halign"))
	}
	if accidental_text.Valign != accidental_textOther.Valign {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Valign"))
	}
	if accidental_text.Underline != accidental_textOther.Underline {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Underline"))
	}
	if accidental_text.Overline != accidental_textOther.Overline {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Overline"))
	}
	if accidental_text.Line_through != accidental_textOther.Line_through {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Line_through"))
	}
	if accidental_text.Rotation != accidental_textOther.Rotation {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Rotation"))
	}
	if accidental_text.Letter_spacing != accidental_textOther.Letter_spacing {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Letter_spacing"))
	}
	if accidental_text.Line_height != accidental_textOther.Line_height {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Line_height"))
	}
	if accidental_text.Dir != accidental_textOther.Dir {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Dir"))
	}
	if accidental_text.Enclosure != accidental_textOther.Enclosure {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "Enclosure"))
	}
	if accidental_text.EnclosedText != accidental_textOther.EnclosedText {
		diffs = append(diffs, accidental_text.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (accord *Accord) GongDiff(stage *Stage, accordOther *Accord) (diffs []string) {
	// insertion point for field diffs
	if accord.Name != accordOther.Name {
		diffs = append(diffs, accord.GongMarshallField(stage, "Name"))
	}
	if accord.String != accordOther.String {
		diffs = append(diffs, accord.GongMarshallField(stage, "String"))
	}
	if accord.Tuning_step != accordOther.Tuning_step {
		diffs = append(diffs, accord.GongMarshallField(stage, "Tuning_step"))
	}
	if accord.Tuning_alter != accordOther.Tuning_alter {
		diffs = append(diffs, accord.GongMarshallField(stage, "Tuning_alter"))
	}
	if accord.Tuning_octave != accordOther.Tuning_octave {
		diffs = append(diffs, accord.GongMarshallField(stage, "Tuning_octave"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (accordion_registration *Accordion_registration) GongDiff(stage *Stage, accordion_registrationOther *Accordion_registration) (diffs []string) {
	// insertion point for field diffs
	if accordion_registration.Name != accordion_registrationOther.Name {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Name"))
	}
	if accordion_registration.Default_x != accordion_registrationOther.Default_x {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Default_x"))
	}
	if accordion_registration.Default_y != accordion_registrationOther.Default_y {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Default_y"))
	}
	if accordion_registration.Relative_x != accordion_registrationOther.Relative_x {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Relative_x"))
	}
	if accordion_registration.Relative_y != accordion_registrationOther.Relative_y {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Relative_y"))
	}
	if accordion_registration.Font_family != accordion_registrationOther.Font_family {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Font_family"))
	}
	if accordion_registration.Font_style != accordion_registrationOther.Font_style {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Font_style"))
	}
	if accordion_registration.Font_size != accordion_registrationOther.Font_size {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Font_size"))
	}
	if accordion_registration.Font_weight != accordion_registrationOther.Font_weight {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Font_weight"))
	}
	if accordion_registration.Color != accordion_registrationOther.Color {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Color"))
	}
	if accordion_registration.Halign != accordion_registrationOther.Halign {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Halign"))
	}
	if accordion_registration.Valign != accordion_registrationOther.Valign {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Valign"))
	}
	if accordion_registration.Id != accordion_registrationOther.Id {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Id"))
	}
	if accordion_registration.Accordion_high != accordion_registrationOther.Accordion_high {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Accordion_high"))
	}
	if accordion_registration.Accordion_middle != accordion_registrationOther.Accordion_middle {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Accordion_middle"))
	}
	if accordion_registration.Accordion_low != accordion_registrationOther.Accordion_low {
		diffs = append(diffs, accordion_registration.GongMarshallField(stage, "Accordion_low"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (appearance *Appearance) GongDiff(stage *Stage, appearanceOther *Appearance) (diffs []string) {
	// insertion point for field diffs
	if appearance.Name != appearanceOther.Name {
		diffs = append(diffs, appearance.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, appearance, "Line_width", appearanceOther.Line_width, appearance.Line_width); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, appearance, "Note_size", appearanceOther.Note_size, appearance.Note_size); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, appearance, "Distance", appearanceOther.Distance, appearance.Distance); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, appearance, "Glyph", appearanceOther.Glyph, appearance.Glyph); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, appearance, "Other_appearance", appearanceOther.Other_appearance, appearance.Other_appearance); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (arpeggiate *Arpeggiate) GongDiff(stage *Stage, arpeggiateOther *Arpeggiate) (diffs []string) {
	// insertion point for field diffs
	if arpeggiate.Name != arpeggiateOther.Name {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Name"))
	}
	if arpeggiate.Number != arpeggiateOther.Number {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Number"))
	}
	if arpeggiate.Direction != arpeggiateOther.Direction {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Direction"))
	}
	if arpeggiate.Unbroken != arpeggiateOther.Unbroken {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Unbroken"))
	}
	if arpeggiate.Default_x != arpeggiateOther.Default_x {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Default_x"))
	}
	if arpeggiate.Default_y != arpeggiateOther.Default_y {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Default_y"))
	}
	if arpeggiate.Relative_x != arpeggiateOther.Relative_x {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Relative_x"))
	}
	if arpeggiate.Relative_y != arpeggiateOther.Relative_y {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Relative_y"))
	}
	if arpeggiate.Placement != arpeggiateOther.Placement {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Placement"))
	}
	if arpeggiate.Color != arpeggiateOther.Color {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Color"))
	}
	if arpeggiate.Id != arpeggiateOther.Id {
		diffs = append(diffs, arpeggiate.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (arrow *Arrow) GongDiff(stage *Stage, arrowOther *Arrow) (diffs []string) {
	// insertion point for field diffs
	if arrow.Name != arrowOther.Name {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Name"))
	}
	if arrow.Default_x != arrowOther.Default_x {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Default_x"))
	}
	if arrow.Default_y != arrowOther.Default_y {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Default_y"))
	}
	if arrow.Relative_x != arrowOther.Relative_x {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Relative_x"))
	}
	if arrow.Relative_y != arrowOther.Relative_y {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Relative_y"))
	}
	if arrow.Font_family != arrowOther.Font_family {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Font_family"))
	}
	if arrow.Font_style != arrowOther.Font_style {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Font_style"))
	}
	if arrow.Font_size != arrowOther.Font_size {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Font_size"))
	}
	if arrow.Font_weight != arrowOther.Font_weight {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Font_weight"))
	}
	if arrow.Color != arrowOther.Color {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Color"))
	}
	if arrow.Placement != arrowOther.Placement {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Placement"))
	}
	if arrow.Smufl != arrowOther.Smufl {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Smufl"))
	}
	if arrow.Arrow_direction != arrowOther.Arrow_direction {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Arrow_direction"))
	}
	if arrow.Arrow_style != arrowOther.Arrow_style {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Arrow_style"))
	}
	if arrow.Arrowhead != arrowOther.Arrowhead {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Arrowhead"))
	}
	if arrow.Circular_arrow != arrowOther.Circular_arrow {
		diffs = append(diffs, arrow.GongMarshallField(stage, "Circular_arrow"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (articulations *Articulations) GongDiff(stage *Stage, articulationsOther *Articulations) (diffs []string) {
	// insertion point for field diffs
	if articulations.Name != articulationsOther.Name {
		diffs = append(diffs, articulations.GongMarshallField(stage, "Name"))
	}
	if articulations.Id != articulationsOther.Id {
		diffs = append(diffs, articulations.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Accent", articulationsOther.Accent, articulations.Accent); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Strong_accent", articulationsOther.Strong_accent, articulations.Strong_accent); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Staccato", articulationsOther.Staccato, articulations.Staccato); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Tenuto", articulationsOther.Tenuto, articulations.Tenuto); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Detached_legato", articulationsOther.Detached_legato, articulations.Detached_legato); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Staccatissimo", articulationsOther.Staccatissimo, articulations.Staccatissimo); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Spiccato", articulationsOther.Spiccato, articulations.Spiccato); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Scoop", articulationsOther.Scoop, articulations.Scoop); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Plop", articulationsOther.Plop, articulations.Plop); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Doit", articulationsOther.Doit, articulations.Doit); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Falloff", articulationsOther.Falloff, articulations.Falloff); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Breath_mark", articulationsOther.Breath_mark, articulations.Breath_mark); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Caesura", articulationsOther.Caesura, articulations.Caesura); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Stress", articulationsOther.Stress, articulations.Stress); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Unstress", articulationsOther.Unstress, articulations.Unstress); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Soft_accent", articulationsOther.Soft_accent, articulations.Soft_accent); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, articulations, "Other_articulation", articulationsOther.Other_articulation, articulations.Other_articulation); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (assess *Assess) GongDiff(stage *Stage, assessOther *Assess) (diffs []string) {
	// insertion point for field diffs
	if assess.Name != assessOther.Name {
		diffs = append(diffs, assess.GongMarshallField(stage, "Name"))
	}
	if assess.Type != assessOther.Type {
		diffs = append(diffs, assess.GongMarshallField(stage, "Type"))
	}
	if assess.Player != assessOther.Player {
		diffs = append(diffs, assess.GongMarshallField(stage, "Player"))
	}
	if assess.Time_only != assessOther.Time_only {
		diffs = append(diffs, assess.GongMarshallField(stage, "Time_only"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attributes *Attributes) GongDiff(stage *Stage, attributesOther *Attributes) (diffs []string) {
	// insertion point for field diffs
	if attributes.Name != attributesOther.Name {
		diffs = append(diffs, attributes.GongMarshallField(stage, "Name"))
	}
	if attributes.Footnote != attributesOther.Footnote {
		diffs = append(diffs, attributes.GongMarshallField(stage, "Footnote"))
	}
	if attributes.Level != attributesOther.Level {
		diffs = append(diffs, attributes.GongMarshallField(stage, "Level"))
	}
	if attributes.Divisions != attributesOther.Divisions {
		diffs = append(diffs, attributes.GongMarshallField(stage, "Divisions"))
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "Key", attributesOther.Key, attributes.Key); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "Time", attributesOther.Time, attributes.Time); ops != "" {
		diffs = append(diffs, ops)
	}
	if attributes.Staves != attributesOther.Staves {
		diffs = append(diffs, attributes.GongMarshallField(stage, "Staves"))
	}
	if attributes.Part_symbol != attributesOther.Part_symbol {
		diffs = append(diffs, attributes.GongMarshallField(stage, "Part_symbol"))
	}
	if attributes.Instruments != attributesOther.Instruments {
		diffs = append(diffs, attributes.GongMarshallField(stage, "Instruments"))
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "Clef", attributesOther.Clef, attributes.Clef); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "Staff_details", attributesOther.Staff_details, attributes.Staff_details); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "Transpose", attributesOther.Transpose, attributes.Transpose); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "For_part", attributesOther.For_part, attributes.For_part); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "Directive", attributesOther.Directive, attributes.Directive); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, attributes, "Measure_style", attributesOther.Measure_style, attributes.Measure_style); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (backup *Backup) GongDiff(stage *Stage, backupOther *Backup) (diffs []string) {
	// insertion point for field diffs
	if backup.Name != backupOther.Name {
		diffs = append(diffs, backup.GongMarshallField(stage, "Name"))
	}
	if backup.Duration != backupOther.Duration {
		diffs = append(diffs, backup.GongMarshallField(stage, "Duration"))
	}
	if backup.Footnote != backupOther.Footnote {
		diffs = append(diffs, backup.GongMarshallField(stage, "Footnote"))
	}
	if backup.Level != backupOther.Level {
		diffs = append(diffs, backup.GongMarshallField(stage, "Level"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bar_style_color *Bar_style_color) GongDiff(stage *Stage, bar_style_colorOther *Bar_style_color) (diffs []string) {
	// insertion point for field diffs
	if bar_style_color.Name != bar_style_colorOther.Name {
		diffs = append(diffs, bar_style_color.GongMarshallField(stage, "Name"))
	}
	if bar_style_color.Color != bar_style_colorOther.Color {
		diffs = append(diffs, bar_style_color.GongMarshallField(stage, "Color"))
	}
	if bar_style_color.EnclosedText != bar_style_colorOther.EnclosedText {
		diffs = append(diffs, bar_style_color.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (barline *Barline) GongDiff(stage *Stage, barlineOther *Barline) (diffs []string) {
	// insertion point for field diffs
	if barline.Name != barlineOther.Name {
		diffs = append(diffs, barline.GongMarshallField(stage, "Name"))
	}
	if barline.Location != barlineOther.Location {
		diffs = append(diffs, barline.GongMarshallField(stage, "Location"))
	}
	if barline.Segno != barlineOther.Segno {
		diffs = append(diffs, barline.GongMarshallField(stage, "Segno"))
	}
	if barline.Coda != barlineOther.Coda {
		diffs = append(diffs, barline.GongMarshallField(stage, "Coda"))
	}
	if barline.Divisions != barlineOther.Divisions {
		diffs = append(diffs, barline.GongMarshallField(stage, "Divisions"))
	}
	if barline.Id != barlineOther.Id {
		diffs = append(diffs, barline.GongMarshallField(stage, "Id"))
	}
	if barline.Bar_style != barlineOther.Bar_style {
		diffs = append(diffs, barline.GongMarshallField(stage, "Bar_style"))
	}
	if barline.Footnote != barlineOther.Footnote {
		diffs = append(diffs, barline.GongMarshallField(stage, "Footnote"))
	}
	if barline.Level != barlineOther.Level {
		diffs = append(diffs, barline.GongMarshallField(stage, "Level"))
	}
	if barline.Wavy_line != barlineOther.Wavy_line {
		diffs = append(diffs, barline.GongMarshallField(stage, "Wavy_line"))
	}
	if barline.Segno_1 != barlineOther.Segno_1 {
		diffs = append(diffs, barline.GongMarshallField(stage, "Segno_1"))
	}
	if barline.Coda_1 != barlineOther.Coda_1 {
		diffs = append(diffs, barline.GongMarshallField(stage, "Coda_1"))
	}
	if barline.Fermata != barlineOther.Fermata {
		diffs = append(diffs, barline.GongMarshallField(stage, "Fermata"))
	}
	if barline.Ending != barlineOther.Ending {
		diffs = append(diffs, barline.GongMarshallField(stage, "Ending"))
	}
	if barline.Repeat != barlineOther.Repeat {
		diffs = append(diffs, barline.GongMarshallField(stage, "Repeat"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (barre *Barre) GongDiff(stage *Stage, barreOther *Barre) (diffs []string) {
	// insertion point for field diffs
	if barre.Name != barreOther.Name {
		diffs = append(diffs, barre.GongMarshallField(stage, "Name"))
	}
	if barre.Type != barreOther.Type {
		diffs = append(diffs, barre.GongMarshallField(stage, "Type"))
	}
	if barre.Color != barreOther.Color {
		diffs = append(diffs, barre.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bass *Bass) GongDiff(stage *Stage, bassOther *Bass) (diffs []string) {
	// insertion point for field diffs
	if bass.Name != bassOther.Name {
		diffs = append(diffs, bass.GongMarshallField(stage, "Name"))
	}
	if bass.Arrangement != bassOther.Arrangement {
		diffs = append(diffs, bass.GongMarshallField(stage, "Arrangement"))
	}
	if bass.Bass_separator != bassOther.Bass_separator {
		diffs = append(diffs, bass.GongMarshallField(stage, "Bass_separator"))
	}
	if bass.Bass_step != bassOther.Bass_step {
		diffs = append(diffs, bass.GongMarshallField(stage, "Bass_step"))
	}
	if bass.Bass_alter != bassOther.Bass_alter {
		diffs = append(diffs, bass.GongMarshallField(stage, "Bass_alter"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bass_step *Bass_step) GongDiff(stage *Stage, bass_stepOther *Bass_step) (diffs []string) {
	// insertion point for field diffs
	if bass_step.Name != bass_stepOther.Name {
		diffs = append(diffs, bass_step.GongMarshallField(stage, "Name"))
	}
	if bass_step.Text != bass_stepOther.Text {
		diffs = append(diffs, bass_step.GongMarshallField(stage, "Text"))
	}
	if bass_step.Default_x != bass_stepOther.Default_x {
		diffs = append(diffs, bass_step.GongMarshallField(stage, "Default_x"))
	}
	if bass_step.Default_y != bass_stepOther.Default_y {
		diffs = append(diffs, bass_step.GongMarshallField(stage, "Default_y"))
	}
	if bass_step.Relative_x != bass_stepOther.Relative_x {
		diffs = append(diffs, bass_step.GongMarshallField(stage, "Relative_x"))
	}
	if bass_step.Relative_y != bass_stepOther.Relative_y {
		diffs = append(diffs, bass_step.GongMarshallField(stage, "Relative_y"))
	}
	if bass_step.Font_family != bass_stepOther.Font_family {
		diffs = append(diffs, bass_step.GongMarshallField(stage, "Font_family"))
	}
	if bass_step.Font_style != bass_stepOther.Font_style {
		diffs = append(diffs, bass_step.GongMarshallField(stage, "Font_style"))
	}
	if bass_step.Font_size != bass_stepOther.Font_size {
		diffs = append(diffs, bass_step.GongMarshallField(stage, "Font_size"))
	}
	if bass_step.Font_weight != bass_stepOther.Font_weight {
		diffs = append(diffs, bass_step.GongMarshallField(stage, "Font_weight"))
	}
	if bass_step.Color != bass_stepOther.Color {
		diffs = append(diffs, bass_step.GongMarshallField(stage, "Color"))
	}
	if bass_step.EnclosedText != bass_stepOther.EnclosedText {
		diffs = append(diffs, bass_step.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (beam *Beam) GongDiff(stage *Stage, beamOther *Beam) (diffs []string) {
	// insertion point for field diffs
	if beam.Name != beamOther.Name {
		diffs = append(diffs, beam.GongMarshallField(stage, "Name"))
	}
	if beam.Number != beamOther.Number {
		diffs = append(diffs, beam.GongMarshallField(stage, "Number"))
	}
	if beam.Repeater != beamOther.Repeater {
		diffs = append(diffs, beam.GongMarshallField(stage, "Repeater"))
	}
	if beam.Fan != beamOther.Fan {
		diffs = append(diffs, beam.GongMarshallField(stage, "Fan"))
	}
	if beam.Color != beamOther.Color {
		diffs = append(diffs, beam.GongMarshallField(stage, "Color"))
	}
	if beam.Id != beamOther.Id {
		diffs = append(diffs, beam.GongMarshallField(stage, "Id"))
	}
	if beam.EnclosedText != beamOther.EnclosedText {
		diffs = append(diffs, beam.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (beat_repeat *Beat_repeat) GongDiff(stage *Stage, beat_repeatOther *Beat_repeat) (diffs []string) {
	// insertion point for field diffs
	if beat_repeat.Name != beat_repeatOther.Name {
		diffs = append(diffs, beat_repeat.GongMarshallField(stage, "Name"))
	}
	if beat_repeat.Type != beat_repeatOther.Type {
		diffs = append(diffs, beat_repeat.GongMarshallField(stage, "Type"))
	}
	if beat_repeat.Slashes != beat_repeatOther.Slashes {
		diffs = append(diffs, beat_repeat.GongMarshallField(stage, "Slashes"))
	}
	if beat_repeat.Use_dots != beat_repeatOther.Use_dots {
		diffs = append(diffs, beat_repeat.GongMarshallField(stage, "Use_dots"))
	}
	if beat_repeat.Slash_type != beat_repeatOther.Slash_type {
		diffs = append(diffs, beat_repeat.GongMarshallField(stage, "Slash_type"))
	}
	if beat_repeat.Slash_dot != beat_repeatOther.Slash_dot {
		diffs = append(diffs, beat_repeat.GongMarshallField(stage, "Slash_dot"))
	}
	if beat_repeat.Except_voice != beat_repeatOther.Except_voice {
		diffs = append(diffs, beat_repeat.GongMarshallField(stage, "Except_voice"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (beat_unit_tied *Beat_unit_tied) GongDiff(stage *Stage, beat_unit_tiedOther *Beat_unit_tied) (diffs []string) {
	// insertion point for field diffs
	if beat_unit_tied.Name != beat_unit_tiedOther.Name {
		diffs = append(diffs, beat_unit_tied.GongMarshallField(stage, "Name"))
	}
	if beat_unit_tied.Beat_unit != beat_unit_tiedOther.Beat_unit {
		diffs = append(diffs, beat_unit_tied.GongMarshallField(stage, "Beat_unit"))
	}
	if beat_unit_tied.Beat_unit_dot != beat_unit_tiedOther.Beat_unit_dot {
		diffs = append(diffs, beat_unit_tied.GongMarshallField(stage, "Beat_unit_dot"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (beater *Beater) GongDiff(stage *Stage, beaterOther *Beater) (diffs []string) {
	// insertion point for field diffs
	if beater.Name != beaterOther.Name {
		diffs = append(diffs, beater.GongMarshallField(stage, "Name"))
	}
	if beater.Tip != beaterOther.Tip {
		diffs = append(diffs, beater.GongMarshallField(stage, "Tip"))
	}
	if beater.EnclosedText != beaterOther.EnclosedText {
		diffs = append(diffs, beater.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bend *Bend) GongDiff(stage *Stage, bendOther *Bend) (diffs []string) {
	// insertion point for field diffs
	if bend.Name != bendOther.Name {
		diffs = append(diffs, bend.GongMarshallField(stage, "Name"))
	}
	if bend.Shape != bendOther.Shape {
		diffs = append(diffs, bend.GongMarshallField(stage, "Shape"))
	}
	if bend.Default_x != bendOther.Default_x {
		diffs = append(diffs, bend.GongMarshallField(stage, "Default_x"))
	}
	if bend.Default_y != bendOther.Default_y {
		diffs = append(diffs, bend.GongMarshallField(stage, "Default_y"))
	}
	if bend.Relative_x != bendOther.Relative_x {
		diffs = append(diffs, bend.GongMarshallField(stage, "Relative_x"))
	}
	if bend.Relative_y != bendOther.Relative_y {
		diffs = append(diffs, bend.GongMarshallField(stage, "Relative_y"))
	}
	if bend.Font_family != bendOther.Font_family {
		diffs = append(diffs, bend.GongMarshallField(stage, "Font_family"))
	}
	if bend.Font_style != bendOther.Font_style {
		diffs = append(diffs, bend.GongMarshallField(stage, "Font_style"))
	}
	if bend.Font_size != bendOther.Font_size {
		diffs = append(diffs, bend.GongMarshallField(stage, "Font_size"))
	}
	if bend.Font_weight != bendOther.Font_weight {
		diffs = append(diffs, bend.GongMarshallField(stage, "Font_weight"))
	}
	if bend.Color != bendOther.Color {
		diffs = append(diffs, bend.GongMarshallField(stage, "Color"))
	}
	if bend.Accelerate != bendOther.Accelerate {
		diffs = append(diffs, bend.GongMarshallField(stage, "Accelerate"))
	}
	if bend.Beats != bendOther.Beats {
		diffs = append(diffs, bend.GongMarshallField(stage, "Beats"))
	}
	if bend.First_beat != bendOther.First_beat {
		diffs = append(diffs, bend.GongMarshallField(stage, "First_beat"))
	}
	if bend.Last_beat != bendOther.Last_beat {
		diffs = append(diffs, bend.GongMarshallField(stage, "Last_beat"))
	}
	if bend.Bend_alter != bendOther.Bend_alter {
		diffs = append(diffs, bend.GongMarshallField(stage, "Bend_alter"))
	}
	if bend.Pre_bend != bendOther.Pre_bend {
		diffs = append(diffs, bend.GongMarshallField(stage, "Pre_bend"))
	}
	if bend.Release != bendOther.Release {
		diffs = append(diffs, bend.GongMarshallField(stage, "Release"))
	}
	if bend.With_bar != bendOther.With_bar {
		diffs = append(diffs, bend.GongMarshallField(stage, "With_bar"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bookmark *Bookmark) GongDiff(stage *Stage, bookmarkOther *Bookmark) (diffs []string) {
	// insertion point for field diffs
	if bookmark.Name != bookmarkOther.Name {
		diffs = append(diffs, bookmark.GongMarshallField(stage, "Name"))
	}
	if bookmark.Id != bookmarkOther.Id {
		diffs = append(diffs, bookmark.GongMarshallField(stage, "Id"))
	}
	if bookmark.NameXSD != bookmarkOther.NameXSD {
		diffs = append(diffs, bookmark.GongMarshallField(stage, "NameXSD"))
	}
	if bookmark.Element != bookmarkOther.Element {
		diffs = append(diffs, bookmark.GongMarshallField(stage, "Element"))
	}
	if bookmark.Position != bookmarkOther.Position {
		diffs = append(diffs, bookmark.GongMarshallField(stage, "Position"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bracket *Bracket) GongDiff(stage *Stage, bracketOther *Bracket) (diffs []string) {
	// insertion point for field diffs
	if bracket.Name != bracketOther.Name {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Name"))
	}
	if bracket.Type != bracketOther.Type {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Type"))
	}
	if bracket.Number != bracketOther.Number {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Number"))
	}
	if bracket.Line_end != bracketOther.Line_end {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Line_end"))
	}
	if bracket.End_length != bracketOther.End_length {
		diffs = append(diffs, bracket.GongMarshallField(stage, "End_length"))
	}
	if bracket.Line_type != bracketOther.Line_type {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Line_type"))
	}
	if bracket.Dash_length != bracketOther.Dash_length {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Dash_length"))
	}
	if bracket.Space_length != bracketOther.Space_length {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Space_length"))
	}
	if bracket.Default_x != bracketOther.Default_x {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Default_x"))
	}
	if bracket.Default_y != bracketOther.Default_y {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Default_y"))
	}
	if bracket.Relative_x != bracketOther.Relative_x {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Relative_x"))
	}
	if bracket.Relative_y != bracketOther.Relative_y {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Relative_y"))
	}
	if bracket.Color != bracketOther.Color {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Color"))
	}
	if bracket.Id != bracketOther.Id {
		diffs = append(diffs, bracket.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (breath_mark *Breath_mark) GongDiff(stage *Stage, breath_markOther *Breath_mark) (diffs []string) {
	// insertion point for field diffs
	if breath_mark.Name != breath_markOther.Name {
		diffs = append(diffs, breath_mark.GongMarshallField(stage, "Name"))
	}
	if breath_mark.Default_x != breath_markOther.Default_x {
		diffs = append(diffs, breath_mark.GongMarshallField(stage, "Default_x"))
	}
	if breath_mark.Default_y != breath_markOther.Default_y {
		diffs = append(diffs, breath_mark.GongMarshallField(stage, "Default_y"))
	}
	if breath_mark.Relative_x != breath_markOther.Relative_x {
		diffs = append(diffs, breath_mark.GongMarshallField(stage, "Relative_x"))
	}
	if breath_mark.Relative_y != breath_markOther.Relative_y {
		diffs = append(diffs, breath_mark.GongMarshallField(stage, "Relative_y"))
	}
	if breath_mark.Font_family != breath_markOther.Font_family {
		diffs = append(diffs, breath_mark.GongMarshallField(stage, "Font_family"))
	}
	if breath_mark.Font_style != breath_markOther.Font_style {
		diffs = append(diffs, breath_mark.GongMarshallField(stage, "Font_style"))
	}
	if breath_mark.Font_size != breath_markOther.Font_size {
		diffs = append(diffs, breath_mark.GongMarshallField(stage, "Font_size"))
	}
	if breath_mark.Font_weight != breath_markOther.Font_weight {
		diffs = append(diffs, breath_mark.GongMarshallField(stage, "Font_weight"))
	}
	if breath_mark.Color != breath_markOther.Color {
		diffs = append(diffs, breath_mark.GongMarshallField(stage, "Color"))
	}
	if breath_mark.Placement != breath_markOther.Placement {
		diffs = append(diffs, breath_mark.GongMarshallField(stage, "Placement"))
	}
	if breath_mark.EnclosedText != breath_markOther.EnclosedText {
		diffs = append(diffs, breath_mark.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (caesura *Caesura) GongDiff(stage *Stage, caesuraOther *Caesura) (diffs []string) {
	// insertion point for field diffs
	if caesura.Name != caesuraOther.Name {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Name"))
	}
	if caesura.Default_x != caesuraOther.Default_x {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Default_x"))
	}
	if caesura.Default_y != caesuraOther.Default_y {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Default_y"))
	}
	if caesura.Relative_x != caesuraOther.Relative_x {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Relative_x"))
	}
	if caesura.Relative_y != caesuraOther.Relative_y {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Relative_y"))
	}
	if caesura.Font_family != caesuraOther.Font_family {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Font_family"))
	}
	if caesura.Font_style != caesuraOther.Font_style {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Font_style"))
	}
	if caesura.Font_size != caesuraOther.Font_size {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Font_size"))
	}
	if caesura.Font_weight != caesuraOther.Font_weight {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Font_weight"))
	}
	if caesura.Color != caesuraOther.Color {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Color"))
	}
	if caesura.Placement != caesuraOther.Placement {
		diffs = append(diffs, caesura.GongMarshallField(stage, "Placement"))
	}
	if caesura.EnclosedText != caesuraOther.EnclosedText {
		diffs = append(diffs, caesura.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (cancel *Cancel) GongDiff(stage *Stage, cancelOther *Cancel) (diffs []string) {
	// insertion point for field diffs
	if cancel.Name != cancelOther.Name {
		diffs = append(diffs, cancel.GongMarshallField(stage, "Name"))
	}
	if cancel.Location != cancelOther.Location {
		diffs = append(diffs, cancel.GongMarshallField(stage, "Location"))
	}
	if cancel.EnclosedText != cancelOther.EnclosedText {
		diffs = append(diffs, cancel.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (clef *Clef) GongDiff(stage *Stage, clefOther *Clef) (diffs []string) {
	// insertion point for field diffs
	if clef.Name != clefOther.Name {
		diffs = append(diffs, clef.GongMarshallField(stage, "Name"))
	}
	if clef.Number != clefOther.Number {
		diffs = append(diffs, clef.GongMarshallField(stage, "Number"))
	}
	if clef.Additional != clefOther.Additional {
		diffs = append(diffs, clef.GongMarshallField(stage, "Additional"))
	}
	if clef.Size != clefOther.Size {
		diffs = append(diffs, clef.GongMarshallField(stage, "Size"))
	}
	if clef.After_barline != clefOther.After_barline {
		diffs = append(diffs, clef.GongMarshallField(stage, "After_barline"))
	}
	if clef.Default_x != clefOther.Default_x {
		diffs = append(diffs, clef.GongMarshallField(stage, "Default_x"))
	}
	if clef.Default_y != clefOther.Default_y {
		diffs = append(diffs, clef.GongMarshallField(stage, "Default_y"))
	}
	if clef.Relative_x != clefOther.Relative_x {
		diffs = append(diffs, clef.GongMarshallField(stage, "Relative_x"))
	}
	if clef.Relative_y != clefOther.Relative_y {
		diffs = append(diffs, clef.GongMarshallField(stage, "Relative_y"))
	}
	if clef.Font_family != clefOther.Font_family {
		diffs = append(diffs, clef.GongMarshallField(stage, "Font_family"))
	}
	if clef.Font_style != clefOther.Font_style {
		diffs = append(diffs, clef.GongMarshallField(stage, "Font_style"))
	}
	if clef.Font_size != clefOther.Font_size {
		diffs = append(diffs, clef.GongMarshallField(stage, "Font_size"))
	}
	if clef.Font_weight != clefOther.Font_weight {
		diffs = append(diffs, clef.GongMarshallField(stage, "Font_weight"))
	}
	if clef.Color != clefOther.Color {
		diffs = append(diffs, clef.GongMarshallField(stage, "Color"))
	}
	if clef.Print_object != clefOther.Print_object {
		diffs = append(diffs, clef.GongMarshallField(stage, "Print_object"))
	}
	if clef.Id != clefOther.Id {
		diffs = append(diffs, clef.GongMarshallField(stage, "Id"))
	}
	if clef.Sign != clefOther.Sign {
		diffs = append(diffs, clef.GongMarshallField(stage, "Sign"))
	}
	if clef.Line != clefOther.Line {
		diffs = append(diffs, clef.GongMarshallField(stage, "Line"))
	}
	if clef.Clef_octave_change != clefOther.Clef_octave_change {
		diffs = append(diffs, clef.GongMarshallField(stage, "Clef_octave_change"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coda *Coda) GongDiff(stage *Stage, codaOther *Coda) (diffs []string) {
	// insertion point for field diffs
	if coda.Name != codaOther.Name {
		diffs = append(diffs, coda.GongMarshallField(stage, "Name"))
	}
	if coda.Smufl != codaOther.Smufl {
		diffs = append(diffs, coda.GongMarshallField(stage, "Smufl"))
	}
	if coda.Default_x != codaOther.Default_x {
		diffs = append(diffs, coda.GongMarshallField(stage, "Default_x"))
	}
	if coda.Default_y != codaOther.Default_y {
		diffs = append(diffs, coda.GongMarshallField(stage, "Default_y"))
	}
	if coda.Relative_x != codaOther.Relative_x {
		diffs = append(diffs, coda.GongMarshallField(stage, "Relative_x"))
	}
	if coda.Relative_y != codaOther.Relative_y {
		diffs = append(diffs, coda.GongMarshallField(stage, "Relative_y"))
	}
	if coda.Font_family != codaOther.Font_family {
		diffs = append(diffs, coda.GongMarshallField(stage, "Font_family"))
	}
	if coda.Font_style != codaOther.Font_style {
		diffs = append(diffs, coda.GongMarshallField(stage, "Font_style"))
	}
	if coda.Font_size != codaOther.Font_size {
		diffs = append(diffs, coda.GongMarshallField(stage, "Font_size"))
	}
	if coda.Font_weight != codaOther.Font_weight {
		diffs = append(diffs, coda.GongMarshallField(stage, "Font_weight"))
	}
	if coda.Color != codaOther.Color {
		diffs = append(diffs, coda.GongMarshallField(stage, "Color"))
	}
	if coda.Halign != codaOther.Halign {
		diffs = append(diffs, coda.GongMarshallField(stage, "Halign"))
	}
	if coda.Valign != codaOther.Valign {
		diffs = append(diffs, coda.GongMarshallField(stage, "Valign"))
	}
	if coda.Id != codaOther.Id {
		diffs = append(diffs, coda.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (credit *Credit) GongDiff(stage *Stage, creditOther *Credit) (diffs []string) {
	// insertion point for field diffs
	if credit.Name != creditOther.Name {
		diffs = append(diffs, credit.GongMarshallField(stage, "Name"))
	}
	if credit.Page != creditOther.Page {
		diffs = append(diffs, credit.GongMarshallField(stage, "Page"))
	}
	if credit.Id != creditOther.Id {
		diffs = append(diffs, credit.GongMarshallField(stage, "Id"))
	}
	if credit.Credit_type != creditOther.Credit_type {
		diffs = append(diffs, credit.GongMarshallField(stage, "Credit_type"))
	}
	if credit.Credit_image != creditOther.Credit_image {
		diffs = append(diffs, credit.GongMarshallField(stage, "Credit_image"))
	}
	if ops := __gong__diffSliceOfPointers(stage, credit, "Link", creditOther.Link, credit.Link); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, credit, "Bookmark", creditOther.Bookmark, credit.Bookmark); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, credit, "Credit_words", creditOther.Credit_words, credit.Credit_words); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, credit, "Credit_symbol", creditOther.Credit_symbol, credit.Credit_symbol); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dashes *Dashes) GongDiff(stage *Stage, dashesOther *Dashes) (diffs []string) {
	// insertion point for field diffs
	if dashes.Name != dashesOther.Name {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Name"))
	}
	if dashes.Type != dashesOther.Type {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Type"))
	}
	if dashes.Number != dashesOther.Number {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Number"))
	}
	if dashes.Dash_length != dashesOther.Dash_length {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Dash_length"))
	}
	if dashes.Space_length != dashesOther.Space_length {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Space_length"))
	}
	if dashes.Default_x != dashesOther.Default_x {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Default_x"))
	}
	if dashes.Default_y != dashesOther.Default_y {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Default_y"))
	}
	if dashes.Relative_x != dashesOther.Relative_x {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Relative_x"))
	}
	if dashes.Relative_y != dashesOther.Relative_y {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Relative_y"))
	}
	if dashes.Color != dashesOther.Color {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Color"))
	}
	if dashes.Id != dashesOther.Id {
		diffs = append(diffs, dashes.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (defaults *Defaults) GongDiff(stage *Stage, defaultsOther *Defaults) (diffs []string) {
	// insertion point for field diffs
	if defaults.Name != defaultsOther.Name {
		diffs = append(diffs, defaults.GongMarshallField(stage, "Name"))
	}
	if defaults.Scaling != defaultsOther.Scaling {
		diffs = append(diffs, defaults.GongMarshallField(stage, "Scaling"))
	}
	if defaults.Concert_score != defaultsOther.Concert_score {
		diffs = append(diffs, defaults.GongMarshallField(stage, "Concert_score"))
	}
	if defaults.Page_layout != defaultsOther.Page_layout {
		diffs = append(diffs, defaults.GongMarshallField(stage, "Page_layout"))
	}
	if defaults.System_layout != defaultsOther.System_layout {
		diffs = append(diffs, defaults.GongMarshallField(stage, "System_layout"))
	}
	if ops := __gong__diffSliceOfPointers(stage, defaults, "Staff_layout", defaultsOther.Staff_layout, defaults.Staff_layout); ops != "" {
		diffs = append(diffs, ops)
	}
	if defaults.Appearance != defaultsOther.Appearance {
		diffs = append(diffs, defaults.GongMarshallField(stage, "Appearance"))
	}
	if defaults.Music_font != defaultsOther.Music_font {
		diffs = append(diffs, defaults.GongMarshallField(stage, "Music_font"))
	}
	if defaults.Word_font != defaultsOther.Word_font {
		diffs = append(diffs, defaults.GongMarshallField(stage, "Word_font"))
	}
	if ops := __gong__diffSliceOfPointers(stage, defaults, "Lyric_font", defaultsOther.Lyric_font, defaults.Lyric_font); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, defaults, "Lyric_language", defaultsOther.Lyric_language, defaults.Lyric_language); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (degree *Degree) GongDiff(stage *Stage, degreeOther *Degree) (diffs []string) {
	// insertion point for field diffs
	if degree.Name != degreeOther.Name {
		diffs = append(diffs, degree.GongMarshallField(stage, "Name"))
	}
	if degree.Print_object != degreeOther.Print_object {
		diffs = append(diffs, degree.GongMarshallField(stage, "Print_object"))
	}
	if degree.Degree_value != degreeOther.Degree_value {
		diffs = append(diffs, degree.GongMarshallField(stage, "Degree_value"))
	}
	if degree.Degree_alter != degreeOther.Degree_alter {
		diffs = append(diffs, degree.GongMarshallField(stage, "Degree_alter"))
	}
	if degree.Degree_type != degreeOther.Degree_type {
		diffs = append(diffs, degree.GongMarshallField(stage, "Degree_type"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (degree_alter *Degree_alter) GongDiff(stage *Stage, degree_alterOther *Degree_alter) (diffs []string) {
	// insertion point for field diffs
	if degree_alter.Name != degree_alterOther.Name {
		diffs = append(diffs, degree_alter.GongMarshallField(stage, "Name"))
	}
	if degree_alter.Plus_minus != degree_alterOther.Plus_minus {
		diffs = append(diffs, degree_alter.GongMarshallField(stage, "Plus_minus"))
	}
	if degree_alter.Default_x != degree_alterOther.Default_x {
		diffs = append(diffs, degree_alter.GongMarshallField(stage, "Default_x"))
	}
	if degree_alter.Default_y != degree_alterOther.Default_y {
		diffs = append(diffs, degree_alter.GongMarshallField(stage, "Default_y"))
	}
	if degree_alter.Relative_x != degree_alterOther.Relative_x {
		diffs = append(diffs, degree_alter.GongMarshallField(stage, "Relative_x"))
	}
	if degree_alter.Relative_y != degree_alterOther.Relative_y {
		diffs = append(diffs, degree_alter.GongMarshallField(stage, "Relative_y"))
	}
	if degree_alter.Font_family != degree_alterOther.Font_family {
		diffs = append(diffs, degree_alter.GongMarshallField(stage, "Font_family"))
	}
	if degree_alter.Font_style != degree_alterOther.Font_style {
		diffs = append(diffs, degree_alter.GongMarshallField(stage, "Font_style"))
	}
	if degree_alter.Font_size != degree_alterOther.Font_size {
		diffs = append(diffs, degree_alter.GongMarshallField(stage, "Font_size"))
	}
	if degree_alter.Font_weight != degree_alterOther.Font_weight {
		diffs = append(diffs, degree_alter.GongMarshallField(stage, "Font_weight"))
	}
	if degree_alter.Color != degree_alterOther.Color {
		diffs = append(diffs, degree_alter.GongMarshallField(stage, "Color"))
	}
	if degree_alter.EnclosedText != degree_alterOther.EnclosedText {
		diffs = append(diffs, degree_alter.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (degree_type *Degree_type) GongDiff(stage *Stage, degree_typeOther *Degree_type) (diffs []string) {
	// insertion point for field diffs
	if degree_type.Name != degree_typeOther.Name {
		diffs = append(diffs, degree_type.GongMarshallField(stage, "Name"))
	}
	if degree_type.Text != degree_typeOther.Text {
		diffs = append(diffs, degree_type.GongMarshallField(stage, "Text"))
	}
	if degree_type.Default_x != degree_typeOther.Default_x {
		diffs = append(diffs, degree_type.GongMarshallField(stage, "Default_x"))
	}
	if degree_type.Default_y != degree_typeOther.Default_y {
		diffs = append(diffs, degree_type.GongMarshallField(stage, "Default_y"))
	}
	if degree_type.Relative_x != degree_typeOther.Relative_x {
		diffs = append(diffs, degree_type.GongMarshallField(stage, "Relative_x"))
	}
	if degree_type.Relative_y != degree_typeOther.Relative_y {
		diffs = append(diffs, degree_type.GongMarshallField(stage, "Relative_y"))
	}
	if degree_type.Font_family != degree_typeOther.Font_family {
		diffs = append(diffs, degree_type.GongMarshallField(stage, "Font_family"))
	}
	if degree_type.Font_style != degree_typeOther.Font_style {
		diffs = append(diffs, degree_type.GongMarshallField(stage, "Font_style"))
	}
	if degree_type.Font_size != degree_typeOther.Font_size {
		diffs = append(diffs, degree_type.GongMarshallField(stage, "Font_size"))
	}
	if degree_type.Font_weight != degree_typeOther.Font_weight {
		diffs = append(diffs, degree_type.GongMarshallField(stage, "Font_weight"))
	}
	if degree_type.Color != degree_typeOther.Color {
		diffs = append(diffs, degree_type.GongMarshallField(stage, "Color"))
	}
	if degree_type.EnclosedText != degree_typeOther.EnclosedText {
		diffs = append(diffs, degree_type.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (degree_value *Degree_value) GongDiff(stage *Stage, degree_valueOther *Degree_value) (diffs []string) {
	// insertion point for field diffs
	if degree_value.Name != degree_valueOther.Name {
		diffs = append(diffs, degree_value.GongMarshallField(stage, "Name"))
	}
	if degree_value.Symbol != degree_valueOther.Symbol {
		diffs = append(diffs, degree_value.GongMarshallField(stage, "Symbol"))
	}
	if degree_value.Text != degree_valueOther.Text {
		diffs = append(diffs, degree_value.GongMarshallField(stage, "Text"))
	}
	if degree_value.Default_x != degree_valueOther.Default_x {
		diffs = append(diffs, degree_value.GongMarshallField(stage, "Default_x"))
	}
	if degree_value.Default_y != degree_valueOther.Default_y {
		diffs = append(diffs, degree_value.GongMarshallField(stage, "Default_y"))
	}
	if degree_value.Relative_x != degree_valueOther.Relative_x {
		diffs = append(diffs, degree_value.GongMarshallField(stage, "Relative_x"))
	}
	if degree_value.Relative_y != degree_valueOther.Relative_y {
		diffs = append(diffs, degree_value.GongMarshallField(stage, "Relative_y"))
	}
	if degree_value.Font_family != degree_valueOther.Font_family {
		diffs = append(diffs, degree_value.GongMarshallField(stage, "Font_family"))
	}
	if degree_value.Font_style != degree_valueOther.Font_style {
		diffs = append(diffs, degree_value.GongMarshallField(stage, "Font_style"))
	}
	if degree_value.Font_size != degree_valueOther.Font_size {
		diffs = append(diffs, degree_value.GongMarshallField(stage, "Font_size"))
	}
	if degree_value.Font_weight != degree_valueOther.Font_weight {
		diffs = append(diffs, degree_value.GongMarshallField(stage, "Font_weight"))
	}
	if degree_value.Color != degree_valueOther.Color {
		diffs = append(diffs, degree_value.GongMarshallField(stage, "Color"))
	}
	if degree_value.EnclosedText != degree_valueOther.EnclosedText {
		diffs = append(diffs, degree_value.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (direction *Direction) GongDiff(stage *Stage, directionOther *Direction) (diffs []string) {
	// insertion point for field diffs
	if direction.Name != directionOther.Name {
		diffs = append(diffs, direction.GongMarshallField(stage, "Name"))
	}
	if direction.Placement != directionOther.Placement {
		diffs = append(diffs, direction.GongMarshallField(stage, "Placement"))
	}
	if direction.Directive != directionOther.Directive {
		diffs = append(diffs, direction.GongMarshallField(stage, "Directive"))
	}
	if direction.System != directionOther.System {
		diffs = append(diffs, direction.GongMarshallField(stage, "System"))
	}
	if direction.Id != directionOther.Id {
		diffs = append(diffs, direction.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, direction, "Direction_type", directionOther.Direction_type, direction.Direction_type); ops != "" {
		diffs = append(diffs, ops)
	}
	if direction.Offset != directionOther.Offset {
		diffs = append(diffs, direction.GongMarshallField(stage, "Offset"))
	}
	if direction.Footnote != directionOther.Footnote {
		diffs = append(diffs, direction.GongMarshallField(stage, "Footnote"))
	}
	if direction.Level != directionOther.Level {
		diffs = append(diffs, direction.GongMarshallField(stage, "Level"))
	}
	if direction.Voice != directionOther.Voice {
		diffs = append(diffs, direction.GongMarshallField(stage, "Voice"))
	}
	if direction.Staff != directionOther.Staff {
		diffs = append(diffs, direction.GongMarshallField(stage, "Staff"))
	}
	if direction.Sound != directionOther.Sound {
		diffs = append(diffs, direction.GongMarshallField(stage, "Sound"))
	}
	if direction.Listening != directionOther.Listening {
		diffs = append(diffs, direction.GongMarshallField(stage, "Listening"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (direction_type *Direction_type) GongDiff(stage *Stage, direction_typeOther *Direction_type) (diffs []string) {
	// insertion point for field diffs
	if direction_type.Name != direction_typeOther.Name {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Name"))
	}
	if direction_type.Id != direction_typeOther.Id {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, direction_type, "Rehearsal", direction_typeOther.Rehearsal, direction_type.Rehearsal); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, direction_type, "Segno", direction_typeOther.Segno, direction_type.Segno); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, direction_type, "Coda", direction_typeOther.Coda, direction_type.Coda); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, direction_type, "Words", direction_typeOther.Words, direction_type.Words); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, direction_type, "Symbol", direction_typeOther.Symbol, direction_type.Symbol); ops != "" {
		diffs = append(diffs, ops)
	}
	if direction_type.Wedge != direction_typeOther.Wedge {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Wedge"))
	}
	if ops := __gong__diffSliceOfPointers(stage, direction_type, "Dynamics", direction_typeOther.Dynamics, direction_type.Dynamics); ops != "" {
		diffs = append(diffs, ops)
	}
	if direction_type.Dashes != direction_typeOther.Dashes {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Dashes"))
	}
	if direction_type.Bracket != direction_typeOther.Bracket {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Bracket"))
	}
	if direction_type.Pedal != direction_typeOther.Pedal {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Pedal"))
	}
	if direction_type.Metronome != direction_typeOther.Metronome {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Metronome"))
	}
	if direction_type.Octave_shift != direction_typeOther.Octave_shift {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Octave_shift"))
	}
	if direction_type.Harp_pedals != direction_typeOther.Harp_pedals {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Harp_pedals"))
	}
	if direction_type.Damp != direction_typeOther.Damp {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Damp"))
	}
	if direction_type.Damp_all != direction_typeOther.Damp_all {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Damp_all"))
	}
	if direction_type.Eyeglasses != direction_typeOther.Eyeglasses {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Eyeglasses"))
	}
	if direction_type.String_mute != direction_typeOther.String_mute {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "String_mute"))
	}
	if direction_type.Scordatura != direction_typeOther.Scordatura {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Scordatura"))
	}
	if direction_type.Image != direction_typeOther.Image {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Image"))
	}
	if direction_type.Principal_voice != direction_typeOther.Principal_voice {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Principal_voice"))
	}
	if ops := __gong__diffSliceOfPointers(stage, direction_type, "Percussion", direction_typeOther.Percussion, direction_type.Percussion); ops != "" {
		diffs = append(diffs, ops)
	}
	if direction_type.Accordion_registration != direction_typeOther.Accordion_registration {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Accordion_registration"))
	}
	if direction_type.Staff_divide != direction_typeOther.Staff_divide {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Staff_divide"))
	}
	if direction_type.Other_direction != direction_typeOther.Other_direction {
		diffs = append(diffs, direction_type.GongMarshallField(stage, "Other_direction"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (distance *Distance) GongDiff(stage *Stage, distanceOther *Distance) (diffs []string) {
	// insertion point for field diffs
	if distance.Name != distanceOther.Name {
		diffs = append(diffs, distance.GongMarshallField(stage, "Name"))
	}
	if distance.Type != distanceOther.Type {
		diffs = append(diffs, distance.GongMarshallField(stage, "Type"))
	}
	if distance.EnclosedText != distanceOther.EnclosedText {
		diffs = append(diffs, distance.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (double *Double) GongDiff(stage *Stage, doubleOther *Double) (diffs []string) {
	// insertion point for field diffs
	if double.Name != doubleOther.Name {
		diffs = append(diffs, double.GongMarshallField(stage, "Name"))
	}
	if double.Above != doubleOther.Above {
		diffs = append(diffs, double.GongMarshallField(stage, "Above"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dynamics *Dynamics) GongDiff(stage *Stage, dynamicsOther *Dynamics) (diffs []string) {
	// insertion point for field diffs
	if dynamics.Name != dynamicsOther.Name {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Name"))
	}
	if dynamics.Default_x != dynamicsOther.Default_x {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Default_x"))
	}
	if dynamics.Default_y != dynamicsOther.Default_y {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Default_y"))
	}
	if dynamics.Relative_x != dynamicsOther.Relative_x {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Relative_x"))
	}
	if dynamics.Relative_y != dynamicsOther.Relative_y {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Relative_y"))
	}
	if dynamics.Font_family != dynamicsOther.Font_family {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Font_family"))
	}
	if dynamics.Font_style != dynamicsOther.Font_style {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Font_style"))
	}
	if dynamics.Font_size != dynamicsOther.Font_size {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Font_size"))
	}
	if dynamics.Font_weight != dynamicsOther.Font_weight {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Font_weight"))
	}
	if dynamics.Color != dynamicsOther.Color {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Color"))
	}
	if dynamics.Halign != dynamicsOther.Halign {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Halign"))
	}
	if dynamics.Valign != dynamicsOther.Valign {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Valign"))
	}
	if dynamics.Placement != dynamicsOther.Placement {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Placement"))
	}
	if dynamics.Underline != dynamicsOther.Underline {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Underline"))
	}
	if dynamics.Overline != dynamicsOther.Overline {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Overline"))
	}
	if dynamics.Line_through != dynamicsOther.Line_through {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Line_through"))
	}
	if dynamics.Enclosure != dynamicsOther.Enclosure {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Enclosure"))
	}
	if dynamics.Id != dynamicsOther.Id {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Id"))
	}
	if dynamics.P != dynamicsOther.P {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "P"))
	}
	if dynamics.Pp != dynamicsOther.Pp {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Pp"))
	}
	if dynamics.Ppp != dynamicsOther.Ppp {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Ppp"))
	}
	if dynamics.Pppp != dynamicsOther.Pppp {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Pppp"))
	}
	if dynamics.Ppppp != dynamicsOther.Ppppp {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Ppppp"))
	}
	if dynamics.Pppppp != dynamicsOther.Pppppp {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Pppppp"))
	}
	if dynamics.F != dynamicsOther.F {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "F"))
	}
	if dynamics.Ff != dynamicsOther.Ff {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Ff"))
	}
	if dynamics.Fff != dynamicsOther.Fff {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Fff"))
	}
	if dynamics.Ffff != dynamicsOther.Ffff {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Ffff"))
	}
	if dynamics.Fffff != dynamicsOther.Fffff {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Fffff"))
	}
	if dynamics.Ffffff != dynamicsOther.Ffffff {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Ffffff"))
	}
	if dynamics.Mp != dynamicsOther.Mp {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Mp"))
	}
	if dynamics.Mf != dynamicsOther.Mf {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Mf"))
	}
	if dynamics.Sf != dynamicsOther.Sf {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Sf"))
	}
	if dynamics.Sfp != dynamicsOther.Sfp {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Sfp"))
	}
	if dynamics.Sfpp != dynamicsOther.Sfpp {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Sfpp"))
	}
	if dynamics.Fp != dynamicsOther.Fp {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Fp"))
	}
	if dynamics.Rf != dynamicsOther.Rf {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Rf"))
	}
	if dynamics.Rfz != dynamicsOther.Rfz {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Rfz"))
	}
	if dynamics.Sfz != dynamicsOther.Sfz {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Sfz"))
	}
	if dynamics.Sffz != dynamicsOther.Sffz {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Sffz"))
	}
	if dynamics.Fz != dynamicsOther.Fz {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Fz"))
	}
	if dynamics.N != dynamicsOther.N {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "N"))
	}
	if dynamics.Pf != dynamicsOther.Pf {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Pf"))
	}
	if dynamics.Sfzp != dynamicsOther.Sfzp {
		diffs = append(diffs, dynamics.GongMarshallField(stage, "Sfzp"))
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamics, "Other_dynamics", dynamicsOther.Other_dynamics, dynamics.Other_dynamics); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (effect *Effect) GongDiff(stage *Stage, effectOther *Effect) (diffs []string) {
	// insertion point for field diffs
	if effect.Name != effectOther.Name {
		diffs = append(diffs, effect.GongMarshallField(stage, "Name"))
	}
	if effect.Smufl != effectOther.Smufl {
		diffs = append(diffs, effect.GongMarshallField(stage, "Smufl"))
	}
	if effect.EnclosedText != effectOther.EnclosedText {
		diffs = append(diffs, effect.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (elision *Elision) GongDiff(stage *Stage, elisionOther *Elision) (diffs []string) {
	// insertion point for field diffs
	if elision.Name != elisionOther.Name {
		diffs = append(diffs, elision.GongMarshallField(stage, "Name"))
	}
	if elision.Smufl != elisionOther.Smufl {
		diffs = append(diffs, elision.GongMarshallField(stage, "Smufl"))
	}
	if elision.Font_family != elisionOther.Font_family {
		diffs = append(diffs, elision.GongMarshallField(stage, "Font_family"))
	}
	if elision.Font_style != elisionOther.Font_style {
		diffs = append(diffs, elision.GongMarshallField(stage, "Font_style"))
	}
	if elision.Font_size != elisionOther.Font_size {
		diffs = append(diffs, elision.GongMarshallField(stage, "Font_size"))
	}
	if elision.Font_weight != elisionOther.Font_weight {
		diffs = append(diffs, elision.GongMarshallField(stage, "Font_weight"))
	}
	if elision.Color != elisionOther.Color {
		diffs = append(diffs, elision.GongMarshallField(stage, "Color"))
	}
	if elision.EnclosedText != elisionOther.EnclosedText {
		diffs = append(diffs, elision.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (empty *Empty) GongDiff(stage *Stage, emptyOther *Empty) (diffs []string) {
	// insertion point for field diffs
	if empty.Name != emptyOther.Name {
		diffs = append(diffs, empty.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (empty_font *Empty_font) GongDiff(stage *Stage, empty_fontOther *Empty_font) (diffs []string) {
	// insertion point for field diffs
	if empty_font.Name != empty_fontOther.Name {
		diffs = append(diffs, empty_font.GongMarshallField(stage, "Name"))
	}
	if empty_font.Font_family != empty_fontOther.Font_family {
		diffs = append(diffs, empty_font.GongMarshallField(stage, "Font_family"))
	}
	if empty_font.Font_style != empty_fontOther.Font_style {
		diffs = append(diffs, empty_font.GongMarshallField(stage, "Font_style"))
	}
	if empty_font.Font_size != empty_fontOther.Font_size {
		diffs = append(diffs, empty_font.GongMarshallField(stage, "Font_size"))
	}
	if empty_font.Font_weight != empty_fontOther.Font_weight {
		diffs = append(diffs, empty_font.GongMarshallField(stage, "Font_weight"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (empty_line *Empty_line) GongDiff(stage *Stage, empty_lineOther *Empty_line) (diffs []string) {
	// insertion point for field diffs
	if empty_line.Name != empty_lineOther.Name {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Name"))
	}
	if empty_line.Line_shape != empty_lineOther.Line_shape {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Line_shape"))
	}
	if empty_line.Line_type != empty_lineOther.Line_type {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Line_type"))
	}
	if empty_line.Line_length != empty_lineOther.Line_length {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Line_length"))
	}
	if empty_line.Dash_length != empty_lineOther.Dash_length {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Dash_length"))
	}
	if empty_line.Space_length != empty_lineOther.Space_length {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Space_length"))
	}
	if empty_line.Default_x != empty_lineOther.Default_x {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Default_x"))
	}
	if empty_line.Default_y != empty_lineOther.Default_y {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Default_y"))
	}
	if empty_line.Relative_x != empty_lineOther.Relative_x {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Relative_x"))
	}
	if empty_line.Relative_y != empty_lineOther.Relative_y {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Relative_y"))
	}
	if empty_line.Font_family != empty_lineOther.Font_family {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Font_family"))
	}
	if empty_line.Font_style != empty_lineOther.Font_style {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Font_style"))
	}
	if empty_line.Font_size != empty_lineOther.Font_size {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Font_size"))
	}
	if empty_line.Font_weight != empty_lineOther.Font_weight {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Font_weight"))
	}
	if empty_line.Color != empty_lineOther.Color {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Color"))
	}
	if empty_line.Placement != empty_lineOther.Placement {
		diffs = append(diffs, empty_line.GongMarshallField(stage, "Placement"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (empty_placement *Empty_placement) GongDiff(stage *Stage, empty_placementOther *Empty_placement) (diffs []string) {
	// insertion point for field diffs
	if empty_placement.Name != empty_placementOther.Name {
		diffs = append(diffs, empty_placement.GongMarshallField(stage, "Name"))
	}
	if empty_placement.Default_x != empty_placementOther.Default_x {
		diffs = append(diffs, empty_placement.GongMarshallField(stage, "Default_x"))
	}
	if empty_placement.Default_y != empty_placementOther.Default_y {
		diffs = append(diffs, empty_placement.GongMarshallField(stage, "Default_y"))
	}
	if empty_placement.Relative_x != empty_placementOther.Relative_x {
		diffs = append(diffs, empty_placement.GongMarshallField(stage, "Relative_x"))
	}
	if empty_placement.Relative_y != empty_placementOther.Relative_y {
		diffs = append(diffs, empty_placement.GongMarshallField(stage, "Relative_y"))
	}
	if empty_placement.Font_family != empty_placementOther.Font_family {
		diffs = append(diffs, empty_placement.GongMarshallField(stage, "Font_family"))
	}
	if empty_placement.Font_style != empty_placementOther.Font_style {
		diffs = append(diffs, empty_placement.GongMarshallField(stage, "Font_style"))
	}
	if empty_placement.Font_size != empty_placementOther.Font_size {
		diffs = append(diffs, empty_placement.GongMarshallField(stage, "Font_size"))
	}
	if empty_placement.Font_weight != empty_placementOther.Font_weight {
		diffs = append(diffs, empty_placement.GongMarshallField(stage, "Font_weight"))
	}
	if empty_placement.Color != empty_placementOther.Color {
		diffs = append(diffs, empty_placement.GongMarshallField(stage, "Color"))
	}
	if empty_placement.Placement != empty_placementOther.Placement {
		diffs = append(diffs, empty_placement.GongMarshallField(stage, "Placement"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (empty_placement_smufl *Empty_placement_smufl) GongDiff(stage *Stage, empty_placement_smuflOther *Empty_placement_smufl) (diffs []string) {
	// insertion point for field diffs
	if empty_placement_smufl.Name != empty_placement_smuflOther.Name {
		diffs = append(diffs, empty_placement_smufl.GongMarshallField(stage, "Name"))
	}
	if empty_placement_smufl.Default_x != empty_placement_smuflOther.Default_x {
		diffs = append(diffs, empty_placement_smufl.GongMarshallField(stage, "Default_x"))
	}
	if empty_placement_smufl.Default_y != empty_placement_smuflOther.Default_y {
		diffs = append(diffs, empty_placement_smufl.GongMarshallField(stage, "Default_y"))
	}
	if empty_placement_smufl.Relative_x != empty_placement_smuflOther.Relative_x {
		diffs = append(diffs, empty_placement_smufl.GongMarshallField(stage, "Relative_x"))
	}
	if empty_placement_smufl.Relative_y != empty_placement_smuflOther.Relative_y {
		diffs = append(diffs, empty_placement_smufl.GongMarshallField(stage, "Relative_y"))
	}
	if empty_placement_smufl.Font_family != empty_placement_smuflOther.Font_family {
		diffs = append(diffs, empty_placement_smufl.GongMarshallField(stage, "Font_family"))
	}
	if empty_placement_smufl.Font_style != empty_placement_smuflOther.Font_style {
		diffs = append(diffs, empty_placement_smufl.GongMarshallField(stage, "Font_style"))
	}
	if empty_placement_smufl.Font_size != empty_placement_smuflOther.Font_size {
		diffs = append(diffs, empty_placement_smufl.GongMarshallField(stage, "Font_size"))
	}
	if empty_placement_smufl.Font_weight != empty_placement_smuflOther.Font_weight {
		diffs = append(diffs, empty_placement_smufl.GongMarshallField(stage, "Font_weight"))
	}
	if empty_placement_smufl.Color != empty_placement_smuflOther.Color {
		diffs = append(diffs, empty_placement_smufl.GongMarshallField(stage, "Color"))
	}
	if empty_placement_smufl.Placement != empty_placement_smuflOther.Placement {
		diffs = append(diffs, empty_placement_smufl.GongMarshallField(stage, "Placement"))
	}
	if empty_placement_smufl.Smufl != empty_placement_smuflOther.Smufl {
		diffs = append(diffs, empty_placement_smufl.GongMarshallField(stage, "Smufl"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (empty_print_object_style_align *Empty_print_object_style_align) GongDiff(stage *Stage, empty_print_object_style_alignOther *Empty_print_object_style_align) (diffs []string) {
	// insertion point for field diffs
	if empty_print_object_style_align.Name != empty_print_object_style_alignOther.Name {
		diffs = append(diffs, empty_print_object_style_align.GongMarshallField(stage, "Name"))
	}
	if empty_print_object_style_align.Print_object != empty_print_object_style_alignOther.Print_object {
		diffs = append(diffs, empty_print_object_style_align.GongMarshallField(stage, "Print_object"))
	}
	if empty_print_object_style_align.Default_x != empty_print_object_style_alignOther.Default_x {
		diffs = append(diffs, empty_print_object_style_align.GongMarshallField(stage, "Default_x"))
	}
	if empty_print_object_style_align.Default_y != empty_print_object_style_alignOther.Default_y {
		diffs = append(diffs, empty_print_object_style_align.GongMarshallField(stage, "Default_y"))
	}
	if empty_print_object_style_align.Relative_x != empty_print_object_style_alignOther.Relative_x {
		diffs = append(diffs, empty_print_object_style_align.GongMarshallField(stage, "Relative_x"))
	}
	if empty_print_object_style_align.Relative_y != empty_print_object_style_alignOther.Relative_y {
		diffs = append(diffs, empty_print_object_style_align.GongMarshallField(stage, "Relative_y"))
	}
	if empty_print_object_style_align.Font_family != empty_print_object_style_alignOther.Font_family {
		diffs = append(diffs, empty_print_object_style_align.GongMarshallField(stage, "Font_family"))
	}
	if empty_print_object_style_align.Font_style != empty_print_object_style_alignOther.Font_style {
		diffs = append(diffs, empty_print_object_style_align.GongMarshallField(stage, "Font_style"))
	}
	if empty_print_object_style_align.Font_size != empty_print_object_style_alignOther.Font_size {
		diffs = append(diffs, empty_print_object_style_align.GongMarshallField(stage, "Font_size"))
	}
	if empty_print_object_style_align.Font_weight != empty_print_object_style_alignOther.Font_weight {
		diffs = append(diffs, empty_print_object_style_align.GongMarshallField(stage, "Font_weight"))
	}
	if empty_print_object_style_align.Color != empty_print_object_style_alignOther.Color {
		diffs = append(diffs, empty_print_object_style_align.GongMarshallField(stage, "Color"))
	}
	if empty_print_object_style_align.Halign != empty_print_object_style_alignOther.Halign {
		diffs = append(diffs, empty_print_object_style_align.GongMarshallField(stage, "Halign"))
	}
	if empty_print_object_style_align.Valign != empty_print_object_style_alignOther.Valign {
		diffs = append(diffs, empty_print_object_style_align.GongMarshallField(stage, "Valign"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (empty_print_style *Empty_print_style) GongDiff(stage *Stage, empty_print_styleOther *Empty_print_style) (diffs []string) {
	// insertion point for field diffs
	if empty_print_style.Name != empty_print_styleOther.Name {
		diffs = append(diffs, empty_print_style.GongMarshallField(stage, "Name"))
	}
	if empty_print_style.Default_x != empty_print_styleOther.Default_x {
		diffs = append(diffs, empty_print_style.GongMarshallField(stage, "Default_x"))
	}
	if empty_print_style.Default_y != empty_print_styleOther.Default_y {
		diffs = append(diffs, empty_print_style.GongMarshallField(stage, "Default_y"))
	}
	if empty_print_style.Relative_x != empty_print_styleOther.Relative_x {
		diffs = append(diffs, empty_print_style.GongMarshallField(stage, "Relative_x"))
	}
	if empty_print_style.Relative_y != empty_print_styleOther.Relative_y {
		diffs = append(diffs, empty_print_style.GongMarshallField(stage, "Relative_y"))
	}
	if empty_print_style.Font_family != empty_print_styleOther.Font_family {
		diffs = append(diffs, empty_print_style.GongMarshallField(stage, "Font_family"))
	}
	if empty_print_style.Font_style != empty_print_styleOther.Font_style {
		diffs = append(diffs, empty_print_style.GongMarshallField(stage, "Font_style"))
	}
	if empty_print_style.Font_size != empty_print_styleOther.Font_size {
		diffs = append(diffs, empty_print_style.GongMarshallField(stage, "Font_size"))
	}
	if empty_print_style.Font_weight != empty_print_styleOther.Font_weight {
		diffs = append(diffs, empty_print_style.GongMarshallField(stage, "Font_weight"))
	}
	if empty_print_style.Color != empty_print_styleOther.Color {
		diffs = append(diffs, empty_print_style.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (empty_print_style_align *Empty_print_style_align) GongDiff(stage *Stage, empty_print_style_alignOther *Empty_print_style_align) (diffs []string) {
	// insertion point for field diffs
	if empty_print_style_align.Name != empty_print_style_alignOther.Name {
		diffs = append(diffs, empty_print_style_align.GongMarshallField(stage, "Name"))
	}
	if empty_print_style_align.Default_x != empty_print_style_alignOther.Default_x {
		diffs = append(diffs, empty_print_style_align.GongMarshallField(stage, "Default_x"))
	}
	if empty_print_style_align.Default_y != empty_print_style_alignOther.Default_y {
		diffs = append(diffs, empty_print_style_align.GongMarshallField(stage, "Default_y"))
	}
	if empty_print_style_align.Relative_x != empty_print_style_alignOther.Relative_x {
		diffs = append(diffs, empty_print_style_align.GongMarshallField(stage, "Relative_x"))
	}
	if empty_print_style_align.Relative_y != empty_print_style_alignOther.Relative_y {
		diffs = append(diffs, empty_print_style_align.GongMarshallField(stage, "Relative_y"))
	}
	if empty_print_style_align.Font_family != empty_print_style_alignOther.Font_family {
		diffs = append(diffs, empty_print_style_align.GongMarshallField(stage, "Font_family"))
	}
	if empty_print_style_align.Font_style != empty_print_style_alignOther.Font_style {
		diffs = append(diffs, empty_print_style_align.GongMarshallField(stage, "Font_style"))
	}
	if empty_print_style_align.Font_size != empty_print_style_alignOther.Font_size {
		diffs = append(diffs, empty_print_style_align.GongMarshallField(stage, "Font_size"))
	}
	if empty_print_style_align.Font_weight != empty_print_style_alignOther.Font_weight {
		diffs = append(diffs, empty_print_style_align.GongMarshallField(stage, "Font_weight"))
	}
	if empty_print_style_align.Color != empty_print_style_alignOther.Color {
		diffs = append(diffs, empty_print_style_align.GongMarshallField(stage, "Color"))
	}
	if empty_print_style_align.Halign != empty_print_style_alignOther.Halign {
		diffs = append(diffs, empty_print_style_align.GongMarshallField(stage, "Halign"))
	}
	if empty_print_style_align.Valign != empty_print_style_alignOther.Valign {
		diffs = append(diffs, empty_print_style_align.GongMarshallField(stage, "Valign"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (empty_print_style_align_id *Empty_print_style_align_id) GongDiff(stage *Stage, empty_print_style_align_idOther *Empty_print_style_align_id) (diffs []string) {
	// insertion point for field diffs
	if empty_print_style_align_id.Name != empty_print_style_align_idOther.Name {
		diffs = append(diffs, empty_print_style_align_id.GongMarshallField(stage, "Name"))
	}
	if empty_print_style_align_id.Default_x != empty_print_style_align_idOther.Default_x {
		diffs = append(diffs, empty_print_style_align_id.GongMarshallField(stage, "Default_x"))
	}
	if empty_print_style_align_id.Default_y != empty_print_style_align_idOther.Default_y {
		diffs = append(diffs, empty_print_style_align_id.GongMarshallField(stage, "Default_y"))
	}
	if empty_print_style_align_id.Relative_x != empty_print_style_align_idOther.Relative_x {
		diffs = append(diffs, empty_print_style_align_id.GongMarshallField(stage, "Relative_x"))
	}
	if empty_print_style_align_id.Relative_y != empty_print_style_align_idOther.Relative_y {
		diffs = append(diffs, empty_print_style_align_id.GongMarshallField(stage, "Relative_y"))
	}
	if empty_print_style_align_id.Font_family != empty_print_style_align_idOther.Font_family {
		diffs = append(diffs, empty_print_style_align_id.GongMarshallField(stage, "Font_family"))
	}
	if empty_print_style_align_id.Font_style != empty_print_style_align_idOther.Font_style {
		diffs = append(diffs, empty_print_style_align_id.GongMarshallField(stage, "Font_style"))
	}
	if empty_print_style_align_id.Font_size != empty_print_style_align_idOther.Font_size {
		diffs = append(diffs, empty_print_style_align_id.GongMarshallField(stage, "Font_size"))
	}
	if empty_print_style_align_id.Font_weight != empty_print_style_align_idOther.Font_weight {
		diffs = append(diffs, empty_print_style_align_id.GongMarshallField(stage, "Font_weight"))
	}
	if empty_print_style_align_id.Color != empty_print_style_align_idOther.Color {
		diffs = append(diffs, empty_print_style_align_id.GongMarshallField(stage, "Color"))
	}
	if empty_print_style_align_id.Halign != empty_print_style_align_idOther.Halign {
		diffs = append(diffs, empty_print_style_align_id.GongMarshallField(stage, "Halign"))
	}
	if empty_print_style_align_id.Valign != empty_print_style_align_idOther.Valign {
		diffs = append(diffs, empty_print_style_align_id.GongMarshallField(stage, "Valign"))
	}
	if empty_print_style_align_id.Id != empty_print_style_align_idOther.Id {
		diffs = append(diffs, empty_print_style_align_id.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (empty_trill_sound *Empty_trill_sound) GongDiff(stage *Stage, empty_trill_soundOther *Empty_trill_sound) (diffs []string) {
	// insertion point for field diffs
	if empty_trill_sound.Name != empty_trill_soundOther.Name {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Name"))
	}
	if empty_trill_sound.Default_x != empty_trill_soundOther.Default_x {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Default_x"))
	}
	if empty_trill_sound.Default_y != empty_trill_soundOther.Default_y {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Default_y"))
	}
	if empty_trill_sound.Relative_x != empty_trill_soundOther.Relative_x {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Relative_x"))
	}
	if empty_trill_sound.Relative_y != empty_trill_soundOther.Relative_y {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Relative_y"))
	}
	if empty_trill_sound.Font_family != empty_trill_soundOther.Font_family {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Font_family"))
	}
	if empty_trill_sound.Font_style != empty_trill_soundOther.Font_style {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Font_style"))
	}
	if empty_trill_sound.Font_size != empty_trill_soundOther.Font_size {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Font_size"))
	}
	if empty_trill_sound.Font_weight != empty_trill_soundOther.Font_weight {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Font_weight"))
	}
	if empty_trill_sound.Color != empty_trill_soundOther.Color {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Color"))
	}
	if empty_trill_sound.Placement != empty_trill_soundOther.Placement {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Placement"))
	}
	if empty_trill_sound.Start_note != empty_trill_soundOther.Start_note {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Start_note"))
	}
	if empty_trill_sound.Trill_step != empty_trill_soundOther.Trill_step {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Trill_step"))
	}
	if empty_trill_sound.Two_note_turn != empty_trill_soundOther.Two_note_turn {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Two_note_turn"))
	}
	if empty_trill_sound.Accelerate != empty_trill_soundOther.Accelerate {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Accelerate"))
	}
	if empty_trill_sound.Beats != empty_trill_soundOther.Beats {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Beats"))
	}
	if empty_trill_sound.Second_beat != empty_trill_soundOther.Second_beat {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Second_beat"))
	}
	if empty_trill_sound.Last_beat != empty_trill_soundOther.Last_beat {
		diffs = append(diffs, empty_trill_sound.GongMarshallField(stage, "Last_beat"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (encoding *Encoding) GongDiff(stage *Stage, encodingOther *Encoding) (diffs []string) {
	// insertion point for field diffs
	if encoding.Name != encodingOther.Name {
		diffs = append(diffs, encoding.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, encoding, "Encoder", encodingOther.Encoder, encoding.Encoder); ops != "" {
		diffs = append(diffs, ops)
	}
	if encoding.Software != encodingOther.Software {
		diffs = append(diffs, encoding.GongMarshallField(stage, "Software"))
	}
	if encoding.Encoding_description != encodingOther.Encoding_description {
		diffs = append(diffs, encoding.GongMarshallField(stage, "Encoding_description"))
	}
	if ops := __gong__diffSliceOfPointers(stage, encoding, "Supports", encodingOther.Supports, encoding.Supports); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ending *Ending) GongDiff(stage *Stage, endingOther *Ending) (diffs []string) {
	// insertion point for field diffs
	if ending.Name != endingOther.Name {
		diffs = append(diffs, ending.GongMarshallField(stage, "Name"))
	}
	if ending.Number != endingOther.Number {
		diffs = append(diffs, ending.GongMarshallField(stage, "Number"))
	}
	if ending.Type != endingOther.Type {
		diffs = append(diffs, ending.GongMarshallField(stage, "Type"))
	}
	if ending.End_length != endingOther.End_length {
		diffs = append(diffs, ending.GongMarshallField(stage, "End_length"))
	}
	if ending.Text_x != endingOther.Text_x {
		diffs = append(diffs, ending.GongMarshallField(stage, "Text_x"))
	}
	if ending.Text_y != endingOther.Text_y {
		diffs = append(diffs, ending.GongMarshallField(stage, "Text_y"))
	}
	if ending.Print_object != endingOther.Print_object {
		diffs = append(diffs, ending.GongMarshallField(stage, "Print_object"))
	}
	if ending.Default_x != endingOther.Default_x {
		diffs = append(diffs, ending.GongMarshallField(stage, "Default_x"))
	}
	if ending.Default_y != endingOther.Default_y {
		diffs = append(diffs, ending.GongMarshallField(stage, "Default_y"))
	}
	if ending.Relative_x != endingOther.Relative_x {
		diffs = append(diffs, ending.GongMarshallField(stage, "Relative_x"))
	}
	if ending.Relative_y != endingOther.Relative_y {
		diffs = append(diffs, ending.GongMarshallField(stage, "Relative_y"))
	}
	if ending.Font_family != endingOther.Font_family {
		diffs = append(diffs, ending.GongMarshallField(stage, "Font_family"))
	}
	if ending.Font_style != endingOther.Font_style {
		diffs = append(diffs, ending.GongMarshallField(stage, "Font_style"))
	}
	if ending.Font_size != endingOther.Font_size {
		diffs = append(diffs, ending.GongMarshallField(stage, "Font_size"))
	}
	if ending.Font_weight != endingOther.Font_weight {
		diffs = append(diffs, ending.GongMarshallField(stage, "Font_weight"))
	}
	if ending.Color != endingOther.Color {
		diffs = append(diffs, ending.GongMarshallField(stage, "Color"))
	}
	if ending.System != endingOther.System {
		diffs = append(diffs, ending.GongMarshallField(stage, "System"))
	}
	if ending.EnclosedText != endingOther.EnclosedText {
		diffs = append(diffs, ending.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (extend *Extend) GongDiff(stage *Stage, extendOther *Extend) (diffs []string) {
	// insertion point for field diffs
	if extend.Name != extendOther.Name {
		diffs = append(diffs, extend.GongMarshallField(stage, "Name"))
	}
	if extend.Type != extendOther.Type {
		diffs = append(diffs, extend.GongMarshallField(stage, "Type"))
	}
	if extend.Default_x != extendOther.Default_x {
		diffs = append(diffs, extend.GongMarshallField(stage, "Default_x"))
	}
	if extend.Default_y != extendOther.Default_y {
		diffs = append(diffs, extend.GongMarshallField(stage, "Default_y"))
	}
	if extend.Relative_x != extendOther.Relative_x {
		diffs = append(diffs, extend.GongMarshallField(stage, "Relative_x"))
	}
	if extend.Relative_y != extendOther.Relative_y {
		diffs = append(diffs, extend.GongMarshallField(stage, "Relative_y"))
	}
	if extend.Color != extendOther.Color {
		diffs = append(diffs, extend.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (feature *Feature) GongDiff(stage *Stage, featureOther *Feature) (diffs []string) {
	// insertion point for field diffs
	if feature.Name != featureOther.Name {
		diffs = append(diffs, feature.GongMarshallField(stage, "Name"))
	}
	if feature.Type != featureOther.Type {
		diffs = append(diffs, feature.GongMarshallField(stage, "Type"))
	}
	if feature.EnclosedText != featureOther.EnclosedText {
		diffs = append(diffs, feature.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (fermata *Fermata) GongDiff(stage *Stage, fermataOther *Fermata) (diffs []string) {
	// insertion point for field diffs
	if fermata.Name != fermataOther.Name {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Name"))
	}
	if fermata.Type != fermataOther.Type {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Type"))
	}
	if fermata.Default_x != fermataOther.Default_x {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Default_x"))
	}
	if fermata.Default_y != fermataOther.Default_y {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Default_y"))
	}
	if fermata.Relative_x != fermataOther.Relative_x {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Relative_x"))
	}
	if fermata.Relative_y != fermataOther.Relative_y {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Relative_y"))
	}
	if fermata.Font_family != fermataOther.Font_family {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Font_family"))
	}
	if fermata.Font_style != fermataOther.Font_style {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Font_style"))
	}
	if fermata.Font_size != fermataOther.Font_size {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Font_size"))
	}
	if fermata.Font_weight != fermataOther.Font_weight {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Font_weight"))
	}
	if fermata.Color != fermataOther.Color {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Color"))
	}
	if fermata.Id != fermataOther.Id {
		diffs = append(diffs, fermata.GongMarshallField(stage, "Id"))
	}
	if fermata.EnclosedText != fermataOther.EnclosedText {
		diffs = append(diffs, fermata.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (figure *Figure) GongDiff(stage *Stage, figureOther *Figure) (diffs []string) {
	// insertion point for field diffs
	if figure.Name != figureOther.Name {
		diffs = append(diffs, figure.GongMarshallField(stage, "Name"))
	}
	if figure.Prefix != figureOther.Prefix {
		diffs = append(diffs, figure.GongMarshallField(stage, "Prefix"))
	}
	if figure.Figure_number != figureOther.Figure_number {
		diffs = append(diffs, figure.GongMarshallField(stage, "Figure_number"))
	}
	if figure.Suffix != figureOther.Suffix {
		diffs = append(diffs, figure.GongMarshallField(stage, "Suffix"))
	}
	if figure.Extend != figureOther.Extend {
		diffs = append(diffs, figure.GongMarshallField(stage, "Extend"))
	}
	if figure.Footnote != figureOther.Footnote {
		diffs = append(diffs, figure.GongMarshallField(stage, "Footnote"))
	}
	if figure.Level != figureOther.Level {
		diffs = append(diffs, figure.GongMarshallField(stage, "Level"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (figured_bass *Figured_bass) GongDiff(stage *Stage, figured_bassOther *Figured_bass) (diffs []string) {
	// insertion point for field diffs
	if figured_bass.Name != figured_bassOther.Name {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Name"))
	}
	if figured_bass.Parentheses != figured_bassOther.Parentheses {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Parentheses"))
	}
	if figured_bass.Default_x != figured_bassOther.Default_x {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Default_x"))
	}
	if figured_bass.Default_y != figured_bassOther.Default_y {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Default_y"))
	}
	if figured_bass.Relative_x != figured_bassOther.Relative_x {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Relative_x"))
	}
	if figured_bass.Relative_y != figured_bassOther.Relative_y {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Relative_y"))
	}
	if figured_bass.Font_family != figured_bassOther.Font_family {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Font_family"))
	}
	if figured_bass.Font_style != figured_bassOther.Font_style {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Font_style"))
	}
	if figured_bass.Font_size != figured_bassOther.Font_size {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Font_size"))
	}
	if figured_bass.Font_weight != figured_bassOther.Font_weight {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Font_weight"))
	}
	if figured_bass.Color != figured_bassOther.Color {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Color"))
	}
	if figured_bass.Halign != figured_bassOther.Halign {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Halign"))
	}
	if figured_bass.Valign != figured_bassOther.Valign {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Valign"))
	}
	if figured_bass.Placement != figured_bassOther.Placement {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Placement"))
	}
	if figured_bass.Print_dot != figured_bassOther.Print_dot {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Print_dot"))
	}
	if figured_bass.Print_lyric != figured_bassOther.Print_lyric {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Print_lyric"))
	}
	if figured_bass.Print_object != figured_bassOther.Print_object {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Print_object"))
	}
	if figured_bass.Print_spacing != figured_bassOther.Print_spacing {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Print_spacing"))
	}
	if figured_bass.Id != figured_bassOther.Id {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, figured_bass, "Figure", figured_bassOther.Figure, figured_bass.Figure); ops != "" {
		diffs = append(diffs, ops)
	}
	if figured_bass.Duration != figured_bassOther.Duration {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Duration"))
	}
	if figured_bass.Footnote != figured_bassOther.Footnote {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Footnote"))
	}
	if figured_bass.Level != figured_bassOther.Level {
		diffs = append(diffs, figured_bass.GongMarshallField(stage, "Level"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (fingering *Fingering) GongDiff(stage *Stage, fingeringOther *Fingering) (diffs []string) {
	// insertion point for field diffs
	if fingering.Name != fingeringOther.Name {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Name"))
	}
	if fingering.Substitution != fingeringOther.Substitution {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Substitution"))
	}
	if fingering.Alternate != fingeringOther.Alternate {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Alternate"))
	}
	if fingering.Default_x != fingeringOther.Default_x {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Default_x"))
	}
	if fingering.Default_y != fingeringOther.Default_y {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Default_y"))
	}
	if fingering.Relative_x != fingeringOther.Relative_x {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Relative_x"))
	}
	if fingering.Relative_y != fingeringOther.Relative_y {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Relative_y"))
	}
	if fingering.Font_family != fingeringOther.Font_family {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Font_family"))
	}
	if fingering.Font_style != fingeringOther.Font_style {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Font_style"))
	}
	if fingering.Font_size != fingeringOther.Font_size {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Font_size"))
	}
	if fingering.Font_weight != fingeringOther.Font_weight {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Font_weight"))
	}
	if fingering.Color != fingeringOther.Color {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Color"))
	}
	if fingering.Placement != fingeringOther.Placement {
		diffs = append(diffs, fingering.GongMarshallField(stage, "Placement"))
	}
	if fingering.EnclosedText != fingeringOther.EnclosedText {
		diffs = append(diffs, fingering.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (first_fret *First_fret) GongDiff(stage *Stage, first_fretOther *First_fret) (diffs []string) {
	// insertion point for field diffs
	if first_fret.Name != first_fretOther.Name {
		diffs = append(diffs, first_fret.GongMarshallField(stage, "Name"))
	}
	if first_fret.Text != first_fretOther.Text {
		diffs = append(diffs, first_fret.GongMarshallField(stage, "Text"))
	}
	if first_fret.Location != first_fretOther.Location {
		diffs = append(diffs, first_fret.GongMarshallField(stage, "Location"))
	}
	if first_fret.EnclosedText != first_fretOther.EnclosedText {
		diffs = append(diffs, first_fret.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (for_part *For_part) GongDiff(stage *Stage, for_partOther *For_part) (diffs []string) {
	// insertion point for field diffs
	if for_part.Name != for_partOther.Name {
		diffs = append(diffs, for_part.GongMarshallField(stage, "Name"))
	}
	if for_part.Number != for_partOther.Number {
		diffs = append(diffs, for_part.GongMarshallField(stage, "Number"))
	}
	if for_part.Id != for_partOther.Id {
		diffs = append(diffs, for_part.GongMarshallField(stage, "Id"))
	}
	if for_part.Part_clef != for_partOther.Part_clef {
		diffs = append(diffs, for_part.GongMarshallField(stage, "Part_clef"))
	}
	if for_part.Part_transpose != for_partOther.Part_transpose {
		diffs = append(diffs, for_part.GongMarshallField(stage, "Part_transpose"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formatted_symbol *Formatted_symbol) GongDiff(stage *Stage, formatted_symbolOther *Formatted_symbol) (diffs []string) {
	// insertion point for field diffs
	if formatted_symbol.Name != formatted_symbolOther.Name {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Name"))
	}
	if formatted_symbol.Justify != formatted_symbolOther.Justify {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Justify"))
	}
	if formatted_symbol.Default_x != formatted_symbolOther.Default_x {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Default_x"))
	}
	if formatted_symbol.Default_y != formatted_symbolOther.Default_y {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Default_y"))
	}
	if formatted_symbol.Relative_x != formatted_symbolOther.Relative_x {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Relative_x"))
	}
	if formatted_symbol.Relative_y != formatted_symbolOther.Relative_y {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Relative_y"))
	}
	if formatted_symbol.Font_family != formatted_symbolOther.Font_family {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Font_family"))
	}
	if formatted_symbol.Font_style != formatted_symbolOther.Font_style {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Font_style"))
	}
	if formatted_symbol.Font_size != formatted_symbolOther.Font_size {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Font_size"))
	}
	if formatted_symbol.Font_weight != formatted_symbolOther.Font_weight {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Font_weight"))
	}
	if formatted_symbol.Color != formatted_symbolOther.Color {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Color"))
	}
	if formatted_symbol.Halign != formatted_symbolOther.Halign {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Halign"))
	}
	if formatted_symbol.Valign != formatted_symbolOther.Valign {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Valign"))
	}
	if formatted_symbol.Underline != formatted_symbolOther.Underline {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Underline"))
	}
	if formatted_symbol.Overline != formatted_symbolOther.Overline {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Overline"))
	}
	if formatted_symbol.Line_through != formatted_symbolOther.Line_through {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Line_through"))
	}
	if formatted_symbol.Rotation != formatted_symbolOther.Rotation {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Rotation"))
	}
	if formatted_symbol.Letter_spacing != formatted_symbolOther.Letter_spacing {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Letter_spacing"))
	}
	if formatted_symbol.Line_height != formatted_symbolOther.Line_height {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Line_height"))
	}
	if formatted_symbol.Dir != formatted_symbolOther.Dir {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Dir"))
	}
	if formatted_symbol.Enclosure != formatted_symbolOther.Enclosure {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "Enclosure"))
	}
	if formatted_symbol.EnclosedText != formatted_symbolOther.EnclosedText {
		diffs = append(diffs, formatted_symbol.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formatted_symbol_id *Formatted_symbol_id) GongDiff(stage *Stage, formatted_symbol_idOther *Formatted_symbol_id) (diffs []string) {
	// insertion point for field diffs
	if formatted_symbol_id.Name != formatted_symbol_idOther.Name {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Name"))
	}
	if formatted_symbol_id.Justify != formatted_symbol_idOther.Justify {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Justify"))
	}
	if formatted_symbol_id.Default_x != formatted_symbol_idOther.Default_x {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Default_x"))
	}
	if formatted_symbol_id.Default_y != formatted_symbol_idOther.Default_y {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Default_y"))
	}
	if formatted_symbol_id.Relative_x != formatted_symbol_idOther.Relative_x {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Relative_x"))
	}
	if formatted_symbol_id.Relative_y != formatted_symbol_idOther.Relative_y {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Relative_y"))
	}
	if formatted_symbol_id.Font_family != formatted_symbol_idOther.Font_family {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Font_family"))
	}
	if formatted_symbol_id.Font_style != formatted_symbol_idOther.Font_style {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Font_style"))
	}
	if formatted_symbol_id.Font_size != formatted_symbol_idOther.Font_size {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Font_size"))
	}
	if formatted_symbol_id.Font_weight != formatted_symbol_idOther.Font_weight {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Font_weight"))
	}
	if formatted_symbol_id.Color != formatted_symbol_idOther.Color {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Color"))
	}
	if formatted_symbol_id.Halign != formatted_symbol_idOther.Halign {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Halign"))
	}
	if formatted_symbol_id.Valign != formatted_symbol_idOther.Valign {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Valign"))
	}
	if formatted_symbol_id.Underline != formatted_symbol_idOther.Underline {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Underline"))
	}
	if formatted_symbol_id.Overline != formatted_symbol_idOther.Overline {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Overline"))
	}
	if formatted_symbol_id.Line_through != formatted_symbol_idOther.Line_through {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Line_through"))
	}
	if formatted_symbol_id.Rotation != formatted_symbol_idOther.Rotation {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Rotation"))
	}
	if formatted_symbol_id.Letter_spacing != formatted_symbol_idOther.Letter_spacing {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Letter_spacing"))
	}
	if formatted_symbol_id.Line_height != formatted_symbol_idOther.Line_height {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Line_height"))
	}
	if formatted_symbol_id.Dir != formatted_symbol_idOther.Dir {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Dir"))
	}
	if formatted_symbol_id.Enclosure != formatted_symbol_idOther.Enclosure {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Enclosure"))
	}
	if formatted_symbol_id.Id != formatted_symbol_idOther.Id {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "Id"))
	}
	if formatted_symbol_id.EnclosedText != formatted_symbol_idOther.EnclosedText {
		diffs = append(diffs, formatted_symbol_id.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formatted_text *Formatted_text) GongDiff(stage *Stage, formatted_textOther *Formatted_text) (diffs []string) {
	// insertion point for field diffs
	if formatted_text.Name != formatted_textOther.Name {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Name"))
	}
	if formatted_text.Lang != formatted_textOther.Lang {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Lang"))
	}
	if formatted_text.Space != formatted_textOther.Space {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Space"))
	}
	if formatted_text.Justify != formatted_textOther.Justify {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Justify"))
	}
	if formatted_text.Default_x != formatted_textOther.Default_x {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Default_x"))
	}
	if formatted_text.Default_y != formatted_textOther.Default_y {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Default_y"))
	}
	if formatted_text.Relative_x != formatted_textOther.Relative_x {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Relative_x"))
	}
	if formatted_text.Relative_y != formatted_textOther.Relative_y {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Relative_y"))
	}
	if formatted_text.Font_family != formatted_textOther.Font_family {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Font_family"))
	}
	if formatted_text.Font_style != formatted_textOther.Font_style {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Font_style"))
	}
	if formatted_text.Font_size != formatted_textOther.Font_size {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Font_size"))
	}
	if formatted_text.Font_weight != formatted_textOther.Font_weight {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Font_weight"))
	}
	if formatted_text.Color != formatted_textOther.Color {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Color"))
	}
	if formatted_text.Halign != formatted_textOther.Halign {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Halign"))
	}
	if formatted_text.Valign != formatted_textOther.Valign {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Valign"))
	}
	if formatted_text.Underline != formatted_textOther.Underline {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Underline"))
	}
	if formatted_text.Overline != formatted_textOther.Overline {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Overline"))
	}
	if formatted_text.Line_through != formatted_textOther.Line_through {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Line_through"))
	}
	if formatted_text.Rotation != formatted_textOther.Rotation {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Rotation"))
	}
	if formatted_text.Letter_spacing != formatted_textOther.Letter_spacing {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Letter_spacing"))
	}
	if formatted_text.Line_height != formatted_textOther.Line_height {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Line_height"))
	}
	if formatted_text.Dir != formatted_textOther.Dir {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Dir"))
	}
	if formatted_text.Enclosure != formatted_textOther.Enclosure {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "Enclosure"))
	}
	if formatted_text.EnclosedText != formatted_textOther.EnclosedText {
		diffs = append(diffs, formatted_text.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formatted_text_id *Formatted_text_id) GongDiff(stage *Stage, formatted_text_idOther *Formatted_text_id) (diffs []string) {
	// insertion point for field diffs
	if formatted_text_id.Name != formatted_text_idOther.Name {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Name"))
	}
	if formatted_text_id.Lang != formatted_text_idOther.Lang {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Lang"))
	}
	if formatted_text_id.Space != formatted_text_idOther.Space {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Space"))
	}
	if formatted_text_id.Justify != formatted_text_idOther.Justify {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Justify"))
	}
	if formatted_text_id.Default_x != formatted_text_idOther.Default_x {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Default_x"))
	}
	if formatted_text_id.Default_y != formatted_text_idOther.Default_y {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Default_y"))
	}
	if formatted_text_id.Relative_x != formatted_text_idOther.Relative_x {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Relative_x"))
	}
	if formatted_text_id.Relative_y != formatted_text_idOther.Relative_y {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Relative_y"))
	}
	if formatted_text_id.Font_family != formatted_text_idOther.Font_family {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Font_family"))
	}
	if formatted_text_id.Font_style != formatted_text_idOther.Font_style {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Font_style"))
	}
	if formatted_text_id.Font_size != formatted_text_idOther.Font_size {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Font_size"))
	}
	if formatted_text_id.Font_weight != formatted_text_idOther.Font_weight {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Font_weight"))
	}
	if formatted_text_id.Color != formatted_text_idOther.Color {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Color"))
	}
	if formatted_text_id.Halign != formatted_text_idOther.Halign {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Halign"))
	}
	if formatted_text_id.Valign != formatted_text_idOther.Valign {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Valign"))
	}
	if formatted_text_id.Underline != formatted_text_idOther.Underline {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Underline"))
	}
	if formatted_text_id.Overline != formatted_text_idOther.Overline {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Overline"))
	}
	if formatted_text_id.Line_through != formatted_text_idOther.Line_through {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Line_through"))
	}
	if formatted_text_id.Rotation != formatted_text_idOther.Rotation {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Rotation"))
	}
	if formatted_text_id.Letter_spacing != formatted_text_idOther.Letter_spacing {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Letter_spacing"))
	}
	if formatted_text_id.Line_height != formatted_text_idOther.Line_height {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Line_height"))
	}
	if formatted_text_id.Dir != formatted_text_idOther.Dir {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Dir"))
	}
	if formatted_text_id.Enclosure != formatted_text_idOther.Enclosure {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Enclosure"))
	}
	if formatted_text_id.Id != formatted_text_idOther.Id {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "Id"))
	}
	if formatted_text_id.EnclosedText != formatted_text_idOther.EnclosedText {
		diffs = append(diffs, formatted_text_id.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (forward *Forward) GongDiff(stage *Stage, forwardOther *Forward) (diffs []string) {
	// insertion point for field diffs
	if forward.Name != forwardOther.Name {
		diffs = append(diffs, forward.GongMarshallField(stage, "Name"))
	}
	if forward.Duration != forwardOther.Duration {
		diffs = append(diffs, forward.GongMarshallField(stage, "Duration"))
	}
	if forward.Footnote != forwardOther.Footnote {
		diffs = append(diffs, forward.GongMarshallField(stage, "Footnote"))
	}
	if forward.Level != forwardOther.Level {
		diffs = append(diffs, forward.GongMarshallField(stage, "Level"))
	}
	if forward.Voice != forwardOther.Voice {
		diffs = append(diffs, forward.GongMarshallField(stage, "Voice"))
	}
	if forward.Staff != forwardOther.Staff {
		diffs = append(diffs, forward.GongMarshallField(stage, "Staff"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (frame *Frame) GongDiff(stage *Stage, frameOther *Frame) (diffs []string) {
	// insertion point for field diffs
	if frame.Name != frameOther.Name {
		diffs = append(diffs, frame.GongMarshallField(stage, "Name"))
	}
	if frame.Height != frameOther.Height {
		diffs = append(diffs, frame.GongMarshallField(stage, "Height"))
	}
	if frame.Width != frameOther.Width {
		diffs = append(diffs, frame.GongMarshallField(stage, "Width"))
	}
	if frame.Unplayed != frameOther.Unplayed {
		diffs = append(diffs, frame.GongMarshallField(stage, "Unplayed"))
	}
	if frame.Default_x != frameOther.Default_x {
		diffs = append(diffs, frame.GongMarshallField(stage, "Default_x"))
	}
	if frame.Default_y != frameOther.Default_y {
		diffs = append(diffs, frame.GongMarshallField(stage, "Default_y"))
	}
	if frame.Relative_x != frameOther.Relative_x {
		diffs = append(diffs, frame.GongMarshallField(stage, "Relative_x"))
	}
	if frame.Relative_y != frameOther.Relative_y {
		diffs = append(diffs, frame.GongMarshallField(stage, "Relative_y"))
	}
	if frame.Color != frameOther.Color {
		diffs = append(diffs, frame.GongMarshallField(stage, "Color"))
	}
	if frame.Halign != frameOther.Halign {
		diffs = append(diffs, frame.GongMarshallField(stage, "Halign"))
	}
	if frame.Valign != frameOther.Valign {
		diffs = append(diffs, frame.GongMarshallField(stage, "Valign"))
	}
	if frame.Id != frameOther.Id {
		diffs = append(diffs, frame.GongMarshallField(stage, "Id"))
	}
	if frame.Frame_strings != frameOther.Frame_strings {
		diffs = append(diffs, frame.GongMarshallField(stage, "Frame_strings"))
	}
	if frame.Frame_frets != frameOther.Frame_frets {
		diffs = append(diffs, frame.GongMarshallField(stage, "Frame_frets"))
	}
	if frame.First_fret != frameOther.First_fret {
		diffs = append(diffs, frame.GongMarshallField(stage, "First_fret"))
	}
	if ops := __gong__diffSliceOfPointers(stage, frame, "Frame_note", frameOther.Frame_note, frame.Frame_note); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (frame_note *Frame_note) GongDiff(stage *Stage, frame_noteOther *Frame_note) (diffs []string) {
	// insertion point for field diffs
	if frame_note.Name != frame_noteOther.Name {
		diffs = append(diffs, frame_note.GongMarshallField(stage, "Name"))
	}
	if frame_note.String != frame_noteOther.String {
		diffs = append(diffs, frame_note.GongMarshallField(stage, "String"))
	}
	if frame_note.Fret != frame_noteOther.Fret {
		diffs = append(diffs, frame_note.GongMarshallField(stage, "Fret"))
	}
	if frame_note.Fingering != frame_noteOther.Fingering {
		diffs = append(diffs, frame_note.GongMarshallField(stage, "Fingering"))
	}
	if frame_note.Barre != frame_noteOther.Barre {
		diffs = append(diffs, frame_note.GongMarshallField(stage, "Barre"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (fret *Fret) GongDiff(stage *Stage, fretOther *Fret) (diffs []string) {
	// insertion point for field diffs
	if fret.Name != fretOther.Name {
		diffs = append(diffs, fret.GongMarshallField(stage, "Name"))
	}
	if fret.Font_family != fretOther.Font_family {
		diffs = append(diffs, fret.GongMarshallField(stage, "Font_family"))
	}
	if fret.Font_style != fretOther.Font_style {
		diffs = append(diffs, fret.GongMarshallField(stage, "Font_style"))
	}
	if fret.Font_size != fretOther.Font_size {
		diffs = append(diffs, fret.GongMarshallField(stage, "Font_size"))
	}
	if fret.Font_weight != fretOther.Font_weight {
		diffs = append(diffs, fret.GongMarshallField(stage, "Font_weight"))
	}
	if fret.Color != fretOther.Color {
		diffs = append(diffs, fret.GongMarshallField(stage, "Color"))
	}
	if fret.EnclosedText != fretOther.EnclosedText {
		diffs = append(diffs, fret.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (glass *Glass) GongDiff(stage *Stage, glassOther *Glass) (diffs []string) {
	// insertion point for field diffs
	if glass.Name != glassOther.Name {
		diffs = append(diffs, glass.GongMarshallField(stage, "Name"))
	}
	if glass.Smufl != glassOther.Smufl {
		diffs = append(diffs, glass.GongMarshallField(stage, "Smufl"))
	}
	if glass.EnclosedText != glassOther.EnclosedText {
		diffs = append(diffs, glass.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (glissando *Glissando) GongDiff(stage *Stage, glissandoOther *Glissando) (diffs []string) {
	// insertion point for field diffs
	if glissando.Name != glissandoOther.Name {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Name"))
	}
	if glissando.Type != glissandoOther.Type {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Type"))
	}
	if glissando.Number != glissandoOther.Number {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Number"))
	}
	if glissando.Line_type != glissandoOther.Line_type {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Line_type"))
	}
	if glissando.Dash_length != glissandoOther.Dash_length {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Dash_length"))
	}
	if glissando.Space_length != glissandoOther.Space_length {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Space_length"))
	}
	if glissando.Default_x != glissandoOther.Default_x {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Default_x"))
	}
	if glissando.Default_y != glissandoOther.Default_y {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Default_y"))
	}
	if glissando.Relative_x != glissandoOther.Relative_x {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Relative_x"))
	}
	if glissando.Relative_y != glissandoOther.Relative_y {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Relative_y"))
	}
	if glissando.Font_family != glissandoOther.Font_family {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Font_family"))
	}
	if glissando.Font_style != glissandoOther.Font_style {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Font_style"))
	}
	if glissando.Font_size != glissandoOther.Font_size {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Font_size"))
	}
	if glissando.Font_weight != glissandoOther.Font_weight {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Font_weight"))
	}
	if glissando.Color != glissandoOther.Color {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Color"))
	}
	if glissando.Id != glissandoOther.Id {
		diffs = append(diffs, glissando.GongMarshallField(stage, "Id"))
	}
	if glissando.EnclosedText != glissandoOther.EnclosedText {
		diffs = append(diffs, glissando.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (glyph *Glyph) GongDiff(stage *Stage, glyphOther *Glyph) (diffs []string) {
	// insertion point for field diffs
	if glyph.Name != glyphOther.Name {
		diffs = append(diffs, glyph.GongMarshallField(stage, "Name"))
	}
	if glyph.Type != glyphOther.Type {
		diffs = append(diffs, glyph.GongMarshallField(stage, "Type"))
	}
	if glyph.EnclosedText != glyphOther.EnclosedText {
		diffs = append(diffs, glyph.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (grace *Grace) GongDiff(stage *Stage, graceOther *Grace) (diffs []string) {
	// insertion point for field diffs
	if grace.Name != graceOther.Name {
		diffs = append(diffs, grace.GongMarshallField(stage, "Name"))
	}
	if grace.Steal_time_previous != graceOther.Steal_time_previous {
		diffs = append(diffs, grace.GongMarshallField(stage, "Steal_time_previous"))
	}
	if grace.Steal_time_following != graceOther.Steal_time_following {
		diffs = append(diffs, grace.GongMarshallField(stage, "Steal_time_following"))
	}
	if grace.Make_time != graceOther.Make_time {
		diffs = append(diffs, grace.GongMarshallField(stage, "Make_time"))
	}
	if grace.Slash != graceOther.Slash {
		diffs = append(diffs, grace.GongMarshallField(stage, "Slash"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (group_barline *Group_barline) GongDiff(stage *Stage, group_barlineOther *Group_barline) (diffs []string) {
	// insertion point for field diffs
	if group_barline.Name != group_barlineOther.Name {
		diffs = append(diffs, group_barline.GongMarshallField(stage, "Name"))
	}
	if group_barline.Color != group_barlineOther.Color {
		diffs = append(diffs, group_barline.GongMarshallField(stage, "Color"))
	}
	if group_barline.EnclosedText != group_barlineOther.EnclosedText {
		diffs = append(diffs, group_barline.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (group_name *Group_name) GongDiff(stage *Stage, group_nameOther *Group_name) (diffs []string) {
	// insertion point for field diffs
	if group_name.Name != group_nameOther.Name {
		diffs = append(diffs, group_name.GongMarshallField(stage, "Name"))
	}
	if group_name.Default_x != group_nameOther.Default_x {
		diffs = append(diffs, group_name.GongMarshallField(stage, "Default_x"))
	}
	if group_name.Default_y != group_nameOther.Default_y {
		diffs = append(diffs, group_name.GongMarshallField(stage, "Default_y"))
	}
	if group_name.Relative_x != group_nameOther.Relative_x {
		diffs = append(diffs, group_name.GongMarshallField(stage, "Relative_x"))
	}
	if group_name.Relative_y != group_nameOther.Relative_y {
		diffs = append(diffs, group_name.GongMarshallField(stage, "Relative_y"))
	}
	if group_name.Font_family != group_nameOther.Font_family {
		diffs = append(diffs, group_name.GongMarshallField(stage, "Font_family"))
	}
	if group_name.Font_style != group_nameOther.Font_style {
		diffs = append(diffs, group_name.GongMarshallField(stage, "Font_style"))
	}
	if group_name.Font_size != group_nameOther.Font_size {
		diffs = append(diffs, group_name.GongMarshallField(stage, "Font_size"))
	}
	if group_name.Font_weight != group_nameOther.Font_weight {
		diffs = append(diffs, group_name.GongMarshallField(stage, "Font_weight"))
	}
	if group_name.Color != group_nameOther.Color {
		diffs = append(diffs, group_name.GongMarshallField(stage, "Color"))
	}
	if group_name.Justify != group_nameOther.Justify {
		diffs = append(diffs, group_name.GongMarshallField(stage, "Justify"))
	}
	if group_name.EnclosedText != group_nameOther.EnclosedText {
		diffs = append(diffs, group_name.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (group_symbol *Group_symbol) GongDiff(stage *Stage, group_symbolOther *Group_symbol) (diffs []string) {
	// insertion point for field diffs
	if group_symbol.Name != group_symbolOther.Name {
		diffs = append(diffs, group_symbol.GongMarshallField(stage, "Name"))
	}
	if group_symbol.Default_x != group_symbolOther.Default_x {
		diffs = append(diffs, group_symbol.GongMarshallField(stage, "Default_x"))
	}
	if group_symbol.Default_y != group_symbolOther.Default_y {
		diffs = append(diffs, group_symbol.GongMarshallField(stage, "Default_y"))
	}
	if group_symbol.Relative_x != group_symbolOther.Relative_x {
		diffs = append(diffs, group_symbol.GongMarshallField(stage, "Relative_x"))
	}
	if group_symbol.Relative_y != group_symbolOther.Relative_y {
		diffs = append(diffs, group_symbol.GongMarshallField(stage, "Relative_y"))
	}
	if group_symbol.Color != group_symbolOther.Color {
		diffs = append(diffs, group_symbol.GongMarshallField(stage, "Color"))
	}
	if group_symbol.EnclosedText != group_symbolOther.EnclosedText {
		diffs = append(diffs, group_symbol.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (grouping *Grouping) GongDiff(stage *Stage, groupingOther *Grouping) (diffs []string) {
	// insertion point for field diffs
	if grouping.Name != groupingOther.Name {
		diffs = append(diffs, grouping.GongMarshallField(stage, "Name"))
	}
	if grouping.Type != groupingOther.Type {
		diffs = append(diffs, grouping.GongMarshallField(stage, "Type"))
	}
	if grouping.Number != groupingOther.Number {
		diffs = append(diffs, grouping.GongMarshallField(stage, "Number"))
	}
	if grouping.Member_of != groupingOther.Member_of {
		diffs = append(diffs, grouping.GongMarshallField(stage, "Member_of"))
	}
	if grouping.Id != groupingOther.Id {
		diffs = append(diffs, grouping.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, grouping, "Feature", groupingOther.Feature, grouping.Feature); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (hammer_on_pull_off *Hammer_on_pull_off) GongDiff(stage *Stage, hammer_on_pull_offOther *Hammer_on_pull_off) (diffs []string) {
	// insertion point for field diffs
	if hammer_on_pull_off.Name != hammer_on_pull_offOther.Name {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "Name"))
	}
	if hammer_on_pull_off.Type != hammer_on_pull_offOther.Type {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "Type"))
	}
	if hammer_on_pull_off.Number != hammer_on_pull_offOther.Number {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "Number"))
	}
	if hammer_on_pull_off.Default_x != hammer_on_pull_offOther.Default_x {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "Default_x"))
	}
	if hammer_on_pull_off.Default_y != hammer_on_pull_offOther.Default_y {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "Default_y"))
	}
	if hammer_on_pull_off.Relative_x != hammer_on_pull_offOther.Relative_x {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "Relative_x"))
	}
	if hammer_on_pull_off.Relative_y != hammer_on_pull_offOther.Relative_y {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "Relative_y"))
	}
	if hammer_on_pull_off.Font_family != hammer_on_pull_offOther.Font_family {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "Font_family"))
	}
	if hammer_on_pull_off.Font_style != hammer_on_pull_offOther.Font_style {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "Font_style"))
	}
	if hammer_on_pull_off.Font_size != hammer_on_pull_offOther.Font_size {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "Font_size"))
	}
	if hammer_on_pull_off.Font_weight != hammer_on_pull_offOther.Font_weight {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "Font_weight"))
	}
	if hammer_on_pull_off.Color != hammer_on_pull_offOther.Color {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "Color"))
	}
	if hammer_on_pull_off.Placement != hammer_on_pull_offOther.Placement {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "Placement"))
	}
	if hammer_on_pull_off.EnclosedText != hammer_on_pull_offOther.EnclosedText {
		diffs = append(diffs, hammer_on_pull_off.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (handbell *Handbell) GongDiff(stage *Stage, handbellOther *Handbell) (diffs []string) {
	// insertion point for field diffs
	if handbell.Name != handbellOther.Name {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Name"))
	}
	if handbell.Default_x != handbellOther.Default_x {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Default_x"))
	}
	if handbell.Default_y != handbellOther.Default_y {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Default_y"))
	}
	if handbell.Relative_x != handbellOther.Relative_x {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Relative_x"))
	}
	if handbell.Relative_y != handbellOther.Relative_y {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Relative_y"))
	}
	if handbell.Font_family != handbellOther.Font_family {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Font_family"))
	}
	if handbell.Font_style != handbellOther.Font_style {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Font_style"))
	}
	if handbell.Font_size != handbellOther.Font_size {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Font_size"))
	}
	if handbell.Font_weight != handbellOther.Font_weight {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Font_weight"))
	}
	if handbell.Color != handbellOther.Color {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Color"))
	}
	if handbell.Placement != handbellOther.Placement {
		diffs = append(diffs, handbell.GongMarshallField(stage, "Placement"))
	}
	if handbell.EnclosedText != handbellOther.EnclosedText {
		diffs = append(diffs, handbell.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (harmon_closed *Harmon_closed) GongDiff(stage *Stage, harmon_closedOther *Harmon_closed) (diffs []string) {
	// insertion point for field diffs
	if harmon_closed.Name != harmon_closedOther.Name {
		diffs = append(diffs, harmon_closed.GongMarshallField(stage, "Name"))
	}
	if harmon_closed.Location != harmon_closedOther.Location {
		diffs = append(diffs, harmon_closed.GongMarshallField(stage, "Location"))
	}
	if harmon_closed.EnclosedText != harmon_closedOther.EnclosedText {
		diffs = append(diffs, harmon_closed.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (harmon_mute *Harmon_mute) GongDiff(stage *Stage, harmon_muteOther *Harmon_mute) (diffs []string) {
	// insertion point for field diffs
	if harmon_mute.Name != harmon_muteOther.Name {
		diffs = append(diffs, harmon_mute.GongMarshallField(stage, "Name"))
	}
	if harmon_mute.Default_x != harmon_muteOther.Default_x {
		diffs = append(diffs, harmon_mute.GongMarshallField(stage, "Default_x"))
	}
	if harmon_mute.Default_y != harmon_muteOther.Default_y {
		diffs = append(diffs, harmon_mute.GongMarshallField(stage, "Default_y"))
	}
	if harmon_mute.Relative_x != harmon_muteOther.Relative_x {
		diffs = append(diffs, harmon_mute.GongMarshallField(stage, "Relative_x"))
	}
	if harmon_mute.Relative_y != harmon_muteOther.Relative_y {
		diffs = append(diffs, harmon_mute.GongMarshallField(stage, "Relative_y"))
	}
	if harmon_mute.Font_family != harmon_muteOther.Font_family {
		diffs = append(diffs, harmon_mute.GongMarshallField(stage, "Font_family"))
	}
	if harmon_mute.Font_style != harmon_muteOther.Font_style {
		diffs = append(diffs, harmon_mute.GongMarshallField(stage, "Font_style"))
	}
	if harmon_mute.Font_size != harmon_muteOther.Font_size {
		diffs = append(diffs, harmon_mute.GongMarshallField(stage, "Font_size"))
	}
	if harmon_mute.Font_weight != harmon_muteOther.Font_weight {
		diffs = append(diffs, harmon_mute.GongMarshallField(stage, "Font_weight"))
	}
	if harmon_mute.Color != harmon_muteOther.Color {
		diffs = append(diffs, harmon_mute.GongMarshallField(stage, "Color"))
	}
	if harmon_mute.Placement != harmon_muteOther.Placement {
		diffs = append(diffs, harmon_mute.GongMarshallField(stage, "Placement"))
	}
	if harmon_mute.Harmon_closed != harmon_muteOther.Harmon_closed {
		diffs = append(diffs, harmon_mute.GongMarshallField(stage, "Harmon_closed"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (harmonic *Harmonic) GongDiff(stage *Stage, harmonicOther *Harmonic) (diffs []string) {
	// insertion point for field diffs
	if harmonic.Name != harmonicOther.Name {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Name"))
	}
	if harmonic.Print_object != harmonicOther.Print_object {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Print_object"))
	}
	if harmonic.Default_x != harmonicOther.Default_x {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Default_x"))
	}
	if harmonic.Default_y != harmonicOther.Default_y {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Default_y"))
	}
	if harmonic.Relative_x != harmonicOther.Relative_x {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Relative_x"))
	}
	if harmonic.Relative_y != harmonicOther.Relative_y {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Relative_y"))
	}
	if harmonic.Font_family != harmonicOther.Font_family {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Font_family"))
	}
	if harmonic.Font_style != harmonicOther.Font_style {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Font_style"))
	}
	if harmonic.Font_size != harmonicOther.Font_size {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Font_size"))
	}
	if harmonic.Font_weight != harmonicOther.Font_weight {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Font_weight"))
	}
	if harmonic.Color != harmonicOther.Color {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Color"))
	}
	if harmonic.Placement != harmonicOther.Placement {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Placement"))
	}
	if harmonic.Natural != harmonicOther.Natural {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Natural"))
	}
	if harmonic.Artificial != harmonicOther.Artificial {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Artificial"))
	}
	if harmonic.Base_pitch != harmonicOther.Base_pitch {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Base_pitch"))
	}
	if harmonic.Touching_pitch != harmonicOther.Touching_pitch {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Touching_pitch"))
	}
	if harmonic.Sounding_pitch != harmonicOther.Sounding_pitch {
		diffs = append(diffs, harmonic.GongMarshallField(stage, "Sounding_pitch"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (harmony *Harmony) GongDiff(stage *Stage, harmonyOther *Harmony) (diffs []string) {
	// insertion point for field diffs
	if harmony.Name != harmonyOther.Name {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Name"))
	}
	if harmony.Type != harmonyOther.Type {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Type"))
	}
	if harmony.Print_frame != harmonyOther.Print_frame {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Print_frame"))
	}
	if harmony.Arrangement != harmonyOther.Arrangement {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Arrangement"))
	}
	if harmony.Print_object != harmonyOther.Print_object {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Print_object"))
	}
	if harmony.Default_x != harmonyOther.Default_x {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Default_x"))
	}
	if harmony.Default_y != harmonyOther.Default_y {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Default_y"))
	}
	if harmony.Relative_x != harmonyOther.Relative_x {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Relative_x"))
	}
	if harmony.Relative_y != harmonyOther.Relative_y {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Relative_y"))
	}
	if harmony.Font_family != harmonyOther.Font_family {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Font_family"))
	}
	if harmony.Font_style != harmonyOther.Font_style {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Font_style"))
	}
	if harmony.Font_size != harmonyOther.Font_size {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Font_size"))
	}
	if harmony.Font_weight != harmonyOther.Font_weight {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Font_weight"))
	}
	if harmony.Color != harmonyOther.Color {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Color"))
	}
	if harmony.Placement != harmonyOther.Placement {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Placement"))
	}
	if harmony.System != harmonyOther.System {
		diffs = append(diffs, harmony.GongMarshallField(stage, "System"))
	}
	if harmony.Id != harmonyOther.Id {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Id"))
	}
	if harmony.Root != harmonyOther.Root {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Root"))
	}
	if harmony.Numeral != harmonyOther.Numeral {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Numeral"))
	}
	if harmony.Function != harmonyOther.Function {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Function"))
	}
	if harmony.Kind != harmonyOther.Kind {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Kind"))
	}
	if harmony.Inversion != harmonyOther.Inversion {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Inversion"))
	}
	if harmony.Bass != harmonyOther.Bass {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Bass"))
	}
	if ops := __gong__diffSliceOfPointers(stage, harmony, "Degree", harmonyOther.Degree, harmony.Degree); ops != "" {
		diffs = append(diffs, ops)
	}
	if harmony.Frame != harmonyOther.Frame {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Frame"))
	}
	if harmony.Offset != harmonyOther.Offset {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Offset"))
	}
	if harmony.Footnote != harmonyOther.Footnote {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Footnote"))
	}
	if harmony.Level != harmonyOther.Level {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Level"))
	}
	if harmony.Staff != harmonyOther.Staff {
		diffs = append(diffs, harmony.GongMarshallField(stage, "Staff"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (harmony_alter *Harmony_alter) GongDiff(stage *Stage, harmony_alterOther *Harmony_alter) (diffs []string) {
	// insertion point for field diffs
	if harmony_alter.Name != harmony_alterOther.Name {
		diffs = append(diffs, harmony_alter.GongMarshallField(stage, "Name"))
	}
	if harmony_alter.Location != harmony_alterOther.Location {
		diffs = append(diffs, harmony_alter.GongMarshallField(stage, "Location"))
	}
	if harmony_alter.Print_object != harmony_alterOther.Print_object {
		diffs = append(diffs, harmony_alter.GongMarshallField(stage, "Print_object"))
	}
	if harmony_alter.Default_x != harmony_alterOther.Default_x {
		diffs = append(diffs, harmony_alter.GongMarshallField(stage, "Default_x"))
	}
	if harmony_alter.Default_y != harmony_alterOther.Default_y {
		diffs = append(diffs, harmony_alter.GongMarshallField(stage, "Default_y"))
	}
	if harmony_alter.Relative_x != harmony_alterOther.Relative_x {
		diffs = append(diffs, harmony_alter.GongMarshallField(stage, "Relative_x"))
	}
	if harmony_alter.Relative_y != harmony_alterOther.Relative_y {
		diffs = append(diffs, harmony_alter.GongMarshallField(stage, "Relative_y"))
	}
	if harmony_alter.Font_family != harmony_alterOther.Font_family {
		diffs = append(diffs, harmony_alter.GongMarshallField(stage, "Font_family"))
	}
	if harmony_alter.Font_style != harmony_alterOther.Font_style {
		diffs = append(diffs, harmony_alter.GongMarshallField(stage, "Font_style"))
	}
	if harmony_alter.Font_size != harmony_alterOther.Font_size {
		diffs = append(diffs, harmony_alter.GongMarshallField(stage, "Font_size"))
	}
	if harmony_alter.Font_weight != harmony_alterOther.Font_weight {
		diffs = append(diffs, harmony_alter.GongMarshallField(stage, "Font_weight"))
	}
	if harmony_alter.Color != harmony_alterOther.Color {
		diffs = append(diffs, harmony_alter.GongMarshallField(stage, "Color"))
	}
	if harmony_alter.EnclosedText != harmony_alterOther.EnclosedText {
		diffs = append(diffs, harmony_alter.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (harp_pedals *Harp_pedals) GongDiff(stage *Stage, harp_pedalsOther *Harp_pedals) (diffs []string) {
	// insertion point for field diffs
	if harp_pedals.Name != harp_pedalsOther.Name {
		diffs = append(diffs, harp_pedals.GongMarshallField(stage, "Name"))
	}
	if harp_pedals.Default_x != harp_pedalsOther.Default_x {
		diffs = append(diffs, harp_pedals.GongMarshallField(stage, "Default_x"))
	}
	if harp_pedals.Default_y != harp_pedalsOther.Default_y {
		diffs = append(diffs, harp_pedals.GongMarshallField(stage, "Default_y"))
	}
	if harp_pedals.Relative_x != harp_pedalsOther.Relative_x {
		diffs = append(diffs, harp_pedals.GongMarshallField(stage, "Relative_x"))
	}
	if harp_pedals.Relative_y != harp_pedalsOther.Relative_y {
		diffs = append(diffs, harp_pedals.GongMarshallField(stage, "Relative_y"))
	}
	if harp_pedals.Font_family != harp_pedalsOther.Font_family {
		diffs = append(diffs, harp_pedals.GongMarshallField(stage, "Font_family"))
	}
	if harp_pedals.Font_style != harp_pedalsOther.Font_style {
		diffs = append(diffs, harp_pedals.GongMarshallField(stage, "Font_style"))
	}
	if harp_pedals.Font_size != harp_pedalsOther.Font_size {
		diffs = append(diffs, harp_pedals.GongMarshallField(stage, "Font_size"))
	}
	if harp_pedals.Font_weight != harp_pedalsOther.Font_weight {
		diffs = append(diffs, harp_pedals.GongMarshallField(stage, "Font_weight"))
	}
	if harp_pedals.Color != harp_pedalsOther.Color {
		diffs = append(diffs, harp_pedals.GongMarshallField(stage, "Color"))
	}
	if harp_pedals.Halign != harp_pedalsOther.Halign {
		diffs = append(diffs, harp_pedals.GongMarshallField(stage, "Halign"))
	}
	if harp_pedals.Valign != harp_pedalsOther.Valign {
		diffs = append(diffs, harp_pedals.GongMarshallField(stage, "Valign"))
	}
	if harp_pedals.Id != harp_pedalsOther.Id {
		diffs = append(diffs, harp_pedals.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, harp_pedals, "Pedal_tuning", harp_pedalsOther.Pedal_tuning, harp_pedals.Pedal_tuning); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (heel_toe *Heel_toe) GongDiff(stage *Stage, heel_toeOther *Heel_toe) (diffs []string) {
	// insertion point for field diffs
	if heel_toe.Name != heel_toeOther.Name {
		diffs = append(diffs, heel_toe.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (hole *Hole) GongDiff(stage *Stage, holeOther *Hole) (diffs []string) {
	// insertion point for field diffs
	if hole.Name != holeOther.Name {
		diffs = append(diffs, hole.GongMarshallField(stage, "Name"))
	}
	if hole.Default_x != holeOther.Default_x {
		diffs = append(diffs, hole.GongMarshallField(stage, "Default_x"))
	}
	if hole.Default_y != holeOther.Default_y {
		diffs = append(diffs, hole.GongMarshallField(stage, "Default_y"))
	}
	if hole.Relative_x != holeOther.Relative_x {
		diffs = append(diffs, hole.GongMarshallField(stage, "Relative_x"))
	}
	if hole.Relative_y != holeOther.Relative_y {
		diffs = append(diffs, hole.GongMarshallField(stage, "Relative_y"))
	}
	if hole.Font_family != holeOther.Font_family {
		diffs = append(diffs, hole.GongMarshallField(stage, "Font_family"))
	}
	if hole.Font_style != holeOther.Font_style {
		diffs = append(diffs, hole.GongMarshallField(stage, "Font_style"))
	}
	if hole.Font_size != holeOther.Font_size {
		diffs = append(diffs, hole.GongMarshallField(stage, "Font_size"))
	}
	if hole.Font_weight != holeOther.Font_weight {
		diffs = append(diffs, hole.GongMarshallField(stage, "Font_weight"))
	}
	if hole.Color != holeOther.Color {
		diffs = append(diffs, hole.GongMarshallField(stage, "Color"))
	}
	if hole.Placement != holeOther.Placement {
		diffs = append(diffs, hole.GongMarshallField(stage, "Placement"))
	}
	if hole.Hole_type != holeOther.Hole_type {
		diffs = append(diffs, hole.GongMarshallField(stage, "Hole_type"))
	}
	if hole.Hole_closed != holeOther.Hole_closed {
		diffs = append(diffs, hole.GongMarshallField(stage, "Hole_closed"))
	}
	if hole.Hole_shape != holeOther.Hole_shape {
		diffs = append(diffs, hole.GongMarshallField(stage, "Hole_shape"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (hole_closed *Hole_closed) GongDiff(stage *Stage, hole_closedOther *Hole_closed) (diffs []string) {
	// insertion point for field diffs
	if hole_closed.Name != hole_closedOther.Name {
		diffs = append(diffs, hole_closed.GongMarshallField(stage, "Name"))
	}
	if hole_closed.Location != hole_closedOther.Location {
		diffs = append(diffs, hole_closed.GongMarshallField(stage, "Location"))
	}
	if hole_closed.EnclosedText != hole_closedOther.EnclosedText {
		diffs = append(diffs, hole_closed.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (horizontal_turn *Horizontal_turn) GongDiff(stage *Stage, horizontal_turnOther *Horizontal_turn) (diffs []string) {
	// insertion point for field diffs
	if horizontal_turn.Name != horizontal_turnOther.Name {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Name"))
	}
	if horizontal_turn.Slash != horizontal_turnOther.Slash {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Slash"))
	}
	if horizontal_turn.Default_x != horizontal_turnOther.Default_x {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Default_x"))
	}
	if horizontal_turn.Default_y != horizontal_turnOther.Default_y {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Default_y"))
	}
	if horizontal_turn.Relative_x != horizontal_turnOther.Relative_x {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Relative_x"))
	}
	if horizontal_turn.Relative_y != horizontal_turnOther.Relative_y {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Relative_y"))
	}
	if horizontal_turn.Font_family != horizontal_turnOther.Font_family {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Font_family"))
	}
	if horizontal_turn.Font_style != horizontal_turnOther.Font_style {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Font_style"))
	}
	if horizontal_turn.Font_size != horizontal_turnOther.Font_size {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Font_size"))
	}
	if horizontal_turn.Font_weight != horizontal_turnOther.Font_weight {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Font_weight"))
	}
	if horizontal_turn.Color != horizontal_turnOther.Color {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Color"))
	}
	if horizontal_turn.Placement != horizontal_turnOther.Placement {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Placement"))
	}
	if horizontal_turn.Start_note != horizontal_turnOther.Start_note {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Start_note"))
	}
	if horizontal_turn.Trill_step != horizontal_turnOther.Trill_step {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Trill_step"))
	}
	if horizontal_turn.Two_note_turn != horizontal_turnOther.Two_note_turn {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Two_note_turn"))
	}
	if horizontal_turn.Accelerate != horizontal_turnOther.Accelerate {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Accelerate"))
	}
	if horizontal_turn.Beats != horizontal_turnOther.Beats {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Beats"))
	}
	if horizontal_turn.Second_beat != horizontal_turnOther.Second_beat {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Second_beat"))
	}
	if horizontal_turn.Last_beat != horizontal_turnOther.Last_beat {
		diffs = append(diffs, horizontal_turn.GongMarshallField(stage, "Last_beat"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (identification *Identification) GongDiff(stage *Stage, identificationOther *Identification) (diffs []string) {
	// insertion point for field diffs
	if identification.Name != identificationOther.Name {
		diffs = append(diffs, identification.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, identification, "Creator", identificationOther.Creator, identification.Creator); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, identification, "Rights", identificationOther.Rights, identification.Rights); ops != "" {
		diffs = append(diffs, ops)
	}
	if identification.Encoding != identificationOther.Encoding {
		diffs = append(diffs, identification.GongMarshallField(stage, "Encoding"))
	}
	if identification.Source != identificationOther.Source {
		diffs = append(diffs, identification.GongMarshallField(stage, "Source"))
	}
	if ops := __gong__diffSliceOfPointers(stage, identification, "Relation", identificationOther.Relation, identification.Relation); ops != "" {
		diffs = append(diffs, ops)
	}
	if identification.Miscellaneous != identificationOther.Miscellaneous {
		diffs = append(diffs, identification.GongMarshallField(stage, "Miscellaneous"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (image *Image) GongDiff(stage *Stage, imageOther *Image) (diffs []string) {
	// insertion point for field diffs
	if image.Name != imageOther.Name {
		diffs = append(diffs, image.GongMarshallField(stage, "Name"))
	}
	if image.Source != imageOther.Source {
		diffs = append(diffs, image.GongMarshallField(stage, "Source"))
	}
	if image.Type != imageOther.Type {
		diffs = append(diffs, image.GongMarshallField(stage, "Type"))
	}
	if image.Height != imageOther.Height {
		diffs = append(diffs, image.GongMarshallField(stage, "Height"))
	}
	if image.Width != imageOther.Width {
		diffs = append(diffs, image.GongMarshallField(stage, "Width"))
	}
	if image.Default_x != imageOther.Default_x {
		diffs = append(diffs, image.GongMarshallField(stage, "Default_x"))
	}
	if image.Default_y != imageOther.Default_y {
		diffs = append(diffs, image.GongMarshallField(stage, "Default_y"))
	}
	if image.Relative_x != imageOther.Relative_x {
		diffs = append(diffs, image.GongMarshallField(stage, "Relative_x"))
	}
	if image.Relative_y != imageOther.Relative_y {
		diffs = append(diffs, image.GongMarshallField(stage, "Relative_y"))
	}
	if image.Halign != imageOther.Halign {
		diffs = append(diffs, image.GongMarshallField(stage, "Halign"))
	}
	if image.Valign != imageOther.Valign {
		diffs = append(diffs, image.GongMarshallField(stage, "Valign"))
	}
	if image.Id != imageOther.Id {
		diffs = append(diffs, image.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (instrument *Instrument) GongDiff(stage *Stage, instrumentOther *Instrument) (diffs []string) {
	// insertion point for field diffs
	if instrument.Name != instrumentOther.Name {
		diffs = append(diffs, instrument.GongMarshallField(stage, "Name"))
	}
	if instrument.Id != instrumentOther.Id {
		diffs = append(diffs, instrument.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (instrument_change *Instrument_change) GongDiff(stage *Stage, instrument_changeOther *Instrument_change) (diffs []string) {
	// insertion point for field diffs
	if instrument_change.Name != instrument_changeOther.Name {
		diffs = append(diffs, instrument_change.GongMarshallField(stage, "Name"))
	}
	if instrument_change.Id != instrument_changeOther.Id {
		diffs = append(diffs, instrument_change.GongMarshallField(stage, "Id"))
	}
	if instrument_change.Instrument_sound != instrument_changeOther.Instrument_sound {
		diffs = append(diffs, instrument_change.GongMarshallField(stage, "Instrument_sound"))
	}
	if instrument_change.Solo != instrument_changeOther.Solo {
		diffs = append(diffs, instrument_change.GongMarshallField(stage, "Solo"))
	}
	if instrument_change.Ensemble != instrument_changeOther.Ensemble {
		diffs = append(diffs, instrument_change.GongMarshallField(stage, "Ensemble"))
	}
	if instrument_change.Virtual_instrument != instrument_changeOther.Virtual_instrument {
		diffs = append(diffs, instrument_change.GongMarshallField(stage, "Virtual_instrument"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (instrument_link *Instrument_link) GongDiff(stage *Stage, instrument_linkOther *Instrument_link) (diffs []string) {
	// insertion point for field diffs
	if instrument_link.Name != instrument_linkOther.Name {
		diffs = append(diffs, instrument_link.GongMarshallField(stage, "Name"))
	}
	if instrument_link.Id != instrument_linkOther.Id {
		diffs = append(diffs, instrument_link.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (interchangeable *Interchangeable) GongDiff(stage *Stage, interchangeableOther *Interchangeable) (diffs []string) {
	// insertion point for field diffs
	if interchangeable.Name != interchangeableOther.Name {
		diffs = append(diffs, interchangeable.GongMarshallField(stage, "Name"))
	}
	if interchangeable.Symbol != interchangeableOther.Symbol {
		diffs = append(diffs, interchangeable.GongMarshallField(stage, "Symbol"))
	}
	if interchangeable.Separator != interchangeableOther.Separator {
		diffs = append(diffs, interchangeable.GongMarshallField(stage, "Separator"))
	}
	if interchangeable.Time_relation != interchangeableOther.Time_relation {
		diffs = append(diffs, interchangeable.GongMarshallField(stage, "Time_relation"))
	}
	if interchangeable.Beats != interchangeableOther.Beats {
		diffs = append(diffs, interchangeable.GongMarshallField(stage, "Beats"))
	}
	if interchangeable.Beat_type != interchangeableOther.Beat_type {
		diffs = append(diffs, interchangeable.GongMarshallField(stage, "Beat_type"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (inversion *Inversion) GongDiff(stage *Stage, inversionOther *Inversion) (diffs []string) {
	// insertion point for field diffs
	if inversion.Name != inversionOther.Name {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Name"))
	}
	if inversion.Text != inversionOther.Text {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Text"))
	}
	if inversion.Default_x != inversionOther.Default_x {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Default_x"))
	}
	if inversion.Default_y != inversionOther.Default_y {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Default_y"))
	}
	if inversion.Relative_x != inversionOther.Relative_x {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Relative_x"))
	}
	if inversion.Relative_y != inversionOther.Relative_y {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Relative_y"))
	}
	if inversion.Font_family != inversionOther.Font_family {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Font_family"))
	}
	if inversion.Font_style != inversionOther.Font_style {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Font_style"))
	}
	if inversion.Font_size != inversionOther.Font_size {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Font_size"))
	}
	if inversion.Font_weight != inversionOther.Font_weight {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Font_weight"))
	}
	if inversion.Color != inversionOther.Color {
		diffs = append(diffs, inversion.GongMarshallField(stage, "Color"))
	}
	if inversion.EnclosedText != inversionOther.EnclosedText {
		diffs = append(diffs, inversion.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (key *Key) GongDiff(stage *Stage, keyOther *Key) (diffs []string) {
	// insertion point for field diffs
	if key.Name != keyOther.Name {
		diffs = append(diffs, key.GongMarshallField(stage, "Name"))
	}
	if key.Number != keyOther.Number {
		diffs = append(diffs, key.GongMarshallField(stage, "Number"))
	}
	if key.Default_x != keyOther.Default_x {
		diffs = append(diffs, key.GongMarshallField(stage, "Default_x"))
	}
	if key.Default_y != keyOther.Default_y {
		diffs = append(diffs, key.GongMarshallField(stage, "Default_y"))
	}
	if key.Relative_x != keyOther.Relative_x {
		diffs = append(diffs, key.GongMarshallField(stage, "Relative_x"))
	}
	if key.Relative_y != keyOther.Relative_y {
		diffs = append(diffs, key.GongMarshallField(stage, "Relative_y"))
	}
	if key.Font_family != keyOther.Font_family {
		diffs = append(diffs, key.GongMarshallField(stage, "Font_family"))
	}
	if key.Font_style != keyOther.Font_style {
		diffs = append(diffs, key.GongMarshallField(stage, "Font_style"))
	}
	if key.Font_size != keyOther.Font_size {
		diffs = append(diffs, key.GongMarshallField(stage, "Font_size"))
	}
	if key.Font_weight != keyOther.Font_weight {
		diffs = append(diffs, key.GongMarshallField(stage, "Font_weight"))
	}
	if key.Color != keyOther.Color {
		diffs = append(diffs, key.GongMarshallField(stage, "Color"))
	}
	if key.Print_object != keyOther.Print_object {
		diffs = append(diffs, key.GongMarshallField(stage, "Print_object"))
	}
	if key.Id != keyOther.Id {
		diffs = append(diffs, key.GongMarshallField(stage, "Id"))
	}
	if key.Cancel != keyOther.Cancel {
		diffs = append(diffs, key.GongMarshallField(stage, "Cancel"))
	}
	if key.Fifths != keyOther.Fifths {
		diffs = append(diffs, key.GongMarshallField(stage, "Fifths"))
	}
	if key.Mode != keyOther.Mode {
		diffs = append(diffs, key.GongMarshallField(stage, "Mode"))
	}
	if key.Key_step != keyOther.Key_step {
		diffs = append(diffs, key.GongMarshallField(stage, "Key_step"))
	}
	if key.Key_alter != keyOther.Key_alter {
		diffs = append(diffs, key.GongMarshallField(stage, "Key_alter"))
	}
	if key.Key_accidental != keyOther.Key_accidental {
		diffs = append(diffs, key.GongMarshallField(stage, "Key_accidental"))
	}
	if ops := __gong__diffSliceOfPointers(stage, key, "Key_octave", keyOther.Key_octave, key.Key_octave); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (key_accidental *Key_accidental) GongDiff(stage *Stage, key_accidentalOther *Key_accidental) (diffs []string) {
	// insertion point for field diffs
	if key_accidental.Name != key_accidentalOther.Name {
		diffs = append(diffs, key_accidental.GongMarshallField(stage, "Name"))
	}
	if key_accidental.Smufl != key_accidentalOther.Smufl {
		diffs = append(diffs, key_accidental.GongMarshallField(stage, "Smufl"))
	}
	if key_accidental.EnclosedText != key_accidentalOther.EnclosedText {
		diffs = append(diffs, key_accidental.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (key_octave *Key_octave) GongDiff(stage *Stage, key_octaveOther *Key_octave) (diffs []string) {
	// insertion point for field diffs
	if key_octave.Name != key_octaveOther.Name {
		diffs = append(diffs, key_octave.GongMarshallField(stage, "Name"))
	}
	if key_octave.Number != key_octaveOther.Number {
		diffs = append(diffs, key_octave.GongMarshallField(stage, "Number"))
	}
	if key_octave.Cancel != key_octaveOther.Cancel {
		diffs = append(diffs, key_octave.GongMarshallField(stage, "Cancel"))
	}
	if key_octave.EnclosedText != key_octaveOther.EnclosedText {
		diffs = append(diffs, key_octave.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (kind *Kind) GongDiff(stage *Stage, kindOther *Kind) (diffs []string) {
	// insertion point for field diffs
	if kind.Name != kindOther.Name {
		diffs = append(diffs, kind.GongMarshallField(stage, "Name"))
	}
	if kind.Use_symbols != kindOther.Use_symbols {
		diffs = append(diffs, kind.GongMarshallField(stage, "Use_symbols"))
	}
	if kind.Text != kindOther.Text {
		diffs = append(diffs, kind.GongMarshallField(stage, "Text"))
	}
	if kind.Stack_degrees != kindOther.Stack_degrees {
		diffs = append(diffs, kind.GongMarshallField(stage, "Stack_degrees"))
	}
	if kind.Parentheses_degrees != kindOther.Parentheses_degrees {
		diffs = append(diffs, kind.GongMarshallField(stage, "Parentheses_degrees"))
	}
	if kind.Bracket_degrees != kindOther.Bracket_degrees {
		diffs = append(diffs, kind.GongMarshallField(stage, "Bracket_degrees"))
	}
	if kind.Default_x != kindOther.Default_x {
		diffs = append(diffs, kind.GongMarshallField(stage, "Default_x"))
	}
	if kind.Default_y != kindOther.Default_y {
		diffs = append(diffs, kind.GongMarshallField(stage, "Default_y"))
	}
	if kind.Relative_x != kindOther.Relative_x {
		diffs = append(diffs, kind.GongMarshallField(stage, "Relative_x"))
	}
	if kind.Relative_y != kindOther.Relative_y {
		diffs = append(diffs, kind.GongMarshallField(stage, "Relative_y"))
	}
	if kind.Font_family != kindOther.Font_family {
		diffs = append(diffs, kind.GongMarshallField(stage, "Font_family"))
	}
	if kind.Font_style != kindOther.Font_style {
		diffs = append(diffs, kind.GongMarshallField(stage, "Font_style"))
	}
	if kind.Font_size != kindOther.Font_size {
		diffs = append(diffs, kind.GongMarshallField(stage, "Font_size"))
	}
	if kind.Font_weight != kindOther.Font_weight {
		diffs = append(diffs, kind.GongMarshallField(stage, "Font_weight"))
	}
	if kind.Color != kindOther.Color {
		diffs = append(diffs, kind.GongMarshallField(stage, "Color"))
	}
	if kind.Halign != kindOther.Halign {
		diffs = append(diffs, kind.GongMarshallField(stage, "Halign"))
	}
	if kind.Valign != kindOther.Valign {
		diffs = append(diffs, kind.GongMarshallField(stage, "Valign"))
	}
	if kind.EnclosedText != kindOther.EnclosedText {
		diffs = append(diffs, kind.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (level *Level) GongDiff(stage *Stage, levelOther *Level) (diffs []string) {
	// insertion point for field diffs
	if level.Name != levelOther.Name {
		diffs = append(diffs, level.GongMarshallField(stage, "Name"))
	}
	if level.Reference != levelOther.Reference {
		diffs = append(diffs, level.GongMarshallField(stage, "Reference"))
	}
	if level.Type != levelOther.Type {
		diffs = append(diffs, level.GongMarshallField(stage, "Type"))
	}
	if level.Parentheses != levelOther.Parentheses {
		diffs = append(diffs, level.GongMarshallField(stage, "Parentheses"))
	}
	if level.Bracket != levelOther.Bracket {
		diffs = append(diffs, level.GongMarshallField(stage, "Bracket"))
	}
	if level.Size != levelOther.Size {
		diffs = append(diffs, level.GongMarshallField(stage, "Size"))
	}
	if level.EnclosedText != levelOther.EnclosedText {
		diffs = append(diffs, level.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (line_detail *Line_detail) GongDiff(stage *Stage, line_detailOther *Line_detail) (diffs []string) {
	// insertion point for field diffs
	if line_detail.Name != line_detailOther.Name {
		diffs = append(diffs, line_detail.GongMarshallField(stage, "Name"))
	}
	if line_detail.Line != line_detailOther.Line {
		diffs = append(diffs, line_detail.GongMarshallField(stage, "Line"))
	}
	if line_detail.Width != line_detailOther.Width {
		diffs = append(diffs, line_detail.GongMarshallField(stage, "Width"))
	}
	if line_detail.Color != line_detailOther.Color {
		diffs = append(diffs, line_detail.GongMarshallField(stage, "Color"))
	}
	if line_detail.Line_type != line_detailOther.Line_type {
		diffs = append(diffs, line_detail.GongMarshallField(stage, "Line_type"))
	}
	if line_detail.Print_object != line_detailOther.Print_object {
		diffs = append(diffs, line_detail.GongMarshallField(stage, "Print_object"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (line_width *Line_width) GongDiff(stage *Stage, line_widthOther *Line_width) (diffs []string) {
	// insertion point for field diffs
	if line_width.Name != line_widthOther.Name {
		diffs = append(diffs, line_width.GongMarshallField(stage, "Name"))
	}
	if line_width.Type != line_widthOther.Type {
		diffs = append(diffs, line_width.GongMarshallField(stage, "Type"))
	}
	if line_width.EnclosedText != line_widthOther.EnclosedText {
		diffs = append(diffs, line_width.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (link *Link) GongDiff(stage *Stage, linkOther *Link) (diffs []string) {
	// insertion point for field diffs
	if link.Name != linkOther.Name {
		diffs = append(diffs, link.GongMarshallField(stage, "Name"))
	}
	if link.NameXSD != linkOther.NameXSD {
		diffs = append(diffs, link.GongMarshallField(stage, "NameXSD"))
	}
	if link.Href != linkOther.Href {
		diffs = append(diffs, link.GongMarshallField(stage, "Href"))
	}
	if link.Type != linkOther.Type {
		diffs = append(diffs, link.GongMarshallField(stage, "Type"))
	}
	if link.Role != linkOther.Role {
		diffs = append(diffs, link.GongMarshallField(stage, "Role"))
	}
	if link.Title != linkOther.Title {
		diffs = append(diffs, link.GongMarshallField(stage, "Title"))
	}
	if link.Show != linkOther.Show {
		diffs = append(diffs, link.GongMarshallField(stage, "Show"))
	}
	if link.Actuate != linkOther.Actuate {
		diffs = append(diffs, link.GongMarshallField(stage, "Actuate"))
	}
	if link.Element != linkOther.Element {
		diffs = append(diffs, link.GongMarshallField(stage, "Element"))
	}
	if link.Position != linkOther.Position {
		diffs = append(diffs, link.GongMarshallField(stage, "Position"))
	}
	if link.Default_x != linkOther.Default_x {
		diffs = append(diffs, link.GongMarshallField(stage, "Default_x"))
	}
	if link.Default_y != linkOther.Default_y {
		diffs = append(diffs, link.GongMarshallField(stage, "Default_y"))
	}
	if link.Relative_x != linkOther.Relative_x {
		diffs = append(diffs, link.GongMarshallField(stage, "Relative_x"))
	}
	if link.Relative_y != linkOther.Relative_y {
		diffs = append(diffs, link.GongMarshallField(stage, "Relative_y"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (listen *Listen) GongDiff(stage *Stage, listenOther *Listen) (diffs []string) {
	// insertion point for field diffs
	if listen.Name != listenOther.Name {
		diffs = append(diffs, listen.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, listen, "Assess", listenOther.Assess, listen.Assess); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, listen, "Wait", listenOther.Wait, listen.Wait); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, listen, "Other_listen", listenOther.Other_listen, listen.Other_listen); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (listening *Listening) GongDiff(stage *Stage, listeningOther *Listening) (diffs []string) {
	// insertion point for field diffs
	if listening.Name != listeningOther.Name {
		diffs = append(diffs, listening.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, listening, "Sync", listeningOther.Sync, listening.Sync); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, listening, "Other_listening", listeningOther.Other_listening, listening.Other_listening); ops != "" {
		diffs = append(diffs, ops)
	}
	if listening.Offset != listeningOther.Offset {
		diffs = append(diffs, listening.GongMarshallField(stage, "Offset"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lyric *Lyric) GongDiff(stage *Stage, lyricOther *Lyric) (diffs []string) {
	// insertion point for field diffs
	if lyric.Name != lyricOther.Name {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Name"))
	}
	if lyric.Number != lyricOther.Number {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Number"))
	}
	if lyric.NameXSD != lyricOther.NameXSD {
		diffs = append(diffs, lyric.GongMarshallField(stage, "NameXSD"))
	}
	if lyric.Time_only != lyricOther.Time_only {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Time_only"))
	}
	if lyric.Justify != lyricOther.Justify {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Justify"))
	}
	if lyric.Default_x != lyricOther.Default_x {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Default_x"))
	}
	if lyric.Default_y != lyricOther.Default_y {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Default_y"))
	}
	if lyric.Relative_x != lyricOther.Relative_x {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Relative_x"))
	}
	if lyric.Relative_y != lyricOther.Relative_y {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Relative_y"))
	}
	if lyric.Placement != lyricOther.Placement {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Placement"))
	}
	if lyric.Color != lyricOther.Color {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Color"))
	}
	if lyric.Print_object != lyricOther.Print_object {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Print_object"))
	}
	if lyric.Id != lyricOther.Id {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, lyric, "Elision", lyricOther.Elision, lyric.Elision); ops != "" {
		diffs = append(diffs, ops)
	}
	if lyric.Syllabic != lyricOther.Syllabic {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Syllabic"))
	}
	if ops := __gong__diffSliceOfPointers(stage, lyric, "Text", lyricOther.Text, lyric.Text); ops != "" {
		diffs = append(diffs, ops)
	}
	if lyric.Extend != lyricOther.Extend {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Extend"))
	}
	if lyric.Laughing != lyricOther.Laughing {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Laughing"))
	}
	if lyric.Humming != lyricOther.Humming {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Humming"))
	}
	if lyric.End_line != lyricOther.End_line {
		diffs = append(diffs, lyric.GongMarshallField(stage, "End_line"))
	}
	if lyric.End_paragraph != lyricOther.End_paragraph {
		diffs = append(diffs, lyric.GongMarshallField(stage, "End_paragraph"))
	}
	if lyric.Footnote != lyricOther.Footnote {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Footnote"))
	}
	if lyric.Level != lyricOther.Level {
		diffs = append(diffs, lyric.GongMarshallField(stage, "Level"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lyric_font *Lyric_font) GongDiff(stage *Stage, lyric_fontOther *Lyric_font) (diffs []string) {
	// insertion point for field diffs
	if lyric_font.Name != lyric_fontOther.Name {
		diffs = append(diffs, lyric_font.GongMarshallField(stage, "Name"))
	}
	if lyric_font.Number != lyric_fontOther.Number {
		diffs = append(diffs, lyric_font.GongMarshallField(stage, "Number"))
	}
	if lyric_font.NameXSD != lyric_fontOther.NameXSD {
		diffs = append(diffs, lyric_font.GongMarshallField(stage, "NameXSD"))
	}
	if lyric_font.Font_family != lyric_fontOther.Font_family {
		diffs = append(diffs, lyric_font.GongMarshallField(stage, "Font_family"))
	}
	if lyric_font.Font_style != lyric_fontOther.Font_style {
		diffs = append(diffs, lyric_font.GongMarshallField(stage, "Font_style"))
	}
	if lyric_font.Font_size != lyric_fontOther.Font_size {
		diffs = append(diffs, lyric_font.GongMarshallField(stage, "Font_size"))
	}
	if lyric_font.Font_weight != lyric_fontOther.Font_weight {
		diffs = append(diffs, lyric_font.GongMarshallField(stage, "Font_weight"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (lyric_language *Lyric_language) GongDiff(stage *Stage, lyric_languageOther *Lyric_language) (diffs []string) {
	// insertion point for field diffs
	if lyric_language.Name != lyric_languageOther.Name {
		diffs = append(diffs, lyric_language.GongMarshallField(stage, "Name"))
	}
	if lyric_language.Number != lyric_languageOther.Number {
		diffs = append(diffs, lyric_language.GongMarshallField(stage, "Number"))
	}
	if lyric_language.NameXSD != lyric_languageOther.NameXSD {
		diffs = append(diffs, lyric_language.GongMarshallField(stage, "NameXSD"))
	}
	if lyric_language.Lang != lyric_languageOther.Lang {
		diffs = append(diffs, lyric_language.GongMarshallField(stage, "Lang"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measure_layout *Measure_layout) GongDiff(stage *Stage, measure_layoutOther *Measure_layout) (diffs []string) {
	// insertion point for field diffs
	if measure_layout.Name != measure_layoutOther.Name {
		diffs = append(diffs, measure_layout.GongMarshallField(stage, "Name"))
	}
	if measure_layout.Measure_distance != measure_layoutOther.Measure_distance {
		diffs = append(diffs, measure_layout.GongMarshallField(stage, "Measure_distance"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measure_numbering *Measure_numbering) GongDiff(stage *Stage, measure_numberingOther *Measure_numbering) (diffs []string) {
	// insertion point for field diffs
	if measure_numbering.Name != measure_numberingOther.Name {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Name"))
	}
	if measure_numbering.System != measure_numberingOther.System {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "System"))
	}
	if measure_numbering.Staff != measure_numberingOther.Staff {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Staff"))
	}
	if measure_numbering.Multiple_rest_always != measure_numberingOther.Multiple_rest_always {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Multiple_rest_always"))
	}
	if measure_numbering.Multiple_rest_range != measure_numberingOther.Multiple_rest_range {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Multiple_rest_range"))
	}
	if measure_numbering.Default_x != measure_numberingOther.Default_x {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Default_x"))
	}
	if measure_numbering.Default_y != measure_numberingOther.Default_y {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Default_y"))
	}
	if measure_numbering.Relative_x != measure_numberingOther.Relative_x {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Relative_x"))
	}
	if measure_numbering.Relative_y != measure_numberingOther.Relative_y {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Relative_y"))
	}
	if measure_numbering.Font_family != measure_numberingOther.Font_family {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Font_family"))
	}
	if measure_numbering.Font_style != measure_numberingOther.Font_style {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Font_style"))
	}
	if measure_numbering.Font_size != measure_numberingOther.Font_size {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Font_size"))
	}
	if measure_numbering.Font_weight != measure_numberingOther.Font_weight {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Font_weight"))
	}
	if measure_numbering.Color != measure_numberingOther.Color {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Color"))
	}
	if measure_numbering.Halign != measure_numberingOther.Halign {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Halign"))
	}
	if measure_numbering.Valign != measure_numberingOther.Valign {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "Valign"))
	}
	if measure_numbering.EnclosedText != measure_numberingOther.EnclosedText {
		diffs = append(diffs, measure_numbering.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measure_repeat *Measure_repeat) GongDiff(stage *Stage, measure_repeatOther *Measure_repeat) (diffs []string) {
	// insertion point for field diffs
	if measure_repeat.Name != measure_repeatOther.Name {
		diffs = append(diffs, measure_repeat.GongMarshallField(stage, "Name"))
	}
	if measure_repeat.Type != measure_repeatOther.Type {
		diffs = append(diffs, measure_repeat.GongMarshallField(stage, "Type"))
	}
	if measure_repeat.Slashes != measure_repeatOther.Slashes {
		diffs = append(diffs, measure_repeat.GongMarshallField(stage, "Slashes"))
	}
	if measure_repeat.EnclosedText != measure_repeatOther.EnclosedText {
		diffs = append(diffs, measure_repeat.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measure_style *Measure_style) GongDiff(stage *Stage, measure_styleOther *Measure_style) (diffs []string) {
	// insertion point for field diffs
	if measure_style.Name != measure_styleOther.Name {
		diffs = append(diffs, measure_style.GongMarshallField(stage, "Name"))
	}
	if measure_style.Number != measure_styleOther.Number {
		diffs = append(diffs, measure_style.GongMarshallField(stage, "Number"))
	}
	if measure_style.Font_family != measure_styleOther.Font_family {
		diffs = append(diffs, measure_style.GongMarshallField(stage, "Font_family"))
	}
	if measure_style.Font_style != measure_styleOther.Font_style {
		diffs = append(diffs, measure_style.GongMarshallField(stage, "Font_style"))
	}
	if measure_style.Font_size != measure_styleOther.Font_size {
		diffs = append(diffs, measure_style.GongMarshallField(stage, "Font_size"))
	}
	if measure_style.Font_weight != measure_styleOther.Font_weight {
		diffs = append(diffs, measure_style.GongMarshallField(stage, "Font_weight"))
	}
	if measure_style.Color != measure_styleOther.Color {
		diffs = append(diffs, measure_style.GongMarshallField(stage, "Color"))
	}
	if measure_style.Id != measure_styleOther.Id {
		diffs = append(diffs, measure_style.GongMarshallField(stage, "Id"))
	}
	if measure_style.Multiple_rest != measure_styleOther.Multiple_rest {
		diffs = append(diffs, measure_style.GongMarshallField(stage, "Multiple_rest"))
	}
	if measure_style.Measure_repeat != measure_styleOther.Measure_repeat {
		diffs = append(diffs, measure_style.GongMarshallField(stage, "Measure_repeat"))
	}
	if measure_style.Beat_repeat != measure_styleOther.Beat_repeat {
		diffs = append(diffs, measure_style.GongMarshallField(stage, "Beat_repeat"))
	}
	if measure_style.Slash != measure_styleOther.Slash {
		diffs = append(diffs, measure_style.GongMarshallField(stage, "Slash"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (membrane *Membrane) GongDiff(stage *Stage, membraneOther *Membrane) (diffs []string) {
	// insertion point for field diffs
	if membrane.Name != membraneOther.Name {
		diffs = append(diffs, membrane.GongMarshallField(stage, "Name"))
	}
	if membrane.Smufl != membraneOther.Smufl {
		diffs = append(diffs, membrane.GongMarshallField(stage, "Smufl"))
	}
	if membrane.EnclosedText != membraneOther.EnclosedText {
		diffs = append(diffs, membrane.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metal *Metal) GongDiff(stage *Stage, metalOther *Metal) (diffs []string) {
	// insertion point for field diffs
	if metal.Name != metalOther.Name {
		diffs = append(diffs, metal.GongMarshallField(stage, "Name"))
	}
	if metal.Smufl != metalOther.Smufl {
		diffs = append(diffs, metal.GongMarshallField(stage, "Smufl"))
	}
	if metal.EnclosedText != metalOther.EnclosedText {
		diffs = append(diffs, metal.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metronome *Metronome) GongDiff(stage *Stage, metronomeOther *Metronome) (diffs []string) {
	// insertion point for field diffs
	if metronome.Name != metronomeOther.Name {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Name"))
	}
	if metronome.Parentheses != metronomeOther.Parentheses {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Parentheses"))
	}
	if metronome.Default_x != metronomeOther.Default_x {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Default_x"))
	}
	if metronome.Default_y != metronomeOther.Default_y {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Default_y"))
	}
	if metronome.Relative_x != metronomeOther.Relative_x {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Relative_x"))
	}
	if metronome.Relative_y != metronomeOther.Relative_y {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Relative_y"))
	}
	if metronome.Font_family != metronomeOther.Font_family {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Font_family"))
	}
	if metronome.Font_style != metronomeOther.Font_style {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Font_style"))
	}
	if metronome.Font_size != metronomeOther.Font_size {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Font_size"))
	}
	if metronome.Font_weight != metronomeOther.Font_weight {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Font_weight"))
	}
	if metronome.Color != metronomeOther.Color {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Color"))
	}
	if metronome.Halign != metronomeOther.Halign {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Halign"))
	}
	if metronome.Valign != metronomeOther.Valign {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Valign"))
	}
	if metronome.Print_object != metronomeOther.Print_object {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Print_object"))
	}
	if metronome.Justify != metronomeOther.Justify {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Justify"))
	}
	if metronome.Id != metronomeOther.Id {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Id"))
	}
	if metronome.Beat_unit != metronomeOther.Beat_unit {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Beat_unit"))
	}
	if metronome.Beat_unit_dot != metronomeOther.Beat_unit_dot {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Beat_unit_dot"))
	}
	if metronome.Per_minute != metronomeOther.Per_minute {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Per_minute"))
	}
	if ops := __gong__diffSliceOfPointers(stage, metronome, "Beat_unit_tied", metronomeOther.Beat_unit_tied, metronome.Beat_unit_tied); ops != "" {
		diffs = append(diffs, ops)
	}
	if metronome.Metronome_arrows != metronomeOther.Metronome_arrows {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Metronome_arrows"))
	}
	if metronome.Metronome_relation != metronomeOther.Metronome_relation {
		diffs = append(diffs, metronome.GongMarshallField(stage, "Metronome_relation"))
	}
	if ops := __gong__diffSliceOfPointers(stage, metronome, "Metronome_note", metronomeOther.Metronome_note, metronome.Metronome_note); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metronome_beam *Metronome_beam) GongDiff(stage *Stage, metronome_beamOther *Metronome_beam) (diffs []string) {
	// insertion point for field diffs
	if metronome_beam.Name != metronome_beamOther.Name {
		diffs = append(diffs, metronome_beam.GongMarshallField(stage, "Name"))
	}
	if metronome_beam.Number != metronome_beamOther.Number {
		diffs = append(diffs, metronome_beam.GongMarshallField(stage, "Number"))
	}
	if metronome_beam.EnclosedText != metronome_beamOther.EnclosedText {
		diffs = append(diffs, metronome_beam.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metronome_note *Metronome_note) GongDiff(stage *Stage, metronome_noteOther *Metronome_note) (diffs []string) {
	// insertion point for field diffs
	if metronome_note.Name != metronome_noteOther.Name {
		diffs = append(diffs, metronome_note.GongMarshallField(stage, "Name"))
	}
	if metronome_note.Metronome_type != metronome_noteOther.Metronome_type {
		diffs = append(diffs, metronome_note.GongMarshallField(stage, "Metronome_type"))
	}
	if metronome_note.Metronome_dot != metronome_noteOther.Metronome_dot {
		diffs = append(diffs, metronome_note.GongMarshallField(stage, "Metronome_dot"))
	}
	if ops := __gong__diffSliceOfPointers(stage, metronome_note, "Metronome_beam", metronome_noteOther.Metronome_beam, metronome_note.Metronome_beam); ops != "" {
		diffs = append(diffs, ops)
	}
	if metronome_note.Metronome_tied != metronome_noteOther.Metronome_tied {
		diffs = append(diffs, metronome_note.GongMarshallField(stage, "Metronome_tied"))
	}
	if metronome_note.Metronome_tuplet != metronome_noteOther.Metronome_tuplet {
		diffs = append(diffs, metronome_note.GongMarshallField(stage, "Metronome_tuplet"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metronome_tied *Metronome_tied) GongDiff(stage *Stage, metronome_tiedOther *Metronome_tied) (diffs []string) {
	// insertion point for field diffs
	if metronome_tied.Name != metronome_tiedOther.Name {
		diffs = append(diffs, metronome_tied.GongMarshallField(stage, "Name"))
	}
	if metronome_tied.Type != metronome_tiedOther.Type {
		diffs = append(diffs, metronome_tied.GongMarshallField(stage, "Type"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metronome_tuplet *Metronome_tuplet) GongDiff(stage *Stage, metronome_tupletOther *Metronome_tuplet) (diffs []string) {
	// insertion point for field diffs
	if metronome_tuplet.Name != metronome_tupletOther.Name {
		diffs = append(diffs, metronome_tuplet.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (midi_device *Midi_device) GongDiff(stage *Stage, midi_deviceOther *Midi_device) (diffs []string) {
	// insertion point for field diffs
	if midi_device.Name != midi_deviceOther.Name {
		diffs = append(diffs, midi_device.GongMarshallField(stage, "Name"))
	}
	if midi_device.Port != midi_deviceOther.Port {
		diffs = append(diffs, midi_device.GongMarshallField(stage, "Port"))
	}
	if midi_device.Id != midi_deviceOther.Id {
		diffs = append(diffs, midi_device.GongMarshallField(stage, "Id"))
	}
	if midi_device.EnclosedText != midi_deviceOther.EnclosedText {
		diffs = append(diffs, midi_device.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (midi_instrument *Midi_instrument) GongDiff(stage *Stage, midi_instrumentOther *Midi_instrument) (diffs []string) {
	// insertion point for field diffs
	if midi_instrument.Name != midi_instrumentOther.Name {
		diffs = append(diffs, midi_instrument.GongMarshallField(stage, "Name"))
	}
	if midi_instrument.Id != midi_instrumentOther.Id {
		diffs = append(diffs, midi_instrument.GongMarshallField(stage, "Id"))
	}
	if midi_instrument.Midi_channel != midi_instrumentOther.Midi_channel {
		diffs = append(diffs, midi_instrument.GongMarshallField(stage, "Midi_channel"))
	}
	if midi_instrument.Midi_name != midi_instrumentOther.Midi_name {
		diffs = append(diffs, midi_instrument.GongMarshallField(stage, "Midi_name"))
	}
	if midi_instrument.Midi_bank != midi_instrumentOther.Midi_bank {
		diffs = append(diffs, midi_instrument.GongMarshallField(stage, "Midi_bank"))
	}
	if midi_instrument.Midi_program != midi_instrumentOther.Midi_program {
		diffs = append(diffs, midi_instrument.GongMarshallField(stage, "Midi_program"))
	}
	if midi_instrument.Midi_unpitched != midi_instrumentOther.Midi_unpitched {
		diffs = append(diffs, midi_instrument.GongMarshallField(stage, "Midi_unpitched"))
	}
	if midi_instrument.Volume != midi_instrumentOther.Volume {
		diffs = append(diffs, midi_instrument.GongMarshallField(stage, "Volume"))
	}
	if midi_instrument.Pan != midi_instrumentOther.Pan {
		diffs = append(diffs, midi_instrument.GongMarshallField(stage, "Pan"))
	}
	if midi_instrument.Elevation != midi_instrumentOther.Elevation {
		diffs = append(diffs, midi_instrument.GongMarshallField(stage, "Elevation"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (miscellaneous *Miscellaneous) GongDiff(stage *Stage, miscellaneousOther *Miscellaneous) (diffs []string) {
	// insertion point for field diffs
	if miscellaneous.Name != miscellaneousOther.Name {
		diffs = append(diffs, miscellaneous.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, miscellaneous, "Miscellaneous_field", miscellaneousOther.Miscellaneous_field, miscellaneous.Miscellaneous_field); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (miscellaneous_field *Miscellaneous_field) GongDiff(stage *Stage, miscellaneous_fieldOther *Miscellaneous_field) (diffs []string) {
	// insertion point for field diffs
	if miscellaneous_field.Name != miscellaneous_fieldOther.Name {
		diffs = append(diffs, miscellaneous_field.GongMarshallField(stage, "Name"))
	}
	if miscellaneous_field.NameXSD != miscellaneous_fieldOther.NameXSD {
		diffs = append(diffs, miscellaneous_field.GongMarshallField(stage, "NameXSD"))
	}
	if miscellaneous_field.EnclosedText != miscellaneous_fieldOther.EnclosedText {
		diffs = append(diffs, miscellaneous_field.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (mordent *Mordent) GongDiff(stage *Stage, mordentOther *Mordent) (diffs []string) {
	// insertion point for field diffs
	if mordent.Name != mordentOther.Name {
		diffs = append(diffs, mordent.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (multiple_rest *Multiple_rest) GongDiff(stage *Stage, multiple_restOther *Multiple_rest) (diffs []string) {
	// insertion point for field diffs
	if multiple_rest.Name != multiple_restOther.Name {
		diffs = append(diffs, multiple_rest.GongMarshallField(stage, "Name"))
	}
	if multiple_rest.Use_symbols != multiple_restOther.Use_symbols {
		diffs = append(diffs, multiple_rest.GongMarshallField(stage, "Use_symbols"))
	}
	if multiple_rest.EnclosedText != multiple_restOther.EnclosedText {
		diffs = append(diffs, multiple_rest.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (name_display *Name_display) GongDiff(stage *Stage, name_displayOther *Name_display) (diffs []string) {
	// insertion point for field diffs
	if name_display.Name != name_displayOther.Name {
		diffs = append(diffs, name_display.GongMarshallField(stage, "Name"))
	}
	if name_display.Print_object != name_displayOther.Print_object {
		diffs = append(diffs, name_display.GongMarshallField(stage, "Print_object"))
	}
	if ops := __gong__diffSliceOfPointers(stage, name_display, "Display_text", name_displayOther.Display_text, name_display.Display_text); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, name_display, "Accidental_text", name_displayOther.Accidental_text, name_display.Accidental_text); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (non_arpeggiate *Non_arpeggiate) GongDiff(stage *Stage, non_arpeggiateOther *Non_arpeggiate) (diffs []string) {
	// insertion point for field diffs
	if non_arpeggiate.Name != non_arpeggiateOther.Name {
		diffs = append(diffs, non_arpeggiate.GongMarshallField(stage, "Name"))
	}
	if non_arpeggiate.Type != non_arpeggiateOther.Type {
		diffs = append(diffs, non_arpeggiate.GongMarshallField(stage, "Type"))
	}
	if non_arpeggiate.Number != non_arpeggiateOther.Number {
		diffs = append(diffs, non_arpeggiate.GongMarshallField(stage, "Number"))
	}
	if non_arpeggiate.Default_x != non_arpeggiateOther.Default_x {
		diffs = append(diffs, non_arpeggiate.GongMarshallField(stage, "Default_x"))
	}
	if non_arpeggiate.Default_y != non_arpeggiateOther.Default_y {
		diffs = append(diffs, non_arpeggiate.GongMarshallField(stage, "Default_y"))
	}
	if non_arpeggiate.Relative_x != non_arpeggiateOther.Relative_x {
		diffs = append(diffs, non_arpeggiate.GongMarshallField(stage, "Relative_x"))
	}
	if non_arpeggiate.Relative_y != non_arpeggiateOther.Relative_y {
		diffs = append(diffs, non_arpeggiate.GongMarshallField(stage, "Relative_y"))
	}
	if non_arpeggiate.Placement != non_arpeggiateOther.Placement {
		diffs = append(diffs, non_arpeggiate.GongMarshallField(stage, "Placement"))
	}
	if non_arpeggiate.Color != non_arpeggiateOther.Color {
		diffs = append(diffs, non_arpeggiate.GongMarshallField(stage, "Color"))
	}
	if non_arpeggiate.Id != non_arpeggiateOther.Id {
		diffs = append(diffs, non_arpeggiate.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notations *Notations) GongDiff(stage *Stage, notationsOther *Notations) (diffs []string) {
	// insertion point for field diffs
	if notations.Name != notationsOther.Name {
		diffs = append(diffs, notations.GongMarshallField(stage, "Name"))
	}
	if notations.Print_object != notationsOther.Print_object {
		diffs = append(diffs, notations.GongMarshallField(stage, "Print_object"))
	}
	if notations.Id != notationsOther.Id {
		diffs = append(diffs, notations.GongMarshallField(stage, "Id"))
	}
	if notations.Footnote != notationsOther.Footnote {
		diffs = append(diffs, notations.GongMarshallField(stage, "Footnote"))
	}
	if notations.Level != notationsOther.Level {
		diffs = append(diffs, notations.GongMarshallField(stage, "Level"))
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Tied", notationsOther.Tied, notations.Tied); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Slur", notationsOther.Slur, notations.Slur); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Tuplet", notationsOther.Tuplet, notations.Tuplet); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Glissando", notationsOther.Glissando, notations.Glissando); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Slide", notationsOther.Slide, notations.Slide); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Ornaments", notationsOther.Ornaments, notations.Ornaments); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Technical", notationsOther.Technical, notations.Technical); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Articulations", notationsOther.Articulations, notations.Articulations); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Dynamics", notationsOther.Dynamics, notations.Dynamics); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Fermata", notationsOther.Fermata, notations.Fermata); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Arpeggiate", notationsOther.Arpeggiate, notations.Arpeggiate); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Non_arpeggiate", notationsOther.Non_arpeggiate, notations.Non_arpeggiate); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Accidental_mark", notationsOther.Accidental_mark, notations.Accidental_mark); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notations, "Other_notation", notationsOther.Other_notation, notations.Other_notation); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (note *Note) GongDiff(stage *Stage, noteOther *Note) (diffs []string) {
	// insertion point for field diffs
	if note.Name != noteOther.Name {
		diffs = append(diffs, note.GongMarshallField(stage, "Name"))
	}
	if note.Print_leger != noteOther.Print_leger {
		diffs = append(diffs, note.GongMarshallField(stage, "Print_leger"))
	}
	if note.Dynamics != noteOther.Dynamics {
		diffs = append(diffs, note.GongMarshallField(stage, "Dynamics"))
	}
	if note.End_dynamics != noteOther.End_dynamics {
		diffs = append(diffs, note.GongMarshallField(stage, "End_dynamics"))
	}
	if note.Attack != noteOther.Attack {
		diffs = append(diffs, note.GongMarshallField(stage, "Attack"))
	}
	if note.Release != noteOther.Release {
		diffs = append(diffs, note.GongMarshallField(stage, "Release"))
	}
	if note.Time_only != noteOther.Time_only {
		diffs = append(diffs, note.GongMarshallField(stage, "Time_only"))
	}
	if note.Pizzicato != noteOther.Pizzicato {
		diffs = append(diffs, note.GongMarshallField(stage, "Pizzicato"))
	}
	if note.Default_x != noteOther.Default_x {
		diffs = append(diffs, note.GongMarshallField(stage, "Default_x"))
	}
	if note.Default_y != noteOther.Default_y {
		diffs = append(diffs, note.GongMarshallField(stage, "Default_y"))
	}
	if note.Relative_x != noteOther.Relative_x {
		diffs = append(diffs, note.GongMarshallField(stage, "Relative_x"))
	}
	if note.Relative_y != noteOther.Relative_y {
		diffs = append(diffs, note.GongMarshallField(stage, "Relative_y"))
	}
	if note.Font_family != noteOther.Font_family {
		diffs = append(diffs, note.GongMarshallField(stage, "Font_family"))
	}
	if note.Font_style != noteOther.Font_style {
		diffs = append(diffs, note.GongMarshallField(stage, "Font_style"))
	}
	if note.Font_size != noteOther.Font_size {
		diffs = append(diffs, note.GongMarshallField(stage, "Font_size"))
	}
	if note.Font_weight != noteOther.Font_weight {
		diffs = append(diffs, note.GongMarshallField(stage, "Font_weight"))
	}
	if note.Color != noteOther.Color {
		diffs = append(diffs, note.GongMarshallField(stage, "Color"))
	}
	if note.Print_dot != noteOther.Print_dot {
		diffs = append(diffs, note.GongMarshallField(stage, "Print_dot"))
	}
	if note.Print_lyric != noteOther.Print_lyric {
		diffs = append(diffs, note.GongMarshallField(stage, "Print_lyric"))
	}
	if note.Print_object != noteOther.Print_object {
		diffs = append(diffs, note.GongMarshallField(stage, "Print_object"))
	}
	if note.Print_spacing != noteOther.Print_spacing {
		diffs = append(diffs, note.GongMarshallField(stage, "Print_spacing"))
	}
	if note.Id != noteOther.Id {
		diffs = append(diffs, note.GongMarshallField(stage, "Id"))
	}
	if note.Grace != noteOther.Grace {
		diffs = append(diffs, note.GongMarshallField(stage, "Grace"))
	}
	if note.Chord != noteOther.Chord {
		diffs = append(diffs, note.GongMarshallField(stage, "Chord"))
	}
	if note.Pitch != noteOther.Pitch {
		diffs = append(diffs, note.GongMarshallField(stage, "Pitch"))
	}
	if note.Unpitched != noteOther.Unpitched {
		diffs = append(diffs, note.GongMarshallField(stage, "Unpitched"))
	}
	if note.Rest != noteOther.Rest {
		diffs = append(diffs, note.GongMarshallField(stage, "Rest"))
	}
	if note.Tie != noteOther.Tie {
		diffs = append(diffs, note.GongMarshallField(stage, "Tie"))
	}
	if note.Cue != noteOther.Cue {
		diffs = append(diffs, note.GongMarshallField(stage, "Cue"))
	}
	if note.Duration != noteOther.Duration {
		diffs = append(diffs, note.GongMarshallField(stage, "Duration"))
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Instrument", noteOther.Instrument, note.Instrument); ops != "" {
		diffs = append(diffs, ops)
	}
	if note.Footnote != noteOther.Footnote {
		diffs = append(diffs, note.GongMarshallField(stage, "Footnote"))
	}
	if note.Level != noteOther.Level {
		diffs = append(diffs, note.GongMarshallField(stage, "Level"))
	}
	if note.Voice != noteOther.Voice {
		diffs = append(diffs, note.GongMarshallField(stage, "Voice"))
	}
	if note.Type != noteOther.Type {
		diffs = append(diffs, note.GongMarshallField(stage, "Type"))
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Dot", noteOther.Dot, note.Dot); ops != "" {
		diffs = append(diffs, ops)
	}
	if note.Accidental != noteOther.Accidental {
		diffs = append(diffs, note.GongMarshallField(stage, "Accidental"))
	}
	if note.Time_modification != noteOther.Time_modification {
		diffs = append(diffs, note.GongMarshallField(stage, "Time_modification"))
	}
	if note.Stem != noteOther.Stem {
		diffs = append(diffs, note.GongMarshallField(stage, "Stem"))
	}
	if note.Notehead != noteOther.Notehead {
		diffs = append(diffs, note.GongMarshallField(stage, "Notehead"))
	}
	if note.Notehead_text != noteOther.Notehead_text {
		diffs = append(diffs, note.GongMarshallField(stage, "Notehead_text"))
	}
	if note.Staff != noteOther.Staff {
		diffs = append(diffs, note.GongMarshallField(stage, "Staff"))
	}
	if note.Beam != noteOther.Beam {
		diffs = append(diffs, note.GongMarshallField(stage, "Beam"))
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Notations", noteOther.Notations, note.Notations); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Lyric", noteOther.Lyric, note.Lyric); ops != "" {
		diffs = append(diffs, ops)
	}
	if note.Play != noteOther.Play {
		diffs = append(diffs, note.GongMarshallField(stage, "Play"))
	}
	if note.Listen != noteOther.Listen {
		diffs = append(diffs, note.GongMarshallField(stage, "Listen"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (note_size *Note_size) GongDiff(stage *Stage, note_sizeOther *Note_size) (diffs []string) {
	// insertion point for field diffs
	if note_size.Name != note_sizeOther.Name {
		diffs = append(diffs, note_size.GongMarshallField(stage, "Name"))
	}
	if note_size.Type != note_sizeOther.Type {
		diffs = append(diffs, note_size.GongMarshallField(stage, "Type"))
	}
	if note_size.EnclosedText != note_sizeOther.EnclosedText {
		diffs = append(diffs, note_size.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (note_type *Note_type) GongDiff(stage *Stage, note_typeOther *Note_type) (diffs []string) {
	// insertion point for field diffs
	if note_type.Name != note_typeOther.Name {
		diffs = append(diffs, note_type.GongMarshallField(stage, "Name"))
	}
	if note_type.Size != note_typeOther.Size {
		diffs = append(diffs, note_type.GongMarshallField(stage, "Size"))
	}
	if note_type.EnclosedText != note_typeOther.EnclosedText {
		diffs = append(diffs, note_type.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notehead *Notehead) GongDiff(stage *Stage, noteheadOther *Notehead) (diffs []string) {
	// insertion point for field diffs
	if notehead.Name != noteheadOther.Name {
		diffs = append(diffs, notehead.GongMarshallField(stage, "Name"))
	}
	if notehead.Filled != noteheadOther.Filled {
		diffs = append(diffs, notehead.GongMarshallField(stage, "Filled"))
	}
	if notehead.Parentheses != noteheadOther.Parentheses {
		diffs = append(diffs, notehead.GongMarshallField(stage, "Parentheses"))
	}
	if notehead.Font_family != noteheadOther.Font_family {
		diffs = append(diffs, notehead.GongMarshallField(stage, "Font_family"))
	}
	if notehead.Font_style != noteheadOther.Font_style {
		diffs = append(diffs, notehead.GongMarshallField(stage, "Font_style"))
	}
	if notehead.Font_size != noteheadOther.Font_size {
		diffs = append(diffs, notehead.GongMarshallField(stage, "Font_size"))
	}
	if notehead.Font_weight != noteheadOther.Font_weight {
		diffs = append(diffs, notehead.GongMarshallField(stage, "Font_weight"))
	}
	if notehead.Color != noteheadOther.Color {
		diffs = append(diffs, notehead.GongMarshallField(stage, "Color"))
	}
	if notehead.Smufl != noteheadOther.Smufl {
		diffs = append(diffs, notehead.GongMarshallField(stage, "Smufl"))
	}
	if notehead.EnclosedText != noteheadOther.EnclosedText {
		diffs = append(diffs, notehead.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notehead_text *Notehead_text) GongDiff(stage *Stage, notehead_textOther *Notehead_text) (diffs []string) {
	// insertion point for field diffs
	if notehead_text.Name != notehead_textOther.Name {
		diffs = append(diffs, notehead_text.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, notehead_text, "Display_text", notehead_textOther.Display_text, notehead_text.Display_text); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, notehead_text, "Accidental_text", notehead_textOther.Accidental_text, notehead_text.Accidental_text); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (numeral *Numeral) GongDiff(stage *Stage, numeralOther *Numeral) (diffs []string) {
	// insertion point for field diffs
	if numeral.Name != numeralOther.Name {
		diffs = append(diffs, numeral.GongMarshallField(stage, "Name"))
	}
	if numeral.Numeral_root != numeralOther.Numeral_root {
		diffs = append(diffs, numeral.GongMarshallField(stage, "Numeral_root"))
	}
	if numeral.Numeral_alter != numeralOther.Numeral_alter {
		diffs = append(diffs, numeral.GongMarshallField(stage, "Numeral_alter"))
	}
	if numeral.Numeral_key != numeralOther.Numeral_key {
		diffs = append(diffs, numeral.GongMarshallField(stage, "Numeral_key"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (numeral_key *Numeral_key) GongDiff(stage *Stage, numeral_keyOther *Numeral_key) (diffs []string) {
	// insertion point for field diffs
	if numeral_key.Name != numeral_keyOther.Name {
		diffs = append(diffs, numeral_key.GongMarshallField(stage, "Name"))
	}
	if numeral_key.Print_object != numeral_keyOther.Print_object {
		diffs = append(diffs, numeral_key.GongMarshallField(stage, "Print_object"))
	}
	if numeral_key.Numeral_fifths != numeral_keyOther.Numeral_fifths {
		diffs = append(diffs, numeral_key.GongMarshallField(stage, "Numeral_fifths"))
	}
	if numeral_key.Numeral_mode != numeral_keyOther.Numeral_mode {
		diffs = append(diffs, numeral_key.GongMarshallField(stage, "Numeral_mode"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (numeral_root *Numeral_root) GongDiff(stage *Stage, numeral_rootOther *Numeral_root) (diffs []string) {
	// insertion point for field diffs
	if numeral_root.Name != numeral_rootOther.Name {
		diffs = append(diffs, numeral_root.GongMarshallField(stage, "Name"))
	}
	if numeral_root.Text != numeral_rootOther.Text {
		diffs = append(diffs, numeral_root.GongMarshallField(stage, "Text"))
	}
	if numeral_root.Default_x != numeral_rootOther.Default_x {
		diffs = append(diffs, numeral_root.GongMarshallField(stage, "Default_x"))
	}
	if numeral_root.Default_y != numeral_rootOther.Default_y {
		diffs = append(diffs, numeral_root.GongMarshallField(stage, "Default_y"))
	}
	if numeral_root.Relative_x != numeral_rootOther.Relative_x {
		diffs = append(diffs, numeral_root.GongMarshallField(stage, "Relative_x"))
	}
	if numeral_root.Relative_y != numeral_rootOther.Relative_y {
		diffs = append(diffs, numeral_root.GongMarshallField(stage, "Relative_y"))
	}
	if numeral_root.Font_family != numeral_rootOther.Font_family {
		diffs = append(diffs, numeral_root.GongMarshallField(stage, "Font_family"))
	}
	if numeral_root.Font_style != numeral_rootOther.Font_style {
		diffs = append(diffs, numeral_root.GongMarshallField(stage, "Font_style"))
	}
	if numeral_root.Font_size != numeral_rootOther.Font_size {
		diffs = append(diffs, numeral_root.GongMarshallField(stage, "Font_size"))
	}
	if numeral_root.Font_weight != numeral_rootOther.Font_weight {
		diffs = append(diffs, numeral_root.GongMarshallField(stage, "Font_weight"))
	}
	if numeral_root.Color != numeral_rootOther.Color {
		diffs = append(diffs, numeral_root.GongMarshallField(stage, "Color"))
	}
	if numeral_root.EnclosedText != numeral_rootOther.EnclosedText {
		diffs = append(diffs, numeral_root.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (octave_shift *Octave_shift) GongDiff(stage *Stage, octave_shiftOther *Octave_shift) (diffs []string) {
	// insertion point for field diffs
	if octave_shift.Name != octave_shiftOther.Name {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Name"))
	}
	if octave_shift.Type != octave_shiftOther.Type {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Type"))
	}
	if octave_shift.Number != octave_shiftOther.Number {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Number"))
	}
	if octave_shift.Size != octave_shiftOther.Size {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Size"))
	}
	if octave_shift.Dash_length != octave_shiftOther.Dash_length {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Dash_length"))
	}
	if octave_shift.Space_length != octave_shiftOther.Space_length {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Space_length"))
	}
	if octave_shift.Default_x != octave_shiftOther.Default_x {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Default_x"))
	}
	if octave_shift.Default_y != octave_shiftOther.Default_y {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Default_y"))
	}
	if octave_shift.Relative_x != octave_shiftOther.Relative_x {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Relative_x"))
	}
	if octave_shift.Relative_y != octave_shiftOther.Relative_y {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Relative_y"))
	}
	if octave_shift.Font_family != octave_shiftOther.Font_family {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Font_family"))
	}
	if octave_shift.Font_style != octave_shiftOther.Font_style {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Font_style"))
	}
	if octave_shift.Font_size != octave_shiftOther.Font_size {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Font_size"))
	}
	if octave_shift.Font_weight != octave_shiftOther.Font_weight {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Font_weight"))
	}
	if octave_shift.Color != octave_shiftOther.Color {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Color"))
	}
	if octave_shift.Id != octave_shiftOther.Id {
		diffs = append(diffs, octave_shift.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (offset *Offset) GongDiff(stage *Stage, offsetOther *Offset) (diffs []string) {
	// insertion point for field diffs
	if offset.Name != offsetOther.Name {
		diffs = append(diffs, offset.GongMarshallField(stage, "Name"))
	}
	if offset.Sound != offsetOther.Sound {
		diffs = append(diffs, offset.GongMarshallField(stage, "Sound"))
	}
	if offset.EnclosedText != offsetOther.EnclosedText {
		diffs = append(diffs, offset.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (opus *Opus) GongDiff(stage *Stage, opusOther *Opus) (diffs []string) {
	// insertion point for field diffs
	if opus.Name != opusOther.Name {
		diffs = append(diffs, opus.GongMarshallField(stage, "Name"))
	}
	if opus.Href != opusOther.Href {
		diffs = append(diffs, opus.GongMarshallField(stage, "Href"))
	}
	if opus.Type != opusOther.Type {
		diffs = append(diffs, opus.GongMarshallField(stage, "Type"))
	}
	if opus.Role != opusOther.Role {
		diffs = append(diffs, opus.GongMarshallField(stage, "Role"))
	}
	if opus.Title != opusOther.Title {
		diffs = append(diffs, opus.GongMarshallField(stage, "Title"))
	}
	if opus.Show != opusOther.Show {
		diffs = append(diffs, opus.GongMarshallField(stage, "Show"))
	}
	if opus.Actuate != opusOther.Actuate {
		diffs = append(diffs, opus.GongMarshallField(stage, "Actuate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ornaments *Ornaments) GongDiff(stage *Stage, ornamentsOther *Ornaments) (diffs []string) {
	// insertion point for field diffs
	if ornaments.Name != ornamentsOther.Name {
		diffs = append(diffs, ornaments.GongMarshallField(stage, "Name"))
	}
	if ornaments.Id != ornamentsOther.Id {
		diffs = append(diffs, ornaments.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Trill_mark", ornamentsOther.Trill_mark, ornaments.Trill_mark); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Turn", ornamentsOther.Turn, ornaments.Turn); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Delayed_turn", ornamentsOther.Delayed_turn, ornaments.Delayed_turn); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Inverted_turn", ornamentsOther.Inverted_turn, ornaments.Inverted_turn); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Delayed_inverted_turn", ornamentsOther.Delayed_inverted_turn, ornaments.Delayed_inverted_turn); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Vertical_turn", ornamentsOther.Vertical_turn, ornaments.Vertical_turn); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Inverted_vertical_turn", ornamentsOther.Inverted_vertical_turn, ornaments.Inverted_vertical_turn); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Shake", ornamentsOther.Shake, ornaments.Shake); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Wavy_line", ornamentsOther.Wavy_line, ornaments.Wavy_line); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Mordent", ornamentsOther.Mordent, ornaments.Mordent); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Inverted_mordent", ornamentsOther.Inverted_mordent, ornaments.Inverted_mordent); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Schleifer", ornamentsOther.Schleifer, ornaments.Schleifer); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Tremolo", ornamentsOther.Tremolo, ornaments.Tremolo); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Haydn", ornamentsOther.Haydn, ornaments.Haydn); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Other_ornament", ornamentsOther.Other_ornament, ornaments.Other_ornament); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, ornaments, "Accidental_mark", ornamentsOther.Accidental_mark, ornaments.Accidental_mark); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (other_appearance *Other_appearance) GongDiff(stage *Stage, other_appearanceOther *Other_appearance) (diffs []string) {
	// insertion point for field diffs
	if other_appearance.Name != other_appearanceOther.Name {
		diffs = append(diffs, other_appearance.GongMarshallField(stage, "Name"))
	}
	if other_appearance.Type != other_appearanceOther.Type {
		diffs = append(diffs, other_appearance.GongMarshallField(stage, "Type"))
	}
	if other_appearance.EnclosedText != other_appearanceOther.EnclosedText {
		diffs = append(diffs, other_appearance.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (other_direction *Other_direction) GongDiff(stage *Stage, other_directionOther *Other_direction) (diffs []string) {
	// insertion point for field diffs
	if other_direction.Name != other_directionOther.Name {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Name"))
	}
	if other_direction.Print_object != other_directionOther.Print_object {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Print_object"))
	}
	if other_direction.Default_x != other_directionOther.Default_x {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Default_x"))
	}
	if other_direction.Default_y != other_directionOther.Default_y {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Default_y"))
	}
	if other_direction.Relative_x != other_directionOther.Relative_x {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Relative_x"))
	}
	if other_direction.Relative_y != other_directionOther.Relative_y {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Relative_y"))
	}
	if other_direction.Font_family != other_directionOther.Font_family {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Font_family"))
	}
	if other_direction.Font_style != other_directionOther.Font_style {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Font_style"))
	}
	if other_direction.Font_size != other_directionOther.Font_size {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Font_size"))
	}
	if other_direction.Font_weight != other_directionOther.Font_weight {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Font_weight"))
	}
	if other_direction.Color != other_directionOther.Color {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Color"))
	}
	if other_direction.Halign != other_directionOther.Halign {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Halign"))
	}
	if other_direction.Valign != other_directionOther.Valign {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Valign"))
	}
	if other_direction.Smufl != other_directionOther.Smufl {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Smufl"))
	}
	if other_direction.Id != other_directionOther.Id {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "Id"))
	}
	if other_direction.EnclosedText != other_directionOther.EnclosedText {
		diffs = append(diffs, other_direction.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (other_listening *Other_listening) GongDiff(stage *Stage, other_listeningOther *Other_listening) (diffs []string) {
	// insertion point for field diffs
	if other_listening.Name != other_listeningOther.Name {
		diffs = append(diffs, other_listening.GongMarshallField(stage, "Name"))
	}
	if other_listening.Type != other_listeningOther.Type {
		diffs = append(diffs, other_listening.GongMarshallField(stage, "Type"))
	}
	if other_listening.Player != other_listeningOther.Player {
		diffs = append(diffs, other_listening.GongMarshallField(stage, "Player"))
	}
	if other_listening.Time_only != other_listeningOther.Time_only {
		diffs = append(diffs, other_listening.GongMarshallField(stage, "Time_only"))
	}
	if other_listening.EnclosedText != other_listeningOther.EnclosedText {
		diffs = append(diffs, other_listening.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (other_notation *Other_notation) GongDiff(stage *Stage, other_notationOther *Other_notation) (diffs []string) {
	// insertion point for field diffs
	if other_notation.Name != other_notationOther.Name {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Name"))
	}
	if other_notation.Type != other_notationOther.Type {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Type"))
	}
	if other_notation.Number != other_notationOther.Number {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Number"))
	}
	if other_notation.Print_object != other_notationOther.Print_object {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Print_object"))
	}
	if other_notation.Default_x != other_notationOther.Default_x {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Default_x"))
	}
	if other_notation.Default_y != other_notationOther.Default_y {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Default_y"))
	}
	if other_notation.Relative_x != other_notationOther.Relative_x {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Relative_x"))
	}
	if other_notation.Relative_y != other_notationOther.Relative_y {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Relative_y"))
	}
	if other_notation.Font_family != other_notationOther.Font_family {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Font_family"))
	}
	if other_notation.Font_style != other_notationOther.Font_style {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Font_style"))
	}
	if other_notation.Font_size != other_notationOther.Font_size {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Font_size"))
	}
	if other_notation.Font_weight != other_notationOther.Font_weight {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Font_weight"))
	}
	if other_notation.Color != other_notationOther.Color {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Color"))
	}
	if other_notation.Placement != other_notationOther.Placement {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Placement"))
	}
	if other_notation.Smufl != other_notationOther.Smufl {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Smufl"))
	}
	if other_notation.Id != other_notationOther.Id {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "Id"))
	}
	if other_notation.EnclosedText != other_notationOther.EnclosedText {
		diffs = append(diffs, other_notation.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (other_placement_text *Other_placement_text) GongDiff(stage *Stage, other_placement_textOther *Other_placement_text) (diffs []string) {
	// insertion point for field diffs
	if other_placement_text.Name != other_placement_textOther.Name {
		diffs = append(diffs, other_placement_text.GongMarshallField(stage, "Name"))
	}
	if other_placement_text.Default_x != other_placement_textOther.Default_x {
		diffs = append(diffs, other_placement_text.GongMarshallField(stage, "Default_x"))
	}
	if other_placement_text.Default_y != other_placement_textOther.Default_y {
		diffs = append(diffs, other_placement_text.GongMarshallField(stage, "Default_y"))
	}
	if other_placement_text.Relative_x != other_placement_textOther.Relative_x {
		diffs = append(diffs, other_placement_text.GongMarshallField(stage, "Relative_x"))
	}
	if other_placement_text.Relative_y != other_placement_textOther.Relative_y {
		diffs = append(diffs, other_placement_text.GongMarshallField(stage, "Relative_y"))
	}
	if other_placement_text.Font_family != other_placement_textOther.Font_family {
		diffs = append(diffs, other_placement_text.GongMarshallField(stage, "Font_family"))
	}
	if other_placement_text.Font_style != other_placement_textOther.Font_style {
		diffs = append(diffs, other_placement_text.GongMarshallField(stage, "Font_style"))
	}
	if other_placement_text.Font_size != other_placement_textOther.Font_size {
		diffs = append(diffs, other_placement_text.GongMarshallField(stage, "Font_size"))
	}
	if other_placement_text.Font_weight != other_placement_textOther.Font_weight {
		diffs = append(diffs, other_placement_text.GongMarshallField(stage, "Font_weight"))
	}
	if other_placement_text.Color != other_placement_textOther.Color {
		diffs = append(diffs, other_placement_text.GongMarshallField(stage, "Color"))
	}
	if other_placement_text.Placement != other_placement_textOther.Placement {
		diffs = append(diffs, other_placement_text.GongMarshallField(stage, "Placement"))
	}
	if other_placement_text.Smufl != other_placement_textOther.Smufl {
		diffs = append(diffs, other_placement_text.GongMarshallField(stage, "Smufl"))
	}
	if other_placement_text.EnclosedText != other_placement_textOther.EnclosedText {
		diffs = append(diffs, other_placement_text.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (other_play *Other_play) GongDiff(stage *Stage, other_playOther *Other_play) (diffs []string) {
	// insertion point for field diffs
	if other_play.Name != other_playOther.Name {
		diffs = append(diffs, other_play.GongMarshallField(stage, "Name"))
	}
	if other_play.Type != other_playOther.Type {
		diffs = append(diffs, other_play.GongMarshallField(stage, "Type"))
	}
	if other_play.EnclosedText != other_playOther.EnclosedText {
		diffs = append(diffs, other_play.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (other_text *Other_text) GongDiff(stage *Stage, other_textOther *Other_text) (diffs []string) {
	// insertion point for field diffs
	if other_text.Name != other_textOther.Name {
		diffs = append(diffs, other_text.GongMarshallField(stage, "Name"))
	}
	if other_text.Smufl != other_textOther.Smufl {
		diffs = append(diffs, other_text.GongMarshallField(stage, "Smufl"))
	}
	if other_text.EnclosedText != other_textOther.EnclosedText {
		diffs = append(diffs, other_text.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (page_layout *Page_layout) GongDiff(stage *Stage, page_layoutOther *Page_layout) (diffs []string) {
	// insertion point for field diffs
	if page_layout.Name != page_layoutOther.Name {
		diffs = append(diffs, page_layout.GongMarshallField(stage, "Name"))
	}
	if page_layout.Page_height != page_layoutOther.Page_height {
		diffs = append(diffs, page_layout.GongMarshallField(stage, "Page_height"))
	}
	if page_layout.Page_width != page_layoutOther.Page_width {
		diffs = append(diffs, page_layout.GongMarshallField(stage, "Page_width"))
	}
	if page_layout.Page_margins != page_layoutOther.Page_margins {
		diffs = append(diffs, page_layout.GongMarshallField(stage, "Page_margins"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (page_margins *Page_margins) GongDiff(stage *Stage, page_marginsOther *Page_margins) (diffs []string) {
	// insertion point for field diffs
	if page_margins.Name != page_marginsOther.Name {
		diffs = append(diffs, page_margins.GongMarshallField(stage, "Name"))
	}
	if page_margins.Type != page_marginsOther.Type {
		diffs = append(diffs, page_margins.GongMarshallField(stage, "Type"))
	}
	if page_margins.Left_margin != page_marginsOther.Left_margin {
		diffs = append(diffs, page_margins.GongMarshallField(stage, "Left_margin"))
	}
	if page_margins.Right_margin != page_marginsOther.Right_margin {
		diffs = append(diffs, page_margins.GongMarshallField(stage, "Right_margin"))
	}
	if page_margins.Top_margin != page_marginsOther.Top_margin {
		diffs = append(diffs, page_margins.GongMarshallField(stage, "Top_margin"))
	}
	if page_margins.Bottom_margin != page_marginsOther.Bottom_margin {
		diffs = append(diffs, page_margins.GongMarshallField(stage, "Bottom_margin"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (part_clef *Part_clef) GongDiff(stage *Stage, part_clefOther *Part_clef) (diffs []string) {
	// insertion point for field diffs
	if part_clef.Name != part_clefOther.Name {
		diffs = append(diffs, part_clef.GongMarshallField(stage, "Name"))
	}
	if part_clef.Sign != part_clefOther.Sign {
		diffs = append(diffs, part_clef.GongMarshallField(stage, "Sign"))
	}
	if part_clef.Line != part_clefOther.Line {
		diffs = append(diffs, part_clef.GongMarshallField(stage, "Line"))
	}
	if part_clef.Clef_octave_change != part_clefOther.Clef_octave_change {
		diffs = append(diffs, part_clef.GongMarshallField(stage, "Clef_octave_change"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (part_group *Part_group) GongDiff(stage *Stage, part_groupOther *Part_group) (diffs []string) {
	// insertion point for field diffs
	if part_group.Name != part_groupOther.Name {
		diffs = append(diffs, part_group.GongMarshallField(stage, "Name"))
	}
	if part_group.Type != part_groupOther.Type {
		diffs = append(diffs, part_group.GongMarshallField(stage, "Type"))
	}
	if part_group.Number != part_groupOther.Number {
		diffs = append(diffs, part_group.GongMarshallField(stage, "Number"))
	}
	if part_group.Group_name != part_groupOther.Group_name {
		diffs = append(diffs, part_group.GongMarshallField(stage, "Group_name"))
	}
	if part_group.Group_name_display != part_groupOther.Group_name_display {
		diffs = append(diffs, part_group.GongMarshallField(stage, "Group_name_display"))
	}
	if part_group.Group_abbreviation != part_groupOther.Group_abbreviation {
		diffs = append(diffs, part_group.GongMarshallField(stage, "Group_abbreviation"))
	}
	if part_group.Group_abbreviation_display != part_groupOther.Group_abbreviation_display {
		diffs = append(diffs, part_group.GongMarshallField(stage, "Group_abbreviation_display"))
	}
	if part_group.Group_symbol != part_groupOther.Group_symbol {
		diffs = append(diffs, part_group.GongMarshallField(stage, "Group_symbol"))
	}
	if part_group.Group_barline != part_groupOther.Group_barline {
		diffs = append(diffs, part_group.GongMarshallField(stage, "Group_barline"))
	}
	if part_group.Group_time != part_groupOther.Group_time {
		diffs = append(diffs, part_group.GongMarshallField(stage, "Group_time"))
	}
	if part_group.Footnote != part_groupOther.Footnote {
		diffs = append(diffs, part_group.GongMarshallField(stage, "Footnote"))
	}
	if part_group.Level != part_groupOther.Level {
		diffs = append(diffs, part_group.GongMarshallField(stage, "Level"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (part_link *Part_link) GongDiff(stage *Stage, part_linkOther *Part_link) (diffs []string) {
	// insertion point for field diffs
	if part_link.Name != part_linkOther.Name {
		diffs = append(diffs, part_link.GongMarshallField(stage, "Name"))
	}
	if part_link.Href != part_linkOther.Href {
		diffs = append(diffs, part_link.GongMarshallField(stage, "Href"))
	}
	if part_link.Type != part_linkOther.Type {
		diffs = append(diffs, part_link.GongMarshallField(stage, "Type"))
	}
	if part_link.Role != part_linkOther.Role {
		diffs = append(diffs, part_link.GongMarshallField(stage, "Role"))
	}
	if part_link.Title != part_linkOther.Title {
		diffs = append(diffs, part_link.GongMarshallField(stage, "Title"))
	}
	if part_link.Show != part_linkOther.Show {
		diffs = append(diffs, part_link.GongMarshallField(stage, "Show"))
	}
	if part_link.Actuate != part_linkOther.Actuate {
		diffs = append(diffs, part_link.GongMarshallField(stage, "Actuate"))
	}
	if ops := __gong__diffSliceOfPointers(stage, part_link, "Instrument_link", part_linkOther.Instrument_link, part_link.Instrument_link); ops != "" {
		diffs = append(diffs, ops)
	}
	if part_link.Group_link != part_linkOther.Group_link {
		diffs = append(diffs, part_link.GongMarshallField(stage, "Group_link"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (part_list *Part_list) GongDiff(stage *Stage, part_listOther *Part_list) (diffs []string) {
	// insertion point for field diffs
	if part_list.Name != part_listOther.Name {
		diffs = append(diffs, part_list.GongMarshallField(stage, "Name"))
	}
	if part_list.Part_group != part_listOther.Part_group {
		diffs = append(diffs, part_list.GongMarshallField(stage, "Part_group"))
	}
	if part_list.Score_part != part_listOther.Score_part {
		diffs = append(diffs, part_list.GongMarshallField(stage, "Score_part"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (part_name *Part_name) GongDiff(stage *Stage, part_nameOther *Part_name) (diffs []string) {
	// insertion point for field diffs
	if part_name.Name != part_nameOther.Name {
		diffs = append(diffs, part_name.GongMarshallField(stage, "Name"))
	}
	if part_name.Default_x != part_nameOther.Default_x {
		diffs = append(diffs, part_name.GongMarshallField(stage, "Default_x"))
	}
	if part_name.Default_y != part_nameOther.Default_y {
		diffs = append(diffs, part_name.GongMarshallField(stage, "Default_y"))
	}
	if part_name.Relative_x != part_nameOther.Relative_x {
		diffs = append(diffs, part_name.GongMarshallField(stage, "Relative_x"))
	}
	if part_name.Relative_y != part_nameOther.Relative_y {
		diffs = append(diffs, part_name.GongMarshallField(stage, "Relative_y"))
	}
	if part_name.Font_family != part_nameOther.Font_family {
		diffs = append(diffs, part_name.GongMarshallField(stage, "Font_family"))
	}
	if part_name.Font_style != part_nameOther.Font_style {
		diffs = append(diffs, part_name.GongMarshallField(stage, "Font_style"))
	}
	if part_name.Font_size != part_nameOther.Font_size {
		diffs = append(diffs, part_name.GongMarshallField(stage, "Font_size"))
	}
	if part_name.Font_weight != part_nameOther.Font_weight {
		diffs = append(diffs, part_name.GongMarshallField(stage, "Font_weight"))
	}
	if part_name.Color != part_nameOther.Color {
		diffs = append(diffs, part_name.GongMarshallField(stage, "Color"))
	}
	if part_name.Print_object != part_nameOther.Print_object {
		diffs = append(diffs, part_name.GongMarshallField(stage, "Print_object"))
	}
	if part_name.Justify != part_nameOther.Justify {
		diffs = append(diffs, part_name.GongMarshallField(stage, "Justify"))
	}
	if part_name.EnclosedText != part_nameOther.EnclosedText {
		diffs = append(diffs, part_name.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (part_symbol *Part_symbol) GongDiff(stage *Stage, part_symbolOther *Part_symbol) (diffs []string) {
	// insertion point for field diffs
	if part_symbol.Name != part_symbolOther.Name {
		diffs = append(diffs, part_symbol.GongMarshallField(stage, "Name"))
	}
	if part_symbol.Top_staff != part_symbolOther.Top_staff {
		diffs = append(diffs, part_symbol.GongMarshallField(stage, "Top_staff"))
	}
	if part_symbol.Bottom_staff != part_symbolOther.Bottom_staff {
		diffs = append(diffs, part_symbol.GongMarshallField(stage, "Bottom_staff"))
	}
	if part_symbol.Default_x != part_symbolOther.Default_x {
		diffs = append(diffs, part_symbol.GongMarshallField(stage, "Default_x"))
	}
	if part_symbol.Default_y != part_symbolOther.Default_y {
		diffs = append(diffs, part_symbol.GongMarshallField(stage, "Default_y"))
	}
	if part_symbol.Relative_x != part_symbolOther.Relative_x {
		diffs = append(diffs, part_symbol.GongMarshallField(stage, "Relative_x"))
	}
	if part_symbol.Relative_y != part_symbolOther.Relative_y {
		diffs = append(diffs, part_symbol.GongMarshallField(stage, "Relative_y"))
	}
	if part_symbol.Color != part_symbolOther.Color {
		diffs = append(diffs, part_symbol.GongMarshallField(stage, "Color"))
	}
	if part_symbol.EnclosedText != part_symbolOther.EnclosedText {
		diffs = append(diffs, part_symbol.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (part_transpose *Part_transpose) GongDiff(stage *Stage, part_transposeOther *Part_transpose) (diffs []string) {
	// insertion point for field diffs
	if part_transpose.Name != part_transposeOther.Name {
		diffs = append(diffs, part_transpose.GongMarshallField(stage, "Name"))
	}
	if part_transpose.Diatonic != part_transposeOther.Diatonic {
		diffs = append(diffs, part_transpose.GongMarshallField(stage, "Diatonic"))
	}
	if part_transpose.Chromatic != part_transposeOther.Chromatic {
		diffs = append(diffs, part_transpose.GongMarshallField(stage, "Chromatic"))
	}
	if part_transpose.Octave_change != part_transposeOther.Octave_change {
		diffs = append(diffs, part_transpose.GongMarshallField(stage, "Octave_change"))
	}
	if part_transpose.Double != part_transposeOther.Double {
		diffs = append(diffs, part_transpose.GongMarshallField(stage, "Double"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pedal *Pedal) GongDiff(stage *Stage, pedalOther *Pedal) (diffs []string) {
	// insertion point for field diffs
	if pedal.Name != pedalOther.Name {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Name"))
	}
	if pedal.Type != pedalOther.Type {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Type"))
	}
	if pedal.Number != pedalOther.Number {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Number"))
	}
	if pedal.Line != pedalOther.Line {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Line"))
	}
	if pedal.Sign != pedalOther.Sign {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Sign"))
	}
	if pedal.Abbreviated != pedalOther.Abbreviated {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Abbreviated"))
	}
	if pedal.Default_x != pedalOther.Default_x {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Default_x"))
	}
	if pedal.Default_y != pedalOther.Default_y {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Default_y"))
	}
	if pedal.Relative_x != pedalOther.Relative_x {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Relative_x"))
	}
	if pedal.Relative_y != pedalOther.Relative_y {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Relative_y"))
	}
	if pedal.Font_family != pedalOther.Font_family {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Font_family"))
	}
	if pedal.Font_style != pedalOther.Font_style {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Font_style"))
	}
	if pedal.Font_size != pedalOther.Font_size {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Font_size"))
	}
	if pedal.Font_weight != pedalOther.Font_weight {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Font_weight"))
	}
	if pedal.Color != pedalOther.Color {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Color"))
	}
	if pedal.Halign != pedalOther.Halign {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Halign"))
	}
	if pedal.Valign != pedalOther.Valign {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Valign"))
	}
	if pedal.Id != pedalOther.Id {
		diffs = append(diffs, pedal.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pedal_tuning *Pedal_tuning) GongDiff(stage *Stage, pedal_tuningOther *Pedal_tuning) (diffs []string) {
	// insertion point for field diffs
	if pedal_tuning.Name != pedal_tuningOther.Name {
		diffs = append(diffs, pedal_tuning.GongMarshallField(stage, "Name"))
	}
	if pedal_tuning.Pedal_step != pedal_tuningOther.Pedal_step {
		diffs = append(diffs, pedal_tuning.GongMarshallField(stage, "Pedal_step"))
	}
	if pedal_tuning.Pedal_alter != pedal_tuningOther.Pedal_alter {
		diffs = append(diffs, pedal_tuning.GongMarshallField(stage, "Pedal_alter"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (per_minute *Per_minute) GongDiff(stage *Stage, per_minuteOther *Per_minute) (diffs []string) {
	// insertion point for field diffs
	if per_minute.Name != per_minuteOther.Name {
		diffs = append(diffs, per_minute.GongMarshallField(stage, "Name"))
	}
	if per_minute.Font_family != per_minuteOther.Font_family {
		diffs = append(diffs, per_minute.GongMarshallField(stage, "Font_family"))
	}
	if per_minute.Font_style != per_minuteOther.Font_style {
		diffs = append(diffs, per_minute.GongMarshallField(stage, "Font_style"))
	}
	if per_minute.Font_size != per_minuteOther.Font_size {
		diffs = append(diffs, per_minute.GongMarshallField(stage, "Font_size"))
	}
	if per_minute.Font_weight != per_minuteOther.Font_weight {
		diffs = append(diffs, per_minute.GongMarshallField(stage, "Font_weight"))
	}
	if per_minute.EnclosedText != per_minuteOther.EnclosedText {
		diffs = append(diffs, per_minute.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (percussion *Percussion) GongDiff(stage *Stage, percussionOther *Percussion) (diffs []string) {
	// insertion point for field diffs
	if percussion.Name != percussionOther.Name {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Name"))
	}
	if percussion.Default_x != percussionOther.Default_x {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Default_x"))
	}
	if percussion.Default_y != percussionOther.Default_y {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Default_y"))
	}
	if percussion.Relative_x != percussionOther.Relative_x {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Relative_x"))
	}
	if percussion.Relative_y != percussionOther.Relative_y {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Relative_y"))
	}
	if percussion.Font_family != percussionOther.Font_family {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Font_family"))
	}
	if percussion.Font_style != percussionOther.Font_style {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Font_style"))
	}
	if percussion.Font_size != percussionOther.Font_size {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Font_size"))
	}
	if percussion.Font_weight != percussionOther.Font_weight {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Font_weight"))
	}
	if percussion.Color != percussionOther.Color {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Color"))
	}
	if percussion.Halign != percussionOther.Halign {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Halign"))
	}
	if percussion.Valign != percussionOther.Valign {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Valign"))
	}
	if percussion.Enclosure != percussionOther.Enclosure {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Enclosure"))
	}
	if percussion.Id != percussionOther.Id {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Id"))
	}
	if percussion.Glass != percussionOther.Glass {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Glass"))
	}
	if percussion.Metal != percussionOther.Metal {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Metal"))
	}
	if percussion.Wood != percussionOther.Wood {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Wood"))
	}
	if percussion.Pitched != percussionOther.Pitched {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Pitched"))
	}
	if percussion.Membrane != percussionOther.Membrane {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Membrane"))
	}
	if percussion.Effect != percussionOther.Effect {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Effect"))
	}
	if percussion.Timpani != percussionOther.Timpani {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Timpani"))
	}
	if percussion.Beater != percussionOther.Beater {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Beater"))
	}
	if percussion.Stick != percussionOther.Stick {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Stick"))
	}
	if percussion.Stick_location != percussionOther.Stick_location {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Stick_location"))
	}
	if percussion.Other_percussion != percussionOther.Other_percussion {
		diffs = append(diffs, percussion.GongMarshallField(stage, "Other_percussion"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pitch *Pitch) GongDiff(stage *Stage, pitchOther *Pitch) (diffs []string) {
	// insertion point for field diffs
	if pitch.Name != pitchOther.Name {
		diffs = append(diffs, pitch.GongMarshallField(stage, "Name"))
	}
	if pitch.Step != pitchOther.Step {
		diffs = append(diffs, pitch.GongMarshallField(stage, "Step"))
	}
	if pitch.Alter != pitchOther.Alter {
		diffs = append(diffs, pitch.GongMarshallField(stage, "Alter"))
	}
	if pitch.Octave != pitchOther.Octave {
		diffs = append(diffs, pitch.GongMarshallField(stage, "Octave"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pitched *Pitched) GongDiff(stage *Stage, pitchedOther *Pitched) (diffs []string) {
	// insertion point for field diffs
	if pitched.Name != pitchedOther.Name {
		diffs = append(diffs, pitched.GongMarshallField(stage, "Name"))
	}
	if pitched.Smufl != pitchedOther.Smufl {
		diffs = append(diffs, pitched.GongMarshallField(stage, "Smufl"))
	}
	if pitched.EnclosedText != pitchedOther.EnclosedText {
		diffs = append(diffs, pitched.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (placement_text *Placement_text) GongDiff(stage *Stage, placement_textOther *Placement_text) (diffs []string) {
	// insertion point for field diffs
	if placement_text.Name != placement_textOther.Name {
		diffs = append(diffs, placement_text.GongMarshallField(stage, "Name"))
	}
	if placement_text.Default_x != placement_textOther.Default_x {
		diffs = append(diffs, placement_text.GongMarshallField(stage, "Default_x"))
	}
	if placement_text.Default_y != placement_textOther.Default_y {
		diffs = append(diffs, placement_text.GongMarshallField(stage, "Default_y"))
	}
	if placement_text.Relative_x != placement_textOther.Relative_x {
		diffs = append(diffs, placement_text.GongMarshallField(stage, "Relative_x"))
	}
	if placement_text.Relative_y != placement_textOther.Relative_y {
		diffs = append(diffs, placement_text.GongMarshallField(stage, "Relative_y"))
	}
	if placement_text.Font_family != placement_textOther.Font_family {
		diffs = append(diffs, placement_text.GongMarshallField(stage, "Font_family"))
	}
	if placement_text.Font_style != placement_textOther.Font_style {
		diffs = append(diffs, placement_text.GongMarshallField(stage, "Font_style"))
	}
	if placement_text.Font_size != placement_textOther.Font_size {
		diffs = append(diffs, placement_text.GongMarshallField(stage, "Font_size"))
	}
	if placement_text.Font_weight != placement_textOther.Font_weight {
		diffs = append(diffs, placement_text.GongMarshallField(stage, "Font_weight"))
	}
	if placement_text.Color != placement_textOther.Color {
		diffs = append(diffs, placement_text.GongMarshallField(stage, "Color"))
	}
	if placement_text.Placement != placement_textOther.Placement {
		diffs = append(diffs, placement_text.GongMarshallField(stage, "Placement"))
	}
	if placement_text.EnclosedText != placement_textOther.EnclosedText {
		diffs = append(diffs, placement_text.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (play *Play) GongDiff(stage *Stage, playOther *Play) (diffs []string) {
	// insertion point for field diffs
	if play.Name != playOther.Name {
		diffs = append(diffs, play.GongMarshallField(stage, "Name"))
	}
	if play.Id != playOther.Id {
		diffs = append(diffs, play.GongMarshallField(stage, "Id"))
	}
	if play.Ipa != playOther.Ipa {
		diffs = append(diffs, play.GongMarshallField(stage, "Ipa"))
	}
	if play.Mute != playOther.Mute {
		diffs = append(diffs, play.GongMarshallField(stage, "Mute"))
	}
	if play.Semi_pitched != playOther.Semi_pitched {
		diffs = append(diffs, play.GongMarshallField(stage, "Semi_pitched"))
	}
	if ops := __gong__diffSliceOfPointers(stage, play, "Other_play", playOther.Other_play, play.Other_play); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (player *Player) GongDiff(stage *Stage, playerOther *Player) (diffs []string) {
	// insertion point for field diffs
	if player.Name != playerOther.Name {
		diffs = append(diffs, player.GongMarshallField(stage, "Name"))
	}
	if player.Id != playerOther.Id {
		diffs = append(diffs, player.GongMarshallField(stage, "Id"))
	}
	if player.Player_name != playerOther.Player_name {
		diffs = append(diffs, player.GongMarshallField(stage, "Player_name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (principal_voice *Principal_voice) GongDiff(stage *Stage, principal_voiceOther *Principal_voice) (diffs []string) {
	// insertion point for field diffs
	if principal_voice.Name != principal_voiceOther.Name {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Name"))
	}
	if principal_voice.Type != principal_voiceOther.Type {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Type"))
	}
	if principal_voice.Symbol != principal_voiceOther.Symbol {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Symbol"))
	}
	if principal_voice.Default_x != principal_voiceOther.Default_x {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Default_x"))
	}
	if principal_voice.Default_y != principal_voiceOther.Default_y {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Default_y"))
	}
	if principal_voice.Relative_x != principal_voiceOther.Relative_x {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Relative_x"))
	}
	if principal_voice.Relative_y != principal_voiceOther.Relative_y {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Relative_y"))
	}
	if principal_voice.Font_family != principal_voiceOther.Font_family {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Font_family"))
	}
	if principal_voice.Font_style != principal_voiceOther.Font_style {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Font_style"))
	}
	if principal_voice.Font_size != principal_voiceOther.Font_size {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Font_size"))
	}
	if principal_voice.Font_weight != principal_voiceOther.Font_weight {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Font_weight"))
	}
	if principal_voice.Color != principal_voiceOther.Color {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Color"))
	}
	if principal_voice.Halign != principal_voiceOther.Halign {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Halign"))
	}
	if principal_voice.Valign != principal_voiceOther.Valign {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Valign"))
	}
	if principal_voice.Id != principal_voiceOther.Id {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "Id"))
	}
	if principal_voice.EnclosedText != principal_voiceOther.EnclosedText {
		diffs = append(diffs, principal_voice.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (print *Print) GongDiff(stage *Stage, printOther *Print) (diffs []string) {
	// insertion point for field diffs
	if print.Name != printOther.Name {
		diffs = append(diffs, print.GongMarshallField(stage, "Name"))
	}
	if print.Staff_spacing != printOther.Staff_spacing {
		diffs = append(diffs, print.GongMarshallField(stage, "Staff_spacing"))
	}
	if print.New_system != printOther.New_system {
		diffs = append(diffs, print.GongMarshallField(stage, "New_system"))
	}
	if print.New_page != printOther.New_page {
		diffs = append(diffs, print.GongMarshallField(stage, "New_page"))
	}
	if print.Blank_page != printOther.Blank_page {
		diffs = append(diffs, print.GongMarshallField(stage, "Blank_page"))
	}
	if print.Page_number != printOther.Page_number {
		diffs = append(diffs, print.GongMarshallField(stage, "Page_number"))
	}
	if print.Id != printOther.Id {
		diffs = append(diffs, print.GongMarshallField(stage, "Id"))
	}
	if print.Page_layout != printOther.Page_layout {
		diffs = append(diffs, print.GongMarshallField(stage, "Page_layout"))
	}
	if print.System_layout != printOther.System_layout {
		diffs = append(diffs, print.GongMarshallField(stage, "System_layout"))
	}
	if ops := __gong__diffSliceOfPointers(stage, print, "Staff_layout", printOther.Staff_layout, print.Staff_layout); ops != "" {
		diffs = append(diffs, ops)
	}
	if print.Measure_layout != printOther.Measure_layout {
		diffs = append(diffs, print.GongMarshallField(stage, "Measure_layout"))
	}
	if print.Measure_numbering != printOther.Measure_numbering {
		diffs = append(diffs, print.GongMarshallField(stage, "Measure_numbering"))
	}
	if print.Part_name_display != printOther.Part_name_display {
		diffs = append(diffs, print.GongMarshallField(stage, "Part_name_display"))
	}
	if print.Part_abbreviation_display != printOther.Part_abbreviation_display {
		diffs = append(diffs, print.GongMarshallField(stage, "Part_abbreviation_display"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (release *Release) GongDiff(stage *Stage, releaseOther *Release) (diffs []string) {
	// insertion point for field diffs
	if release.Name != releaseOther.Name {
		diffs = append(diffs, release.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (repeat *Repeat) GongDiff(stage *Stage, repeatOther *Repeat) (diffs []string) {
	// insertion point for field diffs
	if repeat.Name != repeatOther.Name {
		diffs = append(diffs, repeat.GongMarshallField(stage, "Name"))
	}
	if repeat.Direction != repeatOther.Direction {
		diffs = append(diffs, repeat.GongMarshallField(stage, "Direction"))
	}
	if repeat.Times != repeatOther.Times {
		diffs = append(diffs, repeat.GongMarshallField(stage, "Times"))
	}
	if repeat.After_jump != repeatOther.After_jump {
		diffs = append(diffs, repeat.GongMarshallField(stage, "After_jump"))
	}
	if repeat.Winged != repeatOther.Winged {
		diffs = append(diffs, repeat.GongMarshallField(stage, "Winged"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rest *Rest) GongDiff(stage *Stage, restOther *Rest) (diffs []string) {
	// insertion point for field diffs
	if rest.Name != restOther.Name {
		diffs = append(diffs, rest.GongMarshallField(stage, "Name"))
	}
	if rest.Measure != restOther.Measure {
		diffs = append(diffs, rest.GongMarshallField(stage, "Measure"))
	}
	if rest.Display_step != restOther.Display_step {
		diffs = append(diffs, rest.GongMarshallField(stage, "Display_step"))
	}
	if rest.Display_octave != restOther.Display_octave {
		diffs = append(diffs, rest.GongMarshallField(stage, "Display_octave"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (root *Root) GongDiff(stage *Stage, rootOther *Root) (diffs []string) {
	// insertion point for field diffs
	if root.Name != rootOther.Name {
		diffs = append(diffs, root.GongMarshallField(stage, "Name"))
	}
	if root.Root_step != rootOther.Root_step {
		diffs = append(diffs, root.GongMarshallField(stage, "Root_step"))
	}
	if root.Root_alter != rootOther.Root_alter {
		diffs = append(diffs, root.GongMarshallField(stage, "Root_alter"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (root_step *Root_step) GongDiff(stage *Stage, root_stepOther *Root_step) (diffs []string) {
	// insertion point for field diffs
	if root_step.Name != root_stepOther.Name {
		diffs = append(diffs, root_step.GongMarshallField(stage, "Name"))
	}
	if root_step.Text != root_stepOther.Text {
		diffs = append(diffs, root_step.GongMarshallField(stage, "Text"))
	}
	if root_step.Default_x != root_stepOther.Default_x {
		diffs = append(diffs, root_step.GongMarshallField(stage, "Default_x"))
	}
	if root_step.Default_y != root_stepOther.Default_y {
		diffs = append(diffs, root_step.GongMarshallField(stage, "Default_y"))
	}
	if root_step.Relative_x != root_stepOther.Relative_x {
		diffs = append(diffs, root_step.GongMarshallField(stage, "Relative_x"))
	}
	if root_step.Relative_y != root_stepOther.Relative_y {
		diffs = append(diffs, root_step.GongMarshallField(stage, "Relative_y"))
	}
	if root_step.Font_family != root_stepOther.Font_family {
		diffs = append(diffs, root_step.GongMarshallField(stage, "Font_family"))
	}
	if root_step.Font_style != root_stepOther.Font_style {
		diffs = append(diffs, root_step.GongMarshallField(stage, "Font_style"))
	}
	if root_step.Font_size != root_stepOther.Font_size {
		diffs = append(diffs, root_step.GongMarshallField(stage, "Font_size"))
	}
	if root_step.Font_weight != root_stepOther.Font_weight {
		diffs = append(diffs, root_step.GongMarshallField(stage, "Font_weight"))
	}
	if root_step.Color != root_stepOther.Color {
		diffs = append(diffs, root_step.GongMarshallField(stage, "Color"))
	}
	if root_step.EnclosedText != root_stepOther.EnclosedText {
		diffs = append(diffs, root_step.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (scaling *Scaling) GongDiff(stage *Stage, scalingOther *Scaling) (diffs []string) {
	// insertion point for field diffs
	if scaling.Name != scalingOther.Name {
		diffs = append(diffs, scaling.GongMarshallField(stage, "Name"))
	}
	if scaling.Millimeters != scalingOther.Millimeters {
		diffs = append(diffs, scaling.GongMarshallField(stage, "Millimeters"))
	}
	if scaling.Tenths != scalingOther.Tenths {
		diffs = append(diffs, scaling.GongMarshallField(stage, "Tenths"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (scordatura *Scordatura) GongDiff(stage *Stage, scordaturaOther *Scordatura) (diffs []string) {
	// insertion point for field diffs
	if scordatura.Name != scordaturaOther.Name {
		diffs = append(diffs, scordatura.GongMarshallField(stage, "Name"))
	}
	if scordatura.Id != scordaturaOther.Id {
		diffs = append(diffs, scordatura.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, scordatura, "Accord", scordaturaOther.Accord, scordatura.Accord); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (score_instrument *Score_instrument) GongDiff(stage *Stage, score_instrumentOther *Score_instrument) (diffs []string) {
	// insertion point for field diffs
	if score_instrument.Name != score_instrumentOther.Name {
		diffs = append(diffs, score_instrument.GongMarshallField(stage, "Name"))
	}
	if score_instrument.Id != score_instrumentOther.Id {
		diffs = append(diffs, score_instrument.GongMarshallField(stage, "Id"))
	}
	if score_instrument.Instrument_name != score_instrumentOther.Instrument_name {
		diffs = append(diffs, score_instrument.GongMarshallField(stage, "Instrument_name"))
	}
	if score_instrument.Instrument_abbreviation != score_instrumentOther.Instrument_abbreviation {
		diffs = append(diffs, score_instrument.GongMarshallField(stage, "Instrument_abbreviation"))
	}
	if score_instrument.Instrument_sound != score_instrumentOther.Instrument_sound {
		diffs = append(diffs, score_instrument.GongMarshallField(stage, "Instrument_sound"))
	}
	if score_instrument.Solo != score_instrumentOther.Solo {
		diffs = append(diffs, score_instrument.GongMarshallField(stage, "Solo"))
	}
	if score_instrument.Ensemble != score_instrumentOther.Ensemble {
		diffs = append(diffs, score_instrument.GongMarshallField(stage, "Ensemble"))
	}
	if score_instrument.Virtual_instrument != score_instrumentOther.Virtual_instrument {
		diffs = append(diffs, score_instrument.GongMarshallField(stage, "Virtual_instrument"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (score_part *Score_part) GongDiff(stage *Stage, score_partOther *Score_part) (diffs []string) {
	// insertion point for field diffs
	if score_part.Name != score_partOther.Name {
		diffs = append(diffs, score_part.GongMarshallField(stage, "Name"))
	}
	if score_part.Id != score_partOther.Id {
		diffs = append(diffs, score_part.GongMarshallField(stage, "Id"))
	}
	if score_part.Identification != score_partOther.Identification {
		diffs = append(diffs, score_part.GongMarshallField(stage, "Identification"))
	}
	if ops := __gong__diffSliceOfPointers(stage, score_part, "Part_link", score_partOther.Part_link, score_part.Part_link); ops != "" {
		diffs = append(diffs, ops)
	}
	if score_part.Part_name != score_partOther.Part_name {
		diffs = append(diffs, score_part.GongMarshallField(stage, "Part_name"))
	}
	if score_part.Part_name_display != score_partOther.Part_name_display {
		diffs = append(diffs, score_part.GongMarshallField(stage, "Part_name_display"))
	}
	if score_part.Part_abbreviation != score_partOther.Part_abbreviation {
		diffs = append(diffs, score_part.GongMarshallField(stage, "Part_abbreviation"))
	}
	if score_part.Part_abbreviation_display != score_partOther.Part_abbreviation_display {
		diffs = append(diffs, score_part.GongMarshallField(stage, "Part_abbreviation_display"))
	}
	if score_part.Group != score_partOther.Group {
		diffs = append(diffs, score_part.GongMarshallField(stage, "Group"))
	}
	if ops := __gong__diffSliceOfPointers(stage, score_part, "Score_instrument", score_partOther.Score_instrument, score_part.Score_instrument); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, score_part, "Player", score_partOther.Player, score_part.Player); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, score_part, "Midi_device", score_partOther.Midi_device, score_part.Midi_device); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, score_part, "Midi_instrument", score_partOther.Midi_instrument, score_part.Midi_instrument); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (score_partwise *Score_partwise) GongDiff(stage *Stage, score_partwiseOther *Score_partwise) (diffs []string) {
	// insertion point for field diffs
	if score_partwise.Name != score_partwiseOther.Name {
		diffs = append(diffs, score_partwise.GongMarshallField(stage, "Name"))
	}
	if score_partwise.Version != score_partwiseOther.Version {
		diffs = append(diffs, score_partwise.GongMarshallField(stage, "Version"))
	}
	if score_partwise.Work != score_partwiseOther.Work {
		diffs = append(diffs, score_partwise.GongMarshallField(stage, "Work"))
	}
	if score_partwise.Movement_number != score_partwiseOther.Movement_number {
		diffs = append(diffs, score_partwise.GongMarshallField(stage, "Movement_number"))
	}
	if score_partwise.Movement_title != score_partwiseOther.Movement_title {
		diffs = append(diffs, score_partwise.GongMarshallField(stage, "Movement_title"))
	}
	if score_partwise.Identification != score_partwiseOther.Identification {
		diffs = append(diffs, score_partwise.GongMarshallField(stage, "Identification"))
	}
	if score_partwise.Defaults != score_partwiseOther.Defaults {
		diffs = append(diffs, score_partwise.GongMarshallField(stage, "Defaults"))
	}
	if ops := __gong__diffSliceOfPointers(stage, score_partwise, "Credit", score_partwiseOther.Credit, score_partwise.Credit); ops != "" {
		diffs = append(diffs, ops)
	}
	if score_partwise.Part_list != score_partwiseOther.Part_list {
		diffs = append(diffs, score_partwise.GongMarshallField(stage, "Part_list"))
	}
	if ops := __gong__diffSliceOfPointers(stage, score_partwise, "Part", score_partwiseOther.Part, score_partwise.Part); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (score_timewise *Score_timewise) GongDiff(stage *Stage, score_timewiseOther *Score_timewise) (diffs []string) {
	// insertion point for field diffs
	if score_timewise.Name != score_timewiseOther.Name {
		diffs = append(diffs, score_timewise.GongMarshallField(stage, "Name"))
	}
	if score_timewise.Version != score_timewiseOther.Version {
		diffs = append(diffs, score_timewise.GongMarshallField(stage, "Version"))
	}
	if score_timewise.Work != score_timewiseOther.Work {
		diffs = append(diffs, score_timewise.GongMarshallField(stage, "Work"))
	}
	if score_timewise.Movement_number != score_timewiseOther.Movement_number {
		diffs = append(diffs, score_timewise.GongMarshallField(stage, "Movement_number"))
	}
	if score_timewise.Movement_title != score_timewiseOther.Movement_title {
		diffs = append(diffs, score_timewise.GongMarshallField(stage, "Movement_title"))
	}
	if score_timewise.Identification != score_timewiseOther.Identification {
		diffs = append(diffs, score_timewise.GongMarshallField(stage, "Identification"))
	}
	if score_timewise.Defaults != score_timewiseOther.Defaults {
		diffs = append(diffs, score_timewise.GongMarshallField(stage, "Defaults"))
	}
	if ops := __gong__diffSliceOfPointers(stage, score_timewise, "Credit", score_timewiseOther.Credit, score_timewise.Credit); ops != "" {
		diffs = append(diffs, ops)
	}
	if score_timewise.Part_list != score_timewiseOther.Part_list {
		diffs = append(diffs, score_timewise.GongMarshallField(stage, "Part_list"))
	}
	if ops := __gong__diffSliceOfPointers(stage, score_timewise, "Measure", score_timewiseOther.Measure, score_timewise.Measure); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (segno *Segno) GongDiff(stage *Stage, segnoOther *Segno) (diffs []string) {
	// insertion point for field diffs
	if segno.Name != segnoOther.Name {
		diffs = append(diffs, segno.GongMarshallField(stage, "Name"))
	}
	if segno.Smufl != segnoOther.Smufl {
		diffs = append(diffs, segno.GongMarshallField(stage, "Smufl"))
	}
	if segno.Default_x != segnoOther.Default_x {
		diffs = append(diffs, segno.GongMarshallField(stage, "Default_x"))
	}
	if segno.Default_y != segnoOther.Default_y {
		diffs = append(diffs, segno.GongMarshallField(stage, "Default_y"))
	}
	if segno.Relative_x != segnoOther.Relative_x {
		diffs = append(diffs, segno.GongMarshallField(stage, "Relative_x"))
	}
	if segno.Relative_y != segnoOther.Relative_y {
		diffs = append(diffs, segno.GongMarshallField(stage, "Relative_y"))
	}
	if segno.Font_family != segnoOther.Font_family {
		diffs = append(diffs, segno.GongMarshallField(stage, "Font_family"))
	}
	if segno.Font_style != segnoOther.Font_style {
		diffs = append(diffs, segno.GongMarshallField(stage, "Font_style"))
	}
	if segno.Font_size != segnoOther.Font_size {
		diffs = append(diffs, segno.GongMarshallField(stage, "Font_size"))
	}
	if segno.Font_weight != segnoOther.Font_weight {
		diffs = append(diffs, segno.GongMarshallField(stage, "Font_weight"))
	}
	if segno.Color != segnoOther.Color {
		diffs = append(diffs, segno.GongMarshallField(stage, "Color"))
	}
	if segno.Halign != segnoOther.Halign {
		diffs = append(diffs, segno.GongMarshallField(stage, "Halign"))
	}
	if segno.Valign != segnoOther.Valign {
		diffs = append(diffs, segno.GongMarshallField(stage, "Valign"))
	}
	if segno.Id != segnoOther.Id {
		diffs = append(diffs, segno.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (slash *Slash) GongDiff(stage *Stage, slashOther *Slash) (diffs []string) {
	// insertion point for field diffs
	if slash.Name != slashOther.Name {
		diffs = append(diffs, slash.GongMarshallField(stage, "Name"))
	}
	if slash.Type != slashOther.Type {
		diffs = append(diffs, slash.GongMarshallField(stage, "Type"))
	}
	if slash.Use_dots != slashOther.Use_dots {
		diffs = append(diffs, slash.GongMarshallField(stage, "Use_dots"))
	}
	if slash.Use_stems != slashOther.Use_stems {
		diffs = append(diffs, slash.GongMarshallField(stage, "Use_stems"))
	}
	if slash.Slash_type != slashOther.Slash_type {
		diffs = append(diffs, slash.GongMarshallField(stage, "Slash_type"))
	}
	if slash.Slash_dot != slashOther.Slash_dot {
		diffs = append(diffs, slash.GongMarshallField(stage, "Slash_dot"))
	}
	if slash.Except_voice != slashOther.Except_voice {
		diffs = append(diffs, slash.GongMarshallField(stage, "Except_voice"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (slide *Slide) GongDiff(stage *Stage, slideOther *Slide) (diffs []string) {
	// insertion point for field diffs
	if slide.Name != slideOther.Name {
		diffs = append(diffs, slide.GongMarshallField(stage, "Name"))
	}
	if slide.Type != slideOther.Type {
		diffs = append(diffs, slide.GongMarshallField(stage, "Type"))
	}
	if slide.Number != slideOther.Number {
		diffs = append(diffs, slide.GongMarshallField(stage, "Number"))
	}
	if slide.Line_type != slideOther.Line_type {
		diffs = append(diffs, slide.GongMarshallField(stage, "Line_type"))
	}
	if slide.Dash_length != slideOther.Dash_length {
		diffs = append(diffs, slide.GongMarshallField(stage, "Dash_length"))
	}
	if slide.Space_length != slideOther.Space_length {
		diffs = append(diffs, slide.GongMarshallField(stage, "Space_length"))
	}
	if slide.Default_x != slideOther.Default_x {
		diffs = append(diffs, slide.GongMarshallField(stage, "Default_x"))
	}
	if slide.Default_y != slideOther.Default_y {
		diffs = append(diffs, slide.GongMarshallField(stage, "Default_y"))
	}
	if slide.Relative_x != slideOther.Relative_x {
		diffs = append(diffs, slide.GongMarshallField(stage, "Relative_x"))
	}
	if slide.Relative_y != slideOther.Relative_y {
		diffs = append(diffs, slide.GongMarshallField(stage, "Relative_y"))
	}
	if slide.Font_family != slideOther.Font_family {
		diffs = append(diffs, slide.GongMarshallField(stage, "Font_family"))
	}
	if slide.Font_style != slideOther.Font_style {
		diffs = append(diffs, slide.GongMarshallField(stage, "Font_style"))
	}
	if slide.Font_size != slideOther.Font_size {
		diffs = append(diffs, slide.GongMarshallField(stage, "Font_size"))
	}
	if slide.Font_weight != slideOther.Font_weight {
		diffs = append(diffs, slide.GongMarshallField(stage, "Font_weight"))
	}
	if slide.Color != slideOther.Color {
		diffs = append(diffs, slide.GongMarshallField(stage, "Color"))
	}
	if slide.Accelerate != slideOther.Accelerate {
		diffs = append(diffs, slide.GongMarshallField(stage, "Accelerate"))
	}
	if slide.Beats != slideOther.Beats {
		diffs = append(diffs, slide.GongMarshallField(stage, "Beats"))
	}
	if slide.First_beat != slideOther.First_beat {
		diffs = append(diffs, slide.GongMarshallField(stage, "First_beat"))
	}
	if slide.Last_beat != slideOther.Last_beat {
		diffs = append(diffs, slide.GongMarshallField(stage, "Last_beat"))
	}
	if slide.Id != slideOther.Id {
		diffs = append(diffs, slide.GongMarshallField(stage, "Id"))
	}
	if slide.EnclosedText != slideOther.EnclosedText {
		diffs = append(diffs, slide.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (slur *Slur) GongDiff(stage *Stage, slurOther *Slur) (diffs []string) {
	// insertion point for field diffs
	if slur.Name != slurOther.Name {
		diffs = append(diffs, slur.GongMarshallField(stage, "Name"))
	}
	if slur.Type != slurOther.Type {
		diffs = append(diffs, slur.GongMarshallField(stage, "Type"))
	}
	if slur.Number != slurOther.Number {
		diffs = append(diffs, slur.GongMarshallField(stage, "Number"))
	}
	if slur.Line_type != slurOther.Line_type {
		diffs = append(diffs, slur.GongMarshallField(stage, "Line_type"))
	}
	if slur.Dash_length != slurOther.Dash_length {
		diffs = append(diffs, slur.GongMarshallField(stage, "Dash_length"))
	}
	if slur.Space_length != slurOther.Space_length {
		diffs = append(diffs, slur.GongMarshallField(stage, "Space_length"))
	}
	if slur.Default_x != slurOther.Default_x {
		diffs = append(diffs, slur.GongMarshallField(stage, "Default_x"))
	}
	if slur.Default_y != slurOther.Default_y {
		diffs = append(diffs, slur.GongMarshallField(stage, "Default_y"))
	}
	if slur.Relative_x != slurOther.Relative_x {
		diffs = append(diffs, slur.GongMarshallField(stage, "Relative_x"))
	}
	if slur.Relative_y != slurOther.Relative_y {
		diffs = append(diffs, slur.GongMarshallField(stage, "Relative_y"))
	}
	if slur.Placement != slurOther.Placement {
		diffs = append(diffs, slur.GongMarshallField(stage, "Placement"))
	}
	if slur.Orientation != slurOther.Orientation {
		diffs = append(diffs, slur.GongMarshallField(stage, "Orientation"))
	}
	if slur.Bezier_x != slurOther.Bezier_x {
		diffs = append(diffs, slur.GongMarshallField(stage, "Bezier_x"))
	}
	if slur.Bezier_y != slurOther.Bezier_y {
		diffs = append(diffs, slur.GongMarshallField(stage, "Bezier_y"))
	}
	if slur.Bezier_x2 != slurOther.Bezier_x2 {
		diffs = append(diffs, slur.GongMarshallField(stage, "Bezier_x2"))
	}
	if slur.Bezier_y2 != slurOther.Bezier_y2 {
		diffs = append(diffs, slur.GongMarshallField(stage, "Bezier_y2"))
	}
	if slur.Bezier_offset != slurOther.Bezier_offset {
		diffs = append(diffs, slur.GongMarshallField(stage, "Bezier_offset"))
	}
	if slur.Bezier_offset2 != slurOther.Bezier_offset2 {
		diffs = append(diffs, slur.GongMarshallField(stage, "Bezier_offset2"))
	}
	if slur.Color != slurOther.Color {
		diffs = append(diffs, slur.GongMarshallField(stage, "Color"))
	}
	if slur.Id != slurOther.Id {
		diffs = append(diffs, slur.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (sound *Sound) GongDiff(stage *Stage, soundOther *Sound) (diffs []string) {
	// insertion point for field diffs
	if sound.Name != soundOther.Name {
		diffs = append(diffs, sound.GongMarshallField(stage, "Name"))
	}
	if sound.Tempo != soundOther.Tempo {
		diffs = append(diffs, sound.GongMarshallField(stage, "Tempo"))
	}
	if sound.Dynamics != soundOther.Dynamics {
		diffs = append(diffs, sound.GongMarshallField(stage, "Dynamics"))
	}
	if sound.Dacapo != soundOther.Dacapo {
		diffs = append(diffs, sound.GongMarshallField(stage, "Dacapo"))
	}
	if sound.Segno != soundOther.Segno {
		diffs = append(diffs, sound.GongMarshallField(stage, "Segno"))
	}
	if sound.Dalsegno != soundOther.Dalsegno {
		diffs = append(diffs, sound.GongMarshallField(stage, "Dalsegno"))
	}
	if sound.Coda != soundOther.Coda {
		diffs = append(diffs, sound.GongMarshallField(stage, "Coda"))
	}
	if sound.Tocoda != soundOther.Tocoda {
		diffs = append(diffs, sound.GongMarshallField(stage, "Tocoda"))
	}
	if sound.Divisions != soundOther.Divisions {
		diffs = append(diffs, sound.GongMarshallField(stage, "Divisions"))
	}
	if sound.Forward_repeat != soundOther.Forward_repeat {
		diffs = append(diffs, sound.GongMarshallField(stage, "Forward_repeat"))
	}
	if sound.Fine != soundOther.Fine {
		diffs = append(diffs, sound.GongMarshallField(stage, "Fine"))
	}
	if sound.Time_only != soundOther.Time_only {
		diffs = append(diffs, sound.GongMarshallField(stage, "Time_only"))
	}
	if sound.Pizzicato != soundOther.Pizzicato {
		diffs = append(diffs, sound.GongMarshallField(stage, "Pizzicato"))
	}
	if sound.Pan != soundOther.Pan {
		diffs = append(diffs, sound.GongMarshallField(stage, "Pan"))
	}
	if sound.Elevation != soundOther.Elevation {
		diffs = append(diffs, sound.GongMarshallField(stage, "Elevation"))
	}
	if sound.Damper_pedal != soundOther.Damper_pedal {
		diffs = append(diffs, sound.GongMarshallField(stage, "Damper_pedal"))
	}
	if sound.Soft_pedal != soundOther.Soft_pedal {
		diffs = append(diffs, sound.GongMarshallField(stage, "Soft_pedal"))
	}
	if sound.Sostenuto_pedal != soundOther.Sostenuto_pedal {
		diffs = append(diffs, sound.GongMarshallField(stage, "Sostenuto_pedal"))
	}
	if sound.Id != soundOther.Id {
		diffs = append(diffs, sound.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, sound, "Instrument_change", soundOther.Instrument_change, sound.Instrument_change); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, sound, "Midi_device", soundOther.Midi_device, sound.Midi_device); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, sound, "Midi_instrument", soundOther.Midi_instrument, sound.Midi_instrument); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, sound, "Play", soundOther.Play, sound.Play); ops != "" {
		diffs = append(diffs, ops)
	}
	if sound.Swing != soundOther.Swing {
		diffs = append(diffs, sound.GongMarshallField(stage, "Swing"))
	}
	if sound.Offset != soundOther.Offset {
		diffs = append(diffs, sound.GongMarshallField(stage, "Offset"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (staff_details *Staff_details) GongDiff(stage *Stage, staff_detailsOther *Staff_details) (diffs []string) {
	// insertion point for field diffs
	if staff_details.Name != staff_detailsOther.Name {
		diffs = append(diffs, staff_details.GongMarshallField(stage, "Name"))
	}
	if staff_details.Number != staff_detailsOther.Number {
		diffs = append(diffs, staff_details.GongMarshallField(stage, "Number"))
	}
	if staff_details.Show_frets != staff_detailsOther.Show_frets {
		diffs = append(diffs, staff_details.GongMarshallField(stage, "Show_frets"))
	}
	if staff_details.Print_object != staff_detailsOther.Print_object {
		diffs = append(diffs, staff_details.GongMarshallField(stage, "Print_object"))
	}
	if staff_details.Print_spacing != staff_detailsOther.Print_spacing {
		diffs = append(diffs, staff_details.GongMarshallField(stage, "Print_spacing"))
	}
	if staff_details.Staff_type != staff_detailsOther.Staff_type {
		diffs = append(diffs, staff_details.GongMarshallField(stage, "Staff_type"))
	}
	if staff_details.Staff_lines != staff_detailsOther.Staff_lines {
		diffs = append(diffs, staff_details.GongMarshallField(stage, "Staff_lines"))
	}
	if ops := __gong__diffSliceOfPointers(stage, staff_details, "Line_detail", staff_detailsOther.Line_detail, staff_details.Line_detail); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, staff_details, "Staff_tuning", staff_detailsOther.Staff_tuning, staff_details.Staff_tuning); ops != "" {
		diffs = append(diffs, ops)
	}
	if staff_details.Capo != staff_detailsOther.Capo {
		diffs = append(diffs, staff_details.GongMarshallField(stage, "Capo"))
	}
	if staff_details.Staff_size != staff_detailsOther.Staff_size {
		diffs = append(diffs, staff_details.GongMarshallField(stage, "Staff_size"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (staff_divide *Staff_divide) GongDiff(stage *Stage, staff_divideOther *Staff_divide) (diffs []string) {
	// insertion point for field diffs
	if staff_divide.Name != staff_divideOther.Name {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Name"))
	}
	if staff_divide.Type != staff_divideOther.Type {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Type"))
	}
	if staff_divide.Default_x != staff_divideOther.Default_x {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Default_x"))
	}
	if staff_divide.Default_y != staff_divideOther.Default_y {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Default_y"))
	}
	if staff_divide.Relative_x != staff_divideOther.Relative_x {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Relative_x"))
	}
	if staff_divide.Relative_y != staff_divideOther.Relative_y {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Relative_y"))
	}
	if staff_divide.Font_family != staff_divideOther.Font_family {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Font_family"))
	}
	if staff_divide.Font_style != staff_divideOther.Font_style {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Font_style"))
	}
	if staff_divide.Font_size != staff_divideOther.Font_size {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Font_size"))
	}
	if staff_divide.Font_weight != staff_divideOther.Font_weight {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Font_weight"))
	}
	if staff_divide.Color != staff_divideOther.Color {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Color"))
	}
	if staff_divide.Halign != staff_divideOther.Halign {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Halign"))
	}
	if staff_divide.Valign != staff_divideOther.Valign {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Valign"))
	}
	if staff_divide.Id != staff_divideOther.Id {
		diffs = append(diffs, staff_divide.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (staff_layout *Staff_layout) GongDiff(stage *Stage, staff_layoutOther *Staff_layout) (diffs []string) {
	// insertion point for field diffs
	if staff_layout.Name != staff_layoutOther.Name {
		diffs = append(diffs, staff_layout.GongMarshallField(stage, "Name"))
	}
	if staff_layout.Number != staff_layoutOther.Number {
		diffs = append(diffs, staff_layout.GongMarshallField(stage, "Number"))
	}
	if staff_layout.Staff_distance != staff_layoutOther.Staff_distance {
		diffs = append(diffs, staff_layout.GongMarshallField(stage, "Staff_distance"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (staff_size *Staff_size) GongDiff(stage *Stage, staff_sizeOther *Staff_size) (diffs []string) {
	// insertion point for field diffs
	if staff_size.Name != staff_sizeOther.Name {
		diffs = append(diffs, staff_size.GongMarshallField(stage, "Name"))
	}
	if staff_size.Scaling != staff_sizeOther.Scaling {
		diffs = append(diffs, staff_size.GongMarshallField(stage, "Scaling"))
	}
	if staff_size.EnclosedText != staff_sizeOther.EnclosedText {
		diffs = append(diffs, staff_size.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (staff_tuning *Staff_tuning) GongDiff(stage *Stage, staff_tuningOther *Staff_tuning) (diffs []string) {
	// insertion point for field diffs
	if staff_tuning.Name != staff_tuningOther.Name {
		diffs = append(diffs, staff_tuning.GongMarshallField(stage, "Name"))
	}
	if staff_tuning.Line != staff_tuningOther.Line {
		diffs = append(diffs, staff_tuning.GongMarshallField(stage, "Line"))
	}
	if staff_tuning.Tuning_step != staff_tuningOther.Tuning_step {
		diffs = append(diffs, staff_tuning.GongMarshallField(stage, "Tuning_step"))
	}
	if staff_tuning.Tuning_alter != staff_tuningOther.Tuning_alter {
		diffs = append(diffs, staff_tuning.GongMarshallField(stage, "Tuning_alter"))
	}
	if staff_tuning.Tuning_octave != staff_tuningOther.Tuning_octave {
		diffs = append(diffs, staff_tuning.GongMarshallField(stage, "Tuning_octave"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stem *Stem) GongDiff(stage *Stage, stemOther *Stem) (diffs []string) {
	// insertion point for field diffs
	if stem.Name != stemOther.Name {
		diffs = append(diffs, stem.GongMarshallField(stage, "Name"))
	}
	if stem.Default_x != stemOther.Default_x {
		diffs = append(diffs, stem.GongMarshallField(stage, "Default_x"))
	}
	if stem.Default_y != stemOther.Default_y {
		diffs = append(diffs, stem.GongMarshallField(stage, "Default_y"))
	}
	if stem.Relative_x != stemOther.Relative_x {
		diffs = append(diffs, stem.GongMarshallField(stage, "Relative_x"))
	}
	if stem.Relative_y != stemOther.Relative_y {
		diffs = append(diffs, stem.GongMarshallField(stage, "Relative_y"))
	}
	if stem.Color != stemOther.Color {
		diffs = append(diffs, stem.GongMarshallField(stage, "Color"))
	}
	if stem.EnclosedText != stemOther.EnclosedText {
		diffs = append(diffs, stem.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stick *Stick) GongDiff(stage *Stage, stickOther *Stick) (diffs []string) {
	// insertion point for field diffs
	if stick.Name != stickOther.Name {
		diffs = append(diffs, stick.GongMarshallField(stage, "Name"))
	}
	if stick.Tip != stickOther.Tip {
		diffs = append(diffs, stick.GongMarshallField(stage, "Tip"))
	}
	if stick.Parentheses != stickOther.Parentheses {
		diffs = append(diffs, stick.GongMarshallField(stage, "Parentheses"))
	}
	if stick.Dashed_circle != stickOther.Dashed_circle {
		diffs = append(diffs, stick.GongMarshallField(stage, "Dashed_circle"))
	}
	if stick.Stick_type != stickOther.Stick_type {
		diffs = append(diffs, stick.GongMarshallField(stage, "Stick_type"))
	}
	if stick.Stick_material != stickOther.Stick_material {
		diffs = append(diffs, stick.GongMarshallField(stage, "Stick_material"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (string_mute *String_mute) GongDiff(stage *Stage, string_muteOther *String_mute) (diffs []string) {
	// insertion point for field diffs
	if string_mute.Name != string_muteOther.Name {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Name"))
	}
	if string_mute.Type != string_muteOther.Type {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Type"))
	}
	if string_mute.Default_x != string_muteOther.Default_x {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Default_x"))
	}
	if string_mute.Default_y != string_muteOther.Default_y {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Default_y"))
	}
	if string_mute.Relative_x != string_muteOther.Relative_x {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Relative_x"))
	}
	if string_mute.Relative_y != string_muteOther.Relative_y {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Relative_y"))
	}
	if string_mute.Font_family != string_muteOther.Font_family {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Font_family"))
	}
	if string_mute.Font_style != string_muteOther.Font_style {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Font_style"))
	}
	if string_mute.Font_size != string_muteOther.Font_size {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Font_size"))
	}
	if string_mute.Font_weight != string_muteOther.Font_weight {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Font_weight"))
	}
	if string_mute.Color != string_muteOther.Color {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Color"))
	}
	if string_mute.Halign != string_muteOther.Halign {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Halign"))
	}
	if string_mute.Valign != string_muteOther.Valign {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Valign"))
	}
	if string_mute.Id != string_muteOther.Id {
		diffs = append(diffs, string_mute.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (string_type *String_type) GongDiff(stage *Stage, string_typeOther *String_type) (diffs []string) {
	// insertion point for field diffs
	if string_type.Name != string_typeOther.Name {
		diffs = append(diffs, string_type.GongMarshallField(stage, "Name"))
	}
	if string_type.Default_x != string_typeOther.Default_x {
		diffs = append(diffs, string_type.GongMarshallField(stage, "Default_x"))
	}
	if string_type.Default_y != string_typeOther.Default_y {
		diffs = append(diffs, string_type.GongMarshallField(stage, "Default_y"))
	}
	if string_type.Relative_x != string_typeOther.Relative_x {
		diffs = append(diffs, string_type.GongMarshallField(stage, "Relative_x"))
	}
	if string_type.Relative_y != string_typeOther.Relative_y {
		diffs = append(diffs, string_type.GongMarshallField(stage, "Relative_y"))
	}
	if string_type.Font_family != string_typeOther.Font_family {
		diffs = append(diffs, string_type.GongMarshallField(stage, "Font_family"))
	}
	if string_type.Font_style != string_typeOther.Font_style {
		diffs = append(diffs, string_type.GongMarshallField(stage, "Font_style"))
	}
	if string_type.Font_size != string_typeOther.Font_size {
		diffs = append(diffs, string_type.GongMarshallField(stage, "Font_size"))
	}
	if string_type.Font_weight != string_typeOther.Font_weight {
		diffs = append(diffs, string_type.GongMarshallField(stage, "Font_weight"))
	}
	if string_type.Color != string_typeOther.Color {
		diffs = append(diffs, string_type.GongMarshallField(stage, "Color"))
	}
	if string_type.Placement != string_typeOther.Placement {
		diffs = append(diffs, string_type.GongMarshallField(stage, "Placement"))
	}
	if string_type.EnclosedText != string_typeOther.EnclosedText {
		diffs = append(diffs, string_type.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (strong_accent *Strong_accent) GongDiff(stage *Stage, strong_accentOther *Strong_accent) (diffs []string) {
	// insertion point for field diffs
	if strong_accent.Name != strong_accentOther.Name {
		diffs = append(diffs, strong_accent.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (style_text *Style_text) GongDiff(stage *Stage, style_textOther *Style_text) (diffs []string) {
	// insertion point for field diffs
	if style_text.Name != style_textOther.Name {
		diffs = append(diffs, style_text.GongMarshallField(stage, "Name"))
	}
	if style_text.Default_x != style_textOther.Default_x {
		diffs = append(diffs, style_text.GongMarshallField(stage, "Default_x"))
	}
	if style_text.Default_y != style_textOther.Default_y {
		diffs = append(diffs, style_text.GongMarshallField(stage, "Default_y"))
	}
	if style_text.Relative_x != style_textOther.Relative_x {
		diffs = append(diffs, style_text.GongMarshallField(stage, "Relative_x"))
	}
	if style_text.Relative_y != style_textOther.Relative_y {
		diffs = append(diffs, style_text.GongMarshallField(stage, "Relative_y"))
	}
	if style_text.Font_family != style_textOther.Font_family {
		diffs = append(diffs, style_text.GongMarshallField(stage, "Font_family"))
	}
	if style_text.Font_style != style_textOther.Font_style {
		diffs = append(diffs, style_text.GongMarshallField(stage, "Font_style"))
	}
	if style_text.Font_size != style_textOther.Font_size {
		diffs = append(diffs, style_text.GongMarshallField(stage, "Font_size"))
	}
	if style_text.Font_weight != style_textOther.Font_weight {
		diffs = append(diffs, style_text.GongMarshallField(stage, "Font_weight"))
	}
	if style_text.Color != style_textOther.Color {
		diffs = append(diffs, style_text.GongMarshallField(stage, "Color"))
	}
	if style_text.EnclosedText != style_textOther.EnclosedText {
		diffs = append(diffs, style_text.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (supports *Supports) GongDiff(stage *Stage, supportsOther *Supports) (diffs []string) {
	// insertion point for field diffs
	if supports.Name != supportsOther.Name {
		diffs = append(diffs, supports.GongMarshallField(stage, "Name"))
	}
	if supports.Type != supportsOther.Type {
		diffs = append(diffs, supports.GongMarshallField(stage, "Type"))
	}
	if supports.Element != supportsOther.Element {
		diffs = append(diffs, supports.GongMarshallField(stage, "Element"))
	}
	if supports.Attribute != supportsOther.Attribute {
		diffs = append(diffs, supports.GongMarshallField(stage, "Attribute"))
	}
	if supports.Value != supportsOther.Value {
		diffs = append(diffs, supports.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (swing *Swing) GongDiff(stage *Stage, swingOther *Swing) (diffs []string) {
	// insertion point for field diffs
	if swing.Name != swingOther.Name {
		diffs = append(diffs, swing.GongMarshallField(stage, "Name"))
	}
	if swing.Straight != swingOther.Straight {
		diffs = append(diffs, swing.GongMarshallField(stage, "Straight"))
	}
	if swing.First != swingOther.First {
		diffs = append(diffs, swing.GongMarshallField(stage, "First"))
	}
	if swing.Second != swingOther.Second {
		diffs = append(diffs, swing.GongMarshallField(stage, "Second"))
	}
	if swing.Swing_type != swingOther.Swing_type {
		diffs = append(diffs, swing.GongMarshallField(stage, "Swing_type"))
	}
	if swing.Swing_style != swingOther.Swing_style {
		diffs = append(diffs, swing.GongMarshallField(stage, "Swing_style"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (sync *Sync) GongDiff(stage *Stage, syncOther *Sync) (diffs []string) {
	// insertion point for field diffs
	if sync.Name != syncOther.Name {
		diffs = append(diffs, sync.GongMarshallField(stage, "Name"))
	}
	if sync.Type != syncOther.Type {
		diffs = append(diffs, sync.GongMarshallField(stage, "Type"))
	}
	if sync.Latency != syncOther.Latency {
		diffs = append(diffs, sync.GongMarshallField(stage, "Latency"))
	}
	if sync.Player != syncOther.Player {
		diffs = append(diffs, sync.GongMarshallField(stage, "Player"))
	}
	if sync.Time_only != syncOther.Time_only {
		diffs = append(diffs, sync.GongMarshallField(stage, "Time_only"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (system_dividers *System_dividers) GongDiff(stage *Stage, system_dividersOther *System_dividers) (diffs []string) {
	// insertion point for field diffs
	if system_dividers.Name != system_dividersOther.Name {
		diffs = append(diffs, system_dividers.GongMarshallField(stage, "Name"))
	}
	if system_dividers.Left_divider != system_dividersOther.Left_divider {
		diffs = append(diffs, system_dividers.GongMarshallField(stage, "Left_divider"))
	}
	if system_dividers.Right_divider != system_dividersOther.Right_divider {
		diffs = append(diffs, system_dividers.GongMarshallField(stage, "Right_divider"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (system_layout *System_layout) GongDiff(stage *Stage, system_layoutOther *System_layout) (diffs []string) {
	// insertion point for field diffs
	if system_layout.Name != system_layoutOther.Name {
		diffs = append(diffs, system_layout.GongMarshallField(stage, "Name"))
	}
	if system_layout.System_margins != system_layoutOther.System_margins {
		diffs = append(diffs, system_layout.GongMarshallField(stage, "System_margins"))
	}
	if system_layout.System_distance != system_layoutOther.System_distance {
		diffs = append(diffs, system_layout.GongMarshallField(stage, "System_distance"))
	}
	if system_layout.Top_system_distance != system_layoutOther.Top_system_distance {
		diffs = append(diffs, system_layout.GongMarshallField(stage, "Top_system_distance"))
	}
	if system_layout.System_dividers != system_layoutOther.System_dividers {
		diffs = append(diffs, system_layout.GongMarshallField(stage, "System_dividers"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (system_margins *System_margins) GongDiff(stage *Stage, system_marginsOther *System_margins) (diffs []string) {
	// insertion point for field diffs
	if system_margins.Name != system_marginsOther.Name {
		diffs = append(diffs, system_margins.GongMarshallField(stage, "Name"))
	}
	if system_margins.Left_margin != system_marginsOther.Left_margin {
		diffs = append(diffs, system_margins.GongMarshallField(stage, "Left_margin"))
	}
	if system_margins.Right_margin != system_marginsOther.Right_margin {
		diffs = append(diffs, system_margins.GongMarshallField(stage, "Right_margin"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tap *Tap) GongDiff(stage *Stage, tapOther *Tap) (diffs []string) {
	// insertion point for field diffs
	if tap.Name != tapOther.Name {
		diffs = append(diffs, tap.GongMarshallField(stage, "Name"))
	}
	if tap.Hand != tapOther.Hand {
		diffs = append(diffs, tap.GongMarshallField(stage, "Hand"))
	}
	if tap.Default_x != tapOther.Default_x {
		diffs = append(diffs, tap.GongMarshallField(stage, "Default_x"))
	}
	if tap.Default_y != tapOther.Default_y {
		diffs = append(diffs, tap.GongMarshallField(stage, "Default_y"))
	}
	if tap.Relative_x != tapOther.Relative_x {
		diffs = append(diffs, tap.GongMarshallField(stage, "Relative_x"))
	}
	if tap.Relative_y != tapOther.Relative_y {
		diffs = append(diffs, tap.GongMarshallField(stage, "Relative_y"))
	}
	if tap.Font_family != tapOther.Font_family {
		diffs = append(diffs, tap.GongMarshallField(stage, "Font_family"))
	}
	if tap.Font_style != tapOther.Font_style {
		diffs = append(diffs, tap.GongMarshallField(stage, "Font_style"))
	}
	if tap.Font_size != tapOther.Font_size {
		diffs = append(diffs, tap.GongMarshallField(stage, "Font_size"))
	}
	if tap.Font_weight != tapOther.Font_weight {
		diffs = append(diffs, tap.GongMarshallField(stage, "Font_weight"))
	}
	if tap.Color != tapOther.Color {
		diffs = append(diffs, tap.GongMarshallField(stage, "Color"))
	}
	if tap.Placement != tapOther.Placement {
		diffs = append(diffs, tap.GongMarshallField(stage, "Placement"))
	}
	if tap.EnclosedText != tapOther.EnclosedText {
		diffs = append(diffs, tap.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (technical *Technical) GongDiff(stage *Stage, technicalOther *Technical) (diffs []string) {
	// insertion point for field diffs
	if technical.Name != technicalOther.Name {
		diffs = append(diffs, technical.GongMarshallField(stage, "Name"))
	}
	if technical.Id != technicalOther.Id {
		diffs = append(diffs, technical.GongMarshallField(stage, "Id"))
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Up_bow", technicalOther.Up_bow, technical.Up_bow); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Down_bow", technicalOther.Down_bow, technical.Down_bow); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Harmonic", technicalOther.Harmonic, technical.Harmonic); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Open_string", technicalOther.Open_string, technical.Open_string); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Thumb_position", technicalOther.Thumb_position, technical.Thumb_position); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Fingering", technicalOther.Fingering, technical.Fingering); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Pluck", technicalOther.Pluck, technical.Pluck); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Double_tongue", technicalOther.Double_tongue, technical.Double_tongue); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Triple_tongue", technicalOther.Triple_tongue, technical.Triple_tongue); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Stopped", technicalOther.Stopped, technical.Stopped); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Snap_pizzicato", technicalOther.Snap_pizzicato, technical.Snap_pizzicato); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Fret", technicalOther.Fret, technical.Fret); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "String", technicalOther.String, technical.String); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Hammer_on", technicalOther.Hammer_on, technical.Hammer_on); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Pull_off", technicalOther.Pull_off, technical.Pull_off); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Bend", technicalOther.Bend, technical.Bend); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Tap", technicalOther.Tap, technical.Tap); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Heel", technicalOther.Heel, technical.Heel); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Toe", technicalOther.Toe, technical.Toe); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Fingernails", technicalOther.Fingernails, technical.Fingernails); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Hole", technicalOther.Hole, technical.Hole); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Arrow", technicalOther.Arrow, technical.Arrow); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Handbell", technicalOther.Handbell, technical.Handbell); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Brass_bend", technicalOther.Brass_bend, technical.Brass_bend); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Flip", technicalOther.Flip, technical.Flip); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Smear", technicalOther.Smear, technical.Smear); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Open", technicalOther.Open, technical.Open); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Half_muted", technicalOther.Half_muted, technical.Half_muted); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Harmon_mute", technicalOther.Harmon_mute, technical.Harmon_mute); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Golpe", technicalOther.Golpe, technical.Golpe); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, technical, "Other_technical", technicalOther.Other_technical, technical.Other_technical); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (text_element_data *Text_element_data) GongDiff(stage *Stage, text_element_dataOther *Text_element_data) (diffs []string) {
	// insertion point for field diffs
	if text_element_data.Name != text_element_dataOther.Name {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "Name"))
	}
	if text_element_data.Lang != text_element_dataOther.Lang {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "Lang"))
	}
	if text_element_data.Font_family != text_element_dataOther.Font_family {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "Font_family"))
	}
	if text_element_data.Font_style != text_element_dataOther.Font_style {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "Font_style"))
	}
	if text_element_data.Font_size != text_element_dataOther.Font_size {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "Font_size"))
	}
	if text_element_data.Font_weight != text_element_dataOther.Font_weight {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "Font_weight"))
	}
	if text_element_data.Color != text_element_dataOther.Color {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "Color"))
	}
	if text_element_data.Underline != text_element_dataOther.Underline {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "Underline"))
	}
	if text_element_data.Overline != text_element_dataOther.Overline {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "Overline"))
	}
	if text_element_data.Line_through != text_element_dataOther.Line_through {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "Line_through"))
	}
	if text_element_data.Rotation != text_element_dataOther.Rotation {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "Rotation"))
	}
	if text_element_data.Letter_spacing != text_element_dataOther.Letter_spacing {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "Letter_spacing"))
	}
	if text_element_data.Dir != text_element_dataOther.Dir {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "Dir"))
	}
	if text_element_data.EnclosedText != text_element_dataOther.EnclosedText {
		diffs = append(diffs, text_element_data.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tie *Tie) GongDiff(stage *Stage, tieOther *Tie) (diffs []string) {
	// insertion point for field diffs
	if tie.Name != tieOther.Name {
		diffs = append(diffs, tie.GongMarshallField(stage, "Name"))
	}
	if tie.Type != tieOther.Type {
		diffs = append(diffs, tie.GongMarshallField(stage, "Type"))
	}
	if tie.Time_only != tieOther.Time_only {
		diffs = append(diffs, tie.GongMarshallField(stage, "Time_only"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tied *Tied) GongDiff(stage *Stage, tiedOther *Tied) (diffs []string) {
	// insertion point for field diffs
	if tied.Name != tiedOther.Name {
		diffs = append(diffs, tied.GongMarshallField(stage, "Name"))
	}
	if tied.Type != tiedOther.Type {
		diffs = append(diffs, tied.GongMarshallField(stage, "Type"))
	}
	if tied.Number != tiedOther.Number {
		diffs = append(diffs, tied.GongMarshallField(stage, "Number"))
	}
	if tied.Line_type != tiedOther.Line_type {
		diffs = append(diffs, tied.GongMarshallField(stage, "Line_type"))
	}
	if tied.Dash_length != tiedOther.Dash_length {
		diffs = append(diffs, tied.GongMarshallField(stage, "Dash_length"))
	}
	if tied.Space_length != tiedOther.Space_length {
		diffs = append(diffs, tied.GongMarshallField(stage, "Space_length"))
	}
	if tied.Default_x != tiedOther.Default_x {
		diffs = append(diffs, tied.GongMarshallField(stage, "Default_x"))
	}
	if tied.Default_y != tiedOther.Default_y {
		diffs = append(diffs, tied.GongMarshallField(stage, "Default_y"))
	}
	if tied.Relative_x != tiedOther.Relative_x {
		diffs = append(diffs, tied.GongMarshallField(stage, "Relative_x"))
	}
	if tied.Relative_y != tiedOther.Relative_y {
		diffs = append(diffs, tied.GongMarshallField(stage, "Relative_y"))
	}
	if tied.Placement != tiedOther.Placement {
		diffs = append(diffs, tied.GongMarshallField(stage, "Placement"))
	}
	if tied.Orientation != tiedOther.Orientation {
		diffs = append(diffs, tied.GongMarshallField(stage, "Orientation"))
	}
	if tied.Bezier_x != tiedOther.Bezier_x {
		diffs = append(diffs, tied.GongMarshallField(stage, "Bezier_x"))
	}
	if tied.Bezier_y != tiedOther.Bezier_y {
		diffs = append(diffs, tied.GongMarshallField(stage, "Bezier_y"))
	}
	if tied.Bezier_x2 != tiedOther.Bezier_x2 {
		diffs = append(diffs, tied.GongMarshallField(stage, "Bezier_x2"))
	}
	if tied.Bezier_y2 != tiedOther.Bezier_y2 {
		diffs = append(diffs, tied.GongMarshallField(stage, "Bezier_y2"))
	}
	if tied.Bezier_offset != tiedOther.Bezier_offset {
		diffs = append(diffs, tied.GongMarshallField(stage, "Bezier_offset"))
	}
	if tied.Bezier_offset2 != tiedOther.Bezier_offset2 {
		diffs = append(diffs, tied.GongMarshallField(stage, "Bezier_offset2"))
	}
	if tied.Color != tiedOther.Color {
		diffs = append(diffs, tied.GongMarshallField(stage, "Color"))
	}
	if tied.Id != tiedOther.Id {
		diffs = append(diffs, tied.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (time *Time) GongDiff(stage *Stage, timeOther *Time) (diffs []string) {
	// insertion point for field diffs
	if time.Name != timeOther.Name {
		diffs = append(diffs, time.GongMarshallField(stage, "Name"))
	}
	if time.Number != timeOther.Number {
		diffs = append(diffs, time.GongMarshallField(stage, "Number"))
	}
	if time.Symbol != timeOther.Symbol {
		diffs = append(diffs, time.GongMarshallField(stage, "Symbol"))
	}
	if time.Separator != timeOther.Separator {
		diffs = append(diffs, time.GongMarshallField(stage, "Separator"))
	}
	if time.Default_x != timeOther.Default_x {
		diffs = append(diffs, time.GongMarshallField(stage, "Default_x"))
	}
	if time.Default_y != timeOther.Default_y {
		diffs = append(diffs, time.GongMarshallField(stage, "Default_y"))
	}
	if time.Relative_x != timeOther.Relative_x {
		diffs = append(diffs, time.GongMarshallField(stage, "Relative_x"))
	}
	if time.Relative_y != timeOther.Relative_y {
		diffs = append(diffs, time.GongMarshallField(stage, "Relative_y"))
	}
	if time.Font_family != timeOther.Font_family {
		diffs = append(diffs, time.GongMarshallField(stage, "Font_family"))
	}
	if time.Font_style != timeOther.Font_style {
		diffs = append(diffs, time.GongMarshallField(stage, "Font_style"))
	}
	if time.Font_size != timeOther.Font_size {
		diffs = append(diffs, time.GongMarshallField(stage, "Font_size"))
	}
	if time.Font_weight != timeOther.Font_weight {
		diffs = append(diffs, time.GongMarshallField(stage, "Font_weight"))
	}
	if time.Color != timeOther.Color {
		diffs = append(diffs, time.GongMarshallField(stage, "Color"))
	}
	if time.Halign != timeOther.Halign {
		diffs = append(diffs, time.GongMarshallField(stage, "Halign"))
	}
	if time.Valign != timeOther.Valign {
		diffs = append(diffs, time.GongMarshallField(stage, "Valign"))
	}
	if time.Print_object != timeOther.Print_object {
		diffs = append(diffs, time.GongMarshallField(stage, "Print_object"))
	}
	if time.Id != timeOther.Id {
		diffs = append(diffs, time.GongMarshallField(stage, "Id"))
	}
	if time.Beats != timeOther.Beats {
		diffs = append(diffs, time.GongMarshallField(stage, "Beats"))
	}
	if time.Beat_type != timeOther.Beat_type {
		diffs = append(diffs, time.GongMarshallField(stage, "Beat_type"))
	}
	if time.Interchangeable != timeOther.Interchangeable {
		diffs = append(diffs, time.GongMarshallField(stage, "Interchangeable"))
	}
	if time.Senza_misura != timeOther.Senza_misura {
		diffs = append(diffs, time.GongMarshallField(stage, "Senza_misura"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (time_modification *Time_modification) GongDiff(stage *Stage, time_modificationOther *Time_modification) (diffs []string) {
	// insertion point for field diffs
	if time_modification.Name != time_modificationOther.Name {
		diffs = append(diffs, time_modification.GongMarshallField(stage, "Name"))
	}
	if time_modification.Actual_notes != time_modificationOther.Actual_notes {
		diffs = append(diffs, time_modification.GongMarshallField(stage, "Actual_notes"))
	}
	if time_modification.Normal_notes != time_modificationOther.Normal_notes {
		diffs = append(diffs, time_modification.GongMarshallField(stage, "Normal_notes"))
	}
	if time_modification.Normal_type != time_modificationOther.Normal_type {
		diffs = append(diffs, time_modification.GongMarshallField(stage, "Normal_type"))
	}
	if time_modification.Normal_dot != time_modificationOther.Normal_dot {
		diffs = append(diffs, time_modification.GongMarshallField(stage, "Normal_dot"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timpani *Timpani) GongDiff(stage *Stage, timpaniOther *Timpani) (diffs []string) {
	// insertion point for field diffs
	if timpani.Name != timpaniOther.Name {
		diffs = append(diffs, timpani.GongMarshallField(stage, "Name"))
	}
	if timpani.Smufl != timpaniOther.Smufl {
		diffs = append(diffs, timpani.GongMarshallField(stage, "Smufl"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (transpose *Transpose) GongDiff(stage *Stage, transposeOther *Transpose) (diffs []string) {
	// insertion point for field diffs
	if transpose.Name != transposeOther.Name {
		diffs = append(diffs, transpose.GongMarshallField(stage, "Name"))
	}
	if transpose.Number != transposeOther.Number {
		diffs = append(diffs, transpose.GongMarshallField(stage, "Number"))
	}
	if transpose.Id != transposeOther.Id {
		diffs = append(diffs, transpose.GongMarshallField(stage, "Id"))
	}
	if transpose.Diatonic != transposeOther.Diatonic {
		diffs = append(diffs, transpose.GongMarshallField(stage, "Diatonic"))
	}
	if transpose.Chromatic != transposeOther.Chromatic {
		diffs = append(diffs, transpose.GongMarshallField(stage, "Chromatic"))
	}
	if transpose.Octave_change != transposeOther.Octave_change {
		diffs = append(diffs, transpose.GongMarshallField(stage, "Octave_change"))
	}
	if transpose.Double != transposeOther.Double {
		diffs = append(diffs, transpose.GongMarshallField(stage, "Double"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tremolo *Tremolo) GongDiff(stage *Stage, tremoloOther *Tremolo) (diffs []string) {
	// insertion point for field diffs
	if tremolo.Name != tremoloOther.Name {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Name"))
	}
	if tremolo.Type != tremoloOther.Type {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Type"))
	}
	if tremolo.Default_x != tremoloOther.Default_x {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Default_x"))
	}
	if tremolo.Default_y != tremoloOther.Default_y {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Default_y"))
	}
	if tremolo.Relative_x != tremoloOther.Relative_x {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Relative_x"))
	}
	if tremolo.Relative_y != tremoloOther.Relative_y {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Relative_y"))
	}
	if tremolo.Font_family != tremoloOther.Font_family {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Font_family"))
	}
	if tremolo.Font_style != tremoloOther.Font_style {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Font_style"))
	}
	if tremolo.Font_size != tremoloOther.Font_size {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Font_size"))
	}
	if tremolo.Font_weight != tremoloOther.Font_weight {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Font_weight"))
	}
	if tremolo.Color != tremoloOther.Color {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Color"))
	}
	if tremolo.Placement != tremoloOther.Placement {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Placement"))
	}
	if tremolo.Smufl != tremoloOther.Smufl {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "Smufl"))
	}
	if tremolo.EnclosedText != tremoloOther.EnclosedText {
		diffs = append(diffs, tremolo.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tuplet *Tuplet) GongDiff(stage *Stage, tupletOther *Tuplet) (diffs []string) {
	// insertion point for field diffs
	if tuplet.Name != tupletOther.Name {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Name"))
	}
	if tuplet.Type != tupletOther.Type {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Type"))
	}
	if tuplet.Number != tupletOther.Number {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Number"))
	}
	if tuplet.Bracket != tupletOther.Bracket {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Bracket"))
	}
	if tuplet.Show_number != tupletOther.Show_number {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Show_number"))
	}
	if tuplet.Show_type != tupletOther.Show_type {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Show_type"))
	}
	if tuplet.Line_shape != tupletOther.Line_shape {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Line_shape"))
	}
	if tuplet.Default_x != tupletOther.Default_x {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Default_x"))
	}
	if tuplet.Default_y != tupletOther.Default_y {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Default_y"))
	}
	if tuplet.Relative_x != tupletOther.Relative_x {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Relative_x"))
	}
	if tuplet.Relative_y != tupletOther.Relative_y {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Relative_y"))
	}
	if tuplet.Placement != tupletOther.Placement {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Placement"))
	}
	if tuplet.Id != tupletOther.Id {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Id"))
	}
	if tuplet.Tuplet_actual != tupletOther.Tuplet_actual {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Tuplet_actual"))
	}
	if tuplet.Tuplet_normal != tupletOther.Tuplet_normal {
		diffs = append(diffs, tuplet.GongMarshallField(stage, "Tuplet_normal"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tuplet_dot *Tuplet_dot) GongDiff(stage *Stage, tuplet_dotOther *Tuplet_dot) (diffs []string) {
	// insertion point for field diffs
	if tuplet_dot.Name != tuplet_dotOther.Name {
		diffs = append(diffs, tuplet_dot.GongMarshallField(stage, "Name"))
	}
	if tuplet_dot.Font_family != tuplet_dotOther.Font_family {
		diffs = append(diffs, tuplet_dot.GongMarshallField(stage, "Font_family"))
	}
	if tuplet_dot.Font_style != tuplet_dotOther.Font_style {
		diffs = append(diffs, tuplet_dot.GongMarshallField(stage, "Font_style"))
	}
	if tuplet_dot.Font_size != tuplet_dotOther.Font_size {
		diffs = append(diffs, tuplet_dot.GongMarshallField(stage, "Font_size"))
	}
	if tuplet_dot.Font_weight != tuplet_dotOther.Font_weight {
		diffs = append(diffs, tuplet_dot.GongMarshallField(stage, "Font_weight"))
	}
	if tuplet_dot.Color != tuplet_dotOther.Color {
		diffs = append(diffs, tuplet_dot.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tuplet_number *Tuplet_number) GongDiff(stage *Stage, tuplet_numberOther *Tuplet_number) (diffs []string) {
	// insertion point for field diffs
	if tuplet_number.Name != tuplet_numberOther.Name {
		diffs = append(diffs, tuplet_number.GongMarshallField(stage, "Name"))
	}
	if tuplet_number.Font_family != tuplet_numberOther.Font_family {
		diffs = append(diffs, tuplet_number.GongMarshallField(stage, "Font_family"))
	}
	if tuplet_number.Font_style != tuplet_numberOther.Font_style {
		diffs = append(diffs, tuplet_number.GongMarshallField(stage, "Font_style"))
	}
	if tuplet_number.Font_size != tuplet_numberOther.Font_size {
		diffs = append(diffs, tuplet_number.GongMarshallField(stage, "Font_size"))
	}
	if tuplet_number.Font_weight != tuplet_numberOther.Font_weight {
		diffs = append(diffs, tuplet_number.GongMarshallField(stage, "Font_weight"))
	}
	if tuplet_number.Color != tuplet_numberOther.Color {
		diffs = append(diffs, tuplet_number.GongMarshallField(stage, "Color"))
	}
	if tuplet_number.EnclosedText != tuplet_numberOther.EnclosedText {
		diffs = append(diffs, tuplet_number.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tuplet_portion *Tuplet_portion) GongDiff(stage *Stage, tuplet_portionOther *Tuplet_portion) (diffs []string) {
	// insertion point for field diffs
	if tuplet_portion.Name != tuplet_portionOther.Name {
		diffs = append(diffs, tuplet_portion.GongMarshallField(stage, "Name"))
	}
	if tuplet_portion.Tuplet_number != tuplet_portionOther.Tuplet_number {
		diffs = append(diffs, tuplet_portion.GongMarshallField(stage, "Tuplet_number"))
	}
	if tuplet_portion.Tuplet_type != tuplet_portionOther.Tuplet_type {
		diffs = append(diffs, tuplet_portion.GongMarshallField(stage, "Tuplet_type"))
	}
	if ops := __gong__diffSliceOfPointers(stage, tuplet_portion, "Tuplet_dot", tuplet_portionOther.Tuplet_dot, tuplet_portion.Tuplet_dot); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tuplet_type *Tuplet_type) GongDiff(stage *Stage, tuplet_typeOther *Tuplet_type) (diffs []string) {
	// insertion point for field diffs
	if tuplet_type.Name != tuplet_typeOther.Name {
		diffs = append(diffs, tuplet_type.GongMarshallField(stage, "Name"))
	}
	if tuplet_type.Font_family != tuplet_typeOther.Font_family {
		diffs = append(diffs, tuplet_type.GongMarshallField(stage, "Font_family"))
	}
	if tuplet_type.Font_style != tuplet_typeOther.Font_style {
		diffs = append(diffs, tuplet_type.GongMarshallField(stage, "Font_style"))
	}
	if tuplet_type.Font_size != tuplet_typeOther.Font_size {
		diffs = append(diffs, tuplet_type.GongMarshallField(stage, "Font_size"))
	}
	if tuplet_type.Font_weight != tuplet_typeOther.Font_weight {
		diffs = append(diffs, tuplet_type.GongMarshallField(stage, "Font_weight"))
	}
	if tuplet_type.Color != tuplet_typeOther.Color {
		diffs = append(diffs, tuplet_type.GongMarshallField(stage, "Color"))
	}
	if tuplet_type.EnclosedText != tuplet_typeOther.EnclosedText {
		diffs = append(diffs, tuplet_type.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (typed_text *Typed_text) GongDiff(stage *Stage, typed_textOther *Typed_text) (diffs []string) {
	// insertion point for field diffs
	if typed_text.Name != typed_textOther.Name {
		diffs = append(diffs, typed_text.GongMarshallField(stage, "Name"))
	}
	if typed_text.Type != typed_textOther.Type {
		diffs = append(diffs, typed_text.GongMarshallField(stage, "Type"))
	}
	if typed_text.EnclosedText != typed_textOther.EnclosedText {
		diffs = append(diffs, typed_text.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (unpitched *Unpitched) GongDiff(stage *Stage, unpitchedOther *Unpitched) (diffs []string) {
	// insertion point for field diffs
	if unpitched.Name != unpitchedOther.Name {
		diffs = append(diffs, unpitched.GongMarshallField(stage, "Name"))
	}
	if unpitched.Display_step != unpitchedOther.Display_step {
		diffs = append(diffs, unpitched.GongMarshallField(stage, "Display_step"))
	}
	if unpitched.Display_octave != unpitchedOther.Display_octave {
		diffs = append(diffs, unpitched.GongMarshallField(stage, "Display_octave"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (virtual_instrument *Virtual_instrument) GongDiff(stage *Stage, virtual_instrumentOther *Virtual_instrument) (diffs []string) {
	// insertion point for field diffs
	if virtual_instrument.Name != virtual_instrumentOther.Name {
		diffs = append(diffs, virtual_instrument.GongMarshallField(stage, "Name"))
	}
	if virtual_instrument.Virtual_library != virtual_instrumentOther.Virtual_library {
		diffs = append(diffs, virtual_instrument.GongMarshallField(stage, "Virtual_library"))
	}
	if virtual_instrument.Virtual_name != virtual_instrumentOther.Virtual_name {
		diffs = append(diffs, virtual_instrument.GongMarshallField(stage, "Virtual_name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (wait *Wait) GongDiff(stage *Stage, waitOther *Wait) (diffs []string) {
	// insertion point for field diffs
	if wait.Name != waitOther.Name {
		diffs = append(diffs, wait.GongMarshallField(stage, "Name"))
	}
	if wait.Player != waitOther.Player {
		diffs = append(diffs, wait.GongMarshallField(stage, "Player"))
	}
	if wait.Time_only != waitOther.Time_only {
		diffs = append(diffs, wait.GongMarshallField(stage, "Time_only"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (wavy_line *Wavy_line) GongDiff(stage *Stage, wavy_lineOther *Wavy_line) (diffs []string) {
	// insertion point for field diffs
	if wavy_line.Name != wavy_lineOther.Name {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Name"))
	}
	if wavy_line.Type != wavy_lineOther.Type {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Type"))
	}
	if wavy_line.Number != wavy_lineOther.Number {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Number"))
	}
	if wavy_line.Smufl != wavy_lineOther.Smufl {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Smufl"))
	}
	if wavy_line.Default_x != wavy_lineOther.Default_x {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Default_x"))
	}
	if wavy_line.Default_y != wavy_lineOther.Default_y {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Default_y"))
	}
	if wavy_line.Relative_x != wavy_lineOther.Relative_x {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Relative_x"))
	}
	if wavy_line.Relative_y != wavy_lineOther.Relative_y {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Relative_y"))
	}
	if wavy_line.Placement != wavy_lineOther.Placement {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Placement"))
	}
	if wavy_line.Color != wavy_lineOther.Color {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Color"))
	}
	if wavy_line.Start_note != wavy_lineOther.Start_note {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Start_note"))
	}
	if wavy_line.Trill_step != wavy_lineOther.Trill_step {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Trill_step"))
	}
	if wavy_line.Two_note_turn != wavy_lineOther.Two_note_turn {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Two_note_turn"))
	}
	if wavy_line.Accelerate != wavy_lineOther.Accelerate {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Accelerate"))
	}
	if wavy_line.Beats != wavy_lineOther.Beats {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Beats"))
	}
	if wavy_line.Second_beat != wavy_lineOther.Second_beat {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Second_beat"))
	}
	if wavy_line.Last_beat != wavy_lineOther.Last_beat {
		diffs = append(diffs, wavy_line.GongMarshallField(stage, "Last_beat"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (wedge *Wedge) GongDiff(stage *Stage, wedgeOther *Wedge) (diffs []string) {
	// insertion point for field diffs
	if wedge.Name != wedgeOther.Name {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Name"))
	}
	if wedge.Type != wedgeOther.Type {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Type"))
	}
	if wedge.Number != wedgeOther.Number {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Number"))
	}
	if wedge.Spread != wedgeOther.Spread {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Spread"))
	}
	if wedge.Niente != wedgeOther.Niente {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Niente"))
	}
	if wedge.Line_type != wedgeOther.Line_type {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Line_type"))
	}
	if wedge.Dash_length != wedgeOther.Dash_length {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Dash_length"))
	}
	if wedge.Space_length != wedgeOther.Space_length {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Space_length"))
	}
	if wedge.Default_x != wedgeOther.Default_x {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Default_x"))
	}
	if wedge.Default_y != wedgeOther.Default_y {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Default_y"))
	}
	if wedge.Relative_x != wedgeOther.Relative_x {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Relative_x"))
	}
	if wedge.Relative_y != wedgeOther.Relative_y {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Relative_y"))
	}
	if wedge.Color != wedgeOther.Color {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Color"))
	}
	if wedge.Id != wedgeOther.Id {
		diffs = append(diffs, wedge.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (wood *Wood) GongDiff(stage *Stage, woodOther *Wood) (diffs []string) {
	// insertion point for field diffs
	if wood.Name != woodOther.Name {
		diffs = append(diffs, wood.GongMarshallField(stage, "Name"))
	}
	if wood.Smufl != woodOther.Smufl {
		diffs = append(diffs, wood.GongMarshallField(stage, "Smufl"))
	}
	if wood.EnclosedText != woodOther.EnclosedText {
		diffs = append(diffs, wood.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (work *Work) GongDiff(stage *Stage, workOther *Work) (diffs []string) {
	// insertion point for field diffs
	if work.Name != workOther.Name {
		diffs = append(diffs, work.GongMarshallField(stage, "Name"))
	}
	if work.Work_number != workOther.Work_number {
		diffs = append(diffs, work.GongMarshallField(stage, "Work_number"))
	}
	if work.Work_title != workOther.Work_title {
		diffs = append(diffs, work.GongMarshallField(stage, "Work_title"))
	}
	if work.Opus != workOther.Opus {
		diffs = append(diffs, work.GongMarshallField(stage, "Opus"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
