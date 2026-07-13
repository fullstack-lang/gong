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

	case *models.BottomEndArcShapeV2:
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
			AssociationReverseSliceToForm[*models.BottomEndArcShapeV2Grid, *models.BottomEndArcShapeV2](
				"BottomEndArcShapeV2Grid",
				"BottomEndArcShapesV2",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.BottomEndArcShapeV2Grid) []*models.BottomEndArcShapeV2 {
					return owner.BottomEndArcShapesV2
				})
		}

	case *models.BottomEndArcShapeV2Grid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("BottomEndArcShapesV2", instanceWithInferedType, &instanceWithInferedType.BottomEndArcShapesV2, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.BottomStackGrowthCurveEndArcShapeV2:
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
			AssociationReverseSliceToForm[*models.BottomStackOfGrowthCurveV2, *models.BottomStackGrowthCurveEndArcShapeV2](
				"BottomStackOfGrowthCurveV2",
				"BottomStackGrowthCurveEndArcShapeV2s",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.BottomStackOfGrowthCurveV2) []*models.BottomStackGrowthCurveEndArcShapeV2 {
					return owner.BottomStackGrowthCurveEndArcShapeV2s
				})
		}

	case *models.BottomStackGrowthCurveStartArcShapeV2:
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
			AssociationReverseSliceToForm[*models.BottomStackOfGrowthCurveV2, *models.BottomStackGrowthCurveStartArcShapeV2](
				"BottomStackOfGrowthCurveV2",
				"BottomStackGrowthCurveStartArcShapeV2s",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.BottomStackOfGrowthCurveV2) []*models.BottomStackGrowthCurveStartArcShapeV2 {
					return owner.BottomStackGrowthCurveStartArcShapeV2s
				})
		}

	case *models.BottomStackOfGrowthCurveV2:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("BottomStackGrowthCurveStartArcShapeV2s", instanceWithInferedType, &instanceWithInferedType.BottomStackGrowthCurveStartArcShapeV2s, formGroup, probe)
		AssociationSliceToForm("BottomStackGrowthCurveEndArcShapeV2s", instanceWithInferedType, &instanceWithInferedType.BottomStackGrowthCurveEndArcShapeV2s, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.BottomStartArcShapeV2:
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
			AssociationReverseSliceToForm[*models.BottomStartArcShapeV2Grid, *models.BottomStartArcShapeV2](
				"BottomStartArcShapeV2Grid",
				"BottomStartArcShapesV2",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.BottomStartArcShapeV2Grid) []*models.BottomStartArcShapeV2 {
					return owner.BottomStartArcShapesV2
				})
		}

	case *models.BottomStartArcShapeV2Grid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("BottomStartArcShapesV2", instanceWithInferedType, &instanceWithInferedType.BottomStartArcShapesV2, formGroup, probe)
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

	case *models.EndArcShapeV2:
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
			AssociationReverseSliceToForm[*models.EndArcShapeV2Grid, *models.EndArcShapeV2](
				"EndArcShapeV2Grid",
				"EndArcShapesV2",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.EndArcShapeV2Grid) []*models.EndArcShapeV2 {
					return owner.EndArcShapesV2
				})
		}

	case *models.EndArcShapeV2Grid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("EndArcShapesV2", instanceWithInferedType, &instanceWithInferedType.EndArcShapesV2, formGroup, probe)
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
		AssociationFieldToForm("StartArcShapeV2Grid", instanceWithInferedType.StartArcShapeV2Grid, formGroup, probe)
		AssociationFieldToForm("EndArcShapeV2Grid", instanceWithInferedType.EndArcShapeV2Grid, formGroup, probe)
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
		BasicFieldtoForm("Thickness", instanceWithInferedType.Thickness, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RhombusSideLength", instanceWithInferedType.RhombusSideLength, instanceWithInferedType, probe.formStage, formGroup,
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
		AssociationFieldToForm("BaseVectorShapeGrid", instanceWithInferedType.BaseVectorShapeGrid, formGroup, probe)
		AssociationFieldToForm("ArcNormalVectorShapeGrid", instanceWithInferedType.ArcNormalVectorShapeGrid, formGroup, probe)
		AssociationFieldToForm("StartArcShapeV2Grid", instanceWithInferedType.StartArcShapeV2Grid, formGroup, probe)
		AssociationFieldToForm("TopStartArcShapeV2Grid", instanceWithInferedType.TopStartArcShapeV2Grid, formGroup, probe)
		AssociationFieldToForm("EndArcShapeV2Grid", instanceWithInferedType.EndArcShapeV2Grid, formGroup, probe)
		AssociationFieldToForm("TopEndArcShapeV2Grid", instanceWithInferedType.TopEndArcShapeV2Grid, formGroup, probe)
		AssociationFieldToForm("BottomStartArcShapeV2Grid", instanceWithInferedType.BottomStartArcShapeV2Grid, formGroup, probe)
		AssociationFieldToForm("BottomEndArcShapeV2Grid", instanceWithInferedType.BottomEndArcShapeV2Grid, formGroup, probe)
		AssociationFieldToForm("GrowthCurveBezierShapeGrid", instanceWithInferedType.GrowthCurveBezierShapeGrid, formGroup, probe)
		AssociationFieldToForm("StackOfGrowthCurveV2", instanceWithInferedType.StackOfGrowthCurveV2, formGroup, probe)
		AssociationFieldToForm("TopStackOfGrowthCurveV2", instanceWithInferedType.TopStackOfGrowthCurveV2, formGroup, probe)
		AssociationFieldToForm("BottomStackOfGrowthCurveV2", instanceWithInferedType.BottomStackOfGrowthCurveV2, formGroup, probe)
		AssociationFieldToForm("GrowthCurve2D", instanceWithInferedType.GrowthCurve2D, formGroup, probe)
		AssociationFieldToForm("TopGrowthCurve2D", instanceWithInferedType.TopGrowthCurve2D, formGroup, probe)
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
		BasicFieldtoForm("IsHiddenBaseVectorShapeGrid", instanceWithInferedType.IsHiddenBaseVectorShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenArcNormalVectorShapeGrid", instanceWithInferedType.IsHiddenArcNormalVectorShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenStartArcShapeV2Grid", instanceWithInferedType.IsHiddenStartArcShapeV2Grid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenTopStartArcShapeV2Grid", instanceWithInferedType.IsHiddenTopStartArcShapeV2Grid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenEndArcShapeV2Grid", instanceWithInferedType.IsHiddenEndArcShapeV2Grid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenTopEndArcShapeV2Grid", instanceWithInferedType.IsHiddenTopEndArcShapeV2Grid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenBottomStartArcShapeV2Grid", instanceWithInferedType.IsHiddenBottomStartArcShapeV2Grid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenBottomEndArcShapeV2Grid", instanceWithInferedType.IsHiddenBottomEndArcShapeV2Grid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenGrowthCurveBezierShapeGrid", instanceWithInferedType.IsHiddenGrowthCurveBezierShapeGrid, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenStackOfGrowthCurveV2", instanceWithInferedType.IsHiddenStackOfGrowthCurveV2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenTopStackOfGrowthCurveV2", instanceWithInferedType.IsHiddenTopStackOfGrowthCurveV2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenBottomStackOfGrowthCurveV2", instanceWithInferedType.IsHiddenBottomStackOfGrowthCurveV2, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenGrowthCurve2D", instanceWithInferedType.IsHiddenGrowthCurve2D, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsHiddenTopGrowthCurve2D", instanceWithInferedType.IsHiddenTopGrowthCurve2D, instanceWithInferedType, probe.formStage, formGroup,
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

	case *models.StackGrowthCurveEndArcShapeV2:
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
			AssociationReverseSliceToForm[*models.StackOfGrowthCurveV2, *models.StackGrowthCurveEndArcShapeV2](
				"StackOfGrowthCurveV2",
				"StackGrowthCurveEndArcShapeV2s",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StackOfGrowthCurveV2) []*models.StackGrowthCurveEndArcShapeV2 {
					return owner.StackGrowthCurveEndArcShapeV2s
				})
		}

	case *models.StackGrowthCurveStartArcShapeV2:
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
			AssociationReverseSliceToForm[*models.StackOfGrowthCurveV2, *models.StackGrowthCurveStartArcShapeV2](
				"StackOfGrowthCurveV2",
				"StackGrowthCurveStartArcShapeV2s",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StackOfGrowthCurveV2) []*models.StackGrowthCurveStartArcShapeV2 {
					return owner.StackGrowthCurveStartArcShapeV2s
				})
		}

	case *models.StackOfGrowthCurveV2:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("StackGrowthCurveStartArcShapeV2s", instanceWithInferedType, &instanceWithInferedType.StackGrowthCurveStartArcShapeV2s, formGroup, probe)
		AssociationSliceToForm("StackGrowthCurveEndArcShapeV2s", instanceWithInferedType, &instanceWithInferedType.StackGrowthCurveEndArcShapeV2s, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.StartArcShapeV2:
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
			AssociationReverseSliceToForm[*models.StartArcShapeV2Grid, *models.StartArcShapeV2](
				"StartArcShapeV2Grid",
				"StartArcShapesV2",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.StartArcShapeV2Grid) []*models.StartArcShapeV2 {
					return owner.StartArcShapesV2
				})
		}

	case *models.StartArcShapeV2Grid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("StartArcShapesV2", instanceWithInferedType, &instanceWithInferedType.StartArcShapesV2, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TopEndArcShapeV2:
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
			AssociationReverseSliceToForm[*models.TopEndArcShapeV2Grid, *models.TopEndArcShapeV2](
				"TopEndArcShapeV2Grid",
				"TopEndArcShapesV2",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TopEndArcShapeV2Grid) []*models.TopEndArcShapeV2 {
					return owner.TopEndArcShapesV2
				})
		}

	case *models.TopEndArcShapeV2Grid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("TopEndArcShapesV2", instanceWithInferedType, &instanceWithInferedType.TopEndArcShapesV2, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TopGrowthCurve2D:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("TopStartArcShapeV2Grid", instanceWithInferedType.TopStartArcShapeV2Grid, formGroup, probe)
		AssociationFieldToForm("TopEndArcShapeV2Grid", instanceWithInferedType.TopEndArcShapeV2Grid, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TopStackGrowthCurveEndArcShapeV2:
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
			AssociationReverseSliceToForm[*models.TopStackOfGrowthCurveV2, *models.TopStackGrowthCurveEndArcShapeV2](
				"TopStackOfGrowthCurveV2",
				"TopStackGrowthCurveEndArcShapeV2s",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TopStackOfGrowthCurveV2) []*models.TopStackGrowthCurveEndArcShapeV2 {
					return owner.TopStackGrowthCurveEndArcShapeV2s
				})
		}

	case *models.TopStackGrowthCurveStartArcShapeV2:
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
			AssociationReverseSliceToForm[*models.TopStackOfGrowthCurveV2, *models.TopStackGrowthCurveStartArcShapeV2](
				"TopStackOfGrowthCurveV2",
				"TopStackGrowthCurveStartArcShapeV2s",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TopStackOfGrowthCurveV2) []*models.TopStackGrowthCurveStartArcShapeV2 {
					return owner.TopStackGrowthCurveStartArcShapeV2s
				})
		}

	case *models.TopStackOfGrowthCurveV2:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("TopStackGrowthCurveStartArcShapeV2s", instanceWithInferedType, &instanceWithInferedType.TopStackGrowthCurveStartArcShapeV2s, formGroup, probe)
		AssociationSliceToForm("TopStackGrowthCurveEndArcShapeV2s", instanceWithInferedType, &instanceWithInferedType.TopStackGrowthCurveEndArcShapeV2s, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TopStartArcShapeV2:
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
			AssociationReverseSliceToForm[*models.TopStartArcShapeV2Grid, *models.TopStartArcShapeV2](
				"TopStartArcShapeV2Grid",
				"TopStartArcShapesV2",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TopStartArcShapeV2Grid) []*models.TopStartArcShapeV2 {
					return owner.TopStartArcShapesV2
				})
		}

	case *models.TopStartArcShapeV2Grid:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("TopStartArcShapesV2", instanceWithInferedType, &instanceWithInferedType.TopStartArcShapesV2, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	default:
		_ = instanceWithInferedType
	}
}
