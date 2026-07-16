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
				"Rhombus Side Length",
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
				"Rhombus Inside Angle",
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
				10,
				1,
				&plant.StackHeight,
			),
		)

		group1.Sliders = append(
			group1.Sliders,
			m.NewSlider(
				stager,
				"Relative Vertical Thickness",
				0.01,
				0.3,
				0.01,
				&plant.RelativeVerticalThickness,
			),
		)

	}

	// Add sliders for PlantDiagram 3D view
	var checkedDiagram *PlantDiagram
	for _, diagram := range plant.PlantDiagrams {
		if diagram.IsChecked {
			checkedDiagram = diagram
			break
		}
	}

	if checkedDiagram != nil && checkedDiagram.Rendered3DShape != nil {
		group2 := new(m.Group).Stage(stager.sliderStage)
		group2.Percentage = 25
		layout.Groups = append(layout.Groups, group2)

		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"View X",
				-500,
				500,
				1,
				&checkedDiagram.Rendered3DShape.ViewX,
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
				&checkedDiagram.Rendered3DShape.ViewY,
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
				&checkedDiagram.Rendered3DShape.ViewZ,
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
				&checkedDiagram.Rendered3DShape.TargetX,
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
				&checkedDiagram.Rendered3DShape.TargetY,
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
				&checkedDiagram.Rendered3DShape.TargetZ,
			),
		)
		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"Field of View",
				10,
				120,
				1,
				&checkedDiagram.Rendered3DShape.Fov,
			),
		)
	}

	stager.sliderStage.Commit()
}

func (stager *Stager) OnAfterUpdateSliderElement() {

	stager.enforceSemantic()
	stager.ux_svg_plant_diagram()
	stager.ux_3d_plant_diagram()

	stager.stage.CommitWithSuspendedCallbacks()
}
