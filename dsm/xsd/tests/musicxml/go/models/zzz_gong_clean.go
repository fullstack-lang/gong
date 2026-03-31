// generated code - do not edit
package models

import "time"

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !IsStagedPointerToGongstruct(stage, *element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by A_directive
func (a_directive *A_directive) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_measure
func (a_measure *A_measure) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_measure.Note) || modified
	modified = GongCleanSlice(stage, &a_measure.Backup) || modified
	modified = GongCleanSlice(stage, &a_measure.Forward) || modified
	modified = GongCleanSlice(stage, &a_measure.Direction) || modified
	modified = GongCleanSlice(stage, &a_measure.Attributes) || modified
	modified = GongCleanSlice(stage, &a_measure.Harmony) || modified
	modified = GongCleanSlice(stage, &a_measure.Figured_bass) || modified
	modified = GongCleanSlice(stage, &a_measure.Print) || modified
	modified = GongCleanSlice(stage, &a_measure.Sound) || modified
	modified = GongCleanSlice(stage, &a_measure.Listening) || modified
	modified = GongCleanSlice(stage, &a_measure.Barline) || modified
	modified = GongCleanSlice(stage, &a_measure.Grouping) || modified
	modified = GongCleanSlice(stage, &a_measure.Link) || modified
	modified = GongCleanSlice(stage, &a_measure.Bookmark) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_measure_1
func (a_measure_1 *A_measure_1) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_measure_1.Part) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_part
func (a_part *A_part) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_part.Measure) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_part_1
func (a_part_1 *A_part_1) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_part_1.Note) || modified
	modified = GongCleanSlice(stage, &a_part_1.Backup) || modified
	modified = GongCleanSlice(stage, &a_part_1.Forward) || modified
	modified = GongCleanSlice(stage, &a_part_1.Direction) || modified
	modified = GongCleanSlice(stage, &a_part_1.Attributes) || modified
	modified = GongCleanSlice(stage, &a_part_1.Harmony) || modified
	modified = GongCleanSlice(stage, &a_part_1.Figured_bass) || modified
	modified = GongCleanSlice(stage, &a_part_1.Print) || modified
	modified = GongCleanSlice(stage, &a_part_1.Sound) || modified
	modified = GongCleanSlice(stage, &a_part_1.Listening) || modified
	modified = GongCleanSlice(stage, &a_part_1.Barline) || modified
	modified = GongCleanSlice(stage, &a_part_1.Grouping) || modified
	modified = GongCleanSlice(stage, &a_part_1.Link) || modified
	modified = GongCleanSlice(stage, &a_part_1.Bookmark) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Accidental
func (accidental *Accidental) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Accidental_mark
func (accidental_mark *Accidental_mark) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Accidental_text
func (accidental_text *Accidental_text) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Accord
func (accord *Accord) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Accordion_registration
func (accordion_registration *Accordion_registration) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Appearance
func (appearance *Appearance) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &appearance.Line_width) || modified
	modified = GongCleanSlice(stage, &appearance.Note_size) || modified
	modified = GongCleanSlice(stage, &appearance.Distance) || modified
	modified = GongCleanSlice(stage, &appearance.Glyph) || modified
	modified = GongCleanSlice(stage, &appearance.Other_appearance) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Arpeggiate
func (arpeggiate *Arpeggiate) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Arrow
func (arrow *Arrow) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Articulations
func (articulations *Articulations) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &articulations.Accent) || modified
	modified = GongCleanSlice(stage, &articulations.Strong_accent) || modified
	modified = GongCleanSlice(stage, &articulations.Staccato) || modified
	modified = GongCleanSlice(stage, &articulations.Tenuto) || modified
	modified = GongCleanSlice(stage, &articulations.Detached_legato) || modified
	modified = GongCleanSlice(stage, &articulations.Staccatissimo) || modified
	modified = GongCleanSlice(stage, &articulations.Spiccato) || modified
	modified = GongCleanSlice(stage, &articulations.Scoop) || modified
	modified = GongCleanSlice(stage, &articulations.Plop) || modified
	modified = GongCleanSlice(stage, &articulations.Doit) || modified
	modified = GongCleanSlice(stage, &articulations.Falloff) || modified
	modified = GongCleanSlice(stage, &articulations.Breath_mark) || modified
	modified = GongCleanSlice(stage, &articulations.Caesura) || modified
	modified = GongCleanSlice(stage, &articulations.Stress) || modified
	modified = GongCleanSlice(stage, &articulations.Unstress) || modified
	modified = GongCleanSlice(stage, &articulations.Soft_accent) || modified
	modified = GongCleanSlice(stage, &articulations.Other_articulation) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Assess
