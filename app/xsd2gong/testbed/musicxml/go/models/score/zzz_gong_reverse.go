// generated code - do not edit
package score

// insertion point
func (inst *Credit) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ScorePartwise":
		switch reverseField.Fieldname {
		case "Credit":
			if _scorepartwise, ok := stage.ScorePartwise_Credit_reverseMap[inst]; ok {
				res = _scorepartwise.Name
			}
		}
	case "ScoreTimewise":
		switch reverseField.Fieldname {
		case "Credit":
			if _scoretimewise, ok := stage.ScoreTimewise_Credit_reverseMap[inst]; ok {
				res = _scoretimewise.Name
			}
		}
	}
	return
}

func (inst *Defaults) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EmptyFont) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GroupBarline) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GroupName) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GroupSymbol) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *InstrumentLink) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "PartLink":
		switch reverseField.Fieldname {
		case "InstrumentLink":
			if _partlink, ok := stage.PartLink_InstrumentLink_reverseMap[inst]; ok {
				res = _partlink.Name
			}
		}
	}
	return
}

func (inst *LyricFont) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Defaults":
		switch reverseField.Fieldname {
		case "LyricFont":
			if _defaults, ok := stage.Defaults_LyricFont_reverseMap[inst]; ok {
				res = _defaults.Name
			}
		}
	}
	return
}

func (inst *LyricLanguage) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Defaults":
		switch reverseField.Fieldname {
		case "LyricLanguage":
			if _defaults, ok := stage.Defaults_LyricLanguage_reverseMap[inst]; ok {
				res = _defaults.Name
			}
		}
	}
	return
}

func (inst *Measure) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Part":
		switch reverseField.Fieldname {
		case "Measure":
			if _part, ok := stage.Part_Measure_reverseMap[inst]; ok {
				res = _part.Name
			}
		}
	}
	return
}

func (inst *MeasureMusicDataChoice) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Measure":
		switch reverseField.Fieldname {
		case "MusicData":
			if _measure, ok := stage.Measure_MusicData_reverseMap[inst]; ok {
				res = _measure.Name
			}
		}
	}
	return
}

func (inst *Opus) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Part) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ScorePartwise":
		switch reverseField.Fieldname {
		case "Part":
			if _scorepartwise, ok := stage.ScorePartwise_Part_reverseMap[inst]; ok {
				res = _scorepartwise.Name
			}
		}
	}
	return
}

func (inst *PartGroup) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "PartList":
		switch reverseField.Fieldname {
		case "PartGroup":
			if _partlist, ok := stage.PartList_PartGroup_reverseMap[inst]; ok {
				res = _partlist.Name
			}
		}
	}
	return
}

func (inst *PartLink) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ScorePart":
		switch reverseField.Fieldname {
		case "PartLink":
			if _scorepart, ok := stage.ScorePart_PartLink_reverseMap[inst]; ok {
				res = _scorepart.Name
			}
		}
	}
	return
}

func (inst *PartList) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PartName) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Player) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ScorePart":
		switch reverseField.Fieldname {
		case "Player":
			if _scorepart, ok := stage.ScorePart_Player_reverseMap[inst]; ok {
				res = _scorepart.Name
			}
		}
	}
	return
}

func (inst *ScoreInstrument) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ScorePart":
		switch reverseField.Fieldname {
		case "ScoreInstrument":
			if _scorepart, ok := stage.ScorePart_ScoreInstrument_reverseMap[inst]; ok {
				res = _scorepart.Name
			}
		}
	}
	return
}

func (inst *ScorePart) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "PartList":
		switch reverseField.Fieldname {
		case "ScorePart":
			if _partlist, ok := stage.PartList_ScorePart_reverseMap[inst]; ok {
				res = _partlist.Name
			}
		}
	}
	return
}

func (inst *ScorePartwise) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ScoreTimewise) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ScoreTimewiseMeasure) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ScoreTimewise":
		switch reverseField.Fieldname {
		case "Measure":
			if _scoretimewise, ok := stage.ScoreTimewise_Measure_reverseMap[inst]; ok {
				res = _scoretimewise.Name
			}
		}
	}
	return
}

func (inst *ScoreTimewiseMeasurePart) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ScoreTimewiseMeasure":
		switch reverseField.Fieldname {
		case "Part":
			if _scoretimewisemeasure, ok := stage.ScoreTimewiseMeasure_Part_reverseMap[inst]; ok {
				res = _scoretimewisemeasure.Name
			}
		}
	}
	return
}

func (inst *Work) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}
