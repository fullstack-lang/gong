// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/docx/go/models"
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
	case *models.Body:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Paragraphs
			div := (&form.FormDiv{Name: "Paragraphs"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Paragraphs {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Paragraphs",
				Label: "Paragraphs",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Tables
			div := (&form.FormDiv{Name: "Tables"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Tables {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Tables",
				Label: "Tables",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("LastParagraph", inst.LastParagraph, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Paragraph](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Documents {
				if src.Body == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Document", "Body", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Paragraphs {
				if src.EnclosingBody == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Paragraph", "EnclosingBody", refNames, formGroup, probe.formStage)
		}
	case *models.Document:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("File", inst.File, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.File](), probe.formStage)
		StageSetAssociationFieldToForm("Root", inst.Root, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Node](), probe.formStage)
		StageSetAssociationFieldToForm("Body", inst.Body, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Body](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Docxs {
				if src.Document == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Docx", "Document", refNames, formGroup, probe.formStage)
		}
	case *models.Docx:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Files
			div := (&form.FormDiv{Name: "Files"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Files {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Files",
				Label: "Files",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("Document", inst.Document, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Document](), probe.formStage)
	case *models.File:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Documents {
				if src.File == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Document", "File", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Docxs {
				for _, target := range src.Files {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Docx", "Files", refNames, formGroup, probe.formStage)
		}
	case *models.Node:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Nodes
			div := (&form.FormDiv{Name: "Nodes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Nodes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Nodes",
				Label: "Nodes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Documents {
				if src.Root == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Document", "Root", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Nodes {
				for _, target := range src.Nodes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Node", "Nodes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Paragraphs {
				if src.Node == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Paragraph", "Node", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ParagraphPropertiess {
				if src.Node == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ParagraphProperties", "Node", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ParagraphStyles {
				if src.Node == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ParagraphStyle", "Node", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Runes {
				if src.Node == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rune", "Node", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.RunePropertiess {
				if src.Node == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.RuneProperties", "Node", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tables {
				if src.Node == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Table", "Node", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TableColumns {
				if src.Node == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TableColumn", "Node", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TablePropertiess {
				if src.Node == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TableProperties", "Node", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TableRows {
				if src.Node == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TableRow", "Node", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TableStyles {
				if src.Node == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TableStyle", "Node", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Texts {
				if src.Node == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Text", "Node", refNames, formGroup, probe.formStage)
		}
	case *models.Paragraph:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Node", inst.Node, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Node](), probe.formStage)
		StageSetAssociationFieldToForm("ParagraphProperties", inst.ParagraphProperties, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ParagraphProperties](), probe.formStage)

		{
			// Slice of pointers: Runes
			div := (&form.FormDiv{Name: "Runes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Runes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Runes",
				Label: "Runes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("CollatedText", inst.CollatedText, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Next", inst.Next, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Paragraph](), probe.formStage)
		StageSetAssociationFieldToForm("Previous", inst.Previous, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Paragraph](), probe.formStage)
		StageSetAssociationFieldToForm("EnclosingBody", inst.EnclosingBody, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Body](), probe.formStage)
		StageSetAssociationFieldToForm("EnclosingTableColumn", inst.EnclosingTableColumn, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TableColumn](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Bodys {
				for _, target := range src.Paragraphs {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Body", "Paragraphs", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Bodys {
				if src.LastParagraph == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Body", "LastParagraph", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Paragraphs {
				if src.Next == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Paragraph", "Next", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Paragraphs {
				if src.Previous == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Paragraph", "Previous", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Runes {
				if src.EnclosingParagraph == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rune", "EnclosingParagraph", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TableColumns {
				for _, target := range src.Paragraphs {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TableColumn", "Paragraphs", refNames, formGroup, probe.formStage)
		}
	case *models.ParagraphProperties:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ParagraphStyle", inst.ParagraphStyle, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.ParagraphStyle](), probe.formStage)
		StageSetAssociationFieldToForm("Node", inst.Node, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Node](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Paragraphs {
				if src.ParagraphProperties == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Paragraph", "ParagraphProperties", refNames, formGroup, probe.formStage)
		}
	case *models.ParagraphStyle:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Node", inst.Node, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Node](), probe.formStage)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ValAttr", inst.ValAttr, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.ParagraphPropertiess {
				if src.ParagraphStyle == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ParagraphProperties", "ParagraphStyle", refNames, formGroup, probe.formStage)
		}
	case *models.Rune:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Node", inst.Node, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Node](), probe.formStage)
		StageSetAssociationFieldToForm("Text", inst.Text, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Text](), probe.formStage)
		StageSetAssociationFieldToForm("RuneProperties", inst.RuneProperties, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.RuneProperties](), probe.formStage)
		StageSetAssociationFieldToForm("EnclosingParagraph", inst.EnclosingParagraph, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Paragraph](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Paragraphs {
				for _, target := range src.Runes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Paragraph", "Runes", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Texts {
				if src.EnclosingRune == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Text", "EnclosingRune", refNames, formGroup, probe.formStage)
		}
	case *models.RuneProperties:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Node", inst.Node, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Node](), probe.formStage)
		StageSetBasicFieldtoForm("IsBold", inst.IsBold, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsStrike", inst.IsStrike, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsItalic", inst.IsItalic, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Runes {
				if src.RuneProperties == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rune", "RuneProperties", refNames, formGroup, probe.formStage)
		}
	case *models.Table:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Node", inst.Node, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Node](), probe.formStage)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TableProperties", inst.TableProperties, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TableProperties](), probe.formStage)

		{
			// Slice of pointers: TableRows
			div := (&form.FormDiv{Name: "TableRows"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TableRows {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TableRows",
				Label: "TableRows",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Bodys {
				for _, target := range src.Tables {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Body", "Tables", refNames, formGroup, probe.formStage)
		}
	case *models.TableColumn:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Node", inst.Node, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Node](), probe.formStage)

		{
			// Slice of pointers: Paragraphs
			div := (&form.FormDiv{Name: "Paragraphs"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Paragraphs {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Paragraphs",
				Label: "Paragraphs",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Paragraphs {
				if src.EnclosingTableColumn == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Paragraph", "EnclosingTableColumn", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TableRows {
				for _, target := range src.TableColumns {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TableRow", "TableColumns", refNames, formGroup, probe.formStage)
		}
	case *models.TableProperties:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Node", inst.Node, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Node](), probe.formStage)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("TableStyle", inst.TableStyle, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.TableStyle](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tables {
				if src.TableProperties == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Table", "TableProperties", refNames, formGroup, probe.formStage)
		}
	case *models.TableRow:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Node", inst.Node, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Node](), probe.formStage)

		{
			// Slice of pointers: TableColumns
			div := (&form.FormDiv{Name: "TableColumns"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.TableColumns {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "TableColumns",
				Label: "TableColumns",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tables {
				for _, target := range src.TableRows {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Table", "TableRows", refNames, formGroup, probe.formStage)
		}
	case *models.TableStyle:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Node", inst.Node, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Node](), probe.formStage)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Val", inst.Val, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.TablePropertiess {
				if src.TableStyle == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.TableProperties", "TableStyle", refNames, formGroup, probe.formStage)
		}
	case *models.Text:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Content", inst.Content, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Node", inst.Node, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Node](), probe.formStage)
		StageSetBasicFieldtoForm("PreserveWhiteSpace", inst.PreserveWhiteSpace, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("EnclosingRune", inst.EnclosingRune, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Rune](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Runes {
				if src.Text == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Rune", "Text", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
