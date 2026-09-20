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

	if stageSet.Stage != nil {
		bodyOrdered := []*Body{}
		for body := range stageSet.Stage.Bodys {
			bodyOrdered = append(bodyOrdered, body)
		}
		sort.Slice(bodyOrdered, func(i, j int) bool {
			return stageSet.Stage.Body_stagedOrder[bodyOrdered[i]] < stageSet.Stage.Body_stagedOrder[bodyOrdered[j]]
		})
		for _, body := range bodyOrdered {
			bodyIdent := "__stage_0" + body.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Body{Name: %s}).Stage(stageSet.Stage)", bodyIdent, __gong__toRawStringLiteral(body.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bodyIdent, __gong__toRawStringLiteral(body.Name)))
			for _, elem := range body.Paragraphs {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Paragraphs = append(%s.Paragraphs, %s)", bodyIdent, bodyIdent, targetIdent))
			}
			for _, elem := range body.Tables {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tables = append(%s.Tables, %s)", bodyIdent, bodyIdent, targetIdent))
			}
			if body.LastParagraph != nil {
				targetIdent := "__stage_0" + body.LastParagraph.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LastParagraph = %s", bodyIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		documentOrdered := []*Document{}
		for document := range stageSet.Stage.Documents {
			documentOrdered = append(documentOrdered, document)
		}
		sort.Slice(documentOrdered, func(i, j int) bool {
			return stageSet.Stage.Document_stagedOrder[documentOrdered[i]] < stageSet.Stage.Document_stagedOrder[documentOrdered[j]]
		})
		for _, document := range documentOrdered {
			documentIdent := "__stage_0" + document.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Document{Name: %s}).Stage(stageSet.Stage)", documentIdent, __gong__toRawStringLiteral(document.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", documentIdent, __gong__toRawStringLiteral(document.Name)))
			if document.File != nil {
				targetIdent := "__stage_0" + document.File.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.File = %s", documentIdent, targetIdent))
			}
			if document.Root != nil {
				targetIdent := "__stage_0" + document.Root.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Root = %s", documentIdent, targetIdent))
			}
			if document.Body != nil {
				targetIdent := "__stage_0" + document.Body.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Body = %s", documentIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		docxOrdered := []*Docx{}
		for docx := range stageSet.Stage.Docxs {
			docxOrdered = append(docxOrdered, docx)
		}
		sort.Slice(docxOrdered, func(i, j int) bool {
			return stageSet.Stage.Docx_stagedOrder[docxOrdered[i]] < stageSet.Stage.Docx_stagedOrder[docxOrdered[j]]
		})
		for _, docx := range docxOrdered {
			docxIdent := "__stage_0" + docx.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Docx{Name: %s}).Stage(stageSet.Stage)", docxIdent, __gong__toRawStringLiteral(docx.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", docxIdent, __gong__toRawStringLiteral(docx.Name)))
			for _, elem := range docx.Files {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Files = append(%s.Files, %s)", docxIdent, docxIdent, targetIdent))
			}
			if docx.Document != nil {
				targetIdent := "__stage_0" + docx.Document.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Document = %s", docxIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		fileOrdered := []*File{}
		for file := range stageSet.Stage.Files {
			fileOrdered = append(fileOrdered, file)
		}
		sort.Slice(fileOrdered, func(i, j int) bool {
			return stageSet.Stage.File_stagedOrder[fileOrdered[i]] < stageSet.Stage.File_stagedOrder[fileOrdered[j]]
		})
		for _, file := range fileOrdered {
			fileIdent := "__stage_0" + file.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.File{Name: %s}).Stage(stageSet.Stage)", fileIdent, __gong__toRawStringLiteral(file.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fileIdent, __gong__toRawStringLiteral(file.Name)))
		}
	}
	if stageSet.Stage != nil {
		nodeOrdered := []*Node{}
		for node := range stageSet.Stage.Nodes {
			nodeOrdered = append(nodeOrdered, node)
		}
		sort.Slice(nodeOrdered, func(i, j int) bool {
			return stageSet.Stage.Node_stagedOrder[nodeOrdered[i]] < stageSet.Stage.Node_stagedOrder[nodeOrdered[j]]
		})
		for _, node := range nodeOrdered {
			nodeIdent := "__stage_0" + node.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Node{Name: %s}).Stage(stageSet.Stage)", nodeIdent, __gong__toRawStringLiteral(node.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nodeIdent, __gong__toRawStringLiteral(node.Name)))
			for _, elem := range node.Nodes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Nodes = append(%s.Nodes, %s)", nodeIdent, nodeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		paragraphOrdered := []*Paragraph{}
		for paragraph := range stageSet.Stage.Paragraphs {
			paragraphOrdered = append(paragraphOrdered, paragraph)
		}
		sort.Slice(paragraphOrdered, func(i, j int) bool {
			return stageSet.Stage.Paragraph_stagedOrder[paragraphOrdered[i]] < stageSet.Stage.Paragraph_stagedOrder[paragraphOrdered[j]]
		})
		for _, paragraph := range paragraphOrdered {
			paragraphIdent := "__stage_0" + paragraph.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Paragraph{Name: %s}).Stage(stageSet.Stage)", paragraphIdent, __gong__toRawStringLiteral(paragraph.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", paragraphIdent, __gong__toRawStringLiteral(paragraph.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", paragraphIdent, __gong__toRawStringLiteral(paragraph.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.CollatedText = %s", paragraphIdent, __gong__toRawStringLiteral(paragraph.CollatedText)))
			if paragraph.Node != nil {
				targetIdent := "__stage_0" + paragraph.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", paragraphIdent, targetIdent))
			}
			if paragraph.ParagraphProperties != nil {
				targetIdent := "__stage_0" + paragraph.ParagraphProperties.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParagraphProperties = %s", paragraphIdent, targetIdent))
			}
			for _, elem := range paragraph.Runes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Runes = append(%s.Runes, %s)", paragraphIdent, paragraphIdent, targetIdent))
			}
			if paragraph.Next != nil {
				targetIdent := "__stage_0" + paragraph.Next.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Next = %s", paragraphIdent, targetIdent))
			}
			if paragraph.Previous != nil {
				targetIdent := "__stage_0" + paragraph.Previous.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Previous = %s", paragraphIdent, targetIdent))
			}
			if paragraph.EnclosingBody != nil {
				targetIdent := "__stage_0" + paragraph.EnclosingBody.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingBody = %s", paragraphIdent, targetIdent))
			}
			if paragraph.EnclosingTableColumn != nil {
				targetIdent := "__stage_0" + paragraph.EnclosingTableColumn.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingTableColumn = %s", paragraphIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		paragraphpropertiesOrdered := []*ParagraphProperties{}
		for paragraphproperties := range stageSet.Stage.ParagraphPropertiess {
			paragraphpropertiesOrdered = append(paragraphpropertiesOrdered, paragraphproperties)
		}
		sort.Slice(paragraphpropertiesOrdered, func(i, j int) bool {
			return stageSet.Stage.ParagraphProperties_stagedOrder[paragraphpropertiesOrdered[i]] < stageSet.Stage.ParagraphProperties_stagedOrder[paragraphpropertiesOrdered[j]]
		})
		for _, paragraphproperties := range paragraphpropertiesOrdered {
			paragraphpropertiesIdent := "__stage_0" + paragraphproperties.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ParagraphProperties{Name: %s}).Stage(stageSet.Stage)", paragraphpropertiesIdent, __gong__toRawStringLiteral(paragraphproperties.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", paragraphpropertiesIdent, __gong__toRawStringLiteral(paragraphproperties.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", paragraphpropertiesIdent, __gong__toRawStringLiteral(paragraphproperties.Content)))
			if paragraphproperties.ParagraphStyle != nil {
				targetIdent := "__stage_0" + paragraphproperties.ParagraphStyle.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParagraphStyle = %s", paragraphpropertiesIdent, targetIdent))
			}
			if paragraphproperties.Node != nil {
				targetIdent := "__stage_0" + paragraphproperties.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", paragraphpropertiesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		paragraphstyleOrdered := []*ParagraphStyle{}
		for paragraphstyle := range stageSet.Stage.ParagraphStyles {
			paragraphstyleOrdered = append(paragraphstyleOrdered, paragraphstyle)
		}
		sort.Slice(paragraphstyleOrdered, func(i, j int) bool {
			return stageSet.Stage.ParagraphStyle_stagedOrder[paragraphstyleOrdered[i]] < stageSet.Stage.ParagraphStyle_stagedOrder[paragraphstyleOrdered[j]]
		})
		for _, paragraphstyle := range paragraphstyleOrdered {
			paragraphstyleIdent := "__stage_0" + paragraphstyle.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ParagraphStyle{Name: %s}).Stage(stageSet.Stage)", paragraphstyleIdent, __gong__toRawStringLiteral(paragraphstyle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", paragraphstyleIdent, __gong__toRawStringLiteral(paragraphstyle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", paragraphstyleIdent, __gong__toRawStringLiteral(paragraphstyle.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.ValAttr = %s", paragraphstyleIdent, __gong__toRawStringLiteral(paragraphstyle.ValAttr)))
			if paragraphstyle.Node != nil {
				targetIdent := "__stage_0" + paragraphstyle.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", paragraphstyleIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		runeOrdered := []*Rune{}
		for rune := range stageSet.Stage.Runes {
			runeOrdered = append(runeOrdered, rune)
		}
		sort.Slice(runeOrdered, func(i, j int) bool {
			return stageSet.Stage.Rune_stagedOrder[runeOrdered[i]] < stageSet.Stage.Rune_stagedOrder[runeOrdered[j]]
		})
		for _, rune := range runeOrdered {
			runeIdent := "__stage_0" + rune.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Rune{Name: %s}).Stage(stageSet.Stage)", runeIdent, __gong__toRawStringLiteral(rune.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", runeIdent, __gong__toRawStringLiteral(rune.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", runeIdent, __gong__toRawStringLiteral(rune.Content)))
			if rune.Node != nil {
				targetIdent := "__stage_0" + rune.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", runeIdent, targetIdent))
			}
			if rune.Text != nil {
				targetIdent := "__stage_0" + rune.Text.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Text = %s", runeIdent, targetIdent))
			}
			if rune.RuneProperties != nil {
				targetIdent := "__stage_0" + rune.RuneProperties.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RuneProperties = %s", runeIdent, targetIdent))
			}
			if rune.EnclosingParagraph != nil {
				targetIdent := "__stage_0" + rune.EnclosingParagraph.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingParagraph = %s", runeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		runepropertiesOrdered := []*RuneProperties{}
		for runeproperties := range stageSet.Stage.RunePropertiess {
			runepropertiesOrdered = append(runepropertiesOrdered, runeproperties)
		}
		sort.Slice(runepropertiesOrdered, func(i, j int) bool {
			return stageSet.Stage.RuneProperties_stagedOrder[runepropertiesOrdered[i]] < stageSet.Stage.RuneProperties_stagedOrder[runepropertiesOrdered[j]]
		})
		for _, runeproperties := range runepropertiesOrdered {
			runepropertiesIdent := "__stage_0" + runeproperties.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.RuneProperties{Name: %s}).Stage(stageSet.Stage)", runepropertiesIdent, __gong__toRawStringLiteral(runeproperties.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", runepropertiesIdent, __gong__toRawStringLiteral(runeproperties.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsBold = %t", runepropertiesIdent, runeproperties.IsBold))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStrike = %t", runepropertiesIdent, runeproperties.IsStrike))
			values.WriteString(fmt.Sprintf("\n\t%s.IsItalic = %t", runepropertiesIdent, runeproperties.IsItalic))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", runepropertiesIdent, __gong__toRawStringLiteral(runeproperties.Content)))
			if runeproperties.Node != nil {
				targetIdent := "__stage_0" + runeproperties.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", runepropertiesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		tableOrdered := []*Table{}
		for table := range stageSet.Stage.Tables {
			tableOrdered = append(tableOrdered, table)
		}
		sort.Slice(tableOrdered, func(i, j int) bool {
			return stageSet.Stage.Table_stagedOrder[tableOrdered[i]] < stageSet.Stage.Table_stagedOrder[tableOrdered[j]]
		})
		for _, table := range tableOrdered {
			tableIdent := "__stage_0" + table.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Table{Name: %s}).Stage(stageSet.Stage)", tableIdent, __gong__toRawStringLiteral(table.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tableIdent, __gong__toRawStringLiteral(table.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tableIdent, __gong__toRawStringLiteral(table.Content)))
			if table.Node != nil {
				targetIdent := "__stage_0" + table.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tableIdent, targetIdent))
			}
			if table.TableProperties != nil {
				targetIdent := "__stage_0" + table.TableProperties.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TableProperties = %s", tableIdent, targetIdent))
			}
			for _, elem := range table.TableRows {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TableRows = append(%s.TableRows, %s)", tableIdent, tableIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		tablecolumnOrdered := []*TableColumn{}
		for tablecolumn := range stageSet.Stage.TableColumns {
			tablecolumnOrdered = append(tablecolumnOrdered, tablecolumn)
		}
		sort.Slice(tablecolumnOrdered, func(i, j int) bool {
			return stageSet.Stage.TableColumn_stagedOrder[tablecolumnOrdered[i]] < stageSet.Stage.TableColumn_stagedOrder[tablecolumnOrdered[j]]
		})
		for _, tablecolumn := range tablecolumnOrdered {
			tablecolumnIdent := "__stage_0" + tablecolumn.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.TableColumn{Name: %s}).Stage(stageSet.Stage)", tablecolumnIdent, __gong__toRawStringLiteral(tablecolumn.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tablecolumnIdent, __gong__toRawStringLiteral(tablecolumn.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tablecolumnIdent, __gong__toRawStringLiteral(tablecolumn.Content)))
			if tablecolumn.Node != nil {
				targetIdent := "__stage_0" + tablecolumn.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tablecolumnIdent, targetIdent))
			}
			for _, elem := range tablecolumn.Paragraphs {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Paragraphs = append(%s.Paragraphs, %s)", tablecolumnIdent, tablecolumnIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		tablepropertiesOrdered := []*TableProperties{}
		for tableproperties := range stageSet.Stage.TablePropertiess {
			tablepropertiesOrdered = append(tablepropertiesOrdered, tableproperties)
		}
		sort.Slice(tablepropertiesOrdered, func(i, j int) bool {
			return stageSet.Stage.TableProperties_stagedOrder[tablepropertiesOrdered[i]] < stageSet.Stage.TableProperties_stagedOrder[tablepropertiesOrdered[j]]
		})
		for _, tableproperties := range tablepropertiesOrdered {
			tablepropertiesIdent := "__stage_0" + tableproperties.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.TableProperties{Name: %s}).Stage(stageSet.Stage)", tablepropertiesIdent, __gong__toRawStringLiteral(tableproperties.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tablepropertiesIdent, __gong__toRawStringLiteral(tableproperties.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tablepropertiesIdent, __gong__toRawStringLiteral(tableproperties.Content)))
			if tableproperties.Node != nil {
				targetIdent := "__stage_0" + tableproperties.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tablepropertiesIdent, targetIdent))
			}
			if tableproperties.TableStyle != nil {
				targetIdent := "__stage_0" + tableproperties.TableStyle.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TableStyle = %s", tablepropertiesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		tablerowOrdered := []*TableRow{}
		for tablerow := range stageSet.Stage.TableRows {
			tablerowOrdered = append(tablerowOrdered, tablerow)
		}
		sort.Slice(tablerowOrdered, func(i, j int) bool {
			return stageSet.Stage.TableRow_stagedOrder[tablerowOrdered[i]] < stageSet.Stage.TableRow_stagedOrder[tablerowOrdered[j]]
		})
		for _, tablerow := range tablerowOrdered {
			tablerowIdent := "__stage_0" + tablerow.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.TableRow{Name: %s}).Stage(stageSet.Stage)", tablerowIdent, __gong__toRawStringLiteral(tablerow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tablerowIdent, __gong__toRawStringLiteral(tablerow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tablerowIdent, __gong__toRawStringLiteral(tablerow.Content)))
			if tablerow.Node != nil {
				targetIdent := "__stage_0" + tablerow.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tablerowIdent, targetIdent))
			}
			for _, elem := range tablerow.TableColumns {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TableColumns = append(%s.TableColumns, %s)", tablerowIdent, tablerowIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		tablestyleOrdered := []*TableStyle{}
		for tablestyle := range stageSet.Stage.TableStyles {
			tablestyleOrdered = append(tablestyleOrdered, tablestyle)
		}
		sort.Slice(tablestyleOrdered, func(i, j int) bool {
			return stageSet.Stage.TableStyle_stagedOrder[tablestyleOrdered[i]] < stageSet.Stage.TableStyle_stagedOrder[tablestyleOrdered[j]]
		})
		for _, tablestyle := range tablestyleOrdered {
			tablestyleIdent := "__stage_0" + tablestyle.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.TableStyle{Name: %s}).Stage(stageSet.Stage)", tablestyleIdent, __gong__toRawStringLiteral(tablestyle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tablestyleIdent, __gong__toRawStringLiteral(tablestyle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tablestyleIdent, __gong__toRawStringLiteral(tablestyle.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.Val = %s", tablestyleIdent, __gong__toRawStringLiteral(tablestyle.Val)))
			if tablestyle.Node != nil {
				targetIdent := "__stage_0" + tablestyle.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tablestyleIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		textOrdered := []*Text{}
		for text := range stageSet.Stage.Texts {
			textOrdered = append(textOrdered, text)
		}
		sort.Slice(textOrdered, func(i, j int) bool {
			return stageSet.Stage.Text_stagedOrder[textOrdered[i]] < stageSet.Stage.Text_stagedOrder[textOrdered[j]]
		})
		for _, text := range textOrdered {
			textIdent := "__stage_0" + text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Text{Name: %s}).Stage(stageSet.Stage)", textIdent, __gong__toRawStringLiteral(text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textIdent, __gong__toRawStringLiteral(text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", textIdent, __gong__toRawStringLiteral(text.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.PreserveWhiteSpace = %t", textIdent, text.PreserveWhiteSpace))
			if text.Node != nil {
				targetIdent := "__stage_0" + text.Node.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", textIdent, targetIdent))
			}
			if text.EnclosingRune != nil {
				targetIdent := "__stage_0" + text.EnclosingRune.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingRune = %s", textIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/lib/docx/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *__stage_0__.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *__stage_0__.StageSet) {

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

					switch pkgAlias {
			case "__stage_0__":
				switch typeName {
				case "Body":
					if !preserveOrder {
						inst := (&Body{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Body)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Document":
					if !preserveOrder {
						inst := (&Document{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Document)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Docx":
					if !preserveOrder {
						inst := (&Docx{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Docx)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "File":
					if !preserveOrder {
						inst := (&File{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(File)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Node":
					if !preserveOrder {
						inst := (&Node{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Node)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Paragraph":
					if !preserveOrder {
						inst := (&Paragraph{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Paragraph)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ParagraphProperties":
					if !preserveOrder {
						inst := (&ParagraphProperties{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ParagraphProperties)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ParagraphStyle":
					if !preserveOrder {
						inst := (&ParagraphStyle{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ParagraphStyle)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Rune":
					if !preserveOrder {
						inst := (&Rune{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Rune)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "RuneProperties":
					if !preserveOrder {
						inst := (&RuneProperties{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(RuneProperties)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Table":
					if !preserveOrder {
						inst := (&Table{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Table)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TableColumn":
					if !preserveOrder {
						inst := (&TableColumn{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TableColumn)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TableProperties":
					if !preserveOrder {
						inst := (&TableProperties{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TableProperties)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TableRow":
					if !preserveOrder {
						inst := (&TableRow{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TableRow)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TableStyle":
					if !preserveOrder {
						inst := (&TableStyle{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TableStyle)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Text":
					if !preserveOrder {
						inst := (&Text{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Text)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Paragraph); ok {
										inst.Paragraphs = append(inst.Paragraphs, typedTarget)
									}
								}
							}
						}
					case "Tables":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Table); ok {
										inst.Tables = append(inst.Tables, typedTarget)
									}
								}
							}
						}
					case "LastParagraph":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Paragraph); ok {
									inst.LastParagraph = typedTarget
								}
							}
						}
					}
				case *Document:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "File":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*File); ok {
									inst.File = typedTarget
								}
							}
						}
					case "Root":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Node); ok {
									inst.Root = typedTarget
								}
							}
						}
					case "Body":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Body); ok {
									inst.Body = typedTarget
								}
							}
						}
					}
				case *Docx:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Files":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*File); ok {
										inst.Files = append(inst.Files, typedTarget)
									}
								}
							}
						}
					case "Document":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Document); ok {
									inst.Document = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Node); ok {
										inst.Nodes = append(inst.Nodes, typedTarget)
									}
								}
							}
						}
					}
				case *Paragraph:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "Node":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Node); ok {
									inst.Node = typedTarget
								}
							}
						}
					case "ParagraphProperties":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ParagraphProperties); ok {
									inst.ParagraphProperties = typedTarget
								}
							}
						}
					case "Runes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Rune); ok {
										inst.Runes = append(inst.Runes, typedTarget)
									}
								}
							}
						}
					case "CollatedText":
						inst.CollatedText = GongExtractString(rhs)
					case "Next":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Paragraph); ok {
									inst.Next = typedTarget
								}
							}
						}
					case "Previous":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Paragraph); ok {
									inst.Previous = typedTarget
								}
							}
						}
					case "EnclosingBody":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Body); ok {
									inst.EnclosingBody = typedTarget
								}
							}
						}
					case "EnclosingTableColumn":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TableColumn); ok {
									inst.EnclosingTableColumn = typedTarget
								}
							}
						}
					}
				case *ParagraphProperties:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "ParagraphStyle":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ParagraphStyle); ok {
									inst.ParagraphStyle = typedTarget
								}
							}
						}
					case "Node":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Node); ok {
									inst.Node = typedTarget
								}
							}
						}
					}
				case *ParagraphStyle:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Node":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Node); ok {
									inst.Node = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Node); ok {
									inst.Node = typedTarget
								}
							}
						}
					case "Text":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Text); ok {
									inst.Text = typedTarget
								}
							}
						}
					case "RuneProperties":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*RuneProperties); ok {
									inst.RuneProperties = typedTarget
								}
							}
						}
					case "EnclosingParagraph":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Paragraph); ok {
									inst.EnclosingParagraph = typedTarget
								}
							}
						}
					}
				case *RuneProperties:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Node":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Node); ok {
									inst.Node = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Node); ok {
									inst.Node = typedTarget
								}
							}
						}
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "TableProperties":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TableProperties); ok {
									inst.TableProperties = typedTarget
								}
							}
						}
					case "TableRows":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TableRow); ok {
										inst.TableRows = append(inst.TableRows, typedTarget)
									}
								}
							}
						}
					}
				case *TableColumn:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "Node":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Node); ok {
									inst.Node = typedTarget
								}
							}
						}
					case "Paragraphs":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Paragraph); ok {
										inst.Paragraphs = append(inst.Paragraphs, typedTarget)
									}
								}
							}
						}
					}
				case *TableProperties:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Node":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Node); ok {
									inst.Node = typedTarget
								}
							}
						}
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "TableStyle":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TableStyle); ok {
									inst.TableStyle = typedTarget
								}
							}
						}
					}
				case *TableRow:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "Node":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Node); ok {
									inst.Node = typedTarget
								}
							}
						}
					case "TableColumns":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TableColumn); ok {
										inst.TableColumns = append(inst.TableColumns, typedTarget)
									}
								}
							}
						}
					}
				case *TableStyle:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Node":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Node); ok {
									inst.Node = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Node); ok {
									inst.Node = typedTarget
								}
							}
						}
					case "PreserveWhiteSpace":
						inst.PreserveWhiteSpace = GongExtractBool(rhs)
					case "EnclosingRune":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Rune); ok {
									inst.EnclosingRune = typedTarget
								}
							}
						}
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
