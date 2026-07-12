// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.AxesShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LengthX", instanceWithInferedType.LengthX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LengthY", instanceWithInferedType.LengthY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsWithHiddenHandle", instanceWithInferedType.IsWithHiddenHandle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.CircleGridShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.EndArcShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StartX", instanceWithInferedType.StartX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StartY", instanceWithInferedType.StartY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndX", instanceWithInferedType.EndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndY", instanceWithInferedType.EndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("XAxisRotation", instanceWithInferedType.XAxisRotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LargeArcFlag", instanceWithInferedType.LargeArcFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SweepFlag", instanceWithInferedType.SweepFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RadiusX", instanceWithInferedType.RadiusX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RadiusY", instanceWithInferedType.RadiusY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.EndArcShapeGrid, *models.EndArcShape](
				"EndArcShapeGrid",
				"EndArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.EndArcShapeGrid) []*models.EndArcShape {
					return owner.EndArcShapes
				})
		}

	case *models.EndArcShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("EndArcShapes", instanceWithInferedType, &instanceWithInferedType.EndArcShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.ExplanationTextShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.GridPathShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.GrowthCurveBezierShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StartX", instanceWithInferedType.StartX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StartY", instanceWithInferedType.StartY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ControlPointStartX", instanceWithInferedType.ControlPointStartX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ControlPointStartY", instanceWithInferedType.ControlPointStartY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndX", instanceWithInferedType.EndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndY", instanceWithInferedType.EndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ControlPointEndX", instanceWithInferedType.ControlPointEndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ControlPointEndY", instanceWithInferedType.ControlPointEndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.GrowthCurveBezierShapeGrid, *models.GrowthCurveBezierShape](
				"GrowthCurveBezierShapeGrid",
				"GrowthCurveBezierShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.GrowthCurveBezierShapeGrid) []*models.GrowthCurveBezierShape {
					return owner.GrowthCurveBezierShapes
				})
		}

	case *models.GrowthCurveBezierShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("GrowthCurveBezierShapes", instanceWithInferedType, &instanceWithInferedType.GrowthCurveBezierShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.GrowthCurveRhombusGridShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("GrowthCurveRhombusShapes", instanceWithInferedType, &instanceWithInferedType.GrowthCurveRhombusShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.GrowthCurveRhombusShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.GrowthCurveRhombusGridShape, *models.GrowthCurveRhombusShape](
				"GrowthCurveRhombusGridShape",
				"GrowthCurveRhombusShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.GrowthCurveRhombusGridShape) []*models.GrowthCurveRhombusShape {
					return owner.GrowthCurveRhombusShapes
				})
		}

	case *models.GrowthVectorShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.InitialRhombusGridShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("InitialRhombusShapes", instanceWithInferedType, &instanceWithInferedType.InitialRhombusShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.InitialRhombusShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.InitialRhombusGridShape, *models.InitialRhombusShape](
				"InitialRhombusGridShape",
				"InitialRhombusShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.InitialRhombusGridShape) []*models.InitialRhombusShape {
					return owner.InitialRhombusShapes
				})
		}

	case *models.Library:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SubLibraries", instanceWithInferedType, &instanceWithInferedType.SubLibraries, formGroup, probe)
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LogoSVGFile", instanceWithInferedType.LogoSVGFile, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsRootLibrary", instanceWithInferedType.IsRootLibrary, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Plants", instanceWithInferedType, &instanceWithInferedType.Plants, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Library, *models.Library](
				"Library",
				"SubLibraries",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Library {
					return owner.SubLibraries
				})
		}

	case *models.NextCircleShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.PerpendicularVector:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StartX", instanceWithInferedType.StartX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StartY", instanceWithInferedType.StartY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndX", instanceWithInferedType.EndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndY", instanceWithInferedType.EndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.PerpendicularVectorGrid, *models.PerpendicularVector](
				"PerpendicularVectorGrid",
				"PerpendicularVectors",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.PerpendicularVectorGrid) []*models.PerpendicularVector {
					return owner.PerpendicularVectors
				})
		}

	case *models.PerpendicularVectorGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("PerpendicularVectors", instanceWithInferedType, &instanceWithInferedType.PerpendicularVectors, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.PerpendicularVectorGridHalfway:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("PerpendicularVectorHalfways", instanceWithInferedType, &instanceWithInferedType.PerpendicularVectorHalfways, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.PerpendicularVectorHalfway:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StartX", instanceWithInferedType.StartX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StartY", instanceWithInferedType.StartY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndX", instanceWithInferedType.EndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndY", instanceWithInferedType.EndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.PerpendicularVectorGridHalfway, *models.PerpendicularVectorHalfway](
				"PerpendicularVectorGridHalfway",
				"PerpendicularVectorHalfways",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.PerpendicularVectorGridHalfway) []*models.PerpendicularVectorHalfway {
					return owner.PerpendicularVectorHalfways
				})
		}

	case *models.Plant:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("N", instanceWithInferedType.N, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("M", instanceWithInferedType.M, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StackHeight", instanceWithInferedType.StackHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RhombusInsideAngle", instanceWithInferedType.RhombusInsideAngle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RhombusSideLength", instanceWithInferedType.RhombusSideLength, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ShiftToNearestCircle", instanceWithInferedType.ShiftToNearestCircle, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsPlantDiagramsNodeExpanded", instanceWithInferedType.IsPlantDiagramsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("PlantDiagrams", instanceWithInferedType, &instanceWithInferedType.PlantDiagrams, formGroup, probe)
		AssociationFieldToForm("AxesShape", instanceWithInferedType.AxesShape, formGroup, probe)
		AssociationFieldToForm("ReferenceRhombus", instanceWithInferedType.ReferenceRhombus, formGroup, probe)
		AssociationFieldToForm("PlantCircumferenceShape", instanceWithInferedType.PlantCircumferenceShape, formGroup, probe)
		AssociationFieldToForm("GridPathShape", instanceWithInferedType.GridPathShape, formGroup, probe)
		AssociationFieldToForm("InitialRhombusGridShape", instanceWithInferedType.InitialRhombusGridShape, formGroup, probe)
		AssociationFieldToForm("ExplanationTextShape", instanceWithInferedType.ExplanationTextShape, formGroup, probe)
		AssociationFieldToForm("RotatedReferenceRhombus", instanceWithInferedType.RotatedReferenceRhombus, formGroup, probe)
		AssociationFieldToForm("RotatedPlantCircumferenceShape", instanceWithInferedType.RotatedPlantCircumferenceShape, formGroup, probe)
		AssociationFieldToForm("RotatedGridPathShape", instanceWithInferedType.RotatedGridPathShape, formGroup, probe)
		AssociationFieldToForm("RotatedRhombusGridShape2", instanceWithInferedType.RotatedRhombusGridShape2, formGroup, probe)
		AssociationFieldToForm("GrowthCurveRhombusGridShape", instanceWithInferedType.GrowthCurveRhombusGridShape, formGroup, probe)
		AssociationFieldToForm("GrowthVectorShape", instanceWithInferedType.GrowthVectorShape, formGroup, probe)
		AssociationFieldToForm("PerpendicularVectorGrid", instanceWithInferedType.PerpendicularVectorGrid, formGroup, probe)
		AssociationFieldToForm("PerpendicularVectorGridHalfway", instanceWithInferedType.PerpendicularVectorGridHalfway, formGroup, probe)
		AssociationFieldToForm("StartArcShapeGrid", instanceWithInferedType.StartArcShapeGrid, formGroup, probe)
		AssociationFieldToForm("EndArcShapeGrid", instanceWithInferedType.EndArcShapeGrid, formGroup, probe)
		AssociationFieldToForm("GrowthCurveBezierShapeGrid", instanceWithInferedType.GrowthCurveBezierShapeGrid, formGroup, probe)
		AssociationFieldToForm("StackOfGrowthCurve", instanceWithInferedType.StackOfGrowthCurve, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Library, *models.Plant](
				"Library",
				"Plants",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Library) []*models.Plant {
					return owner.Plants
				})
		}

	case *models.PlantCircumferenceShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("AngleDegree", instanceWithInferedType.AngleDegree, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Length", instanceWithInferedType.Length, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.PlantDiagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("OriginX", instanceWithInferedType.OriginX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("OriginY", instanceWithInferedType.OriginY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenAxesShape", instanceWithInferedType.IsHiddenAxesShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenReferenceRhombus", instanceWithInferedType.IsHiddenReferenceRhombus, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenPlantCircumferenceShape", instanceWithInferedType.IsHiddenPlantCircumferenceShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenGridPathShape", instanceWithInferedType.IsHiddenGridPathShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenRhombusGridShape", instanceWithInferedType.IsHiddenRhombusGridShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenExplanationTextShape", instanceWithInferedType.IsHiddenExplanationTextShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenRotatedReferenceRhombus", instanceWithInferedType.IsHiddenRotatedReferenceRhombus, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenRotatedPlantCircumferenceShape", instanceWithInferedType.IsHiddenRotatedPlantCircumferenceShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenRotatedGridPathShape", instanceWithInferedType.IsHiddenRotatedGridPathShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenRotatedRhombusGridShape", instanceWithInferedType.IsHiddenRotatedRhombusGridShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenGrowthPathRhombusGridShape", instanceWithInferedType.IsHiddenGrowthPathRhombusGridShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenGrowthVectorShape", instanceWithInferedType.IsHiddenGrowthVectorShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenPerpendicularVectorGrid", instanceWithInferedType.IsHiddenPerpendicularVectorGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenPerpendicularVectorGridHalfway", instanceWithInferedType.IsHiddenPerpendicularVectorGridHalfway, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenStartArcShapeGrid", instanceWithInferedType.IsHiddenStartArcShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenEndArcShapeGrid", instanceWithInferedType.IsHiddenEndArcShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenGrowthCurveBezierShapeGrid", instanceWithInferedType.IsHiddenGrowthCurveBezierShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenStackOfGrowthCurve", instanceWithInferedType.IsHiddenStackOfGrowthCurve, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Rendered3DShape", instanceWithInferedType.Rendered3DShape, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Plant, *models.PlantDiagram](
				"Plant",
				"PlantDiagrams",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Plant) []*models.PlantDiagram {
					return owner.PlantDiagrams
				})
		}

	case *models.Rendered3DShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ViewX", instanceWithInferedType.ViewX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ViewY", instanceWithInferedType.ViewY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ViewZ", instanceWithInferedType.ViewZ, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("TargetX", instanceWithInferedType.TargetX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("TargetY", instanceWithInferedType.TargetY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("TargetZ", instanceWithInferedType.TargetZ, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Fov", instanceWithInferedType.Fov, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.RhombusShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.RotatedRhombusGridShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("RotatedRhombusShapes", instanceWithInferedType, &instanceWithInferedType.RotatedRhombusShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.RotatedRhombusShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.RotatedRhombusGridShape, *models.RotatedRhombusShape](
				"RotatedRhombusGridShape",
				"RotatedRhombusShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.RotatedRhombusGridShape) []*models.RotatedRhombusShape {
					return owner.RotatedRhombusShapes
				})
		}

	case *models.StackGrowthCurveBezierShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StartX", instanceWithInferedType.StartX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StartY", instanceWithInferedType.StartY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ControlPointStartX", instanceWithInferedType.ControlPointStartX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ControlPointStartY", instanceWithInferedType.ControlPointStartY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndX", instanceWithInferedType.EndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndY", instanceWithInferedType.EndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ControlPointEndX", instanceWithInferedType.ControlPointEndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ControlPointEndY", instanceWithInferedType.ControlPointEndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.StackOfGrowthCurve, *models.StackGrowthCurveBezierShape](
				"StackOfGrowthCurve",
				"StackGrowthCurveBezierShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StackOfGrowthCurve) []*models.StackGrowthCurveBezierShape {
					return owner.StackGrowthCurveBezierShapes
				})
		}

	case *models.StackOfGrowthCurve:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("StackGrowthCurveBezierShapes", instanceWithInferedType, &instanceWithInferedType.StackGrowthCurveBezierShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.StartArcShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StartX", instanceWithInferedType.StartX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StartY", instanceWithInferedType.StartY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndX", instanceWithInferedType.EndX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EndY", instanceWithInferedType.EndY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("XAxisRotation", instanceWithInferedType.XAxisRotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LargeArcFlag", instanceWithInferedType.LargeArcFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SweepFlag", instanceWithInferedType.SweepFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RadiusX", instanceWithInferedType.RadiusX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RadiusY", instanceWithInferedType.RadiusY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.StartArcShapeGrid, *models.StartArcShape](
				"StartArcShapeGrid",
				"StartArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StartArcShapeGrid) []*models.StartArcShape {
					return owner.StartArcShapes
				})
		}

	case *models.StartArcShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("StartArcShapes", instanceWithInferedType, &instanceWithInferedType.StartArcShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	default:
		_ = instanceWithInferedType
	}
}
