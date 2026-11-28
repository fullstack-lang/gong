package models

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"
)

func appendPlantTypeSelect(stager *Stager, formGroup *form.FormGroup, plant *PlantAbstract) {
	codes := plant.PlantType.Codes()
	codeValues := plant.PlantType.CodeValues()
	currentCode := ""
	for idx, code := range codes {
		if codeValues[idx] == plant.PlantType.ToString() {
			currentCode = code
			break
		}
	}

	form.AppendSelect(
		stager.sliderStage,
		formGroup,
		"PlantType",
		"Plant Type",
		codes,
		currentCode,
		func(selected string) {
			var newType PlantType
			if err := (&newType).FromCodeString(selected); err == nil {
				if plant.PlantType != newType {
					plant.PlantType = newType
					if newType == Stool {
						if plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_STOOL_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
							plant.CurrentView = VIEW_PLANT_2D
						}
					} else if newType == Clock {
						if plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_CLOCK_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
							plant.CurrentView = VIEW_PLANT_2D
						}
					} else if newType != TubeVase && newType != TrapezeVase {
						if plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
							plant.CurrentView = VIEW_PLANT_2D
						}
					}
					stager.stage.Commit()
				}
			}
		},
	)
}

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

	formGroup := form.NewLiveFormGroup(stager.sliderStage, "Plant Controls")
	onUpdate := stager.OnAfterUpdateSliderElement

	if plant.CurrentView == VIEW_PLANT_2D {
		appendPlantTypeSelect(stager, formGroup, plant)
	}

	form.StartAccordion(stager.sliderStage, formGroup, "Plant Lattice", true)

	form.AppendSlider(stager.sliderStage, formGroup, "N", 1, 20, 1, &plant.N, onUpdate)
	form.AppendSlider(stager.sliderStage, formGroup, "M", 1, 20, 1, &plant.M, onUpdate)
	form.AppendSlider(stager.sliderStage, formGroup, "Side Length", 5, 600, 5, &plant.RhombusSideLength, onUpdate)

	if zoom := stager.getActive2DDiagramZoom(plant); zoom != nil {
		form.AppendSlider(stager.sliderStage, formGroup, "Zoom", 0.05, 5.0, 0.05, zoom, onUpdate)
	}

	form.AppendSlider(stager.sliderStage, formGroup, "Inside Angle", 0, 180, 1, &plant.RhombusInsideAngle, onUpdate)
	form.AppendSlider(stager.sliderStage, formGroup, "Stack Height", 0, 20, 1, &plant.StackHeight, onUpdate)

	if plant.CurrentView == VIEW_PLANT_3D {
		if transparency := stager.getActivePlant3DCylinderTransparency(plant); transparency != nil {
			form.AppendSlider(stager.sliderStage, formGroup, "Cylinder Transparency", 0.0, 1.0, 0.05, transparency, onUpdate)
		}
	}

	if plant.CurrentView != VIEW_PLANT_2D && (plant.PlantType == TubeVase || plant.PlantType == VaseTrapeze) && plant.TubeVaseAbstract != nil {
		form.StartAccordion(stager.sliderStage, formGroup, "Vase Core", true)

		form.AppendSlider(stager.sliderStage, formGroup, "Thickness", 0.01, 0.3, 0.01, &plant.TubeVaseAbstract.RelativeVerticalThickness, onUpdate)

		if plant.PlantType == TubeVase {
			form.AppendSlider(stager.sliderStage, formGroup, "2D Separation", 0.0, 0.1, 0.002, &plant.TubeVaseAbstract.RelativeCuttedStackFloorHeight, onUpdate)
		}

		form.AppendSlider(stager.sliderStage, formGroup, "3D Separation", 0.0, 1.0, 0.01, &plant.TubeVaseAbstract.RelativeRotatedTorusSeparation, onUpdate)
		form.AppendSlider(stager.sliderStage, formGroup, "Radial Thickness", 0.01, 0.5, 0.01, &plant.TubeVaseAbstract.RelativeRadialThickness, onUpdate)
		form.AppendSlider(stager.sliderStage, formGroup, "Rot Ratio", 0.0, 1.0, 0.005, &plant.TubeVaseAbstract.RotationRatio, onUpdate)
		form.AppendSlider(stager.sliderStage, formGroup, "Radial Repetition", 1, 4, 1, &plant.TubeVaseAbstract.RadialRepetitions, onUpdate)
		form.AppendSlider(stager.sliderStage, formGroup, "Transparency", 0.0, 1.0, 0.05, &plant.TubeVaseAbstract.Transparency, onUpdate)
		form.AppendSlider(stager.sliderStage, formGroup, "Color Combination", 1, 10, 1, &plant.TubeVaseAbstract.ColorCombination, onUpdate)

		if plant.PlantType == VaseTrapeze {
			form.StartAccordion(stager.sliderStage, formGroup, "Ribbon & Planes", true)
			form.AppendSlider(stager.sliderStage, formGroup, "Z Ribbon", -2.0, 10.0, 0.01, &plant.TubeVaseAbstract.Z_Ribbon, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Ribbon Vertical Scale", -5.0, 20.0, 0.01, &plant.TubeVaseAbstract.RibbonVerticalScale, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Top Plane", -2.0, 8.0, 0.01, &plant.TubeVaseAbstract.TopPlaneHeight, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Bottom Plane", -2.0, 8.0, 0.01, &plant.TubeVaseAbstract.BottomPlaneHeight, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Base Ring Bottom Plane", -2.0, 8.0, 0.01, &plant.TubeVaseAbstract.BaseRingBottomPlaneHeight, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Projection Angle", -90.0, 90.0, 0.1, &plant.TubeVaseAbstract.ProjectionAngle, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Base Plate Height", 0.0, 5.0, 0.01, &plant.TubeVaseAbstract.BasePlateHeight, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Carved Out Top Rings Parameter", 0.0, 1.0, 0.01, &plant.TubeVaseAbstract.CarvedOutTopRingsParameter, onUpdate)

			form.StartAccordion(stager.sliderStage, formGroup, "Bulbous, Wall & Lip", true)
			form.AppendSlider(stager.sliderStage, formGroup, "Bulbous Start Tangent Magnitude", 0.0, 10.0, 0.01, &plant.TubeVaseAbstract.BulbousStartTangentMagnitude, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Bulbous End Angle", -89.0, 89.0, 1.0, &plant.TubeVaseAbstract.BulbousEndAngle, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Bulbous End Tangent Magnitude", 0.0, 5.0, 0.01, &plant.TubeVaseAbstract.BulbousEndTangentMagnitude, onUpdate)
			form.AppendCheckBox(stager.sliderStage, formGroup, "Constant Thickness", &plant.TubeVaseAbstract.HasConstantThickness, onUpdate)
			form.AppendCheckBox(stager.sliderStage, formGroup, "Ring Bottom", &plant.TubeVaseAbstract.HasRingBottom, onUpdate)
			form.AppendCheckBox(stager.sliderStage, formGroup, "Curved Wall", &plant.TubeVaseAbstract.HasCurvedWall, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Wall Curvature", -1.0, 0.0, 0.01, &plant.TubeVaseAbstract.WallCurvature, onUpdate)
			form.AppendCheckBox(stager.sliderStage, formGroup, "Lip", &plant.TubeVaseAbstract.HasLip, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Lip Extra Height", 0.002, 0.50, 0.002, &plant.TubeVaseAbstract.LipHeight, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Lip Width", 0.01, 1.0, 0.01, &plant.TubeVaseAbstract.LipWidth, onUpdate)

			form.StartAccordion(stager.sliderStage, formGroup, "Fill Line & Drainage", true)
			form.AppendCheckBox(stager.sliderStage, formGroup, "Fill Level Line", &plant.TubeVaseAbstract.HasFillLevelLine, onUpdate)
			form.AppendCheckBox(stager.sliderStage, formGroup, "Drainage Hole", &plant.TubeVaseAbstract.HasDrainageHole, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Drainage Hole Radius", 0.02, 0.30, 0.005, &plant.TubeVaseAbstract.DrainageHoleRadius, onUpdate)
			form.AppendCheckBox(stager.sliderStage, formGroup, "Foot Rim", &plant.TubeVaseAbstract.HasFootRim, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Foot Rim Height", 0.005, 0.15, 0.005, &plant.TubeVaseAbstract.FootRimHeight, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Foot Rim Width", 0.01, 0.20, 0.005, &plant.TubeVaseAbstract.FootRimWidth, onUpdate)
		}

		if plant.PlantType == TubeVase {
			form.StartAccordion(stager.sliderStage, formGroup, "Trajectory, Key & Movie", true)
			form.AppendSlider(stager.sliderStage, formGroup, "Traj offset X", -0.15, 0.15, 0.001, &plant.TubeVaseAbstract.RelativeTrajectoryOffsetX, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Traj offset Y", -0.15, 0.15, 0.001, &plant.TubeVaseAbstract.RelativeTrajectoryOffsetY, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Nb Step P1 P2", 1, 30, 1, &plant.TubeVaseAbstract.NbStepP1P2, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Chosen Step", 1, plant.TubeVaseAbstract.NbStepP1P2, 1, &plant.TubeVaseAbstract.ChosenStep, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Horiz Ring Height", 0.0, 1.0, 0.005, &plant.TubeVaseAbstract.RelativeHorizontalRingsHeight, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Offset Key X", -500, 500, 1, &plant.TubeVaseAbstract.OffsetKeyX, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Offset Key Y", -500, 500, 1, &plant.TubeVaseAbstract.OffsetKeyY, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Width Key", 0, 500, 1, &plant.TubeVaseAbstract.WidthKey, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Height Key", 0, 500, 1, &plant.TubeVaseAbstract.HeightKey, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Key Size Reduction", 0.0, 1.0, 0.01, &plant.TubeVaseAbstract.RelativeKeySize, onUpdate)
			form.AppendSlider(stager.sliderStage, formGroup, "Movie Nb Frames", 0, 1000, 1, &plant.TubeVaseAbstract.MovieNbFrames, onUpdate)
		}

		if plant.CurrentView == VIEW_TUBE_VASE_3D || plant.CurrentView == VIEW_VASE_TRAPEZE_3D {
			if diagram := getActiveTubeVase3DDiagram(plant); diagram != nil {
				onUpdateLighting := func() {
					stager.UpdateAmbientLight()
					stager.stage.CommitWithSuspendedCallbacks()
				}
				form.StartAccordion(stager.sliderStage, formGroup, "Lighting", true)
				form.AppendSlider(stager.sliderStage, formGroup, "Ambient Light", 0.05, 3.0, 0.05, &diagram.AmbientLightLevel, onUpdateLighting)

				onUpdateCutoff := func() {
					stager.UpdateCutoffPlane()
					stager.stage.CommitWithSuspendedCallbacks()
				}
				axis := diagram.CutoffAxis
				if axis != CutoffAxisX && axis != CutoffAxisY && axis != CutoffAxisZ {
					axis = CutoffAxisX
				}

				form.StartAccordion(stager.sliderStage, formGroup, "Cutoff Plane", true)
				form.AppendCheckBox(stager.sliderStage, formGroup, "Enable Cutoff", &diagram.IsCutoffEnabled, onUpdateCutoff)
				form.AppendCheckBox(stager.sliderStage, formGroup, "Show Cutoff Plane", &diagram.IsCutoffHelperVisible, onUpdateCutoff)
				form.AppendCheckBox(stager.sliderStage, formGroup, "Flip Cutoff", &diagram.IsCutoffFlipped, onUpdateCutoff)
				form.AppendSelect(
					stager.sliderStage,
					formGroup,
					"Cutoff Axis",
					"Cutoff Axis",
					axis.CodeValues(),
					axis.ToString(),
					func(selected string) {
						var newAxis CutoffAxis
						if err := (&newAxis).FromString(selected); err == nil {
							if diagram.CutoffAxis != newAxis {
								diagram.CutoffAxis = newAxis
								onUpdateCutoff()
							}
						}
					},
				)
				form.AppendSlider(stager.sliderStage, formGroup, "Cutoff Offset", -1.0, 1.0, 0.01, &diagram.CutoffOffset, onUpdateCutoff)
			}
		}
	}

	form.EndAccordion(stager.sliderStage, formGroup)

	stager.sliderStage.Commit()
}

func appendCommonCylinderSliders(
	stage *form.Stage,
	formGroup *form.FormGroup,
	onUpdate func(),
	plant *PlantAbstract,
	radialReps *int,
	transparency *float64,
	relTubeDiam *float64,
	relHeight3DTorus *float64,
	torusVertScale *float64,
	relHeight *float64,
	projAngle *float64,
	projAngleMax float64,
) {
	form.AppendSlider(stage, formGroup, "N", 1, 20, 1, &plant.N, onUpdate)
	form.AppendSlider(stage, formGroup, "M", 1, 20, 1, &plant.M, onUpdate)
	form.AppendSlider(stage, formGroup, "Side Length", 5, 600, 5, &plant.RhombusSideLength, onUpdate)
	form.AppendSlider(stage, formGroup, "Inside Angle", 0, 180, 1, &plant.RhombusInsideAngle, onUpdate)
	form.AppendSlider(stage, formGroup, "Stack Height", 0, 20, 1, &plant.StackHeight, onUpdate)
	form.AppendSlider(stage, formGroup, "Radial Repetition", 1, 10, 1, radialReps, onUpdate)
	form.AppendSlider(stage, formGroup, "Transparency", 0.0, 1.0, 0.05, transparency, onUpdate)
	form.AppendSlider(stage, formGroup, "Tube Rel Diameter", 0.005, 0.1, 0.001, relTubeDiam, onUpdate)
	form.AppendSlider(stage, formGroup, "Rel Height 3D Torus", 0.0, 6.0, 0.01, relHeight3DTorus, onUpdate)
	form.AppendSlider(stage, formGroup, "Torus Vert Scale", 0.0, 5.0, 0.01, torusVertScale, onUpdate)
	form.AppendSlider(stage, formGroup, "Rel Height", 0.0, 6.0, 0.01, relHeight, onUpdate)
	form.AppendSlider(stage, formGroup, "Projection Angle", 0.0, projAngleMax, 0.1, projAngle, onUpdate)
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

	formGroup := form.NewLiveFormGroup(stager.sliderStoolStage, "Stool Controls")
	onUpdate := stager.OnAfterUpdateSliderElement

	form.StartAccordion(stager.sliderStoolStage, formGroup, "Plant Lattice & Cylinder", true)
	appendCommonCylinderSliders(
		stager.sliderStoolStage,
		formGroup,
		onUpdate,
		plant,
		&plant.StoolAbstract.RadialRepetitions,
		&plant.StoolAbstract.Transparency,
		&plant.StoolAbstract.RelativeTubeDiameter,
		&plant.StoolAbstract.RelativeHeight3DTorus,
		&plant.StoolAbstract.StoolTorusVerticalScale,
		&plant.StoolAbstract.RelativeHeight,
		&plant.StoolAbstract.ProjectionAngle,
		20.0,
	)

	form.StartAccordion(stager.sliderStoolStage, formGroup, "Seat & Eye", true)
	form.AppendCheckBox(stager.sliderStoolStage, formGroup, "Grasshopper Mode", &plant.StoolAbstract.IsGrasshopperMode, onUpdate)
	if plant.StoolAbstract.IsGrasshopperMode {
		form.AppendSlider(stager.sliderStoolStage, formGroup, "Leg Fillet", 0.0, 1.5, 0.01, &plant.StoolAbstract.LegFilletRadius, onUpdate)
		form.AppendSlider(stager.sliderStoolStage, formGroup, "Inner Leg Offset", 0.0, 1.0, 0.01, &plant.StoolAbstract.InnerLegOffset, onUpdate)
		form.AppendSlider(stager.sliderStoolStage, formGroup, "Seat Top Fillet", 0.0, 2.0, 0.02, &plant.StoolAbstract.SeatTopFilletRadius, onUpdate)
		form.AppendCheckBox(stager.sliderStoolStage, formGroup, "Legs Reach Top", &plant.StoolAbstract.IsLegsReachTop, onUpdate)
	} else {
		form.AppendSlider(stager.sliderStoolStage, formGroup, "Rel Seat Thickness", 0.01, 1.0, 0.01, &plant.StoolAbstract.RelativeSeatThickness, onUpdate)
		form.AppendSlider(stager.sliderStoolStage, formGroup, "Rel Eye Separation", 0.0, 0.5, 0.005, &plant.StoolAbstract.RelativeEyeSeparationCriteria, onUpdate)
		form.AppendSlider(stager.sliderStoolStage, formGroup, "Rel Eye Corner Strength", 0.0, 2.0, 0.01, &plant.StoolAbstract.RelativeEyeCornerControlVectorStrength, onUpdate)
	}
	form.EndAccordion(stager.sliderStoolStage, formGroup)

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

	formGroup := form.NewLiveFormGroup(stager.sliderClockStage, "Clock Controls")
	form.StartAccordion(stager.sliderClockStage, formGroup, "Plant Lattice & Cylinder", true)
	appendCommonCylinderSliders(
		stager.sliderClockStage,
		formGroup,
		stager.OnAfterUpdateSliderElement,
		plant,
		&plant.ClockAbstract.RadialRepetitions,
		&plant.ClockAbstract.Transparency,
		&plant.ClockAbstract.RelativeTubeDiameter,
		&plant.ClockAbstract.RelativeHeight3DTorus,
		&plant.ClockAbstract.ClockTorusVerticalScale,
		&plant.ClockAbstract.RelativeHeight,
		&plant.ClockAbstract.ProjectionAngle,
		45.0,
	)
	form.EndAccordion(stager.sliderClockStage, formGroup)

	stager.sliderClockStage.Commit()
}

func (stager *Stager) OnAfterUpdateSliderElement() {

	stager.enforceSemantic()
	stager.ux_tree()
	stager.ux_svg_plant_diagram()
	stager.ux_svg_music()
	stager.ux_tone_music()
	stager.ux_cursor_music()
	stager.UpdateThreeJSStage()
	stager.UpdateStool3DStage()
	stager.UpdateClock3DStage()
	stager.UpdatePlant3DStage()

	stager.stage.CommitWithSuspendedCallbacks()
	stager.button()
}

func getActive2DDiagramZoom(plant *PlantAbstract) *float64 {
	if plant == nil {
		return nil
	}

	if plant.CurrentView == VIEW_VASE_2D || plant.CurrentView == VIEW_TUBE_VASE_2D || plant.CurrentView == VIEW_VASE_TRAPEZE_2D {
		for _, d := range plant.Vase2DDiagrams {
			if d.IsChecked {
				return &d.Zoom
			}
		}
		if len(plant.Vase2DDiagrams) > 0 {
			return &plant.Vase2DDiagrams[0].Zoom
		}
	}

	if plant.PlantType == Stool {
		for _, d := range plant.Stool2DDiagrams {
			if d.IsChecked {
				return &d.Zoom
			}
		}
	}

	if plant.PlantType == Clock {
		for _, d := range plant.Clock2DDiagrams {
			if d.IsChecked {
				return &d.Zoom
			}
		}
	}

	for _, d := range plant.Plant2DDiagrams {
		if d.IsChecked {
			return &d.Zoom
		}
	}
	if len(plant.Plant2DDiagrams) > 0 {
		return &plant.Plant2DDiagrams[0].Zoom
	}
	if len(plant.Vase2DDiagrams) > 0 {
		return &plant.Vase2DDiagrams[0].Zoom
	}
	if len(plant.Stool2DDiagrams) > 0 {
		return &plant.Stool2DDiagrams[0].Zoom
	}
	if len(plant.Clock2DDiagrams) > 0 {
		return &plant.Clock2DDiagrams[0].Zoom
	}

	return nil
}

func (stager *Stager) getActive2DDiagramZoom(plant *PlantAbstract) *float64 {
	return getActive2DDiagramZoom(plant)
}

func GetPlant2DZoom(plant *PlantAbstract) float64 {
	if plant == nil {
		return 1.0
	}
	zPtr := getActive2DDiagramZoom(plant)
	if zPtr != nil && *zPtr > 0 {
		return *zPtr
	}
	return 1.0
}

func (stager *Stager) GetPlant2DZoom(plant *PlantAbstract) float64 {
	return GetPlant2DZoom(plant)
}

func getActivePlant3DCylinderTransparency(plant *PlantAbstract) *float64 {
	if plant == nil {
		return nil
	}
	for _, d := range plant.Plant3DDiagrams {
		if d.IsChecked && d.StemCylinder3DShape != nil {
			return &d.StemCylinder3DShape.Transparency
		}
	}
	if len(plant.Plant3DDiagrams) > 0 && plant.Plant3DDiagrams[0].StemCylinder3DShape != nil {
		return &plant.Plant3DDiagrams[0].StemCylinder3DShape.Transparency
	}
	return nil
}

func (stager *Stager) getActivePlant3DCylinderTransparency(plant *PlantAbstract) *float64 {
	return getActivePlant3DCylinderTransparency(plant)
}

func getActiveTubeVase3DDiagram(plant *PlantAbstract) *TubeVase3DDiagram {
	if plant == nil {
		return nil
	}
	for _, d := range plant.TubeVase3DDiagrams {
		if d.IsChecked {
			return d
		}
	}
	if len(plant.TubeVase3DDiagrams) > 0 {
		return plant.TubeVase3DDiagrams[0]
	}
	return nil
}
