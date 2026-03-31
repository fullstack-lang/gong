// generated code - do not edit
package models

// insertion point
func (inst *A_directive) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *A_measure) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *A_measure_1) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *A_part) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *A_part_1) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Accidental) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Accidental_mark) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Accidental_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Accord) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Accordion_registration) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Appearance) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Arpeggiate) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Arrow) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Articulations) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Assess) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Attributes) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Backup) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Bar_style_color) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Barline) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Barre) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Bass) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Bass_step) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Beam) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Beat_repeat) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Beat_unit_tied) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Beater) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Bend) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Bookmark) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Bracket) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Breath_mark) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Caesura) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Cancel) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Clef) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Coda) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Credit) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Dashes) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Defaults) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Degree) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Degree_alter) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Degree_type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Degree_value) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Direction) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Direction_type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Distance) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Double) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Dynamics) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Effect) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Elision) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Empty) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Empty_font) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Empty_line) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Empty_placement) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Empty_placement_smufl) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Empty_print_object_style_align) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Empty_print_style) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Empty_print_style_align) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Empty_print_style_align_id) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Empty_trill_sound) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Encoding) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Ending) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Extend) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Feature) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Fermata) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Figure) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Figured_bass) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Fingering) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *First_fret) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *For_part) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Formatted_symbol) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Formatted_symbol_id) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Formatted_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Formatted_text_id) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Forward) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Frame) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Frame_note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Fret) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Glass) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Glissando) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Glyph) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Grace) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Group_barline) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Group_name) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Group_symbol) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Grouping) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Hammer_on_pull_off) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Handbell) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Harmon_closed) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Harmon_mute) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Harmonic) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Harmony) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Harmony_alter) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Harp_pedals) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Heel_toe) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Hole) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Hole_closed) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Horizontal_turn) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Identification) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Image) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Instrument) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Instrument_change) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Instrument_link) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Interchangeable) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Inversion) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Key) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Key_accidental) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Key_octave) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Kind) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Level) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Line_detail) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Line_width) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Link) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Listen) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Listening) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Lyric) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Lyric_font) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Lyric_language) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Measure_layout) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Measure_numbering) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Measure_repeat) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Measure_style) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Membrane) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Metal) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Metronome) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Metronome_beam) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Metronome_note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Metronome_tied) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Metronome_tuplet) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Midi_device) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Midi_instrument) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Miscellaneous) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Miscellaneous_field) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Mordent) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Multiple_rest) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Name_display) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Non_arpeggiate) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Notations) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Note_size) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Note_type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Notehead) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Notehead_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Numeral) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Numeral_key) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Numeral_root) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Octave_shift) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Offset) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Opus) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Ornaments) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Other_appearance) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Other_direction) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Other_listening) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Other_notation) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Other_placement_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Other_play) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Other_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Page_layout) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Page_margins) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Part_clef) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Part_group) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Part_link) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Part_list) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Part_name) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Part_symbol) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Part_transpose) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Pedal) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Pedal_tuning) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Per_minute) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Percussion) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Pitch) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Pitched) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Placement_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Play) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Player) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Principal_voice) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Print) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Release) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Repeat) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Rest) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Root) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Root_step) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Scaling) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Scordatura) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Score_instrument) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Score_part) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Score_partwise) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Score_timewise) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Segno) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Slash) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Slide) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Slur) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Sound) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Staff_details) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Staff_divide) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Staff_layout) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Staff_size) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Staff_tuning) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Stem) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Stick) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *String_mute) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *String_type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Strong_accent) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Style_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Supports) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Swing) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Sync) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *System_dividers) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *System_layout) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *System_margins) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Tap) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Technical) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Text_element_data) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Tie) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Tied) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Time) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Time_modification) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Timpani) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Transpose) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Tremolo) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Tuplet) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Tuplet_dot) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Tuplet_number) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Tuplet_portion) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Tuplet_type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Typed_text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Unpitched) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Virtual_instrument) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Wait) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Wavy_line) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Wedge) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Wood) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Work) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

// insertion point
func (inst *A_directive) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Directive":
			res = stage.Attributes_Directive_reverseMap[inst]
		}
	}
	return res
}

