package models

import "fmt"

func (stager *Stager) enforcePlantHasMusicAbstract() (needCommit bool) {
	for plant := range *stager.stage.GetInstancesSet[*PlantAbstract]() {
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
					OriginY:                        600.0,
					ScoreScale:                     1.0,
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
				stager.logAndNotify(fmt.Sprintf("Plant %s: created missing MusicAbstract", plant.Name))
			} else {
				if plant.MusicAbstract.ScoreScale <= 0 {
					plant.MusicAbstract.ScoreScale = 1.0
					needCommit = true
					stager.logAndNotify(fmt.Sprintf("Plant %s Music: default ScoreScale set to 1.0", plant.Name))
				}
				if plant.MusicAbstract.PitchHeight <= 0 {
					plant.MusicAbstract.PitchHeight = 0.138
					needCommit = true
					stager.logAndNotify(fmt.Sprintf("Plant %s Music: default PitchHeight set to 0.138", plant.Name))
				}
				if plant.MusicAbstract.NbOfBeatsInTheme <= 0 {
					plant.MusicAbstract.NbOfBeatsInTheme = 16
					needCommit = true
					stager.logAndNotify(fmt.Sprintf("Plant %s Music: default NbOfBeatsInTheme set to 16", plant.Name))
				}
				if plant.MusicAbstract.BeatsPerSecond <= 0 {
					plant.MusicAbstract.BeatsPerSecond = 6.0
					needCommit = true
					stager.logAndNotify(fmt.Sprintf("Plant %s Music: default BeatsPerSecond set to 6.0", plant.Name))
				}
				if plant.MusicAbstract.BezierControlLengthRatio <= 0 {
					plant.MusicAbstract.BezierControlLengthRatio = 0.56
					needCommit = true
					stager.logAndNotify(fmt.Sprintf("Plant %s Music: default BezierControlLengthRatio set to 0.56", plant.Name))
				}
				if plant.MusicAbstract.NbPitchLines <= 0 {
					plant.MusicAbstract.NbPitchLines = 50
					needCommit = true
					stager.logAndNotify(fmt.Sprintf("Plant %s Music: default NbPitchLines set to 50", plant.Name))
				}
				if plant.MusicAbstract.NbBeatLines <= 0 {
					plant.MusicAbstract.NbBeatLines = 64
					needCommit = true
					stager.logAndNotify(fmt.Sprintf("Plant %s Music: default NbBeatLines set to 64", plant.Name))
				}
				if plant.MusicAbstract.OriginY <= 0 {
					plant.MusicAbstract.OriginY = 600.0
					needCommit = true
					stager.logAndNotify(fmt.Sprintf("Plant %s Music: default OriginY set to 600.0", plant.Name))
				}
				if plant.MusicAbstract.ThemeBinaryEncoding == 0 {
					plant.MusicAbstract.ThemeBinaryEncoding = 0xFFFF
					needCommit = true
					stager.logAndNotify(fmt.Sprintf("Plant %s Music: default ThemeBinaryEncoding set to 0xFFFF", plant.Name))
				}
			}
		} else {
			if plant.MusicAbstract != nil {
				plant.MusicAbstract = nil
				needCommit = true
				stager.logAndNotify(fmt.Sprintf("Plant %s: removed MusicAbstract because PlantType is %s", plant.Name, plant.PlantType))
			}
		}
	}

	// Unstage unreferenced MusicAbstract
	for ma := range *stager.stage.GetInstancesSet[*MusicAbstract]() {
		hasOwner := false
		for plant := range *stager.stage.GetInstancesSet[*PlantAbstract]() {
			if plant.MusicAbstract == ma {
				hasOwner = true
				break
			}
		}
		if !hasOwner {
			ma.Unstage(stager.stage)
			needCommit = true
			stager.logAndNotify(fmt.Sprintf("Removed orphaned MusicAbstract %s", ma.Name))
		}
	}

	return needCommit
}

func (stager *Stager) enforceMusicAbstractName() (needCommit bool) {
	for plant := range *stager.stage.GetInstancesSet[*PlantAbstract]() {
		if plant.PlantType == Music && plant.MusicAbstract != nil {
			expectedName := plant.Name + "-MusicAbstract"
			if plant.MusicAbstract.Name != expectedName {
				oldName := plant.MusicAbstract.Name
				plant.MusicAbstract.Name = expectedName
				needCommit = true
				stager.logAndNotify(fmt.Sprintf("Renamed MusicAbstract from '%s' to '%s'", oldName, expectedName))
			}
		}
	}
	return needCommit
}
