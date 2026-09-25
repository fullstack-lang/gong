// generated code - do not edit
package models

// insertion point
func (inst *A_directive) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Directive":
			if _attributes, ok := stage.Attributes_Directive_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *A_measure) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_part":
		switch reverseField.Fieldname {
		case "Measure":
			if _a_part, ok := stage.A_part_Measure_reverseMap[inst]; ok {
				res = _a_part.Name
			}
		}
	}
	return
}

func (inst *A_measure_1) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Score_timewise":
		switch reverseField.Fieldname {
		case "Measure":
			if _score_timewise, ok := stage.Score_timewise_Measure_reverseMap[inst]; ok {
				res = _score_timewise.Name
			}
		}
	}
	return
}

func (inst *A_part) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Score_partwise":
		switch reverseField.Fieldname {
		case "Part":
			if _score_partwise, ok := stage.Score_partwise_Part_reverseMap[inst]; ok {
				res = _score_partwise.Name
			}
		}
	}
	return
}

func (inst *A_part_1) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure_1":
		switch reverseField.Fieldname {
		case "Part":
			if _a_measure_1, ok := stage.A_measure_1_Part_reverseMap[inst]; ok {
				res = _a_measure_1.Name
			}
		}
	}
	return
}

func (inst *Accidental) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Accidental_mark) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Accidental_mark":
			if _notations, ok := stage.Notations_Accidental_mark_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Accidental_mark":
			if _ornaments, ok := stage.Ornaments_Accidental_mark_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		}
	}
	return
}

func (inst *Accidental_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Name_display":
		switch reverseField.Fieldname {
		case "Accidental_text":
			if _name_display, ok := stage.Name_display_Accidental_text_reverseMap[inst]; ok {
				res = _name_display.Name
			}
		}
	case "Notehead_text":
		switch reverseField.Fieldname {
		case "Accidental_text":
			if _notehead_text, ok := stage.Notehead_text_Accidental_text_reverseMap[inst]; ok {
				res = _notehead_text.Name
			}
		}
	}
	return
}

func (inst *Accord) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Scordatura":
		switch reverseField.Fieldname {
		case "Accord":
			if _scordatura, ok := stage.Scordatura_Accord_reverseMap[inst]; ok {
				res = _scordatura.Name
			}
		}
	}
	return
}

func (inst *Accordion_registration) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Appearance) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Arpeggiate) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Arpeggiate":
			if _notations, ok := stage.Notations_Arpeggiate_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *Arrow) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Arrow":
			if _technical, ok := stage.Technical_Arrow_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Articulations) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Articulations":
			if _notations, ok := stage.Notations_Articulations_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *Assess) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Listen":
		switch reverseField.Fieldname {
		case "Assess":
			if _listen, ok := stage.Listen_Assess_reverseMap[inst]; ok {
				res = _listen.Name
			}
		}
	}
	return
}

func (inst *Attributes) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Attributes":
			if _a_measure, ok := stage.A_measure_Attributes_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Attributes":
			if _a_part_1, ok := stage.A_part_1_Attributes_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	}
	return
}

func (inst *Backup) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Backup":
			if _a_measure, ok := stage.A_measure_Backup_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Backup":
			if _a_part_1, ok := stage.A_part_1_Backup_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	}
	return
}

func (inst *Bar_style_color) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Barline) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Barline":
			if _a_measure, ok := stage.A_measure_Barline_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Barline":
			if _a_part_1, ok := stage.A_part_1_Barline_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	}
	return
}

func (inst *Barre) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Bass) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Bass_step) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Beam) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Beat_repeat) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Beat_unit_tied) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Metronome":
		switch reverseField.Fieldname {
		case "Beat_unit_tied":
			if _metronome, ok := stage.Metronome_Beat_unit_tied_reverseMap[inst]; ok {
				res = _metronome.Name
			}
		}
	}
	return
}