func (assess *Assess) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Attributes
func (attributes *Attributes) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &attributes.Key) || modified
	modified = GongCleanSlice(stage, &attributes.Time) || modified
	modified = GongCleanSlice(stage, &attributes.Clef) || modified
	modified = GongCleanSlice(stage, &attributes.Staff_details) || modified
	modified = GongCleanSlice(stage, &attributes.Transpose) || modified
	modified = GongCleanSlice(stage, &attributes.For_part) || modified
	modified = GongCleanSlice(stage, &attributes.Directive) || modified
	modified = GongCleanSlice(stage, &attributes.Measure_style) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &attributes.Footnote) || modified
	modified = GongCleanPointer(stage, &attributes.Level) || modified
	modified = GongCleanPointer(stage, &attributes.Part_symbol) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Backup
func (backup *Backup) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &backup.Footnote) || modified
	modified = GongCleanPointer(stage, &backup.Level) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Bar_style_color
func (bar_style_color *Bar_style_color) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Barline
func (barline *Barline) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &barline.Bar_style) || modified
	modified = GongCleanPointer(stage, &barline.Footnote) || modified
	modified = GongCleanPointer(stage, &barline.Level) || modified
	modified = GongCleanPointer(stage, &barline.Wavy_line) || modified
	modified = GongCleanPointer(stage, &barline.Segno_1) || modified
	modified = GongCleanPointer(stage, &barline.Coda_1) || modified
	modified = GongCleanPointer(stage, &barline.Fermata) || modified
	modified = GongCleanPointer(stage, &barline.Ending) || modified
	modified = GongCleanPointer(stage, &barline.Repeat) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Barre
func (barre *Barre) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Bass
func (bass *Bass) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &bass.Bass_separator) || modified
	modified = GongCleanPointer(stage, &bass.Bass_step) || modified
	modified = GongCleanPointer(stage, &bass.Bass_alter) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Bass_step
func (bass_step *Bass_step) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Beam
func (beam *Beam) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Beat_repeat
func (beat_repeat *Beat_repeat) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Beat_unit_tied
func (beat_unit_tied *Beat_unit_tied) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Beater
func (beater *Beater) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Bend
func (bend *Bend) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &bend.Release) || modified
	modified = GongCleanPointer(stage, &bend.With_bar) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Bookmark
func (bookmark *Bookmark) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Bracket
func (bracket *Bracket) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Breath_mark
func (breath_mark *Breath_mark) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Caesura
func (caesura *Caesura) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Cancel
func (cancel *Cancel) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Clef
func (clef *Clef) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Coda
func (coda *Coda) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Credit
func (credit *Credit) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &credit.Link) || modified
	modified = GongCleanSlice(stage, &credit.Bookmark) || modified
	modified = GongCleanSlice(stage, &credit.Credit_words) || modified
	modified = GongCleanSlice(stage, &credit.Credit_symbol) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &credit.Credit_image) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Dashes
func (dashes *Dashes) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Defaults
func (defaults *Defaults) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &defaults.Staff_layout) || modified
	modified = GongCleanSlice(stage, &defaults.Lyric_font) || modified
	modified = GongCleanSlice(stage, &defaults.Lyric_language) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &defaults.Scaling) || modified
	modified = GongCleanPointer(stage, &defaults.Page_layout) || modified
	modified = GongCleanPointer(stage, &defaults.System_layout) || modified
	modified = GongCleanPointer(stage, &defaults.Appearance) || modified
	modified = GongCleanPointer(stage, &defaults.Music_font) || modified
	modified = GongCleanPointer(stage, &defaults.Word_font) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Degree
func (degree *Degree) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &degree.Degree_value) || modified
	modified = GongCleanPointer(stage, &degree.Degree_alter) || modified
	modified = GongCleanPointer(stage, &degree.Degree_type) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Degree_alter
