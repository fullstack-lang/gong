// generated code - do not edit
package note

// insertion point
func (inst *Accidental) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AccidentalMark) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "AccidentalMark":
			if _notations, ok := stage.Notations_AccidentalMark_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	case "Ornaments":
		switch reverseField.Fieldname {
		case "AccidentalMark":
			if _ornaments, ok := stage.Ornaments_AccidentalMark_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		}
	}
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

func (inst *Backup) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Beam) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Note":
		switch reverseField.Fieldname {
		case "Beam":
			if _note, ok := stage.Note_Beam_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	}
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

func (inst *BreathMark) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "BreathMark":
			if _articulations, ok := stage.Articulations_BreathMark_reverseMap[inst]; ok {
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

func (inst *EmptyLine) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Extend) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Figure) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "FiguredBass":
		switch reverseField.Fieldname {
		case "Figure":
			if _figuredbass, ok := stage.FiguredBass_Figure_reverseMap[inst]; ok {
				res = _figuredbass.Name
			}
		}
	}
	return
}

func (inst *FiguredBass) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Forward) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
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

func (inst *Grace) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *HammerOnPullOff) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "HammerOn":
			if _technical, ok := stage.Technical_HammerOn_reverseMap[inst]; ok {
				res = _technical.Name
			}
		case "PullOff":
			if _technical, ok := stage.Technical_PullOff_reverseMap[inst]; ok {
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

func (inst *HarmonClosed) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *HarmonMute) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Technical":
		switch reverseField.Fieldname {
		case "HarmonMute":
			if _technical, ok := stage.Technical_HarmonMute_reverseMap[inst]; ok {
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

func (inst *HeelToe) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *HoleClosed) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
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

func (inst *Listen) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
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
		case "InvertedMordent":
			if _ornaments, ok := stage.Ornaments_InvertedMordent_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		}
	}
	return
}

func (inst *NonArpeggiate) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "NonArpeggiate":
			if _notations, ok := stage.Notations_NonArpeggiate_reverseMap[inst]; ok {
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
	return
}

func (inst *NoteType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Notehead) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *NoteheadText) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
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

func (inst *OtherNotation) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Notations":
		switch reverseField.Fieldname {
		case "OtherNotation":
			if _notations, ok := stage.Notations_OtherNotation_reverseMap[inst]; ok {
				res = _notations.Name
			}
		}
	}
	return
}

func (inst *OtherPlacementText) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "OtherArticulation":
			if _articulations, ok := stage.Articulations_OtherArticulation_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		}
	case "Ornaments":
		switch reverseField.Fieldname {
		case "OtherOrnament":
			if _ornaments, ok := stage.Ornaments_OtherOrnament_reverseMap[inst]; ok {
				res = _ornaments.Name
			}
		}
	case "Technical":
		switch reverseField.Fieldname {
		case "OtherTechnical":
			if _technical, ok := stage.Technical_OtherTechnical_reverseMap[inst]; ok {
				res = _technical.Name
			}
		}
	}
	return
}

func (inst *Pitch) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PlacementText) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Release) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Rest) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
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

func (inst *Stem) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *StrongAccent) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Articulations":
		switch reverseField.Fieldname {
		case "StrongAccent":
			if _articulations, ok := stage.Articulations_StrongAccent_reverseMap[inst]; ok {
				res = _articulations.Name
			}
		}
	}
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

func (inst *TextElementData) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Tie) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Note":
		switch reverseField.Fieldname {
		case "Tie":
			if _note, ok := stage.Note_Tie_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	}
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

func (inst *TupletDot) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "TupletPortion":
		switch reverseField.Fieldname {
		case "TupletDot":
			if _tupletportion, ok := stage.TupletPortion_TupletDot_reverseMap[inst]; ok {
				res = _tupletportion.Name
			}
		}
	}
	return
}

func (inst *TupletNumber) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TupletPortion) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TupletType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Unpitched) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
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
