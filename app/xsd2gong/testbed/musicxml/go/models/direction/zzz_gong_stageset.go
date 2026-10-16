// generated code - do not edit
package direction

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

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/barline"
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
		for _, accord := range __gong__sortStageSetInstances(stageSet.Stage.Accords, stageSet.Stage.Accord_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			accordIdent := "__direction" + accord.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Accord{Name: %s}).Stage(stageSet.Stage)", accordIdent, __gong__toRawStringLiteral(accord.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", accordIdent, __gong__toRawStringLiteral(accord.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.String = %d", accordIdent, accord.String))
			values.WriteString(fmt.Sprintf("\n\t%s.TuningStep = %s", accordIdent, __gong__toRawStringLiteral(string(accord.TuningStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.TuningAlter = %s", accordIdent, __gong__toRawStringLiteral(accord.TuningAlter)))
			values.WriteString(fmt.Sprintf("\n\t%s.TuningOctave = %d", accordIdent, accord.TuningOctave))
		}
	}
	if stageSet.Stage != nil {
		for _, accordionregistration := range __gong__sortStageSetInstances(stageSet.Stage.AccordionRegistrations, stageSet.Stage.AccordionRegistration_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			accordionregistrationIdent := "__direction" + accordionregistration.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.AccordionRegistration{Name: %s}).Stage(stageSet.Stage)", accordionregistrationIdent, __gong__toRawStringLiteral(accordionregistration.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", accordionregistrationIdent, __gong__toRawStringLiteral(accordionregistration.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", accordionregistrationIdent, __gong__toRawStringLiteral(accordionregistration.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", accordionregistrationIdent, __gong__toRawStringLiteral(accordionregistration.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", accordionregistrationIdent, __gong__toRawStringLiteral(accordionregistration.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", accordionregistrationIdent, __gong__toRawStringLiteral(accordionregistration.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", accordionregistrationIdent, __gong__toRawStringLiteral(accordionregistration.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", accordionregistrationIdent, __gong__toRawStringLiteral(string(accordionregistration.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", accordionregistrationIdent, __gong__toRawStringLiteral(accordionregistration.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", accordionregistrationIdent, __gong__toRawStringLiteral(string(accordionregistration.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", accordionregistrationIdent, __gong__toRawStringLiteral(accordionregistration.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", accordionregistrationIdent, __gong__toRawStringLiteral(string(accordionregistration.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", accordionregistrationIdent, __gong__toRawStringLiteral(string(accordionregistration.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", accordionregistrationIdent, __gong__toRawStringLiteral(accordionregistration.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.AccordionMiddle = %d", accordionregistrationIdent, accordionregistration.AccordionMiddle))
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
			barreIdent := "__direction" + barre.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Barre{Name: %s}).Stage(stageSet.Stage)", barreIdent, __gong__toRawStringLiteral(barre.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", barreIdent, __gong__toRawStringLiteral(barre.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", barreIdent, __gong__toRawStringLiteral(string(barre.Type))))
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
			bassIdent := "__direction" + bass.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Bass{Name: %s}).Stage(stageSet.Stage)", bassIdent, __gong__toRawStringLiteral(bass.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bassIdent, __gong__toRawStringLiteral(bass.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arrangement = %s", bassIdent, __gong__toRawStringLiteral(string(bass.Arrangement))))
			if bass.BassSeparator != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + bass.BassSeparator.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BassSeparator = %s", bassIdent, targetIdent))
			}
			if bass.BassStep != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + bass.BassStep.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BassStep = %s", bassIdent, targetIdent))
			}
			if bass.BassAlter != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + bass.BassAlter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BassAlter = %s", bassIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, bassstep := range __gong__sortStageSetInstances(stageSet.Stage.BassSteps, stageSet.Stage.BassStep_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bassstepIdent := "__direction" + bassstep.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.BassStep{Name: %s}).Stage(stageSet.Stage)", bassstepIdent, __gong__toRawStringLiteral(bassstep.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bassstepIdent, __gong__toRawStringLiteral(bassstep.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", bassstepIdent, __gong__toRawStringLiteral(bassstep.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", bassstepIdent, __gong__toRawStringLiteral(bassstep.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", bassstepIdent, __gong__toRawStringLiteral(bassstep.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", bassstepIdent, __gong__toRawStringLiteral(bassstep.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", bassstepIdent, __gong__toRawStringLiteral(bassstep.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", bassstepIdent, __gong__toRawStringLiteral(bassstep.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", bassstepIdent, __gong__toRawStringLiteral(string(bassstep.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", bassstepIdent, __gong__toRawStringLiteral(bassstep.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", bassstepIdent, __gong__toRawStringLiteral(string(bassstep.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", bassstepIdent, __gong__toRawStringLiteral(bassstep.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", bassstepIdent, __gong__toRawStringLiteral(string(bassstep.Value))))
		}
	}
	if stageSet.Stage != nil {
		for _, beatunittied := range __gong__sortStageSetInstances(stageSet.Stage.BeatUnitTieds, stageSet.Stage.BeatUnitTied_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			beatunittiedIdent := "__direction" + beatunittied.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.BeatUnitTied{Name: %s}).Stage(stageSet.Stage)", beatunittiedIdent, __gong__toRawStringLiteral(beatunittied.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", beatunittiedIdent, __gong__toRawStringLiteral(beatunittied.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.BeatUnit = %s", beatunittiedIdent, __gong__toRawStringLiteral(string(beatunittied.BeatUnit))))
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
			beaterIdent := "__direction" + beater.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Beater{Name: %s}).Stage(stageSet.Stage)", beaterIdent, __gong__toRawStringLiteral(beater.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", beaterIdent, __gong__toRawStringLiteral(beater.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tip = %s", beaterIdent, __gong__toRawStringLiteral(string(beater.Tip))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", beaterIdent, __gong__toRawStringLiteral(string(beater.Value))))
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
			bracketIdent := "__direction" + bracket.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Bracket{Name: %s}).Stage(stageSet.Stage)", bracketIdent, __gong__toRawStringLiteral(bracket.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", bracketIdent, __gong__toRawStringLiteral(string(bracket.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", bracketIdent, bracket.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.LineEnd = %s", bracketIdent, __gong__toRawStringLiteral(string(bracket.LineEnd))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndLength = %s", bracketIdent, __gong__toRawStringLiteral(bracket.EndLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.LineType = %s", bracketIdent, __gong__toRawStringLiteral(string(bracket.LineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.DashLength = %s", bracketIdent, __gong__toRawStringLiteral(bracket.DashLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.SpaceLength = %s", bracketIdent, __gong__toRawStringLiteral(bracket.SpaceLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", bracketIdent, __gong__toRawStringLiteral(bracket.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", bracketIdent, __gong__toRawStringLiteral(bracket.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", bracketIdent, __gong__toRawStringLiteral(bracket.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", bracketIdent, __gong__toRawStringLiteral(bracket.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", bracketIdent, __gong__toRawStringLiteral(bracket.Id)))
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
			dashesIdent := "__direction" + dashes.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Dashes{Name: %s}).Stage(stageSet.Stage)", dashesIdent, __gong__toRawStringLiteral(dashes.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", dashesIdent, __gong__toRawStringLiteral(dashes.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", dashesIdent, __gong__toRawStringLiteral(string(dashes.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", dashesIdent, dashes.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.DashLength = %s", dashesIdent, __gong__toRawStringLiteral(dashes.DashLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.SpaceLength = %s", dashesIdent, __gong__toRawStringLiteral(dashes.SpaceLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", dashesIdent, __gong__toRawStringLiteral(dashes.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", dashesIdent, __gong__toRawStringLiteral(dashes.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", dashesIdent, __gong__toRawStringLiteral(dashes.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", dashesIdent, __gong__toRawStringLiteral(dashes.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", dashesIdent, __gong__toRawStringLiteral(dashes.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", dashesIdent, __gong__toRawStringLiteral(dashes.Id)))
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
			degreeIdent := "__direction" + degree.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Degree{Name: %s}).Stage(stageSet.Stage)", degreeIdent, __gong__toRawStringLiteral(degree.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", degreeIdent, __gong__toRawStringLiteral(degree.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", degreeIdent, __gong__toRawStringLiteral(string(degree.PrintObject))))
			if degree.DegreeValue != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + degree.DegreeValue.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DegreeValue = %s", degreeIdent, targetIdent))
			}
			if degree.DegreeAlter != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + degree.DegreeAlter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DegreeAlter = %s", degreeIdent, targetIdent))
			}
			if degree.DegreeType != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + degree.DegreeType.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DegreeType = %s", degreeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, degreealter := range __gong__sortStageSetInstances(stageSet.Stage.DegreeAlters, stageSet.Stage.DegreeAlter_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			degreealterIdent := "__direction" + degreealter.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.DegreeAlter{Name: %s}).Stage(stageSet.Stage)", degreealterIdent, __gong__toRawStringLiteral(degreealter.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", degreealterIdent, __gong__toRawStringLiteral(degreealter.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PlusMinus = %s", degreealterIdent, __gong__toRawStringLiteral(string(degreealter.PlusMinus))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", degreealterIdent, __gong__toRawStringLiteral(degreealter.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", degreealterIdent, __gong__toRawStringLiteral(degreealter.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", degreealterIdent, __gong__toRawStringLiteral(degreealter.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", degreealterIdent, __gong__toRawStringLiteral(degreealter.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", degreealterIdent, __gong__toRawStringLiteral(degreealter.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", degreealterIdent, __gong__toRawStringLiteral(string(degreealter.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", degreealterIdent, __gong__toRawStringLiteral(degreealter.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", degreealterIdent, __gong__toRawStringLiteral(string(degreealter.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", degreealterIdent, __gong__toRawStringLiteral(degreealter.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", degreealterIdent, __gong__toRawStringLiteral(degreealter.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, degreetype := range __gong__sortStageSetInstances(stageSet.Stage.DegreeTypes, stageSet.Stage.DegreeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			degreetypeIdent := "__direction" + degreetype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.DegreeType{Name: %s}).Stage(stageSet.Stage)", degreetypeIdent, __gong__toRawStringLiteral(degreetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", degreetypeIdent, __gong__toRawStringLiteral(degreetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", degreetypeIdent, __gong__toRawStringLiteral(degreetype.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", degreetypeIdent, __gong__toRawStringLiteral(degreetype.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", degreetypeIdent, __gong__toRawStringLiteral(degreetype.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", degreetypeIdent, __gong__toRawStringLiteral(degreetype.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", degreetypeIdent, __gong__toRawStringLiteral(degreetype.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", degreetypeIdent, __gong__toRawStringLiteral(degreetype.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", degreetypeIdent, __gong__toRawStringLiteral(string(degreetype.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", degreetypeIdent, __gong__toRawStringLiteral(degreetype.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", degreetypeIdent, __gong__toRawStringLiteral(string(degreetype.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", degreetypeIdent, __gong__toRawStringLiteral(degreetype.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", degreetypeIdent, __gong__toRawStringLiteral(string(degreetype.Value))))
		}
	}
	if stageSet.Stage != nil {
		for _, degreevalue := range __gong__sortStageSetInstances(stageSet.Stage.DegreeValues, stageSet.Stage.DegreeValue_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			degreevalueIdent := "__direction" + degreevalue.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.DegreeValue{Name: %s}).Stage(stageSet.Stage)", degreevalueIdent, __gong__toRawStringLiteral(degreevalue.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", degreevalueIdent, __gong__toRawStringLiteral(degreevalue.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Symbol = %s", degreevalueIdent, __gong__toRawStringLiteral(string(degreevalue.Symbol))))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", degreevalueIdent, __gong__toRawStringLiteral(degreevalue.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", degreevalueIdent, __gong__toRawStringLiteral(degreevalue.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", degreevalueIdent, __gong__toRawStringLiteral(degreevalue.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", degreevalueIdent, __gong__toRawStringLiteral(degreevalue.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", degreevalueIdent, __gong__toRawStringLiteral(degreevalue.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", degreevalueIdent, __gong__toRawStringLiteral(degreevalue.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", degreevalueIdent, __gong__toRawStringLiteral(string(degreevalue.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", degreevalueIdent, __gong__toRawStringLiteral(degreevalue.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", degreevalueIdent, __gong__toRawStringLiteral(string(degreevalue.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", degreevalueIdent, __gong__toRawStringLiteral(degreevalue.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", degreevalueIdent, degreevalue.Value))
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
			directionIdent := "__direction" + direction.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Direction{Name: %s}).Stage(stageSet.Stage)", directionIdent, __gong__toRawStringLiteral(direction.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionIdent, __gong__toRawStringLiteral(direction.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", directionIdent, __gong__toRawStringLiteral(string(direction.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Directive = %s", directionIdent, __gong__toRawStringLiteral(string(direction.Directive))))
			values.WriteString(fmt.Sprintf("\n\t%s.System = %s", directionIdent, __gong__toRawStringLiteral(string(direction.System))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directionIdent, __gong__toRawStringLiteral(direction.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Voice = %s", directionIdent, __gong__toRawStringLiteral(direction.Voice)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff = %d", directionIdent, direction.Staff))
			for _, elem := range direction.DirectionType {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectionType = append(%s.DirectionType, %s)", directionIdent, directionIdent, targetIdent))
			}
			if direction.Offset != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + direction.Offset.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", directionIdent, targetIdent))
			}
			if direction.Sound != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + direction.Sound.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sound = %s", directionIdent, targetIdent))
			}
			if direction.Listening != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + direction.Listening.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listening = %s", directionIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, directiontype := range __gong__sortStageSetInstances(stageSet.Stage.DirectionTypes, stageSet.Stage.DirectionType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			directiontypeIdent := "__direction" + directiontype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.DirectionType{Name: %s}).Stage(stageSet.Stage)", directiontypeIdent, __gong__toRawStringLiteral(directiontype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directiontypeIdent, __gong__toRawStringLiteral(directiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directiontypeIdent, __gong__toRawStringLiteral(directiontype.Id)))
			if directiontype.Wedge != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.Wedge.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wedge = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.Dashes != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.Dashes.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dashes = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.Bracket != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.Bracket.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bracket = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.Pedal != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.Pedal.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pedal = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.Metronome != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.Metronome.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metronome = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.OctaveShift != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.OctaveShift.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OctaveShift = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.HarpPedals != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.HarpPedals.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HarpPedals = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.StringMute != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.StringMute.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StringMute = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.Scordatura != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.Scordatura.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Scordatura = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.Image != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.Image.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Image = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.PrincipalVoice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.PrincipalVoice.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PrincipalVoice = %s", directiontypeIdent, targetIdent))
			}
			for _, elem := range directiontype.Percussion {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Percussion = append(%s.Percussion, %s)", directiontypeIdent, directiontypeIdent, targetIdent))
			}
			if directiontype.AccordionRegistration != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.AccordionRegistration.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AccordionRegistration = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.StaffDivide != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.StaffDivide.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StaffDivide = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.OtherDirection != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + directiontype.OtherDirection.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherDirection = %s", directiontypeIdent, targetIdent))
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
			effectIdent := "__direction" + effect.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Effect{Name: %s}).Stage(stageSet.Stage)", effectIdent, __gong__toRawStringLiteral(effect.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", effectIdent, __gong__toRawStringLiteral(effect.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", effectIdent, __gong__toRawStringLiteral(effect.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", effectIdent, __gong__toRawStringLiteral(string(effect.Value))))
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
			featureIdent := "__direction" + feature.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Feature{Name: %s}).Stage(stageSet.Stage)", featureIdent, __gong__toRawStringLiteral(feature.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", featureIdent, __gong__toRawStringLiteral(feature.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", featureIdent, __gong__toRawStringLiteral(feature.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", featureIdent, __gong__toRawStringLiteral(feature.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, firstfret := range __gong__sortStageSetInstances(stageSet.Stage.FirstFrets, stageSet.Stage.FirstFret_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			firstfretIdent := "__direction" + firstfret.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.FirstFret{Name: %s}).Stage(stageSet.Stage)", firstfretIdent, __gong__toRawStringLiteral(firstfret.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", firstfretIdent, __gong__toRawStringLiteral(firstfret.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", firstfretIdent, __gong__toRawStringLiteral(firstfret.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", firstfretIdent, __gong__toRawStringLiteral(string(firstfret.Location))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", firstfretIdent, firstfret.Value))
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
			frameIdent := "__direction" + frame.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Frame{Name: %s}).Stage(stageSet.Stage)", frameIdent, __gong__toRawStringLiteral(frame.Name)))
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
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", frameIdent, __gong__toRawStringLiteral(frame.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", frameIdent, __gong__toRawStringLiteral(frame.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", frameIdent, __gong__toRawStringLiteral(frame.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", frameIdent, __gong__toRawStringLiteral(frame.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", frameIdent, __gong__toRawStringLiteral(frame.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", frameIdent, __gong__toRawStringLiteral(string(frame.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", frameIdent, __gong__toRawStringLiteral(string(frame.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", frameIdent, __gong__toRawStringLiteral(frame.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.FrameStrings = %d", frameIdent, frame.FrameStrings))
			values.WriteString(fmt.Sprintf("\n\t%s.FrameFrets = %d", frameIdent, frame.FrameFrets))
			if frame.FirstFret != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + frame.FirstFret.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FirstFret = %s", frameIdent, targetIdent))
			}
			for _, elem := range frame.FrameNote {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FrameNote = append(%s.FrameNote, %s)", frameIdent, frameIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, framenote := range __gong__sortStageSetInstances(stageSet.Stage.FrameNotes, stageSet.Stage.FrameNote_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			framenoteIdent := "__direction" + framenote.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.FrameNote{Name: %s}).Stage(stageSet.Stage)", framenoteIdent, __gong__toRawStringLiteral(framenote.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", framenoteIdent, __gong__toRawStringLiteral(framenote.Name)))
			if framenote.Barre != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + framenote.Barre.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Barre = %s", framenoteIdent, targetIdent))
			}
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
			glassIdent := "__direction" + glass.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Glass{Name: %s}).Stage(stageSet.Stage)", glassIdent, __gong__toRawStringLiteral(glass.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glassIdent, __gong__toRawStringLiteral(glass.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", glassIdent, __gong__toRawStringLiteral(glass.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", glassIdent, __gong__toRawStringLiteral(string(glass.Value))))
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
			groupingIdent := "__direction" + grouping.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Grouping{Name: %s}).Stage(stageSet.Stage)", groupingIdent, __gong__toRawStringLiteral(grouping.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", groupingIdent, __gong__toRawStringLiteral(string(grouping.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.MemberOf = %s", groupingIdent, __gong__toRawStringLiteral(grouping.MemberOf)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Id)))
			for _, elem := range grouping.Feature {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Feature = append(%s.Feature, %s)", groupingIdent, groupingIdent, targetIdent))
			}
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
			harmonyIdent := "__direction" + harmony.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Harmony{Name: %s}).Stage(stageSet.Stage)", harmonyIdent, __gong__toRawStringLiteral(harmony.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", harmonyIdent, __gong__toRawStringLiteral(string(harmony.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintFrame = %s", harmonyIdent, __gong__toRawStringLiteral(string(harmony.PrintFrame))))
			values.WriteString(fmt.Sprintf("\n\t%s.Arrangement = %s", harmonyIdent, __gong__toRawStringLiteral(string(harmony.Arrangement))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", harmonyIdent, __gong__toRawStringLiteral(string(harmony.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", harmonyIdent, __gong__toRawStringLiteral(string(harmony.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", harmonyIdent, __gong__toRawStringLiteral(string(harmony.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", harmonyIdent, __gong__toRawStringLiteral(string(harmony.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.System = %s", harmonyIdent, __gong__toRawStringLiteral(string(harmony.System))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", harmonyIdent, __gong__toRawStringLiteral(harmony.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff = %d", harmonyIdent, harmony.Staff))
			for _, elem := range harmony.Root {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Root = append(%s.Root, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Numeral {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Numeral = append(%s.Numeral, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Function {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Function = append(%s.Function, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Kind {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Kind = append(%s.Kind, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Inversion {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inversion = append(%s.Inversion, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Bass {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bass = append(%s.Bass, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Degree {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Degree = append(%s.Degree, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			if harmony.Frame != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + harmony.Frame.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frame = %s", harmonyIdent, targetIdent))
			}
			if harmony.Offset != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + harmony.Offset.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", harmonyIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, harmonyalter := range __gong__sortStageSetInstances(stageSet.Stage.HarmonyAlters, stageSet.Stage.HarmonyAlter_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			harmonyalterIdent := "__direction" + harmonyalter.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.HarmonyAlter{Name: %s}).Stage(stageSet.Stage)", harmonyalterIdent, __gong__toRawStringLiteral(harmonyalter.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harmonyalterIdent, __gong__toRawStringLiteral(harmonyalter.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", harmonyalterIdent, __gong__toRawStringLiteral(string(harmonyalter.Location))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", harmonyalterIdent, __gong__toRawStringLiteral(string(harmonyalter.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", harmonyalterIdent, __gong__toRawStringLiteral(harmonyalter.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", harmonyalterIdent, __gong__toRawStringLiteral(harmonyalter.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", harmonyalterIdent, __gong__toRawStringLiteral(harmonyalter.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", harmonyalterIdent, __gong__toRawStringLiteral(harmonyalter.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", harmonyalterIdent, __gong__toRawStringLiteral(harmonyalter.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", harmonyalterIdent, __gong__toRawStringLiteral(string(harmonyalter.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", harmonyalterIdent, __gong__toRawStringLiteral(harmonyalter.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", harmonyalterIdent, __gong__toRawStringLiteral(string(harmonyalter.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", harmonyalterIdent, __gong__toRawStringLiteral(harmonyalter.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", harmonyalterIdent, __gong__toRawStringLiteral(harmonyalter.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, harppedals := range __gong__sortStageSetInstances(stageSet.Stage.HarpPedalss, stageSet.Stage.HarpPedals_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			harppedalsIdent := "__direction" + harppedals.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.HarpPedals{Name: %s}).Stage(stageSet.Stage)", harppedalsIdent, __gong__toRawStringLiteral(harppedals.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harppedalsIdent, __gong__toRawStringLiteral(harppedals.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", harppedalsIdent, __gong__toRawStringLiteral(harppedals.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", harppedalsIdent, __gong__toRawStringLiteral(harppedals.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", harppedalsIdent, __gong__toRawStringLiteral(harppedals.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", harppedalsIdent, __gong__toRawStringLiteral(harppedals.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", harppedalsIdent, __gong__toRawStringLiteral(harppedals.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", harppedalsIdent, __gong__toRawStringLiteral(string(harppedals.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", harppedalsIdent, __gong__toRawStringLiteral(harppedals.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", harppedalsIdent, __gong__toRawStringLiteral(string(harppedals.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", harppedalsIdent, __gong__toRawStringLiteral(harppedals.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", harppedalsIdent, __gong__toRawStringLiteral(string(harppedals.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", harppedalsIdent, __gong__toRawStringLiteral(string(harppedals.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", harppedalsIdent, __gong__toRawStringLiteral(harppedals.Id)))
			for _, elem := range harppedals.PedalTuning {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PedalTuning = append(%s.PedalTuning, %s)", harppedalsIdent, harppedalsIdent, targetIdent))
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
			imageIdent := "__direction" + image.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Image{Name: %s}).Stage(stageSet.Stage)", imageIdent, __gong__toRawStringLiteral(image.Name)))
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
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", imageIdent, __gong__toRawStringLiteral(image.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", imageIdent, __gong__toRawStringLiteral(image.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", imageIdent, __gong__toRawStringLiteral(image.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", imageIdent, __gong__toRawStringLiteral(image.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", imageIdent, __gong__toRawStringLiteral(string(image.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", imageIdent, __gong__toRawStringLiteral(string(image.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", imageIdent, __gong__toRawStringLiteral(image.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, instrumentchange := range __gong__sortStageSetInstances(stageSet.Stage.InstrumentChanges, stageSet.Stage.InstrumentChange_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			instrumentchangeIdent := "__direction" + instrumentchange.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.InstrumentChange{Name: %s}).Stage(stageSet.Stage)", instrumentchangeIdent, __gong__toRawStringLiteral(instrumentchange.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrumentchangeIdent, __gong__toRawStringLiteral(instrumentchange.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", instrumentchangeIdent, __gong__toRawStringLiteral(instrumentchange.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.InstrumentSound = %s", instrumentchangeIdent, __gong__toRawStringLiteral(instrumentchange.InstrumentSound)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ensemble = %s", instrumentchangeIdent, __gong__toRawStringLiteral(instrumentchange.Ensemble)))
			if instrumentchange.VirtualInstrument != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + instrumentchange.VirtualInstrument.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VirtualInstrument = %s", instrumentchangeIdent, targetIdent))
			}
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
			inversionIdent := "__direction" + inversion.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Inversion{Name: %s}).Stage(stageSet.Stage)", inversionIdent, __gong__toRawStringLiteral(inversion.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", inversionIdent, __gong__toRawStringLiteral(inversion.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", inversionIdent, __gong__toRawStringLiteral(inversion.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", inversionIdent, __gong__toRawStringLiteral(inversion.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", inversionIdent, __gong__toRawStringLiteral(inversion.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", inversionIdent, __gong__toRawStringLiteral(inversion.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", inversionIdent, __gong__toRawStringLiteral(string(inversion.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", inversionIdent, __gong__toRawStringLiteral(inversion.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", inversionIdent, __gong__toRawStringLiteral(string(inversion.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", inversionIdent, __gong__toRawStringLiteral(inversion.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", inversionIdent, inversion.Value))
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
			kindIdent := "__direction" + kind.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Kind{Name: %s}).Stage(stageSet.Stage)", kindIdent, __gong__toRawStringLiteral(kind.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", kindIdent, __gong__toRawStringLiteral(kind.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.UseSymbols = %s", kindIdent, __gong__toRawStringLiteral(string(kind.UseSymbols))))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", kindIdent, __gong__toRawStringLiteral(kind.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackDegrees = %s", kindIdent, __gong__toRawStringLiteral(string(kind.StackDegrees))))
			values.WriteString(fmt.Sprintf("\n\t%s.ParenthesesDegrees = %s", kindIdent, __gong__toRawStringLiteral(string(kind.ParenthesesDegrees))))
			values.WriteString(fmt.Sprintf("\n\t%s.BracketDegrees = %s", kindIdent, __gong__toRawStringLiteral(string(kind.BracketDegrees))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", kindIdent, __gong__toRawStringLiteral(kind.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", kindIdent, __gong__toRawStringLiteral(kind.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", kindIdent, __gong__toRawStringLiteral(kind.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", kindIdent, __gong__toRawStringLiteral(kind.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", kindIdent, __gong__toRawStringLiteral(kind.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", kindIdent, __gong__toRawStringLiteral(string(kind.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", kindIdent, __gong__toRawStringLiteral(kind.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", kindIdent, __gong__toRawStringLiteral(string(kind.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", kindIdent, __gong__toRawStringLiteral(kind.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", kindIdent, __gong__toRawStringLiteral(string(kind.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", kindIdent, __gong__toRawStringLiteral(string(kind.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", kindIdent, __gong__toRawStringLiteral(string(kind.Value))))
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
			listeningIdent := "__direction" + listening.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Listening{Name: %s}).Stage(stageSet.Stage)", listeningIdent, __gong__toRawStringLiteral(listening.Name)))
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
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sync = append(%s.Sync, %s)", listeningIdent, listeningIdent, targetIdent))
			}
			for _, elem := range listening.OtherListening {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherListening = append(%s.OtherListening, %s)", listeningIdent, listeningIdent, targetIdent))
			}
			if listening.Offset != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + listening.Offset.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", listeningIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, measurenumbering := range __gong__sortStageSetInstances(stageSet.Stage.MeasureNumberings, stageSet.Stage.MeasureNumbering_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			measurenumberingIdent := "__direction" + measurenumbering.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.MeasureNumbering{Name: %s}).Stage(stageSet.Stage)", measurenumberingIdent, __gong__toRawStringLiteral(measurenumbering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measurenumberingIdent, __gong__toRawStringLiteral(measurenumbering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.System = %s", measurenumberingIdent, __gong__toRawStringLiteral(string(measurenumbering.System))))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff = %d", measurenumberingIdent, measurenumbering.Staff))
			values.WriteString(fmt.Sprintf("\n\t%s.MultipleRestAlways = %s", measurenumberingIdent, __gong__toRawStringLiteral(string(measurenumbering.MultipleRestAlways))))
			values.WriteString(fmt.Sprintf("\n\t%s.MultipleRestRange = %s", measurenumberingIdent, __gong__toRawStringLiteral(string(measurenumbering.MultipleRestRange))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", measurenumberingIdent, __gong__toRawStringLiteral(measurenumbering.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", measurenumberingIdent, __gong__toRawStringLiteral(measurenumbering.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", measurenumberingIdent, __gong__toRawStringLiteral(measurenumbering.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", measurenumberingIdent, __gong__toRawStringLiteral(measurenumbering.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", measurenumberingIdent, __gong__toRawStringLiteral(measurenumbering.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", measurenumberingIdent, __gong__toRawStringLiteral(string(measurenumbering.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", measurenumberingIdent, __gong__toRawStringLiteral(measurenumbering.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", measurenumberingIdent, __gong__toRawStringLiteral(string(measurenumbering.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", measurenumberingIdent, __gong__toRawStringLiteral(measurenumbering.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", measurenumberingIdent, __gong__toRawStringLiteral(string(measurenumbering.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", measurenumberingIdent, __gong__toRawStringLiteral(string(measurenumbering.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", measurenumberingIdent, __gong__toRawStringLiteral(string(measurenumbering.Value))))
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
			membraneIdent := "__direction" + membrane.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Membrane{Name: %s}).Stage(stageSet.Stage)", membraneIdent, __gong__toRawStringLiteral(membrane.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", membraneIdent, __gong__toRawStringLiteral(membrane.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", membraneIdent, __gong__toRawStringLiteral(membrane.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", membraneIdent, __gong__toRawStringLiteral(string(membrane.Value))))
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
			metalIdent := "__direction" + metal.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Metal{Name: %s}).Stage(stageSet.Stage)", metalIdent, __gong__toRawStringLiteral(metal.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metalIdent, __gong__toRawStringLiteral(metal.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", metalIdent, __gong__toRawStringLiteral(metal.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", metalIdent, __gong__toRawStringLiteral(string(metal.Value))))
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
			metronomeIdent := "__direction" + metronome.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Metronome{Name: %s}).Stage(stageSet.Stage)", metronomeIdent, __gong__toRawStringLiteral(metronome.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", metronomeIdent, __gong__toRawStringLiteral(string(metronome.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", metronomeIdent, __gong__toRawStringLiteral(string(metronome.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", metronomeIdent, __gong__toRawStringLiteral(string(metronome.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", metronomeIdent, __gong__toRawStringLiteral(string(metronome.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", metronomeIdent, __gong__toRawStringLiteral(string(metronome.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", metronomeIdent, __gong__toRawStringLiteral(string(metronome.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.Justify = %s", metronomeIdent, __gong__toRawStringLiteral(string(metronome.Justify))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.MetronomeRelation = %s", metronomeIdent, __gong__toRawStringLiteral(metronome.MetronomeRelation)))
			for _, elem := range metronome.BeatUnitTied {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BeatUnitTied = append(%s.BeatUnitTied, %s)", metronomeIdent, metronomeIdent, targetIdent))
			}
			if metronome.PerMinute != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + metronome.PerMinute.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PerMinute = %s", metronomeIdent, targetIdent))
			}
			for _, elem := range metronome.MetronomeNote {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MetronomeNote = append(%s.MetronomeNote, %s)", metronomeIdent, metronomeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, metronomebeam := range __gong__sortStageSetInstances(stageSet.Stage.MetronomeBeams, stageSet.Stage.MetronomeBeam_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			metronomebeamIdent := "__direction" + metronomebeam.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.MetronomeBeam{Name: %s}).Stage(stageSet.Stage)", metronomebeamIdent, __gong__toRawStringLiteral(metronomebeam.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronomebeamIdent, __gong__toRawStringLiteral(metronomebeam.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", metronomebeamIdent, metronomebeam.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", metronomebeamIdent, __gong__toRawStringLiteral(string(metronomebeam.Value))))
		}
	}
	if stageSet.Stage != nil {
		for _, metronomenote := range __gong__sortStageSetInstances(stageSet.Stage.MetronomeNotes, stageSet.Stage.MetronomeNote_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			metronomenoteIdent := "__direction" + metronomenote.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.MetronomeNote{Name: %s}).Stage(stageSet.Stage)", metronomenoteIdent, __gong__toRawStringLiteral(metronomenote.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronomenoteIdent, __gong__toRawStringLiteral(metronomenote.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MetronomeType = %s", metronomenoteIdent, __gong__toRawStringLiteral(string(metronomenote.MetronomeType))))
			for _, elem := range metronomenote.MetronomeBeam {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MetronomeBeam = append(%s.MetronomeBeam, %s)", metronomenoteIdent, metronomenoteIdent, targetIdent))
			}
			if metronomenote.MetronomeTied != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + metronomenote.MetronomeTied.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MetronomeTied = %s", metronomenoteIdent, targetIdent))
			}
			if metronomenote.MetronomeTuplet != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + metronomenote.MetronomeTuplet.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MetronomeTuplet = %s", metronomenoteIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, metronometied := range __gong__sortStageSetInstances(stageSet.Stage.MetronomeTieds, stageSet.Stage.MetronomeTied_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			metronometiedIdent := "__direction" + metronometied.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.MetronomeTied{Name: %s}).Stage(stageSet.Stage)", metronometiedIdent, __gong__toRawStringLiteral(metronometied.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronometiedIdent, __gong__toRawStringLiteral(metronometied.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", metronometiedIdent, __gong__toRawStringLiteral(string(metronometied.Type))))
		}
	}
	if stageSet.Stage != nil {
		for _, metronometuplet := range __gong__sortStageSetInstances(stageSet.Stage.MetronomeTuplets, stageSet.Stage.MetronomeTuplet_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			metronometupletIdent := "__direction" + metronometuplet.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.MetronomeTuplet{Name: %s}).Stage(stageSet.Stage)", metronometupletIdent, __gong__toRawStringLiteral(metronometuplet.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronometupletIdent, __gong__toRawStringLiteral(metronometuplet.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", metronometupletIdent, __gong__toRawStringLiteral(string(metronometuplet.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Bracket = %s", metronometupletIdent, __gong__toRawStringLiteral(string(metronometuplet.Bracket))))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowNumber = %s", metronometupletIdent, __gong__toRawStringLiteral(string(metronometuplet.ShowNumber))))
			values.WriteString(fmt.Sprintf("\n\t%s.ActualNotes = %d", metronometupletIdent, metronometuplet.ActualNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.NormalNotes = %d", metronometupletIdent, metronometuplet.NormalNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.NormalType = %s", metronometupletIdent, __gong__toRawStringLiteral(string(metronometuplet.NormalType))))
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
			numeralIdent := "__direction" + numeral.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Numeral{Name: %s}).Stage(stageSet.Stage)", numeralIdent, __gong__toRawStringLiteral(numeral.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", numeralIdent, __gong__toRawStringLiteral(numeral.Name)))
			if numeral.NumeralRoot != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + numeral.NumeralRoot.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NumeralRoot = %s", numeralIdent, targetIdent))
			}
			if numeral.NumeralAlter != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + numeral.NumeralAlter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NumeralAlter = %s", numeralIdent, targetIdent))
			}
			if numeral.NumeralKey != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + numeral.NumeralKey.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NumeralKey = %s", numeralIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, numeralkey := range __gong__sortStageSetInstances(stageSet.Stage.NumeralKeys, stageSet.Stage.NumeralKey_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			numeralkeyIdent := "__direction" + numeralkey.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.NumeralKey{Name: %s}).Stage(stageSet.Stage)", numeralkeyIdent, __gong__toRawStringLiteral(numeralkey.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", numeralkeyIdent, __gong__toRawStringLiteral(numeralkey.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", numeralkeyIdent, __gong__toRawStringLiteral(string(numeralkey.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.NumeralFifths = %d", numeralkeyIdent, numeralkey.NumeralFifths))
			values.WriteString(fmt.Sprintf("\n\t%s.NumeralMode = %s", numeralkeyIdent, __gong__toRawStringLiteral(string(numeralkey.NumeralMode))))
		}
	}
	if stageSet.Stage != nil {
		for _, numeralroot := range __gong__sortStageSetInstances(stageSet.Stage.NumeralRoots, stageSet.Stage.NumeralRoot_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			numeralrootIdent := "__direction" + numeralroot.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.NumeralRoot{Name: %s}).Stage(stageSet.Stage)", numeralrootIdent, __gong__toRawStringLiteral(numeralroot.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", numeralrootIdent, __gong__toRawStringLiteral(numeralroot.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", numeralrootIdent, __gong__toRawStringLiteral(numeralroot.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", numeralrootIdent, __gong__toRawStringLiteral(numeralroot.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", numeralrootIdent, __gong__toRawStringLiteral(numeralroot.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", numeralrootIdent, __gong__toRawStringLiteral(numeralroot.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", numeralrootIdent, __gong__toRawStringLiteral(numeralroot.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", numeralrootIdent, __gong__toRawStringLiteral(numeralroot.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", numeralrootIdent, __gong__toRawStringLiteral(string(numeralroot.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", numeralrootIdent, __gong__toRawStringLiteral(numeralroot.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", numeralrootIdent, __gong__toRawStringLiteral(string(numeralroot.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", numeralrootIdent, __gong__toRawStringLiteral(numeralroot.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", numeralrootIdent, numeralroot.Value))
		}
	}
	if stageSet.Stage != nil {
		for _, octaveshift := range __gong__sortStageSetInstances(stageSet.Stage.OctaveShifts, stageSet.Stage.OctaveShift_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			octaveshiftIdent := "__direction" + octaveshift.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.OctaveShift{Name: %s}).Stage(stageSet.Stage)", octaveshiftIdent, __gong__toRawStringLiteral(octaveshift.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", octaveshiftIdent, __gong__toRawStringLiteral(octaveshift.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", octaveshiftIdent, __gong__toRawStringLiteral(string(octaveshift.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", octaveshiftIdent, octaveshift.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %d", octaveshiftIdent, octaveshift.Size))
			values.WriteString(fmt.Sprintf("\n\t%s.DashLength = %s", octaveshiftIdent, __gong__toRawStringLiteral(octaveshift.DashLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.SpaceLength = %s", octaveshiftIdent, __gong__toRawStringLiteral(octaveshift.SpaceLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", octaveshiftIdent, __gong__toRawStringLiteral(octaveshift.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", octaveshiftIdent, __gong__toRawStringLiteral(octaveshift.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", octaveshiftIdent, __gong__toRawStringLiteral(octaveshift.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", octaveshiftIdent, __gong__toRawStringLiteral(octaveshift.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", octaveshiftIdent, __gong__toRawStringLiteral(octaveshift.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", octaveshiftIdent, __gong__toRawStringLiteral(string(octaveshift.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", octaveshiftIdent, __gong__toRawStringLiteral(octaveshift.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", octaveshiftIdent, __gong__toRawStringLiteral(string(octaveshift.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", octaveshiftIdent, __gong__toRawStringLiteral(octaveshift.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", octaveshiftIdent, __gong__toRawStringLiteral(octaveshift.Id)))
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
			offsetIdent := "__direction" + offset.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Offset{Name: %s}).Stage(stageSet.Stage)", offsetIdent, __gong__toRawStringLiteral(offset.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", offsetIdent, __gong__toRawStringLiteral(offset.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sound = %s", offsetIdent, __gong__toRawStringLiteral(string(offset.Sound))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", offsetIdent, __gong__toRawStringLiteral(offset.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, otherdirection := range __gong__sortStageSetInstances(stageSet.Stage.OtherDirections, stageSet.Stage.OtherDirection_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			otherdirectionIdent := "__direction" + otherdirection.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.OtherDirection{Name: %s}).Stage(stageSet.Stage)", otherdirectionIdent, __gong__toRawStringLiteral(otherdirection.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", otherdirectionIdent, __gong__toRawStringLiteral(otherdirection.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", otherdirectionIdent, __gong__toRawStringLiteral(string(otherdirection.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", otherdirectionIdent, __gong__toRawStringLiteral(otherdirection.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", otherdirectionIdent, __gong__toRawStringLiteral(otherdirection.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", otherdirectionIdent, __gong__toRawStringLiteral(otherdirection.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", otherdirectionIdent, __gong__toRawStringLiteral(otherdirection.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", otherdirectionIdent, __gong__toRawStringLiteral(otherdirection.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", otherdirectionIdent, __gong__toRawStringLiteral(string(otherdirection.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", otherdirectionIdent, __gong__toRawStringLiteral(otherdirection.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", otherdirectionIdent, __gong__toRawStringLiteral(string(otherdirection.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", otherdirectionIdent, __gong__toRawStringLiteral(otherdirection.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", otherdirectionIdent, __gong__toRawStringLiteral(string(otherdirection.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", otherdirectionIdent, __gong__toRawStringLiteral(string(otherdirection.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", otherdirectionIdent, __gong__toRawStringLiteral(otherdirection.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", otherdirectionIdent, __gong__toRawStringLiteral(otherdirection.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", otherdirectionIdent, __gong__toRawStringLiteral(otherdirection.Value)))
		}
	}
	if stageSet.Stage != nil {
		for _, otherlistening := range __gong__sortStageSetInstances(stageSet.Stage.OtherListenings, stageSet.Stage.OtherListening_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			otherlisteningIdent := "__direction" + otherlistening.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.OtherListening{Name: %s}).Stage(stageSet.Stage)", otherlisteningIdent, __gong__toRawStringLiteral(otherlistening.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", otherlisteningIdent, __gong__toRawStringLiteral(otherlistening.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", otherlisteningIdent, __gong__toRawStringLiteral(otherlistening.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", otherlisteningIdent, __gong__toRawStringLiteral(otherlistening.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeOnly = %s", otherlisteningIdent, __gong__toRawStringLiteral(otherlistening.TimeOnly)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", otherlisteningIdent, __gong__toRawStringLiteral(otherlistening.Value)))
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
			pedalIdent := "__direction" + pedal.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Pedal{Name: %s}).Stage(stageSet.Stage)", pedalIdent, __gong__toRawStringLiteral(pedal.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", pedalIdent, __gong__toRawStringLiteral(string(pedal.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", pedalIdent, pedal.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %s", pedalIdent, __gong__toRawStringLiteral(string(pedal.Line))))
			values.WriteString(fmt.Sprintf("\n\t%s.Sign = %s", pedalIdent, __gong__toRawStringLiteral(string(pedal.Sign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Abbreviated = %s", pedalIdent, __gong__toRawStringLiteral(string(pedal.Abbreviated))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", pedalIdent, __gong__toRawStringLiteral(pedal.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", pedalIdent, __gong__toRawStringLiteral(pedal.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", pedalIdent, __gong__toRawStringLiteral(pedal.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", pedalIdent, __gong__toRawStringLiteral(pedal.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", pedalIdent, __gong__toRawStringLiteral(pedal.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", pedalIdent, __gong__toRawStringLiteral(string(pedal.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", pedalIdent, __gong__toRawStringLiteral(pedal.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", pedalIdent, __gong__toRawStringLiteral(string(pedal.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", pedalIdent, __gong__toRawStringLiteral(string(pedal.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", pedalIdent, __gong__toRawStringLiteral(string(pedal.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", pedalIdent, __gong__toRawStringLiteral(pedal.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, pedaltuning := range __gong__sortStageSetInstances(stageSet.Stage.PedalTunings, stageSet.Stage.PedalTuning_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pedaltuningIdent := "__direction" + pedaltuning.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.PedalTuning{Name: %s}).Stage(stageSet.Stage)", pedaltuningIdent, __gong__toRawStringLiteral(pedaltuning.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pedaltuningIdent, __gong__toRawStringLiteral(pedaltuning.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PedalStep = %s", pedaltuningIdent, __gong__toRawStringLiteral(string(pedaltuning.PedalStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.PedalAlter = %s", pedaltuningIdent, __gong__toRawStringLiteral(pedaltuning.PedalAlter)))
		}
	}
	if stageSet.Stage != nil {
		for _, perminute := range __gong__sortStageSetInstances(stageSet.Stage.PerMinutes, stageSet.Stage.PerMinute_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			perminuteIdent := "__direction" + perminute.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.PerMinute{Name: %s}).Stage(stageSet.Stage)", perminuteIdent, __gong__toRawStringLiteral(perminute.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", perminuteIdent, __gong__toRawStringLiteral(perminute.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", perminuteIdent, __gong__toRawStringLiteral(perminute.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", perminuteIdent, __gong__toRawStringLiteral(string(perminute.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", perminuteIdent, __gong__toRawStringLiteral(perminute.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", perminuteIdent, __gong__toRawStringLiteral(string(perminute.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", perminuteIdent, __gong__toRawStringLiteral(perminute.Value)))
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
			percussionIdent := "__direction" + percussion.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Percussion{Name: %s}).Stage(stageSet.Stage)", percussionIdent, __gong__toRawStringLiteral(percussion.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", percussionIdent, __gong__toRawStringLiteral(percussion.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", percussionIdent, __gong__toRawStringLiteral(percussion.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", percussionIdent, __gong__toRawStringLiteral(percussion.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", percussionIdent, __gong__toRawStringLiteral(percussion.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", percussionIdent, __gong__toRawStringLiteral(percussion.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", percussionIdent, __gong__toRawStringLiteral(string(percussion.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", percussionIdent, __gong__toRawStringLiteral(percussion.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", percussionIdent, __gong__toRawStringLiteral(string(percussion.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", percussionIdent, __gong__toRawStringLiteral(string(percussion.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", percussionIdent, __gong__toRawStringLiteral(string(percussion.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Enclosure = %s", percussionIdent, __gong__toRawStringLiteral(string(percussion.Enclosure))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", percussionIdent, __gong__toRawStringLiteral(percussion.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.StickLocation = %s", percussionIdent, __gong__toRawStringLiteral(string(percussion.StickLocation))))
			if percussion.Glass != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + percussion.Glass.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Glass = %s", percussionIdent, targetIdent))
			}
			if percussion.Metal != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + percussion.Metal.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metal = %s", percussionIdent, targetIdent))
			}
			if percussion.Wood != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + percussion.Wood.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wood = %s", percussionIdent, targetIdent))
			}
			if percussion.Pitched != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + percussion.Pitched.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pitched = %s", percussionIdent, targetIdent))
			}
			if percussion.Membrane != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + percussion.Membrane.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Membrane = %s", percussionIdent, targetIdent))
			}
			if percussion.Effect != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + percussion.Effect.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Effect = %s", percussionIdent, targetIdent))
			}
			if percussion.Timpani != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + percussion.Timpani.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Timpani = %s", percussionIdent, targetIdent))
			}
			if percussion.Beater != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + percussion.Beater.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Beater = %s", percussionIdent, targetIdent))
			}
			if percussion.Stick != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + percussion.Stick.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stick = %s", percussionIdent, targetIdent))
			}
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
			pitchedIdent := "__direction" + pitched.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Pitched{Name: %s}).Stage(stageSet.Stage)", pitchedIdent, __gong__toRawStringLiteral(pitched.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pitchedIdent, __gong__toRawStringLiteral(pitched.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", pitchedIdent, __gong__toRawStringLiteral(pitched.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", pitchedIdent, __gong__toRawStringLiteral(string(pitched.Value))))
		}
	}
	if stageSet.Stage != nil {
		for _, principalvoice := range __gong__sortStageSetInstances(stageSet.Stage.PrincipalVoices, stageSet.Stage.PrincipalVoice_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			principalvoiceIdent := "__direction" + principalvoice.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.PrincipalVoice{Name: %s}).Stage(stageSet.Stage)", principalvoiceIdent, __gong__toRawStringLiteral(principalvoice.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", principalvoiceIdent, __gong__toRawStringLiteral(principalvoice.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", principalvoiceIdent, __gong__toRawStringLiteral(string(principalvoice.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Symbol = %s", principalvoiceIdent, __gong__toRawStringLiteral(string(principalvoice.Symbol))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", principalvoiceIdent, __gong__toRawStringLiteral(principalvoice.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", principalvoiceIdent, __gong__toRawStringLiteral(principalvoice.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", principalvoiceIdent, __gong__toRawStringLiteral(principalvoice.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", principalvoiceIdent, __gong__toRawStringLiteral(principalvoice.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", principalvoiceIdent, __gong__toRawStringLiteral(principalvoice.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", principalvoiceIdent, __gong__toRawStringLiteral(string(principalvoice.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", principalvoiceIdent, __gong__toRawStringLiteral(principalvoice.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", principalvoiceIdent, __gong__toRawStringLiteral(string(principalvoice.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", principalvoiceIdent, __gong__toRawStringLiteral(principalvoice.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", principalvoiceIdent, __gong__toRawStringLiteral(string(principalvoice.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", principalvoiceIdent, __gong__toRawStringLiteral(string(principalvoice.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", principalvoiceIdent, __gong__toRawStringLiteral(principalvoice.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", principalvoiceIdent, __gong__toRawStringLiteral(principalvoice.Value)))
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
			printIdent := "__direction" + print.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Print{Name: %s}).Stage(stageSet.Stage)", printIdent, __gong__toRawStringLiteral(print.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", printIdent, __gong__toRawStringLiteral(print.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StaffSpacing = %s", printIdent, __gong__toRawStringLiteral(print.StaffSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.NewSystem = %s", printIdent, __gong__toRawStringLiteral(string(print.NewSystem))))
			values.WriteString(fmt.Sprintf("\n\t%s.NewPage = %s", printIdent, __gong__toRawStringLiteral(string(print.NewPage))))
			values.WriteString(fmt.Sprintf("\n\t%s.BlankPage = %d", printIdent, print.BlankPage))
			values.WriteString(fmt.Sprintf("\n\t%s.PageNumber = %s", printIdent, __gong__toRawStringLiteral(print.PageNumber)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", printIdent, __gong__toRawStringLiteral(print.Id)))
			if print.MeasureNumbering != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + print.MeasureNumbering.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MeasureNumbering = %s", printIdent, targetIdent))
			}
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
			rootIdent := "__direction" + root.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Root{Name: %s}).Stage(stageSet.Stage)", rootIdent, __gong__toRawStringLiteral(root.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rootIdent, __gong__toRawStringLiteral(root.Name)))
			if root.RootStep != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + root.RootStep.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootStep = %s", rootIdent, targetIdent))
			}
			if root.RootAlter != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + root.RootAlter.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootAlter = %s", rootIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, rootstep := range __gong__sortStageSetInstances(stageSet.Stage.RootSteps, stageSet.Stage.RootStep_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rootstepIdent := "__direction" + rootstep.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.RootStep{Name: %s}).Stage(stageSet.Stage)", rootstepIdent, __gong__toRawStringLiteral(rootstep.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rootstepIdent, __gong__toRawStringLiteral(rootstep.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", rootstepIdent, __gong__toRawStringLiteral(rootstep.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", rootstepIdent, __gong__toRawStringLiteral(rootstep.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", rootstepIdent, __gong__toRawStringLiteral(rootstep.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", rootstepIdent, __gong__toRawStringLiteral(rootstep.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", rootstepIdent, __gong__toRawStringLiteral(rootstep.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", rootstepIdent, __gong__toRawStringLiteral(rootstep.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", rootstepIdent, __gong__toRawStringLiteral(string(rootstep.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", rootstepIdent, __gong__toRawStringLiteral(rootstep.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", rootstepIdent, __gong__toRawStringLiteral(string(rootstep.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", rootstepIdent, __gong__toRawStringLiteral(rootstep.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", rootstepIdent, __gong__toRawStringLiteral(string(rootstep.Value))))
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
			scordaturaIdent := "__direction" + scordatura.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Scordatura{Name: %s}).Stage(stageSet.Stage)", scordaturaIdent, __gong__toRawStringLiteral(scordatura.Name)))
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
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accord = append(%s.Accord, %s)", scordaturaIdent, scordaturaIdent, targetIdent))
			}
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
			soundIdent := "__direction" + sound.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Sound{Name: %s}).Stage(stageSet.Stage)", soundIdent, __gong__toRawStringLiteral(sound.Name)))
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
			values.WriteString(fmt.Sprintf("\n\t%s.ForwardRepeat = %s", soundIdent, __gong__toRawStringLiteral(string(sound.ForwardRepeat))))
			values.WriteString(fmt.Sprintf("\n\t%s.Fine = %s", soundIdent, __gong__toRawStringLiteral(sound.Fine)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeOnly = %s", soundIdent, __gong__toRawStringLiteral(sound.TimeOnly)))
			values.WriteString(fmt.Sprintf("\n\t%s.Pizzicato = %s", soundIdent, __gong__toRawStringLiteral(string(sound.Pizzicato))))
			values.WriteString(fmt.Sprintf("\n\t%s.Pan = %s", soundIdent, __gong__toRawStringLiteral(sound.Pan)))
			values.WriteString(fmt.Sprintf("\n\t%s.Elevation = %s", soundIdent, __gong__toRawStringLiteral(sound.Elevation)))
			values.WriteString(fmt.Sprintf("\n\t%s.DamperPedal = %s", soundIdent, __gong__toRawStringLiteral(sound.DamperPedal)))
			values.WriteString(fmt.Sprintf("\n\t%s.SoftPedal = %s", soundIdent, __gong__toRawStringLiteral(sound.SoftPedal)))
			values.WriteString(fmt.Sprintf("\n\t%s.SostenutoPedal = %s", soundIdent, __gong__toRawStringLiteral(sound.SostenutoPedal)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", soundIdent, __gong__toRawStringLiteral(sound.Id)))
			for _, elem := range sound.InstrumentChange {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InstrumentChange = append(%s.InstrumentChange, %s)", soundIdent, soundIdent, targetIdent))
			}
			if sound.Swing != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + sound.Swing.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Swing = %s", soundIdent, targetIdent))
			}
			if sound.Offset != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__direction" + sound.Offset.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", soundIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, staffdivide := range __gong__sortStageSetInstances(stageSet.Stage.StaffDivides, stageSet.Stage.StaffDivide_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staffdivideIdent := "__direction" + staffdivide.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.StaffDivide{Name: %s}).Stage(stageSet.Stage)", staffdivideIdent, __gong__toRawStringLiteral(staffdivide.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staffdivideIdent, __gong__toRawStringLiteral(staffdivide.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", staffdivideIdent, __gong__toRawStringLiteral(string(staffdivide.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", staffdivideIdent, __gong__toRawStringLiteral(staffdivide.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", staffdivideIdent, __gong__toRawStringLiteral(staffdivide.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", staffdivideIdent, __gong__toRawStringLiteral(staffdivide.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", staffdivideIdent, __gong__toRawStringLiteral(staffdivide.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", staffdivideIdent, __gong__toRawStringLiteral(staffdivide.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", staffdivideIdent, __gong__toRawStringLiteral(string(staffdivide.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", staffdivideIdent, __gong__toRawStringLiteral(staffdivide.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", staffdivideIdent, __gong__toRawStringLiteral(string(staffdivide.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", staffdivideIdent, __gong__toRawStringLiteral(staffdivide.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", staffdivideIdent, __gong__toRawStringLiteral(string(staffdivide.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", staffdivideIdent, __gong__toRawStringLiteral(string(staffdivide.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", staffdivideIdent, __gong__toRawStringLiteral(staffdivide.Id)))
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
			stickIdent := "__direction" + stick.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Stick{Name: %s}).Stage(stageSet.Stage)", stickIdent, __gong__toRawStringLiteral(stick.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stickIdent, __gong__toRawStringLiteral(stick.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tip = %s", stickIdent, __gong__toRawStringLiteral(string(stick.Tip))))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", stickIdent, __gong__toRawStringLiteral(string(stick.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.DashedCircle = %s", stickIdent, __gong__toRawStringLiteral(string(stick.DashedCircle))))
			values.WriteString(fmt.Sprintf("\n\t%s.StickType = %s", stickIdent, __gong__toRawStringLiteral(string(stick.StickType))))
			values.WriteString(fmt.Sprintf("\n\t%s.StickMaterial = %s", stickIdent, __gong__toRawStringLiteral(string(stick.StickMaterial))))
		}
	}
	if stageSet.Stage != nil {
		for _, stringmute := range __gong__sortStageSetInstances(stageSet.Stage.StringMutes, stageSet.Stage.StringMute_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stringmuteIdent := "__direction" + stringmute.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.StringMute{Name: %s}).Stage(stageSet.Stage)", stringmuteIdent, __gong__toRawStringLiteral(stringmute.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stringmuteIdent, __gong__toRawStringLiteral(stringmute.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", stringmuteIdent, __gong__toRawStringLiteral(string(stringmute.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", stringmuteIdent, __gong__toRawStringLiteral(stringmute.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", stringmuteIdent, __gong__toRawStringLiteral(stringmute.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", stringmuteIdent, __gong__toRawStringLiteral(stringmute.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", stringmuteIdent, __gong__toRawStringLiteral(stringmute.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", stringmuteIdent, __gong__toRawStringLiteral(stringmute.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", stringmuteIdent, __gong__toRawStringLiteral(string(stringmute.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", stringmuteIdent, __gong__toRawStringLiteral(stringmute.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", stringmuteIdent, __gong__toRawStringLiteral(string(stringmute.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", stringmuteIdent, __gong__toRawStringLiteral(stringmute.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Halign = %s", stringmuteIdent, __gong__toRawStringLiteral(string(stringmute.Halign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Valign = %s", stringmuteIdent, __gong__toRawStringLiteral(string(stringmute.Valign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", stringmuteIdent, __gong__toRawStringLiteral(stringmute.Id)))
		}
	}
	if stageSet.Stage != nil {
		for _, styletext := range __gong__sortStageSetInstances(stageSet.Stage.StyleTexts, stageSet.Stage.StyleText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			styletextIdent := "__direction" + styletext.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.StyleText{Name: %s}).Stage(stageSet.Stage)", styletextIdent, __gong__toRawStringLiteral(styletext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", styletextIdent, __gong__toRawStringLiteral(styletext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", styletextIdent, __gong__toRawStringLiteral(styletext.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", styletextIdent, __gong__toRawStringLiteral(styletext.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", styletextIdent, __gong__toRawStringLiteral(styletext.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", styletextIdent, __gong__toRawStringLiteral(styletext.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", styletextIdent, __gong__toRawStringLiteral(styletext.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", styletextIdent, __gong__toRawStringLiteral(string(styletext.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", styletextIdent, __gong__toRawStringLiteral(styletext.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", styletextIdent, __gong__toRawStringLiteral(string(styletext.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", styletextIdent, __gong__toRawStringLiteral(styletext.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", styletextIdent, __gong__toRawStringLiteral(styletext.Value)))
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
			swingIdent := "__direction" + swing.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Swing{Name: %s}).Stage(stageSet.Stage)", swingIdent, __gong__toRawStringLiteral(swing.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", swingIdent, __gong__toRawStringLiteral(swing.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.First = %d", swingIdent, swing.First))
			values.WriteString(fmt.Sprintf("\n\t%s.Second = %d", swingIdent, swing.Second))
			values.WriteString(fmt.Sprintf("\n\t%s.SwingType = %s", swingIdent, __gong__toRawStringLiteral(string(swing.SwingType))))
			values.WriteString(fmt.Sprintf("\n\t%s.SwingStyle = %s", swingIdent, __gong__toRawStringLiteral(swing.SwingStyle)))
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
			syncIdent := "__direction" + sync.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Sync{Name: %s}).Stage(stageSet.Stage)", syncIdent, __gong__toRawStringLiteral(sync.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", syncIdent, __gong__toRawStringLiteral(sync.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", syncIdent, __gong__toRawStringLiteral(string(sync.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Latency = %d", syncIdent, sync.Latency))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", syncIdent, __gong__toRawStringLiteral(sync.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeOnly = %s", syncIdent, __gong__toRawStringLiteral(sync.TimeOnly)))
		}
	}
	if stageSet.Stage != nil {
		for _, timemodification := range __gong__sortStageSetInstances(stageSet.Stage.TimeModifications, stageSet.Stage.TimeModification_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			timemodificationIdent := "__direction" + timemodification.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.TimeModification{Name: %s}).Stage(stageSet.Stage)", timemodificationIdent, __gong__toRawStringLiteral(timemodification.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timemodificationIdent, __gong__toRawStringLiteral(timemodification.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ActualNotes = %d", timemodificationIdent, timemodification.ActualNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.NormalNotes = %d", timemodificationIdent, timemodification.NormalNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.NormalType = %s", timemodificationIdent, __gong__toRawStringLiteral(string(timemodification.NormalType))))
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
			timpaniIdent := "__direction" + timpani.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Timpani{Name: %s}).Stage(stageSet.Stage)", timpaniIdent, __gong__toRawStringLiteral(timpani.Name)))
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
		for _, virtualinstrument := range __gong__sortStageSetInstances(stageSet.Stage.VirtualInstruments, stageSet.Stage.VirtualInstrument_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			virtualinstrumentIdent := "__direction" + virtualinstrument.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.VirtualInstrument{Name: %s}).Stage(stageSet.Stage)", virtualinstrumentIdent, __gong__toRawStringLiteral(virtualinstrument.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", virtualinstrumentIdent, __gong__toRawStringLiteral(virtualinstrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.VirtualLibrary = %s", virtualinstrumentIdent, __gong__toRawStringLiteral(virtualinstrument.VirtualLibrary)))
			values.WriteString(fmt.Sprintf("\n\t%s.VirtualName = %s", virtualinstrumentIdent, __gong__toRawStringLiteral(virtualinstrument.VirtualName)))
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
			wedgeIdent := "__direction" + wedge.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Wedge{Name: %s}).Stage(stageSet.Stage)", wedgeIdent, __gong__toRawStringLiteral(wedge.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", wedgeIdent, __gong__toRawStringLiteral(string(wedge.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", wedgeIdent, wedge.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Spread = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.Spread)))
			values.WriteString(fmt.Sprintf("\n\t%s.Niente = %s", wedgeIdent, __gong__toRawStringLiteral(string(wedge.Niente))))
			values.WriteString(fmt.Sprintf("\n\t%s.LineType = %s", wedgeIdent, __gong__toRawStringLiteral(string(wedge.LineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.DashLength = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.DashLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.SpaceLength = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.SpaceLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", wedgeIdent, __gong__toRawStringLiteral(wedge.RelativeY)))
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
			woodIdent := "__direction" + wood.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Wood{Name: %s}).Stage(stageSet.Stage)", woodIdent, __gong__toRawStringLiteral(wood.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", woodIdent, __gong__toRawStringLiteral(wood.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", woodIdent, __gong__toRawStringLiteral(wood.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", woodIdent, __gong__toRawStringLiteral(string(wood.Value))))
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
go run github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *direction.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *direction.StageSet) {

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
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction":
			aliasToCanonical[alias] = "direction"
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
			case "direction":
				switch typeName {
				case "Accord":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Accord), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "AccordionRegistration":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AccordionRegistration), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Barre":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Barre), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Bass":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bass), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "BassStep":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BassStep), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "BeatUnitTied":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BeatUnitTied), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Beater":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Beater), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Bracket":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Bracket), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Dashes":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Dashes), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Degree":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Degree), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DegreeAlter":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DegreeAlter), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DegreeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DegreeType), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DegreeValue":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DegreeValue), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Direction":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Direction), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "DirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DirectionType), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Effect":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Effect), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Feature":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Feature), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "FirstFret":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FirstFret), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Frame":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Frame), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "FrameNote":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FrameNote), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Glass":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Glass), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Grouping":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Grouping), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Harmony":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Harmony), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "HarmonyAlter":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HarmonyAlter), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "HarpPedals":
					identifierMap[ident.Name] = __gong__stageSetInit(new(HarpPedals), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Image":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Image), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "InstrumentChange":
					identifierMap[ident.Name] = __gong__stageSetInit(new(InstrumentChange), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Inversion":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Inversion), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Kind":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Kind), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Listening":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Listening), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MeasureNumbering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MeasureNumbering), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Membrane":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Membrane), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Metal":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metal), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Metronome":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Metronome), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MetronomeBeam":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MetronomeBeam), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MetronomeNote":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MetronomeNote), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MetronomeTied":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MetronomeTied), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "MetronomeTuplet":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MetronomeTuplet), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Numeral":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Numeral), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NumeralKey":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NumeralKey), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "NumeralRoot":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NumeralRoot), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "OctaveShift":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OctaveShift), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Offset":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Offset), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "OtherDirection":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OtherDirection), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "OtherListening":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OtherListening), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Pedal":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Pedal), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PedalTuning":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PedalTuning), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PerMinute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PerMinute), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Percussion":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Percussion), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Pitched":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Pitched), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PrincipalVoice":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PrincipalVoice), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Print":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Print), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Root":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Root), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "RootStep":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RootStep), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Scordatura":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Scordatura), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Sound":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Sound), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StaffDivide":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StaffDivide), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Stick":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Stick), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StringMute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StringMute), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StyleText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StyleText), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Swing":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Swing), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Sync":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Sync), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TimeModification":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimeModification), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Timpani":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Timpani), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "VirtualInstrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(VirtualInstrument), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Wedge":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Wedge), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Wood":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Wood), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *Accord:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "String":
						inst.String = GongExtractInt(rhs)
					case "TuningStep":
						inst.TuningStep = attributes.EnumStep(GongExtractString(rhs))
					case "TuningAlter":
						inst.TuningAlter = GongExtractString(rhs)
					case "TuningOctave":
						inst.TuningOctave = GongExtractInt(rhs)
					}
				case *AccordionRegistration:
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
					case "Halign":
						inst.Halign = common.EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = common.EnumValign(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "AccordionMiddle":
						inst.AccordionMiddle = GongExtractInt(rhs)
					}
				case *Barre:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *Bass:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Arrangement":
						inst.Arrangement = EnumHarmonyArrangement(GongExtractString(rhs))
					case "BassSeparator":
						__gong__assignPointer(&inst.BassSeparator, rhs, identifierMap)
					case "BassStep":
						__gong__assignPointer(&inst.BassStep, rhs, identifierMap)
					case "BassAlter":
						__gong__assignPointer(&inst.BassAlter, rhs, identifierMap)
					}
				case *BassStep:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
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
						inst.Value = attributes.EnumStep(GongExtractString(rhs))
					}
				case *BeatUnitTied:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "BeatUnit":
						inst.BeatUnit = common.EnumNoteTypeValue(GongExtractString(rhs))
					}
				case *Beater:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tip":
						inst.Tip = EnumTipDirection(GongExtractString(rhs))
					case "Value":
						inst.Value = EnumBeaterValue(GongExtractString(rhs))
					}
				case *Bracket:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStopContinue(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "LineEnd":
						inst.LineEnd = EnumLineEnd(GongExtractString(rhs))
					case "EndLength":
						inst.EndLength = GongExtractString(rhs)
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
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Dashes:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStopContinue(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
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
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Degree:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "DegreeValue":
						__gong__assignPointer(&inst.DegreeValue, rhs, identifierMap)
					case "DegreeAlter":
						__gong__assignPointer(&inst.DegreeAlter, rhs, identifierMap)
					case "DegreeType":
						__gong__assignPointer(&inst.DegreeType, rhs, identifierMap)
					}
				case *DegreeAlter:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PlusMinus":
						inst.PlusMinus = common.EnumYesNo(GongExtractString(rhs))
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
				case *DegreeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
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
						inst.Value = EnumDegreeTypeValue(GongExtractString(rhs))
					}
				case *DegreeValue:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Symbol":
						inst.Symbol = EnumDegreeSymbolValue(GongExtractString(rhs))
					case "Text":
						inst.Text = GongExtractString(rhs)
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
						inst.Value = GongExtractInt(rhs)
					}
				case *Direction:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Directive":
						inst.Directive = common.EnumYesNo(GongExtractString(rhs))
					case "System":
						inst.System = barline.EnumSystemRelation(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "DirectionType":
						__gong__assignSliceOfPointers(&inst.DirectionType, rhs, identifierMap)
					case "Offset":
						__gong__assignPointer(&inst.Offset, rhs, identifierMap)
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					case "Sound":
						__gong__assignPointer(&inst.Sound, rhs, identifierMap)
					case "Listening":
						__gong__assignPointer(&inst.Listening, rhs, identifierMap)
					}
				case *DirectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Wedge":
						__gong__assignPointer(&inst.Wedge, rhs, identifierMap)
					case "Dashes":
						__gong__assignPointer(&inst.Dashes, rhs, identifierMap)
					case "Bracket":
						__gong__assignPointer(&inst.Bracket, rhs, identifierMap)
					case "Pedal":
						__gong__assignPointer(&inst.Pedal, rhs, identifierMap)
					case "Metronome":
						__gong__assignPointer(&inst.Metronome, rhs, identifierMap)
					case "OctaveShift":
						__gong__assignPointer(&inst.OctaveShift, rhs, identifierMap)
					case "HarpPedals":
						__gong__assignPointer(&inst.HarpPedals, rhs, identifierMap)
					case "StringMute":
						__gong__assignPointer(&inst.StringMute, rhs, identifierMap)
					case "Scordatura":
						__gong__assignPointer(&inst.Scordatura, rhs, identifierMap)
					case "Image":
						__gong__assignPointer(&inst.Image, rhs, identifierMap)
					case "PrincipalVoice":
						__gong__assignPointer(&inst.PrincipalVoice, rhs, identifierMap)
					case "Percussion":
						__gong__assignSliceOfPointers(&inst.Percussion, rhs, identifierMap)
					case "AccordionRegistration":
						__gong__assignPointer(&inst.AccordionRegistration, rhs, identifierMap)
					case "StaffDivide":
						__gong__assignPointer(&inst.StaffDivide, rhs, identifierMap)
					case "OtherDirection":
						__gong__assignPointer(&inst.OtherDirection, rhs, identifierMap)
					}
				case *Effect:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = EnumEffectValue(GongExtractString(rhs))
					}
				case *Feature:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *FirstFret:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "Location":
						inst.Location = common.EnumLeftRight(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractInt(rhs)
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
					case "Halign":
						inst.Halign = common.EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = common.EnumValignImage(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "FrameStrings":
						inst.FrameStrings = GongExtractInt(rhs)
					case "FrameFrets":
						inst.FrameFrets = GongExtractInt(rhs)
					case "FirstFret":
						__gong__assignPointer(&inst.FirstFret, rhs, identifierMap)
					case "FrameNote":
						__gong__assignSliceOfPointers(&inst.FrameNote, rhs, identifierMap)
					}
				case *FrameNote:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Barre":
						__gong__assignPointer(&inst.Barre, rhs, identifierMap)
					}
				case *Glass:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = EnumGlassValue(GongExtractString(rhs))
					}
				case *Grouping:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStopSingle(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "MemberOf":
						inst.MemberOf = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Feature":
						__gong__assignSliceOfPointers(&inst.Feature, rhs, identifierMap)
					}
				case *Harmony:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumHarmonyType(GongExtractString(rhs))
					case "PrintFrame":
						inst.PrintFrame = common.EnumYesNo(GongExtractString(rhs))
					case "Arrangement":
						inst.Arrangement = EnumHarmonyArrangement(GongExtractString(rhs))
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
					case "System":
						inst.System = barline.EnumSystemRelation(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Root":
						__gong__assignSliceOfPointers(&inst.Root, rhs, identifierMap)
					case "Numeral":
						__gong__assignSliceOfPointers(&inst.Numeral, rhs, identifierMap)
					case "Function":
						__gong__assignSliceOfPointers(&inst.Function, rhs, identifierMap)
					case "Kind":
						__gong__assignSliceOfPointers(&inst.Kind, rhs, identifierMap)
					case "Inversion":
						__gong__assignSliceOfPointers(&inst.Inversion, rhs, identifierMap)
					case "Bass":
						__gong__assignSliceOfPointers(&inst.Bass, rhs, identifierMap)
					case "Degree":
						__gong__assignSliceOfPointers(&inst.Degree, rhs, identifierMap)
					case "Frame":
						__gong__assignPointer(&inst.Frame, rhs, identifierMap)
					case "Offset":
						__gong__assignPointer(&inst.Offset, rhs, identifierMap)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					}
				case *HarmonyAlter:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						inst.Location = common.EnumLeftRight(GongExtractString(rhs))
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
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *HarpPedals:
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
					case "Halign":
						inst.Halign = common.EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = common.EnumValign(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "PedalTuning":
						__gong__assignSliceOfPointers(&inst.PedalTuning, rhs, identifierMap)
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
					case "DefaultX":
						inst.DefaultX = GongExtractString(rhs)
					case "DefaultY":
						inst.DefaultY = GongExtractString(rhs)
					case "RelativeX":
						inst.RelativeX = GongExtractString(rhs)
					case "RelativeY":
						inst.RelativeY = GongExtractString(rhs)
					case "Halign":
						inst.Halign = common.EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = common.EnumValignImage(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *InstrumentChange:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "InstrumentSound":
						inst.InstrumentSound = GongExtractString(rhs)
					case "Ensemble":
						inst.Ensemble = GongExtractString(rhs)
					case "VirtualInstrument":
						__gong__assignPointer(&inst.VirtualInstrument, rhs, identifierMap)
					}
				case *Inversion:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
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
						inst.Value = GongExtractInt(rhs)
					}
				case *Kind:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "UseSymbols":
						inst.UseSymbols = common.EnumYesNo(GongExtractString(rhs))
					case "Text":
						inst.Text = GongExtractString(rhs)
					case "StackDegrees":
						inst.StackDegrees = common.EnumYesNo(GongExtractString(rhs))
					case "ParenthesesDegrees":
						inst.ParenthesesDegrees = common.EnumYesNo(GongExtractString(rhs))
					case "BracketDegrees":
						inst.BracketDegrees = common.EnumYesNo(GongExtractString(rhs))
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
					case "Value":
						inst.Value = EnumKindValue(GongExtractString(rhs))
					}
				case *Listening:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Sync":
						__gong__assignSliceOfPointers(&inst.Sync, rhs, identifierMap)
					case "OtherListening":
						__gong__assignSliceOfPointers(&inst.OtherListening, rhs, identifierMap)
					case "Offset":
						__gong__assignPointer(&inst.Offset, rhs, identifierMap)
					}
				case *MeasureNumbering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "System":
						inst.System = barline.EnumSystemRelationNumber(GongExtractString(rhs))
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					case "MultipleRestAlways":
						inst.MultipleRestAlways = common.EnumYesNo(GongExtractString(rhs))
					case "MultipleRestRange":
						inst.MultipleRestRange = common.EnumYesNo(GongExtractString(rhs))
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
					case "Value":
						inst.Value = EnumMeasureNumberingValue(GongExtractString(rhs))
					}
				case *Membrane:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = EnumMembraneValue(GongExtractString(rhs))
					}
				case *Metal:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = EnumMetalValue(GongExtractString(rhs))
					}
				case *Metronome:
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
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "Justify":
						inst.Justify = common.EnumLeftCenterRight(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "BeatUnitTied":
						__gong__assignSliceOfPointers(&inst.BeatUnitTied, rhs, identifierMap)
					case "PerMinute":
						__gong__assignPointer(&inst.PerMinute, rhs, identifierMap)
					case "MetronomeNote":
						__gong__assignSliceOfPointers(&inst.MetronomeNote, rhs, identifierMap)
					case "MetronomeRelation":
						inst.MetronomeRelation = GongExtractString(rhs)
					}
				case *MetronomeBeam:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Value":
						inst.Value = EnumBeamValue(GongExtractString(rhs))
					}
				case *MetronomeNote:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MetronomeType":
						inst.MetronomeType = common.EnumNoteTypeValue(GongExtractString(rhs))
					case "MetronomeBeam":
						__gong__assignSliceOfPointers(&inst.MetronomeBeam, rhs, identifierMap)
					case "MetronomeTied":
						__gong__assignPointer(&inst.MetronomeTied, rhs, identifierMap)
					case "MetronomeTuplet":
						__gong__assignPointer(&inst.MetronomeTuplet, rhs, identifierMap)
					}
				case *MetronomeTied:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					}
				case *MetronomeTuplet:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Bracket":
						inst.Bracket = common.EnumYesNo(GongExtractString(rhs))
					case "ShowNumber":
						inst.ShowNumber = EnumShowTuplet(GongExtractString(rhs))
					case "ActualNotes":
						inst.ActualNotes = GongExtractInt(rhs)
					case "NormalNotes":
						inst.NormalNotes = GongExtractInt(rhs)
					case "NormalType":
						inst.NormalType = common.EnumNoteTypeValue(GongExtractString(rhs))
					}
				case *Numeral:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumeralRoot":
						__gong__assignPointer(&inst.NumeralRoot, rhs, identifierMap)
					case "NumeralAlter":
						__gong__assignPointer(&inst.NumeralAlter, rhs, identifierMap)
					case "NumeralKey":
						__gong__assignPointer(&inst.NumeralKey, rhs, identifierMap)
					}
				case *NumeralKey:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "NumeralFifths":
						inst.NumeralFifths = GongExtractInt(rhs)
					case "NumeralMode":
						inst.NumeralMode = EnumNumeralMode(GongExtractString(rhs))
					}
				case *NumeralRoot:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
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
						inst.Value = GongExtractInt(rhs)
					}
				case *OctaveShift:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumUpDownStopContinue(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Size":
						inst.Size = GongExtractInt(rhs)
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
					}
				case *Offset:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Sound":
						inst.Sound = common.EnumYesNo(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *OtherDirection:
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
					case "Halign":
						inst.Halign = common.EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = common.EnumValign(GongExtractString(rhs))
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *OtherListening:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "TimeOnly":
						inst.TimeOnly = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Pedal:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumPedalType(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Line":
						inst.Line = common.EnumYesNo(GongExtractString(rhs))
					case "Sign":
						inst.Sign = common.EnumYesNo(GongExtractString(rhs))
					case "Abbreviated":
						inst.Abbreviated = common.EnumYesNo(GongExtractString(rhs))
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
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *PedalTuning:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PedalStep":
						inst.PedalStep = attributes.EnumStep(GongExtractString(rhs))
					case "PedalAlter":
						inst.PedalAlter = GongExtractString(rhs)
					}
				case *PerMinute:
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
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Percussion:
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
					case "Halign":
						inst.Halign = common.EnumLeftCenterRight(GongExtractString(rhs))
					case "Valign":
						inst.Valign = common.EnumValign(GongExtractString(rhs))
					case "Enclosure":
						inst.Enclosure = common.EnumEnclosureShape(GongExtractString(rhs))
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
					case "StickLocation":
						inst.StickLocation = EnumStickLocation(GongExtractString(rhs))
					}
				case *Pitched:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = EnumPitchedValue(GongExtractString(rhs))
					}
				case *PrincipalVoice:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Symbol":
						inst.Symbol = EnumPrincipalVoiceSymbol(GongExtractString(rhs))
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
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Print:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StaffSpacing":
						inst.StaffSpacing = GongExtractString(rhs)
					case "NewSystem":
						inst.NewSystem = common.EnumYesNo(GongExtractString(rhs))
					case "NewPage":
						inst.NewPage = common.EnumYesNo(GongExtractString(rhs))
					case "BlankPage":
						inst.BlankPage = GongExtractInt(rhs)
					case "PageNumber":
						inst.PageNumber = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "MeasureNumbering":
						__gong__assignPointer(&inst.MeasureNumbering, rhs, identifierMap)
					}
				case *Root:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RootStep":
						__gong__assignPointer(&inst.RootStep, rhs, identifierMap)
					case "RootAlter":
						__gong__assignPointer(&inst.RootAlter, rhs, identifierMap)
					}
				case *RootStep:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
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
						inst.Value = attributes.EnumStep(GongExtractString(rhs))
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
				case *Sound:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tempo":
						inst.Tempo = GongExtractString(rhs)
					case "Dynamics":
						inst.Dynamics = GongExtractString(rhs)
					case "Dacapo":
						inst.Dacapo = common.EnumYesNo(GongExtractString(rhs))
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
					case "ForwardRepeat":
						inst.ForwardRepeat = common.EnumYesNo(GongExtractString(rhs))
					case "Fine":
						inst.Fine = GongExtractString(rhs)
					case "TimeOnly":
						inst.TimeOnly = GongExtractString(rhs)
					case "Pizzicato":
						inst.Pizzicato = common.EnumYesNo(GongExtractString(rhs))
					case "Pan":
						inst.Pan = GongExtractString(rhs)
					case "Elevation":
						inst.Elevation = GongExtractString(rhs)
					case "DamperPedal":
						inst.DamperPedal = GongExtractString(rhs)
					case "SoftPedal":
						inst.SoftPedal = GongExtractString(rhs)
					case "SostenutoPedal":
						inst.SostenutoPedal = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "InstrumentChange":
						__gong__assignSliceOfPointers(&inst.InstrumentChange, rhs, identifierMap)
					case "Swing":
						__gong__assignPointer(&inst.Swing, rhs, identifierMap)
					case "Offset":
						__gong__assignPointer(&inst.Offset, rhs, identifierMap)
					}
				case *StaffDivide:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumStaffDivideSymbol(GongExtractString(rhs))
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
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *Stick:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tip":
						inst.Tip = EnumTipDirection(GongExtractString(rhs))
					case "Parentheses":
						inst.Parentheses = common.EnumYesNo(GongExtractString(rhs))
					case "DashedCircle":
						inst.DashedCircle = common.EnumYesNo(GongExtractString(rhs))
					case "StickType":
						inst.StickType = EnumStickType(GongExtractString(rhs))
					case "StickMaterial":
						inst.StickMaterial = EnumStickMaterial(GongExtractString(rhs))
					}
				case *StringMute:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumOnOff(GongExtractString(rhs))
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
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *StyleText:
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
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *Swing:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "First":
						inst.First = GongExtractInt(rhs)
					case "Second":
						inst.Second = GongExtractInt(rhs)
					case "SwingType":
						inst.SwingType = common.EnumSwingTypeValue(GongExtractString(rhs))
					case "SwingStyle":
						inst.SwingStyle = GongExtractString(rhs)
					}
				case *Sync:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumSyncType(GongExtractString(rhs))
					case "Latency":
						inst.Latency = GongExtractInt(rhs)
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "TimeOnly":
						inst.TimeOnly = GongExtractString(rhs)
					}
				case *TimeModification:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ActualNotes":
						inst.ActualNotes = GongExtractInt(rhs)
					case "NormalNotes":
						inst.NormalNotes = GongExtractInt(rhs)
					case "NormalType":
						inst.NormalType = common.EnumNoteTypeValue(GongExtractString(rhs))
					}
				case *Timpani:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					}
				case *VirtualInstrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "VirtualLibrary":
						inst.VirtualLibrary = GongExtractString(rhs)
					case "VirtualName":
						inst.VirtualName = GongExtractString(rhs)
					}
				case *Wedge:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = EnumWedgeType(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Spread":
						inst.Spread = GongExtractString(rhs)
					case "Niente":
						inst.Niente = common.EnumYesNo(GongExtractString(rhs))
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
					case "Value":
						inst.Value = EnumWoodValue(GongExtractString(rhs))
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
