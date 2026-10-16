// generated code - do not edit
package attributes

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
		for _, attributes := range __gong__sortStageSetInstances(stageSet.Stage.Attributess, stageSet.Stage.Attributes_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attributesIdent := "__attributes" + attributes.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Attributes{Name: %s}).Stage(stageSet.Stage)", attributesIdent, __gong__toRawStringLiteral(attributes.Name)))
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
			for _, elem := range attributes.Key {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Key = append(%s.Key, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Time {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Time = append(%s.Time, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			if attributes.PartSymbol != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + attributes.PartSymbol.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartSymbol = %s", attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Clef {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Clef = append(%s.Clef, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.StaffDetails {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StaffDetails = append(%s.StaffDetails, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Transpose {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Transpose = append(%s.Transpose, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.ForPart {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ForPart = append(%s.ForPart, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Directive {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Directive = append(%s.Directive, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.MeasureStyle {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MeasureStyle = append(%s.MeasureStyle, %s)", attributesIdent, attributesIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, attributesdirective := range __gong__sortStageSetInstances(stageSet.Stage.AttributesDirectives, stageSet.Stage.AttributesDirective_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attributesdirectiveIdent := "__attributes" + attributesdirective.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.AttributesDirective{Name: %s}).Stage(stageSet.Stage)", attributesdirectiveIdent, __gong__toRawStringLiteral(attributesdirective.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attributesdirectiveIdent, __gong__toRawStringLiteral(attributesdirective.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", attributesdirectiveIdent, __gong__toRawStringLiteral(attributesdirective.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", attributesdirectiveIdent, __gong__toRawStringLiteral(attributesdirective.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", attributesdirectiveIdent, __gong__toRawStringLiteral(attributesdirective.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", attributesdirectiveIdent, __gong__toRawStringLiteral(attributesdirective.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", attributesdirectiveIdent, __gong__toRawStringLiteral(attributesdirective.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", attributesdirectiveIdent, __gong__toRawStringLiteral(attributesdirective.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", attributesdirectiveIdent, __gong__toRawStringLiteral(string(attributesdirective.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", attributesdirectiveIdent, __gong__toRawStringLiteral(attributesdirective.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", attributesdirectiveIdent, __gong__toRawStringLiteral(string(attributesdirective.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", attributesdirectiveIdent, __gong__toRawStringLiteral(attributesdirective.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", attributesdirectiveIdent, __gong__toRawStringLiteral(attributesdirective.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, beatrepeat := range __gong__sortStageSetInstances(stageSet.Stage.BeatRepeats, stageSet.Stage.BeatRepeat_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			beatrepeatIdent := "__attributes" + beatrepeat.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.BeatRepeat{Name: %s}).Stage(stageSet.Stage)", beatrepeatIdent, __gong__toRawStringLiteral(beatrepeat.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", beatrepeatIdent, __gong__toRawStringLiteral(beatrepeat.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", beatrepeatIdent, __gong__toRawStringLiteral(string(beatrepeat.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Slashes = %d", beatrepeatIdent, beatrepeat.Slashes))
			values.WriteString(fmt.Sprintf("\n\t%s.UseDots = %s", beatrepeatIdent, __gong__toRawStringLiteral(string(beatrepeat.UseDots))))
			values.WriteString(fmt.Sprintf("\n\t%s.SlashType = %s", beatrepeatIdent, __gong__toRawStringLiteral(string(beatrepeat.SlashType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ExceptVoice = %s", beatrepeatIdent, __gong__toRawStringLiteral(beatrepeat.ExceptVoice)))
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
			cancelIdent := "__attributes" + cancel.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Cancel{Name: %s}).Stage(stageSet.Stage)", cancelIdent, __gong__toRawStringLiteral(cancel.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cancelIdent, __gong__toRawStringLiteral(cancel.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", cancelIdent, __gong__toRawStringLiteral(string(cancel.Location))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", cancelIdent, cancel.Value))
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
			clefIdent := "__attributes" + clef.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Clef{Name: %s}).Stage(stageSet.Stage)", clefIdent, __gong__toRawStringLiteral(clef.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", clefIdent, __gong__toRawStringLiteral(clef.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", clefIdent, clef.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Additional = %s", clefIdent, __gong__toRawStringLiteral(string(clef.Additional))))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %s", clefIdent, __gong__toRawStringLiteral(string(clef.Size))))
			values.WriteString(fmt.Sprintf("\n\t%s.AfterBarline = %s", clefIdent, __gong__toRawStringLiteral(string(clef.AfterBarline))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", clefIdent, __gong__toRawStringLiteral(clef.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", clefIdent, __gong__toRawStringLiteral(clef.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", clefIdent, __gong__toRawStringLiteral(clef.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", clefIdent, __gong__toRawStringLiteral(clef.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", clefIdent, __gong__toRawStringLiteral(clef.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", clefIdent, __gong__toRawStringLiteral(string(clef.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", clefIdent, __gong__toRawStringLiteral(clef.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", clefIdent, __gong__toRawStringLiteral(string(clef.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", clefIdent, __gong__toRawStringLiteral(clef.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", clefIdent, __gong__toRawStringLiteral(string(clef.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", clefIdent, __gong__toRawStringLiteral(clef.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sign = %s", clefIdent, __gong__toRawStringLiteral(string(clef.Sign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", clefIdent, clef.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.ClefOctaveChange = %d", clefIdent, clef.ClefOctaveChange))
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
			doubleIdent := "__attributes" + double.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Double{Name: %s}).Stage(stageSet.Stage)", doubleIdent, __gong__toRawStringLiteral(double.Name)))
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
		for _, forpart := range __gong__sortStageSetInstances(stageSet.Stage.ForParts, stageSet.Stage.ForPart_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			forpartIdent := "__attributes" + forpart.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.ForPart{Name: %s}).Stage(stageSet.Stage)", forpartIdent, __gong__toRawStringLiteral(forpart.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", forpartIdent, __gong__toRawStringLiteral(forpart.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", forpartIdent, forpart.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", forpartIdent, __gong__toRawStringLiteral(forpart.Id)))
			if forpart.PartClef != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + forpart.PartClef.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartClef = %s", forpartIdent, targetIdent))
			}
			if forpart.PartTranspose != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + forpart.PartTranspose.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartTranspose = %s", forpartIdent, targetIdent))
			}
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
			interchangeableIdent := "__attributes" + interchangeable.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Interchangeable{Name: %s}).Stage(stageSet.Stage)", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Symbol = %s", interchangeableIdent, __gong__toRawStringLiteral(string(interchangeable.Symbol))))
			values.WriteString(fmt.Sprintf("\n\t%s.Separator = %s", interchangeableIdent, __gong__toRawStringLiteral(string(interchangeable.Separator))))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeRelation = %s", interchangeableIdent, __gong__toRawStringLiteral(string(interchangeable.TimeRelation))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.BeatType = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.BeatType)))
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
			keyIdent := "__attributes" + key.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Key{Name: %s}).Stage(stageSet.Stage)", keyIdent, __gong__toRawStringLiteral(key.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", keyIdent, __gong__toRawStringLiteral(key.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", keyIdent, key.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", keyIdent, __gong__toRawStringLiteral(key.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", keyIdent, __gong__toRawStringLiteral(key.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", keyIdent, __gong__toRawStringLiteral(key.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", keyIdent, __gong__toRawStringLiteral(key.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", keyIdent, __gong__toRawStringLiteral(key.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", keyIdent, __gong__toRawStringLiteral(string(key.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", keyIdent, __gong__toRawStringLiteral(key.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", keyIdent, __gong__toRawStringLiteral(string(key.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", keyIdent, __gong__toRawStringLiteral(key.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", keyIdent, __gong__toRawStringLiteral(string(key.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", keyIdent, __gong__toRawStringLiteral(key.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Fifths = %d", keyIdent, key.Fifths))
			values.WriteString(fmt.Sprintf("\n\t%s.Mode = %s", keyIdent, __gong__toRawStringLiteral(key.Mode)))
			values.WriteString(fmt.Sprintf("\n\t%s.KeyStep = %s", keyIdent, __gong__toRawStringLiteral(string(key.KeyStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.KeyAlter = %s", keyIdent, __gong__toRawStringLiteral(key.KeyAlter)))
			if key.Cancel != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + key.Cancel.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Cancel = %s", keyIdent, targetIdent))
			}
			for _, elem := range key.KeyAccidental {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.KeyAccidental = append(%s.KeyAccidental, %s)", keyIdent, keyIdent, targetIdent))
			}
			for _, elem := range key.KeyOctave {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.KeyOctave = append(%s.KeyOctave, %s)", keyIdent, keyIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, keyaccidental := range __gong__sortStageSetInstances(stageSet.Stage.KeyAccidentals, stageSet.Stage.KeyAccidental_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			keyaccidentalIdent := "__attributes" + keyaccidental.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.KeyAccidental{Name: %s}).Stage(stageSet.Stage)", keyaccidentalIdent, __gong__toRawStringLiteral(keyaccidental.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", keyaccidentalIdent, __gong__toRawStringLiteral(keyaccidental.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", keyaccidentalIdent, __gong__toRawStringLiteral(keyaccidental.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", keyaccidentalIdent, __gong__toRawStringLiteral(string(keyaccidental.Value))))
		}
	}
	if stageSet.Stage != nil {
		for _, keyoctave := range __gong__sortStageSetInstances(stageSet.Stage.KeyOctaves, stageSet.Stage.KeyOctave_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			keyoctaveIdent := "__attributes" + keyoctave.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.KeyOctave{Name: %s}).Stage(stageSet.Stage)", keyoctaveIdent, __gong__toRawStringLiteral(keyoctave.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", keyoctaveIdent, __gong__toRawStringLiteral(keyoctave.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", keyoctaveIdent, keyoctave.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Cancel = %s", keyoctaveIdent, __gong__toRawStringLiteral(string(keyoctave.Cancel))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", keyoctaveIdent, keyoctave.Value))
		}
	}
	if stageSet.Stage != nil {
		for _, linedetail := range __gong__sortStageSetInstances(stageSet.Stage.LineDetails, stageSet.Stage.LineDetail_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linedetailIdent := "__attributes" + linedetail.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.LineDetail{Name: %s}).Stage(stageSet.Stage)", linedetailIdent, __gong__toRawStringLiteral(linedetail.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linedetailIdent, __gong__toRawStringLiteral(linedetail.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", linedetailIdent, linedetail.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", linedetailIdent, __gong__toRawStringLiteral(linedetail.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", linedetailIdent, __gong__toRawStringLiteral(linedetail.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.LineType = %s", linedetailIdent, __gong__toRawStringLiteral(string(linedetail.LineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", linedetailIdent, __gong__toRawStringLiteral(string(linedetail.PrintObject))))
		}
	}
	if stageSet.Stage != nil {
		for _, measurerepeat := range __gong__sortStageSetInstances(stageSet.Stage.MeasureRepeats, stageSet.Stage.MeasureRepeat_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			measurerepeatIdent := "__attributes" + measurerepeat.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.MeasureRepeat{Name: %s}).Stage(stageSet.Stage)", measurerepeatIdent, __gong__toRawStringLiteral(measurerepeat.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measurerepeatIdent, __gong__toRawStringLiteral(measurerepeat.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", measurerepeatIdent, __gong__toRawStringLiteral(string(measurerepeat.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Slashes = %d", measurerepeatIdent, measurerepeat.Slashes))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", measurerepeatIdent, __gong__toRawStringLiteral(measurerepeat.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, measurestyle := range __gong__sortStageSetInstances(stageSet.Stage.MeasureStyles, stageSet.Stage.MeasureStyle_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			measurestyleIdent := "__attributes" + measurestyle.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.MeasureStyle{Name: %s}).Stage(stageSet.Stage)", measurestyleIdent, __gong__toRawStringLiteral(measurestyle.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measurestyleIdent, __gong__toRawStringLiteral(measurestyle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", measurestyleIdent, measurestyle.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", measurestyleIdent, __gong__toRawStringLiteral(measurestyle.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", measurestyleIdent, __gong__toRawStringLiteral(string(measurestyle.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", measurestyleIdent, __gong__toRawStringLiteral(measurestyle.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", measurestyleIdent, __gong__toRawStringLiteral(string(measurestyle.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", measurestyleIdent, __gong__toRawStringLiteral(measurestyle.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", measurestyleIdent, __gong__toRawStringLiteral(measurestyle.Id)))
			if measurestyle.MultipleRest != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + measurestyle.MultipleRest.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MultipleRest = %s", measurestyleIdent, targetIdent))
			}
			if measurestyle.MeasureRepeat != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + measurestyle.MeasureRepeat.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MeasureRepeat = %s", measurestyleIdent, targetIdent))
			}
			if measurestyle.BeatRepeat != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + measurestyle.BeatRepeat.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BeatRepeat = %s", measurestyleIdent, targetIdent))
			}
			if measurestyle.Slash != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + measurestyle.Slash.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slash = %s", measurestyleIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, multiplerest := range __gong__sortStageSetInstances(stageSet.Stage.MultipleRests, stageSet.Stage.MultipleRest_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			multiplerestIdent := "__attributes" + multiplerest.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.MultipleRest{Name: %s}).Stage(stageSet.Stage)", multiplerestIdent, __gong__toRawStringLiteral(multiplerest.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", multiplerestIdent, __gong__toRawStringLiteral(multiplerest.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.UseSymbols = %s", multiplerestIdent, __gong__toRawStringLiteral(string(multiplerest.UseSymbols))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", multiplerestIdent, multiplerest.Value))
		}
	}
	if stageSet.Stage != nil {
		for _, partclef := range __gong__sortStageSetInstances(stageSet.Stage.PartClefs, stageSet.Stage.PartClef_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			partclefIdent := "__attributes" + partclef.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.PartClef{Name: %s}).Stage(stageSet.Stage)", partclefIdent, __gong__toRawStringLiteral(partclef.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partclefIdent, __gong__toRawStringLiteral(partclef.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sign = %s", partclefIdent, __gong__toRawStringLiteral(string(partclef.Sign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", partclefIdent, partclef.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.ClefOctaveChange = %d", partclefIdent, partclef.ClefOctaveChange))
		}
	}
	if stageSet.Stage != nil {
		for _, partsymbol := range __gong__sortStageSetInstances(stageSet.Stage.PartSymbols, stageSet.Stage.PartSymbol_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			partsymbolIdent := "__attributes" + partsymbol.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.PartSymbol{Name: %s}).Stage(stageSet.Stage)", partsymbolIdent, __gong__toRawStringLiteral(partsymbol.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partsymbolIdent, __gong__toRawStringLiteral(partsymbol.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.TopStaff = %d", partsymbolIdent, partsymbol.TopStaff))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomStaff = %d", partsymbolIdent, partsymbol.BottomStaff))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", partsymbolIdent, __gong__toRawStringLiteral(partsymbol.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", partsymbolIdent, __gong__toRawStringLiteral(partsymbol.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", partsymbolIdent, __gong__toRawStringLiteral(partsymbol.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", partsymbolIdent, __gong__toRawStringLiteral(partsymbol.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", partsymbolIdent, __gong__toRawStringLiteral(partsymbol.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", partsymbolIdent, __gong__toRawStringLiteral(string(partsymbol.Value))))
		}
	}
	if stageSet.Stage != nil {
		for _, parttranspose := range __gong__sortStageSetInstances(stageSet.Stage.PartTransposes, stageSet.Stage.PartTranspose_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parttransposeIdent := "__attributes" + parttranspose.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.PartTranspose{Name: %s}).Stage(stageSet.Stage)", parttransposeIdent, __gong__toRawStringLiteral(parttranspose.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parttransposeIdent, __gong__toRawStringLiteral(parttranspose.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Diatonic = %d", parttransposeIdent, parttranspose.Diatonic))
			values.WriteString(fmt.Sprintf("\n\t%s.Chromatic = %s", parttransposeIdent, __gong__toRawStringLiteral(parttranspose.Chromatic)))
			values.WriteString(fmt.Sprintf("\n\t%s.OctaveChange = %d", parttransposeIdent, parttranspose.OctaveChange))
			if parttranspose.Double != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + parttranspose.Double.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Double = %s", parttransposeIdent, targetIdent))
			}
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
			slashIdent := "__attributes" + slash.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Slash{Name: %s}).Stage(stageSet.Stage)", slashIdent, __gong__toRawStringLiteral(slash.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", slashIdent, __gong__toRawStringLiteral(slash.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", slashIdent, __gong__toRawStringLiteral(string(slash.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.UseDots = %s", slashIdent, __gong__toRawStringLiteral(string(slash.UseDots))))
			values.WriteString(fmt.Sprintf("\n\t%s.UseStems = %s", slashIdent, __gong__toRawStringLiteral(string(slash.UseStems))))
			values.WriteString(fmt.Sprintf("\n\t%s.SlashType = %s", slashIdent, __gong__toRawStringLiteral(string(slash.SlashType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ExceptVoice = %s", slashIdent, __gong__toRawStringLiteral(slash.ExceptVoice)))
		}
	}
	if stageSet.Stage != nil {
		for _, staffdetails := range __gong__sortStageSetInstances(stageSet.Stage.StaffDetailss, stageSet.Stage.StaffDetails_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staffdetailsIdent := "__attributes" + staffdetails.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.StaffDetails{Name: %s}).Stage(stageSet.Stage)", staffdetailsIdent, __gong__toRawStringLiteral(staffdetails.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staffdetailsIdent, __gong__toRawStringLiteral(staffdetails.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", staffdetailsIdent, staffdetails.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowFrets = %s", staffdetailsIdent, __gong__toRawStringLiteral(string(staffdetails.ShowFrets))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", staffdetailsIdent, __gong__toRawStringLiteral(string(staffdetails.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintSpacing = %s", staffdetailsIdent, __gong__toRawStringLiteral(string(staffdetails.PrintSpacing))))
			values.WriteString(fmt.Sprintf("\n\t%s.StaffType = %s", staffdetailsIdent, __gong__toRawStringLiteral(string(staffdetails.StaffType))))
			values.WriteString(fmt.Sprintf("\n\t%s.StaffLines = %d", staffdetailsIdent, staffdetails.StaffLines))
			values.WriteString(fmt.Sprintf("\n\t%s.Capo = %d", staffdetailsIdent, staffdetails.Capo))
			for _, elem := range staffdetails.LineDetail {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LineDetail = append(%s.LineDetail, %s)", staffdetailsIdent, staffdetailsIdent, targetIdent))
			}
			for _, elem := range staffdetails.StaffTuning {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StaffTuning = append(%s.StaffTuning, %s)", staffdetailsIdent, staffdetailsIdent, targetIdent))
			}
			if staffdetails.StaffSize != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + staffdetails.StaffSize.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StaffSize = %s", staffdetailsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, staffsize := range __gong__sortStageSetInstances(stageSet.Stage.StaffSizes, stageSet.Stage.StaffSize_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staffsizeIdent := "__attributes" + staffsize.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.StaffSize{Name: %s}).Stage(stageSet.Stage)", staffsizeIdent, __gong__toRawStringLiteral(staffsize.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staffsizeIdent, __gong__toRawStringLiteral(staffsize.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Scaling = %s", staffsizeIdent, __gong__toRawStringLiteral(staffsize.Scaling)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", staffsizeIdent, __gong__toRawStringLiteral(staffsize.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, stafftuning := range __gong__sortStageSetInstances(stageSet.Stage.StaffTunings, stageSet.Stage.StaffTuning_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stafftuningIdent := "__attributes" + stafftuning.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.StaffTuning{Name: %s}).Stage(stageSet.Stage)", stafftuningIdent, __gong__toRawStringLiteral(stafftuning.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stafftuningIdent, __gong__toRawStringLiteral(stafftuning.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", stafftuningIdent, stafftuning.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.TuningStep = %s", stafftuningIdent, __gong__toRawStringLiteral(string(stafftuning.TuningStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.TuningAlter = %s", stafftuningIdent, __gong__toRawStringLiteral(stafftuning.TuningAlter)))
			values.WriteString(fmt.Sprintf("\n\t%s.TuningOctave = %d", stafftuningIdent, stafftuning.TuningOctave))
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
			timeIdent := "__attributes" + time.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Time{Name: %s}).Stage(stageSet.Stage)", timeIdent, __gong__toRawStringLiteral(time.Name)))
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
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", timeIdent, __gong__toRawStringLiteral(time.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", timeIdent, __gong__toRawStringLiteral(time.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", timeIdent, __gong__toRawStringLiteral(time.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", timeIdent, __gong__toRawStringLiteral(time.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", timeIdent, __gong__toRawStringLiteral(time.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", timeIdent, __gong__toRawStringLiteral(string(time.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", timeIdent, __gong__toRawStringLiteral(time.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", timeIdent, __gong__toRawStringLiteral(string(time.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", timeIdent, __gong__toRawStringLiteral(time.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", timeIdent, __gong__toRawStringLiteral(string(time.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", timeIdent, __gong__toRawStringLiteral(string(time.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", timeIdent, __gong__toRawStringLiteral(string(time.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", timeIdent, __gong__toRawStringLiteral(time.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", timeIdent, __gong__toRawStringLiteral(time.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.BeatType = %s", timeIdent, __gong__toRawStringLiteral(time.BeatType)))
			values.WriteString(fmt.Sprintf("\n\t%s.SenzaMisura = %s", timeIdent, __gong__toRawStringLiteral(time.SenzaMisura)))
			if time.Interchangeable != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + time.Interchangeable.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Interchangeable = %s", timeIdent, targetIdent))
			}
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
			transposeIdent := "__attributes" + transpose.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Transpose{Name: %s}).Stage(stageSet.Stage)", transposeIdent, __gong__toRawStringLiteral(transpose.Name)))
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
			values.WriteString(fmt.Sprintf("\n\t%s.OctaveChange = %d", transposeIdent, transpose.OctaveChange))
			if transpose.Double != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__attributes" + transpose.Double.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Double = %s", transposeIdent, targetIdent))
			}
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
go run github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *attributes.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *attributes.StageSet) {

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
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes":
			aliasToCanonical[alias] = "attributes"
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
			case "attributes":
				switch typeName {
				case "Attributes":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Attributes), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "AttributesDirective":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AttributesDirective), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "BeatRepeat":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BeatRepeat), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Cancel":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Cancel), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Clef":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Clef), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Double":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Double), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ForPart":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ForPart), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Interchangeable":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Interchangeable), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Key":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Key), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "KeyAccidental":
					identifierMap[ident.Name] = __gong__stageSetInit(new(KeyAccidental), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "KeyOctave":
					identifierMap[ident.Name] = __gong__stageSetInit(new(KeyOctave), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "LineDetail":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LineDetail), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MeasureRepeat":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeasureRepeat), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MeasureStyle":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeasureStyle), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MultipleRest":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MultipleRest), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PartClef":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PartClef), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PartSymbol":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PartSymbol), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PartTranspose":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PartTranspose), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Slash":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Slash), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StaffDetails":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StaffDetails), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StaffSize":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StaffSize), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StaffTuning":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StaffTuning), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Time":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Time), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Transpose":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Transpose), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *Attributes:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Divisions":
						inst.Divisions = GongExtractString(rhs)
					case "Key":
						__gong__assignSliceOfPointers(&inst.Key, rhs, identifierMap)
					case "Time":
						__gong__assignSliceOfPointers(&inst.Time, rhs, identifierMap)
					case "Staves":
						inst.Staves = GongExtractInt(rhs)
					case "PartSymbol":
						__gong__assignPointer(&inst.PartSymbol, rhs, identifierMap)
					case "Instruments":
						inst.Instruments = GongExtractInt(rhs)
					case "Clef":
						__gong__assignSliceOfPointers(&inst.Clef, rhs, identifierMap)
					case "StaffDetails":
						__gong__assignSliceOfPointers(&inst.StaffDetails, rhs, identifierMap)
					case "Transpose":
						__gong__assignSliceOfPointers(&inst.Transpose, rhs, identifierMap)
					case "ForPart":
						__gong__assignSliceOfPointers(&inst.ForPart, rhs, identifierMap)
					case "Directive":
						__gong__assignSliceOfPointers(&inst.Directive, rhs, identifierMap)
					case "MeasureStyle":
						__gong__assignSliceOfPointers(&inst.MeasureStyle, rhs, identifierMap)
					}
				case *AttributesDirective:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
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
						inst.Value = GongExtractString(rhs)
					}
				case *BeatRepeat:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Slashes":
						inst.Slashes = GongExtractInt(rhs)
					case "UseDots":
						inst.UseDots = common.EnumYesNo(GongExtractString(rhs))
					case "SlashType":
						inst.SlashType = common.EnumNoteTypeValue(GongExtractString(rhs))
					case "ExceptVoice":
						inst.ExceptVoice = GongExtractString(rhs)
					}
				case *Cancel:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						inst.Location = EnumCancelLocation(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractInt(rhs)
					}
				case *Clef:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Additional":
						inst.Additional = common.EnumYesNo(GongExtractString(rhs))
					case "Size":
						inst.Size = common.EnumSymbolSize(GongExtractString(rhs))
					case "AfterBarline":
						inst.AfterBarline = common.EnumYesNo(GongExtractString(rhs))
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
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Sign":
						inst.Sign = EnumClefSign(GongExtractString(rhs))
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "ClefOctaveChange":
						inst.ClefOctaveChange = GongExtractInt(rhs)
					}
				case *Double:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Above":
						inst.Above = common.EnumYesNo(GongExtractString(rhs))
					}
				case *ForPart:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "PartClef":
						__gong__assignPointer(&inst.PartClef, rhs, identifierMap)
					case "PartTranspose":
						__gong__assignPointer(&inst.PartTranspose, rhs, identifierMap)
					}
				case *Interchangeable:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Symbol":
						inst.Symbol = EnumTimeSymbol(GongExtractString(rhs))
					case "Separator":
						inst.Separator = EnumTimeSeparator(GongExtractString(rhs))
					case "TimeRelation":
						inst.TimeRelation = EnumTimeRelation(GongExtractString(rhs))
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "BeatType":
						inst.BeatType = GongExtractString(rhs)
					}
				case *Key:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Cancel":
						__gong__assignPointer(&inst.Cancel, rhs, identifierMap)
					case "Fifths":
						inst.Fifths = GongExtractInt(rhs)
					case "Mode":
						inst.Mode = GongExtractString(rhs)
					case "KeyStep":
						inst.KeyStep = EnumStep(GongExtractString(rhs))
					case "KeyAlter":
						inst.KeyAlter = GongExtractString(rhs)
					case "KeyAccidental":
						__gong__assignSliceOfPointers(&inst.KeyAccidental, rhs, identifierMap)
					case "KeyOctave":
						__gong__assignSliceOfPointers(&inst.KeyOctave, rhs, identifierMap)
					}
				case *KeyAccidental:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = common.EnumAccidentalValue(GongExtractString(rhs))
					}
				case *KeyOctave:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Cancel":
						inst.Cancel = common.EnumYesNo(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractInt(rhs)
					}
				case *LineDetail:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "Width":
						inst.Width = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "LineType":
						inst.LineType = common.EnumLineType(GongExtractString(rhs))
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					}
				case *MeasureRepeat:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Slashes":
						inst.Slashes = GongExtractInt(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MeasureStyle:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
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
					case "MultipleRest":
						__gong__assignPointer(&inst.MultipleRest, rhs, identifierMap)
					case "MeasureRepeat":
						__gong__assignPointer(&inst.MeasureRepeat, rhs, identifierMap)
					case "BeatRepeat":
						__gong__assignPointer(&inst.BeatRepeat, rhs, identifierMap)
					case "Slash":
						__gong__assignPointer(&inst.Slash, rhs, identifierMap)
					}
				case *MultipleRest:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "UseSymbols":
						inst.UseSymbols = common.EnumYesNo(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractInt(rhs)
					}
				case *PartClef:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Sign":
						inst.Sign = EnumClefSign(GongExtractString(rhs))
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "ClefOctaveChange":
						inst.ClefOctaveChange = GongExtractInt(rhs)
					}
				case *PartSymbol:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "TopStaff":
						inst.TopStaff = GongExtractInt(rhs)
					case "BottomStaff":
						inst.BottomStaff = GongExtractInt(rhs)
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
						inst.Value = EnumGroupSymbolValue(GongExtractString(rhs))
					}
				case *PartTranspose:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Diatonic":
						inst.Diatonic = GongExtractInt(rhs)
					case "Chromatic":
						inst.Chromatic = GongExtractString(rhs)
					case "OctaveChange":
						inst.OctaveChange = GongExtractInt(rhs)
					case "Double":
						__gong__assignPointer(&inst.Double, rhs, identifierMap)
					}
				case *Slash:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "UseDots":
						inst.UseDots = common.EnumYesNo(GongExtractString(rhs))
					case "UseStems":
						inst.UseStems = common.EnumYesNo(GongExtractString(rhs))
					case "SlashType":
						inst.SlashType = common.EnumNoteTypeValue(GongExtractString(rhs))
					case "ExceptVoice":
						inst.ExceptVoice = GongExtractString(rhs)
					}
				case *StaffDetails:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "ShowFrets":
						inst.ShowFrets = EnumShowFrets(GongExtractString(rhs))
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "PrintSpacing":
						inst.PrintSpacing = common.EnumYesNo(GongExtractString(rhs))
					case "StaffType":
						inst.StaffType = EnumStaffType(GongExtractString(rhs))
					case "StaffLines":
						inst.StaffLines = GongExtractInt(rhs)
					case "LineDetail":
						__gong__assignSliceOfPointers(&inst.LineDetail, rhs, identifierMap)
					case "StaffTuning":
						__gong__assignSliceOfPointers(&inst.StaffTuning, rhs, identifierMap)
					case "Capo":
						inst.Capo = GongExtractInt(rhs)
					case "StaffSize":
						__gong__assignPointer(&inst.StaffSize, rhs, identifierMap)
					}
				case *StaffSize:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Scaling":
						inst.Scaling = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *StaffTuning:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "TuningStep":
						inst.TuningStep = EnumStep(GongExtractString(rhs))
					case "TuningAlter":
						inst.TuningAlter = GongExtractString(rhs)
					case "TuningOctave":
						inst.TuningOctave = GongExtractInt(rhs)
					}
				case *Time:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Symbol":
						inst.Symbol = EnumTimeSymbol(GongExtractString(rhs))
					case "Separator":
						inst.Separator = EnumTimeSeparator(GongExtractString(rhs))
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
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "BeatType":
						inst.BeatType = GongExtractString(rhs)
					case "Interchangeable":
						__gong__assignPointer(&inst.Interchangeable, rhs, identifierMap)
					case "SenzaMisura":
						inst.SenzaMisura = GongExtractString(rhs)
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
					case "OctaveChange":
						inst.OctaveChange = GongExtractInt(rhs)
					case "Double":
						__gong__assignPointer(&inst.Double, rhs, identifierMap)
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
