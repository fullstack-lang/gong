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
		attributeshapeOrdered := []*AttributeShape{}
		for attributeshape := range stageSet.Stage.AttributeShapes {
			attributeshapeOrdered = append(attributeshapeOrdered, attributeshape)
		}
		sort.Slice(attributeshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.AttributeShape_stagedOrder[attributeshapeOrdered[i]] < stageSet.Stage.AttributeShape_stagedOrder[attributeshapeOrdered[j]]
		})
		for _, attributeshape := range attributeshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attributeshapeIdent := "__models" + attributeshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AttributeShape{Name: %s}).Stage(stageSet.Stage)", attributeshapeIdent, __gong__toRawStringLiteral(attributeshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attributeshapeIdent, __gong__toRawStringLiteral(attributeshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.FieldTypeAsString = %s", attributeshapeIdent, __gong__toRawStringLiteral(attributeshape.FieldTypeAsString)))
			values.WriteString(fmt.Sprintf("\n\t%s.Structname = %s", attributeshapeIdent, __gong__toRawStringLiteral(attributeshape.Structname)))
			values.WriteString(fmt.Sprintf("\n\t%s.Fieldtypename = %s", attributeshapeIdent, __gong__toRawStringLiteral(attributeshape.Fieldtypename)))
		}
	}
	if stageSet.Stage != nil {
		classdiagramOrdered := []*Classdiagram{}
		for classdiagram := range stageSet.Stage.Classdiagrams {
			classdiagramOrdered = append(classdiagramOrdered, classdiagram)
		}
		sort.Slice(classdiagramOrdered, func(i, j int) bool {
			return stageSet.Stage.Classdiagram_stagedOrder[classdiagramOrdered[i]] < stageSet.Stage.Classdiagram_stagedOrder[classdiagramOrdered[j]]
		})
		for _, classdiagram := range classdiagramOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			classdiagramIdent := "__models" + classdiagram.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Classdiagram{Name: %s}).Stage(stageSet.Stage)", classdiagramIdent, __gong__toRawStringLiteral(classdiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", classdiagramIdent, __gong__toRawStringLiteral(classdiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", classdiagramIdent, __gong__toRawStringLiteral(classdiagram.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsIncludedInStaticWebSite = %t", classdiagramIdent, classdiagram.IsIncludedInStaticWebSite))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowNbInstances = %t", classdiagramIdent, classdiagram.ShowNbInstances))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowMultiplicity = %t", classdiagramIdent, classdiagram.ShowMultiplicity))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowLinkNames = %t", classdiagramIdent, classdiagram.ShowLinkNames))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInRenameMode = %t", classdiagramIdent, classdiagram.IsInRenameMode))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", classdiagramIdent, classdiagram.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NodeGongStructsIsExpanded = %t", classdiagramIdent, classdiagram.NodeGongStructsIsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NodeGongStructNodeExpansion = %s", classdiagramIdent, __gong__toRawStringLiteral(classdiagram.NodeGongStructNodeExpansion)))
			values.WriteString(fmt.Sprintf("\n\t%s.NodeGongEnumsIsExpanded = %t", classdiagramIdent, classdiagram.NodeGongEnumsIsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NodeGongEnumNodeExpansion = %s", classdiagramIdent, __gong__toRawStringLiteral(classdiagram.NodeGongEnumNodeExpansion)))
			values.WriteString(fmt.Sprintf("\n\t%s.NodeGongNotesIsExpanded = %t", classdiagramIdent, classdiagram.NodeGongNotesIsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NodeGongNoteNodeExpansion = %s", classdiagramIdent, __gong__toRawStringLiteral(classdiagram.NodeGongNoteNodeExpansion)))
			for _, elem := range classdiagram.GongStructShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongStructShapes = append(%s.GongStructShapes, %s)", classdiagramIdent, classdiagramIdent, targetIdent))
			}
			for _, elem := range classdiagram.GongEnumShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongEnumShapes = append(%s.GongEnumShapes, %s)", classdiagramIdent, classdiagramIdent, targetIdent))
			}
			for _, elem := range classdiagram.GongNoteShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongNoteShapes = append(%s.GongNoteShapes, %s)", classdiagramIdent, classdiagramIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		diagrampackageOrdered := []*DiagramPackage{}
		for diagrampackage := range stageSet.Stage.DiagramPackages {
			diagrampackageOrdered = append(diagrampackageOrdered, diagrampackage)
		}
		sort.Slice(diagrampackageOrdered, func(i, j int) bool {
			return stageSet.Stage.DiagramPackage_stagedOrder[diagrampackageOrdered[i]] < stageSet.Stage.DiagramPackage_stagedOrder[diagrampackageOrdered[j]]
		})
		for _, diagrampackage := range diagrampackageOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagrampackageIdent := "__models" + diagrampackage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DiagramPackage{Name: %s}).Stage(stageSet.Stage)", diagrampackageIdent, __gong__toRawStringLiteral(diagrampackage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagrampackageIdent, __gong__toRawStringLiteral(diagrampackage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Path = %s", diagrampackageIdent, __gong__toRawStringLiteral(diagrampackage.Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.GongModelPath = %s", diagrampackageIdent, __gong__toRawStringLiteral(diagrampackage.GongModelPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.AbsolutePathToDiagramPackage = %s", diagrampackageIdent, __gong__toRawStringLiteral(diagrampackage.AbsolutePathToDiagramPackage)))
			for _, elem := range diagrampackage.Classdiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Classdiagrams = append(%s.Classdiagrams, %s)", diagrampackageIdent, diagrampackageIdent, targetIdent))
			}
			if diagrampackage.SelectedClassdiagram != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + diagrampackage.SelectedClassdiagram.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SelectedClassdiagram = %s", diagrampackageIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		gongenumshapeOrdered := []*GongEnumShape{}
		for gongenumshape := range stageSet.Stage.GongEnumShapes {
			gongenumshapeOrdered = append(gongenumshapeOrdered, gongenumshape)
		}
		sort.Slice(gongenumshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.GongEnumShape_stagedOrder[gongenumshapeOrdered[i]] < stageSet.Stage.GongEnumShape_stagedOrder[gongenumshapeOrdered[j]]
		})
		for _, gongenumshape := range gongenumshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongenumshapeIdent := "__models" + gongenumshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongEnumShape{Name: %s}).Stage(stageSet.Stage)", gongenumshapeIdent, __gong__toRawStringLiteral(gongenumshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongenumshapeIdent, __gong__toRawStringLiteral(gongenumshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", gongenumshapeIdent, gongenumshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", gongenumshapeIdent, gongenumshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", gongenumshapeIdent, gongenumshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", gongenumshapeIdent, gongenumshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", gongenumshapeIdent, gongenumshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", gongenumshapeIdent, gongenumshape.IsExpanded))
			for _, elem := range gongenumshape.GongEnumValueShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongEnumValueShapes = append(%s.GongEnumValueShapes, %s)", gongenumshapeIdent, gongenumshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		gongenumvalueshapeOrdered := []*GongEnumValueShape{}
		for gongenumvalueshape := range stageSet.Stage.GongEnumValueShapes {
			gongenumvalueshapeOrdered = append(gongenumvalueshapeOrdered, gongenumvalueshape)
		}
		sort.Slice(gongenumvalueshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.GongEnumValueShape_stagedOrder[gongenumvalueshapeOrdered[i]] < stageSet.Stage.GongEnumValueShape_stagedOrder[gongenumvalueshapeOrdered[j]]
		})
		for _, gongenumvalueshape := range gongenumvalueshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongenumvalueshapeIdent := "__models" + gongenumvalueshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongEnumValueShape{Name: %s}).Stage(stageSet.Stage)", gongenumvalueshapeIdent, __gong__toRawStringLiteral(gongenumvalueshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongenumvalueshapeIdent, __gong__toRawStringLiteral(gongenumvalueshape.Name)))
		}
	}
	if stageSet.Stage != nil {
		gongnotelinkshapeOrdered := []*GongNoteLinkShape{}
		for gongnotelinkshape := range stageSet.Stage.GongNoteLinkShapes {
			gongnotelinkshapeOrdered = append(gongnotelinkshapeOrdered, gongnotelinkshape)
		}
		sort.Slice(gongnotelinkshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.GongNoteLinkShape_stagedOrder[gongnotelinkshapeOrdered[i]] < stageSet.Stage.GongNoteLinkShape_stagedOrder[gongnotelinkshapeOrdered[j]]
		})
		for _, gongnotelinkshape := range gongnotelinkshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongnotelinkshapeIdent := "__models" + gongnotelinkshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongNoteLinkShape{Name: %s}).Stage(stageSet.Stage)", gongnotelinkshapeIdent, __gong__toRawStringLiteral(gongnotelinkshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongnotelinkshapeIdent, __gong__toRawStringLiteral(gongnotelinkshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", gongnotelinkshapeIdent, __gong__toRawStringLiteral(gongnotelinkshape.Identifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", gongnotelinkshapeIdent, __gong__toRawStringLiteral(string(gongnotelinkshape.Type))))
		}
	}
	if stageSet.Stage != nil {
		gongnoteshapeOrdered := []*GongNoteShape{}
		for gongnoteshape := range stageSet.Stage.GongNoteShapes {
			gongnoteshapeOrdered = append(gongnoteshapeOrdered, gongnoteshape)
		}
		sort.Slice(gongnoteshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.GongNoteShape_stagedOrder[gongnoteshapeOrdered[i]] < stageSet.Stage.GongNoteShape_stagedOrder[gongnoteshapeOrdered[j]]
		})
		for _, gongnoteshape := range gongnoteshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongnoteshapeIdent := "__models" + gongnoteshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongNoteShape{Name: %s}).Stage(stageSet.Stage)", gongnoteshapeIdent, __gong__toRawStringLiteral(gongnoteshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongnoteshapeIdent, __gong__toRawStringLiteral(gongnoteshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", gongnoteshapeIdent, __gong__toRawStringLiteral(gongnoteshape.Identifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.Body = %s", gongnoteshapeIdent, __gong__toRawStringLiteral(gongnoteshape.Body)))
			values.WriteString(fmt.Sprintf("\n\t%s.BodyHTML = %s", gongnoteshapeIdent, __gong__toRawStringLiteral(gongnoteshape.BodyHTML)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", gongnoteshapeIdent, gongnoteshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", gongnoteshapeIdent, gongnoteshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", gongnoteshapeIdent, gongnoteshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", gongnoteshapeIdent, gongnoteshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", gongnoteshapeIdent, gongnoteshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.Matched = %t", gongnoteshapeIdent, gongnoteshape.Matched))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", gongnoteshapeIdent, gongnoteshape.IsExpanded))
			for _, elem := range gongnoteshape.GongNoteLinkShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongNoteLinkShapes = append(%s.GongNoteLinkShapes, %s)", gongnoteshapeIdent, gongnoteshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		gongstructshapeOrdered := []*GongStructShape{}
		for gongstructshape := range stageSet.Stage.GongStructShapes {
			gongstructshapeOrdered = append(gongstructshapeOrdered, gongstructshape)
		}
		sort.Slice(gongstructshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.GongStructShape_stagedOrder[gongstructshapeOrdered[i]] < stageSet.Stage.GongStructShape_stagedOrder[gongstructshapeOrdered[j]]
		})
		for _, gongstructshape := range gongstructshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongstructshapeIdent := "__models" + gongstructshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongStructShape{Name: %s}).Stage(stageSet.Stage)", gongstructshapeIdent, __gong__toRawStringLiteral(gongstructshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongstructshapeIdent, __gong__toRawStringLiteral(gongstructshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", gongstructshapeIdent, gongstructshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", gongstructshapeIdent, gongstructshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", gongstructshapeIdent, gongstructshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", gongstructshapeIdent, gongstructshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", gongstructshapeIdent, gongstructshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelected = %t", gongstructshapeIdent, gongstructshape.IsSelected))
			for _, elem := range gongstructshape.AttributeShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeShapes = append(%s.AttributeShapes, %s)", gongstructshapeIdent, gongstructshapeIdent, targetIdent))
			}
			for _, elem := range gongstructshape.LinkShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LinkShapes = append(%s.LinkShapes, %s)", gongstructshapeIdent, gongstructshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		linkshapeOrdered := []*LinkShape{}
		for linkshape := range stageSet.Stage.LinkShapes {
			linkshapeOrdered = append(linkshapeOrdered, linkshape)
		}
		sort.Slice(linkshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.LinkShape_stagedOrder[linkshapeOrdered[i]] < stageSet.Stage.LinkShape_stagedOrder[linkshapeOrdered[j]]
		})
		for _, linkshape := range linkshapeOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linkshapeIdent := "__models" + linkshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LinkShape{Name: %s}).Stage(stageSet.Stage)", linkshapeIdent, __gong__toRawStringLiteral(linkshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linkshapeIdent, __gong__toRawStringLiteral(linkshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.FieldOffsetX = %f", linkshapeIdent, linkshape.FieldOffsetX))
			values.WriteString(fmt.Sprintf("\n\t%s.FieldOffsetY = %f", linkshapeIdent, linkshape.FieldOffsetY))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetMultiplicity = %s", linkshapeIdent, __gong__toRawStringLiteral(string(linkshape.TargetMultiplicity))))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetMultiplicityOffsetX = %f", linkshapeIdent, linkshape.TargetMultiplicityOffsetX))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetMultiplicityOffsetY = %f", linkshapeIdent, linkshape.TargetMultiplicityOffsetY))
			values.WriteString(fmt.Sprintf("\n\t%s.SourceMultiplicity = %s", linkshapeIdent, __gong__toRawStringLiteral(string(linkshape.SourceMultiplicity))))
			values.WriteString(fmt.Sprintf("\n\t%s.SourceMultiplicityOffsetX = %f", linkshapeIdent, linkshape.SourceMultiplicityOffsetX))
			values.WriteString(fmt.Sprintf("\n\t%s.SourceMultiplicityOffsetY = %f", linkshapeIdent, linkshape.SourceMultiplicityOffsetY))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", linkshapeIdent, linkshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", linkshapeIdent, linkshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", linkshapeIdent, __gong__toRawStringLiteral(string(linkshape.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", linkshapeIdent, linkshape.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", linkshapeIdent, __gong__toRawStringLiteral(string(linkshape.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", linkshapeIdent, linkshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", linkshapeIdent, linkshape.CornerOffsetRatio))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
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
		case "github.com/fullstack-lang/gong/lib/doc/go/models":
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
				case "AttributeShape":
					if !preserveOrder {
						inst := (&AttributeShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(AttributeShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Classdiagram":
					if !preserveOrder {
						inst := (&Classdiagram{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Classdiagram)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "DiagramPackage":
					if !preserveOrder {
						inst := (&DiagramPackage{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(DiagramPackage)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GongEnumShape":
					if !preserveOrder {
						inst := (&GongEnumShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GongEnumShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GongEnumValueShape":
					if !preserveOrder {
						inst := (&GongEnumValueShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GongEnumValueShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GongNoteLinkShape":
					if !preserveOrder {
						inst := (&GongNoteLinkShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GongNoteLinkShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GongNoteShape":
					if !preserveOrder {
						inst := (&GongNoteShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GongNoteShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "GongStructShape":
					if !preserveOrder {
						inst := (&GongStructShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(GongStructShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "LinkShape":
					if !preserveOrder {
						inst := (&LinkShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(LinkShape)
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
				case *AttributeShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "FieldTypeAsString":
						inst.FieldTypeAsString = GongExtractString(rhs)
					case "Structname":
						inst.Structname = GongExtractString(rhs)
					case "Fieldtypename":
						inst.Fieldtypename = GongExtractString(rhs)
					}
				case *Classdiagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "IsIncludedInStaticWebSite":
						inst.IsIncludedInStaticWebSite = GongExtractBool(rhs)
					case "GongStructShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GongStructShape); ok {
										inst.GongStructShapes = append(inst.GongStructShapes, typedTarget)
									}
								}
							}
						}
					case "GongEnumShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GongEnumShape); ok {
										inst.GongEnumShapes = append(inst.GongEnumShapes, typedTarget)
									}
								}
							}
						}
					case "GongNoteShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GongNoteShape); ok {
										inst.GongNoteShapes = append(inst.GongNoteShapes, typedTarget)
									}
								}
							}
						}
					case "ShowNbInstances":
						inst.ShowNbInstances = GongExtractBool(rhs)
					case "ShowMultiplicity":
						inst.ShowMultiplicity = GongExtractBool(rhs)
					case "ShowLinkNames":
						inst.ShowLinkNames = GongExtractBool(rhs)
					case "IsInRenameMode":
						inst.IsInRenameMode = GongExtractBool(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "NodeGongStructsIsExpanded":
						inst.NodeGongStructsIsExpanded = GongExtractBool(rhs)
					case "NodeGongStructNodeExpansion":
						inst.NodeGongStructNodeExpansion = GongExtractString(rhs)
					case "NodeGongEnumsIsExpanded":
						inst.NodeGongEnumsIsExpanded = GongExtractBool(rhs)
					case "NodeGongEnumNodeExpansion":
						inst.NodeGongEnumNodeExpansion = GongExtractString(rhs)
					case "NodeGongNotesIsExpanded":
						inst.NodeGongNotesIsExpanded = GongExtractBool(rhs)
					case "NodeGongNoteNodeExpansion":
						inst.NodeGongNoteNodeExpansion = GongExtractString(rhs)
					}
				case *DiagramPackage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Path":
						inst.Path = GongExtractString(rhs)
					case "GongModelPath":
						inst.GongModelPath = GongExtractString(rhs)
					case "Classdiagrams":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Classdiagram); ok {
										inst.Classdiagrams = append(inst.Classdiagrams, typedTarget)
									}
								}
							}
						}
					case "SelectedClassdiagram":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Classdiagram); ok {
									inst.SelectedClassdiagram = typedTarget
								}
							}
						}
					case "AbsolutePathToDiagramPackage":
						inst.AbsolutePathToDiagramPackage = GongExtractString(rhs)
					}
				case *GongEnumShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					case "GongEnumValueShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GongEnumValueShape); ok {
										inst.GongEnumValueShapes = append(inst.GongEnumValueShapes, typedTarget)
									}
								}
							}
						}
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *GongEnumValueShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *GongNoteLinkShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Identifier":
						inst.Identifier = GongExtractString(rhs)
					case "Type":
						inst.Type = NoteShapeLinkType(GongExtractString(rhs))
					}
				case *GongNoteShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Identifier":
						inst.Identifier = GongExtractString(rhs)
					case "Body":
						inst.Body = GongExtractString(rhs)
					case "BodyHTML":
						inst.BodyHTML = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					case "Matched":
						inst.Matched = GongExtractBool(rhs)
					case "GongNoteLinkShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*GongNoteLinkShape); ok {
										inst.GongNoteLinkShapes = append(inst.GongNoteLinkShapes, typedTarget)
									}
								}
							}
						}
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *GongStructShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					case "AttributeShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*AttributeShape); ok {
										inst.AttributeShapes = append(inst.AttributeShapes, typedTarget)
									}
								}
							}
						}
					case "LinkShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*LinkShape); ok {
										inst.LinkShapes = append(inst.LinkShapes, typedTarget)
									}
								}
							}
						}
					case "IsSelected":
						inst.IsSelected = GongExtractBool(rhs)
					}
				case *LinkShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "FieldOffsetX":
						inst.FieldOffsetX = GongExtractFloat(rhs)
					case "FieldOffsetY":
						inst.FieldOffsetY = GongExtractFloat(rhs)
					case "TargetMultiplicity":
						inst.TargetMultiplicity = MultiplicityType(GongExtractString(rhs))
					case "TargetMultiplicityOffsetX":
						inst.TargetMultiplicityOffsetX = GongExtractFloat(rhs)
					case "TargetMultiplicityOffsetY":
						inst.TargetMultiplicityOffsetY = GongExtractFloat(rhs)
					case "SourceMultiplicity":
						inst.SourceMultiplicity = MultiplicityType(GongExtractString(rhs))
					case "SourceMultiplicityOffsetX":
						inst.SourceMultiplicityOffsetX = GongExtractFloat(rhs)
					case "SourceMultiplicityOffsetY":
						inst.SourceMultiplicityOffsetY = GongExtractFloat(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
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
