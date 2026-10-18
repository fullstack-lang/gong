// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
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
	case *models.AttributeShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FieldTypeAsString", inst.FieldTypeAsString, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Structname", inst.Structname, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Fieldtypename", inst.Fieldtypename, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GongStructShapes {
				for _, target := range src.AttributeShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GongStructShape", "AttributeShapes", refNames, formGroup, probe.formStage)
		}
	case *models.Classdiagram:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Description", inst.Description, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsIncludedInStaticWebSite", inst.IsIncludedInStaticWebSite, probe.formStage, formGroup)

		{
			// Slice of pointers: GongStructShapes
			div := (&form.FormDiv{Name: "GongStructShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GongStructShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GongStructShapes",
				Label: "GongStructShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GongEnumShapes
			div := (&form.FormDiv{Name: "GongEnumShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GongEnumShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GongEnumShapes",
				Label: "GongEnumShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: GongNoteShapes
			div := (&form.FormDiv{Name: "GongNoteShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GongNoteShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GongNoteShapes",
				Label: "GongNoteShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("ShowNbInstances", inst.ShowNbInstances, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowMultiplicity", inst.ShowMultiplicity, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowLinkNames", inst.ShowLinkNames, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsInRenameMode", inst.IsInRenameMode, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NodeGongStructsIsExpanded", inst.NodeGongStructsIsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NodeGongStructNodeExpansion", inst.NodeGongStructNodeExpansion, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NodeGongEnumsIsExpanded", inst.NodeGongEnumsIsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NodeGongEnumNodeExpansion", inst.NodeGongEnumNodeExpansion, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NodeGongNotesIsExpanded", inst.NodeGongNotesIsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NodeGongNoteNodeExpansion", inst.NodeGongNoteNodeExpansion, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramPackages {
				for _, target := range src.Classdiagrams {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramPackage", "Classdiagrams", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.DiagramPackages {
				if src.SelectedClassdiagram == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DiagramPackage", "SelectedClassdiagram", refNames, formGroup, probe.formStage)
		}
	case *models.DiagramPackage:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Path", inst.Path, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("GongModelPath", inst.GongModelPath, probe.formStage, formGroup)

		{
			// Slice of pointers: Classdiagrams
			div := (&form.FormDiv{Name: "Classdiagrams"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Classdiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Classdiagrams",
				Label: "Classdiagrams",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("SelectedClassdiagram", inst.SelectedClassdiagram, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Classdiagram](), probe.formStage)
		StageSetBasicFieldtoForm("AbsolutePathToDiagramPackage", inst.AbsolutePathToDiagramPackage, probe.formStage, formGroup)
	case *models.GongEnumShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: GongEnumValueShapes
			div := (&form.FormDiv{Name: "GongEnumValueShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GongEnumValueShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GongEnumValueShapes",
				Label: "GongEnumValueShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Classdiagrams {
				for _, target := range src.GongEnumShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Classdiagram", "GongEnumShapes", refNames, formGroup, probe.formStage)
		}
	case *models.GongEnumValueShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GongEnumShapes {
				for _, target := range src.GongEnumValueShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GongEnumShape", "GongEnumValueShapes", refNames, formGroup, probe.formStage)
		}
	case *models.GongNoteLinkShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Identifier", inst.Identifier, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Type", inst.Type, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GongNoteShapes {
				for _, target := range src.GongNoteLinkShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GongNoteShape", "GongNoteLinkShapes", refNames, formGroup, probe.formStage)
		}
	case *models.GongNoteShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Identifier", inst.Identifier, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Body", inst.Body, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BodyHTML", inst.BodyHTML, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Matched", inst.Matched, probe.formStage, formGroup)

		{
			// Slice of pointers: GongNoteLinkShapes
			div := (&form.FormDiv{Name: "GongNoteLinkShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GongNoteLinkShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GongNoteLinkShapes",
				Label: "GongNoteLinkShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Classdiagrams {
				for _, target := range src.GongNoteShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Classdiagram", "GongNoteShapes", refNames, formGroup, probe.formStage)
		}
	case *models.GongStructShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			// Slice of pointers: AttributeShapes
			div := (&form.FormDiv{Name: "AttributeShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AttributeShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AttributeShapes",
				Label: "AttributeShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: LinkShapes
			div := (&form.FormDiv{Name: "LinkShapes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.LinkShapes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "LinkShapes",
				Label: "LinkShapes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsSelected", inst.IsSelected, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Classdiagrams {
				for _, target := range src.GongStructShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Classdiagram", "GongStructShapes", refNames, formGroup, probe.formStage)
		}
	case *models.LinkShape:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FieldOffsetX", inst.FieldOffsetX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FieldOffsetY", inst.FieldOffsetY, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("TargetMultiplicity", inst.TargetMultiplicity, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("TargetMultiplicityOffsetX", inst.TargetMultiplicityOffsetX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TargetMultiplicityOffsetY", inst.TargetMultiplicityOffsetY, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("SourceMultiplicity", inst.SourceMultiplicity, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("SourceMultiplicityOffsetX", inst.SourceMultiplicityOffsetX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SourceMultiplicityOffsetY", inst.SourceMultiplicityOffsetY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("X", inst.X, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Y", inst.Y, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("StartOrientation", inst.StartOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("StartRatio", inst.StartRatio, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EndOrientation", inst.EndOrientation, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("EndRatio", inst.EndRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CornerOffsetRatio", inst.CornerOffsetRatio, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsHidden", inst.IsHidden, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GongStructShapes {
				for _, target := range src.LinkShapes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GongStructShape", "LinkShapes", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
