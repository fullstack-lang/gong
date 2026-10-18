// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func StageSetFillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *StageSetProbe,
) {
	switch inst := instance.(type) {
	case *models.Chapter:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MardownContent", inst.MardownContent, probe.formStage, formGroup)

		{
			// Slice of pointers: Sections
			div := (&form.FormDiv{Name: "Sections"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sections {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sections",
				Label: "Sections",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Pages
			div := (&form.FormDiv{Name: "Pages"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Pages {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Pages",
				Label: "Pages",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: SubChapters
			div := (&form.FormDiv{Name: "SubChapters"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.SubChapters {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "SubChapters",
				Label: "SubChapters",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Chapters {
				for _, target := range src.SubChapters {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Chapter", "SubChapters", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Contents {
				for _, target := range src.Chapters {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Content", "Chapters", refNames, formGroup, probe.formStage)
		}
	case *models.Content:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MardownContent", inst.MardownContent, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ContentPath", inst.ContentPath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OutputPath", inst.OutputPath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StaticPath", inst.StaticPath, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LogoSVGFile", inst.LogoSVGFile, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsBespokeLogoFileName", inst.IsBespokeLogoFileName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BespokeLogoFileName", inst.BespokeLogoFileName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsBespokePageTileLogoFileName", inst.IsBespokePageTileLogoFileName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BespokePageTileLogoFileName", inst.BespokePageTileLogoFileName, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Target", inst.Target, formGroup, probe.formStage)

		{
			// Slice of pointers: Chapters
			div := (&form.FormDiv{Name: "Chapters"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Chapters {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Chapters",
				Label: "Chapters",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("VersionInfo", inst.VersionInfo, probe.formStage, formGroup)
	case *models.DownloadableFile:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Base64Content", inst.Base64Content, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sections {
				if src.DownloadableFile == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Section", "DownloadableFile", refNames, formGroup, probe.formStage)
		}
	case *models.JpgImage:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Base64Content", inst.Base64Content, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sections {
				if src.JpgImage == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Section", "JpgImage", refNames, formGroup, probe.formStage)
		}
	case *models.Page:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MardownContent", inst.MardownContent, probe.formStage, formGroup)

		{
			// Slice of pointers: Sections
			div := (&form.FormDiv{Name: "Sections"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sections {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sections",
				Label: "Sections",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Chapters {
				for _, target := range src.Pages {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Chapter", "Pages", refNames, formGroup, probe.formStage)
		}
	case *models.PngImage:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Base64Content", inst.Base64Content, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sections {
				if src.PngImage == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Section", "PngImage", refNames, formGroup, probe.formStage)
		}
	case *models.Section:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MardownContent", inst.MardownContent, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsImage", inst.IsImage, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SvgImage", inst.SvgImage, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.SvgImage](), probe.formStage)
		StageSetAssociationFieldToForm("PngImage", inst.PngImage, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.PngImage](), probe.formStage)
		StageSetAssociationFieldToForm("JpgImage", inst.JpgImage, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.JpgImage](), probe.formStage)
		StageSetBasicFieldtoForm("IsDownloadableFile", inst.IsDownloadableFile, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DownloadableFile", inst.DownloadableFile, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.DownloadableFile](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Chapters {
				for _, target := range src.Sections {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Chapter", "Sections", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Pages {
				for _, target := range src.Sections {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Page", "Sections", refNames, formGroup, probe.formStage)
		}
	case *models.SvgImage:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Sections {
				if src.SvgImage == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Section", "SvgImage", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