func (inst *Beater) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Bend) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Bend":
			if _technical, ok := stage.Technical_Bend_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Bookmark) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Bookmark":
			if _a_measure, ok := stage.A_measure_Bookmark_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Bookmark":
			if _a_part_1, ok := stage.A_part_1_Bookmark_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	case "Credit":
		switch reverseField.Fieldname {
		case "Bookmark":
			if _credit, ok := stage.Credit_Bookmark_reverseMap[inst]; ok {
				res = _credit.Name
			}
		}
	}
	return
}

func (inst *Bracket) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Breath_mark) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "Breath_mark":
			if _articulations, ok := stage.Articulations_Breath_mark_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		}
	}
	return
}

func (inst *Caesura) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "Caesura":
			if _articulations, ok := stage.Articulations_Caesura_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		}
	}
	return
}

func (inst *Cancel) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Clef) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Clef":
			if _attributes, ok := stage.Attributes_Clef_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *Coda) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Direction_type":
		switch reverseField.Fieldname {
		case "Coda":
			if _direction_type, ok := stage.Direction_type_Coda_reverseMap[inst]; ok {
				res = _direction_type.Name
			}
		}
	}
	return
}

func (inst *Credit) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Score_partwise":
		switch reverseField.Fieldname {
		case "Credit":
			if _score_partwise, ok := stage.Score_partwise_Credit_reverseMap[inst]; ok {
				res = _score_partwise.Name
			}
		}
	case "Score_timewise":
		switch reverseField.Fieldname {
		case "Credit":
			if _score_timewise, ok := stage.Score_timewise_Credit_reverseMap[inst]; ok {
				res = _score_timewise.Name
			}
		}
	}
	return
}

func (inst *Dashes) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Defaults) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Degree) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Harmony":
		switch reverseField.Fieldname {
		case "Degree":
			if _harmony, ok := stage.Harmony_Degree_reverseMap[inst]; ok {
				res = _harmony.Name
			}
		}
	}
	return
}

func (inst *Degree_alter) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Degree_type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Degree_value) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Direction) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Direction":
			if _a_measure, ok := stage.A_measure_Direction_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Direction":
			if _a_part_1, ok := stage.A_part_1_Direction_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	}
	return
}

func (inst *Direction_type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Direction":
		switch reverseField.Fieldname {
		case "Direction_type":
			if _direction, ok := stage.Direction_Direction_type_reverseMap[inst]; ok {
				res = _direction.Name
			}
		}
	}
	return
}

func (inst *Distance) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Appearance":
		switch reverseField.Fieldname {
		case "Distance":
			if _appearance, ok := stage.Appearance_Distance_reverseMap[inst]; ok {
				res = _appearance.Name
			}
		}
	}
	return
}

func (inst *Double) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Dynamics) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Direction_type":
		switch reverseField.Fieldname {
		case "Dynamics":
			if _direction_type, ok := stage.Direction_type_Dynamics_reverseMap[inst]; ok {
				res = _direction_type.Name
			}
		}
	case "Notations":
		switch reverseField.Fieldname {
		case "Dynamics":
			if _notations, ok := stage.Notations_Dynamics_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *Effect) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Elision) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Lyric":
		switch reverseField.Fieldname {
		case "Elision":
			if _lyric, ok := stage.Lyric_Elision_reverseMap[inst]; ok {
				res = _lyric.Name
			}
		}
	}
	return
}

func (inst *Empty) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Empty_font) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Empty_line) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "Scoop":
			if _articulations, ok := stage.Articulations_Scoop_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		case "Plop":
			if _articulations, ok := stage.Articulations_Plop_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		case "Doit":
			if _articulations, ok := stage.Articulations_Doit_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		case "Falloff":
			if _articulations, ok := stage.Articulations_Falloff_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		}
	}
	return
}

