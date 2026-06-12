// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

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
		AssociationSliceToForm("Sections", instanceWithInferedType, &instanceWithInferedType.Sections, formGroup, probe)
		AssociationSliceToForm("Pages", instanceWithInferedType, &instanceWithInferedType.Pages, formGroup, probe)
		AssociationSliceToForm("SubChapters", instanceWithInferedType, &instanceWithInferedType.SubChapters, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Chapter, *models.Chapter](
				"Chapter",
				"SubChapters",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Chapter) []*models.Chapter {
					return owner.SubChapters
				})
		}
		{
			AssociationReverseSliceToForm[*models.Content, *models.Chapter](
				"Content",
				"Chapters",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Content) []*models.Chapter {
					return owner.Chapters
				})
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
		BasicFieldtoForm("StaticPath", instanceWithInferedType.StaticPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LogoSVGFile", instanceWithInferedType.LogoSVGFile, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300)
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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.DownloadableFile:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Base64Content", instanceWithInferedType.Base64Content, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.JpgImage:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Base64Content", instanceWithInferedType.Base64Content, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Page:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MardownContent", instanceWithInferedType.MardownContent, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 300)
		AssociationSliceToForm("Sections", instanceWithInferedType, &instanceWithInferedType.Sections, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Chapter, *models.Page](
				"Chapter",
				"Pages",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Chapter) []*models.Page {
					return owner.Pages
				})
		}

	case *models.PngImage:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Base64Content", instanceWithInferedType.Base64Content, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

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
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Chapter, *models.Section](
				"Chapter",
				"Sections",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Chapter) []*models.Section {
					return owner.Sections
				})
		}
		{
			AssociationReverseSliceToForm[*models.Page, *models.Section](
				"Page",
				"Sections",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Page) []*models.Section {
					return owner.Sections
				})
		}

	case *models.SvgImage:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 400)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	default:
		_ = instanceWithInferedType
	}
}
