// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsme/cld/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Category1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Category1Shape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Category1", instanceWithInferedType.Category1, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "Category1Shapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"Category1Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"Category1Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Category2:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Category2Shape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Category2", instanceWithInferedType.Category2, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "Category2Shapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"Category2Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"Category2Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Category3:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Category3Shape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Category3", instanceWithInferedType.Category3, formGroup, probe)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "Category3Shapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"Category3Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"Category3Shapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ControlPointShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X_Relative", instanceWithInferedType.X_Relative, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y_Relative", instanceWithInferedType.Y_Relative, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsStartShapeTheClosestShape", instanceWithInferedType.IsStartShapeTheClosestShape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "InfluenceShape"
			rf.Fieldname = "ControlPointShapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.InfluenceShape),
					"ControlPointShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.InfluenceShape](
					nil,
					"ControlPointShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Desk:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SelectedDiagram", instanceWithInferedType.SelectedDiagram, formGroup, probe)

	case *models.Diagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Category1Shapes", instanceWithInferedType, &instanceWithInferedType.Category1Shapes, formGroup, probe)
		AssociationSliceToForm("Category2Shapes", instanceWithInferedType, &instanceWithInferedType.Category2Shapes, formGroup, probe)
		AssociationSliceToForm("Category3Shapes", instanceWithInferedType, &instanceWithInferedType.Category3Shapes, formGroup, probe)
		AssociationSliceToForm("InfluenceShapes", instanceWithInferedType, &instanceWithInferedType.InfluenceShapes, formGroup, probe)
		BasicFieldtoForm("IsEditable", instanceWithInferedType.IsEditable, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsNodeExpanded", instanceWithInferedType.IsNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsCategory1NodeExpanded", instanceWithInferedType.IsCategory1NodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsCategory2NodeExpanded", instanceWithInferedType.IsCategory2NodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsCategory3NodeExpanded", instanceWithInferedType.IsCategory3NodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInfluenceCategoryNodeExpanded", instanceWithInferedType.IsInfluenceCategoryNodeExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsCategory1Shown", instanceWithInferedType.IsCategory1Shown, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsCategory2Shown", instanceWithInferedType.IsCategory2Shown, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsCategory3Shown", instanceWithInferedType.IsCategory3Shown, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInfluenceCategoryShown", instanceWithInferedType.IsInfluenceCategoryShown, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("XMargin", instanceWithInferedType.XMargin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("YMargin", instanceWithInferedType.YMargin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NbPixPerCharacter", instanceWithInferedType.NbPixPerCharacter, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("RedColorCode", instanceWithInferedType.RedColorCode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BackgroundGreyColorCode", instanceWithInferedType.BackgroundGreyColorCode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("GrayColorCode", instanceWithInferedType.GrayColorCode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Category1RectAnchorType", instanceWithInferedType.Category1RectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Category1TextAnchorType", instanceWithInferedType.Category1TextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Category1DominantBaselineType", instanceWithInferedType.Category1DominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Category1FontSize", instanceWithInferedType.Category1FontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Category1FontWeigth", instanceWithInferedType.Category1FontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Category1FontFamily", instanceWithInferedType.Category1FontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Category1LetterSpacing", instanceWithInferedType.Category1LetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Category2TypeFontSize", instanceWithInferedType.Category2TypeFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Category2TypeFontWeigth", instanceWithInferedType.Category2TypeFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Category2TypeFontFamily", instanceWithInferedType.Category2TypeFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Category2TypeLetterSpacing", instanceWithInferedType.Category2TypeLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Category2TypeRectAnchorType", instanceWithInferedType.Category2TypeRectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Category2DominantBaselineType", instanceWithInferedType.Category2DominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Category2StrokeWidth", instanceWithInferedType.Category2StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Category3RectAnchorType", instanceWithInferedType.Category3RectAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Category3TextAnchorType", instanceWithInferedType.Category3TextAnchorType, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Category3DominantBaselineType", instanceWithInferedType.Category3DominantBaselineType, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Category3FontSize", instanceWithInferedType.Category3FontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Category3FontWeigth", instanceWithInferedType.Category3FontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Category3FontFamily", instanceWithInferedType.Category3FontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Category3LetterSpacing", instanceWithInferedType.Category3LetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceStrokeWidth", instanceWithInferedType.InfluenceStrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceArrowSize", instanceWithInferedType.InfluenceArrowSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceArrowStartOffset", instanceWithInferedType.InfluenceArrowStartOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceArrowEndOffset", instanceWithInferedType.InfluenceArrowEndOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceCornerRadius", instanceWithInferedType.InfluenceCornerRadius, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceFontSize", instanceWithInferedType.InfluenceFontSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceFontWeigth", instanceWithInferedType.InfluenceFontWeigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceFontFamily", instanceWithInferedType.InfluenceFontFamily, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceLetterSpacing", instanceWithInferedType.InfluenceLetterSpacing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InfluenceDashedLinePattern", instanceWithInferedType.InfluenceDashedLinePattern, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Influence:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SourceCategory1", instanceWithInferedType.SourceCategory1, formGroup, probe)
		AssociationFieldToForm("SourceCategory2", instanceWithInferedType.SourceCategory2, formGroup, probe)
		AssociationFieldToForm("SourceCategory3", instanceWithInferedType.SourceCategory3, formGroup, probe)
		AssociationFieldToForm("TargetCategory1", instanceWithInferedType.TargetCategory1, formGroup, probe)
		AssociationFieldToForm("TargetCategory2", instanceWithInferedType.TargetCategory2, formGroup, probe)
		AssociationFieldToForm("TargetCategory3", instanceWithInferedType.TargetCategory3, formGroup, probe)
		BasicFieldtoForm("IsHypothtical", instanceWithInferedType.IsHypothtical, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TextAtEndOfArrow", instanceWithInferedType.TextAtEndOfArrow, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.InfluenceShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Influence", instanceWithInferedType.Influence, formGroup, probe)
		AssociationSliceToForm("ControlPointShapes", instanceWithInferedType, &instanceWithInferedType.ControlPointShapes, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Diagram"
			rf.Fieldname = "InfluenceShapes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Diagram),
					"InfluenceShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Diagram](
					nil,
					"InfluenceShapes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
