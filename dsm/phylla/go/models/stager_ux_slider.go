package models

import (
	m "github.com/fullstack-lang/gong/lib/slider/go/models"
)

func (stager *Stager) ux_slider() {

	plant := stager.selectedPlant
	if plant == nil {
		plant = stager.GetCurrentPlant()
	}

	stager.sliderStage.Reset()

	if plant == nil {
		stager.sliderStage.Commit()
		return
	}

	layout := new(m.Layout).Stage(stager.sliderStage)

	group1 := new(m.Group).Stage(stager.sliderStage)
	group1.Percentage = 65
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
				"Side Length",
				5,
				600,
				5,
				&plant.RhombusSideLength,
			),
		)

		group1.Sliders = append(
			group1.Sliders,
			m.NewSlider(
				stager,
				"Inside Angle",
				0,
				180,
				1,
				&plant.RhombusInsideAngle,
			),
		)

		group1.Sliders = append(
			group1.Sliders,
			m.NewSlider(
				stager,
				"Stack Height",
				0,
				20,
				1,
				&plant.StackHeight,
			),
		)

		if plant.CurrentView != VIEW_PLANT_2D && plant.PlantType == Vase {

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Thickness",
					0.01,
					0.3,
					0.01,
					&plant.VaseAbstract.RelativeVerticalThickness,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"2D Separation",
					0.0,
					0.1,
					0.002,
					&plant.VaseAbstract.RelativeCuttedStackFloorHeight,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"3D Separation",
					0.0,
					1.0,
					0.01,
					&plant.VaseAbstract.RelativeRotatedTorusSeparation,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Radial Thickness",
					0.01,
					0.5,
					0.01,
					&plant.VaseAbstract.RelativeRadialThickness,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Rot Ratio",
					0.0,
					1.0,
					0.005,
					&plant.VaseAbstract.RotationRatio,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Radial Repetition",
					1,
					4,
					1,
					&plant.VaseAbstract.RadialRepetitions,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Transparency",
					0.0,
					1.0,
					0.05,
					&plant.VaseAbstract.Transparency,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				NewBoolSlider(
					stager,
					"Alternating Ring Colors",
					&plant.VaseAbstract.HasAlternatingRingColors,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Traj offset X",
					-0.15,
					0.15,
					0.001,
					&plant.VaseAbstract.RelativeTrajectoryOffsetX,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Traj offset Y",
					-0.15,
					0.15,
					0.001,
					&plant.VaseAbstract.RelativeTrajectoryOffsetY,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Nb Step P1 P2",
					1,
					30,
					1,
					&plant.VaseAbstract.NbStepP1P2,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Chosen Step",
					1,
					plant.VaseAbstract.NbStepP1P2,
					1,
					&plant.VaseAbstract.ChosenStep,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Horiz Ring Height",
					0.0,
					1.0,
					0.005,
					&plant.VaseAbstract.RelativeHorizontalRingsHeight,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Offset Key X",
					-500,
					500,
					1,
					&plant.VaseAbstract.OffsetKeyX,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Offset Key Y",
					-500,
					500,
					1,
					&plant.VaseAbstract.OffsetKeyY,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Width Key",
					0,
					500,
					1,
					&plant.VaseAbstract.WidthKey,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Height Key",
					0,
					500,
					1,
					&plant.VaseAbstract.HeightKey,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Key Size Reduction",
					0.0,
					1.0,
					0.01,
					&plant.VaseAbstract.RelativeKeySize,
				),
			)

			group1.Sliders = append(
				group1.Sliders,
				m.NewSlider(
					stager,
					"Movie Nb Frames",
					0,
					1000,
					1,
					&plant.VaseAbstract.MovieNbFrames,
				),
			)
		}

	}

	stager.sliderStage.Commit()
}

type SliderStageTarget struct {
	stager      *Stager
	sliderStage *m.Stage
}

func (t *SliderStageTarget) GetSliderStage() *m.Stage {
	return t.sliderStage
}

func (t *SliderStageTarget) OnAfterUpdateSliderElement() {
	t.stager.OnAfterUpdateSliderElement()
}

func appendCommonCylinderSliders(
	target m.Target,
	group *m.Group,
	plant *PlantAbstract,
	radialReps *int,
	transparency *float64,
	relTubeDiam *float64,
	relHeight3DTorus *float64,
	torusVertScale *float64,
	relHeight *float64,
	projAngle *float64,
) {
	group.Sliders = append(
		group.Sliders,
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
			"Stack Height",
			0,
			20,
			1,
			&plant.StackHeight,
		),
		m.NewSlider(
			target,
			"Radial Repetition",
			1,
			10,
			1,
			radialReps,
		),
		m.NewSlider(
			target,
			"Transparency",
			0.0,
			1.0,
			0.05,
			transparency,
		),
		m.NewSlider(
			target,
			"Tube Rel Diameter",
			0.005,
			0.1,
			0.001,
			relTubeDiam,
		),
		m.NewSlider(
			target,
			"Rel Height 3D Torus",
			0.0,
			6.0,
			0.01,
			relHeight3DTorus,
		),
		m.NewSlider(
			target,
			"Torus Vert Scale",
			0.0,
			5.0,
			0.01,
			torusVertScale,
		),
		m.NewSlider(
			target,
			"Rel Height",
			0.0,
			6.0,
			0.01,
			relHeight,
		),
		m.NewSlider(
			target,
			"Projection Angle",
			0.0,
			45,
			0.1,
			projAngle,
		),
	)
}

