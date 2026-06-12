// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/pkg/docx/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Body:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Paragraphs", instanceWithInferedType, &instanceWithInferedType.Paragraphs, formGroup, probe)
		AssociationSliceToForm("Tables", instanceWithInferedType, &instanceWithInferedType.Tables, formGroup, probe)
		AssociationFieldToForm("LastParagraph", instanceWithInferedType.LastParagraph, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Document:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("File", instanceWithInferedType.File, formGroup, probe)
		AssociationFieldToForm("Root", instanceWithInferedType.Root, formGroup, probe)
		AssociationFieldToForm("Body", instanceWithInferedType.Body, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Docx:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Files", instanceWithInferedType, &instanceWithInferedType.Files, formGroup, probe)
		AssociationFieldToForm("Document", instanceWithInferedType.Document, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.File:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Docx, *models.File](
				"Docx",
				"Files",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Docx) []*models.File {
					return owner.Files
				})
		}

	case *models.Node:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Nodes", instanceWithInferedType, &instanceWithInferedType.Nodes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Node, *models.Node](
				"Node",
				"Nodes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Node) []*models.Node {
					return owner.Nodes
				})
		}

	case *models.Paragraph:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		AssociationFieldToForm("ParagraphProperties", instanceWithInferedType.ParagraphProperties, formGroup, probe)
		AssociationSliceToForm("Runes", instanceWithInferedType, &instanceWithInferedType.Runes, formGroup, probe)
		BasicFieldtoForm("CollatedText", instanceWithInferedType.CollatedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Next", instanceWithInferedType.Next, formGroup, probe)
		AssociationFieldToForm("Previous", instanceWithInferedType.Previous, formGroup, probe)
		AssociationFieldToForm("EnclosingBody", instanceWithInferedType.EnclosingBody, formGroup, probe)
		AssociationFieldToForm("EnclosingTableColumn", instanceWithInferedType.EnclosingTableColumn, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Body, *models.Paragraph](
				"Body",
				"Paragraphs",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Body) []*models.Paragraph {
					return owner.Paragraphs
				})
		}
		{
			AssociationReverseSliceToForm[*models.TableColumn, *models.Paragraph](
				"TableColumn",
				"Paragraphs",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TableColumn) []*models.Paragraph {
					return owner.Paragraphs
				})
		}

	case *models.ParagraphProperties:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ParagraphStyle", instanceWithInferedType.ParagraphStyle, formGroup, probe)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.ParagraphStyle:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ValAttr", instanceWithInferedType.ValAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Rune:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		AssociationFieldToForm("Text", instanceWithInferedType.Text, formGroup, probe)
		AssociationFieldToForm("RuneProperties", instanceWithInferedType.RuneProperties, formGroup, probe)
		AssociationFieldToForm("EnclosingParagraph", instanceWithInferedType.EnclosingParagraph, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Paragraph, *models.Rune](
				"Paragraph",
				"Runes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Paragraph) []*models.Rune {
					return owner.Runes
				})
		}

	case *models.RuneProperties:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		BasicFieldtoForm("IsBold", instanceWithInferedType.IsBold, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsStrike", instanceWithInferedType.IsStrike, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsItalic", instanceWithInferedType.IsItalic, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Table:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("TableProperties", instanceWithInferedType.TableProperties, formGroup, probe)
		AssociationSliceToForm("TableRows", instanceWithInferedType, &instanceWithInferedType.TableRows, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Body, *models.Table](
				"Body",
				"Tables",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Body) []*models.Table {
					return owner.Tables
				})
		}

	case *models.TableColumn:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		AssociationSliceToForm("Paragraphs", instanceWithInferedType, &instanceWithInferedType.Paragraphs, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.TableRow, *models.TableColumn](
				"TableRow",
				"TableColumns",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.TableRow) []*models.TableColumn {
					return owner.TableColumns
				})
		}

	case *models.TableProperties:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("TableStyle", instanceWithInferedType.TableStyle, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.TableRow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		AssociationSliceToForm("TableColumns", instanceWithInferedType, &instanceWithInferedType.TableColumns, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Table, *models.TableRow](
				"Table",
				"TableRows",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Table) []*models.TableRow {
					return owner.TableRows
				})
		}

	case *models.TableStyle:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Val", instanceWithInferedType.Val, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		BasicFieldtoForm("PreserveWhiteSpace", instanceWithInferedType.PreserveWhiteSpace, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("EnclosingRune", instanceWithInferedType.EnclosingRune, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	default:
		_ = instanceWithInferedType
	}
}
