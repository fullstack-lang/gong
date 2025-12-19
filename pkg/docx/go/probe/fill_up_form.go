// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

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

	case *models.Document:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("File", instanceWithInferedType.File, formGroup, probe)
		AssociationFieldToForm("Root", instanceWithInferedType.Root, formGroup, probe)
		AssociationFieldToForm("Body", instanceWithInferedType.Body, formGroup, probe)

	case *models.Docx:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Files", instanceWithInferedType, &instanceWithInferedType.Files, formGroup, probe)
		AssociationFieldToForm("Document", instanceWithInferedType.Document, formGroup, probe)

	case *models.File:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Docx"
			rf.Fieldname = "Files"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Docx),
					"Files",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Docx](
					nil,
					"Files",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Node:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Nodes", instanceWithInferedType, &instanceWithInferedType.Nodes, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Node"
			rf.Fieldname = "Nodes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Node),
					"Nodes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Node](
					nil,
					"Nodes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Next", instanceWithInferedType.Next, formGroup, probe)
		AssociationFieldToForm("Previous", instanceWithInferedType.Previous, formGroup, probe)
		AssociationFieldToForm("EnclosingBody", instanceWithInferedType.EnclosingBody, formGroup, probe)
		AssociationFieldToForm("EnclosingTableColumn", instanceWithInferedType.EnclosingTableColumn, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Body"
			rf.Fieldname = "Paragraphs"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Body),
					"Paragraphs",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Body](
					nil,
					"Paragraphs",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TableColumn"
			rf.Fieldname = "Paragraphs"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.TableColumn),
					"Paragraphs",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.TableColumn](
					nil,
					"Paragraphs",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.ParagraphProperties:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("ParagraphStyle", instanceWithInferedType.ParagraphStyle, formGroup, probe)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)

	case *models.ParagraphStyle:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ValAttr", instanceWithInferedType.ValAttr, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

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
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Paragraph"
			rf.Fieldname = "Runes"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Paragraph),
					"Runes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Paragraph](
					nil,
					"Runes",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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

	case *models.Table:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("TableProperties", instanceWithInferedType.TableProperties, formGroup, probe)
		AssociationSliceToForm("TableRows", instanceWithInferedType, &instanceWithInferedType.TableRows, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Body"
			rf.Fieldname = "Tables"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Body),
					"Tables",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Body](
					nil,
					"Tables",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.TableColumn:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		AssociationSliceToForm("Paragraphs", instanceWithInferedType, &instanceWithInferedType.Paragraphs, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "TableRow"
			rf.Fieldname = "TableColumns"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.TableRow),
					"TableColumns",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.TableRow](
					nil,
					"TableColumns",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.TableProperties:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("TableStyle", instanceWithInferedType.TableStyle, formGroup, probe)

	case *models.TableRow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Content", instanceWithInferedType.Content, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Node", instanceWithInferedType.Node, formGroup, probe)
		AssociationSliceToForm("TableColumns", instanceWithInferedType, &instanceWithInferedType.TableColumns, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Table"
			rf.Fieldname = "TableRows"
			reverseFieldOwner := instanceWithInferedType.GongGetReverseFieldOwner(probe.stageOfInterest, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Table),
					"TableRows",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Table](
					nil,
					"TableRows",
					instanceWithInferedType,
					formGroup,
					probe)
			}
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

	default:
		_ = instanceWithInferedType
	}
}
