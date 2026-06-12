// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.AttributeShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FieldTypeAsString", instanceWithInferedType.FieldTypeAsString, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Structname", instanceWithInferedType.Structname, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Fieldtypename", instanceWithInferedType.Fieldtypename, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.GongStructShape, *models.AttributeShape](
				"GongStructShape",
				"AttributeShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.GongStructShape) []*models.AttributeShape {
					return owner.AttributeShapes
				})
		}

	case *models.Classdiagram:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Description", instanceWithInferedType.Description, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		BasicFieldtoForm("IsIncludedInStaticWebSite", instanceWithInferedType.IsIncludedInStaticWebSite, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("GongStructShapes", instanceWithInferedType, &instanceWithInferedType.GongStructShapes, formGroup, probe)
		AssociationSliceToForm("GongEnumShapes", instanceWithInferedType, &instanceWithInferedType.GongEnumShapes, formGroup, probe)
		AssociationSliceToForm("GongNoteShapes", instanceWithInferedType, &instanceWithInferedType.GongNoteShapes, formGroup, probe)
		BasicFieldtoForm("ShowNbInstances", instanceWithInferedType.ShowNbInstances, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ShowMultiplicity", instanceWithInferedType.ShowMultiplicity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ShowLinkNames", instanceWithInferedType.ShowLinkNames, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInRenameMode", instanceWithInferedType.IsInRenameMode, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NodeGongStructsIsExpanded", instanceWithInferedType.NodeGongStructsIsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NodeGongStructNodeExpansion", instanceWithInferedType.NodeGongStructNodeExpansion, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NodeGongEnumsIsExpanded", instanceWithInferedType.NodeGongEnumsIsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NodeGongEnumNodeExpansion", instanceWithInferedType.NodeGongEnumNodeExpansion, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NodeGongNotesIsExpanded", instanceWithInferedType.NodeGongNotesIsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NodeGongNoteNodeExpansion", instanceWithInferedType.NodeGongNoteNodeExpansion, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.DiagramPackage, *models.Classdiagram](
				"DiagramPackage",
				"Classdiagrams",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.DiagramPackage) []*models.Classdiagram {
					return owner.Classdiagrams
				})
		}

	case *models.DiagramPackage:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Path", instanceWithInferedType.Path, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("GongModelPath", instanceWithInferedType.GongModelPath, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Classdiagrams", instanceWithInferedType, &instanceWithInferedType.Classdiagrams, formGroup, probe)
		AssociationFieldToForm("SelectedClassdiagram", instanceWithInferedType.SelectedClassdiagram, formGroup, probe)
		BasicFieldtoForm("AbsolutePathToDiagramPackage", instanceWithInferedType.AbsolutePathToDiagramPackage, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.GongEnumShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("GongEnumValueShapes", instanceWithInferedType, &instanceWithInferedType.GongEnumValueShapes, formGroup, probe)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Classdiagram, *models.GongEnumShape](
				"Classdiagram",
				"GongEnumShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Classdiagram) []*models.GongEnumShape {
					return owner.GongEnumShapes
				})
		}

	case *models.GongEnumValueShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.GongEnumShape, *models.GongEnumValueShape](
				"GongEnumShape",
				"GongEnumValueShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.GongEnumShape) []*models.GongEnumValueShape {
					return owner.GongEnumValueShapes
				})
		}

	case *models.GongNoteLinkShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.GongNoteShape, *models.GongNoteLinkShape](
				"GongNoteShape",
				"GongNoteLinkShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.GongNoteShape) []*models.GongNoteLinkShape {
					return owner.GongNoteLinkShapes
				})
		}

	case *models.GongNoteShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Identifier", instanceWithInferedType.Identifier, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Body", instanceWithInferedType.Body, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BodyHTML", instanceWithInferedType.BodyHTML, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Matched", instanceWithInferedType.Matched, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("GongNoteLinkShapes", instanceWithInferedType, &instanceWithInferedType.GongNoteLinkShapes, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Classdiagram, *models.GongNoteShape](
				"Classdiagram",
				"GongNoteShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Classdiagram) []*models.GongNoteShape {
					return owner.GongNoteShapes
				})
		}

	case *models.GongStructShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("AttributeShapes", instanceWithInferedType, &instanceWithInferedType.AttributeShapes, formGroup, probe)
		AssociationSliceToForm("LinkShapes", instanceWithInferedType, &instanceWithInferedType.LinkShapes, formGroup, probe)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Height", instanceWithInferedType.Height, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsSelected", instanceWithInferedType.IsSelected, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Classdiagram, *models.GongStructShape](
				"Classdiagram",
				"GongStructShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Classdiagram) []*models.GongStructShape {
					return owner.GongStructShapes
				})
		}

	case *models.LinkShape:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FieldOffsetX", instanceWithInferedType.FieldOffsetX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("FieldOffsetY", instanceWithInferedType.FieldOffsetY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("TargetMultiplicity", instanceWithInferedType.TargetMultiplicity, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("TargetMultiplicityOffsetX", instanceWithInferedType.TargetMultiplicityOffsetX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TargetMultiplicityOffsetY", instanceWithInferedType.TargetMultiplicityOffsetY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("SourceMultiplicity", instanceWithInferedType.SourceMultiplicity, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("SourceMultiplicityOffsetX", instanceWithInferedType.SourceMultiplicityOffsetX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SourceMultiplicityOffsetY", instanceWithInferedType.SourceMultiplicityOffsetY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("X", instanceWithInferedType.X, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Y", instanceWithInferedType.Y, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("StartOrientation", instanceWithInferedType.StartOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("StartRatio", instanceWithInferedType.StartRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("EndOrientation", instanceWithInferedType.EndOrientation, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("EndRatio", instanceWithInferedType.EndRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("CornerOffsetRatio", instanceWithInferedType.CornerOffsetRatio, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.GongStructShape, *models.LinkShape](
				"GongStructShape",
				"LinkShapes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.GongStructShape) []*models.LinkShape {
					return owner.LinkShapes
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
