// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

)

var (
	_ = time.Hour
	_ = slices.Index[[]int, int]
	_ = sort.Slice
	_ = strconv.Itoa
)

// StageSet coordinates multiple stages across packages
type StageSet struct {
	Stage *Stage
}


// Commit commits all stages in StageSet in dependency order
func (stageSet *StageSet) Commit() {
	if stageSet.Stage != nil {
		stageSet.Stage.Commit()
	}
}

// Checkout checkouts all stages in StageSet
func (stageSet *StageSet) Checkout() {
	if stageSet.Stage != nil {
		stageSet.Stage.Checkout()
	}
}

// Reset resets all stages in StageSet
func (stageSet *StageSet) Reset() {
	if stageSet.Stage != nil {
		stageSet.Stage.Reset()
	}
}

// Clean cleans all stages in StageSet in dependency order
func (stageSet *StageSet) Clean() {
	if stageSet.Stage != nil {
		stageSet.Stage.Clean()
	}
}

// ComputeReverseMaps computes reverse maps on all stages in StageSet
func (stageSet *StageSet) ComputeReverseMaps() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReverseMaps()
	}
}

// ComputeInstancesNb computes instances nb on all stages in StageSet
func (stageSet *StageSet) ComputeInstancesNb() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeInstancesNb()
	}
}

// ComputeReferenceAndOrders computes reference and orders on all stages in StageSet
func (stageSet *StageSet) ComputeReferenceAndOrders() {
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReferenceAndOrders()
	}
}

// NewStageSet creates a StageSet with all stages initialized
func NewStageSet(path string) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = NewStage(path)
	return stageSet
}

// NewStageSetFromStage creates a StageSet using an existing root stage
func NewStageSetFromStage(stage *Stage) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = stage
	return stageSet
}

// GetProbeSplitStageName returns the split stage name for the StageSet probe
func (stageSet *StageSet) GetProbeSplitStageName() string {
	if stageSet.Stage != nil {
		return stageSet.Stage.GetProbeSplitStageName() + "_stageset"
	}
	return "stageset_probe_split"
}

// MarshallFile marshalls all stages into a file
func (stageSet *StageSet) MarshallFile(filename, packageName string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stageSet.Marshall(file, packageName)
}

// Marshall marshalls all stages into an open file
func (stageSet *StageSet) Marshall(file *os.File, packageName string) {
	res, err := stageSet.MarshallToString(packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}
	fmt.Fprintln(file, res)
}