func (inst *A_measure) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_part":
		switch reverseField.Fieldname {
		case "Measure":
			res = stage.A_part_Measure_reverseMap[inst]
		}
	}
	return res
}

func (inst *A_measure_1) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Score_timewise":
		switch reverseField.Fieldname {
		case "Measure":
			res = stage.Score_timewise_Measure_reverseMap[inst]
		}
	}
	return res
}

func (inst *A_part) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Score_partwise":
		switch reverseField.Fieldname {
		case "Part":
			res = stage.Score_partwise_Part_reverseMap[inst]
		}
	}
	return res
}

func (inst *A_part_1) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure_1":
		switch reverseField.Fieldname {
		case "Part":
			res = stage.A_measure_1_Part_reverseMap[inst]
		}
	}
	return res
}

func (inst *Accidental) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Accidental_mark) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Accidental_mark":
			res = stage.Notations_Accidental_mark_reverseMap[inst]
		}
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Accidental_mark":
			res = stage.Ornaments_Accidental_mark_reverseMap[inst]
		}
	}
	return res
}

func (inst *Accidental_text) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Name_display":
		switch reverseField.Fieldname {
		case "Accidental_text":
			res = stage.Name_display_Accidental_text_reverseMap[inst]
		}
	case "Notehead_text":
		switch reverseField.Fieldname {
		case "Accidental_text":
			res = stage.Notehead_text_Accidental_text_reverseMap[inst]
		}
	}
	return res
}

func (inst *Accord) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Scordatura":
		switch reverseField.Fieldname {
		case "Accord":
			res = stage.Scordatura_Accord_reverseMap[inst]
		}
	}
	return res
}

func (inst *Accordion_registration) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Appearance) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Arpeggiate) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Arpeggiate":
			res = stage.Notations_Arpeggiate_reverseMap[inst]
		}
	}
	return res
}

func (inst *Arrow) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Arrow":
			res = stage.Technical_Arrow_reverseMap[inst]
		}
	}
	return res
}

func (inst *Articulations) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Articulations":
			res = stage.Notations_Articulations_reverseMap[inst]
		}
	}
	return res
}

func (inst *Assess) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Listen":
		switch reverseField.Fieldname {
		case "Assess":
			res = stage.Listen_Assess_reverseMap[inst]
		}
	}
	return res
}

func (inst *Attributes) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Attributes":
			res = stage.A_measure_Attributes_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Attributes":
			res = stage.A_part_1_Attributes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Backup) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Backup":
			res = stage.A_measure_Backup_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Backup":
			res = stage.A_part_1_Backup_reverseMap[inst]
		}
	}
	return res
}

func (inst *Bar_style_color) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Barline) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Barline":
			res = stage.A_measure_Barline_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Barline":
			res = stage.A_part_1_Barline_reverseMap[inst]
		}
	}
	return res
}

func (inst *Barre) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Bass) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Bass_step) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Beam) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Beat_repeat) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Beat_unit_tied) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Metronome":
		switch reverseField.Fieldname {
		case "Beat_unit_tied":
			res = stage.Metronome_Beat_unit_tied_reverseMap[inst]
		}
	}
	return res
}

func (inst *Beater) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Bend) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Bend":
			res = stage.Technical_Bend_reverseMap[inst]
		}
	}
	return res
}

func (inst *Bookmark) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Bookmark":
			res = stage.A_measure_Bookmark_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Bookmark":
			res = stage.A_part_1_Bookmark_reverseMap[inst]
		}
	case "Credit":
		switch reverseField.Fieldname {
		case "Bookmark":
			res = stage.Credit_Bookmark_reverseMap[inst]
		}
	}
	return res
}

func (inst *Bracket) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Breath_mark) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "Breath_mark":
			res = stage.Articulations_Breath_mark_reverseMap[inst]
		}
	}
	return res
}

func (inst *Caesura) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "Caesura":
			res = stage.Articulations_Caesura_reverseMap[inst]
		}
	}
	return res
}

func (inst *Cancel) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Clef) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Clef":
			res = stage.Attributes_Clef_reverseMap[inst]
		}
	}
	return res
}

