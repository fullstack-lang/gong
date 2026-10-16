// generated code - do not edit
package direction

// insertion point
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

func (inst *AccordionRegistration) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Barre) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Bass) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Harmony":
		switch reverseField.Fieldname {
		case "Bass":
			if _harmony, ok := stage.Harmony_Bass_reverseMap[inst]; ok {
				res = _harmony.Name
			}
		}
	}
	return
}

func (inst *BassStep) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *BeatUnitTied) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Metronome":
		switch reverseField.Fieldname {
		case "BeatUnitTied":
			if _metronome, ok := stage.Metronome_BeatUnitTied_reverseMap[inst]; ok {
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

func (inst *Bracket) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Dashes) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
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

func (inst *DegreeAlter) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DegreeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DegreeValue) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Direction) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DirectionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Direction":
		switch reverseField.Fieldname {
		case "DirectionType":
			if _direction, ok := stage.Direction_DirectionType_reverseMap[inst]; ok {
				res = _direction.Name
			}
		}
	}
	return
}

func (inst *Effect) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
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

func (inst *FirstFret) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Frame) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *FrameNote) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Frame":
		switch reverseField.Fieldname {
		case "FrameNote":
			if _frame, ok := stage.Frame_FrameNote_reverseMap[inst]; ok {
				res = _frame.Name
			}
		}
	}
	return
}

func (inst *Glass) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Grouping) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Harmony) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *HarmonyAlter) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *HarpPedals) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Image) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *InstrumentChange) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Sound":
		switch reverseField.Fieldname {
		case "InstrumentChange":
			if _sound, ok := stage.Sound_InstrumentChange_reverseMap[inst]; ok {
				res = _sound.Name
			}
		}
	}
	return
}

func (inst *Inversion) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Harmony":
		switch reverseField.Fieldname {
		case "Inversion":
			if _harmony, ok := stage.Harmony_Inversion_reverseMap[inst]; ok {
				res = _harmony.Name
			}
		}
	}
	return
}

func (inst *Kind) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Harmony":
		switch reverseField.Fieldname {
		case "Kind":
			if _harmony, ok := stage.Harmony_Kind_reverseMap[inst]; ok {
				res = _harmony.Name
			}
		}
	}
	return
}

func (inst *Listening) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MeasureNumbering) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
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

func (inst *MetronomeBeam) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MetronomeNote":
		switch reverseField.Fieldname {
		case "MetronomeBeam":
			if _metronomenote, ok := stage.MetronomeNote_MetronomeBeam_reverseMap[inst]; ok {
				res = _metronomenote.Name
			}
		}
	}
	return
}

func (inst *MetronomeNote) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Metronome":
		switch reverseField.Fieldname {
		case "MetronomeNote":
			if _metronome, ok := stage.Metronome_MetronomeNote_reverseMap[inst]; ok {
				res = _metronome.Name
			}
		}
	}
	return
}

func (inst *MetronomeTied) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MetronomeTuplet) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Numeral) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Harmony":
		switch reverseField.Fieldname {
		case "Numeral":
			if _harmony, ok := stage.Harmony_Numeral_reverseMap[inst]; ok {
				res = _harmony.Name
			}
		}
	}
	return
}

func (inst *NumeralKey) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *NumeralRoot) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *OctaveShift) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Offset) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *OtherDirection) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *OtherListening) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Listening":
		switch reverseField.Fieldname {
		case "OtherListening":
			if _listening, ok := stage.Listening_OtherListening_reverseMap[inst]; ok {
				res = _listening.Name
			}
		}
	}
	return
}

func (inst *Pedal) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PedalTuning) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "HarpPedals":
		switch reverseField.Fieldname {
		case "PedalTuning":
			if _harppedals, ok := stage.HarpPedals_PedalTuning_reverseMap[inst]; ok {
				res = _harppedals.Name
			}
		}
	}
	return
}

func (inst *PerMinute) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Percussion) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DirectionType":
		switch reverseField.Fieldname {
		case "Percussion":
			if _directiontype, ok := stage.DirectionType_Percussion_reverseMap[inst]; ok {
				res = _directiontype.Name
			}
		}
	}
	return
}

func (inst *Pitched) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PrincipalVoice) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Print) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Root) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Harmony":
		switch reverseField.Fieldname {
		case "Root":
			if _harmony, ok := stage.Harmony_Root_reverseMap[inst]; ok {
				res = _harmony.Name
			}
		}
	}
	return
}

func (inst *RootStep) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Scordatura) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Sound) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *StaffDivide) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Stick) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *StringMute) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *StyleText) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Harmony":
		switch reverseField.Fieldname {
		case "Function":
			if _harmony, ok := stage.Harmony_Function_reverseMap[inst]; ok {
				res = _harmony.Name
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

func (inst *TimeModification) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Timpani) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *VirtualInstrument) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
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