// MarshallToString marshalls all stages into a Go code string
func (stageSet *StageSet) MarshallToString(packageName string) (res string, err error) {
	var declarations strings.Builder
	var values strings.Builder
	var pointers strings.Builder
	var lastStageDecl string
	var lastStageVal string
	var lastStagePtr string
	_ = lastStageDecl
	_ = lastStageVal
	_ = lastStagePtr

	if stageSet.Stage != nil {
		for _, body := range __gong__sortStageSetInstances(stageSet.Stage.Bodys, stageSet.Stage.Body_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bodyIdent := "__models" + body.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Body{Name: %s}).Stage(stageSet.Stage)", bodyIdent, __gong__toRawStringLiteral(body.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bodyIdent, __gong__toRawStringLiteral(body.Name)))
			for _, elem := range body.Paragraphs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Paragraphs = append(%s.Paragraphs, %s)", bodyIdent, bodyIdent, targetIdent))
			}
			for _, elem := range body.Tables {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tables = append(%s.Tables, %s)", bodyIdent, bodyIdent, targetIdent))
			}
			if body.LastParagraph != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + body.LastParagraph.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LastParagraph = %s", bodyIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, document := range __gong__sortStageSetInstances(stageSet.Stage.Documents, stageSet.Stage.Document_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			documentIdent := "__models" + document.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Document{Name: %s}).Stage(stageSet.Stage)", documentIdent, __gong__toRawStringLiteral(document.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", documentIdent, __gong__toRawStringLiteral(document.Name)))
			if document.File != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + document.File.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.File = %s", documentIdent, targetIdent))
			}
			if document.Root != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + document.Root.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Root = %s", documentIdent, targetIdent))
			}
			if document.Body != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + document.Body.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Body = %s", documentIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, docx := range __gong__sortStageSetInstances(stageSet.Stage.Docxs, stageSet.Stage.Docx_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			docxIdent := "__models" + docx.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Docx{Name: %s}).Stage(stageSet.Stage)", docxIdent, __gong__toRawStringLiteral(docx.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", docxIdent, __gong__toRawStringLiteral(docx.Name)))
			for _, elem := range docx.Files {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Files = append(%s.Files, %s)", docxIdent, docxIdent, targetIdent))
			}
			if docx.Document != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + docx.Document.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Document = %s", docxIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, file := range __gong__sortStageSetInstances(stageSet.Stage.Files, stageSet.Stage.File_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			fileIdent := "__models" + file.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.File{Name: %s}).Stage(stageSet.Stage)", fileIdent, __gong__toRawStringLiteral(file.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fileIdent, __gong__toRawStringLiteral(file.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, node := range __gong__sortStageSetInstances(stageSet.Stage.Nodes, stageSet.Stage.Node_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			nodeIdent := "__models" + node.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Node{Name: %s}).Stage(stageSet.Stage)", nodeIdent, __gong__toRawStringLiteral(node.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nodeIdent, __gong__toRawStringLiteral(node.Name)))
			for _, elem := range node.Nodes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Nodes = append(%s.Nodes, %s)", nodeIdent, nodeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, paragraph := range __gong__sortStageSetInstances(stageSet.Stage.Paragraphs, stageSet.Stage.Paragraph_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			paragraphIdent := "__models" + paragraph.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Paragraph{Name: %s}).Stage(stageSet.Stage)", paragraphIdent, __gong__toRawStringLiteral(paragraph.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", paragraphIdent, __gong__toRawStringLiteral(paragraph.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", paragraphIdent, __gong__toRawStringLiteral(paragraph.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.CollatedText = %s", paragraphIdent, __gong__toRawStringLiteral(paragraph.CollatedText)))
			if paragraph.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + paragraph.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", paragraphIdent, targetIdent))
			}
			if paragraph.ParagraphProperties != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + paragraph.ParagraphProperties.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParagraphProperties = %s", paragraphIdent, targetIdent))
			}
			for _, elem := range paragraph.Runes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Runes = append(%s.Runes, %s)", paragraphIdent, paragraphIdent, targetIdent))
			}
			if paragraph.Next != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + paragraph.Next.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Next = %s", paragraphIdent, targetIdent))
			}
			if paragraph.Previous != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + paragraph.Previous.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Previous = %s", paragraphIdent, targetIdent))
			}
			if paragraph.EnclosingBody != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + paragraph.EnclosingBody.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingBody = %s", paragraphIdent, targetIdent))
			}
			if paragraph.EnclosingTableColumn != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + paragraph.EnclosingTableColumn.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingTableColumn = %s", paragraphIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, paragraphproperties := range __gong__sortStageSetInstances(stageSet.Stage.ParagraphPropertiess, stageSet.Stage.ParagraphProperties_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			paragraphpropertiesIdent := "__models" + paragraphproperties.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParagraphProperties{Name: %s}).Stage(stageSet.Stage)", paragraphpropertiesIdent, __gong__toRawStringLiteral(paragraphproperties.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", paragraphpropertiesIdent, __gong__toRawStringLiteral(paragraphproperties.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", paragraphpropertiesIdent, __gong__toRawStringLiteral(paragraphproperties.Content)))
			if paragraphproperties.ParagraphStyle != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + paragraphproperties.ParagraphStyle.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParagraphStyle = %s", paragraphpropertiesIdent, targetIdent))
			}
			if paragraphproperties.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + paragraphproperties.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", paragraphpropertiesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, paragraphstyle := range __gong__sortStageSetInstances(stageSet.Stage.ParagraphStyles, stageSet.Stage.ParagraphStyle_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			paragraphstyleIdent := "__models" + paragraphstyle.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParagraphStyle{Name: %s}).Stage(stageSet.Stage)", paragraphstyleIdent, __gong__toRawStringLiteral(paragraphstyle.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", paragraphstyleIdent, __gong__toRawStringLiteral(paragraphstyle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", paragraphstyleIdent, __gong__toRawStringLiteral(paragraphstyle.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.ValAttr = %s", paragraphstyleIdent, __gong__toRawStringLiteral(paragraphstyle.ValAttr)))
			if paragraphstyle.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + paragraphstyle.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", paragraphstyleIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, rune := range __gong__sortStageSetInstances(stageSet.Stage.Runes, stageSet.Stage.Rune_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			runeIdent := "__models" + rune.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Rune{Name: %s}).Stage(stageSet.Stage)", runeIdent, __gong__toRawStringLiteral(rune.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", runeIdent, __gong__toRawStringLiteral(rune.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", runeIdent, __gong__toRawStringLiteral(rune.Content)))
			if rune.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + rune.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", runeIdent, targetIdent))
			}
			if rune.Text != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + rune.Text.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Text = %s", runeIdent, targetIdent))
			}
			if rune.RuneProperties != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + rune.RuneProperties.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RuneProperties = %s", runeIdent, targetIdent))
			}
			if rune.EnclosingParagraph != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + rune.EnclosingParagraph.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingParagraph = %s", runeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, runeproperties := range __gong__sortStageSetInstances(stageSet.Stage.RunePropertiess, stageSet.Stage.RuneProperties_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			runepropertiesIdent := "__models" + runeproperties.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RuneProperties{Name: %s}).Stage(stageSet.Stage)", runepropertiesIdent, __gong__toRawStringLiteral(runeproperties.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", runepropertiesIdent, __gong__toRawStringLiteral(runeproperties.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsBold = %t", runepropertiesIdent, runeproperties.IsBold))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStrike = %t", runepropertiesIdent, runeproperties.IsStrike))
			values.WriteString(fmt.Sprintf("\n\t%s.IsItalic = %t", runepropertiesIdent, runeproperties.IsItalic))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", runepropertiesIdent, __gong__toRawStringLiteral(runeproperties.Content)))
			if runeproperties.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + runeproperties.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", runepropertiesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, table := range __gong__sortStageSetInstances(stageSet.Stage.Tables, stageSet.Stage.Table_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tableIdent := "__models" + table.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Table{Name: %s}).Stage(stageSet.Stage)", tableIdent, __gong__toRawStringLiteral(table.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tableIdent, __gong__toRawStringLiteral(table.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tableIdent, __gong__toRawStringLiteral(table.Content)))
			if table.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + table.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tableIdent, targetIdent))
			}
			if table.TableProperties != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + table.TableProperties.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TableProperties = %s", tableIdent, targetIdent))
			}
			for _, elem := range table.TableRows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TableRows = append(%s.TableRows, %s)", tableIdent, tableIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, tablecolumn := range __gong__sortStageSetInstances(stageSet.Stage.TableColumns, stageSet.Stage.TableColumn_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tablecolumnIdent := "__models" + tablecolumn.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TableColumn{Name: %s}).Stage(stageSet.Stage)", tablecolumnIdent, __gong__toRawStringLiteral(tablecolumn.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tablecolumnIdent, __gong__toRawStringLiteral(tablecolumn.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tablecolumnIdent, __gong__toRawStringLiteral(tablecolumn.Content)))
			if tablecolumn.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tablecolumn.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tablecolumnIdent, targetIdent))
			}
			for _, elem := range tablecolumn.Paragraphs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Paragraphs = append(%s.Paragraphs, %s)", tablecolumnIdent, tablecolumnIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, tableproperties := range __gong__sortStageSetInstances(stageSet.Stage.TablePropertiess, stageSet.Stage.TableProperties_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tablepropertiesIdent := "__models" + tableproperties.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TableProperties{Name: %s}).Stage(stageSet.Stage)", tablepropertiesIdent, __gong__toRawStringLiteral(tableproperties.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tablepropertiesIdent, __gong__toRawStringLiteral(tableproperties.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tablepropertiesIdent, __gong__toRawStringLiteral(tableproperties.Content)))
			if tableproperties.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tableproperties.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tablepropertiesIdent, targetIdent))
			}
			if tableproperties.TableStyle != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tableproperties.TableStyle.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TableStyle = %s", tablepropertiesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, tablerow := range __gong__sortStageSetInstances(stageSet.Stage.TableRows, stageSet.Stage.TableRow_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tablerowIdent := "__models" + tablerow.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TableRow{Name: %s}).Stage(stageSet.Stage)", tablerowIdent, __gong__toRawStringLiteral(tablerow.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tablerowIdent, __gong__toRawStringLiteral(tablerow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tablerowIdent, __gong__toRawStringLiteral(tablerow.Content)))
			if tablerow.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tablerow.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tablerowIdent, targetIdent))
			}
			for _, elem := range tablerow.TableColumns {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TableColumns = append(%s.TableColumns, %s)", tablerowIdent, tablerowIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, tablestyle := range __gong__sortStageSetInstances(stageSet.Stage.TableStyles, stageSet.Stage.TableStyle_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tablestyleIdent := "__models" + tablestyle.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TableStyle{Name: %s}).Stage(stageSet.Stage)", tablestyleIdent, __gong__toRawStringLiteral(tablestyle.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tablestyleIdent, __gong__toRawStringLiteral(tablestyle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tablestyleIdent, __gong__toRawStringLiteral(tablestyle.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.Val = %s", tablestyleIdent, __gong__toRawStringLiteral(tablestyle.Val)))
			if tablestyle.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tablestyle.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tablestyleIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, text := range __gong__sortStageSetInstances(stageSet.Stage.Texts, stageSet.Stage.Text_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textIdent := "__models" + text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Text{Name: %s}).Stage(stageSet.Stage)", textIdent, __gong__toRawStringLiteral(text.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textIdent, __gong__toRawStringLiteral(text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", textIdent, __gong__toRawStringLiteral(text.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.PreserveWhiteSpace = %t", textIdent, text.PreserveWhiteSpace))
			if text.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + text.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", textIdent, targetIdent))
			}
			if text.EnclosingRune != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + text.EnclosingRune.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingRune = %s", textIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/docx/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *models.StageSet) {

	// ------------------------------------------------------------------------
	// Phase 1: Declarations (in topological order: leaves first)
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 2: Value Initializations
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)
	// ------------------------------------------------------------------------%s
}
`, packageName, declarations.String(), values.String(), pointers.String())

	return res, nil
}

// ParseAstFile Parse pathToFile and stages all instances declared in the file into stageSet
func (stageSet *StageSet) ParseAstFile(pathToFile string, preserveOrder bool) error {
	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist " + pathToFile + " ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file into stageSet
func (stageSet *StageSet) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		return errors.New("Unable to read embedded file " + err.Error())
	}

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string into stageSet
func (stageSet *StageSet) ParseAstString(blob string, preserveOrder bool) error {
	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", blob, parser.ParseComments)
	if errParser != nil {
		return errors.New("Unable to parse " + errParser.Error())
	}

	return stageSet.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances into stageSet
func (stageSet *StageSet) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	identifierMap := make(map[string]any)

	aliasToCanonical := make(map[string]string)
	for _, imp := range inFile.Imports {
		p := strings.Trim(imp.Path.Value, "\"`")
		alias := filepath.Base(p)
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		switch p {
		case "github.com/fullstack-lang/gong/lib/docx/go/models":
			aliasToCanonical[alias] = "models"
		}
	}

	ast.Inspect(inFile, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			if len(node.Lhs) < 1 || len(node.Rhs) < 1 {
				return true
			}

			// CASE 1: Initialization ( := )
			if node.Tok == token.DEFINE {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok {
					var pkgAlias string
					var typeName string
					var instanceName string

					ast.Inspect(node.Rhs[0], func(expr ast.Node) bool {
						if compLit, ok := expr.(*ast.CompositeLit); ok {
							if selExpr, ok := compLit.Type.(*ast.SelectorExpr); ok {
								if pkgId, ok := selExpr.X.(*ast.Ident); ok {
									pkgAlias = pkgId.Name
								}
								typeName = selExpr.Sel.Name
								for _, elt := range compLit.Elts {
									if kv, ok := elt.(*ast.KeyValueExpr); ok {
										if k, ok := kv.Key.(*ast.Ident); ok && k.Name == "Name" {
											if v, ok := kv.Value.(*ast.BasicLit); ok {
												instanceName = strings.Trim(v.Value, "\"`")
											}
										}
									}
								}
								return false
							}
						}
						return true
					})

					if canonical, ok := aliasToCanonical[pkgAlias]; ok {
						pkgAlias = canonical
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "Body":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Body), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Document":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Document), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Docx":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Docx), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "File":
					identifierMap[ident.Name] = __gong__stageSetInit(new(File), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Node":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Node), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Paragraph":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Paragraph), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ParagraphProperties":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParagraphProperties), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ParagraphStyle":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParagraphStyle), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Rune":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Rune), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "RuneProperties":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RuneProperties), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Table":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Table), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TableColumn":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TableColumn), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TableProperties":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TableProperties), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TableRow":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TableRow), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TableStyle":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TableStyle), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Text), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				}
					}
				}
				return false
			}

			// CASE 2: Assignment ( = )
			if node.Tok == token.ASSIGN {
				if selExpr, ok := node.Lhs[0].(*ast.SelectorExpr); ok {
					if ident, ok := selExpr.X.(*ast.Ident); ok {
						if instance, exists := identifierMap[ident.Name]; exists {
							fieldName := selExpr.Sel.Name
							rhs := node.Rhs[0]
							switch inst := instance.(type) {
				case *Body:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Paragraphs":
						__gong__assignSliceOfPointers(&inst.Paragraphs, rhs, identifierMap)
					case "Tables":
						__gong__assignSliceOfPointers(&inst.Tables, rhs, identifierMap)
					case "LastParagraph":
						__gong__assignPointer(&inst.LastParagraph, rhs, identifierMap)
					}
				case *Document:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "File":
						__gong__assignPointer(&inst.File, rhs, identifierMap)
					case "Root":
						__gong__assignPointer(&inst.Root, rhs, identifierMap)
					case "Body":
						__gong__assignPointer(&inst.Body, rhs, identifierMap)
					}
				case *Docx:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Files":
						__gong__assignSliceOfPointers(&inst.Files, rhs, identifierMap)
					case "Document":
						__gong__assignPointer(&inst.Document, rhs, identifierMap)
					}
				case *File:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Node:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Nodes":
						__gong__assignSliceOfPointers(&inst.Nodes, rhs, identifierMap)
					}
				case *Paragraph:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "Node":
						__gong__assignPointer(&inst.Node, rhs, identifierMap)
					case "ParagraphProperties":
						__gong__assignPointer(&inst.ParagraphProperties, rhs, identifierMap)
					case "Runes":
						__gong__assignSliceOfPointers(&inst.Runes, rhs, identifierMap)
					case "CollatedText":
						inst.CollatedText = GongExtractString(rhs)
					case "Next":
						__gong__assignPointer(&inst.Next, rhs, identifierMap)
					case "Previous":
						__gong__assignPointer(&inst.Previous, rhs, identifierMap)
					case "EnclosingBody":
						__gong__assignPointer(&inst.EnclosingBody, rhs, identifierMap)
					case "EnclosingTableColumn":
						__gong__assignPointer(&inst.EnclosingTableColumn, rhs, identifierMap)
					}
				case *ParagraphProperties:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "ParagraphStyle":
						__gong__assignPointer(&inst.ParagraphStyle, rhs, identifierMap)
					case "Node":
						__gong__assignPointer(&inst.Node, rhs, identifierMap)
					}
				case *ParagraphStyle:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Node":
						__gong__assignPointer(&inst.Node, rhs, identifierMap)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "ValAttr":
						inst.ValAttr = GongExtractString(rhs)
					}
				case *Rune:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "Node":
						__gong__assignPointer(&inst.Node, rhs, identifierMap)
					case "Text":
						__gong__assignPointer(&inst.Text, rhs, identifierMap)
					case "RuneProperties":
						__gong__assignPointer(&inst.RuneProperties, rhs, identifierMap)
					case "EnclosingParagraph":
						__gong__assignPointer(&inst.EnclosingParagraph, rhs, identifierMap)
					}
				case *RuneProperties:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Node":
						__gong__assignPointer(&inst.Node, rhs, identifierMap)
					case "IsBold":
						inst.IsBold = GongExtractBool(rhs)
					case "IsStrike":
						inst.IsStrike = GongExtractBool(rhs)
					case "IsItalic":
						inst.IsItalic = GongExtractBool(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					}
				case *Table:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Node":
						__gong__assignPointer(&inst.Node, rhs, identifierMap)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "TableProperties":
						__gong__assignPointer(&inst.TableProperties, rhs, identifierMap)
					case "TableRows":
						__gong__assignSliceOfPointers(&inst.TableRows, rhs, identifierMap)
					}
				case *TableColumn:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "Node":
						__gong__assignPointer(&inst.Node, rhs, identifierMap)
					case "Paragraphs":
						__gong__assignSliceOfPointers(&inst.Paragraphs, rhs, identifierMap)
					}
				case *TableProperties:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Node":
						__gong__assignPointer(&inst.Node, rhs, identifierMap)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "TableStyle":
						__gong__assignPointer(&inst.TableStyle, rhs, identifierMap)
					}
				case *TableRow:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "Node":
						__gong__assignPointer(&inst.Node, rhs, identifierMap)
					case "TableColumns":
						__gong__assignSliceOfPointers(&inst.TableColumns, rhs, identifierMap)
					}
				case *TableStyle:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Node":
						__gong__assignPointer(&inst.Node, rhs, identifierMap)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "Val":
						inst.Val = GongExtractString(rhs)
					}
				case *Text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "Node":
						__gong__assignPointer(&inst.Node, rhs, identifierMap)
					case "PreserveWhiteSpace":
						inst.PreserveWhiteSpace = GongExtractBool(rhs)
					case "EnclosingRune":
						__gong__assignPointer(&inst.EnclosingRune, rhs, identifierMap)
					}
							}
						}
					}
				}
			}
		}
		return true
	})

	return nil
}