func (inst *Coda) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Direction_type":
		switch reverseField.Fieldname {
		case "Coda":
			res = stage.Direction_type_Coda_reverseMap[inst]
		}
	}
	return res
}

func (inst *Credit) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Score_partwise":
		switch reverseField.Fieldname {
		case "Credit":
			res = stage.Score_partwise_Credit_reverseMap[inst]
		}
	case "Score_timewise":
		switch reverseField.Fieldname {
		case "Credit":
			res = stage.Score_timewise_Credit_reverseMap[inst]
		}
	}
	return res
}

func (inst *Dashes) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Defaults) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Degree) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Harmony":
		switch reverseField.Fieldname {
		case "Degree":
			res = stage.Harmony_Degree_reverseMap[inst]
		}
	}
	return res
}

func (inst *Degree_alter) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Degree_type) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Degree_value) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Direction) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Direction":
			res = stage.A_measure_Direction_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Direction":
			res = stage.A_part_1_Direction_reverseMap[inst]
		}
	}
	return res
}

func (inst *Direction_type) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Direction":
		switch reverseField.Fieldname {
		case "Direction_type":
			res = stage.Direction_Direction_type_reverseMap[inst]
		}
	}
	return res
}

func (inst *Distance) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Appearance":
		switch reverseField.Fieldname {
		case "Distance":
			res = stage.Appearance_Distance_reverseMap[inst]
		}
	}
	return res
}

func (inst *Double) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Dynamics) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Direction_type":
		switch reverseField.Fieldname {
		case "Dynamics":
			res = stage.Direction_type_Dynamics_reverseMap[inst]
		}
	case "Notations":
		switch reverseField.Fieldname {
		case "Dynamics":
			res = stage.Notations_Dynamics_reverseMap[inst]
		}
	}
	return res
}

func (inst *Effect) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Elision) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Lyric":
		switch reverseField.Fieldname {
		case "Elision":
			res = stage.Lyric_Elision_reverseMap[inst]
		}
	}
	return res
}

func (inst *Empty) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Empty_font) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Empty_line) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "Scoop":
			res = stage.Articulations_Scoop_reverseMap[inst]
		case "Plop":
			res = stage.Articulations_Plop_reverseMap[inst]
		case "Doit":
			res = stage.Articulations_Doit_reverseMap[inst]
		case "Falloff":
			res = stage.Articulations_Falloff_reverseMap[inst]
		}
	}
	return res
}

func (inst *Empty_placement) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "Accent":
			res = stage.Articulations_Accent_reverseMap[inst]
		case "Staccato":
			res = stage.Articulations_Staccato_reverseMap[inst]
		case "Tenuto":
			res = stage.Articulations_Tenuto_reverseMap[inst]
		case "Detached_legato":
			res = stage.Articulations_Detached_legato_reverseMap[inst]
		case "Staccatissimo":
			res = stage.Articulations_Staccatissimo_reverseMap[inst]
		case "Spiccato":
			res = stage.Articulations_Spiccato_reverseMap[inst]
		case "Stress":
			res = stage.Articulations_Stress_reverseMap[inst]
		case "Unstress":
			res = stage.Articulations_Unstress_reverseMap[inst]
		case "Soft_accent":
			res = stage.Articulations_Soft_accent_reverseMap[inst]
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Dot":
			res = stage.Note_Dot_reverseMap[inst]
		}
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Schleifer":
			res = stage.Ornaments_Schleifer_reverseMap[inst]
		}
	case "Technical":
		switch reverseField.Fieldname {
		case "Up_bow":
			res = stage.Technical_Up_bow_reverseMap[inst]
		case "Down_bow":
			res = stage.Technical_Down_bow_reverseMap[inst]
		case "Open_string":
			res = stage.Technical_Open_string_reverseMap[inst]
		case "Thumb_position":
			res = stage.Technical_Thumb_position_reverseMap[inst]
		case "Double_tongue":
			res = stage.Technical_Double_tongue_reverseMap[inst]
		case "Triple_tongue":
			res = stage.Technical_Triple_tongue_reverseMap[inst]
		case "Snap_pizzicato":
			res = stage.Technical_Snap_pizzicato_reverseMap[inst]
		case "Fingernails":
			res = stage.Technical_Fingernails_reverseMap[inst]
		case "Brass_bend":
			res = stage.Technical_Brass_bend_reverseMap[inst]
		case "Flip":
			res = stage.Technical_Flip_reverseMap[inst]
		case "Smear":
			res = stage.Technical_Smear_reverseMap[inst]
		case "Golpe":
			res = stage.Technical_Golpe_reverseMap[inst]
		}
	}
	return res
}

