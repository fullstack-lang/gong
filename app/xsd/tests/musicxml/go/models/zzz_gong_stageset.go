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
		for _, a_directive := range __gong__sortStageSetInstances(stageSet.Stage.A_directives, stageSet.Stage.A_directive_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_directiveIdent := "__models" + a_directive.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_directive{Name: %s}).Stage(stageSet.Stage)", a_directiveIdent, __gong__toRawStringLiteral(a_directive.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_directiveIdent, __gong__toRawStringLiteral(a_directive.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", a_directiveIdent, __gong__toRawStringLiteral(a_directive.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", a_directiveIdent, __gong__toRawStringLiteral(a_directive.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", a_directiveIdent, __gong__toRawStringLiteral(a_directive.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", a_directiveIdent, __gong__toRawStringLiteral(a_directive.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", a_directiveIdent, __gong__toRawStringLiteral(a_directive.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", a_directiveIdent, __gong__toRawStringLiteral(a_directive.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", a_directiveIdent, __gong__toRawStringLiteral(a_directive.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", a_directiveIdent, __gong__toRawStringLiteral(a_directive.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", a_directiveIdent, __gong__toRawStringLiteral(a_directive.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", a_directiveIdent, __gong__toRawStringLiteral(a_directive.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", a_directiveIdent, __gong__toRawStringLiteral(a_directive.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, a_measure := range __gong__sortStageSetInstances(stageSet.Stage.A_measures, stageSet.Stage.A_measure_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_measureIdent := "__models" + a_measure.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_measure{Name: %s}).Stage(stageSet.Stage)", a_measureIdent, __gong__toRawStringLiteral(a_measure.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_measureIdent, __gong__toRawStringLiteral(a_measure.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", a_measureIdent, __gong__toRawStringLiteral(a_measure.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", a_measureIdent, __gong__toRawStringLiteral(a_measure.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Implicit = %s", a_measureIdent, __gong__toRawStringLiteral(string(a_measure.Implicit))))
			values.WriteString(fmt.Sprintf("\n\t%s.Non_controlling = %s", a_measureIdent, __gong__toRawStringLiteral(string(a_measure.Non_controlling))))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", a_measureIdent, __gong__toRawStringLiteral(a_measure.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", a_measureIdent, __gong__toRawStringLiteral(a_measure.Id)))
			for _, elem := range a_measure.Note {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = append(%s.Note, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Backup {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Backup = append(%s.Backup, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Forward {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Forward = append(%s.Forward, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Direction {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Direction = append(%s.Direction, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Attributes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Harmony {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmony = append(%s.Harmony, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Figured_bass {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Figured_bass = append(%s.Figured_bass, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Print {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Print = append(%s.Print, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Sound {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sound = append(%s.Sound, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Listening {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listening = append(%s.Listening, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Barline {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Barline = append(%s.Barline, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Grouping {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Grouping = append(%s.Grouping, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Link {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Link = append(%s.Link, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Bookmark {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bookmark = append(%s.Bookmark, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_measure_1 := range __gong__sortStageSetInstances(stageSet.Stage.A_measure_1s, stageSet.Stage.A_measure_1_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_measure_1Ident := "__models" + a_measure_1.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_measure_1{Name: %s}).Stage(stageSet.Stage)", a_measure_1Ident, __gong__toRawStringLiteral(a_measure_1.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_measure_1Ident, __gong__toRawStringLiteral(a_measure_1.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", a_measure_1Ident, __gong__toRawStringLiteral(a_measure_1.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", a_measure_1Ident, __gong__toRawStringLiteral(a_measure_1.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Implicit = %s", a_measure_1Ident, __gong__toRawStringLiteral(string(a_measure_1.Implicit))))
			values.WriteString(fmt.Sprintf("\n\t%s.Non_controlling = %s", a_measure_1Ident, __gong__toRawStringLiteral(string(a_measure_1.Non_controlling))))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", a_measure_1Ident, __gong__toRawStringLiteral(a_measure_1.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", a_measure_1Ident, __gong__toRawStringLiteral(a_measure_1.Id)))
			for _, elem := range a_measure_1.Part {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = append(%s.Part, %s)", a_measure_1Ident, a_measure_1Ident, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_part := range __gong__sortStageSetInstances(stageSet.Stage.A_parts, stageSet.Stage.A_part_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_partIdent := "__models" + a_part.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_part{Name: %s}).Stage(stageSet.Stage)", a_partIdent, __gong__toRawStringLiteral(a_part.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_partIdent, __gong__toRawStringLiteral(a_part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", a_partIdent, __gong__toRawStringLiteral(a_part.Id)))
			for _, elem := range a_part.Measure {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure = append(%s.Measure, %s)", a_partIdent, a_partIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, a_part_1 := range __gong__sortStageSetInstances(stageSet.Stage.A_part_1s, stageSet.Stage.A_part_1_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_part_1Ident := "__models" + a_part_1.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_part_1{Name: %s}).Stage(stageSet.Stage)", a_part_1Ident, __gong__toRawStringLiteral(a_part_1.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_part_1Ident, __gong__toRawStringLiteral(a_part_1.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", a_part_1Ident, __gong__toRawStringLiteral(a_part_1.Id)))
			for _, elem := range a_part_1.Note {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = append(%s.Note, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Backup {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Backup = append(%s.Backup, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Forward {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Forward = append(%s.Forward, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Direction {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Direction = append(%s.Direction, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Attributes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Harmony {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmony = append(%s.Harmony, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Figured_bass {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Figured_bass = append(%s.Figured_bass, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Print {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Print = append(%s.Print, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Sound {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sound = append(%s.Sound, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Listening {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listening = append(%s.Listening, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Barline {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Barline = append(%s.Barline, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Grouping {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Grouping = append(%s.Grouping, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Link {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Link = append(%s.Link, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Bookmark {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bookmark = append(%s.Bookmark, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, accidental := range __gong__sortStageSetInstances(stageSet.Stage.Accidentals, stageSet.Stage.Accidental_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			accidentalIdent := "__models" + accidental.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Accidental{Name: %s}).Stage(stageSet.Stage)", accidentalIdent, __gong__toRawStringLiteral(accidental.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Cautionary = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Cautionary)))
			values.WriteString(fmt.Sprintf("\n\t%s.Editorial = %s", accidentalIdent, __gong__toRawStringLiteral(string(accidental.Editorial))))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", accidentalIdent, __gong__toRawStringLiteral(string(accidental.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.Bracket = %s", accidentalIdent, __gong__toRawStringLiteral(string(accidental.Bracket))))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %s", accidentalIdent, __gong__toRawStringLiteral(string(accidental.Size))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, accidental_mark := range __gong__sortStageSetInstances(stageSet.Stage.Accidental_marks, stageSet.Stage.Accidental_mark_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			accidental_markIdent := "__models" + accidental_mark.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Accidental_mark{Name: %s}).Stage(stageSet.Stage)", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", accidental_markIdent, __gong__toRawStringLiteral(string(accidental_mark.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.Bracket = %s", accidental_markIdent, __gong__toRawStringLiteral(string(accidental_mark.Bracket))))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %s", accidental_markIdent, __gong__toRawStringLiteral(string(accidental_mark.Size))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", accidental_markIdent, __gong__toRawStringLiteral(string(accidental_mark.EnclosedText))))
		}
	}
	if stageSet.Stage != nil {
		for _, accidental_text := range __gong__sortStageSetInstances(stageSet.Stage.Accidental_texts, stageSet.Stage.Accidental_text_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			accidental_textIdent := "__models" + accidental_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Accidental_text{Name: %s}).Stage(stageSet.Stage)", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Space)))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", accidental_textIdent, __gong__toRawStringLiteral(string(accidental_text.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", accidental_textIdent, accidental_text.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", accidental_textIdent, accidental_text.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_through = %d", accidental_textIdent, accidental_text.Line_through))
			values.WriteString(fmt.Sprintf("\n\t%s.Rotation = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Rotation)))
			values.WriteString(fmt.Sprintf("\n\t%s.Letter_spacing = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Letter_spacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_height = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Line_height)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dir = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Dir)))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Enclosure)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", accidental_textIdent, __gong__toRawStringLiteral(string(accidental_text.EnclosedText))))
		}
	}
	if stageSet.Stage != nil {
		for _, accord := range __gong__sortStageSetInstances(stageSet.Stage.Accords, stageSet.Stage.Accord_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			accordIdent := "__models" + accord.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Accord{Name: %s}).Stage(stageSet.Stage)", accordIdent, __gong__toRawStringLiteral(accord.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", accordIdent, __gong__toRawStringLiteral(accord.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.String = %d", accordIdent, accord.String))
			values.WriteString(fmt.Sprintf("\n\t%s.Tuning_step = %s", accordIdent, __gong__toRawStringLiteral(string(accord.Tuning_step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Tuning_alter = %s", accordIdent, __gong__toRawStringLiteral(accord.Tuning_alter)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tuning_octave = %d", accordIdent, accord.Tuning_octave))
		}
	}
	if stageSet.Stage != nil {
		for _, accordion_registration := range __gong__sortStageSetInstances(stageSet.Stage.Accordion_registrations, stageSet.Stage.Accordion_registration_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			accordion_registrationIdent := "__models" + accordion_registration.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Accordion_registration{Name: %s}).Stage(stageSet.Stage)", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Accordion_high = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Accordion_high)))
			values.WriteString(fmt.Sprintf("\n\t%s.Accordion_middle = %d", accordion_registrationIdent, accordion_registration.Accordion_middle))
			values.WriteString(fmt.Sprintf("\n\t%s.Accordion_low = %s", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Accordion_low)))
		}
	}
	if stageSet.Stage != nil {
		for _, appearance := range __gong__sortStageSetInstances(stageSet.Stage.Appearances, stageSet.Stage.Appearance_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			appearanceIdent := "__models" + appearance.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Appearance{Name: %s}).Stage(stageSet.Stage)", appearanceIdent, __gong__toRawStringLiteral(appearance.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", appearanceIdent, __gong__toRawStringLiteral(appearance.Name)))
			for _, elem := range appearance.Line_width {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Line_width = append(%s.Line_width, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.Note_size {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_size = append(%s.Note_size, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.Distance {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = append(%s.Distance, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.Glyph {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Glyph = append(%s.Glyph, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.Other_appearance {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_appearance = append(%s.Other_appearance, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, arpeggiate := range __gong__sortStageSetInstances(stageSet.Stage.Arpeggiates, stageSet.Stage.Arpeggiate_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arpeggiateIdent := "__models" + arpeggiate.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Arpeggiate{Name: %s}).Stage(stageSet.Stage)", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", arpeggiateIdent, arpeggiate.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Direction)))
			values.WriteString(fmt.Sprintf("\n\t%s.Unbroken = %s", arpeggiateIdent, __gong__toRawStringLiteral(string(arpeggiate.Unbroken))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, arrow := range __gong__sortStageSetInstances(stageSet.Stage.Arrows, stageSet.Stage.Arrow_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			arrowIdent := "__models" + arrow.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Arrow{Name: %s}).Stage(stageSet.Stage)", arrowIdent, __gong__toRawStringLiteral(arrow.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arrow_direction = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Arrow_direction)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arrow_style = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Arrow_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arrowhead = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Arrowhead)))
			values.WriteString(fmt.Sprintf("\n\t%s.Circular_arrow = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Circular_arrow)))
		}
	}
	if stageSet.Stage != nil {
		for _, articulations := range __gong__sortStageSetInstances(stageSet.Stage.Articulationss, stageSet.Stage.Articulations_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			articulationsIdent := "__models" + articulations.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Articulations{Name: %s}).Stage(stageSet.Stage)", articulationsIdent, __gong__toRawStringLiteral(articulations.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", articulationsIdent, __gong__toRawStringLiteral(articulations.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", articulationsIdent, __gong__toRawStringLiteral(articulations.Id)))
			for _, elem := range articulations.Accent {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accent = append(%s.Accent, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Strong_accent {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Strong_accent = append(%s.Strong_accent, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Staccato {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staccato = append(%s.Staccato, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Tenuto {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tenuto = append(%s.Tenuto, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Detached_legato {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Detached_legato = append(%s.Detached_legato, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Staccatissimo {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staccatissimo = append(%s.Staccatissimo, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Spiccato {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Spiccato = append(%s.Spiccato, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Scoop {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Scoop = append(%s.Scoop, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Plop {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Plop = append(%s.Plop, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Doit {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Doit = append(%s.Doit, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Falloff {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Falloff = append(%s.Falloff, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Breath_mark {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Breath_mark = append(%s.Breath_mark, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Caesura {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Caesura = append(%s.Caesura, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Stress {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stress = append(%s.Stress, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Unstress {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Unstress = append(%s.Unstress, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Soft_accent {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Soft_accent = append(%s.Soft_accent, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Other_articulation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_articulation = append(%s.Other_articulation, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, assess := range __gong__sortStageSetInstances(stageSet.Stage.Assesss, stageSet.Stage.Assess_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			assessIdent := "__models" + assess.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Assess{Name: %s}).Stage(stageSet.Stage)", assessIdent, __gong__toRawStringLiteral(assess.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", assessIdent, __gong__toRawStringLiteral(assess.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", assessIdent, __gong__toRawStringLiteral(string(assess.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", assessIdent, __gong__toRawStringLiteral(assess.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_only = %s", assessIdent, __gong__toRawStringLiteral(assess.Time_only)))
		}
	}
	if stageSet.Stage != nil {
		for _, attributes := range __gong__sortStageSetInstances(stageSet.Stage.Attributess, stageSet.Stage.Attributes_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attributesIdent := "__models" + attributes.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Attributes{Name: %s}).Stage(stageSet.Stage)", attributesIdent, __gong__toRawStringLiteral(attributes.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attributesIdent, __gong__toRawStringLiteral(attributes.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Divisions = %s", attributesIdent, __gong__toRawStringLiteral(attributes.Divisions)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staves = %d", attributesIdent, attributes.Staves))
			values.WriteString(fmt.Sprintf("\n\t%s.Instruments = %d", attributesIdent, attributes.Instruments))
			if attributes.Footnote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attributes.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", attributesIdent, targetIdent))
			}
			if attributes.Level != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attributes.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Key {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Key = append(%s.Key, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Time {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Time = append(%s.Time, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			if attributes.Part_symbol != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + attributes.Part_symbol.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_symbol = %s", attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Clef {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Clef = append(%s.Clef, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Staff_details {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staff_details = append(%s.Staff_details, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Transpose {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Transpose = append(%s.Transpose, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.For_part {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.For_part = append(%s.For_part, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Directive {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Directive = append(%s.Directive, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Measure_style {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure_style = append(%s.Measure_style, %s)", attributesIdent, attributesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, backup := range __gong__sortStageSetInstances(stageSet.Stage.Backups, stageSet.Stage.Backup_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			backupIdent := "__models" + backup.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Backup{Name: %s}).Stage(stageSet.Stage)", backupIdent, __gong__toRawStringLiteral(backup.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", backupIdent, __gong__toRawStringLiteral(backup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = %s", backupIdent, __gong__toRawStringLiteral(backup.Duration)))
			if backup.Footnote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + backup.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", backupIdent, targetIdent))
			}
			if backup.Level != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + backup.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", backupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, bar_style_color := range __gong__sortStageSetInstances(stageSet.Stage.Bar_style_colors, stageSet.Stage.Bar_style_color_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bar_style_colorIdent := "__models" + bar_style_color.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Bar_style_color{Name: %s}).Stage(stageSet.Stage)", bar_style_colorIdent, __gong__toRawStringLiteral(bar_style_color.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bar_style_colorIdent, __gong__toRawStringLiteral(bar_style_color.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", bar_style_colorIdent, __gong__toRawStringLiteral(bar_style_color.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", bar_style_colorIdent, __gong__toRawStringLiteral(bar_style_color.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, barline := range __gong__sortStageSetInstances(stageSet.Stage.Barlines, stageSet.Stage.Barline_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			barlineIdent := "__models" + barline.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Barline{Name: %s}).Stage(stageSet.Stage)", barlineIdent, __gong__toRawStringLiteral(barline.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", barlineIdent, __gong__toRawStringLiteral(barline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", barlineIdent, __gong__toRawStringLiteral(barline.Location)))
			values.WriteString(fmt.Sprintf("\n\t%s.Segno = %s", barlineIdent, __gong__toRawStringLiteral(barline.Segno)))
			values.WriteString(fmt.Sprintf("\n\t%s.Coda = %s", barlineIdent, __gong__toRawStringLiteral(barline.Coda)))
			values.WriteString(fmt.Sprintf("\n\t%s.Divisions = %s", barlineIdent, __gong__toRawStringLiteral(barline.Divisions)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", barlineIdent, __gong__toRawStringLiteral(barline.Id)))
			if barline.Bar_style != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + barline.Bar_style.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bar_style = %s", barlineIdent, targetIdent))
			}
			if barline.Footnote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + barline.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", barlineIdent, targetIdent))
			}
			if barline.Level != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + barline.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", barlineIdent, targetIdent))
			}
			if barline.Wavy_line != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + barline.Wavy_line.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wavy_line = %s", barlineIdent, targetIdent))
			}
			if barline.Segno_1 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + barline.Segno_1.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Segno_1 = %s", barlineIdent, targetIdent))
			}
			if barline.Coda_1 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + barline.Coda_1.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Coda_1 = %s", barlineIdent, targetIdent))
			}
			if barline.Fermata != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + barline.Fermata.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fermata = %s", barlineIdent, targetIdent))
			}
			if barline.Ending != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + barline.Ending.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ending = %s", barlineIdent, targetIdent))
			}
			if barline.Repeat != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + barline.Repeat.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Repeat = %s", barlineIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, barre := range __gong__sortStageSetInstances(stageSet.Stage.Barres, stageSet.Stage.Barre_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			barreIdent := "__models" + barre.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Barre{Name: %s}).Stage(stageSet.Stage)", barreIdent, __gong__toRawStringLiteral(barre.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", barreIdent, __gong__toRawStringLiteral(barre.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", barreIdent, __gong__toRawStringLiteral(barre.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", barreIdent, __gong__toRawStringLiteral(barre.Color)))
		}
	}
	if stageSet.Stage != nil {
		for _, bass := range __gong__sortStageSetInstances(stageSet.Stage.Basss, stageSet.Stage.Bass_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bassIdent := "__models" + bass.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Bass{Name: %s}).Stage(stageSet.Stage)", bassIdent, __gong__toRawStringLiteral(bass.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bassIdent, __gong__toRawStringLiteral(bass.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arrangement = %s", bassIdent, __gong__toRawStringLiteral(bass.Arrangement)))
			if bass.Bass_separator != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + bass.Bass_separator.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bass_separator = %s", bassIdent, targetIdent))
			}
			if bass.Bass_step != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + bass.Bass_step.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bass_step = %s", bassIdent, targetIdent))
			}
			if bass.Bass_alter != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + bass.Bass_alter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bass_alter = %s", bassIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, bass_step := range __gong__sortStageSetInstances(stageSet.Stage.Bass_steps, stageSet.Stage.Bass_step_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bass_stepIdent := "__models" + bass_step.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Bass_step{Name: %s}).Stage(stageSet.Stage)", bass_stepIdent, __gong__toRawStringLiteral(bass_step.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bass_stepIdent, __gong__toRawStringLiteral(bass_step.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", bass_stepIdent, __gong__toRawStringLiteral(bass_step.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", bass_stepIdent, __gong__toRawStringLiteral(bass_step.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", bass_stepIdent, __gong__toRawStringLiteral(bass_step.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", bass_stepIdent, __gong__toRawStringLiteral(bass_step.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", bass_stepIdent, __gong__toRawStringLiteral(bass_step.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", bass_stepIdent, __gong__toRawStringLiteral(bass_step.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", bass_stepIdent, __gong__toRawStringLiteral(bass_step.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", bass_stepIdent, __gong__toRawStringLiteral(bass_step.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", bass_stepIdent, __gong__toRawStringLiteral(bass_step.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", bass_stepIdent, __gong__toRawStringLiteral(bass_step.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", bass_stepIdent, __gong__toRawStringLiteral(bass_step.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, beam := range __gong__sortStageSetInstances(stageSet.Stage.Beams, stageSet.Stage.Beam_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			beamIdent := "__models" + beam.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Beam{Name: %s}).Stage(stageSet.Stage)", beamIdent, __gong__toRawStringLiteral(beam.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", beamIdent, __gong__toRawStringLiteral(beam.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", beamIdent, beam.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Repeater = %s", beamIdent, __gong__toRawStringLiteral(string(beam.Repeater))))
			values.WriteString(fmt.Sprintf("\n\t%s.Fan = %s", beamIdent, __gong__toRawStringLiteral(beam.Fan)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", beamIdent, __gong__toRawStringLiteral(beam.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", beamIdent, __gong__toRawStringLiteral(beam.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", beamIdent, __gong__toRawStringLiteral(beam.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, beat_repeat := range __gong__sortStageSetInstances(stageSet.Stage.Beat_repeats, stageSet.Stage.Beat_repeat_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			beat_repeatIdent := "__models" + beat_repeat.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Beat_repeat{Name: %s}).Stage(stageSet.Stage)", beat_repeatIdent, __gong__toRawStringLiteral(beat_repeat.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", beat_repeatIdent, __gong__toRawStringLiteral(beat_repeat.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", beat_repeatIdent, __gong__toRawStringLiteral(string(beat_repeat.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Slashes = %d", beat_repeatIdent, beat_repeat.Slashes))
			values.WriteString(fmt.Sprintf("\n\t%s.Use_dots = %s", beat_repeatIdent, __gong__toRawStringLiteral(string(beat_repeat.Use_dots))))
			values.WriteString(fmt.Sprintf("\n\t%s.Slash_type = %s", beat_repeatIdent, __gong__toRawStringLiteral(string(beat_repeat.Slash_type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Slash_dot = %s", beat_repeatIdent, __gong__toRawStringLiteral(beat_repeat.Slash_dot)))
			values.WriteString(fmt.Sprintf("\n\t%s.Except_voice = %s", beat_repeatIdent, __gong__toRawStringLiteral(beat_repeat.Except_voice)))
		}
	}
	if stageSet.Stage != nil {
		for _, beat_unit_tied := range __gong__sortStageSetInstances(stageSet.Stage.Beat_unit_tieds, stageSet.Stage.Beat_unit_tied_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			beat_unit_tiedIdent := "__models" + beat_unit_tied.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Beat_unit_tied{Name: %s}).Stage(stageSet.Stage)", beat_unit_tiedIdent, __gong__toRawStringLiteral(beat_unit_tied.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", beat_unit_tiedIdent, __gong__toRawStringLiteral(beat_unit_tied.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Beat_unit = %s", beat_unit_tiedIdent, __gong__toRawStringLiteral(string(beat_unit_tied.Beat_unit))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beat_unit_dot = %s", beat_unit_tiedIdent, __gong__toRawStringLiteral(beat_unit_tied.Beat_unit_dot)))
		}
	}
	if stageSet.Stage != nil {
		for _, beater := range __gong__sortStageSetInstances(stageSet.Stage.Beaters, stageSet.Stage.Beater_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			beaterIdent := "__models" + beater.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Beater{Name: %s}).Stage(stageSet.Stage)", beaterIdent, __gong__toRawStringLiteral(beater.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", beaterIdent, __gong__toRawStringLiteral(beater.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tip = %s", beaterIdent, __gong__toRawStringLiteral(beater.Tip)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", beaterIdent, __gong__toRawStringLiteral(beater.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, bend := range __gong__sortStageSetInstances(stageSet.Stage.Bends, stageSet.Stage.Bend_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bendIdent := "__models" + bend.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Bend{Name: %s}).Stage(stageSet.Stage)", bendIdent, __gong__toRawStringLiteral(bend.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bendIdent, __gong__toRawStringLiteral(bend.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Shape = %s", bendIdent, __gong__toRawStringLiteral(bend.Shape)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", bendIdent, __gong__toRawStringLiteral(bend.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", bendIdent, __gong__toRawStringLiteral(bend.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", bendIdent, __gong__toRawStringLiteral(bend.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", bendIdent, __gong__toRawStringLiteral(bend.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", bendIdent, __gong__toRawStringLiteral(bend.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", bendIdent, __gong__toRawStringLiteral(bend.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", bendIdent, __gong__toRawStringLiteral(bend.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", bendIdent, __gong__toRawStringLiteral(bend.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", bendIdent, __gong__toRawStringLiteral(bend.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Accelerate = %s", bendIdent, __gong__toRawStringLiteral(string(bend.Accelerate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", bendIdent, __gong__toRawStringLiteral(bend.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.First_beat = %s", bendIdent, __gong__toRawStringLiteral(bend.First_beat)))
			values.WriteString(fmt.Sprintf("\n\t%s.Last_beat = %s", bendIdent, __gong__toRawStringLiteral(bend.Last_beat)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bend_alter = %s", bendIdent, __gong__toRawStringLiteral(bend.Bend_alter)))
			values.WriteString(fmt.Sprintf("\n\t%s.Pre_bend = %s", bendIdent, __gong__toRawStringLiteral(bend.Pre_bend)))
			if bend.Release != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + bend.Release.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Release = %s", bendIdent, targetIdent))
			}
			if bend.With_bar != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + bend.With_bar.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.With_bar = %s", bendIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, bookmark := range __gong__sortStageSetInstances(stageSet.Stage.Bookmarks, stageSet.Stage.Bookmark_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bookmarkIdent := "__models" + bookmark.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Bookmark{Name: %s}).Stage(stageSet.Stage)", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Element = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Element)))
			values.WriteString(fmt.Sprintf("\n\t%s.Position = %d", bookmarkIdent, bookmark.Position))
		}
	}
	if stageSet.Stage != nil {
		for _, bracket := range __gong__sortStageSetInstances(stageSet.Stage.Brackets, stageSet.Stage.Bracket_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bracketIdent := "__models" + bracket.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Bracket{Name: %s}).Stage(stageSet.Stage)", bracketIdent, __gong__toRawStringLiteral(bracket.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", bracketIdent, bracket.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_end = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Line_end)))
			values.WriteString(fmt.Sprintf("\n\t%s.End_length = %s", bracketIdent, __gong__toRawStringLiteral(bracket.End_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_type = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Line_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dash_length = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Dash_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space_length = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Space_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, breath_mark := range __gong__sortStageSetInstances(stageSet.Stage.Breath_marks, stageSet.Stage.Breath_mark_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			breath_markIdent := "__models" + breath_mark.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Breath_mark{Name: %s}).Stage(stageSet.Stage)", breath_markIdent, __gong__toRawStringLiteral(breath_mark.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", breath_markIdent, __gong__toRawStringLiteral(breath_mark.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", breath_markIdent, __gong__toRawStringLiteral(breath_mark.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", breath_markIdent, __gong__toRawStringLiteral(breath_mark.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", breath_markIdent, __gong__toRawStringLiteral(breath_mark.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", breath_markIdent, __gong__toRawStringLiteral(breath_mark.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", breath_markIdent, __gong__toRawStringLiteral(breath_mark.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", breath_markIdent, __gong__toRawStringLiteral(breath_mark.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", breath_markIdent, __gong__toRawStringLiteral(breath_mark.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", breath_markIdent, __gong__toRawStringLiteral(breath_mark.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", breath_markIdent, __gong__toRawStringLiteral(breath_mark.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", breath_markIdent, __gong__toRawStringLiteral(breath_mark.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", breath_markIdent, __gong__toRawStringLiteral(breath_mark.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, caesura := range __gong__sortStageSetInstances(stageSet.Stage.Caesuras, stageSet.Stage.Caesura_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			caesuraIdent := "__models" + caesura.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Caesura{Name: %s}).Stage(stageSet.Stage)", caesuraIdent, __gong__toRawStringLiteral(caesura.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, cancel := range __gong__sortStageSetInstances(stageSet.Stage.Cancels, stageSet.Stage.Cancel_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cancelIdent := "__models" + cancel.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Cancel{Name: %s}).Stage(stageSet.Stage)", cancelIdent, __gong__toRawStringLiteral(cancel.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cancelIdent, __gong__toRawStringLiteral(cancel.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", cancelIdent, __gong__toRawStringLiteral(cancel.Location)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", cancelIdent, cancel.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		for _, clef := range __gong__sortStageSetInstances(stageSet.Stage.Clefs, stageSet.Stage.Clef_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			clefIdent := "__models" + clef.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Clef{Name: %s}).Stage(stageSet.Stage)", clefIdent, __gong__toRawStringLiteral(clef.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", clefIdent, __gong__toRawStringLiteral(clef.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", clefIdent, clef.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Additional = %s", clefIdent, __gong__toRawStringLiteral(string(clef.Additional))))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %s", clefIdent, __gong__toRawStringLiteral(clef.Size)))
			values.WriteString(fmt.Sprintf("\n\t%s.After_barline = %s", clefIdent, __gong__toRawStringLiteral(string(clef.After_barline))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", clefIdent, __gong__toRawStringLiteral(clef.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", clefIdent, __gong__toRawStringLiteral(clef.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", clefIdent, __gong__toRawStringLiteral(clef.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", clefIdent, __gong__toRawStringLiteral(clef.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", clefIdent, __gong__toRawStringLiteral(clef.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", clefIdent, __gong__toRawStringLiteral(clef.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", clefIdent, __gong__toRawStringLiteral(clef.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", clefIdent, __gong__toRawStringLiteral(clef.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", clefIdent, __gong__toRawStringLiteral(clef.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", clefIdent, __gong__toRawStringLiteral(string(clef.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", clefIdent, __gong__toRawStringLiteral(clef.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sign = %s", clefIdent, __gong__toRawStringLiteral(clef.Sign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", clefIdent, clef.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.Clef_octave_change = %d", clefIdent, clef.Clef_octave_change))
		}
	}
	if stageSet.Stage != nil {
		for _, coda := range __gong__sortStageSetInstances(stageSet.Stage.Codas, stageSet.Stage.Coda_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			codaIdent := "__models" + coda.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Coda{Name: %s}).Stage(stageSet.Stage)", codaIdent, __gong__toRawStringLiteral(coda.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codaIdent, __gong__toRawStringLiteral(coda.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", codaIdent, __gong__toRawStringLiteral(coda.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", codaIdent, __gong__toRawStringLiteral(coda.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", codaIdent, __gong__toRawStringLiteral(coda.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", codaIdent, __gong__toRawStringLiteral(coda.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", codaIdent, __gong__toRawStringLiteral(coda.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", codaIdent, __gong__toRawStringLiteral(coda.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", codaIdent, __gong__toRawStringLiteral(coda.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", codaIdent, __gong__toRawStringLiteral(coda.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", codaIdent, __gong__toRawStringLiteral(coda.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", codaIdent, __gong__toRawStringLiteral(coda.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", codaIdent, __gong__toRawStringLiteral(coda.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", codaIdent, __gong__toRawStringLiteral(coda.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", codaIdent, __gong__toRawStringLiteral(coda.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, credit := range __gong__sortStageSetInstances(stageSet.Stage.Credits, stageSet.Stage.Credit_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			creditIdent := "__models" + credit.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Credit{Name: %s}).Stage(stageSet.Stage)", creditIdent, __gong__toRawStringLiteral(credit.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", creditIdent, __gong__toRawStringLiteral(credit.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Page = %d", creditIdent, credit.Page))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", creditIdent, __gong__toRawStringLiteral(credit.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Credit_type = %s", creditIdent, __gong__toRawStringLiteral(credit.Credit_type)))
			if credit.Credit_image != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + credit.Credit_image.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit_image = %s", creditIdent, targetIdent))
			}
			for _, elem := range credit.Link {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Link = append(%s.Link, %s)", creditIdent, creditIdent, targetIdent))
			}
			for _, elem := range credit.Bookmark {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bookmark = append(%s.Bookmark, %s)", creditIdent, creditIdent, targetIdent))
			}
			for _, elem := range credit.Credit_words {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit_words = append(%s.Credit_words, %s)", creditIdent, creditIdent, targetIdent))
			}
			for _, elem := range credit.Credit_symbol {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit_symbol = append(%s.Credit_symbol, %s)", creditIdent, creditIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, dashes := range __gong__sortStageSetInstances(stageSet.Stage.Dashess, stageSet.Stage.Dashes_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dashesIdent := "__models" + dashes.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Dashes{Name: %s}).Stage(stageSet.Stage)", dashesIdent, __gong__toRawStringLiteral(dashes.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dashesIdent, __gong__toRawStringLiteral(dashes.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", dashesIdent, __gong__toRawStringLiteral(string(dashes.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", dashesIdent, dashes.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Dash_length = %s", dashesIdent, __gong__toRawStringLiteral(dashes.Dash_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space_length = %s", dashesIdent, __gong__toRawStringLiteral(dashes.Space_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", dashesIdent, __gong__toRawStringLiteral(dashes.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", dashesIdent, __gong__toRawStringLiteral(dashes.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", dashesIdent, __gong__toRawStringLiteral(dashes.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", dashesIdent, __gong__toRawStringLiteral(dashes.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", dashesIdent, __gong__toRawStringLiteral(dashes.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", dashesIdent, __gong__toRawStringLiteral(dashes.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, defaults := range __gong__sortStageSetInstances(stageSet.Stage.Defaultss, stageSet.Stage.Defaults_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			defaultsIdent := "__models" + defaults.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Defaults{Name: %s}).Stage(stageSet.Stage)", defaultsIdent, __gong__toRawStringLiteral(defaults.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", defaultsIdent, __gong__toRawStringLiteral(defaults.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Concert_score = %s", defaultsIdent, __gong__toRawStringLiteral(defaults.Concert_score)))
			if defaults.Scaling != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + defaults.Scaling.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Scaling = %s", defaultsIdent, targetIdent))
			}
			if defaults.Page_layout != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + defaults.Page_layout.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Page_layout = %s", defaultsIdent, targetIdent))
			}
			if defaults.System_layout != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + defaults.System_layout.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System_layout = %s", defaultsIdent, targetIdent))
			}
			for _, elem := range defaults.Staff_layout {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staff_layout = append(%s.Staff_layout, %s)", defaultsIdent, defaultsIdent, targetIdent))
			}
			if defaults.Appearance != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + defaults.Appearance.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Appearance = %s", defaultsIdent, targetIdent))
			}
			if defaults.Music_font != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + defaults.Music_font.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Music_font = %s", defaultsIdent, targetIdent))
			}
			if defaults.Word_font != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + defaults.Word_font.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Word_font = %s", defaultsIdent, targetIdent))
			}
			for _, elem := range defaults.Lyric_font {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lyric_font = append(%s.Lyric_font, %s)", defaultsIdent, defaultsIdent, targetIdent))
			}
			for _, elem := range defaults.Lyric_language {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lyric_language = append(%s.Lyric_language, %s)", defaultsIdent, defaultsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, degree := range __gong__sortStageSetInstances(stageSet.Stage.Degrees, stageSet.Stage.Degree_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			degreeIdent := "__models" + degree.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Degree{Name: %s}).Stage(stageSet.Stage)", degreeIdent, __gong__toRawStringLiteral(degree.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", degreeIdent, __gong__toRawStringLiteral(degree.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", degreeIdent, __gong__toRawStringLiteral(string(degree.Print_object))))
			if degree.Degree_value != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + degree.Degree_value.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Degree_value = %s", degreeIdent, targetIdent))
			}
			if degree.Degree_alter != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + degree.Degree_alter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Degree_alter = %s", degreeIdent, targetIdent))
			}
			if degree.Degree_type != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + degree.Degree_type.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Degree_type = %s", degreeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, degree_alter := range __gong__sortStageSetInstances(stageSet.Stage.Degree_alters, stageSet.Stage.Degree_alter_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			degree_alterIdent := "__models" + degree_alter.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Degree_alter{Name: %s}).Stage(stageSet.Stage)", degree_alterIdent, __gong__toRawStringLiteral(degree_alter.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", degree_alterIdent, __gong__toRawStringLiteral(degree_alter.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Plus_minus = %s", degree_alterIdent, __gong__toRawStringLiteral(string(degree_alter.Plus_minus))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", degree_alterIdent, __gong__toRawStringLiteral(degree_alter.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", degree_alterIdent, __gong__toRawStringLiteral(degree_alter.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", degree_alterIdent, __gong__toRawStringLiteral(degree_alter.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", degree_alterIdent, __gong__toRawStringLiteral(degree_alter.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", degree_alterIdent, __gong__toRawStringLiteral(degree_alter.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", degree_alterIdent, __gong__toRawStringLiteral(degree_alter.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", degree_alterIdent, __gong__toRawStringLiteral(degree_alter.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", degree_alterIdent, __gong__toRawStringLiteral(degree_alter.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", degree_alterIdent, __gong__toRawStringLiteral(degree_alter.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", degree_alterIdent, __gong__toRawStringLiteral(degree_alter.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, degree_type := range __gong__sortStageSetInstances(stageSet.Stage.Degree_types, stageSet.Stage.Degree_type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			degree_typeIdent := "__models" + degree_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Degree_type{Name: %s}).Stage(stageSet.Stage)", degree_typeIdent, __gong__toRawStringLiteral(degree_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", degree_typeIdent, __gong__toRawStringLiteral(degree_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", degree_typeIdent, __gong__toRawStringLiteral(degree_type.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", degree_typeIdent, __gong__toRawStringLiteral(degree_type.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", degree_typeIdent, __gong__toRawStringLiteral(degree_type.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", degree_typeIdent, __gong__toRawStringLiteral(degree_type.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", degree_typeIdent, __gong__toRawStringLiteral(degree_type.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", degree_typeIdent, __gong__toRawStringLiteral(degree_type.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", degree_typeIdent, __gong__toRawStringLiteral(degree_type.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", degree_typeIdent, __gong__toRawStringLiteral(degree_type.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", degree_typeIdent, __gong__toRawStringLiteral(degree_type.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", degree_typeIdent, __gong__toRawStringLiteral(degree_type.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", degree_typeIdent, __gong__toRawStringLiteral(degree_type.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, degree_value := range __gong__sortStageSetInstances(stageSet.Stage.Degree_values, stageSet.Stage.Degree_value_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			degree_valueIdent := "__models" + degree_value.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Degree_value{Name: %s}).Stage(stageSet.Stage)", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Symbol = %s", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Symbol)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", degree_valueIdent, degree_value.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		for _, direction := range __gong__sortStageSetInstances(stageSet.Stage.Directions, stageSet.Stage.Direction_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directionIdent := "__models" + direction.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Direction{Name: %s}).Stage(stageSet.Stage)", directionIdent, __gong__toRawStringLiteral(direction.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionIdent, __gong__toRawStringLiteral(direction.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", directionIdent, __gong__toRawStringLiteral(direction.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Directive = %s", directionIdent, __gong__toRawStringLiteral(string(direction.Directive))))
			values.WriteString(fmt.Sprintf("\n\t%s.System = %s", directionIdent, __gong__toRawStringLiteral(string(direction.System))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directionIdent, __gong__toRawStringLiteral(direction.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Voice = %s", directionIdent, __gong__toRawStringLiteral(direction.Voice)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff = %d", directionIdent, direction.Staff))
			for _, elem := range direction.Direction_type {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Direction_type = append(%s.Direction_type, %s)", directionIdent, directionIdent, targetIdent))
			}
			if direction.Offset != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction.Offset.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", directionIdent, targetIdent))
			}
			if direction.Footnote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", directionIdent, targetIdent))
			}
			if direction.Level != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", directionIdent, targetIdent))
			}
			if direction.Sound != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction.Sound.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sound = %s", directionIdent, targetIdent))
			}
			if direction.Listening != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction.Listening.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listening = %s", directionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, direction_type := range __gong__sortStageSetInstances(stageSet.Stage.Direction_types, stageSet.Stage.Direction_type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			direction_typeIdent := "__models" + direction_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Direction_type{Name: %s}).Stage(stageSet.Stage)", direction_typeIdent, __gong__toRawStringLiteral(direction_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", direction_typeIdent, __gong__toRawStringLiteral(direction_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", direction_typeIdent, __gong__toRawStringLiteral(direction_type.Id)))
			for _, elem := range direction_type.Rehearsal {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rehearsal = append(%s.Rehearsal, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			for _, elem := range direction_type.Segno {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Segno = append(%s.Segno, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			for _, elem := range direction_type.Coda {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Coda = append(%s.Coda, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			for _, elem := range direction_type.Words {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Words = append(%s.Words, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			for _, elem := range direction_type.Symbol {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Symbol = append(%s.Symbol, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			if direction_type.Wedge != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Wedge.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wedge = %s", direction_typeIdent, targetIdent))
			}
			for _, elem := range direction_type.Dynamics {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dynamics = append(%s.Dynamics, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			if direction_type.Dashes != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Dashes.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dashes = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Bracket != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Bracket.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bracket = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Pedal != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Pedal.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pedal = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Metronome != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Metronome.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metronome = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Octave_shift != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Octave_shift.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Octave_shift = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Harp_pedals != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Harp_pedals.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harp_pedals = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Damp != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Damp.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Damp = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Damp_all != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Damp_all.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Damp_all = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Eyeglasses != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Eyeglasses.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Eyeglasses = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.String_mute != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.String_mute.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.String_mute = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Scordatura != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Scordatura.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Scordatura = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Image != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Image.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Image = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Principal_voice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Principal_voice.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Principal_voice = %s", direction_typeIdent, targetIdent))
			}
			for _, elem := range direction_type.Percussion {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Percussion = append(%s.Percussion, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			if direction_type.Accordion_registration != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Accordion_registration.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accordion_registration = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Staff_divide != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Staff_divide.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staff_divide = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Other_direction != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + direction_type.Other_direction.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_direction = %s", direction_typeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, distance := range __gong__sortStageSetInstances(stageSet.Stage.Distances, stageSet.Stage.Distance_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			distanceIdent := "__models" + distance.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Distance{Name: %s}).Stage(stageSet.Stage)", distanceIdent, __gong__toRawStringLiteral(distance.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", distanceIdent, __gong__toRawStringLiteral(distance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", distanceIdent, __gong__toRawStringLiteral(distance.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", distanceIdent, __gong__toRawStringLiteral(distance.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, double := range __gong__sortStageSetInstances(stageSet.Stage.Doubles, stageSet.Stage.Double_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			doubleIdent := "__models" + double.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Double{Name: %s}).Stage(stageSet.Stage)", doubleIdent, __gong__toRawStringLiteral(double.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", doubleIdent, __gong__toRawStringLiteral(double.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Above = %s", doubleIdent, __gong__toRawStringLiteral(string(double.Above))))
		}
	}
	if stageSet.Stage != nil {
		for _, dynamics := range __gong__sortStageSetInstances(stageSet.Stage.Dynamicss, stageSet.Stage.Dynamics_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			dynamicsIdent := "__models" + dynamics.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Dynamics{Name: %s}).Stage(stageSet.Stage)", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", dynamicsIdent, dynamics.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", dynamicsIdent, dynamics.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_through = %d", dynamicsIdent, dynamics.Line_through))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Enclosure)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.P = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.P)))
			values.WriteString(fmt.Sprintf("\n\t%s.Pp = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Pp)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ppp = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Ppp)))
			values.WriteString(fmt.Sprintf("\n\t%s.Pppp = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Pppp)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ppppp = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Ppppp)))
			values.WriteString(fmt.Sprintf("\n\t%s.Pppppp = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Pppppp)))
			values.WriteString(fmt.Sprintf("\n\t%s.F = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.F)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ff = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Ff)))
			values.WriteString(fmt.Sprintf("\n\t%s.Fff = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Fff)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ffff = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Ffff)))
			values.WriteString(fmt.Sprintf("\n\t%s.Fffff = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Fffff)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ffffff = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Ffffff)))
			values.WriteString(fmt.Sprintf("\n\t%s.Mp = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Mp)))
			values.WriteString(fmt.Sprintf("\n\t%s.Mf = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Mf)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sf = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Sf)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sfp = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Sfp)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sfpp = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Sfpp)))
			values.WriteString(fmt.Sprintf("\n\t%s.Fp = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Fp)))
			values.WriteString(fmt.Sprintf("\n\t%s.Rf = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Rf)))
			values.WriteString(fmt.Sprintf("\n\t%s.Rfz = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Rfz)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sfz = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Sfz)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sffz = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Sffz)))
			values.WriteString(fmt.Sprintf("\n\t%s.Fz = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Fz)))
			values.WriteString(fmt.Sprintf("\n\t%s.N = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.N)))
			values.WriteString(fmt.Sprintf("\n\t%s.Pf = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Pf)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sfzp = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Sfzp)))
			for _, elem := range dynamics.Other_dynamics {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_dynamics = append(%s.Other_dynamics, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, effect := range __gong__sortStageSetInstances(stageSet.Stage.Effects, stageSet.Stage.Effect_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			effectIdent := "__models" + effect.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Effect{Name: %s}).Stage(stageSet.Stage)", effectIdent, __gong__toRawStringLiteral(effect.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", effectIdent, __gong__toRawStringLiteral(effect.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", effectIdent, __gong__toRawStringLiteral(effect.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", effectIdent, __gong__toRawStringLiteral(effect.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, elision := range __gong__sortStageSetInstances(stageSet.Stage.Elisions, stageSet.Stage.Elision_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elisionIdent := "__models" + elision.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Elision{Name: %s}).Stage(stageSet.Stage)", elisionIdent, __gong__toRawStringLiteral(elision.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elisionIdent, __gong__toRawStringLiteral(elision.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", elisionIdent, __gong__toRawStringLiteral(elision.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", elisionIdent, __gong__toRawStringLiteral(elision.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", elisionIdent, __gong__toRawStringLiteral(elision.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", elisionIdent, __gong__toRawStringLiteral(elision.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", elisionIdent, __gong__toRawStringLiteral(elision.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", elisionIdent, __gong__toRawStringLiteral(elision.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", elisionIdent, __gong__toRawStringLiteral(elision.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, empty := range __gong__sortStageSetInstances(stageSet.Stage.Emptys, stageSet.Stage.Empty_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			emptyIdent := "__models" + empty.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Empty{Name: %s}).Stage(stageSet.Stage)", emptyIdent, __gong__toRawStringLiteral(empty.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", emptyIdent, __gong__toRawStringLiteral(empty.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, empty_font := range __gong__sortStageSetInstances(stageSet.Stage.Empty_fonts, stageSet.Stage.Empty_font_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			empty_fontIdent := "__models" + empty_font.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Empty_font{Name: %s}).Stage(stageSet.Stage)", empty_fontIdent, __gong__toRawStringLiteral(empty_font.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", empty_fontIdent, __gong__toRawStringLiteral(empty_font.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", empty_fontIdent, __gong__toRawStringLiteral(empty_font.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", empty_fontIdent, __gong__toRawStringLiteral(empty_font.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", empty_fontIdent, __gong__toRawStringLiteral(empty_font.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", empty_fontIdent, __gong__toRawStringLiteral(empty_font.Font_weight)))
		}
	}
	if stageSet.Stage != nil {
		for _, empty_line := range __gong__sortStageSetInstances(stageSet.Stage.Empty_lines, stageSet.Stage.Empty_line_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			empty_lineIdent := "__models" + empty_line.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Empty_line{Name: %s}).Stage(stageSet.Stage)", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_shape = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Line_shape)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_type = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Line_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_length = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Line_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dash_length = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Dash_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space_length = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Space_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Placement)))
		}
	}
	if stageSet.Stage != nil {
		for _, empty_placement := range __gong__sortStageSetInstances(stageSet.Stage.Empty_placements, stageSet.Stage.Empty_placement_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			empty_placementIdent := "__models" + empty_placement.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Empty_placement{Name: %s}).Stage(stageSet.Stage)", empty_placementIdent, __gong__toRawStringLiteral(empty_placement.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", empty_placementIdent, __gong__toRawStringLiteral(empty_placement.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", empty_placementIdent, __gong__toRawStringLiteral(empty_placement.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", empty_placementIdent, __gong__toRawStringLiteral(empty_placement.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", empty_placementIdent, __gong__toRawStringLiteral(empty_placement.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", empty_placementIdent, __gong__toRawStringLiteral(empty_placement.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", empty_placementIdent, __gong__toRawStringLiteral(empty_placement.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", empty_placementIdent, __gong__toRawStringLiteral(empty_placement.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", empty_placementIdent, __gong__toRawStringLiteral(empty_placement.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", empty_placementIdent, __gong__toRawStringLiteral(empty_placement.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", empty_placementIdent, __gong__toRawStringLiteral(empty_placement.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", empty_placementIdent, __gong__toRawStringLiteral(empty_placement.Placement)))
		}
	}
	if stageSet.Stage != nil {
		for _, empty_placement_smufl := range __gong__sortStageSetInstances(stageSet.Stage.Empty_placement_smufls, stageSet.Stage.Empty_placement_smufl_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			empty_placement_smuflIdent := "__models" + empty_placement_smufl.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Empty_placement_smufl{Name: %s}).Stage(stageSet.Stage)", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Smufl)))
		}
	}
	if stageSet.Stage != nil {
		for _, empty_print_object_style_align := range __gong__sortStageSetInstances(stageSet.Stage.Empty_print_object_style_aligns, stageSet.Stage.Empty_print_object_style_align_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			empty_print_object_style_alignIdent := "__models" + empty_print_object_style_align.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Empty_print_object_style_align{Name: %s}).Stage(stageSet.Stage)", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(string(empty_print_object_style_align.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Valign)))
		}
	}
	if stageSet.Stage != nil {
		for _, empty_print_style := range __gong__sortStageSetInstances(stageSet.Stage.Empty_print_styles, stageSet.Stage.Empty_print_style_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			empty_print_styleIdent := "__models" + empty_print_style.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Empty_print_style{Name: %s}).Stage(stageSet.Stage)", empty_print_styleIdent, __gong__toRawStringLiteral(empty_print_style.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", empty_print_styleIdent, __gong__toRawStringLiteral(empty_print_style.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", empty_print_styleIdent, __gong__toRawStringLiteral(empty_print_style.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", empty_print_styleIdent, __gong__toRawStringLiteral(empty_print_style.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", empty_print_styleIdent, __gong__toRawStringLiteral(empty_print_style.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", empty_print_styleIdent, __gong__toRawStringLiteral(empty_print_style.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", empty_print_styleIdent, __gong__toRawStringLiteral(empty_print_style.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", empty_print_styleIdent, __gong__toRawStringLiteral(empty_print_style.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", empty_print_styleIdent, __gong__toRawStringLiteral(empty_print_style.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", empty_print_styleIdent, __gong__toRawStringLiteral(empty_print_style.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", empty_print_styleIdent, __gong__toRawStringLiteral(empty_print_style.Color)))
		}
	}
	if stageSet.Stage != nil {
		for _, empty_print_style_align := range __gong__sortStageSetInstances(stageSet.Stage.Empty_print_style_aligns, stageSet.Stage.Empty_print_style_align_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			empty_print_style_alignIdent := "__models" + empty_print_style_align.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Empty_print_style_align{Name: %s}).Stage(stageSet.Stage)", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Valign)))
		}
	}
	if stageSet.Stage != nil {
		for _, empty_print_style_align_id := range __gong__sortStageSetInstances(stageSet.Stage.Empty_print_style_align_ids, stageSet.Stage.Empty_print_style_align_id_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			empty_print_style_align_idIdent := "__models" + empty_print_style_align_id.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Empty_print_style_align_id{Name: %s}).Stage(stageSet.Stage)", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, empty_trill_sound := range __gong__sortStageSetInstances(stageSet.Stage.Empty_trill_sounds, stageSet.Stage.Empty_trill_sound_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			empty_trill_soundIdent := "__models" + empty_trill_sound.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Empty_trill_sound{Name: %s}).Stage(stageSet.Stage)", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Start_note = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Start_note)))
			values.WriteString(fmt.Sprintf("\n\t%s.Trill_step = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Trill_step)))
			values.WriteString(fmt.Sprintf("\n\t%s.Two_note_turn = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Two_note_turn)))
			values.WriteString(fmt.Sprintf("\n\t%s.Accelerate = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(string(empty_trill_sound.Accelerate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.Second_beat = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Second_beat)))
			values.WriteString(fmt.Sprintf("\n\t%s.Last_beat = %s", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Last_beat)))
		}
	}
	if stageSet.Stage != nil {
		for _, encoding := range __gong__sortStageSetInstances(stageSet.Stage.Encodings, stageSet.Stage.Encoding_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			encodingIdent := "__models" + encoding.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Encoding{Name: %s}).Stage(stageSet.Stage)", encodingIdent, __gong__toRawStringLiteral(encoding.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", encodingIdent, __gong__toRawStringLiteral(encoding.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Software = %s", encodingIdent, __gong__toRawStringLiteral(encoding.Software)))
			values.WriteString(fmt.Sprintf("\n\t%s.Encoding_description = %s", encodingIdent, __gong__toRawStringLiteral(encoding.Encoding_description)))
			for _, elem := range encoding.Encoder {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Encoder = append(%s.Encoder, %s)", encodingIdent, encodingIdent, targetIdent))
			}
			for _, elem := range encoding.Supports {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Supports = append(%s.Supports, %s)", encodingIdent, encodingIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, ending := range __gong__sortStageSetInstances(stageSet.Stage.Endings, stageSet.Stage.Ending_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			endingIdent := "__models" + ending.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Ending{Name: %s}).Stage(stageSet.Stage)", endingIdent, __gong__toRawStringLiteral(ending.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", endingIdent, __gong__toRawStringLiteral(ending.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", endingIdent, __gong__toRawStringLiteral(ending.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", endingIdent, __gong__toRawStringLiteral(ending.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.End_length = %s", endingIdent, __gong__toRawStringLiteral(ending.End_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text_x = %s", endingIdent, __gong__toRawStringLiteral(ending.Text_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text_y = %s", endingIdent, __gong__toRawStringLiteral(ending.Text_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", endingIdent, __gong__toRawStringLiteral(string(ending.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", endingIdent, __gong__toRawStringLiteral(ending.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", endingIdent, __gong__toRawStringLiteral(ending.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", endingIdent, __gong__toRawStringLiteral(ending.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", endingIdent, __gong__toRawStringLiteral(ending.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", endingIdent, __gong__toRawStringLiteral(ending.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", endingIdent, __gong__toRawStringLiteral(ending.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", endingIdent, __gong__toRawStringLiteral(ending.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", endingIdent, __gong__toRawStringLiteral(ending.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", endingIdent, __gong__toRawStringLiteral(ending.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.System = %s", endingIdent, __gong__toRawStringLiteral(string(ending.System))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", endingIdent, __gong__toRawStringLiteral(ending.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, extend := range __gong__sortStageSetInstances(stageSet.Stage.Extends, stageSet.Stage.Extend_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			extendIdent := "__models" + extend.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Extend{Name: %s}).Stage(stageSet.Stage)", extendIdent, __gong__toRawStringLiteral(extend.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", extendIdent, __gong__toRawStringLiteral(extend.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", extendIdent, __gong__toRawStringLiteral(string(extend.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", extendIdent, __gong__toRawStringLiteral(extend.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", extendIdent, __gong__toRawStringLiteral(extend.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", extendIdent, __gong__toRawStringLiteral(extend.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", extendIdent, __gong__toRawStringLiteral(extend.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", extendIdent, __gong__toRawStringLiteral(extend.Color)))
		}
	}
	if stageSet.Stage != nil {
		for _, feature := range __gong__sortStageSetInstances(stageSet.Stage.Features, stageSet.Stage.Feature_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			featureIdent := "__models" + feature.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Feature{Name: %s}).Stage(stageSet.Stage)", featureIdent, __gong__toRawStringLiteral(feature.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", featureIdent, __gong__toRawStringLiteral(feature.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", featureIdent, __gong__toRawStringLiteral(feature.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", featureIdent, __gong__toRawStringLiteral(feature.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, fermata := range __gong__sortStageSetInstances(stageSet.Stage.Fermatas, stageSet.Stage.Fermata_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			fermataIdent := "__models" + fermata.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Fermata{Name: %s}).Stage(stageSet.Stage)", fermataIdent, __gong__toRawStringLiteral(fermata.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", fermataIdent, __gong__toRawStringLiteral(fermata.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, figure := range __gong__sortStageSetInstances(stageSet.Stage.Figures, stageSet.Stage.Figure_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			figureIdent := "__models" + figure.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Figure{Name: %s}).Stage(stageSet.Stage)", figureIdent, __gong__toRawStringLiteral(figure.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", figureIdent, __gong__toRawStringLiteral(figure.Name)))
			if figure.Prefix != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + figure.Prefix.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Prefix = %s", figureIdent, targetIdent))
			}
			if figure.Figure_number != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + figure.Figure_number.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Figure_number = %s", figureIdent, targetIdent))
			}
			if figure.Suffix != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + figure.Suffix.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Suffix = %s", figureIdent, targetIdent))
			}
			if figure.Extend != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + figure.Extend.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extend = %s", figureIdent, targetIdent))
			}
			if figure.Footnote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + figure.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", figureIdent, targetIdent))
			}
			if figure.Level != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + figure.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", figureIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, figured_bass := range __gong__sortStageSetInstances(stageSet.Stage.Figured_basss, stageSet.Stage.Figured_bass_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			figured_bassIdent := "__models" + figured_bass.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Figured_bass{Name: %s}).Stage(stageSet.Stage)", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", figured_bassIdent, __gong__toRawStringLiteral(string(figured_bass.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_dot = %s", figured_bassIdent, __gong__toRawStringLiteral(string(figured_bass.Print_dot))))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_lyric = %s", figured_bassIdent, __gong__toRawStringLiteral(string(figured_bass.Print_lyric))))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", figured_bassIdent, __gong__toRawStringLiteral(string(figured_bass.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_spacing = %s", figured_bassIdent, __gong__toRawStringLiteral(string(figured_bass.Print_spacing))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = %s", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Duration)))
			for _, elem := range figured_bass.Figure {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Figure = append(%s.Figure, %s)", figured_bassIdent, figured_bassIdent, targetIdent))
			}
			if figured_bass.Footnote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + figured_bass.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", figured_bassIdent, targetIdent))
			}
			if figured_bass.Level != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + figured_bass.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", figured_bassIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, fingering := range __gong__sortStageSetInstances(stageSet.Stage.Fingerings, stageSet.Stage.Fingering_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			fingeringIdent := "__models" + fingering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Fingering{Name: %s}).Stage(stageSet.Stage)", fingeringIdent, __gong__toRawStringLiteral(fingering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Substitution = %s", fingeringIdent, __gong__toRawStringLiteral(string(fingering.Substitution))))
			values.WriteString(fmt.Sprintf("\n\t%s.Alternate = %s", fingeringIdent, __gong__toRawStringLiteral(string(fingering.Alternate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, first_fret := range __gong__sortStageSetInstances(stageSet.Stage.First_frets, stageSet.Stage.First_fret_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			first_fretIdent := "__models" + first_fret.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.First_fret{Name: %s}).Stage(stageSet.Stage)", first_fretIdent, __gong__toRawStringLiteral(first_fret.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", first_fretIdent, __gong__toRawStringLiteral(first_fret.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", first_fretIdent, __gong__toRawStringLiteral(first_fret.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", first_fretIdent, __gong__toRawStringLiteral(first_fret.Location)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", first_fretIdent, first_fret.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		for _, for_part := range __gong__sortStageSetInstances(stageSet.Stage.For_parts, stageSet.Stage.For_part_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			for_partIdent := "__models" + for_part.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.For_part{Name: %s}).Stage(stageSet.Stage)", for_partIdent, __gong__toRawStringLiteral(for_part.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", for_partIdent, __gong__toRawStringLiteral(for_part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", for_partIdent, for_part.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", for_partIdent, __gong__toRawStringLiteral(for_part.Id)))
			if for_part.Part_clef != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + for_part.Part_clef.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_clef = %s", for_partIdent, targetIdent))
			}
			if for_part.Part_transpose != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + for_part.Part_transpose.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_transpose = %s", for_partIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, formatted_symbol := range __gong__sortStageSetInstances(stageSet.Stage.Formatted_symbols, stageSet.Stage.Formatted_symbol_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formatted_symbolIdent := "__models" + formatted_symbol.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Formatted_symbol{Name: %s}).Stage(stageSet.Stage)", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", formatted_symbolIdent, __gong__toRawStringLiteral(string(formatted_symbol.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", formatted_symbolIdent, formatted_symbol.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", formatted_symbolIdent, formatted_symbol.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_through = %d", formatted_symbolIdent, formatted_symbol.Line_through))
			values.WriteString(fmt.Sprintf("\n\t%s.Rotation = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Rotation)))
			values.WriteString(fmt.Sprintf("\n\t%s.Letter_spacing = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Letter_spacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_height = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Line_height)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dir = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Dir)))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Enclosure)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, formatted_symbol_id := range __gong__sortStageSetInstances(stageSet.Stage.Formatted_symbol_ids, stageSet.Stage.Formatted_symbol_id_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formatted_symbol_idIdent := "__models" + formatted_symbol_id.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Formatted_symbol_id{Name: %s}).Stage(stageSet.Stage)", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(string(formatted_symbol_id.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", formatted_symbol_idIdent, formatted_symbol_id.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", formatted_symbol_idIdent, formatted_symbol_id.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_through = %d", formatted_symbol_idIdent, formatted_symbol_id.Line_through))
			values.WriteString(fmt.Sprintf("\n\t%s.Rotation = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Rotation)))
			values.WriteString(fmt.Sprintf("\n\t%s.Letter_spacing = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Letter_spacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_height = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Line_height)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dir = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Dir)))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Enclosure)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, formatted_text := range __gong__sortStageSetInstances(stageSet.Stage.Formatted_texts, stageSet.Stage.Formatted_text_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formatted_textIdent := "__models" + formatted_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Formatted_text{Name: %s}).Stage(stageSet.Stage)", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Space)))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", formatted_textIdent, __gong__toRawStringLiteral(string(formatted_text.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", formatted_textIdent, formatted_text.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", formatted_textIdent, formatted_text.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_through = %d", formatted_textIdent, formatted_text.Line_through))
			values.WriteString(fmt.Sprintf("\n\t%s.Rotation = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Rotation)))
			values.WriteString(fmt.Sprintf("\n\t%s.Letter_spacing = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Letter_spacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_height = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Line_height)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dir = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Dir)))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Enclosure)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, formatted_text_id := range __gong__sortStageSetInstances(stageSet.Stage.Formatted_text_ids, stageSet.Stage.Formatted_text_id_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formatted_text_idIdent := "__models" + formatted_text_id.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Formatted_text_id{Name: %s}).Stage(stageSet.Stage)", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Space)))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", formatted_text_idIdent, __gong__toRawStringLiteral(string(formatted_text_id.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", formatted_text_idIdent, formatted_text_id.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", formatted_text_idIdent, formatted_text_id.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_through = %d", formatted_text_idIdent, formatted_text_id.Line_through))
			values.WriteString(fmt.Sprintf("\n\t%s.Rotation = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Rotation)))
			values.WriteString(fmt.Sprintf("\n\t%s.Letter_spacing = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Letter_spacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_height = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Line_height)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dir = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Dir)))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Enclosure)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, forward := range __gong__sortStageSetInstances(stageSet.Stage.Forwards, stageSet.Stage.Forward_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			forwardIdent := "__models" + forward.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Forward{Name: %s}).Stage(stageSet.Stage)", forwardIdent, __gong__toRawStringLiteral(forward.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", forwardIdent, __gong__toRawStringLiteral(forward.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = %s", forwardIdent, __gong__toRawStringLiteral(forward.Duration)))
			values.WriteString(fmt.Sprintf("\n\t%s.Voice = %s", forwardIdent, __gong__toRawStringLiteral(forward.Voice)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff = %d", forwardIdent, forward.Staff))
			if forward.Footnote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + forward.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", forwardIdent, targetIdent))
			}
			if forward.Level != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + forward.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", forwardIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, frame := range __gong__sortStageSetInstances(stageSet.Stage.Frames, stageSet.Stage.Frame_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			frameIdent := "__models" + frame.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Frame{Name: %s}).Stage(stageSet.Stage)", frameIdent, __gong__toRawStringLiteral(frame.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", frameIdent, __gong__toRawStringLiteral(frame.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %s", frameIdent, __gong__toRawStringLiteral(frame.Height)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", frameIdent, __gong__toRawStringLiteral(frame.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Unplayed = %s", frameIdent, __gong__toRawStringLiteral(frame.Unplayed)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", frameIdent, __gong__toRawStringLiteral(frame.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", frameIdent, __gong__toRawStringLiteral(frame.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", frameIdent, __gong__toRawStringLiteral(frame.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", frameIdent, __gong__toRawStringLiteral(frame.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", frameIdent, __gong__toRawStringLiteral(frame.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", frameIdent, __gong__toRawStringLiteral(frame.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", frameIdent, __gong__toRawStringLiteral(frame.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", frameIdent, __gong__toRawStringLiteral(frame.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Frame_strings = %d", frameIdent, frame.Frame_strings))
			values.WriteString(fmt.Sprintf("\n\t%s.Frame_frets = %d", frameIdent, frame.Frame_frets))
			if frame.First_fret != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + frame.First_fret.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.First_fret = %s", frameIdent, targetIdent))
			}
			for _, elem := range frame.Frame_note {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frame_note = append(%s.Frame_note, %s)", frameIdent, frameIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, frame_note := range __gong__sortStageSetInstances(stageSet.Stage.Frame_notes, stageSet.Stage.Frame_note_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			frame_noteIdent := "__models" + frame_note.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Frame_note{Name: %s}).Stage(stageSet.Stage)", frame_noteIdent, __gong__toRawStringLiteral(frame_note.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", frame_noteIdent, __gong__toRawStringLiteral(frame_note.Name)))
			if frame_note.String != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + frame_note.String.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.String = %s", frame_noteIdent, targetIdent))
			}
			if frame_note.Fret != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + frame_note.Fret.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fret = %s", frame_noteIdent, targetIdent))
			}
			if frame_note.Fingering != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + frame_note.Fingering.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fingering = %s", frame_noteIdent, targetIdent))
			}
			if frame_note.Barre != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + frame_note.Barre.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Barre = %s", frame_noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, fret := range __gong__sortStageSetInstances(stageSet.Stage.Frets, stageSet.Stage.Fret_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			fretIdent := "__models" + fret.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Fret{Name: %s}).Stage(stageSet.Stage)", fretIdent, __gong__toRawStringLiteral(fret.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fretIdent, __gong__toRawStringLiteral(fret.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", fretIdent, __gong__toRawStringLiteral(fret.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", fretIdent, __gong__toRawStringLiteral(fret.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", fretIdent, __gong__toRawStringLiteral(fret.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", fretIdent, __gong__toRawStringLiteral(fret.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", fretIdent, __gong__toRawStringLiteral(fret.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", fretIdent, fret.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		for _, glass := range __gong__sortStageSetInstances(stageSet.Stage.Glasss, stageSet.Stage.Glass_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			glassIdent := "__models" + glass.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Glass{Name: %s}).Stage(stageSet.Stage)", glassIdent, __gong__toRawStringLiteral(glass.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glassIdent, __gong__toRawStringLiteral(glass.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", glassIdent, __gong__toRawStringLiteral(glass.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", glassIdent, __gong__toRawStringLiteral(glass.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, glissando := range __gong__sortStageSetInstances(stageSet.Stage.Glissandos, stageSet.Stage.Glissando_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			glissandoIdent := "__models" + glissando.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Glissando{Name: %s}).Stage(stageSet.Stage)", glissandoIdent, __gong__toRawStringLiteral(glissando.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", glissandoIdent, __gong__toRawStringLiteral(string(glissando.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", glissandoIdent, glissando.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_type = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Line_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dash_length = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Dash_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space_length = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Space_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, glyph := range __gong__sortStageSetInstances(stageSet.Stage.Glyphs, stageSet.Stage.Glyph_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			glyphIdent := "__models" + glyph.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Glyph{Name: %s}).Stage(stageSet.Stage)", glyphIdent, __gong__toRawStringLiteral(glyph.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glyphIdent, __gong__toRawStringLiteral(glyph.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", glyphIdent, __gong__toRawStringLiteral(glyph.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", glyphIdent, __gong__toRawStringLiteral(glyph.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, grace := range __gong__sortStageSetInstances(stageSet.Stage.Graces, stageSet.Stage.Grace_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			graceIdent := "__models" + grace.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Grace{Name: %s}).Stage(stageSet.Stage)", graceIdent, __gong__toRawStringLiteral(grace.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", graceIdent, __gong__toRawStringLiteral(grace.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Steal_time_previous = %s", graceIdent, __gong__toRawStringLiteral(grace.Steal_time_previous)))
			values.WriteString(fmt.Sprintf("\n\t%s.Steal_time_following = %s", graceIdent, __gong__toRawStringLiteral(grace.Steal_time_following)))
			values.WriteString(fmt.Sprintf("\n\t%s.Make_time = %s", graceIdent, __gong__toRawStringLiteral(grace.Make_time)))
			values.WriteString(fmt.Sprintf("\n\t%s.Slash = %s", graceIdent, __gong__toRawStringLiteral(string(grace.Slash))))
		}
	}
	if stageSet.Stage != nil {
		for _, group_barline := range __gong__sortStageSetInstances(stageSet.Stage.Group_barlines, stageSet.Stage.Group_barline_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			group_barlineIdent := "__models" + group_barline.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Group_barline{Name: %s}).Stage(stageSet.Stage)", group_barlineIdent, __gong__toRawStringLiteral(group_barline.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", group_barlineIdent, __gong__toRawStringLiteral(group_barline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", group_barlineIdent, __gong__toRawStringLiteral(group_barline.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", group_barlineIdent, __gong__toRawStringLiteral(group_barline.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, group_name := range __gong__sortStageSetInstances(stageSet.Stage.Group_names, stageSet.Stage.Group_name_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			group_nameIdent := "__models" + group_name.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Group_name{Name: %s}).Stage(stageSet.Stage)", group_nameIdent, __gong__toRawStringLiteral(group_name.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", group_nameIdent, __gong__toRawStringLiteral(group_name.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", group_nameIdent, __gong__toRawStringLiteral(group_name.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", group_nameIdent, __gong__toRawStringLiteral(group_name.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", group_nameIdent, __gong__toRawStringLiteral(group_name.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", group_nameIdent, __gong__toRawStringLiteral(group_name.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", group_nameIdent, __gong__toRawStringLiteral(group_name.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", group_nameIdent, __gong__toRawStringLiteral(group_name.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", group_nameIdent, __gong__toRawStringLiteral(group_name.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", group_nameIdent, __gong__toRawStringLiteral(group_name.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", group_nameIdent, __gong__toRawStringLiteral(group_name.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", group_nameIdent, __gong__toRawStringLiteral(string(group_name.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", group_nameIdent, __gong__toRawStringLiteral(group_name.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, group_symbol := range __gong__sortStageSetInstances(stageSet.Stage.Group_symbols, stageSet.Stage.Group_symbol_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			group_symbolIdent := "__models" + group_symbol.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Group_symbol{Name: %s}).Stage(stageSet.Stage)", group_symbolIdent, __gong__toRawStringLiteral(group_symbol.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", group_symbolIdent, __gong__toRawStringLiteral(group_symbol.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", group_symbolIdent, __gong__toRawStringLiteral(group_symbol.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", group_symbolIdent, __gong__toRawStringLiteral(group_symbol.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", group_symbolIdent, __gong__toRawStringLiteral(group_symbol.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", group_symbolIdent, __gong__toRawStringLiteral(group_symbol.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", group_symbolIdent, __gong__toRawStringLiteral(group_symbol.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", group_symbolIdent, __gong__toRawStringLiteral(group_symbol.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, grouping := range __gong__sortStageSetInstances(stageSet.Stage.Groupings, stageSet.Stage.Grouping_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groupingIdent := "__models" + grouping.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Grouping{Name: %s}).Stage(stageSet.Stage)", groupingIdent, __gong__toRawStringLiteral(grouping.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Member_of = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Member_of)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Id)))
			for _, elem := range grouping.Feature {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Feature = append(%s.Feature, %s)", groupingIdent, groupingIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, hammer_on_pull_off := range __gong__sortStageSetInstances(stageSet.Stage.Hammer_on_pull_offs, stageSet.Stage.Hammer_on_pull_off_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			hammer_on_pull_offIdent := "__models" + hammer_on_pull_off.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Hammer_on_pull_off{Name: %s}).Stage(stageSet.Stage)", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", hammer_on_pull_offIdent, __gong__toRawStringLiteral(string(hammer_on_pull_off.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", hammer_on_pull_offIdent, hammer_on_pull_off.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, handbell := range __gong__sortStageSetInstances(stageSet.Stage.Handbells, stageSet.Stage.Handbell_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			handbellIdent := "__models" + handbell.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Handbell{Name: %s}).Stage(stageSet.Stage)", handbellIdent, __gong__toRawStringLiteral(handbell.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", handbellIdent, __gong__toRawStringLiteral(handbell.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", handbellIdent, __gong__toRawStringLiteral(handbell.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", handbellIdent, __gong__toRawStringLiteral(handbell.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", handbellIdent, __gong__toRawStringLiteral(handbell.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", handbellIdent, __gong__toRawStringLiteral(handbell.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", handbellIdent, __gong__toRawStringLiteral(handbell.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", handbellIdent, __gong__toRawStringLiteral(handbell.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", handbellIdent, __gong__toRawStringLiteral(handbell.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", handbellIdent, __gong__toRawStringLiteral(handbell.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", handbellIdent, __gong__toRawStringLiteral(handbell.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", handbellIdent, __gong__toRawStringLiteral(handbell.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", handbellIdent, __gong__toRawStringLiteral(handbell.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, harmon_closed := range __gong__sortStageSetInstances(stageSet.Stage.Harmon_closeds, stageSet.Stage.Harmon_closed_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			harmon_closedIdent := "__models" + harmon_closed.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Harmon_closed{Name: %s}).Stage(stageSet.Stage)", harmon_closedIdent, __gong__toRawStringLiteral(harmon_closed.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harmon_closedIdent, __gong__toRawStringLiteral(harmon_closed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", harmon_closedIdent, __gong__toRawStringLiteral(harmon_closed.Location)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", harmon_closedIdent, __gong__toRawStringLiteral(harmon_closed.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, harmon_mute := range __gong__sortStageSetInstances(stageSet.Stage.Harmon_mutes, stageSet.Stage.Harmon_mute_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			harmon_muteIdent := "__models" + harmon_mute.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Harmon_mute{Name: %s}).Stage(stageSet.Stage)", harmon_muteIdent, __gong__toRawStringLiteral(harmon_mute.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harmon_muteIdent, __gong__toRawStringLiteral(harmon_mute.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", harmon_muteIdent, __gong__toRawStringLiteral(harmon_mute.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", harmon_muteIdent, __gong__toRawStringLiteral(harmon_mute.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", harmon_muteIdent, __gong__toRawStringLiteral(harmon_mute.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", harmon_muteIdent, __gong__toRawStringLiteral(harmon_mute.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", harmon_muteIdent, __gong__toRawStringLiteral(harmon_mute.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", harmon_muteIdent, __gong__toRawStringLiteral(harmon_mute.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", harmon_muteIdent, __gong__toRawStringLiteral(harmon_mute.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", harmon_muteIdent, __gong__toRawStringLiteral(harmon_mute.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", harmon_muteIdent, __gong__toRawStringLiteral(harmon_mute.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", harmon_muteIdent, __gong__toRawStringLiteral(harmon_mute.Placement)))
			if harmon_mute.Harmon_closed != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + harmon_mute.Harmon_closed.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmon_closed = %s", harmon_muteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, harmonic := range __gong__sortStageSetInstances(stageSet.Stage.Harmonics, stageSet.Stage.Harmonic_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			harmonicIdent := "__models" + harmonic.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Harmonic{Name: %s}).Stage(stageSet.Stage)", harmonicIdent, __gong__toRawStringLiteral(harmonic.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", harmonicIdent, __gong__toRawStringLiteral(string(harmonic.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Natural = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Natural)))
			values.WriteString(fmt.Sprintf("\n\t%s.Artificial = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Artificial)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base_pitch = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Base_pitch)))
			values.WriteString(fmt.Sprintf("\n\t%s.Touching_pitch = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Touching_pitch)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sounding_pitch = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Sounding_pitch)))
		}
	}
	if stageSet.Stage != nil {
		for _, harmony := range __gong__sortStageSetInstances(stageSet.Stage.Harmonys, stageSet.Stage.Harmony_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			harmonyIdent := "__models" + harmony.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Harmony{Name: %s}).Stage(stageSet.Stage)", harmonyIdent, __gong__toRawStringLiteral(harmony.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_frame = %s", harmonyIdent, __gong__toRawStringLiteral(string(harmony.Print_frame))))
			values.WriteString(fmt.Sprintf("\n\t%s.Arrangement = %s", harmonyIdent, __gong__toRawStringLiteral(string(harmony.Arrangement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", harmonyIdent, __gong__toRawStringLiteral(string(harmony.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.System = %s", harmonyIdent, __gong__toRawStringLiteral(string(harmony.System))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff = %d", harmonyIdent, harmony.Staff))
			if harmony.Root != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + harmony.Root.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Root = %s", harmonyIdent, targetIdent))
			}
			if harmony.Numeral != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + harmony.Numeral.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Numeral = %s", harmonyIdent, targetIdent))
			}
			if harmony.Function != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + harmony.Function.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Function = %s", harmonyIdent, targetIdent))
			}
			if harmony.Kind != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + harmony.Kind.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Kind = %s", harmonyIdent, targetIdent))
			}
			if harmony.Inversion != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + harmony.Inversion.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inversion = %s", harmonyIdent, targetIdent))
			}
			if harmony.Bass != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + harmony.Bass.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bass = %s", harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Degree {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Degree = append(%s.Degree, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			if harmony.Frame != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + harmony.Frame.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frame = %s", harmonyIdent, targetIdent))
			}
			if harmony.Offset != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + harmony.Offset.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", harmonyIdent, targetIdent))
			}
			if harmony.Footnote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + harmony.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", harmonyIdent, targetIdent))
			}
			if harmony.Level != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + harmony.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", harmonyIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, harmony_alter := range __gong__sortStageSetInstances(stageSet.Stage.Harmony_alters, stageSet.Stage.Harmony_alter_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			harmony_alterIdent := "__models" + harmony_alter.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Harmony_alter{Name: %s}).Stage(stageSet.Stage)", harmony_alterIdent, __gong__toRawStringLiteral(harmony_alter.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harmony_alterIdent, __gong__toRawStringLiteral(harmony_alter.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", harmony_alterIdent, __gong__toRawStringLiteral(string(harmony_alter.Location))))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", harmony_alterIdent, __gong__toRawStringLiteral(string(harmony_alter.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", harmony_alterIdent, __gong__toRawStringLiteral(harmony_alter.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", harmony_alterIdent, __gong__toRawStringLiteral(harmony_alter.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", harmony_alterIdent, __gong__toRawStringLiteral(harmony_alter.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", harmony_alterIdent, __gong__toRawStringLiteral(harmony_alter.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", harmony_alterIdent, __gong__toRawStringLiteral(harmony_alter.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", harmony_alterIdent, __gong__toRawStringLiteral(harmony_alter.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", harmony_alterIdent, __gong__toRawStringLiteral(harmony_alter.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", harmony_alterIdent, __gong__toRawStringLiteral(harmony_alter.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", harmony_alterIdent, __gong__toRawStringLiteral(harmony_alter.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", harmony_alterIdent, __gong__toRawStringLiteral(harmony_alter.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, harp_pedals := range __gong__sortStageSetInstances(stageSet.Stage.Harp_pedalss, stageSet.Stage.Harp_pedals_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			harp_pedalsIdent := "__models" + harp_pedals.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Harp_pedals{Name: %s}).Stage(stageSet.Stage)", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Id)))
			for _, elem := range harp_pedals.Pedal_tuning {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pedal_tuning = append(%s.Pedal_tuning, %s)", harp_pedalsIdent, harp_pedalsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, heel_toe := range __gong__sortStageSetInstances(stageSet.Stage.Heel_toes, stageSet.Stage.Heel_toe_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			heel_toeIdent := "__models" + heel_toe.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Heel_toe{Name: %s}).Stage(stageSet.Stage)", heel_toeIdent, __gong__toRawStringLiteral(heel_toe.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", heel_toeIdent, __gong__toRawStringLiteral(heel_toe.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, hole := range __gong__sortStageSetInstances(stageSet.Stage.Holes, stageSet.Stage.Hole_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holeIdent := "__models" + hole.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Hole{Name: %s}).Stage(stageSet.Stage)", holeIdent, __gong__toRawStringLiteral(hole.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holeIdent, __gong__toRawStringLiteral(hole.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", holeIdent, __gong__toRawStringLiteral(hole.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", holeIdent, __gong__toRawStringLiteral(hole.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", holeIdent, __gong__toRawStringLiteral(hole.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", holeIdent, __gong__toRawStringLiteral(hole.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", holeIdent, __gong__toRawStringLiteral(hole.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", holeIdent, __gong__toRawStringLiteral(hole.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", holeIdent, __gong__toRawStringLiteral(hole.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", holeIdent, __gong__toRawStringLiteral(hole.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", holeIdent, __gong__toRawStringLiteral(hole.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", holeIdent, __gong__toRawStringLiteral(hole.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Hole_type = %s", holeIdent, __gong__toRawStringLiteral(hole.Hole_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Hole_shape = %s", holeIdent, __gong__toRawStringLiteral(hole.Hole_shape)))
			if hole.Hole_closed != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + hole.Hole_closed.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Hole_closed = %s", holeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, hole_closed := range __gong__sortStageSetInstances(stageSet.Stage.Hole_closeds, stageSet.Stage.Hole_closed_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			hole_closedIdent := "__models" + hole_closed.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Hole_closed{Name: %s}).Stage(stageSet.Stage)", hole_closedIdent, __gong__toRawStringLiteral(hole_closed.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", hole_closedIdent, __gong__toRawStringLiteral(hole_closed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", hole_closedIdent, __gong__toRawStringLiteral(hole_closed.Location)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", hole_closedIdent, __gong__toRawStringLiteral(hole_closed.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, horizontal_turn := range __gong__sortStageSetInstances(stageSet.Stage.Horizontal_turns, stageSet.Stage.Horizontal_turn_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			horizontal_turnIdent := "__models" + horizontal_turn.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Horizontal_turn{Name: %s}).Stage(stageSet.Stage)", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Slash = %s", horizontal_turnIdent, __gong__toRawStringLiteral(string(horizontal_turn.Slash))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Start_note = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Start_note)))
			values.WriteString(fmt.Sprintf("\n\t%s.Trill_step = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Trill_step)))
			values.WriteString(fmt.Sprintf("\n\t%s.Two_note_turn = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Two_note_turn)))
			values.WriteString(fmt.Sprintf("\n\t%s.Accelerate = %s", horizontal_turnIdent, __gong__toRawStringLiteral(string(horizontal_turn.Accelerate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.Second_beat = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Second_beat)))
			values.WriteString(fmt.Sprintf("\n\t%s.Last_beat = %s", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Last_beat)))
		}
	}
	if stageSet.Stage != nil {
		for _, identification := range __gong__sortStageSetInstances(stageSet.Stage.Identifications, stageSet.Stage.Identification_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			identificationIdent := "__models" + identification.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Identification{Name: %s}).Stage(stageSet.Stage)", identificationIdent, __gong__toRawStringLiteral(identification.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", identificationIdent, __gong__toRawStringLiteral(identification.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Source = %s", identificationIdent, __gong__toRawStringLiteral(identification.Source)))
			for _, elem := range identification.Creator {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Creator = append(%s.Creator, %s)", identificationIdent, identificationIdent, targetIdent))
			}
			for _, elem := range identification.Rights {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rights = append(%s.Rights, %s)", identificationIdent, identificationIdent, targetIdent))
			}
			if identification.Encoding != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + identification.Encoding.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Encoding = %s", identificationIdent, targetIdent))
			}
			for _, elem := range identification.Relation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Relation = append(%s.Relation, %s)", identificationIdent, identificationIdent, targetIdent))
			}
			if identification.Miscellaneous != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + identification.Miscellaneous.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Miscellaneous = %s", identificationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, image := range __gong__sortStageSetInstances(stageSet.Stage.Images, stageSet.Stage.Image_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			imageIdent := "__models" + image.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Image{Name: %s}).Stage(stageSet.Stage)", imageIdent, __gong__toRawStringLiteral(image.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", imageIdent, __gong__toRawStringLiteral(image.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Source = %s", imageIdent, __gong__toRawStringLiteral(image.Source)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", imageIdent, __gong__toRawStringLiteral(image.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %s", imageIdent, __gong__toRawStringLiteral(image.Height)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", imageIdent, __gong__toRawStringLiteral(image.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", imageIdent, __gong__toRawStringLiteral(image.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", imageIdent, __gong__toRawStringLiteral(image.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", imageIdent, __gong__toRawStringLiteral(image.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", imageIdent, __gong__toRawStringLiteral(image.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", imageIdent, __gong__toRawStringLiteral(image.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", imageIdent, __gong__toRawStringLiteral(image.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", imageIdent, __gong__toRawStringLiteral(image.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, instrument := range __gong__sortStageSetInstances(stageSet.Stage.Instruments, stageSet.Stage.Instrument_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			instrumentIdent := "__models" + instrument.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Instrument{Name: %s}).Stage(stageSet.Stage)", instrumentIdent, __gong__toRawStringLiteral(instrument.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrumentIdent, __gong__toRawStringLiteral(instrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", instrumentIdent, __gong__toRawStringLiteral(instrument.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, instrument_change := range __gong__sortStageSetInstances(stageSet.Stage.Instrument_changes, stageSet.Stage.Instrument_change_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			instrument_changeIdent := "__models" + instrument_change.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Instrument_change{Name: %s}).Stage(stageSet.Stage)", instrument_changeIdent, __gong__toRawStringLiteral(instrument_change.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrument_changeIdent, __gong__toRawStringLiteral(instrument_change.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", instrument_changeIdent, __gong__toRawStringLiteral(instrument_change.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Instrument_sound = %s", instrument_changeIdent, __gong__toRawStringLiteral(instrument_change.Instrument_sound)))
			values.WriteString(fmt.Sprintf("\n\t%s.Solo = %s", instrument_changeIdent, __gong__toRawStringLiteral(instrument_change.Solo)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ensemble = %s", instrument_changeIdent, __gong__toRawStringLiteral(instrument_change.Ensemble)))
			if instrument_change.Virtual_instrument != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + instrument_change.Virtual_instrument.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Virtual_instrument = %s", instrument_changeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, instrument_link := range __gong__sortStageSetInstances(stageSet.Stage.Instrument_links, stageSet.Stage.Instrument_link_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			instrument_linkIdent := "__models" + instrument_link.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Instrument_link{Name: %s}).Stage(stageSet.Stage)", instrument_linkIdent, __gong__toRawStringLiteral(instrument_link.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrument_linkIdent, __gong__toRawStringLiteral(instrument_link.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", instrument_linkIdent, __gong__toRawStringLiteral(instrument_link.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, interchangeable := range __gong__sortStageSetInstances(stageSet.Stage.Interchangeables, stageSet.Stage.Interchangeable_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			interchangeableIdent := "__models" + interchangeable.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Interchangeable{Name: %s}).Stage(stageSet.Stage)", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Symbol = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Symbol)))
			values.WriteString(fmt.Sprintf("\n\t%s.Separator = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Separator)))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_relation = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Time_relation)))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.Beat_type = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Beat_type)))
		}
	}
	if stageSet.Stage != nil {
		for _, inversion := range __gong__sortStageSetInstances(stageSet.Stage.Inversions, stageSet.Stage.Inversion_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			inversionIdent := "__models" + inversion.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Inversion{Name: %s}).Stage(stageSet.Stage)", inversionIdent, __gong__toRawStringLiteral(inversion.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", inversionIdent, inversion.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		for _, key := range __gong__sortStageSetInstances(stageSet.Stage.Keys, stageSet.Stage.Key_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			keyIdent := "__models" + key.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Key{Name: %s}).Stage(stageSet.Stage)", keyIdent, __gong__toRawStringLiteral(key.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", keyIdent, __gong__toRawStringLiteral(key.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", keyIdent, key.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", keyIdent, __gong__toRawStringLiteral(key.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", keyIdent, __gong__toRawStringLiteral(key.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", keyIdent, __gong__toRawStringLiteral(key.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", keyIdent, __gong__toRawStringLiteral(key.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", keyIdent, __gong__toRawStringLiteral(key.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", keyIdent, __gong__toRawStringLiteral(key.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", keyIdent, __gong__toRawStringLiteral(key.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", keyIdent, __gong__toRawStringLiteral(key.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", keyIdent, __gong__toRawStringLiteral(key.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", keyIdent, __gong__toRawStringLiteral(string(key.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", keyIdent, __gong__toRawStringLiteral(key.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Fifths = %d", keyIdent, key.Fifths))
			values.WriteString(fmt.Sprintf("\n\t%s.Mode = %s", keyIdent, __gong__toRawStringLiteral(key.Mode)))
			values.WriteString(fmt.Sprintf("\n\t%s.Key_step = %s", keyIdent, __gong__toRawStringLiteral(string(key.Key_step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Key_alter = %s", keyIdent, __gong__toRawStringLiteral(key.Key_alter)))
			if key.Cancel != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + key.Cancel.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Cancel = %s", keyIdent, targetIdent))
			}
			if key.Key_accidental != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + key.Key_accidental.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Key_accidental = %s", keyIdent, targetIdent))
			}
			for _, elem := range key.Key_octave {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Key_octave = append(%s.Key_octave, %s)", keyIdent, keyIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, key_accidental := range __gong__sortStageSetInstances(stageSet.Stage.Key_accidentals, stageSet.Stage.Key_accidental_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			key_accidentalIdent := "__models" + key_accidental.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Key_accidental{Name: %s}).Stage(stageSet.Stage)", key_accidentalIdent, __gong__toRawStringLiteral(key_accidental.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", key_accidentalIdent, __gong__toRawStringLiteral(key_accidental.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", key_accidentalIdent, __gong__toRawStringLiteral(key_accidental.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", key_accidentalIdent, __gong__toRawStringLiteral(string(key_accidental.EnclosedText))))
		}
	}
	if stageSet.Stage != nil {
		for _, key_octave := range __gong__sortStageSetInstances(stageSet.Stage.Key_octaves, stageSet.Stage.Key_octave_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			key_octaveIdent := "__models" + key_octave.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Key_octave{Name: %s}).Stage(stageSet.Stage)", key_octaveIdent, __gong__toRawStringLiteral(key_octave.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", key_octaveIdent, __gong__toRawStringLiteral(key_octave.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", key_octaveIdent, key_octave.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Cancel = %s", key_octaveIdent, __gong__toRawStringLiteral(string(key_octave.Cancel))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", key_octaveIdent, key_octave.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		for _, kind := range __gong__sortStageSetInstances(stageSet.Stage.Kinds, stageSet.Stage.Kind_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			kindIdent := "__models" + kind.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Kind{Name: %s}).Stage(stageSet.Stage)", kindIdent, __gong__toRawStringLiteral(kind.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", kindIdent, __gong__toRawStringLiteral(kind.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Use_symbols = %s", kindIdent, __gong__toRawStringLiteral(string(kind.Use_symbols))))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", kindIdent, __gong__toRawStringLiteral(kind.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Stack_degrees = %s", kindIdent, __gong__toRawStringLiteral(string(kind.Stack_degrees))))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses_degrees = %s", kindIdent, __gong__toRawStringLiteral(string(kind.Parentheses_degrees))))
			values.WriteString(fmt.Sprintf("\n\t%s.Bracket_degrees = %s", kindIdent, __gong__toRawStringLiteral(string(kind.Bracket_degrees))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", kindIdent, __gong__toRawStringLiteral(kind.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", kindIdent, __gong__toRawStringLiteral(kind.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", kindIdent, __gong__toRawStringLiteral(kind.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", kindIdent, __gong__toRawStringLiteral(kind.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", kindIdent, __gong__toRawStringLiteral(kind.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", kindIdent, __gong__toRawStringLiteral(kind.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", kindIdent, __gong__toRawStringLiteral(kind.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", kindIdent, __gong__toRawStringLiteral(kind.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", kindIdent, __gong__toRawStringLiteral(kind.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", kindIdent, __gong__toRawStringLiteral(kind.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", kindIdent, __gong__toRawStringLiteral(kind.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", kindIdent, __gong__toRawStringLiteral(kind.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, level := range __gong__sortStageSetInstances(stageSet.Stage.Levels, stageSet.Stage.Level_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			levelIdent := "__models" + level.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Level{Name: %s}).Stage(stageSet.Stage)", levelIdent, __gong__toRawStringLiteral(level.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", levelIdent, __gong__toRawStringLiteral(level.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Reference = %s", levelIdent, __gong__toRawStringLiteral(string(level.Reference))))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", levelIdent, __gong__toRawStringLiteral(string(level.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", levelIdent, __gong__toRawStringLiteral(string(level.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.Bracket = %s", levelIdent, __gong__toRawStringLiteral(string(level.Bracket))))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %s", levelIdent, __gong__toRawStringLiteral(string(level.Size))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", levelIdent, __gong__toRawStringLiteral(level.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, line_detail := range __gong__sortStageSetInstances(stageSet.Stage.Line_details, stageSet.Stage.Line_detail_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			line_detailIdent := "__models" + line_detail.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Line_detail{Name: %s}).Stage(stageSet.Stage)", line_detailIdent, __gong__toRawStringLiteral(line_detail.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", line_detailIdent, __gong__toRawStringLiteral(line_detail.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", line_detailIdent, line_detail.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", line_detailIdent, __gong__toRawStringLiteral(line_detail.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", line_detailIdent, __gong__toRawStringLiteral(line_detail.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_type = %s", line_detailIdent, __gong__toRawStringLiteral(line_detail.Line_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", line_detailIdent, __gong__toRawStringLiteral(string(line_detail.Print_object))))
		}
	}
	if stageSet.Stage != nil {
		for _, line_width := range __gong__sortStageSetInstances(stageSet.Stage.Line_widths, stageSet.Stage.Line_width_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			line_widthIdent := "__models" + line_width.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Line_width{Name: %s}).Stage(stageSet.Stage)", line_widthIdent, __gong__toRawStringLiteral(line_width.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", line_widthIdent, __gong__toRawStringLiteral(line_width.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", line_widthIdent, __gong__toRawStringLiteral(line_width.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", line_widthIdent, __gong__toRawStringLiteral(line_width.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, link := range __gong__sortStageSetInstances(stageSet.Stage.Links, stageSet.Stage.Link_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linkIdent := "__models" + link.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Link{Name: %s}).Stage(stageSet.Stage)", linkIdent, __gong__toRawStringLiteral(link.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linkIdent, __gong__toRawStringLiteral(link.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", linkIdent, __gong__toRawStringLiteral(link.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", linkIdent, __gong__toRawStringLiteral(link.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", linkIdent, __gong__toRawStringLiteral(link.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", linkIdent, __gong__toRawStringLiteral(link.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", linkIdent, __gong__toRawStringLiteral(link.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", linkIdent, __gong__toRawStringLiteral(link.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", linkIdent, __gong__toRawStringLiteral(link.Actuate)))
			values.WriteString(fmt.Sprintf("\n\t%s.Element = %s", linkIdent, __gong__toRawStringLiteral(link.Element)))
			values.WriteString(fmt.Sprintf("\n\t%s.Position = %d", linkIdent, link.Position))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", linkIdent, __gong__toRawStringLiteral(link.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", linkIdent, __gong__toRawStringLiteral(link.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", linkIdent, __gong__toRawStringLiteral(link.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", linkIdent, __gong__toRawStringLiteral(link.Relative_y)))
		}
	}
	if stageSet.Stage != nil {
		for _, listen := range __gong__sortStageSetInstances(stageSet.Stage.Listens, stageSet.Stage.Listen_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			listenIdent := "__models" + listen.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Listen{Name: %s}).Stage(stageSet.Stage)", listenIdent, __gong__toRawStringLiteral(listen.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", listenIdent, __gong__toRawStringLiteral(listen.Name)))
			for _, elem := range listen.Assess {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Assess = append(%s.Assess, %s)", listenIdent, listenIdent, targetIdent))
			}
			for _, elem := range listen.Wait {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wait = append(%s.Wait, %s)", listenIdent, listenIdent, targetIdent))
			}
			for _, elem := range listen.Other_listen {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_listen = append(%s.Other_listen, %s)", listenIdent, listenIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, listening := range __gong__sortStageSetInstances(stageSet.Stage.Listenings, stageSet.Stage.Listening_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			listeningIdent := "__models" + listening.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Listening{Name: %s}).Stage(stageSet.Stage)", listeningIdent, __gong__toRawStringLiteral(listening.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", listeningIdent, __gong__toRawStringLiteral(listening.Name)))
			for _, elem := range listening.Sync {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sync = append(%s.Sync, %s)", listeningIdent, listeningIdent, targetIdent))
			}
			for _, elem := range listening.Other_listening {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_listening = append(%s.Other_listening, %s)", listeningIdent, listeningIdent, targetIdent))
			}
			if listening.Offset != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + listening.Offset.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", listeningIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, lyric := range __gong__sortStageSetInstances(stageSet.Stage.Lyrics, stageSet.Stage.Lyric_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lyricIdent := "__models" + lyric.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Lyric{Name: %s}).Stage(stageSet.Stage)", lyricIdent, __gong__toRawStringLiteral(lyric.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", lyricIdent, __gong__toRawStringLiteral(lyric.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_only = %s", lyricIdent, __gong__toRawStringLiteral(string(lyric.Time_only))))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", lyricIdent, __gong__toRawStringLiteral(string(lyric.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", lyricIdent, __gong__toRawStringLiteral(string(lyric.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Syllabic = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Syllabic)))
			values.WriteString(fmt.Sprintf("\n\t%s.Laughing = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Laughing)))
			values.WriteString(fmt.Sprintf("\n\t%s.Humming = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Humming)))
			values.WriteString(fmt.Sprintf("\n\t%s.End_line = %s", lyricIdent, __gong__toRawStringLiteral(lyric.End_line)))
			values.WriteString(fmt.Sprintf("\n\t%s.End_paragraph = %s", lyricIdent, __gong__toRawStringLiteral(lyric.End_paragraph)))
			for _, elem := range lyric.Elision {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elision = append(%s.Elision, %s)", lyricIdent, lyricIdent, targetIdent))
			}
			for _, elem := range lyric.Text {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Text = append(%s.Text, %s)", lyricIdent, lyricIdent, targetIdent))
			}
			if lyric.Extend != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + lyric.Extend.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extend = %s", lyricIdent, targetIdent))
			}
			if lyric.Footnote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + lyric.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", lyricIdent, targetIdent))
			}
			if lyric.Level != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + lyric.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", lyricIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, lyric_font := range __gong__sortStageSetInstances(stageSet.Stage.Lyric_fonts, stageSet.Stage.Lyric_font_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lyric_fontIdent := "__models" + lyric_font.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Lyric_font{Name: %s}).Stage(stageSet.Stage)", lyric_fontIdent, __gong__toRawStringLiteral(lyric_font.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lyric_fontIdent, __gong__toRawStringLiteral(lyric_font.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", lyric_fontIdent, __gong__toRawStringLiteral(lyric_font.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", lyric_fontIdent, __gong__toRawStringLiteral(lyric_font.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", lyric_fontIdent, __gong__toRawStringLiteral(lyric_font.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", lyric_fontIdent, __gong__toRawStringLiteral(lyric_font.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", lyric_fontIdent, __gong__toRawStringLiteral(lyric_font.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", lyric_fontIdent, __gong__toRawStringLiteral(lyric_font.Font_weight)))
		}
	}
	if stageSet.Stage != nil {
		for _, lyric_language := range __gong__sortStageSetInstances(stageSet.Stage.Lyric_languages, stageSet.Stage.Lyric_language_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lyric_languageIdent := "__models" + lyric_language.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Lyric_language{Name: %s}).Stage(stageSet.Stage)", lyric_languageIdent, __gong__toRawStringLiteral(lyric_language.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lyric_languageIdent, __gong__toRawStringLiteral(lyric_language.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", lyric_languageIdent, __gong__toRawStringLiteral(lyric_language.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", lyric_languageIdent, __gong__toRawStringLiteral(lyric_language.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", lyric_languageIdent, __gong__toRawStringLiteral(lyric_language.Lang)))
		}
	}
	if stageSet.Stage != nil {
		for _, measure_layout := range __gong__sortStageSetInstances(stageSet.Stage.Measure_layouts, stageSet.Stage.Measure_layout_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			measure_layoutIdent := "__models" + measure_layout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Measure_layout{Name: %s}).Stage(stageSet.Stage)", measure_layoutIdent, __gong__toRawStringLiteral(measure_layout.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measure_layoutIdent, __gong__toRawStringLiteral(measure_layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Measure_distance = %s", measure_layoutIdent, __gong__toRawStringLiteral(measure_layout.Measure_distance)))
		}
	}
	if stageSet.Stage != nil {
		for _, measure_numbering := range __gong__sortStageSetInstances(stageSet.Stage.Measure_numberings, stageSet.Stage.Measure_numbering_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			measure_numberingIdent := "__models" + measure_numbering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Measure_numbering{Name: %s}).Stage(stageSet.Stage)", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.System = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.System)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff = %d", measure_numberingIdent, measure_numbering.Staff))
			values.WriteString(fmt.Sprintf("\n\t%s.Multiple_rest_always = %s", measure_numberingIdent, __gong__toRawStringLiteral(string(measure_numbering.Multiple_rest_always))))
			values.WriteString(fmt.Sprintf("\n\t%s.Multiple_rest_range = %s", measure_numberingIdent, __gong__toRawStringLiteral(string(measure_numbering.Multiple_rest_range))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, measure_repeat := range __gong__sortStageSetInstances(stageSet.Stage.Measure_repeats, stageSet.Stage.Measure_repeat_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			measure_repeatIdent := "__models" + measure_repeat.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Measure_repeat{Name: %s}).Stage(stageSet.Stage)", measure_repeatIdent, __gong__toRawStringLiteral(measure_repeat.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measure_repeatIdent, __gong__toRawStringLiteral(measure_repeat.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", measure_repeatIdent, __gong__toRawStringLiteral(string(measure_repeat.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Slashes = %d", measure_repeatIdent, measure_repeat.Slashes))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", measure_repeatIdent, __gong__toRawStringLiteral(measure_repeat.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, measure_style := range __gong__sortStageSetInstances(stageSet.Stage.Measure_styles, stageSet.Stage.Measure_style_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			measure_styleIdent := "__models" + measure_style.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Measure_style{Name: %s}).Stage(stageSet.Stage)", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", measure_styleIdent, measure_style.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Id)))
			if measure_style.Multiple_rest != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + measure_style.Multiple_rest.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Multiple_rest = %s", measure_styleIdent, targetIdent))
			}
			if measure_style.Measure_repeat != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + measure_style.Measure_repeat.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure_repeat = %s", measure_styleIdent, targetIdent))
			}
			if measure_style.Beat_repeat != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + measure_style.Beat_repeat.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Beat_repeat = %s", measure_styleIdent, targetIdent))
			}
			if measure_style.Slash != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + measure_style.Slash.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slash = %s", measure_styleIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, membrane := range __gong__sortStageSetInstances(stageSet.Stage.Membranes, stageSet.Stage.Membrane_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			membraneIdent := "__models" + membrane.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Membrane{Name: %s}).Stage(stageSet.Stage)", membraneIdent, __gong__toRawStringLiteral(membrane.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", membraneIdent, __gong__toRawStringLiteral(membrane.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", membraneIdent, __gong__toRawStringLiteral(membrane.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", membraneIdent, __gong__toRawStringLiteral(membrane.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, metal := range __gong__sortStageSetInstances(stageSet.Stage.Metals, stageSet.Stage.Metal_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			metalIdent := "__models" + metal.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Metal{Name: %s}).Stage(stageSet.Stage)", metalIdent, __gong__toRawStringLiteral(metal.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metalIdent, __gong__toRawStringLiteral(metal.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", metalIdent, __gong__toRawStringLiteral(metal.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", metalIdent, __gong__toRawStringLiteral(metal.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, metronome := range __gong__sortStageSetInstances(stageSet.Stage.Metronomes, stageSet.Stage.Metronome_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			metronomeIdent := "__models" + metronome.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Metronome{Name: %s}).Stage(stageSet.Stage)", metronomeIdent, __gong__toRawStringLiteral(metronome.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", metronomeIdent, __gong__toRawStringLiteral(string(metronome.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", metronomeIdent, __gong__toRawStringLiteral(string(metronome.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", metronomeIdent, __gong__toRawStringLiteral(string(metronome.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Beat_unit = %s", metronomeIdent, __gong__toRawStringLiteral(string(metronome.Beat_unit))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beat_unit_dot = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Beat_unit_dot)))
			values.WriteString(fmt.Sprintf("\n\t%s.Metronome_arrows = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Metronome_arrows)))
			values.WriteString(fmt.Sprintf("\n\t%s.Metronome_relation = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Metronome_relation)))
			if metronome.Per_minute != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + metronome.Per_minute.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Per_minute = %s", metronomeIdent, targetIdent))
			}
			for _, elem := range metronome.Beat_unit_tied {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Beat_unit_tied = append(%s.Beat_unit_tied, %s)", metronomeIdent, metronomeIdent, targetIdent))
			}
			for _, elem := range metronome.Metronome_note {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metronome_note = append(%s.Metronome_note, %s)", metronomeIdent, metronomeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, metronome_beam := range __gong__sortStageSetInstances(stageSet.Stage.Metronome_beams, stageSet.Stage.Metronome_beam_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			metronome_beamIdent := "__models" + metronome_beam.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Metronome_beam{Name: %s}).Stage(stageSet.Stage)", metronome_beamIdent, __gong__toRawStringLiteral(metronome_beam.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronome_beamIdent, __gong__toRawStringLiteral(metronome_beam.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", metronome_beamIdent, metronome_beam.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", metronome_beamIdent, __gong__toRawStringLiteral(string(metronome_beam.EnclosedText))))
		}
	}
	if stageSet.Stage != nil {
		for _, metronome_note := range __gong__sortStageSetInstances(stageSet.Stage.Metronome_notes, stageSet.Stage.Metronome_note_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			metronome_noteIdent := "__models" + metronome_note.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Metronome_note{Name: %s}).Stage(stageSet.Stage)", metronome_noteIdent, __gong__toRawStringLiteral(metronome_note.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronome_noteIdent, __gong__toRawStringLiteral(metronome_note.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Metronome_type = %s", metronome_noteIdent, __gong__toRawStringLiteral(metronome_note.Metronome_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Metronome_dot = %s", metronome_noteIdent, __gong__toRawStringLiteral(metronome_note.Metronome_dot)))
			for _, elem := range metronome_note.Metronome_beam {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metronome_beam = append(%s.Metronome_beam, %s)", metronome_noteIdent, metronome_noteIdent, targetIdent))
			}
			if metronome_note.Metronome_tied != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + metronome_note.Metronome_tied.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metronome_tied = %s", metronome_noteIdent, targetIdent))
			}
			if metronome_note.Metronome_tuplet != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + metronome_note.Metronome_tuplet.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metronome_tuplet = %s", metronome_noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, metronome_tied := range __gong__sortStageSetInstances(stageSet.Stage.Metronome_tieds, stageSet.Stage.Metronome_tied_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			metronome_tiedIdent := "__models" + metronome_tied.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Metronome_tied{Name: %s}).Stage(stageSet.Stage)", metronome_tiedIdent, __gong__toRawStringLiteral(metronome_tied.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronome_tiedIdent, __gong__toRawStringLiteral(metronome_tied.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", metronome_tiedIdent, __gong__toRawStringLiteral(string(metronome_tied.Type))))
		}
	}
	if stageSet.Stage != nil {
		for _, metronome_tuplet := range __gong__sortStageSetInstances(stageSet.Stage.Metronome_tuplets, stageSet.Stage.Metronome_tuplet_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			metronome_tupletIdent := "__models" + metronome_tuplet.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Metronome_tuplet{Name: %s}).Stage(stageSet.Stage)", metronome_tupletIdent, __gong__toRawStringLiteral(metronome_tuplet.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronome_tupletIdent, __gong__toRawStringLiteral(metronome_tuplet.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, midi_device := range __gong__sortStageSetInstances(stageSet.Stage.Midi_devices, stageSet.Stage.Midi_device_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			midi_deviceIdent := "__models" + midi_device.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Midi_device{Name: %s}).Stage(stageSet.Stage)", midi_deviceIdent, __gong__toRawStringLiteral(midi_device.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", midi_deviceIdent, __gong__toRawStringLiteral(midi_device.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Port = %d", midi_deviceIdent, midi_device.Port))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", midi_deviceIdent, __gong__toRawStringLiteral(midi_device.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", midi_deviceIdent, __gong__toRawStringLiteral(midi_device.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, midi_instrument := range __gong__sortStageSetInstances(stageSet.Stage.Midi_instruments, stageSet.Stage.Midi_instrument_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			midi_instrumentIdent := "__models" + midi_instrument.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Midi_instrument{Name: %s}).Stage(stageSet.Stage)", midi_instrumentIdent, __gong__toRawStringLiteral(midi_instrument.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", midi_instrumentIdent, __gong__toRawStringLiteral(midi_instrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", midi_instrumentIdent, __gong__toRawStringLiteral(midi_instrument.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Midi_channel = %d", midi_instrumentIdent, midi_instrument.Midi_channel))
			values.WriteString(fmt.Sprintf("\n\t%s.Midi_name = %s", midi_instrumentIdent, __gong__toRawStringLiteral(midi_instrument.Midi_name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Midi_bank = %d", midi_instrumentIdent, midi_instrument.Midi_bank))
			values.WriteString(fmt.Sprintf("\n\t%s.Midi_program = %d", midi_instrumentIdent, midi_instrument.Midi_program))
			values.WriteString(fmt.Sprintf("\n\t%s.Midi_unpitched = %d", midi_instrumentIdent, midi_instrument.Midi_unpitched))
			values.WriteString(fmt.Sprintf("\n\t%s.Volume = %s", midi_instrumentIdent, __gong__toRawStringLiteral(midi_instrument.Volume)))
			values.WriteString(fmt.Sprintf("\n\t%s.Pan = %s", midi_instrumentIdent, __gong__toRawStringLiteral(midi_instrument.Pan)))
			values.WriteString(fmt.Sprintf("\n\t%s.Elevation = %s", midi_instrumentIdent, __gong__toRawStringLiteral(midi_instrument.Elevation)))
		}
	}
	if stageSet.Stage != nil {
		for _, miscellaneous := range __gong__sortStageSetInstances(stageSet.Stage.Miscellaneouss, stageSet.Stage.Miscellaneous_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			miscellaneousIdent := "__models" + miscellaneous.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Miscellaneous{Name: %s}).Stage(stageSet.Stage)", miscellaneousIdent, __gong__toRawStringLiteral(miscellaneous.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", miscellaneousIdent, __gong__toRawStringLiteral(miscellaneous.Name)))
			for _, elem := range miscellaneous.Miscellaneous_field {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Miscellaneous_field = append(%s.Miscellaneous_field, %s)", miscellaneousIdent, miscellaneousIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, miscellaneous_field := range __gong__sortStageSetInstances(stageSet.Stage.Miscellaneous_fields, stageSet.Stage.Miscellaneous_field_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			miscellaneous_fieldIdent := "__models" + miscellaneous_field.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Miscellaneous_field{Name: %s}).Stage(stageSet.Stage)", miscellaneous_fieldIdent, __gong__toRawStringLiteral(miscellaneous_field.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", miscellaneous_fieldIdent, __gong__toRawStringLiteral(miscellaneous_field.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", miscellaneous_fieldIdent, __gong__toRawStringLiteral(miscellaneous_field.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", miscellaneous_fieldIdent, __gong__toRawStringLiteral(miscellaneous_field.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, mordent := range __gong__sortStageSetInstances(stageSet.Stage.Mordents, stageSet.Stage.Mordent_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			mordentIdent := "__models" + mordent.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Mordent{Name: %s}).Stage(stageSet.Stage)", mordentIdent, __gong__toRawStringLiteral(mordent.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", mordentIdent, __gong__toRawStringLiteral(mordent.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, multiple_rest := range __gong__sortStageSetInstances(stageSet.Stage.Multiple_rests, stageSet.Stage.Multiple_rest_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			multiple_restIdent := "__models" + multiple_rest.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Multiple_rest{Name: %s}).Stage(stageSet.Stage)", multiple_restIdent, __gong__toRawStringLiteral(multiple_rest.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", multiple_restIdent, __gong__toRawStringLiteral(multiple_rest.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Use_symbols = %s", multiple_restIdent, __gong__toRawStringLiteral(string(multiple_rest.Use_symbols))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", multiple_restIdent, multiple_rest.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		for _, name_display := range __gong__sortStageSetInstances(stageSet.Stage.Name_displays, stageSet.Stage.Name_display_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			name_displayIdent := "__models" + name_display.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Name_display{Name: %s}).Stage(stageSet.Stage)", name_displayIdent, __gong__toRawStringLiteral(name_display.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", name_displayIdent, __gong__toRawStringLiteral(name_display.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", name_displayIdent, __gong__toRawStringLiteral(string(name_display.Print_object))))
			for _, elem := range name_display.Display_text {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Display_text = append(%s.Display_text, %s)", name_displayIdent, name_displayIdent, targetIdent))
			}
			for _, elem := range name_display.Accidental_text {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accidental_text = append(%s.Accidental_text, %s)", name_displayIdent, name_displayIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, non_arpeggiate := range __gong__sortStageSetInstances(stageSet.Stage.Non_arpeggiates, stageSet.Stage.Non_arpeggiate_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			non_arpeggiateIdent := "__models" + non_arpeggiate.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Non_arpeggiate{Name: %s}).Stage(stageSet.Stage)", non_arpeggiateIdent, __gong__toRawStringLiteral(non_arpeggiate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", non_arpeggiateIdent, __gong__toRawStringLiteral(non_arpeggiate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", non_arpeggiateIdent, __gong__toRawStringLiteral(non_arpeggiate.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", non_arpeggiateIdent, non_arpeggiate.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", non_arpeggiateIdent, __gong__toRawStringLiteral(non_arpeggiate.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", non_arpeggiateIdent, __gong__toRawStringLiteral(non_arpeggiate.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", non_arpeggiateIdent, __gong__toRawStringLiteral(non_arpeggiate.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", non_arpeggiateIdent, __gong__toRawStringLiteral(non_arpeggiate.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", non_arpeggiateIdent, __gong__toRawStringLiteral(non_arpeggiate.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", non_arpeggiateIdent, __gong__toRawStringLiteral(non_arpeggiate.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", non_arpeggiateIdent, __gong__toRawStringLiteral(non_arpeggiate.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, notations := range __gong__sortStageSetInstances(stageSet.Stage.Notationss, stageSet.Stage.Notations_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notationsIdent := "__models" + notations.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Notations{Name: %s}).Stage(stageSet.Stage)", notationsIdent, __gong__toRawStringLiteral(notations.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notationsIdent, __gong__toRawStringLiteral(notations.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", notationsIdent, __gong__toRawStringLiteral(string(notations.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", notationsIdent, __gong__toRawStringLiteral(notations.Id)))
			if notations.Footnote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notations.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", notationsIdent, targetIdent))
			}
			if notations.Level != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + notations.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", notationsIdent, targetIdent))
			}
			for _, elem := range notations.Tied {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tied = append(%s.Tied, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Slur {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slur = append(%s.Slur, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Tuplet {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet = append(%s.Tuplet, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Glissando {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Glissando = append(%s.Glissando, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Slide {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slide = append(%s.Slide, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Ornaments {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ornaments = append(%s.Ornaments, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Technical {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Technical = append(%s.Technical, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Articulations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Articulations = append(%s.Articulations, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Dynamics {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dynamics = append(%s.Dynamics, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Fermata {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fermata = append(%s.Fermata, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Arpeggiate {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arpeggiate = append(%s.Arpeggiate, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Non_arpeggiate {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Non_arpeggiate = append(%s.Non_arpeggiate, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Accidental_mark {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accidental_mark = append(%s.Accidental_mark, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Other_notation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_notation = append(%s.Other_notation, %s)", notationsIdent, notationsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, note := range __gong__sortStageSetInstances(stageSet.Stage.Notes, stageSet.Stage.Note_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteIdent := "__models" + note.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Note{Name: %s}).Stage(stageSet.Stage)", noteIdent, __gong__toRawStringLiteral(note.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteIdent, __gong__toRawStringLiteral(note.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_leger = %s", noteIdent, __gong__toRawStringLiteral(string(note.Print_leger))))
			values.WriteString(fmt.Sprintf("\n\t%s.Dynamics = %s", noteIdent, __gong__toRawStringLiteral(note.Dynamics)))
			values.WriteString(fmt.Sprintf("\n\t%s.End_dynamics = %s", noteIdent, __gong__toRawStringLiteral(note.End_dynamics)))
			values.WriteString(fmt.Sprintf("\n\t%s.Attack = %s", noteIdent, __gong__toRawStringLiteral(note.Attack)))
			values.WriteString(fmt.Sprintf("\n\t%s.Release = %s", noteIdent, __gong__toRawStringLiteral(note.Release)))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_only = %s", noteIdent, __gong__toRawStringLiteral(string(note.Time_only))))
			values.WriteString(fmt.Sprintf("\n\t%s.Pizzicato = %s", noteIdent, __gong__toRawStringLiteral(string(note.Pizzicato))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", noteIdent, __gong__toRawStringLiteral(note.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", noteIdent, __gong__toRawStringLiteral(note.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", noteIdent, __gong__toRawStringLiteral(note.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", noteIdent, __gong__toRawStringLiteral(note.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", noteIdent, __gong__toRawStringLiteral(note.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", noteIdent, __gong__toRawStringLiteral(note.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", noteIdent, __gong__toRawStringLiteral(note.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", noteIdent, __gong__toRawStringLiteral(note.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", noteIdent, __gong__toRawStringLiteral(note.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_dot = %s", noteIdent, __gong__toRawStringLiteral(string(note.Print_dot))))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_lyric = %s", noteIdent, __gong__toRawStringLiteral(string(note.Print_lyric))))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", noteIdent, __gong__toRawStringLiteral(string(note.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_spacing = %s", noteIdent, __gong__toRawStringLiteral(string(note.Print_spacing))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", noteIdent, __gong__toRawStringLiteral(note.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Chord = %s", noteIdent, __gong__toRawStringLiteral(note.Chord)))
			values.WriteString(fmt.Sprintf("\n\t%s.Cue = %s", noteIdent, __gong__toRawStringLiteral(note.Cue)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = %s", noteIdent, __gong__toRawStringLiteral(note.Duration)))
			values.WriteString(fmt.Sprintf("\n\t%s.Voice = %s", noteIdent, __gong__toRawStringLiteral(note.Voice)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff = %d", noteIdent, note.Staff))
			if note.Grace != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Grace.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Grace = %s", noteIdent, targetIdent))
			}
			if note.Pitch != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Pitch.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pitch = %s", noteIdent, targetIdent))
			}
			if note.Unpitched != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Unpitched.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Unpitched = %s", noteIdent, targetIdent))
			}
			if note.Rest != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Rest.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rest = %s", noteIdent, targetIdent))
			}
			if note.Tie != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Tie.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tie = %s", noteIdent, targetIdent))
			}
			for _, elem := range note.Instrument {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Instrument = append(%s.Instrument, %s)", noteIdent, noteIdent, targetIdent))
			}
			if note.Footnote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", noteIdent, targetIdent))
			}
			if note.Level != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", noteIdent, targetIdent))
			}
			if note.Type != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Type.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type = %s", noteIdent, targetIdent))
			}
			for _, elem := range note.Dot {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dot = append(%s.Dot, %s)", noteIdent, noteIdent, targetIdent))
			}
			if note.Accidental != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Accidental.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accidental = %s", noteIdent, targetIdent))
			}
			if note.Time_modification != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Time_modification.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Time_modification = %s", noteIdent, targetIdent))
			}
			if note.Stem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Stem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stem = %s", noteIdent, targetIdent))
			}
			if note.Notehead != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Notehead.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notehead = %s", noteIdent, targetIdent))
			}
			if note.Notehead_text != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Notehead_text.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notehead_text = %s", noteIdent, targetIdent))
			}
			if note.Beam != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Beam.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Beam = %s", noteIdent, targetIdent))
			}
			for _, elem := range note.Notations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notations = append(%s.Notations, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Lyric {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lyric = append(%s.Lyric, %s)", noteIdent, noteIdent, targetIdent))
			}
			if note.Play != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Play.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Play = %s", noteIdent, targetIdent))
			}
			if note.Listen != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + note.Listen.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listen = %s", noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, note_size := range __gong__sortStageSetInstances(stageSet.Stage.Note_sizes, stageSet.Stage.Note_size_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			note_sizeIdent := "__models" + note_size.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Note_size{Name: %s}).Stage(stageSet.Stage)", note_sizeIdent, __gong__toRawStringLiteral(note_size.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", note_sizeIdent, __gong__toRawStringLiteral(note_size.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", note_sizeIdent, __gong__toRawStringLiteral(note_size.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", note_sizeIdent, __gong__toRawStringLiteral(note_size.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, note_type := range __gong__sortStageSetInstances(stageSet.Stage.Note_types, stageSet.Stage.Note_type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			note_typeIdent := "__models" + note_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Note_type{Name: %s}).Stage(stageSet.Stage)", note_typeIdent, __gong__toRawStringLiteral(note_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", note_typeIdent, __gong__toRawStringLiteral(note_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %s", note_typeIdent, __gong__toRawStringLiteral(string(note_type.Size))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", note_typeIdent, __gong__toRawStringLiteral(string(note_type.EnclosedText))))
		}
	}
	if stageSet.Stage != nil {
		for _, notehead := range __gong__sortStageSetInstances(stageSet.Stage.Noteheads, stageSet.Stage.Notehead_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteheadIdent := "__models" + notehead.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Notehead{Name: %s}).Stage(stageSet.Stage)", noteheadIdent, __gong__toRawStringLiteral(notehead.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteheadIdent, __gong__toRawStringLiteral(notehead.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Filled = %s", noteheadIdent, __gong__toRawStringLiteral(string(notehead.Filled))))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", noteheadIdent, __gong__toRawStringLiteral(string(notehead.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", noteheadIdent, __gong__toRawStringLiteral(notehead.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", noteheadIdent, __gong__toRawStringLiteral(notehead.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", noteheadIdent, __gong__toRawStringLiteral(notehead.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", noteheadIdent, __gong__toRawStringLiteral(notehead.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", noteheadIdent, __gong__toRawStringLiteral(notehead.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", noteheadIdent, __gong__toRawStringLiteral(notehead.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", noteheadIdent, __gong__toRawStringLiteral(notehead.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, notehead_text := range __gong__sortStageSetInstances(stageSet.Stage.Notehead_texts, stageSet.Stage.Notehead_text_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notehead_textIdent := "__models" + notehead_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Notehead_text{Name: %s}).Stage(stageSet.Stage)", notehead_textIdent, __gong__toRawStringLiteral(notehead_text.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notehead_textIdent, __gong__toRawStringLiteral(notehead_text.Name)))
			for _, elem := range notehead_text.Display_text {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Display_text = append(%s.Display_text, %s)", notehead_textIdent, notehead_textIdent, targetIdent))
			}
			for _, elem := range notehead_text.Accidental_text {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accidental_text = append(%s.Accidental_text, %s)", notehead_textIdent, notehead_textIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, numeral := range __gong__sortStageSetInstances(stageSet.Stage.Numerals, stageSet.Stage.Numeral_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			numeralIdent := "__models" + numeral.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Numeral{Name: %s}).Stage(stageSet.Stage)", numeralIdent, __gong__toRawStringLiteral(numeral.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", numeralIdent, __gong__toRawStringLiteral(numeral.Name)))
			if numeral.Numeral_root != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + numeral.Numeral_root.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Numeral_root = %s", numeralIdent, targetIdent))
			}
			if numeral.Numeral_alter != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + numeral.Numeral_alter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Numeral_alter = %s", numeralIdent, targetIdent))
			}
			if numeral.Numeral_key != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + numeral.Numeral_key.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Numeral_key = %s", numeralIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, numeral_key := range __gong__sortStageSetInstances(stageSet.Stage.Numeral_keys, stageSet.Stage.Numeral_key_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			numeral_keyIdent := "__models" + numeral_key.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Numeral_key{Name: %s}).Stage(stageSet.Stage)", numeral_keyIdent, __gong__toRawStringLiteral(numeral_key.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", numeral_keyIdent, __gong__toRawStringLiteral(numeral_key.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", numeral_keyIdent, __gong__toRawStringLiteral(string(numeral_key.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Numeral_fifths = %d", numeral_keyIdent, numeral_key.Numeral_fifths))
			values.WriteString(fmt.Sprintf("\n\t%s.Numeral_mode = %s", numeral_keyIdent, __gong__toRawStringLiteral(numeral_key.Numeral_mode)))
		}
	}
	if stageSet.Stage != nil {
		for _, numeral_root := range __gong__sortStageSetInstances(stageSet.Stage.Numeral_roots, stageSet.Stage.Numeral_root_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			numeral_rootIdent := "__models" + numeral_root.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Numeral_root{Name: %s}).Stage(stageSet.Stage)", numeral_rootIdent, __gong__toRawStringLiteral(numeral_root.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", numeral_rootIdent, __gong__toRawStringLiteral(numeral_root.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", numeral_rootIdent, __gong__toRawStringLiteral(numeral_root.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", numeral_rootIdent, __gong__toRawStringLiteral(numeral_root.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", numeral_rootIdent, __gong__toRawStringLiteral(numeral_root.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", numeral_rootIdent, __gong__toRawStringLiteral(numeral_root.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", numeral_rootIdent, __gong__toRawStringLiteral(numeral_root.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", numeral_rootIdent, __gong__toRawStringLiteral(numeral_root.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", numeral_rootIdent, __gong__toRawStringLiteral(numeral_root.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", numeral_rootIdent, __gong__toRawStringLiteral(numeral_root.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", numeral_rootIdent, __gong__toRawStringLiteral(numeral_root.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", numeral_rootIdent, __gong__toRawStringLiteral(numeral_root.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", numeral_rootIdent, numeral_root.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		for _, octave_shift := range __gong__sortStageSetInstances(stageSet.Stage.Octave_shifts, stageSet.Stage.Octave_shift_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			octave_shiftIdent := "__models" + octave_shift.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Octave_shift{Name: %s}).Stage(stageSet.Stage)", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", octave_shiftIdent, octave_shift.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %d", octave_shiftIdent, octave_shift.Size))
			values.WriteString(fmt.Sprintf("\n\t%s.Dash_length = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Dash_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space_length = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Space_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, offset := range __gong__sortStageSetInstances(stageSet.Stage.Offsets, stageSet.Stage.Offset_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			offsetIdent := "__models" + offset.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Offset{Name: %s}).Stage(stageSet.Stage)", offsetIdent, __gong__toRawStringLiteral(offset.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", offsetIdent, __gong__toRawStringLiteral(offset.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sound = %s", offsetIdent, __gong__toRawStringLiteral(string(offset.Sound))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", offsetIdent, __gong__toRawStringLiteral(offset.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, opus := range __gong__sortStageSetInstances(stageSet.Stage.Opuss, stageSet.Stage.Opus_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			opusIdent := "__models" + opus.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Opus{Name: %s}).Stage(stageSet.Stage)", opusIdent, __gong__toRawStringLiteral(opus.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", opusIdent, __gong__toRawStringLiteral(opus.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", opusIdent, __gong__toRawStringLiteral(opus.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", opusIdent, __gong__toRawStringLiteral(opus.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", opusIdent, __gong__toRawStringLiteral(opus.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", opusIdent, __gong__toRawStringLiteral(opus.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", opusIdent, __gong__toRawStringLiteral(opus.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", opusIdent, __gong__toRawStringLiteral(opus.Actuate)))
		}
	}
	if stageSet.Stage != nil {
		for _, ornaments := range __gong__sortStageSetInstances(stageSet.Stage.Ornamentss, stageSet.Stage.Ornaments_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ornamentsIdent := "__models" + ornaments.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Ornaments{Name: %s}).Stage(stageSet.Stage)", ornamentsIdent, __gong__toRawStringLiteral(ornaments.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ornamentsIdent, __gong__toRawStringLiteral(ornaments.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ornamentsIdent, __gong__toRawStringLiteral(ornaments.Id)))
			for _, elem := range ornaments.Trill_mark {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Trill_mark = append(%s.Trill_mark, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Turn {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Turn = append(%s.Turn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Delayed_turn {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Delayed_turn = append(%s.Delayed_turn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Inverted_turn {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inverted_turn = append(%s.Inverted_turn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Delayed_inverted_turn {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Delayed_inverted_turn = append(%s.Delayed_inverted_turn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Vertical_turn {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Vertical_turn = append(%s.Vertical_turn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Inverted_vertical_turn {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inverted_vertical_turn = append(%s.Inverted_vertical_turn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Shake {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Shake = append(%s.Shake, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Wavy_line {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wavy_line = append(%s.Wavy_line, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Mordent {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mordent = append(%s.Mordent, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Inverted_mordent {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inverted_mordent = append(%s.Inverted_mordent, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Schleifer {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Schleifer = append(%s.Schleifer, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Tremolo {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tremolo = append(%s.Tremolo, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Haydn {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Haydn = append(%s.Haydn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Other_ornament {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_ornament = append(%s.Other_ornament, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Accidental_mark {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accidental_mark = append(%s.Accidental_mark, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, other_appearance := range __gong__sortStageSetInstances(stageSet.Stage.Other_appearances, stageSet.Stage.Other_appearance_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			other_appearanceIdent := "__models" + other_appearance.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Other_appearance{Name: %s}).Stage(stageSet.Stage)", other_appearanceIdent, __gong__toRawStringLiteral(other_appearance.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", other_appearanceIdent, __gong__toRawStringLiteral(other_appearance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", other_appearanceIdent, __gong__toRawStringLiteral(other_appearance.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", other_appearanceIdent, __gong__toRawStringLiteral(other_appearance.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, other_direction := range __gong__sortStageSetInstances(stageSet.Stage.Other_directions, stageSet.Stage.Other_direction_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			other_directionIdent := "__models" + other_direction.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Other_direction{Name: %s}).Stage(stageSet.Stage)", other_directionIdent, __gong__toRawStringLiteral(other_direction.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", other_directionIdent, __gong__toRawStringLiteral(string(other_direction.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", other_directionIdent, __gong__toRawStringLiteral(other_direction.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, other_listening := range __gong__sortStageSetInstances(stageSet.Stage.Other_listenings, stageSet.Stage.Other_listening_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			other_listeningIdent := "__models" + other_listening.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Other_listening{Name: %s}).Stage(stageSet.Stage)", other_listeningIdent, __gong__toRawStringLiteral(other_listening.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", other_listeningIdent, __gong__toRawStringLiteral(other_listening.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", other_listeningIdent, __gong__toRawStringLiteral(other_listening.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", other_listeningIdent, __gong__toRawStringLiteral(other_listening.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_only = %s", other_listeningIdent, __gong__toRawStringLiteral(string(other_listening.Time_only))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", other_listeningIdent, __gong__toRawStringLiteral(other_listening.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, other_notation := range __gong__sortStageSetInstances(stageSet.Stage.Other_notations, stageSet.Stage.Other_notation_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			other_notationIdent := "__models" + other_notation.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Other_notation{Name: %s}).Stage(stageSet.Stage)", other_notationIdent, __gong__toRawStringLiteral(other_notation.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", other_notationIdent, __gong__toRawStringLiteral(string(other_notation.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", other_notationIdent, other_notation.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", other_notationIdent, __gong__toRawStringLiteral(string(other_notation.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", other_notationIdent, __gong__toRawStringLiteral(other_notation.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, other_placement_text := range __gong__sortStageSetInstances(stageSet.Stage.Other_placement_texts, stageSet.Stage.Other_placement_text_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			other_placement_textIdent := "__models" + other_placement_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Other_placement_text{Name: %s}).Stage(stageSet.Stage)", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, other_play := range __gong__sortStageSetInstances(stageSet.Stage.Other_plays, stageSet.Stage.Other_play_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			other_playIdent := "__models" + other_play.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Other_play{Name: %s}).Stage(stageSet.Stage)", other_playIdent, __gong__toRawStringLiteral(other_play.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", other_playIdent, __gong__toRawStringLiteral(other_play.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", other_playIdent, __gong__toRawStringLiteral(other_play.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", other_playIdent, __gong__toRawStringLiteral(other_play.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, other_text := range __gong__sortStageSetInstances(stageSet.Stage.Other_texts, stageSet.Stage.Other_text_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			other_textIdent := "__models" + other_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Other_text{Name: %s}).Stage(stageSet.Stage)", other_textIdent, __gong__toRawStringLiteral(other_text.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", other_textIdent, __gong__toRawStringLiteral(other_text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", other_textIdent, __gong__toRawStringLiteral(other_text.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", other_textIdent, __gong__toRawStringLiteral(other_text.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, page_layout := range __gong__sortStageSetInstances(stageSet.Stage.Page_layouts, stageSet.Stage.Page_layout_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			page_layoutIdent := "__models" + page_layout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Page_layout{Name: %s}).Stage(stageSet.Stage)", page_layoutIdent, __gong__toRawStringLiteral(page_layout.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", page_layoutIdent, __gong__toRawStringLiteral(page_layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Page_height = %s", page_layoutIdent, __gong__toRawStringLiteral(page_layout.Page_height)))
			values.WriteString(fmt.Sprintf("\n\t%s.Page_width = %s", page_layoutIdent, __gong__toRawStringLiteral(page_layout.Page_width)))
			if page_layout.Page_margins != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + page_layout.Page_margins.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Page_margins = %s", page_layoutIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, page_margins := range __gong__sortStageSetInstances(stageSet.Stage.Page_marginss, stageSet.Stage.Page_margins_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			page_marginsIdent := "__models" + page_margins.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Page_margins{Name: %s}).Stage(stageSet.Stage)", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Left_margin = %s", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Left_margin)))
			values.WriteString(fmt.Sprintf("\n\t%s.Right_margin = %s", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Right_margin)))
			values.WriteString(fmt.Sprintf("\n\t%s.Top_margin = %s", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Top_margin)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bottom_margin = %s", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Bottom_margin)))
		}
	}
	if stageSet.Stage != nil {
		for _, part_clef := range __gong__sortStageSetInstances(stageSet.Stage.Part_clefs, stageSet.Stage.Part_clef_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			part_clefIdent := "__models" + part_clef.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Part_clef{Name: %s}).Stage(stageSet.Stage)", part_clefIdent, __gong__toRawStringLiteral(part_clef.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", part_clefIdent, __gong__toRawStringLiteral(part_clef.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sign = %s", part_clefIdent, __gong__toRawStringLiteral(part_clef.Sign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", part_clefIdent, part_clef.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.Clef_octave_change = %d", part_clefIdent, part_clef.Clef_octave_change))
		}
	}
	if stageSet.Stage != nil {
		for _, part_group := range __gong__sortStageSetInstances(stageSet.Stage.Part_groups, stageSet.Stage.Part_group_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			part_groupIdent := "__models" + part_group.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Part_group{Name: %s}).Stage(stageSet.Stage)", part_groupIdent, __gong__toRawStringLiteral(part_group.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", part_groupIdent, __gong__toRawStringLiteral(part_group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", part_groupIdent, __gong__toRawStringLiteral(string(part_group.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", part_groupIdent, __gong__toRawStringLiteral(part_group.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_time = %s", part_groupIdent, __gong__toRawStringLiteral(part_group.Group_time)))
			if part_group.Group_name != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + part_group.Group_name.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group_name = %s", part_groupIdent, targetIdent))
			}
			if part_group.Group_name_display != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + part_group.Group_name_display.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group_name_display = %s", part_groupIdent, targetIdent))
			}
			if part_group.Group_abbreviation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + part_group.Group_abbreviation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group_abbreviation = %s", part_groupIdent, targetIdent))
			}
			if part_group.Group_abbreviation_display != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + part_group.Group_abbreviation_display.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group_abbreviation_display = %s", part_groupIdent, targetIdent))
			}
			if part_group.Group_symbol != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + part_group.Group_symbol.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group_symbol = %s", part_groupIdent, targetIdent))
			}
			if part_group.Group_barline != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + part_group.Group_barline.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group_barline = %s", part_groupIdent, targetIdent))
			}
			if part_group.Footnote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + part_group.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", part_groupIdent, targetIdent))
			}
			if part_group.Level != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + part_group.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", part_groupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, part_link := range __gong__sortStageSetInstances(stageSet.Stage.Part_links, stageSet.Stage.Part_link_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			part_linkIdent := "__models" + part_link.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Part_link{Name: %s}).Stage(stageSet.Stage)", part_linkIdent, __gong__toRawStringLiteral(part_link.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Actuate)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_link = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Group_link)))
			for _, elem := range part_link.Instrument_link {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Instrument_link = append(%s.Instrument_link, %s)", part_linkIdent, part_linkIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, part_list := range __gong__sortStageSetInstances(stageSet.Stage.Part_lists, stageSet.Stage.Part_list_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			part_listIdent := "__models" + part_list.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Part_list{Name: %s}).Stage(stageSet.Stage)", part_listIdent, __gong__toRawStringLiteral(part_list.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", part_listIdent, __gong__toRawStringLiteral(part_list.Name)))
			if part_list.Part_group != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + part_list.Part_group.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_group = %s", part_listIdent, targetIdent))
			}
			if part_list.Score_part != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + part_list.Score_part.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Score_part = %s", part_listIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, part_name := range __gong__sortStageSetInstances(stageSet.Stage.Part_names, stageSet.Stage.Part_name_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			part_nameIdent := "__models" + part_name.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Part_name{Name: %s}).Stage(stageSet.Stage)", part_nameIdent, __gong__toRawStringLiteral(part_name.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", part_nameIdent, __gong__toRawStringLiteral(part_name.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", part_nameIdent, __gong__toRawStringLiteral(part_name.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", part_nameIdent, __gong__toRawStringLiteral(part_name.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", part_nameIdent, __gong__toRawStringLiteral(part_name.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", part_nameIdent, __gong__toRawStringLiteral(part_name.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", part_nameIdent, __gong__toRawStringLiteral(part_name.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", part_nameIdent, __gong__toRawStringLiteral(part_name.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", part_nameIdent, __gong__toRawStringLiteral(part_name.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", part_nameIdent, __gong__toRawStringLiteral(part_name.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", part_nameIdent, __gong__toRawStringLiteral(part_name.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", part_nameIdent, __gong__toRawStringLiteral(string(part_name.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", part_nameIdent, __gong__toRawStringLiteral(string(part_name.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", part_nameIdent, __gong__toRawStringLiteral(part_name.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, part_symbol := range __gong__sortStageSetInstances(stageSet.Stage.Part_symbols, stageSet.Stage.Part_symbol_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			part_symbolIdent := "__models" + part_symbol.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Part_symbol{Name: %s}).Stage(stageSet.Stage)", part_symbolIdent, __gong__toRawStringLiteral(part_symbol.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", part_symbolIdent, __gong__toRawStringLiteral(part_symbol.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Top_staff = %d", part_symbolIdent, part_symbol.Top_staff))
			values.WriteString(fmt.Sprintf("\n\t%s.Bottom_staff = %d", part_symbolIdent, part_symbol.Bottom_staff))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", part_symbolIdent, __gong__toRawStringLiteral(part_symbol.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", part_symbolIdent, __gong__toRawStringLiteral(part_symbol.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", part_symbolIdent, __gong__toRawStringLiteral(part_symbol.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", part_symbolIdent, __gong__toRawStringLiteral(part_symbol.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", part_symbolIdent, __gong__toRawStringLiteral(part_symbol.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", part_symbolIdent, __gong__toRawStringLiteral(string(part_symbol.EnclosedText))))
		}
	}
	if stageSet.Stage != nil {
		for _, part_transpose := range __gong__sortStageSetInstances(stageSet.Stage.Part_transposes, stageSet.Stage.Part_transpose_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			part_transposeIdent := "__models" + part_transpose.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Part_transpose{Name: %s}).Stage(stageSet.Stage)", part_transposeIdent, __gong__toRawStringLiteral(part_transpose.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", part_transposeIdent, __gong__toRawStringLiteral(part_transpose.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Diatonic = %d", part_transposeIdent, part_transpose.Diatonic))
			values.WriteString(fmt.Sprintf("\n\t%s.Chromatic = %s", part_transposeIdent, __gong__toRawStringLiteral(part_transpose.Chromatic)))
			values.WriteString(fmt.Sprintf("\n\t%s.Octave_change = %d", part_transposeIdent, part_transpose.Octave_change))
			values.WriteString(fmt.Sprintf("\n\t%s.Double = %f", part_transposeIdent, part_transpose.Double))
		}
	}
	if stageSet.Stage != nil {
		for _, pedal := range __gong__sortStageSetInstances(stageSet.Stage.Pedals, stageSet.Stage.Pedal_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pedalIdent := "__models" + pedal.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Pedal{Name: %s}).Stage(stageSet.Stage)", pedalIdent, __gong__toRawStringLiteral(pedal.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", pedalIdent, pedal.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %s", pedalIdent, __gong__toRawStringLiteral(string(pedal.Line))))
			values.WriteString(fmt.Sprintf("\n\t%s.Sign = %s", pedalIdent, __gong__toRawStringLiteral(string(pedal.Sign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Abbreviated = %s", pedalIdent, __gong__toRawStringLiteral(string(pedal.Abbreviated))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, pedal_tuning := range __gong__sortStageSetInstances(stageSet.Stage.Pedal_tunings, stageSet.Stage.Pedal_tuning_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pedal_tuningIdent := "__models" + pedal_tuning.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Pedal_tuning{Name: %s}).Stage(stageSet.Stage)", pedal_tuningIdent, __gong__toRawStringLiteral(pedal_tuning.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pedal_tuningIdent, __gong__toRawStringLiteral(pedal_tuning.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Pedal_step = %s", pedal_tuningIdent, __gong__toRawStringLiteral(string(pedal_tuning.Pedal_step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Pedal_alter = %s", pedal_tuningIdent, __gong__toRawStringLiteral(pedal_tuning.Pedal_alter)))
		}
	}
	if stageSet.Stage != nil {
		for _, per_minute := range __gong__sortStageSetInstances(stageSet.Stage.Per_minutes, stageSet.Stage.Per_minute_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			per_minuteIdent := "__models" + per_minute.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Per_minute{Name: %s}).Stage(stageSet.Stage)", per_minuteIdent, __gong__toRawStringLiteral(per_minute.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", per_minuteIdent, __gong__toRawStringLiteral(per_minute.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", per_minuteIdent, __gong__toRawStringLiteral(per_minute.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", per_minuteIdent, __gong__toRawStringLiteral(per_minute.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", per_minuteIdent, __gong__toRawStringLiteral(per_minute.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", per_minuteIdent, __gong__toRawStringLiteral(per_minute.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", per_minuteIdent, __gong__toRawStringLiteral(per_minute.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, percussion := range __gong__sortStageSetInstances(stageSet.Stage.Percussions, stageSet.Stage.Percussion_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			percussionIdent := "__models" + percussion.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Percussion{Name: %s}).Stage(stageSet.Stage)", percussionIdent, __gong__toRawStringLiteral(percussion.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Enclosure)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Stick_location = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Stick_location)))
			if percussion.Glass != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + percussion.Glass.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Glass = %s", percussionIdent, targetIdent))
			}
			if percussion.Metal != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + percussion.Metal.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metal = %s", percussionIdent, targetIdent))
			}
			if percussion.Wood != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + percussion.Wood.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wood = %s", percussionIdent, targetIdent))
			}
			if percussion.Pitched != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + percussion.Pitched.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pitched = %s", percussionIdent, targetIdent))
			}
			if percussion.Membrane != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + percussion.Membrane.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Membrane = %s", percussionIdent, targetIdent))
			}
			if percussion.Effect != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + percussion.Effect.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Effect = %s", percussionIdent, targetIdent))
			}
			if percussion.Timpani != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + percussion.Timpani.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Timpani = %s", percussionIdent, targetIdent))
			}
			if percussion.Beater != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + percussion.Beater.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Beater = %s", percussionIdent, targetIdent))
			}
			if percussion.Stick != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + percussion.Stick.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stick = %s", percussionIdent, targetIdent))
			}
			if percussion.Other_percussion != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + percussion.Other_percussion.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_percussion = %s", percussionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, pitch := range __gong__sortStageSetInstances(stageSet.Stage.Pitchs, stageSet.Stage.Pitch_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pitchIdent := "__models" + pitch.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Pitch{Name: %s}).Stage(stageSet.Stage)", pitchIdent, __gong__toRawStringLiteral(pitch.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pitchIdent, __gong__toRawStringLiteral(pitch.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Step = %s", pitchIdent, __gong__toRawStringLiteral(string(pitch.Step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Alter = %s", pitchIdent, __gong__toRawStringLiteral(pitch.Alter)))
			values.WriteString(fmt.Sprintf("\n\t%s.Octave = %d", pitchIdent, pitch.Octave))
		}
	}
	if stageSet.Stage != nil {
		for _, pitched := range __gong__sortStageSetInstances(stageSet.Stage.Pitcheds, stageSet.Stage.Pitched_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pitchedIdent := "__models" + pitched.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Pitched{Name: %s}).Stage(stageSet.Stage)", pitchedIdent, __gong__toRawStringLiteral(pitched.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pitchedIdent, __gong__toRawStringLiteral(pitched.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", pitchedIdent, __gong__toRawStringLiteral(pitched.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", pitchedIdent, __gong__toRawStringLiteral(pitched.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, placement_text := range __gong__sortStageSetInstances(stageSet.Stage.Placement_texts, stageSet.Stage.Placement_text_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			placement_textIdent := "__models" + placement_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Placement_text{Name: %s}).Stage(stageSet.Stage)", placement_textIdent, __gong__toRawStringLiteral(placement_text.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", placement_textIdent, __gong__toRawStringLiteral(placement_text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", placement_textIdent, __gong__toRawStringLiteral(placement_text.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", placement_textIdent, __gong__toRawStringLiteral(placement_text.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", placement_textIdent, __gong__toRawStringLiteral(placement_text.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", placement_textIdent, __gong__toRawStringLiteral(placement_text.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", placement_textIdent, __gong__toRawStringLiteral(placement_text.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", placement_textIdent, __gong__toRawStringLiteral(placement_text.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", placement_textIdent, __gong__toRawStringLiteral(placement_text.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", placement_textIdent, __gong__toRawStringLiteral(placement_text.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", placement_textIdent, __gong__toRawStringLiteral(placement_text.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", placement_textIdent, __gong__toRawStringLiteral(placement_text.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", placement_textIdent, __gong__toRawStringLiteral(placement_text.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, play := range __gong__sortStageSetInstances(stageSet.Stage.Plays, stageSet.Stage.Play_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			playIdent := "__models" + play.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Play{Name: %s}).Stage(stageSet.Stage)", playIdent, __gong__toRawStringLiteral(play.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", playIdent, __gong__toRawStringLiteral(play.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", playIdent, __gong__toRawStringLiteral(play.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ipa = %s", playIdent, __gong__toRawStringLiteral(play.Ipa)))
			values.WriteString(fmt.Sprintf("\n\t%s.Mute = %s", playIdent, __gong__toRawStringLiteral(play.Mute)))
			values.WriteString(fmt.Sprintf("\n\t%s.Semi_pitched = %s", playIdent, __gong__toRawStringLiteral(play.Semi_pitched)))
			for _, elem := range play.Other_play {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_play = append(%s.Other_play, %s)", playIdent, playIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, player := range __gong__sortStageSetInstances(stageSet.Stage.Players, stageSet.Stage.Player_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			playerIdent := "__models" + player.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Player{Name: %s}).Stage(stageSet.Stage)", playerIdent, __gong__toRawStringLiteral(player.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", playerIdent, __gong__toRawStringLiteral(player.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", playerIdent, __gong__toRawStringLiteral(player.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Player_name = %s", playerIdent, __gong__toRawStringLiteral(player.Player_name)))
		}
	}
	if stageSet.Stage != nil {
		for _, principal_voice := range __gong__sortStageSetInstances(stageSet.Stage.Principal_voices, stageSet.Stage.Principal_voice_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			principal_voiceIdent := "__models" + principal_voice.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Principal_voice{Name: %s}).Stage(stageSet.Stage)", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", principal_voiceIdent, __gong__toRawStringLiteral(string(principal_voice.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Symbol = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Symbol)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, print := range __gong__sortStageSetInstances(stageSet.Stage.Prints, stageSet.Stage.Print_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			printIdent := "__models" + print.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Print{Name: %s}).Stage(stageSet.Stage)", printIdent, __gong__toRawStringLiteral(print.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", printIdent, __gong__toRawStringLiteral(print.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff_spacing = %s", printIdent, __gong__toRawStringLiteral(print.Staff_spacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.New_system = %s", printIdent, __gong__toRawStringLiteral(string(print.New_system))))
			values.WriteString(fmt.Sprintf("\n\t%s.New_page = %s", printIdent, __gong__toRawStringLiteral(string(print.New_page))))
			values.WriteString(fmt.Sprintf("\n\t%s.Blank_page = %d", printIdent, print.Blank_page))
			values.WriteString(fmt.Sprintf("\n\t%s.Page_number = %s", printIdent, __gong__toRawStringLiteral(print.Page_number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", printIdent, __gong__toRawStringLiteral(print.Id)))
			if print.Page_layout != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + print.Page_layout.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Page_layout = %s", printIdent, targetIdent))
			}
			if print.System_layout != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + print.System_layout.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System_layout = %s", printIdent, targetIdent))
			}
			for _, elem := range print.Staff_layout {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staff_layout = append(%s.Staff_layout, %s)", printIdent, printIdent, targetIdent))
			}
			if print.Measure_layout != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + print.Measure_layout.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure_layout = %s", printIdent, targetIdent))
			}
			if print.Measure_numbering != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + print.Measure_numbering.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure_numbering = %s", printIdent, targetIdent))
			}
			if print.Part_name_display != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + print.Part_name_display.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_name_display = %s", printIdent, targetIdent))
			}
			if print.Part_abbreviation_display != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + print.Part_abbreviation_display.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_abbreviation_display = %s", printIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, release := range __gong__sortStageSetInstances(stageSet.Stage.Releases, stageSet.Stage.Release_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			releaseIdent := "__models" + release.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Release{Name: %s}).Stage(stageSet.Stage)", releaseIdent, __gong__toRawStringLiteral(release.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", releaseIdent, __gong__toRawStringLiteral(release.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, repeat := range __gong__sortStageSetInstances(stageSet.Stage.Repeats, stageSet.Stage.Repeat_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			repeatIdent := "__models" + repeat.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Repeat{Name: %s}).Stage(stageSet.Stage)", repeatIdent, __gong__toRawStringLiteral(repeat.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", repeatIdent, __gong__toRawStringLiteral(repeat.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", repeatIdent, __gong__toRawStringLiteral(repeat.Direction)))
			values.WriteString(fmt.Sprintf("\n\t%s.Times = %d", repeatIdent, repeat.Times))
			values.WriteString(fmt.Sprintf("\n\t%s.After_jump = %s", repeatIdent, __gong__toRawStringLiteral(string(repeat.After_jump))))
			values.WriteString(fmt.Sprintf("\n\t%s.Winged = %s", repeatIdent, __gong__toRawStringLiteral(repeat.Winged)))
		}
	}
	if stageSet.Stage != nil {
		for _, rest := range __gong__sortStageSetInstances(stageSet.Stage.Rests, stageSet.Stage.Rest_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			restIdent := "__models" + rest.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Rest{Name: %s}).Stage(stageSet.Stage)", restIdent, __gong__toRawStringLiteral(rest.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", restIdent, __gong__toRawStringLiteral(rest.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Measure = %s", restIdent, __gong__toRawStringLiteral(string(rest.Measure))))
			values.WriteString(fmt.Sprintf("\n\t%s.Display_step = %s", restIdent, __gong__toRawStringLiteral(string(rest.Display_step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Display_octave = %d", restIdent, rest.Display_octave))
		}
	}
	if stageSet.Stage != nil {
		for _, root := range __gong__sortStageSetInstances(stageSet.Stage.Roots, stageSet.Stage.Root_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rootIdent := "__models" + root.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Root{Name: %s}).Stage(stageSet.Stage)", rootIdent, __gong__toRawStringLiteral(root.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rootIdent, __gong__toRawStringLiteral(root.Name)))
			if root.Root_step != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + root.Root_step.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Root_step = %s", rootIdent, targetIdent))
			}
			if root.Root_alter != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + root.Root_alter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Root_alter = %s", rootIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, root_step := range __gong__sortStageSetInstances(stageSet.Stage.Root_steps, stageSet.Stage.Root_step_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			root_stepIdent := "__models" + root_step.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Root_step{Name: %s}).Stage(stageSet.Stage)", root_stepIdent, __gong__toRawStringLiteral(root_step.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", root_stepIdent, __gong__toRawStringLiteral(root_step.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", root_stepIdent, __gong__toRawStringLiteral(root_step.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", root_stepIdent, __gong__toRawStringLiteral(root_step.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", root_stepIdent, __gong__toRawStringLiteral(root_step.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", root_stepIdent, __gong__toRawStringLiteral(root_step.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", root_stepIdent, __gong__toRawStringLiteral(root_step.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", root_stepIdent, __gong__toRawStringLiteral(root_step.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", root_stepIdent, __gong__toRawStringLiteral(root_step.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", root_stepIdent, __gong__toRawStringLiteral(root_step.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", root_stepIdent, __gong__toRawStringLiteral(root_step.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", root_stepIdent, __gong__toRawStringLiteral(root_step.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", root_stepIdent, __gong__toRawStringLiteral(string(root_step.EnclosedText))))
		}
	}
	if stageSet.Stage != nil {
		for _, scaling := range __gong__sortStageSetInstances(stageSet.Stage.Scalings, stageSet.Stage.Scaling_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			scalingIdent := "__models" + scaling.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Scaling{Name: %s}).Stage(stageSet.Stage)", scalingIdent, __gong__toRawStringLiteral(scaling.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scalingIdent, __gong__toRawStringLiteral(scaling.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Millimeters = %s", scalingIdent, __gong__toRawStringLiteral(scaling.Millimeters)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tenths = %s", scalingIdent, __gong__toRawStringLiteral(scaling.Tenths)))
		}
	}
	if stageSet.Stage != nil {
		for _, scordatura := range __gong__sortStageSetInstances(stageSet.Stage.Scordaturas, stageSet.Stage.Scordatura_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			scordaturaIdent := "__models" + scordatura.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Scordatura{Name: %s}).Stage(stageSet.Stage)", scordaturaIdent, __gong__toRawStringLiteral(scordatura.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scordaturaIdent, __gong__toRawStringLiteral(scordatura.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", scordaturaIdent, __gong__toRawStringLiteral(scordatura.Id)))
			for _, elem := range scordatura.Accord {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accord = append(%s.Accord, %s)", scordaturaIdent, scordaturaIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, score_instrument := range __gong__sortStageSetInstances(stageSet.Stage.Score_instruments, stageSet.Stage.Score_instrument_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			score_instrumentIdent := "__models" + score_instrument.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Score_instrument{Name: %s}).Stage(stageSet.Stage)", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Instrument_name = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Instrument_name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Instrument_abbreviation = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Instrument_abbreviation)))
			values.WriteString(fmt.Sprintf("\n\t%s.Instrument_sound = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Instrument_sound)))
			values.WriteString(fmt.Sprintf("\n\t%s.Solo = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Solo)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ensemble = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Ensemble)))
			if score_instrument.Virtual_instrument != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_instrument.Virtual_instrument.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Virtual_instrument = %s", score_instrumentIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, score_part := range __gong__sortStageSetInstances(stageSet.Stage.Score_parts, stageSet.Stage.Score_part_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			score_partIdent := "__models" + score_part.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Score_part{Name: %s}).Stage(stageSet.Stage)", score_partIdent, __gong__toRawStringLiteral(score_part.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", score_partIdent, __gong__toRawStringLiteral(score_part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", score_partIdent, __gong__toRawStringLiteral(score_part.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group = %s", score_partIdent, __gong__toRawStringLiteral(score_part.Group)))
			if score_part.Identification != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_part.Identification.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identification = %s", score_partIdent, targetIdent))
			}
			for _, elem := range score_part.Part_link {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_link = append(%s.Part_link, %s)", score_partIdent, score_partIdent, targetIdent))
			}
			if score_part.Part_name != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_part.Part_name.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_name = %s", score_partIdent, targetIdent))
			}
			if score_part.Part_name_display != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_part.Part_name_display.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_name_display = %s", score_partIdent, targetIdent))
			}
			if score_part.Part_abbreviation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_part.Part_abbreviation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_abbreviation = %s", score_partIdent, targetIdent))
			}
			if score_part.Part_abbreviation_display != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_part.Part_abbreviation_display.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_abbreviation_display = %s", score_partIdent, targetIdent))
			}
			for _, elem := range score_part.Score_instrument {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Score_instrument = append(%s.Score_instrument, %s)", score_partIdent, score_partIdent, targetIdent))
			}
			for _, elem := range score_part.Player {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Player = append(%s.Player, %s)", score_partIdent, score_partIdent, targetIdent))
			}
			for _, elem := range score_part.Midi_device {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Midi_device = append(%s.Midi_device, %s)", score_partIdent, score_partIdent, targetIdent))
			}
			for _, elem := range score_part.Midi_instrument {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Midi_instrument = append(%s.Midi_instrument, %s)", score_partIdent, score_partIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, score_partwise := range __gong__sortStageSetInstances(stageSet.Stage.Score_partwises, stageSet.Stage.Score_partwise_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			score_partwiseIdent := "__models" + score_partwise.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Score_partwise{Name: %s}).Stage(stageSet.Stage)", score_partwiseIdent, __gong__toRawStringLiteral(score_partwise.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", score_partwiseIdent, __gong__toRawStringLiteral(score_partwise.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Version = %s", score_partwiseIdent, __gong__toRawStringLiteral(score_partwise.Version)))
			values.WriteString(fmt.Sprintf("\n\t%s.Movement_number = %s", score_partwiseIdent, __gong__toRawStringLiteral(score_partwise.Movement_number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Movement_title = %s", score_partwiseIdent, __gong__toRawStringLiteral(score_partwise.Movement_title)))
			if score_partwise.Work != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_partwise.Work.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Work = %s", score_partwiseIdent, targetIdent))
			}
			if score_partwise.Identification != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_partwise.Identification.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identification = %s", score_partwiseIdent, targetIdent))
			}
			if score_partwise.Defaults != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_partwise.Defaults.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Defaults = %s", score_partwiseIdent, targetIdent))
			}
			for _, elem := range score_partwise.Credit {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit = append(%s.Credit, %s)", score_partwiseIdent, score_partwiseIdent, targetIdent))
			}
			if score_partwise.Part_list != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_partwise.Part_list.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_list = %s", score_partwiseIdent, targetIdent))
			}
			for _, elem := range score_partwise.Part {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = append(%s.Part, %s)", score_partwiseIdent, score_partwiseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, score_timewise := range __gong__sortStageSetInstances(stageSet.Stage.Score_timewises, stageSet.Stage.Score_timewise_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			score_timewiseIdent := "__models" + score_timewise.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Score_timewise{Name: %s}).Stage(stageSet.Stage)", score_timewiseIdent, __gong__toRawStringLiteral(score_timewise.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", score_timewiseIdent, __gong__toRawStringLiteral(score_timewise.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Version = %s", score_timewiseIdent, __gong__toRawStringLiteral(score_timewise.Version)))
			values.WriteString(fmt.Sprintf("\n\t%s.Movement_number = %s", score_timewiseIdent, __gong__toRawStringLiteral(score_timewise.Movement_number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Movement_title = %s", score_timewiseIdent, __gong__toRawStringLiteral(score_timewise.Movement_title)))
			if score_timewise.Work != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_timewise.Work.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Work = %s", score_timewiseIdent, targetIdent))
			}
			if score_timewise.Identification != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_timewise.Identification.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identification = %s", score_timewiseIdent, targetIdent))
			}
			if score_timewise.Defaults != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_timewise.Defaults.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Defaults = %s", score_timewiseIdent, targetIdent))
			}
			for _, elem := range score_timewise.Credit {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit = append(%s.Credit, %s)", score_timewiseIdent, score_timewiseIdent, targetIdent))
			}
			if score_timewise.Part_list != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + score_timewise.Part_list.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_list = %s", score_timewiseIdent, targetIdent))
			}
			for _, elem := range score_timewise.Measure {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure = append(%s.Measure, %s)", score_timewiseIdent, score_timewiseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, segno := range __gong__sortStageSetInstances(stageSet.Stage.Segnos, stageSet.Stage.Segno_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			segnoIdent := "__models" + segno.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Segno{Name: %s}).Stage(stageSet.Stage)", segnoIdent, __gong__toRawStringLiteral(segno.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", segnoIdent, __gong__toRawStringLiteral(segno.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", segnoIdent, __gong__toRawStringLiteral(segno.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", segnoIdent, __gong__toRawStringLiteral(segno.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", segnoIdent, __gong__toRawStringLiteral(segno.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", segnoIdent, __gong__toRawStringLiteral(segno.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", segnoIdent, __gong__toRawStringLiteral(segno.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", segnoIdent, __gong__toRawStringLiteral(segno.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", segnoIdent, __gong__toRawStringLiteral(segno.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", segnoIdent, __gong__toRawStringLiteral(segno.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", segnoIdent, __gong__toRawStringLiteral(segno.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", segnoIdent, __gong__toRawStringLiteral(segno.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", segnoIdent, __gong__toRawStringLiteral(segno.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", segnoIdent, __gong__toRawStringLiteral(segno.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", segnoIdent, __gong__toRawStringLiteral(segno.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, slash := range __gong__sortStageSetInstances(stageSet.Stage.Slashs, stageSet.Stage.Slash_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			slashIdent := "__models" + slash.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Slash{Name: %s}).Stage(stageSet.Stage)", slashIdent, __gong__toRawStringLiteral(slash.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", slashIdent, __gong__toRawStringLiteral(slash.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", slashIdent, __gong__toRawStringLiteral(string(slash.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Use_dots = %s", slashIdent, __gong__toRawStringLiteral(string(slash.Use_dots))))
			values.WriteString(fmt.Sprintf("\n\t%s.Use_stems = %s", slashIdent, __gong__toRawStringLiteral(string(slash.Use_stems))))
			values.WriteString(fmt.Sprintf("\n\t%s.Slash_type = %s", slashIdent, __gong__toRawStringLiteral(string(slash.Slash_type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Slash_dot = %s", slashIdent, __gong__toRawStringLiteral(slash.Slash_dot)))
			values.WriteString(fmt.Sprintf("\n\t%s.Except_voice = %s", slashIdent, __gong__toRawStringLiteral(slash.Except_voice)))
		}
	}
	if stageSet.Stage != nil {
		for _, slide := range __gong__sortStageSetInstances(stageSet.Stage.Slides, stageSet.Stage.Slide_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			slideIdent := "__models" + slide.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Slide{Name: %s}).Stage(stageSet.Stage)", slideIdent, __gong__toRawStringLiteral(slide.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", slideIdent, __gong__toRawStringLiteral(slide.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", slideIdent, __gong__toRawStringLiteral(string(slide.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", slideIdent, slide.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_type = %s", slideIdent, __gong__toRawStringLiteral(slide.Line_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dash_length = %s", slideIdent, __gong__toRawStringLiteral(slide.Dash_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space_length = %s", slideIdent, __gong__toRawStringLiteral(slide.Space_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", slideIdent, __gong__toRawStringLiteral(slide.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", slideIdent, __gong__toRawStringLiteral(slide.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", slideIdent, __gong__toRawStringLiteral(slide.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", slideIdent, __gong__toRawStringLiteral(slide.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", slideIdent, __gong__toRawStringLiteral(slide.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", slideIdent, __gong__toRawStringLiteral(slide.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", slideIdent, __gong__toRawStringLiteral(slide.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", slideIdent, __gong__toRawStringLiteral(slide.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", slideIdent, __gong__toRawStringLiteral(slide.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Accelerate = %s", slideIdent, __gong__toRawStringLiteral(string(slide.Accelerate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", slideIdent, __gong__toRawStringLiteral(slide.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.First_beat = %s", slideIdent, __gong__toRawStringLiteral(slide.First_beat)))
			values.WriteString(fmt.Sprintf("\n\t%s.Last_beat = %s", slideIdent, __gong__toRawStringLiteral(slide.Last_beat)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", slideIdent, __gong__toRawStringLiteral(slide.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", slideIdent, __gong__toRawStringLiteral(slide.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, slur := range __gong__sortStageSetInstances(stageSet.Stage.Slurs, stageSet.Stage.Slur_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			slurIdent := "__models" + slur.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Slur{Name: %s}).Stage(stageSet.Stage)", slurIdent, __gong__toRawStringLiteral(slur.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", slurIdent, __gong__toRawStringLiteral(slur.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", slurIdent, __gong__toRawStringLiteral(string(slur.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", slurIdent, slur.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_type = %s", slurIdent, __gong__toRawStringLiteral(slur.Line_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dash_length = %s", slurIdent, __gong__toRawStringLiteral(slur.Dash_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space_length = %s", slurIdent, __gong__toRawStringLiteral(slur.Space_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", slurIdent, __gong__toRawStringLiteral(slur.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", slurIdent, __gong__toRawStringLiteral(slur.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", slurIdent, __gong__toRawStringLiteral(slur.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", slurIdent, __gong__toRawStringLiteral(slur.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", slurIdent, __gong__toRawStringLiteral(slur.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Orientation = %s", slurIdent, __gong__toRawStringLiteral(slur.Orientation)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bezier_x = %s", slurIdent, __gong__toRawStringLiteral(slur.Bezier_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bezier_y = %s", slurIdent, __gong__toRawStringLiteral(slur.Bezier_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bezier_x2 = %s", slurIdent, __gong__toRawStringLiteral(slur.Bezier_x2)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bezier_y2 = %s", slurIdent, __gong__toRawStringLiteral(slur.Bezier_y2)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bezier_offset = %s", slurIdent, __gong__toRawStringLiteral(slur.Bezier_offset)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bezier_offset2 = %s", slurIdent, __gong__toRawStringLiteral(slur.Bezier_offset2)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", slurIdent, __gong__toRawStringLiteral(slur.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", slurIdent, __gong__toRawStringLiteral(slur.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, sound := range __gong__sortStageSetInstances(stageSet.Stage.Sounds, stageSet.Stage.Sound_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			soundIdent := "__models" + sound.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Sound{Name: %s}).Stage(stageSet.Stage)", soundIdent, __gong__toRawStringLiteral(sound.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", soundIdent, __gong__toRawStringLiteral(sound.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tempo = %s", soundIdent, __gong__toRawStringLiteral(sound.Tempo)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dynamics = %s", soundIdent, __gong__toRawStringLiteral(sound.Dynamics)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dacapo = %s", soundIdent, __gong__toRawStringLiteral(string(sound.Dacapo))))
			values.WriteString(fmt.Sprintf("\n\t%s.Segno = %s", soundIdent, __gong__toRawStringLiteral(sound.Segno)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dalsegno = %s", soundIdent, __gong__toRawStringLiteral(sound.Dalsegno)))
			values.WriteString(fmt.Sprintf("\n\t%s.Coda = %s", soundIdent, __gong__toRawStringLiteral(sound.Coda)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tocoda = %s", soundIdent, __gong__toRawStringLiteral(sound.Tocoda)))
			values.WriteString(fmt.Sprintf("\n\t%s.Divisions = %s", soundIdent, __gong__toRawStringLiteral(sound.Divisions)))
			values.WriteString(fmt.Sprintf("\n\t%s.Forward_repeat = %s", soundIdent, __gong__toRawStringLiteral(string(sound.Forward_repeat))))
			values.WriteString(fmt.Sprintf("\n\t%s.Fine = %s", soundIdent, __gong__toRawStringLiteral(sound.Fine)))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_only = %s", soundIdent, __gong__toRawStringLiteral(string(sound.Time_only))))
			values.WriteString(fmt.Sprintf("\n\t%s.Pizzicato = %s", soundIdent, __gong__toRawStringLiteral(string(sound.Pizzicato))))
			values.WriteString(fmt.Sprintf("\n\t%s.Pan = %s", soundIdent, __gong__toRawStringLiteral(sound.Pan)))
			values.WriteString(fmt.Sprintf("\n\t%s.Elevation = %s", soundIdent, __gong__toRawStringLiteral(sound.Elevation)))
			values.WriteString(fmt.Sprintf("\n\t%s.Damper_pedal = %s", soundIdent, __gong__toRawStringLiteral(sound.Damper_pedal)))
			values.WriteString(fmt.Sprintf("\n\t%s.Soft_pedal = %s", soundIdent, __gong__toRawStringLiteral(sound.Soft_pedal)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sostenuto_pedal = %s", soundIdent, __gong__toRawStringLiteral(sound.Sostenuto_pedal)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", soundIdent, __gong__toRawStringLiteral(sound.Id)))
			for _, elem := range sound.Instrument_change {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Instrument_change = append(%s.Instrument_change, %s)", soundIdent, soundIdent, targetIdent))
			}
			for _, elem := range sound.Midi_device {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Midi_device = append(%s.Midi_device, %s)", soundIdent, soundIdent, targetIdent))
			}
			for _, elem := range sound.Midi_instrument {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Midi_instrument = append(%s.Midi_instrument, %s)", soundIdent, soundIdent, targetIdent))
			}
			for _, elem := range sound.Play {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Play = append(%s.Play, %s)", soundIdent, soundIdent, targetIdent))
			}
			if sound.Swing != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + sound.Swing.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Swing = %s", soundIdent, targetIdent))
			}
			if sound.Offset != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + sound.Offset.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", soundIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, staff_details := range __gong__sortStageSetInstances(stageSet.Stage.Staff_detailss, stageSet.Stage.Staff_details_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staff_detailsIdent := "__models" + staff_details.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Staff_details{Name: %s}).Stage(stageSet.Stage)", staff_detailsIdent, __gong__toRawStringLiteral(staff_details.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staff_detailsIdent, __gong__toRawStringLiteral(staff_details.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", staff_detailsIdent, staff_details.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Show_frets = %s", staff_detailsIdent, __gong__toRawStringLiteral(staff_details.Show_frets)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", staff_detailsIdent, __gong__toRawStringLiteral(string(staff_details.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_spacing = %s", staff_detailsIdent, __gong__toRawStringLiteral(string(staff_details.Print_spacing))))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff_type = %s", staff_detailsIdent, __gong__toRawStringLiteral(staff_details.Staff_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff_lines = %d", staff_detailsIdent, staff_details.Staff_lines))
			values.WriteString(fmt.Sprintf("\n\t%s.Capo = %d", staff_detailsIdent, staff_details.Capo))
			for _, elem := range staff_details.Line_detail {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Line_detail = append(%s.Line_detail, %s)", staff_detailsIdent, staff_detailsIdent, targetIdent))
			}
			for _, elem := range staff_details.Staff_tuning {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staff_tuning = append(%s.Staff_tuning, %s)", staff_detailsIdent, staff_detailsIdent, targetIdent))
			}
			if staff_details.Staff_size != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + staff_details.Staff_size.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staff_size = %s", staff_detailsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, staff_divide := range __gong__sortStageSetInstances(stageSet.Stage.Staff_divides, stageSet.Stage.Staff_divide_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staff_divideIdent := "__models" + staff_divide.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Staff_divide{Name: %s}).Stage(stageSet.Stage)", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, staff_layout := range __gong__sortStageSetInstances(stageSet.Stage.Staff_layouts, stageSet.Stage.Staff_layout_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staff_layoutIdent := "__models" + staff_layout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Staff_layout{Name: %s}).Stage(stageSet.Stage)", staff_layoutIdent, __gong__toRawStringLiteral(staff_layout.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staff_layoutIdent, __gong__toRawStringLiteral(staff_layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", staff_layoutIdent, staff_layout.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff_distance = %s", staff_layoutIdent, __gong__toRawStringLiteral(staff_layout.Staff_distance)))
		}
	}
	if stageSet.Stage != nil {
		for _, staff_size := range __gong__sortStageSetInstances(stageSet.Stage.Staff_sizes, stageSet.Stage.Staff_size_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staff_sizeIdent := "__models" + staff_size.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Staff_size{Name: %s}).Stage(stageSet.Stage)", staff_sizeIdent, __gong__toRawStringLiteral(staff_size.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staff_sizeIdent, __gong__toRawStringLiteral(staff_size.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Scaling = %s", staff_sizeIdent, __gong__toRawStringLiteral(staff_size.Scaling)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", staff_sizeIdent, __gong__toRawStringLiteral(staff_size.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, staff_tuning := range __gong__sortStageSetInstances(stageSet.Stage.Staff_tunings, stageSet.Stage.Staff_tuning_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staff_tuningIdent := "__models" + staff_tuning.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Staff_tuning{Name: %s}).Stage(stageSet.Stage)", staff_tuningIdent, __gong__toRawStringLiteral(staff_tuning.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staff_tuningIdent, __gong__toRawStringLiteral(staff_tuning.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", staff_tuningIdent, staff_tuning.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.Tuning_step = %s", staff_tuningIdent, __gong__toRawStringLiteral(string(staff_tuning.Tuning_step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Tuning_alter = %s", staff_tuningIdent, __gong__toRawStringLiteral(staff_tuning.Tuning_alter)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tuning_octave = %d", staff_tuningIdent, staff_tuning.Tuning_octave))
		}
	}
	if stageSet.Stage != nil {
		for _, stem := range __gong__sortStageSetInstances(stageSet.Stage.Stems, stageSet.Stage.Stem_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stemIdent := "__models" + stem.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Stem{Name: %s}).Stage(stageSet.Stage)", stemIdent, __gong__toRawStringLiteral(stem.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stemIdent, __gong__toRawStringLiteral(stem.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", stemIdent, __gong__toRawStringLiteral(stem.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", stemIdent, __gong__toRawStringLiteral(stem.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", stemIdent, __gong__toRawStringLiteral(stem.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", stemIdent, __gong__toRawStringLiteral(stem.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", stemIdent, __gong__toRawStringLiteral(stem.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", stemIdent, __gong__toRawStringLiteral(stem.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, stick := range __gong__sortStageSetInstances(stageSet.Stage.Sticks, stageSet.Stage.Stick_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stickIdent := "__models" + stick.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Stick{Name: %s}).Stage(stageSet.Stage)", stickIdent, __gong__toRawStringLiteral(stick.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stickIdent, __gong__toRawStringLiteral(stick.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tip = %s", stickIdent, __gong__toRawStringLiteral(string(stick.Tip))))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", stickIdent, __gong__toRawStringLiteral(string(stick.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.Dashed_circle = %s", stickIdent, __gong__toRawStringLiteral(string(stick.Dashed_circle))))
			values.WriteString(fmt.Sprintf("\n\t%s.Stick_type = %s", stickIdent, __gong__toRawStringLiteral(stick.Stick_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Stick_material = %s", stickIdent, __gong__toRawStringLiteral(stick.Stick_material)))
		}
	}
	if stageSet.Stage != nil {
		for _, string_mute := range __gong__sortStageSetInstances(stageSet.Stage.String_mutes, stageSet.Stage.String_mute_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			string_muteIdent := "__models" + string_mute.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.String_mute{Name: %s}).Stage(stageSet.Stage)", string_muteIdent, __gong__toRawStringLiteral(string_mute.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", string_muteIdent, __gong__toRawStringLiteral(string_mute.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, string_type := range __gong__sortStageSetInstances(stageSet.Stage.String_types, stageSet.Stage.String_type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			string_typeIdent := "__models" + string_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.String_type{Name: %s}).Stage(stageSet.Stage)", string_typeIdent, __gong__toRawStringLiteral(string_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", string_typeIdent, __gong__toRawStringLiteral(string_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", string_typeIdent, __gong__toRawStringLiteral(string_type.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", string_typeIdent, __gong__toRawStringLiteral(string_type.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", string_typeIdent, __gong__toRawStringLiteral(string_type.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", string_typeIdent, __gong__toRawStringLiteral(string_type.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", string_typeIdent, __gong__toRawStringLiteral(string_type.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", string_typeIdent, __gong__toRawStringLiteral(string_type.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", string_typeIdent, __gong__toRawStringLiteral(string_type.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", string_typeIdent, __gong__toRawStringLiteral(string_type.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", string_typeIdent, __gong__toRawStringLiteral(string_type.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", string_typeIdent, __gong__toRawStringLiteral(string_type.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", string_typeIdent, string_type.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		for _, strong_accent := range __gong__sortStageSetInstances(stageSet.Stage.Strong_accents, stageSet.Stage.Strong_accent_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			strong_accentIdent := "__models" + strong_accent.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Strong_accent{Name: %s}).Stage(stageSet.Stage)", strong_accentIdent, __gong__toRawStringLiteral(strong_accent.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", strong_accentIdent, __gong__toRawStringLiteral(strong_accent.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, style_text := range __gong__sortStageSetInstances(stageSet.Stage.Style_texts, stageSet.Stage.Style_text_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			style_textIdent := "__models" + style_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Style_text{Name: %s}).Stage(stageSet.Stage)", style_textIdent, __gong__toRawStringLiteral(style_text.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", style_textIdent, __gong__toRawStringLiteral(style_text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", style_textIdent, __gong__toRawStringLiteral(style_text.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", style_textIdent, __gong__toRawStringLiteral(style_text.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", style_textIdent, __gong__toRawStringLiteral(style_text.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", style_textIdent, __gong__toRawStringLiteral(style_text.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", style_textIdent, __gong__toRawStringLiteral(style_text.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", style_textIdent, __gong__toRawStringLiteral(style_text.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", style_textIdent, __gong__toRawStringLiteral(style_text.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", style_textIdent, __gong__toRawStringLiteral(style_text.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", style_textIdent, __gong__toRawStringLiteral(style_text.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", style_textIdent, __gong__toRawStringLiteral(style_text.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, supports := range __gong__sortStageSetInstances(stageSet.Stage.Supportss, stageSet.Stage.Supports_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			supportsIdent := "__models" + supports.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Supports{Name: %s}).Stage(stageSet.Stage)", supportsIdent, __gong__toRawStringLiteral(supports.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", supportsIdent, __gong__toRawStringLiteral(supports.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", supportsIdent, __gong__toRawStringLiteral(string(supports.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Element = %s", supportsIdent, __gong__toRawStringLiteral(supports.Element)))
			values.WriteString(fmt.Sprintf("\n\t%s.Attribute = %s", supportsIdent, __gong__toRawStringLiteral(supports.Attribute)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", supportsIdent, __gong__toRawStringLiteral(supports.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, swing := range __gong__sortStageSetInstances(stageSet.Stage.Swings, stageSet.Stage.Swing_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			swingIdent := "__models" + swing.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Swing{Name: %s}).Stage(stageSet.Stage)", swingIdent, __gong__toRawStringLiteral(swing.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", swingIdent, __gong__toRawStringLiteral(swing.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Straight = %s", swingIdent, __gong__toRawStringLiteral(swing.Straight)))
			values.WriteString(fmt.Sprintf("\n\t%s.First = %d", swingIdent, swing.First))
			values.WriteString(fmt.Sprintf("\n\t%s.Second = %d", swingIdent, swing.Second))
			values.WriteString(fmt.Sprintf("\n\t%s.Swing_type = %s", swingIdent, __gong__toRawStringLiteral(string(swing.Swing_type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Swing_style = %s", swingIdent, __gong__toRawStringLiteral(swing.Swing_style)))
		}
	}
	if stageSet.Stage != nil {
		for _, sync := range __gong__sortStageSetInstances(stageSet.Stage.Syncs, stageSet.Stage.Sync_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			syncIdent := "__models" + sync.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Sync{Name: %s}).Stage(stageSet.Stage)", syncIdent, __gong__toRawStringLiteral(sync.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", syncIdent, __gong__toRawStringLiteral(sync.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", syncIdent, __gong__toRawStringLiteral(sync.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Latency = %d", syncIdent, sync.Latency))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", syncIdent, __gong__toRawStringLiteral(sync.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_only = %s", syncIdent, __gong__toRawStringLiteral(string(sync.Time_only))))
		}
	}
	if stageSet.Stage != nil {
		for _, system_dividers := range __gong__sortStageSetInstances(stageSet.Stage.System_dividerss, stageSet.Stage.System_dividers_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			system_dividersIdent := "__models" + system_dividers.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.System_dividers{Name: %s}).Stage(stageSet.Stage)", system_dividersIdent, __gong__toRawStringLiteral(system_dividers.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", system_dividersIdent, __gong__toRawStringLiteral(system_dividers.Name)))
			if system_dividers.Left_divider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + system_dividers.Left_divider.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Left_divider = %s", system_dividersIdent, targetIdent))
			}
			if system_dividers.Right_divider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + system_dividers.Right_divider.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Right_divider = %s", system_dividersIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, system_layout := range __gong__sortStageSetInstances(stageSet.Stage.System_layouts, stageSet.Stage.System_layout_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			system_layoutIdent := "__models" + system_layout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.System_layout{Name: %s}).Stage(stageSet.Stage)", system_layoutIdent, __gong__toRawStringLiteral(system_layout.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", system_layoutIdent, __gong__toRawStringLiteral(system_layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.System_distance = %s", system_layoutIdent, __gong__toRawStringLiteral(system_layout.System_distance)))
			values.WriteString(fmt.Sprintf("\n\t%s.Top_system_distance = %s", system_layoutIdent, __gong__toRawStringLiteral(system_layout.Top_system_distance)))
			if system_layout.System_margins != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + system_layout.System_margins.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System_margins = %s", system_layoutIdent, targetIdent))
			}
			if system_layout.System_dividers != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + system_layout.System_dividers.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System_dividers = %s", system_layoutIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, system_margins := range __gong__sortStageSetInstances(stageSet.Stage.System_marginss, stageSet.Stage.System_margins_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			system_marginsIdent := "__models" + system_margins.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.System_margins{Name: %s}).Stage(stageSet.Stage)", system_marginsIdent, __gong__toRawStringLiteral(system_margins.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", system_marginsIdent, __gong__toRawStringLiteral(system_margins.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Left_margin = %s", system_marginsIdent, __gong__toRawStringLiteral(system_margins.Left_margin)))
			values.WriteString(fmt.Sprintf("\n\t%s.Right_margin = %s", system_marginsIdent, __gong__toRawStringLiteral(system_margins.Right_margin)))
		}
	}
	if stageSet.Stage != nil {
		for _, tap := range __gong__sortStageSetInstances(stageSet.Stage.Taps, stageSet.Stage.Tap_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tapIdent := "__models" + tap.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tap{Name: %s}).Stage(stageSet.Stage)", tapIdent, __gong__toRawStringLiteral(tap.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tapIdent, __gong__toRawStringLiteral(tap.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Hand = %s", tapIdent, __gong__toRawStringLiteral(tap.Hand)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", tapIdent, __gong__toRawStringLiteral(tap.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", tapIdent, __gong__toRawStringLiteral(tap.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", tapIdent, __gong__toRawStringLiteral(tap.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", tapIdent, __gong__toRawStringLiteral(tap.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", tapIdent, __gong__toRawStringLiteral(tap.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", tapIdent, __gong__toRawStringLiteral(tap.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", tapIdent, __gong__toRawStringLiteral(tap.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", tapIdent, __gong__toRawStringLiteral(tap.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tapIdent, __gong__toRawStringLiteral(tap.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", tapIdent, __gong__toRawStringLiteral(tap.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", tapIdent, __gong__toRawStringLiteral(tap.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, technical := range __gong__sortStageSetInstances(stageSet.Stage.Technicals, stageSet.Stage.Technical_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			technicalIdent := "__models" + technical.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Technical{Name: %s}).Stage(stageSet.Stage)", technicalIdent, __gong__toRawStringLiteral(technical.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", technicalIdent, __gong__toRawStringLiteral(technical.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", technicalIdent, __gong__toRawStringLiteral(technical.Id)))
			for _, elem := range technical.Up_bow {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Up_bow = append(%s.Up_bow, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Down_bow {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Down_bow = append(%s.Down_bow, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Harmonic {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmonic = append(%s.Harmonic, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Open_string {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Open_string = append(%s.Open_string, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Thumb_position {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Thumb_position = append(%s.Thumb_position, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Fingering {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fingering = append(%s.Fingering, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Pluck {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pluck = append(%s.Pluck, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Double_tongue {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Double_tongue = append(%s.Double_tongue, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Triple_tongue {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Triple_tongue = append(%s.Triple_tongue, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Stopped {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stopped = append(%s.Stopped, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Snap_pizzicato {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Snap_pizzicato = append(%s.Snap_pizzicato, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Fret {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fret = append(%s.Fret, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.String {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.String = append(%s.String, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Hammer_on {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Hammer_on = append(%s.Hammer_on, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Pull_off {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pull_off = append(%s.Pull_off, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Bend {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bend = append(%s.Bend, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Tap {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tap = append(%s.Tap, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Heel {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Heel = append(%s.Heel, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Toe {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Toe = append(%s.Toe, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Fingernails {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fingernails = append(%s.Fingernails, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Hole {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Hole = append(%s.Hole, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Arrow {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arrow = append(%s.Arrow, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Handbell {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Handbell = append(%s.Handbell, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Brass_bend {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Brass_bend = append(%s.Brass_bend, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Flip {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Flip = append(%s.Flip, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Smear {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Smear = append(%s.Smear, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Open {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Open = append(%s.Open, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Half_muted {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Half_muted = append(%s.Half_muted, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Harmon_mute {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmon_mute = append(%s.Harmon_mute, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Golpe {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Golpe = append(%s.Golpe, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Other_technical {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_technical = append(%s.Other_technical, %s)", technicalIdent, technicalIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, text_element_data := range __gong__sortStageSetInstances(stageSet.Stage.Text_element_datas, stageSet.Stage.Text_element_data_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			text_element_dataIdent := "__models" + text_element_data.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Text_element_data{Name: %s}).Stage(stageSet.Stage)", text_element_dataIdent, __gong__toRawStringLiteral(text_element_data.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", text_element_dataIdent, __gong__toRawStringLiteral(text_element_data.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", text_element_dataIdent, __gong__toRawStringLiteral(text_element_data.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", text_element_dataIdent, __gong__toRawStringLiteral(text_element_data.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", text_element_dataIdent, __gong__toRawStringLiteral(text_element_data.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", text_element_dataIdent, __gong__toRawStringLiteral(text_element_data.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", text_element_dataIdent, __gong__toRawStringLiteral(text_element_data.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", text_element_dataIdent, __gong__toRawStringLiteral(text_element_data.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", text_element_dataIdent, text_element_data.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", text_element_dataIdent, text_element_data.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_through = %d", text_element_dataIdent, text_element_data.Line_through))
			values.WriteString(fmt.Sprintf("\n\t%s.Rotation = %s", text_element_dataIdent, __gong__toRawStringLiteral(text_element_data.Rotation)))
			values.WriteString(fmt.Sprintf("\n\t%s.Letter_spacing = %s", text_element_dataIdent, __gong__toRawStringLiteral(text_element_data.Letter_spacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dir = %s", text_element_dataIdent, __gong__toRawStringLiteral(text_element_data.Dir)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", text_element_dataIdent, __gong__toRawStringLiteral(text_element_data.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, tie := range __gong__sortStageSetInstances(stageSet.Stage.Ties, stageSet.Stage.Tie_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tieIdent := "__models" + tie.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tie{Name: %s}).Stage(stageSet.Stage)", tieIdent, __gong__toRawStringLiteral(tie.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tieIdent, __gong__toRawStringLiteral(tie.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", tieIdent, __gong__toRawStringLiteral(string(tie.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_only = %s", tieIdent, __gong__toRawStringLiteral(string(tie.Time_only))))
		}
	}
	if stageSet.Stage != nil {
		for _, tied := range __gong__sortStageSetInstances(stageSet.Stage.Tieds, stageSet.Stage.Tied_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tiedIdent := "__models" + tied.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tied{Name: %s}).Stage(stageSet.Stage)", tiedIdent, __gong__toRawStringLiteral(tied.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tiedIdent, __gong__toRawStringLiteral(tied.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", tiedIdent, __gong__toRawStringLiteral(tied.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", tiedIdent, tied.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_type = %s", tiedIdent, __gong__toRawStringLiteral(tied.Line_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dash_length = %s", tiedIdent, __gong__toRawStringLiteral(tied.Dash_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space_length = %s", tiedIdent, __gong__toRawStringLiteral(tied.Space_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", tiedIdent, __gong__toRawStringLiteral(tied.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", tiedIdent, __gong__toRawStringLiteral(tied.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", tiedIdent, __gong__toRawStringLiteral(tied.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", tiedIdent, __gong__toRawStringLiteral(tied.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", tiedIdent, __gong__toRawStringLiteral(tied.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Orientation = %s", tiedIdent, __gong__toRawStringLiteral(tied.Orientation)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bezier_x = %s", tiedIdent, __gong__toRawStringLiteral(tied.Bezier_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bezier_y = %s", tiedIdent, __gong__toRawStringLiteral(tied.Bezier_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bezier_x2 = %s", tiedIdent, __gong__toRawStringLiteral(tied.Bezier_x2)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bezier_y2 = %s", tiedIdent, __gong__toRawStringLiteral(tied.Bezier_y2)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bezier_offset = %s", tiedIdent, __gong__toRawStringLiteral(tied.Bezier_offset)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bezier_offset2 = %s", tiedIdent, __gong__toRawStringLiteral(tied.Bezier_offset2)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tiedIdent, __gong__toRawStringLiteral(tied.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", tiedIdent, __gong__toRawStringLiteral(tied.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, time := range __gong__sortStageSetInstances(stageSet.Stage.Times, stageSet.Stage.Time_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			timeIdent := "__models" + time.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Time{Name: %s}).Stage(stageSet.Stage)", timeIdent, __gong__toRawStringLiteral(time.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timeIdent, __gong__toRawStringLiteral(time.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", timeIdent, time.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Symbol = %s", timeIdent, __gong__toRawStringLiteral(string(time.Symbol))))
			values.WriteString(fmt.Sprintf("\n\t%s.Separator = %s", timeIdent, __gong__toRawStringLiteral(string(time.Separator))))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", timeIdent, __gong__toRawStringLiteral(time.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", timeIdent, __gong__toRawStringLiteral(time.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", timeIdent, __gong__toRawStringLiteral(time.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", timeIdent, __gong__toRawStringLiteral(time.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", timeIdent, __gong__toRawStringLiteral(time.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", timeIdent, __gong__toRawStringLiteral(time.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", timeIdent, __gong__toRawStringLiteral(time.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", timeIdent, __gong__toRawStringLiteral(time.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", timeIdent, __gong__toRawStringLiteral(time.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", timeIdent, __gong__toRawStringLiteral(time.Halign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", timeIdent, __gong__toRawStringLiteral(time.Valign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", timeIdent, __gong__toRawStringLiteral(string(time.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", timeIdent, __gong__toRawStringLiteral(time.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", timeIdent, __gong__toRawStringLiteral(time.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.Beat_type = %s", timeIdent, __gong__toRawStringLiteral(time.Beat_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Senza_misura = %s", timeIdent, __gong__toRawStringLiteral(time.Senza_misura)))
			if time.Interchangeable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + time.Interchangeable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Interchangeable = %s", timeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, time_modification := range __gong__sortStageSetInstances(stageSet.Stage.Time_modifications, stageSet.Stage.Time_modification_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			time_modificationIdent := "__models" + time_modification.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Time_modification{Name: %s}).Stage(stageSet.Stage)", time_modificationIdent, __gong__toRawStringLiteral(time_modification.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", time_modificationIdent, __gong__toRawStringLiteral(time_modification.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actual_notes = %d", time_modificationIdent, time_modification.Actual_notes))
			values.WriteString(fmt.Sprintf("\n\t%s.Normal_notes = %d", time_modificationIdent, time_modification.Normal_notes))
			values.WriteString(fmt.Sprintf("\n\t%s.Normal_type = %s", time_modificationIdent, __gong__toRawStringLiteral(string(time_modification.Normal_type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Normal_dot = %s", time_modificationIdent, __gong__toRawStringLiteral(time_modification.Normal_dot)))
		}
	}
	if stageSet.Stage != nil {
		for _, timpani := range __gong__sortStageSetInstances(stageSet.Stage.Timpanis, stageSet.Stage.Timpani_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			timpaniIdent := "__models" + timpani.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Timpani{Name: %s}).Stage(stageSet.Stage)", timpaniIdent, __gong__toRawStringLiteral(timpani.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timpaniIdent, __gong__toRawStringLiteral(timpani.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", timpaniIdent, __gong__toRawStringLiteral(timpani.Smufl)))
		}
	}
	if stageSet.Stage != nil {
		for _, transpose := range __gong__sortStageSetInstances(stageSet.Stage.Transposes, stageSet.Stage.Transpose_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			transposeIdent := "__models" + transpose.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Transpose{Name: %s}).Stage(stageSet.Stage)", transposeIdent, __gong__toRawStringLiteral(transpose.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", transposeIdent, __gong__toRawStringLiteral(transpose.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", transposeIdent, transpose.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", transposeIdent, __gong__toRawStringLiteral(transpose.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Diatonic = %d", transposeIdent, transpose.Diatonic))
			values.WriteString(fmt.Sprintf("\n\t%s.Chromatic = %s", transposeIdent, __gong__toRawStringLiteral(transpose.Chromatic)))
			values.WriteString(fmt.Sprintf("\n\t%s.Octave_change = %d", transposeIdent, transpose.Octave_change))
			values.WriteString(fmt.Sprintf("\n\t%s.Double = %f", transposeIdent, transpose.Double))
		}
	}
	if stageSet.Stage != nil {
		for _, tremolo := range __gong__sortStageSetInstances(stageSet.Stage.Tremolos, stageSet.Stage.Tremolo_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tremoloIdent := "__models" + tremolo.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tremolo{Name: %s}).Stage(stageSet.Stage)", tremoloIdent, __gong__toRawStringLiteral(tremolo.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", tremoloIdent, tremolo.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		for _, tuplet := range __gong__sortStageSetInstances(stageSet.Stage.Tuplets, stageSet.Stage.Tuplet_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tupletIdent := "__models" + tuplet.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tuplet{Name: %s}).Stage(stageSet.Stage)", tupletIdent, __gong__toRawStringLiteral(tuplet.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", tupletIdent, __gong__toRawStringLiteral(string(tuplet.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", tupletIdent, tuplet.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Bracket = %s", tupletIdent, __gong__toRawStringLiteral(string(tuplet.Bracket))))
			values.WriteString(fmt.Sprintf("\n\t%s.Show_number = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.Show_number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show_type = %s", tupletIdent, __gong__toRawStringLiteral(string(tuplet.Show_type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_shape = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.Line_shape)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.Id)))
			if tuplet.Tuplet_actual != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tuplet.Tuplet_actual.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet_actual = %s", tupletIdent, targetIdent))
			}
			if tuplet.Tuplet_normal != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tuplet.Tuplet_normal.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet_normal = %s", tupletIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, tuplet_dot := range __gong__sortStageSetInstances(stageSet.Stage.Tuplet_dots, stageSet.Stage.Tuplet_dot_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tuplet_dotIdent := "__models" + tuplet_dot.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tuplet_dot{Name: %s}).Stage(stageSet.Stage)", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Color)))
		}
	}
	if stageSet.Stage != nil {
		for _, tuplet_number := range __gong__sortStageSetInstances(stageSet.Stage.Tuplet_numbers, stageSet.Stage.Tuplet_number_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tuplet_numberIdent := "__models" + tuplet_number.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tuplet_number{Name: %s}).Stage(stageSet.Stage)", tuplet_numberIdent, __gong__toRawStringLiteral(tuplet_number.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tuplet_numberIdent, __gong__toRawStringLiteral(tuplet_number.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", tuplet_numberIdent, __gong__toRawStringLiteral(tuplet_number.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", tuplet_numberIdent, __gong__toRawStringLiteral(tuplet_number.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", tuplet_numberIdent, __gong__toRawStringLiteral(tuplet_number.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", tuplet_numberIdent, __gong__toRawStringLiteral(tuplet_number.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tuplet_numberIdent, __gong__toRawStringLiteral(tuplet_number.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", tuplet_numberIdent, tuplet_number.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		for _, tuplet_portion := range __gong__sortStageSetInstances(stageSet.Stage.Tuplet_portions, stageSet.Stage.Tuplet_portion_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tuplet_portionIdent := "__models" + tuplet_portion.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tuplet_portion{Name: %s}).Stage(stageSet.Stage)", tuplet_portionIdent, __gong__toRawStringLiteral(tuplet_portion.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tuplet_portionIdent, __gong__toRawStringLiteral(tuplet_portion.Name)))
			if tuplet_portion.Tuplet_number != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tuplet_portion.Tuplet_number.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet_number = %s", tuplet_portionIdent, targetIdent))
			}
			if tuplet_portion.Tuplet_type != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tuplet_portion.Tuplet_type.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet_type = %s", tuplet_portionIdent, targetIdent))
			}
			for _, elem := range tuplet_portion.Tuplet_dot {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet_dot = append(%s.Tuplet_dot, %s)", tuplet_portionIdent, tuplet_portionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, tuplet_type := range __gong__sortStageSetInstances(stageSet.Stage.Tuplet_types, stageSet.Stage.Tuplet_type_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tuplet_typeIdent := "__models" + tuplet_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tuplet_type{Name: %s}).Stage(stageSet.Stage)", tuplet_typeIdent, __gong__toRawStringLiteral(tuplet_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tuplet_typeIdent, __gong__toRawStringLiteral(tuplet_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", tuplet_typeIdent, __gong__toRawStringLiteral(tuplet_type.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", tuplet_typeIdent, __gong__toRawStringLiteral(tuplet_type.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", tuplet_typeIdent, __gong__toRawStringLiteral(tuplet_type.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", tuplet_typeIdent, __gong__toRawStringLiteral(tuplet_type.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tuplet_typeIdent, __gong__toRawStringLiteral(tuplet_type.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", tuplet_typeIdent, __gong__toRawStringLiteral(string(tuplet_type.EnclosedText))))
		}
	}
	if stageSet.Stage != nil {
		for _, typed_text := range __gong__sortStageSetInstances(stageSet.Stage.Typed_texts, stageSet.Stage.Typed_text_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			typed_textIdent := "__models" + typed_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Typed_text{Name: %s}).Stage(stageSet.Stage)", typed_textIdent, __gong__toRawStringLiteral(typed_text.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", typed_textIdent, __gong__toRawStringLiteral(typed_text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", typed_textIdent, __gong__toRawStringLiteral(typed_text.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", typed_textIdent, __gong__toRawStringLiteral(typed_text.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, unpitched := range __gong__sortStageSetInstances(stageSet.Stage.Unpitcheds, stageSet.Stage.Unpitched_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unpitchedIdent := "__models" + unpitched.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Unpitched{Name: %s}).Stage(stageSet.Stage)", unpitchedIdent, __gong__toRawStringLiteral(unpitched.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unpitchedIdent, __gong__toRawStringLiteral(unpitched.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Display_step = %s", unpitchedIdent, __gong__toRawStringLiteral(string(unpitched.Display_step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Display_octave = %d", unpitchedIdent, unpitched.Display_octave))
		}
	}
	if stageSet.Stage != nil {
		for _, virtual_instrument := range __gong__sortStageSetInstances(stageSet.Stage.Virtual_instruments, stageSet.Stage.Virtual_instrument_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			virtual_instrumentIdent := "__models" + virtual_instrument.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Virtual_instrument{Name: %s}).Stage(stageSet.Stage)", virtual_instrumentIdent, __gong__toRawStringLiteral(virtual_instrument.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", virtual_instrumentIdent, __gong__toRawStringLiteral(virtual_instrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Virtual_library = %s", virtual_instrumentIdent, __gong__toRawStringLiteral(virtual_instrument.Virtual_library)))
			values.WriteString(fmt.Sprintf("\n\t%s.Virtual_name = %s", virtual_instrumentIdent, __gong__toRawStringLiteral(virtual_instrument.Virtual_name)))
		}
	}
	if stageSet.Stage != nil {
		for _, wait := range __gong__sortStageSetInstances(stageSet.Stage.Waits, stageSet.Stage.Wait_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			waitIdent := "__models" + wait.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Wait{Name: %s}).Stage(stageSet.Stage)", waitIdent, __gong__toRawStringLiteral(wait.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", waitIdent, __gong__toRawStringLiteral(wait.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", waitIdent, __gong__toRawStringLiteral(wait.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_only = %s", waitIdent, __gong__toRawStringLiteral(string(wait.Time_only))))
		}
	}
	if stageSet.Stage != nil {
		for _, wavy_line := range __gong__sortStageSetInstances(stageSet.Stage.Wavy_lines, stageSet.Stage.Wavy_line_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			wavy_lineIdent := "__models" + wavy_line.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Wavy_line{Name: %s}).Stage(stageSet.Stage)", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", wavy_lineIdent, __gong__toRawStringLiteral(string(wavy_line.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", wavy_lineIdent, wavy_line.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Start_note = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Start_note)))
			values.WriteString(fmt.Sprintf("\n\t%s.Trill_step = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Trill_step)))
			values.WriteString(fmt.Sprintf("\n\t%s.Two_note_turn = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Two_note_turn)))
			values.WriteString(fmt.Sprintf("\n\t%s.Accelerate = %s", wavy_lineIdent, __gong__toRawStringLiteral(string(wavy_line.Accelerate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.Second_beat = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Second_beat)))
			values.WriteString(fmt.Sprintf("\n\t%s.Last_beat = %s", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Last_beat)))
		}
	}
	if stageSet.Stage != nil {
		for _, wedge := range __gong__sortStageSetInstances(stageSet.Stage.Wedges, stageSet.Stage.Wedge_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			wedgeIdent := "__models" + wedge.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Wedge{Name: %s}).Stage(stageSet.Stage)", wedgeIdent, __gong__toRawStringLiteral(wedge.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", wedgeIdent, wedge.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Spread = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Spread)))
			values.WriteString(fmt.Sprintf("\n\t%s.Niente = %s", wedgeIdent, __gong__toRawStringLiteral(string(wedge.Niente))))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_type = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Line_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dash_length = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Dash_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space_length = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Space_length)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_x = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Default_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Default_y = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Default_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_x = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Relative_x)))
			values.WriteString(fmt.Sprintf("\n\t%s.Relative_y = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Relative_y)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, wood := range __gong__sortStageSetInstances(stageSet.Stage.Woods, stageSet.Stage.Wood_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			woodIdent := "__models" + wood.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Wood{Name: %s}).Stage(stageSet.Stage)", woodIdent, __gong__toRawStringLiteral(wood.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", woodIdent, __gong__toRawStringLiteral(wood.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", woodIdent, __gong__toRawStringLiteral(wood.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", woodIdent, __gong__toRawStringLiteral(wood.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		for _, work := range __gong__sortStageSetInstances(stageSet.Stage.Works, stageSet.Stage.Work_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			workIdent := "__models" + work.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Work{Name: %s}).Stage(stageSet.Stage)", workIdent, __gong__toRawStringLiteral(work.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", workIdent, __gong__toRawStringLiteral(work.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Work_number = %s", workIdent, __gong__toRawStringLiteral(work.Work_number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Work_title = %s", workIdent, __gong__toRawStringLiteral(work.Work_title)))
			if work.Opus != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + work.Opus.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Opus = %s", workIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models"
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
		case "github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models":
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
				case "A_directive":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_directive), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_measure":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_measure), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_measure_1":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_measure_1), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_part":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_part), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "A_part_1":
					identifierMap[ident.Name] = __gong__stageSetInit(new(A_part_1), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Accidental":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Accidental), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Accidental_mark":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Accidental_mark), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Accidental_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Accidental_text), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Accord":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Accord), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Accordion_registration":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Accordion_registration), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Appearance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Appearance), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Arpeggiate":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Arpeggiate), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Arrow":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Arrow), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Articulations":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Articulations), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Assess":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Assess), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Attributes":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Attributes), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Backup":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Backup), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Bar_style_color":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bar_style_color), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Barline":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Barline), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Barre":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Barre), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Bass":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bass), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Bass_step":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bass_step), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Beam":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Beam), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Beat_repeat":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Beat_repeat), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Beat_unit_tied":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Beat_unit_tied), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Beater":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Beater), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Bend":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bend), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Bookmark":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bookmark), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Bracket":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bracket), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Breath_mark":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Breath_mark), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Caesura":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Caesura), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Cancel":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Cancel), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Clef":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Clef), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Coda":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Coda), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Credit":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Credit), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Dashes":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Dashes), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Defaults":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Defaults), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Degree":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Degree), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Degree_alter":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Degree_alter), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Degree_type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Degree_type), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Degree_value":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Degree_value), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Direction":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Direction), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Direction_type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Direction_type), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Distance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Distance), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Double":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Double), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Dynamics":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Dynamics), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Effect":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Effect), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Elision":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Elision), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Empty":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Empty_font":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_font), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Empty_line":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_line), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Empty_placement":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_placement), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Empty_placement_smufl":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_placement_smufl), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Empty_print_object_style_align":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_print_object_style_align), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Empty_print_style":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_print_style), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Empty_print_style_align":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_print_style_align), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Empty_print_style_align_id":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_print_style_align_id), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Empty_trill_sound":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty_trill_sound), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Encoding":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Encoding), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Ending":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Ending), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Extend":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Extend), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Feature":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Feature), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Fermata":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Fermata), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Figure":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Figure), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Figured_bass":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Figured_bass), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Fingering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Fingering), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "First_fret":
					identifierMap[ident.Name] = __gong__stageSetInit(new(First_fret), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "For_part":
					identifierMap[ident.Name] = __gong__stageSetInit(new(For_part), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Formatted_symbol":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Formatted_symbol), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Formatted_symbol_id":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Formatted_symbol_id), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Formatted_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Formatted_text), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Formatted_text_id":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Formatted_text_id), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Forward":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Forward), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Frame":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Frame), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Frame_note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Frame_note), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Fret":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Fret), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Glass":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Glass), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Glissando":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Glissando), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Glyph":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Glyph), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Grace":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Grace), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Group_barline":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Group_barline), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Group_name":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Group_name), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Group_symbol":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Group_symbol), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Grouping":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Grouping), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Hammer_on_pull_off":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Hammer_on_pull_off), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Handbell":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Handbell), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Harmon_closed":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harmon_closed), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Harmon_mute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harmon_mute), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Harmonic":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harmonic), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Harmony":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harmony), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Harmony_alter":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harmony_alter), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Harp_pedals":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harp_pedals), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Heel_toe":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Heel_toe), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Hole":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Hole), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Hole_closed":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Hole_closed), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Horizontal_turn":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Horizontal_turn), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Identification":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Identification), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Image":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Image), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Instrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Instrument), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Instrument_change":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Instrument_change), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Instrument_link":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Instrument_link), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Interchangeable":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Interchangeable), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Inversion":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Inversion), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Key":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Key), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Key_accidental":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Key_accidental), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Key_octave":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Key_octave), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Kind":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Kind), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Level":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Level), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Line_detail":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Line_detail), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Line_width":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Line_width), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Link":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Link), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Listen":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Listen), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Listening":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Listening), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Lyric":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Lyric), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Lyric_font":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Lyric_font), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Lyric_language":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Lyric_language), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Measure_layout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Measure_layout), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Measure_numbering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Measure_numbering), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Measure_repeat":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Measure_repeat), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Measure_style":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Measure_style), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Membrane":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Membrane), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Metal":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metal), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Metronome":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metronome), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Metronome_beam":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metronome_beam), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Metronome_note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metronome_note), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Metronome_tied":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metronome_tied), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Metronome_tuplet":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metronome_tuplet), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Midi_device":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Midi_device), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Midi_instrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Midi_instrument), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Miscellaneous":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Miscellaneous), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Miscellaneous_field":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Miscellaneous_field), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Mordent":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Mordent), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Multiple_rest":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Multiple_rest), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Name_display":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Name_display), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Non_arpeggiate":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Non_arpeggiate), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Notations":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Notations), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Note_size":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note_size), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Note_type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note_type), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Notehead":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Notehead), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Notehead_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Notehead_text), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Numeral":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Numeral), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Numeral_key":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Numeral_key), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Numeral_root":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Numeral_root), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Octave_shift":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Octave_shift), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Offset":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Offset), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Opus":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Opus), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Ornaments":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Ornaments), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Other_appearance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_appearance), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Other_direction":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_direction), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Other_listening":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_listening), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Other_notation":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_notation), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Other_placement_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_placement_text), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Other_play":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_play), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Other_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Other_text), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Page_layout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Page_layout), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Page_margins":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Page_margins), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Part_clef":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_clef), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Part_group":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_group), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Part_link":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_link), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Part_list":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_list), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Part_name":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_name), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Part_symbol":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_symbol), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Part_transpose":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part_transpose), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Pedal":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Pedal), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Pedal_tuning":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Pedal_tuning), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Per_minute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Per_minute), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Percussion":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Percussion), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Pitch":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Pitch), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Pitched":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Pitched), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Placement_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Placement_text), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Play":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Play), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Player":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Player), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Principal_voice":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Principal_voice), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Print":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Print), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Release":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Release), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Repeat":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Repeat), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Rest":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Rest), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Root":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Root), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Root_step":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Root_step), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Scaling":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Scaling), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Scordatura":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Scordatura), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Score_instrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Score_instrument), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Score_part":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Score_part), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Score_partwise":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Score_partwise), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Score_timewise":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Score_timewise), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Segno":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Segno), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Slash":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Slash), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Slide":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Slide), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Slur":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Slur), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Sound":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Sound), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Staff_details":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Staff_details), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Staff_divide":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Staff_divide), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Staff_layout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Staff_layout), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Staff_size":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Staff_size), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Staff_tuning":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Staff_tuning), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Stem":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Stem), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Stick":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Stick), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "String_mute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(String_mute), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "String_type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(String_type), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Strong_accent":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Strong_accent), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Style_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Style_text), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Supports":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Supports), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Swing":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Swing), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Sync":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Sync), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "System_dividers":
					identifierMap[ident.Name] = __gong__stageSetInit(new(System_dividers), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "System_layout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(System_layout), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "System_margins":
					identifierMap[ident.Name] = __gong__stageSetInit(new(System_margins), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tap":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tap), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Technical":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Technical), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Text_element_data":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Text_element_data), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tie":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tie), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tied":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tied), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Time":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Time), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Time_modification":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Time_modification), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Timpani":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Timpani), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Transpose":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Transpose), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tremolo":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tremolo), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tuplet":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tuplet), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tuplet_dot":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tuplet_dot), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tuplet_number":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tuplet_number), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tuplet_portion":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tuplet_portion), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tuplet_type":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tuplet_type), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Typed_text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Typed_text), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Unpitched":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Unpitched), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Virtual_instrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Virtual_instrument), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Wait":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Wait), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Wavy_line":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Wavy_line), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Wedge":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Wedge), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Wood":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Wood), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Work":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Work), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *A_directive:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *A_measure:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Implicit":
						inst.Implicit = Enum_Yes_no(GongExtractString(rhs))
					case "Non_controlling":
						inst.Non_controlling = Enum_Yes_no(GongExtractString(rhs))
					case "Width":
						inst.Width = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Note":
						__gong__assignSliceOfPointers(&inst.Note, rhs, identifierMap)
					case "Backup":
						__gong__assignSliceOfPointers(&inst.Backup, rhs, identifierMap)
					case "Forward":
						__gong__assignSliceOfPointers(&inst.Forward, rhs, identifierMap)
					case "Direction":
						__gong__assignSliceOfPointers(&inst.Direction, rhs, identifierMap)
					case "Attributes":
						__gong__assignSliceOfPointers(&inst.Attributes, rhs, identifierMap)
					case "Harmony":
						__gong__assignSliceOfPointers(&inst.Harmony, rhs, identifierMap)
					case "Figured_bass":
						__gong__assignSliceOfPointers(&inst.Figured_bass, rhs, identifierMap)
					case "Print":
						__gong__assignSliceOfPointers(&inst.Print, rhs, identifierMap)
					case "Sound":
						__gong__assignSliceOfPointers(&inst.Sound, rhs, identifierMap)
					case "Listening":
						__gong__assignSliceOfPointers(&inst.Listening, rhs, identifierMap)
					case "Barline":
						__gong__assignSliceOfPointers(&inst.Barline, rhs, identifierMap)
					case "Grouping":
						__gong__assignSliceOfPointers(&inst.Grouping, rhs, identifierMap)
					case "Link":
						__gong__assignSliceOfPointers(&inst.Link, rhs, identifierMap)
					case "Bookmark":
						__gong__assignSliceOfPointers(&inst.Bookmark, rhs, identifierMap)
					}
				case *A_measure_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Implicit":
						inst.Implicit = Enum_Yes_no(GongExtractString(rhs))
					case "Non_controlling":
						inst.Non_controlling = Enum_Yes_no(GongExtractString(rhs))
					case "Width":
						inst.Width = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Part":
						__gong__assignSliceOfPointers(&inst.Part, rhs, identifierMap)
					}
				case *A_part:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Measure":
						__gong__assignSliceOfPointers(&inst.Measure, rhs, identifierMap)
					}
				case *A_part_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Note":
						__gong__assignSliceOfPointers(&inst.Note, rhs, identifierMap)
					case "Backup":
						__gong__assignSliceOfPointers(&inst.Backup, rhs, identifierMap)
					case "Forward":
						__gong__assignSliceOfPointers(&inst.Forward, rhs, identifierMap)
					case "Direction":
						__gong__assignSliceOfPointers(&inst.Direction, rhs, identifierMap)
					case "Attributes":
						__gong__assignSliceOfPointers(&inst.Attributes, rhs, identifierMap)
					case "Harmony":
						__gong__assignSliceOfPointers(&inst.Harmony, rhs, identifierMap)
					case "Figured_bass":
						__gong__assignSliceOfPointers(&inst.Figured_bass, rhs, identifierMap)
					case "Print":
						__gong__assignSliceOfPointers(&inst.Print, rhs, identifierMap)
					case "Sound":
						__gong__assignSliceOfPointers(&inst.Sound, rhs, identifierMap)
					case "Listening":
						__gong__assignSliceOfPointers(&inst.Listening, rhs, identifierMap)
					case "Barline":
						__gong__assignSliceOfPointers(&inst.Barline, rhs, identifierMap)
					case "Grouping":
						__gong__assignSliceOfPointers(&inst.Grouping, rhs, identifierMap)
					case "Link":
						__gong__assignSliceOfPointers(&inst.Link, rhs, identifierMap)
					case "Bookmark":
						__gong__assignSliceOfPointers(&inst.Bookmark, rhs, identifierMap)
					}
				case *Accidental:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Cautionary":
						inst.Cautionary = GongExtractString(rhs)
					case "Editorial":
						inst.Editorial = Enum_Yes_no(GongExtractString(rhs))
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Parentheses":
						inst.Parentheses = Enum_Yes_no(GongExtractString(rhs))
					case "Bracket":
						inst.Bracket = Enum_Yes_no(GongExtractString(rhs))
					case "Size":
						inst.Size = Enum_Symbol_size(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Accidental_mark:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Parentheses":
						inst.Parentheses = Enum_Yes_no(GongExtractString(rhs))
					case "Bracket":
						inst.Bracket = Enum_Yes_no(GongExtractString(rhs))
					case "Size":
						inst.Size = Enum_Symbol_size(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = Enum_Accidental_value(GongExtractString(rhs))
					}
				case *Accidental_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Space":
						inst.Space = GongExtractString(rhs)
					case "Justify":
						inst.Justify = Enum_Left_center_right(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Rotation":
						inst.Rotation = GongExtractString(rhs)
					case "Letter_spacing":
						inst.Letter_spacing = GongExtractString(rhs)
					case "Line_height":
						inst.Line_height = GongExtractString(rhs)
					case "Dir":
						inst.Dir = GongExtractString(rhs)
					case "Enclosure":
						inst.Enclosure = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = Enum_Accidental_value(GongExtractString(rhs))
					}
				case *Accord:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "String":
						inst.String = GongExtractInt(rhs)
					case "Tuning_step":
						inst.Tuning_step = Enum_Step(GongExtractString(rhs))
					case "Tuning_alter":
						inst.Tuning_alter = GongExtractString(rhs)
					case "Tuning_octave":
						inst.Tuning_octave = GongExtractInt(rhs)
					}
				case *Accordion_registration:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Accordion_high":
						inst.Accordion_high = GongExtractString(rhs)
					case "Accordion_middle":
						inst.Accordion_middle = GongExtractInt(rhs)
					case "Accordion_low":
						inst.Accordion_low = GongExtractString(rhs)
					}
				case *Appearance:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Line_width":
						__gong__assignSliceOfPointers(&inst.Line_width, rhs, identifierMap)
					case "Note_size":
						__gong__assignSliceOfPointers(&inst.Note_size, rhs, identifierMap)
					case "Distance":
						__gong__assignSliceOfPointers(&inst.Distance, rhs, identifierMap)
					case "Glyph":
						__gong__assignSliceOfPointers(&inst.Glyph, rhs, identifierMap)
					case "Other_appearance":
						__gong__assignSliceOfPointers(&inst.Other_appearance, rhs, identifierMap)
					}
				case *Arpeggiate:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Direction":
						inst.Direction = GongExtractString(rhs)
					case "Unbroken":
						inst.Unbroken = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Arrow:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Arrow_direction":
						inst.Arrow_direction = GongExtractString(rhs)
					case "Arrow_style":
						inst.Arrow_style = GongExtractString(rhs)
					case "Arrowhead":
						inst.Arrowhead = GongExtractString(rhs)
					case "Circular_arrow":
						inst.Circular_arrow = GongExtractString(rhs)
					}
				case *Articulations:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Accent":
						__gong__assignSliceOfPointers(&inst.Accent, rhs, identifierMap)
					case "Strong_accent":
						__gong__assignSliceOfPointers(&inst.Strong_accent, rhs, identifierMap)
					case "Staccato":
						__gong__assignSliceOfPointers(&inst.Staccato, rhs, identifierMap)
					case "Tenuto":
						__gong__assignSliceOfPointers(&inst.Tenuto, rhs, identifierMap)
					case "Detached_legato":
						__gong__assignSliceOfPointers(&inst.Detached_legato, rhs, identifierMap)
					case "Staccatissimo":
						__gong__assignSliceOfPointers(&inst.Staccatissimo, rhs, identifierMap)
					case "Spiccato":
						__gong__assignSliceOfPointers(&inst.Spiccato, rhs, identifierMap)
					case "Scoop":
						__gong__assignSliceOfPointers(&inst.Scoop, rhs, identifierMap)
					case "Plop":
						__gong__assignSliceOfPointers(&inst.Plop, rhs, identifierMap)
					case "Doit":
						__gong__assignSliceOfPointers(&inst.Doit, rhs, identifierMap)
					case "Falloff":
						__gong__assignSliceOfPointers(&inst.Falloff, rhs, identifierMap)
					case "Breath_mark":
						__gong__assignSliceOfPointers(&inst.Breath_mark, rhs, identifierMap)
					case "Caesura":
						__gong__assignSliceOfPointers(&inst.Caesura, rhs, identifierMap)
					case "Stress":
						__gong__assignSliceOfPointers(&inst.Stress, rhs, identifierMap)
					case "Unstress":
						__gong__assignSliceOfPointers(&inst.Unstress, rhs, identifierMap)
					case "Soft_accent":
						__gong__assignSliceOfPointers(&inst.Soft_accent, rhs, identifierMap)
					case "Other_articulation":
						__gong__assignSliceOfPointers(&inst.Other_articulation, rhs, identifierMap)
					}
				case *Assess:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Yes_no(GongExtractString(rhs))
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = GongExtractString(rhs)
					}
				case *Attributes:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Divisions":
						inst.Divisions = GongExtractString(rhs)
					case "Key":
						__gong__assignSliceOfPointers(&inst.Key, rhs, identifierMap)
					case "Time":
						__gong__assignSliceOfPointers(&inst.Time, rhs, identifierMap)
					case "Staves":
						inst.Staves = GongExtractInt(rhs)
					case "Part_symbol":
						__gong__assignPointer(&inst.Part_symbol, rhs, identifierMap)
					case "Instruments":
						inst.Instruments = GongExtractInt(rhs)
					case "Clef":
						__gong__assignSliceOfPointers(&inst.Clef, rhs, identifierMap)
					case "Staff_details":
						__gong__assignSliceOfPointers(&inst.Staff_details, rhs, identifierMap)
					case "Transpose":
						__gong__assignSliceOfPointers(&inst.Transpose, rhs, identifierMap)
					case "For_part":
						__gong__assignSliceOfPointers(&inst.For_part, rhs, identifierMap)
					case "Directive":
						__gong__assignSliceOfPointers(&inst.Directive, rhs, identifierMap)
					case "Measure_style":
						__gong__assignSliceOfPointers(&inst.Measure_style, rhs, identifierMap)
					}
				case *Backup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *Bar_style_color:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Barline:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						inst.Location = GongExtractString(rhs)
					case "Segno":
						inst.Segno = GongExtractString(rhs)
					case "Coda":
						inst.Coda = GongExtractString(rhs)
					case "Divisions":
						inst.Divisions = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Bar_style":
						__gong__assignPointer(&inst.Bar_style, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Wavy_line":
						__gong__assignPointer(&inst.Wavy_line, rhs, identifierMap)
					case "Segno_1":
						__gong__assignPointer(&inst.Segno_1, rhs, identifierMap)
					case "Coda_1":
						__gong__assignPointer(&inst.Coda_1, rhs, identifierMap)
					case "Fermata":
						__gong__assignPointer(&inst.Fermata, rhs, identifierMap)
					case "Ending":
						__gong__assignPointer(&inst.Ending, rhs, identifierMap)
					case "Repeat":
						__gong__assignPointer(&inst.Repeat, rhs, identifierMap)
					}
				case *Barre:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *Bass:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Arrangement":
						inst.Arrangement = GongExtractString(rhs)
					case "Bass_separator":
						__gong__assignPointer(&inst.Bass_separator, rhs, identifierMap)
					case "Bass_step":
						__gong__assignPointer(&inst.Bass_step, rhs, identifierMap)
					case "Bass_alter":
						__gong__assignPointer(&inst.Bass_alter, rhs, identifierMap)
					}
				case *Bass_step:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Beam:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Repeater":
						inst.Repeater = Enum_Yes_no(GongExtractString(rhs))
					case "Fan":
						inst.Fan = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Beat_repeat:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop(GongExtractString(rhs))
					case "Slashes":
						inst.Slashes = GongExtractInt(rhs)
					case "Use_dots":
						inst.Use_dots = Enum_Yes_no(GongExtractString(rhs))
					case "Slash_type":
						inst.Slash_type = Enum_Note_type_value(GongExtractString(rhs))
					case "Slash_dot":
						inst.Slash_dot = GongExtractString(rhs)
					case "Except_voice":
						inst.Except_voice = GongExtractString(rhs)
					}
				case *Beat_unit_tied:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Beat_unit":
						inst.Beat_unit = Enum_Note_type_value(GongExtractString(rhs))
					case "Beat_unit_dot":
						inst.Beat_unit_dot = GongExtractString(rhs)
					}
				case *Beater:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tip":
						inst.Tip = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Bend:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Shape":
						inst.Shape = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Accelerate":
						inst.Accelerate = Enum_Yes_no(GongExtractString(rhs))
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "First_beat":
						inst.First_beat = GongExtractString(rhs)
					case "Last_beat":
						inst.Last_beat = GongExtractString(rhs)
					case "Bend_alter":
						inst.Bend_alter = GongExtractString(rhs)
					case "Pre_bend":
						inst.Pre_bend = GongExtractString(rhs)
					case "Release":
						__gong__assignPointer(&inst.Release, rhs, identifierMap)
					case "With_bar":
						__gong__assignPointer(&inst.With_bar, rhs, identifierMap)
					}
				case *Bookmark:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Element":
						inst.Element = GongExtractString(rhs)
					case "Position":
						inst.Position = GongExtractInt(rhs)
					}
				case *Bracket:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Line_end":
						inst.Line_end = GongExtractString(rhs)
					case "End_length":
						inst.End_length = GongExtractString(rhs)
					case "Line_type":
						inst.Line_type = GongExtractString(rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Breath_mark:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Caesura:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Cancel:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						inst.Location = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Clef:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Additional":
						inst.Additional = Enum_Yes_no(GongExtractString(rhs))
					case "Size":
						inst.Size = GongExtractString(rhs)
					case "After_barline":
						inst.After_barline = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Sign":
						inst.Sign = GongExtractString(rhs)
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "Clef_octave_change":
						inst.Clef_octave_change = GongExtractInt(rhs)
					}
				case *Coda:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Credit:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Page":
						inst.Page = GongExtractInt(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Credit_type":
						inst.Credit_type = GongExtractString(rhs)
					case "Credit_image":
						__gong__assignPointer(&inst.Credit_image, rhs, identifierMap)
					case "Link":
						__gong__assignSliceOfPointers(&inst.Link, rhs, identifierMap)
					case "Bookmark":
						__gong__assignSliceOfPointers(&inst.Bookmark, rhs, identifierMap)
					case "Credit_words":
						__gong__assignSliceOfPointers(&inst.Credit_words, rhs, identifierMap)
					case "Credit_symbol":
						__gong__assignSliceOfPointers(&inst.Credit_symbol, rhs, identifierMap)
					}
				case *Dashes:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop_continue(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Defaults:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Scaling":
						__gong__assignPointer(&inst.Scaling, rhs, identifierMap)
					case "Concert_score":
						inst.Concert_score = GongExtractString(rhs)
					case "Page_layout":
						__gong__assignPointer(&inst.Page_layout, rhs, identifierMap)
					case "System_layout":
						__gong__assignPointer(&inst.System_layout, rhs, identifierMap)
					case "Staff_layout":
						__gong__assignSliceOfPointers(&inst.Staff_layout, rhs, identifierMap)
					case "Appearance":
						__gong__assignPointer(&inst.Appearance, rhs, identifierMap)
					case "Music_font":
						__gong__assignPointer(&inst.Music_font, rhs, identifierMap)
					case "Word_font":
						__gong__assignPointer(&inst.Word_font, rhs, identifierMap)
					case "Lyric_font":
						__gong__assignSliceOfPointers(&inst.Lyric_font, rhs, identifierMap)
					case "Lyric_language":
						__gong__assignSliceOfPointers(&inst.Lyric_language, rhs, identifierMap)
					}
				case *Degree:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Degree_value":
						__gong__assignPointer(&inst.Degree_value, rhs, identifierMap)
					case "Degree_alter":
						__gong__assignPointer(&inst.Degree_alter, rhs, identifierMap)
					case "Degree_type":
						__gong__assignPointer(&inst.Degree_type, rhs, identifierMap)
					}
				case *Degree_alter:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Plus_minus":
						inst.Plus_minus = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Degree_type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Degree_value:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Symbol":
						inst.Symbol = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Direction:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Directive":
						inst.Directive = Enum_Yes_no(GongExtractString(rhs))
					case "System":
						inst.System = Enum_System_relation_number(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Direction_type":
						__gong__assignSliceOfPointers(&inst.Direction_type, rhs, identifierMap)
					case "Offset":
						__gong__assignPointer(&inst.Offset, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					case "Sound":
						__gong__assignPointer(&inst.Sound, rhs, identifierMap)
					case "Listening":
						__gong__assignPointer(&inst.Listening, rhs, identifierMap)
					}
				case *Direction_type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Rehearsal":
						__gong__assignSliceOfPointers(&inst.Rehearsal, rhs, identifierMap)
					case "Segno":
						__gong__assignSliceOfPointers(&inst.Segno, rhs, identifierMap)
					case "Coda":
						__gong__assignSliceOfPointers(&inst.Coda, rhs, identifierMap)
					case "Words":
						__gong__assignSliceOfPointers(&inst.Words, rhs, identifierMap)
					case "Symbol":
						__gong__assignSliceOfPointers(&inst.Symbol, rhs, identifierMap)
					case "Wedge":
						__gong__assignPointer(&inst.Wedge, rhs, identifierMap)
					case "Dynamics":
						__gong__assignSliceOfPointers(&inst.Dynamics, rhs, identifierMap)
					case "Dashes":
						__gong__assignPointer(&inst.Dashes, rhs, identifierMap)
					case "Bracket":
						__gong__assignPointer(&inst.Bracket, rhs, identifierMap)
					case "Pedal":
						__gong__assignPointer(&inst.Pedal, rhs, identifierMap)
					case "Metronome":
						__gong__assignPointer(&inst.Metronome, rhs, identifierMap)
					case "Octave_shift":
						__gong__assignPointer(&inst.Octave_shift, rhs, identifierMap)
					case "Harp_pedals":
						__gong__assignPointer(&inst.Harp_pedals, rhs, identifierMap)
					case "Damp":
						__gong__assignPointer(&inst.Damp, rhs, identifierMap)
					case "Damp_all":
						__gong__assignPointer(&inst.Damp_all, rhs, identifierMap)
					case "Eyeglasses":
						__gong__assignPointer(&inst.Eyeglasses, rhs, identifierMap)
					case "String_mute":
						__gong__assignPointer(&inst.String_mute, rhs, identifierMap)
					case "Scordatura":
						__gong__assignPointer(&inst.Scordatura, rhs, identifierMap)
					case "Image":
						__gong__assignPointer(&inst.Image, rhs, identifierMap)
					case "Principal_voice":
						__gong__assignPointer(&inst.Principal_voice, rhs, identifierMap)
					case "Percussion":
						__gong__assignSliceOfPointers(&inst.Percussion, rhs, identifierMap)
					case "Accordion_registration":
						__gong__assignPointer(&inst.Accordion_registration, rhs, identifierMap)
					case "Staff_divide":
						__gong__assignPointer(&inst.Staff_divide, rhs, identifierMap)
					case "Other_direction":
						__gong__assignPointer(&inst.Other_direction, rhs, identifierMap)
					}
				case *Distance:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Double:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Above":
						inst.Above = Enum_Yes_no(GongExtractString(rhs))
					}
				case *Dynamics:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Enclosure":
						inst.Enclosure = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "P":
						inst.P = GongExtractString(rhs)
					case "Pp":
						inst.Pp = GongExtractString(rhs)
					case "Ppp":
						inst.Ppp = GongExtractString(rhs)
					case "Pppp":
						inst.Pppp = GongExtractString(rhs)
					case "Ppppp":
						inst.Ppppp = GongExtractString(rhs)
					case "Pppppp":
						inst.Pppppp = GongExtractString(rhs)
					case "F":
						inst.F = GongExtractString(rhs)
					case "Ff":
						inst.Ff = GongExtractString(rhs)
					case "Fff":
						inst.Fff = GongExtractString(rhs)
					case "Ffff":
						inst.Ffff = GongExtractString(rhs)
					case "Fffff":
						inst.Fffff = GongExtractString(rhs)
					case "Ffffff":
						inst.Ffffff = GongExtractString(rhs)
					case "Mp":
						inst.Mp = GongExtractString(rhs)
					case "Mf":
						inst.Mf = GongExtractString(rhs)
					case "Sf":
						inst.Sf = GongExtractString(rhs)
					case "Sfp":
						inst.Sfp = GongExtractString(rhs)
					case "Sfpp":
						inst.Sfpp = GongExtractString(rhs)
					case "Fp":
						inst.Fp = GongExtractString(rhs)
					case "Rf":
						inst.Rf = GongExtractString(rhs)
					case "Rfz":
						inst.Rfz = GongExtractString(rhs)
					case "Sfz":
						inst.Sfz = GongExtractString(rhs)
					case "Sffz":
						inst.Sffz = GongExtractString(rhs)
					case "Fz":
						inst.Fz = GongExtractString(rhs)
					case "N":
						inst.N = GongExtractString(rhs)
					case "Pf":
						inst.Pf = GongExtractString(rhs)
					case "Sfzp":
						inst.Sfzp = GongExtractString(rhs)
					case "Other_dynamics":
						__gong__assignSliceOfPointers(&inst.Other_dynamics, rhs, identifierMap)
					}
				case *Effect:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Elision:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Empty:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Empty_font:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					}
				case *Empty_line:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Line_shape":
						inst.Line_shape = GongExtractString(rhs)
					case "Line_type":
						inst.Line_type = GongExtractString(rhs)
					case "Line_length":
						inst.Line_length = GongExtractString(rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					}
				case *Empty_placement:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					}
				case *Empty_placement_smufl:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					}
				case *Empty_print_object_style_align:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					}
				case *Empty_print_style:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *Empty_print_style_align:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					}
				case *Empty_print_style_align_id:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Empty_trill_sound:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Start_note":
						inst.Start_note = GongExtractString(rhs)
					case "Trill_step":
						inst.Trill_step = GongExtractString(rhs)
					case "Two_note_turn":
						inst.Two_note_turn = GongExtractString(rhs)
					case "Accelerate":
						inst.Accelerate = Enum_Yes_no(GongExtractString(rhs))
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "Second_beat":
						inst.Second_beat = GongExtractString(rhs)
					case "Last_beat":
						inst.Last_beat = GongExtractString(rhs)
					}
				case *Encoding:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Encoder":
						__gong__assignSliceOfPointers(&inst.Encoder, rhs, identifierMap)
					case "Software":
						inst.Software = GongExtractString(rhs)
					case "Encoding_description":
						inst.Encoding_description = GongExtractString(rhs)
					case "Supports":
						__gong__assignSliceOfPointers(&inst.Supports, rhs, identifierMap)
					}
				case *Ending:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "End_length":
						inst.End_length = GongExtractString(rhs)
					case "Text_x":
						inst.Text_x = GongExtractString(rhs)
					case "Text_y":
						inst.Text_y = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "System":
						inst.System = Enum_System_relation_number(GongExtractString(rhs))
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Extend:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop_continue(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *Feature:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Fermata:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Figure:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Prefix":
						__gong__assignPointer(&inst.Prefix, rhs, identifierMap)
					case "Figure_number":
						__gong__assignPointer(&inst.Figure_number, rhs, identifierMap)
					case "Suffix":
						__gong__assignPointer(&inst.Suffix, rhs, identifierMap)
					case "Extend":
						__gong__assignPointer(&inst.Extend, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *Figured_bass:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Parentheses":
						inst.Parentheses = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Print_dot":
						inst.Print_dot = Enum_Yes_no(GongExtractString(rhs))
					case "Print_lyric":
						inst.Print_lyric = Enum_Yes_no(GongExtractString(rhs))
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Print_spacing":
						inst.Print_spacing = Enum_Yes_no(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Figure":
						__gong__assignSliceOfPointers(&inst.Figure, rhs, identifierMap)
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *Fingering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Substitution":
						inst.Substitution = Enum_Yes_no(GongExtractString(rhs))
					case "Alternate":
						inst.Alternate = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *First_fret:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Location":
						inst.Location = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *For_part:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Part_clef":
						__gong__assignPointer(&inst.Part_clef, rhs, identifierMap)
					case "Part_transpose":
						__gong__assignPointer(&inst.Part_transpose, rhs, identifierMap)
					}
				case *Formatted_symbol:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Justify":
						inst.Justify = Enum_Left_center_right(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Rotation":
						inst.Rotation = GongExtractString(rhs)
					case "Letter_spacing":
						inst.Letter_spacing = GongExtractString(rhs)
					case "Line_height":
						inst.Line_height = GongExtractString(rhs)
					case "Dir":
						inst.Dir = GongExtractString(rhs)
					case "Enclosure":
						inst.Enclosure = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Formatted_symbol_id:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Justify":
						inst.Justify = Enum_Left_center_right(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Rotation":
						inst.Rotation = GongExtractString(rhs)
					case "Letter_spacing":
						inst.Letter_spacing = GongExtractString(rhs)
					case "Line_height":
						inst.Line_height = GongExtractString(rhs)
					case "Dir":
						inst.Dir = GongExtractString(rhs)
					case "Enclosure":
						inst.Enclosure = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Formatted_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Space":
						inst.Space = GongExtractString(rhs)
					case "Justify":
						inst.Justify = Enum_Left_center_right(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Rotation":
						inst.Rotation = GongExtractString(rhs)
					case "Letter_spacing":
						inst.Letter_spacing = GongExtractString(rhs)
					case "Line_height":
						inst.Line_height = GongExtractString(rhs)
					case "Dir":
						inst.Dir = GongExtractString(rhs)
					case "Enclosure":
						inst.Enclosure = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Formatted_text_id:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Space":
						inst.Space = GongExtractString(rhs)
					case "Justify":
						inst.Justify = Enum_Left_center_right(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Rotation":
						inst.Rotation = GongExtractString(rhs)
					case "Letter_spacing":
						inst.Letter_spacing = GongExtractString(rhs)
					case "Line_height":
						inst.Line_height = GongExtractString(rhs)
					case "Dir":
						inst.Dir = GongExtractString(rhs)
					case "Enclosure":
						inst.Enclosure = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Forward:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					}
				case *Frame:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Height":
						inst.Height = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractString(rhs)
					case "Unplayed":
						inst.Unplayed = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Frame_strings":
						inst.Frame_strings = GongExtractInt(rhs)
					case "Frame_frets":
						inst.Frame_frets = GongExtractInt(rhs)
					case "First_fret":
						__gong__assignPointer(&inst.First_fret, rhs, identifierMap)
					case "Frame_note":
						__gong__assignSliceOfPointers(&inst.Frame_note, rhs, identifierMap)
					}
				case *Frame_note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "String":
						__gong__assignPointer(&inst.String, rhs, identifierMap)
					case "Fret":
						__gong__assignPointer(&inst.Fret, rhs, identifierMap)
					case "Fingering":
						__gong__assignPointer(&inst.Fingering, rhs, identifierMap)
					case "Barre":
						__gong__assignPointer(&inst.Barre, rhs, identifierMap)
					}
				case *Fret:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Glass:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Glissando:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Line_type":
						inst.Line_type = GongExtractString(rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Glyph:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Grace:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Steal_time_previous":
						inst.Steal_time_previous = GongExtractString(rhs)
					case "Steal_time_following":
						inst.Steal_time_following = GongExtractString(rhs)
					case "Make_time":
						inst.Make_time = GongExtractString(rhs)
					case "Slash":
						inst.Slash = Enum_Yes_no(GongExtractString(rhs))
					}
				case *Group_barline:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Group_name:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Justify":
						inst.Justify = Enum_Left_center_right(GongExtractString(rhs))
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Group_symbol:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Grouping:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "Member_of":
						inst.Member_of = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Feature":
						__gong__assignSliceOfPointers(&inst.Feature, rhs, identifierMap)
					}
				case *Hammer_on_pull_off:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Handbell:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Harmon_closed:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						inst.Location = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Harmon_mute:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Harmon_closed":
						__gong__assignPointer(&inst.Harmon_closed, rhs, identifierMap)
					}
				case *Harmonic:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Natural":
						inst.Natural = GongExtractString(rhs)
					case "Artificial":
						inst.Artificial = GongExtractString(rhs)
					case "Base_pitch":
						inst.Base_pitch = GongExtractString(rhs)
					case "Touching_pitch":
						inst.Touching_pitch = GongExtractString(rhs)
					case "Sounding_pitch":
						inst.Sounding_pitch = GongExtractString(rhs)
					}
				case *Harmony:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Print_frame":
						inst.Print_frame = Enum_Yes_no(GongExtractString(rhs))
					case "Arrangement":
						inst.Arrangement = Enum_Harmony_arrangement(GongExtractString(rhs))
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "System":
						inst.System = Enum_System_relation_number(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Root":
						__gong__assignPointer(&inst.Root, rhs, identifierMap)
					case "Numeral":
						__gong__assignPointer(&inst.Numeral, rhs, identifierMap)
					case "Function":
						__gong__assignPointer(&inst.Function, rhs, identifierMap)
					case "Kind":
						__gong__assignPointer(&inst.Kind, rhs, identifierMap)
					case "Inversion":
						__gong__assignPointer(&inst.Inversion, rhs, identifierMap)
					case "Bass":
						__gong__assignPointer(&inst.Bass, rhs, identifierMap)
					case "Degree":
						__gong__assignSliceOfPointers(&inst.Degree, rhs, identifierMap)
					case "Frame":
						__gong__assignPointer(&inst.Frame, rhs, identifierMap)
					case "Offset":
						__gong__assignPointer(&inst.Offset, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					}
				case *Harmony_alter:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						inst.Location = Enum_Left_right(GongExtractString(rhs))
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Harp_pedals:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Pedal_tuning":
						__gong__assignSliceOfPointers(&inst.Pedal_tuning, rhs, identifierMap)
					}
				case *Heel_toe:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Hole:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Hole_type":
						inst.Hole_type = GongExtractString(rhs)
					case "Hole_closed":
						__gong__assignPointer(&inst.Hole_closed, rhs, identifierMap)
					case "Hole_shape":
						inst.Hole_shape = GongExtractString(rhs)
					}
				case *Hole_closed:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						inst.Location = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Horizontal_turn:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Slash":
						inst.Slash = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Start_note":
						inst.Start_note = GongExtractString(rhs)
					case "Trill_step":
						inst.Trill_step = GongExtractString(rhs)
					case "Two_note_turn":
						inst.Two_note_turn = GongExtractString(rhs)
					case "Accelerate":
						inst.Accelerate = Enum_Yes_no(GongExtractString(rhs))
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "Second_beat":
						inst.Second_beat = GongExtractString(rhs)
					case "Last_beat":
						inst.Last_beat = GongExtractString(rhs)
					}
				case *Identification:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Creator":
						__gong__assignSliceOfPointers(&inst.Creator, rhs, identifierMap)
					case "Rights":
						__gong__assignSliceOfPointers(&inst.Rights, rhs, identifierMap)
					case "Encoding":
						__gong__assignPointer(&inst.Encoding, rhs, identifierMap)
					case "Source":
						inst.Source = GongExtractString(rhs)
					case "Relation":
						__gong__assignSliceOfPointers(&inst.Relation, rhs, identifierMap)
					case "Miscellaneous":
						__gong__assignPointer(&inst.Miscellaneous, rhs, identifierMap)
					}
				case *Image:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Source":
						inst.Source = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Height":
						inst.Height = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Instrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Instrument_change:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Instrument_sound":
						inst.Instrument_sound = GongExtractString(rhs)
					case "Solo":
						inst.Solo = GongExtractString(rhs)
					case "Ensemble":
						inst.Ensemble = GongExtractString(rhs)
					case "Virtual_instrument":
						__gong__assignPointer(&inst.Virtual_instrument, rhs, identifierMap)
					}
				case *Instrument_link:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Interchangeable:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Symbol":
						inst.Symbol = GongExtractString(rhs)
					case "Separator":
						inst.Separator = GongExtractString(rhs)
					case "Time_relation":
						inst.Time_relation = GongExtractString(rhs)
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "Beat_type":
						inst.Beat_type = GongExtractString(rhs)
					}
				case *Inversion:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Key:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Cancel":
						__gong__assignPointer(&inst.Cancel, rhs, identifierMap)
					case "Fifths":
						inst.Fifths = GongExtractInt(rhs)
					case "Mode":
						inst.Mode = GongExtractString(rhs)
					case "Key_step":
						inst.Key_step = Enum_Step(GongExtractString(rhs))
					case "Key_alter":
						inst.Key_alter = GongExtractString(rhs)
					case "Key_accidental":
						__gong__assignPointer(&inst.Key_accidental, rhs, identifierMap)
					case "Key_octave":
						__gong__assignSliceOfPointers(&inst.Key_octave, rhs, identifierMap)
					}
				case *Key_accidental:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = Enum_Accidental_value(GongExtractString(rhs))
					}
				case *Key_octave:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Cancel":
						inst.Cancel = Enum_Yes_no(GongExtractString(rhs))
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Kind:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Use_symbols":
						inst.Use_symbols = Enum_Yes_no(GongExtractString(rhs))
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Stack_degrees":
						inst.Stack_degrees = Enum_Yes_no(GongExtractString(rhs))
					case "Parentheses_degrees":
						inst.Parentheses_degrees = Enum_Yes_no(GongExtractString(rhs))
					case "Bracket_degrees":
						inst.Bracket_degrees = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Level:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Reference":
						inst.Reference = Enum_Yes_no(GongExtractString(rhs))
					case "Type":
						inst.Type = Enum_Start_stop_single(GongExtractString(rhs))
					case "Parentheses":
						inst.Parentheses = Enum_Yes_no(GongExtractString(rhs))
					case "Bracket":
						inst.Bracket = Enum_Yes_no(GongExtractString(rhs))
					case "Size":
						inst.Size = Enum_Symbol_size(GongExtractString(rhs))
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Line_detail:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "Width":
						inst.Width = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Line_type":
						inst.Line_type = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					}
				case *Line_width:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Link:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					case "Element":
						inst.Element = GongExtractString(rhs)
					case "Position":
						inst.Position = GongExtractInt(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					}
				case *Listen:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Assess":
						__gong__assignSliceOfPointers(&inst.Assess, rhs, identifierMap)
					case "Wait":
						__gong__assignSliceOfPointers(&inst.Wait, rhs, identifierMap)
					case "Other_listen":
						__gong__assignSliceOfPointers(&inst.Other_listen, rhs, identifierMap)
					}
				case *Listening:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Sync":
						__gong__assignSliceOfPointers(&inst.Sync, rhs, identifierMap)
					case "Other_listening":
						__gong__assignSliceOfPointers(&inst.Other_listening, rhs, identifierMap)
					case "Offset":
						__gong__assignPointer(&inst.Offset, rhs, identifierMap)
					}
				case *Lyric:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = Enum_Time_only(GongExtractString(rhs))
					case "Justify":
						inst.Justify = Enum_Left_center_right(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Elision":
						__gong__assignSliceOfPointers(&inst.Elision, rhs, identifierMap)
					case "Syllabic":
						inst.Syllabic = GongExtractString(rhs)
					case "Text":
						__gong__assignSliceOfPointers(&inst.Text, rhs, identifierMap)
					case "Extend":
						__gong__assignPointer(&inst.Extend, rhs, identifierMap)
					case "Laughing":
						inst.Laughing = GongExtractString(rhs)
					case "Humming":
						inst.Humming = GongExtractString(rhs)
					case "End_line":
						inst.End_line = GongExtractString(rhs)
					case "End_paragraph":
						inst.End_paragraph = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *Lyric_font:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					}
				case *Lyric_language:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					}
				case *Measure_layout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Measure_distance":
						inst.Measure_distance = GongExtractString(rhs)
					}
				case *Measure_numbering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "System":
						inst.System = GongExtractString(rhs)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					case "Multiple_rest_always":
						inst.Multiple_rest_always = Enum_Yes_no(GongExtractString(rhs))
					case "Multiple_rest_range":
						inst.Multiple_rest_range = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Measure_repeat:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop(GongExtractString(rhs))
					case "Slashes":
						inst.Slashes = GongExtractInt(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Measure_style:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Multiple_rest":
						__gong__assignPointer(&inst.Multiple_rest, rhs, identifierMap)
					case "Measure_repeat":
						__gong__assignPointer(&inst.Measure_repeat, rhs, identifierMap)
					case "Beat_repeat":
						__gong__assignPointer(&inst.Beat_repeat, rhs, identifierMap)
					case "Slash":
						__gong__assignPointer(&inst.Slash, rhs, identifierMap)
					}
				case *Membrane:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Metal:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Metronome:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Parentheses":
						inst.Parentheses = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Justify":
						inst.Justify = Enum_Left_center_right(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Beat_unit":
						inst.Beat_unit = Enum_Note_type_value(GongExtractString(rhs))
					case "Beat_unit_dot":
						inst.Beat_unit_dot = GongExtractString(rhs)
					case "Per_minute":
						__gong__assignPointer(&inst.Per_minute, rhs, identifierMap)
					case "Beat_unit_tied":
						__gong__assignSliceOfPointers(&inst.Beat_unit_tied, rhs, identifierMap)
					case "Metronome_arrows":
						inst.Metronome_arrows = GongExtractString(rhs)
					case "Metronome_relation":
						inst.Metronome_relation = GongExtractString(rhs)
					case "Metronome_note":
						__gong__assignSliceOfPointers(&inst.Metronome_note, rhs, identifierMap)
					}
				case *Metronome_beam:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "EnclosedText":
						inst.EnclosedText = Enum_Beam_value(GongExtractString(rhs))
					}
				case *Metronome_note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Metronome_type":
						inst.Metronome_type = GongExtractString(rhs)
					case "Metronome_dot":
						inst.Metronome_dot = GongExtractString(rhs)
					case "Metronome_beam":
						__gong__assignSliceOfPointers(&inst.Metronome_beam, rhs, identifierMap)
					case "Metronome_tied":
						__gong__assignPointer(&inst.Metronome_tied, rhs, identifierMap)
					case "Metronome_tuplet":
						__gong__assignPointer(&inst.Metronome_tuplet, rhs, identifierMap)
					}
				case *Metronome_tied:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop(GongExtractString(rhs))
					}
				case *Metronome_tuplet:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Midi_device:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Port":
						inst.Port = GongExtractInt(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Midi_instrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Midi_channel":
						inst.Midi_channel = GongExtractInt(rhs)
					case "Midi_name":
						inst.Midi_name = GongExtractString(rhs)
					case "Midi_bank":
						inst.Midi_bank = GongExtractInt(rhs)
					case "Midi_program":
						inst.Midi_program = GongExtractInt(rhs)
					case "Midi_unpitched":
						inst.Midi_unpitched = GongExtractInt(rhs)
					case "Volume":
						inst.Volume = GongExtractString(rhs)
					case "Pan":
						inst.Pan = GongExtractString(rhs)
					case "Elevation":
						inst.Elevation = GongExtractString(rhs)
					}
				case *Miscellaneous:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Miscellaneous_field":
						__gong__assignSliceOfPointers(&inst.Miscellaneous_field, rhs, identifierMap)
					}
				case *Miscellaneous_field:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NameXSD":
						inst.NameXSD = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Mordent:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Multiple_rest:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Use_symbols":
						inst.Use_symbols = Enum_Yes_no(GongExtractString(rhs))
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Name_display:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Display_text":
						__gong__assignSliceOfPointers(&inst.Display_text, rhs, identifierMap)
					case "Accidental_text":
						__gong__assignSliceOfPointers(&inst.Accidental_text, rhs, identifierMap)
					}
				case *Non_arpeggiate:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Notations:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Tied":
						__gong__assignSliceOfPointers(&inst.Tied, rhs, identifierMap)
					case "Slur":
						__gong__assignSliceOfPointers(&inst.Slur, rhs, identifierMap)
					case "Tuplet":
						__gong__assignSliceOfPointers(&inst.Tuplet, rhs, identifierMap)
					case "Glissando":
						__gong__assignSliceOfPointers(&inst.Glissando, rhs, identifierMap)
					case "Slide":
						__gong__assignSliceOfPointers(&inst.Slide, rhs, identifierMap)
					case "Ornaments":
						__gong__assignSliceOfPointers(&inst.Ornaments, rhs, identifierMap)
					case "Technical":
						__gong__assignSliceOfPointers(&inst.Technical, rhs, identifierMap)
					case "Articulations":
						__gong__assignSliceOfPointers(&inst.Articulations, rhs, identifierMap)
					case "Dynamics":
						__gong__assignSliceOfPointers(&inst.Dynamics, rhs, identifierMap)
					case "Fermata":
						__gong__assignSliceOfPointers(&inst.Fermata, rhs, identifierMap)
					case "Arpeggiate":
						__gong__assignSliceOfPointers(&inst.Arpeggiate, rhs, identifierMap)
					case "Non_arpeggiate":
						__gong__assignSliceOfPointers(&inst.Non_arpeggiate, rhs, identifierMap)
					case "Accidental_mark":
						__gong__assignSliceOfPointers(&inst.Accidental_mark, rhs, identifierMap)
					case "Other_notation":
						__gong__assignSliceOfPointers(&inst.Other_notation, rhs, identifierMap)
					}
				case *Note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_leger":
						inst.Print_leger = Enum_Yes_no(GongExtractString(rhs))
					case "Dynamics":
						inst.Dynamics = GongExtractString(rhs)
					case "End_dynamics":
						inst.End_dynamics = GongExtractString(rhs)
					case "Attack":
						inst.Attack = GongExtractString(rhs)
					case "Release":
						inst.Release = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = Enum_Time_only(GongExtractString(rhs))
					case "Pizzicato":
						inst.Pizzicato = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Print_dot":
						inst.Print_dot = Enum_Yes_no(GongExtractString(rhs))
					case "Print_lyric":
						inst.Print_lyric = Enum_Yes_no(GongExtractString(rhs))
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Print_spacing":
						inst.Print_spacing = Enum_Yes_no(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Grace":
						__gong__assignPointer(&inst.Grace, rhs, identifierMap)
					case "Chord":
						inst.Chord = GongExtractString(rhs)
					case "Pitch":
						__gong__assignPointer(&inst.Pitch, rhs, identifierMap)
					case "Unpitched":
						__gong__assignPointer(&inst.Unpitched, rhs, identifierMap)
					case "Rest":
						__gong__assignPointer(&inst.Rest, rhs, identifierMap)
					case "Tie":
						__gong__assignPointer(&inst.Tie, rhs, identifierMap)
					case "Cue":
						inst.Cue = GongExtractString(rhs)
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					case "Instrument":
						__gong__assignSliceOfPointers(&inst.Instrument, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Type":
						__gong__assignPointer(&inst.Type, rhs, identifierMap)
					case "Dot":
						__gong__assignSliceOfPointers(&inst.Dot, rhs, identifierMap)
					case "Accidental":
						__gong__assignPointer(&inst.Accidental, rhs, identifierMap)
					case "Time_modification":
						__gong__assignPointer(&inst.Time_modification, rhs, identifierMap)
					case "Stem":
						__gong__assignPointer(&inst.Stem, rhs, identifierMap)
					case "Notehead":
						__gong__assignPointer(&inst.Notehead, rhs, identifierMap)
					case "Notehead_text":
						__gong__assignPointer(&inst.Notehead_text, rhs, identifierMap)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					case "Beam":
						__gong__assignPointer(&inst.Beam, rhs, identifierMap)
					case "Notations":
						__gong__assignSliceOfPointers(&inst.Notations, rhs, identifierMap)
					case "Lyric":
						__gong__assignSliceOfPointers(&inst.Lyric, rhs, identifierMap)
					case "Play":
						__gong__assignPointer(&inst.Play, rhs, identifierMap)
					case "Listen":
						__gong__assignPointer(&inst.Listen, rhs, identifierMap)
					}
				case *Note_size:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Note_type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Size":
						inst.Size = Enum_Symbol_size(GongExtractString(rhs))
					case "EnclosedText":
						inst.EnclosedText = Enum_Note_type_value(GongExtractString(rhs))
					}
				case *Notehead:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Filled":
						inst.Filled = Enum_Yes_no(GongExtractString(rhs))
					case "Parentheses":
						inst.Parentheses = Enum_Yes_no(GongExtractString(rhs))
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Notehead_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Display_text":
						__gong__assignSliceOfPointers(&inst.Display_text, rhs, identifierMap)
					case "Accidental_text":
						__gong__assignSliceOfPointers(&inst.Accidental_text, rhs, identifierMap)
					}
				case *Numeral:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Numeral_root":
						__gong__assignPointer(&inst.Numeral_root, rhs, identifierMap)
					case "Numeral_alter":
						__gong__assignPointer(&inst.Numeral_alter, rhs, identifierMap)
					case "Numeral_key":
						__gong__assignPointer(&inst.Numeral_key, rhs, identifierMap)
					}
				case *Numeral_key:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Numeral_fifths":
						inst.Numeral_fifths = GongExtractInt(rhs)
					case "Numeral_mode":
						inst.Numeral_mode = GongExtractString(rhs)
					}
				case *Numeral_root:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Octave_shift:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Size":
						inst.Size = GongExtractInt(rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Offset:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Sound":
						inst.Sound = Enum_Yes_no(GongExtractString(rhs))
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Opus:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *Ornaments:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Trill_mark":
						__gong__assignSliceOfPointers(&inst.Trill_mark, rhs, identifierMap)
					case "Turn":
						__gong__assignSliceOfPointers(&inst.Turn, rhs, identifierMap)
					case "Delayed_turn":
						__gong__assignSliceOfPointers(&inst.Delayed_turn, rhs, identifierMap)
					case "Inverted_turn":
						__gong__assignSliceOfPointers(&inst.Inverted_turn, rhs, identifierMap)
					case "Delayed_inverted_turn":
						__gong__assignSliceOfPointers(&inst.Delayed_inverted_turn, rhs, identifierMap)
					case "Vertical_turn":
						__gong__assignSliceOfPointers(&inst.Vertical_turn, rhs, identifierMap)
					case "Inverted_vertical_turn":
						__gong__assignSliceOfPointers(&inst.Inverted_vertical_turn, rhs, identifierMap)
					case "Shake":
						__gong__assignSliceOfPointers(&inst.Shake, rhs, identifierMap)
					case "Wavy_line":
						__gong__assignSliceOfPointers(&inst.Wavy_line, rhs, identifierMap)
					case "Mordent":
						__gong__assignSliceOfPointers(&inst.Mordent, rhs, identifierMap)
					case "Inverted_mordent":
						__gong__assignSliceOfPointers(&inst.Inverted_mordent, rhs, identifierMap)
					case "Schleifer":
						__gong__assignSliceOfPointers(&inst.Schleifer, rhs, identifierMap)
					case "Tremolo":
						__gong__assignSliceOfPointers(&inst.Tremolo, rhs, identifierMap)
					case "Haydn":
						__gong__assignSliceOfPointers(&inst.Haydn, rhs, identifierMap)
					case "Other_ornament":
						__gong__assignSliceOfPointers(&inst.Other_ornament, rhs, identifierMap)
					case "Accidental_mark":
						__gong__assignSliceOfPointers(&inst.Accidental_mark, rhs, identifierMap)
					}
				case *Other_appearance:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Other_direction:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Other_listening:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = Enum_Time_only(GongExtractString(rhs))
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Other_notation:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop_single(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Other_placement_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Other_play:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Other_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Page_layout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Page_height":
						inst.Page_height = GongExtractString(rhs)
					case "Page_width":
						inst.Page_width = GongExtractString(rhs)
					case "Page_margins":
						__gong__assignPointer(&inst.Page_margins, rhs, identifierMap)
					}
				case *Page_margins:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Left_margin":
						inst.Left_margin = GongExtractString(rhs)
					case "Right_margin":
						inst.Right_margin = GongExtractString(rhs)
					case "Top_margin":
						inst.Top_margin = GongExtractString(rhs)
					case "Bottom_margin":
						inst.Bottom_margin = GongExtractString(rhs)
					}
				case *Part_clef:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Sign":
						inst.Sign = GongExtractString(rhs)
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "Clef_octave_change":
						inst.Clef_octave_change = GongExtractInt(rhs)
					}
				case *Part_group:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "Group_name":
						__gong__assignPointer(&inst.Group_name, rhs, identifierMap)
					case "Group_name_display":
						__gong__assignPointer(&inst.Group_name_display, rhs, identifierMap)
					case "Group_abbreviation":
						__gong__assignPointer(&inst.Group_abbreviation, rhs, identifierMap)
					case "Group_abbreviation_display":
						__gong__assignPointer(&inst.Group_abbreviation_display, rhs, identifierMap)
					case "Group_symbol":
						__gong__assignPointer(&inst.Group_symbol, rhs, identifierMap)
					case "Group_barline":
						__gong__assignPointer(&inst.Group_barline, rhs, identifierMap)
					case "Group_time":
						inst.Group_time = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *Part_link:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					case "Instrument_link":
						__gong__assignSliceOfPointers(&inst.Instrument_link, rhs, identifierMap)
					case "Group_link":
						inst.Group_link = GongExtractString(rhs)
					}
				case *Part_list:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Part_group":
						__gong__assignPointer(&inst.Part_group, rhs, identifierMap)
					case "Score_part":
						__gong__assignPointer(&inst.Score_part, rhs, identifierMap)
					}
				case *Part_name:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Justify":
						inst.Justify = Enum_Left_center_right(GongExtractString(rhs))
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Part_symbol:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Top_staff":
						inst.Top_staff = GongExtractInt(rhs)
					case "Bottom_staff":
						inst.Bottom_staff = GongExtractInt(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = Enum_Group_symbol_value(GongExtractString(rhs))
					}
				case *Part_transpose:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Diatonic":
						inst.Diatonic = GongExtractInt(rhs)
					case "Chromatic":
						inst.Chromatic = GongExtractString(rhs)
					case "Octave_change":
						inst.Octave_change = GongExtractInt(rhs)
					case "Double":
						inst.Double = GongExtractFloat(rhs)
					}
				case *Pedal:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Line":
						inst.Line = Enum_Yes_no(GongExtractString(rhs))
					case "Sign":
						inst.Sign = Enum_Yes_no(GongExtractString(rhs))
					case "Abbreviated":
						inst.Abbreviated = Enum_Yes_no(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Pedal_tuning:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Pedal_step":
						inst.Pedal_step = Enum_Step(GongExtractString(rhs))
					case "Pedal_alter":
						inst.Pedal_alter = GongExtractString(rhs)
					}
				case *Per_minute:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Percussion:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Enclosure":
						inst.Enclosure = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Glass":
						__gong__assignPointer(&inst.Glass, rhs, identifierMap)
					case "Metal":
						__gong__assignPointer(&inst.Metal, rhs, identifierMap)
					case "Wood":
						__gong__assignPointer(&inst.Wood, rhs, identifierMap)
					case "Pitched":
						__gong__assignPointer(&inst.Pitched, rhs, identifierMap)
					case "Membrane":
						__gong__assignPointer(&inst.Membrane, rhs, identifierMap)
					case "Effect":
						__gong__assignPointer(&inst.Effect, rhs, identifierMap)
					case "Timpani":
						__gong__assignPointer(&inst.Timpani, rhs, identifierMap)
					case "Beater":
						__gong__assignPointer(&inst.Beater, rhs, identifierMap)
					case "Stick":
						__gong__assignPointer(&inst.Stick, rhs, identifierMap)
					case "Stick_location":
						inst.Stick_location = GongExtractString(rhs)
					case "Other_percussion":
						__gong__assignPointer(&inst.Other_percussion, rhs, identifierMap)
					}
				case *Pitch:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Step":
						inst.Step = Enum_Step(GongExtractString(rhs))
					case "Alter":
						inst.Alter = GongExtractString(rhs)
					case "Octave":
						inst.Octave = GongExtractInt(rhs)
					}
				case *Pitched:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Placement_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Play:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Ipa":
						inst.Ipa = GongExtractString(rhs)
					case "Mute":
						inst.Mute = GongExtractString(rhs)
					case "Semi_pitched":
						inst.Semi_pitched = GongExtractString(rhs)
					case "Other_play":
						__gong__assignSliceOfPointers(&inst.Other_play, rhs, identifierMap)
					}
				case *Player:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Player_name":
						inst.Player_name = GongExtractString(rhs)
					}
				case *Principal_voice:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop(GongExtractString(rhs))
					case "Symbol":
						inst.Symbol = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Print:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Staff_spacing":
						inst.Staff_spacing = GongExtractString(rhs)
					case "New_system":
						inst.New_system = Enum_Yes_no(GongExtractString(rhs))
					case "New_page":
						inst.New_page = Enum_Yes_no(GongExtractString(rhs))
					case "Blank_page":
						inst.Blank_page = GongExtractInt(rhs)
					case "Page_number":
						inst.Page_number = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Page_layout":
						__gong__assignPointer(&inst.Page_layout, rhs, identifierMap)
					case "System_layout":
						__gong__assignPointer(&inst.System_layout, rhs, identifierMap)
					case "Staff_layout":
						__gong__assignSliceOfPointers(&inst.Staff_layout, rhs, identifierMap)
					case "Measure_layout":
						__gong__assignPointer(&inst.Measure_layout, rhs, identifierMap)
					case "Measure_numbering":
						__gong__assignPointer(&inst.Measure_numbering, rhs, identifierMap)
					case "Part_name_display":
						__gong__assignPointer(&inst.Part_name_display, rhs, identifierMap)
					case "Part_abbreviation_display":
						__gong__assignPointer(&inst.Part_abbreviation_display, rhs, identifierMap)
					}
				case *Release:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Repeat:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Direction":
						inst.Direction = GongExtractString(rhs)
					case "Times":
						inst.Times = GongExtractInt(rhs)
					case "After_jump":
						inst.After_jump = Enum_Yes_no(GongExtractString(rhs))
					case "Winged":
						inst.Winged = GongExtractString(rhs)
					}
				case *Rest:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Measure":
						inst.Measure = Enum_Yes_no(GongExtractString(rhs))
					case "Display_step":
						inst.Display_step = Enum_Step(GongExtractString(rhs))
					case "Display_octave":
						inst.Display_octave = GongExtractInt(rhs)
					}
				case *Root:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Root_step":
						__gong__assignPointer(&inst.Root_step, rhs, identifierMap)
					case "Root_alter":
						__gong__assignPointer(&inst.Root_alter, rhs, identifierMap)
					}
				case *Root_step:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = Enum_Step(GongExtractString(rhs))
					}
				case *Scaling:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Millimeters":
						inst.Millimeters = GongExtractString(rhs)
					case "Tenths":
						inst.Tenths = GongExtractString(rhs)
					}
				case *Scordatura:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Accord":
						__gong__assignSliceOfPointers(&inst.Accord, rhs, identifierMap)
					}
				case *Score_instrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Instrument_name":
						inst.Instrument_name = GongExtractString(rhs)
					case "Instrument_abbreviation":
						inst.Instrument_abbreviation = GongExtractString(rhs)
					case "Instrument_sound":
						inst.Instrument_sound = GongExtractString(rhs)
					case "Solo":
						inst.Solo = GongExtractString(rhs)
					case "Ensemble":
						inst.Ensemble = GongExtractString(rhs)
					case "Virtual_instrument":
						__gong__assignPointer(&inst.Virtual_instrument, rhs, identifierMap)
					}
				case *Score_part:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Identification":
						__gong__assignPointer(&inst.Identification, rhs, identifierMap)
					case "Part_link":
						__gong__assignSliceOfPointers(&inst.Part_link, rhs, identifierMap)
					case "Part_name":
						__gong__assignPointer(&inst.Part_name, rhs, identifierMap)
					case "Part_name_display":
						__gong__assignPointer(&inst.Part_name_display, rhs, identifierMap)
					case "Part_abbreviation":
						__gong__assignPointer(&inst.Part_abbreviation, rhs, identifierMap)
					case "Part_abbreviation_display":
						__gong__assignPointer(&inst.Part_abbreviation_display, rhs, identifierMap)
					case "Group":
						inst.Group = GongExtractString(rhs)
					case "Score_instrument":
						__gong__assignSliceOfPointers(&inst.Score_instrument, rhs, identifierMap)
					case "Player":
						__gong__assignSliceOfPointers(&inst.Player, rhs, identifierMap)
					case "Midi_device":
						__gong__assignSliceOfPointers(&inst.Midi_device, rhs, identifierMap)
					case "Midi_instrument":
						__gong__assignSliceOfPointers(&inst.Midi_instrument, rhs, identifierMap)
					}
				case *Score_partwise:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Version":
						inst.Version = GongExtractString(rhs)
					case "Work":
						__gong__assignPointer(&inst.Work, rhs, identifierMap)
					case "Movement_number":
						inst.Movement_number = GongExtractString(rhs)
					case "Movement_title":
						inst.Movement_title = GongExtractString(rhs)
					case "Identification":
						__gong__assignPointer(&inst.Identification, rhs, identifierMap)
					case "Defaults":
						__gong__assignPointer(&inst.Defaults, rhs, identifierMap)
					case "Credit":
						__gong__assignSliceOfPointers(&inst.Credit, rhs, identifierMap)
					case "Part_list":
						__gong__assignPointer(&inst.Part_list, rhs, identifierMap)
					case "Part":
						__gong__assignSliceOfPointers(&inst.Part, rhs, identifierMap)
					}
				case *Score_timewise:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Version":
						inst.Version = GongExtractString(rhs)
					case "Work":
						__gong__assignPointer(&inst.Work, rhs, identifierMap)
					case "Movement_number":
						inst.Movement_number = GongExtractString(rhs)
					case "Movement_title":
						inst.Movement_title = GongExtractString(rhs)
					case "Identification":
						__gong__assignPointer(&inst.Identification, rhs, identifierMap)
					case "Defaults":
						__gong__assignPointer(&inst.Defaults, rhs, identifierMap)
					case "Credit":
						__gong__assignSliceOfPointers(&inst.Credit, rhs, identifierMap)
					case "Part_list":
						__gong__assignPointer(&inst.Part_list, rhs, identifierMap)
					case "Measure":
						__gong__assignSliceOfPointers(&inst.Measure, rhs, identifierMap)
					}
				case *Segno:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Slash:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop(GongExtractString(rhs))
					case "Use_dots":
						inst.Use_dots = Enum_Yes_no(GongExtractString(rhs))
					case "Use_stems":
						inst.Use_stems = Enum_Yes_no(GongExtractString(rhs))
					case "Slash_type":
						inst.Slash_type = Enum_Note_type_value(GongExtractString(rhs))
					case "Slash_dot":
						inst.Slash_dot = GongExtractString(rhs)
					case "Except_voice":
						inst.Except_voice = GongExtractString(rhs)
					}
				case *Slide:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Line_type":
						inst.Line_type = GongExtractString(rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Accelerate":
						inst.Accelerate = Enum_Yes_no(GongExtractString(rhs))
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "First_beat":
						inst.First_beat = GongExtractString(rhs)
					case "Last_beat":
						inst.Last_beat = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Slur:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop_continue(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Line_type":
						inst.Line_type = GongExtractString(rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Orientation":
						inst.Orientation = GongExtractString(rhs)
					case "Bezier_x":
						inst.Bezier_x = GongExtractString(rhs)
					case "Bezier_y":
						inst.Bezier_y = GongExtractString(rhs)
					case "Bezier_x2":
						inst.Bezier_x2 = GongExtractString(rhs)
					case "Bezier_y2":
						inst.Bezier_y2 = GongExtractString(rhs)
					case "Bezier_offset":
						inst.Bezier_offset = GongExtractString(rhs)
					case "Bezier_offset2":
						inst.Bezier_offset2 = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Sound:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tempo":
						inst.Tempo = GongExtractString(rhs)
					case "Dynamics":
						inst.Dynamics = GongExtractString(rhs)
					case "Dacapo":
						inst.Dacapo = Enum_Yes_no(GongExtractString(rhs))
					case "Segno":
						inst.Segno = GongExtractString(rhs)
					case "Dalsegno":
						inst.Dalsegno = GongExtractString(rhs)
					case "Coda":
						inst.Coda = GongExtractString(rhs)
					case "Tocoda":
						inst.Tocoda = GongExtractString(rhs)
					case "Divisions":
						inst.Divisions = GongExtractString(rhs)
					case "Forward_repeat":
						inst.Forward_repeat = Enum_Yes_no(GongExtractString(rhs))
					case "Fine":
						inst.Fine = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = Enum_Time_only(GongExtractString(rhs))
					case "Pizzicato":
						inst.Pizzicato = Enum_Yes_no(GongExtractString(rhs))
					case "Pan":
						inst.Pan = GongExtractString(rhs)
					case "Elevation":
						inst.Elevation = GongExtractString(rhs)
					case "Damper_pedal":
						inst.Damper_pedal = GongExtractString(rhs)
					case "Soft_pedal":
						inst.Soft_pedal = GongExtractString(rhs)
					case "Sostenuto_pedal":
						inst.Sostenuto_pedal = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Instrument_change":
						__gong__assignSliceOfPointers(&inst.Instrument_change, rhs, identifierMap)
					case "Midi_device":
						__gong__assignSliceOfPointers(&inst.Midi_device, rhs, identifierMap)
					case "Midi_instrument":
						__gong__assignSliceOfPointers(&inst.Midi_instrument, rhs, identifierMap)
					case "Play":
						__gong__assignSliceOfPointers(&inst.Play, rhs, identifierMap)
					case "Swing":
						__gong__assignPointer(&inst.Swing, rhs, identifierMap)
					case "Offset":
						__gong__assignPointer(&inst.Offset, rhs, identifierMap)
					}
				case *Staff_details:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Show_frets":
						inst.Show_frets = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Print_spacing":
						inst.Print_spacing = Enum_Yes_no(GongExtractString(rhs))
					case "Staff_type":
						inst.Staff_type = GongExtractString(rhs)
					case "Staff_lines":
						inst.Staff_lines = GongExtractInt(rhs)
					case "Line_detail":
						__gong__assignSliceOfPointers(&inst.Line_detail, rhs, identifierMap)
					case "Staff_tuning":
						__gong__assignSliceOfPointers(&inst.Staff_tuning, rhs, identifierMap)
					case "Capo":
						inst.Capo = GongExtractInt(rhs)
					case "Staff_size":
						__gong__assignPointer(&inst.Staff_size, rhs, identifierMap)
					}
				case *Staff_divide:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Staff_layout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Staff_distance":
						inst.Staff_distance = GongExtractString(rhs)
					}
				case *Staff_size:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Scaling":
						inst.Scaling = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Staff_tuning:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "Tuning_step":
						inst.Tuning_step = Enum_Step(GongExtractString(rhs))
					case "Tuning_alter":
						inst.Tuning_alter = GongExtractString(rhs)
					case "Tuning_octave":
						inst.Tuning_octave = GongExtractInt(rhs)
					}
				case *Stem:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Stick:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tip":
						inst.Tip = Enum_Tip_direction(GongExtractString(rhs))
					case "Parentheses":
						inst.Parentheses = Enum_Yes_no(GongExtractString(rhs))
					case "Dashed_circle":
						inst.Dashed_circle = Enum_Yes_no(GongExtractString(rhs))
					case "Stick_type":
						inst.Stick_type = GongExtractString(rhs)
					case "Stick_material":
						inst.Stick_material = GongExtractString(rhs)
					}
				case *String_mute:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *String_type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Strong_accent:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Style_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Supports:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Yes_no(GongExtractString(rhs))
					case "Element":
						inst.Element = GongExtractString(rhs)
					case "Attribute":
						inst.Attribute = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Swing:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Straight":
						inst.Straight = GongExtractString(rhs)
					case "First":
						inst.First = GongExtractInt(rhs)
					case "Second":
						inst.Second = GongExtractInt(rhs)
					case "Swing_type":
						inst.Swing_type = Enum_Note_type_value(GongExtractString(rhs))
					case "Swing_style":
						inst.Swing_style = GongExtractString(rhs)
					}
				case *Sync:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Latency":
						inst.Latency = GongExtractInt(rhs)
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = Enum_Time_only(GongExtractString(rhs))
					}
				case *System_dividers:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Left_divider":
						__gong__assignPointer(&inst.Left_divider, rhs, identifierMap)
					case "Right_divider":
						__gong__assignPointer(&inst.Right_divider, rhs, identifierMap)
					}
				case *System_layout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "System_margins":
						__gong__assignPointer(&inst.System_margins, rhs, identifierMap)
					case "System_distance":
						inst.System_distance = GongExtractString(rhs)
					case "Top_system_distance":
						inst.Top_system_distance = GongExtractString(rhs)
					case "System_dividers":
						__gong__assignPointer(&inst.System_dividers, rhs, identifierMap)
					}
				case *System_margins:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Left_margin":
						inst.Left_margin = GongExtractString(rhs)
					case "Right_margin":
						inst.Right_margin = GongExtractString(rhs)
					}
				case *Tap:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Hand":
						inst.Hand = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Technical:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Up_bow":
						__gong__assignSliceOfPointers(&inst.Up_bow, rhs, identifierMap)
					case "Down_bow":
						__gong__assignSliceOfPointers(&inst.Down_bow, rhs, identifierMap)
					case "Harmonic":
						__gong__assignSliceOfPointers(&inst.Harmonic, rhs, identifierMap)
					case "Open_string":
						__gong__assignSliceOfPointers(&inst.Open_string, rhs, identifierMap)
					case "Thumb_position":
						__gong__assignSliceOfPointers(&inst.Thumb_position, rhs, identifierMap)
					case "Fingering":
						__gong__assignSliceOfPointers(&inst.Fingering, rhs, identifierMap)
					case "Pluck":
						__gong__assignSliceOfPointers(&inst.Pluck, rhs, identifierMap)
					case "Double_tongue":
						__gong__assignSliceOfPointers(&inst.Double_tongue, rhs, identifierMap)
					case "Triple_tongue":
						__gong__assignSliceOfPointers(&inst.Triple_tongue, rhs, identifierMap)
					case "Stopped":
						__gong__assignSliceOfPointers(&inst.Stopped, rhs, identifierMap)
					case "Snap_pizzicato":
						__gong__assignSliceOfPointers(&inst.Snap_pizzicato, rhs, identifierMap)
					case "Fret":
						__gong__assignSliceOfPointers(&inst.Fret, rhs, identifierMap)
					case "String":
						__gong__assignSliceOfPointers(&inst.String, rhs, identifierMap)
					case "Hammer_on":
						__gong__assignSliceOfPointers(&inst.Hammer_on, rhs, identifierMap)
					case "Pull_off":
						__gong__assignSliceOfPointers(&inst.Pull_off, rhs, identifierMap)
					case "Bend":
						__gong__assignSliceOfPointers(&inst.Bend, rhs, identifierMap)
					case "Tap":
						__gong__assignSliceOfPointers(&inst.Tap, rhs, identifierMap)
					case "Heel":
						__gong__assignSliceOfPointers(&inst.Heel, rhs, identifierMap)
					case "Toe":
						__gong__assignSliceOfPointers(&inst.Toe, rhs, identifierMap)
					case "Fingernails":
						__gong__assignSliceOfPointers(&inst.Fingernails, rhs, identifierMap)
					case "Hole":
						__gong__assignSliceOfPointers(&inst.Hole, rhs, identifierMap)
					case "Arrow":
						__gong__assignSliceOfPointers(&inst.Arrow, rhs, identifierMap)
					case "Handbell":
						__gong__assignSliceOfPointers(&inst.Handbell, rhs, identifierMap)
					case "Brass_bend":
						__gong__assignSliceOfPointers(&inst.Brass_bend, rhs, identifierMap)
					case "Flip":
						__gong__assignSliceOfPointers(&inst.Flip, rhs, identifierMap)
					case "Smear":
						__gong__assignSliceOfPointers(&inst.Smear, rhs, identifierMap)
					case "Open":
						__gong__assignSliceOfPointers(&inst.Open, rhs, identifierMap)
					case "Half_muted":
						__gong__assignSliceOfPointers(&inst.Half_muted, rhs, identifierMap)
					case "Harmon_mute":
						__gong__assignSliceOfPointers(&inst.Harmon_mute, rhs, identifierMap)
					case "Golpe":
						__gong__assignSliceOfPointers(&inst.Golpe, rhs, identifierMap)
					case "Other_technical":
						__gong__assignSliceOfPointers(&inst.Other_technical, rhs, identifierMap)
					}
				case *Text_element_data:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "Line_through":
						inst.Line_through = GongExtractInt(rhs)
					case "Rotation":
						inst.Rotation = GongExtractString(rhs)
					case "Letter_spacing":
						inst.Letter_spacing = GongExtractString(rhs)
					case "Dir":
						inst.Dir = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Tie:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop(GongExtractString(rhs))
					case "Time_only":
						inst.Time_only = Enum_Time_only(GongExtractString(rhs))
					}
				case *Tied:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Line_type":
						inst.Line_type = GongExtractString(rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Orientation":
						inst.Orientation = GongExtractString(rhs)
					case "Bezier_x":
						inst.Bezier_x = GongExtractString(rhs)
					case "Bezier_y":
						inst.Bezier_y = GongExtractString(rhs)
					case "Bezier_x2":
						inst.Bezier_x2 = GongExtractString(rhs)
					case "Bezier_y2":
						inst.Bezier_y2 = GongExtractString(rhs)
					case "Bezier_offset":
						inst.Bezier_offset = GongExtractString(rhs)
					case "Bezier_offset2":
						inst.Bezier_offset2 = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Time:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Symbol":
						inst.Symbol = Enum_Time_symbol(GongExtractString(rhs))
					case "Separator":
						inst.Separator = Enum_Time_separator(GongExtractString(rhs))
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = GongExtractString(rhs)
					case "Valign":
						inst.Valign = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "Beat_type":
						inst.Beat_type = GongExtractString(rhs)
					case "Interchangeable":
						__gong__assignPointer(&inst.Interchangeable, rhs, identifierMap)
					case "Senza_misura":
						inst.Senza_misura = GongExtractString(rhs)
					}
				case *Time_modification:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Actual_notes":
						inst.Actual_notes = GongExtractInt(rhs)
					case "Normal_notes":
						inst.Normal_notes = GongExtractInt(rhs)
					case "Normal_type":
						inst.Normal_type = Enum_Note_type_value(GongExtractString(rhs))
					case "Normal_dot":
						inst.Normal_dot = GongExtractString(rhs)
					}
				case *Timpani:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					}
				case *Transpose:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Diatonic":
						inst.Diatonic = GongExtractInt(rhs)
					case "Chromatic":
						inst.Chromatic = GongExtractString(rhs)
					case "Octave_change":
						inst.Octave_change = GongExtractInt(rhs)
					case "Double":
						inst.Double = GongExtractFloat(rhs)
					}
				case *Tremolo:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Tuplet:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Bracket":
						inst.Bracket = Enum_Yes_no(GongExtractString(rhs))
					case "Show_number":
						inst.Show_number = GongExtractString(rhs)
					case "Show_type":
						inst.Show_type = Enum_Show_tuplet(GongExtractString(rhs))
					case "Line_shape":
						inst.Line_shape = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Tuplet_actual":
						__gong__assignPointer(&inst.Tuplet_actual, rhs, identifierMap)
					case "Tuplet_normal":
						__gong__assignPointer(&inst.Tuplet_normal, rhs, identifierMap)
					}
				case *Tuplet_dot:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *Tuplet_number:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *Tuplet_portion:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tuplet_number":
						__gong__assignPointer(&inst.Tuplet_number, rhs, identifierMap)
					case "Tuplet_type":
						__gong__assignPointer(&inst.Tuplet_type, rhs, identifierMap)
					case "Tuplet_dot":
						__gong__assignSliceOfPointers(&inst.Tuplet_dot, rhs, identifierMap)
					}
				case *Tuplet_type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Font_family":
						inst.Font_family = GongExtractString(rhs)
					case "Font_style":
						inst.Font_style = GongExtractString(rhs)
					case "Font_size":
						inst.Font_size = GongExtractString(rhs)
					case "Font_weight":
						inst.Font_weight = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = Enum_Note_type_value(GongExtractString(rhs))
					}
				case *Typed_text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Unpitched:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Display_step":
						inst.Display_step = Enum_Step(GongExtractString(rhs))
					case "Display_octave":
						inst.Display_octave = GongExtractInt(rhs)
					}
				case *Virtual_instrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Virtual_library":
						inst.Virtual_library = GongExtractString(rhs)
					case "Virtual_name":
						inst.Virtual_name = GongExtractString(rhs)
					}
				case *Wait:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "Time_only":
						inst.Time_only = Enum_Time_only(GongExtractString(rhs))
					}
				case *Wavy_line:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = Enum_Start_stop_continue(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Placement":
						inst.Placement = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Start_note":
						inst.Start_note = GongExtractString(rhs)
					case "Trill_step":
						inst.Trill_step = GongExtractString(rhs)
					case "Two_note_turn":
						inst.Two_note_turn = GongExtractString(rhs)
					case "Accelerate":
						inst.Accelerate = Enum_Yes_no(GongExtractString(rhs))
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "Second_beat":
						inst.Second_beat = GongExtractString(rhs)
					case "Last_beat":
						inst.Last_beat = GongExtractString(rhs)
					}
				case *Wedge:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Spread":
						inst.Spread = GongExtractString(rhs)
					case "Niente":
						inst.Niente = Enum_Yes_no(GongExtractString(rhs))
					case "Line_type":
						inst.Line_type = GongExtractString(rhs)
					case "Dash_length":
						inst.Dash_length = GongExtractString(rhs)
					case "Space_length":
						inst.Space_length = GongExtractString(rhs)
					case "Default_x":
						inst.Default_x = GongExtractString(rhs)
					case "Default_y":
						inst.Default_y = GongExtractString(rhs)
					case "Relative_x":
						inst.Relative_x = GongExtractString(rhs)
					case "Relative_y":
						inst.Relative_y = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Wood:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *Work:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Work_number":
						inst.Work_number = GongExtractString(rhs)
					case "Work_title":
						inst.Work_title = GongExtractString(rhs)
					case "Opus":
						__gong__assignPointer(&inst.Opus, rhs, identifierMap)
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