func (degree_alter *Degree_alter) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Degree_type
func (degree_type *Degree_type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Degree_value
func (degree_value *Degree_value) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Direction
func (direction *Direction) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &direction.Direction_type) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &direction.Offset) || modified
	modified = GongCleanPointer(stage, &direction.Footnote) || modified
	modified = GongCleanPointer(stage, &direction.Level) || modified
	modified = GongCleanPointer(stage, &direction.Sound) || modified
	modified = GongCleanPointer(stage, &direction.Listening) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Direction_type
func (direction_type *Direction_type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &direction_type.Rehearsal) || modified
	modified = GongCleanSlice(stage, &direction_type.Segno) || modified
	modified = GongCleanSlice(stage, &direction_type.Coda) || modified
	modified = GongCleanSlice(stage, &direction_type.Words) || modified
	modified = GongCleanSlice(stage, &direction_type.Symbol) || modified
	modified = GongCleanSlice(stage, &direction_type.Dynamics) || modified
	modified = GongCleanSlice(stage, &direction_type.Percussion) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &direction_type.Wedge) || modified
	modified = GongCleanPointer(stage, &direction_type.Dashes) || modified
	modified = GongCleanPointer(stage, &direction_type.Bracket) || modified
	modified = GongCleanPointer(stage, &direction_type.Pedal) || modified
	modified = GongCleanPointer(stage, &direction_type.Metronome) || modified
	modified = GongCleanPointer(stage, &direction_type.Octave_shift) || modified
	modified = GongCleanPointer(stage, &direction_type.Harp_pedals) || modified
	modified = GongCleanPointer(stage, &direction_type.Damp) || modified
	modified = GongCleanPointer(stage, &direction_type.Damp_all) || modified
	modified = GongCleanPointer(stage, &direction_type.Eyeglasses) || modified
	modified = GongCleanPointer(stage, &direction_type.String_mute) || modified
	modified = GongCleanPointer(stage, &direction_type.Scordatura) || modified
	modified = GongCleanPointer(stage, &direction_type.Image) || modified
	modified = GongCleanPointer(stage, &direction_type.Principal_voice) || modified
	modified = GongCleanPointer(stage, &direction_type.Accordion_registration) || modified
	modified = GongCleanPointer(stage, &direction_type.Staff_divide) || modified
	modified = GongCleanPointer(stage, &direction_type.Other_direction) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Distance
func (distance *Distance) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Double
func (double *Double) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Dynamics
func (dynamics *Dynamics) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &dynamics.Other_dynamics) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Effect
func (effect *Effect) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Elision
func (elision *Elision) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Empty
func (empty *Empty) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Empty_font
func (empty_font *Empty_font) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Empty_line
func (empty_line *Empty_line) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Empty_placement
func (empty_placement *Empty_placement) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Empty_placement_smufl
func (empty_placement_smufl *Empty_placement_smufl) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Empty_print_object_style_align
func (empty_print_object_style_align *Empty_print_object_style_align) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Empty_print_style
func (empty_print_style *Empty_print_style) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Empty_print_style_align
func (empty_print_style_align *Empty_print_style_align) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Empty_print_style_align_id
func (empty_print_style_align_id *Empty_print_style_align_id) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Empty_trill_sound
func (empty_trill_sound *Empty_trill_sound) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Encoding
func (encoding *Encoding) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &encoding.Encoder) || modified
	modified = GongCleanSlice(stage, &encoding.Supports) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Ending
func (ending *Ending) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Extend
func (extend *Extend) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Feature
func (feature *Feature) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Fermata
func (fermata *Fermata) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Figure
func (figure *Figure) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &figure.Prefix) || modified
	modified = GongCleanPointer(stage, &figure.Figure_number) || modified
	modified = GongCleanPointer(stage, &figure.Suffix) || modified
	modified = GongCleanPointer(stage, &figure.Extend) || modified
	modified = GongCleanPointer(stage, &figure.Footnote) || modified
	modified = GongCleanPointer(stage, &figure.Level) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Figured_bass
func (figured_bass *Figured_bass) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &figured_bass.Figure) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &figured_bass.Footnote) || modified
	modified = GongCleanPointer(stage, &figured_bass.Level) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Fingering
func (fingering *Fingering) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by First_fret
func (first_fret *First_fret) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by For_part
func (for_part *For_part) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &for_part.Part_clef) || modified
	modified = GongCleanPointer(stage, &for_part.Part_transpose) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Formatted_symbol