func (inst *Empty_placement_smufl) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Stopped":
			res = stage.Technical_Stopped_reverseMap[inst]
		case "Open":
			res = stage.Technical_Open_reverseMap[inst]
		case "Half_muted":
			res = stage.Technical_Half_muted_reverseMap[inst]
		}
	}
	return res
}

func (inst *Empty_print_object_style_align) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Empty_print_style) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Empty_print_style_align) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Empty_print_style_align_id) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Empty_trill_sound) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Trill_mark":
			res = stage.Ornaments_Trill_mark_reverseMap[inst]
		case "Vertical_turn":
			res = stage.Ornaments_Vertical_turn_reverseMap[inst]
		case "Inverted_vertical_turn":
			res = stage.Ornaments_Inverted_vertical_turn_reverseMap[inst]
		case "Shake":
			res = stage.Ornaments_Shake_reverseMap[inst]
		case "Haydn":
			res = stage.Ornaments_Haydn_reverseMap[inst]
		}
	}
	return res
}

func (inst *Encoding) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Ending) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Extend) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Feature) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Grouping":
		switch reverseField.Fieldname {
		case "Feature":
			res = stage.Grouping_Feature_reverseMap[inst]
		}
	}
	return res
}

func (inst *Fermata) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Fermata":
			res = stage.Notations_Fermata_reverseMap[inst]
		}
	}
	return res
}

func (inst *Figure) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Figured_bass":
		switch reverseField.Fieldname {
		case "Figure":
			res = stage.Figured_bass_Figure_reverseMap[inst]
		}
	}
	return res
}

func (inst *Figured_bass) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Figured_bass":
			res = stage.A_measure_Figured_bass_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Figured_bass":
			res = stage.A_part_1_Figured_bass_reverseMap[inst]
		}
	}
	return res
}

func (inst *Fingering) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Fingering":
			res = stage.Technical_Fingering_reverseMap[inst]
		}
	}
	return res
}

func (inst *First_fret) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *For_part) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "For_part":
			res = stage.Attributes_For_part_reverseMap[inst]
		}
	}
	return res
}

func (inst *Formatted_symbol) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Formatted_symbol_id) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Credit":
		switch reverseField.Fieldname {
		case "Credit_symbol":
			res = stage.Credit_Credit_symbol_reverseMap[inst]
		}
	case "Direction_type":
		switch reverseField.Fieldname {
		case "Symbol":
			res = stage.Direction_type_Symbol_reverseMap[inst]
		}
	}
	return res
}

func (inst *Formatted_text) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Name_display":
		switch reverseField.Fieldname {
		case "Display_text":
			res = stage.Name_display_Display_text_reverseMap[inst]
		}
	case "Notehead_text":
		switch reverseField.Fieldname {
		case "Display_text":
			res = stage.Notehead_text_Display_text_reverseMap[inst]
		}
	}
	return res
}

func (inst *Formatted_text_id) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Credit":
		switch reverseField.Fieldname {
		case "Credit_words":
			res = stage.Credit_Credit_words_reverseMap[inst]
		}
	case "Direction_type":
		switch reverseField.Fieldname {
		case "Rehearsal":
			res = stage.Direction_type_Rehearsal_reverseMap[inst]
		case "Words":
			res = stage.Direction_type_Words_reverseMap[inst]
		}
	}
	return res
}

func (inst *Forward) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Forward":
			res = stage.A_measure_Forward_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Forward":
			res = stage.A_part_1_Forward_reverseMap[inst]
		}
	}
	return res
}

func (inst *Frame) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Frame_note) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Frame":
		switch reverseField.Fieldname {
		case "Frame_note":
			res = stage.Frame_Frame_note_reverseMap[inst]
		}
	}
	return res
}

