// generated code - do not edit
package common

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

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/xlink"
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
		for _, accidentaltext := range __gong__sortStageSetInstances(stageSet.Stage.AccidentalTexts, stageSet.Stage.AccidentalText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			accidentaltextIdent := "__common" + accidentaltext.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.AccidentalText{Name: %s}).Stage(stageSet.Stage)", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.Space)))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", accidentaltextIdent, __gong__toRawStringLiteral(string(accidentaltext.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", accidentaltextIdent, __gong__toRawStringLiteral(string(accidentaltext.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", accidentaltextIdent, __gong__toRawStringLiteral(string(accidentaltext.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", accidentaltextIdent, __gong__toRawStringLiteral(string(accidentaltext.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", accidentaltextIdent, __gong__toRawStringLiteral(string(accidentaltext.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", accidentaltextIdent, accidentaltext.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", accidentaltextIdent, accidentaltext.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.LineThrough = %d", accidentaltextIdent, accidentaltext.LineThrough))
			values.WriteString(fmt.Sprintf("\n\t%s.Rotation = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.Rotation)))
			values.WriteString(fmt.Sprintf("\n\t%s.LetterSpacing = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.LetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.LineHeight = %s", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.LineHeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dir = %s", accidentaltextIdent, __gong__toRawStringLiteral(string(accidentaltext.Dir))))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", accidentaltextIdent, __gong__toRawStringLiteral(string(accidentaltext.Enclosure))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", accidentaltextIdent, __gong__toRawStringLiteral(string(accidentaltext.Value))))
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
			appearanceIdent := "__common" + appearance.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Appearance{Name: %s}).Stage(stageSet.Stage)", appearanceIdent, __gong__toRawStringLiteral(appearance.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", appearanceIdent, __gong__toRawStringLiteral(appearance.Name)))
			for _, elem := range appearance.LineWidth {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LineWidth = append(%s.LineWidth, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.NoteSize {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteSize = append(%s.NoteSize, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.Distance {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = append(%s.Distance, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.Glyph {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Glyph = append(%s.Glyph, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.OtherAppearance {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherAppearance = append(%s.OtherAppearance, %s)", appearanceIdent, appearanceIdent, targetIdent))
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
			bookmarkIdent := "__common" + bookmark.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Bookmark{Name: %s}).Stage(stageSet.Stage)", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameAttr = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.NameAttr)))
			values.WriteString(fmt.Sprintf("\n\t%s.Element = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Element)))
			values.WriteString(fmt.Sprintf("\n\t%s.Position = %d", bookmarkIdent, bookmark.Position))
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
			codaIdent := "__common" + coda.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Coda{Name: %s}).Stage(stageSet.Stage)", codaIdent, __gong__toRawStringLiteral(coda.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", codaIdent, __gong__toRawStringLiteral(coda.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", codaIdent, __gong__toRawStringLiteral(coda.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", codaIdent, __gong__toRawStringLiteral(coda.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", codaIdent, __gong__toRawStringLiteral(coda.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", codaIdent, __gong__toRawStringLiteral(coda.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", codaIdent, __gong__toRawStringLiteral(coda.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", codaIdent, __gong__toRawStringLiteral(coda.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", codaIdent, __gong__toRawStringLiteral(string(coda.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", codaIdent, __gong__toRawStringLiteral(coda.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", codaIdent, __gong__toRawStringLiteral(string(coda.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", codaIdent, __gong__toRawStringLiteral(coda.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", codaIdent, __gong__toRawStringLiteral(string(coda.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", codaIdent, __gong__toRawStringLiteral(string(coda.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", codaIdent, __gong__toRawStringLiteral(coda.Id)))
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
			distanceIdent := "__common" + distance.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Distance{Name: %s}).Stage(stageSet.Stage)", distanceIdent, __gong__toRawStringLiteral(distance.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", distanceIdent, __gong__toRawStringLiteral(distance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", distanceIdent, __gong__toRawStringLiteral(distance.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", distanceIdent, __gong__toRawStringLiteral(distance.Value)))
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
			dynamicsIdent := "__common" + dynamics.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Dynamics{Name: %s}).Stage(stageSet.Stage)", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", dynamicsIdent, __gong__toRawStringLiteral(string(dynamics.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", dynamicsIdent, __gong__toRawStringLiteral(string(dynamics.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", dynamicsIdent, __gong__toRawStringLiteral(string(dynamics.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", dynamicsIdent, __gong__toRawStringLiteral(string(dynamics.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", dynamicsIdent, __gong__toRawStringLiteral(string(dynamics.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", dynamicsIdent, dynamics.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", dynamicsIdent, dynamics.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.LineThrough = %d", dynamicsIdent, dynamics.LineThrough))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", dynamicsIdent, __gong__toRawStringLiteral(string(dynamics.Enclosure))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Id)))
			for _, elem := range dynamics.P {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.P = append(%s.P, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Pp {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pp = append(%s.Pp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Ppp {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ppp = append(%s.Ppp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Pppp {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pppp = append(%s.Pppp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Ppppp {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ppppp = append(%s.Ppppp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Pppppp {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pppppp = append(%s.Pppppp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.F {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.F = append(%s.F, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Ff {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ff = append(%s.Ff, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Fff {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fff = append(%s.Fff, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Ffff {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ffff = append(%s.Ffff, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Fffff {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fffff = append(%s.Fffff, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Ffffff {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ffffff = append(%s.Ffffff, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Mp {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mp = append(%s.Mp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Mf {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mf = append(%s.Mf, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Sf {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sf = append(%s.Sf, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Sfp {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sfp = append(%s.Sfp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Sfpp {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sfpp = append(%s.Sfpp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Fp {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fp = append(%s.Fp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Rf {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rf = append(%s.Rf, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Rfz {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rfz = append(%s.Rfz, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Sfz {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sfz = append(%s.Sfz, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Sffz {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sffz = append(%s.Sffz, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Fz {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fz = append(%s.Fz, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.N {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.N = append(%s.N, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Pf {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pf = append(%s.Pf, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Sfzp {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sfzp = append(%s.Sfzp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.OtherDynamics {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherDynamics = append(%s.OtherDynamics, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
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
			emptyIdent := "__common" + empty.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Empty{Name: %s}).Stage(stageSet.Stage)", emptyIdent, __gong__toRawStringLiteral(empty.Name)))
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
		for _, emptyplacement := range __gong__sortStageSetInstances(stageSet.Stage.EmptyPlacements, stageSet.Stage.EmptyPlacement_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			emptyplacementIdent := "__common" + emptyplacement.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyPlacement{Name: %s}).Stage(stageSet.Stage)", emptyplacementIdent, __gong__toRawStringLiteral(emptyplacement.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", emptyplacementIdent, __gong__toRawStringLiteral(emptyplacement.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", emptyplacementIdent, __gong__toRawStringLiteral(emptyplacement.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", emptyplacementIdent, __gong__toRawStringLiteral(emptyplacement.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", emptyplacementIdent, __gong__toRawStringLiteral(emptyplacement.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", emptyplacementIdent, __gong__toRawStringLiteral(emptyplacement.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", emptyplacementIdent, __gong__toRawStringLiteral(emptyplacement.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", emptyplacementIdent, __gong__toRawStringLiteral(string(emptyplacement.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", emptyplacementIdent, __gong__toRawStringLiteral(emptyplacement.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", emptyplacementIdent, __gong__toRawStringLiteral(string(emptyplacement.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", emptyplacementIdent, __gong__toRawStringLiteral(emptyplacement.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", emptyplacementIdent, __gong__toRawStringLiteral(string(emptyplacement.Placement))))
		}
	}
	if stageSet.Stage != nil {
		for _, emptyplacementsmufl := range __gong__sortStageSetInstances(stageSet.Stage.EmptyPlacementSmufls, stageSet.Stage.EmptyPlacementSmufl_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			emptyplacementsmuflIdent := "__common" + emptyplacementsmufl.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyPlacementSmufl{Name: %s}).Stage(stageSet.Stage)", emptyplacementsmuflIdent, __gong__toRawStringLiteral(emptyplacementsmufl.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", emptyplacementsmuflIdent, __gong__toRawStringLiteral(emptyplacementsmufl.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", emptyplacementsmuflIdent, __gong__toRawStringLiteral(emptyplacementsmufl.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", emptyplacementsmuflIdent, __gong__toRawStringLiteral(emptyplacementsmufl.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", emptyplacementsmuflIdent, __gong__toRawStringLiteral(emptyplacementsmufl.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", emptyplacementsmuflIdent, __gong__toRawStringLiteral(emptyplacementsmufl.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", emptyplacementsmuflIdent, __gong__toRawStringLiteral(emptyplacementsmufl.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", emptyplacementsmuflIdent, __gong__toRawStringLiteral(string(emptyplacementsmufl.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", emptyplacementsmuflIdent, __gong__toRawStringLiteral(emptyplacementsmufl.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", emptyplacementsmuflIdent, __gong__toRawStringLiteral(string(emptyplacementsmufl.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", emptyplacementsmuflIdent, __gong__toRawStringLiteral(emptyplacementsmufl.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", emptyplacementsmuflIdent, __gong__toRawStringLiteral(string(emptyplacementsmufl.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", emptyplacementsmuflIdent, __gong__toRawStringLiteral(emptyplacementsmufl.Smufl)))
		}
	}
	if stageSet.Stage != nil {
		for _, emptyprintobjectstylealign := range __gong__sortStageSetInstances(stageSet.Stage.EmptyPrintObjectStyleAligns, stageSet.Stage.EmptyPrintObjectStyleAlign_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			emptyprintobjectstylealignIdent := "__common" + emptyprintobjectstylealign.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyPrintObjectStyleAlign{Name: %s}).Stage(stageSet.Stage)", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(emptyprintobjectstylealign.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(emptyprintobjectstylealign.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(string(emptyprintobjectstylealign.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(emptyprintobjectstylealign.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(emptyprintobjectstylealign.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(emptyprintobjectstylealign.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(emptyprintobjectstylealign.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(emptyprintobjectstylealign.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(string(emptyprintobjectstylealign.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(emptyprintobjectstylealign.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(string(emptyprintobjectstylealign.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(emptyprintobjectstylealign.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(string(emptyprintobjectstylealign.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(string(emptyprintobjectstylealign.Valign))))
		}
	}
	if stageSet.Stage != nil {
		for _, emptyprintstyle := range __gong__sortStageSetInstances(stageSet.Stage.EmptyPrintStyles, stageSet.Stage.EmptyPrintStyle_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			emptyprintstyleIdent := "__common" + emptyprintstyle.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyPrintStyle{Name: %s}).Stage(stageSet.Stage)", emptyprintstyleIdent, __gong__toRawStringLiteral(emptyprintstyle.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", emptyprintstyleIdent, __gong__toRawStringLiteral(emptyprintstyle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", emptyprintstyleIdent, __gong__toRawStringLiteral(emptyprintstyle.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", emptyprintstyleIdent, __gong__toRawStringLiteral(emptyprintstyle.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", emptyprintstyleIdent, __gong__toRawStringLiteral(emptyprintstyle.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", emptyprintstyleIdent, __gong__toRawStringLiteral(emptyprintstyle.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", emptyprintstyleIdent, __gong__toRawStringLiteral(emptyprintstyle.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", emptyprintstyleIdent, __gong__toRawStringLiteral(string(emptyprintstyle.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", emptyprintstyleIdent, __gong__toRawStringLiteral(emptyprintstyle.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", emptyprintstyleIdent, __gong__toRawStringLiteral(string(emptyprintstyle.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", emptyprintstyleIdent, __gong__toRawStringLiteral(emptyprintstyle.Color)))
		}
	}
	if stageSet.Stage != nil {
		for _, emptyprintstylealign := range __gong__sortStageSetInstances(stageSet.Stage.EmptyPrintStyleAligns, stageSet.Stage.EmptyPrintStyleAlign_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			emptyprintstylealignIdent := "__common" + emptyprintstylealign.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyPrintStyleAlign{Name: %s}).Stage(stageSet.Stage)", emptyprintstylealignIdent, __gong__toRawStringLiteral(emptyprintstylealign.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", emptyprintstylealignIdent, __gong__toRawStringLiteral(emptyprintstylealign.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", emptyprintstylealignIdent, __gong__toRawStringLiteral(emptyprintstylealign.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", emptyprintstylealignIdent, __gong__toRawStringLiteral(emptyprintstylealign.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", emptyprintstylealignIdent, __gong__toRawStringLiteral(emptyprintstylealign.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", emptyprintstylealignIdent, __gong__toRawStringLiteral(emptyprintstylealign.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", emptyprintstylealignIdent, __gong__toRawStringLiteral(emptyprintstylealign.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", emptyprintstylealignIdent, __gong__toRawStringLiteral(string(emptyprintstylealign.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", emptyprintstylealignIdent, __gong__toRawStringLiteral(emptyprintstylealign.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", emptyprintstylealignIdent, __gong__toRawStringLiteral(string(emptyprintstylealign.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", emptyprintstylealignIdent, __gong__toRawStringLiteral(emptyprintstylealign.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", emptyprintstylealignIdent, __gong__toRawStringLiteral(string(emptyprintstylealign.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", emptyprintstylealignIdent, __gong__toRawStringLiteral(string(emptyprintstylealign.Valign))))
		}
	}
	if stageSet.Stage != nil {
		for _, emptyprintstylealignid := range __gong__sortStageSetInstances(stageSet.Stage.EmptyPrintStyleAlignIds, stageSet.Stage.EmptyPrintStyleAlignId_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			emptyprintstylealignidIdent := "__common" + emptyprintstylealignid.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyPrintStyleAlignId{Name: %s}).Stage(stageSet.Stage)", emptyprintstylealignidIdent, __gong__toRawStringLiteral(emptyprintstylealignid.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", emptyprintstylealignidIdent, __gong__toRawStringLiteral(emptyprintstylealignid.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", emptyprintstylealignidIdent, __gong__toRawStringLiteral(emptyprintstylealignid.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", emptyprintstylealignidIdent, __gong__toRawStringLiteral(emptyprintstylealignid.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", emptyprintstylealignidIdent, __gong__toRawStringLiteral(emptyprintstylealignid.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", emptyprintstylealignidIdent, __gong__toRawStringLiteral(emptyprintstylealignid.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", emptyprintstylealignidIdent, __gong__toRawStringLiteral(emptyprintstylealignid.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", emptyprintstylealignidIdent, __gong__toRawStringLiteral(string(emptyprintstylealignid.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", emptyprintstylealignidIdent, __gong__toRawStringLiteral(emptyprintstylealignid.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", emptyprintstylealignidIdent, __gong__toRawStringLiteral(string(emptyprintstylealignid.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", emptyprintstylealignidIdent, __gong__toRawStringLiteral(emptyprintstylealignid.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", emptyprintstylealignidIdent, __gong__toRawStringLiteral(string(emptyprintstylealignid.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", emptyprintstylealignidIdent, __gong__toRawStringLiteral(string(emptyprintstylealignid.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", emptyprintstylealignidIdent, __gong__toRawStringLiteral(emptyprintstylealignid.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, emptytrillsound := range __gong__sortStageSetInstances(stageSet.Stage.EmptyTrillSounds, stageSet.Stage.EmptyTrillSound_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			emptytrillsoundIdent := "__common" + emptytrillsound.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyTrillSound{Name: %s}).Stage(stageSet.Stage)", emptytrillsoundIdent, __gong__toRawStringLiteral(emptytrillsound.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(emptytrillsound.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(emptytrillsound.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(emptytrillsound.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(emptytrillsound.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(emptytrillsound.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(emptytrillsound.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(string(emptytrillsound.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(emptytrillsound.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(string(emptytrillsound.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(emptytrillsound.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(string(emptytrillsound.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.StartNote = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(string(emptytrillsound.StartNote))))
			values.WriteString(fmt.Sprintf("\n\t%s.TrillStep = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(string(emptytrillsound.TrillStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.TwoNoteTurn = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(string(emptytrillsound.TwoNoteTurn))))
			values.WriteString(fmt.Sprintf("\n\t%s.Accelerate = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(string(emptytrillsound.Accelerate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(emptytrillsound.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.SecondBeat = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(emptytrillsound.SecondBeat)))
			values.WriteString(fmt.Sprintf("\n\t%s.LastBeat = %s", emptytrillsoundIdent, __gong__toRawStringLiteral(emptytrillsound.LastBeat)))
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
			encodingIdent := "__common" + encoding.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Encoding{Name: %s}).Stage(stageSet.Stage)", encodingIdent, __gong__toRawStringLiteral(encoding.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", encodingIdent, __gong__toRawStringLiteral(encoding.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.EncodingDate = %s", encodingIdent, __gong__toRawStringLiteral(encoding.EncodingDate)))
			values.WriteString(fmt.Sprintf("\n\t%s.Software = %s", encodingIdent, __gong__toRawStringLiteral(encoding.Software)))
			values.WriteString(fmt.Sprintf("\n\t%s.EncodingDescription = %s", encodingIdent, __gong__toRawStringLiteral(encoding.EncodingDescription)))
			for _, elem := range encoding.Encoder {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Encoder = append(%s.Encoder, %s)", encodingIdent, encodingIdent, targetIdent))
			}
			for _, elem := range encoding.Supports {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Supports = append(%s.Supports, %s)", encodingIdent, encodingIdent, targetIdent))
			}
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
			fermataIdent := "__common" + fermata.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Fermata{Name: %s}).Stage(stageSet.Stage)", fermataIdent, __gong__toRawStringLiteral(fermata.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", fermataIdent, __gong__toRawStringLiteral(string(fermata.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", fermataIdent, __gong__toRawStringLiteral(fermata.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", fermataIdent, __gong__toRawStringLiteral(fermata.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", fermataIdent, __gong__toRawStringLiteral(fermata.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", fermataIdent, __gong__toRawStringLiteral(fermata.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", fermataIdent, __gong__toRawStringLiteral(fermata.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", fermataIdent, __gong__toRawStringLiteral(string(fermata.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", fermataIdent, __gong__toRawStringLiteral(fermata.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", fermataIdent, __gong__toRawStringLiteral(string(fermata.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", fermataIdent, __gong__toRawStringLiteral(fermata.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", fermataIdent, __gong__toRawStringLiteral(string(fermata.Value))))
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
			fingeringIdent := "__common" + fingering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Fingering{Name: %s}).Stage(stageSet.Stage)", fingeringIdent, __gong__toRawStringLiteral(fingering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Substitution = %s", fingeringIdent, __gong__toRawStringLiteral(string(fingering.Substitution))))
			values.WriteString(fmt.Sprintf("\n\t%s.Alternate = %s", fingeringIdent, __gong__toRawStringLiteral(string(fingering.Alternate))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", fingeringIdent, __gong__toRawStringLiteral(string(fingering.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", fingeringIdent, __gong__toRawStringLiteral(string(fingering.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", fingeringIdent, __gong__toRawStringLiteral(string(fingering.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", fingeringIdent, __gong__toRawStringLiteral(fingering.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, formattedsymbol := range __gong__sortStageSetInstances(stageSet.Stage.FormattedSymbols, stageSet.Stage.FormattedSymbol_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formattedsymbolIdent := "__common" + formattedsymbol.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.FormattedSymbol{Name: %s}).Stage(stageSet.Stage)", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", formattedsymbolIdent, __gong__toRawStringLiteral(string(formattedsymbol.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", formattedsymbolIdent, __gong__toRawStringLiteral(string(formattedsymbol.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", formattedsymbolIdent, __gong__toRawStringLiteral(string(formattedsymbol.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", formattedsymbolIdent, __gong__toRawStringLiteral(string(formattedsymbol.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", formattedsymbolIdent, __gong__toRawStringLiteral(string(formattedsymbol.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", formattedsymbolIdent, formattedsymbol.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", formattedsymbolIdent, formattedsymbol.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.LineThrough = %d", formattedsymbolIdent, formattedsymbol.LineThrough))
			values.WriteString(fmt.Sprintf("\n\t%s.Rotation = %s", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.Rotation)))
			values.WriteString(fmt.Sprintf("\n\t%s.LetterSpacing = %s", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.LetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.LineHeight = %s", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.LineHeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dir = %s", formattedsymbolIdent, __gong__toRawStringLiteral(string(formattedsymbol.Dir))))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", formattedsymbolIdent, __gong__toRawStringLiteral(string(formattedsymbol.Enclosure))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, formattedsymbolid := range __gong__sortStageSetInstances(stageSet.Stage.FormattedSymbolIds, stageSet.Stage.FormattedSymbolId_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formattedsymbolidIdent := "__common" + formattedsymbolid.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.FormattedSymbolId{Name: %s}).Stage(stageSet.Stage)", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(string(formattedsymbolid.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(string(formattedsymbolid.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(string(formattedsymbolid.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(string(formattedsymbolid.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(string(formattedsymbolid.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", formattedsymbolidIdent, formattedsymbolid.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", formattedsymbolidIdent, formattedsymbolid.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.LineThrough = %d", formattedsymbolidIdent, formattedsymbolid.LineThrough))
			values.WriteString(fmt.Sprintf("\n\t%s.Rotation = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.Rotation)))
			values.WriteString(fmt.Sprintf("\n\t%s.LetterSpacing = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.LetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.LineHeight = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.LineHeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dir = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(string(formattedsymbolid.Dir))))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(string(formattedsymbolid.Enclosure))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, formattedtext := range __gong__sortStageSetInstances(stageSet.Stage.FormattedTexts, stageSet.Stage.FormattedText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formattedtextIdent := "__common" + formattedtext.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.FormattedText{Name: %s}).Stage(stageSet.Stage)", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.Space)))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", formattedtextIdent, __gong__toRawStringLiteral(string(formattedtext.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", formattedtextIdent, __gong__toRawStringLiteral(string(formattedtext.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", formattedtextIdent, __gong__toRawStringLiteral(string(formattedtext.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", formattedtextIdent, __gong__toRawStringLiteral(string(formattedtext.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", formattedtextIdent, __gong__toRawStringLiteral(string(formattedtext.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", formattedtextIdent, formattedtext.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", formattedtextIdent, formattedtext.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.LineThrough = %d", formattedtextIdent, formattedtext.LineThrough))
			values.WriteString(fmt.Sprintf("\n\t%s.Rotation = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.Rotation)))
			values.WriteString(fmt.Sprintf("\n\t%s.LetterSpacing = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.LetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.LineHeight = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.LineHeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dir = %s", formattedtextIdent, __gong__toRawStringLiteral(string(formattedtext.Dir))))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", formattedtextIdent, __gong__toRawStringLiteral(string(formattedtext.Enclosure))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, formattedtextid := range __gong__sortStageSetInstances(stageSet.Stage.FormattedTextIds, stageSet.Stage.FormattedTextId_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formattedtextidIdent := "__common" + formattedtextid.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.FormattedTextId{Name: %s}).Stage(stageSet.Stage)", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.Space = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.Space)))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", formattedtextidIdent, __gong__toRawStringLiteral(string(formattedtextid.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", formattedtextidIdent, __gong__toRawStringLiteral(string(formattedtextid.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", formattedtextidIdent, __gong__toRawStringLiteral(string(formattedtextid.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", formattedtextidIdent, __gong__toRawStringLiteral(string(formattedtextid.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", formattedtextidIdent, __gong__toRawStringLiteral(string(formattedtextid.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Underline = %d", formattedtextidIdent, formattedtextid.Underline))
			values.WriteString(fmt.Sprintf("\n\t%s.Overline = %d", formattedtextidIdent, formattedtextid.Overline))
			values.WriteString(fmt.Sprintf("\n\t%s.LineThrough = %d", formattedtextidIdent, formattedtextid.LineThrough))
			values.WriteString(fmt.Sprintf("\n\t%s.Rotation = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.Rotation)))
			values.WriteString(fmt.Sprintf("\n\t%s.LetterSpacing = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.LetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.LineHeight = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.LineHeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dir = %s", formattedtextidIdent, __gong__toRawStringLiteral(string(formattedtextid.Dir))))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", formattedtextidIdent, __gong__toRawStringLiteral(string(formattedtextid.Enclosure))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.Value)))
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
			fretIdent := "__common" + fret.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Fret{Name: %s}).Stage(stageSet.Stage)", fretIdent, __gong__toRawStringLiteral(fret.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fretIdent, __gong__toRawStringLiteral(fret.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", fretIdent, __gong__toRawStringLiteral(fret.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", fretIdent, __gong__toRawStringLiteral(string(fret.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", fretIdent, __gong__toRawStringLiteral(fret.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", fretIdent, __gong__toRawStringLiteral(string(fret.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", fretIdent, __gong__toRawStringLiteral(fret.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", fretIdent, fret.Value))
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
			glyphIdent := "__common" + glyph.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Glyph{Name: %s}).Stage(stageSet.Stage)", glyphIdent, __gong__toRawStringLiteral(glyph.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glyphIdent, __gong__toRawStringLiteral(glyph.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", glyphIdent, __gong__toRawStringLiteral(glyph.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", glyphIdent, __gong__toRawStringLiteral(glyph.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, horizontalturn := range __gong__sortStageSetInstances(stageSet.Stage.HorizontalTurns, stageSet.Stage.HorizontalTurn_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			horizontalturnIdent := "__common" + horizontalturn.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.HorizontalTurn{Name: %s}).Stage(stageSet.Stage)", horizontalturnIdent, __gong__toRawStringLiteral(horizontalturn.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", horizontalturnIdent, __gong__toRawStringLiteral(horizontalturn.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Slash = %s", horizontalturnIdent, __gong__toRawStringLiteral(string(horizontalturn.Slash))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", horizontalturnIdent, __gong__toRawStringLiteral(horizontalturn.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", horizontalturnIdent, __gong__toRawStringLiteral(horizontalturn.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", horizontalturnIdent, __gong__toRawStringLiteral(horizontalturn.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", horizontalturnIdent, __gong__toRawStringLiteral(horizontalturn.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", horizontalturnIdent, __gong__toRawStringLiteral(horizontalturn.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", horizontalturnIdent, __gong__toRawStringLiteral(string(horizontalturn.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", horizontalturnIdent, __gong__toRawStringLiteral(horizontalturn.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", horizontalturnIdent, __gong__toRawStringLiteral(string(horizontalturn.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", horizontalturnIdent, __gong__toRawStringLiteral(horizontalturn.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", horizontalturnIdent, __gong__toRawStringLiteral(string(horizontalturn.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.StartNote = %s", horizontalturnIdent, __gong__toRawStringLiteral(string(horizontalturn.StartNote))))
			values.WriteString(fmt.Sprintf("\n\t%s.TrillStep = %s", horizontalturnIdent, __gong__toRawStringLiteral(string(horizontalturn.TrillStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.TwoNoteTurn = %s", horizontalturnIdent, __gong__toRawStringLiteral(string(horizontalturn.TwoNoteTurn))))
			values.WriteString(fmt.Sprintf("\n\t%s.Accelerate = %s", horizontalturnIdent, __gong__toRawStringLiteral(string(horizontalturn.Accelerate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", horizontalturnIdent, __gong__toRawStringLiteral(horizontalturn.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.SecondBeat = %s", horizontalturnIdent, __gong__toRawStringLiteral(horizontalturn.SecondBeat)))
			values.WriteString(fmt.Sprintf("\n\t%s.LastBeat = %s", horizontalturnIdent, __gong__toRawStringLiteral(horizontalturn.LastBeat)))
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
			identificationIdent := "__common" + identification.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Identification{Name: %s}).Stage(stageSet.Stage)", identificationIdent, __gong__toRawStringLiteral(identification.Name)))
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
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Creator = append(%s.Creator, %s)", identificationIdent, identificationIdent, targetIdent))
			}
			for _, elem := range identification.Rights {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rights = append(%s.Rights, %s)", identificationIdent, identificationIdent, targetIdent))
			}
			if identification.Encoding != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + identification.Encoding.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Encoding = %s", identificationIdent, targetIdent))
			}
			for _, elem := range identification.Relation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Relation = append(%s.Relation, %s)", identificationIdent, identificationIdent, targetIdent))
			}
			if identification.Miscellaneous != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + identification.Miscellaneous.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Miscellaneous = %s", identificationIdent, targetIdent))
			}
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
			levelIdent := "__common" + level.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Level{Name: %s}).Stage(stageSet.Stage)", levelIdent, __gong__toRawStringLiteral(level.Name)))
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
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", levelIdent, __gong__toRawStringLiteral(level.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, linewidth := range __gong__sortStageSetInstances(stageSet.Stage.LineWidths, stageSet.Stage.LineWidth_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linewidthIdent := "__common" + linewidth.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.LineWidth{Name: %s}).Stage(stageSet.Stage)", linewidthIdent, __gong__toRawStringLiteral(linewidth.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linewidthIdent, __gong__toRawStringLiteral(linewidth.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", linewidthIdent, __gong__toRawStringLiteral(linewidth.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", linewidthIdent, __gong__toRawStringLiteral(linewidth.Value)))
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
			linkIdent := "__common" + link.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Link{Name: %s}).Stage(stageSet.Stage)", linkIdent, __gong__toRawStringLiteral(link.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linkIdent, __gong__toRawStringLiteral(link.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameAttr = %s", linkIdent, __gong__toRawStringLiteral(link.NameAttr)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", linkIdent, __gong__toRawStringLiteral(link.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", linkIdent, __gong__toRawStringLiteral(string(link.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", linkIdent, __gong__toRawStringLiteral(link.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", linkIdent, __gong__toRawStringLiteral(link.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", linkIdent, __gong__toRawStringLiteral(string(link.Show))))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", linkIdent, __gong__toRawStringLiteral(string(link.Actuate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Element = %s", linkIdent, __gong__toRawStringLiteral(link.Element)))
			values.WriteString(fmt.Sprintf("\n\t%s.Position = %d", linkIdent, link.Position))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", linkIdent, __gong__toRawStringLiteral(link.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", linkIdent, __gong__toRawStringLiteral(link.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", linkIdent, __gong__toRawStringLiteral(link.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", linkIdent, __gong__toRawStringLiteral(link.RelativeY)))
		}
	}
	if stageSet.Stage != nil {
		for _, measurelayout := range __gong__sortStageSetInstances(stageSet.Stage.MeasureLayouts, stageSet.Stage.MeasureLayout_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			measurelayoutIdent := "__common" + measurelayout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.MeasureLayout{Name: %s}).Stage(stageSet.Stage)", measurelayoutIdent, __gong__toRawStringLiteral(measurelayout.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measurelayoutIdent, __gong__toRawStringLiteral(measurelayout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MeasureDistance = %s", measurelayoutIdent, __gong__toRawStringLiteral(measurelayout.MeasureDistance)))
		}
	}
	if stageSet.Stage != nil {
		for _, mididevice := range __gong__sortStageSetInstances(stageSet.Stage.MidiDevices, stageSet.Stage.MidiDevice_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			midideviceIdent := "__common" + mididevice.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.MidiDevice{Name: %s}).Stage(stageSet.Stage)", midideviceIdent, __gong__toRawStringLiteral(mididevice.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", midideviceIdent, __gong__toRawStringLiteral(mididevice.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Port = %d", midideviceIdent, mididevice.Port))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", midideviceIdent, __gong__toRawStringLiteral(mididevice.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", midideviceIdent, __gong__toRawStringLiteral(mididevice.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, midiinstrument := range __gong__sortStageSetInstances(stageSet.Stage.MidiInstruments, stageSet.Stage.MidiInstrument_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			midiinstrumentIdent := "__common" + midiinstrument.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.MidiInstrument{Name: %s}).Stage(stageSet.Stage)", midiinstrumentIdent, __gong__toRawStringLiteral(midiinstrument.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", midiinstrumentIdent, __gong__toRawStringLiteral(midiinstrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", midiinstrumentIdent, __gong__toRawStringLiteral(midiinstrument.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.MidiChannel = %d", midiinstrumentIdent, midiinstrument.MidiChannel))
			values.WriteString(fmt.Sprintf("\n\t%s.MidiName = %s", midiinstrumentIdent, __gong__toRawStringLiteral(midiinstrument.MidiName)))
			values.WriteString(fmt.Sprintf("\n\t%s.MidiBank = %d", midiinstrumentIdent, midiinstrument.MidiBank))
			values.WriteString(fmt.Sprintf("\n\t%s.MidiProgram = %d", midiinstrumentIdent, midiinstrument.MidiProgram))
			values.WriteString(fmt.Sprintf("\n\t%s.MidiUnpitched = %d", midiinstrumentIdent, midiinstrument.MidiUnpitched))
			values.WriteString(fmt.Sprintf("\n\t%s.Volume = %s", midiinstrumentIdent, __gong__toRawStringLiteral(midiinstrument.Volume)))
			values.WriteString(fmt.Sprintf("\n\t%s.Pan = %s", midiinstrumentIdent, __gong__toRawStringLiteral(midiinstrument.Pan)))
			values.WriteString(fmt.Sprintf("\n\t%s.Elevation = %s", midiinstrumentIdent, __gong__toRawStringLiteral(midiinstrument.Elevation)))
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
			miscellaneousIdent := "__common" + miscellaneous.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Miscellaneous{Name: %s}).Stage(stageSet.Stage)", miscellaneousIdent, __gong__toRawStringLiteral(miscellaneous.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", miscellaneousIdent, __gong__toRawStringLiteral(miscellaneous.Name)))
			for _, elem := range miscellaneous.MiscellaneousField {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MiscellaneousField = append(%s.MiscellaneousField, %s)", miscellaneousIdent, miscellaneousIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, miscellaneousfield := range __gong__sortStageSetInstances(stageSet.Stage.MiscellaneousFields, stageSet.Stage.MiscellaneousField_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			miscellaneousfieldIdent := "__common" + miscellaneousfield.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.MiscellaneousField{Name: %s}).Stage(stageSet.Stage)", miscellaneousfieldIdent, __gong__toRawStringLiteral(miscellaneousfield.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", miscellaneousfieldIdent, __gong__toRawStringLiteral(miscellaneousfield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameAttr = %s", miscellaneousfieldIdent, __gong__toRawStringLiteral(miscellaneousfield.NameAttr)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", miscellaneousfieldIdent, __gong__toRawStringLiteral(miscellaneousfield.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, namedisplay := range __gong__sortStageSetInstances(stageSet.Stage.NameDisplays, stageSet.Stage.NameDisplay_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			namedisplayIdent := "__common" + namedisplay.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.NameDisplay{Name: %s}).Stage(stageSet.Stage)", namedisplayIdent, __gong__toRawStringLiteral(namedisplay.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", namedisplayIdent, __gong__toRawStringLiteral(namedisplay.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", namedisplayIdent, __gong__toRawStringLiteral(string(namedisplay.PrintObject))))
			for _, elem := range namedisplay.DisplayText {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DisplayText = append(%s.DisplayText, %s)", namedisplayIdent, namedisplayIdent, targetIdent))
			}
			for _, elem := range namedisplay.AccidentalText {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AccidentalText = append(%s.AccidentalText, %s)", namedisplayIdent, namedisplayIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, notesize := range __gong__sortStageSetInstances(stageSet.Stage.NoteSizes, stageSet.Stage.NoteSize_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notesizeIdent := "__common" + notesize.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.NoteSize{Name: %s}).Stage(stageSet.Stage)", notesizeIdent, __gong__toRawStringLiteral(notesize.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notesizeIdent, __gong__toRawStringLiteral(notesize.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", notesizeIdent, __gong__toRawStringLiteral(string(notesize.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", notesizeIdent, __gong__toRawStringLiteral(notesize.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, otherappearance := range __gong__sortStageSetInstances(stageSet.Stage.OtherAppearances, stageSet.Stage.OtherAppearance_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			otherappearanceIdent := "__common" + otherappearance.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.OtherAppearance{Name: %s}).Stage(stageSet.Stage)", otherappearanceIdent, __gong__toRawStringLiteral(otherappearance.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", otherappearanceIdent, __gong__toRawStringLiteral(otherappearance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", otherappearanceIdent, __gong__toRawStringLiteral(otherappearance.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", otherappearanceIdent, __gong__toRawStringLiteral(otherappearance.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, otherplay := range __gong__sortStageSetInstances(stageSet.Stage.OtherPlays, stageSet.Stage.OtherPlay_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			otherplayIdent := "__common" + otherplay.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.OtherPlay{Name: %s}).Stage(stageSet.Stage)", otherplayIdent, __gong__toRawStringLiteral(otherplay.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", otherplayIdent, __gong__toRawStringLiteral(otherplay.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", otherplayIdent, __gong__toRawStringLiteral(otherplay.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", otherplayIdent, __gong__toRawStringLiteral(otherplay.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, othertext := range __gong__sortStageSetInstances(stageSet.Stage.OtherTexts, stageSet.Stage.OtherText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			othertextIdent := "__common" + othertext.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.OtherText{Name: %s}).Stage(stageSet.Stage)", othertextIdent, __gong__toRawStringLiteral(othertext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", othertextIdent, __gong__toRawStringLiteral(othertext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", othertextIdent, __gong__toRawStringLiteral(othertext.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", othertextIdent, __gong__toRawStringLiteral(othertext.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, pagelayout := range __gong__sortStageSetInstances(stageSet.Stage.PageLayouts, stageSet.Stage.PageLayout_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pagelayoutIdent := "__common" + pagelayout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.PageLayout{Name: %s}).Stage(stageSet.Stage)", pagelayoutIdent, __gong__toRawStringLiteral(pagelayout.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pagelayoutIdent, __gong__toRawStringLiteral(pagelayout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PageHeight = %s", pagelayoutIdent, __gong__toRawStringLiteral(pagelayout.PageHeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.PageWidth = %s", pagelayoutIdent, __gong__toRawStringLiteral(pagelayout.PageWidth)))
			for _, elem := range pagelayout.PageMargins {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PageMargins = append(%s.PageMargins, %s)", pagelayoutIdent, pagelayoutIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, pagemargins := range __gong__sortStageSetInstances(stageSet.Stage.PageMarginss, stageSet.Stage.PageMargins_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pagemarginsIdent := "__common" + pagemargins.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.PageMargins{Name: %s}).Stage(stageSet.Stage)", pagemarginsIdent, __gong__toRawStringLiteral(pagemargins.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pagemarginsIdent, __gong__toRawStringLiteral(pagemargins.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", pagemarginsIdent, __gong__toRawStringLiteral(string(pagemargins.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.LeftMargin = %s", pagemarginsIdent, __gong__toRawStringLiteral(pagemargins.LeftMargin)))
			values.WriteString(fmt.Sprintf("\n\t%s.RightMargin = %s", pagemarginsIdent, __gong__toRawStringLiteral(pagemargins.RightMargin)))
			values.WriteString(fmt.Sprintf("\n\t%s.TopMargin = %s", pagemarginsIdent, __gong__toRawStringLiteral(pagemargins.TopMargin)))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomMargin = %s", pagemarginsIdent, __gong__toRawStringLiteral(pagemargins.BottomMargin)))
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
			playIdent := "__common" + play.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Play{Name: %s}).Stage(stageSet.Stage)", playIdent, __gong__toRawStringLiteral(play.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", playIdent, __gong__toRawStringLiteral(play.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", playIdent, __gong__toRawStringLiteral(play.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ipa = %s", playIdent, __gong__toRawStringLiteral(play.Ipa)))
			values.WriteString(fmt.Sprintf("\n\t%s.Mute = %s", playIdent, __gong__toRawStringLiteral(string(play.Mute))))
			values.WriteString(fmt.Sprintf("\n\t%s.SemiPitched = %s", playIdent, __gong__toRawStringLiteral(string(play.SemiPitched))))
			for _, elem := range play.OtherPlay {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherPlay = append(%s.OtherPlay, %s)", playIdent, playIdent, targetIdent))
			}
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
			scalingIdent := "__common" + scaling.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Scaling{Name: %s}).Stage(stageSet.Stage)", scalingIdent, __gong__toRawStringLiteral(scaling.Name)))
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
		for _, segno := range __gong__sortStageSetInstances(stageSet.Stage.Segnos, stageSet.Stage.Segno_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			segnoIdent := "__common" + segno.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Segno{Name: %s}).Stage(stageSet.Stage)", segnoIdent, __gong__toRawStringLiteral(segno.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", segnoIdent, __gong__toRawStringLiteral(segno.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", segnoIdent, __gong__toRawStringLiteral(segno.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", segnoIdent, __gong__toRawStringLiteral(segno.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", segnoIdent, __gong__toRawStringLiteral(segno.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", segnoIdent, __gong__toRawStringLiteral(segno.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", segnoIdent, __gong__toRawStringLiteral(segno.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", segnoIdent, __gong__toRawStringLiteral(segno.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", segnoIdent, __gong__toRawStringLiteral(string(segno.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", segnoIdent, __gong__toRawStringLiteral(segno.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", segnoIdent, __gong__toRawStringLiteral(string(segno.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", segnoIdent, __gong__toRawStringLiteral(segno.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", segnoIdent, __gong__toRawStringLiteral(string(segno.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", segnoIdent, __gong__toRawStringLiteral(string(segno.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", segnoIdent, __gong__toRawStringLiteral(segno.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, stafflayout := range __gong__sortStageSetInstances(stageSet.Stage.StaffLayouts, stageSet.Stage.StaffLayout_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stafflayoutIdent := "__common" + stafflayout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.StaffLayout{Name: %s}).Stage(stageSet.Stage)", stafflayoutIdent, __gong__toRawStringLiteral(stafflayout.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stafflayoutIdent, __gong__toRawStringLiteral(stafflayout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", stafflayoutIdent, stafflayout.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.StaffDistance = %s", stafflayoutIdent, __gong__toRawStringLiteral(stafflayout.StaffDistance)))
		}
	}
	if stageSet.Stage != nil {
		for _, stringtype := range __gong__sortStageSetInstances(stageSet.Stage.StringTypes, stageSet.Stage.StringType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stringtypeIdent := "__common" + stringtype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.StringType{Name: %s}).Stage(stageSet.Stage)", stringtypeIdent, __gong__toRawStringLiteral(stringtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stringtypeIdent, __gong__toRawStringLiteral(stringtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", stringtypeIdent, __gong__toRawStringLiteral(stringtype.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", stringtypeIdent, __gong__toRawStringLiteral(stringtype.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", stringtypeIdent, __gong__toRawStringLiteral(stringtype.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", stringtypeIdent, __gong__toRawStringLiteral(stringtype.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", stringtypeIdent, __gong__toRawStringLiteral(stringtype.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", stringtypeIdent, __gong__toRawStringLiteral(string(stringtype.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", stringtypeIdent, __gong__toRawStringLiteral(stringtype.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", stringtypeIdent, __gong__toRawStringLiteral(string(stringtype.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", stringtypeIdent, __gong__toRawStringLiteral(stringtype.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", stringtypeIdent, __gong__toRawStringLiteral(string(stringtype.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", stringtypeIdent, stringtype.Value))
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
			supportsIdent := "__common" + supports.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Supports{Name: %s}).Stage(stageSet.Stage)", supportsIdent, __gong__toRawStringLiteral(supports.Name)))
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
		for _, systemdividers := range __gong__sortStageSetInstances(stageSet.Stage.SystemDividerss, stageSet.Stage.SystemDividers_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			systemdividersIdent := "__common" + systemdividers.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.SystemDividers{Name: %s}).Stage(stageSet.Stage)", systemdividersIdent, __gong__toRawStringLiteral(systemdividers.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", systemdividersIdent, __gong__toRawStringLiteral(systemdividers.Name)))
			if systemdividers.LeftDivider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + systemdividers.LeftDivider.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LeftDivider = %s", systemdividersIdent, targetIdent))
			}
			if systemdividers.RightDivider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + systemdividers.RightDivider.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RightDivider = %s", systemdividersIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, systemlayout := range __gong__sortStageSetInstances(stageSet.Stage.SystemLayouts, stageSet.Stage.SystemLayout_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			systemlayoutIdent := "__common" + systemlayout.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.SystemLayout{Name: %s}).Stage(stageSet.Stage)", systemlayoutIdent, __gong__toRawStringLiteral(systemlayout.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", systemlayoutIdent, __gong__toRawStringLiteral(systemlayout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SystemDistance = %s", systemlayoutIdent, __gong__toRawStringLiteral(systemlayout.SystemDistance)))
			values.WriteString(fmt.Sprintf("\n\t%s.TopSystemDistance = %s", systemlayoutIdent, __gong__toRawStringLiteral(systemlayout.TopSystemDistance)))
			if systemlayout.SystemMargins != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + systemlayout.SystemMargins.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemMargins = %s", systemlayoutIdent, targetIdent))
			}
			if systemlayout.SystemDividers != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__common" + systemlayout.SystemDividers.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemDividers = %s", systemlayoutIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, systemmargins := range __gong__sortStageSetInstances(stageSet.Stage.SystemMarginss, stageSet.Stage.SystemMargins_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			systemmarginsIdent := "__common" + systemmargins.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.SystemMargins{Name: %s}).Stage(stageSet.Stage)", systemmarginsIdent, __gong__toRawStringLiteral(systemmargins.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", systemmarginsIdent, __gong__toRawStringLiteral(systemmargins.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.LeftMargin = %s", systemmarginsIdent, __gong__toRawStringLiteral(systemmargins.LeftMargin)))
			values.WriteString(fmt.Sprintf("\n\t%s.RightMargin = %s", systemmarginsIdent, __gong__toRawStringLiteral(systemmargins.RightMargin)))
		}
	}
	if stageSet.Stage != nil {
		for _, typedtext := range __gong__sortStageSetInstances(stageSet.Stage.TypedTexts, stageSet.Stage.TypedText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			typedtextIdent := "__common" + typedtext.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.TypedText{Name: %s}).Stage(stageSet.Stage)", typedtextIdent, __gong__toRawStringLiteral(typedtext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", typedtextIdent, __gong__toRawStringLiteral(typedtext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", typedtextIdent, __gong__toRawStringLiteral(typedtext.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", typedtextIdent, __gong__toRawStringLiteral(typedtext.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, wavyline := range __gong__sortStageSetInstances(stageSet.Stage.WavyLines, stageSet.Stage.WavyLine_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			wavylineIdent := "__common" + wavyline.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.WavyLine{Name: %s}).Stage(stageSet.Stage)", wavylineIdent, __gong__toRawStringLiteral(wavyline.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", wavylineIdent, __gong__toRawStringLiteral(wavyline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", wavylineIdent, __gong__toRawStringLiteral(string(wavyline.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", wavylineIdent, wavyline.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", wavylineIdent, __gong__toRawStringLiteral(wavyline.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", wavylineIdent, __gong__toRawStringLiteral(wavyline.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", wavylineIdent, __gong__toRawStringLiteral(wavyline.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", wavylineIdent, __gong__toRawStringLiteral(wavyline.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", wavylineIdent, __gong__toRawStringLiteral(wavyline.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", wavylineIdent, __gong__toRawStringLiteral(string(wavyline.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", wavylineIdent, __gong__toRawStringLiteral(wavyline.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.StartNote = %s", wavylineIdent, __gong__toRawStringLiteral(string(wavyline.StartNote))))
			values.WriteString(fmt.Sprintf("\n\t%s.TrillStep = %s", wavylineIdent, __gong__toRawStringLiteral(string(wavyline.TrillStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.TwoNoteTurn = %s", wavylineIdent, __gong__toRawStringLiteral(string(wavyline.TwoNoteTurn))))
			values.WriteString(fmt.Sprintf("\n\t%s.Accelerate = %s", wavylineIdent, __gong__toRawStringLiteral(string(wavyline.Accelerate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", wavylineIdent, __gong__toRawStringLiteral(wavyline.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.SecondBeat = %s", wavylineIdent, __gong__toRawStringLiteral(wavyline.SecondBeat)))
			values.WriteString(fmt.Sprintf("\n\t%s.LastBeat = %s", wavylineIdent, __gong__toRawStringLiteral(wavyline.LastBeat)))
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
go run github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *common.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *common.StageSet) {

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
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common":
			aliasToCanonical[alias] = "common"
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
			case "common":
				switch typeName {
				case "AccidentalText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AccidentalText), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Appearance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Appearance), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Bookmark":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bookmark), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Coda":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Coda), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Distance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Distance), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Dynamics":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Dynamics), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Empty":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Empty), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "EmptyPlacement":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EmptyPlacement), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "EmptyPlacementSmufl":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EmptyPlacementSmufl), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "EmptyPrintObjectStyleAlign":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EmptyPrintObjectStyleAlign), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "EmptyPrintStyle":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EmptyPrintStyle), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "EmptyPrintStyleAlign":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EmptyPrintStyleAlign), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "EmptyPrintStyleAlignId":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EmptyPrintStyleAlignId), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "EmptyTrillSound":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EmptyTrillSound), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Encoding":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Encoding), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Fermata":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Fermata), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Fingering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Fingering), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "FormattedSymbol":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormattedSymbol), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "FormattedSymbolId":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormattedSymbolId), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "FormattedText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormattedText), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "FormattedTextId":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FormattedTextId), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Fret":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Fret), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Glyph":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Glyph), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "HorizontalTurn":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HorizontalTurn), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Identification":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Identification), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Level":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Level), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "LineWidth":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LineWidth), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Link":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Link), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MeasureLayout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeasureLayout), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MidiDevice":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MidiDevice), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MidiInstrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MidiInstrument), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Miscellaneous":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Miscellaneous), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MiscellaneousField":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MiscellaneousField), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NameDisplay":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NameDisplay), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NoteSize":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoteSize), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "OtherAppearance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OtherAppearance), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "OtherPlay":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OtherPlay), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "OtherText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OtherText), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PageLayout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PageLayout), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PageMargins":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PageMargins), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Play":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Play), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Scaling":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Scaling), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Segno":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Segno), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StaffLayout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StaffLayout), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StringType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StringType), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Supports":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Supports), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SystemDividers":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SystemDividers), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SystemLayout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SystemLayout), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SystemMargins":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SystemMargins), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TypedText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TypedText), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "WavyLine":
					identifierMap[ident.Name] = __gong__stageSetInit(new(WavyLine), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *AccidentalText:
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
						inst.Justify = EnumLeftCenterRight(GongExtractString(rhs))
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = EnumValign(GongExtractString(rhs))
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
					case "LineHeight":
						inst.LineHeight = GongExtractString(rhs)
					case "Dir":
						inst.Dir = EnumTextDirection(GongExtractString(rhs))
					case "Enclosure":
						inst.Enclosure = EnumEnclosureShape(GongExtractString(rhs))
					case "Value":
						inst.Value = EnumAccidentalValue(GongExtractString(rhs))
					}
				case *Appearance:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "LineWidth":
						__gong__assignSliceOfPointers(&inst.LineWidth, rhs, identifierMap)
					case "NoteSize":
						__gong__assignSliceOfPointers(&inst.NoteSize, rhs, identifierMap)
					case "Distance":
						__gong__assignSliceOfPointers(&inst.Distance, rhs, identifierMap)
					case "Glyph":
						__gong__assignSliceOfPointers(&inst.Glyph, rhs, identifierMap)
					case "OtherAppearance":
						__gong__assignSliceOfPointers(&inst.OtherAppearance, rhs, identifierMap)
					}
				case *Bookmark:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "NameAttr":
						inst.NameAttr = GongExtractString(rhs)
					case "Element":
						inst.Element = GongExtractString(rhs)
					case "Position":
						inst.Position = GongExtractInt(rhs)
					}
				case *Coda:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = EnumValign(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Distance:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Dynamics:
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = EnumValign(GongExtractString(rhs))
					case "Placement":
						inst.Placement = EnumAboveBelow(GongExtractString(rhs))
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "LineThrough":
						inst.LineThrough = GongExtractInt(rhs)
					case "Enclosure":
						inst.Enclosure = EnumEnclosureShape(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "P":
						__gong__assignSliceOfPointers(&inst.P, rhs, identifierMap)
					case "Pp":
						__gong__assignSliceOfPointers(&inst.Pp, rhs, identifierMap)
					case "Ppp":
						__gong__assignSliceOfPointers(&inst.Ppp, rhs, identifierMap)
					case "Pppp":
						__gong__assignSliceOfPointers(&inst.Pppp, rhs, identifierMap)
					case "Ppppp":
						__gong__assignSliceOfPointers(&inst.Ppppp, rhs, identifierMap)
					case "Pppppp":
						__gong__assignSliceOfPointers(&inst.Pppppp, rhs, identifierMap)
					case "F":
						__gong__assignSliceOfPointers(&inst.F, rhs, identifierMap)
					case "Ff":
						__gong__assignSliceOfPointers(&inst.Ff, rhs, identifierMap)
					case "Fff":
						__gong__assignSliceOfPointers(&inst.Fff, rhs, identifierMap)
					case "Ffff":
						__gong__assignSliceOfPointers(&inst.Ffff, rhs, identifierMap)
					case "Fffff":
						__gong__assignSliceOfPointers(&inst.Fffff, rhs, identifierMap)
					case "Ffffff":
						__gong__assignSliceOfPointers(&inst.Ffffff, rhs, identifierMap)
					case "Mp":
						__gong__assignSliceOfPointers(&inst.Mp, rhs, identifierMap)
					case "Mf":
						__gong__assignSliceOfPointers(&inst.Mf, rhs, identifierMap)
					case "Sf":
						__gong__assignSliceOfPointers(&inst.Sf, rhs, identifierMap)
					case "Sfp":
						__gong__assignSliceOfPointers(&inst.Sfp, rhs, identifierMap)
					case "Sfpp":
						__gong__assignSliceOfPointers(&inst.Sfpp, rhs, identifierMap)
					case "Fp":
						__gong__assignSliceOfPointers(&inst.Fp, rhs, identifierMap)
					case "Rf":
						__gong__assignSliceOfPointers(&inst.Rf, rhs, identifierMap)
					case "Rfz":
						__gong__assignSliceOfPointers(&inst.Rfz, rhs, identifierMap)
					case "Sfz":
						__gong__assignSliceOfPointers(&inst.Sfz, rhs, identifierMap)
					case "Sffz":
						__gong__assignSliceOfPointers(&inst.Sffz, rhs, identifierMap)
					case "Fz":
						__gong__assignSliceOfPointers(&inst.Fz, rhs, identifierMap)
					case "N":
						__gong__assignSliceOfPointers(&inst.N, rhs, identifierMap)
					case "Pf":
						__gong__assignSliceOfPointers(&inst.Pf, rhs, identifierMap)
					case "Sfzp":
						__gong__assignSliceOfPointers(&inst.Sfzp, rhs, identifierMap)
					case "OtherDynamics":
						__gong__assignSliceOfPointers(&inst.OtherDynamics, rhs, identifierMap)
					}
				case *Empty:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *EmptyPlacement:
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = EnumAboveBelow(GongExtractString(rhs))
					}
				case *EmptyPlacementSmufl:
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = EnumAboveBelow(GongExtractString(rhs))
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					}
				case *EmptyPrintObjectStyleAlign:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PrintObject":
						inst.PrintObject = EnumYesNo(GongExtractString(rhs))
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = EnumValign(GongExtractString(rhs))
					}
				case *EmptyPrintStyle:
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *EmptyPrintStyleAlign:
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = EnumValign(GongExtractString(rhs))
					}
				case *EmptyPrintStyleAlignId:
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = EnumValign(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *EmptyTrillSound:
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = EnumAboveBelow(GongExtractString(rhs))
					case "StartNote":
						inst.StartNote = EnumStartNote(GongExtractString(rhs))
					case "TrillStep":
						inst.TrillStep = EnumTrillStep(GongExtractString(rhs))
					case "TwoNoteTurn":
						inst.TwoNoteTurn = EnumTwoNoteTurn(GongExtractString(rhs))
					case "Accelerate":
						inst.Accelerate = EnumYesNo(GongExtractString(rhs))
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "SecondBeat":
						inst.SecondBeat = GongExtractString(rhs)
					case "LastBeat":
						inst.LastBeat = GongExtractString(rhs)
					}
				case *Encoding:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "EncodingDate":
						inst.EncodingDate = GongExtractString(rhs)
					case "Encoder":
						__gong__assignSliceOfPointers(&inst.Encoder, rhs, identifierMap)
					case "Software":
						inst.Software = GongExtractString(rhs)
					case "EncodingDescription":
						inst.EncodingDescription = GongExtractString(rhs)
					case "Supports":
						__gong__assignSliceOfPointers(&inst.Supports, rhs, identifierMap)
					}
				case *Fermata:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumUprightInverted(GongExtractString(rhs))
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = EnumFermataShape(GongExtractString(rhs))
					}
				case *Fingering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Substitution":
						inst.Substitution = EnumYesNo(GongExtractString(rhs))
					case "Alternate":
						inst.Alternate = EnumYesNo(GongExtractString(rhs))
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = EnumAboveBelow(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *FormattedSymbol:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Justify":
						inst.Justify = EnumLeftCenterRight(GongExtractString(rhs))
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = EnumValign(GongExtractString(rhs))
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
					case "LineHeight":
						inst.LineHeight = GongExtractString(rhs)
					case "Dir":
						inst.Dir = EnumTextDirection(GongExtractString(rhs))
					case "Enclosure":
						inst.Enclosure = EnumEnclosureShape(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *FormattedSymbolId:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Justify":
						inst.Justify = EnumLeftCenterRight(GongExtractString(rhs))
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = EnumValign(GongExtractString(rhs))
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
					case "LineHeight":
						inst.LineHeight = GongExtractString(rhs)
					case "Dir":
						inst.Dir = EnumTextDirection(GongExtractString(rhs))
					case "Enclosure":
						inst.Enclosure = EnumEnclosureShape(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *FormattedText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Space":
						inst.Space = GongExtractString(rhs)
					case "Justify":
						inst.Justify = EnumLeftCenterRight(GongExtractString(rhs))
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = EnumValign(GongExtractString(rhs))
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
					case "LineHeight":
						inst.LineHeight = GongExtractString(rhs)
					case "Dir":
						inst.Dir = EnumTextDirection(GongExtractString(rhs))
					case "Enclosure":
						inst.Enclosure = EnumEnclosureShape(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *FormattedTextId:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Space":
						inst.Space = GongExtractString(rhs)
					case "Justify":
						inst.Justify = EnumLeftCenterRight(GongExtractString(rhs))
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = EnumValign(GongExtractString(rhs))
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
					case "LineHeight":
						inst.LineHeight = GongExtractString(rhs)
					case "Dir":
						inst.Dir = EnumTextDirection(GongExtractString(rhs))
					case "Enclosure":
						inst.Enclosure = EnumEnclosureShape(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Fret:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractInt(rhs)
					}
				case *Glyph:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *HorizontalTurn:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Slash":
						inst.Slash = EnumYesNo(GongExtractString(rhs))
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = EnumAboveBelow(GongExtractString(rhs))
					case "StartNote":
						inst.StartNote = EnumStartNote(GongExtractString(rhs))
					case "TrillStep":
						inst.TrillStep = EnumTrillStep(GongExtractString(rhs))
					case "TwoNoteTurn":
						inst.TwoNoteTurn = EnumTwoNoteTurn(GongExtractString(rhs))
					case "Accelerate":
						inst.Accelerate = EnumYesNo(GongExtractString(rhs))
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "SecondBeat":
						inst.SecondBeat = GongExtractString(rhs)
					case "LastBeat":
						inst.LastBeat = GongExtractString(rhs)
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
				case *Level:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Reference":
						inst.Reference = EnumYesNo(GongExtractString(rhs))
					case "Type":
						inst.Type = EnumStartStopSingle(GongExtractString(rhs))
					case "Parentheses":
						inst.Parentheses = EnumYesNo(GongExtractString(rhs))
					case "Bracket":
						inst.Bracket = EnumYesNo(GongExtractString(rhs))
					case "Size":
						inst.Size = EnumSymbolSize(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *LineWidth:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Link:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NameAttr":
						inst.NameAttr = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Type":
						inst.Type = xlink.EnumTypeType(GongExtractString(rhs))
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = xlink.EnumShowType(GongExtractString(rhs))
					case "Actuate":
						inst.Actuate = xlink.EnumActuateType(GongExtractString(rhs))
					case "Element":
						inst.Element = GongExtractString(rhs)
					case "Position":
						inst.Position = GongExtractInt(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					}
				case *MeasureLayout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MeasureDistance":
						inst.MeasureDistance = GongExtractString(rhs)
					}
				case *MidiDevice:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Port":
						inst.Port = GongExtractInt(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *MidiInstrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "MidiChannel":
						inst.MidiChannel = GongExtractInt(rhs)
					case "MidiName":
						inst.MidiName = GongExtractString(rhs)
					case "MidiBank":
						inst.MidiBank = GongExtractInt(rhs)
					case "MidiProgram":
						inst.MidiProgram = GongExtractInt(rhs)
					case "MidiUnpitched":
						inst.MidiUnpitched = GongExtractInt(rhs)
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
					case "MiscellaneousField":
						__gong__assignSliceOfPointers(&inst.MiscellaneousField, rhs, identifierMap)
					}
				case *MiscellaneousField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NameAttr":
						inst.NameAttr = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *NameDisplay:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PrintObject":
						inst.PrintObject = EnumYesNo(GongExtractString(rhs))
					case "DisplayText":
						__gong__assignSliceOfPointers(&inst.DisplayText, rhs, identifierMap)
					case "AccidentalText":
						__gong__assignSliceOfPointers(&inst.AccidentalText, rhs, identifierMap)
					}
				case *NoteSize:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumNoteSizeType(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *OtherAppearance:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *OtherPlay:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *OtherText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *PageLayout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PageHeight":
						inst.PageHeight = GongExtractString(rhs)
					case "PageWidth":
						inst.PageWidth = GongExtractString(rhs)
					case "PageMargins":
						__gong__assignSliceOfPointers(&inst.PageMargins, rhs, identifierMap)
					}
				case *PageMargins:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumMarginType(GongExtractString(rhs))
					case "LeftMargin":
						inst.LeftMargin = GongExtractString(rhs)
					case "RightMargin":
						inst.RightMargin = GongExtractString(rhs)
					case "TopMargin":
						inst.TopMargin = GongExtractString(rhs)
					case "BottomMargin":
						inst.BottomMargin = GongExtractString(rhs)
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
						inst.Mute = EnumMute(GongExtractString(rhs))
					case "SemiPitched":
						inst.SemiPitched = EnumSemiPitched(GongExtractString(rhs))
					case "OtherPlay":
						__gong__assignSliceOfPointers(&inst.OtherPlay, rhs, identifierMap)
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
				case *Segno:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Halign":
						inst.Halign = EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = EnumValign(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *StaffLayout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "StaffDistance":
						inst.StaffDistance = GongExtractString(rhs)
					}
				case *StringType:
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
						inst.FontStyle = EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = EnumFontWeight(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Placement":
						inst.Placement = EnumAboveBelow(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractInt(rhs)
					}
				case *Supports:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumYesNo(GongExtractString(rhs))
					case "Element":
						inst.Element = GongExtractString(rhs)
					case "Attribute":
						inst.Attribute = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *SystemDividers:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "LeftDivider":
						__gong__assignPointer(&inst.LeftDivider, rhs, identifierMap)
					case "RightDivider":
						__gong__assignPointer(&inst.RightDivider, rhs, identifierMap)
					}
				case *SystemLayout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SystemMargins":
						__gong__assignPointer(&inst.SystemMargins, rhs, identifierMap)
					case "SystemDistance":
						inst.SystemDistance = GongExtractString(rhs)
					case "TopSystemDistance":
						inst.TopSystemDistance = GongExtractString(rhs)
					case "SystemDividers":
						__gong__assignPointer(&inst.SystemDividers, rhs, identifierMap)
					}
				case *SystemMargins:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "LeftMargin":
						inst.LeftMargin = GongExtractString(rhs)
					case "RightMargin":
						inst.RightMargin = GongExtractString(rhs)
					}
				case *TypedText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *WavyLine:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumStartStopContinue(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "Placement":
						inst.Placement = EnumAboveBelow(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "StartNote":
						inst.StartNote = EnumStartNote(GongExtractString(rhs))
					case "TrillStep":
						inst.TrillStep = EnumTrillStep(GongExtractString(rhs))
					case "TwoNoteTurn":
						inst.TwoNoteTurn = EnumTwoNoteTurn(GongExtractString(rhs))
					case "Accelerate":
						inst.Accelerate = EnumYesNo(GongExtractString(rhs))
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "SecondBeat":
						inst.SecondBeat = GongExtractString(rhs)
					case "LastBeat":
						inst.LastBeat = GongExtractString(rhs)
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
