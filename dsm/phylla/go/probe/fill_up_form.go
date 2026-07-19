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
	case *models.ArcNormalVectorShape:
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
			AssociationReverseSliceToForm[*models.ArcNormalVectorShapeGrid, *models.ArcNormalVectorShape](
				"ArcNormalVectorShapeGrid",
				"ArcNormalVectorShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ArcNormalVectorShapeGrid) []*models.ArcNormalVectorShape {
					return owner.ArcNormalVectorShapes
				})
		}

	case *models.ArcNormalVectorShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ArcNormalVectorShapes", instanceWithInferedType, &instanceWithInferedType.ArcNormalVectorShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

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

	case *models.BaseVectorShape:
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
			AssociationReverseSliceToForm[*models.BaseVectorShapeGrid, *models.BaseVectorShape](
				"BaseVectorShapeGrid",
				"BaseVectorShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.BaseVectorShapeGrid) []*models.BaseVectorShape {
					return owner.BaseVectorShapes
				})
		}

	case *models.BaseVectorShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("BaseVectorShapes", instanceWithInferedType, &instanceWithInferedType.BaseVectorShapes, formGroup, probe)
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

	case *models.EndHalfwayArcShape:
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
		BasicFieldtoForm("RadiusX", instanceWithInferedType.RadiusX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RadiusY", instanceWithInferedType.RadiusY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("XAxisRotation", instanceWithInferedType.XAxisRotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LargeArcFlag", instanceWithInferedType.LargeArcFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SweepFlag", instanceWithInferedType.SweepFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.EndHalfwayArcShapeGrid, *models.EndHalfwayArcShape](
				"EndHalfwayArcShapeGrid",
				"EndHalfwayArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.EndHalfwayArcShapeGrid) []*models.EndHalfwayArcShape {
					return owner.EndHalfwayArcShapes
				})
		}

	case *models.EndHalfwayArcShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("EndHalfwayArcShapes", instanceWithInferedType, &instanceWithInferedType.EndHalfwayArcShapes, formGroup, probe)
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

	case *models.GrowthCurve2D:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("StartHalfwayArcShapeGrid", instanceWithInferedType.StartHalfwayArcShapeGrid, formGroup, probe)
		AssociationFieldToForm("EndHalfwayArcShapeGrid", instanceWithInferedType.EndHalfwayArcShapeGrid, formGroup, probe)
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

	case *models.MidArcVectorShape:
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
			AssociationReverseSliceToForm[*models.MidArcVectorShapeGrid, *models.MidArcVectorShape](
				"MidArcVectorShapeGrid",
				"MidArcVectorShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.MidArcVectorShapeGrid) []*models.MidArcVectorShape {
					return owner.MidArcVectorShapes
				})
		}

	case *models.MidArcVectorShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("MidArcVectorShapes", instanceWithInferedType, &instanceWithInferedType.MidArcVectorShapes, formGroup, probe)
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
		BasicFieldtoForm("RelativeVerticalThickness", instanceWithInferedType.RelativeVerticalThickness, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RadialThickness", instanceWithInferedType.RadialThickness, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RhombusSideLength", instanceWithInferedType.RhombusSideLength, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("CuttedStackFloorHeight", instanceWithInferedType.CuttedStackFloorHeight, instanceWithInferedType, probe.formStage, formGroup,
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
		AssociationFieldToForm("RhombusStuff", instanceWithInferedType.RhombusStuff, formGroup, probe)
		AssociationFieldToForm("GrowthVectorShape", instanceWithInferedType.GrowthVectorShape, formGroup, probe)
		AssociationFieldToForm("PerpendicularVectorGrid", instanceWithInferedType.PerpendicularVectorGrid, formGroup, probe)
		AssociationFieldToForm("PerpendicularVectorGridHalfway", instanceWithInferedType.PerpendicularVectorGridHalfway, formGroup, probe)
		AssociationFieldToForm("BaseVectorShapeGrid", instanceWithInferedType.BaseVectorShapeGrid, formGroup, probe)
		AssociationFieldToForm("ArcNormalVectorShapeGrid", instanceWithInferedType.ArcNormalVectorShapeGrid, formGroup, probe)
		AssociationFieldToForm("StartArcShapeGrid", instanceWithInferedType.StartArcShapeGrid, formGroup, probe)
		AssociationFieldToForm("TopStartArcShapeGrid", instanceWithInferedType.TopStartArcShapeGrid, formGroup, probe)
		AssociationFieldToForm("EndArcShapeGrid", instanceWithInferedType.EndArcShapeGrid, formGroup, probe)
		AssociationFieldToForm("TopEndArcShapeGrid", instanceWithInferedType.TopEndArcShapeGrid, formGroup, probe)
		AssociationFieldToForm("ShiftedBottomTopStartArcShapeGrid", instanceWithInferedType.ShiftedBottomTopStartArcShapeGrid, formGroup, probe)
		AssociationFieldToForm("MidArcVectorShapeGrid", instanceWithInferedType.MidArcVectorShapeGrid, formGroup, probe)
		AssociationFieldToForm("TopMidArcVectorShapeGrid", instanceWithInferedType.TopMidArcVectorShapeGrid, formGroup, probe)
		AssociationFieldToForm("StartHalfwayArcShapeGrid", instanceWithInferedType.StartHalfwayArcShapeGrid, formGroup, probe)
		AssociationFieldToForm("TopStartHalfwayArcShapeGrid", instanceWithInferedType.TopStartHalfwayArcShapeGrid, formGroup, probe)
		AssociationFieldToForm("EndHalfwayArcShapeGrid", instanceWithInferedType.EndHalfwayArcShapeGrid, formGroup, probe)
		AssociationFieldToForm("TopEndHalfwayArcShapeGrid", instanceWithInferedType.TopEndHalfwayArcShapeGrid, formGroup, probe)
		AssociationFieldToForm("StackOfGrowthCurve", instanceWithInferedType.StackOfGrowthCurve, formGroup, probe)
		AssociationFieldToForm("TopStackOfGrowthCurve", instanceWithInferedType.TopStackOfGrowthCurve, formGroup, probe)
		AssociationFieldToForm("ShiftedLeftStackOfGrowthCurve", instanceWithInferedType.ShiftedLeftStackOfGrowthCurve, formGroup, probe)
		AssociationFieldToForm("ShiftedLeftStackOfNormalVector", instanceWithInferedType.ShiftedLeftStackOfNormalVector, formGroup, probe)
		AssociationFieldToForm("GrowthCurve2D", instanceWithInferedType.GrowthCurve2D, formGroup, probe)
		AssociationFieldToForm("TopGrowthCurve2D", instanceWithInferedType.TopGrowthCurve2D, formGroup, probe)
		AssociationFieldToForm("StackOfGrowthCurve2D", instanceWithInferedType.StackOfGrowthCurve2D, formGroup, probe)
		AssociationFieldToForm("TopStackOfGrowthCurve2D", instanceWithInferedType.TopStackOfGrowthCurve2D, formGroup, probe)
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
		BasicFieldtoForm("IsRhombusNodesExpanded", instanceWithInferedType.IsRhombusNodesExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsArcNodesExpanded", instanceWithInferedType.IsArcNodesExpanded, instanceWithInferedType, probe.formStage, formGroup,
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
		BasicFieldtoForm("IsHiddenBaseVectorShapeGrid", instanceWithInferedType.IsHiddenBaseVectorShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenArcNormalVectorShapeGrid", instanceWithInferedType.IsHiddenArcNormalVectorShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenStartArcShapeGrid", instanceWithInferedType.IsHiddenStartArcShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenTopStartArcShapeGrid", instanceWithInferedType.IsHiddenTopStartArcShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenShiftedBottomTopStartArcShapeGrid", instanceWithInferedType.IsHiddenShiftedBottomTopStartArcShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenMidArcVectorShapeGrid", instanceWithInferedType.IsHiddenMidArcVectorShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenTopMidArcVectorShapeGrid", instanceWithInferedType.IsHiddenTopMidArcVectorShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenStartHalfwayArcShapeGrid", instanceWithInferedType.IsHiddenStartHalfwayArcShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenTopStartHalfwayArcShapeGrid", instanceWithInferedType.IsHiddenTopStartHalfwayArcShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenEndHalfwayArcShapeGrid", instanceWithInferedType.IsHiddenEndHalfwayArcShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenTopEndHalfwayArcShapeGrid", instanceWithInferedType.IsHiddenTopEndHalfwayArcShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenEndArcShapeGrid", instanceWithInferedType.IsHiddenEndArcShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenTopEndArcShapeGrid", instanceWithInferedType.IsHiddenTopEndArcShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenBottomStartArcShapeGrid", instanceWithInferedType.IsHiddenBottomStartArcShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenBottomEndArcShapeGrid", instanceWithInferedType.IsHiddenBottomEndArcShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenStackOfGrowthCurve", instanceWithInferedType.IsHiddenStackOfGrowthCurve, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenTopStackOfGrowthCurve", instanceWithInferedType.IsHiddenTopStackOfGrowthCurve, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenBottomStackOfGrowthCurve", instanceWithInferedType.IsHiddenBottomStackOfGrowthCurve, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenShiftedLeftStackOfGrowthCurve", instanceWithInferedType.IsHiddenShiftedLeftStackOfGrowthCurve, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenShiftedLeftStackOfNormalVector", instanceWithInferedType.IsHiddenShiftedLeftStackOfNormalVector, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenGrowthCurve2D", instanceWithInferedType.IsHiddenGrowthCurve2D, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenTopGrowthCurve2D", instanceWithInferedType.IsHiddenTopGrowthCurve2D, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenStackOfGrowthCurve2D", instanceWithInferedType.IsHiddenStackOfGrowthCurve2D, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenTopStackOfGrowthCurve2D", instanceWithInferedType.IsHiddenTopStackOfGrowthCurve2D, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenTorusStackShape", instanceWithInferedType.IsHiddenTorusStackShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ComputedPrefix", instanceWithInferedType.ComputedPrefix, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("Rendered3DShape", instanceWithInferedType.Rendered3DShape, formGroup, probe)
		AssociationFieldToForm("TorusStackShape", instanceWithInferedType.TorusStackShape, formGroup, probe)
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

	case *models.RhombusStuff:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
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

	case *models.ShiftedBottomTopStartArcShape:
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
			AssociationReverseSliceToForm[*models.ShiftedBottomTopStartArcShapeGrid, *models.ShiftedBottomTopStartArcShape](
				"ShiftedBottomTopStartArcShapeGrid",
				"ShiftedBottomTopStartArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ShiftedBottomTopStartArcShapeGrid) []*models.ShiftedBottomTopStartArcShape {
					return owner.ShiftedBottomTopStartArcShapes
				})
		}

	case *models.ShiftedBottomTopStartArcShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ShiftedBottomTopStartArcShapes", instanceWithInferedType, &instanceWithInferedType.ShiftedBottomTopStartArcShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.ShiftedLeftStackGrowthCurveEndArcShape:
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
			AssociationReverseSliceToForm[*models.ShiftedLeftStackOfGrowthCurve, *models.ShiftedLeftStackGrowthCurveEndArcShape](
				"ShiftedLeftStackOfGrowthCurve",
				"ShiftedLeftStackGrowthCurveEndArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ShiftedLeftStackOfGrowthCurve) []*models.ShiftedLeftStackGrowthCurveEndArcShape {
					return owner.ShiftedLeftStackGrowthCurveEndArcShapes
				})
		}

	case *models.ShiftedLeftStackGrowthCurveStartArcShape:
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
			AssociationReverseSliceToForm[*models.ShiftedLeftStackOfGrowthCurve, *models.ShiftedLeftStackGrowthCurveStartArcShape](
				"ShiftedLeftStackOfGrowthCurve",
				"ShiftedLeftStackGrowthCurveStartArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ShiftedLeftStackOfGrowthCurve) []*models.ShiftedLeftStackGrowthCurveStartArcShape {
					return owner.ShiftedLeftStackGrowthCurveStartArcShapes
				})
		}

	case *models.ShiftedLeftStackNormalVector:
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
			AssociationReverseSliceToForm[*models.ShiftedLeftStackOfNormalVector, *models.ShiftedLeftStackNormalVector](
				"ShiftedLeftStackOfNormalVector",
				"ShiftedLeftStackNormalVectors",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.ShiftedLeftStackOfNormalVector) []*models.ShiftedLeftStackNormalVector {
					return owner.ShiftedLeftStackNormalVectors
				})
		}

	case *models.ShiftedLeftStackOfGrowthCurve:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ShiftedLeftStackGrowthCurveStartArcShapes", instanceWithInferedType, &instanceWithInferedType.ShiftedLeftStackGrowthCurveStartArcShapes, formGroup, probe)
		AssociationSliceToForm("ShiftedLeftStackGrowthCurveEndArcShapes", instanceWithInferedType, &instanceWithInferedType.ShiftedLeftStackGrowthCurveEndArcShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.ShiftedLeftStackOfNormalVector:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ShiftedLeftStackNormalVectors", instanceWithInferedType, &instanceWithInferedType.ShiftedLeftStackNormalVectors, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.StackGrowthCurve2DEndHalfwayArcShape:
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
		BasicFieldtoForm("RadiusX", instanceWithInferedType.RadiusX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RadiusY", instanceWithInferedType.RadiusY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("XAxisRotation", instanceWithInferedType.XAxisRotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LargeArcFlag", instanceWithInferedType.LargeArcFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SweepFlag", instanceWithInferedType.SweepFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.StackOfGrowthCurve2D, *models.StackGrowthCurve2DEndHalfwayArcShape](
				"StackOfGrowthCurve2D",
				"StackGrowthCurve2DEndHalfwayArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StackOfGrowthCurve2D) []*models.StackGrowthCurve2DEndHalfwayArcShape {
					return owner.StackGrowthCurve2DEndHalfwayArcShapes
				})
		}

	case *models.StackGrowthCurve2DStartHalfwayArcShape:
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
		BasicFieldtoForm("RadiusX", instanceWithInferedType.RadiusX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RadiusY", instanceWithInferedType.RadiusY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("XAxisRotation", instanceWithInferedType.XAxisRotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LargeArcFlag", instanceWithInferedType.LargeArcFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SweepFlag", instanceWithInferedType.SweepFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.StackOfGrowthCurve2D, *models.StackGrowthCurve2DStartHalfwayArcShape](
				"StackOfGrowthCurve2D",
				"StackGrowthCurve2DStartHalfwayArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StackOfGrowthCurve2D) []*models.StackGrowthCurve2DStartHalfwayArcShape {
					return owner.StackGrowthCurve2DStartHalfwayArcShapes
				})
		}

	case *models.StackGrowthCurveEndArcShape:
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
			AssociationReverseSliceToForm[*models.StackOfGrowthCurve, *models.StackGrowthCurveEndArcShape](
				"StackOfGrowthCurve",
				"StackGrowthCurveEndArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StackOfGrowthCurve) []*models.StackGrowthCurveEndArcShape {
					return owner.StackGrowthCurveEndArcShapes
				})
		}

	case *models.StackGrowthCurveStartArcShape:
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
			AssociationReverseSliceToForm[*models.StackOfGrowthCurve, *models.StackGrowthCurveStartArcShape](
				"StackOfGrowthCurve",
				"StackGrowthCurveStartArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StackOfGrowthCurve) []*models.StackGrowthCurveStartArcShape {
					return owner.StackGrowthCurveStartArcShapes
				})
		}

	case *models.StackOfGrowthCurve:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("StackGrowthCurveStartArcShapes", instanceWithInferedType, &instanceWithInferedType.StackGrowthCurveStartArcShapes, formGroup, probe)
		AssociationSliceToForm("StackGrowthCurveEndArcShapes", instanceWithInferedType, &instanceWithInferedType.StackGrowthCurveEndArcShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.StackOfGrowthCurve2D:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("StackGrowthCurve2DStartHalfwayArcShapes", instanceWithInferedType, &instanceWithInferedType.StackGrowthCurve2DStartHalfwayArcShapes, formGroup, probe)
		AssociationSliceToForm("StackGrowthCurve2DEndHalfwayArcShapes", instanceWithInferedType, &instanceWithInferedType.StackGrowthCurve2DEndHalfwayArcShapes, formGroup, probe)
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

	case *models.StartHalfwayArcShape:
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
		BasicFieldtoForm("RadiusX", instanceWithInferedType.RadiusX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RadiusY", instanceWithInferedType.RadiusY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("XAxisRotation", instanceWithInferedType.XAxisRotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LargeArcFlag", instanceWithInferedType.LargeArcFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SweepFlag", instanceWithInferedType.SweepFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.StartHalfwayArcShapeGrid, *models.StartHalfwayArcShape](
				"StartHalfwayArcShapeGrid",
				"StartHalfwayArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StartHalfwayArcShapeGrid) []*models.StartHalfwayArcShape {
					return owner.StartHalfwayArcShapes
				})
		}

	case *models.StartHalfwayArcShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("StartHalfwayArcShapes", instanceWithInferedType, &instanceWithInferedType.StartHalfwayArcShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TopEndArcShape:
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
			AssociationReverseSliceToForm[*models.TopEndArcShapeGrid, *models.TopEndArcShape](
				"TopEndArcShapeGrid",
				"TopEndArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TopEndArcShapeGrid) []*models.TopEndArcShape {
					return owner.TopEndArcShapes
				})
		}

	case *models.TopEndArcShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("TopEndArcShapes", instanceWithInferedType, &instanceWithInferedType.TopEndArcShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TopEndHalfwayArcShape:
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
		BasicFieldtoForm("RadiusX", instanceWithInferedType.RadiusX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RadiusY", instanceWithInferedType.RadiusY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("XAxisRotation", instanceWithInferedType.XAxisRotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LargeArcFlag", instanceWithInferedType.LargeArcFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SweepFlag", instanceWithInferedType.SweepFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.TopEndHalfwayArcShapeGrid, *models.TopEndHalfwayArcShape](
				"TopEndHalfwayArcShapeGrid",
				"TopEndHalfwayArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TopEndHalfwayArcShapeGrid) []*models.TopEndHalfwayArcShape {
					return owner.TopEndHalfwayArcShapes
				})
		}

	case *models.TopEndHalfwayArcShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("TopEndHalfwayArcShapes", instanceWithInferedType, &instanceWithInferedType.TopEndHalfwayArcShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TopGrowthCurve2D:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("TopStartHalfwayArcShapeGrid", instanceWithInferedType.TopStartHalfwayArcShapeGrid, formGroup, probe)
		AssociationFieldToForm("TopEndHalfwayArcShapeGrid", instanceWithInferedType.TopEndHalfwayArcShapeGrid, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TopMidArcVectorShape:
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
			AssociationReverseSliceToForm[*models.TopMidArcVectorShapeGrid, *models.TopMidArcVectorShape](
				"TopMidArcVectorShapeGrid",
				"TopMidArcVectorShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TopMidArcVectorShapeGrid) []*models.TopMidArcVectorShape {
					return owner.TopMidArcVectorShapes
				})
		}

	case *models.TopMidArcVectorShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("TopMidArcVectorShapes", instanceWithInferedType, &instanceWithInferedType.TopMidArcVectorShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TopStackGrowthCurve2DEndHalfwayArcShape:
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
		BasicFieldtoForm("RadiusX", instanceWithInferedType.RadiusX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RadiusY", instanceWithInferedType.RadiusY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("XAxisRotation", instanceWithInferedType.XAxisRotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LargeArcFlag", instanceWithInferedType.LargeArcFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SweepFlag", instanceWithInferedType.SweepFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.TopStackOfGrowthCurve2D, *models.TopStackGrowthCurve2DEndHalfwayArcShape](
				"TopStackOfGrowthCurve2D",
				"TopStackGrowthCurve2DEndHalfwayArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TopStackOfGrowthCurve2D) []*models.TopStackGrowthCurve2DEndHalfwayArcShape {
					return owner.TopStackGrowthCurve2DEndHalfwayArcShapes
				})
		}

	case *models.TopStackGrowthCurve2DStartHalfwayArcShape:
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
		BasicFieldtoForm("RadiusX", instanceWithInferedType.RadiusX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RadiusY", instanceWithInferedType.RadiusY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("XAxisRotation", instanceWithInferedType.XAxisRotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LargeArcFlag", instanceWithInferedType.LargeArcFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SweepFlag", instanceWithInferedType.SweepFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.TopStackOfGrowthCurve2D, *models.TopStackGrowthCurve2DStartHalfwayArcShape](
				"TopStackOfGrowthCurve2D",
				"TopStackGrowthCurve2DStartHalfwayArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TopStackOfGrowthCurve2D) []*models.TopStackGrowthCurve2DStartHalfwayArcShape {
					return owner.TopStackGrowthCurve2DStartHalfwayArcShapes
				})
		}

	case *models.TopStackGrowthCurveEndArcShape:
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
			AssociationReverseSliceToForm[*models.TopStackOfGrowthCurve, *models.TopStackGrowthCurveEndArcShape](
				"TopStackOfGrowthCurve",
				"TopStackGrowthCurveEndArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TopStackOfGrowthCurve) []*models.TopStackGrowthCurveEndArcShape {
					return owner.TopStackGrowthCurveEndArcShapes
				})
		}

	case *models.TopStackGrowthCurveStartArcShape:
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
			AssociationReverseSliceToForm[*models.TopStackOfGrowthCurve, *models.TopStackGrowthCurveStartArcShape](
				"TopStackOfGrowthCurve",
				"TopStackGrowthCurveStartArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TopStackOfGrowthCurve) []*models.TopStackGrowthCurveStartArcShape {
					return owner.TopStackGrowthCurveStartArcShapes
				})
		}

	case *models.TopStackOfGrowthCurve:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("TopStackGrowthCurveStartArcShapes", instanceWithInferedType, &instanceWithInferedType.TopStackGrowthCurveStartArcShapes, formGroup, probe)
		AssociationSliceToForm("TopStackGrowthCurveEndArcShapes", instanceWithInferedType, &instanceWithInferedType.TopStackGrowthCurveEndArcShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TopStackOfGrowthCurve2D:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("TopStackGrowthCurve2DStartHalfwayArcShapes", instanceWithInferedType, &instanceWithInferedType.TopStackGrowthCurve2DStartHalfwayArcShapes, formGroup, probe)
		AssociationSliceToForm("TopStackGrowthCurve2DEndHalfwayArcShapes", instanceWithInferedType, &instanceWithInferedType.TopStackGrowthCurve2DEndHalfwayArcShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TopStartArcShape:
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
			AssociationReverseSliceToForm[*models.TopStartArcShapeGrid, *models.TopStartArcShape](
				"TopStartArcShapeGrid",
				"TopStartArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TopStartArcShapeGrid) []*models.TopStartArcShape {
					return owner.TopStartArcShapes
				})
		}

	case *models.TopStartArcShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("TopStartArcShapes", instanceWithInferedType, &instanceWithInferedType.TopStartArcShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TopStartHalfwayArcShape:
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
		BasicFieldtoForm("RadiusX", instanceWithInferedType.RadiusX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RadiusY", instanceWithInferedType.RadiusY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("XAxisRotation", instanceWithInferedType.XAxisRotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LargeArcFlag", instanceWithInferedType.LargeArcFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SweepFlag", instanceWithInferedType.SweepFlag, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.TopStartHalfwayArcShapeGrid, *models.TopStartHalfwayArcShape](
				"TopStartHalfwayArcShapeGrid",
				"TopStartHalfwayArcShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TopStartHalfwayArcShapeGrid) []*models.TopStartHalfwayArcShape {
					return owner.TopStartHalfwayArcShapes
				})
		}

	case *models.TopStartHalfwayArcShapeGrid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("TopStartHalfwayArcShapes", instanceWithInferedType, &instanceWithInferedType.TopStartHalfwayArcShapes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TorusStackShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	default:
		_ = instanceWithInferedType
	}
}
