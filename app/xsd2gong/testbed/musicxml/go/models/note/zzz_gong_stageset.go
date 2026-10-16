// generated code - do not edit
package note

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
	"runtime/debug"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
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
		for _, accidental := range __gong__sortStageSetInstances(stageSet.Stage.Accidentals, stageSet.Stage.Accidental_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			accidentalIdent := "__note" + accidental.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Accidental{Name: %s}).Stage(stageSet.Stage)", accidentalIdent, __gong__toRawStringLiteral(accidental.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Cautionary = %s", accidentalIdent, __gong__toRawStringLiteral(string(accidental.Cautionary))))
			values.WriteString(fmt.Sprintf("\n\t%s.Editorial = %s", accidentalIdent, __gong__toRawStringLiteral(string(accidental.Editorial))))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", accidentalIdent, __gong__toRawStringLiteral(string(accidental.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.Bracket = %s", accidentalIdent, __gong__toRawStringLiteral(string(accidental.Bracket))))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %s", accidentalIdent, __gong__toRawStringLiteral(string(accidental.Size))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", accidentalIdent, __gong__toRawStringLiteral(string(accidental.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", accidentalIdent, __gong__toRawStringLiteral(string(accidental.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", accidentalIdent, __gong__toRawStringLiteral(accidental.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", accidentalIdent, __gong__toRawStringLiteral(string(accidental.Value))))
		}
	}
	if stageSet.Stage != nil {
		for _, accidentalmark := range __gong__sortStageSetInstances(stageSet.Stage.AccidentalMarks, stageSet.Stage.AccidentalMark_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			accidentalmarkIdent := "__note" + accidentalmark.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.AccidentalMark{Name: %s}).Stage(stageSet.Stage)", accidentalmarkIdent, __gong__toRawStringLiteral(accidentalmark.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", accidentalmarkIdent, __gong__toRawStringLiteral(accidentalmark.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", accidentalmarkIdent, __gong__toRawStringLiteral(accidentalmark.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", accidentalmarkIdent, __gong__toRawStringLiteral(string(accidentalmark.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.Bracket = %s", accidentalmarkIdent, __gong__toRawStringLiteral(string(accidentalmark.Bracket))))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %s", accidentalmarkIdent, __gong__toRawStringLiteral(string(accidentalmark.Size))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", accidentalmarkIdent, __gong__toRawStringLiteral(accidentalmark.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", accidentalmarkIdent, __gong__toRawStringLiteral(accidentalmark.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", accidentalmarkIdent, __gong__toRawStringLiteral(accidentalmark.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", accidentalmarkIdent, __gong__toRawStringLiteral(accidentalmark.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", accidentalmarkIdent, __gong__toRawStringLiteral(accidentalmark.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", accidentalmarkIdent, __gong__toRawStringLiteral(string(accidentalmark.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", accidentalmarkIdent, __gong__toRawStringLiteral(accidentalmark.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", accidentalmarkIdent, __gong__toRawStringLiteral(string(accidentalmark.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", accidentalmarkIdent, __gong__toRawStringLiteral(accidentalmark.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", accidentalmarkIdent, __gong__toRawStringLiteral(string(accidentalmark.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", accidentalmarkIdent, __gong__toRawStringLiteral(accidentalmark.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", accidentalmarkIdent, __gong__toRawStringLiteral(string(accidentalmark.Value))))
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
			arpeggiateIdent := "__note" + arpeggiate.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Arpeggiate{Name: %s}).Stage(stageSet.Stage)", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", arpeggiateIdent, arpeggiate.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", arpeggiateIdent, __gong__toRawStringLiteral(string(arpeggiate.Direction))))
			values.WriteString(fmt.Sprintf("\n\t%s.Unbroken = %s", arpeggiateIdent, __gong__toRawStringLiteral(string(arpeggiate.Unbroken))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", arpeggiateIdent, __gong__toRawStringLiteral(string(arpeggiate.Placement))))
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
			arrowIdent := "__note" + arrow.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Arrow{Name: %s}).Stage(stageSet.Stage)", arrowIdent, __gong__toRawStringLiteral(arrow.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", arrowIdent, __gong__toRawStringLiteral(arrow.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", arrowIdent, __gong__toRawStringLiteral(arrow.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", arrowIdent, __gong__toRawStringLiteral(arrow.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", arrowIdent, __gong__toRawStringLiteral(arrow.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", arrowIdent, __gong__toRawStringLiteral(arrow.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", arrowIdent, __gong__toRawStringLiteral(string(arrow.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", arrowIdent, __gong__toRawStringLiteral(arrow.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", arrowIdent, __gong__toRawStringLiteral(string(arrow.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", arrowIdent, __gong__toRawStringLiteral(string(arrow.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", arrowIdent, __gong__toRawStringLiteral(arrow.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArrowDirection = %s", arrowIdent, __gong__toRawStringLiteral(string(arrow.ArrowDirection))))
			values.WriteString(fmt.Sprintf("\n\t%s.ArrowStyle = %s", arrowIdent, __gong__toRawStringLiteral(string(arrow.ArrowStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.CircularArrow = %s", arrowIdent, __gong__toRawStringLiteral(string(arrow.CircularArrow))))
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
			articulationsIdent := "__note" + articulations.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Articulations{Name: %s}).Stage(stageSet.Stage)", articulationsIdent, __gong__toRawStringLiteral(articulations.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", articulationsIdent, __gong__toRawStringLiteral(articulations.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", articulationsIdent, __gong__toRawStringLiteral(articulations.Id)))
			for _, elem := range articulations.StrongAccent {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StrongAccent = append(%s.StrongAccent, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Scoop {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Scoop = append(%s.Scoop, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Plop {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Plop = append(%s.Plop, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Doit {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Doit = append(%s.Doit, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Falloff {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Falloff = append(%s.Falloff, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.BreathMark {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BreathMark = append(%s.BreathMark, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Caesura {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Caesura = append(%s.Caesura, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.OtherArticulation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherArticulation = append(%s.OtherArticulation, %s)", articulationsIdent, articulationsIdent, targetIdent))
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
			assessIdent := "__note" + assess.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Assess{Name: %s}).Stage(stageSet.Stage)", assessIdent, __gong__toRawStringLiteral(assess.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", assessIdent, __gong__toRawStringLiteral(assess.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", assessIdent, __gong__toRawStringLiteral(string(assess.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", assessIdent, __gong__toRawStringLiteral(assess.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeOnly = %s", assessIdent, __gong__toRawStringLiteral(assess.TimeOnly)))
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
			backupIdent := "__note" + backup.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Backup{Name: %s}).Stage(stageSet.Stage)", backupIdent, __gong__toRawStringLiteral(backup.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", backupIdent, __gong__toRawStringLiteral(backup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = %s", backupIdent, __gong__toRawStringLiteral(backup.Duration)))
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
			beamIdent := "__note" + beam.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Beam{Name: %s}).Stage(stageSet.Stage)", beamIdent, __gong__toRawStringLiteral(beam.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", beamIdent, __gong__toRawStringLiteral(beam.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", beamIdent, beam.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Repeater = %s", beamIdent, __gong__toRawStringLiteral(string(beam.Repeater))))
			values.WriteString(fmt.Sprintf("\n\t%s.Fan = %s", beamIdent, __gong__toRawStringLiteral(string(beam.Fan))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", beamIdent, __gong__toRawStringLiteral(beam.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", beamIdent, __gong__toRawStringLiteral(beam.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", beamIdent, __gong__toRawStringLiteral(string(beam.Value))))
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
			bendIdent := "__note" + bend.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Bend{Name: %s}).Stage(stageSet.Stage)", bendIdent, __gong__toRawStringLiteral(bend.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bendIdent, __gong__toRawStringLiteral(bend.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Shape = %s", bendIdent, __gong__toRawStringLiteral(string(bend.Shape))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", bendIdent, __gong__toRawStringLiteral(bend.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", bendIdent, __gong__toRawStringLiteral(bend.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", bendIdent, __gong__toRawStringLiteral(bend.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", bendIdent, __gong__toRawStringLiteral(bend.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", bendIdent, __gong__toRawStringLiteral(bend.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", bendIdent, __gong__toRawStringLiteral(string(bend.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", bendIdent, __gong__toRawStringLiteral(bend.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", bendIdent, __gong__toRawStringLiteral(string(bend.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", bendIdent, __gong__toRawStringLiteral(bend.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Accelerate = %s", bendIdent, __gong__toRawStringLiteral(string(bend.Accelerate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", bendIdent, __gong__toRawStringLiteral(bend.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.FirstBeat = %s", bendIdent, __gong__toRawStringLiteral(bend.FirstBeat)))
			values.WriteString(fmt.Sprintf("\n\t%s.LastBeat = %s", bendIdent, __gong__toRawStringLiteral(bend.LastBeat)))
			values.WriteString(fmt.Sprintf("\n\t%s.BendAlter = %s", bendIdent, __gong__toRawStringLiteral(bend.BendAlter)))
			if bend.Release != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + bend.Release.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Release = %s", bendIdent, targetIdent))
			}
			if bend.WithBar != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + bend.WithBar.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WithBar = %s", bendIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, breathmark := range __gong__sortStageSetInstances(stageSet.Stage.BreathMarks, stageSet.Stage.BreathMark_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			breathmarkIdent := "__note" + breathmark.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.BreathMark{Name: %s}).Stage(stageSet.Stage)", breathmarkIdent, __gong__toRawStringLiteral(breathmark.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", breathmarkIdent, __gong__toRawStringLiteral(breathmark.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", breathmarkIdent, __gong__toRawStringLiteral(breathmark.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", breathmarkIdent, __gong__toRawStringLiteral(breathmark.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", breathmarkIdent, __gong__toRawStringLiteral(breathmark.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", breathmarkIdent, __gong__toRawStringLiteral(breathmark.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", breathmarkIdent, __gong__toRawStringLiteral(breathmark.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", breathmarkIdent, __gong__toRawStringLiteral(string(breathmark.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", breathmarkIdent, __gong__toRawStringLiteral(breathmark.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", breathmarkIdent, __gong__toRawStringLiteral(string(breathmark.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", breathmarkIdent, __gong__toRawStringLiteral(breathmark.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", breathmarkIdent, __gong__toRawStringLiteral(string(breathmark.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", breathmarkIdent, __gong__toRawStringLiteral(string(breathmark.Value))))
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
			caesuraIdent := "__note" + caesura.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Caesura{Name: %s}).Stage(stageSet.Stage)", caesuraIdent, __gong__toRawStringLiteral(caesura.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", caesuraIdent, __gong__toRawStringLiteral(string(caesura.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", caesuraIdent, __gong__toRawStringLiteral(string(caesura.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", caesuraIdent, __gong__toRawStringLiteral(caesura.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", caesuraIdent, __gong__toRawStringLiteral(string(caesura.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", caesuraIdent, __gong__toRawStringLiteral(string(caesura.Value))))
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
			elisionIdent := "__note" + elision.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Elision{Name: %s}).Stage(stageSet.Stage)", elisionIdent, __gong__toRawStringLiteral(elision.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elisionIdent, __gong__toRawStringLiteral(elision.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", elisionIdent, __gong__toRawStringLiteral(elision.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", elisionIdent, __gong__toRawStringLiteral(elision.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", elisionIdent, __gong__toRawStringLiteral(string(elision.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", elisionIdent, __gong__toRawStringLiteral(elision.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", elisionIdent, __gong__toRawStringLiteral(string(elision.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", elisionIdent, __gong__toRawStringLiteral(elision.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", elisionIdent, __gong__toRawStringLiteral(elision.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, emptyline := range __gong__sortStageSetInstances(stageSet.Stage.EmptyLines, stageSet.Stage.EmptyLine_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			emptylineIdent := "__note" + emptyline.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.EmptyLine{Name: %s}).Stage(stageSet.Stage)", emptylineIdent, __gong__toRawStringLiteral(emptyline.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", emptylineIdent, __gong__toRawStringLiteral(emptyline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.LineShape = %s", emptylineIdent, __gong__toRawStringLiteral(string(emptyline.LineShape))))
			values.WriteString(fmt.Sprintf("\n\t%s.LineType = %s", emptylineIdent, __gong__toRawStringLiteral(string(emptyline.LineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.LineLength = %s", emptylineIdent, __gong__toRawStringLiteral(string(emptyline.LineLength))))
			values.WriteString(fmt.Sprintf("\n\t%s.DashLength = %s", emptylineIdent, __gong__toRawStringLiteral(emptyline.DashLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.SpaceLength = %s", emptylineIdent, __gong__toRawStringLiteral(emptyline.SpaceLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", emptylineIdent, __gong__toRawStringLiteral(emptyline.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", emptylineIdent, __gong__toRawStringLiteral(emptyline.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", emptylineIdent, __gong__toRawStringLiteral(emptyline.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", emptylineIdent, __gong__toRawStringLiteral(emptyline.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", emptylineIdent, __gong__toRawStringLiteral(emptyline.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", emptylineIdent, __gong__toRawStringLiteral(string(emptyline.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", emptylineIdent, __gong__toRawStringLiteral(emptyline.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", emptylineIdent, __gong__toRawStringLiteral(string(emptyline.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", emptylineIdent, __gong__toRawStringLiteral(emptyline.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", emptylineIdent, __gong__toRawStringLiteral(string(emptyline.Placement))))
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
			extendIdent := "__note" + extend.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Extend{Name: %s}).Stage(stageSet.Stage)", extendIdent, __gong__toRawStringLiteral(extend.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", extendIdent, __gong__toRawStringLiteral(extend.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", extendIdent, __gong__toRawStringLiteral(string(extend.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", extendIdent, __gong__toRawStringLiteral(extend.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", extendIdent, __gong__toRawStringLiteral(extend.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", extendIdent, __gong__toRawStringLiteral(extend.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", extendIdent, __gong__toRawStringLiteral(extend.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", extendIdent, __gong__toRawStringLiteral(extend.Color)))
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
			figureIdent := "__note" + figure.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Figure{Name: %s}).Stage(stageSet.Stage)", figureIdent, __gong__toRawStringLiteral(figure.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", figureIdent, __gong__toRawStringLiteral(figure.Name)))
			if figure.Extend != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + figure.Extend.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extend = %s", figureIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, figuredbass := range __gong__sortStageSetInstances(stageSet.Stage.FiguredBasss, stageSet.Stage.FiguredBass_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			figuredbassIdent := "__note" + figuredbass.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.FiguredBass{Name: %s}).Stage(stageSet.Stage)", figuredbassIdent, __gong__toRawStringLiteral(figuredbass.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", figuredbassIdent, __gong__toRawStringLiteral(figuredbass.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", figuredbassIdent, __gong__toRawStringLiteral(string(figuredbass.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", figuredbassIdent, __gong__toRawStringLiteral(figuredbass.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", figuredbassIdent, __gong__toRawStringLiteral(figuredbass.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", figuredbassIdent, __gong__toRawStringLiteral(figuredbass.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", figuredbassIdent, __gong__toRawStringLiteral(figuredbass.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", figuredbassIdent, __gong__toRawStringLiteral(figuredbass.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", figuredbassIdent, __gong__toRawStringLiteral(string(figuredbass.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", figuredbassIdent, __gong__toRawStringLiteral(figuredbass.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", figuredbassIdent, __gong__toRawStringLiteral(string(figuredbass.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", figuredbassIdent, __gong__toRawStringLiteral(figuredbass.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", figuredbassIdent, __gong__toRawStringLiteral(string(figuredbass.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", figuredbassIdent, __gong__toRawStringLiteral(string(figuredbass.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", figuredbassIdent, __gong__toRawStringLiteral(string(figuredbass.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintDot = %s", figuredbassIdent, __gong__toRawStringLiteral(string(figuredbass.PrintDot))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintLyric = %s", figuredbassIdent, __gong__toRawStringLiteral(string(figuredbass.PrintLyric))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", figuredbassIdent, __gong__toRawStringLiteral(string(figuredbass.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintSpacing = %s", figuredbassIdent, __gong__toRawStringLiteral(string(figuredbass.PrintSpacing))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", figuredbassIdent, __gong__toRawStringLiteral(figuredbass.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = %s", figuredbassIdent, __gong__toRawStringLiteral(figuredbass.Duration)))
			for _, elem := range figuredbass.Figure {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Figure = append(%s.Figure, %s)", figuredbassIdent, figuredbassIdent, targetIdent))
			}
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
			forwardIdent := "__note" + forward.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Forward{Name: %s}).Stage(stageSet.Stage)", forwardIdent, __gong__toRawStringLiteral(forward.Name)))
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
			glissandoIdent := "__note" + glissando.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Glissando{Name: %s}).Stage(stageSet.Stage)", glissandoIdent, __gong__toRawStringLiteral(glissando.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", glissandoIdent, __gong__toRawStringLiteral(string(glissando.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", glissandoIdent, glissando.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.LineType = %s", glissandoIdent, __gong__toRawStringLiteral(string(glissando.LineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.DashLength = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.DashLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.SpaceLength = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.SpaceLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", glissandoIdent, __gong__toRawStringLiteral(string(glissando.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", glissandoIdent, __gong__toRawStringLiteral(string(glissando.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", glissandoIdent, __gong__toRawStringLiteral(glissando.Value)))
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
			graceIdent := "__note" + grace.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Grace{Name: %s}).Stage(stageSet.Stage)", graceIdent, __gong__toRawStringLiteral(grace.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", graceIdent, __gong__toRawStringLiteral(grace.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StealTimePrevious = %s", graceIdent, __gong__toRawStringLiteral(grace.StealTimePrevious)))
			values.WriteString(fmt.Sprintf("\n\t%s.StealTimeFollowing = %s", graceIdent, __gong__toRawStringLiteral(grace.StealTimeFollowing)))
			values.WriteString(fmt.Sprintf("\n\t%s.MakeTime = %s", graceIdent, __gong__toRawStringLiteral(grace.MakeTime)))
			values.WriteString(fmt.Sprintf("\n\t%s.Slash = %s", graceIdent, __gong__toRawStringLiteral(string(grace.Slash))))
		}
	}
	if stageSet.Stage != nil {
		for _, hammeronpulloff := range __gong__sortStageSetInstances(stageSet.Stage.HammerOnPullOffs, stageSet.Stage.HammerOnPullOff_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			hammeronpulloffIdent := "__note" + hammeronpulloff.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.HammerOnPullOff{Name: %s}).Stage(stageSet.Stage)", hammeronpulloffIdent, __gong__toRawStringLiteral(hammeronpulloff.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", hammeronpulloffIdent, __gong__toRawStringLiteral(hammeronpulloff.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", hammeronpulloffIdent, __gong__toRawStringLiteral(string(hammeronpulloff.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", hammeronpulloffIdent, hammeronpulloff.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", hammeronpulloffIdent, __gong__toRawStringLiteral(hammeronpulloff.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", hammeronpulloffIdent, __gong__toRawStringLiteral(hammeronpulloff.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", hammeronpulloffIdent, __gong__toRawStringLiteral(hammeronpulloff.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", hammeronpulloffIdent, __gong__toRawStringLiteral(hammeronpulloff.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", hammeronpulloffIdent, __gong__toRawStringLiteral(hammeronpulloff.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", hammeronpulloffIdent, __gong__toRawStringLiteral(string(hammeronpulloff.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", hammeronpulloffIdent, __gong__toRawStringLiteral(hammeronpulloff.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", hammeronpulloffIdent, __gong__toRawStringLiteral(string(hammeronpulloff.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", hammeronpulloffIdent, __gong__toRawStringLiteral(hammeronpulloff.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", hammeronpulloffIdent, __gong__toRawStringLiteral(string(hammeronpulloff.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", hammeronpulloffIdent, __gong__toRawStringLiteral(hammeronpulloff.Value)))
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
			handbellIdent := "__note" + handbell.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Handbell{Name: %s}).Stage(stageSet.Stage)", handbellIdent, __gong__toRawStringLiteral(handbell.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", handbellIdent, __gong__toRawStringLiteral(handbell.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", handbellIdent, __gong__toRawStringLiteral(handbell.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", handbellIdent, __gong__toRawStringLiteral(handbell.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", handbellIdent, __gong__toRawStringLiteral(handbell.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", handbellIdent, __gong__toRawStringLiteral(handbell.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", handbellIdent, __gong__toRawStringLiteral(handbell.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", handbellIdent, __gong__toRawStringLiteral(string(handbell.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", handbellIdent, __gong__toRawStringLiteral(handbell.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", handbellIdent, __gong__toRawStringLiteral(string(handbell.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", handbellIdent, __gong__toRawStringLiteral(handbell.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", handbellIdent, __gong__toRawStringLiteral(string(handbell.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", handbellIdent, __gong__toRawStringLiteral(string(handbell.Value))))
		}
	}
	if stageSet.Stage != nil {
		for _, harmonclosed := range __gong__sortStageSetInstances(stageSet.Stage.HarmonCloseds, stageSet.Stage.HarmonClosed_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			harmonclosedIdent := "__note" + harmonclosed.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.HarmonClosed{Name: %s}).Stage(stageSet.Stage)", harmonclosedIdent, __gong__toRawStringLiteral(harmonclosed.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harmonclosedIdent, __gong__toRawStringLiteral(harmonclosed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", harmonclosedIdent, __gong__toRawStringLiteral(string(harmonclosed.Location))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", harmonclosedIdent, __gong__toRawStringLiteral(string(harmonclosed.Value))))
		}
	}
	if stageSet.Stage != nil {
		for _, harmonmute := range __gong__sortStageSetInstances(stageSet.Stage.HarmonMutes, stageSet.Stage.HarmonMute_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			harmonmuteIdent := "__note" + harmonmute.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.HarmonMute{Name: %s}).Stage(stageSet.Stage)", harmonmuteIdent, __gong__toRawStringLiteral(harmonmute.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harmonmuteIdent, __gong__toRawStringLiteral(harmonmute.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", harmonmuteIdent, __gong__toRawStringLiteral(harmonmute.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", harmonmuteIdent, __gong__toRawStringLiteral(harmonmute.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", harmonmuteIdent, __gong__toRawStringLiteral(harmonmute.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", harmonmuteIdent, __gong__toRawStringLiteral(harmonmute.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", harmonmuteIdent, __gong__toRawStringLiteral(harmonmute.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", harmonmuteIdent, __gong__toRawStringLiteral(string(harmonmute.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", harmonmuteIdent, __gong__toRawStringLiteral(harmonmute.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", harmonmuteIdent, __gong__toRawStringLiteral(string(harmonmute.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", harmonmuteIdent, __gong__toRawStringLiteral(harmonmute.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", harmonmuteIdent, __gong__toRawStringLiteral(string(harmonmute.Placement))))
			if harmonmute.HarmonClosed != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + harmonmute.HarmonClosed.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HarmonClosed = %s", harmonmuteIdent, targetIdent))
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
			harmonicIdent := "__note" + harmonic.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Harmonic{Name: %s}).Stage(stageSet.Stage)", harmonicIdent, __gong__toRawStringLiteral(harmonic.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", harmonicIdent, __gong__toRawStringLiteral(string(harmonic.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", harmonicIdent, __gong__toRawStringLiteral(string(harmonic.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", harmonicIdent, __gong__toRawStringLiteral(string(harmonic.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", harmonicIdent, __gong__toRawStringLiteral(harmonic.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", harmonicIdent, __gong__toRawStringLiteral(string(harmonic.Placement))))
		}
	}
	if stageSet.Stage != nil {
		for _, heeltoe := range __gong__sortStageSetInstances(stageSet.Stage.HeelToes, stageSet.Stage.HeelToe_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			heeltoeIdent := "__note" + heeltoe.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.HeelToe{Name: %s}).Stage(stageSet.Stage)", heeltoeIdent, __gong__toRawStringLiteral(heeltoe.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", heeltoeIdent, __gong__toRawStringLiteral(heeltoe.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", heeltoeIdent, __gong__toRawStringLiteral(heeltoe.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", heeltoeIdent, __gong__toRawStringLiteral(heeltoe.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", heeltoeIdent, __gong__toRawStringLiteral(heeltoe.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", heeltoeIdent, __gong__toRawStringLiteral(heeltoe.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", heeltoeIdent, __gong__toRawStringLiteral(heeltoe.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", heeltoeIdent, __gong__toRawStringLiteral(string(heeltoe.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", heeltoeIdent, __gong__toRawStringLiteral(heeltoe.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", heeltoeIdent, __gong__toRawStringLiteral(string(heeltoe.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", heeltoeIdent, __gong__toRawStringLiteral(heeltoe.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", heeltoeIdent, __gong__toRawStringLiteral(string(heeltoe.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Substitution = %s", heeltoeIdent, __gong__toRawStringLiteral(string(heeltoe.Substitution))))
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
			holeIdent := "__note" + hole.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Hole{Name: %s}).Stage(stageSet.Stage)", holeIdent, __gong__toRawStringLiteral(hole.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holeIdent, __gong__toRawStringLiteral(hole.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", holeIdent, __gong__toRawStringLiteral(hole.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", holeIdent, __gong__toRawStringLiteral(hole.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", holeIdent, __gong__toRawStringLiteral(hole.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", holeIdent, __gong__toRawStringLiteral(hole.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", holeIdent, __gong__toRawStringLiteral(hole.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", holeIdent, __gong__toRawStringLiteral(string(hole.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", holeIdent, __gong__toRawStringLiteral(hole.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", holeIdent, __gong__toRawStringLiteral(string(hole.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", holeIdent, __gong__toRawStringLiteral(hole.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", holeIdent, __gong__toRawStringLiteral(string(hole.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.HoleType = %s", holeIdent, __gong__toRawStringLiteral(hole.HoleType)))
			values.WriteString(fmt.Sprintf("\n\t%s.HoleShape = %s", holeIdent, __gong__toRawStringLiteral(hole.HoleShape)))
			if hole.HoleClosed != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + hole.HoleClosed.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HoleClosed = %s", holeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, holeclosed := range __gong__sortStageSetInstances(stageSet.Stage.HoleCloseds, stageSet.Stage.HoleClosed_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			holeclosedIdent := "__note" + holeclosed.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.HoleClosed{Name: %s}).Stage(stageSet.Stage)", holeclosedIdent, __gong__toRawStringLiteral(holeclosed.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holeclosedIdent, __gong__toRawStringLiteral(holeclosed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", holeclosedIdent, __gong__toRawStringLiteral(string(holeclosed.Location))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", holeclosedIdent, __gong__toRawStringLiteral(string(holeclosed.Value))))
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
			instrumentIdent := "__note" + instrument.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Instrument{Name: %s}).Stage(stageSet.Stage)", instrumentIdent, __gong__toRawStringLiteral(instrument.Name)))
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
		for _, listen := range __gong__sortStageSetInstances(stageSet.Stage.Listens, stageSet.Stage.Listen_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			listenIdent := "__note" + listen.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Listen{Name: %s}).Stage(stageSet.Stage)", listenIdent, __gong__toRawStringLiteral(listen.Name)))
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
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Assess = append(%s.Assess, %s)", listenIdent, listenIdent, targetIdent))
			}
			for _, elem := range listen.Wait {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wait = append(%s.Wait, %s)", listenIdent, listenIdent, targetIdent))
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
			lyricIdent := "__note" + lyric.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Lyric{Name: %s}).Stage(stageSet.Stage)", lyricIdent, __gong__toRawStringLiteral(lyric.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameAttr = %s", lyricIdent, __gong__toRawStringLiteral(lyric.NameAttr)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeOnly = %s", lyricIdent, __gong__toRawStringLiteral(lyric.TimeOnly)))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", lyricIdent, __gong__toRawStringLiteral(string(lyric.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", lyricIdent, __gong__toRawStringLiteral(lyric.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", lyricIdent, __gong__toRawStringLiteral(lyric.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", lyricIdent, __gong__toRawStringLiteral(lyric.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", lyricIdent, __gong__toRawStringLiteral(lyric.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", lyricIdent, __gong__toRawStringLiteral(string(lyric.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", lyricIdent, __gong__toRawStringLiteral(string(lyric.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", lyricIdent, __gong__toRawStringLiteral(lyric.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Syllabic = %s", lyricIdent, __gong__toRawStringLiteral(string(lyric.Syllabic))))
			if lyric.Text != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + lyric.Text.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Text = %s", lyricIdent, targetIdent))
			}
			for _, elem := range lyric.Elision {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elision = append(%s.Elision, %s)", lyricIdent, lyricIdent, targetIdent))
			}
			if lyric.Extend != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + lyric.Extend.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extend = %s", lyricIdent, targetIdent))
			}
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
			mordentIdent := "__note" + mordent.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Mordent{Name: %s}).Stage(stageSet.Stage)", mordentIdent, __gong__toRawStringLiteral(mordent.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", mordentIdent, __gong__toRawStringLiteral(mordent.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", mordentIdent, __gong__toRawStringLiteral(mordent.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", mordentIdent, __gong__toRawStringLiteral(mordent.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", mordentIdent, __gong__toRawStringLiteral(mordent.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", mordentIdent, __gong__toRawStringLiteral(mordent.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", mordentIdent, __gong__toRawStringLiteral(mordent.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", mordentIdent, __gong__toRawStringLiteral(string(mordent.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", mordentIdent, __gong__toRawStringLiteral(mordent.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", mordentIdent, __gong__toRawStringLiteral(string(mordent.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", mordentIdent, __gong__toRawStringLiteral(mordent.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", mordentIdent, __gong__toRawStringLiteral(string(mordent.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.StartNote = %s", mordentIdent, __gong__toRawStringLiteral(string(mordent.StartNote))))
			values.WriteString(fmt.Sprintf("\n\t%s.TrillStep = %s", mordentIdent, __gong__toRawStringLiteral(string(mordent.TrillStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.TwoNoteTurn = %s", mordentIdent, __gong__toRawStringLiteral(string(mordent.TwoNoteTurn))))
			values.WriteString(fmt.Sprintf("\n\t%s.Accelerate = %s", mordentIdent, __gong__toRawStringLiteral(string(mordent.Accelerate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", mordentIdent, __gong__toRawStringLiteral(mordent.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.SecondBeat = %s", mordentIdent, __gong__toRawStringLiteral(mordent.SecondBeat)))
			values.WriteString(fmt.Sprintf("\n\t%s.LastBeat = %s", mordentIdent, __gong__toRawStringLiteral(mordent.LastBeat)))
			values.WriteString(fmt.Sprintf("\n\t%s.Long = %s", mordentIdent, __gong__toRawStringLiteral(string(mordent.Long))))
			values.WriteString(fmt.Sprintf("\n\t%s.Approach = %s", mordentIdent, __gong__toRawStringLiteral(string(mordent.Approach))))
			values.WriteString(fmt.Sprintf("\n\t%s.Departure = %s", mordentIdent, __gong__toRawStringLiteral(string(mordent.Departure))))
		}
	}
	if stageSet.Stage != nil {
		for _, nonarpeggiate := range __gong__sortStageSetInstances(stageSet.Stage.NonArpeggiates, stageSet.Stage.NonArpeggiate_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			nonarpeggiateIdent := "__note" + nonarpeggiate.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.NonArpeggiate{Name: %s}).Stage(stageSet.Stage)", nonarpeggiateIdent, __gong__toRawStringLiteral(nonarpeggiate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nonarpeggiateIdent, __gong__toRawStringLiteral(nonarpeggiate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", nonarpeggiateIdent, __gong__toRawStringLiteral(string(nonarpeggiate.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", nonarpeggiateIdent, nonarpeggiate.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", nonarpeggiateIdent, __gong__toRawStringLiteral(nonarpeggiate.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", nonarpeggiateIdent, __gong__toRawStringLiteral(nonarpeggiate.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", nonarpeggiateIdent, __gong__toRawStringLiteral(nonarpeggiate.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", nonarpeggiateIdent, __gong__toRawStringLiteral(nonarpeggiate.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", nonarpeggiateIdent, __gong__toRawStringLiteral(string(nonarpeggiate.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", nonarpeggiateIdent, __gong__toRawStringLiteral(nonarpeggiate.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", nonarpeggiateIdent, __gong__toRawStringLiteral(nonarpeggiate.Id)))
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
			notationsIdent := "__note" + notations.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Notations{Name: %s}).Stage(stageSet.Stage)", notationsIdent, __gong__toRawStringLiteral(notations.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notationsIdent, __gong__toRawStringLiteral(notations.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", notationsIdent, __gong__toRawStringLiteral(string(notations.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", notationsIdent, __gong__toRawStringLiteral(notations.Id)))
			for _, elem := range notations.Tied {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tied = append(%s.Tied, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Slur {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slur = append(%s.Slur, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Tuplet {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet = append(%s.Tuplet, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Glissando {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Glissando = append(%s.Glissando, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Slide {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slide = append(%s.Slide, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Ornaments {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ornaments = append(%s.Ornaments, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Technical {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Technical = append(%s.Technical, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Articulations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Articulations = append(%s.Articulations, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Arpeggiate {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arpeggiate = append(%s.Arpeggiate, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.NonArpeggiate {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NonArpeggiate = append(%s.NonArpeggiate, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.AccidentalMark {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AccidentalMark = append(%s.AccidentalMark, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.OtherNotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherNotation = append(%s.OtherNotation, %s)", notationsIdent, notationsIdent, targetIdent))
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
			noteIdent := "__note" + note.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Note{Name: %s}).Stage(stageSet.Stage)", noteIdent, __gong__toRawStringLiteral(note.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteIdent, __gong__toRawStringLiteral(note.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintLeger = %s", noteIdent, __gong__toRawStringLiteral(string(note.PrintLeger))))
			values.WriteString(fmt.Sprintf("\n\t%s.Dynamics = %s", noteIdent, __gong__toRawStringLiteral(note.Dynamics)))
			values.WriteString(fmt.Sprintf("\n\t%s.EndDynamics = %s", noteIdent, __gong__toRawStringLiteral(note.EndDynamics)))
			values.WriteString(fmt.Sprintf("\n\t%s.Attack = %s", noteIdent, __gong__toRawStringLiteral(note.Attack)))
			values.WriteString(fmt.Sprintf("\n\t%s.Release = %s", noteIdent, __gong__toRawStringLiteral(note.Release)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeOnly = %s", noteIdent, __gong__toRawStringLiteral(note.TimeOnly)))
			values.WriteString(fmt.Sprintf("\n\t%s.Pizzicato = %s", noteIdent, __gong__toRawStringLiteral(string(note.Pizzicato))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", noteIdent, __gong__toRawStringLiteral(note.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", noteIdent, __gong__toRawStringLiteral(note.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", noteIdent, __gong__toRawStringLiteral(note.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", noteIdent, __gong__toRawStringLiteral(note.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", noteIdent, __gong__toRawStringLiteral(note.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", noteIdent, __gong__toRawStringLiteral(string(note.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", noteIdent, __gong__toRawStringLiteral(note.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", noteIdent, __gong__toRawStringLiteral(string(note.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", noteIdent, __gong__toRawStringLiteral(note.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintDot = %s", noteIdent, __gong__toRawStringLiteral(string(note.PrintDot))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintLyric = %s", noteIdent, __gong__toRawStringLiteral(string(note.PrintLyric))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", noteIdent, __gong__toRawStringLiteral(string(note.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintSpacing = %s", noteIdent, __gong__toRawStringLiteral(string(note.PrintSpacing))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", noteIdent, __gong__toRawStringLiteral(note.Id)))
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
				targetIdent := "__note" + note.Grace.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Grace = %s", noteIdent, targetIdent))
			}
			if note.Pitch != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + note.Pitch.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pitch = %s", noteIdent, targetIdent))
			}
			if note.Unpitched != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + note.Unpitched.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Unpitched = %s", noteIdent, targetIdent))
			}
			if note.Rest != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + note.Rest.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rest = %s", noteIdent, targetIdent))
			}
			for _, elem := range note.Tie {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tie = append(%s.Tie, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Instrument {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Instrument = append(%s.Instrument, %s)", noteIdent, noteIdent, targetIdent))
			}
			if note.Type != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + note.Type.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type = %s", noteIdent, targetIdent))
			}
			if note.Accidental != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + note.Accidental.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accidental = %s", noteIdent, targetIdent))
			}
			if note.Stem != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + note.Stem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stem = %s", noteIdent, targetIdent))
			}
			if note.Notehead != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + note.Notehead.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notehead = %s", noteIdent, targetIdent))
			}
			if note.NoteheadText != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + note.NoteheadText.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteheadText = %s", noteIdent, targetIdent))
			}
			for _, elem := range note.Beam {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Beam = append(%s.Beam, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Notations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notations = append(%s.Notations, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Lyric {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lyric = append(%s.Lyric, %s)", noteIdent, noteIdent, targetIdent))
			}
			if note.Listen != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + note.Listen.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listen = %s", noteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, notetype := range __gong__sortStageSetInstances(stageSet.Stage.NoteTypes, stageSet.Stage.NoteType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notetypeIdent := "__note" + notetype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.NoteType{Name: %s}).Stage(stageSet.Stage)", notetypeIdent, __gong__toRawStringLiteral(notetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notetypeIdent, __gong__toRawStringLiteral(notetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %s", notetypeIdent, __gong__toRawStringLiteral(string(notetype.Size))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", notetypeIdent, __gong__toRawStringLiteral(string(notetype.Value))))
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
			noteheadIdent := "__note" + notehead.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Notehead{Name: %s}).Stage(stageSet.Stage)", noteheadIdent, __gong__toRawStringLiteral(notehead.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteheadIdent, __gong__toRawStringLiteral(notehead.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Filled = %s", noteheadIdent, __gong__toRawStringLiteral(string(notehead.Filled))))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", noteheadIdent, __gong__toRawStringLiteral(string(notehead.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", noteheadIdent, __gong__toRawStringLiteral(notehead.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", noteheadIdent, __gong__toRawStringLiteral(string(notehead.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", noteheadIdent, __gong__toRawStringLiteral(notehead.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", noteheadIdent, __gong__toRawStringLiteral(string(notehead.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", noteheadIdent, __gong__toRawStringLiteral(notehead.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", noteheadIdent, __gong__toRawStringLiteral(notehead.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", noteheadIdent, __gong__toRawStringLiteral(string(notehead.Value))))
		}
	}
	if stageSet.Stage != nil {
		for _, noteheadtext := range __gong__sortStageSetInstances(stageSet.Stage.NoteheadTexts, stageSet.Stage.NoteheadText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			noteheadtextIdent := "__note" + noteheadtext.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.NoteheadText{Name: %s}).Stage(stageSet.Stage)", noteheadtextIdent, __gong__toRawStringLiteral(noteheadtext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteheadtextIdent, __gong__toRawStringLiteral(noteheadtext.Name)))
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
			ornamentsIdent := "__note" + ornaments.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Ornaments{Name: %s}).Stage(stageSet.Stage)", ornamentsIdent, __gong__toRawStringLiteral(ornaments.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ornamentsIdent, __gong__toRawStringLiteral(ornaments.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ornamentsIdent, __gong__toRawStringLiteral(ornaments.Id)))
			for _, elem := range ornaments.Mordent {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mordent = append(%s.Mordent, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.InvertedMordent {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InvertedMordent = append(%s.InvertedMordent, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Tremolo {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tremolo = append(%s.Tremolo, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.OtherOrnament {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherOrnament = append(%s.OtherOrnament, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.AccidentalMark {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AccidentalMark = append(%s.AccidentalMark, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, othernotation := range __gong__sortStageSetInstances(stageSet.Stage.OtherNotations, stageSet.Stage.OtherNotation_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			othernotationIdent := "__note" + othernotation.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.OtherNotation{Name: %s}).Stage(stageSet.Stage)", othernotationIdent, __gong__toRawStringLiteral(othernotation.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", othernotationIdent, __gong__toRawStringLiteral(othernotation.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", othernotationIdent, __gong__toRawStringLiteral(string(othernotation.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", othernotationIdent, othernotation.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", othernotationIdent, __gong__toRawStringLiteral(string(othernotation.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", othernotationIdent, __gong__toRawStringLiteral(othernotation.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", othernotationIdent, __gong__toRawStringLiteral(othernotation.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", othernotationIdent, __gong__toRawStringLiteral(othernotation.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", othernotationIdent, __gong__toRawStringLiteral(othernotation.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", othernotationIdent, __gong__toRawStringLiteral(othernotation.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", othernotationIdent, __gong__toRawStringLiteral(string(othernotation.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", othernotationIdent, __gong__toRawStringLiteral(othernotation.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", othernotationIdent, __gong__toRawStringLiteral(string(othernotation.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", othernotationIdent, __gong__toRawStringLiteral(othernotation.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", othernotationIdent, __gong__toRawStringLiteral(string(othernotation.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", othernotationIdent, __gong__toRawStringLiteral(othernotation.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", othernotationIdent, __gong__toRawStringLiteral(othernotation.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", othernotationIdent, __gong__toRawStringLiteral(othernotation.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, otherplacementtext := range __gong__sortStageSetInstances(stageSet.Stage.OtherPlacementTexts, stageSet.Stage.OtherPlacementText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			otherplacementtextIdent := "__note" + otherplacementtext.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.OtherPlacementText{Name: %s}).Stage(stageSet.Stage)", otherplacementtextIdent, __gong__toRawStringLiteral(otherplacementtext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", otherplacementtextIdent, __gong__toRawStringLiteral(otherplacementtext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", otherplacementtextIdent, __gong__toRawStringLiteral(otherplacementtext.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", otherplacementtextIdent, __gong__toRawStringLiteral(otherplacementtext.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", otherplacementtextIdent, __gong__toRawStringLiteral(otherplacementtext.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", otherplacementtextIdent, __gong__toRawStringLiteral(otherplacementtext.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", otherplacementtextIdent, __gong__toRawStringLiteral(otherplacementtext.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", otherplacementtextIdent, __gong__toRawStringLiteral(string(otherplacementtext.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", otherplacementtextIdent, __gong__toRawStringLiteral(otherplacementtext.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", otherplacementtextIdent, __gong__toRawStringLiteral(string(otherplacementtext.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", otherplacementtextIdent, __gong__toRawStringLiteral(otherplacementtext.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", otherplacementtextIdent, __gong__toRawStringLiteral(string(otherplacementtext.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", otherplacementtextIdent, __gong__toRawStringLiteral(otherplacementtext.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", otherplacementtextIdent, __gong__toRawStringLiteral(otherplacementtext.Value)))
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
			pitchIdent := "__note" + pitch.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Pitch{Name: %s}).Stage(stageSet.Stage)", pitchIdent, __gong__toRawStringLiteral(pitch.Name)))
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
		for _, placementtext := range __gong__sortStageSetInstances(stageSet.Stage.PlacementTexts, stageSet.Stage.PlacementText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			placementtextIdent := "__note" + placementtext.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.PlacementText{Name: %s}).Stage(stageSet.Stage)", placementtextIdent, __gong__toRawStringLiteral(placementtext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", placementtextIdent, __gong__toRawStringLiteral(placementtext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", placementtextIdent, __gong__toRawStringLiteral(placementtext.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", placementtextIdent, __gong__toRawStringLiteral(placementtext.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", placementtextIdent, __gong__toRawStringLiteral(placementtext.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", placementtextIdent, __gong__toRawStringLiteral(placementtext.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", placementtextIdent, __gong__toRawStringLiteral(placementtext.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", placementtextIdent, __gong__toRawStringLiteral(string(placementtext.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", placementtextIdent, __gong__toRawStringLiteral(placementtext.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", placementtextIdent, __gong__toRawStringLiteral(string(placementtext.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", placementtextIdent, __gong__toRawStringLiteral(placementtext.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", placementtextIdent, __gong__toRawStringLiteral(string(placementtext.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", placementtextIdent, __gong__toRawStringLiteral(placementtext.Value)))
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
			releaseIdent := "__note" + release.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Release{Name: %s}).Stage(stageSet.Stage)", releaseIdent, __gong__toRawStringLiteral(release.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", releaseIdent, __gong__toRawStringLiteral(release.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", releaseIdent, __gong__toRawStringLiteral(release.Offset)))
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
			restIdent := "__note" + rest.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Rest{Name: %s}).Stage(stageSet.Stage)", restIdent, __gong__toRawStringLiteral(rest.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", restIdent, __gong__toRawStringLiteral(rest.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Measure = %s", restIdent, __gong__toRawStringLiteral(string(rest.Measure))))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayStep = %s", restIdent, __gong__toRawStringLiteral(string(rest.DisplayStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayOctave = %d", restIdent, rest.DisplayOctave))
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
			slideIdent := "__note" + slide.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Slide{Name: %s}).Stage(stageSet.Stage)", slideIdent, __gong__toRawStringLiteral(slide.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", slideIdent, __gong__toRawStringLiteral(slide.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", slideIdent, __gong__toRawStringLiteral(string(slide.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", slideIdent, slide.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.LineType = %s", slideIdent, __gong__toRawStringLiteral(string(slide.LineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.DashLength = %s", slideIdent, __gong__toRawStringLiteral(slide.DashLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.SpaceLength = %s", slideIdent, __gong__toRawStringLiteral(slide.SpaceLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", slideIdent, __gong__toRawStringLiteral(slide.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", slideIdent, __gong__toRawStringLiteral(slide.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", slideIdent, __gong__toRawStringLiteral(slide.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", slideIdent, __gong__toRawStringLiteral(slide.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", slideIdent, __gong__toRawStringLiteral(slide.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", slideIdent, __gong__toRawStringLiteral(string(slide.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", slideIdent, __gong__toRawStringLiteral(slide.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", slideIdent, __gong__toRawStringLiteral(string(slide.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", slideIdent, __gong__toRawStringLiteral(slide.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Accelerate = %s", slideIdent, __gong__toRawStringLiteral(string(slide.Accelerate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", slideIdent, __gong__toRawStringLiteral(slide.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.FirstBeat = %s", slideIdent, __gong__toRawStringLiteral(slide.FirstBeat)))
			values.WriteString(fmt.Sprintf("\n\t%s.LastBeat = %s", slideIdent, __gong__toRawStringLiteral(slide.LastBeat)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", slideIdent, __gong__toRawStringLiteral(slide.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", slideIdent, __gong__toRawStringLiteral(slide.Value)))
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
			slurIdent := "__note" + slur.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Slur{Name: %s}).Stage(stageSet.Stage)", slurIdent, __gong__toRawStringLiteral(slur.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", slurIdent, __gong__toRawStringLiteral(slur.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", slurIdent, __gong__toRawStringLiteral(string(slur.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", slurIdent, slur.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.LineType = %s", slurIdent, __gong__toRawStringLiteral(string(slur.LineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.DashLength = %s", slurIdent, __gong__toRawStringLiteral(slur.DashLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.SpaceLength = %s", slurIdent, __gong__toRawStringLiteral(slur.SpaceLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", slurIdent, __gong__toRawStringLiteral(slur.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", slurIdent, __gong__toRawStringLiteral(slur.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", slurIdent, __gong__toRawStringLiteral(slur.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", slurIdent, __gong__toRawStringLiteral(slur.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", slurIdent, __gong__toRawStringLiteral(string(slur.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Orientation = %s", slurIdent, __gong__toRawStringLiteral(string(slur.Orientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierX = %s", slurIdent, __gong__toRawStringLiteral(slur.BezierX)))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierY = %s", slurIdent, __gong__toRawStringLiteral(slur.BezierY)))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierX2 = %s", slurIdent, __gong__toRawStringLiteral(slur.BezierX2)))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierY2 = %s", slurIdent, __gong__toRawStringLiteral(slur.BezierY2)))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierOffset = %s", slurIdent, __gong__toRawStringLiteral(slur.BezierOffset)))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierOffset2 = %s", slurIdent, __gong__toRawStringLiteral(slur.BezierOffset2)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", slurIdent, __gong__toRawStringLiteral(slur.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", slurIdent, __gong__toRawStringLiteral(slur.Id)))
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
			stemIdent := "__note" + stem.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Stem{Name: %s}).Stage(stageSet.Stage)", stemIdent, __gong__toRawStringLiteral(stem.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stemIdent, __gong__toRawStringLiteral(stem.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", stemIdent, __gong__toRawStringLiteral(stem.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", stemIdent, __gong__toRawStringLiteral(stem.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", stemIdent, __gong__toRawStringLiteral(stem.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", stemIdent, __gong__toRawStringLiteral(stem.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", stemIdent, __gong__toRawStringLiteral(stem.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", stemIdent, __gong__toRawStringLiteral(string(stem.Value))))
		}
	}
	if stageSet.Stage != nil {
		for _, strongaccent := range __gong__sortStageSetInstances(stageSet.Stage.StrongAccents, stageSet.Stage.StrongAccent_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			strongaccentIdent := "__note" + strongaccent.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.StrongAccent{Name: %s}).Stage(stageSet.Stage)", strongaccentIdent, __gong__toRawStringLiteral(strongaccent.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", strongaccentIdent, __gong__toRawStringLiteral(strongaccent.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", strongaccentIdent, __gong__toRawStringLiteral(strongaccent.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", strongaccentIdent, __gong__toRawStringLiteral(strongaccent.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", strongaccentIdent, __gong__toRawStringLiteral(strongaccent.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", strongaccentIdent, __gong__toRawStringLiteral(strongaccent.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", strongaccentIdent, __gong__toRawStringLiteral(strongaccent.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", strongaccentIdent, __gong__toRawStringLiteral(string(strongaccent.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", strongaccentIdent, __gong__toRawStringLiteral(strongaccent.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", strongaccentIdent, __gong__toRawStringLiteral(string(strongaccent.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", strongaccentIdent, __gong__toRawStringLiteral(strongaccent.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", strongaccentIdent, __gong__toRawStringLiteral(string(strongaccent.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", strongaccentIdent, __gong__toRawStringLiteral(string(strongaccent.Type))))
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
			tapIdent := "__note" + tap.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Tap{Name: %s}).Stage(stageSet.Stage)", tapIdent, __gong__toRawStringLiteral(tap.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tapIdent, __gong__toRawStringLiteral(tap.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Hand = %s", tapIdent, __gong__toRawStringLiteral(string(tap.Hand))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", tapIdent, __gong__toRawStringLiteral(tap.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", tapIdent, __gong__toRawStringLiteral(tap.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", tapIdent, __gong__toRawStringLiteral(tap.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", tapIdent, __gong__toRawStringLiteral(tap.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", tapIdent, __gong__toRawStringLiteral(tap.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", tapIdent, __gong__toRawStringLiteral(string(tap.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", tapIdent, __gong__toRawStringLiteral(tap.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", tapIdent, __gong__toRawStringLiteral(string(tap.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tapIdent, __gong__toRawStringLiteral(tap.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", tapIdent, __gong__toRawStringLiteral(string(tap.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", tapIdent, __gong__toRawStringLiteral(tap.Value)))
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
			technicalIdent := "__note" + technical.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Technical{Name: %s}).Stage(stageSet.Stage)", technicalIdent, __gong__toRawStringLiteral(technical.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", technicalIdent, __gong__toRawStringLiteral(technical.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", technicalIdent, __gong__toRawStringLiteral(technical.Id)))
			for _, elem := range technical.Harmonic {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmonic = append(%s.Harmonic, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Pluck {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pluck = append(%s.Pluck, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.HammerOn {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HammerOn = append(%s.HammerOn, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.PullOff {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PullOff = append(%s.PullOff, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Bend {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bend = append(%s.Bend, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Tap {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tap = append(%s.Tap, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Heel {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Heel = append(%s.Heel, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Toe {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Toe = append(%s.Toe, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Hole {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Hole = append(%s.Hole, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Arrow {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arrow = append(%s.Arrow, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Handbell {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Handbell = append(%s.Handbell, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.HarmonMute {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HarmonMute = append(%s.HarmonMute, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.OtherTechnical {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherTechnical = append(%s.OtherTechnical, %s)", technicalIdent, technicalIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, textelementdata := range __gong__sortStageSetInstances(stageSet.Stage.TextElementDatas, stageSet.Stage.TextElementData_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textelementdataIdent := "__note" + textelementdata.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.TextElementData{Name: %s}).Stage(stageSet.Stage)", textelementdataIdent, __gong__toRawStringLiteral(textelementdata.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textelementdataIdent, __gong__toRawStringLiteral(textelementdata.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", textelementdataIdent, __gong__toRawStringLiteral(textelementdata.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", textelementdataIdent, __gong__toRawStringLiteral(textelementdata.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", textelementdataIdent, __gong__toRawStringLiteral(string(textelementdata.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", textelementdataIdent, __gong__toRawStringLiteral(textelementdata.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", textelementdataIdent, __gong__toRawStringLiteral(string(textelementdata.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", textelementdataIdent, __gong__toRawStringLiteral(textelementdata.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", textelementdataIdent, textelementdata.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", textelementdataIdent, textelementdata.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.LineThrough = %d", textelementdataIdent, textelementdata.LineThrough))
			values.WriteString(fmt.Sprintf("\n\t%s.Rotation = %s", textelementdataIdent, __gong__toRawStringLiteral(textelementdata.Rotation)))
			values.WriteString(fmt.Sprintf("\n\t%s.LetterSpacing = %s", textelementdataIdent, __gong__toRawStringLiteral(textelementdata.LetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dir = %s", textelementdataIdent, __gong__toRawStringLiteral(string(textelementdata.Dir))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", textelementdataIdent, __gong__toRawStringLiteral(textelementdata.Value)))
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
			tieIdent := "__note" + tie.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Tie{Name: %s}).Stage(stageSet.Stage)", tieIdent, __gong__toRawStringLiteral(tie.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tieIdent, __gong__toRawStringLiteral(tie.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", tieIdent, __gong__toRawStringLiteral(string(tie.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeOnly = %s", tieIdent, __gong__toRawStringLiteral(tie.TimeOnly)))
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
			tiedIdent := "__note" + tied.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Tied{Name: %s}).Stage(stageSet.Stage)", tiedIdent, __gong__toRawStringLiteral(tied.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tiedIdent, __gong__toRawStringLiteral(tied.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", tiedIdent, __gong__toRawStringLiteral(string(tied.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", tiedIdent, tied.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.LineType = %s", tiedIdent, __gong__toRawStringLiteral(string(tied.LineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.DashLength = %s", tiedIdent, __gong__toRawStringLiteral(tied.DashLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.SpaceLength = %s", tiedIdent, __gong__toRawStringLiteral(tied.SpaceLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", tiedIdent, __gong__toRawStringLiteral(tied.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", tiedIdent, __gong__toRawStringLiteral(tied.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", tiedIdent, __gong__toRawStringLiteral(tied.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", tiedIdent, __gong__toRawStringLiteral(tied.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", tiedIdent, __gong__toRawStringLiteral(string(tied.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Orientation = %s", tiedIdent, __gong__toRawStringLiteral(string(tied.Orientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierX = %s", tiedIdent, __gong__toRawStringLiteral(tied.BezierX)))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierY = %s", tiedIdent, __gong__toRawStringLiteral(tied.BezierY)))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierX2 = %s", tiedIdent, __gong__toRawStringLiteral(tied.BezierX2)))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierY2 = %s", tiedIdent, __gong__toRawStringLiteral(tied.BezierY2)))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierOffset = %s", tiedIdent, __gong__toRawStringLiteral(tied.BezierOffset)))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierOffset2 = %s", tiedIdent, __gong__toRawStringLiteral(tied.BezierOffset2)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tiedIdent, __gong__toRawStringLiteral(tied.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", tiedIdent, __gong__toRawStringLiteral(tied.Id)))
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
			tremoloIdent := "__note" + tremolo.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Tremolo{Name: %s}).Stage(stageSet.Stage)", tremoloIdent, __gong__toRawStringLiteral(tremolo.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", tremoloIdent, __gong__toRawStringLiteral(string(tremolo.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", tremoloIdent, __gong__toRawStringLiteral(string(tremolo.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", tremoloIdent, __gong__toRawStringLiteral(string(tremolo.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", tremoloIdent, __gong__toRawStringLiteral(string(tremolo.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", tremoloIdent, __gong__toRawStringLiteral(tremolo.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", tremoloIdent, tremolo.Value))
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
			tupletIdent := "__note" + tuplet.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Tuplet{Name: %s}).Stage(stageSet.Stage)", tupletIdent, __gong__toRawStringLiteral(tuplet.Name)))
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
			values.WriteString(fmt.Sprintf("\n\t%s.ShowNumber = %s", tupletIdent, __gong__toRawStringLiteral(string(tuplet.ShowNumber))))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowType = %s", tupletIdent, __gong__toRawStringLiteral(string(tuplet.ShowType))))
			values.WriteString(fmt.Sprintf("\n\t%s.LineShape = %s", tupletIdent, __gong__toRawStringLiteral(string(tuplet.LineShape))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", tupletIdent, __gong__toRawStringLiteral(string(tuplet.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", tupletIdent, __gong__toRawStringLiteral(tuplet.Id)))
			if tuplet.TupletActual != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + tuplet.TupletActual.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TupletActual = %s", tupletIdent, targetIdent))
			}
			if tuplet.TupletNormal != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + tuplet.TupletNormal.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TupletNormal = %s", tupletIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, tupletdot := range __gong__sortStageSetInstances(stageSet.Stage.TupletDots, stageSet.Stage.TupletDot_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tupletdotIdent := "__note" + tupletdot.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.TupletDot{Name: %s}).Stage(stageSet.Stage)", tupletdotIdent, __gong__toRawStringLiteral(tupletdot.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tupletdotIdent, __gong__toRawStringLiteral(tupletdot.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", tupletdotIdent, __gong__toRawStringLiteral(tupletdot.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", tupletdotIdent, __gong__toRawStringLiteral(string(tupletdot.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", tupletdotIdent, __gong__toRawStringLiteral(tupletdot.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", tupletdotIdent, __gong__toRawStringLiteral(string(tupletdot.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tupletdotIdent, __gong__toRawStringLiteral(tupletdot.Color)))
		}
	}
	if stageSet.Stage != nil {
		for _, tupletnumber := range __gong__sortStageSetInstances(stageSet.Stage.TupletNumbers, stageSet.Stage.TupletNumber_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tupletnumberIdent := "__note" + tupletnumber.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.TupletNumber{Name: %s}).Stage(stageSet.Stage)", tupletnumberIdent, __gong__toRawStringLiteral(tupletnumber.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tupletnumberIdent, __gong__toRawStringLiteral(tupletnumber.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", tupletnumberIdent, __gong__toRawStringLiteral(tupletnumber.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", tupletnumberIdent, __gong__toRawStringLiteral(string(tupletnumber.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", tupletnumberIdent, __gong__toRawStringLiteral(tupletnumber.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", tupletnumberIdent, __gong__toRawStringLiteral(string(tupletnumber.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tupletnumberIdent, __gong__toRawStringLiteral(tupletnumber.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", tupletnumberIdent, tupletnumber.Value))
		}
	}
	if stageSet.Stage != nil {
		for _, tupletportion := range __gong__sortStageSetInstances(stageSet.Stage.TupletPortions, stageSet.Stage.TupletPortion_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tupletportionIdent := "__note" + tupletportion.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.TupletPortion{Name: %s}).Stage(stageSet.Stage)", tupletportionIdent, __gong__toRawStringLiteral(tupletportion.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tupletportionIdent, __gong__toRawStringLiteral(tupletportion.Name)))
			if tupletportion.TupletNumber != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + tupletportion.TupletNumber.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TupletNumber = %s", tupletportionIdent, targetIdent))
			}
			if tupletportion.TupletType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + tupletportion.TupletType.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TupletType = %s", tupletportionIdent, targetIdent))
			}
			for _, elem := range tupletportion.TupletDot {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TupletDot = append(%s.TupletDot, %s)", tupletportionIdent, tupletportionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, tuplettype := range __gong__sortStageSetInstances(stageSet.Stage.TupletTypes, stageSet.Stage.TupletType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tuplettypeIdent := "__note" + tuplettype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.TupletType{Name: %s}).Stage(stageSet.Stage)", tuplettypeIdent, __gong__toRawStringLiteral(tuplettype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tuplettypeIdent, __gong__toRawStringLiteral(tuplettype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", tuplettypeIdent, __gong__toRawStringLiteral(tuplettype.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", tuplettypeIdent, __gong__toRawStringLiteral(string(tuplettype.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", tuplettypeIdent, __gong__toRawStringLiteral(tuplettype.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", tuplettypeIdent, __gong__toRawStringLiteral(string(tuplettype.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tuplettypeIdent, __gong__toRawStringLiteral(tuplettype.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", tuplettypeIdent, __gong__toRawStringLiteral(string(tuplettype.Value))))
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
			unpitchedIdent := "__note" + unpitched.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Unpitched{Name: %s}).Stage(stageSet.Stage)", unpitchedIdent, __gong__toRawStringLiteral(unpitched.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unpitchedIdent, __gong__toRawStringLiteral(unpitched.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayStep = %s", unpitchedIdent, __gong__toRawStringLiteral(string(unpitched.DisplayStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayOctave = %d", unpitchedIdent, unpitched.DisplayOctave))
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
			waitIdent := "__note" + wait.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Wait{Name: %s}).Stage(stageSet.Stage)", waitIdent, __gong__toRawStringLiteral(wait.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", waitIdent, __gong__toRawStringLiteral(wait.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", waitIdent, __gong__toRawStringLiteral(wait.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeOnly = %s", waitIdent, __gong__toRawStringLiteral(wait.TimeOnly)))
		}
	}


	var goModuleVersion string
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		goModuleVersion = buildInfo.Main.Version
	}

	goModuleVersionWithoutDirty := strings.ReplaceAll(goModuleVersion, "+dirty", "")
	if goModuleVersionWithoutDirty == "" || goModuleVersionWithoutDirty == "(devel)" {
		goModuleVersionWithoutDirty = "latest"
	}

	res = fmt.Sprintf(`// file generated by gong
// go module version: %s
/* to edit this file with the default probe, use the command
go run github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *note.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *note.StageSet) {

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
`, goModuleVersion, goModuleVersionWithoutDirty, packageName, declarations.String(), values.String(), pointers.String())

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
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note":
			aliasToCanonical[alias] = "note"
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
			case "note":
				switch typeName {
				case "Accidental":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Accidental), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "AccidentalMark":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AccidentalMark), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Arpeggiate":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Arpeggiate), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Arrow":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Arrow), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Articulations":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Articulations), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Assess":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Assess), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Backup":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Backup), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Beam":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Beam), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Bend":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bend), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "BreathMark":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BreathMark), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Caesura":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Caesura), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Elision":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Elision), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "EmptyLine":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EmptyLine), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Extend":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Extend), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Figure":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Figure), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "FiguredBass":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FiguredBass), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Forward":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Forward), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Glissando":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Glissando), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Grace":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Grace), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "HammerOnPullOff":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HammerOnPullOff), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Handbell":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Handbell), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "HarmonClosed":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HarmonClosed), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "HarmonMute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HarmonMute), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Harmonic":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harmonic), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "HeelToe":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HeelToe), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Hole":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Hole), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "HoleClosed":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HoleClosed), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Instrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Instrument), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Listen":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Listen), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Lyric":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Lyric), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Mordent":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Mordent), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NonArpeggiate":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NonArpeggiate), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Notations":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Notations), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Note), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteType), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Notehead":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Notehead), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteheadText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteheadText), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Ornaments":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Ornaments), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "OtherNotation":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OtherNotation), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "OtherPlacementText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OtherPlacementText), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Pitch":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Pitch), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PlacementText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PlacementText), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Release":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Release), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Rest":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Rest), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Slide":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Slide), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Slur":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Slur), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Stem":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Stem), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StrongAccent":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StrongAccent), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tap":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tap), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Technical":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Technical), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TextElementData":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextElementData), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tie":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tie), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tied":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tied), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tremolo":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tremolo), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tuplet":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tuplet), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TupletDot":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TupletDot), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TupletNumber":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TupletNumber), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TupletPortion":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TupletPortion), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TupletType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TupletType), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Unpitched":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Unpitched), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Wait":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Wait), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *Accidental:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Cautionary":
						inst.Cautionary = common.EnumYesNo(GongExtractString(rhs))
					case "Editorial":
						inst.Editorial = common.EnumYesNo(GongExtractString(rhs))
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Parentheses":
						inst.Parentheses = common.EnumYesNo(GongExtractString(rhs))
					case "Bracket":
						inst.Bracket = common.EnumYesNo(GongExtractString(rhs))
					case "Size":
						inst.Size = common.EnumSymbolSize(GongExtractString(rhs))
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Value":
						inst.Value = common.EnumAccidentalValue(GongExtractString(rhs))
					}
				case *AccidentalMark:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Parentheses":
						inst.Parentheses = common.EnumYesNo(GongExtractString(rhs))
					case "Bracket":
						inst.Bracket = common.EnumYesNo(GongExtractString(rhs))
					case "Size":
						inst.Size = common.EnumSymbolSize(GongExtractString(rhs))
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = common.EnumAccidentalValue(GongExtractString(rhs))
					}
				case *Arpeggiate:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Direction":
						inst.Direction = common.EnumUpDown(GongExtractString(rhs))
					case "Unbroken":
						inst.Unbroken = common.EnumYesNo(GongExtractString(rhs))
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Arrow:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "ArrowDirection":
						inst.ArrowDirection = EnumArrowDirection(GongExtractString(rhs))
					case "ArrowStyle":
						inst.ArrowStyle = EnumArrowStyle(GongExtractString(rhs))
					case "CircularArrow":
						inst.CircularArrow = EnumCircularArrow(GongExtractString(rhs))
					}
				case *Articulations:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "StrongAccent":
						__gong__assignSliceOfPointers(&inst.StrongAccent, rhs, identifierMap)
					case "Scoop":
						__gong__assignSliceOfPointers(&inst.Scoop, rhs, identifierMap)
					case "Plop":
						__gong__assignSliceOfPointers(&inst.Plop, rhs, identifierMap)
					case "Doit":
						__gong__assignSliceOfPointers(&inst.Doit, rhs, identifierMap)
					case "Falloff":
						__gong__assignSliceOfPointers(&inst.Falloff, rhs, identifierMap)
					case "BreathMark":
						__gong__assignSliceOfPointers(&inst.BreathMark, rhs, identifierMap)
					case "Caesura":
						__gong__assignSliceOfPointers(&inst.Caesura, rhs, identifierMap)
					case "OtherArticulation":
						__gong__assignSliceOfPointers(&inst.OtherArticulation, rhs, identifierMap)
					}
				case *Assess:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumYesNo(GongExtractString(rhs))
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "TimeOnly":
						inst.TimeOnly = GongExtractString(rhs)
					}
				case *Backup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					}
				case *Beam:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Repeater":
						inst.Repeater = common.EnumYesNo(GongExtractString(rhs))
					case "Fan":
						inst.Fan = EnumFan(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = direction.EnumBeamValue(GongExtractString(rhs))
					}
				case *Bend:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Shape":
						inst.Shape = EnumBendShape(GongExtractString(rhs))
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Accelerate":
						inst.Accelerate = common.EnumYesNo(GongExtractString(rhs))
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "FirstBeat":
						inst.FirstBeat = GongExtractString(rhs)
					case "LastBeat":
						inst.LastBeat = GongExtractString(rhs)
					case "BendAlter":
						inst.BendAlter = GongExtractString(rhs)
					case "Release":
						__gong__assignPointer(&inst.Release, rhs, identifierMap)
					case "WithBar":
						__gong__assignPointer(&inst.WithBar, rhs, identifierMap)
					}
				case *BreathMark:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Value":
						inst.Value = EnumBreathMarkValue(GongExtractString(rhs))
					}
				case *Caesura:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Value":
						inst.Value = EnumCaesuraValue(GongExtractString(rhs))
					}
				case *Elision:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *EmptyLine:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "LineShape":
						inst.LineShape = common.EnumLineShape(GongExtractString(rhs))
					case "LineType":
						inst.LineType = common.EnumLineType(GongExtractString(rhs))
					case "LineLength":
						inst.LineLength = common.EnumLineLength(GongExtractString(rhs))
					case "DashLength":
						inst.DashLength = GongExtractString(rhs)
					case "SpaceLength":
						inst.SpaceLength = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					}
				case *Extend:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStopContinue(GongExtractString(rhs))
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *Figure:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Extend":
						__gong__assignPointer(&inst.Extend, rhs, identifierMap)
					}
				case *FiguredBass:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Parentheses":
						inst.Parentheses = common.EnumYesNo(GongExtractString(rhs))
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = common.EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = common.EnumValign(GongExtractString(rhs))
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "PrintDot":
						inst.PrintDot = common.EnumYesNo(GongExtractString(rhs))
					case "PrintLyric":
						inst.PrintLyric = common.EnumYesNo(GongExtractString(rhs))
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "PrintSpacing":
						inst.PrintSpacing = common.EnumYesNo(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Figure":
						__gong__assignSliceOfPointers(&inst.Figure, rhs, identifierMap)
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					}
				case *Forward:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					}
				case *Glissando:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "LineType":
						inst.LineType = common.EnumLineType(GongExtractString(rhs))
					case "DashLength":
						inst.DashLength = GongExtractString(rhs)
					case "SpaceLength":
						inst.SpaceLength = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Grace:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StealTimePrevious":
						inst.StealTimePrevious = GongExtractString(rhs)
					case "StealTimeFollowing":
						inst.StealTimeFollowing = GongExtractString(rhs)
					case "MakeTime":
						inst.MakeTime = GongExtractString(rhs)
					case "Slash":
						inst.Slash = common.EnumYesNo(GongExtractString(rhs))
					}
				case *HammerOnPullOff:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Handbell:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Value":
						inst.Value = EnumHandbellValue(GongExtractString(rhs))
					}
				case *HarmonClosed:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						inst.Location = EnumHarmonClosedLocation(GongExtractString(rhs))
					case "Value":
						inst.Value = EnumHarmonClosedValue(GongExtractString(rhs))
					}
				case *HarmonMute:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "HarmonClosed":
						__gong__assignPointer(&inst.HarmonClosed, rhs, identifierMap)
					}
				case *Harmonic:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					}
				case *HeelToe:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Substitution":
						inst.Substitution = common.EnumYesNo(GongExtractString(rhs))
					}
				case *Hole:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "HoleType":
						inst.HoleType = GongExtractString(rhs)
					case "HoleClosed":
						__gong__assignPointer(&inst.HoleClosed, rhs, identifierMap)
					case "HoleShape":
						inst.HoleShape = GongExtractString(rhs)
					}
				case *HoleClosed:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						inst.Location = EnumHoleClosedLocation(GongExtractString(rhs))
					case "Value":
						inst.Value = EnumHoleClosedValue(GongExtractString(rhs))
					}
				case *Instrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Listen:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Assess":
						__gong__assignSliceOfPointers(&inst.Assess, rhs, identifierMap)
					case "Wait":
						__gong__assignSliceOfPointers(&inst.Wait, rhs, identifierMap)
					}
				case *Lyric:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "NameAttr":
						inst.NameAttr = GongExtractString(rhs)
					case "TimeOnly":
						inst.TimeOnly = GongExtractString(rhs)
					case "Justify":
						inst.Justify = common.EnumLeftCenterRight(GongExtractString(rhs))
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Syllabic":
						inst.Syllabic = EnumSyllabic(GongExtractString(rhs))
					case "Text":
						__gong__assignPointer(&inst.Text, rhs, identifierMap)
					case "Elision":
						__gong__assignSliceOfPointers(&inst.Elision, rhs, identifierMap)
					case "Extend":
						__gong__assignPointer(&inst.Extend, rhs, identifierMap)
					}
				case *Mordent:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "StartNote":
						inst.StartNote = common.EnumStartNote(GongExtractString(rhs))
					case "TrillStep":
						inst.TrillStep = common.EnumTrillStep(GongExtractString(rhs))
					case "TwoNoteTurn":
						inst.TwoNoteTurn = common.EnumTwoNoteTurn(GongExtractString(rhs))
					case "Accelerate":
						inst.Accelerate = common.EnumYesNo(GongExtractString(rhs))
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "SecondBeat":
						inst.SecondBeat = GongExtractString(rhs)
					case "LastBeat":
						inst.LastBeat = GongExtractString(rhs)
					case "Long":
						inst.Long = common.EnumYesNo(GongExtractString(rhs))
					case "Approach":
						inst.Approach = common.EnumAboveBelow(GongExtractString(rhs))
					case "Departure":
						inst.Departure = common.EnumAboveBelow(GongExtractString(rhs))
					}
				case *NonArpeggiate:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumTopBottom(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Notations:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
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
					case "Arpeggiate":
						__gong__assignSliceOfPointers(&inst.Arpeggiate, rhs, identifierMap)
					case "NonArpeggiate":
						__gong__assignSliceOfPointers(&inst.NonArpeggiate, rhs, identifierMap)
					case "AccidentalMark":
						__gong__assignSliceOfPointers(&inst.AccidentalMark, rhs, identifierMap)
					case "OtherNotation":
						__gong__assignSliceOfPointers(&inst.OtherNotation, rhs, identifierMap)
					}
				case *Note:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PrintLeger":
						inst.PrintLeger = common.EnumYesNo(GongExtractString(rhs))
					case "Dynamics":
						inst.Dynamics = GongExtractString(rhs)
					case "EndDynamics":
						inst.EndDynamics = GongExtractString(rhs)
					case "Attack":
						inst.Attack = GongExtractString(rhs)
					case "Release":
						inst.Release = GongExtractString(rhs)
					case "TimeOnly":
						inst.TimeOnly = GongExtractString(rhs)
					case "Pizzicato":
						inst.Pizzicato = common.EnumYesNo(GongExtractString(rhs))
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "PrintDot":
						inst.PrintDot = common.EnumYesNo(GongExtractString(rhs))
					case "PrintLyric":
						inst.PrintLyric = common.EnumYesNo(GongExtractString(rhs))
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "PrintSpacing":
						inst.PrintSpacing = common.EnumYesNo(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Grace":
						__gong__assignPointer(&inst.Grace, rhs, identifierMap)
					case "Pitch":
						__gong__assignPointer(&inst.Pitch, rhs, identifierMap)
					case "Unpitched":
						__gong__assignPointer(&inst.Unpitched, rhs, identifierMap)
					case "Rest":
						__gong__assignPointer(&inst.Rest, rhs, identifierMap)
					case "Tie":
						__gong__assignSliceOfPointers(&inst.Tie, rhs, identifierMap)
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					case "Instrument":
						__gong__assignSliceOfPointers(&inst.Instrument, rhs, identifierMap)
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Type":
						__gong__assignPointer(&inst.Type, rhs, identifierMap)
					case "Accidental":
						__gong__assignPointer(&inst.Accidental, rhs, identifierMap)
					case "Stem":
						__gong__assignPointer(&inst.Stem, rhs, identifierMap)
					case "Notehead":
						__gong__assignPointer(&inst.Notehead, rhs, identifierMap)
					case "NoteheadText":
						__gong__assignPointer(&inst.NoteheadText, rhs, identifierMap)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					case "Beam":
						__gong__assignSliceOfPointers(&inst.Beam, rhs, identifierMap)
					case "Notations":
						__gong__assignSliceOfPointers(&inst.Notations, rhs, identifierMap)
					case "Lyric":
						__gong__assignSliceOfPointers(&inst.Lyric, rhs, identifierMap)
					case "Listen":
						__gong__assignPointer(&inst.Listen, rhs, identifierMap)
					}
				case *NoteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Size":
						inst.Size = common.EnumSymbolSize(GongExtractString(rhs))
					case "Value":
						inst.Value = common.EnumNoteTypeValue(GongExtractString(rhs))
					}
				case *Notehead:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Filled":
						inst.Filled = common.EnumYesNo(GongExtractString(rhs))
					case "Parentheses":
						inst.Parentheses = common.EnumYesNo(GongExtractString(rhs))
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = EnumNoteheadValue(GongExtractString(rhs))
					}
				case *NoteheadText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Ornaments:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Mordent":
						__gong__assignSliceOfPointers(&inst.Mordent, rhs, identifierMap)
					case "InvertedMordent":
						__gong__assignSliceOfPointers(&inst.InvertedMordent, rhs, identifierMap)
					case "Tremolo":
						__gong__assignSliceOfPointers(&inst.Tremolo, rhs, identifierMap)
					case "OtherOrnament":
						__gong__assignSliceOfPointers(&inst.OtherOrnament, rhs, identifierMap)
					case "AccidentalMark":
						__gong__assignSliceOfPointers(&inst.AccidentalMark, rhs, identifierMap)
					}
				case *OtherNotation:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStopSingle(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *OtherPlacementText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Pitch:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Step":
						inst.Step = attributes.EnumStep(GongExtractString(rhs))
					case "Alter":
						inst.Alter = GongExtractString(rhs)
					case "Octave":
						inst.Octave = GongExtractInt(rhs)
					}
				case *PlacementText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Release:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Offset":
						inst.Offset = GongExtractString(rhs)
					}
				case *Rest:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Measure":
						inst.Measure = common.EnumYesNo(GongExtractString(rhs))
					case "DisplayStep":
						inst.DisplayStep = attributes.EnumStep(GongExtractString(rhs))
					case "DisplayOctave":
						inst.DisplayOctave = GongExtractInt(rhs)
					}
				case *Slide:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "LineType":
						inst.LineType = common.EnumLineType(GongExtractString(rhs))
					case "DashLength":
						inst.DashLength = GongExtractString(rhs)
					case "SpaceLength":
						inst.SpaceLength = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Accelerate":
						inst.Accelerate = common.EnumYesNo(GongExtractString(rhs))
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "FirstBeat":
						inst.FirstBeat = GongExtractString(rhs)
					case "LastBeat":
						inst.LastBeat = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Slur:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStopContinue(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "LineType":
						inst.LineType = common.EnumLineType(GongExtractString(rhs))
					case "DashLength":
						inst.DashLength = GongExtractString(rhs)
					case "SpaceLength":
						inst.SpaceLength = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Orientation":
						inst.Orientation = common.EnumOverUnder(GongExtractString(rhs))
					case "BezierX":
						inst.BezierX = GongExtractString(rhs)
					case "BezierY":
						inst.BezierY = GongExtractString(rhs)
					case "BezierX2":
						inst.BezierX2 = GongExtractString(rhs)
					case "BezierY2":
						inst.BezierY2 = GongExtractString(rhs)
					case "BezierOffset":
						inst.BezierOffset = GongExtractString(rhs)
					case "BezierOffset2":
						inst.BezierOffset2 = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Stem:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Value":
						inst.Value = EnumStemValue(GongExtractString(rhs))
					}
				case *StrongAccent:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Type":
						inst.Type = common.EnumUpDown(GongExtractString(rhs))
					}
				case *Tap:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Hand":
						inst.Hand = EnumTapHand(GongExtractString(rhs))
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Technical:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Harmonic":
						__gong__assignSliceOfPointers(&inst.Harmonic, rhs, identifierMap)
					case "Pluck":
						__gong__assignSliceOfPointers(&inst.Pluck, rhs, identifierMap)
					case "HammerOn":
						__gong__assignSliceOfPointers(&inst.HammerOn, rhs, identifierMap)
					case "PullOff":
						__gong__assignSliceOfPointers(&inst.PullOff, rhs, identifierMap)
					case "Bend":
						__gong__assignSliceOfPointers(&inst.Bend, rhs, identifierMap)
					case "Tap":
						__gong__assignSliceOfPointers(&inst.Tap, rhs, identifierMap)
					case "Heel":
						__gong__assignSliceOfPointers(&inst.Heel, rhs, identifierMap)
					case "Toe":
						__gong__assignSliceOfPointers(&inst.Toe, rhs, identifierMap)
					case "Hole":
						__gong__assignSliceOfPointers(&inst.Hole, rhs, identifierMap)
					case "Arrow":
						__gong__assignSliceOfPointers(&inst.Arrow, rhs, identifierMap)
					case "Handbell":
						__gong__assignSliceOfPointers(&inst.Handbell, rhs, identifierMap)
					case "HarmonMute":
						__gong__assignSliceOfPointers(&inst.HarmonMute, rhs, identifierMap)
					case "OtherTechnical":
						__gong__assignSliceOfPointers(&inst.OtherTechnical, rhs, identifierMap)
					}
				case *TextElementData:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "LineThrough":
						inst.LineThrough = GongExtractInt(rhs)
					case "Rotation":
						inst.Rotation = GongExtractString(rhs)
					case "LetterSpacing":
						inst.LetterSpacing = GongExtractString(rhs)
					case "Dir":
						inst.Dir = common.EnumTextDirection(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Tie:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "TimeOnly":
						inst.TimeOnly = GongExtractString(rhs)
					}
				case *Tied:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumTiedType(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "LineType":
						inst.LineType = common.EnumLineType(GongExtractString(rhs))
					case "DashLength":
						inst.DashLength = GongExtractString(rhs)
					case "SpaceLength":
						inst.SpaceLength = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Orientation":
						inst.Orientation = common.EnumOverUnder(GongExtractString(rhs))
					case "BezierX":
						inst.BezierX = GongExtractString(rhs)
					case "BezierY":
						inst.BezierY = GongExtractString(rhs)
					case "BezierX2":
						inst.BezierX2 = GongExtractString(rhs)
					case "BezierY2":
						inst.BezierY2 = GongExtractString(rhs)
					case "BezierOffset":
						inst.BezierOffset = GongExtractString(rhs)
					case "BezierOffset2":
						inst.BezierOffset2 = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Tremolo:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumTremoloType(GongExtractString(rhs))
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractInt(rhs)
					}
				case *Tuplet:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Bracket":
						inst.Bracket = common.EnumYesNo(GongExtractString(rhs))
					case "ShowNumber":
						inst.ShowNumber = direction.EnumShowTuplet(GongExtractString(rhs))
					case "ShowType":
						inst.ShowType = direction.EnumShowTuplet(GongExtractString(rhs))
					case "LineShape":
						inst.LineShape = common.EnumLineShape(GongExtractString(rhs))
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "TupletActual":
						__gong__assignPointer(&inst.TupletActual, rhs, identifierMap)
					case "TupletNormal":
						__gong__assignPointer(&inst.TupletNormal, rhs, identifierMap)
					}
				case *TupletDot:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *TupletNumber:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractInt(rhs)
					}
				case *TupletPortion:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "TupletNumber":
						__gong__assignPointer(&inst.TupletNumber, rhs, identifierMap)
					case "TupletType":
						__gong__assignPointer(&inst.TupletType, rhs, identifierMap)
					case "TupletDot":
						__gong__assignSliceOfPointers(&inst.TupletDot, rhs, identifierMap)
					}
				case *TupletType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Value":
						inst.Value = common.EnumNoteTypeValue(GongExtractString(rhs))
					}
				case *Unpitched:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DisplayStep":
						inst.DisplayStep = attributes.EnumStep(GongExtractString(rhs))
					case "DisplayOctave":
						inst.DisplayOctave = GongExtractInt(rhs)
					}
				case *Wait:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "TimeOnly":
						inst.TimeOnly = GongExtractString(rhs)
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