func (formatted_symbol *Formatted_symbol) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Formatted_symbol_id
func (formatted_symbol_id *Formatted_symbol_id) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Formatted_text
func (formatted_text *Formatted_text) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Formatted_text_id
func (formatted_text_id *Formatted_text_id) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Forward
func (forward *Forward) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &forward.Footnote) || modified
	modified = GongCleanPointer(stage, &forward.Level) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Frame
func (frame *Frame) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &frame.Frame_note) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &frame.First_fret) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Frame_note
func (frame_note *Frame_note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &frame_note.String) || modified
	modified = GongCleanPointer(stage, &frame_note.Fret) || modified
	modified = GongCleanPointer(stage, &frame_note.Fingering) || modified
	modified = GongCleanPointer(stage, &frame_note.Barre) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Fret
func (fret *Fret) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Glass
func (glass *Glass) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Glissando
func (glissando *Glissando) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Glyph
func (glyph *Glyph) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Grace
func (grace *Grace) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Group_barline
func (group_barline *Group_barline) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Group_name
func (group_name *Group_name) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Group_symbol
func (group_symbol *Group_symbol) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Grouping
func (grouping *Grouping) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &grouping.Feature) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Hammer_on_pull_off
func (hammer_on_pull_off *Hammer_on_pull_off) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Handbell
func (handbell *Handbell) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Harmon_closed
func (harmon_closed *Harmon_closed) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Harmon_mute
func (harmon_mute *Harmon_mute) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &harmon_mute.Harmon_closed) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Harmonic
func (harmonic *Harmonic) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Harmony
func (harmony *Harmony) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &harmony.Degree) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &harmony.Root) || modified
	modified = GongCleanPointer(stage, &harmony.Numeral) || modified
	modified = GongCleanPointer(stage, &harmony.Function) || modified
	modified = GongCleanPointer(stage, &harmony.Kind) || modified
	modified = GongCleanPointer(stage, &harmony.Inversion) || modified
	modified = GongCleanPointer(stage, &harmony.Bass) || modified
	modified = GongCleanPointer(stage, &harmony.Frame) || modified
	modified = GongCleanPointer(stage, &harmony.Offset) || modified
	modified = GongCleanPointer(stage, &harmony.Footnote) || modified
	modified = GongCleanPointer(stage, &harmony.Level) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Harmony_alter
func (harmony_alter *Harmony_alter) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Harp_pedals
func (harp_pedals *Harp_pedals) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &harp_pedals.Pedal_tuning) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Heel_toe
func (heel_toe *Heel_toe) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Hole
func (hole *Hole) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &hole.Hole_closed) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Hole_closed
func (hole_closed *Hole_closed) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Horizontal_turn
func (horizontal_turn *Horizontal_turn) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Identification
func (identification *Identification) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &identification.Creator) || modified
	modified = GongCleanSlice(stage, &identification.Rights) || modified
	modified = GongCleanSlice(stage, &identification.Relation) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &identification.Encoding) || modified
	modified = GongCleanPointer(stage, &identification.Miscellaneous) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Image
func (image *Image) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Instrument
func (instrument *Instrument) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Instrument_change
func (instrument_change *Instrument_change) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &instrument_change.Virtual_instrument) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Instrument_link
func (instrument_link *Instrument_link) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Interchangeable
func (interchangeable *Interchangeable) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Inversion
func (inversion *Inversion) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Key
func (key *Key) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &key.Key_octave) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &key.Cancel) || modified
	modified = GongCleanPointer(stage, &key.Key_accidental) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Key_accidental
func (key_accidental *Key_accidental) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Key_octave
func (key_octave *Key_octave) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Kind
func (kind *Kind) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Level
func (level *Level) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Line_detail
func (line_detail *Line_detail) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Line_width
func (line_width *Line_width) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Link
func (link *Link) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Listen
func (listen *Listen) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &listen.Assess) || modified
	modified = GongCleanSlice(stage, &listen.Wait) || modified
	modified = GongCleanSlice(stage, &listen.Other_listen) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Listening
func (listening *Listening) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &listening.Sync) || modified
	modified = GongCleanSlice(stage, &listening.Other_listening) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &listening.Offset) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Lyric
func (lyric *Lyric) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &lyric.Elision) || modified
	modified = GongCleanSlice(stage, &lyric.Text) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &lyric.Extend) || modified
	modified = GongCleanPointer(stage, &lyric.Footnote) || modified
	modified = GongCleanPointer(stage, &lyric.Level) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Lyric_font