func (inst *Empty_placement) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "Accent":
			if _articulations, ok := stage.Articulations_Accent_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		case "Staccato":
			if _articulations, ok := stage.Articulations_Staccato_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		case "Tenuto":
			if _articulations, ok := stage.Articulations_Tenuto_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		case "Detached_legato":
			if _articulations, ok := stage.Articulations_Detached_legato_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		case "Staccatissimo":
			if _articulations, ok := stage.Articulations_Staccatissimo_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		case "Spiccato":
			if _articulations, ok := stage.Articulations_Spiccato_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		case "Stress":
			if _articulations, ok := stage.Articulations_Stress_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		case "Unstress":
			if _articulations, ok := stage.Articulations_Unstress_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		case "Soft_accent":
			if _articulations, ok := stage.Articulations_Soft_accent_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Dot":
			if _note, ok := stage.Note_Dot_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Schleifer":
			if _ornaments, ok := stage.Ornaments_Schleifer_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		}
	case "Technical":
		switch reverseField.Fieldname {
		case "Up_bow":
			if _technical, ok := stage.Technical_Up_bow_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Down_bow":
			if _technical, ok := stage.Technical_Down_bow_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Open_string":
			if _technical, ok := stage.Technical_Open_string_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Thumb_position":
			if _technical, ok := stage.Technical_Thumb_position_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Double_tongue":
			if _technical, ok := stage.Technical_Double_tongue_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Triple_tongue":
			if _technical, ok := stage.Technical_Triple_tongue_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Snap_pizzicato":
			if _technical, ok := stage.Technical_Snap_pizzicato_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Fingernails":
			if _technical, ok := stage.Technical_Fingernails_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Brass_bend":
			if _technical, ok := stage.Technical_Brass_bend_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Flip":
			if _technical, ok := stage.Technical_Flip_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Smear":
			if _technical, ok := stage.Technical_Smear_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Golpe":
			if _technical, ok := stage.Technical_Golpe_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Empty_placement_smufl) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Stopped":
			if _technical, ok := stage.Technical_Stopped_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Open":
			if _technical, ok := stage.Technical_Open_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Half_muted":
			if _technical, ok := stage.Technical_Half_muted_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Empty_print_object_style_align) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Empty_print_style) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Empty_print_style_align) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Empty_print_style_align_id) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Empty_trill_sound) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Trill_mark":
			if _ornaments, ok := stage.Ornaments_Trill_mark_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		case "Vertical_turn":
			if _ornaments, ok := stage.Ornaments_Vertical_turn_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		case "Inverted_vertical_turn":
			if _ornaments, ok := stage.Ornaments_Inverted_vertical_turn_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		case "Shake":
			if _ornaments, ok := stage.Ornaments_Shake_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		case "Haydn":
			if _ornaments, ok := stage.Ornaments_Haydn_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		}
	}
	return
}

func (inst *Encoding) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Ending) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Extend) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Feature) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Grouping":
		switch reverseField.Fieldname {
		case "Feature":
			if _grouping, ok := stage.Grouping_Feature_reverseMap[inst]; ok {
				res = _grouping.Name
			}
		}
	}
	return
}

func (inst *Fermata) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Fermata":
			if _notations, ok := stage.Notations_Fermata_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *Figure) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Figured_bass":
		switch reverseField.Fieldname {
		case "Figure":
			if _figured_bass, ok := stage.Figured_bass_Figure_reverseMap[inst]; ok {
				res = _figured_bass.Name
			}
		}
	}
	return
}

func (inst *Figured_bass) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Figured_bass":
			if _a_measure, ok := stage.A_measure_Figured_bass_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Figured_bass":
			if _a_part_1, ok := stage.A_part_1_Figured_bass_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	}
	return
}

func (inst *Fingering) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Fingering":
			if _technical, ok := stage.Technical_Fingering_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *First_fret) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *For_part) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "For_part":
			if _attributes, ok := stage.Attributes_For_part_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *Formatted_symbol) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Formatted_symbol_id) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Credit":
		switch reverseField.Fieldname {
		case "Credit_symbol":
			if _credit, ok := stage.Credit_Credit_symbol_reverseMap[inst]; ok {
				res = _credit.Name
			}
		}
	case "Direction_type":
		switch reverseField.Fieldname {
		case "Symbol":
			if _direction_type, ok := stage.Direction_type_Symbol_reverseMap[inst]; ok {
				res = _direction_type.Name
			}
		}
	}
	return
}

