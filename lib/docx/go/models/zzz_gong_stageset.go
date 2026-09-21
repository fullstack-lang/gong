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
		bodyOrdered := []*Body{}
		for body := range stageSet.Stage.Bodys {
			bodyOrdered = append(bodyOrdered, body)
		}
		sort.Slice(bodyOrdered, func(i, j int) bool {
			return stageSet.Stage.Body_stagedOrder[bodyOrdered[i]] < stageSet.Stage.Body_stagedOrder[bodyOrdered[j]]
		})
		for _, body := range bodyOrdered {
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
		documentOrdered := []*Document{}
		for document := range stageSet.Stage.Documents {
			documentOrdered = append(documentOrdered, document)
		}
		sort.Slice(documentOrdered, func(i, j int) bool {
			return stageSet.Stage.Document_stagedOrder[documentOrdered[i]] < stageSet.Stage.Document_stagedOrder[documentOrdered[j]]
		})
		for _, document := range documentOrdered {
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
		docxOrdered := []*Docx{}
		for docx := range stageSet.Stage.Docxs {
			docxOrdered = append(docxOrdered, docx)
		}
		sort.Slice(docxOrdered, func(i, j int) bool {
			return stageSet.Stage.Docx_stagedOrder[docxOrdered[i]] < stageSet.Stage.Docx_stagedOrder[docxOrdered[j]]
		})
		for _, docx := range docxOrdered {
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
		fileOrdered := []*File{}
		for file := range stageSet.Stage.Files {
			fileOrdered = append(fileOrdered, file)
		}
		sort.Slice(fileOrdered, func(i, j int) bool {
			return stageSet.Stage.File_stagedOrder[fileOrdered[i]] < stageSet.Stage.File_stagedOrder[fileOrdered[j]]
		})
		for _, file := range fileOrdered {
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
		nodeOrdered := []*Node{}
		for node := range stageSet.Stage.Nodes {
			nodeOrdered = append(nodeOrdered, node)
		}
		sort.Slice(nodeOrdered, func(i, j int) bool {
			return stageSet.Stage.Node_stagedOrder[nodeOrdered[i]] < stageSet.Stage.Node_stagedOrder[nodeOrdered[j]]
		})
		for _, node := range nodeOrdered {
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
		paragraphOrdered := []*Paragraph{}
		for paragraph := range stageSet.Stage.Paragraphs {
			paragraphOrdered = append(paragraphOrdered, paragraph)
		}
		sort.Slice(paragraphOrdered, func(i, j int) bool {
			return stageSet.Stage.Paragraph_stagedOrder[paragraphOrdered[i]] < stageSet.Stage.Paragraph_stagedOrder[paragraphOrdered[j]]
		})
		for _, paragraph := range paragraphOrdered {
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
		paragraphpropertiesOrdered := []*ParagraphProperties{}
		for paragraphproperties := range stageSet.Stage.ParagraphPropertiess {
			paragraphpropertiesOrdered = append(paragraphpropertiesOrdered, paragraphproperties)
		}
		sort.Slice(paragraphpropertiesOrdered, func(i, j int) bool {
			return stageSet.Stage.ParagraphProperties_stagedOrder[paragraphpropertiesOrdered[i]] < stageSet.Stage.ParagraphProperties_stagedOrder[paragraphpropertiesOrdered[j]]
		})
		for _, paragraphproperties := range paragraphpropertiesOrdered {
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
		paragraphstyleOrdered := []*ParagraphStyle{}
		for paragraphstyle := range stageSet.Stage.ParagraphStyles {
			paragraphstyleOrdered = append(paragraphstyleOrdered, paragraphstyle)
		}
		sort.Slice(paragraphstyleOrdered, func(i, j int) bool {
			return stageSet.Stage.ParagraphStyle_stagedOrder[paragraphstyleOrdered[i]] < stageSet.Stage.ParagraphStyle_stagedOrder[paragraphstyleOrdered[j]]
		})
		for _, paragraphstyle := range paragraphstyleOrdered {
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
		runeOrdered := []*Rune{}
		for rune := range stageSet.Stage.Runes {
			runeOrdered = append(runeOrdered, rune)
		}
		sort.Slice(runeOrdered, func(i, j int) bool {
			return stageSet.Stage.Rune_stagedOrder[runeOrdered[i]] < stageSet.Stage.Rune_stagedOrder[runeOrdered[j]]
		})
		for _, rune := range runeOrdered {
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
		runepropertiesOrdered := []*RuneProperties{}
		for runeproperties := range stageSet.Stage.RunePropertiess {
			runepropertiesOrdered = append(runepropertiesOrdered, runeproperties)
		}
		sort.Slice(runepropertiesOrdered, func(i, j int) bool {
			return stageSet.Stage.RuneProperties_stagedOrder[runepropertiesOrdered[i]] < stageSet.Stage.RuneProperties_stagedOrder[runepropertiesOrdered[j]]
		})
		for _, runeproperties := range runepropertiesOrdered {
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
		tableOrdered := []*Table{}
		for table := range stageSet.Stage.Tables {
			tableOrdered = append(tableOrdered, table)
		}
		sort.Slice(tableOrdered, func(i, j int) bool {
			return stageSet.Stage.Table_stagedOrder[tableOrdered[i]] < stageSet.Stage.Table_stagedOrder[tableOrdered[j]]
		})
		for _, table := range tableOrdered {
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
		tablecolumnOrdered := []*TableColumn{}
		for tablecolumn := range stageSet.Stage.TableColumns {
			tablecolumnOrdered = append(tablecolumnOrdered, tablecolumn)
		}
		sort.Slice(tablecolumnOrdered, func(i, j int) bool {
			return stageSet.Stage.TableColumn_stagedOrder[tablecolumnOrdered[i]] < stageSet.Stage.TableColumn_stagedOrder[tablecolumnOrdered[j]]
		})
		for _, tablecolumn := range tablecolumnOrdered {
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
		tablepropertiesOrdered := []*TableProperties{}
		for tableproperties := range stageSet.Stage.TablePropertiess {
			tablepropertiesOrdered = append(tablepropertiesOrdered, tableproperties)
		}
		sort.Slice(tablepropertiesOrdered, func(i, j int) bool {
			return stageSet.Stage.TableProperties_stagedOrder[tablepropertiesOrdered[i]] < stageSet.Stage.TableProperties_stagedOrder[tablepropertiesOrdered[j]]
		})
		for _, tableproperties := range tablepropertiesOrdered {
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
		tablerowOrdered := []*TableRow{}
		for tablerow := range stageSet.Stage.TableRows {
			tablerowOrdered = append(tablerowOrdered, tablerow)
		}
		sort.Slice(tablerowOrdered, func(i, j int) bool {
			return stageSet.Stage.TableRow_stagedOrder[tablerowOrdered[i]] < stageSet.Stage.TableRow_stagedOrder[tablerowOrdered[j]]
		})
		for _, tablerow := range tablerowOrdered {
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
		tablestyleOrdered := []*TableStyle{}
		for tablestyle := range stageSet.Stage.TableStyles {
			tablestyleOrdered = append(tablestyleOrdered, tablestyle)
		}
		sort.Slice(tablestyleOrdered, func(i, j int) bool {
			return stageSet.Stage.TableStyle_stagedOrder[tablestyleOrdered[i]] < stageSet.Stage.TableStyle_stagedOrder[tablestyleOrdered[j]]
		})
		for _, tablestyle := range tablestyleOrdered {
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
		textOrdered := []*Text{}
		for text := range stageSet.Stage.Texts {
			textOrdered = append(textOrdered, text)
		}
		sort.Slice(textOrdered, func(i, j int) bool {
			return stageSet.Stage.Text_stagedOrder[textOrdered[i]] < stageSet.Stage.Text_stagedOrder[textOrdered[j]]
		})
		for _, text := range textOrdered {
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