func (lyric_font *Lyric_font) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Lyric_language
func (lyric_language *Lyric_language) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Measure_layout
func (measure_layout *Measure_layout) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Measure_numbering
func (measure_numbering *Measure_numbering) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Measure_repeat
func (measure_repeat *Measure_repeat) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Measure_style
func (measure_style *Measure_style) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &measure_style.Multiple_rest) || modified
	modified = GongCleanPointer(stage, &measure_style.Measure_repeat) || modified
	modified = GongCleanPointer(stage, &measure_style.Beat_repeat) || modified
	modified = GongCleanPointer(stage, &measure_style.Slash) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Membrane
func (membrane *Membrane) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Metal
func (metal *Metal) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Metronome
func (metronome *Metronome) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &metronome.Beat_unit_tied) || modified
	modified = GongCleanSlice(stage, &metronome.Metronome_note) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &metronome.Per_minute) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Metronome_beam
func (metronome_beam *Metronome_beam) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Metronome_note
func (metronome_note *Metronome_note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &metronome_note.Metronome_beam) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &metronome_note.Metronome_tied) || modified
	modified = GongCleanPointer(stage, &metronome_note.Metronome_tuplet) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Metronome_tied
func (metronome_tied *Metronome_tied) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Metronome_tuplet
func (metronome_tuplet *Metronome_tuplet) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Midi_device
func (midi_device *Midi_device) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Midi_instrument
func (midi_instrument *Midi_instrument) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Miscellaneous
func (miscellaneous *Miscellaneous) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &miscellaneous.Miscellaneous_field) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Miscellaneous_field
func (miscellaneous_field *Miscellaneous_field) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Mordent
func (mordent *Mordent) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Multiple_rest
func (multiple_rest *Multiple_rest) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Name_display
func (name_display *Name_display) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &name_display.Display_text) || modified
	modified = GongCleanSlice(stage, &name_display.Accidental_text) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Non_arpeggiate
func (non_arpeggiate *Non_arpeggiate) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Notations
func (notations *Notations) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &notations.Tied) || modified
	modified = GongCleanSlice(stage, &notations.Slur) || modified
	modified = GongCleanSlice(stage, &notations.Tuplet) || modified
	modified = GongCleanSlice(stage, &notations.Glissando) || modified
	modified = GongCleanSlice(stage, &notations.Slide) || modified
	modified = GongCleanSlice(stage, &notations.Ornaments) || modified
	modified = GongCleanSlice(stage, &notations.Technical) || modified
	modified = GongCleanSlice(stage, &notations.Articulations) || modified
	modified = GongCleanSlice(stage, &notations.Dynamics) || modified
	modified = GongCleanSlice(stage, &notations.Fermata) || modified
	modified = GongCleanSlice(stage, &notations.Arpeggiate) || modified
	modified = GongCleanSlice(stage, &notations.Non_arpeggiate) || modified
	modified = GongCleanSlice(stage, &notations.Accidental_mark) || modified
	modified = GongCleanSlice(stage, &notations.Other_notation) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &notations.Footnote) || modified
	modified = GongCleanPointer(stage, &notations.Level) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Note
func (note *Note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &note.Instrument) || modified
	modified = GongCleanSlice(stage, &note.Dot) || modified
	modified = GongCleanSlice(stage, &note.Notations) || modified
	modified = GongCleanSlice(stage, &note.Lyric) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &note.Grace) || modified
	modified = GongCleanPointer(stage, &note.Pitch) || modified
	modified = GongCleanPointer(stage, &note.Unpitched) || modified
	modified = GongCleanPointer(stage, &note.Rest) || modified
	modified = GongCleanPointer(stage, &note.Tie) || modified
	modified = GongCleanPointer(stage, &note.Footnote) || modified
	modified = GongCleanPointer(stage, &note.Level) || modified
	modified = GongCleanPointer(stage, &note.Type) || modified
	modified = GongCleanPointer(stage, &note.Accidental) || modified
	modified = GongCleanPointer(stage, &note.Time_modification) || modified
	modified = GongCleanPointer(stage, &note.Stem) || modified
	modified = GongCleanPointer(stage, &note.Notehead) || modified
	modified = GongCleanPointer(stage, &note.Notehead_text) || modified
	modified = GongCleanPointer(stage, &note.Beam) || modified
	modified = GongCleanPointer(stage, &note.Play) || modified
	modified = GongCleanPointer(stage, &note.Listen) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Note_size