func (inst *Formatted_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Name_display":
		switch reverseField.Fieldname {
		case "Display_text":
			if _name_display, ok := stage.Name_display_Display_text_reverseMap[inst]; ok {
				res = _name_display.Name
			}
		}
	case "Notehead_text":
		switch reverseField.Fieldname {
		case "Display_text":
			if _notehead_text, ok := stage.Notehead_text_Display_text_reverseMap[inst]; ok {
				res = _notehead_text.Name
			}
		}
	}
	return
}

func (inst *Formatted_text_id) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Credit":
		switch reverseField.Fieldname {
		case "Credit_words":
			if _credit, ok := stage.Credit_Credit_words_reverseMap[inst]; ok {
				res = _credit.Name
			}
		}
	case "Direction_type":
		switch reverseField.Fieldname {
		case "Rehearsal":
			if _direction_type, ok := stage.Direction_type_Rehearsal_reverseMap[inst]; ok {
				res = _direction_type.Name
			}
		case "Words":
			if _direction_type, ok := stage.Direction_type_Words_reverseMap[inst]; ok {
				res = _direction_type.Name
			}
		}
	}
	return
}

func (inst *Forward) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Forward":
			if _a_measure, ok := stage.A_measure_Forward_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Forward":
			if _a_part_1, ok := stage.A_part_1_Forward_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	}
	return
}

func (inst *Frame) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Frame_note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Frame":
		switch reverseField.Fieldname {
		case "Frame_note":
			if _frame, ok := stage.Frame_Frame_note_reverseMap[inst]; ok {
				res = _frame.Name
			}
		}
	}
	return
}

func (inst *Fret) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Fret":
			if _technical, ok := stage.Technical_Fret_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Glass) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Glissando) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Glissando":
			if _notations, ok := stage.Notations_Glissando_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *Glyph) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Appearance":
		switch reverseField.Fieldname {
		case "Glyph":
			if _appearance, ok := stage.Appearance_Glyph_reverseMap[inst]; ok {
				res = _appearance.Name
			}
		}
	}
	return
}

func (inst *Grace) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Group_barline) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Group_name) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Group_symbol) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Grouping) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Grouping":
			if _a_measure, ok := stage.A_measure_Grouping_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Grouping":
			if _a_part_1, ok := stage.A_part_1_Grouping_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	}
	return
}

func (inst *Hammer_on_pull_off) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Hammer_on":
			if _technical, ok := stage.Technical_Hammer_on_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Pull_off":
			if _technical, ok := stage.Technical_Pull_off_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Handbell) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Handbell":
			if _technical, ok := stage.Technical_Handbell_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Harmon_closed) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Harmon_mute) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Harmon_mute":
			if _technical, ok := stage.Technical_Harmon_mute_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Harmonic) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Harmonic":
			if _technical, ok := stage.Technical_Harmonic_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Harmony) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Harmony":
			if _a_measure, ok := stage.A_measure_Harmony_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Harmony":
			if _a_part_1, ok := stage.A_part_1_Harmony_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	}
	return
}

func (inst *Harmony_alter) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Harp_pedals) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Heel_toe) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Heel":
			if _technical, ok := stage.Technical_Heel_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "Toe":
			if _technical, ok := stage.Technical_Toe_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Hole) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Hole":
			if _technical, ok := stage.Technical_Hole_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Hole_closed) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Horizontal_turn) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Turn":
			if _ornaments, ok := stage.Ornaments_Turn_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		case "Delayed_turn":
			if _ornaments, ok := stage.Ornaments_Delayed_turn_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		case "Inverted_turn":
			if _ornaments, ok := stage.Ornaments_Inverted_turn_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		case "Delayed_inverted_turn":
			if _ornaments, ok := stage.Ornaments_Delayed_inverted_turn_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		}
	}
	return
}

