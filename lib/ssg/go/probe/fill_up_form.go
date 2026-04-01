// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Chapter:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MardownContent", instanceWithInferedType.MardownContent, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 300)
		AssociationSliceToForm("Pages", instanceWithInferedType, &instanceWithInferedType.Pages, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Content"
			rf.Fieldname = "Chapters"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Content),
					"Chapters",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Content](
					nil,
					"Chapters",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Content:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MardownContent", instanceWithInferedType.MardownContent, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 300)
		BasicFieldtoForm("ContentPath", instanceWithInferedType.ContentPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("OutputPath", instanceWithInferedType.OutputPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LayoutPath", instanceWithInferedType.LayoutPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StaticPath", instanceWithInferedType.StaticPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsBespokeLogoFileName", instanceWithInferedType.IsBespokeLogoFileName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BespokeLogoFileName", instanceWithInferedType.BespokeLogoFileName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsBespokePageTileLogoFileName", instanceWithInferedType.IsBespokePageTileLogoFileName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BespokePageTileLogoFileName", instanceWithInferedType.BespokePageTileLogoFileName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Target", instanceWithInferedType.Target, instanceWithInferedType, probe.formStage, formGroup)
		AssociationSliceToForm("Chapters", instanceWithInferedType, &instanceWithInferedType.Chapters, formGroup, probe)
		BasicFieldtoForm("VersionInfo", instanceWithInferedType.VersionInfo, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.DownloadableFile:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Base64Content", instanceWithInferedType.Base64Content, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)

	case *models.JpgImage:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Base64Content", instanceWithInferedType.Base64Content, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)

	case *models.Page:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MardownContent", instanceWithInferedType.MardownContent, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 300)
		AssociationSliceToForm("Sections", instanceWithInferedType, &instanceWithInferedType.Sections, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Chapter"
			rf.Fieldname = "Pages"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Chapter),
					"Pages",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Chapter](
					nil,
					"Pages",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.PngImage:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Base64Content", instanceWithInferedType.Base64Content, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)

	case *models.Section:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MardownContent", instanceWithInferedType.MardownContent, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 300)
		BasicFieldtoForm("IsImage", instanceWithInferedType.IsImage, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SvgImage", instanceWithInferedType.SvgImage, formGroup, probe)
		AssociationFieldToForm("PngImage", instanceWithInferedType.PngImage, formGroup, probe)
		AssociationFieldToForm("JpgImage", instanceWithInferedType.JpgImage, formGroup, probe)
		BasicFieldtoForm("IsDownloadableFile", instanceWithInferedType.IsDownloadableFile, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("DownloadableFile", instanceWithInferedType.DownloadableFile, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Page"
			rf.Fieldname = "Sections"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Page),
					"Sections",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Page](
					nil,
					"Sections",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.SvgImage:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)

	default:
		_ = instanceWithInferedType
	}
}