func (note_size *Note_size) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Note_type
func (note_type *Note_type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Notehead
func (notehead *Notehead) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Notehead_text
func (notehead_text *Notehead_text) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &notehead_text.Display_text) || modified
	modified = GongCleanSlice(stage, &notehead_text.Accidental_text) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Numeral
func (numeral *Numeral) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &numeral.Numeral_root) || modified
	modified = GongCleanPointer(stage, &numeral.Numeral_alter) || modified
	modified = GongCleanPointer(stage, &numeral.Numeral_key) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Numeral_key
func (numeral_key *Numeral_key) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Numeral_root
func (numeral_root *Numeral_root) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Octave_shift
func (octave_shift *Octave_shift) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Offset
func (offset *Offset) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Opus
func (opus *Opus) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Ornaments
func (ornaments *Ornaments) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &ornaments.Trill_mark) || modified
	modified = GongCleanSlice(stage, &ornaments.Turn) || modified
	modified = GongCleanSlice(stage, &ornaments.Delayed_turn) || modified
	modified = GongCleanSlice(stage, &ornaments.Inverted_turn) || modified
	modified = GongCleanSlice(stage, &ornaments.Delayed_inverted_turn) || modified
	modified = GongCleanSlice(stage, &ornaments.Vertical_turn) || modified
	modified = GongCleanSlice(stage, &ornaments.Inverted_vertical_turn) || modified
	modified = GongCleanSlice(stage, &ornaments.Shake) || modified
	modified = GongCleanSlice(stage, &ornaments.Wavy_line) || modified
	modified = GongCleanSlice(stage, &ornaments.Mordent) || modified
	modified = GongCleanSlice(stage, &ornaments.Inverted_mordent) || modified
	modified = GongCleanSlice(stage, &ornaments.Schleifer) || modified
	modified = GongCleanSlice(stage, &ornaments.Tremolo) || modified
	modified = GongCleanSlice(stage, &ornaments.Haydn) || modified
	modified = GongCleanSlice(stage, &ornaments.Other_ornament) || modified
	modified = GongCleanSlice(stage, &ornaments.Accidental_mark) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Other_appearance
func (other_appearance *Other_appearance) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Other_direction
func (other_direction *Other_direction) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Other_listening
func (other_listening *Other_listening) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Other_notation
func (other_notation *Other_notation) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Other_placement_text
func (other_placement_text *Other_placement_text) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Other_play
func (other_play *Other_play) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Other_text
func (other_text *Other_text) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Page_layout
func (page_layout *Page_layout) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &page_layout.Page_margins) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Page_margins
func (page_margins *Page_margins) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Part_clef
func (part_clef *Part_clef) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Part_group
func (part_group *Part_group) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &part_group.Group_name) || modified
	modified = GongCleanPointer(stage, &part_group.Group_name_display) || modified
	modified = GongCleanPointer(stage, &part_group.Group_abbreviation) || modified
	modified = GongCleanPointer(stage, &part_group.Group_abbreviation_display) || modified
	modified = GongCleanPointer(stage, &part_group.Group_symbol) || modified
	modified = GongCleanPointer(stage, &part_group.Group_barline) || modified
	modified = GongCleanPointer(stage, &part_group.Footnote) || modified
	modified = GongCleanPointer(stage, &part_group.Level) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Part_link
func (part_link *Part_link) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &part_link.Instrument_link) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Part_list
func (part_list *Part_list) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &part_list.Part_group) || modified
	modified = GongCleanPointer(stage, &part_list.Score_part) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Part_name
func (part_name *Part_name) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Part_symbol
func (part_symbol *Part_symbol) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Part_transpose
func (part_transpose *Part_transpose) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Pedal
func (pedal *Pedal) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Pedal_tuning
func (pedal_tuning *Pedal_tuning) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Per_minute
func (per_minute *Per_minute) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Percussion
func (percussion *Percussion) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &percussion.Glass) || modified
	modified = GongCleanPointer(stage, &percussion.Metal) || modified
	modified = GongCleanPointer(stage, &percussion.Wood) || modified
	modified = GongCleanPointer(stage, &percussion.Pitched) || modified
	modified = GongCleanPointer(stage, &percussion.Membrane) || modified
	modified = GongCleanPointer(stage, &percussion.Effect) || modified
	modified = GongCleanPointer(stage, &percussion.Timpani) || modified
	modified = GongCleanPointer(stage, &percussion.Beater) || modified
	modified = GongCleanPointer(stage, &percussion.Stick) || modified
	modified = GongCleanPointer(stage, &percussion.Other_percussion) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Pitch