func (inst *Identification) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Image) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Instrument) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Note":
		switch reverseField.Fieldname {
		case "Instrument":
			if _note, ok := stage.Note_Instrument_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	}
	return
}

func (inst *Instrument_change) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Sound":
		switch reverseField.Fieldname {
		case "Instrument_change":
			if _sound, ok := stage.Sound_Instrument_change_reverseMap[inst]; ok {
				res = _sound.Name
			}
		}
	}
	return
}

func (inst *Instrument_link) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Part_link":
		switch reverseField.Fieldname {
		case "Instrument_link":
			if _part_link, ok := stage.Part_link_Instrument_link_reverseMap[inst]; ok {
				res = _part_link.Name
			}
		}
	}
	return
}

func (inst *Interchangeable) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Inversion) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Key) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Key":
			if _attributes, ok := stage.Attributes_Key_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *Key_accidental) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Key_octave) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Key":
		switch reverseField.Fieldname {
		case "Key_octave":
			if _key, ok := stage.Key_Key_octave_reverseMap[inst]; ok {
				res = _key.Name
			}
		}
	}
	return
}

func (inst *Kind) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Level) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Line_detail) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Staff_details":
		switch reverseField.Fieldname {
		case "Line_detail":
			if _staff_details, ok := stage.Staff_details_Line_detail_reverseMap[inst]; ok {
				res = _staff_details.Name
			}
		}
	}
	return
}

func (inst *Line_width) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Appearance":
		switch reverseField.Fieldname {
		case "Line_width":
			if _appearance, ok := stage.Appearance_Line_width_reverseMap[inst]; ok {
				res = _appearance.Name
			}
		}
	}
	return
}

func (inst *Link) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Link":
			if _a_measure, ok := stage.A_measure_Link_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Link":
			if _a_part_1, ok := stage.A_part_1_Link_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	case "Credit":
		switch reverseField.Fieldname {
		case "Link":
			if _credit, ok := stage.Credit_Link_reverseMap[inst]; ok {
				res = _credit.Name
			}
		}
	}
	return
}

func (inst *Listen) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Listening) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Listening":
			if _a_measure, ok := stage.A_measure_Listening_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Listening":
			if _a_part_1, ok := stage.A_part_1_Listening_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	}
	return
}

func (inst *Lyric) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Note":
		switch reverseField.Fieldname {
		case "Lyric":
			if _note, ok := stage.Note_Lyric_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	}
	return
}

func (inst *Lyric_font) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Defaults":
		switch reverseField.Fieldname {
		case "Lyric_font":
			if _defaults, ok := stage.Defaults_Lyric_font_reverseMap[inst]; ok {
				res = _defaults.Name
			}
		}
	}
	return
}

func (inst *Lyric_language) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Defaults":
		switch reverseField.Fieldname {
		case "Lyric_language":
			if _defaults, ok := stage.Defaults_Lyric_language_reverseMap[inst]; ok {
				res = _defaults.Name
			}
		}
	}
	return
}

func (inst *Measure_layout) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Measure_numbering) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Measure_repeat) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Measure_style) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Measure_style":
			if _attributes, ok := stage.Attributes_Measure_style_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *Membrane) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Metal) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Metronome) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Metronome_beam) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Metronome_note":
		switch reverseField.Fieldname {
		case "Metronome_beam":
			if _metronome_note, ok := stage.Metronome_note_Metronome_beam_reverseMap[inst]; ok {
				res = _metronome_note.Name
			}
		}
	}
	return
}

func (inst *Metronome_note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Metronome":
		switch reverseField.Fieldname {
		case "Metronome_note":
			if _metronome, ok := stage.Metronome_Metronome_note_reverseMap[inst]; ok {
				res = _metronome.Name
			}
		}
	}
	return
}

