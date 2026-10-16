// generated code - do not edit
package score

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
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
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
		for _, credit := range __gong__sortStageSetInstances(stageSet.Stage.Credits, stageSet.Stage.Credit_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			creditIdent := "__score" + credit.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Credit{Name: %s}).Stage(stageSet.Stage)", creditIdent, __gong__toRawStringLiteral(credit.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", creditIdent, __gong__toRawStringLiteral(credit.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Page = %d", creditIdent, credit.Page))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", creditIdent, __gong__toRawStringLiteral(credit.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.CreditType = %s", creditIdent, __gong__toRawStringLiteral(credit.CreditType)))
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
			defaultsIdent := "__score" + defaults.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Defaults{Name: %s}).Stage(stageSet.Stage)", defaultsIdent, __gong__toRawStringLiteral(defaults.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", defaultsIdent, __gong__toRawStringLiteral(defaults.Name)))
			if defaults.MusicFont != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + defaults.MusicFont.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MusicFont = %s", defaultsIdent, targetIdent))
			}
			if defaults.WordFont != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + defaults.WordFont.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WordFont = %s", defaultsIdent, targetIdent))
			}
			for _, elem := range defaults.LyricFont {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LyricFont = append(%s.LyricFont, %s)", defaultsIdent, defaultsIdent, targetIdent))
			}
			for _, elem := range defaults.LyricLanguage {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LyricLanguage = append(%s.LyricLanguage, %s)", defaultsIdent, defaultsIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, emptyfont := range __gong__sortStageSetInstances(stageSet.Stage.EmptyFonts, stageSet.Stage.EmptyFont_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			emptyfontIdent := "__score" + emptyfont.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.EmptyFont{Name: %s}).Stage(stageSet.Stage)", emptyfontIdent, __gong__toRawStringLiteral(emptyfont.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", emptyfontIdent, __gong__toRawStringLiteral(emptyfont.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", emptyfontIdent, __gong__toRawStringLiteral(emptyfont.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", emptyfontIdent, __gong__toRawStringLiteral(string(emptyfont.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", emptyfontIdent, __gong__toRawStringLiteral(emptyfont.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", emptyfontIdent, __gong__toRawStringLiteral(string(emptyfont.FontWeight))))
		}
	}
	if stageSet.Stage != nil {
		for _, groupbarline := range __gong__sortStageSetInstances(stageSet.Stage.GroupBarlines, stageSet.Stage.GroupBarline_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groupbarlineIdent := "__score" + groupbarline.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.GroupBarline{Name: %s}).Stage(stageSet.Stage)", groupbarlineIdent, __gong__toRawStringLiteral(groupbarline.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupbarlineIdent, __gong__toRawStringLiteral(groupbarline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", groupbarlineIdent, __gong__toRawStringLiteral(groupbarline.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", groupbarlineIdent, __gong__toRawStringLiteral(string(groupbarline.Value))))
		}
	}
	if stageSet.Stage != nil {
		for _, groupname := range __gong__sortStageSetInstances(stageSet.Stage.GroupNames, stageSet.Stage.GroupName_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groupnameIdent := "__score" + groupname.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.GroupName{Name: %s}).Stage(stageSet.Stage)", groupnameIdent, __gong__toRawStringLiteral(groupname.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupnameIdent, __gong__toRawStringLiteral(groupname.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", groupnameIdent, __gong__toRawStringLiteral(groupname.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", groupnameIdent, __gong__toRawStringLiteral(groupname.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", groupnameIdent, __gong__toRawStringLiteral(groupname.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", groupnameIdent, __gong__toRawStringLiteral(groupname.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", groupnameIdent, __gong__toRawStringLiteral(groupname.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", groupnameIdent, __gong__toRawStringLiteral(string(groupname.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", groupnameIdent, __gong__toRawStringLiteral(groupname.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", groupnameIdent, __gong__toRawStringLiteral(string(groupname.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", groupnameIdent, __gong__toRawStringLiteral(groupname.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", groupnameIdent, __gong__toRawStringLiteral(string(groupname.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", groupnameIdent, __gong__toRawStringLiteral(groupname.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, groupsymbol := range __gong__sortStageSetInstances(stageSet.Stage.GroupSymbols, stageSet.Stage.GroupSymbol_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groupsymbolIdent := "__score" + groupsymbol.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.GroupSymbol{Name: %s}).Stage(stageSet.Stage)", groupsymbolIdent, __gong__toRawStringLiteral(groupsymbol.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupsymbolIdent, __gong__toRawStringLiteral(groupsymbol.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", groupsymbolIdent, __gong__toRawStringLiteral(groupsymbol.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", groupsymbolIdent, __gong__toRawStringLiteral(groupsymbol.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", groupsymbolIdent, __gong__toRawStringLiteral(groupsymbol.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", groupsymbolIdent, __gong__toRawStringLiteral(groupsymbol.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", groupsymbolIdent, __gong__toRawStringLiteral(groupsymbol.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", groupsymbolIdent, __gong__toRawStringLiteral(string(groupsymbol.Value))))
		}
	}
	if stageSet.Stage != nil {
		for _, instrumentlink := range __gong__sortStageSetInstances(stageSet.Stage.InstrumentLinks, stageSet.Stage.InstrumentLink_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			instrumentlinkIdent := "__score" + instrumentlink.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.InstrumentLink{Name: %s}).Stage(stageSet.Stage)", instrumentlinkIdent, __gong__toRawStringLiteral(instrumentlink.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrumentlinkIdent, __gong__toRawStringLiteral(instrumentlink.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", instrumentlinkIdent, __gong__toRawStringLiteral(instrumentlink.Id)))
			if instrumentlink.RefScoreInstrument != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + instrumentlink.RefScoreInstrument.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RefScoreInstrument = %s", instrumentlinkIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, lyricfont := range __gong__sortStageSetInstances(stageSet.Stage.LyricFonts, stageSet.Stage.LyricFont_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lyricfontIdent := "__score" + lyricfont.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.LyricFont{Name: %s}).Stage(stageSet.Stage)", lyricfontIdent, __gong__toRawStringLiteral(lyricfont.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lyricfontIdent, __gong__toRawStringLiteral(lyricfont.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", lyricfontIdent, __gong__toRawStringLiteral(lyricfont.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameAttr = %s", lyricfontIdent, __gong__toRawStringLiteral(lyricfont.NameAttr)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", lyricfontIdent, __gong__toRawStringLiteral(lyricfont.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", lyricfontIdent, __gong__toRawStringLiteral(string(lyricfont.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", lyricfontIdent, __gong__toRawStringLiteral(lyricfont.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", lyricfontIdent, __gong__toRawStringLiteral(string(lyricfont.FontWeight))))
		}
	}
	if stageSet.Stage != nil {
		for _, lyriclanguage := range __gong__sortStageSetInstances(stageSet.Stage.LyricLanguages, stageSet.Stage.LyricLanguage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lyriclanguageIdent := "__score" + lyriclanguage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.LyricLanguage{Name: %s}).Stage(stageSet.Stage)", lyriclanguageIdent, __gong__toRawStringLiteral(lyriclanguage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lyriclanguageIdent, __gong__toRawStringLiteral(lyriclanguage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", lyriclanguageIdent, __gong__toRawStringLiteral(lyriclanguage.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameAttr = %s", lyriclanguageIdent, __gong__toRawStringLiteral(lyriclanguage.NameAttr)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", lyriclanguageIdent, __gong__toRawStringLiteral(lyriclanguage.Lang)))
		}
	}
	if stageSet.Stage != nil {
		for _, measure := range __gong__sortStageSetInstances(stageSet.Stage.Measures, stageSet.Stage.Measure_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			measureIdent := "__score" + measure.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Measure{Name: %s}).Stage(stageSet.Stage)", measureIdent, __gong__toRawStringLiteral(measure.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measureIdent, __gong__toRawStringLiteral(measure.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", measureIdent, __gong__toRawStringLiteral(measure.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", measureIdent, __gong__toRawStringLiteral(measure.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Implicit = %s", measureIdent, __gong__toRawStringLiteral(string(measure.Implicit))))
			values.WriteString(fmt.Sprintf("\n\t%s.NonControlling = %s", measureIdent, __gong__toRawStringLiteral(string(measure.NonControlling))))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", measureIdent, __gong__toRawStringLiteral(measure.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", measureIdent, __gong__toRawStringLiteral(measure.Id)))
			for _, elem := range measure.MusicData {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MusicData = append(%s.MusicData, %s)", measureIdent, measureIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, measuremusicdatachoice := range __gong__sortStageSetInstances(stageSet.Stage.MeasureMusicDataChoices, stageSet.Stage.MeasureMusicDataChoice_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			measuremusicdatachoiceIdent := "__score" + measuremusicdatachoice.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.MeasureMusicDataChoice{Name: %s}).Stage(stageSet.Stage)", measuremusicdatachoiceIdent, __gong__toRawStringLiteral(measuremusicdatachoice.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measuremusicdatachoiceIdent, __gong__toRawStringLiteral(measuremusicdatachoice.Name)))
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
			opusIdent := "__score" + opus.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Opus{Name: %s}).Stage(stageSet.Stage)", opusIdent, __gong__toRawStringLiteral(opus.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", opusIdent, __gong__toRawStringLiteral(opus.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", opusIdent, __gong__toRawStringLiteral(opus.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", opusIdent, __gong__toRawStringLiteral(string(opus.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", opusIdent, __gong__toRawStringLiteral(opus.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", opusIdent, __gong__toRawStringLiteral(opus.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", opusIdent, __gong__toRawStringLiteral(string(opus.Show))))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", opusIdent, __gong__toRawStringLiteral(string(opus.Actuate))))
		}
	}
	if stageSet.Stage != nil {
		for _, part := range __gong__sortStageSetInstances(stageSet.Stage.Parts, stageSet.Stage.Part_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			partIdent := "__score" + part.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Part{Name: %s}).Stage(stageSet.Stage)", partIdent, __gong__toRawStringLiteral(part.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partIdent, __gong__toRawStringLiteral(part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", partIdent, __gong__toRawStringLiteral(part.Id)))
			if part.RefScorePart != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + part.RefScorePart.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RefScorePart = %s", partIdent, targetIdent))
			}
			for _, elem := range part.Measure {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure = append(%s.Measure, %s)", partIdent, partIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, partgroup := range __gong__sortStageSetInstances(stageSet.Stage.PartGroups, stageSet.Stage.PartGroup_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			partgroupIdent := "__score" + partgroup.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.PartGroup{Name: %s}).Stage(stageSet.Stage)", partgroupIdent, __gong__toRawStringLiteral(partgroup.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partgroupIdent, __gong__toRawStringLiteral(partgroup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", partgroupIdent, __gong__toRawStringLiteral(string(partgroup.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", partgroupIdent, __gong__toRawStringLiteral(partgroup.Number)))
			if partgroup.GroupName != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + partgroup.GroupName.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupName = %s", partgroupIdent, targetIdent))
			}
			if partgroup.GroupAbbreviation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + partgroup.GroupAbbreviation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupAbbreviation = %s", partgroupIdent, targetIdent))
			}
			if partgroup.GroupSymbol != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + partgroup.GroupSymbol.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupSymbol = %s", partgroupIdent, targetIdent))
			}
			if partgroup.GroupBarline != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + partgroup.GroupBarline.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupBarline = %s", partgroupIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, partlink := range __gong__sortStageSetInstances(stageSet.Stage.PartLinks, stageSet.Stage.PartLink_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			partlinkIdent := "__score" + partlink.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.PartLink{Name: %s}).Stage(stageSet.Stage)", partlinkIdent, __gong__toRawStringLiteral(partlink.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partlinkIdent, __gong__toRawStringLiteral(partlink.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", partlinkIdent, __gong__toRawStringLiteral(partlink.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", partlinkIdent, __gong__toRawStringLiteral(string(partlink.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", partlinkIdent, __gong__toRawStringLiteral(partlink.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", partlinkIdent, __gong__toRawStringLiteral(partlink.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", partlinkIdent, __gong__toRawStringLiteral(string(partlink.Show))))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", partlinkIdent, __gong__toRawStringLiteral(string(partlink.Actuate))))
			values.WriteString(fmt.Sprintf("\n\t%s.GroupLink = %s", partlinkIdent, __gong__toRawStringLiteral(partlink.GroupLink)))
			for _, elem := range partlink.InstrumentLink {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InstrumentLink = append(%s.InstrumentLink, %s)", partlinkIdent, partlinkIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, partlist := range __gong__sortStageSetInstances(stageSet.Stage.PartLists, stageSet.Stage.PartList_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			partlistIdent := "__score" + partlist.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.PartList{Name: %s}).Stage(stageSet.Stage)", partlistIdent, __gong__toRawStringLiteral(partlist.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partlistIdent, __gong__toRawStringLiteral(partlist.Name)))
			for _, elem := range partlist.PartGroup {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartGroup = append(%s.PartGroup, %s)", partlistIdent, partlistIdent, targetIdent))
			}
			for _, elem := range partlist.ScorePart {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ScorePart = append(%s.ScorePart, %s)", partlistIdent, partlistIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, partname := range __gong__sortStageSetInstances(stageSet.Stage.PartNames, stageSet.Stage.PartName_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			partnameIdent := "__score" + partname.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.PartName{Name: %s}).Stage(stageSet.Stage)", partnameIdent, __gong__toRawStringLiteral(partname.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partnameIdent, __gong__toRawStringLiteral(partname.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", partnameIdent, __gong__toRawStringLiteral(partname.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", partnameIdent, __gong__toRawStringLiteral(partname.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", partnameIdent, __gong__toRawStringLiteral(partname.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", partnameIdent, __gong__toRawStringLiteral(partname.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", partnameIdent, __gong__toRawStringLiteral(partname.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", partnameIdent, __gong__toRawStringLiteral(string(partname.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", partnameIdent, __gong__toRawStringLiteral(partname.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", partnameIdent, __gong__toRawStringLiteral(string(partname.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", partnameIdent, __gong__toRawStringLiteral(partname.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", partnameIdent, __gong__toRawStringLiteral(string(partname.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", partnameIdent, __gong__toRawStringLiteral(string(partname.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", partnameIdent, __gong__toRawStringLiteral(partname.Value)))
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
			playerIdent := "__score" + player.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Player{Name: %s}).Stage(stageSet.Stage)", playerIdent, __gong__toRawStringLiteral(player.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", playerIdent, __gong__toRawStringLiteral(player.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", playerIdent, __gong__toRawStringLiteral(player.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.PlayerName = %s", playerIdent, __gong__toRawStringLiteral(player.PlayerName)))
		}
	}
	if stageSet.Stage != nil {
		for _, scoreinstrument := range __gong__sortStageSetInstances(stageSet.Stage.ScoreInstruments, stageSet.Stage.ScoreInstrument_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			scoreinstrumentIdent := "__score" + scoreinstrument.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.ScoreInstrument{Name: %s}).Stage(stageSet.Stage)", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.InstrumentName = %s", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.InstrumentName)))
			values.WriteString(fmt.Sprintf("\n\t%s.InstrumentAbbreviation = %s", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.InstrumentAbbreviation)))
			values.WriteString(fmt.Sprintf("\n\t%s.InstrumentSound = %s", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.InstrumentSound)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ensemble = %s", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.Ensemble)))
		}
	}
	if stageSet.Stage != nil {
		for _, scorepart := range __gong__sortStageSetInstances(stageSet.Stage.ScoreParts, stageSet.Stage.ScorePart_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			scorepartIdent := "__score" + scorepart.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.ScorePart{Name: %s}).Stage(stageSet.Stage)", scorepartIdent, __gong__toRawStringLiteral(scorepart.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scorepartIdent, __gong__toRawStringLiteral(scorepart.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", scorepartIdent, __gong__toRawStringLiteral(scorepart.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group = %s", scorepartIdent, __gong__toRawStringLiteral(scorepart.Group)))
			for _, elem := range scorepart.PartLink {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartLink = append(%s.PartLink, %s)", scorepartIdent, scorepartIdent, targetIdent))
			}
			if scorepart.PartName != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + scorepart.PartName.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartName = %s", scorepartIdent, targetIdent))
			}
			if scorepart.PartAbbreviation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + scorepart.PartAbbreviation.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartAbbreviation = %s", scorepartIdent, targetIdent))
			}
			for _, elem := range scorepart.ScoreInstrument {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ScoreInstrument = append(%s.ScoreInstrument, %s)", scorepartIdent, scorepartIdent, targetIdent))
			}
			for _, elem := range scorepart.Player {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Player = append(%s.Player, %s)", scorepartIdent, scorepartIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, scorepartwise := range __gong__sortStageSetInstances(stageSet.Stage.ScorePartwises, stageSet.Stage.ScorePartwise_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			scorepartwiseIdent := "__score" + scorepartwise.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.ScorePartwise{Name: %s}).Stage(stageSet.Stage)", scorepartwiseIdent, __gong__toRawStringLiteral(scorepartwise.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scorepartwiseIdent, __gong__toRawStringLiteral(scorepartwise.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Version = %s", scorepartwiseIdent, __gong__toRawStringLiteral(scorepartwise.Version)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementNumber = %s", scorepartwiseIdent, __gong__toRawStringLiteral(scorepartwise.MovementNumber)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementTitle = %s", scorepartwiseIdent, __gong__toRawStringLiteral(scorepartwise.MovementTitle)))
			if scorepartwise.Work != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + scorepartwise.Work.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Work = %s", scorepartwiseIdent, targetIdent))
			}
			if scorepartwise.Defaults != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + scorepartwise.Defaults.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Defaults = %s", scorepartwiseIdent, targetIdent))
			}
			for _, elem := range scorepartwise.Credit {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit = append(%s.Credit, %s)", scorepartwiseIdent, scorepartwiseIdent, targetIdent))
			}
			if scorepartwise.PartList != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + scorepartwise.PartList.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartList = %s", scorepartwiseIdent, targetIdent))
			}
			for _, elem := range scorepartwise.Part {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = append(%s.Part, %s)", scorepartwiseIdent, scorepartwiseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, scoretimewise := range __gong__sortStageSetInstances(stageSet.Stage.ScoreTimewises, stageSet.Stage.ScoreTimewise_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			scoretimewiseIdent := "__score" + scoretimewise.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.ScoreTimewise{Name: %s}).Stage(stageSet.Stage)", scoretimewiseIdent, __gong__toRawStringLiteral(scoretimewise.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scoretimewiseIdent, __gong__toRawStringLiteral(scoretimewise.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Version = %s", scoretimewiseIdent, __gong__toRawStringLiteral(scoretimewise.Version)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementNumber = %s", scoretimewiseIdent, __gong__toRawStringLiteral(scoretimewise.MovementNumber)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementTitle = %s", scoretimewiseIdent, __gong__toRawStringLiteral(scoretimewise.MovementTitle)))
			if scoretimewise.Work != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + scoretimewise.Work.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Work = %s", scoretimewiseIdent, targetIdent))
			}
			if scoretimewise.Defaults != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + scoretimewise.Defaults.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Defaults = %s", scoretimewiseIdent, targetIdent))
			}
			for _, elem := range scoretimewise.Credit {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit = append(%s.Credit, %s)", scoretimewiseIdent, scoretimewiseIdent, targetIdent))
			}
			if scoretimewise.PartList != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + scoretimewise.PartList.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartList = %s", scoretimewiseIdent, targetIdent))
			}
			for _, elem := range scoretimewise.Measure {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure = append(%s.Measure, %s)", scoretimewiseIdent, scoretimewiseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, scoretimewisemeasure := range __gong__sortStageSetInstances(stageSet.Stage.ScoreTimewiseMeasures, stageSet.Stage.ScoreTimewiseMeasure_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			scoretimewisemeasureIdent := "__score" + scoretimewisemeasure.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.ScoreTimewiseMeasure{Name: %s}).Stage(stageSet.Stage)", scoretimewisemeasureIdent, __gong__toRawStringLiteral(scoretimewisemeasure.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(scoretimewisemeasure.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(scoretimewisemeasure.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(scoretimewisemeasure.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Implicit = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(string(scoretimewisemeasure.Implicit))))
			values.WriteString(fmt.Sprintf("\n\t%s.NonControlling = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(string(scoretimewisemeasure.NonControlling))))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(scoretimewisemeasure.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(scoretimewisemeasure.Id)))
			for _, elem := range scoretimewisemeasure.Part {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = append(%s.Part, %s)", scoretimewisemeasureIdent, scoretimewisemeasureIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, scoretimewisemeasurepart := range __gong__sortStageSetInstances(stageSet.Stage.ScoreTimewiseMeasureParts, stageSet.Stage.ScoreTimewiseMeasurePart_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			scoretimewisemeasurepartIdent := "__score" + scoretimewisemeasurepart.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.ScoreTimewiseMeasurePart{Name: %s}).Stage(stageSet.Stage)", scoretimewisemeasurepartIdent, __gong__toRawStringLiteral(scoretimewisemeasurepart.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scoretimewisemeasurepartIdent, __gong__toRawStringLiteral(scoretimewisemeasurepart.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", scoretimewisemeasurepartIdent, __gong__toRawStringLiteral(scoretimewisemeasurepart.Id)))
			if scoretimewisemeasurepart.RefScorePart != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + scoretimewisemeasurepart.RefScorePart.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RefScorePart = %s", scoretimewisemeasurepartIdent, targetIdent))
			}
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
			workIdent := "__score" + work.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Work{Name: %s}).Stage(stageSet.Stage)", workIdent, __gong__toRawStringLiteral(work.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", workIdent, __gong__toRawStringLiteral(work.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.WorkNumber = %s", workIdent, __gong__toRawStringLiteral(work.WorkNumber)))
			values.WriteString(fmt.Sprintf("\n\t%s.WorkTitle = %s", workIdent, __gong__toRawStringLiteral(work.WorkTitle)))
			if work.Opus != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + work.Opus.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Opus = %s", workIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/score/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/score"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *score.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *score.StageSet) {

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
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/score":
			aliasToCanonical[alias] = "score"
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
			case "score":
				switch typeName {
				case "Credit":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Credit), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Defaults":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Defaults), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "EmptyFont":
					identifierMap[ident.Name] = __gong__stageSetInit(new(EmptyFont), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "GroupBarline":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GroupBarline), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "GroupName":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GroupName), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "GroupSymbol":
					identifierMap[ident.Name] = __gong__stageSetInit(new(GroupSymbol), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "InstrumentLink":
					identifierMap[ident.Name] = __gong__stageSetInit(new(InstrumentLink), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "LyricFont":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LyricFont), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "LyricLanguage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LyricLanguage), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Measure":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Measure), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MeasureMusicDataChoice":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeasureMusicDataChoice), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Opus":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Opus), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Part":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Part), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PartGroup":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PartGroup), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PartLink":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PartLink), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PartList":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PartList), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PartName":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PartName), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Player":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Player), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ScoreInstrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ScoreInstrument), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ScorePart":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ScorePart), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ScorePartwise":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ScorePartwise), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ScoreTimewise":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ScoreTimewise), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ScoreTimewiseMeasure":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ScoreTimewiseMeasure), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ScoreTimewiseMeasurePart":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ScoreTimewiseMeasurePart), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *Credit:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Page":
						inst.Page = GongExtractInt(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "CreditType":
						inst.CreditType = GongExtractString(rhs)
					}
				case *Defaults:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MusicFont":
						__gong__assignPointer(&inst.MusicFont, rhs, identifierMap)
					case "WordFont":
						__gong__assignPointer(&inst.WordFont, rhs, identifierMap)
					case "LyricFont":
						__gong__assignSliceOfPointers(&inst.LyricFont, rhs, identifierMap)
					case "LyricLanguage":
						__gong__assignSliceOfPointers(&inst.LyricLanguage, rhs, identifierMap)
					}
				case *EmptyFont:
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
					}
				case *GroupBarline:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Value":
						inst.Value = EnumGroupBarlineValue(GongExtractString(rhs))
					}
				case *GroupName:
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
					case "Justify":
						inst.Justify = common.EnumLeftCenterRight(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *GroupSymbol:
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
						inst.Value = attributes.EnumGroupSymbolValue(GongExtractString(rhs))
					}
				case *InstrumentLink:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RefScoreInstrument":
						__gong__assignPointer(&inst.RefScoreInstrument, rhs, identifierMap)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *LyricFont:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "NameAttr":
						inst.NameAttr = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = common.EnumFontStyle(GongExtractString(rhs))
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = common.EnumFontWeight(GongExtractString(rhs))
					}
				case *LyricLanguage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "NameAttr":
						inst.NameAttr = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					}
				case *Measure:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Implicit":
						inst.Implicit = common.EnumYesNo(GongExtractString(rhs))
					case "NonControlling":
						inst.NonControlling = common.EnumYesNo(GongExtractString(rhs))
					case "Width":
						inst.Width = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "MusicData":
						__gong__assignSliceOfPointers(&inst.MusicData, rhs, identifierMap)
					}
				case *MeasureMusicDataChoice:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Opus:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
					}
				case *Part:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RefScorePart":
						__gong__assignPointer(&inst.RefScorePart, rhs, identifierMap)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Measure":
						__gong__assignSliceOfPointers(&inst.Measure, rhs, identifierMap)
					}
				case *PartGroup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "GroupName":
						__gong__assignPointer(&inst.GroupName, rhs, identifierMap)
					case "GroupAbbreviation":
						__gong__assignPointer(&inst.GroupAbbreviation, rhs, identifierMap)
					case "GroupSymbol":
						__gong__assignPointer(&inst.GroupSymbol, rhs, identifierMap)
					case "GroupBarline":
						__gong__assignPointer(&inst.GroupBarline, rhs, identifierMap)
					}
				case *PartLink:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
					case "InstrumentLink":
						__gong__assignSliceOfPointers(&inst.InstrumentLink, rhs, identifierMap)
					case "GroupLink":
						inst.GroupLink = GongExtractString(rhs)
					}
				case *PartList:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PartGroup":
						__gong__assignSliceOfPointers(&inst.PartGroup, rhs, identifierMap)
					case "ScorePart":
						__gong__assignSliceOfPointers(&inst.ScorePart, rhs, identifierMap)
					}
				case *PartName:
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
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "Justify":
						inst.Justify = common.EnumLeftCenterRight(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Player:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "PlayerName":
						inst.PlayerName = GongExtractString(rhs)
					}
				case *ScoreInstrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "InstrumentName":
						inst.InstrumentName = GongExtractString(rhs)
					case "InstrumentAbbreviation":
						inst.InstrumentAbbreviation = GongExtractString(rhs)
					case "InstrumentSound":
						inst.InstrumentSound = GongExtractString(rhs)
					case "Ensemble":
						inst.Ensemble = GongExtractString(rhs)
					}
				case *ScorePart:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "PartLink":
						__gong__assignSliceOfPointers(&inst.PartLink, rhs, identifierMap)
					case "PartName":
						__gong__assignPointer(&inst.PartName, rhs, identifierMap)
					case "PartAbbreviation":
						__gong__assignPointer(&inst.PartAbbreviation, rhs, identifierMap)
					case "Group":
						inst.Group = GongExtractString(rhs)
					case "ScoreInstrument":
						__gong__assignSliceOfPointers(&inst.ScoreInstrument, rhs, identifierMap)
					case "Player":
						__gong__assignSliceOfPointers(&inst.Player, rhs, identifierMap)
					}
				case *ScorePartwise:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Version":
						inst.Version = GongExtractString(rhs)
					case "Work":
						__gong__assignPointer(&inst.Work, rhs, identifierMap)
					case "MovementNumber":
						inst.MovementNumber = GongExtractString(rhs)
					case "MovementTitle":
						inst.MovementTitle = GongExtractString(rhs)
					case "Defaults":
						__gong__assignPointer(&inst.Defaults, rhs, identifierMap)
					case "Credit":
						__gong__assignSliceOfPointers(&inst.Credit, rhs, identifierMap)
					case "PartList":
						__gong__assignPointer(&inst.PartList, rhs, identifierMap)
					case "Part":
						__gong__assignSliceOfPointers(&inst.Part, rhs, identifierMap)
					}
				case *ScoreTimewise:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Version":
						inst.Version = GongExtractString(rhs)
					case "Work":
						__gong__assignPointer(&inst.Work, rhs, identifierMap)
					case "MovementNumber":
						inst.MovementNumber = GongExtractString(rhs)
					case "MovementTitle":
						inst.MovementTitle = GongExtractString(rhs)
					case "Defaults":
						__gong__assignPointer(&inst.Defaults, rhs, identifierMap)
					case "Credit":
						__gong__assignSliceOfPointers(&inst.Credit, rhs, identifierMap)
					case "PartList":
						__gong__assignPointer(&inst.PartList, rhs, identifierMap)
					case "Measure":
						__gong__assignSliceOfPointers(&inst.Measure, rhs, identifierMap)
					}
				case *ScoreTimewiseMeasure:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Implicit":
						inst.Implicit = common.EnumYesNo(GongExtractString(rhs))
					case "NonControlling":
						inst.NonControlling = common.EnumYesNo(GongExtractString(rhs))
					case "Width":
						inst.Width = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Part":
						__gong__assignSliceOfPointers(&inst.Part, rhs, identifierMap)
					}
				case *ScoreTimewiseMeasurePart:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RefScorePart":
						__gong__assignPointer(&inst.RefScorePart, rhs, identifierMap)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Work:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "WorkNumber":
						inst.WorkNumber = GongExtractString(rhs)
					case "WorkTitle":
						inst.WorkTitle = GongExtractString(rhs)
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
