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
		a_directiveOrdered := []*A_directive{}
		for a_directive := range stageSet.Stage.A_directives {
			a_directiveOrdered = append(a_directiveOrdered, a_directive)
		}
		sort.Slice(a_directiveOrdered, func(i, j int) bool {
			return stageSet.Stage.A_directive_stagedOrder[a_directiveOrdered[i]] < stageSet.Stage.A_directive_stagedOrder[a_directiveOrdered[j]]
		})
		for _, a_directive := range a_directiveOrdered {
			a_directiveIdent := "__stage_0" + a_directive.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.A_directive{Name: %s}).Stage(stageSet.Stage)", a_directiveIdent, __gong__toRawStringLiteral(a_directive.Name)))
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
		a_measureOrdered := []*A_measure{}
		for a_measure := range stageSet.Stage.A_measures {
			a_measureOrdered = append(a_measureOrdered, a_measure)
		}
		sort.Slice(a_measureOrdered, func(i, j int) bool {
			return stageSet.Stage.A_measure_stagedOrder[a_measureOrdered[i]] < stageSet.Stage.A_measure_stagedOrder[a_measureOrdered[j]]
		})
		for _, a_measure := range a_measureOrdered {
			a_measureIdent := "__stage_0" + a_measure.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.A_measure{Name: %s}).Stage(stageSet.Stage)", a_measureIdent, __gong__toRawStringLiteral(a_measure.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_measureIdent, __gong__toRawStringLiteral(a_measure.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", a_measureIdent, __gong__toRawStringLiteral(a_measure.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", a_measureIdent, __gong__toRawStringLiteral(a_measure.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Implicit = %s", a_measureIdent, __gong__toRawStringLiteral(string(a_measure.Implicit))))
			values.WriteString(fmt.Sprintf("\n\t%s.Non_controlling = %s", a_measureIdent, __gong__toRawStringLiteral(string(a_measure.Non_controlling))))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", a_measureIdent, __gong__toRawStringLiteral(a_measure.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", a_measureIdent, __gong__toRawStringLiteral(a_measure.Id)))
			for _, elem := range a_measure.Note {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = append(%s.Note, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Backup {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Backup = append(%s.Backup, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Forward {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Forward = append(%s.Forward, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Direction {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Direction = append(%s.Direction, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Attributes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Harmony {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmony = append(%s.Harmony, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Figured_bass {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Figured_bass = append(%s.Figured_bass, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Print {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Print = append(%s.Print, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Sound {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sound = append(%s.Sound, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Listening {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listening = append(%s.Listening, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Barline {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Barline = append(%s.Barline, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Grouping {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Grouping = append(%s.Grouping, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Link {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Link = append(%s.Link, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
			for _, elem := range a_measure.Bookmark {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bookmark = append(%s.Bookmark, %s)", a_measureIdent, a_measureIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		a_measure_1Ordered := []*A_measure_1{}
		for a_measure_1 := range stageSet.Stage.A_measure_1s {
			a_measure_1Ordered = append(a_measure_1Ordered, a_measure_1)
		}
		sort.Slice(a_measure_1Ordered, func(i, j int) bool {
			return stageSet.Stage.A_measure_1_stagedOrder[a_measure_1Ordered[i]] < stageSet.Stage.A_measure_1_stagedOrder[a_measure_1Ordered[j]]
		})
		for _, a_measure_1 := range a_measure_1Ordered {
			a_measure_1Ident := "__stage_0" + a_measure_1.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.A_measure_1{Name: %s}).Stage(stageSet.Stage)", a_measure_1Ident, __gong__toRawStringLiteral(a_measure_1.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_measure_1Ident, __gong__toRawStringLiteral(a_measure_1.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", a_measure_1Ident, __gong__toRawStringLiteral(a_measure_1.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", a_measure_1Ident, __gong__toRawStringLiteral(a_measure_1.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Implicit = %s", a_measure_1Ident, __gong__toRawStringLiteral(string(a_measure_1.Implicit))))
			values.WriteString(fmt.Sprintf("\n\t%s.Non_controlling = %s", a_measure_1Ident, __gong__toRawStringLiteral(string(a_measure_1.Non_controlling))))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", a_measure_1Ident, __gong__toRawStringLiteral(a_measure_1.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", a_measure_1Ident, __gong__toRawStringLiteral(a_measure_1.Id)))
			for _, elem := range a_measure_1.Part {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = append(%s.Part, %s)", a_measure_1Ident, a_measure_1Ident, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		a_partOrdered := []*A_part{}
		for a_part := range stageSet.Stage.A_parts {
			a_partOrdered = append(a_partOrdered, a_part)
		}
		sort.Slice(a_partOrdered, func(i, j int) bool {
			return stageSet.Stage.A_part_stagedOrder[a_partOrdered[i]] < stageSet.Stage.A_part_stagedOrder[a_partOrdered[j]]
		})
		for _, a_part := range a_partOrdered {
			a_partIdent := "__stage_0" + a_part.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.A_part{Name: %s}).Stage(stageSet.Stage)", a_partIdent, __gong__toRawStringLiteral(a_part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_partIdent, __gong__toRawStringLiteral(a_part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", a_partIdent, __gong__toRawStringLiteral(a_part.Id)))
			for _, elem := range a_part.Measure {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure = append(%s.Measure, %s)", a_partIdent, a_partIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		a_part_1Ordered := []*A_part_1{}
		for a_part_1 := range stageSet.Stage.A_part_1s {
			a_part_1Ordered = append(a_part_1Ordered, a_part_1)
		}
		sort.Slice(a_part_1Ordered, func(i, j int) bool {
			return stageSet.Stage.A_part_1_stagedOrder[a_part_1Ordered[i]] < stageSet.Stage.A_part_1_stagedOrder[a_part_1Ordered[j]]
		})
		for _, a_part_1 := range a_part_1Ordered {
			a_part_1Ident := "__stage_0" + a_part_1.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.A_part_1{Name: %s}).Stage(stageSet.Stage)", a_part_1Ident, __gong__toRawStringLiteral(a_part_1.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_part_1Ident, __gong__toRawStringLiteral(a_part_1.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", a_part_1Ident, __gong__toRawStringLiteral(a_part_1.Id)))
			for _, elem := range a_part_1.Note {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = append(%s.Note, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Backup {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Backup = append(%s.Backup, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Forward {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Forward = append(%s.Forward, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Direction {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Direction = append(%s.Direction, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Attributes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Harmony {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmony = append(%s.Harmony, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Figured_bass {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Figured_bass = append(%s.Figured_bass, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Print {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Print = append(%s.Print, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Sound {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sound = append(%s.Sound, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Listening {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listening = append(%s.Listening, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Barline {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Barline = append(%s.Barline, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Grouping {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Grouping = append(%s.Grouping, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Link {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Link = append(%s.Link, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
			for _, elem := range a_part_1.Bookmark {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bookmark = append(%s.Bookmark, %s)", a_part_1Ident, a_part_1Ident, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		accidentalOrdered := []*Accidental{}
		for accidental := range stageSet.Stage.Accidentals {
			accidentalOrdered = append(accidentalOrdered, accidental)
		}
		sort.Slice(accidentalOrdered, func(i, j int) bool {
			return stageSet.Stage.Accidental_stagedOrder[accidentalOrdered[i]] < stageSet.Stage.Accidental_stagedOrder[accidentalOrdered[j]]
		})
		for _, accidental := range accidentalOrdered {
			accidentalIdent := "__stage_0" + accidental.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Accidental{Name: %s}).Stage(stageSet.Stage)", accidentalIdent, __gong__toRawStringLiteral(accidental.Name)))
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
		accidental_markOrdered := []*Accidental_mark{}
		for accidental_mark := range stageSet.Stage.Accidental_marks {
			accidental_markOrdered = append(accidental_markOrdered, accidental_mark)
		}
		sort.Slice(accidental_markOrdered, func(i, j int) bool {
			return stageSet.Stage.Accidental_mark_stagedOrder[accidental_markOrdered[i]] < stageSet.Stage.Accidental_mark_stagedOrder[accidental_markOrdered[j]]
		})
		for _, accidental_mark := range accidental_markOrdered {
			accidental_markIdent := "__stage_0" + accidental_mark.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Accidental_mark{Name: %s}).Stage(stageSet.Stage)", accidental_markIdent, __gong__toRawStringLiteral(accidental_mark.Name)))
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
		accidental_textOrdered := []*Accidental_text{}
		for accidental_text := range stageSet.Stage.Accidental_texts {
			accidental_textOrdered = append(accidental_textOrdered, accidental_text)
		}
		sort.Slice(accidental_textOrdered, func(i, j int) bool {
			return stageSet.Stage.Accidental_text_stagedOrder[accidental_textOrdered[i]] < stageSet.Stage.Accidental_text_stagedOrder[accidental_textOrdered[j]]
		})
		for _, accidental_text := range accidental_textOrdered {
			accidental_textIdent := "__stage_0" + accidental_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Accidental_text{Name: %s}).Stage(stageSet.Stage)", accidental_textIdent, __gong__toRawStringLiteral(accidental_text.Name)))
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
		accordOrdered := []*Accord{}
		for accord := range stageSet.Stage.Accords {
			accordOrdered = append(accordOrdered, accord)
		}
		sort.Slice(accordOrdered, func(i, j int) bool {
			return stageSet.Stage.Accord_stagedOrder[accordOrdered[i]] < stageSet.Stage.Accord_stagedOrder[accordOrdered[j]]
		})
		for _, accord := range accordOrdered {
			accordIdent := "__stage_0" + accord.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Accord{Name: %s}).Stage(stageSet.Stage)", accordIdent, __gong__toRawStringLiteral(accord.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", accordIdent, __gong__toRawStringLiteral(accord.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.String = %d", accordIdent, accord.String))
			values.WriteString(fmt.Sprintf("\n\t%s.Tuning_step = %s", accordIdent, __gong__toRawStringLiteral(string(accord.Tuning_step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Tuning_alter = %s", accordIdent, __gong__toRawStringLiteral(accord.Tuning_alter)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tuning_octave = %d", accordIdent, accord.Tuning_octave))
		}
	}
	if stageSet.Stage != nil {
		accordion_registrationOrdered := []*Accordion_registration{}
		for accordion_registration := range stageSet.Stage.Accordion_registrations {
			accordion_registrationOrdered = append(accordion_registrationOrdered, accordion_registration)
		}
		sort.Slice(accordion_registrationOrdered, func(i, j int) bool {
			return stageSet.Stage.Accordion_registration_stagedOrder[accordion_registrationOrdered[i]] < stageSet.Stage.Accordion_registration_stagedOrder[accordion_registrationOrdered[j]]
		})
		for _, accordion_registration := range accordion_registrationOrdered {
			accordion_registrationIdent := "__stage_0" + accordion_registration.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Accordion_registration{Name: %s}).Stage(stageSet.Stage)", accordion_registrationIdent, __gong__toRawStringLiteral(accordion_registration.Name)))
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
		appearanceOrdered := []*Appearance{}
		for appearance := range stageSet.Stage.Appearances {
			appearanceOrdered = append(appearanceOrdered, appearance)
		}
		sort.Slice(appearanceOrdered, func(i, j int) bool {
			return stageSet.Stage.Appearance_stagedOrder[appearanceOrdered[i]] < stageSet.Stage.Appearance_stagedOrder[appearanceOrdered[j]]
		})
		for _, appearance := range appearanceOrdered {
			appearanceIdent := "__stage_0" + appearance.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Appearance{Name: %s}).Stage(stageSet.Stage)", appearanceIdent, __gong__toRawStringLiteral(appearance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", appearanceIdent, __gong__toRawStringLiteral(appearance.Name)))
			for _, elem := range appearance.Line_width {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Line_width = append(%s.Line_width, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.Note_size {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note_size = append(%s.Note_size, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.Distance {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = append(%s.Distance, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.Glyph {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Glyph = append(%s.Glyph, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.Other_appearance {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_appearance = append(%s.Other_appearance, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		arpeggiateOrdered := []*Arpeggiate{}
		for arpeggiate := range stageSet.Stage.Arpeggiates {
			arpeggiateOrdered = append(arpeggiateOrdered, arpeggiate)
		}
		sort.Slice(arpeggiateOrdered, func(i, j int) bool {
			return stageSet.Stage.Arpeggiate_stagedOrder[arpeggiateOrdered[i]] < stageSet.Stage.Arpeggiate_stagedOrder[arpeggiateOrdered[j]]
		})
		for _, arpeggiate := range arpeggiateOrdered {
			arpeggiateIdent := "__stage_0" + arpeggiate.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Arpeggiate{Name: %s}).Stage(stageSet.Stage)", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Name)))
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
		arrowOrdered := []*Arrow{}
		for arrow := range stageSet.Stage.Arrows {
			arrowOrdered = append(arrowOrdered, arrow)
		}
		sort.Slice(arrowOrdered, func(i, j int) bool {
			return stageSet.Stage.Arrow_stagedOrder[arrowOrdered[i]] < stageSet.Stage.Arrow_stagedOrder[arrowOrdered[j]]
		})
		for _, arrow := range arrowOrdered {
			arrowIdent := "__stage_0" + arrow.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Arrow{Name: %s}).Stage(stageSet.Stage)", arrowIdent, __gong__toRawStringLiteral(arrow.Name)))
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
		articulationsOrdered := []*Articulations{}
		for articulations := range stageSet.Stage.Articulationss {
			articulationsOrdered = append(articulationsOrdered, articulations)
		}
		sort.Slice(articulationsOrdered, func(i, j int) bool {
			return stageSet.Stage.Articulations_stagedOrder[articulationsOrdered[i]] < stageSet.Stage.Articulations_stagedOrder[articulationsOrdered[j]]
		})
		for _, articulations := range articulationsOrdered {
			articulationsIdent := "__stage_0" + articulations.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Articulations{Name: %s}).Stage(stageSet.Stage)", articulationsIdent, __gong__toRawStringLiteral(articulations.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", articulationsIdent, __gong__toRawStringLiteral(articulations.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", articulationsIdent, __gong__toRawStringLiteral(articulations.Id)))
			for _, elem := range articulations.Accent {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accent = append(%s.Accent, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Strong_accent {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Strong_accent = append(%s.Strong_accent, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Staccato {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staccato = append(%s.Staccato, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Tenuto {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tenuto = append(%s.Tenuto, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Detached_legato {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Detached_legato = append(%s.Detached_legato, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Staccatissimo {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staccatissimo = append(%s.Staccatissimo, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Spiccato {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Spiccato = append(%s.Spiccato, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Scoop {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Scoop = append(%s.Scoop, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Plop {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Plop = append(%s.Plop, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Doit {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Doit = append(%s.Doit, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Falloff {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Falloff = append(%s.Falloff, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Breath_mark {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Breath_mark = append(%s.Breath_mark, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Caesura {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Caesura = append(%s.Caesura, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Stress {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stress = append(%s.Stress, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Unstress {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Unstress = append(%s.Unstress, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Soft_accent {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Soft_accent = append(%s.Soft_accent, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Other_articulation {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_articulation = append(%s.Other_articulation, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		assessOrdered := []*Assess{}
		for assess := range stageSet.Stage.Assesss {
			assessOrdered = append(assessOrdered, assess)
		}
		sort.Slice(assessOrdered, func(i, j int) bool {
			return stageSet.Stage.Assess_stagedOrder[assessOrdered[i]] < stageSet.Stage.Assess_stagedOrder[assessOrdered[j]]
		})
		for _, assess := range assessOrdered {
			assessIdent := "__stage_0" + assess.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Assess{Name: %s}).Stage(stageSet.Stage)", assessIdent, __gong__toRawStringLiteral(assess.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", assessIdent, __gong__toRawStringLiteral(assess.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", assessIdent, __gong__toRawStringLiteral(string(assess.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", assessIdent, __gong__toRawStringLiteral(assess.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_only = %s", assessIdent, __gong__toRawStringLiteral(assess.Time_only)))
		}
	}
	if stageSet.Stage != nil {
		attributesOrdered := []*Attributes{}
		for attributes := range stageSet.Stage.Attributess {
			attributesOrdered = append(attributesOrdered, attributes)
		}
		sort.Slice(attributesOrdered, func(i, j int) bool {
			return stageSet.Stage.Attributes_stagedOrder[attributesOrdered[i]] < stageSet.Stage.Attributes_stagedOrder[attributesOrdered[j]]
		})
		for _, attributes := range attributesOrdered {
			attributesIdent := "__stage_0" + attributes.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Attributes{Name: %s}).Stage(stageSet.Stage)", attributesIdent, __gong__toRawStringLiteral(attributes.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attributesIdent, __gong__toRawStringLiteral(attributes.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Divisions = %s", attributesIdent, __gong__toRawStringLiteral(attributes.Divisions)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staves = %d", attributesIdent, attributes.Staves))
			values.WriteString(fmt.Sprintf("\n\t%s.Instruments = %d", attributesIdent, attributes.Instruments))
			if attributes.Footnote != nil {
				targetIdent := "__stage_0" + attributes.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", attributesIdent, targetIdent))
			}
			if attributes.Level != nil {
				targetIdent := "__stage_0" + attributes.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Key {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Key = append(%s.Key, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Time {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Time = append(%s.Time, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			if attributes.Part_symbol != nil {
				targetIdent := "__stage_0" + attributes.Part_symbol.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_symbol = %s", attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Clef {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Clef = append(%s.Clef, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Staff_details {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staff_details = append(%s.Staff_details, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Transpose {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Transpose = append(%s.Transpose, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.For_part {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.For_part = append(%s.For_part, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Directive {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Directive = append(%s.Directive, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Measure_style {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure_style = append(%s.Measure_style, %s)", attributesIdent, attributesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		backupOrdered := []*Backup{}
		for backup := range stageSet.Stage.Backups {
			backupOrdered = append(backupOrdered, backup)
		}
		sort.Slice(backupOrdered, func(i, j int) bool {
			return stageSet.Stage.Backup_stagedOrder[backupOrdered[i]] < stageSet.Stage.Backup_stagedOrder[backupOrdered[j]]
		})
		for _, backup := range backupOrdered {
			backupIdent := "__stage_0" + backup.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Backup{Name: %s}).Stage(stageSet.Stage)", backupIdent, __gong__toRawStringLiteral(backup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", backupIdent, __gong__toRawStringLiteral(backup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = %s", backupIdent, __gong__toRawStringLiteral(backup.Duration)))
			if backup.Footnote != nil {
				targetIdent := "__stage_0" + backup.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", backupIdent, targetIdent))
			}
			if backup.Level != nil {
				targetIdent := "__stage_0" + backup.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", backupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		bar_style_colorOrdered := []*Bar_style_color{}
		for bar_style_color := range stageSet.Stage.Bar_style_colors {
			bar_style_colorOrdered = append(bar_style_colorOrdered, bar_style_color)
		}
		sort.Slice(bar_style_colorOrdered, func(i, j int) bool {
			return stageSet.Stage.Bar_style_color_stagedOrder[bar_style_colorOrdered[i]] < stageSet.Stage.Bar_style_color_stagedOrder[bar_style_colorOrdered[j]]
		})
		for _, bar_style_color := range bar_style_colorOrdered {
			bar_style_colorIdent := "__stage_0" + bar_style_color.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Bar_style_color{Name: %s}).Stage(stageSet.Stage)", bar_style_colorIdent, __gong__toRawStringLiteral(bar_style_color.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bar_style_colorIdent, __gong__toRawStringLiteral(bar_style_color.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", bar_style_colorIdent, __gong__toRawStringLiteral(bar_style_color.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", bar_style_colorIdent, __gong__toRawStringLiteral(bar_style_color.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		barlineOrdered := []*Barline{}
		for barline := range stageSet.Stage.Barlines {
			barlineOrdered = append(barlineOrdered, barline)
		}
		sort.Slice(barlineOrdered, func(i, j int) bool {
			return stageSet.Stage.Barline_stagedOrder[barlineOrdered[i]] < stageSet.Stage.Barline_stagedOrder[barlineOrdered[j]]
		})
		for _, barline := range barlineOrdered {
			barlineIdent := "__stage_0" + barline.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Barline{Name: %s}).Stage(stageSet.Stage)", barlineIdent, __gong__toRawStringLiteral(barline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", barlineIdent, __gong__toRawStringLiteral(barline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", barlineIdent, __gong__toRawStringLiteral(barline.Location)))
			values.WriteString(fmt.Sprintf("\n\t%s.Segno = %s", barlineIdent, __gong__toRawStringLiteral(barline.Segno)))
			values.WriteString(fmt.Sprintf("\n\t%s.Coda = %s", barlineIdent, __gong__toRawStringLiteral(barline.Coda)))
			values.WriteString(fmt.Sprintf("\n\t%s.Divisions = %s", barlineIdent, __gong__toRawStringLiteral(barline.Divisions)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", barlineIdent, __gong__toRawStringLiteral(barline.Id)))
			if barline.Bar_style != nil {
				targetIdent := "__stage_0" + barline.Bar_style.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bar_style = %s", barlineIdent, targetIdent))
			}
			if barline.Footnote != nil {
				targetIdent := "__stage_0" + barline.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", barlineIdent, targetIdent))
			}
			if barline.Level != nil {
				targetIdent := "__stage_0" + barline.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", barlineIdent, targetIdent))
			}
			if barline.Wavy_line != nil {
				targetIdent := "__stage_0" + barline.Wavy_line.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wavy_line = %s", barlineIdent, targetIdent))
			}
			if barline.Segno_1 != nil {
				targetIdent := "__stage_0" + barline.Segno_1.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Segno_1 = %s", barlineIdent, targetIdent))
			}
			if barline.Coda_1 != nil {
				targetIdent := "__stage_0" + barline.Coda_1.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Coda_1 = %s", barlineIdent, targetIdent))
			}
			if barline.Fermata != nil {
				targetIdent := "__stage_0" + barline.Fermata.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fermata = %s", barlineIdent, targetIdent))
			}
			if barline.Ending != nil {
				targetIdent := "__stage_0" + barline.Ending.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ending = %s", barlineIdent, targetIdent))
			}
			if barline.Repeat != nil {
				targetIdent := "__stage_0" + barline.Repeat.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Repeat = %s", barlineIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		barreOrdered := []*Barre{}
		for barre := range stageSet.Stage.Barres {
			barreOrdered = append(barreOrdered, barre)
		}
		sort.Slice(barreOrdered, func(i, j int) bool {
			return stageSet.Stage.Barre_stagedOrder[barreOrdered[i]] < stageSet.Stage.Barre_stagedOrder[barreOrdered[j]]
		})
		for _, barre := range barreOrdered {
			barreIdent := "__stage_0" + barre.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Barre{Name: %s}).Stage(stageSet.Stage)", barreIdent, __gong__toRawStringLiteral(barre.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", barreIdent, __gong__toRawStringLiteral(barre.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", barreIdent, __gong__toRawStringLiteral(barre.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", barreIdent, __gong__toRawStringLiteral(barre.Color)))
		}
	}
	if stageSet.Stage != nil {
		bassOrdered := []*Bass{}
		for bass := range stageSet.Stage.Basss {
			bassOrdered = append(bassOrdered, bass)
		}
		sort.Slice(bassOrdered, func(i, j int) bool {
			return stageSet.Stage.Bass_stagedOrder[bassOrdered[i]] < stageSet.Stage.Bass_stagedOrder[bassOrdered[j]]
		})
		for _, bass := range bassOrdered {
			bassIdent := "__stage_0" + bass.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Bass{Name: %s}).Stage(stageSet.Stage)", bassIdent, __gong__toRawStringLiteral(bass.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bassIdent, __gong__toRawStringLiteral(bass.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arrangement = %s", bassIdent, __gong__toRawStringLiteral(bass.Arrangement)))
			if bass.Bass_separator != nil {
				targetIdent := "__stage_0" + bass.Bass_separator.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bass_separator = %s", bassIdent, targetIdent))
			}
			if bass.Bass_step != nil {
				targetIdent := "__stage_0" + bass.Bass_step.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bass_step = %s", bassIdent, targetIdent))
			}
			if bass.Bass_alter != nil {
				targetIdent := "__stage_0" + bass.Bass_alter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bass_alter = %s", bassIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		bass_stepOrdered := []*Bass_step{}
		for bass_step := range stageSet.Stage.Bass_steps {
			bass_stepOrdered = append(bass_stepOrdered, bass_step)
		}
		sort.Slice(bass_stepOrdered, func(i, j int) bool {
			return stageSet.Stage.Bass_step_stagedOrder[bass_stepOrdered[i]] < stageSet.Stage.Bass_step_stagedOrder[bass_stepOrdered[j]]
		})
		for _, bass_step := range bass_stepOrdered {
			bass_stepIdent := "__stage_0" + bass_step.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Bass_step{Name: %s}).Stage(stageSet.Stage)", bass_stepIdent, __gong__toRawStringLiteral(bass_step.Name)))
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
		beamOrdered := []*Beam{}
		for beam := range stageSet.Stage.Beams {
			beamOrdered = append(beamOrdered, beam)
		}
		sort.Slice(beamOrdered, func(i, j int) bool {
			return stageSet.Stage.Beam_stagedOrder[beamOrdered[i]] < stageSet.Stage.Beam_stagedOrder[beamOrdered[j]]
		})
		for _, beam := range beamOrdered {
			beamIdent := "__stage_0" + beam.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Beam{Name: %s}).Stage(stageSet.Stage)", beamIdent, __gong__toRawStringLiteral(beam.Name)))
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
		beat_repeatOrdered := []*Beat_repeat{}
		for beat_repeat := range stageSet.Stage.Beat_repeats {
			beat_repeatOrdered = append(beat_repeatOrdered, beat_repeat)
		}
		sort.Slice(beat_repeatOrdered, func(i, j int) bool {
			return stageSet.Stage.Beat_repeat_stagedOrder[beat_repeatOrdered[i]] < stageSet.Stage.Beat_repeat_stagedOrder[beat_repeatOrdered[j]]
		})
		for _, beat_repeat := range beat_repeatOrdered {
			beat_repeatIdent := "__stage_0" + beat_repeat.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Beat_repeat{Name: %s}).Stage(stageSet.Stage)", beat_repeatIdent, __gong__toRawStringLiteral(beat_repeat.Name)))
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
		beat_unit_tiedOrdered := []*Beat_unit_tied{}
		for beat_unit_tied := range stageSet.Stage.Beat_unit_tieds {
			beat_unit_tiedOrdered = append(beat_unit_tiedOrdered, beat_unit_tied)
		}
		sort.Slice(beat_unit_tiedOrdered, func(i, j int) bool {
			return stageSet.Stage.Beat_unit_tied_stagedOrder[beat_unit_tiedOrdered[i]] < stageSet.Stage.Beat_unit_tied_stagedOrder[beat_unit_tiedOrdered[j]]
		})
		for _, beat_unit_tied := range beat_unit_tiedOrdered {
			beat_unit_tiedIdent := "__stage_0" + beat_unit_tied.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Beat_unit_tied{Name: %s}).Stage(stageSet.Stage)", beat_unit_tiedIdent, __gong__toRawStringLiteral(beat_unit_tied.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", beat_unit_tiedIdent, __gong__toRawStringLiteral(beat_unit_tied.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Beat_unit = %s", beat_unit_tiedIdent, __gong__toRawStringLiteral(string(beat_unit_tied.Beat_unit))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beat_unit_dot = %s", beat_unit_tiedIdent, __gong__toRawStringLiteral(beat_unit_tied.Beat_unit_dot)))
		}
	}
	if stageSet.Stage != nil {
		beaterOrdered := []*Beater{}
		for beater := range stageSet.Stage.Beaters {
			beaterOrdered = append(beaterOrdered, beater)
		}
		sort.Slice(beaterOrdered, func(i, j int) bool {
			return stageSet.Stage.Beater_stagedOrder[beaterOrdered[i]] < stageSet.Stage.Beater_stagedOrder[beaterOrdered[j]]
		})
		for _, beater := range beaterOrdered {
			beaterIdent := "__stage_0" + beater.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Beater{Name: %s}).Stage(stageSet.Stage)", beaterIdent, __gong__toRawStringLiteral(beater.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", beaterIdent, __gong__toRawStringLiteral(beater.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tip = %s", beaterIdent, __gong__toRawStringLiteral(beater.Tip)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", beaterIdent, __gong__toRawStringLiteral(beater.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		bendOrdered := []*Bend{}
		for bend := range stageSet.Stage.Bends {
			bendOrdered = append(bendOrdered, bend)
		}
		sort.Slice(bendOrdered, func(i, j int) bool {
			return stageSet.Stage.Bend_stagedOrder[bendOrdered[i]] < stageSet.Stage.Bend_stagedOrder[bendOrdered[j]]
		})
		for _, bend := range bendOrdered {
			bendIdent := "__stage_0" + bend.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Bend{Name: %s}).Stage(stageSet.Stage)", bendIdent, __gong__toRawStringLiteral(bend.Name)))
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
				targetIdent := "__stage_0" + bend.Release.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Release = %s", bendIdent, targetIdent))
			}
			if bend.With_bar != nil {
				targetIdent := "__stage_0" + bend.With_bar.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.With_bar = %s", bendIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		bookmarkOrdered := []*Bookmark{}
		for bookmark := range stageSet.Stage.Bookmarks {
			bookmarkOrdered = append(bookmarkOrdered, bookmark)
		}
		sort.Slice(bookmarkOrdered, func(i, j int) bool {
			return stageSet.Stage.Bookmark_stagedOrder[bookmarkOrdered[i]] < stageSet.Stage.Bookmark_stagedOrder[bookmarkOrdered[j]]
		})
		for _, bookmark := range bookmarkOrdered {
			bookmarkIdent := "__stage_0" + bookmark.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Bookmark{Name: %s}).Stage(stageSet.Stage)", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Element = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Element)))
			values.WriteString(fmt.Sprintf("\n\t%s.Position = %d", bookmarkIdent, bookmark.Position))
		}
	}
	if stageSet.Stage != nil {
		bracketOrdered := []*Bracket{}
		for bracket := range stageSet.Stage.Brackets {
			bracketOrdered = append(bracketOrdered, bracket)
		}
		sort.Slice(bracketOrdered, func(i, j int) bool {
			return stageSet.Stage.Bracket_stagedOrder[bracketOrdered[i]] < stageSet.Stage.Bracket_stagedOrder[bracketOrdered[j]]
		})
		for _, bracket := range bracketOrdered {
			bracketIdent := "__stage_0" + bracket.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Bracket{Name: %s}).Stage(stageSet.Stage)", bracketIdent, __gong__toRawStringLiteral(bracket.Name)))
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
		breath_markOrdered := []*Breath_mark{}
		for breath_mark := range stageSet.Stage.Breath_marks {
			breath_markOrdered = append(breath_markOrdered, breath_mark)
		}
		sort.Slice(breath_markOrdered, func(i, j int) bool {
			return stageSet.Stage.Breath_mark_stagedOrder[breath_markOrdered[i]] < stageSet.Stage.Breath_mark_stagedOrder[breath_markOrdered[j]]
		})
		for _, breath_mark := range breath_markOrdered {
			breath_markIdent := "__stage_0" + breath_mark.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Breath_mark{Name: %s}).Stage(stageSet.Stage)", breath_markIdent, __gong__toRawStringLiteral(breath_mark.Name)))
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
		caesuraOrdered := []*Caesura{}
		for caesura := range stageSet.Stage.Caesuras {
			caesuraOrdered = append(caesuraOrdered, caesura)
		}
		sort.Slice(caesuraOrdered, func(i, j int) bool {
			return stageSet.Stage.Caesura_stagedOrder[caesuraOrdered[i]] < stageSet.Stage.Caesura_stagedOrder[caesuraOrdered[j]]
		})
		for _, caesura := range caesuraOrdered {
			caesuraIdent := "__stage_0" + caesura.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Caesura{Name: %s}).Stage(stageSet.Stage)", caesuraIdent, __gong__toRawStringLiteral(caesura.Name)))
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
		cancelOrdered := []*Cancel{}
		for cancel := range stageSet.Stage.Cancels {
			cancelOrdered = append(cancelOrdered, cancel)
		}
		sort.Slice(cancelOrdered, func(i, j int) bool {
			return stageSet.Stage.Cancel_stagedOrder[cancelOrdered[i]] < stageSet.Stage.Cancel_stagedOrder[cancelOrdered[j]]
		})
		for _, cancel := range cancelOrdered {
			cancelIdent := "__stage_0" + cancel.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Cancel{Name: %s}).Stage(stageSet.Stage)", cancelIdent, __gong__toRawStringLiteral(cancel.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cancelIdent, __gong__toRawStringLiteral(cancel.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", cancelIdent, __gong__toRawStringLiteral(cancel.Location)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", cancelIdent, cancel.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		clefOrdered := []*Clef{}
		for clef := range stageSet.Stage.Clefs {
			clefOrdered = append(clefOrdered, clef)
		}
		sort.Slice(clefOrdered, func(i, j int) bool {
			return stageSet.Stage.Clef_stagedOrder[clefOrdered[i]] < stageSet.Stage.Clef_stagedOrder[clefOrdered[j]]
		})
		for _, clef := range clefOrdered {
			clefIdent := "__stage_0" + clef.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Clef{Name: %s}).Stage(stageSet.Stage)", clefIdent, __gong__toRawStringLiteral(clef.Name)))
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
		codaOrdered := []*Coda{}
		for coda := range stageSet.Stage.Codas {
			codaOrdered = append(codaOrdered, coda)
		}
		sort.Slice(codaOrdered, func(i, j int) bool {
			return stageSet.Stage.Coda_stagedOrder[codaOrdered[i]] < stageSet.Stage.Coda_stagedOrder[codaOrdered[j]]
		})
		for _, coda := range codaOrdered {
			codaIdent := "__stage_0" + coda.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Coda{Name: %s}).Stage(stageSet.Stage)", codaIdent, __gong__toRawStringLiteral(coda.Name)))
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
		creditOrdered := []*Credit{}
		for credit := range stageSet.Stage.Credits {
			creditOrdered = append(creditOrdered, credit)
		}
		sort.Slice(creditOrdered, func(i, j int) bool {
			return stageSet.Stage.Credit_stagedOrder[creditOrdered[i]] < stageSet.Stage.Credit_stagedOrder[creditOrdered[j]]
		})
		for _, credit := range creditOrdered {
			creditIdent := "__stage_0" + credit.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Credit{Name: %s}).Stage(stageSet.Stage)", creditIdent, __gong__toRawStringLiteral(credit.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", creditIdent, __gong__toRawStringLiteral(credit.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Page = %d", creditIdent, credit.Page))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", creditIdent, __gong__toRawStringLiteral(credit.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Credit_type = %s", creditIdent, __gong__toRawStringLiteral(credit.Credit_type)))
			if credit.Credit_image != nil {
				targetIdent := "__stage_0" + credit.Credit_image.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit_image = %s", creditIdent, targetIdent))
			}
			for _, elem := range credit.Link {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Link = append(%s.Link, %s)", creditIdent, creditIdent, targetIdent))
			}
			for _, elem := range credit.Bookmark {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bookmark = append(%s.Bookmark, %s)", creditIdent, creditIdent, targetIdent))
			}
			for _, elem := range credit.Credit_words {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit_words = append(%s.Credit_words, %s)", creditIdent, creditIdent, targetIdent))
			}
			for _, elem := range credit.Credit_symbol {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit_symbol = append(%s.Credit_symbol, %s)", creditIdent, creditIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		dashesOrdered := []*Dashes{}
		for dashes := range stageSet.Stage.Dashess {
			dashesOrdered = append(dashesOrdered, dashes)
		}
		sort.Slice(dashesOrdered, func(i, j int) bool {
			return stageSet.Stage.Dashes_stagedOrder[dashesOrdered[i]] < stageSet.Stage.Dashes_stagedOrder[dashesOrdered[j]]
		})
		for _, dashes := range dashesOrdered {
			dashesIdent := "__stage_0" + dashes.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Dashes{Name: %s}).Stage(stageSet.Stage)", dashesIdent, __gong__toRawStringLiteral(dashes.Name)))
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
		defaultsOrdered := []*Defaults{}
		for defaults := range stageSet.Stage.Defaultss {
			defaultsOrdered = append(defaultsOrdered, defaults)
		}
		sort.Slice(defaultsOrdered, func(i, j int) bool {
			return stageSet.Stage.Defaults_stagedOrder[defaultsOrdered[i]] < stageSet.Stage.Defaults_stagedOrder[defaultsOrdered[j]]
		})
		for _, defaults := range defaultsOrdered {
			defaultsIdent := "__stage_0" + defaults.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Defaults{Name: %s}).Stage(stageSet.Stage)", defaultsIdent, __gong__toRawStringLiteral(defaults.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", defaultsIdent, __gong__toRawStringLiteral(defaults.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Concert_score = %s", defaultsIdent, __gong__toRawStringLiteral(defaults.Concert_score)))
			if defaults.Scaling != nil {
				targetIdent := "__stage_0" + defaults.Scaling.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Scaling = %s", defaultsIdent, targetIdent))
			}
			if defaults.Page_layout != nil {
				targetIdent := "__stage_0" + defaults.Page_layout.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Page_layout = %s", defaultsIdent, targetIdent))
			}
			if defaults.System_layout != nil {
				targetIdent := "__stage_0" + defaults.System_layout.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System_layout = %s", defaultsIdent, targetIdent))
			}
			for _, elem := range defaults.Staff_layout {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staff_layout = append(%s.Staff_layout, %s)", defaultsIdent, defaultsIdent, targetIdent))
			}
			if defaults.Appearance != nil {
				targetIdent := "__stage_0" + defaults.Appearance.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Appearance = %s", defaultsIdent, targetIdent))
			}
			if defaults.Music_font != nil {
				targetIdent := "__stage_0" + defaults.Music_font.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Music_font = %s", defaultsIdent, targetIdent))
			}
			if defaults.Word_font != nil {
				targetIdent := "__stage_0" + defaults.Word_font.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Word_font = %s", defaultsIdent, targetIdent))
			}
			for _, elem := range defaults.Lyric_font {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lyric_font = append(%s.Lyric_font, %s)", defaultsIdent, defaultsIdent, targetIdent))
			}
			for _, elem := range defaults.Lyric_language {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lyric_language = append(%s.Lyric_language, %s)", defaultsIdent, defaultsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		degreeOrdered := []*Degree{}
		for degree := range stageSet.Stage.Degrees {
			degreeOrdered = append(degreeOrdered, degree)
		}
		sort.Slice(degreeOrdered, func(i, j int) bool {
			return stageSet.Stage.Degree_stagedOrder[degreeOrdered[i]] < stageSet.Stage.Degree_stagedOrder[degreeOrdered[j]]
		})
		for _, degree := range degreeOrdered {
			degreeIdent := "__stage_0" + degree.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Degree{Name: %s}).Stage(stageSet.Stage)", degreeIdent, __gong__toRawStringLiteral(degree.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", degreeIdent, __gong__toRawStringLiteral(degree.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", degreeIdent, __gong__toRawStringLiteral(string(degree.Print_object))))
			if degree.Degree_value != nil {
				targetIdent := "__stage_0" + degree.Degree_value.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Degree_value = %s", degreeIdent, targetIdent))
			}
			if degree.Degree_alter != nil {
				targetIdent := "__stage_0" + degree.Degree_alter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Degree_alter = %s", degreeIdent, targetIdent))
			}
			if degree.Degree_type != nil {
				targetIdent := "__stage_0" + degree.Degree_type.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Degree_type = %s", degreeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		degree_alterOrdered := []*Degree_alter{}
		for degree_alter := range stageSet.Stage.Degree_alters {
			degree_alterOrdered = append(degree_alterOrdered, degree_alter)
		}
		sort.Slice(degree_alterOrdered, func(i, j int) bool {
			return stageSet.Stage.Degree_alter_stagedOrder[degree_alterOrdered[i]] < stageSet.Stage.Degree_alter_stagedOrder[degree_alterOrdered[j]]
		})
		for _, degree_alter := range degree_alterOrdered {
			degree_alterIdent := "__stage_0" + degree_alter.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Degree_alter{Name: %s}).Stage(stageSet.Stage)", degree_alterIdent, __gong__toRawStringLiteral(degree_alter.Name)))
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
		degree_typeOrdered := []*Degree_type{}
		for degree_type := range stageSet.Stage.Degree_types {
			degree_typeOrdered = append(degree_typeOrdered, degree_type)
		}
		sort.Slice(degree_typeOrdered, func(i, j int) bool {
			return stageSet.Stage.Degree_type_stagedOrder[degree_typeOrdered[i]] < stageSet.Stage.Degree_type_stagedOrder[degree_typeOrdered[j]]
		})
		for _, degree_type := range degree_typeOrdered {
			degree_typeIdent := "__stage_0" + degree_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Degree_type{Name: %s}).Stage(stageSet.Stage)", degree_typeIdent, __gong__toRawStringLiteral(degree_type.Name)))
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
		degree_valueOrdered := []*Degree_value{}
		for degree_value := range stageSet.Stage.Degree_values {
			degree_valueOrdered = append(degree_valueOrdered, degree_value)
		}
		sort.Slice(degree_valueOrdered, func(i, j int) bool {
			return stageSet.Stage.Degree_value_stagedOrder[degree_valueOrdered[i]] < stageSet.Stage.Degree_value_stagedOrder[degree_valueOrdered[j]]
		})
		for _, degree_value := range degree_valueOrdered {
			degree_valueIdent := "__stage_0" + degree_value.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Degree_value{Name: %s}).Stage(stageSet.Stage)", degree_valueIdent, __gong__toRawStringLiteral(degree_value.Name)))
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
		directionOrdered := []*Direction{}
		for direction := range stageSet.Stage.Directions {
			directionOrdered = append(directionOrdered, direction)
		}
		sort.Slice(directionOrdered, func(i, j int) bool {
			return stageSet.Stage.Direction_stagedOrder[directionOrdered[i]] < stageSet.Stage.Direction_stagedOrder[directionOrdered[j]]
		})
		for _, direction := range directionOrdered {
			directionIdent := "__stage_0" + direction.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Direction{Name: %s}).Stage(stageSet.Stage)", directionIdent, __gong__toRawStringLiteral(direction.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionIdent, __gong__toRawStringLiteral(direction.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", directionIdent, __gong__toRawStringLiteral(direction.Placement)))
			values.WriteString(fmt.Sprintf("\n\t%s.Directive = %s", directionIdent, __gong__toRawStringLiteral(string(direction.Directive))))
			values.WriteString(fmt.Sprintf("\n\t%s.System = %s", directionIdent, __gong__toRawStringLiteral(string(direction.System))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directionIdent, __gong__toRawStringLiteral(direction.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Voice = %s", directionIdent, __gong__toRawStringLiteral(direction.Voice)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff = %d", directionIdent, direction.Staff))
			for _, elem := range direction.Direction_type {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Direction_type = append(%s.Direction_type, %s)", directionIdent, directionIdent, targetIdent))
			}
			if direction.Offset != nil {
				targetIdent := "__stage_0" + direction.Offset.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", directionIdent, targetIdent))
			}
			if direction.Footnote != nil {
				targetIdent := "__stage_0" + direction.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", directionIdent, targetIdent))
			}
			if direction.Level != nil {
				targetIdent := "__stage_0" + direction.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", directionIdent, targetIdent))
			}
			if direction.Sound != nil {
				targetIdent := "__stage_0" + direction.Sound.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sound = %s", directionIdent, targetIdent))
			}
			if direction.Listening != nil {
				targetIdent := "__stage_0" + direction.Listening.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listening = %s", directionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		direction_typeOrdered := []*Direction_type{}
		for direction_type := range stageSet.Stage.Direction_types {
			direction_typeOrdered = append(direction_typeOrdered, direction_type)
		}
		sort.Slice(direction_typeOrdered, func(i, j int) bool {
			return stageSet.Stage.Direction_type_stagedOrder[direction_typeOrdered[i]] < stageSet.Stage.Direction_type_stagedOrder[direction_typeOrdered[j]]
		})
		for _, direction_type := range direction_typeOrdered {
			direction_typeIdent := "__stage_0" + direction_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Direction_type{Name: %s}).Stage(stageSet.Stage)", direction_typeIdent, __gong__toRawStringLiteral(direction_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", direction_typeIdent, __gong__toRawStringLiteral(direction_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", direction_typeIdent, __gong__toRawStringLiteral(direction_type.Id)))
			for _, elem := range direction_type.Rehearsal {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rehearsal = append(%s.Rehearsal, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			for _, elem := range direction_type.Segno {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Segno = append(%s.Segno, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			for _, elem := range direction_type.Coda {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Coda = append(%s.Coda, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			for _, elem := range direction_type.Words {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Words = append(%s.Words, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			for _, elem := range direction_type.Symbol {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Symbol = append(%s.Symbol, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			if direction_type.Wedge != nil {
				targetIdent := "__stage_0" + direction_type.Wedge.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wedge = %s", direction_typeIdent, targetIdent))
			}
			for _, elem := range direction_type.Dynamics {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dynamics = append(%s.Dynamics, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			if direction_type.Dashes != nil {
				targetIdent := "__stage_0" + direction_type.Dashes.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dashes = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Bracket != nil {
				targetIdent := "__stage_0" + direction_type.Bracket.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bracket = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Pedal != nil {
				targetIdent := "__stage_0" + direction_type.Pedal.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pedal = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Metronome != nil {
				targetIdent := "__stage_0" + direction_type.Metronome.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metronome = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Octave_shift != nil {
				targetIdent := "__stage_0" + direction_type.Octave_shift.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Octave_shift = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Harp_pedals != nil {
				targetIdent := "__stage_0" + direction_type.Harp_pedals.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harp_pedals = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Damp != nil {
				targetIdent := "__stage_0" + direction_type.Damp.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Damp = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Damp_all != nil {
				targetIdent := "__stage_0" + direction_type.Damp_all.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Damp_all = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Eyeglasses != nil {
				targetIdent := "__stage_0" + direction_type.Eyeglasses.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Eyeglasses = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.String_mute != nil {
				targetIdent := "__stage_0" + direction_type.String_mute.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.String_mute = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Scordatura != nil {
				targetIdent := "__stage_0" + direction_type.Scordatura.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Scordatura = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Image != nil {
				targetIdent := "__stage_0" + direction_type.Image.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Image = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Principal_voice != nil {
				targetIdent := "__stage_0" + direction_type.Principal_voice.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Principal_voice = %s", direction_typeIdent, targetIdent))
			}
			for _, elem := range direction_type.Percussion {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Percussion = append(%s.Percussion, %s)", direction_typeIdent, direction_typeIdent, targetIdent))
			}
			if direction_type.Accordion_registration != nil {
				targetIdent := "__stage_0" + direction_type.Accordion_registration.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accordion_registration = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Staff_divide != nil {
				targetIdent := "__stage_0" + direction_type.Staff_divide.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staff_divide = %s", direction_typeIdent, targetIdent))
			}
			if direction_type.Other_direction != nil {
				targetIdent := "__stage_0" + direction_type.Other_direction.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_direction = %s", direction_typeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		distanceOrdered := []*Distance{}
		for distance := range stageSet.Stage.Distances {
			distanceOrdered = append(distanceOrdered, distance)
		}
		sort.Slice(distanceOrdered, func(i, j int) bool {
			return stageSet.Stage.Distance_stagedOrder[distanceOrdered[i]] < stageSet.Stage.Distance_stagedOrder[distanceOrdered[j]]
		})
		for _, distance := range distanceOrdered {
			distanceIdent := "__stage_0" + distance.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Distance{Name: %s}).Stage(stageSet.Stage)", distanceIdent, __gong__toRawStringLiteral(distance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", distanceIdent, __gong__toRawStringLiteral(distance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", distanceIdent, __gong__toRawStringLiteral(distance.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", distanceIdent, __gong__toRawStringLiteral(distance.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		doubleOrdered := []*Double{}
		for double := range stageSet.Stage.Doubles {
			doubleOrdered = append(doubleOrdered, double)
		}
		sort.Slice(doubleOrdered, func(i, j int) bool {
			return stageSet.Stage.Double_stagedOrder[doubleOrdered[i]] < stageSet.Stage.Double_stagedOrder[doubleOrdered[j]]
		})
		for _, double := range doubleOrdered {
			doubleIdent := "__stage_0" + double.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Double{Name: %s}).Stage(stageSet.Stage)", doubleIdent, __gong__toRawStringLiteral(double.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", doubleIdent, __gong__toRawStringLiteral(double.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Above = %s", doubleIdent, __gong__toRawStringLiteral(string(double.Above))))
		}
	}
	if stageSet.Stage != nil {
		dynamicsOrdered := []*Dynamics{}
		for dynamics := range stageSet.Stage.Dynamicss {
			dynamicsOrdered = append(dynamicsOrdered, dynamics)
		}
		sort.Slice(dynamicsOrdered, func(i, j int) bool {
			return stageSet.Stage.Dynamics_stagedOrder[dynamicsOrdered[i]] < stageSet.Stage.Dynamics_stagedOrder[dynamicsOrdered[j]]
		})
		for _, dynamics := range dynamicsOrdered {
			dynamicsIdent := "__stage_0" + dynamics.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Dynamics{Name: %s}).Stage(stageSet.Stage)", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Name)))
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
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_dynamics = append(%s.Other_dynamics, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		effectOrdered := []*Effect{}
		for effect := range stageSet.Stage.Effects {
			effectOrdered = append(effectOrdered, effect)
		}
		sort.Slice(effectOrdered, func(i, j int) bool {
			return stageSet.Stage.Effect_stagedOrder[effectOrdered[i]] < stageSet.Stage.Effect_stagedOrder[effectOrdered[j]]
		})
		for _, effect := range effectOrdered {
			effectIdent := "__stage_0" + effect.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Effect{Name: %s}).Stage(stageSet.Stage)", effectIdent, __gong__toRawStringLiteral(effect.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", effectIdent, __gong__toRawStringLiteral(effect.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", effectIdent, __gong__toRawStringLiteral(effect.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", effectIdent, __gong__toRawStringLiteral(effect.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		elisionOrdered := []*Elision{}
		for elision := range stageSet.Stage.Elisions {
			elisionOrdered = append(elisionOrdered, elision)
		}
		sort.Slice(elisionOrdered, func(i, j int) bool {
			return stageSet.Stage.Elision_stagedOrder[elisionOrdered[i]] < stageSet.Stage.Elision_stagedOrder[elisionOrdered[j]]
		})
		for _, elision := range elisionOrdered {
			elisionIdent := "__stage_0" + elision.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Elision{Name: %s}).Stage(stageSet.Stage)", elisionIdent, __gong__toRawStringLiteral(elision.Name)))
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
		emptyOrdered := []*Empty{}
		for empty := range stageSet.Stage.Emptys {
			emptyOrdered = append(emptyOrdered, empty)
		}
		sort.Slice(emptyOrdered, func(i, j int) bool {
			return stageSet.Stage.Empty_stagedOrder[emptyOrdered[i]] < stageSet.Stage.Empty_stagedOrder[emptyOrdered[j]]
		})
		for _, empty := range emptyOrdered {
			emptyIdent := "__stage_0" + empty.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Empty{Name: %s}).Stage(stageSet.Stage)", emptyIdent, __gong__toRawStringLiteral(empty.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", emptyIdent, __gong__toRawStringLiteral(empty.Name)))
		}
	}
	if stageSet.Stage != nil {
		empty_fontOrdered := []*Empty_font{}
		for empty_font := range stageSet.Stage.Empty_fonts {
			empty_fontOrdered = append(empty_fontOrdered, empty_font)
		}
		sort.Slice(empty_fontOrdered, func(i, j int) bool {
			return stageSet.Stage.Empty_font_stagedOrder[empty_fontOrdered[i]] < stageSet.Stage.Empty_font_stagedOrder[empty_fontOrdered[j]]
		})
		for _, empty_font := range empty_fontOrdered {
			empty_fontIdent := "__stage_0" + empty_font.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Empty_font{Name: %s}).Stage(stageSet.Stage)", empty_fontIdent, __gong__toRawStringLiteral(empty_font.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", empty_fontIdent, __gong__toRawStringLiteral(empty_font.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", empty_fontIdent, __gong__toRawStringLiteral(empty_font.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", empty_fontIdent, __gong__toRawStringLiteral(empty_font.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", empty_fontIdent, __gong__toRawStringLiteral(empty_font.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", empty_fontIdent, __gong__toRawStringLiteral(empty_font.Font_weight)))
		}
	}
	if stageSet.Stage != nil {
		empty_lineOrdered := []*Empty_line{}
		for empty_line := range stageSet.Stage.Empty_lines {
			empty_lineOrdered = append(empty_lineOrdered, empty_line)
		}
		sort.Slice(empty_lineOrdered, func(i, j int) bool {
			return stageSet.Stage.Empty_line_stagedOrder[empty_lineOrdered[i]] < stageSet.Stage.Empty_line_stagedOrder[empty_lineOrdered[j]]
		})
		for _, empty_line := range empty_lineOrdered {
			empty_lineIdent := "__stage_0" + empty_line.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Empty_line{Name: %s}).Stage(stageSet.Stage)", empty_lineIdent, __gong__toRawStringLiteral(empty_line.Name)))
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
		empty_placementOrdered := []*Empty_placement{}
		for empty_placement := range stageSet.Stage.Empty_placements {
			empty_placementOrdered = append(empty_placementOrdered, empty_placement)
		}
		sort.Slice(empty_placementOrdered, func(i, j int) bool {
			return stageSet.Stage.Empty_placement_stagedOrder[empty_placementOrdered[i]] < stageSet.Stage.Empty_placement_stagedOrder[empty_placementOrdered[j]]
		})
		for _, empty_placement := range empty_placementOrdered {
			empty_placementIdent := "__stage_0" + empty_placement.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Empty_placement{Name: %s}).Stage(stageSet.Stage)", empty_placementIdent, __gong__toRawStringLiteral(empty_placement.Name)))
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
		empty_placement_smuflOrdered := []*Empty_placement_smufl{}
		for empty_placement_smufl := range stageSet.Stage.Empty_placement_smufls {
			empty_placement_smuflOrdered = append(empty_placement_smuflOrdered, empty_placement_smufl)
		}
		sort.Slice(empty_placement_smuflOrdered, func(i, j int) bool {
			return stageSet.Stage.Empty_placement_smufl_stagedOrder[empty_placement_smuflOrdered[i]] < stageSet.Stage.Empty_placement_smufl_stagedOrder[empty_placement_smuflOrdered[j]]
		})
		for _, empty_placement_smufl := range empty_placement_smuflOrdered {
			empty_placement_smuflIdent := "__stage_0" + empty_placement_smufl.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Empty_placement_smufl{Name: %s}).Stage(stageSet.Stage)", empty_placement_smuflIdent, __gong__toRawStringLiteral(empty_placement_smufl.Name)))
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
		empty_print_object_style_alignOrdered := []*Empty_print_object_style_align{}
		for empty_print_object_style_align := range stageSet.Stage.Empty_print_object_style_aligns {
			empty_print_object_style_alignOrdered = append(empty_print_object_style_alignOrdered, empty_print_object_style_align)
		}
		sort.Slice(empty_print_object_style_alignOrdered, func(i, j int) bool {
			return stageSet.Stage.Empty_print_object_style_align_stagedOrder[empty_print_object_style_alignOrdered[i]] < stageSet.Stage.Empty_print_object_style_align_stagedOrder[empty_print_object_style_alignOrdered[j]]
		})
		for _, empty_print_object_style_align := range empty_print_object_style_alignOrdered {
			empty_print_object_style_alignIdent := "__stage_0" + empty_print_object_style_align.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Empty_print_object_style_align{Name: %s}).Stage(stageSet.Stage)", empty_print_object_style_alignIdent, __gong__toRawStringLiteral(empty_print_object_style_align.Name)))
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
		empty_print_styleOrdered := []*Empty_print_style{}
		for empty_print_style := range stageSet.Stage.Empty_print_styles {
			empty_print_styleOrdered = append(empty_print_styleOrdered, empty_print_style)
		}
		sort.Slice(empty_print_styleOrdered, func(i, j int) bool {
			return stageSet.Stage.Empty_print_style_stagedOrder[empty_print_styleOrdered[i]] < stageSet.Stage.Empty_print_style_stagedOrder[empty_print_styleOrdered[j]]
		})
		for _, empty_print_style := range empty_print_styleOrdered {
			empty_print_styleIdent := "__stage_0" + empty_print_style.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Empty_print_style{Name: %s}).Stage(stageSet.Stage)", empty_print_styleIdent, __gong__toRawStringLiteral(empty_print_style.Name)))
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
		empty_print_style_alignOrdered := []*Empty_print_style_align{}
		for empty_print_style_align := range stageSet.Stage.Empty_print_style_aligns {
			empty_print_style_alignOrdered = append(empty_print_style_alignOrdered, empty_print_style_align)
		}
		sort.Slice(empty_print_style_alignOrdered, func(i, j int) bool {
			return stageSet.Stage.Empty_print_style_align_stagedOrder[empty_print_style_alignOrdered[i]] < stageSet.Stage.Empty_print_style_align_stagedOrder[empty_print_style_alignOrdered[j]]
		})
		for _, empty_print_style_align := range empty_print_style_alignOrdered {
			empty_print_style_alignIdent := "__stage_0" + empty_print_style_align.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Empty_print_style_align{Name: %s}).Stage(stageSet.Stage)", empty_print_style_alignIdent, __gong__toRawStringLiteral(empty_print_style_align.Name)))
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
		empty_print_style_align_idOrdered := []*Empty_print_style_align_id{}
		for empty_print_style_align_id := range stageSet.Stage.Empty_print_style_align_ids {
			empty_print_style_align_idOrdered = append(empty_print_style_align_idOrdered, empty_print_style_align_id)
		}
		sort.Slice(empty_print_style_align_idOrdered, func(i, j int) bool {
			return stageSet.Stage.Empty_print_style_align_id_stagedOrder[empty_print_style_align_idOrdered[i]] < stageSet.Stage.Empty_print_style_align_id_stagedOrder[empty_print_style_align_idOrdered[j]]
		})
		for _, empty_print_style_align_id := range empty_print_style_align_idOrdered {
			empty_print_style_align_idIdent := "__stage_0" + empty_print_style_align_id.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Empty_print_style_align_id{Name: %s}).Stage(stageSet.Stage)", empty_print_style_align_idIdent, __gong__toRawStringLiteral(empty_print_style_align_id.Name)))
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
		empty_trill_soundOrdered := []*Empty_trill_sound{}
		for empty_trill_sound := range stageSet.Stage.Empty_trill_sounds {
			empty_trill_soundOrdered = append(empty_trill_soundOrdered, empty_trill_sound)
		}
		sort.Slice(empty_trill_soundOrdered, func(i, j int) bool {
			return stageSet.Stage.Empty_trill_sound_stagedOrder[empty_trill_soundOrdered[i]] < stageSet.Stage.Empty_trill_sound_stagedOrder[empty_trill_soundOrdered[j]]
		})
		for _, empty_trill_sound := range empty_trill_soundOrdered {
			empty_trill_soundIdent := "__stage_0" + empty_trill_sound.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Empty_trill_sound{Name: %s}).Stage(stageSet.Stage)", empty_trill_soundIdent, __gong__toRawStringLiteral(empty_trill_sound.Name)))
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
		encodingOrdered := []*Encoding{}
		for encoding := range stageSet.Stage.Encodings {
			encodingOrdered = append(encodingOrdered, encoding)
		}
		sort.Slice(encodingOrdered, func(i, j int) bool {
			return stageSet.Stage.Encoding_stagedOrder[encodingOrdered[i]] < stageSet.Stage.Encoding_stagedOrder[encodingOrdered[j]]
		})
		for _, encoding := range encodingOrdered {
			encodingIdent := "__stage_0" + encoding.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Encoding{Name: %s}).Stage(stageSet.Stage)", encodingIdent, __gong__toRawStringLiteral(encoding.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", encodingIdent, __gong__toRawStringLiteral(encoding.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Software = %s", encodingIdent, __gong__toRawStringLiteral(encoding.Software)))
			values.WriteString(fmt.Sprintf("\n\t%s.Encoding_description = %s", encodingIdent, __gong__toRawStringLiteral(encoding.Encoding_description)))
			for _, elem := range encoding.Encoder {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Encoder = append(%s.Encoder, %s)", encodingIdent, encodingIdent, targetIdent))
			}
			for _, elem := range encoding.Supports {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Supports = append(%s.Supports, %s)", encodingIdent, encodingIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		endingOrdered := []*Ending{}
		for ending := range stageSet.Stage.Endings {
			endingOrdered = append(endingOrdered, ending)
		}
		sort.Slice(endingOrdered, func(i, j int) bool {
			return stageSet.Stage.Ending_stagedOrder[endingOrdered[i]] < stageSet.Stage.Ending_stagedOrder[endingOrdered[j]]
		})
		for _, ending := range endingOrdered {
			endingIdent := "__stage_0" + ending.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Ending{Name: %s}).Stage(stageSet.Stage)", endingIdent, __gong__toRawStringLiteral(ending.Name)))
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
		extendOrdered := []*Extend{}
		for extend := range stageSet.Stage.Extends {
			extendOrdered = append(extendOrdered, extend)
		}
		sort.Slice(extendOrdered, func(i, j int) bool {
			return stageSet.Stage.Extend_stagedOrder[extendOrdered[i]] < stageSet.Stage.Extend_stagedOrder[extendOrdered[j]]
		})
		for _, extend := range extendOrdered {
			extendIdent := "__stage_0" + extend.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Extend{Name: %s}).Stage(stageSet.Stage)", extendIdent, __gong__toRawStringLiteral(extend.Name)))
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
		featureOrdered := []*Feature{}
		for feature := range stageSet.Stage.Features {
			featureOrdered = append(featureOrdered, feature)
		}
		sort.Slice(featureOrdered, func(i, j int) bool {
			return stageSet.Stage.Feature_stagedOrder[featureOrdered[i]] < stageSet.Stage.Feature_stagedOrder[featureOrdered[j]]
		})
		for _, feature := range featureOrdered {
			featureIdent := "__stage_0" + feature.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Feature{Name: %s}).Stage(stageSet.Stage)", featureIdent, __gong__toRawStringLiteral(feature.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", featureIdent, __gong__toRawStringLiteral(feature.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", featureIdent, __gong__toRawStringLiteral(feature.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", featureIdent, __gong__toRawStringLiteral(feature.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		fermataOrdered := []*Fermata{}
		for fermata := range stageSet.Stage.Fermatas {
			fermataOrdered = append(fermataOrdered, fermata)
		}
		sort.Slice(fermataOrdered, func(i, j int) bool {
			return stageSet.Stage.Fermata_stagedOrder[fermataOrdered[i]] < stageSet.Stage.Fermata_stagedOrder[fermataOrdered[j]]
		})
		for _, fermata := range fermataOrdered {
			fermataIdent := "__stage_0" + fermata.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Fermata{Name: %s}).Stage(stageSet.Stage)", fermataIdent, __gong__toRawStringLiteral(fermata.Name)))
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
		figureOrdered := []*Figure{}
		for figure := range stageSet.Stage.Figures {
			figureOrdered = append(figureOrdered, figure)
		}
		sort.Slice(figureOrdered, func(i, j int) bool {
			return stageSet.Stage.Figure_stagedOrder[figureOrdered[i]] < stageSet.Stage.Figure_stagedOrder[figureOrdered[j]]
		})
		for _, figure := range figureOrdered {
			figureIdent := "__stage_0" + figure.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Figure{Name: %s}).Stage(stageSet.Stage)", figureIdent, __gong__toRawStringLiteral(figure.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", figureIdent, __gong__toRawStringLiteral(figure.Name)))
			if figure.Prefix != nil {
				targetIdent := "__stage_0" + figure.Prefix.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Prefix = %s", figureIdent, targetIdent))
			}
			if figure.Figure_number != nil {
				targetIdent := "__stage_0" + figure.Figure_number.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Figure_number = %s", figureIdent, targetIdent))
			}
			if figure.Suffix != nil {
				targetIdent := "__stage_0" + figure.Suffix.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Suffix = %s", figureIdent, targetIdent))
			}
			if figure.Extend != nil {
				targetIdent := "__stage_0" + figure.Extend.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extend = %s", figureIdent, targetIdent))
			}
			if figure.Footnote != nil {
				targetIdent := "__stage_0" + figure.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", figureIdent, targetIdent))
			}
			if figure.Level != nil {
				targetIdent := "__stage_0" + figure.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", figureIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		figured_bassOrdered := []*Figured_bass{}
		for figured_bass := range stageSet.Stage.Figured_basss {
			figured_bassOrdered = append(figured_bassOrdered, figured_bass)
		}
		sort.Slice(figured_bassOrdered, func(i, j int) bool {
			return stageSet.Stage.Figured_bass_stagedOrder[figured_bassOrdered[i]] < stageSet.Stage.Figured_bass_stagedOrder[figured_bassOrdered[j]]
		})
		for _, figured_bass := range figured_bassOrdered {
			figured_bassIdent := "__stage_0" + figured_bass.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Figured_bass{Name: %s}).Stage(stageSet.Stage)", figured_bassIdent, __gong__toRawStringLiteral(figured_bass.Name)))
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
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Figure = append(%s.Figure, %s)", figured_bassIdent, figured_bassIdent, targetIdent))
			}
			if figured_bass.Footnote != nil {
				targetIdent := "__stage_0" + figured_bass.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", figured_bassIdent, targetIdent))
			}
			if figured_bass.Level != nil {
				targetIdent := "__stage_0" + figured_bass.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", figured_bassIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		fingeringOrdered := []*Fingering{}
		for fingering := range stageSet.Stage.Fingerings {
			fingeringOrdered = append(fingeringOrdered, fingering)
		}
		sort.Slice(fingeringOrdered, func(i, j int) bool {
			return stageSet.Stage.Fingering_stagedOrder[fingeringOrdered[i]] < stageSet.Stage.Fingering_stagedOrder[fingeringOrdered[j]]
		})
		for _, fingering := range fingeringOrdered {
			fingeringIdent := "__stage_0" + fingering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Fingering{Name: %s}).Stage(stageSet.Stage)", fingeringIdent, __gong__toRawStringLiteral(fingering.Name)))
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
		first_fretOrdered := []*First_fret{}
		for first_fret := range stageSet.Stage.First_frets {
			first_fretOrdered = append(first_fretOrdered, first_fret)
		}
		sort.Slice(first_fretOrdered, func(i, j int) bool {
			return stageSet.Stage.First_fret_stagedOrder[first_fretOrdered[i]] < stageSet.Stage.First_fret_stagedOrder[first_fretOrdered[j]]
		})
		for _, first_fret := range first_fretOrdered {
			first_fretIdent := "__stage_0" + first_fret.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.First_fret{Name: %s}).Stage(stageSet.Stage)", first_fretIdent, __gong__toRawStringLiteral(first_fret.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", first_fretIdent, __gong__toRawStringLiteral(first_fret.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", first_fretIdent, __gong__toRawStringLiteral(first_fret.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", first_fretIdent, __gong__toRawStringLiteral(first_fret.Location)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", first_fretIdent, first_fret.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		for_partOrdered := []*For_part{}
		for for_part := range stageSet.Stage.For_parts {
			for_partOrdered = append(for_partOrdered, for_part)
		}
		sort.Slice(for_partOrdered, func(i, j int) bool {
			return stageSet.Stage.For_part_stagedOrder[for_partOrdered[i]] < stageSet.Stage.For_part_stagedOrder[for_partOrdered[j]]
		})
		for _, for_part := range for_partOrdered {
			for_partIdent := "__stage_0" + for_part.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.For_part{Name: %s}).Stage(stageSet.Stage)", for_partIdent, __gong__toRawStringLiteral(for_part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", for_partIdent, __gong__toRawStringLiteral(for_part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", for_partIdent, for_part.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", for_partIdent, __gong__toRawStringLiteral(for_part.Id)))
			if for_part.Part_clef != nil {
				targetIdent := "__stage_0" + for_part.Part_clef.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_clef = %s", for_partIdent, targetIdent))
			}
			if for_part.Part_transpose != nil {
				targetIdent := "__stage_0" + for_part.Part_transpose.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_transpose = %s", for_partIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		formatted_symbolOrdered := []*Formatted_symbol{}
		for formatted_symbol := range stageSet.Stage.Formatted_symbols {
			formatted_symbolOrdered = append(formatted_symbolOrdered, formatted_symbol)
		}
		sort.Slice(formatted_symbolOrdered, func(i, j int) bool {
			return stageSet.Stage.Formatted_symbol_stagedOrder[formatted_symbolOrdered[i]] < stageSet.Stage.Formatted_symbol_stagedOrder[formatted_symbolOrdered[j]]
		})
		for _, formatted_symbol := range formatted_symbolOrdered {
			formatted_symbolIdent := "__stage_0" + formatted_symbol.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Formatted_symbol{Name: %s}).Stage(stageSet.Stage)", formatted_symbolIdent, __gong__toRawStringLiteral(formatted_symbol.Name)))
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
		formatted_symbol_idOrdered := []*Formatted_symbol_id{}
		for formatted_symbol_id := range stageSet.Stage.Formatted_symbol_ids {
			formatted_symbol_idOrdered = append(formatted_symbol_idOrdered, formatted_symbol_id)
		}
		sort.Slice(formatted_symbol_idOrdered, func(i, j int) bool {
			return stageSet.Stage.Formatted_symbol_id_stagedOrder[formatted_symbol_idOrdered[i]] < stageSet.Stage.Formatted_symbol_id_stagedOrder[formatted_symbol_idOrdered[j]]
		})
		for _, formatted_symbol_id := range formatted_symbol_idOrdered {
			formatted_symbol_idIdent := "__stage_0" + formatted_symbol_id.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Formatted_symbol_id{Name: %s}).Stage(stageSet.Stage)", formatted_symbol_idIdent, __gong__toRawStringLiteral(formatted_symbol_id.Name)))
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
		formatted_textOrdered := []*Formatted_text{}
		for formatted_text := range stageSet.Stage.Formatted_texts {
			formatted_textOrdered = append(formatted_textOrdered, formatted_text)
		}
		sort.Slice(formatted_textOrdered, func(i, j int) bool {
			return stageSet.Stage.Formatted_text_stagedOrder[formatted_textOrdered[i]] < stageSet.Stage.Formatted_text_stagedOrder[formatted_textOrdered[j]]
		})
		for _, formatted_text := range formatted_textOrdered {
			formatted_textIdent := "__stage_0" + formatted_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Formatted_text{Name: %s}).Stage(stageSet.Stage)", formatted_textIdent, __gong__toRawStringLiteral(formatted_text.Name)))
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
		formatted_text_idOrdered := []*Formatted_text_id{}
		for formatted_text_id := range stageSet.Stage.Formatted_text_ids {
			formatted_text_idOrdered = append(formatted_text_idOrdered, formatted_text_id)
		}
		sort.Slice(formatted_text_idOrdered, func(i, j int) bool {
			return stageSet.Stage.Formatted_text_id_stagedOrder[formatted_text_idOrdered[i]] < stageSet.Stage.Formatted_text_id_stagedOrder[formatted_text_idOrdered[j]]
		})
		for _, formatted_text_id := range formatted_text_idOrdered {
			formatted_text_idIdent := "__stage_0" + formatted_text_id.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Formatted_text_id{Name: %s}).Stage(stageSet.Stage)", formatted_text_idIdent, __gong__toRawStringLiteral(formatted_text_id.Name)))
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
		forwardOrdered := []*Forward{}
		for forward := range stageSet.Stage.Forwards {
			forwardOrdered = append(forwardOrdered, forward)
		}
		sort.Slice(forwardOrdered, func(i, j int) bool {
			return stageSet.Stage.Forward_stagedOrder[forwardOrdered[i]] < stageSet.Stage.Forward_stagedOrder[forwardOrdered[j]]
		})
		for _, forward := range forwardOrdered {
			forwardIdent := "__stage_0" + forward.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Forward{Name: %s}).Stage(stageSet.Stage)", forwardIdent, __gong__toRawStringLiteral(forward.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", forwardIdent, __gong__toRawStringLiteral(forward.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = %s", forwardIdent, __gong__toRawStringLiteral(forward.Duration)))
			values.WriteString(fmt.Sprintf("\n\t%s.Voice = %s", forwardIdent, __gong__toRawStringLiteral(forward.Voice)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff = %d", forwardIdent, forward.Staff))
			if forward.Footnote != nil {
				targetIdent := "__stage_0" + forward.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", forwardIdent, targetIdent))
			}
			if forward.Level != nil {
				targetIdent := "__stage_0" + forward.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", forwardIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		frameOrdered := []*Frame{}
		for frame := range stageSet.Stage.Frames {
			frameOrdered = append(frameOrdered, frame)
		}
		sort.Slice(frameOrdered, func(i, j int) bool {
			return stageSet.Stage.Frame_stagedOrder[frameOrdered[i]] < stageSet.Stage.Frame_stagedOrder[frameOrdered[j]]
		})
		for _, frame := range frameOrdered {
			frameIdent := "__stage_0" + frame.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Frame{Name: %s}).Stage(stageSet.Stage)", frameIdent, __gong__toRawStringLiteral(frame.Name)))
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
				targetIdent := "__stage_0" + frame.First_fret.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.First_fret = %s", frameIdent, targetIdent))
			}
			for _, elem := range frame.Frame_note {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frame_note = append(%s.Frame_note, %s)", frameIdent, frameIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		frame_noteOrdered := []*Frame_note{}
		for frame_note := range stageSet.Stage.Frame_notes {
			frame_noteOrdered = append(frame_noteOrdered, frame_note)
		}
		sort.Slice(frame_noteOrdered, func(i, j int) bool {
			return stageSet.Stage.Frame_note_stagedOrder[frame_noteOrdered[i]] < stageSet.Stage.Frame_note_stagedOrder[frame_noteOrdered[j]]
		})
		for _, frame_note := range frame_noteOrdered {
			frame_noteIdent := "__stage_0" + frame_note.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Frame_note{Name: %s}).Stage(stageSet.Stage)", frame_noteIdent, __gong__toRawStringLiteral(frame_note.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", frame_noteIdent, __gong__toRawStringLiteral(frame_note.Name)))
			if frame_note.String != nil {
				targetIdent := "__stage_0" + frame_note.String.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.String = %s", frame_noteIdent, targetIdent))
			}
			if frame_note.Fret != nil {
				targetIdent := "__stage_0" + frame_note.Fret.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fret = %s", frame_noteIdent, targetIdent))
			}
			if frame_note.Fingering != nil {
				targetIdent := "__stage_0" + frame_note.Fingering.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fingering = %s", frame_noteIdent, targetIdent))
			}
			if frame_note.Barre != nil {
				targetIdent := "__stage_0" + frame_note.Barre.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Barre = %s", frame_noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		fretOrdered := []*Fret{}
		for fret := range stageSet.Stage.Frets {
			fretOrdered = append(fretOrdered, fret)
		}
		sort.Slice(fretOrdered, func(i, j int) bool {
			return stageSet.Stage.Fret_stagedOrder[fretOrdered[i]] < stageSet.Stage.Fret_stagedOrder[fretOrdered[j]]
		})
		for _, fret := range fretOrdered {
			fretIdent := "__stage_0" + fret.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Fret{Name: %s}).Stage(stageSet.Stage)", fretIdent, __gong__toRawStringLiteral(fret.Name)))
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
		glassOrdered := []*Glass{}
		for glass := range stageSet.Stage.Glasss {
			glassOrdered = append(glassOrdered, glass)
		}
		sort.Slice(glassOrdered, func(i, j int) bool {
			return stageSet.Stage.Glass_stagedOrder[glassOrdered[i]] < stageSet.Stage.Glass_stagedOrder[glassOrdered[j]]
		})
		for _, glass := range glassOrdered {
			glassIdent := "__stage_0" + glass.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Glass{Name: %s}).Stage(stageSet.Stage)", glassIdent, __gong__toRawStringLiteral(glass.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glassIdent, __gong__toRawStringLiteral(glass.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", glassIdent, __gong__toRawStringLiteral(glass.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", glassIdent, __gong__toRawStringLiteral(glass.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		glissandoOrdered := []*Glissando{}
		for glissando := range stageSet.Stage.Glissandos {
			glissandoOrdered = append(glissandoOrdered, glissando)
		}
		sort.Slice(glissandoOrdered, func(i, j int) bool {
			return stageSet.Stage.Glissando_stagedOrder[glissandoOrdered[i]] < stageSet.Stage.Glissando_stagedOrder[glissandoOrdered[j]]
		})
		for _, glissando := range glissandoOrdered {
			glissandoIdent := "__stage_0" + glissando.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Glissando{Name: %s}).Stage(stageSet.Stage)", glissandoIdent, __gong__toRawStringLiteral(glissando.Name)))
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
		glyphOrdered := []*Glyph{}
		for glyph := range stageSet.Stage.Glyphs {
			glyphOrdered = append(glyphOrdered, glyph)
		}
		sort.Slice(glyphOrdered, func(i, j int) bool {
			return stageSet.Stage.Glyph_stagedOrder[glyphOrdered[i]] < stageSet.Stage.Glyph_stagedOrder[glyphOrdered[j]]
		})
		for _, glyph := range glyphOrdered {
			glyphIdent := "__stage_0" + glyph.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Glyph{Name: %s}).Stage(stageSet.Stage)", glyphIdent, __gong__toRawStringLiteral(glyph.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glyphIdent, __gong__toRawStringLiteral(glyph.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", glyphIdent, __gong__toRawStringLiteral(glyph.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", glyphIdent, __gong__toRawStringLiteral(glyph.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		graceOrdered := []*Grace{}
		for grace := range stageSet.Stage.Graces {
			graceOrdered = append(graceOrdered, grace)
		}
		sort.Slice(graceOrdered, func(i, j int) bool {
			return stageSet.Stage.Grace_stagedOrder[graceOrdered[i]] < stageSet.Stage.Grace_stagedOrder[graceOrdered[j]]
		})
		for _, grace := range graceOrdered {
			graceIdent := "__stage_0" + grace.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Grace{Name: %s}).Stage(stageSet.Stage)", graceIdent, __gong__toRawStringLiteral(grace.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", graceIdent, __gong__toRawStringLiteral(grace.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Steal_time_previous = %s", graceIdent, __gong__toRawStringLiteral(grace.Steal_time_previous)))
			values.WriteString(fmt.Sprintf("\n\t%s.Steal_time_following = %s", graceIdent, __gong__toRawStringLiteral(grace.Steal_time_following)))
			values.WriteString(fmt.Sprintf("\n\t%s.Make_time = %s", graceIdent, __gong__toRawStringLiteral(grace.Make_time)))
			values.WriteString(fmt.Sprintf("\n\t%s.Slash = %s", graceIdent, __gong__toRawStringLiteral(string(grace.Slash))))
		}
	}
	if stageSet.Stage != nil {
		group_barlineOrdered := []*Group_barline{}
		for group_barline := range stageSet.Stage.Group_barlines {
			group_barlineOrdered = append(group_barlineOrdered, group_barline)
		}
		sort.Slice(group_barlineOrdered, func(i, j int) bool {
			return stageSet.Stage.Group_barline_stagedOrder[group_barlineOrdered[i]] < stageSet.Stage.Group_barline_stagedOrder[group_barlineOrdered[j]]
		})
		for _, group_barline := range group_barlineOrdered {
			group_barlineIdent := "__stage_0" + group_barline.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Group_barline{Name: %s}).Stage(stageSet.Stage)", group_barlineIdent, __gong__toRawStringLiteral(group_barline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", group_barlineIdent, __gong__toRawStringLiteral(group_barline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", group_barlineIdent, __gong__toRawStringLiteral(group_barline.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", group_barlineIdent, __gong__toRawStringLiteral(group_barline.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		group_nameOrdered := []*Group_name{}
		for group_name := range stageSet.Stage.Group_names {
			group_nameOrdered = append(group_nameOrdered, group_name)
		}
		sort.Slice(group_nameOrdered, func(i, j int) bool {
			return stageSet.Stage.Group_name_stagedOrder[group_nameOrdered[i]] < stageSet.Stage.Group_name_stagedOrder[group_nameOrdered[j]]
		})
		for _, group_name := range group_nameOrdered {
			group_nameIdent := "__stage_0" + group_name.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Group_name{Name: %s}).Stage(stageSet.Stage)", group_nameIdent, __gong__toRawStringLiteral(group_name.Name)))
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
		group_symbolOrdered := []*Group_symbol{}
		for group_symbol := range stageSet.Stage.Group_symbols {
			group_symbolOrdered = append(group_symbolOrdered, group_symbol)
		}
		sort.Slice(group_symbolOrdered, func(i, j int) bool {
			return stageSet.Stage.Group_symbol_stagedOrder[group_symbolOrdered[i]] < stageSet.Stage.Group_symbol_stagedOrder[group_symbolOrdered[j]]
		})
		for _, group_symbol := range group_symbolOrdered {
			group_symbolIdent := "__stage_0" + group_symbol.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Group_symbol{Name: %s}).Stage(stageSet.Stage)", group_symbolIdent, __gong__toRawStringLiteral(group_symbol.Name)))
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
		groupingOrdered := []*Grouping{}
		for grouping := range stageSet.Stage.Groupings {
			groupingOrdered = append(groupingOrdered, grouping)
		}
		sort.Slice(groupingOrdered, func(i, j int) bool {
			return stageSet.Stage.Grouping_stagedOrder[groupingOrdered[i]] < stageSet.Stage.Grouping_stagedOrder[groupingOrdered[j]]
		})
		for _, grouping := range groupingOrdered {
			groupingIdent := "__stage_0" + grouping.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Grouping{Name: %s}).Stage(stageSet.Stage)", groupingIdent, __gong__toRawStringLiteral(grouping.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Member_of = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Member_of)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Id)))
			for _, elem := range grouping.Feature {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Feature = append(%s.Feature, %s)", groupingIdent, groupingIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		hammer_on_pull_offOrdered := []*Hammer_on_pull_off{}
		for hammer_on_pull_off := range stageSet.Stage.Hammer_on_pull_offs {
			hammer_on_pull_offOrdered = append(hammer_on_pull_offOrdered, hammer_on_pull_off)
		}
		sort.Slice(hammer_on_pull_offOrdered, func(i, j int) bool {
			return stageSet.Stage.Hammer_on_pull_off_stagedOrder[hammer_on_pull_offOrdered[i]] < stageSet.Stage.Hammer_on_pull_off_stagedOrder[hammer_on_pull_offOrdered[j]]
		})
		for _, hammer_on_pull_off := range hammer_on_pull_offOrdered {
			hammer_on_pull_offIdent := "__stage_0" + hammer_on_pull_off.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Hammer_on_pull_off{Name: %s}).Stage(stageSet.Stage)", hammer_on_pull_offIdent, __gong__toRawStringLiteral(hammer_on_pull_off.Name)))
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
		handbellOrdered := []*Handbell{}
		for handbell := range stageSet.Stage.Handbells {
			handbellOrdered = append(handbellOrdered, handbell)
		}
		sort.Slice(handbellOrdered, func(i, j int) bool {
			return stageSet.Stage.Handbell_stagedOrder[handbellOrdered[i]] < stageSet.Stage.Handbell_stagedOrder[handbellOrdered[j]]
		})
		for _, handbell := range handbellOrdered {
			handbellIdent := "__stage_0" + handbell.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Handbell{Name: %s}).Stage(stageSet.Stage)", handbellIdent, __gong__toRawStringLiteral(handbell.Name)))
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
		harmon_closedOrdered := []*Harmon_closed{}
		for harmon_closed := range stageSet.Stage.Harmon_closeds {
			harmon_closedOrdered = append(harmon_closedOrdered, harmon_closed)
		}
		sort.Slice(harmon_closedOrdered, func(i, j int) bool {
			return stageSet.Stage.Harmon_closed_stagedOrder[harmon_closedOrdered[i]] < stageSet.Stage.Harmon_closed_stagedOrder[harmon_closedOrdered[j]]
		})
		for _, harmon_closed := range harmon_closedOrdered {
			harmon_closedIdent := "__stage_0" + harmon_closed.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Harmon_closed{Name: %s}).Stage(stageSet.Stage)", harmon_closedIdent, __gong__toRawStringLiteral(harmon_closed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harmon_closedIdent, __gong__toRawStringLiteral(harmon_closed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", harmon_closedIdent, __gong__toRawStringLiteral(harmon_closed.Location)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", harmon_closedIdent, __gong__toRawStringLiteral(harmon_closed.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		harmon_muteOrdered := []*Harmon_mute{}
		for harmon_mute := range stageSet.Stage.Harmon_mutes {
			harmon_muteOrdered = append(harmon_muteOrdered, harmon_mute)
		}
		sort.Slice(harmon_muteOrdered, func(i, j int) bool {
			return stageSet.Stage.Harmon_mute_stagedOrder[harmon_muteOrdered[i]] < stageSet.Stage.Harmon_mute_stagedOrder[harmon_muteOrdered[j]]
		})
		for _, harmon_mute := range harmon_muteOrdered {
			harmon_muteIdent := "__stage_0" + harmon_mute.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Harmon_mute{Name: %s}).Stage(stageSet.Stage)", harmon_muteIdent, __gong__toRawStringLiteral(harmon_mute.Name)))
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
				targetIdent := "__stage_0" + harmon_mute.Harmon_closed.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmon_closed = %s", harmon_muteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		harmonicOrdered := []*Harmonic{}
		for harmonic := range stageSet.Stage.Harmonics {
			harmonicOrdered = append(harmonicOrdered, harmonic)
		}
		sort.Slice(harmonicOrdered, func(i, j int) bool {
			return stageSet.Stage.Harmonic_stagedOrder[harmonicOrdered[i]] < stageSet.Stage.Harmonic_stagedOrder[harmonicOrdered[j]]
		})
		for _, harmonic := range harmonicOrdered {
			harmonicIdent := "__stage_0" + harmonic.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Harmonic{Name: %s}).Stage(stageSet.Stage)", harmonicIdent, __gong__toRawStringLiteral(harmonic.Name)))
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
		harmonyOrdered := []*Harmony{}
		for harmony := range stageSet.Stage.Harmonys {
			harmonyOrdered = append(harmonyOrdered, harmony)
		}
		sort.Slice(harmonyOrdered, func(i, j int) bool {
			return stageSet.Stage.Harmony_stagedOrder[harmonyOrdered[i]] < stageSet.Stage.Harmony_stagedOrder[harmonyOrdered[j]]
		})
		for _, harmony := range harmonyOrdered {
			harmonyIdent := "__stage_0" + harmony.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Harmony{Name: %s}).Stage(stageSet.Stage)", harmonyIdent, __gong__toRawStringLiteral(harmony.Name)))
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
				targetIdent := "__stage_0" + harmony.Root.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Root = %s", harmonyIdent, targetIdent))
			}
			if harmony.Numeral != nil {
				targetIdent := "__stage_0" + harmony.Numeral.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Numeral = %s", harmonyIdent, targetIdent))
			}
			if harmony.Function != nil {
				targetIdent := "__stage_0" + harmony.Function.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Function = %s", harmonyIdent, targetIdent))
			}
			if harmony.Kind != nil {
				targetIdent := "__stage_0" + harmony.Kind.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Kind = %s", harmonyIdent, targetIdent))
			}
			if harmony.Inversion != nil {
				targetIdent := "__stage_0" + harmony.Inversion.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inversion = %s", harmonyIdent, targetIdent))
			}
			if harmony.Bass != nil {
				targetIdent := "__stage_0" + harmony.Bass.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bass = %s", harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Degree {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Degree = append(%s.Degree, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			if harmony.Frame != nil {
				targetIdent := "__stage_0" + harmony.Frame.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frame = %s", harmonyIdent, targetIdent))
			}
			if harmony.Offset != nil {
				targetIdent := "__stage_0" + harmony.Offset.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", harmonyIdent, targetIdent))
			}
			if harmony.Footnote != nil {
				targetIdent := "__stage_0" + harmony.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", harmonyIdent, targetIdent))
			}
			if harmony.Level != nil {
				targetIdent := "__stage_0" + harmony.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", harmonyIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		harmony_alterOrdered := []*Harmony_alter{}
		for harmony_alter := range stageSet.Stage.Harmony_alters {
			harmony_alterOrdered = append(harmony_alterOrdered, harmony_alter)
		}
		sort.Slice(harmony_alterOrdered, func(i, j int) bool {
			return stageSet.Stage.Harmony_alter_stagedOrder[harmony_alterOrdered[i]] < stageSet.Stage.Harmony_alter_stagedOrder[harmony_alterOrdered[j]]
		})
		for _, harmony_alter := range harmony_alterOrdered {
			harmony_alterIdent := "__stage_0" + harmony_alter.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Harmony_alter{Name: %s}).Stage(stageSet.Stage)", harmony_alterIdent, __gong__toRawStringLiteral(harmony_alter.Name)))
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
		harp_pedalsOrdered := []*Harp_pedals{}
		for harp_pedals := range stageSet.Stage.Harp_pedalss {
			harp_pedalsOrdered = append(harp_pedalsOrdered, harp_pedals)
		}
		sort.Slice(harp_pedalsOrdered, func(i, j int) bool {
			return stageSet.Stage.Harp_pedals_stagedOrder[harp_pedalsOrdered[i]] < stageSet.Stage.Harp_pedals_stagedOrder[harp_pedalsOrdered[j]]
		})
		for _, harp_pedals := range harp_pedalsOrdered {
			harp_pedalsIdent := "__stage_0" + harp_pedals.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Harp_pedals{Name: %s}).Stage(stageSet.Stage)", harp_pedalsIdent, __gong__toRawStringLiteral(harp_pedals.Name)))
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
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pedal_tuning = append(%s.Pedal_tuning, %s)", harp_pedalsIdent, harp_pedalsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		heel_toeOrdered := []*Heel_toe{}
		for heel_toe := range stageSet.Stage.Heel_toes {
			heel_toeOrdered = append(heel_toeOrdered, heel_toe)
		}
		sort.Slice(heel_toeOrdered, func(i, j int) bool {
			return stageSet.Stage.Heel_toe_stagedOrder[heel_toeOrdered[i]] < stageSet.Stage.Heel_toe_stagedOrder[heel_toeOrdered[j]]
		})
		for _, heel_toe := range heel_toeOrdered {
			heel_toeIdent := "__stage_0" + heel_toe.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Heel_toe{Name: %s}).Stage(stageSet.Stage)", heel_toeIdent, __gong__toRawStringLiteral(heel_toe.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", heel_toeIdent, __gong__toRawStringLiteral(heel_toe.Name)))
		}
	}
	if stageSet.Stage != nil {
		holeOrdered := []*Hole{}
		for hole := range stageSet.Stage.Holes {
			holeOrdered = append(holeOrdered, hole)
		}
		sort.Slice(holeOrdered, func(i, j int) bool {
			return stageSet.Stage.Hole_stagedOrder[holeOrdered[i]] < stageSet.Stage.Hole_stagedOrder[holeOrdered[j]]
		})
		for _, hole := range holeOrdered {
			holeIdent := "__stage_0" + hole.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Hole{Name: %s}).Stage(stageSet.Stage)", holeIdent, __gong__toRawStringLiteral(hole.Name)))
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
				targetIdent := "__stage_0" + hole.Hole_closed.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Hole_closed = %s", holeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		hole_closedOrdered := []*Hole_closed{}
		for hole_closed := range stageSet.Stage.Hole_closeds {
			hole_closedOrdered = append(hole_closedOrdered, hole_closed)
		}
		sort.Slice(hole_closedOrdered, func(i, j int) bool {
			return stageSet.Stage.Hole_closed_stagedOrder[hole_closedOrdered[i]] < stageSet.Stage.Hole_closed_stagedOrder[hole_closedOrdered[j]]
		})
		for _, hole_closed := range hole_closedOrdered {
			hole_closedIdent := "__stage_0" + hole_closed.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Hole_closed{Name: %s}).Stage(stageSet.Stage)", hole_closedIdent, __gong__toRawStringLiteral(hole_closed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", hole_closedIdent, __gong__toRawStringLiteral(hole_closed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", hole_closedIdent, __gong__toRawStringLiteral(hole_closed.Location)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", hole_closedIdent, __gong__toRawStringLiteral(hole_closed.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		horizontal_turnOrdered := []*Horizontal_turn{}
		for horizontal_turn := range stageSet.Stage.Horizontal_turns {
			horizontal_turnOrdered = append(horizontal_turnOrdered, horizontal_turn)
		}
		sort.Slice(horizontal_turnOrdered, func(i, j int) bool {
			return stageSet.Stage.Horizontal_turn_stagedOrder[horizontal_turnOrdered[i]] < stageSet.Stage.Horizontal_turn_stagedOrder[horizontal_turnOrdered[j]]
		})
		for _, horizontal_turn := range horizontal_turnOrdered {
			horizontal_turnIdent := "__stage_0" + horizontal_turn.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Horizontal_turn{Name: %s}).Stage(stageSet.Stage)", horizontal_turnIdent, __gong__toRawStringLiteral(horizontal_turn.Name)))
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
		identificationOrdered := []*Identification{}
		for identification := range stageSet.Stage.Identifications {
			identificationOrdered = append(identificationOrdered, identification)
		}
		sort.Slice(identificationOrdered, func(i, j int) bool {
			return stageSet.Stage.Identification_stagedOrder[identificationOrdered[i]] < stageSet.Stage.Identification_stagedOrder[identificationOrdered[j]]
		})
		for _, identification := range identificationOrdered {
			identificationIdent := "__stage_0" + identification.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Identification{Name: %s}).Stage(stageSet.Stage)", identificationIdent, __gong__toRawStringLiteral(identification.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", identificationIdent, __gong__toRawStringLiteral(identification.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Source = %s", identificationIdent, __gong__toRawStringLiteral(identification.Source)))
			for _, elem := range identification.Creator {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Creator = append(%s.Creator, %s)", identificationIdent, identificationIdent, targetIdent))
			}
			for _, elem := range identification.Rights {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rights = append(%s.Rights, %s)", identificationIdent, identificationIdent, targetIdent))
			}
			if identification.Encoding != nil {
				targetIdent := "__stage_0" + identification.Encoding.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Encoding = %s", identificationIdent, targetIdent))
			}
			for _, elem := range identification.Relation {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Relation = append(%s.Relation, %s)", identificationIdent, identificationIdent, targetIdent))
			}
			if identification.Miscellaneous != nil {
				targetIdent := "__stage_0" + identification.Miscellaneous.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Miscellaneous = %s", identificationIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		imageOrdered := []*Image{}
		for image := range stageSet.Stage.Images {
			imageOrdered = append(imageOrdered, image)
		}
		sort.Slice(imageOrdered, func(i, j int) bool {
			return stageSet.Stage.Image_stagedOrder[imageOrdered[i]] < stageSet.Stage.Image_stagedOrder[imageOrdered[j]]
		})
		for _, image := range imageOrdered {
			imageIdent := "__stage_0" + image.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Image{Name: %s}).Stage(stageSet.Stage)", imageIdent, __gong__toRawStringLiteral(image.Name)))
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
		instrumentOrdered := []*Instrument{}
		for instrument := range stageSet.Stage.Instruments {
			instrumentOrdered = append(instrumentOrdered, instrument)
		}
		sort.Slice(instrumentOrdered, func(i, j int) bool {
			return stageSet.Stage.Instrument_stagedOrder[instrumentOrdered[i]] < stageSet.Stage.Instrument_stagedOrder[instrumentOrdered[j]]
		})
		for _, instrument := range instrumentOrdered {
			instrumentIdent := "__stage_0" + instrument.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Instrument{Name: %s}).Stage(stageSet.Stage)", instrumentIdent, __gong__toRawStringLiteral(instrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrumentIdent, __gong__toRawStringLiteral(instrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", instrumentIdent, __gong__toRawStringLiteral(instrument.Id)))
		}
	}
	if stageSet.Stage != nil {
		instrument_changeOrdered := []*Instrument_change{}
		for instrument_change := range stageSet.Stage.Instrument_changes {
			instrument_changeOrdered = append(instrument_changeOrdered, instrument_change)
		}
		sort.Slice(instrument_changeOrdered, func(i, j int) bool {
			return stageSet.Stage.Instrument_change_stagedOrder[instrument_changeOrdered[i]] < stageSet.Stage.Instrument_change_stagedOrder[instrument_changeOrdered[j]]
		})
		for _, instrument_change := range instrument_changeOrdered {
			instrument_changeIdent := "__stage_0" + instrument_change.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Instrument_change{Name: %s}).Stage(stageSet.Stage)", instrument_changeIdent, __gong__toRawStringLiteral(instrument_change.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrument_changeIdent, __gong__toRawStringLiteral(instrument_change.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", instrument_changeIdent, __gong__toRawStringLiteral(instrument_change.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Instrument_sound = %s", instrument_changeIdent, __gong__toRawStringLiteral(instrument_change.Instrument_sound)))
			values.WriteString(fmt.Sprintf("\n\t%s.Solo = %s", instrument_changeIdent, __gong__toRawStringLiteral(instrument_change.Solo)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ensemble = %s", instrument_changeIdent, __gong__toRawStringLiteral(instrument_change.Ensemble)))
			if instrument_change.Virtual_instrument != nil {
				targetIdent := "__stage_0" + instrument_change.Virtual_instrument.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Virtual_instrument = %s", instrument_changeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		instrument_linkOrdered := []*Instrument_link{}
		for instrument_link := range stageSet.Stage.Instrument_links {
			instrument_linkOrdered = append(instrument_linkOrdered, instrument_link)
		}
		sort.Slice(instrument_linkOrdered, func(i, j int) bool {
			return stageSet.Stage.Instrument_link_stagedOrder[instrument_linkOrdered[i]] < stageSet.Stage.Instrument_link_stagedOrder[instrument_linkOrdered[j]]
		})
		for _, instrument_link := range instrument_linkOrdered {
			instrument_linkIdent := "__stage_0" + instrument_link.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Instrument_link{Name: %s}).Stage(stageSet.Stage)", instrument_linkIdent, __gong__toRawStringLiteral(instrument_link.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrument_linkIdent, __gong__toRawStringLiteral(instrument_link.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", instrument_linkIdent, __gong__toRawStringLiteral(instrument_link.Id)))
		}
	}
	if stageSet.Stage != nil {
		interchangeableOrdered := []*Interchangeable{}
		for interchangeable := range stageSet.Stage.Interchangeables {
			interchangeableOrdered = append(interchangeableOrdered, interchangeable)
		}
		sort.Slice(interchangeableOrdered, func(i, j int) bool {
			return stageSet.Stage.Interchangeable_stagedOrder[interchangeableOrdered[i]] < stageSet.Stage.Interchangeable_stagedOrder[interchangeableOrdered[j]]
		})
		for _, interchangeable := range interchangeableOrdered {
			interchangeableIdent := "__stage_0" + interchangeable.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Interchangeable{Name: %s}).Stage(stageSet.Stage)", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Symbol = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Symbol)))
			values.WriteString(fmt.Sprintf("\n\t%s.Separator = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Separator)))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_relation = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Time_relation)))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.Beat_type = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Beat_type)))
		}
	}
	if stageSet.Stage != nil {
		inversionOrdered := []*Inversion{}
		for inversion := range stageSet.Stage.Inversions {
			inversionOrdered = append(inversionOrdered, inversion)
		}
		sort.Slice(inversionOrdered, func(i, j int) bool {
			return stageSet.Stage.Inversion_stagedOrder[inversionOrdered[i]] < stageSet.Stage.Inversion_stagedOrder[inversionOrdered[j]]
		})
		for _, inversion := range inversionOrdered {
			inversionIdent := "__stage_0" + inversion.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Inversion{Name: %s}).Stage(stageSet.Stage)", inversionIdent, __gong__toRawStringLiteral(inversion.Name)))
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
		keyOrdered := []*Key{}
		for key := range stageSet.Stage.Keys {
			keyOrdered = append(keyOrdered, key)
		}
		sort.Slice(keyOrdered, func(i, j int) bool {
			return stageSet.Stage.Key_stagedOrder[keyOrdered[i]] < stageSet.Stage.Key_stagedOrder[keyOrdered[j]]
		})
		for _, key := range keyOrdered {
			keyIdent := "__stage_0" + key.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Key{Name: %s}).Stage(stageSet.Stage)", keyIdent, __gong__toRawStringLiteral(key.Name)))
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
				targetIdent := "__stage_0" + key.Cancel.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Cancel = %s", keyIdent, targetIdent))
			}
			if key.Key_accidental != nil {
				targetIdent := "__stage_0" + key.Key_accidental.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Key_accidental = %s", keyIdent, targetIdent))
			}
			for _, elem := range key.Key_octave {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Key_octave = append(%s.Key_octave, %s)", keyIdent, keyIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		key_accidentalOrdered := []*Key_accidental{}
		for key_accidental := range stageSet.Stage.Key_accidentals {
			key_accidentalOrdered = append(key_accidentalOrdered, key_accidental)
		}
		sort.Slice(key_accidentalOrdered, func(i, j int) bool {
			return stageSet.Stage.Key_accidental_stagedOrder[key_accidentalOrdered[i]] < stageSet.Stage.Key_accidental_stagedOrder[key_accidentalOrdered[j]]
		})
		for _, key_accidental := range key_accidentalOrdered {
			key_accidentalIdent := "__stage_0" + key_accidental.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Key_accidental{Name: %s}).Stage(stageSet.Stage)", key_accidentalIdent, __gong__toRawStringLiteral(key_accidental.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", key_accidentalIdent, __gong__toRawStringLiteral(key_accidental.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", key_accidentalIdent, __gong__toRawStringLiteral(key_accidental.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", key_accidentalIdent, __gong__toRawStringLiteral(string(key_accidental.EnclosedText))))
		}
	}
	if stageSet.Stage != nil {
		key_octaveOrdered := []*Key_octave{}
		for key_octave := range stageSet.Stage.Key_octaves {
			key_octaveOrdered = append(key_octaveOrdered, key_octave)
		}
		sort.Slice(key_octaveOrdered, func(i, j int) bool {
			return stageSet.Stage.Key_octave_stagedOrder[key_octaveOrdered[i]] < stageSet.Stage.Key_octave_stagedOrder[key_octaveOrdered[j]]
		})
		for _, key_octave := range key_octaveOrdered {
			key_octaveIdent := "__stage_0" + key_octave.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Key_octave{Name: %s}).Stage(stageSet.Stage)", key_octaveIdent, __gong__toRawStringLiteral(key_octave.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", key_octaveIdent, __gong__toRawStringLiteral(key_octave.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", key_octaveIdent, key_octave.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Cancel = %s", key_octaveIdent, __gong__toRawStringLiteral(string(key_octave.Cancel))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", key_octaveIdent, key_octave.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		kindOrdered := []*Kind{}
		for kind := range stageSet.Stage.Kinds {
			kindOrdered = append(kindOrdered, kind)
		}
		sort.Slice(kindOrdered, func(i, j int) bool {
			return stageSet.Stage.Kind_stagedOrder[kindOrdered[i]] < stageSet.Stage.Kind_stagedOrder[kindOrdered[j]]
		})
		for _, kind := range kindOrdered {
			kindIdent := "__stage_0" + kind.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Kind{Name: %s}).Stage(stageSet.Stage)", kindIdent, __gong__toRawStringLiteral(kind.Name)))
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
		levelOrdered := []*Level{}
		for level := range stageSet.Stage.Levels {
			levelOrdered = append(levelOrdered, level)
		}
		sort.Slice(levelOrdered, func(i, j int) bool {
			return stageSet.Stage.Level_stagedOrder[levelOrdered[i]] < stageSet.Stage.Level_stagedOrder[levelOrdered[j]]
		})
		for _, level := range levelOrdered {
			levelIdent := "__stage_0" + level.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Level{Name: %s}).Stage(stageSet.Stage)", levelIdent, __gong__toRawStringLiteral(level.Name)))
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
		line_detailOrdered := []*Line_detail{}
		for line_detail := range stageSet.Stage.Line_details {
			line_detailOrdered = append(line_detailOrdered, line_detail)
		}
		sort.Slice(line_detailOrdered, func(i, j int) bool {
			return stageSet.Stage.Line_detail_stagedOrder[line_detailOrdered[i]] < stageSet.Stage.Line_detail_stagedOrder[line_detailOrdered[j]]
		})
		for _, line_detail := range line_detailOrdered {
			line_detailIdent := "__stage_0" + line_detail.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Line_detail{Name: %s}).Stage(stageSet.Stage)", line_detailIdent, __gong__toRawStringLiteral(line_detail.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", line_detailIdent, __gong__toRawStringLiteral(line_detail.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", line_detailIdent, line_detail.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", line_detailIdent, __gong__toRawStringLiteral(line_detail.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", line_detailIdent, __gong__toRawStringLiteral(line_detail.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line_type = %s", line_detailIdent, __gong__toRawStringLiteral(line_detail.Line_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", line_detailIdent, __gong__toRawStringLiteral(string(line_detail.Print_object))))
		}
	}
	if stageSet.Stage != nil {
		line_widthOrdered := []*Line_width{}
		for line_width := range stageSet.Stage.Line_widths {
			line_widthOrdered = append(line_widthOrdered, line_width)
		}
		sort.Slice(line_widthOrdered, func(i, j int) bool {
			return stageSet.Stage.Line_width_stagedOrder[line_widthOrdered[i]] < stageSet.Stage.Line_width_stagedOrder[line_widthOrdered[j]]
		})
		for _, line_width := range line_widthOrdered {
			line_widthIdent := "__stage_0" + line_width.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Line_width{Name: %s}).Stage(stageSet.Stage)", line_widthIdent, __gong__toRawStringLiteral(line_width.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", line_widthIdent, __gong__toRawStringLiteral(line_width.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", line_widthIdent, __gong__toRawStringLiteral(line_width.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", line_widthIdent, __gong__toRawStringLiteral(line_width.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		linkOrdered := []*Link{}
		for link := range stageSet.Stage.Links {
			linkOrdered = append(linkOrdered, link)
		}
		sort.Slice(linkOrdered, func(i, j int) bool {
			return stageSet.Stage.Link_stagedOrder[linkOrdered[i]] < stageSet.Stage.Link_stagedOrder[linkOrdered[j]]
		})
		for _, link := range linkOrdered {
			linkIdent := "__stage_0" + link.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Link{Name: %s}).Stage(stageSet.Stage)", linkIdent, __gong__toRawStringLiteral(link.Name)))
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
		listenOrdered := []*Listen{}
		for listen := range stageSet.Stage.Listens {
			listenOrdered = append(listenOrdered, listen)
		}
		sort.Slice(listenOrdered, func(i, j int) bool {
			return stageSet.Stage.Listen_stagedOrder[listenOrdered[i]] < stageSet.Stage.Listen_stagedOrder[listenOrdered[j]]
		})
		for _, listen := range listenOrdered {
			listenIdent := "__stage_0" + listen.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Listen{Name: %s}).Stage(stageSet.Stage)", listenIdent, __gong__toRawStringLiteral(listen.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", listenIdent, __gong__toRawStringLiteral(listen.Name)))
			for _, elem := range listen.Assess {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Assess = append(%s.Assess, %s)", listenIdent, listenIdent, targetIdent))
			}
			for _, elem := range listen.Wait {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wait = append(%s.Wait, %s)", listenIdent, listenIdent, targetIdent))
			}
			for _, elem := range listen.Other_listen {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_listen = append(%s.Other_listen, %s)", listenIdent, listenIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		listeningOrdered := []*Listening{}
		for listening := range stageSet.Stage.Listenings {
			listeningOrdered = append(listeningOrdered, listening)
		}
		sort.Slice(listeningOrdered, func(i, j int) bool {
			return stageSet.Stage.Listening_stagedOrder[listeningOrdered[i]] < stageSet.Stage.Listening_stagedOrder[listeningOrdered[j]]
		})
		for _, listening := range listeningOrdered {
			listeningIdent := "__stage_0" + listening.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Listening{Name: %s}).Stage(stageSet.Stage)", listeningIdent, __gong__toRawStringLiteral(listening.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", listeningIdent, __gong__toRawStringLiteral(listening.Name)))
			for _, elem := range listening.Sync {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sync = append(%s.Sync, %s)", listeningIdent, listeningIdent, targetIdent))
			}
			for _, elem := range listening.Other_listening {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_listening = append(%s.Other_listening, %s)", listeningIdent, listeningIdent, targetIdent))
			}
			if listening.Offset != nil {
				targetIdent := "__stage_0" + listening.Offset.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", listeningIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		lyricOrdered := []*Lyric{}
		for lyric := range stageSet.Stage.Lyrics {
			lyricOrdered = append(lyricOrdered, lyric)
		}
		sort.Slice(lyricOrdered, func(i, j int) bool {
			return stageSet.Stage.Lyric_stagedOrder[lyricOrdered[i]] < stageSet.Stage.Lyric_stagedOrder[lyricOrdered[j]]
		})
		for _, lyric := range lyricOrdered {
			lyricIdent := "__stage_0" + lyric.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Lyric{Name: %s}).Stage(stageSet.Stage)", lyricIdent, __gong__toRawStringLiteral(lyric.Name)))
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
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elision = append(%s.Elision, %s)", lyricIdent, lyricIdent, targetIdent))
			}
			for _, elem := range lyric.Text {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Text = append(%s.Text, %s)", lyricIdent, lyricIdent, targetIdent))
			}
			if lyric.Extend != nil {
				targetIdent := "__stage_0" + lyric.Extend.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extend = %s", lyricIdent, targetIdent))
			}
			if lyric.Footnote != nil {
				targetIdent := "__stage_0" + lyric.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", lyricIdent, targetIdent))
			}
			if lyric.Level != nil {
				targetIdent := "__stage_0" + lyric.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", lyricIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		lyric_fontOrdered := []*Lyric_font{}
		for lyric_font := range stageSet.Stage.Lyric_fonts {
			lyric_fontOrdered = append(lyric_fontOrdered, lyric_font)
		}
		sort.Slice(lyric_fontOrdered, func(i, j int) bool {
			return stageSet.Stage.Lyric_font_stagedOrder[lyric_fontOrdered[i]] < stageSet.Stage.Lyric_font_stagedOrder[lyric_fontOrdered[j]]
		})
		for _, lyric_font := range lyric_fontOrdered {
			lyric_fontIdent := "__stage_0" + lyric_font.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Lyric_font{Name: %s}).Stage(stageSet.Stage)", lyric_fontIdent, __gong__toRawStringLiteral(lyric_font.Name)))
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
		lyric_languageOrdered := []*Lyric_language{}
		for lyric_language := range stageSet.Stage.Lyric_languages {
			lyric_languageOrdered = append(lyric_languageOrdered, lyric_language)
		}
		sort.Slice(lyric_languageOrdered, func(i, j int) bool {
			return stageSet.Stage.Lyric_language_stagedOrder[lyric_languageOrdered[i]] < stageSet.Stage.Lyric_language_stagedOrder[lyric_languageOrdered[j]]
		})
		for _, lyric_language := range lyric_languageOrdered {
			lyric_languageIdent := "__stage_0" + lyric_language.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Lyric_language{Name: %s}).Stage(stageSet.Stage)", lyric_languageIdent, __gong__toRawStringLiteral(lyric_language.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lyric_languageIdent, __gong__toRawStringLiteral(lyric_language.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", lyric_languageIdent, __gong__toRawStringLiteral(lyric_language.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", lyric_languageIdent, __gong__toRawStringLiteral(lyric_language.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", lyric_languageIdent, __gong__toRawStringLiteral(lyric_language.Lang)))
		}
	}
	if stageSet.Stage != nil {
		measure_layoutOrdered := []*Measure_layout{}
		for measure_layout := range stageSet.Stage.Measure_layouts {
			measure_layoutOrdered = append(measure_layoutOrdered, measure_layout)
		}
		sort.Slice(measure_layoutOrdered, func(i, j int) bool {
			return stageSet.Stage.Measure_layout_stagedOrder[measure_layoutOrdered[i]] < stageSet.Stage.Measure_layout_stagedOrder[measure_layoutOrdered[j]]
		})
		for _, measure_layout := range measure_layoutOrdered {
			measure_layoutIdent := "__stage_0" + measure_layout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Measure_layout{Name: %s}).Stage(stageSet.Stage)", measure_layoutIdent, __gong__toRawStringLiteral(measure_layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measure_layoutIdent, __gong__toRawStringLiteral(measure_layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Measure_distance = %s", measure_layoutIdent, __gong__toRawStringLiteral(measure_layout.Measure_distance)))
		}
	}
	if stageSet.Stage != nil {
		measure_numberingOrdered := []*Measure_numbering{}
		for measure_numbering := range stageSet.Stage.Measure_numberings {
			measure_numberingOrdered = append(measure_numberingOrdered, measure_numbering)
		}
		sort.Slice(measure_numberingOrdered, func(i, j int) bool {
			return stageSet.Stage.Measure_numbering_stagedOrder[measure_numberingOrdered[i]] < stageSet.Stage.Measure_numbering_stagedOrder[measure_numberingOrdered[j]]
		})
		for _, measure_numbering := range measure_numberingOrdered {
			measure_numberingIdent := "__stage_0" + measure_numbering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Measure_numbering{Name: %s}).Stage(stageSet.Stage)", measure_numberingIdent, __gong__toRawStringLiteral(measure_numbering.Name)))
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
		measure_repeatOrdered := []*Measure_repeat{}
		for measure_repeat := range stageSet.Stage.Measure_repeats {
			measure_repeatOrdered = append(measure_repeatOrdered, measure_repeat)
		}
		sort.Slice(measure_repeatOrdered, func(i, j int) bool {
			return stageSet.Stage.Measure_repeat_stagedOrder[measure_repeatOrdered[i]] < stageSet.Stage.Measure_repeat_stagedOrder[measure_repeatOrdered[j]]
		})
		for _, measure_repeat := range measure_repeatOrdered {
			measure_repeatIdent := "__stage_0" + measure_repeat.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Measure_repeat{Name: %s}).Stage(stageSet.Stage)", measure_repeatIdent, __gong__toRawStringLiteral(measure_repeat.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measure_repeatIdent, __gong__toRawStringLiteral(measure_repeat.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", measure_repeatIdent, __gong__toRawStringLiteral(string(measure_repeat.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Slashes = %d", measure_repeatIdent, measure_repeat.Slashes))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", measure_repeatIdent, __gong__toRawStringLiteral(measure_repeat.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		measure_styleOrdered := []*Measure_style{}
		for measure_style := range stageSet.Stage.Measure_styles {
			measure_styleOrdered = append(measure_styleOrdered, measure_style)
		}
		sort.Slice(measure_styleOrdered, func(i, j int) bool {
			return stageSet.Stage.Measure_style_stagedOrder[measure_styleOrdered[i]] < stageSet.Stage.Measure_style_stagedOrder[measure_styleOrdered[j]]
		})
		for _, measure_style := range measure_styleOrdered {
			measure_styleIdent := "__stage_0" + measure_style.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Measure_style{Name: %s}).Stage(stageSet.Stage)", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", measure_styleIdent, measure_style.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", measure_styleIdent, __gong__toRawStringLiteral(measure_style.Id)))
			if measure_style.Multiple_rest != nil {
				targetIdent := "__stage_0" + measure_style.Multiple_rest.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Multiple_rest = %s", measure_styleIdent, targetIdent))
			}
			if measure_style.Measure_repeat != nil {
				targetIdent := "__stage_0" + measure_style.Measure_repeat.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure_repeat = %s", measure_styleIdent, targetIdent))
			}
			if measure_style.Beat_repeat != nil {
				targetIdent := "__stage_0" + measure_style.Beat_repeat.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Beat_repeat = %s", measure_styleIdent, targetIdent))
			}
			if measure_style.Slash != nil {
				targetIdent := "__stage_0" + measure_style.Slash.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slash = %s", measure_styleIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		membraneOrdered := []*Membrane{}
		for membrane := range stageSet.Stage.Membranes {
			membraneOrdered = append(membraneOrdered, membrane)
		}
		sort.Slice(membraneOrdered, func(i, j int) bool {
			return stageSet.Stage.Membrane_stagedOrder[membraneOrdered[i]] < stageSet.Stage.Membrane_stagedOrder[membraneOrdered[j]]
		})
		for _, membrane := range membraneOrdered {
			membraneIdent := "__stage_0" + membrane.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Membrane{Name: %s}).Stage(stageSet.Stage)", membraneIdent, __gong__toRawStringLiteral(membrane.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", membraneIdent, __gong__toRawStringLiteral(membrane.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", membraneIdent, __gong__toRawStringLiteral(membrane.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", membraneIdent, __gong__toRawStringLiteral(membrane.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		metalOrdered := []*Metal{}
		for metal := range stageSet.Stage.Metals {
			metalOrdered = append(metalOrdered, metal)
		}
		sort.Slice(metalOrdered, func(i, j int) bool {
			return stageSet.Stage.Metal_stagedOrder[metalOrdered[i]] < stageSet.Stage.Metal_stagedOrder[metalOrdered[j]]
		})
		for _, metal := range metalOrdered {
			metalIdent := "__stage_0" + metal.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Metal{Name: %s}).Stage(stageSet.Stage)", metalIdent, __gong__toRawStringLiteral(metal.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metalIdent, __gong__toRawStringLiteral(metal.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", metalIdent, __gong__toRawStringLiteral(metal.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", metalIdent, __gong__toRawStringLiteral(metal.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		metronomeOrdered := []*Metronome{}
		for metronome := range stageSet.Stage.Metronomes {
			metronomeOrdered = append(metronomeOrdered, metronome)
		}
		sort.Slice(metronomeOrdered, func(i, j int) bool {
			return stageSet.Stage.Metronome_stagedOrder[metronomeOrdered[i]] < stageSet.Stage.Metronome_stagedOrder[metronomeOrdered[j]]
		})
		for _, metronome := range metronomeOrdered {
			metronomeIdent := "__stage_0" + metronome.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Metronome{Name: %s}).Stage(stageSet.Stage)", metronomeIdent, __gong__toRawStringLiteral(metronome.Name)))
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
				targetIdent := "__stage_0" + metronome.Per_minute.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Per_minute = %s", metronomeIdent, targetIdent))
			}
			for _, elem := range metronome.Beat_unit_tied {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Beat_unit_tied = append(%s.Beat_unit_tied, %s)", metronomeIdent, metronomeIdent, targetIdent))
			}
			for _, elem := range metronome.Metronome_note {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metronome_note = append(%s.Metronome_note, %s)", metronomeIdent, metronomeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		metronome_beamOrdered := []*Metronome_beam{}
		for metronome_beam := range stageSet.Stage.Metronome_beams {
			metronome_beamOrdered = append(metronome_beamOrdered, metronome_beam)
		}
		sort.Slice(metronome_beamOrdered, func(i, j int) bool {
			return stageSet.Stage.Metronome_beam_stagedOrder[metronome_beamOrdered[i]] < stageSet.Stage.Metronome_beam_stagedOrder[metronome_beamOrdered[j]]
		})
		for _, metronome_beam := range metronome_beamOrdered {
			metronome_beamIdent := "__stage_0" + metronome_beam.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Metronome_beam{Name: %s}).Stage(stageSet.Stage)", metronome_beamIdent, __gong__toRawStringLiteral(metronome_beam.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronome_beamIdent, __gong__toRawStringLiteral(metronome_beam.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", metronome_beamIdent, metronome_beam.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", metronome_beamIdent, __gong__toRawStringLiteral(string(metronome_beam.EnclosedText))))
		}
	}
	if stageSet.Stage != nil {
		metronome_noteOrdered := []*Metronome_note{}
		for metronome_note := range stageSet.Stage.Metronome_notes {
			metronome_noteOrdered = append(metronome_noteOrdered, metronome_note)
		}
		sort.Slice(metronome_noteOrdered, func(i, j int) bool {
			return stageSet.Stage.Metronome_note_stagedOrder[metronome_noteOrdered[i]] < stageSet.Stage.Metronome_note_stagedOrder[metronome_noteOrdered[j]]
		})
		for _, metronome_note := range metronome_noteOrdered {
			metronome_noteIdent := "__stage_0" + metronome_note.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Metronome_note{Name: %s}).Stage(stageSet.Stage)", metronome_noteIdent, __gong__toRawStringLiteral(metronome_note.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronome_noteIdent, __gong__toRawStringLiteral(metronome_note.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Metronome_type = %s", metronome_noteIdent, __gong__toRawStringLiteral(metronome_note.Metronome_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Metronome_dot = %s", metronome_noteIdent, __gong__toRawStringLiteral(metronome_note.Metronome_dot)))
			for _, elem := range metronome_note.Metronome_beam {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metronome_beam = append(%s.Metronome_beam, %s)", metronome_noteIdent, metronome_noteIdent, targetIdent))
			}
			if metronome_note.Metronome_tied != nil {
				targetIdent := "__stage_0" + metronome_note.Metronome_tied.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metronome_tied = %s", metronome_noteIdent, targetIdent))
			}
			if metronome_note.Metronome_tuplet != nil {
				targetIdent := "__stage_0" + metronome_note.Metronome_tuplet.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metronome_tuplet = %s", metronome_noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		metronome_tiedOrdered := []*Metronome_tied{}
		for metronome_tied := range stageSet.Stage.Metronome_tieds {
			metronome_tiedOrdered = append(metronome_tiedOrdered, metronome_tied)
		}
		sort.Slice(metronome_tiedOrdered, func(i, j int) bool {
			return stageSet.Stage.Metronome_tied_stagedOrder[metronome_tiedOrdered[i]] < stageSet.Stage.Metronome_tied_stagedOrder[metronome_tiedOrdered[j]]
		})
		for _, metronome_tied := range metronome_tiedOrdered {
			metronome_tiedIdent := "__stage_0" + metronome_tied.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Metronome_tied{Name: %s}).Stage(stageSet.Stage)", metronome_tiedIdent, __gong__toRawStringLiteral(metronome_tied.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronome_tiedIdent, __gong__toRawStringLiteral(metronome_tied.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", metronome_tiedIdent, __gong__toRawStringLiteral(string(metronome_tied.Type))))
		}
	}
	if stageSet.Stage != nil {
		metronome_tupletOrdered := []*Metronome_tuplet{}
		for metronome_tuplet := range stageSet.Stage.Metronome_tuplets {
			metronome_tupletOrdered = append(metronome_tupletOrdered, metronome_tuplet)
		}
		sort.Slice(metronome_tupletOrdered, func(i, j int) bool {
			return stageSet.Stage.Metronome_tuplet_stagedOrder[metronome_tupletOrdered[i]] < stageSet.Stage.Metronome_tuplet_stagedOrder[metronome_tupletOrdered[j]]
		})
		for _, metronome_tuplet := range metronome_tupletOrdered {
			metronome_tupletIdent := "__stage_0" + metronome_tuplet.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Metronome_tuplet{Name: %s}).Stage(stageSet.Stage)", metronome_tupletIdent, __gong__toRawStringLiteral(metronome_tuplet.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronome_tupletIdent, __gong__toRawStringLiteral(metronome_tuplet.Name)))
		}
	}
	if stageSet.Stage != nil {
		midi_deviceOrdered := []*Midi_device{}
		for midi_device := range stageSet.Stage.Midi_devices {
			midi_deviceOrdered = append(midi_deviceOrdered, midi_device)
		}
		sort.Slice(midi_deviceOrdered, func(i, j int) bool {
			return stageSet.Stage.Midi_device_stagedOrder[midi_deviceOrdered[i]] < stageSet.Stage.Midi_device_stagedOrder[midi_deviceOrdered[j]]
		})
		for _, midi_device := range midi_deviceOrdered {
			midi_deviceIdent := "__stage_0" + midi_device.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Midi_device{Name: %s}).Stage(stageSet.Stage)", midi_deviceIdent, __gong__toRawStringLiteral(midi_device.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", midi_deviceIdent, __gong__toRawStringLiteral(midi_device.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Port = %d", midi_deviceIdent, midi_device.Port))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", midi_deviceIdent, __gong__toRawStringLiteral(midi_device.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", midi_deviceIdent, __gong__toRawStringLiteral(midi_device.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		midi_instrumentOrdered := []*Midi_instrument{}
		for midi_instrument := range stageSet.Stage.Midi_instruments {
			midi_instrumentOrdered = append(midi_instrumentOrdered, midi_instrument)
		}
		sort.Slice(midi_instrumentOrdered, func(i, j int) bool {
			return stageSet.Stage.Midi_instrument_stagedOrder[midi_instrumentOrdered[i]] < stageSet.Stage.Midi_instrument_stagedOrder[midi_instrumentOrdered[j]]
		})
		for _, midi_instrument := range midi_instrumentOrdered {
			midi_instrumentIdent := "__stage_0" + midi_instrument.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Midi_instrument{Name: %s}).Stage(stageSet.Stage)", midi_instrumentIdent, __gong__toRawStringLiteral(midi_instrument.Name)))
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
		miscellaneousOrdered := []*Miscellaneous{}
		for miscellaneous := range stageSet.Stage.Miscellaneouss {
			miscellaneousOrdered = append(miscellaneousOrdered, miscellaneous)
		}
		sort.Slice(miscellaneousOrdered, func(i, j int) bool {
			return stageSet.Stage.Miscellaneous_stagedOrder[miscellaneousOrdered[i]] < stageSet.Stage.Miscellaneous_stagedOrder[miscellaneousOrdered[j]]
		})
		for _, miscellaneous := range miscellaneousOrdered {
			miscellaneousIdent := "__stage_0" + miscellaneous.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Miscellaneous{Name: %s}).Stage(stageSet.Stage)", miscellaneousIdent, __gong__toRawStringLiteral(miscellaneous.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", miscellaneousIdent, __gong__toRawStringLiteral(miscellaneous.Name)))
			for _, elem := range miscellaneous.Miscellaneous_field {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Miscellaneous_field = append(%s.Miscellaneous_field, %s)", miscellaneousIdent, miscellaneousIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		miscellaneous_fieldOrdered := []*Miscellaneous_field{}
		for miscellaneous_field := range stageSet.Stage.Miscellaneous_fields {
			miscellaneous_fieldOrdered = append(miscellaneous_fieldOrdered, miscellaneous_field)
		}
		sort.Slice(miscellaneous_fieldOrdered, func(i, j int) bool {
			return stageSet.Stage.Miscellaneous_field_stagedOrder[miscellaneous_fieldOrdered[i]] < stageSet.Stage.Miscellaneous_field_stagedOrder[miscellaneous_fieldOrdered[j]]
		})
		for _, miscellaneous_field := range miscellaneous_fieldOrdered {
			miscellaneous_fieldIdent := "__stage_0" + miscellaneous_field.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Miscellaneous_field{Name: %s}).Stage(stageSet.Stage)", miscellaneous_fieldIdent, __gong__toRawStringLiteral(miscellaneous_field.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", miscellaneous_fieldIdent, __gong__toRawStringLiteral(miscellaneous_field.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameXSD = %s", miscellaneous_fieldIdent, __gong__toRawStringLiteral(miscellaneous_field.NameXSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", miscellaneous_fieldIdent, __gong__toRawStringLiteral(miscellaneous_field.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		mordentOrdered := []*Mordent{}
		for mordent := range stageSet.Stage.Mordents {
			mordentOrdered = append(mordentOrdered, mordent)
		}
		sort.Slice(mordentOrdered, func(i, j int) bool {
			return stageSet.Stage.Mordent_stagedOrder[mordentOrdered[i]] < stageSet.Stage.Mordent_stagedOrder[mordentOrdered[j]]
		})
		for _, mordent := range mordentOrdered {
			mordentIdent := "__stage_0" + mordent.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Mordent{Name: %s}).Stage(stageSet.Stage)", mordentIdent, __gong__toRawStringLiteral(mordent.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", mordentIdent, __gong__toRawStringLiteral(mordent.Name)))
		}
	}
	if stageSet.Stage != nil {
		multiple_restOrdered := []*Multiple_rest{}
		for multiple_rest := range stageSet.Stage.Multiple_rests {
			multiple_restOrdered = append(multiple_restOrdered, multiple_rest)
		}
		sort.Slice(multiple_restOrdered, func(i, j int) bool {
			return stageSet.Stage.Multiple_rest_stagedOrder[multiple_restOrdered[i]] < stageSet.Stage.Multiple_rest_stagedOrder[multiple_restOrdered[j]]
		})
		for _, multiple_rest := range multiple_restOrdered {
			multiple_restIdent := "__stage_0" + multiple_rest.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Multiple_rest{Name: %s}).Stage(stageSet.Stage)", multiple_restIdent, __gong__toRawStringLiteral(multiple_rest.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", multiple_restIdent, __gong__toRawStringLiteral(multiple_rest.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Use_symbols = %s", multiple_restIdent, __gong__toRawStringLiteral(string(multiple_rest.Use_symbols))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %d", multiple_restIdent, multiple_rest.EnclosedText))
		}
	}
	if stageSet.Stage != nil {
		name_displayOrdered := []*Name_display{}
		for name_display := range stageSet.Stage.Name_displays {
			name_displayOrdered = append(name_displayOrdered, name_display)
		}
		sort.Slice(name_displayOrdered, func(i, j int) bool {
			return stageSet.Stage.Name_display_stagedOrder[name_displayOrdered[i]] < stageSet.Stage.Name_display_stagedOrder[name_displayOrdered[j]]
		})
		for _, name_display := range name_displayOrdered {
			name_displayIdent := "__stage_0" + name_display.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Name_display{Name: %s}).Stage(stageSet.Stage)", name_displayIdent, __gong__toRawStringLiteral(name_display.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", name_displayIdent, __gong__toRawStringLiteral(name_display.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", name_displayIdent, __gong__toRawStringLiteral(string(name_display.Print_object))))
			for _, elem := range name_display.Display_text {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Display_text = append(%s.Display_text, %s)", name_displayIdent, name_displayIdent, targetIdent))
			}
			for _, elem := range name_display.Accidental_text {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accidental_text = append(%s.Accidental_text, %s)", name_displayIdent, name_displayIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		non_arpeggiateOrdered := []*Non_arpeggiate{}
		for non_arpeggiate := range stageSet.Stage.Non_arpeggiates {
			non_arpeggiateOrdered = append(non_arpeggiateOrdered, non_arpeggiate)
		}
		sort.Slice(non_arpeggiateOrdered, func(i, j int) bool {
			return stageSet.Stage.Non_arpeggiate_stagedOrder[non_arpeggiateOrdered[i]] < stageSet.Stage.Non_arpeggiate_stagedOrder[non_arpeggiateOrdered[j]]
		})
		for _, non_arpeggiate := range non_arpeggiateOrdered {
			non_arpeggiateIdent := "__stage_0" + non_arpeggiate.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Non_arpeggiate{Name: %s}).Stage(stageSet.Stage)", non_arpeggiateIdent, __gong__toRawStringLiteral(non_arpeggiate.Name)))
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
		notationsOrdered := []*Notations{}
		for notations := range stageSet.Stage.Notationss {
			notationsOrdered = append(notationsOrdered, notations)
		}
		sort.Slice(notationsOrdered, func(i, j int) bool {
			return stageSet.Stage.Notations_stagedOrder[notationsOrdered[i]] < stageSet.Stage.Notations_stagedOrder[notationsOrdered[j]]
		})
		for _, notations := range notationsOrdered {
			notationsIdent := "__stage_0" + notations.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Notations{Name: %s}).Stage(stageSet.Stage)", notationsIdent, __gong__toRawStringLiteral(notations.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notationsIdent, __gong__toRawStringLiteral(notations.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", notationsIdent, __gong__toRawStringLiteral(string(notations.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", notationsIdent, __gong__toRawStringLiteral(notations.Id)))
			if notations.Footnote != nil {
				targetIdent := "__stage_0" + notations.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", notationsIdent, targetIdent))
			}
			if notations.Level != nil {
				targetIdent := "__stage_0" + notations.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", notationsIdent, targetIdent))
			}
			for _, elem := range notations.Tied {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tied = append(%s.Tied, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Slur {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slur = append(%s.Slur, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Tuplet {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet = append(%s.Tuplet, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Glissando {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Glissando = append(%s.Glissando, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Slide {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slide = append(%s.Slide, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Ornaments {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ornaments = append(%s.Ornaments, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Technical {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Technical = append(%s.Technical, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Articulations {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Articulations = append(%s.Articulations, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Dynamics {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dynamics = append(%s.Dynamics, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Fermata {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fermata = append(%s.Fermata, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Arpeggiate {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arpeggiate = append(%s.Arpeggiate, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Non_arpeggiate {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Non_arpeggiate = append(%s.Non_arpeggiate, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Accidental_mark {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accidental_mark = append(%s.Accidental_mark, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Other_notation {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_notation = append(%s.Other_notation, %s)", notationsIdent, notationsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		noteOrdered := []*Note{}
		for note := range stageSet.Stage.Notes {
			noteOrdered = append(noteOrdered, note)
		}
		sort.Slice(noteOrdered, func(i, j int) bool {
			return stageSet.Stage.Note_stagedOrder[noteOrdered[i]] < stageSet.Stage.Note_stagedOrder[noteOrdered[j]]
		})
		for _, note := range noteOrdered {
			noteIdent := "__stage_0" + note.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Note{Name: %s}).Stage(stageSet.Stage)", noteIdent, __gong__toRawStringLiteral(note.Name)))
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
				targetIdent := "__stage_0" + note.Grace.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Grace = %s", noteIdent, targetIdent))
			}
			if note.Pitch != nil {
				targetIdent := "__stage_0" + note.Pitch.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pitch = %s", noteIdent, targetIdent))
			}
			if note.Unpitched != nil {
				targetIdent := "__stage_0" + note.Unpitched.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Unpitched = %s", noteIdent, targetIdent))
			}
			if note.Rest != nil {
				targetIdent := "__stage_0" + note.Rest.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rest = %s", noteIdent, targetIdent))
			}
			if note.Tie != nil {
				targetIdent := "__stage_0" + note.Tie.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tie = %s", noteIdent, targetIdent))
			}
			for _, elem := range note.Instrument {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Instrument = append(%s.Instrument, %s)", noteIdent, noteIdent, targetIdent))
			}
			if note.Footnote != nil {
				targetIdent := "__stage_0" + note.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", noteIdent, targetIdent))
			}
			if note.Level != nil {
				targetIdent := "__stage_0" + note.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", noteIdent, targetIdent))
			}
			if note.Type != nil {
				targetIdent := "__stage_0" + note.Type.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type = %s", noteIdent, targetIdent))
			}
			for _, elem := range note.Dot {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dot = append(%s.Dot, %s)", noteIdent, noteIdent, targetIdent))
			}
			if note.Accidental != nil {
				targetIdent := "__stage_0" + note.Accidental.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accidental = %s", noteIdent, targetIdent))
			}
			if note.Time_modification != nil {
				targetIdent := "__stage_0" + note.Time_modification.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Time_modification = %s", noteIdent, targetIdent))
			}
			if note.Stem != nil {
				targetIdent := "__stage_0" + note.Stem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stem = %s", noteIdent, targetIdent))
			}
			if note.Notehead != nil {
				targetIdent := "__stage_0" + note.Notehead.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notehead = %s", noteIdent, targetIdent))
			}
			if note.Notehead_text != nil {
				targetIdent := "__stage_0" + note.Notehead_text.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notehead_text = %s", noteIdent, targetIdent))
			}
			if note.Beam != nil {
				targetIdent := "__stage_0" + note.Beam.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Beam = %s", noteIdent, targetIdent))
			}
			for _, elem := range note.Notations {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notations = append(%s.Notations, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Lyric {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lyric = append(%s.Lyric, %s)", noteIdent, noteIdent, targetIdent))
			}
			if note.Play != nil {
				targetIdent := "__stage_0" + note.Play.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Play = %s", noteIdent, targetIdent))
			}
			if note.Listen != nil {
				targetIdent := "__stage_0" + note.Listen.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listen = %s", noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		note_sizeOrdered := []*Note_size{}
		for note_size := range stageSet.Stage.Note_sizes {
			note_sizeOrdered = append(note_sizeOrdered, note_size)
		}
		sort.Slice(note_sizeOrdered, func(i, j int) bool {
			return stageSet.Stage.Note_size_stagedOrder[note_sizeOrdered[i]] < stageSet.Stage.Note_size_stagedOrder[note_sizeOrdered[j]]
		})
		for _, note_size := range note_sizeOrdered {
			note_sizeIdent := "__stage_0" + note_size.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Note_size{Name: %s}).Stage(stageSet.Stage)", note_sizeIdent, __gong__toRawStringLiteral(note_size.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", note_sizeIdent, __gong__toRawStringLiteral(note_size.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", note_sizeIdent, __gong__toRawStringLiteral(note_size.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", note_sizeIdent, __gong__toRawStringLiteral(note_size.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		note_typeOrdered := []*Note_type{}
		for note_type := range stageSet.Stage.Note_types {
			note_typeOrdered = append(note_typeOrdered, note_type)
		}
		sort.Slice(note_typeOrdered, func(i, j int) bool {
			return stageSet.Stage.Note_type_stagedOrder[note_typeOrdered[i]] < stageSet.Stage.Note_type_stagedOrder[note_typeOrdered[j]]
		})
		for _, note_type := range note_typeOrdered {
			note_typeIdent := "__stage_0" + note_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Note_type{Name: %s}).Stage(stageSet.Stage)", note_typeIdent, __gong__toRawStringLiteral(note_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", note_typeIdent, __gong__toRawStringLiteral(note_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %s", note_typeIdent, __gong__toRawStringLiteral(string(note_type.Size))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", note_typeIdent, __gong__toRawStringLiteral(string(note_type.EnclosedText))))
		}
	}
	if stageSet.Stage != nil {
		noteheadOrdered := []*Notehead{}
		for notehead := range stageSet.Stage.Noteheads {
			noteheadOrdered = append(noteheadOrdered, notehead)
		}
		sort.Slice(noteheadOrdered, func(i, j int) bool {
			return stageSet.Stage.Notehead_stagedOrder[noteheadOrdered[i]] < stageSet.Stage.Notehead_stagedOrder[noteheadOrdered[j]]
		})
		for _, notehead := range noteheadOrdered {
			noteheadIdent := "__stage_0" + notehead.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Notehead{Name: %s}).Stage(stageSet.Stage)", noteheadIdent, __gong__toRawStringLiteral(notehead.Name)))
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
		notehead_textOrdered := []*Notehead_text{}
		for notehead_text := range stageSet.Stage.Notehead_texts {
			notehead_textOrdered = append(notehead_textOrdered, notehead_text)
		}
		sort.Slice(notehead_textOrdered, func(i, j int) bool {
			return stageSet.Stage.Notehead_text_stagedOrder[notehead_textOrdered[i]] < stageSet.Stage.Notehead_text_stagedOrder[notehead_textOrdered[j]]
		})
		for _, notehead_text := range notehead_textOrdered {
			notehead_textIdent := "__stage_0" + notehead_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Notehead_text{Name: %s}).Stage(stageSet.Stage)", notehead_textIdent, __gong__toRawStringLiteral(notehead_text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notehead_textIdent, __gong__toRawStringLiteral(notehead_text.Name)))
			for _, elem := range notehead_text.Display_text {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Display_text = append(%s.Display_text, %s)", notehead_textIdent, notehead_textIdent, targetIdent))
			}
			for _, elem := range notehead_text.Accidental_text {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accidental_text = append(%s.Accidental_text, %s)", notehead_textIdent, notehead_textIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		numeralOrdered := []*Numeral{}
		for numeral := range stageSet.Stage.Numerals {
			numeralOrdered = append(numeralOrdered, numeral)
		}
		sort.Slice(numeralOrdered, func(i, j int) bool {
			return stageSet.Stage.Numeral_stagedOrder[numeralOrdered[i]] < stageSet.Stage.Numeral_stagedOrder[numeralOrdered[j]]
		})
		for _, numeral := range numeralOrdered {
			numeralIdent := "__stage_0" + numeral.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Numeral{Name: %s}).Stage(stageSet.Stage)", numeralIdent, __gong__toRawStringLiteral(numeral.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", numeralIdent, __gong__toRawStringLiteral(numeral.Name)))
			if numeral.Numeral_root != nil {
				targetIdent := "__stage_0" + numeral.Numeral_root.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Numeral_root = %s", numeralIdent, targetIdent))
			}
			if numeral.Numeral_alter != nil {
				targetIdent := "__stage_0" + numeral.Numeral_alter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Numeral_alter = %s", numeralIdent, targetIdent))
			}
			if numeral.Numeral_key != nil {
				targetIdent := "__stage_0" + numeral.Numeral_key.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Numeral_key = %s", numeralIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		numeral_keyOrdered := []*Numeral_key{}
		for numeral_key := range stageSet.Stage.Numeral_keys {
			numeral_keyOrdered = append(numeral_keyOrdered, numeral_key)
		}
		sort.Slice(numeral_keyOrdered, func(i, j int) bool {
			return stageSet.Stage.Numeral_key_stagedOrder[numeral_keyOrdered[i]] < stageSet.Stage.Numeral_key_stagedOrder[numeral_keyOrdered[j]]
		})
		for _, numeral_key := range numeral_keyOrdered {
			numeral_keyIdent := "__stage_0" + numeral_key.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Numeral_key{Name: %s}).Stage(stageSet.Stage)", numeral_keyIdent, __gong__toRawStringLiteral(numeral_key.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", numeral_keyIdent, __gong__toRawStringLiteral(numeral_key.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", numeral_keyIdent, __gong__toRawStringLiteral(string(numeral_key.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Numeral_fifths = %d", numeral_keyIdent, numeral_key.Numeral_fifths))
			values.WriteString(fmt.Sprintf("\n\t%s.Numeral_mode = %s", numeral_keyIdent, __gong__toRawStringLiteral(numeral_key.Numeral_mode)))
		}
	}
	if stageSet.Stage != nil {
		numeral_rootOrdered := []*Numeral_root{}
		for numeral_root := range stageSet.Stage.Numeral_roots {
			numeral_rootOrdered = append(numeral_rootOrdered, numeral_root)
		}
		sort.Slice(numeral_rootOrdered, func(i, j int) bool {
			return stageSet.Stage.Numeral_root_stagedOrder[numeral_rootOrdered[i]] < stageSet.Stage.Numeral_root_stagedOrder[numeral_rootOrdered[j]]
		})
		for _, numeral_root := range numeral_rootOrdered {
			numeral_rootIdent := "__stage_0" + numeral_root.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Numeral_root{Name: %s}).Stage(stageSet.Stage)", numeral_rootIdent, __gong__toRawStringLiteral(numeral_root.Name)))
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
		octave_shiftOrdered := []*Octave_shift{}
		for octave_shift := range stageSet.Stage.Octave_shifts {
			octave_shiftOrdered = append(octave_shiftOrdered, octave_shift)
		}
		sort.Slice(octave_shiftOrdered, func(i, j int) bool {
			return stageSet.Stage.Octave_shift_stagedOrder[octave_shiftOrdered[i]] < stageSet.Stage.Octave_shift_stagedOrder[octave_shiftOrdered[j]]
		})
		for _, octave_shift := range octave_shiftOrdered {
			octave_shiftIdent := "__stage_0" + octave_shift.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Octave_shift{Name: %s}).Stage(stageSet.Stage)", octave_shiftIdent, __gong__toRawStringLiteral(octave_shift.Name)))
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
		offsetOrdered := []*Offset{}
		for offset := range stageSet.Stage.Offsets {
			offsetOrdered = append(offsetOrdered, offset)
		}
		sort.Slice(offsetOrdered, func(i, j int) bool {
			return stageSet.Stage.Offset_stagedOrder[offsetOrdered[i]] < stageSet.Stage.Offset_stagedOrder[offsetOrdered[j]]
		})
		for _, offset := range offsetOrdered {
			offsetIdent := "__stage_0" + offset.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Offset{Name: %s}).Stage(stageSet.Stage)", offsetIdent, __gong__toRawStringLiteral(offset.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", offsetIdent, __gong__toRawStringLiteral(offset.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sound = %s", offsetIdent, __gong__toRawStringLiteral(string(offset.Sound))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", offsetIdent, __gong__toRawStringLiteral(offset.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		opusOrdered := []*Opus{}
		for opus := range stageSet.Stage.Opuss {
			opusOrdered = append(opusOrdered, opus)
		}
		sort.Slice(opusOrdered, func(i, j int) bool {
			return stageSet.Stage.Opus_stagedOrder[opusOrdered[i]] < stageSet.Stage.Opus_stagedOrder[opusOrdered[j]]
		})
		for _, opus := range opusOrdered {
			opusIdent := "__stage_0" + opus.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Opus{Name: %s}).Stage(stageSet.Stage)", opusIdent, __gong__toRawStringLiteral(opus.Name)))
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
		ornamentsOrdered := []*Ornaments{}
		for ornaments := range stageSet.Stage.Ornamentss {
			ornamentsOrdered = append(ornamentsOrdered, ornaments)
		}
		sort.Slice(ornamentsOrdered, func(i, j int) bool {
			return stageSet.Stage.Ornaments_stagedOrder[ornamentsOrdered[i]] < stageSet.Stage.Ornaments_stagedOrder[ornamentsOrdered[j]]
		})
		for _, ornaments := range ornamentsOrdered {
			ornamentsIdent := "__stage_0" + ornaments.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Ornaments{Name: %s}).Stage(stageSet.Stage)", ornamentsIdent, __gong__toRawStringLiteral(ornaments.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ornamentsIdent, __gong__toRawStringLiteral(ornaments.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ornamentsIdent, __gong__toRawStringLiteral(ornaments.Id)))
			for _, elem := range ornaments.Trill_mark {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Trill_mark = append(%s.Trill_mark, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Turn {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Turn = append(%s.Turn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Delayed_turn {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Delayed_turn = append(%s.Delayed_turn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Inverted_turn {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inverted_turn = append(%s.Inverted_turn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Delayed_inverted_turn {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Delayed_inverted_turn = append(%s.Delayed_inverted_turn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Vertical_turn {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Vertical_turn = append(%s.Vertical_turn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Inverted_vertical_turn {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inverted_vertical_turn = append(%s.Inverted_vertical_turn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Shake {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Shake = append(%s.Shake, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Wavy_line {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wavy_line = append(%s.Wavy_line, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Mordent {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mordent = append(%s.Mordent, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Inverted_mordent {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inverted_mordent = append(%s.Inverted_mordent, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Schleifer {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Schleifer = append(%s.Schleifer, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Tremolo {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tremolo = append(%s.Tremolo, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Haydn {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Haydn = append(%s.Haydn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Other_ornament {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_ornament = append(%s.Other_ornament, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Accidental_mark {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accidental_mark = append(%s.Accidental_mark, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		other_appearanceOrdered := []*Other_appearance{}
		for other_appearance := range stageSet.Stage.Other_appearances {
			other_appearanceOrdered = append(other_appearanceOrdered, other_appearance)
		}
		sort.Slice(other_appearanceOrdered, func(i, j int) bool {
			return stageSet.Stage.Other_appearance_stagedOrder[other_appearanceOrdered[i]] < stageSet.Stage.Other_appearance_stagedOrder[other_appearanceOrdered[j]]
		})
		for _, other_appearance := range other_appearanceOrdered {
			other_appearanceIdent := "__stage_0" + other_appearance.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Other_appearance{Name: %s}).Stage(stageSet.Stage)", other_appearanceIdent, __gong__toRawStringLiteral(other_appearance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", other_appearanceIdent, __gong__toRawStringLiteral(other_appearance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", other_appearanceIdent, __gong__toRawStringLiteral(other_appearance.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", other_appearanceIdent, __gong__toRawStringLiteral(other_appearance.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		other_directionOrdered := []*Other_direction{}
		for other_direction := range stageSet.Stage.Other_directions {
			other_directionOrdered = append(other_directionOrdered, other_direction)
		}
		sort.Slice(other_directionOrdered, func(i, j int) bool {
			return stageSet.Stage.Other_direction_stagedOrder[other_directionOrdered[i]] < stageSet.Stage.Other_direction_stagedOrder[other_directionOrdered[j]]
		})
		for _, other_direction := range other_directionOrdered {
			other_directionIdent := "__stage_0" + other_direction.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Other_direction{Name: %s}).Stage(stageSet.Stage)", other_directionIdent, __gong__toRawStringLiteral(other_direction.Name)))
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
		other_listeningOrdered := []*Other_listening{}
		for other_listening := range stageSet.Stage.Other_listenings {
			other_listeningOrdered = append(other_listeningOrdered, other_listening)
		}
		sort.Slice(other_listeningOrdered, func(i, j int) bool {
			return stageSet.Stage.Other_listening_stagedOrder[other_listeningOrdered[i]] < stageSet.Stage.Other_listening_stagedOrder[other_listeningOrdered[j]]
		})
		for _, other_listening := range other_listeningOrdered {
			other_listeningIdent := "__stage_0" + other_listening.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Other_listening{Name: %s}).Stage(stageSet.Stage)", other_listeningIdent, __gong__toRawStringLiteral(other_listening.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", other_listeningIdent, __gong__toRawStringLiteral(other_listening.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", other_listeningIdent, __gong__toRawStringLiteral(other_listening.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", other_listeningIdent, __gong__toRawStringLiteral(other_listening.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_only = %s", other_listeningIdent, __gong__toRawStringLiteral(string(other_listening.Time_only))))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", other_listeningIdent, __gong__toRawStringLiteral(other_listening.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		other_notationOrdered := []*Other_notation{}
		for other_notation := range stageSet.Stage.Other_notations {
			other_notationOrdered = append(other_notationOrdered, other_notation)
		}
		sort.Slice(other_notationOrdered, func(i, j int) bool {
			return stageSet.Stage.Other_notation_stagedOrder[other_notationOrdered[i]] < stageSet.Stage.Other_notation_stagedOrder[other_notationOrdered[j]]
		})
		for _, other_notation := range other_notationOrdered {
			other_notationIdent := "__stage_0" + other_notation.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Other_notation{Name: %s}).Stage(stageSet.Stage)", other_notationIdent, __gong__toRawStringLiteral(other_notation.Name)))
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
		other_placement_textOrdered := []*Other_placement_text{}
		for other_placement_text := range stageSet.Stage.Other_placement_texts {
			other_placement_textOrdered = append(other_placement_textOrdered, other_placement_text)
		}
		sort.Slice(other_placement_textOrdered, func(i, j int) bool {
			return stageSet.Stage.Other_placement_text_stagedOrder[other_placement_textOrdered[i]] < stageSet.Stage.Other_placement_text_stagedOrder[other_placement_textOrdered[j]]
		})
		for _, other_placement_text := range other_placement_textOrdered {
			other_placement_textIdent := "__stage_0" + other_placement_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Other_placement_text{Name: %s}).Stage(stageSet.Stage)", other_placement_textIdent, __gong__toRawStringLiteral(other_placement_text.Name)))
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
		other_playOrdered := []*Other_play{}
		for other_play := range stageSet.Stage.Other_plays {
			other_playOrdered = append(other_playOrdered, other_play)
		}
		sort.Slice(other_playOrdered, func(i, j int) bool {
			return stageSet.Stage.Other_play_stagedOrder[other_playOrdered[i]] < stageSet.Stage.Other_play_stagedOrder[other_playOrdered[j]]
		})
		for _, other_play := range other_playOrdered {
			other_playIdent := "__stage_0" + other_play.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Other_play{Name: %s}).Stage(stageSet.Stage)", other_playIdent, __gong__toRawStringLiteral(other_play.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", other_playIdent, __gong__toRawStringLiteral(other_play.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", other_playIdent, __gong__toRawStringLiteral(other_play.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", other_playIdent, __gong__toRawStringLiteral(other_play.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		other_textOrdered := []*Other_text{}
		for other_text := range stageSet.Stage.Other_texts {
			other_textOrdered = append(other_textOrdered, other_text)
		}
		sort.Slice(other_textOrdered, func(i, j int) bool {
			return stageSet.Stage.Other_text_stagedOrder[other_textOrdered[i]] < stageSet.Stage.Other_text_stagedOrder[other_textOrdered[j]]
		})
		for _, other_text := range other_textOrdered {
			other_textIdent := "__stage_0" + other_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Other_text{Name: %s}).Stage(stageSet.Stage)", other_textIdent, __gong__toRawStringLiteral(other_text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", other_textIdent, __gong__toRawStringLiteral(other_text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", other_textIdent, __gong__toRawStringLiteral(other_text.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", other_textIdent, __gong__toRawStringLiteral(other_text.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		page_layoutOrdered := []*Page_layout{}
		for page_layout := range stageSet.Stage.Page_layouts {
			page_layoutOrdered = append(page_layoutOrdered, page_layout)
		}
		sort.Slice(page_layoutOrdered, func(i, j int) bool {
			return stageSet.Stage.Page_layout_stagedOrder[page_layoutOrdered[i]] < stageSet.Stage.Page_layout_stagedOrder[page_layoutOrdered[j]]
		})
		for _, page_layout := range page_layoutOrdered {
			page_layoutIdent := "__stage_0" + page_layout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Page_layout{Name: %s}).Stage(stageSet.Stage)", page_layoutIdent, __gong__toRawStringLiteral(page_layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", page_layoutIdent, __gong__toRawStringLiteral(page_layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Page_height = %s", page_layoutIdent, __gong__toRawStringLiteral(page_layout.Page_height)))
			values.WriteString(fmt.Sprintf("\n\t%s.Page_width = %s", page_layoutIdent, __gong__toRawStringLiteral(page_layout.Page_width)))
			if page_layout.Page_margins != nil {
				targetIdent := "__stage_0" + page_layout.Page_margins.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Page_margins = %s", page_layoutIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		page_marginsOrdered := []*Page_margins{}
		for page_margins := range stageSet.Stage.Page_marginss {
			page_marginsOrdered = append(page_marginsOrdered, page_margins)
		}
		sort.Slice(page_marginsOrdered, func(i, j int) bool {
			return stageSet.Stage.Page_margins_stagedOrder[page_marginsOrdered[i]] < stageSet.Stage.Page_margins_stagedOrder[page_marginsOrdered[j]]
		})
		for _, page_margins := range page_marginsOrdered {
			page_marginsIdent := "__stage_0" + page_margins.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Page_margins{Name: %s}).Stage(stageSet.Stage)", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Left_margin = %s", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Left_margin)))
			values.WriteString(fmt.Sprintf("\n\t%s.Right_margin = %s", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Right_margin)))
			values.WriteString(fmt.Sprintf("\n\t%s.Top_margin = %s", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Top_margin)))
			values.WriteString(fmt.Sprintf("\n\t%s.Bottom_margin = %s", page_marginsIdent, __gong__toRawStringLiteral(page_margins.Bottom_margin)))
		}
	}
	if stageSet.Stage != nil {
		part_clefOrdered := []*Part_clef{}
		for part_clef := range stageSet.Stage.Part_clefs {
			part_clefOrdered = append(part_clefOrdered, part_clef)
		}
		sort.Slice(part_clefOrdered, func(i, j int) bool {
			return stageSet.Stage.Part_clef_stagedOrder[part_clefOrdered[i]] < stageSet.Stage.Part_clef_stagedOrder[part_clefOrdered[j]]
		})
		for _, part_clef := range part_clefOrdered {
			part_clefIdent := "__stage_0" + part_clef.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Part_clef{Name: %s}).Stage(stageSet.Stage)", part_clefIdent, __gong__toRawStringLiteral(part_clef.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", part_clefIdent, __gong__toRawStringLiteral(part_clef.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sign = %s", part_clefIdent, __gong__toRawStringLiteral(part_clef.Sign)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", part_clefIdent, part_clef.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.Clef_octave_change = %d", part_clefIdent, part_clef.Clef_octave_change))
		}
	}
	if stageSet.Stage != nil {
		part_groupOrdered := []*Part_group{}
		for part_group := range stageSet.Stage.Part_groups {
			part_groupOrdered = append(part_groupOrdered, part_group)
		}
		sort.Slice(part_groupOrdered, func(i, j int) bool {
			return stageSet.Stage.Part_group_stagedOrder[part_groupOrdered[i]] < stageSet.Stage.Part_group_stagedOrder[part_groupOrdered[j]]
		})
		for _, part_group := range part_groupOrdered {
			part_groupIdent := "__stage_0" + part_group.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Part_group{Name: %s}).Stage(stageSet.Stage)", part_groupIdent, __gong__toRawStringLiteral(part_group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", part_groupIdent, __gong__toRawStringLiteral(part_group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", part_groupIdent, __gong__toRawStringLiteral(string(part_group.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", part_groupIdent, __gong__toRawStringLiteral(part_group.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_time = %s", part_groupIdent, __gong__toRawStringLiteral(part_group.Group_time)))
			if part_group.Group_name != nil {
				targetIdent := "__stage_0" + part_group.Group_name.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group_name = %s", part_groupIdent, targetIdent))
			}
			if part_group.Group_name_display != nil {
				targetIdent := "__stage_0" + part_group.Group_name_display.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group_name_display = %s", part_groupIdent, targetIdent))
			}
			if part_group.Group_abbreviation != nil {
				targetIdent := "__stage_0" + part_group.Group_abbreviation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group_abbreviation = %s", part_groupIdent, targetIdent))
			}
			if part_group.Group_abbreviation_display != nil {
				targetIdent := "__stage_0" + part_group.Group_abbreviation_display.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group_abbreviation_display = %s", part_groupIdent, targetIdent))
			}
			if part_group.Group_symbol != nil {
				targetIdent := "__stage_0" + part_group.Group_symbol.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group_symbol = %s", part_groupIdent, targetIdent))
			}
			if part_group.Group_barline != nil {
				targetIdent := "__stage_0" + part_group.Group_barline.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Group_barline = %s", part_groupIdent, targetIdent))
			}
			if part_group.Footnote != nil {
				targetIdent := "__stage_0" + part_group.Footnote.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", part_groupIdent, targetIdent))
			}
			if part_group.Level != nil {
				targetIdent := "__stage_0" + part_group.Level.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", part_groupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		part_linkOrdered := []*Part_link{}
		for part_link := range stageSet.Stage.Part_links {
			part_linkOrdered = append(part_linkOrdered, part_link)
		}
		sort.Slice(part_linkOrdered, func(i, j int) bool {
			return stageSet.Stage.Part_link_stagedOrder[part_linkOrdered[i]] < stageSet.Stage.Part_link_stagedOrder[part_linkOrdered[j]]
		})
		for _, part_link := range part_linkOrdered {
			part_linkIdent := "__stage_0" + part_link.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Part_link{Name: %s}).Stage(stageSet.Stage)", part_linkIdent, __gong__toRawStringLiteral(part_link.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Actuate)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group_link = %s", part_linkIdent, __gong__toRawStringLiteral(part_link.Group_link)))
			for _, elem := range part_link.Instrument_link {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Instrument_link = append(%s.Instrument_link, %s)", part_linkIdent, part_linkIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		part_listOrdered := []*Part_list{}
		for part_list := range stageSet.Stage.Part_lists {
			part_listOrdered = append(part_listOrdered, part_list)
		}
		sort.Slice(part_listOrdered, func(i, j int) bool {
			return stageSet.Stage.Part_list_stagedOrder[part_listOrdered[i]] < stageSet.Stage.Part_list_stagedOrder[part_listOrdered[j]]
		})
		for _, part_list := range part_listOrdered {
			part_listIdent := "__stage_0" + part_list.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Part_list{Name: %s}).Stage(stageSet.Stage)", part_listIdent, __gong__toRawStringLiteral(part_list.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", part_listIdent, __gong__toRawStringLiteral(part_list.Name)))
			if part_list.Part_group != nil {
				targetIdent := "__stage_0" + part_list.Part_group.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_group = %s", part_listIdent, targetIdent))
			}
			if part_list.Score_part != nil {
				targetIdent := "__stage_0" + part_list.Score_part.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Score_part = %s", part_listIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		part_nameOrdered := []*Part_name{}
		for part_name := range stageSet.Stage.Part_names {
			part_nameOrdered = append(part_nameOrdered, part_name)
		}
		sort.Slice(part_nameOrdered, func(i, j int) bool {
			return stageSet.Stage.Part_name_stagedOrder[part_nameOrdered[i]] < stageSet.Stage.Part_name_stagedOrder[part_nameOrdered[j]]
		})
		for _, part_name := range part_nameOrdered {
			part_nameIdent := "__stage_0" + part_name.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Part_name{Name: %s}).Stage(stageSet.Stage)", part_nameIdent, __gong__toRawStringLiteral(part_name.Name)))
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
		part_symbolOrdered := []*Part_symbol{}
		for part_symbol := range stageSet.Stage.Part_symbols {
			part_symbolOrdered = append(part_symbolOrdered, part_symbol)
		}
		sort.Slice(part_symbolOrdered, func(i, j int) bool {
			return stageSet.Stage.Part_symbol_stagedOrder[part_symbolOrdered[i]] < stageSet.Stage.Part_symbol_stagedOrder[part_symbolOrdered[j]]
		})
		for _, part_symbol := range part_symbolOrdered {
			part_symbolIdent := "__stage_0" + part_symbol.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Part_symbol{Name: %s}).Stage(stageSet.Stage)", part_symbolIdent, __gong__toRawStringLiteral(part_symbol.Name)))
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
		part_transposeOrdered := []*Part_transpose{}
		for part_transpose := range stageSet.Stage.Part_transposes {
			part_transposeOrdered = append(part_transposeOrdered, part_transpose)
		}
		sort.Slice(part_transposeOrdered, func(i, j int) bool {
			return stageSet.Stage.Part_transpose_stagedOrder[part_transposeOrdered[i]] < stageSet.Stage.Part_transpose_stagedOrder[part_transposeOrdered[j]]
		})
		for _, part_transpose := range part_transposeOrdered {
			part_transposeIdent := "__stage_0" + part_transpose.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Part_transpose{Name: %s}).Stage(stageSet.Stage)", part_transposeIdent, __gong__toRawStringLiteral(part_transpose.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", part_transposeIdent, __gong__toRawStringLiteral(part_transpose.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Diatonic = %d", part_transposeIdent, part_transpose.Diatonic))
			values.WriteString(fmt.Sprintf("\n\t%s.Chromatic = %s", part_transposeIdent, __gong__toRawStringLiteral(part_transpose.Chromatic)))
			values.WriteString(fmt.Sprintf("\n\t%s.Octave_change = %d", part_transposeIdent, part_transpose.Octave_change))
			values.WriteString(fmt.Sprintf("\n\t%s.Double = %f", part_transposeIdent, part_transpose.Double))
		}
	}
	if stageSet.Stage != nil {
		pedalOrdered := []*Pedal{}
		for pedal := range stageSet.Stage.Pedals {
			pedalOrdered = append(pedalOrdered, pedal)
		}
		sort.Slice(pedalOrdered, func(i, j int) bool {
			return stageSet.Stage.Pedal_stagedOrder[pedalOrdered[i]] < stageSet.Stage.Pedal_stagedOrder[pedalOrdered[j]]
		})
		for _, pedal := range pedalOrdered {
			pedalIdent := "__stage_0" + pedal.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Pedal{Name: %s}).Stage(stageSet.Stage)", pedalIdent, __gong__toRawStringLiteral(pedal.Name)))
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
		pedal_tuningOrdered := []*Pedal_tuning{}
		for pedal_tuning := range stageSet.Stage.Pedal_tunings {
			pedal_tuningOrdered = append(pedal_tuningOrdered, pedal_tuning)
		}
		sort.Slice(pedal_tuningOrdered, func(i, j int) bool {
			return stageSet.Stage.Pedal_tuning_stagedOrder[pedal_tuningOrdered[i]] < stageSet.Stage.Pedal_tuning_stagedOrder[pedal_tuningOrdered[j]]
		})
		for _, pedal_tuning := range pedal_tuningOrdered {
			pedal_tuningIdent := "__stage_0" + pedal_tuning.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Pedal_tuning{Name: %s}).Stage(stageSet.Stage)", pedal_tuningIdent, __gong__toRawStringLiteral(pedal_tuning.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pedal_tuningIdent, __gong__toRawStringLiteral(pedal_tuning.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Pedal_step = %s", pedal_tuningIdent, __gong__toRawStringLiteral(string(pedal_tuning.Pedal_step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Pedal_alter = %s", pedal_tuningIdent, __gong__toRawStringLiteral(pedal_tuning.Pedal_alter)))
		}
	}
	if stageSet.Stage != nil {
		per_minuteOrdered := []*Per_minute{}
		for per_minute := range stageSet.Stage.Per_minutes {
			per_minuteOrdered = append(per_minuteOrdered, per_minute)
		}
		sort.Slice(per_minuteOrdered, func(i, j int) bool {
			return stageSet.Stage.Per_minute_stagedOrder[per_minuteOrdered[i]] < stageSet.Stage.Per_minute_stagedOrder[per_minuteOrdered[j]]
		})
		for _, per_minute := range per_minuteOrdered {
			per_minuteIdent := "__stage_0" + per_minute.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Per_minute{Name: %s}).Stage(stageSet.Stage)", per_minuteIdent, __gong__toRawStringLiteral(per_minute.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", per_minuteIdent, __gong__toRawStringLiteral(per_minute.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", per_minuteIdent, __gong__toRawStringLiteral(per_minute.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", per_minuteIdent, __gong__toRawStringLiteral(per_minute.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", per_minuteIdent, __gong__toRawStringLiteral(per_minute.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", per_minuteIdent, __gong__toRawStringLiteral(per_minute.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", per_minuteIdent, __gong__toRawStringLiteral(per_minute.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		percussionOrdered := []*Percussion{}
		for percussion := range stageSet.Stage.Percussions {
			percussionOrdered = append(percussionOrdered, percussion)
		}
		sort.Slice(percussionOrdered, func(i, j int) bool {
			return stageSet.Stage.Percussion_stagedOrder[percussionOrdered[i]] < stageSet.Stage.Percussion_stagedOrder[percussionOrdered[j]]
		})
		for _, percussion := range percussionOrdered {
			percussionIdent := "__stage_0" + percussion.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Percussion{Name: %s}).Stage(stageSet.Stage)", percussionIdent, __gong__toRawStringLiteral(percussion.Name)))
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
				targetIdent := "__stage_0" + percussion.Glass.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Glass = %s", percussionIdent, targetIdent))
			}
			if percussion.Metal != nil {
				targetIdent := "__stage_0" + percussion.Metal.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metal = %s", percussionIdent, targetIdent))
			}
			if percussion.Wood != nil {
				targetIdent := "__stage_0" + percussion.Wood.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wood = %s", percussionIdent, targetIdent))
			}
			if percussion.Pitched != nil {
				targetIdent := "__stage_0" + percussion.Pitched.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pitched = %s", percussionIdent, targetIdent))
			}
			if percussion.Membrane != nil {
				targetIdent := "__stage_0" + percussion.Membrane.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Membrane = %s", percussionIdent, targetIdent))
			}
			if percussion.Effect != nil {
				targetIdent := "__stage_0" + percussion.Effect.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Effect = %s", percussionIdent, targetIdent))
			}
			if percussion.Timpani != nil {
				targetIdent := "__stage_0" + percussion.Timpani.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Timpani = %s", percussionIdent, targetIdent))
			}
			if percussion.Beater != nil {
				targetIdent := "__stage_0" + percussion.Beater.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Beater = %s", percussionIdent, targetIdent))
			}
			if percussion.Stick != nil {
				targetIdent := "__stage_0" + percussion.Stick.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stick = %s", percussionIdent, targetIdent))
			}
			if percussion.Other_percussion != nil {
				targetIdent := "__stage_0" + percussion.Other_percussion.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_percussion = %s", percussionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		pitchOrdered := []*Pitch{}
		for pitch := range stageSet.Stage.Pitchs {
			pitchOrdered = append(pitchOrdered, pitch)
		}
		sort.Slice(pitchOrdered, func(i, j int) bool {
			return stageSet.Stage.Pitch_stagedOrder[pitchOrdered[i]] < stageSet.Stage.Pitch_stagedOrder[pitchOrdered[j]]
		})
		for _, pitch := range pitchOrdered {
			pitchIdent := "__stage_0" + pitch.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Pitch{Name: %s}).Stage(stageSet.Stage)", pitchIdent, __gong__toRawStringLiteral(pitch.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pitchIdent, __gong__toRawStringLiteral(pitch.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Step = %s", pitchIdent, __gong__toRawStringLiteral(string(pitch.Step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Alter = %s", pitchIdent, __gong__toRawStringLiteral(pitch.Alter)))
			values.WriteString(fmt.Sprintf("\n\t%s.Octave = %d", pitchIdent, pitch.Octave))
		}
	}
	if stageSet.Stage != nil {
		pitchedOrdered := []*Pitched{}
		for pitched := range stageSet.Stage.Pitcheds {
			pitchedOrdered = append(pitchedOrdered, pitched)
		}
		sort.Slice(pitchedOrdered, func(i, j int) bool {
			return stageSet.Stage.Pitched_stagedOrder[pitchedOrdered[i]] < stageSet.Stage.Pitched_stagedOrder[pitchedOrdered[j]]
		})
		for _, pitched := range pitchedOrdered {
			pitchedIdent := "__stage_0" + pitched.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Pitched{Name: %s}).Stage(stageSet.Stage)", pitchedIdent, __gong__toRawStringLiteral(pitched.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pitchedIdent, __gong__toRawStringLiteral(pitched.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", pitchedIdent, __gong__toRawStringLiteral(pitched.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", pitchedIdent, __gong__toRawStringLiteral(pitched.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		placement_textOrdered := []*Placement_text{}
		for placement_text := range stageSet.Stage.Placement_texts {
			placement_textOrdered = append(placement_textOrdered, placement_text)
		}
		sort.Slice(placement_textOrdered, func(i, j int) bool {
			return stageSet.Stage.Placement_text_stagedOrder[placement_textOrdered[i]] < stageSet.Stage.Placement_text_stagedOrder[placement_textOrdered[j]]
		})
		for _, placement_text := range placement_textOrdered {
			placement_textIdent := "__stage_0" + placement_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Placement_text{Name: %s}).Stage(stageSet.Stage)", placement_textIdent, __gong__toRawStringLiteral(placement_text.Name)))
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
		playOrdered := []*Play{}
		for play := range stageSet.Stage.Plays {
			playOrdered = append(playOrdered, play)
		}
		sort.Slice(playOrdered, func(i, j int) bool {
			return stageSet.Stage.Play_stagedOrder[playOrdered[i]] < stageSet.Stage.Play_stagedOrder[playOrdered[j]]
		})
		for _, play := range playOrdered {
			playIdent := "__stage_0" + play.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Play{Name: %s}).Stage(stageSet.Stage)", playIdent, __gong__toRawStringLiteral(play.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", playIdent, __gong__toRawStringLiteral(play.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", playIdent, __gong__toRawStringLiteral(play.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ipa = %s", playIdent, __gong__toRawStringLiteral(play.Ipa)))
			values.WriteString(fmt.Sprintf("\n\t%s.Mute = %s", playIdent, __gong__toRawStringLiteral(play.Mute)))
			values.WriteString(fmt.Sprintf("\n\t%s.Semi_pitched = %s", playIdent, __gong__toRawStringLiteral(play.Semi_pitched)))
			for _, elem := range play.Other_play {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_play = append(%s.Other_play, %s)", playIdent, playIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		playerOrdered := []*Player{}
		for player := range stageSet.Stage.Players {
			playerOrdered = append(playerOrdered, player)
		}
		sort.Slice(playerOrdered, func(i, j int) bool {
			return stageSet.Stage.Player_stagedOrder[playerOrdered[i]] < stageSet.Stage.Player_stagedOrder[playerOrdered[j]]
		})
		for _, player := range playerOrdered {
			playerIdent := "__stage_0" + player.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Player{Name: %s}).Stage(stageSet.Stage)", playerIdent, __gong__toRawStringLiteral(player.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", playerIdent, __gong__toRawStringLiteral(player.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", playerIdent, __gong__toRawStringLiteral(player.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Player_name = %s", playerIdent, __gong__toRawStringLiteral(player.Player_name)))
		}
	}
	if stageSet.Stage != nil {
		principal_voiceOrdered := []*Principal_voice{}
		for principal_voice := range stageSet.Stage.Principal_voices {
			principal_voiceOrdered = append(principal_voiceOrdered, principal_voice)
		}
		sort.Slice(principal_voiceOrdered, func(i, j int) bool {
			return stageSet.Stage.Principal_voice_stagedOrder[principal_voiceOrdered[i]] < stageSet.Stage.Principal_voice_stagedOrder[principal_voiceOrdered[j]]
		})
		for _, principal_voice := range principal_voiceOrdered {
			principal_voiceIdent := "__stage_0" + principal_voice.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Principal_voice{Name: %s}).Stage(stageSet.Stage)", principal_voiceIdent, __gong__toRawStringLiteral(principal_voice.Name)))
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
		printOrdered := []*Print{}
		for print := range stageSet.Stage.Prints {
			printOrdered = append(printOrdered, print)
		}
		sort.Slice(printOrdered, func(i, j int) bool {
			return stageSet.Stage.Print_stagedOrder[printOrdered[i]] < stageSet.Stage.Print_stagedOrder[printOrdered[j]]
		})
		for _, print := range printOrdered {
			printIdent := "__stage_0" + print.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Print{Name: %s}).Stage(stageSet.Stage)", printIdent, __gong__toRawStringLiteral(print.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", printIdent, __gong__toRawStringLiteral(print.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff_spacing = %s", printIdent, __gong__toRawStringLiteral(print.Staff_spacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.New_system = %s", printIdent, __gong__toRawStringLiteral(string(print.New_system))))
			values.WriteString(fmt.Sprintf("\n\t%s.New_page = %s", printIdent, __gong__toRawStringLiteral(string(print.New_page))))
			values.WriteString(fmt.Sprintf("\n\t%s.Blank_page = %d", printIdent, print.Blank_page))
			values.WriteString(fmt.Sprintf("\n\t%s.Page_number = %s", printIdent, __gong__toRawStringLiteral(print.Page_number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", printIdent, __gong__toRawStringLiteral(print.Id)))
			if print.Page_layout != nil {
				targetIdent := "__stage_0" + print.Page_layout.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Page_layout = %s", printIdent, targetIdent))
			}
			if print.System_layout != nil {
				targetIdent := "__stage_0" + print.System_layout.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System_layout = %s", printIdent, targetIdent))
			}
			for _, elem := range print.Staff_layout {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staff_layout = append(%s.Staff_layout, %s)", printIdent, printIdent, targetIdent))
			}
			if print.Measure_layout != nil {
				targetIdent := "__stage_0" + print.Measure_layout.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure_layout = %s", printIdent, targetIdent))
			}
			if print.Measure_numbering != nil {
				targetIdent := "__stage_0" + print.Measure_numbering.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure_numbering = %s", printIdent, targetIdent))
			}
			if print.Part_name_display != nil {
				targetIdent := "__stage_0" + print.Part_name_display.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_name_display = %s", printIdent, targetIdent))
			}
			if print.Part_abbreviation_display != nil {
				targetIdent := "__stage_0" + print.Part_abbreviation_display.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_abbreviation_display = %s", printIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		releaseOrdered := []*Release{}
		for release := range stageSet.Stage.Releases {
			releaseOrdered = append(releaseOrdered, release)
		}
		sort.Slice(releaseOrdered, func(i, j int) bool {
			return stageSet.Stage.Release_stagedOrder[releaseOrdered[i]] < stageSet.Stage.Release_stagedOrder[releaseOrdered[j]]
		})
		for _, release := range releaseOrdered {
			releaseIdent := "__stage_0" + release.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Release{Name: %s}).Stage(stageSet.Stage)", releaseIdent, __gong__toRawStringLiteral(release.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", releaseIdent, __gong__toRawStringLiteral(release.Name)))
		}
	}
	if stageSet.Stage != nil {
		repeatOrdered := []*Repeat{}
		for repeat := range stageSet.Stage.Repeats {
			repeatOrdered = append(repeatOrdered, repeat)
		}
		sort.Slice(repeatOrdered, func(i, j int) bool {
			return stageSet.Stage.Repeat_stagedOrder[repeatOrdered[i]] < stageSet.Stage.Repeat_stagedOrder[repeatOrdered[j]]
		})
		for _, repeat := range repeatOrdered {
			repeatIdent := "__stage_0" + repeat.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Repeat{Name: %s}).Stage(stageSet.Stage)", repeatIdent, __gong__toRawStringLiteral(repeat.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", repeatIdent, __gong__toRawStringLiteral(repeat.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", repeatIdent, __gong__toRawStringLiteral(repeat.Direction)))
			values.WriteString(fmt.Sprintf("\n\t%s.Times = %d", repeatIdent, repeat.Times))
			values.WriteString(fmt.Sprintf("\n\t%s.After_jump = %s", repeatIdent, __gong__toRawStringLiteral(string(repeat.After_jump))))
			values.WriteString(fmt.Sprintf("\n\t%s.Winged = %s", repeatIdent, __gong__toRawStringLiteral(repeat.Winged)))
		}
	}
	if stageSet.Stage != nil {
		restOrdered := []*Rest{}
		for rest := range stageSet.Stage.Rests {
			restOrdered = append(restOrdered, rest)
		}
		sort.Slice(restOrdered, func(i, j int) bool {
			return stageSet.Stage.Rest_stagedOrder[restOrdered[i]] < stageSet.Stage.Rest_stagedOrder[restOrdered[j]]
		})
		for _, rest := range restOrdered {
			restIdent := "__stage_0" + rest.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Rest{Name: %s}).Stage(stageSet.Stage)", restIdent, __gong__toRawStringLiteral(rest.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", restIdent, __gong__toRawStringLiteral(rest.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Measure = %s", restIdent, __gong__toRawStringLiteral(string(rest.Measure))))
			values.WriteString(fmt.Sprintf("\n\t%s.Display_step = %s", restIdent, __gong__toRawStringLiteral(string(rest.Display_step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Display_octave = %d", restIdent, rest.Display_octave))
		}
	}
	if stageSet.Stage != nil {
		rootOrdered := []*Root{}
		for root := range stageSet.Stage.Roots {
			rootOrdered = append(rootOrdered, root)
		}
		sort.Slice(rootOrdered, func(i, j int) bool {
			return stageSet.Stage.Root_stagedOrder[rootOrdered[i]] < stageSet.Stage.Root_stagedOrder[rootOrdered[j]]
		})
		for _, root := range rootOrdered {
			rootIdent := "__stage_0" + root.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Root{Name: %s}).Stage(stageSet.Stage)", rootIdent, __gong__toRawStringLiteral(root.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rootIdent, __gong__toRawStringLiteral(root.Name)))
			if root.Root_step != nil {
				targetIdent := "__stage_0" + root.Root_step.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Root_step = %s", rootIdent, targetIdent))
			}
			if root.Root_alter != nil {
				targetIdent := "__stage_0" + root.Root_alter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Root_alter = %s", rootIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		root_stepOrdered := []*Root_step{}
		for root_step := range stageSet.Stage.Root_steps {
			root_stepOrdered = append(root_stepOrdered, root_step)
		}
		sort.Slice(root_stepOrdered, func(i, j int) bool {
			return stageSet.Stage.Root_step_stagedOrder[root_stepOrdered[i]] < stageSet.Stage.Root_step_stagedOrder[root_stepOrdered[j]]
		})
		for _, root_step := range root_stepOrdered {
			root_stepIdent := "__stage_0" + root_step.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Root_step{Name: %s}).Stage(stageSet.Stage)", root_stepIdent, __gong__toRawStringLiteral(root_step.Name)))
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
		scalingOrdered := []*Scaling{}
		for scaling := range stageSet.Stage.Scalings {
			scalingOrdered = append(scalingOrdered, scaling)
		}
		sort.Slice(scalingOrdered, func(i, j int) bool {
			return stageSet.Stage.Scaling_stagedOrder[scalingOrdered[i]] < stageSet.Stage.Scaling_stagedOrder[scalingOrdered[j]]
		})
		for _, scaling := range scalingOrdered {
			scalingIdent := "__stage_0" + scaling.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Scaling{Name: %s}).Stage(stageSet.Stage)", scalingIdent, __gong__toRawStringLiteral(scaling.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scalingIdent, __gong__toRawStringLiteral(scaling.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Millimeters = %s", scalingIdent, __gong__toRawStringLiteral(scaling.Millimeters)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tenths = %s", scalingIdent, __gong__toRawStringLiteral(scaling.Tenths)))
		}
	}
	if stageSet.Stage != nil {
		scordaturaOrdered := []*Scordatura{}
		for scordatura := range stageSet.Stage.Scordaturas {
			scordaturaOrdered = append(scordaturaOrdered, scordatura)
		}
		sort.Slice(scordaturaOrdered, func(i, j int) bool {
			return stageSet.Stage.Scordatura_stagedOrder[scordaturaOrdered[i]] < stageSet.Stage.Scordatura_stagedOrder[scordaturaOrdered[j]]
		})
		for _, scordatura := range scordaturaOrdered {
			scordaturaIdent := "__stage_0" + scordatura.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Scordatura{Name: %s}).Stage(stageSet.Stage)", scordaturaIdent, __gong__toRawStringLiteral(scordatura.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scordaturaIdent, __gong__toRawStringLiteral(scordatura.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", scordaturaIdent, __gong__toRawStringLiteral(scordatura.Id)))
			for _, elem := range scordatura.Accord {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accord = append(%s.Accord, %s)", scordaturaIdent, scordaturaIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		score_instrumentOrdered := []*Score_instrument{}
		for score_instrument := range stageSet.Stage.Score_instruments {
			score_instrumentOrdered = append(score_instrumentOrdered, score_instrument)
		}
		sort.Slice(score_instrumentOrdered, func(i, j int) bool {
			return stageSet.Stage.Score_instrument_stagedOrder[score_instrumentOrdered[i]] < stageSet.Stage.Score_instrument_stagedOrder[score_instrumentOrdered[j]]
		})
		for _, score_instrument := range score_instrumentOrdered {
			score_instrumentIdent := "__stage_0" + score_instrument.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Score_instrument{Name: %s}).Stage(stageSet.Stage)", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Instrument_name = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Instrument_name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Instrument_abbreviation = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Instrument_abbreviation)))
			values.WriteString(fmt.Sprintf("\n\t%s.Instrument_sound = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Instrument_sound)))
			values.WriteString(fmt.Sprintf("\n\t%s.Solo = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Solo)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ensemble = %s", score_instrumentIdent, __gong__toRawStringLiteral(score_instrument.Ensemble)))
			if score_instrument.Virtual_instrument != nil {
				targetIdent := "__stage_0" + score_instrument.Virtual_instrument.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Virtual_instrument = %s", score_instrumentIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		score_partOrdered := []*Score_part{}
		for score_part := range stageSet.Stage.Score_parts {
			score_partOrdered = append(score_partOrdered, score_part)
		}
		sort.Slice(score_partOrdered, func(i, j int) bool {
			return stageSet.Stage.Score_part_stagedOrder[score_partOrdered[i]] < stageSet.Stage.Score_part_stagedOrder[score_partOrdered[j]]
		})
		for _, score_part := range score_partOrdered {
			score_partIdent := "__stage_0" + score_part.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Score_part{Name: %s}).Stage(stageSet.Stage)", score_partIdent, __gong__toRawStringLiteral(score_part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", score_partIdent, __gong__toRawStringLiteral(score_part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", score_partIdent, __gong__toRawStringLiteral(score_part.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group = %s", score_partIdent, __gong__toRawStringLiteral(score_part.Group)))
			if score_part.Identification != nil {
				targetIdent := "__stage_0" + score_part.Identification.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identification = %s", score_partIdent, targetIdent))
			}
			for _, elem := range score_part.Part_link {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_link = append(%s.Part_link, %s)", score_partIdent, score_partIdent, targetIdent))
			}
			if score_part.Part_name != nil {
				targetIdent := "__stage_0" + score_part.Part_name.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_name = %s", score_partIdent, targetIdent))
			}
			if score_part.Part_name_display != nil {
				targetIdent := "__stage_0" + score_part.Part_name_display.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_name_display = %s", score_partIdent, targetIdent))
			}
			if score_part.Part_abbreviation != nil {
				targetIdent := "__stage_0" + score_part.Part_abbreviation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_abbreviation = %s", score_partIdent, targetIdent))
			}
			if score_part.Part_abbreviation_display != nil {
				targetIdent := "__stage_0" + score_part.Part_abbreviation_display.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_abbreviation_display = %s", score_partIdent, targetIdent))
			}
			for _, elem := range score_part.Score_instrument {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Score_instrument = append(%s.Score_instrument, %s)", score_partIdent, score_partIdent, targetIdent))
			}
			for _, elem := range score_part.Player {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Player = append(%s.Player, %s)", score_partIdent, score_partIdent, targetIdent))
			}
			for _, elem := range score_part.Midi_device {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Midi_device = append(%s.Midi_device, %s)", score_partIdent, score_partIdent, targetIdent))
			}
			for _, elem := range score_part.Midi_instrument {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Midi_instrument = append(%s.Midi_instrument, %s)", score_partIdent, score_partIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		score_partwiseOrdered := []*Score_partwise{}
		for score_partwise := range stageSet.Stage.Score_partwises {
			score_partwiseOrdered = append(score_partwiseOrdered, score_partwise)
		}
		sort.Slice(score_partwiseOrdered, func(i, j int) bool {
			return stageSet.Stage.Score_partwise_stagedOrder[score_partwiseOrdered[i]] < stageSet.Stage.Score_partwise_stagedOrder[score_partwiseOrdered[j]]
		})
		for _, score_partwise := range score_partwiseOrdered {
			score_partwiseIdent := "__stage_0" + score_partwise.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Score_partwise{Name: %s}).Stage(stageSet.Stage)", score_partwiseIdent, __gong__toRawStringLiteral(score_partwise.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", score_partwiseIdent, __gong__toRawStringLiteral(score_partwise.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Version = %s", score_partwiseIdent, __gong__toRawStringLiteral(score_partwise.Version)))
			values.WriteString(fmt.Sprintf("\n\t%s.Movement_number = %s", score_partwiseIdent, __gong__toRawStringLiteral(score_partwise.Movement_number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Movement_title = %s", score_partwiseIdent, __gong__toRawStringLiteral(score_partwise.Movement_title)))
			if score_partwise.Work != nil {
				targetIdent := "__stage_0" + score_partwise.Work.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Work = %s", score_partwiseIdent, targetIdent))
			}
			if score_partwise.Identification != nil {
				targetIdent := "__stage_0" + score_partwise.Identification.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identification = %s", score_partwiseIdent, targetIdent))
			}
			if score_partwise.Defaults != nil {
				targetIdent := "__stage_0" + score_partwise.Defaults.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Defaults = %s", score_partwiseIdent, targetIdent))
			}
			for _, elem := range score_partwise.Credit {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit = append(%s.Credit, %s)", score_partwiseIdent, score_partwiseIdent, targetIdent))
			}
			if score_partwise.Part_list != nil {
				targetIdent := "__stage_0" + score_partwise.Part_list.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_list = %s", score_partwiseIdent, targetIdent))
			}
			for _, elem := range score_partwise.Part {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = append(%s.Part, %s)", score_partwiseIdent, score_partwiseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		score_timewiseOrdered := []*Score_timewise{}
		for score_timewise := range stageSet.Stage.Score_timewises {
			score_timewiseOrdered = append(score_timewiseOrdered, score_timewise)
		}
		sort.Slice(score_timewiseOrdered, func(i, j int) bool {
			return stageSet.Stage.Score_timewise_stagedOrder[score_timewiseOrdered[i]] < stageSet.Stage.Score_timewise_stagedOrder[score_timewiseOrdered[j]]
		})
		for _, score_timewise := range score_timewiseOrdered {
			score_timewiseIdent := "__stage_0" + score_timewise.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Score_timewise{Name: %s}).Stage(stageSet.Stage)", score_timewiseIdent, __gong__toRawStringLiteral(score_timewise.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", score_timewiseIdent, __gong__toRawStringLiteral(score_timewise.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Version = %s", score_timewiseIdent, __gong__toRawStringLiteral(score_timewise.Version)))
			values.WriteString(fmt.Sprintf("\n\t%s.Movement_number = %s", score_timewiseIdent, __gong__toRawStringLiteral(score_timewise.Movement_number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Movement_title = %s", score_timewiseIdent, __gong__toRawStringLiteral(score_timewise.Movement_title)))
			if score_timewise.Work != nil {
				targetIdent := "__stage_0" + score_timewise.Work.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Work = %s", score_timewiseIdent, targetIdent))
			}
			if score_timewise.Identification != nil {
				targetIdent := "__stage_0" + score_timewise.Identification.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identification = %s", score_timewiseIdent, targetIdent))
			}
			if score_timewise.Defaults != nil {
				targetIdent := "__stage_0" + score_timewise.Defaults.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Defaults = %s", score_timewiseIdent, targetIdent))
			}
			for _, elem := range score_timewise.Credit {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit = append(%s.Credit, %s)", score_timewiseIdent, score_timewiseIdent, targetIdent))
			}
			if score_timewise.Part_list != nil {
				targetIdent := "__stage_0" + score_timewise.Part_list.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part_list = %s", score_timewiseIdent, targetIdent))
			}
			for _, elem := range score_timewise.Measure {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure = append(%s.Measure, %s)", score_timewiseIdent, score_timewiseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		segnoOrdered := []*Segno{}
		for segno := range stageSet.Stage.Segnos {
			segnoOrdered = append(segnoOrdered, segno)
		}
		sort.Slice(segnoOrdered, func(i, j int) bool {
			return stageSet.Stage.Segno_stagedOrder[segnoOrdered[i]] < stageSet.Stage.Segno_stagedOrder[segnoOrdered[j]]
		})
		for _, segno := range segnoOrdered {
			segnoIdent := "__stage_0" + segno.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Segno{Name: %s}).Stage(stageSet.Stage)", segnoIdent, __gong__toRawStringLiteral(segno.Name)))
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
		slashOrdered := []*Slash{}
		for slash := range stageSet.Stage.Slashs {
			slashOrdered = append(slashOrdered, slash)
		}
		sort.Slice(slashOrdered, func(i, j int) bool {
			return stageSet.Stage.Slash_stagedOrder[slashOrdered[i]] < stageSet.Stage.Slash_stagedOrder[slashOrdered[j]]
		})
		for _, slash := range slashOrdered {
			slashIdent := "__stage_0" + slash.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Slash{Name: %s}).Stage(stageSet.Stage)", slashIdent, __gong__toRawStringLiteral(slash.Name)))
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
		slideOrdered := []*Slide{}
		for slide := range stageSet.Stage.Slides {
			slideOrdered = append(slideOrdered, slide)
		}
		sort.Slice(slideOrdered, func(i, j int) bool {
			return stageSet.Stage.Slide_stagedOrder[slideOrdered[i]] < stageSet.Stage.Slide_stagedOrder[slideOrdered[j]]
		})
		for _, slide := range slideOrdered {
			slideIdent := "__stage_0" + slide.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Slide{Name: %s}).Stage(stageSet.Stage)", slideIdent, __gong__toRawStringLiteral(slide.Name)))
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
		slurOrdered := []*Slur{}
		for slur := range stageSet.Stage.Slurs {
			slurOrdered = append(slurOrdered, slur)
		}
		sort.Slice(slurOrdered, func(i, j int) bool {
			return stageSet.Stage.Slur_stagedOrder[slurOrdered[i]] < stageSet.Stage.Slur_stagedOrder[slurOrdered[j]]
		})
		for _, slur := range slurOrdered {
			slurIdent := "__stage_0" + slur.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Slur{Name: %s}).Stage(stageSet.Stage)", slurIdent, __gong__toRawStringLiteral(slur.Name)))
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
		soundOrdered := []*Sound{}
		for sound := range stageSet.Stage.Sounds {
			soundOrdered = append(soundOrdered, sound)
		}
		sort.Slice(soundOrdered, func(i, j int) bool {
			return stageSet.Stage.Sound_stagedOrder[soundOrdered[i]] < stageSet.Stage.Sound_stagedOrder[soundOrdered[j]]
		})
		for _, sound := range soundOrdered {
			soundIdent := "__stage_0" + sound.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Sound{Name: %s}).Stage(stageSet.Stage)", soundIdent, __gong__toRawStringLiteral(sound.Name)))
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
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Instrument_change = append(%s.Instrument_change, %s)", soundIdent, soundIdent, targetIdent))
			}
			for _, elem := range sound.Midi_device {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Midi_device = append(%s.Midi_device, %s)", soundIdent, soundIdent, targetIdent))
			}
			for _, elem := range sound.Midi_instrument {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Midi_instrument = append(%s.Midi_instrument, %s)", soundIdent, soundIdent, targetIdent))
			}
			for _, elem := range sound.Play {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Play = append(%s.Play, %s)", soundIdent, soundIdent, targetIdent))
			}
			if sound.Swing != nil {
				targetIdent := "__stage_0" + sound.Swing.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Swing = %s", soundIdent, targetIdent))
			}
			if sound.Offset != nil {
				targetIdent := "__stage_0" + sound.Offset.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", soundIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		staff_detailsOrdered := []*Staff_details{}
		for staff_details := range stageSet.Stage.Staff_detailss {
			staff_detailsOrdered = append(staff_detailsOrdered, staff_details)
		}
		sort.Slice(staff_detailsOrdered, func(i, j int) bool {
			return stageSet.Stage.Staff_details_stagedOrder[staff_detailsOrdered[i]] < stageSet.Stage.Staff_details_stagedOrder[staff_detailsOrdered[j]]
		})
		for _, staff_details := range staff_detailsOrdered {
			staff_detailsIdent := "__stage_0" + staff_details.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Staff_details{Name: %s}).Stage(stageSet.Stage)", staff_detailsIdent, __gong__toRawStringLiteral(staff_details.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staff_detailsIdent, __gong__toRawStringLiteral(staff_details.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", staff_detailsIdent, staff_details.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Show_frets = %s", staff_detailsIdent, __gong__toRawStringLiteral(staff_details.Show_frets)))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_object = %s", staff_detailsIdent, __gong__toRawStringLiteral(string(staff_details.Print_object))))
			values.WriteString(fmt.Sprintf("\n\t%s.Print_spacing = %s", staff_detailsIdent, __gong__toRawStringLiteral(string(staff_details.Print_spacing))))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff_type = %s", staff_detailsIdent, __gong__toRawStringLiteral(staff_details.Staff_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff_lines = %d", staff_detailsIdent, staff_details.Staff_lines))
			values.WriteString(fmt.Sprintf("\n\t%s.Capo = %d", staff_detailsIdent, staff_details.Capo))
			for _, elem := range staff_details.Line_detail {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Line_detail = append(%s.Line_detail, %s)", staff_detailsIdent, staff_detailsIdent, targetIdent))
			}
			for _, elem := range staff_details.Staff_tuning {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staff_tuning = append(%s.Staff_tuning, %s)", staff_detailsIdent, staff_detailsIdent, targetIdent))
			}
			if staff_details.Staff_size != nil {
				targetIdent := "__stage_0" + staff_details.Staff_size.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staff_size = %s", staff_detailsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		staff_divideOrdered := []*Staff_divide{}
		for staff_divide := range stageSet.Stage.Staff_divides {
			staff_divideOrdered = append(staff_divideOrdered, staff_divide)
		}
		sort.Slice(staff_divideOrdered, func(i, j int) bool {
			return stageSet.Stage.Staff_divide_stagedOrder[staff_divideOrdered[i]] < stageSet.Stage.Staff_divide_stagedOrder[staff_divideOrdered[j]]
		})
		for _, staff_divide := range staff_divideOrdered {
			staff_divideIdent := "__stage_0" + staff_divide.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Staff_divide{Name: %s}).Stage(stageSet.Stage)", staff_divideIdent, __gong__toRawStringLiteral(staff_divide.Name)))
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
		staff_layoutOrdered := []*Staff_layout{}
		for staff_layout := range stageSet.Stage.Staff_layouts {
			staff_layoutOrdered = append(staff_layoutOrdered, staff_layout)
		}
		sort.Slice(staff_layoutOrdered, func(i, j int) bool {
			return stageSet.Stage.Staff_layout_stagedOrder[staff_layoutOrdered[i]] < stageSet.Stage.Staff_layout_stagedOrder[staff_layoutOrdered[j]]
		})
		for _, staff_layout := range staff_layoutOrdered {
			staff_layoutIdent := "__stage_0" + staff_layout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Staff_layout{Name: %s}).Stage(stageSet.Stage)", staff_layoutIdent, __gong__toRawStringLiteral(staff_layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staff_layoutIdent, __gong__toRawStringLiteral(staff_layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", staff_layoutIdent, staff_layout.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff_distance = %s", staff_layoutIdent, __gong__toRawStringLiteral(staff_layout.Staff_distance)))
		}
	}
	if stageSet.Stage != nil {
		staff_sizeOrdered := []*Staff_size{}
		for staff_size := range stageSet.Stage.Staff_sizes {
			staff_sizeOrdered = append(staff_sizeOrdered, staff_size)
		}
		sort.Slice(staff_sizeOrdered, func(i, j int) bool {
			return stageSet.Stage.Staff_size_stagedOrder[staff_sizeOrdered[i]] < stageSet.Stage.Staff_size_stagedOrder[staff_sizeOrdered[j]]
		})
		for _, staff_size := range staff_sizeOrdered {
			staff_sizeIdent := "__stage_0" + staff_size.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Staff_size{Name: %s}).Stage(stageSet.Stage)", staff_sizeIdent, __gong__toRawStringLiteral(staff_size.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staff_sizeIdent, __gong__toRawStringLiteral(staff_size.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Scaling = %s", staff_sizeIdent, __gong__toRawStringLiteral(staff_size.Scaling)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", staff_sizeIdent, __gong__toRawStringLiteral(staff_size.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		staff_tuningOrdered := []*Staff_tuning{}
		for staff_tuning := range stageSet.Stage.Staff_tunings {
			staff_tuningOrdered = append(staff_tuningOrdered, staff_tuning)
		}
		sort.Slice(staff_tuningOrdered, func(i, j int) bool {
			return stageSet.Stage.Staff_tuning_stagedOrder[staff_tuningOrdered[i]] < stageSet.Stage.Staff_tuning_stagedOrder[staff_tuningOrdered[j]]
		})
		for _, staff_tuning := range staff_tuningOrdered {
			staff_tuningIdent := "__stage_0" + staff_tuning.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Staff_tuning{Name: %s}).Stage(stageSet.Stage)", staff_tuningIdent, __gong__toRawStringLiteral(staff_tuning.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staff_tuningIdent, __gong__toRawStringLiteral(staff_tuning.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", staff_tuningIdent, staff_tuning.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.Tuning_step = %s", staff_tuningIdent, __gong__toRawStringLiteral(string(staff_tuning.Tuning_step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Tuning_alter = %s", staff_tuningIdent, __gong__toRawStringLiteral(staff_tuning.Tuning_alter)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tuning_octave = %d", staff_tuningIdent, staff_tuning.Tuning_octave))
		}
	}
	if stageSet.Stage != nil {
		stemOrdered := []*Stem{}
		for stem := range stageSet.Stage.Stems {
			stemOrdered = append(stemOrdered, stem)
		}
		sort.Slice(stemOrdered, func(i, j int) bool {
			return stageSet.Stage.Stem_stagedOrder[stemOrdered[i]] < stageSet.Stage.Stem_stagedOrder[stemOrdered[j]]
		})
		for _, stem := range stemOrdered {
			stemIdent := "__stage_0" + stem.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Stem{Name: %s}).Stage(stageSet.Stage)", stemIdent, __gong__toRawStringLiteral(stem.Name)))
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
		stickOrdered := []*Stick{}
		for stick := range stageSet.Stage.Sticks {
			stickOrdered = append(stickOrdered, stick)
		}
		sort.Slice(stickOrdered, func(i, j int) bool {
			return stageSet.Stage.Stick_stagedOrder[stickOrdered[i]] < stageSet.Stage.Stick_stagedOrder[stickOrdered[j]]
		})
		for _, stick := range stickOrdered {
			stickIdent := "__stage_0" + stick.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Stick{Name: %s}).Stage(stageSet.Stage)", stickIdent, __gong__toRawStringLiteral(stick.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stickIdent, __gong__toRawStringLiteral(stick.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tip = %s", stickIdent, __gong__toRawStringLiteral(string(stick.Tip))))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", stickIdent, __gong__toRawStringLiteral(string(stick.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.Dashed_circle = %s", stickIdent, __gong__toRawStringLiteral(string(stick.Dashed_circle))))
			values.WriteString(fmt.Sprintf("\n\t%s.Stick_type = %s", stickIdent, __gong__toRawStringLiteral(stick.Stick_type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Stick_material = %s", stickIdent, __gong__toRawStringLiteral(stick.Stick_material)))
		}
	}
	if stageSet.Stage != nil {
		string_muteOrdered := []*String_mute{}
		for string_mute := range stageSet.Stage.String_mutes {
			string_muteOrdered = append(string_muteOrdered, string_mute)
		}
		sort.Slice(string_muteOrdered, func(i, j int) bool {
			return stageSet.Stage.String_mute_stagedOrder[string_muteOrdered[i]] < stageSet.Stage.String_mute_stagedOrder[string_muteOrdered[j]]
		})
		for _, string_mute := range string_muteOrdered {
			string_muteIdent := "__stage_0" + string_mute.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.String_mute{Name: %s}).Stage(stageSet.Stage)", string_muteIdent, __gong__toRawStringLiteral(string_mute.Name)))
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
		string_typeOrdered := []*String_type{}
		for string_type := range stageSet.Stage.String_types {
			string_typeOrdered = append(string_typeOrdered, string_type)
		}
		sort.Slice(string_typeOrdered, func(i, j int) bool {
			return stageSet.Stage.String_type_stagedOrder[string_typeOrdered[i]] < stageSet.Stage.String_type_stagedOrder[string_typeOrdered[j]]
		})
		for _, string_type := range string_typeOrdered {
			string_typeIdent := "__stage_0" + string_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.String_type{Name: %s}).Stage(stageSet.Stage)", string_typeIdent, __gong__toRawStringLiteral(string_type.Name)))
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
		strong_accentOrdered := []*Strong_accent{}
		for strong_accent := range stageSet.Stage.Strong_accents {
			strong_accentOrdered = append(strong_accentOrdered, strong_accent)
		}
		sort.Slice(strong_accentOrdered, func(i, j int) bool {
			return stageSet.Stage.Strong_accent_stagedOrder[strong_accentOrdered[i]] < stageSet.Stage.Strong_accent_stagedOrder[strong_accentOrdered[j]]
		})
		for _, strong_accent := range strong_accentOrdered {
			strong_accentIdent := "__stage_0" + strong_accent.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Strong_accent{Name: %s}).Stage(stageSet.Stage)", strong_accentIdent, __gong__toRawStringLiteral(strong_accent.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", strong_accentIdent, __gong__toRawStringLiteral(strong_accent.Name)))
		}
	}
	if stageSet.Stage != nil {
		style_textOrdered := []*Style_text{}
		for style_text := range stageSet.Stage.Style_texts {
			style_textOrdered = append(style_textOrdered, style_text)
		}
		sort.Slice(style_textOrdered, func(i, j int) bool {
			return stageSet.Stage.Style_text_stagedOrder[style_textOrdered[i]] < stageSet.Stage.Style_text_stagedOrder[style_textOrdered[j]]
		})
		for _, style_text := range style_textOrdered {
			style_textIdent := "__stage_0" + style_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Style_text{Name: %s}).Stage(stageSet.Stage)", style_textIdent, __gong__toRawStringLiteral(style_text.Name)))
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
		supportsOrdered := []*Supports{}
		for supports := range stageSet.Stage.Supportss {
			supportsOrdered = append(supportsOrdered, supports)
		}
		sort.Slice(supportsOrdered, func(i, j int) bool {
			return stageSet.Stage.Supports_stagedOrder[supportsOrdered[i]] < stageSet.Stage.Supports_stagedOrder[supportsOrdered[j]]
		})
		for _, supports := range supportsOrdered {
			supportsIdent := "__stage_0" + supports.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Supports{Name: %s}).Stage(stageSet.Stage)", supportsIdent, __gong__toRawStringLiteral(supports.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", supportsIdent, __gong__toRawStringLiteral(supports.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", supportsIdent, __gong__toRawStringLiteral(string(supports.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Element = %s", supportsIdent, __gong__toRawStringLiteral(supports.Element)))
			values.WriteString(fmt.Sprintf("\n\t%s.Attribute = %s", supportsIdent, __gong__toRawStringLiteral(supports.Attribute)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", supportsIdent, __gong__toRawStringLiteral(supports.Value)))
		}
	}
	if stageSet.Stage != nil {
		swingOrdered := []*Swing{}
		for swing := range stageSet.Stage.Swings {
			swingOrdered = append(swingOrdered, swing)
		}
		sort.Slice(swingOrdered, func(i, j int) bool {
			return stageSet.Stage.Swing_stagedOrder[swingOrdered[i]] < stageSet.Stage.Swing_stagedOrder[swingOrdered[j]]
		})
		for _, swing := range swingOrdered {
			swingIdent := "__stage_0" + swing.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Swing{Name: %s}).Stage(stageSet.Stage)", swingIdent, __gong__toRawStringLiteral(swing.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", swingIdent, __gong__toRawStringLiteral(swing.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Straight = %s", swingIdent, __gong__toRawStringLiteral(swing.Straight)))
			values.WriteString(fmt.Sprintf("\n\t%s.First = %d", swingIdent, swing.First))
			values.WriteString(fmt.Sprintf("\n\t%s.Second = %d", swingIdent, swing.Second))
			values.WriteString(fmt.Sprintf("\n\t%s.Swing_type = %s", swingIdent, __gong__toRawStringLiteral(string(swing.Swing_type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Swing_style = %s", swingIdent, __gong__toRawStringLiteral(swing.Swing_style)))
		}
	}
	if stageSet.Stage != nil {
		syncOrdered := []*Sync{}
		for sync := range stageSet.Stage.Syncs {
			syncOrdered = append(syncOrdered, sync)
		}
		sort.Slice(syncOrdered, func(i, j int) bool {
			return stageSet.Stage.Sync_stagedOrder[syncOrdered[i]] < stageSet.Stage.Sync_stagedOrder[syncOrdered[j]]
		})
		for _, sync := range syncOrdered {
			syncIdent := "__stage_0" + sync.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Sync{Name: %s}).Stage(stageSet.Stage)", syncIdent, __gong__toRawStringLiteral(sync.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", syncIdent, __gong__toRawStringLiteral(sync.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", syncIdent, __gong__toRawStringLiteral(sync.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Latency = %d", syncIdent, sync.Latency))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", syncIdent, __gong__toRawStringLiteral(sync.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_only = %s", syncIdent, __gong__toRawStringLiteral(string(sync.Time_only))))
		}
	}
	if stageSet.Stage != nil {
		system_dividersOrdered := []*System_dividers{}
		for system_dividers := range stageSet.Stage.System_dividerss {
			system_dividersOrdered = append(system_dividersOrdered, system_dividers)
		}
		sort.Slice(system_dividersOrdered, func(i, j int) bool {
			return stageSet.Stage.System_dividers_stagedOrder[system_dividersOrdered[i]] < stageSet.Stage.System_dividers_stagedOrder[system_dividersOrdered[j]]
		})
		for _, system_dividers := range system_dividersOrdered {
			system_dividersIdent := "__stage_0" + system_dividers.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.System_dividers{Name: %s}).Stage(stageSet.Stage)", system_dividersIdent, __gong__toRawStringLiteral(system_dividers.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", system_dividersIdent, __gong__toRawStringLiteral(system_dividers.Name)))
			if system_dividers.Left_divider != nil {
				targetIdent := "__stage_0" + system_dividers.Left_divider.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Left_divider = %s", system_dividersIdent, targetIdent))
			}
			if system_dividers.Right_divider != nil {
				targetIdent := "__stage_0" + system_dividers.Right_divider.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Right_divider = %s", system_dividersIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		system_layoutOrdered := []*System_layout{}
		for system_layout := range stageSet.Stage.System_layouts {
			system_layoutOrdered = append(system_layoutOrdered, system_layout)
		}
		sort.Slice(system_layoutOrdered, func(i, j int) bool {
			return stageSet.Stage.System_layout_stagedOrder[system_layoutOrdered[i]] < stageSet.Stage.System_layout_stagedOrder[system_layoutOrdered[j]]
		})
		for _, system_layout := range system_layoutOrdered {
			system_layoutIdent := "__stage_0" + system_layout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.System_layout{Name: %s}).Stage(stageSet.Stage)", system_layoutIdent, __gong__toRawStringLiteral(system_layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", system_layoutIdent, __gong__toRawStringLiteral(system_layout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.System_distance = %s", system_layoutIdent, __gong__toRawStringLiteral(system_layout.System_distance)))
			values.WriteString(fmt.Sprintf("\n\t%s.Top_system_distance = %s", system_layoutIdent, __gong__toRawStringLiteral(system_layout.Top_system_distance)))
			if system_layout.System_margins != nil {
				targetIdent := "__stage_0" + system_layout.System_margins.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System_margins = %s", system_layoutIdent, targetIdent))
			}
			if system_layout.System_dividers != nil {
				targetIdent := "__stage_0" + system_layout.System_dividers.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.System_dividers = %s", system_layoutIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		system_marginsOrdered := []*System_margins{}
		for system_margins := range stageSet.Stage.System_marginss {
			system_marginsOrdered = append(system_marginsOrdered, system_margins)
		}
		sort.Slice(system_marginsOrdered, func(i, j int) bool {
			return stageSet.Stage.System_margins_stagedOrder[system_marginsOrdered[i]] < stageSet.Stage.System_margins_stagedOrder[system_marginsOrdered[j]]
		})
		for _, system_margins := range system_marginsOrdered {
			system_marginsIdent := "__stage_0" + system_margins.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.System_margins{Name: %s}).Stage(stageSet.Stage)", system_marginsIdent, __gong__toRawStringLiteral(system_margins.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", system_marginsIdent, __gong__toRawStringLiteral(system_margins.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Left_margin = %s", system_marginsIdent, __gong__toRawStringLiteral(system_margins.Left_margin)))
			values.WriteString(fmt.Sprintf("\n\t%s.Right_margin = %s", system_marginsIdent, __gong__toRawStringLiteral(system_margins.Right_margin)))
		}
	}
	if stageSet.Stage != nil {
		tapOrdered := []*Tap{}
		for tap := range stageSet.Stage.Taps {
			tapOrdered = append(tapOrdered, tap)
		}
		sort.Slice(tapOrdered, func(i, j int) bool {
			return stageSet.Stage.Tap_stagedOrder[tapOrdered[i]] < stageSet.Stage.Tap_stagedOrder[tapOrdered[j]]
		})
		for _, tap := range tapOrdered {
			tapIdent := "__stage_0" + tap.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Tap{Name: %s}).Stage(stageSet.Stage)", tapIdent, __gong__toRawStringLiteral(tap.Name)))
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
		technicalOrdered := []*Technical{}
		for technical := range stageSet.Stage.Technicals {
			technicalOrdered = append(technicalOrdered, technical)
		}
		sort.Slice(technicalOrdered, func(i, j int) bool {
			return stageSet.Stage.Technical_stagedOrder[technicalOrdered[i]] < stageSet.Stage.Technical_stagedOrder[technicalOrdered[j]]
		})
		for _, technical := range technicalOrdered {
			technicalIdent := "__stage_0" + technical.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Technical{Name: %s}).Stage(stageSet.Stage)", technicalIdent, __gong__toRawStringLiteral(technical.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", technicalIdent, __gong__toRawStringLiteral(technical.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", technicalIdent, __gong__toRawStringLiteral(technical.Id)))
			for _, elem := range technical.Up_bow {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Up_bow = append(%s.Up_bow, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Down_bow {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Down_bow = append(%s.Down_bow, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Harmonic {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmonic = append(%s.Harmonic, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Open_string {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Open_string = append(%s.Open_string, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Thumb_position {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Thumb_position = append(%s.Thumb_position, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Fingering {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fingering = append(%s.Fingering, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Pluck {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pluck = append(%s.Pluck, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Double_tongue {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Double_tongue = append(%s.Double_tongue, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Triple_tongue {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Triple_tongue = append(%s.Triple_tongue, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Stopped {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stopped = append(%s.Stopped, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Snap_pizzicato {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Snap_pizzicato = append(%s.Snap_pizzicato, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Fret {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fret = append(%s.Fret, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.String {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.String = append(%s.String, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Hammer_on {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Hammer_on = append(%s.Hammer_on, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Pull_off {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pull_off = append(%s.Pull_off, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Bend {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bend = append(%s.Bend, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Tap {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tap = append(%s.Tap, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Heel {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Heel = append(%s.Heel, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Toe {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Toe = append(%s.Toe, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Fingernails {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fingernails = append(%s.Fingernails, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Hole {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Hole = append(%s.Hole, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Arrow {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arrow = append(%s.Arrow, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Handbell {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Handbell = append(%s.Handbell, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Brass_bend {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Brass_bend = append(%s.Brass_bend, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Flip {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Flip = append(%s.Flip, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Smear {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Smear = append(%s.Smear, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Open {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Open = append(%s.Open, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Half_muted {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Half_muted = append(%s.Half_muted, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Harmon_mute {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmon_mute = append(%s.Harmon_mute, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Golpe {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Golpe = append(%s.Golpe, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Other_technical {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Other_technical = append(%s.Other_technical, %s)", technicalIdent, technicalIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		text_element_dataOrdered := []*Text_element_data{}
		for text_element_data := range stageSet.Stage.Text_element_datas {
			text_element_dataOrdered = append(text_element_dataOrdered, text_element_data)
		}
		sort.Slice(text_element_dataOrdered, func(i, j int) bool {
			return stageSet.Stage.Text_element_data_stagedOrder[text_element_dataOrdered[i]] < stageSet.Stage.Text_element_data_stagedOrder[text_element_dataOrdered[j]]
		})
		for _, text_element_data := range text_element_dataOrdered {
			text_element_dataIdent := "__stage_0" + text_element_data.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Text_element_data{Name: %s}).Stage(stageSet.Stage)", text_element_dataIdent, __gong__toRawStringLiteral(text_element_data.Name)))
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
		tieOrdered := []*Tie{}
		for tie := range stageSet.Stage.Ties {
			tieOrdered = append(tieOrdered, tie)
		}
		sort.Slice(tieOrdered, func(i, j int) bool {
			return stageSet.Stage.Tie_stagedOrder[tieOrdered[i]] < stageSet.Stage.Tie_stagedOrder[tieOrdered[j]]
		})
		for _, tie := range tieOrdered {
			tieIdent := "__stage_0" + tie.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Tie{Name: %s}).Stage(stageSet.Stage)", tieIdent, __gong__toRawStringLiteral(tie.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tieIdent, __gong__toRawStringLiteral(tie.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", tieIdent, __gong__toRawStringLiteral(string(tie.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_only = %s", tieIdent, __gong__toRawStringLiteral(string(tie.Time_only))))
		}
	}
	if stageSet.Stage != nil {
		tiedOrdered := []*Tied{}
		for tied := range stageSet.Stage.Tieds {
			tiedOrdered = append(tiedOrdered, tied)
		}
		sort.Slice(tiedOrdered, func(i, j int) bool {
			return stageSet.Stage.Tied_stagedOrder[tiedOrdered[i]] < stageSet.Stage.Tied_stagedOrder[tiedOrdered[j]]
		})
		for _, tied := range tiedOrdered {
			tiedIdent := "__stage_0" + tied.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Tied{Name: %s}).Stage(stageSet.Stage)", tiedIdent, __gong__toRawStringLiteral(tied.Name)))
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
		timeOrdered := []*Time{}
		for time := range stageSet.Stage.Times {
			timeOrdered = append(timeOrdered, time)
		}
		sort.Slice(timeOrdered, func(i, j int) bool {
			return stageSet.Stage.Time_stagedOrder[timeOrdered[i]] < stageSet.Stage.Time_stagedOrder[timeOrdered[j]]
		})
		for _, time := range timeOrdered {
			timeIdent := "__stage_0" + time.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Time{Name: %s}).Stage(stageSet.Stage)", timeIdent, __gong__toRawStringLiteral(time.Name)))
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
				targetIdent := "__stage_0" + time.Interchangeable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Interchangeable = %s", timeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		time_modificationOrdered := []*Time_modification{}
		for time_modification := range stageSet.Stage.Time_modifications {
			time_modificationOrdered = append(time_modificationOrdered, time_modification)
		}
		sort.Slice(time_modificationOrdered, func(i, j int) bool {
			return stageSet.Stage.Time_modification_stagedOrder[time_modificationOrdered[i]] < stageSet.Stage.Time_modification_stagedOrder[time_modificationOrdered[j]]
		})
		for _, time_modification := range time_modificationOrdered {
			time_modificationIdent := "__stage_0" + time_modification.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Time_modification{Name: %s}).Stage(stageSet.Stage)", time_modificationIdent, __gong__toRawStringLiteral(time_modification.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", time_modificationIdent, __gong__toRawStringLiteral(time_modification.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actual_notes = %d", time_modificationIdent, time_modification.Actual_notes))
			values.WriteString(fmt.Sprintf("\n\t%s.Normal_notes = %d", time_modificationIdent, time_modification.Normal_notes))
			values.WriteString(fmt.Sprintf("\n\t%s.Normal_type = %s", time_modificationIdent, __gong__toRawStringLiteral(string(time_modification.Normal_type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Normal_dot = %s", time_modificationIdent, __gong__toRawStringLiteral(time_modification.Normal_dot)))
		}
	}
	if stageSet.Stage != nil {
		timpaniOrdered := []*Timpani{}
		for timpani := range stageSet.Stage.Timpanis {
			timpaniOrdered = append(timpaniOrdered, timpani)
		}
		sort.Slice(timpaniOrdered, func(i, j int) bool {
			return stageSet.Stage.Timpani_stagedOrder[timpaniOrdered[i]] < stageSet.Stage.Timpani_stagedOrder[timpaniOrdered[j]]
		})
		for _, timpani := range timpaniOrdered {
			timpaniIdent := "__stage_0" + timpani.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Timpani{Name: %s}).Stage(stageSet.Stage)", timpaniIdent, __gong__toRawStringLiteral(timpani.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timpaniIdent, __gong__toRawStringLiteral(timpani.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", timpaniIdent, __gong__toRawStringLiteral(timpani.Smufl)))
		}
	}
	if stageSet.Stage != nil {
		transposeOrdered := []*Transpose{}
		for transpose := range stageSet.Stage.Transposes {
			transposeOrdered = append(transposeOrdered, transpose)
		}
		sort.Slice(transposeOrdered, func(i, j int) bool {
			return stageSet.Stage.Transpose_stagedOrder[transposeOrdered[i]] < stageSet.Stage.Transpose_stagedOrder[transposeOrdered[j]]
		})
		for _, transpose := range transposeOrdered {
			transposeIdent := "__stage_0" + transpose.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Transpose{Name: %s}).Stage(stageSet.Stage)", transposeIdent, __gong__toRawStringLiteral(transpose.Name)))
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
		tremoloOrdered := []*Tremolo{}
		for tremolo := range stageSet.Stage.Tremolos {
			tremoloOrdered = append(tremoloOrdered, tremolo)
		}
		sort.Slice(tremoloOrdered, func(i, j int) bool {
			return stageSet.Stage.Tremolo_stagedOrder[tremoloOrdered[i]] < stageSet.Stage.Tremolo_stagedOrder[tremoloOrdered[j]]
		})
		for _, tremolo := range tremoloOrdered {
			tremoloIdent := "__stage_0" + tremolo.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Tremolo{Name: %s}).Stage(stageSet.Stage)", tremoloIdent, __gong__toRawStringLiteral(tremolo.Name)))
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
		tupletOrdered := []*Tuplet{}
		for tuplet := range stageSet.Stage.Tuplets {
			tupletOrdered = append(tupletOrdered, tuplet)
		}
		sort.Slice(tupletOrdered, func(i, j int) bool {
			return stageSet.Stage.Tuplet_stagedOrder[tupletOrdered[i]] < stageSet.Stage.Tuplet_stagedOrder[tupletOrdered[j]]
		})
		for _, tuplet := range tupletOrdered {
			tupletIdent := "__stage_0" + tuplet.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Tuplet{Name: %s}).Stage(stageSet.Stage)", tupletIdent, __gong__toRawStringLiteral(tuplet.Name)))
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
				targetIdent := "__stage_0" + tuplet.Tuplet_actual.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet_actual = %s", tupletIdent, targetIdent))
			}
			if tuplet.Tuplet_normal != nil {
				targetIdent := "__stage_0" + tuplet.Tuplet_normal.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet_normal = %s", tupletIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		tuplet_dotOrdered := []*Tuplet_dot{}
		for tuplet_dot := range stageSet.Stage.Tuplet_dots {
			tuplet_dotOrdered = append(tuplet_dotOrdered, tuplet_dot)
		}
		sort.Slice(tuplet_dotOrdered, func(i, j int) bool {
			return stageSet.Stage.Tuplet_dot_stagedOrder[tuplet_dotOrdered[i]] < stageSet.Stage.Tuplet_dot_stagedOrder[tuplet_dotOrdered[j]]
		})
		for _, tuplet_dot := range tuplet_dotOrdered {
			tuplet_dotIdent := "__stage_0" + tuplet_dot.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Tuplet_dot{Name: %s}).Stage(stageSet.Stage)", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_family = %s", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Font_family)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_style = %s", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Font_style)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_size = %s", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Font_size)))
			values.WriteString(fmt.Sprintf("\n\t%s.Font_weight = %s", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Font_weight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tuplet_dotIdent, __gong__toRawStringLiteral(tuplet_dot.Color)))
		}
	}
	if stageSet.Stage != nil {
		tuplet_numberOrdered := []*Tuplet_number{}
		for tuplet_number := range stageSet.Stage.Tuplet_numbers {
			tuplet_numberOrdered = append(tuplet_numberOrdered, tuplet_number)
		}
		sort.Slice(tuplet_numberOrdered, func(i, j int) bool {
			return stageSet.Stage.Tuplet_number_stagedOrder[tuplet_numberOrdered[i]] < stageSet.Stage.Tuplet_number_stagedOrder[tuplet_numberOrdered[j]]
		})
		for _, tuplet_number := range tuplet_numberOrdered {
			tuplet_numberIdent := "__stage_0" + tuplet_number.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Tuplet_number{Name: %s}).Stage(stageSet.Stage)", tuplet_numberIdent, __gong__toRawStringLiteral(tuplet_number.Name)))
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
		tuplet_portionOrdered := []*Tuplet_portion{}
		for tuplet_portion := range stageSet.Stage.Tuplet_portions {
			tuplet_portionOrdered = append(tuplet_portionOrdered, tuplet_portion)
		}
		sort.Slice(tuplet_portionOrdered, func(i, j int) bool {
			return stageSet.Stage.Tuplet_portion_stagedOrder[tuplet_portionOrdered[i]] < stageSet.Stage.Tuplet_portion_stagedOrder[tuplet_portionOrdered[j]]
		})
		for _, tuplet_portion := range tuplet_portionOrdered {
			tuplet_portionIdent := "__stage_0" + tuplet_portion.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Tuplet_portion{Name: %s}).Stage(stageSet.Stage)", tuplet_portionIdent, __gong__toRawStringLiteral(tuplet_portion.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tuplet_portionIdent, __gong__toRawStringLiteral(tuplet_portion.Name)))
			if tuplet_portion.Tuplet_number != nil {
				targetIdent := "__stage_0" + tuplet_portion.Tuplet_number.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet_number = %s", tuplet_portionIdent, targetIdent))
			}
			if tuplet_portion.Tuplet_type != nil {
				targetIdent := "__stage_0" + tuplet_portion.Tuplet_type.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet_type = %s", tuplet_portionIdent, targetIdent))
			}
			for _, elem := range tuplet_portion.Tuplet_dot {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet_dot = append(%s.Tuplet_dot, %s)", tuplet_portionIdent, tuplet_portionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		tuplet_typeOrdered := []*Tuplet_type{}
		for tuplet_type := range stageSet.Stage.Tuplet_types {
			tuplet_typeOrdered = append(tuplet_typeOrdered, tuplet_type)
		}
		sort.Slice(tuplet_typeOrdered, func(i, j int) bool {
			return stageSet.Stage.Tuplet_type_stagedOrder[tuplet_typeOrdered[i]] < stageSet.Stage.Tuplet_type_stagedOrder[tuplet_typeOrdered[j]]
		})
		for _, tuplet_type := range tuplet_typeOrdered {
			tuplet_typeIdent := "__stage_0" + tuplet_type.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Tuplet_type{Name: %s}).Stage(stageSet.Stage)", tuplet_typeIdent, __gong__toRawStringLiteral(tuplet_type.Name)))
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
		typed_textOrdered := []*Typed_text{}
		for typed_text := range stageSet.Stage.Typed_texts {
			typed_textOrdered = append(typed_textOrdered, typed_text)
		}
		sort.Slice(typed_textOrdered, func(i, j int) bool {
			return stageSet.Stage.Typed_text_stagedOrder[typed_textOrdered[i]] < stageSet.Stage.Typed_text_stagedOrder[typed_textOrdered[j]]
		})
		for _, typed_text := range typed_textOrdered {
			typed_textIdent := "__stage_0" + typed_text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Typed_text{Name: %s}).Stage(stageSet.Stage)", typed_textIdent, __gong__toRawStringLiteral(typed_text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", typed_textIdent, __gong__toRawStringLiteral(typed_text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", typed_textIdent, __gong__toRawStringLiteral(typed_text.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", typed_textIdent, __gong__toRawStringLiteral(typed_text.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		unpitchedOrdered := []*Unpitched{}
		for unpitched := range stageSet.Stage.Unpitcheds {
			unpitchedOrdered = append(unpitchedOrdered, unpitched)
		}
		sort.Slice(unpitchedOrdered, func(i, j int) bool {
			return stageSet.Stage.Unpitched_stagedOrder[unpitchedOrdered[i]] < stageSet.Stage.Unpitched_stagedOrder[unpitchedOrdered[j]]
		})
		for _, unpitched := range unpitchedOrdered {
			unpitchedIdent := "__stage_0" + unpitched.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Unpitched{Name: %s}).Stage(stageSet.Stage)", unpitchedIdent, __gong__toRawStringLiteral(unpitched.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unpitchedIdent, __gong__toRawStringLiteral(unpitched.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Display_step = %s", unpitchedIdent, __gong__toRawStringLiteral(string(unpitched.Display_step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Display_octave = %d", unpitchedIdent, unpitched.Display_octave))
		}
	}
	if stageSet.Stage != nil {
		virtual_instrumentOrdered := []*Virtual_instrument{}
		for virtual_instrument := range stageSet.Stage.Virtual_instruments {
			virtual_instrumentOrdered = append(virtual_instrumentOrdered, virtual_instrument)
		}
		sort.Slice(virtual_instrumentOrdered, func(i, j int) bool {
			return stageSet.Stage.Virtual_instrument_stagedOrder[virtual_instrumentOrdered[i]] < stageSet.Stage.Virtual_instrument_stagedOrder[virtual_instrumentOrdered[j]]
		})
		for _, virtual_instrument := range virtual_instrumentOrdered {
			virtual_instrumentIdent := "__stage_0" + virtual_instrument.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Virtual_instrument{Name: %s}).Stage(stageSet.Stage)", virtual_instrumentIdent, __gong__toRawStringLiteral(virtual_instrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", virtual_instrumentIdent, __gong__toRawStringLiteral(virtual_instrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Virtual_library = %s", virtual_instrumentIdent, __gong__toRawStringLiteral(virtual_instrument.Virtual_library)))
			values.WriteString(fmt.Sprintf("\n\t%s.Virtual_name = %s", virtual_instrumentIdent, __gong__toRawStringLiteral(virtual_instrument.Virtual_name)))
		}
	}
	if stageSet.Stage != nil {
		waitOrdered := []*Wait{}
		for wait := range stageSet.Stage.Waits {
			waitOrdered = append(waitOrdered, wait)
		}
		sort.Slice(waitOrdered, func(i, j int) bool {
			return stageSet.Stage.Wait_stagedOrder[waitOrdered[i]] < stageSet.Stage.Wait_stagedOrder[waitOrdered[j]]
		})
		for _, wait := range waitOrdered {
			waitIdent := "__stage_0" + wait.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Wait{Name: %s}).Stage(stageSet.Stage)", waitIdent, __gong__toRawStringLiteral(wait.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", waitIdent, __gong__toRawStringLiteral(wait.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", waitIdent, __gong__toRawStringLiteral(wait.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.Time_only = %s", waitIdent, __gong__toRawStringLiteral(string(wait.Time_only))))
		}
	}
	if stageSet.Stage != nil {
		wavy_lineOrdered := []*Wavy_line{}
		for wavy_line := range stageSet.Stage.Wavy_lines {
			wavy_lineOrdered = append(wavy_lineOrdered, wavy_line)
		}
		sort.Slice(wavy_lineOrdered, func(i, j int) bool {
			return stageSet.Stage.Wavy_line_stagedOrder[wavy_lineOrdered[i]] < stageSet.Stage.Wavy_line_stagedOrder[wavy_lineOrdered[j]]
		})
		for _, wavy_line := range wavy_lineOrdered {
			wavy_lineIdent := "__stage_0" + wavy_line.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Wavy_line{Name: %s}).Stage(stageSet.Stage)", wavy_lineIdent, __gong__toRawStringLiteral(wavy_line.Name)))
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
		wedgeOrdered := []*Wedge{}
		for wedge := range stageSet.Stage.Wedges {
			wedgeOrdered = append(wedgeOrdered, wedge)
		}
		sort.Slice(wedgeOrdered, func(i, j int) bool {
			return stageSet.Stage.Wedge_stagedOrder[wedgeOrdered[i]] < stageSet.Stage.Wedge_stagedOrder[wedgeOrdered[j]]
		})
		for _, wedge := range wedgeOrdered {
			wedgeIdent := "__stage_0" + wedge.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Wedge{Name: %s}).Stage(stageSet.Stage)", wedgeIdent, __gong__toRawStringLiteral(wedge.Name)))
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
		woodOrdered := []*Wood{}
		for wood := range stageSet.Stage.Woods {
			woodOrdered = append(woodOrdered, wood)
		}
		sort.Slice(woodOrdered, func(i, j int) bool {
			return stageSet.Stage.Wood_stagedOrder[woodOrdered[i]] < stageSet.Stage.Wood_stagedOrder[woodOrdered[j]]
		})
		for _, wood := range woodOrdered {
			woodIdent := "__stage_0" + wood.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Wood{Name: %s}).Stage(stageSet.Stage)", woodIdent, __gong__toRawStringLiteral(wood.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", woodIdent, __gong__toRawStringLiteral(wood.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", woodIdent, __gong__toRawStringLiteral(wood.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", woodIdent, __gong__toRawStringLiteral(wood.EnclosedText)))
		}
	}
	if stageSet.Stage != nil {
		workOrdered := []*Work{}
		for work := range stageSet.Stage.Works {
			workOrdered = append(workOrdered, work)
		}
		sort.Slice(workOrdered, func(i, j int) bool {
			return stageSet.Stage.Work_stagedOrder[workOrdered[i]] < stageSet.Stage.Work_stagedOrder[workOrdered[j]]
		})
		for _, work := range workOrdered {
			workIdent := "__stage_0" + work.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Work{Name: %s}).Stage(stageSet.Stage)", workIdent, __gong__toRawStringLiteral(work.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", workIdent, __gong__toRawStringLiteral(work.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Work_number = %s", workIdent, __gong__toRawStringLiteral(work.Work_number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Work_title = %s", workIdent, __gong__toRawStringLiteral(work.Work_title)))
			if work.Opus != nil {
				targetIdent := "__stage_0" + work.Opus.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Opus = %s", workIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models"
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
				case "A_directive":
					if !preserveOrder {
						inst := (&A_directive{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_directive)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_measure":
					if !preserveOrder {
						inst := (&A_measure{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_measure)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_measure_1":
					if !preserveOrder {
						inst := (&A_measure_1{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_measure_1)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_part":
					if !preserveOrder {
						inst := (&A_part{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_part)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "A_part_1":
					if !preserveOrder {
						inst := (&A_part_1{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(A_part_1)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Accidental":
					if !preserveOrder {
						inst := (&Accidental{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Accidental)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Accidental_mark":
					if !preserveOrder {
						inst := (&Accidental_mark{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Accidental_mark)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Accidental_text":
					if !preserveOrder {
						inst := (&Accidental_text{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Accidental_text)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Accord":
					if !preserveOrder {
						inst := (&Accord{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Accord)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Accordion_registration":
					if !preserveOrder {
						inst := (&Accordion_registration{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Accordion_registration)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Appearance":
					if !preserveOrder {
						inst := (&Appearance{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Appearance)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Arpeggiate":
					if !preserveOrder {
						inst := (&Arpeggiate{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Arpeggiate)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Arrow":
					if !preserveOrder {
						inst := (&Arrow{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Arrow)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Articulations":
					if !preserveOrder {
						inst := (&Articulations{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Articulations)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Assess":
					if !preserveOrder {
						inst := (&Assess{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Assess)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Attributes":
					if !preserveOrder {
						inst := (&Attributes{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Attributes)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Backup":
					if !preserveOrder {
						inst := (&Backup{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Backup)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Bar_style_color":
					if !preserveOrder {
						inst := (&Bar_style_color{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Bar_style_color)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Barline":
					if !preserveOrder {
						inst := (&Barline{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Barline)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Barre":
					if !preserveOrder {
						inst := (&Barre{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Barre)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Bass":
					if !preserveOrder {
						inst := (&Bass{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Bass)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Bass_step":
					if !preserveOrder {
						inst := (&Bass_step{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Bass_step)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Beam":
					if !preserveOrder {
						inst := (&Beam{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Beam)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Beat_repeat":
					if !preserveOrder {
						inst := (&Beat_repeat{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Beat_repeat)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Beat_unit_tied":
					if !preserveOrder {
						inst := (&Beat_unit_tied{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Beat_unit_tied)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Beater":
					if !preserveOrder {
						inst := (&Beater{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Beater)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Bend":
					if !preserveOrder {
						inst := (&Bend{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Bend)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Bookmark":
					if !preserveOrder {
						inst := (&Bookmark{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Bookmark)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Bracket":
					if !preserveOrder {
						inst := (&Bracket{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Bracket)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Breath_mark":
					if !preserveOrder {
						inst := (&Breath_mark{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Breath_mark)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Caesura":
					if !preserveOrder {
						inst := (&Caesura{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Caesura)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Cancel":
					if !preserveOrder {
						inst := (&Cancel{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Cancel)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Clef":
					if !preserveOrder {
						inst := (&Clef{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Clef)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Coda":
					if !preserveOrder {
						inst := (&Coda{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Coda)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Credit":
					if !preserveOrder {
						inst := (&Credit{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Credit)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Dashes":
					if !preserveOrder {
						inst := (&Dashes{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Dashes)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Defaults":
					if !preserveOrder {
						inst := (&Defaults{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Defaults)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Degree":
					if !preserveOrder {
						inst := (&Degree{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Degree)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Degree_alter":
					if !preserveOrder {
						inst := (&Degree_alter{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Degree_alter)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Degree_type":
					if !preserveOrder {
						inst := (&Degree_type{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Degree_type)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Degree_value":
					if !preserveOrder {
						inst := (&Degree_value{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Degree_value)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Direction":
					if !preserveOrder {
						inst := (&Direction{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Direction)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Direction_type":
					if !preserveOrder {
						inst := (&Direction_type{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Direction_type)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Distance":
					if !preserveOrder {
						inst := (&Distance{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Distance)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Double":
					if !preserveOrder {
						inst := (&Double{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Double)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Dynamics":
					if !preserveOrder {
						inst := (&Dynamics{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Dynamics)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Effect":
					if !preserveOrder {
						inst := (&Effect{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Effect)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Elision":
					if !preserveOrder {
						inst := (&Elision{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Elision)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Empty":
					if !preserveOrder {
						inst := (&Empty{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Empty)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Empty_font":
					if !preserveOrder {
						inst := (&Empty_font{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Empty_font)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Empty_line":
					if !preserveOrder {
						inst := (&Empty_line{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Empty_line)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Empty_placement":
					if !preserveOrder {
						inst := (&Empty_placement{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Empty_placement)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Empty_placement_smufl":
					if !preserveOrder {
						inst := (&Empty_placement_smufl{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Empty_placement_smufl)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Empty_print_object_style_align":
					if !preserveOrder {
						inst := (&Empty_print_object_style_align{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Empty_print_object_style_align)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Empty_print_style":
					if !preserveOrder {
						inst := (&Empty_print_style{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Empty_print_style)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Empty_print_style_align":
					if !preserveOrder {
						inst := (&Empty_print_style_align{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Empty_print_style_align)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Empty_print_style_align_id":
					if !preserveOrder {
						inst := (&Empty_print_style_align_id{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Empty_print_style_align_id)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Empty_trill_sound":
					if !preserveOrder {
						inst := (&Empty_trill_sound{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Empty_trill_sound)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Encoding":
					if !preserveOrder {
						inst := (&Encoding{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Encoding)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Ending":
					if !preserveOrder {
						inst := (&Ending{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Ending)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Extend":
					if !preserveOrder {
						inst := (&Extend{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Extend)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Feature":
					if !preserveOrder {
						inst := (&Feature{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Feature)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Fermata":
					if !preserveOrder {
						inst := (&Fermata{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Fermata)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Figure":
					if !preserveOrder {
						inst := (&Figure{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Figure)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Figured_bass":
					if !preserveOrder {
						inst := (&Figured_bass{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Figured_bass)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Fingering":
					if !preserveOrder {
						inst := (&Fingering{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Fingering)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "First_fret":
					if !preserveOrder {
						inst := (&First_fret{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(First_fret)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "For_part":
					if !preserveOrder {
						inst := (&For_part{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(For_part)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Formatted_symbol":
					if !preserveOrder {
						inst := (&Formatted_symbol{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Formatted_symbol)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Formatted_symbol_id":
					if !preserveOrder {
						inst := (&Formatted_symbol_id{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Formatted_symbol_id)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Formatted_text":
					if !preserveOrder {
						inst := (&Formatted_text{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Formatted_text)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Formatted_text_id":
					if !preserveOrder {
						inst := (&Formatted_text_id{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Formatted_text_id)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Forward":
					if !preserveOrder {
						inst := (&Forward{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Forward)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Frame":
					if !preserveOrder {
						inst := (&Frame{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Frame)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Frame_note":
					if !preserveOrder {
						inst := (&Frame_note{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Frame_note)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Fret":
					if !preserveOrder {
						inst := (&Fret{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Fret)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Glass":
					if !preserveOrder {
						inst := (&Glass{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Glass)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Glissando":
					if !preserveOrder {
						inst := (&Glissando{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Glissando)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Glyph":
					if !preserveOrder {
						inst := (&Glyph{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Glyph)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Grace":
					if !preserveOrder {
						inst := (&Grace{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Grace)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Group_barline":
					if !preserveOrder {
						inst := (&Group_barline{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Group_barline)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Group_name":
					if !preserveOrder {
						inst := (&Group_name{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Group_name)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Group_symbol":
					if !preserveOrder {
						inst := (&Group_symbol{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Group_symbol)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Grouping":
					if !preserveOrder {
						inst := (&Grouping{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Grouping)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Hammer_on_pull_off":
					if !preserveOrder {
						inst := (&Hammer_on_pull_off{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Hammer_on_pull_off)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Handbell":
					if !preserveOrder {
						inst := (&Handbell{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Handbell)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Harmon_closed":
					if !preserveOrder {
						inst := (&Harmon_closed{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Harmon_closed)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Harmon_mute":
					if !preserveOrder {
						inst := (&Harmon_mute{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Harmon_mute)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Harmonic":
					if !preserveOrder {
						inst := (&Harmonic{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Harmonic)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Harmony":
					if !preserveOrder {
						inst := (&Harmony{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Harmony)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Harmony_alter":
					if !preserveOrder {
						inst := (&Harmony_alter{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Harmony_alter)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Harp_pedals":
					if !preserveOrder {
						inst := (&Harp_pedals{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Harp_pedals)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Heel_toe":
					if !preserveOrder {
						inst := (&Heel_toe{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Heel_toe)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Hole":
					if !preserveOrder {
						inst := (&Hole{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Hole)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Hole_closed":
					if !preserveOrder {
						inst := (&Hole_closed{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Hole_closed)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Horizontal_turn":
					if !preserveOrder {
						inst := (&Horizontal_turn{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Horizontal_turn)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Identification":
					if !preserveOrder {
						inst := (&Identification{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Identification)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Image":
					if !preserveOrder {
						inst := (&Image{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Image)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Instrument":
					if !preserveOrder {
						inst := (&Instrument{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Instrument)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Instrument_change":
					if !preserveOrder {
						inst := (&Instrument_change{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Instrument_change)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Instrument_link":
					if !preserveOrder {
						inst := (&Instrument_link{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Instrument_link)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Interchangeable":
					if !preserveOrder {
						inst := (&Interchangeable{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Interchangeable)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Inversion":
					if !preserveOrder {
						inst := (&Inversion{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Inversion)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Key":
					if !preserveOrder {
						inst := (&Key{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Key)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Key_accidental":
					if !preserveOrder {
						inst := (&Key_accidental{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Key_accidental)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Key_octave":
					if !preserveOrder {
						inst := (&Key_octave{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Key_octave)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Kind":
					if !preserveOrder {
						inst := (&Kind{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Kind)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Level":
					if !preserveOrder {
						inst := (&Level{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Level)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Line_detail":
					if !preserveOrder {
						inst := (&Line_detail{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Line_detail)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Line_width":
					if !preserveOrder {
						inst := (&Line_width{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Line_width)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Link":
					if !preserveOrder {
						inst := (&Link{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Link)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Listen":
					if !preserveOrder {
						inst := (&Listen{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Listen)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Listening":
					if !preserveOrder {
						inst := (&Listening{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Listening)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Lyric":
					if !preserveOrder {
						inst := (&Lyric{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Lyric)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Lyric_font":
					if !preserveOrder {
						inst := (&Lyric_font{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Lyric_font)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Lyric_language":
					if !preserveOrder {
						inst := (&Lyric_language{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Lyric_language)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Measure_layout":
					if !preserveOrder {
						inst := (&Measure_layout{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Measure_layout)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Measure_numbering":
					if !preserveOrder {
						inst := (&Measure_numbering{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Measure_numbering)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Measure_repeat":
					if !preserveOrder {
						inst := (&Measure_repeat{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Measure_repeat)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Measure_style":
					if !preserveOrder {
						inst := (&Measure_style{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Measure_style)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Membrane":
					if !preserveOrder {
						inst := (&Membrane{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Membrane)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Metal":
					if !preserveOrder {
						inst := (&Metal{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Metal)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Metronome":
					if !preserveOrder {
						inst := (&Metronome{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Metronome)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Metronome_beam":
					if !preserveOrder {
						inst := (&Metronome_beam{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Metronome_beam)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Metronome_note":
					if !preserveOrder {
						inst := (&Metronome_note{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Metronome_note)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Metronome_tied":
					if !preserveOrder {
						inst := (&Metronome_tied{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Metronome_tied)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Metronome_tuplet":
					if !preserveOrder {
						inst := (&Metronome_tuplet{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Metronome_tuplet)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Midi_device":
					if !preserveOrder {
						inst := (&Midi_device{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Midi_device)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Midi_instrument":
					if !preserveOrder {
						inst := (&Midi_instrument{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Midi_instrument)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Miscellaneous":
					if !preserveOrder {
						inst := (&Miscellaneous{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Miscellaneous)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Miscellaneous_field":
					if !preserveOrder {
						inst := (&Miscellaneous_field{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Miscellaneous_field)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Mordent":
					if !preserveOrder {
						inst := (&Mordent{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Mordent)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Multiple_rest":
					if !preserveOrder {
						inst := (&Multiple_rest{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Multiple_rest)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Name_display":
					if !preserveOrder {
						inst := (&Name_display{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Name_display)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Non_arpeggiate":
					if !preserveOrder {
						inst := (&Non_arpeggiate{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Non_arpeggiate)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Notations":
					if !preserveOrder {
						inst := (&Notations{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Notations)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Note":
					if !preserveOrder {
						inst := (&Note{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Note)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Note_size":
					if !preserveOrder {
						inst := (&Note_size{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Note_size)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Note_type":
					if !preserveOrder {
						inst := (&Note_type{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Note_type)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Notehead":
					if !preserveOrder {
						inst := (&Notehead{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Notehead)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Notehead_text":
					if !preserveOrder {
						inst := (&Notehead_text{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Notehead_text)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Numeral":
					if !preserveOrder {
						inst := (&Numeral{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Numeral)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Numeral_key":
					if !preserveOrder {
						inst := (&Numeral_key{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Numeral_key)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Numeral_root":
					if !preserveOrder {
						inst := (&Numeral_root{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Numeral_root)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Octave_shift":
					if !preserveOrder {
						inst := (&Octave_shift{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Octave_shift)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Offset":
					if !preserveOrder {
						inst := (&Offset{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Offset)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Opus":
					if !preserveOrder {
						inst := (&Opus{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Opus)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Ornaments":
					if !preserveOrder {
						inst := (&Ornaments{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Ornaments)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Other_appearance":
					if !preserveOrder {
						inst := (&Other_appearance{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Other_appearance)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Other_direction":
					if !preserveOrder {
						inst := (&Other_direction{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Other_direction)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Other_listening":
					if !preserveOrder {
						inst := (&Other_listening{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Other_listening)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Other_notation":
					if !preserveOrder {
						inst := (&Other_notation{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Other_notation)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Other_placement_text":
					if !preserveOrder {
						inst := (&Other_placement_text{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Other_placement_text)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Other_play":
					if !preserveOrder {
						inst := (&Other_play{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Other_play)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Other_text":
					if !preserveOrder {
						inst := (&Other_text{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Other_text)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Page_layout":
					if !preserveOrder {
						inst := (&Page_layout{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Page_layout)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Page_margins":
					if !preserveOrder {
						inst := (&Page_margins{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Page_margins)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Part_clef":
					if !preserveOrder {
						inst := (&Part_clef{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Part_clef)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Part_group":
					if !preserveOrder {
						inst := (&Part_group{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Part_group)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Part_link":
					if !preserveOrder {
						inst := (&Part_link{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Part_link)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Part_list":
					if !preserveOrder {
						inst := (&Part_list{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Part_list)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Part_name":
					if !preserveOrder {
						inst := (&Part_name{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Part_name)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Part_symbol":
					if !preserveOrder {
						inst := (&Part_symbol{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Part_symbol)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Part_transpose":
					if !preserveOrder {
						inst := (&Part_transpose{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Part_transpose)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Pedal":
					if !preserveOrder {
						inst := (&Pedal{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Pedal)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Pedal_tuning":
					if !preserveOrder {
						inst := (&Pedal_tuning{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Pedal_tuning)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Per_minute":
					if !preserveOrder {
						inst := (&Per_minute{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Per_minute)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Percussion":
					if !preserveOrder {
						inst := (&Percussion{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Percussion)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Pitch":
					if !preserveOrder {
						inst := (&Pitch{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Pitch)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Pitched":
					if !preserveOrder {
						inst := (&Pitched{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Pitched)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Placement_text":
					if !preserveOrder {
						inst := (&Placement_text{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Placement_text)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Play":
					if !preserveOrder {
						inst := (&Play{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Play)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Player":
					if !preserveOrder {
						inst := (&Player{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Player)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Principal_voice":
					if !preserveOrder {
						inst := (&Principal_voice{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Principal_voice)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Print":
					if !preserveOrder {
						inst := (&Print{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Print)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Release":
					if !preserveOrder {
						inst := (&Release{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Release)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Repeat":
					if !preserveOrder {
						inst := (&Repeat{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Repeat)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Rest":
					if !preserveOrder {
						inst := (&Rest{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Rest)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Root":
					if !preserveOrder {
						inst := (&Root{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Root)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Root_step":
					if !preserveOrder {
						inst := (&Root_step{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Root_step)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Scaling":
					if !preserveOrder {
						inst := (&Scaling{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Scaling)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Scordatura":
					if !preserveOrder {
						inst := (&Scordatura{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Scordatura)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Score_instrument":
					if !preserveOrder {
						inst := (&Score_instrument{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Score_instrument)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Score_part":
					if !preserveOrder {
						inst := (&Score_part{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Score_part)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Score_partwise":
					if !preserveOrder {
						inst := (&Score_partwise{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Score_partwise)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Score_timewise":
					if !preserveOrder {
						inst := (&Score_timewise{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Score_timewise)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Segno":
					if !preserveOrder {
						inst := (&Segno{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Segno)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Slash":
					if !preserveOrder {
						inst := (&Slash{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Slash)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Slide":
					if !preserveOrder {
						inst := (&Slide{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Slide)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Slur":
					if !preserveOrder {
						inst := (&Slur{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Slur)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Sound":
					if !preserveOrder {
						inst := (&Sound{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Sound)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Staff_details":
					if !preserveOrder {
						inst := (&Staff_details{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Staff_details)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Staff_divide":
					if !preserveOrder {
						inst := (&Staff_divide{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Staff_divide)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Staff_layout":
					if !preserveOrder {
						inst := (&Staff_layout{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Staff_layout)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Staff_size":
					if !preserveOrder {
						inst := (&Staff_size{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Staff_size)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Staff_tuning":
					if !preserveOrder {
						inst := (&Staff_tuning{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Staff_tuning)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Stem":
					if !preserveOrder {
						inst := (&Stem{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Stem)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Stick":
					if !preserveOrder {
						inst := (&Stick{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Stick)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "String_mute":
					if !preserveOrder {
						inst := (&String_mute{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(String_mute)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "String_type":
					if !preserveOrder {
						inst := (&String_type{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(String_type)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Strong_accent":
					if !preserveOrder {
						inst := (&Strong_accent{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Strong_accent)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Style_text":
					if !preserveOrder {
						inst := (&Style_text{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Style_text)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Supports":
					if !preserveOrder {
						inst := (&Supports{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Supports)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Swing":
					if !preserveOrder {
						inst := (&Swing{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Swing)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Sync":
					if !preserveOrder {
						inst := (&Sync{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Sync)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "System_dividers":
					if !preserveOrder {
						inst := (&System_dividers{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(System_dividers)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "System_layout":
					if !preserveOrder {
						inst := (&System_layout{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(System_layout)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "System_margins":
					if !preserveOrder {
						inst := (&System_margins{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(System_margins)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Tap":
					if !preserveOrder {
						inst := (&Tap{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Tap)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Technical":
					if !preserveOrder {
						inst := (&Technical{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Technical)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Text_element_data":
					if !preserveOrder {
						inst := (&Text_element_data{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Text_element_data)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Tie":
					if !preserveOrder {
						inst := (&Tie{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Tie)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Tied":
					if !preserveOrder {
						inst := (&Tied{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Tied)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Time":
					if !preserveOrder {
						inst := (&Time{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Time)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Time_modification":
					if !preserveOrder {
						inst := (&Time_modification{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Time_modification)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Timpani":
					if !preserveOrder {
						inst := (&Timpani{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Timpani)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Transpose":
					if !preserveOrder {
						inst := (&Transpose{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Transpose)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Tremolo":
					if !preserveOrder {
						inst := (&Tremolo{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Tremolo)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Tuplet":
					if !preserveOrder {
						inst := (&Tuplet{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Tuplet)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Tuplet_dot":
					if !preserveOrder {
						inst := (&Tuplet_dot{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Tuplet_dot)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Tuplet_number":
					if !preserveOrder {
						inst := (&Tuplet_number{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Tuplet_number)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Tuplet_portion":
					if !preserveOrder {
						inst := (&Tuplet_portion{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Tuplet_portion)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Tuplet_type":
					if !preserveOrder {
						inst := (&Tuplet_type{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Tuplet_type)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Typed_text":
					if !preserveOrder {
						inst := (&Typed_text{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Typed_text)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Unpitched":
					if !preserveOrder {
						inst := (&Unpitched{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Unpitched)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Virtual_instrument":
					if !preserveOrder {
						inst := (&Virtual_instrument{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Virtual_instrument)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Wait":
					if !preserveOrder {
						inst := (&Wait{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Wait)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Wavy_line":
					if !preserveOrder {
						inst := (&Wavy_line{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Wavy_line)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Wedge":
					if !preserveOrder {
						inst := (&Wedge{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Wedge)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Wood":
					if !preserveOrder {
						inst := (&Wood{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Wood)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Work":
					if !preserveOrder {
						inst := (&Work{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Work)
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note); ok {
										inst.Note = append(inst.Note, typedTarget)
									}
								}
							}
						}
					case "Backup":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Backup); ok {
										inst.Backup = append(inst.Backup, typedTarget)
									}
								}
							}
						}
					case "Forward":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Forward); ok {
										inst.Forward = append(inst.Forward, typedTarget)
									}
								}
							}
						}
					case "Direction":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Direction); ok {
										inst.Direction = append(inst.Direction, typedTarget)
									}
								}
							}
						}
					case "Attributes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Attributes); ok {
										inst.Attributes = append(inst.Attributes, typedTarget)
									}
								}
							}
						}
					case "Harmony":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Harmony); ok {
										inst.Harmony = append(inst.Harmony, typedTarget)
									}
								}
							}
						}
					case "Figured_bass":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Figured_bass); ok {
										inst.Figured_bass = append(inst.Figured_bass, typedTarget)
									}
								}
							}
						}
					case "Print":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Print); ok {
										inst.Print = append(inst.Print, typedTarget)
									}
								}
							}
						}
					case "Sound":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Sound); ok {
										inst.Sound = append(inst.Sound, typedTarget)
									}
								}
							}
						}
					case "Listening":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Listening); ok {
										inst.Listening = append(inst.Listening, typedTarget)
									}
								}
							}
						}
					case "Barline":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Barline); ok {
										inst.Barline = append(inst.Barline, typedTarget)
									}
								}
							}
						}
					case "Grouping":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Grouping); ok {
										inst.Grouping = append(inst.Grouping, typedTarget)
									}
								}
							}
						}
					case "Link":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Link); ok {
										inst.Link = append(inst.Link, typedTarget)
									}
								}
							}
						}
					case "Bookmark":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bookmark); ok {
										inst.Bookmark = append(inst.Bookmark, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_part_1); ok {
										inst.Part = append(inst.Part, typedTarget)
									}
								}
							}
						}
					}
				case *A_part:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Measure":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_measure); ok {
										inst.Measure = append(inst.Measure, typedTarget)
									}
								}
							}
						}
					}
				case *A_part_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Note":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note); ok {
										inst.Note = append(inst.Note, typedTarget)
									}
								}
							}
						}
					case "Backup":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Backup); ok {
										inst.Backup = append(inst.Backup, typedTarget)
									}
								}
							}
						}
					case "Forward":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Forward); ok {
										inst.Forward = append(inst.Forward, typedTarget)
									}
								}
							}
						}
					case "Direction":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Direction); ok {
										inst.Direction = append(inst.Direction, typedTarget)
									}
								}
							}
						}
					case "Attributes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Attributes); ok {
										inst.Attributes = append(inst.Attributes, typedTarget)
									}
								}
							}
						}
					case "Harmony":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Harmony); ok {
										inst.Harmony = append(inst.Harmony, typedTarget)
									}
								}
							}
						}
					case "Figured_bass":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Figured_bass); ok {
										inst.Figured_bass = append(inst.Figured_bass, typedTarget)
									}
								}
							}
						}
					case "Print":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Print); ok {
										inst.Print = append(inst.Print, typedTarget)
									}
								}
							}
						}
					case "Sound":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Sound); ok {
										inst.Sound = append(inst.Sound, typedTarget)
									}
								}
							}
						}
					case "Listening":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Listening); ok {
										inst.Listening = append(inst.Listening, typedTarget)
									}
								}
							}
						}
					case "Barline":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Barline); ok {
										inst.Barline = append(inst.Barline, typedTarget)
									}
								}
							}
						}
					case "Grouping":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Grouping); ok {
										inst.Grouping = append(inst.Grouping, typedTarget)
									}
								}
							}
						}
					case "Link":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Link); ok {
										inst.Link = append(inst.Link, typedTarget)
									}
								}
							}
						}
					case "Bookmark":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bookmark); ok {
										inst.Bookmark = append(inst.Bookmark, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Line_width); ok {
										inst.Line_width = append(inst.Line_width, typedTarget)
									}
								}
							}
						}
					case "Note_size":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note_size); ok {
										inst.Note_size = append(inst.Note_size, typedTarget)
									}
								}
							}
						}
					case "Distance":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Distance); ok {
										inst.Distance = append(inst.Distance, typedTarget)
									}
								}
							}
						}
					case "Glyph":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Glyph); ok {
										inst.Glyph = append(inst.Glyph, typedTarget)
									}
								}
							}
						}
					case "Other_appearance":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Other_appearance); ok {
										inst.Other_appearance = append(inst.Other_appearance, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Accent = append(inst.Accent, typedTarget)
									}
								}
							}
						}
					case "Strong_accent":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Strong_accent); ok {
										inst.Strong_accent = append(inst.Strong_accent, typedTarget)
									}
								}
							}
						}
					case "Staccato":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Staccato = append(inst.Staccato, typedTarget)
									}
								}
							}
						}
					case "Tenuto":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Tenuto = append(inst.Tenuto, typedTarget)
									}
								}
							}
						}
					case "Detached_legato":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Detached_legato = append(inst.Detached_legato, typedTarget)
									}
								}
							}
						}
					case "Staccatissimo":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Staccatissimo = append(inst.Staccatissimo, typedTarget)
									}
								}
							}
						}
					case "Spiccato":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Spiccato = append(inst.Spiccato, typedTarget)
									}
								}
							}
						}
					case "Scoop":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_line); ok {
										inst.Scoop = append(inst.Scoop, typedTarget)
									}
								}
							}
						}
					case "Plop":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_line); ok {
										inst.Plop = append(inst.Plop, typedTarget)
									}
								}
							}
						}
					case "Doit":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_line); ok {
										inst.Doit = append(inst.Doit, typedTarget)
									}
								}
							}
						}
					case "Falloff":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_line); ok {
										inst.Falloff = append(inst.Falloff, typedTarget)
									}
								}
							}
						}
					case "Breath_mark":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Breath_mark); ok {
										inst.Breath_mark = append(inst.Breath_mark, typedTarget)
									}
								}
							}
						}
					case "Caesura":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Caesura); ok {
										inst.Caesura = append(inst.Caesura, typedTarget)
									}
								}
							}
						}
					case "Stress":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Stress = append(inst.Stress, typedTarget)
									}
								}
							}
						}
					case "Unstress":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Unstress = append(inst.Unstress, typedTarget)
									}
								}
							}
						}
					case "Soft_accent":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Soft_accent = append(inst.Soft_accent, typedTarget)
									}
								}
							}
						}
					case "Other_articulation":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Other_placement_text); ok {
										inst.Other_articulation = append(inst.Other_articulation, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
									inst.Footnote = typedTarget
								}
							}
						}
					case "Level":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Level); ok {
									inst.Level = typedTarget
								}
							}
						}
					case "Divisions":
						inst.Divisions = GongExtractString(rhs)
					case "Key":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Key); ok {
										inst.Key = append(inst.Key, typedTarget)
									}
								}
							}
						}
					case "Time":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Time); ok {
										inst.Time = append(inst.Time, typedTarget)
									}
								}
							}
						}
					case "Staves":
						inst.Staves = GongExtractInt(rhs)
					case "Part_symbol":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part_symbol); ok {
									inst.Part_symbol = typedTarget
								}
							}
						}
					case "Instruments":
						inst.Instruments = GongExtractInt(rhs)
					case "Clef":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Clef); ok {
										inst.Clef = append(inst.Clef, typedTarget)
									}
								}
							}
						}
					case "Staff_details":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Staff_details); ok {
										inst.Staff_details = append(inst.Staff_details, typedTarget)
									}
								}
							}
						}
					case "Transpose":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Transpose); ok {
										inst.Transpose = append(inst.Transpose, typedTarget)
									}
								}
							}
						}
					case "For_part":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*For_part); ok {
										inst.For_part = append(inst.For_part, typedTarget)
									}
								}
							}
						}
					case "Directive":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_directive); ok {
										inst.Directive = append(inst.Directive, typedTarget)
									}
								}
							}
						}
					case "Measure_style":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Measure_style); ok {
										inst.Measure_style = append(inst.Measure_style, typedTarget)
									}
								}
							}
						}
					}
				case *Backup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					case "Footnote":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
									inst.Footnote = typedTarget
								}
							}
						}
					case "Level":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Level); ok {
									inst.Level = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bar_style_color); ok {
									inst.Bar_style = typedTarget
								}
							}
						}
					case "Footnote":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
									inst.Footnote = typedTarget
								}
							}
						}
					case "Level":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Level); ok {
									inst.Level = typedTarget
								}
							}
						}
					case "Wavy_line":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Wavy_line); ok {
									inst.Wavy_line = typedTarget
								}
							}
						}
					case "Segno_1":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Segno); ok {
									inst.Segno_1 = typedTarget
								}
							}
						}
					case "Coda_1":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Coda); ok {
									inst.Coda_1 = typedTarget
								}
							}
						}
					case "Fermata":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Fermata); ok {
									inst.Fermata = typedTarget
								}
							}
						}
					case "Ending":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Ending); ok {
									inst.Ending = typedTarget
								}
							}
						}
					case "Repeat":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Repeat); ok {
									inst.Repeat = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Style_text); ok {
									inst.Bass_separator = typedTarget
								}
							}
						}
					case "Bass_step":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bass_step); ok {
									inst.Bass_step = typedTarget
								}
							}
						}
					case "Bass_alter":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Harmony_alter); ok {
									inst.Bass_alter = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Release); ok {
									inst.Release = typedTarget
								}
							}
						}
					case "With_bar":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Placement_text); ok {
									inst.With_bar = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Image); ok {
									inst.Credit_image = typedTarget
								}
							}
						}
					case "Link":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Link); ok {
										inst.Link = append(inst.Link, typedTarget)
									}
								}
							}
						}
					case "Bookmark":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bookmark); ok {
										inst.Bookmark = append(inst.Bookmark, typedTarget)
									}
								}
							}
						}
					case "Credit_words":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text_id); ok {
										inst.Credit_words = append(inst.Credit_words, typedTarget)
									}
								}
							}
						}
					case "Credit_symbol":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_symbol_id); ok {
										inst.Credit_symbol = append(inst.Credit_symbol, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Scaling); ok {
									inst.Scaling = typedTarget
								}
							}
						}
					case "Concert_score":
						inst.Concert_score = GongExtractString(rhs)
					case "Page_layout":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Page_layout); ok {
									inst.Page_layout = typedTarget
								}
							}
						}
					case "System_layout":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System_layout); ok {
									inst.System_layout = typedTarget
								}
							}
						}
					case "Staff_layout":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Staff_layout); ok {
										inst.Staff_layout = append(inst.Staff_layout, typedTarget)
									}
								}
							}
						}
					case "Appearance":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Appearance); ok {
									inst.Appearance = typedTarget
								}
							}
						}
					case "Music_font":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_font); ok {
									inst.Music_font = typedTarget
								}
							}
						}
					case "Word_font":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_font); ok {
									inst.Word_font = typedTarget
								}
							}
						}
					case "Lyric_font":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Lyric_font); ok {
										inst.Lyric_font = append(inst.Lyric_font, typedTarget)
									}
								}
							}
						}
					case "Lyric_language":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Lyric_language); ok {
										inst.Lyric_language = append(inst.Lyric_language, typedTarget)
									}
								}
							}
						}
					}
				case *Degree:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Print_object":
						inst.Print_object = Enum_Yes_no(GongExtractString(rhs))
					case "Degree_value":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Degree_value); ok {
									inst.Degree_value = typedTarget
								}
							}
						}
					case "Degree_alter":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Degree_alter); ok {
									inst.Degree_alter = typedTarget
								}
							}
						}
					case "Degree_type":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Degree_type); ok {
									inst.Degree_type = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Direction_type); ok {
										inst.Direction_type = append(inst.Direction_type, typedTarget)
									}
								}
							}
						}
					case "Offset":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Offset); ok {
									inst.Offset = typedTarget
								}
							}
						}
					case "Footnote":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
									inst.Footnote = typedTarget
								}
							}
						}
					case "Level":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Level); ok {
									inst.Level = typedTarget
								}
							}
						}
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					case "Sound":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Sound); ok {
									inst.Sound = typedTarget
								}
							}
						}
					case "Listening":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Listening); ok {
									inst.Listening = typedTarget
								}
							}
						}
					}
				case *Direction_type:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Rehearsal":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text_id); ok {
										inst.Rehearsal = append(inst.Rehearsal, typedTarget)
									}
								}
							}
						}
					case "Segno":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Segno); ok {
										inst.Segno = append(inst.Segno, typedTarget)
									}
								}
							}
						}
					case "Coda":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Coda); ok {
										inst.Coda = append(inst.Coda, typedTarget)
									}
								}
							}
						}
					case "Words":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text_id); ok {
										inst.Words = append(inst.Words, typedTarget)
									}
								}
							}
						}
					case "Symbol":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_symbol_id); ok {
										inst.Symbol = append(inst.Symbol, typedTarget)
									}
								}
							}
						}
					case "Wedge":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Wedge); ok {
									inst.Wedge = typedTarget
								}
							}
						}
					case "Dynamics":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Dynamics); ok {
										inst.Dynamics = append(inst.Dynamics, typedTarget)
									}
								}
							}
						}
					case "Dashes":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Dashes); ok {
									inst.Dashes = typedTarget
								}
							}
						}
					case "Bracket":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bracket); ok {
									inst.Bracket = typedTarget
								}
							}
						}
					case "Pedal":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Pedal); ok {
									inst.Pedal = typedTarget
								}
							}
						}
					case "Metronome":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Metronome); ok {
									inst.Metronome = typedTarget
								}
							}
						}
					case "Octave_shift":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Octave_shift); ok {
									inst.Octave_shift = typedTarget
								}
							}
						}
					case "Harp_pedals":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Harp_pedals); ok {
									inst.Harp_pedals = typedTarget
								}
							}
						}
					case "Damp":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_print_style_align_id); ok {
									inst.Damp = typedTarget
								}
							}
						}
					case "Damp_all":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_print_style_align_id); ok {
									inst.Damp_all = typedTarget
								}
							}
						}
					case "Eyeglasses":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_print_style_align_id); ok {
									inst.Eyeglasses = typedTarget
								}
							}
						}
					case "String_mute":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*String_mute); ok {
									inst.String_mute = typedTarget
								}
							}
						}
					case "Scordatura":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Scordatura); ok {
									inst.Scordatura = typedTarget
								}
							}
						}
					case "Image":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Image); ok {
									inst.Image = typedTarget
								}
							}
						}
					case "Principal_voice":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Principal_voice); ok {
									inst.Principal_voice = typedTarget
								}
							}
						}
					case "Percussion":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Percussion); ok {
										inst.Percussion = append(inst.Percussion, typedTarget)
									}
								}
							}
						}
					case "Accordion_registration":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Accordion_registration); ok {
									inst.Accordion_registration = typedTarget
								}
							}
						}
					case "Staff_divide":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Staff_divide); ok {
									inst.Staff_divide = typedTarget
								}
							}
						}
					case "Other_direction":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Other_direction); ok {
									inst.Other_direction = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Other_text); ok {
										inst.Other_dynamics = append(inst.Other_dynamics, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Typed_text); ok {
										inst.Encoder = append(inst.Encoder, typedTarget)
									}
								}
							}
						}
					case "Software":
						inst.Software = GongExtractString(rhs)
					case "Encoding_description":
						inst.Encoding_description = GongExtractString(rhs)
					case "Supports":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Supports); ok {
										inst.Supports = append(inst.Supports, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Style_text); ok {
									inst.Prefix = typedTarget
								}
							}
						}
					case "Figure_number":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Style_text); ok {
									inst.Figure_number = typedTarget
								}
							}
						}
					case "Suffix":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Style_text); ok {
									inst.Suffix = typedTarget
								}
							}
						}
					case "Extend":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Extend); ok {
									inst.Extend = typedTarget
								}
							}
						}
					case "Footnote":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
									inst.Footnote = typedTarget
								}
							}
						}
					case "Level":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Level); ok {
									inst.Level = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Figure); ok {
										inst.Figure = append(inst.Figure, typedTarget)
									}
								}
							}
						}
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					case "Footnote":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
									inst.Footnote = typedTarget
								}
							}
						}
					case "Level":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Level); ok {
									inst.Level = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part_clef); ok {
									inst.Part_clef = typedTarget
								}
							}
						}
					case "Part_transpose":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part_transpose); ok {
									inst.Part_transpose = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
									inst.Footnote = typedTarget
								}
							}
						}
					case "Level":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Level); ok {
									inst.Level = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*First_fret); ok {
									inst.First_fret = typedTarget
								}
							}
						}
					case "Frame_note":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Frame_note); ok {
										inst.Frame_note = append(inst.Frame_note, typedTarget)
									}
								}
							}
						}
					}
				case *Frame_note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "String":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*String_type); ok {
									inst.String = typedTarget
								}
							}
						}
					case "Fret":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Fret); ok {
									inst.Fret = typedTarget
								}
							}
						}
					case "Fingering":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Fingering); ok {
									inst.Fingering = typedTarget
								}
							}
						}
					case "Barre":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Barre); ok {
									inst.Barre = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Feature); ok {
										inst.Feature = append(inst.Feature, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Harmon_closed); ok {
									inst.Harmon_closed = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Root); ok {
									inst.Root = typedTarget
								}
							}
						}
					case "Numeral":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Numeral); ok {
									inst.Numeral = typedTarget
								}
							}
						}
					case "Function":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Style_text); ok {
									inst.Function = typedTarget
								}
							}
						}
					case "Kind":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Kind); ok {
									inst.Kind = typedTarget
								}
							}
						}
					case "Inversion":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Inversion); ok {
									inst.Inversion = typedTarget
								}
							}
						}
					case "Bass":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bass); ok {
									inst.Bass = typedTarget
								}
							}
						}
					case "Degree":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Degree); ok {
										inst.Degree = append(inst.Degree, typedTarget)
									}
								}
							}
						}
					case "Frame":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Frame); ok {
									inst.Frame = typedTarget
								}
							}
						}
					case "Offset":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Offset); ok {
									inst.Offset = typedTarget
								}
							}
						}
					case "Footnote":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
									inst.Footnote = typedTarget
								}
							}
						}
					case "Level":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Level); ok {
									inst.Level = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Pedal_tuning); ok {
										inst.Pedal_tuning = append(inst.Pedal_tuning, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Hole_closed); ok {
									inst.Hole_closed = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Typed_text); ok {
										inst.Creator = append(inst.Creator, typedTarget)
									}
								}
							}
						}
					case "Rights":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Typed_text); ok {
										inst.Rights = append(inst.Rights, typedTarget)
									}
								}
							}
						}
					case "Encoding":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Encoding); ok {
									inst.Encoding = typedTarget
								}
							}
						}
					case "Source":
						inst.Source = GongExtractString(rhs)
					case "Relation":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Typed_text); ok {
										inst.Relation = append(inst.Relation, typedTarget)
									}
								}
							}
						}
					case "Miscellaneous":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Miscellaneous); ok {
									inst.Miscellaneous = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Virtual_instrument); ok {
									inst.Virtual_instrument = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Cancel); ok {
									inst.Cancel = typedTarget
								}
							}
						}
					case "Fifths":
						inst.Fifths = GongExtractInt(rhs)
					case "Mode":
						inst.Mode = GongExtractString(rhs)
					case "Key_step":
						inst.Key_step = Enum_Step(GongExtractString(rhs))
					case "Key_alter":
						inst.Key_alter = GongExtractString(rhs)
					case "Key_accidental":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Key_accidental); ok {
									inst.Key_accidental = typedTarget
								}
							}
						}
					case "Key_octave":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Key_octave); ok {
										inst.Key_octave = append(inst.Key_octave, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Assess); ok {
										inst.Assess = append(inst.Assess, typedTarget)
									}
								}
							}
						}
					case "Wait":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Wait); ok {
										inst.Wait = append(inst.Wait, typedTarget)
									}
								}
							}
						}
					case "Other_listen":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Other_listening); ok {
										inst.Other_listen = append(inst.Other_listen, typedTarget)
									}
								}
							}
						}
					}
				case *Listening:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Sync":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Sync); ok {
										inst.Sync = append(inst.Sync, typedTarget)
									}
								}
							}
						}
					case "Other_listening":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Other_listening); ok {
										inst.Other_listening = append(inst.Other_listening, typedTarget)
									}
								}
							}
						}
					case "Offset":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Offset); ok {
									inst.Offset = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Elision); ok {
										inst.Elision = append(inst.Elision, typedTarget)
									}
								}
							}
						}
					case "Syllabic":
						inst.Syllabic = GongExtractString(rhs)
					case "Text":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Text_element_data); ok {
										inst.Text = append(inst.Text, typedTarget)
									}
								}
							}
						}
					case "Extend":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Extend); ok {
									inst.Extend = typedTarget
								}
							}
						}
					case "Laughing":
						inst.Laughing = GongExtractString(rhs)
					case "Humming":
						inst.Humming = GongExtractString(rhs)
					case "End_line":
						inst.End_line = GongExtractString(rhs)
					case "End_paragraph":
						inst.End_paragraph = GongExtractString(rhs)
					case "Footnote":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
									inst.Footnote = typedTarget
								}
							}
						}
					case "Level":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Level); ok {
									inst.Level = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Multiple_rest); ok {
									inst.Multiple_rest = typedTarget
								}
							}
						}
					case "Measure_repeat":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Measure_repeat); ok {
									inst.Measure_repeat = typedTarget
								}
							}
						}
					case "Beat_repeat":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Beat_repeat); ok {
									inst.Beat_repeat = typedTarget
								}
							}
						}
					case "Slash":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Slash); ok {
									inst.Slash = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Per_minute); ok {
									inst.Per_minute = typedTarget
								}
							}
						}
					case "Beat_unit_tied":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Beat_unit_tied); ok {
										inst.Beat_unit_tied = append(inst.Beat_unit_tied, typedTarget)
									}
								}
							}
						}
					case "Metronome_arrows":
						inst.Metronome_arrows = GongExtractString(rhs)
					case "Metronome_relation":
						inst.Metronome_relation = GongExtractString(rhs)
					case "Metronome_note":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Metronome_note); ok {
										inst.Metronome_note = append(inst.Metronome_note, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Metronome_beam); ok {
										inst.Metronome_beam = append(inst.Metronome_beam, typedTarget)
									}
								}
							}
						}
					case "Metronome_tied":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Metronome_tied); ok {
									inst.Metronome_tied = typedTarget
								}
							}
						}
					case "Metronome_tuplet":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Metronome_tuplet); ok {
									inst.Metronome_tuplet = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Miscellaneous_field); ok {
										inst.Miscellaneous_field = append(inst.Miscellaneous_field, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
										inst.Display_text = append(inst.Display_text, typedTarget)
									}
								}
							}
						}
					case "Accidental_text":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Accidental_text); ok {
										inst.Accidental_text = append(inst.Accidental_text, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
									inst.Footnote = typedTarget
								}
							}
						}
					case "Level":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Level); ok {
									inst.Level = typedTarget
								}
							}
						}
					case "Tied":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Tied); ok {
										inst.Tied = append(inst.Tied, typedTarget)
									}
								}
							}
						}
					case "Slur":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Slur); ok {
										inst.Slur = append(inst.Slur, typedTarget)
									}
								}
							}
						}
					case "Tuplet":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Tuplet); ok {
										inst.Tuplet = append(inst.Tuplet, typedTarget)
									}
								}
							}
						}
					case "Glissando":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Glissando); ok {
										inst.Glissando = append(inst.Glissando, typedTarget)
									}
								}
							}
						}
					case "Slide":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Slide); ok {
										inst.Slide = append(inst.Slide, typedTarget)
									}
								}
							}
						}
					case "Ornaments":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Ornaments); ok {
										inst.Ornaments = append(inst.Ornaments, typedTarget)
									}
								}
							}
						}
					case "Technical":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Technical); ok {
										inst.Technical = append(inst.Technical, typedTarget)
									}
								}
							}
						}
					case "Articulations":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Articulations); ok {
										inst.Articulations = append(inst.Articulations, typedTarget)
									}
								}
							}
						}
					case "Dynamics":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Dynamics); ok {
										inst.Dynamics = append(inst.Dynamics, typedTarget)
									}
								}
							}
						}
					case "Fermata":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Fermata); ok {
										inst.Fermata = append(inst.Fermata, typedTarget)
									}
								}
							}
						}
					case "Arpeggiate":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Arpeggiate); ok {
										inst.Arpeggiate = append(inst.Arpeggiate, typedTarget)
									}
								}
							}
						}
					case "Non_arpeggiate":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Non_arpeggiate); ok {
										inst.Non_arpeggiate = append(inst.Non_arpeggiate, typedTarget)
									}
								}
							}
						}
					case "Accidental_mark":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Accidental_mark); ok {
										inst.Accidental_mark = append(inst.Accidental_mark, typedTarget)
									}
								}
							}
						}
					case "Other_notation":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Other_notation); ok {
										inst.Other_notation = append(inst.Other_notation, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Grace); ok {
									inst.Grace = typedTarget
								}
							}
						}
					case "Chord":
						inst.Chord = GongExtractString(rhs)
					case "Pitch":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Pitch); ok {
									inst.Pitch = typedTarget
								}
							}
						}
					case "Unpitched":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Unpitched); ok {
									inst.Unpitched = typedTarget
								}
							}
						}
					case "Rest":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Rest); ok {
									inst.Rest = typedTarget
								}
							}
						}
					case "Tie":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Tie); ok {
									inst.Tie = typedTarget
								}
							}
						}
					case "Cue":
						inst.Cue = GongExtractString(rhs)
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					case "Instrument":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Instrument); ok {
										inst.Instrument = append(inst.Instrument, typedTarget)
									}
								}
							}
						}
					case "Footnote":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
									inst.Footnote = typedTarget
								}
							}
						}
					case "Level":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Level); ok {
									inst.Level = typedTarget
								}
							}
						}
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Type":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Note_type); ok {
									inst.Type = typedTarget
								}
							}
						}
					case "Dot":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Dot = append(inst.Dot, typedTarget)
									}
								}
							}
						}
					case "Accidental":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Accidental); ok {
									inst.Accidental = typedTarget
								}
							}
						}
					case "Time_modification":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Time_modification); ok {
									inst.Time_modification = typedTarget
								}
							}
						}
					case "Stem":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Stem); ok {
									inst.Stem = typedTarget
								}
							}
						}
					case "Notehead":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Notehead); ok {
									inst.Notehead = typedTarget
								}
							}
						}
					case "Notehead_text":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Notehead_text); ok {
									inst.Notehead_text = typedTarget
								}
							}
						}
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					case "Beam":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Beam); ok {
									inst.Beam = typedTarget
								}
							}
						}
					case "Notations":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Notations); ok {
										inst.Notations = append(inst.Notations, typedTarget)
									}
								}
							}
						}
					case "Lyric":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Lyric); ok {
										inst.Lyric = append(inst.Lyric, typedTarget)
									}
								}
							}
						}
					case "Play":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Play); ok {
									inst.Play = typedTarget
								}
							}
						}
					case "Listen":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Listen); ok {
									inst.Listen = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
										inst.Display_text = append(inst.Display_text, typedTarget)
									}
								}
							}
						}
					case "Accidental_text":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Accidental_text); ok {
										inst.Accidental_text = append(inst.Accidental_text, typedTarget)
									}
								}
							}
						}
					}
				case *Numeral:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Numeral_root":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Numeral_root); ok {
									inst.Numeral_root = typedTarget
								}
							}
						}
					case "Numeral_alter":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Harmony_alter); ok {
									inst.Numeral_alter = typedTarget
								}
							}
						}
					case "Numeral_key":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Numeral_key); ok {
									inst.Numeral_key = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_trill_sound); ok {
										inst.Trill_mark = append(inst.Trill_mark, typedTarget)
									}
								}
							}
						}
					case "Turn":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Horizontal_turn); ok {
										inst.Turn = append(inst.Turn, typedTarget)
									}
								}
							}
						}
					case "Delayed_turn":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Horizontal_turn); ok {
										inst.Delayed_turn = append(inst.Delayed_turn, typedTarget)
									}
								}
							}
						}
					case "Inverted_turn":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Horizontal_turn); ok {
										inst.Inverted_turn = append(inst.Inverted_turn, typedTarget)
									}
								}
							}
						}
					case "Delayed_inverted_turn":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Horizontal_turn); ok {
										inst.Delayed_inverted_turn = append(inst.Delayed_inverted_turn, typedTarget)
									}
								}
							}
						}
					case "Vertical_turn":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_trill_sound); ok {
										inst.Vertical_turn = append(inst.Vertical_turn, typedTarget)
									}
								}
							}
						}
					case "Inverted_vertical_turn":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_trill_sound); ok {
										inst.Inverted_vertical_turn = append(inst.Inverted_vertical_turn, typedTarget)
									}
								}
							}
						}
					case "Shake":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_trill_sound); ok {
										inst.Shake = append(inst.Shake, typedTarget)
									}
								}
							}
						}
					case "Wavy_line":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Wavy_line); ok {
										inst.Wavy_line = append(inst.Wavy_line, typedTarget)
									}
								}
							}
						}
					case "Mordent":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Mordent); ok {
										inst.Mordent = append(inst.Mordent, typedTarget)
									}
								}
							}
						}
					case "Inverted_mordent":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Mordent); ok {
										inst.Inverted_mordent = append(inst.Inverted_mordent, typedTarget)
									}
								}
							}
						}
					case "Schleifer":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Schleifer = append(inst.Schleifer, typedTarget)
									}
								}
							}
						}
					case "Tremolo":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Tremolo); ok {
										inst.Tremolo = append(inst.Tremolo, typedTarget)
									}
								}
							}
						}
					case "Haydn":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_trill_sound); ok {
										inst.Haydn = append(inst.Haydn, typedTarget)
									}
								}
							}
						}
					case "Other_ornament":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Other_placement_text); ok {
										inst.Other_ornament = append(inst.Other_ornament, typedTarget)
									}
								}
							}
						}
					case "Accidental_mark":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Accidental_mark); ok {
										inst.Accidental_mark = append(inst.Accidental_mark, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Page_margins); ok {
									inst.Page_margins = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group_name); ok {
									inst.Group_name = typedTarget
								}
							}
						}
					case "Group_name_display":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Name_display); ok {
									inst.Group_name_display = typedTarget
								}
							}
						}
					case "Group_abbreviation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group_name); ok {
									inst.Group_abbreviation = typedTarget
								}
							}
						}
					case "Group_abbreviation_display":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Name_display); ok {
									inst.Group_abbreviation_display = typedTarget
								}
							}
						}
					case "Group_symbol":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group_symbol); ok {
									inst.Group_symbol = typedTarget
								}
							}
						}
					case "Group_barline":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Group_barline); ok {
									inst.Group_barline = typedTarget
								}
							}
						}
					case "Group_time":
						inst.Group_time = GongExtractString(rhs)
					case "Footnote":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Formatted_text); ok {
									inst.Footnote = typedTarget
								}
							}
						}
					case "Level":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Level); ok {
									inst.Level = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Instrument_link); ok {
										inst.Instrument_link = append(inst.Instrument_link, typedTarget)
									}
								}
							}
						}
					case "Group_link":
						inst.Group_link = GongExtractString(rhs)
					}
				case *Part_list:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Part_group":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part_group); ok {
									inst.Part_group = typedTarget
								}
							}
						}
					case "Score_part":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Score_part); ok {
									inst.Score_part = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Glass); ok {
									inst.Glass = typedTarget
								}
							}
						}
					case "Metal":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Metal); ok {
									inst.Metal = typedTarget
								}
							}
						}
					case "Wood":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Wood); ok {
									inst.Wood = typedTarget
								}
							}
						}
					case "Pitched":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Pitched); ok {
									inst.Pitched = typedTarget
								}
							}
						}
					case "Membrane":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Membrane); ok {
									inst.Membrane = typedTarget
								}
							}
						}
					case "Effect":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Effect); ok {
									inst.Effect = typedTarget
								}
							}
						}
					case "Timpani":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Timpani); ok {
									inst.Timpani = typedTarget
								}
							}
						}
					case "Beater":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Beater); ok {
									inst.Beater = typedTarget
								}
							}
						}
					case "Stick":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Stick); ok {
									inst.Stick = typedTarget
								}
							}
						}
					case "Stick_location":
						inst.Stick_location = GongExtractString(rhs)
					case "Other_percussion":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Other_text); ok {
									inst.Other_percussion = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Other_play); ok {
										inst.Other_play = append(inst.Other_play, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Page_layout); ok {
									inst.Page_layout = typedTarget
								}
							}
						}
					case "System_layout":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System_layout); ok {
									inst.System_layout = typedTarget
								}
							}
						}
					case "Staff_layout":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Staff_layout); ok {
										inst.Staff_layout = append(inst.Staff_layout, typedTarget)
									}
								}
							}
						}
					case "Measure_layout":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Measure_layout); ok {
									inst.Measure_layout = typedTarget
								}
							}
						}
					case "Measure_numbering":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Measure_numbering); ok {
									inst.Measure_numbering = typedTarget
								}
							}
						}
					case "Part_name_display":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Name_display); ok {
									inst.Part_name_display = typedTarget
								}
							}
						}
					case "Part_abbreviation_display":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Name_display); ok {
									inst.Part_abbreviation_display = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Root_step); ok {
									inst.Root_step = typedTarget
								}
							}
						}
					case "Root_alter":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Harmony_alter); ok {
									inst.Root_alter = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Accord); ok {
										inst.Accord = append(inst.Accord, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Virtual_instrument); ok {
									inst.Virtual_instrument = typedTarget
								}
							}
						}
					}
				case *Score_part:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Identification":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Identification); ok {
									inst.Identification = typedTarget
								}
							}
						}
					case "Part_link":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part_link); ok {
										inst.Part_link = append(inst.Part_link, typedTarget)
									}
								}
							}
						}
					case "Part_name":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part_name); ok {
									inst.Part_name = typedTarget
								}
							}
						}
					case "Part_name_display":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Name_display); ok {
									inst.Part_name_display = typedTarget
								}
							}
						}
					case "Part_abbreviation":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part_name); ok {
									inst.Part_abbreviation = typedTarget
								}
							}
						}
					case "Part_abbreviation_display":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Name_display); ok {
									inst.Part_abbreviation_display = typedTarget
								}
							}
						}
					case "Group":
						inst.Group = GongExtractString(rhs)
					case "Score_instrument":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Score_instrument); ok {
										inst.Score_instrument = append(inst.Score_instrument, typedTarget)
									}
								}
							}
						}
					case "Player":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Player); ok {
										inst.Player = append(inst.Player, typedTarget)
									}
								}
							}
						}
					case "Midi_device":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Midi_device); ok {
										inst.Midi_device = append(inst.Midi_device, typedTarget)
									}
								}
							}
						}
					case "Midi_instrument":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Midi_instrument); ok {
										inst.Midi_instrument = append(inst.Midi_instrument, typedTarget)
									}
								}
							}
						}
					}
				case *Score_partwise:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Version":
						inst.Version = GongExtractString(rhs)
					case "Work":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Work); ok {
									inst.Work = typedTarget
								}
							}
						}
					case "Movement_number":
						inst.Movement_number = GongExtractString(rhs)
					case "Movement_title":
						inst.Movement_title = GongExtractString(rhs)
					case "Identification":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Identification); ok {
									inst.Identification = typedTarget
								}
							}
						}
					case "Defaults":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Defaults); ok {
									inst.Defaults = typedTarget
								}
							}
						}
					case "Credit":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Credit); ok {
										inst.Credit = append(inst.Credit, typedTarget)
									}
								}
							}
						}
					case "Part_list":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part_list); ok {
									inst.Part_list = typedTarget
								}
							}
						}
					case "Part":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_part); ok {
										inst.Part = append(inst.Part, typedTarget)
									}
								}
							}
						}
					}
				case *Score_timewise:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Version":
						inst.Version = GongExtractString(rhs)
					case "Work":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Work); ok {
									inst.Work = typedTarget
								}
							}
						}
					case "Movement_number":
						inst.Movement_number = GongExtractString(rhs)
					case "Movement_title":
						inst.Movement_title = GongExtractString(rhs)
					case "Identification":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Identification); ok {
									inst.Identification = typedTarget
								}
							}
						}
					case "Defaults":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Defaults); ok {
									inst.Defaults = typedTarget
								}
							}
						}
					case "Credit":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Credit); ok {
										inst.Credit = append(inst.Credit, typedTarget)
									}
								}
							}
						}
					case "Part_list":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Part_list); ok {
									inst.Part_list = typedTarget
								}
							}
						}
					case "Measure":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*A_measure_1); ok {
										inst.Measure = append(inst.Measure, typedTarget)
									}
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Instrument_change); ok {
										inst.Instrument_change = append(inst.Instrument_change, typedTarget)
									}
								}
							}
						}
					case "Midi_device":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Midi_device); ok {
										inst.Midi_device = append(inst.Midi_device, typedTarget)
									}
								}
							}
						}
					case "Midi_instrument":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Midi_instrument); ok {
										inst.Midi_instrument = append(inst.Midi_instrument, typedTarget)
									}
								}
							}
						}
					case "Play":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Play); ok {
										inst.Play = append(inst.Play, typedTarget)
									}
								}
							}
						}
					case "Swing":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Swing); ok {
									inst.Swing = typedTarget
								}
							}
						}
					case "Offset":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Offset); ok {
									inst.Offset = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Line_detail); ok {
										inst.Line_detail = append(inst.Line_detail, typedTarget)
									}
								}
							}
						}
					case "Staff_tuning":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Staff_tuning); ok {
										inst.Staff_tuning = append(inst.Staff_tuning, typedTarget)
									}
								}
							}
						}
					case "Capo":
						inst.Capo = GongExtractInt(rhs)
					case "Staff_size":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Staff_size); ok {
									inst.Staff_size = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_print_object_style_align); ok {
									inst.Left_divider = typedTarget
								}
							}
						}
					case "Right_divider":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_print_object_style_align); ok {
									inst.Right_divider = typedTarget
								}
							}
						}
					}
				case *System_layout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "System_margins":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System_margins); ok {
									inst.System_margins = typedTarget
								}
							}
						}
					case "System_distance":
						inst.System_distance = GongExtractString(rhs)
					case "Top_system_distance":
						inst.Top_system_distance = GongExtractString(rhs)
					case "System_dividers":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*System_dividers); ok {
									inst.System_dividers = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Up_bow = append(inst.Up_bow, typedTarget)
									}
								}
							}
						}
					case "Down_bow":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Down_bow = append(inst.Down_bow, typedTarget)
									}
								}
							}
						}
					case "Harmonic":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Harmonic); ok {
										inst.Harmonic = append(inst.Harmonic, typedTarget)
									}
								}
							}
						}
					case "Open_string":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Open_string = append(inst.Open_string, typedTarget)
									}
								}
							}
						}
					case "Thumb_position":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Thumb_position = append(inst.Thumb_position, typedTarget)
									}
								}
							}
						}
					case "Fingering":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Fingering); ok {
										inst.Fingering = append(inst.Fingering, typedTarget)
									}
								}
							}
						}
					case "Pluck":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Placement_text); ok {
										inst.Pluck = append(inst.Pluck, typedTarget)
									}
								}
							}
						}
					case "Double_tongue":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Double_tongue = append(inst.Double_tongue, typedTarget)
									}
								}
							}
						}
					case "Triple_tongue":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Triple_tongue = append(inst.Triple_tongue, typedTarget)
									}
								}
							}
						}
					case "Stopped":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement_smufl); ok {
										inst.Stopped = append(inst.Stopped, typedTarget)
									}
								}
							}
						}
					case "Snap_pizzicato":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Snap_pizzicato = append(inst.Snap_pizzicato, typedTarget)
									}
								}
							}
						}
					case "Fret":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Fret); ok {
										inst.Fret = append(inst.Fret, typedTarget)
									}
								}
							}
						}
					case "String":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*String_type); ok {
										inst.String = append(inst.String, typedTarget)
									}
								}
							}
						}
					case "Hammer_on":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Hammer_on_pull_off); ok {
										inst.Hammer_on = append(inst.Hammer_on, typedTarget)
									}
								}
							}
						}
					case "Pull_off":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Hammer_on_pull_off); ok {
										inst.Pull_off = append(inst.Pull_off, typedTarget)
									}
								}
							}
						}
					case "Bend":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Bend); ok {
										inst.Bend = append(inst.Bend, typedTarget)
									}
								}
							}
						}
					case "Tap":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Tap); ok {
										inst.Tap = append(inst.Tap, typedTarget)
									}
								}
							}
						}
					case "Heel":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Heel_toe); ok {
										inst.Heel = append(inst.Heel, typedTarget)
									}
								}
							}
						}
					case "Toe":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Heel_toe); ok {
										inst.Toe = append(inst.Toe, typedTarget)
									}
								}
							}
						}
					case "Fingernails":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Fingernails = append(inst.Fingernails, typedTarget)
									}
								}
							}
						}
					case "Hole":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Hole); ok {
										inst.Hole = append(inst.Hole, typedTarget)
									}
								}
							}
						}
					case "Arrow":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Arrow); ok {
										inst.Arrow = append(inst.Arrow, typedTarget)
									}
								}
							}
						}
					case "Handbell":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Handbell); ok {
										inst.Handbell = append(inst.Handbell, typedTarget)
									}
								}
							}
						}
					case "Brass_bend":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Brass_bend = append(inst.Brass_bend, typedTarget)
									}
								}
							}
						}
					case "Flip":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Flip = append(inst.Flip, typedTarget)
									}
								}
							}
						}
					case "Smear":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Smear = append(inst.Smear, typedTarget)
									}
								}
							}
						}
					case "Open":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement_smufl); ok {
										inst.Open = append(inst.Open, typedTarget)
									}
								}
							}
						}
					case "Half_muted":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement_smufl); ok {
										inst.Half_muted = append(inst.Half_muted, typedTarget)
									}
								}
							}
						}
					case "Harmon_mute":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Harmon_mute); ok {
										inst.Harmon_mute = append(inst.Harmon_mute, typedTarget)
									}
								}
							}
						}
					case "Golpe":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Empty_placement); ok {
										inst.Golpe = append(inst.Golpe, typedTarget)
									}
								}
							}
						}
					case "Other_technical":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Other_placement_text); ok {
										inst.Other_technical = append(inst.Other_technical, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Interchangeable); ok {
									inst.Interchangeable = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Tuplet_portion); ok {
									inst.Tuplet_actual = typedTarget
								}
							}
						}
					case "Tuplet_normal":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Tuplet_portion); ok {
									inst.Tuplet_normal = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Tuplet_number); ok {
									inst.Tuplet_number = typedTarget
								}
							}
						}
					case "Tuplet_type":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Tuplet_type); ok {
									inst.Tuplet_type = typedTarget
								}
							}
						}
					case "Tuplet_dot":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Tuplet_dot); ok {
										inst.Tuplet_dot = append(inst.Tuplet_dot, typedTarget)
									}
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Opus); ok {
									inst.Opus = typedTarget
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