func (inst *Metronome_tied) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Metronome_tuplet) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Midi_device) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Score_part":
		switch reverseField.Fieldname {
		case "Midi_device":
			if _score_part, ok := stage.Score_part_Midi_device_reverseMap[inst]; ok {
				res = _score_part.Name
			}
		}
	case "Sound":
		switch reverseField.Fieldname {
		case "Midi_device":
			if _sound, ok := stage.Sound_Midi_device_reverseMap[inst]; ok {
				res = _sound.Name
			}
		}
	}
	return
}

func (inst *Midi_instrument) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Score_part":
		switch reverseField.Fieldname {
		case "Midi_instrument":
			if _score_part, ok := stage.Score_part_Midi_instrument_reverseMap[inst]; ok {
				res = _score_part.Name
			}
		}
	case "Sound":
		switch reverseField.Fieldname {
		case "Midi_instrument":
			if _sound, ok := stage.Sound_Midi_instrument_reverseMap[inst]; ok {
				res = _sound.Name
			}
		}
	}
	return
}

func (inst *Miscellaneous) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Miscellaneous_field) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Miscellaneous":
		switch reverseField.Fieldname {
		case "Miscellaneous_field":
			if _miscellaneous, ok := stage.Miscellaneous_Miscellaneous_field_reverseMap[inst]; ok {
				res = _miscellaneous.Name
			}
		}
	}
	return
}

func (inst *Mordent) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Mordent":
			if _ornaments, ok := stage.Ornaments_Mordent_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		case "Inverted_mordent":
			if _ornaments, ok := stage.Ornaments_Inverted_mordent_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		}
	}
	return
}

func (inst *Multiple_rest) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Name_display) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Non_arpeggiate) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Non_arpeggiate":
			if _notations, ok := stage.Notations_Non_arpeggiate_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *Notations) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Note":
		switch reverseField.Fieldname {
		case "Notations":
			if _note, ok := stage.Note_Notations_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	}
	return
}

func (inst *Note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Note":
			if _a_measure, ok := stage.A_measure_Note_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Note":
			if _a_part_1, ok := stage.A_part_1_Note_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	}
	return
}

func (inst *Note_size) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Appearance":
		switch reverseField.Fieldname {
		case "Note_size":
			if _appearance, ok := stage.Appearance_Note_size_reverseMap[inst]; ok {
				res = _appearance.Name
			}
		}
	}
	return
}

func (inst *Note_type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Notehead) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Notehead_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Numeral) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Numeral_key) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Numeral_root) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Octave_shift) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Offset) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Opus) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Ornaments) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Ornaments":
			if _notations, ok := stage.Notations_Ornaments_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *Other_appearance) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Appearance":
		switch reverseField.Fieldname {
		case "Other_appearance":
			if _appearance, ok := stage.Appearance_Other_appearance_reverseMap[inst]; ok {
				res = _appearance.Name
			}
		}
	}
	return
}

func (inst *Other_direction) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Other_listening) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Listen":
		switch reverseField.Fieldname {
		case "Other_listen":
			if _listen, ok := stage.Listen_Other_listen_reverseMap[inst]; ok {
				res = _listen.Name
			}
		}
	case "Listening":
		switch reverseField.Fieldname {
		case "Other_listening":
			if _listening, ok := stage.Listening_Other_listening_reverseMap[inst]; ok {
				res = _listening.Name
			}
		}
	}
	return
}

func (inst *Other_notation) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Other_notation":
			if _notations, ok := stage.Notations_Other_notation_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *Other_placement_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "Other_articulation":
			if _articulations, ok := stage.Articulations_Other_articulation_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		}
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Other_ornament":
			if _ornaments, ok := stage.Ornaments_Other_ornament_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		}
	case "Technical":
		switch reverseField.Fieldname {
		case "Other_technical":
			if _technical, ok := stage.Technical_Other_technical_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Other_play) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Play":
		switch reverseField.Fieldname {
		case "Other_play":
			if _play, ok := stage.Play_Other_play_reverseMap[inst]; ok {
				res = _play.Name
			}
		}
	}
	return
}

