package models

func (stager *Stager) enforcePlantHasMusicAbstract() (needCommit bool) {
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.PlantType == Music {
			if plant.MusicAbstract == nil {
				ma := (&MusicAbstract{
					Name:                           plant.Name + "-MusicAbstract",
					IsChecked:                      true,
					PitchHeight:                    0.138,
					NbOfBeatsInTheme:               16,
					BeatsPerSecond:                 6.0,
					FirstVoiceShiftX:               0.1,
					FirstVoiceShiftY:               2.40,
					PitchDifference:                12,
					Level:                          11.1,
					ActualBeatsTemporalShift:       6,
					IsMinor:                        true,
					ThemeBinaryEncoding:            0xFFFF,
					BezierControlLengthRatio:       0.56,
					NbPitchLines:                   50,
					NbBeatLines:                    64,
					OriginX:                        50.0,
					OriginY:                        750.0,
					ShowFirstVoice:                 true,
					ShowFirstVoiceShiftRight:       true,
					ShowSecondVoice:                true,
					ShowSecondVoiceShiftRight:      true,
					ShowFirstVoiceNotes:            true,
					ShowFirstVoiceNotesShiftRight:  true,
					ShowSecondVoiceNotes:           true,
					ShowSecondVoiceNotesShiftRight: true,
					IsComposerNodeExpanded:         true,
				}).Stage(stager.stage)
				plant.MusicAbstract = ma
				needCommit = true
			} else {
				if plant.MusicAbstract.PitchHeight <= 0 {
					plant.MusicAbstract.PitchHeight = 0.138
					needCommit = true
				}
				if plant.MusicAbstract.NbOfBeatsInTheme <= 0 {
					plant.MusicAbstract.NbOfBeatsInTheme = 16
					needCommit = true
				}
				if plant.MusicAbstract.BeatsPerSecond <= 0 {
					plant.MusicAbstract.BeatsPerSecond = 6.0
					needCommit = true
				}
				if plant.MusicAbstract.BezierControlLengthRatio <= 0 {
					plant.MusicAbstract.BezierControlLengthRatio = 0.56
					needCommit = true
				}
				if plant.MusicAbstract.NbPitchLines <= 0 {
					plant.MusicAbstract.NbPitchLines = 50
					needCommit = true
				}
				if plant.MusicAbstract.NbBeatLines <= 0 {
					plant.MusicAbstract.NbBeatLines = 64
					needCommit = true
				}
				if plant.MusicAbstract.OriginY <= 0 {
					plant.MusicAbstract.OriginY = 750.0
					needCommit = true
				}
				if plant.MusicAbstract.ThemeBinaryEncoding == 0 {
					plant.MusicAbstract.ThemeBinaryEncoding = 0xFFFF
					needCommit = true
				}
			}
		} else {
			if plant.MusicAbstract != nil {
				plant.MusicAbstract = nil
				needCommit = true
			}
		}
	}

	// Unstage unreferenced MusicAbstract
	for ma := range *GetGongstructInstancesSetFromPointerType[*MusicAbstract](stager.stage) {
		hasOwner := false
		for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
			if plant.MusicAbstract == ma {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			ma.Unstage(stager.stage)
			needCommit = true
		}
	}

	return needCommit
}

func (stager *Stager) enforceMusicAbstractName() (needCommit bool) {
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		if plant.PlantType == Music && plant.MusicAbstract != nil {
			expectedName := plant.Name + "-MusicAbstract"
			if plant.MusicAbstract.Name != expectedName {
				plant.MusicAbstract.Name = expectedName
				needCommit = true
			}
		}
	}
	return needCommit
}