func (pitch *Pitch) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Pitched
func (pitched *Pitched) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Placement_text
func (placement_text *Placement_text) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Play
func (play *Play) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &play.Other_play) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Player
func (player *Player) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Principal_voice
func (principal_voice *Principal_voice) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Print
func (print *Print) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &print.Staff_layout) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &print.Page_layout) || modified
	modified = GongCleanPointer(stage, &print.System_layout) || modified
	modified = GongCleanPointer(stage, &print.Measure_layout) || modified
	modified = GongCleanPointer(stage, &print.Measure_numbering) || modified
	modified = GongCleanPointer(stage, &print.Part_name_display) || modified
	modified = GongCleanPointer(stage, &print.Part_abbreviation_display) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Release
func (release *Release) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Repeat
func (repeat *Repeat) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Rest
func (rest *Rest) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Root
func (root *Root) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &root.Root_step) || modified
	modified = GongCleanPointer(stage, &root.Root_alter) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Root_step
func (root_step *Root_step) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Scaling
func (scaling *Scaling) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Scordatura
func (scordatura *Scordatura) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &scordatura.Accord) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Score_instrument
func (score_instrument *Score_instrument) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &score_instrument.Virtual_instrument) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Score_part
func (score_part *Score_part) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &score_part.Part_link) || modified
	modified = GongCleanSlice(stage, &score_part.Score_instrument) || modified
	modified = GongCleanSlice(stage, &score_part.Player) || modified
	modified = GongCleanSlice(stage, &score_part.Midi_device) || modified
	modified = GongCleanSlice(stage, &score_part.Midi_instrument) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &score_part.Identification) || modified
	modified = GongCleanPointer(stage, &score_part.Part_name) || modified
	modified = GongCleanPointer(stage, &score_part.Part_name_display) || modified
	modified = GongCleanPointer(stage, &score_part.Part_abbreviation) || modified
	modified = GongCleanPointer(stage, &score_part.Part_abbreviation_display) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Score_partwise
func (score_partwise *Score_partwise) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &score_partwise.Credit) || modified
	modified = GongCleanSlice(stage, &score_partwise.Part) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &score_partwise.Work) || modified
	modified = GongCleanPointer(stage, &score_partwise.Identification) || modified
	modified = GongCleanPointer(stage, &score_partwise.Defaults) || modified
	modified = GongCleanPointer(stage, &score_partwise.Part_list) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Score_timewise
func (score_timewise *Score_timewise) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &score_timewise.Credit) || modified
	modified = GongCleanSlice(stage, &score_timewise.Measure) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &score_timewise.Work) || modified
	modified = GongCleanPointer(stage, &score_timewise.Identification) || modified
	modified = GongCleanPointer(stage, &score_timewise.Defaults) || modified
	modified = GongCleanPointer(stage, &score_timewise.Part_list) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Segno
func (segno *Segno) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Slash
func (slash *Slash) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Slide
func (slide *Slide) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Slur
func (slur *Slur) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Sound
func (sound *Sound) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &sound.Instrument_change) || modified
	modified = GongCleanSlice(stage, &sound.Midi_device) || modified
	modified = GongCleanSlice(stage, &sound.Midi_instrument) || modified
	modified = GongCleanSlice(stage, &sound.Play) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &sound.Swing) || modified
	modified = GongCleanPointer(stage, &sound.Offset) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Staff_details
func (staff_details *Staff_details) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &staff_details.Line_detail) || modified
	modified = GongCleanSlice(stage, &staff_details.Staff_tuning) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &staff_details.Staff_size) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Staff_divide