func (stager *Stager) ux_slider_stool() {
	plant := stager.selectedPlant
	if plant == nil {
		plant = stager.GetCurrentPlant()
	}

	stager.sliderStoolStage.Reset()

	if plant == nil || plant.PlantType != Stool || plant.StoolAbstract == nil {
		stager.sliderStoolStage.Commit()
		return
	}

	layout := new(m.Layout).Stage(stager.sliderStoolStage)

	group1 := new(m.Group).Stage(stager.sliderStoolStage)
	group1.Percentage = 65
	layout.Groups = append(layout.Groups, group1)

	target := &SliderStageTarget{stager: stager, sliderStage: stager.sliderStoolStage}

	appendCommonCylinderSliders(
		target,
		group1,
		plant,
		&plant.StoolAbstract.RadialRepetitions,
		&plant.StoolAbstract.Transparency,
		&plant.StoolAbstract.RelativeTubeDiameter,
		&plant.StoolAbstract.RelativeHeight3DTorus,
		&plant.StoolAbstract.StoolTorusVerticalScale,
		&plant.StoolAbstract.RelativeHeight,
		&plant.StoolAbstract.ProjectionAngle,
	)

	group1.Sliders = append(
		group1.Sliders,
		m.NewSlider(
			target,
			"Rel Seat Thickness",
			0.01,
			1.0,
			0.01,
			&plant.StoolAbstract.RelativeSeatThickness,
		),
		m.NewSlider(
			target,
			"Rel Eye Separation",
			0.0,
			0.5,
			0.005,
			&plant.StoolAbstract.RelativeEyeSeparationCriteria,
		),
		m.NewSlider(
			target,
			"Rel Eye Corner Strength",
			0.0,
			2.0,
			0.01,
			&plant.StoolAbstract.RelativeEyeCornerControlVectorStrength,
		),
	)

	stager.sliderStoolStage.Commit()
}

func (stager *Stager) ux_slider_clock() {
	plant := stager.selectedPlant
	if plant == nil {
		plant = stager.GetCurrentPlant()
	}

	stager.sliderClockStage.Reset()

	if plant == nil || plant.PlantType != Clock || plant.ClockAbstract == nil {
		stager.sliderClockStage.Commit()
		return
	}

	layout := new(m.Layout).Stage(stager.sliderClockStage)

	group1 := new(m.Group).Stage(stager.sliderClockStage)
	group1.Percentage = 65
	layout.Groups = append(layout.Groups, group1)

	target := &SliderStageTarget{stager: stager, sliderStage: stager.sliderClockStage}

	appendCommonCylinderSliders(
		target,
		group1,
		plant,
		&plant.ClockAbstract.RadialRepetitions,
		&plant.ClockAbstract.Transparency,
		&plant.ClockAbstract.RelativeTubeDiameter,
		&plant.ClockAbstract.RelativeHeight3DTorus,
		&plant.ClockAbstract.ClockTorusVerticalScale,
		&plant.ClockAbstract.RelativeHeight,
		&plant.ClockAbstract.ProjectionAngle,
	)

	stager.sliderClockStage.Commit()
}

func (stager *Stager) OnAfterUpdateSliderElement() {

	stager.enforceSemantic()
	stager.ux_tree()
	stager.ux_svg_plant_diagram()
	stager.UpdateThreeJSStage()
	stager.UpdateStool3DStage()
	stager.UpdateClock3DStage()

	stager.stage.CommitWithSuspendedCallbacks()
}

type BoolSliderProxy struct {
	slider *m.Slider
	Value  *bool
	stager *Stager
}

func (proxy *BoolSliderProxy) Updated() {
	*proxy.Value = (proxy.slider.ValueInt != 0)
	proxy.stager.OnAfterUpdateSliderElement()
}

func NewBoolSlider(
	stager *Stager,
	name string,
	valueRef *bool,
) *m.Slider {
	slider := new(m.Slider).Stage(stager.sliderStage)
	slider.Name = name
	slider.IsInt = true
	slider.MinInt = 0
	slider.MaxInt = 1
	slider.StepInt = 1
	if *valueRef {
		slider.ValueInt = 1
	} else {
		slider.ValueInt = 0
	}

	proxy := &BoolSliderProxy{
		slider: slider,
		Value:  valueRef,
		stager: stager,
	}

	slider.Proxy = proxy
	return slider
}