func (inst *Fret) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Fret":
			res = stage.Technical_Fret_reverseMap[inst]
		}
	}
	return res
}

func (inst *Glass) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Glissando) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Glissando":
			res = stage.Notations_Glissando_reverseMap[inst]
		}
	}
	return res
}

func (inst *Glyph) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Appearance":
		switch reverseField.Fieldname {
		case "Glyph":
			res = stage.Appearance_Glyph_reverseMap[inst]
		}
	}
	return res
}

func (inst *Grace) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Group_barline) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Group_name) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Group_symbol) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Grouping) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Grouping":
			res = stage.A_measure_Grouping_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Grouping":
			res = stage.A_part_1_Grouping_reverseMap[inst]
		}
	}
	return res
}

func (inst *Hammer_on_pull_off) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Hammer_on":
			res = stage.Technical_Hammer_on_reverseMap[inst]
		case "Pull_off":
			res = stage.Technical_Pull_off_reverseMap[inst]
		}
	}
	return res
}

func (inst *Handbell) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Handbell":
			res = stage.Technical_Handbell_reverseMap[inst]
		}
	}
	return res
}

func (inst *Harmon_closed) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Harmon_mute) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Harmon_mute":
			res = stage.Technical_Harmon_mute_reverseMap[inst]
		}
	}
	return res
}

func (inst *Harmonic) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Harmonic":
			res = stage.Technical_Harmonic_reverseMap[inst]
		}
	}
	return res
}

func (inst *Harmony) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Harmony":
			res = stage.A_measure_Harmony_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Harmony":
			res = stage.A_part_1_Harmony_reverseMap[inst]
		}
	}
	return res
}

func (inst *Harmony_alter) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Harp_pedals) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Heel_toe) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Heel":
			res = stage.Technical_Heel_reverseMap[inst]
		case "Toe":
			res = stage.Technical_Toe_reverseMap[inst]
		}
	}
	return res
}

func (inst *Hole) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Hole":
			res = stage.Technical_Hole_reverseMap[inst]
		}
	}
	return res
}

func (inst *Hole_closed) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Horizontal_turn) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Turn":
			res = stage.Ornaments_Turn_reverseMap[inst]
		case "Delayed_turn":
			res = stage.Ornaments_Delayed_turn_reverseMap[inst]
		case "Inverted_turn":
			res = stage.Ornaments_Inverted_turn_reverseMap[inst]
		case "Delayed_inverted_turn":
			res = stage.Ornaments_Delayed_inverted_turn_reverseMap[inst]
		}
	}
	return res
}

func (inst *Identification) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Image) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Instrument) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Note":
		switch reverseField.Fieldname {
		case "Instrument":
			res = stage.Note_Instrument_reverseMap[inst]
		}
	}
	return res
}

func (inst *Instrument_change) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Sound":
		switch reverseField.Fieldname {
		case "Instrument_change":
			res = stage.Sound_Instrument_change_reverseMap[inst]
		}
	}
	return res
}

func (inst *Instrument_link) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Part_link":
		switch reverseField.Fieldname {
		case "Instrument_link":
			res = stage.Part_link_Instrument_link_reverseMap[inst]
		}
	}
	return res
}

func (inst *Interchangeable) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Inversion) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Key) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Key":
			res = stage.Attributes_Key_reverseMap[inst]
		}
	}
	return res
}

func (inst *Key_accidental) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Key_octave) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Key":
		switch reverseField.Fieldname {
		case "Key_octave":
			res = stage.Key_Key_octave_reverseMap[inst]
		}
	}
	return res
}

func (inst *Kind) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Level) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Line_detail) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Staff_details":
		switch reverseField.Fieldname {
		case "Line_detail":
			res = stage.Staff_details_Line_detail_reverseMap[inst]
		}
	}
	return res
}

func (inst *Line_width) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Appearance":
		switch reverseField.Fieldname {
		case "Line_width":
			res = stage.Appearance_Line_width_reverseMap[inst]
		}
	}
	return res
}