func (staff_divide *Staff_divide) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Staff_layout
func (staff_layout *Staff_layout) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Staff_size
func (staff_size *Staff_size) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Staff_tuning
func (staff_tuning *Staff_tuning) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Stem
func (stem *Stem) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Stick
func (stick *Stick) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by String_mute
func (string_mute *String_mute) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by String_type
func (string_type *String_type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Strong_accent
func (strong_accent *Strong_accent) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Style_text
func (style_text *Style_text) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Supports
func (supports *Supports) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Swing
func (swing *Swing) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Sync
func (sync *Sync) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by System_dividers
func (system_dividers *System_dividers) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &system_dividers.Left_divider) || modified
	modified = GongCleanPointer(stage, &system_dividers.Right_divider) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by System_layout
func (system_layout *System_layout) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &system_layout.System_margins) || modified
	modified = GongCleanPointer(stage, &system_layout.System_dividers) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by System_margins
func (system_margins *System_margins) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Tap
func (tap *Tap) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Technical
func (technical *Technical) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &technical.Up_bow) || modified
	modified = GongCleanSlice(stage, &technical.Down_bow) || modified
	modified = GongCleanSlice(stage, &technical.Harmonic) || modified
	modified = GongCleanSlice(stage, &technical.Open_string) || modified
	modified = GongCleanSlice(stage, &technical.Thumb_position) || modified
	modified = GongCleanSlice(stage, &technical.Fingering) || modified
	modified = GongCleanSlice(stage, &technical.Pluck) || modified
	modified = GongCleanSlice(stage, &technical.Double_tongue) || modified
	modified = GongCleanSlice(stage, &technical.Triple_tongue) || modified
	modified = GongCleanSlice(stage, &technical.Stopped) || modified
	modified = GongCleanSlice(stage, &technical.Snap_pizzicato) || modified
	modified = GongCleanSlice(stage, &technical.Fret) || modified
	modified = GongCleanSlice(stage, &technical.String) || modified
	modified = GongCleanSlice(stage, &technical.Hammer_on) || modified
	modified = GongCleanSlice(stage, &technical.Pull_off) || modified
	modified = GongCleanSlice(stage, &technical.Bend) || modified
	modified = GongCleanSlice(stage, &technical.Tap) || modified
	modified = GongCleanSlice(stage, &technical.Heel) || modified
	modified = GongCleanSlice(stage, &technical.Toe) || modified
	modified = GongCleanSlice(stage, &technical.Fingernails) || modified
	modified = GongCleanSlice(stage, &technical.Hole) || modified
	modified = GongCleanSlice(stage, &technical.Arrow) || modified
	modified = GongCleanSlice(stage, &technical.Handbell) || modified
	modified = GongCleanSlice(stage, &technical.Brass_bend) || modified
	modified = GongCleanSlice(stage, &technical.Flip) || modified
	modified = GongCleanSlice(stage, &technical.Smear) || modified
	modified = GongCleanSlice(stage, &technical.Open) || modified
	modified = GongCleanSlice(stage, &technical.Half_muted) || modified
	modified = GongCleanSlice(stage, &technical.Harmon_mute) || modified
	modified = GongCleanSlice(stage, &technical.Golpe) || modified
	modified = GongCleanSlice(stage, &technical.Other_technical) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Text_element_data
func (text_element_data *Text_element_data) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Tie
func (tie *Tie) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Tied
func (tied *Tied) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Time
func (time *Time) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &time.Interchangeable) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Time_modification
func (time_modification *Time_modification) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Timpani
func (timpani *Timpani) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Transpose
func (transpose *Transpose) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Tremolo
func (tremolo *Tremolo) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Tuplet
func (tuplet *Tuplet) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &tuplet.Tuplet_actual) || modified
	modified = GongCleanPointer(stage, &tuplet.Tuplet_normal) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Tuplet_dot
func (tuplet_dot *Tuplet_dot) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Tuplet_number
func (tuplet_number *Tuplet_number) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Tuplet_portion
func (tuplet_portion *Tuplet_portion) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &tuplet_portion.Tuplet_dot) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &tuplet_portion.Tuplet_number) || modified
	modified = GongCleanPointer(stage, &tuplet_portion.Tuplet_type) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Tuplet_type
func (tuplet_type *Tuplet_type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Typed_text
func (typed_text *Typed_text) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Unpitched
func (unpitched *Unpitched) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Virtual_instrument
func (virtual_instrument *Virtual_instrument) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Wait
func (wait *Wait) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Wavy_line
func (wavy_line *Wavy_line) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Wedge
func (wedge *Wedge) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Wood
func (wood *Wood) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Work
func (work *Work) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &work.Opus) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
