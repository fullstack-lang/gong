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

		if plant.CurrentView != VIEW_PLANT_2D && plant.VaseAbstract != nil {

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

	// Add sliders for PlantDiagram 3D view
	var checkedDiagram *PlantDiagram
	for _, diagram := range plant.PlantDiagrams {
		if diagram.IsChecked {
			checkedDiagram = diagram
			break
		}
	}

	if plant.CurrentView != VIEW_PLANT_2D && checkedDiagram != nil && checkedDiagram.VaseDiagram != nil && checkedDiagram.VaseDiagram.Rendered3DShape != nil {
		group2 := new(m.Group).Stage(stager.sliderStage)
		group2.Percentage = 35
		layout.Groups = append(layout.Groups, group2)

		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"View X",
				-500,
				500,
				1,
				&checkedDiagram.VaseDiagram.Rendered3DShape.ViewX,
			),
		)
		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"View Y",
				-500,
				500,
				1,
				&checkedDiagram.VaseDiagram.Rendered3DShape.ViewY,
			),
		)
		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"View Z",
				-500,
				500,
				1,
				&checkedDiagram.VaseDiagram.Rendered3DShape.ViewZ,
			),
		)
		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"Target X",
				-500,
				500,
				1,
				&checkedDiagram.VaseDiagram.Rendered3DShape.TargetX,
			),
		)
		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"Target Y",
				-500,
				500,
				1,
				&checkedDiagram.VaseDiagram.Rendered3DShape.TargetY,
			),
		)
		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"Target Z",
				-500,
				500,
				1,
				&checkedDiagram.VaseDiagram.Rendered3DShape.TargetZ,
			),
		)
		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"Fov",
				-500,
				500,
				1,
				&checkedDiagram.VaseDiagram.Rendered3DShape.Fov,
			),
		)
	}

	stager.sliderStage.Commit()
}

func (stager *Stager) OnAfterUpdateSliderElement() {

	stager.enforceSemantic()
	stager.ux_tree()
	stager.ux_svg_plant_diagram()
	stager.ux_3d_plant_diagram()

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