func (inst *Link) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Link":
			res = stage.A_measure_Link_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Link":
			res = stage.A_part_1_Link_reverseMap[inst]
		}
	case "Credit":
		switch reverseField.Fieldname {
		case "Link":
			res = stage.Credit_Link_reverseMap[inst]
		}
	}
	return res
}

func (inst *Listen) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Listening) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Listening":
			res = stage.A_measure_Listening_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Listening":
			res = stage.A_part_1_Listening_reverseMap[inst]
		}
	}
	return res
}

func (inst *Lyric) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Note":
		switch reverseField.Fieldname {
		case "Lyric":
			res = stage.Note_Lyric_reverseMap[inst]
		}
	}
	return res
}

func (inst *Lyric_font) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Defaults":
		switch reverseField.Fieldname {
		case "Lyric_font":
			res = stage.Defaults_Lyric_font_reverseMap[inst]
		}
	}
	return res
}

func (inst *Lyric_language) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Defaults":
		switch reverseField.Fieldname {
		case "Lyric_language":
			res = stage.Defaults_Lyric_language_reverseMap[inst]
		}
	}
	return res
}

func (inst *Measure_layout) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Measure_numbering) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Measure_repeat) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Measure_style) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Measure_style":
			res = stage.Attributes_Measure_style_reverseMap[inst]
		}
	}
	return res
}

func (inst *Membrane) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Metal) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Metronome) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Metronome_beam) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Metronome_note":
		switch reverseField.Fieldname {
		case "Metronome_beam":
			res = stage.Metronome_note_Metronome_beam_reverseMap[inst]
		}
	}
	return res
}

func (inst *Metronome_note) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Metronome":
		switch reverseField.Fieldname {
		case "Metronome_note":
			res = stage.Metronome_Metronome_note_reverseMap[inst]
		}
	}
	return res
}

func (inst *Metronome_tied) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Metronome_tuplet) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Midi_device) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Score_part":
		switch reverseField.Fieldname {
		case "Midi_device":
			res = stage.Score_part_Midi_device_reverseMap[inst]
		}
	case "Sound":
		switch reverseField.Fieldname {
		case "Midi_device":
			res = stage.Sound_Midi_device_reverseMap[inst]
		}
	}
	return res
}

func (inst *Midi_instrument) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Score_part":
		switch reverseField.Fieldname {
		case "Midi_instrument":
			res = stage.Score_part_Midi_instrument_reverseMap[inst]
		}
	case "Sound":
		switch reverseField.Fieldname {
		case "Midi_instrument":
			res = stage.Sound_Midi_instrument_reverseMap[inst]
		}
	}
	return res
}

func (inst *Miscellaneous) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Miscellaneous_field) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Miscellaneous":
		switch reverseField.Fieldname {
		case "Miscellaneous_field":
			res = stage.Miscellaneous_Miscellaneous_field_reverseMap[inst]
		}
	}
	return res
}

func (inst *Mordent) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Mordent":
			res = stage.Ornaments_Mordent_reverseMap[inst]
		case "Inverted_mordent":
			res = stage.Ornaments_Inverted_mordent_reverseMap[inst]
		}
	}
	return res
}

func (inst *Multiple_rest) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Name_display) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Non_arpeggiate) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Non_arpeggiate":
			res = stage.Notations_Non_arpeggiate_reverseMap[inst]
		}
	}
	return res
}

func (inst *Notations) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Note":
		switch reverseField.Fieldname {
		case "Notations":
			res = stage.Note_Notations_reverseMap[inst]
		}
	}
	return res
}

func (inst *Note) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Note":
			res = stage.A_measure_Note_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Note":
			res = stage.A_part_1_Note_reverseMap[inst]
		}
	}
	return res
}

func (inst *Note_size) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Appearance":
		switch reverseField.Fieldname {
		case "Note_size":
			res = stage.Appearance_Note_size_reverseMap[inst]
		}
	}
	return res
}

func (inst *Note_type) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Notehead) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Notehead_text) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Numeral) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Numeral_key) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Numeral_root) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Octave_shift) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Offset) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Opus) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Ornaments) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Ornaments":
			res = stage.Notations_Ornaments_reverseMap[inst]
		}
	}
	return res
}