// __gong__sortStageSetInstances sorts instances by their staged order
func __gong__sortStageSetInstances[T comparable](instances map[T]struct{}, orderMap map[T]uint) []T {
	ordered := make([]T, 0, len(instances))
	for inst := range instances {
		ordered = append(ordered, inst)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return orderMap[ordered[i]] < orderMap[ordered[j]]
	})
	return ordered
}

func __gong__stageSetInit[P interface {
	SetName(string)
	StageVoid(S)
	StagePreserveOrder(S, uint)
}, S any](instance P, stage S, identifier string, instanceName string, preserveOrder bool) any {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if order, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifier", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, order)
		}
	}
	return instance
}

func __gong__assignPointer[T any](targetPtr **T, rhs ast.Expr, identifierMap map[string]any) {
	if rIdent, ok := rhs.(*ast.Ident); ok {
		if rIdent.Name == "nil" {
			*targetPtr = nil
			return
		}
		if target, ok := identifierMap[rIdent.Name]; ok {
			if typedTarget, ok := target.(*T); ok {
				*targetPtr = typedTarget
			}
		}
	}
}

func __gong__assignSliceOfPointers[T any](slice *[]*T, rhs ast.Expr, identifierMap map[string]any) {
	if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
		if rIdent, ok := call.Args[1].(*ast.Ident); ok {
			if target, ok := identifierMap[rIdent.Name]; ok {
				if typedTarget, ok := target.(*T); ok {
					*slice = append(*slice, typedTarget)
				}
			}
		}
	}
}