func (inst *Other_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Dynamics":
		switch reverseField.Fieldname {
		case "Other_dynamics":
			if _dynamics, ok := stage.Dynamics_Other_dynamics_reverseMap[inst]; ok {
				res = _dynamics.Name
			}
		}
	}
	return
}

func (inst *Page_layout) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Page_margins) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Part_clef) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Part_group) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Part_link) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Score_part":
		switch reverseField.Fieldname {
		case "Part_link":
			if _score_part, ok := stage.Score_part_Part_link_reverseMap[inst]; ok {
				res = _score_part.Name
			}
		}
	}
	return
}

func (inst *Part_list) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Part_name) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Part_symbol) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Part_transpose) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Pedal) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Pedal_tuning) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Harp_pedals":
		switch reverseField.Fieldname {
		case "Pedal_tuning":
			if _harp_pedals, ok := stage.Harp_pedals_Pedal_tuning_reverseMap[inst]; ok {
				res = _harp_pedals.Name
			}
		}
	}
	return
}

func (inst *Per_minute) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Percussion) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Direction_type":
		switch reverseField.Fieldname {
		case "Percussion":
			if _direction_type, ok := stage.Direction_type_Percussion_reverseMap[inst]; ok {
				res = _direction_type.Name
			}
		}
	}
	return
}

func (inst *Pitch) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Pitched) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Placement_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Pluck":
			if _technical, ok := stage.Technical_Pluck_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Play) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Sound":
		switch reverseField.Fieldname {
		case "Play":
			if _sound, ok := stage.Sound_Play_reverseMap[inst]; ok {
				res = _sound.Name
			}
		}
	}
	return
}

func (inst *Player) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Score_part":
		switch reverseField.Fieldname {
		case "Player":
			if _score_part, ok := stage.Score_part_Player_reverseMap[inst]; ok {
				res = _score_part.Name
			}
		}
	}
	return
}

func (inst *Principal_voice) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Print) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Print":
			if _a_measure, ok := stage.A_measure_Print_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Print":
			if _a_part_1, ok := stage.A_part_1_Print_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	}
	return
}

func (inst *Release) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Repeat) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Rest) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Root) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Root_step) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Scaling) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Scordatura) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Score_instrument) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Score_part":
		switch reverseField.Fieldname {
		case "Score_instrument":
			if _score_part, ok := stage.Score_part_Score_instrument_reverseMap[inst]; ok {
				res = _score_part.Name
			}
		}
	}
	return
}

func (inst *Score_part) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Score_partwise) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Score_timewise) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Segno) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Direction_type":
		switch reverseField.Fieldname {
		case "Segno":
			if _direction_type, ok := stage.Direction_type_Segno_reverseMap[inst]; ok {
				res = _direction_type.Name
			}
		}
	}
	return
}

func (inst *Slash) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Slide) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Slide":
			if _notations, ok := stage.Notations_Slide_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *Slur) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Slur":
			if _notations, ok := stage.Notations_Slur_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *Sound) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Sound":
			if _a_measure, ok := stage.A_measure_Sound_reverseMap[inst]; ok {
				res = _a_measure.Name
			}
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Sound":
			if _a_part_1, ok := stage.A_part_1_Sound_reverseMap[inst]; ok {
				res = _a_part_1.Name
			}
		}
	}
	return
}

func (inst *Staff_details) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Staff_details":
			if _attributes, ok := stage.Attributes_Staff_details_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *Staff_divide) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Staff_layout) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Defaults":
		switch reverseField.Fieldname {
		case "Staff_layout":
			if _defaults, ok := stage.Defaults_Staff_layout_reverseMap[inst]; ok {
				res = _defaults.Name
			}
		}
	case "Print":
		switch reverseField.Fieldname {
		case "Staff_layout":
			if _print, ok := stage.Print_Staff_layout_reverseMap[inst]; ok {
				res = _print.Name
			}
		}
	}
	return
}