func (inst *Other_appearance) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Appearance":
		switch reverseField.Fieldname {
		case "Other_appearance":
			res = stage.Appearance_Other_appearance_reverseMap[inst]
		}
	}
	return res
}

func (inst *Other_direction) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Other_listening) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Listen":
		switch reverseField.Fieldname {
		case "Other_listen":
			res = stage.Listen_Other_listen_reverseMap[inst]
		}
	case "Listening":
		switch reverseField.Fieldname {
		case "Other_listening":
			res = stage.Listening_Other_listening_reverseMap[inst]
		}
	}
	return res
}

func (inst *Other_notation) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Other_notation":
			res = stage.Notations_Other_notation_reverseMap[inst]
		}
	}
	return res
}

func (inst *Other_placement_text) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "Other_articulation":
			res = stage.Articulations_Other_articulation_reverseMap[inst]
		}
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Other_ornament":
			res = stage.Ornaments_Other_ornament_reverseMap[inst]
		}
	case "Technical":
		switch reverseField.Fieldname {
		case "Other_technical":
			res = stage.Technical_Other_technical_reverseMap[inst]
		}
	}
	return res
}

func (inst *Other_play) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Play":
		switch reverseField.Fieldname {
		case "Other_play":
			res = stage.Play_Other_play_reverseMap[inst]
		}
	}
	return res
}

func (inst *Other_text) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Dynamics":
		switch reverseField.Fieldname {
		case "Other_dynamics":
			res = stage.Dynamics_Other_dynamics_reverseMap[inst]
		}
	}
	return res
}

func (inst *Page_layout) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Page_margins) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Part_clef) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Part_group) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Part_link) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Score_part":
		switch reverseField.Fieldname {
		case "Part_link":
			res = stage.Score_part_Part_link_reverseMap[inst]
		}
	}
	return res
}

func (inst *Part_list) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Part_name) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Part_symbol) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Part_transpose) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Pedal) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Pedal_tuning) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Harp_pedals":
		switch reverseField.Fieldname {
		case "Pedal_tuning":
			res = stage.Harp_pedals_Pedal_tuning_reverseMap[inst]
		}
	}
	return res
}

func (inst *Per_minute) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Percussion) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Direction_type":
		switch reverseField.Fieldname {
		case "Percussion":
			res = stage.Direction_type_Percussion_reverseMap[inst]
		}
	}
	return res
}

func (inst *Pitch) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Pitched) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Placement_text) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Pluck":
			res = stage.Technical_Pluck_reverseMap[inst]
		}
	}
	return res
}

func (inst *Play) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Sound":
		switch reverseField.Fieldname {
		case "Play":
			res = stage.Sound_Play_reverseMap[inst]
		}
	}
	return res
}

func (inst *Player) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Score_part":
		switch reverseField.Fieldname {
		case "Player":
			res = stage.Score_part_Player_reverseMap[inst]
		}
	}
	return res
}

func (inst *Principal_voice) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Print) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Print":
			res = stage.A_measure_Print_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Print":
			res = stage.A_part_1_Print_reverseMap[inst]
		}
	}
	return res
}

func (inst *Release) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Repeat) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Rest) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Root) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Root_step) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Scaling) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Scordatura) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Score_instrument) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Score_part":
		switch reverseField.Fieldname {
		case "Score_instrument":
			res = stage.Score_part_Score_instrument_reverseMap[inst]
		}
	}
	return res
}

func (inst *Score_part) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Score_partwise) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Score_timewise) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Segno) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Direction_type":
		switch reverseField.Fieldname {
		case "Segno":
			res = stage.Direction_type_Segno_reverseMap[inst]
		}
	}
	return res
}

func (inst *Slash) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Slide) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Slide":
			res = stage.Notations_Slide_reverseMap[inst]
		}
	}
	return res
}

func (inst *Slur) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Slur":
			res = stage.Notations_Slur_reverseMap[inst]
		}
	}
	return res
}

func (inst *Sound) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_measure":
		switch reverseField.Fieldname {
		case "Sound":
			res = stage.A_measure_Sound_reverseMap[inst]
		}
	case "A_part_1":
		switch reverseField.Fieldname {
		case "Sound":
			res = stage.A_part_1_Sound_reverseMap[inst]
		}
	}
	return res
}

