package models

import (
	m "github.com/fullstack-lang/gong/lib/slider/go/models"
)

func (stager *Stager) ux_slider() {

	plant := stager.selectedPlant

	stager.sliderStage.Reset()

	if plant == nil {
		stager.sliderStage.Commit()
		return
	}

	layout := new(m.Layout).Stage(stager.sliderStage)

	group1 := new(m.Group).Stage(stager.sliderStage)
	group1.Percentage = 25
	layout.Groups = append(layout.Groups, group1)

	// group1 := new(m.Group).Stage(stager.sliderStage)
	// group1.Percentage = 25
	// layout.Groups = append(layout.Groups, group1)

	// group3 := new(m.Group).Stage(stager.sliderStage)
	// group3.Percentage = 50
	// layout.Groups = append(layout.Groups, group3)

	{
		group1.Sliders = append(
			group1.Sliders,
			m.NewSlider(
				stager,
				"Side Length",
				5,
				300,
				5,
				&plant.SideLength,
			),
		)

		group1.Sliders = append(
			group1.Sliders,
			m.NewSlider(
				stager,
				"Inside Angle",
				60,
				120,
				1,
				&plant.InsideAngle,
			),
		)

		group1.Sliders = append(
			group1.Sliders,
			m.NewSlider(
				stager,
				"M",
				1,
				20,
				1,
				&plant.M,
			),
		)

		group1.Sliders = append(
			group1.Sliders,
			m.NewSlider(
				stager,
				"N",
				1,
				20,
				1,
				&plant.N,
			),
		)

		group1.Sliders = append(
			group1.Sliders,
			m.NewSlider(
				stager,
				"Z",
				1,
				120,
				1,
				&plant.Z,
			),
		)

	}

	stager.sliderStage.Commit()
}

func (stager *Stager) OnAfterUpdateSliderElement() {

	stager.enforceSemantic()
	stager.stage.CommitWithSuspendedCallbacks()
}