func (inst *Staff_size) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Staff_tuning) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Staff_details":
		switch reverseField.Fieldname {
		case "Staff_tuning":
			if _staff_details, ok := stage.Staff_details_Staff_tuning_reverseMap[inst]; ok {
				res = _staff_details.Name
			}
		}
	}
	return
}

func (inst *Stem) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Stick) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *String_mute) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *String_type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "String":
			if _technical, ok := stage.Technical_String_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Strong_accent) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "Strong_accent":
			if _articulations, ok := stage.Articulations_Strong_accent_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		}
	}
	return
}

func (inst *Style_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Supports) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Encoding":
		switch reverseField.Fieldname {
		case "Supports":
			if _encoding, ok := stage.Encoding_Supports_reverseMap[inst]; ok {
				res = _encoding.Name
			}
		}
	}
	return
}

func (inst *Swing) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Sync) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Listening":
		switch reverseField.Fieldname {
		case "Sync":
			if _listening, ok := stage.Listening_Sync_reverseMap[inst]; ok {
				res = _listening.Name
			}
		}
	}
	return
}

func (inst *System_dividers) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *System_layout) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *System_margins) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Tap) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Tap":
			if _technical, ok := stage.Technical_Tap_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Technical) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Technical":
			if _notations, ok := stage.Notations_Technical_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *Text_element_data) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Lyric":
		switch reverseField.Fieldname {
		case "Text":
			if _lyric, ok := stage.Lyric_Text_reverseMap[inst]; ok {
				res = _lyric.Name
			}
		}
	}
	return
}

func (inst *Tie) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Tied) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Tied":
			if _notations, ok := stage.Notations_Tied_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *Time) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Time":
			if _attributes, ok := stage.Attributes_Time_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *Time_modification) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Timpani) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Transpose) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Transpose":
			if _attributes, ok := stage.Attributes_Transpose_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *Tremolo) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Tremolo":
			if _ornaments, ok := stage.Ornaments_Tremolo_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		}
	}
	return
}

func (inst *Tuplet) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Tuplet":
			if _notations, ok := stage.Notations_Tuplet_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *Tuplet_dot) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Tuplet_portion":
		switch reverseField.Fieldname {
		case "Tuplet_dot":
			if _tuplet_portion, ok := stage.Tuplet_portion_Tuplet_dot_reverseMap[inst]; ok {
				res = _tuplet_portion.Name
			}
		}
	}
	return
}

func (inst *Tuplet_number) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Tuplet_portion) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Tuplet_type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Typed_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Encoding":
		switch reverseField.Fieldname {
		case "Encoder":
			if _encoding, ok := stage.Encoding_Encoder_reverseMap[inst]; ok {
				res = _encoding.Name
			}
		}
	case "Identification":
		switch reverseField.Fieldname {
		case "Creator":
			if _identification, ok := stage.Identification_Creator_reverseMap[inst]; ok {
				res = _identification.Name
			}
		case "Rights":
			if _identification, ok := stage.Identification_Rights_reverseMap[inst]; ok {
				res = _identification.Name
			}
		case "Relation":
			if _identification, ok := stage.Identification_Relation_reverseMap[inst]; ok {
				res = _identification.Name
			}
		}
	}
	return
}

func (inst *Unpitched) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Virtual_instrument) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Wait) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Listen":
		switch reverseField.Fieldname {
		case "Wait":
			if _listen, ok := stage.Listen_Wait_reverseMap[inst]; ok {
				res = _listen.Name
			}
		}
	}
	return
}

func (inst *Wavy_line) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Wavy_line":
			if _ornaments, ok := stage.Ornaments_Wavy_line_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		}
	}
	return
}

func (inst *Wedge) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Wood) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Work) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}