func (inst *Staff_details) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Staff_details":
			res = stage.Attributes_Staff_details_reverseMap[inst]
		}
	}
	return res
}

func (inst *Staff_divide) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Staff_layout) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Defaults":
		switch reverseField.Fieldname {
		case "Staff_layout":
			res = stage.Defaults_Staff_layout_reverseMap[inst]
		}
	case "Print":
		switch reverseField.Fieldname {
		case "Staff_layout":
			res = stage.Print_Staff_layout_reverseMap[inst]
		}
	}
	return res
}

func (inst *Staff_size) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Staff_tuning) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Staff_details":
		switch reverseField.Fieldname {
		case "Staff_tuning":
			res = stage.Staff_details_Staff_tuning_reverseMap[inst]
		}
	}
	return res
}

func (inst *Stem) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Stick) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *String_mute) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *String_type) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "String":
			res = stage.Technical_String_reverseMap[inst]
		}
	}
	return res
}

func (inst *Strong_accent) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "Strong_accent":
			res = stage.Articulations_Strong_accent_reverseMap[inst]
		}
	}
	return res
}

func (inst *Style_text) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Supports) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Encoding":
		switch reverseField.Fieldname {
		case "Supports":
			res = stage.Encoding_Supports_reverseMap[inst]
		}
	}
	return res
}

func (inst *Swing) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Sync) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Listening":
		switch reverseField.Fieldname {
		case "Sync":
			res = stage.Listening_Sync_reverseMap[inst]
		}
	}
	return res
}

func (inst *System_dividers) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *System_layout) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *System_margins) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Tap) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "Tap":
			res = stage.Technical_Tap_reverseMap[inst]
		}
	}
	return res
}

func (inst *Technical) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Technical":
			res = stage.Notations_Technical_reverseMap[inst]
		}
	}
	return res
}

func (inst *Text_element_data) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Lyric":
		switch reverseField.Fieldname {
		case "Text":
			res = stage.Lyric_Text_reverseMap[inst]
		}
	}
	return res
}

func (inst *Tie) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Tied) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Tied":
			res = stage.Notations_Tied_reverseMap[inst]
		}
	}
	return res
}

func (inst *Time) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Time":
			res = stage.Attributes_Time_reverseMap[inst]
		}
	}
	return res
}

func (inst *Time_modification) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Timpani) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Transpose) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Transpose":
			res = stage.Attributes_Transpose_reverseMap[inst]
		}
	}
	return res
}

func (inst *Tremolo) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Tremolo":
			res = stage.Ornaments_Tremolo_reverseMap[inst]
		}
	}
	return res
}

func (inst *Tuplet) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "Tuplet":
			res = stage.Notations_Tuplet_reverseMap[inst]
		}
	}
	return res
}

func (inst *Tuplet_dot) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Tuplet_portion":
		switch reverseField.Fieldname {
		case "Tuplet_dot":
			res = stage.Tuplet_portion_Tuplet_dot_reverseMap[inst]
		}
	}
	return res
}

func (inst *Tuplet_number) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Tuplet_portion) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Tuplet_type) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Typed_text) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Encoding":
		switch reverseField.Fieldname {
		case "Encoder":
			res = stage.Encoding_Encoder_reverseMap[inst]
		}
	case "Identification":
		switch reverseField.Fieldname {
		case "Creator":
			res = stage.Identification_Creator_reverseMap[inst]
		case "Rights":
			res = stage.Identification_Rights_reverseMap[inst]
		case "Relation":
			res = stage.Identification_Relation_reverseMap[inst]
		}
	}
	return res
}

func (inst *Unpitched) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Virtual_instrument) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Wait) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Listen":
		switch reverseField.Fieldname {
		case "Wait":
			res = stage.Listen_Wait_reverseMap[inst]
		}
	}
	return res
}

func (inst *Wavy_line) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Ornaments":
		switch reverseField.Fieldname {
		case "Wavy_line":
			res = stage.Ornaments_Wavy_line_reverseMap[inst]
		}
	}
	return res
}

func (inst *Wedge) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Wood) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Work) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}
