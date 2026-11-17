package models

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
)

func (stager *Stager) ux_slider_music() {
	plant := stager.selectedPlant
	if plant == nil {
		plant = stager.GetCurrentPlant()
	}

	stager.sliderMusicStage.Reset()

	if plant == nil || plant.PlantType != Music || plant.MusicAbstract == nil {
		stager.sliderMusicStage.Commit()
		return
	}

	formGroup := form.NewLiveFormGroup(stager.sliderMusicStage, "Music Controls")
	onUpdate := stager.OnAfterUpdateSliderElement

	// Group 1: Score Navigation & Scale (Origin X, Y and Scale)
	form.StartAccordion(stager.sliderMusicStage, formGroup, "Score Navigation & Scale", true)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "Score Scale", 0.1, 3.0, 0.05, &plant.MusicAbstract.ScoreScale, onUpdate)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "Score X", -500.0, 2000.0, 10.0, &plant.MusicAbstract.OriginX, onUpdate)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "Score Y", 0.0, 3000.0, 20.0, &plant.MusicAbstract.OriginY, onUpdate)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "Side Length", 5, 600, 5, &plant.RhombusSideLength, onUpdate)

	// Group 2: Plant geometry & lattice parameters
	form.StartAccordion(stager.sliderMusicStage, formGroup, "Plant Lattice", true)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "N", 1, 20, 1, &plant.N, onUpdate)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "M", 1, 20, 1, &plant.M, onUpdate)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "Inside Angle", 0, 180, 1, &plant.RhombusInsideAngle, onUpdate)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "Bezier Strength", 0.0, 4.0, 0.01, &plant.MusicAbstract.BezierControlLengthRatio, onUpdate)

	// Group 3: Music composition parameters (matching phyllotaxymusic)
	form.StartAccordion(stager.sliderMusicStage, formGroup, "Music Composition", true)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "Pitch Height", 0.001, 0.3, 0.001, &plant.MusicAbstract.PitchHeight, onUpdate)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "Nb Beats in Theme", 1, 64, 1, &plant.MusicAbstract.NbOfBeatsInTheme, onUpdate)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "BeatsPerSecond", 0.1, 20.0, 0.05, &plant.MusicAbstract.BeatsPerSecond, onUpdate)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "1st voice X", -1.0, 1.0, 0.01, &plant.MusicAbstract.FirstVoiceShiftX, onUpdate)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "1st voice Y", -1.0, 4.0, 0.01, &plant.MusicAbstract.FirstVoiceShiftY, onUpdate)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "2nd voice pitch diff", -12, 24, 1, &plant.MusicAbstract.PitchDifference, onUpdate)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "Level", 0.0, 20.0, 0.1, &plant.MusicAbstract.Level, onUpdate)
	form.AppendSlider(stager.sliderMusicStage, formGroup, "Actual Beats Shift", 0, 20, 1, &plant.MusicAbstract.ActualBeatsTemporalShift, onUpdate)
	form.AppendRadioCheckBox(stager.sliderMusicStage, formGroup, "Scale", "Minor", "Major", &plant.MusicAbstract.IsMinor, onUpdate)
	form.EndAccordion(stager.sliderMusicStage, formGroup)

	stager.sliderMusicStage.Commit()
}
