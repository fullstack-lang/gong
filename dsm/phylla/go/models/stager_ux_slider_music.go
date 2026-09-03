package models

import (
	m "github.com/fullstack-lang/gong/lib/slider/go/models"
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

	layout := new(m.Layout).Stage(stager.sliderMusicStage)

	group1 := new(m.Group).Stage(stager.sliderMusicStage)
	group1.Percentage = 40
	layout.Groups = append(layout.Groups, group1)

	group2 := new(m.Group).Stage(stager.sliderMusicStage)
	group2.Percentage = 60
	layout.Groups = append(layout.Groups, group2)

	target := &SliderStageTarget{stager: stager, sliderStage: stager.sliderMusicStage}

	// Group 1: Plant geometry & lattice parameters
	group1.Sliders = append(
		group1.Sliders,
		m.NewSlider(
			target,
			"N",
			1,
			20,
			1,
			&plant.N,
		),
		m.NewSlider(
			target,
			"M",
			1,
			20,
			1,
			&plant.M,
		),
		m.NewSlider(
			target,
			"Side Length",
			5,
			600,
			5,
			&plant.RhombusSideLength,
		),
		m.NewSlider(
			target,
			"Inside Angle",
			0,
			180,
			1,
			&plant.RhombusInsideAngle,
		),
		m.NewSlider(
			target,
			"Bezier Strength",
			0,
			4,
			0.01,
			&plant.MusicAbstract.BezierControlLengthRatio,
		),
	)

	// Group 2: Music composition parameters (matching phyllotaxymusic)
	group2.Sliders = append(
		group2.Sliders,
		m.NewSlider(
			target,
			"Pitch Height",
			0.001,
			0.3,
			0.001,
			&plant.MusicAbstract.PitchHeight,
		),
		m.NewSlider(
			target,
			"Nb Beats in Theme",
			1,
			64,
			1,
			&plant.MusicAbstract.NbOfBeatsInTheme,
		),
		m.NewSlider(
			target,
			"BeatsPerSecond",
			0.1,
			20,
			0.05,
			&plant.MusicAbstract.BeatsPerSecond,
		),
		m.NewSlider(
			target,
			"1st voice X",
			-1,
			1,
			0.01,
			&plant.MusicAbstract.FirstVoiceShiftX,
		),
		m.NewSlider(
			target,
			"1st voice Y",
			-1,
			4,
			0.01,
			&plant.MusicAbstract.FirstVoiceShiftY,
		),
		m.NewSlider(
			target,
			"2nd voice pitch diff",
			-12,
			24,
			1,
			&plant.MusicAbstract.PitchDifference,
		),
		m.NewSlider(
			target,
			"Level",
			0,
			20,
			0.1,
			&plant.MusicAbstract.Level,
		),
		m.NewSlider(
			target,
			"Actual Beats Shift",
			0,
			20,
			1,
			&plant.MusicAbstract.ActualBeatsTemporalShift,
		),
	)

	group2.Checkboxes = append(
		group2.Checkboxes,
		m.NewCheckbox(
			target,
			"Scale",
			"Minor",
			"Major",
			&plant.MusicAbstract.IsMinor,
		),
	)

	stager.sliderMusicStage.Commit()
}
