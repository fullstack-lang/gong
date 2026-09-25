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

	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/stool"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/clock"
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
	StoolStage *stool.Stage
	MusicStage *music.Stage
	ClockStage *clock.Stage
}


// Commit commits all stages in StageSet in dependency order
func (stageSet *StageSet) Commit() {
	if stageSet.ClockStage != nil {
		stageSet.ClockStage.Commit()
	}
	if stageSet.MusicStage != nil {
		stageSet.MusicStage.Commit()
	}
	if stageSet.StoolStage != nil {
		stageSet.StoolStage.Commit()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.Commit()
	}
}

// Checkout checkouts all stages in StageSet
func (stageSet *StageSet) Checkout() {
	if stageSet.ClockStage != nil {
		stageSet.ClockStage.Checkout()
	}
	if stageSet.MusicStage != nil {
		stageSet.MusicStage.Checkout()
	}
	if stageSet.StoolStage != nil {
		stageSet.StoolStage.Checkout()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.Checkout()
	}
}

// Reset resets all stages in StageSet
func (stageSet *StageSet) Reset() {
	if stageSet.ClockStage != nil {
		stageSet.ClockStage.Reset()
	}
	if stageSet.MusicStage != nil {
		stageSet.MusicStage.Reset()
	}
	if stageSet.StoolStage != nil {
		stageSet.StoolStage.Reset()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.Reset()
	}
}

// Clean cleans all stages in StageSet in dependency order
func (stageSet *StageSet) Clean() {
	if stageSet.ClockStage != nil {
		stageSet.ClockStage.Clean()
	}
	if stageSet.MusicStage != nil {
		stageSet.MusicStage.Clean()
	}
	if stageSet.StoolStage != nil {
		stageSet.StoolStage.Clean()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.Clean()
	}
}

// ComputeReverseMaps computes reverse maps on all stages in StageSet
func (stageSet *StageSet) ComputeReverseMaps() {
	if stageSet.ClockStage != nil {
		stageSet.ClockStage.ComputeReverseMaps()
	}
	if stageSet.MusicStage != nil {
		stageSet.MusicStage.ComputeReverseMaps()
	}
	if stageSet.StoolStage != nil {
		stageSet.StoolStage.ComputeReverseMaps()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReverseMaps()
	}
}

// ComputeInstancesNb computes instances nb on all stages in StageSet
func (stageSet *StageSet) ComputeInstancesNb() {
	if stageSet.ClockStage != nil {
		stageSet.ClockStage.ComputeInstancesNb()
	}
	if stageSet.MusicStage != nil {
		stageSet.MusicStage.ComputeInstancesNb()
	}
	if stageSet.StoolStage != nil {
		stageSet.StoolStage.ComputeInstancesNb()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeInstancesNb()
	}
}

// ComputeReferenceAndOrders computes reference and orders on all stages in StageSet
func (stageSet *StageSet) ComputeReferenceAndOrders() {
	if stageSet.ClockStage != nil {
		stageSet.ClockStage.ComputeReferenceAndOrders()
	}
	if stageSet.MusicStage != nil {
		stageSet.MusicStage.ComputeReferenceAndOrders()
	}
	if stageSet.StoolStage != nil {
		stageSet.StoolStage.ComputeReferenceAndOrders()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReferenceAndOrders()
	}
}

// NewStageSet creates a StageSet with all stages initialized
func NewStageSet(path string) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = NewStage(path)
	subPath_StoolStage := "stool"
	if path != "" {
	subPath_StoolStage = path + "_stool"
	}
	stageSet.StoolStage = stool.NewStage(subPath_StoolStage)
	subPath_MusicStage := "music"
	if path != "" {
	subPath_MusicStage = path + "_music"
	}
	stageSet.MusicStage = music.NewStage(subPath_MusicStage)
	subPath_ClockStage := "clock"
	if path != "" {
	subPath_ClockStage = path + "_clock"
	}
	stageSet.ClockStage = clock.NewStage(subPath_ClockStage)
	return stageSet
}

// NewStageSetFromStage creates a StageSet using an existing root stage
func NewStageSetFromStage(stage *Stage) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = stage
	subPath_StoolStage := "stool"
	if stage != nil && stage.GetName() != "" {
	subPath_StoolStage = stage.GetName() + "_stool"
	}
	stageSet.StoolStage = stool.NewStage(subPath_StoolStage)
	subPath_MusicStage := "music"
	if stage != nil && stage.GetName() != "" {
	subPath_MusicStage = stage.GetName() + "_music"
	}
	stageSet.MusicStage = music.NewStage(subPath_MusicStage)
	subPath_ClockStage := "clock"
	if stage != nil && stage.GetName() != "" {
	subPath_ClockStage = stage.GetName() + "_clock"
	}
	stageSet.ClockStage = clock.NewStage(subPath_ClockStage)
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

	if stageSet.ClockStage != nil {
		for _, clockabstract := range __gong__sortStageSetInstances(stageSet.ClockStage.ClockAbstracts, stageSet.ClockStage.ClockAbstract_stagedOrder) {
			if lastStageDecl != "ClockStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ClockStage"
			}
			clockabstractIdent := "__clock" + clockabstract.GongGetIdentifier(stageSet.ClockStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&clock.ClockAbstract{Name: %s}).Stage(stageSet.ClockStage)", clockabstractIdent, __gong__toRawStringLiteral(clockabstract.Name)))
			if lastStageVal != "ClockStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ClockStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", clockabstractIdent, __gong__toRawStringLiteral(clockabstract.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialRepetitions = %d", clockabstractIdent, clockabstract.RadialRepetitions))
			values.WriteString(fmt.Sprintf("\n\t%s.Transparency = %f", clockabstractIdent, clockabstract.Transparency))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeTubeDiameter = %f", clockabstractIdent, clockabstract.RelativeTubeDiameter))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeHeight3DTorus = %f", clockabstractIdent, clockabstract.RelativeHeight3DTorus))
			values.WriteString(fmt.Sprintf("\n\t%s.ClockTorusVerticalScale = %f", clockabstractIdent, clockabstract.ClockTorusVerticalScale))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeHeight = %f", clockabstractIdent, clockabstract.RelativeHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.ProjectionAngle = %f", clockabstractIdent, clockabstract.ProjectionAngle))
		}
	}
	if stageSet.MusicStage != nil {
		for _, musicabstract := range __gong__sortStageSetInstances(stageSet.MusicStage.MusicAbstracts, stageSet.MusicStage.MusicAbstract_stagedOrder) {
			if lastStageDecl != "MusicStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "MusicStage"
			}
			musicabstractIdent := "__music" + musicabstract.GongGetIdentifier(stageSet.MusicStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&music.MusicAbstract{Name: %s}).Stage(stageSet.MusicStage)", musicabstractIdent, __gong__toRawStringLiteral(musicabstract.Name)))
			if lastStageVal != "MusicStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "MusicStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", musicabstractIdent, __gong__toRawStringLiteral(musicabstract.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", musicabstractIdent, musicabstract.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.PitchHeight = %f", musicabstractIdent, musicabstract.PitchHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.NbOfBeatsInTheme = %d", musicabstractIdent, musicabstract.NbOfBeatsInTheme))
			values.WriteString(fmt.Sprintf("\n\t%s.BeatsPerSecond = %f", musicabstractIdent, musicabstract.BeatsPerSecond))
			values.WriteString(fmt.Sprintf("\n\t%s.FirstVoiceShiftX = %f", musicabstractIdent, musicabstract.FirstVoiceShiftX))
			values.WriteString(fmt.Sprintf("\n\t%s.FirstVoiceShiftY = %f", musicabstractIdent, musicabstract.FirstVoiceShiftY))
			values.WriteString(fmt.Sprintf("\n\t%s.PitchDifference = %d", musicabstractIdent, musicabstract.PitchDifference))
			values.WriteString(fmt.Sprintf("\n\t%s.Level = %f", musicabstractIdent, musicabstract.Level))
			values.WriteString(fmt.Sprintf("\n\t%s.ActualBeatsTemporalShift = %d", musicabstractIdent, musicabstract.ActualBeatsTemporalShift))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMinor = %t", musicabstractIdent, musicabstract.IsMinor))
			values.WriteString(fmt.Sprintf("\n\t%s.ThemeBinaryEncoding = %d", musicabstractIdent, musicabstract.ThemeBinaryEncoding))
			values.WriteString(fmt.Sprintf("\n\t%s.BezierControlLengthRatio = %f", musicabstractIdent, musicabstract.BezierControlLengthRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPitchLines = %d", musicabstractIdent, musicabstract.NbPitchLines))
			values.WriteString(fmt.Sprintf("\n\t%s.NbBeatLines = %d", musicabstractIdent, musicabstract.NbBeatLines))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginX = %f", musicabstractIdent, musicabstract.OriginX))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginY = %f", musicabstractIdent, musicabstract.OriginY))
			values.WriteString(fmt.Sprintf("\n\t%s.ScoreScale = %f", musicabstractIdent, musicabstract.ScoreScale))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowFirstVoice = %t", musicabstractIdent, musicabstract.ShowFirstVoice))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowFirstVoiceShiftRight = %t", musicabstractIdent, musicabstract.ShowFirstVoiceShiftRight))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowSecondVoice = %t", musicabstractIdent, musicabstract.ShowSecondVoice))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowSecondVoiceShiftRight = %t", musicabstractIdent, musicabstract.ShowSecondVoiceShiftRight))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowFirstVoiceNotes = %t", musicabstractIdent, musicabstract.ShowFirstVoiceNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowFirstVoiceNotesShiftRight = %t", musicabstractIdent, musicabstract.ShowFirstVoiceNotesShiftRight))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowSecondVoiceNotes = %t", musicabstractIdent, musicabstract.ShowSecondVoiceNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowSecondVoiceNotesShiftRight = %t", musicabstractIdent, musicabstract.ShowSecondVoiceNotesShiftRight))
			values.WriteString(fmt.Sprintf("\n\t%s.IsComposerNodeExpanded = %t", musicabstractIdent, musicabstract.IsComposerNodeExpanded))
		}
	}
	if stageSet.StoolStage != nil {
		for _, stoolabstract := range __gong__sortStageSetInstances(stageSet.StoolStage.StoolAbstracts, stageSet.StoolStage.StoolAbstract_stagedOrder) {
			if lastStageDecl != "StoolStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "StoolStage"
			}
			stoolabstractIdent := "__stool" + stoolabstract.GongGetIdentifier(stageSet.StoolStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&stool.StoolAbstract{Name: %s}).Stage(stageSet.StoolStage)", stoolabstractIdent, __gong__toRawStringLiteral(stoolabstract.Name)))
			if lastStageVal != "StoolStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "StoolStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stoolabstractIdent, __gong__toRawStringLiteral(stoolabstract.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialRepetitions = %d", stoolabstractIdent, stoolabstract.RadialRepetitions))
			values.WriteString(fmt.Sprintf("\n\t%s.Transparency = %f", stoolabstractIdent, stoolabstract.Transparency))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeTubeDiameter = %f", stoolabstractIdent, stoolabstract.RelativeTubeDiameter))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeHeight3DTorus = %f", stoolabstractIdent, stoolabstract.RelativeHeight3DTorus))
			values.WriteString(fmt.Sprintf("\n\t%s.StoolTorusVerticalScale = %f", stoolabstractIdent, stoolabstract.StoolTorusVerticalScale))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeHeight = %f", stoolabstractIdent, stoolabstract.RelativeHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeSeatThickness = %f", stoolabstractIdent, stoolabstract.RelativeSeatThickness))
			values.WriteString(fmt.Sprintf("\n\t%s.ProjectionAngle = %f", stoolabstractIdent, stoolabstract.ProjectionAngle))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeEyeSeparationCriteria = %f", stoolabstractIdent, stoolabstract.RelativeEyeSeparationCriteria))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeEyeCornerControlVectorStrength = %f", stoolabstractIdent, stoolabstract.RelativeEyeCornerControlVectorStrength))
		}
	}
	if stageSet.Stage != nil {
		for _, angle0shape := range __gong__sortStageSetInstances(stageSet.Stage.Angle0Shapes, stageSet.Stage.Angle0Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			angle0shapeIdent := "__models" + angle0shape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Angle0Shape{Name: %s}).Stage(stageSet.Stage)", angle0shapeIdent, __gong__toRawStringLiteral(angle0shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", angle0shapeIdent, __gong__toRawStringLiteral(angle0shape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, bottomcurveplane1shape := range __gong__sortStageSetInstances(stageSet.Stage.BottomCurvePlane1Shapes, stageSet.Stage.BottomCurvePlane1Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bottomcurveplane1shapeIdent := "__models" + bottomcurveplane1shape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.BottomCurvePlane1Shape{Name: %s}).Stage(stageSet.Stage)", bottomcurveplane1shapeIdent, __gong__toRawStringLiteral(bottomcurveplane1shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bottomcurveplane1shapeIdent, __gong__toRawStringLiteral(bottomcurveplane1shape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, bottomcurveplane2shape := range __gong__sortStageSetInstances(stageSet.Stage.BottomCurvePlane2Shapes, stageSet.Stage.BottomCurvePlane2Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bottomcurveplane2shapeIdent := "__models" + bottomcurveplane2shape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.BottomCurvePlane2Shape{Name: %s}).Stage(stageSet.Stage)", bottomcurveplane2shapeIdent, __gong__toRawStringLiteral(bottomcurveplane2shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bottomcurveplane2shapeIdent, __gong__toRawStringLiteral(bottomcurveplane2shape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, circumference3dshape := range __gong__sortStageSetInstances(stageSet.Stage.Circumference3DShapes, stageSet.Stage.Circumference3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			circumference3dshapeIdent := "__models" + circumference3dshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Circumference3DShape{Name: %s}).Stage(stageSet.Stage)", circumference3dshapeIdent, __gong__toRawStringLiteral(circumference3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", circumference3dshapeIdent, __gong__toRawStringLiteral(circumference3dshape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, clock2ddiagram := range __gong__sortStageSetInstances(stageSet.Stage.Clock2DDiagrams, stageSet.Stage.Clock2DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			clock2ddiagramIdent := "__models" + clock2ddiagram.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Clock2DDiagram{Name: %s}).Stage(stageSet.Stage)", clock2ddiagramIdent, __gong__toRawStringLiteral(clock2ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", clock2ddiagramIdent, __gong__toRawStringLiteral(clock2ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Zoom = %f", clock2ddiagramIdent, clock2ddiagram.Zoom))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenAxesShape = %t", clock2ddiagramIdent, clock2ddiagram.IsHiddenAxesShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", clock2ddiagramIdent, clock2ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", clock2ddiagramIdent, __gong__toRawStringLiteral(clock2ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", clock2ddiagramIdent, clock2ddiagram.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		for _, clock3ddiagram := range __gong__sortStageSetInstances(stageSet.Stage.Clock3DDiagrams, stageSet.Stage.Clock3DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			clock3ddiagramIdent := "__models" + clock3ddiagram.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Clock3DDiagram{Name: %s}).Stage(stageSet.Stage)", clock3ddiagramIdent, __gong__toRawStringLiteral(clock3ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", clock3ddiagramIdent, __gong__toRawStringLiteral(clock3ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenClockTopCurveShape = %t", clock3ddiagramIdent, clock3ddiagram.IsHiddenClockTopCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTorus3DShape = %t", clock3ddiagramIdent, clock3ddiagram.IsHiddenTorus3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSampledPoints3DShape = %t", clock3ddiagramIdent, clock3ddiagram.IsHiddenSampledPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTiledFloor3DShape = %t", clock3ddiagramIdent, clock3ddiagram.IsHiddenTiledFloor3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", clock3ddiagramIdent, clock3ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", clock3ddiagramIdent, __gong__toRawStringLiteral(clock3ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", clock3ddiagramIdent, clock3ddiagram.IsExpanded))
			if clock3ddiagram.SampledPoints3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + clock3ddiagram.SampledPoints3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SampledPoints3DShape = %s", clock3ddiagramIdent, targetIdent))
			}
			if clock3ddiagram.Rendered3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + clock3ddiagram.Rendered3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rendered3DShape = %s", clock3ddiagramIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, cutline3dshape := range __gong__sortStageSetInstances(stageSet.Stage.CutLine3DShapes, stageSet.Stage.CutLine3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cutline3dshapeIdent := "__models" + cutline3dshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.CutLine3DShape{Name: %s}).Stage(stageSet.Stage)", cutline3dshapeIdent, __gong__toRawStringLiteral(cutline3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cutline3dshapeIdent, __gong__toRawStringLiteral(cutline3dshape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, leaves3dshape := range __gong__sortStageSetInstances(stageSet.Stage.Leaves3DShapes, stageSet.Stage.Leaves3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			leaves3dshapeIdent := "__models" + leaves3dshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Leaves3DShape{Name: %s}).Stage(stageSet.Stage)", leaves3dshapeIdent, __gong__toRawStringLiteral(leaves3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", leaves3dshapeIdent, __gong__toRawStringLiteral(leaves3dshape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, library := range __gong__sortStageSetInstances(stageSet.Stage.Librarys, stageSet.Stage.Library_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			libraryIdent := "__models" + library.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Library{Name: %s}).Stage(stageSet.Stage)", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", libraryIdent, __gong__toRawStringLiteral(library.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", libraryIdent, __gong__toRawStringLiteral(library.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", libraryIdent, library.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRootLibrary = %t", libraryIdent, library.IsRootLibrary))
			for _, elem := range library.Plants {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Plants = append(%s.Plants, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SubLibraries {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibraries = append(%s.SubLibraries, %s)", libraryIdent, libraryIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, originalpoints3dshape := range __gong__sortStageSetInstances(stageSet.Stage.OriginalPoints3DShapes, stageSet.Stage.OriginalPoints3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			originalpoints3dshapeIdent := "__models" + originalpoints3dshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.OriginalPoints3DShape{Name: %s}).Stage(stageSet.Stage)", originalpoints3dshapeIdent, __gong__toRawStringLiteral(originalpoints3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", originalpoints3dshapeIdent, __gong__toRawStringLiteral(originalpoints3dshape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, parastichymcurves3dshape := range __gong__sortStageSetInstances(stageSet.Stage.ParastichyMCurves3DShapes, stageSet.Stage.ParastichyMCurves3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parastichymcurves3dshapeIdent := "__models" + parastichymcurves3dshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParastichyMCurves3DShape{Name: %s}).Stage(stageSet.Stage)", parastichymcurves3dshapeIdent, __gong__toRawStringLiteral(parastichymcurves3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parastichymcurves3dshapeIdent, __gong__toRawStringLiteral(parastichymcurves3dshape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, parastichyncurves3dshape := range __gong__sortStageSetInstances(stageSet.Stage.ParastichyNCurves3DShapes, stageSet.Stage.ParastichyNCurves3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			parastichyncurves3dshapeIdent := "__models" + parastichyncurves3dshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParastichyNCurves3DShape{Name: %s}).Stage(stageSet.Stage)", parastichyncurves3dshapeIdent, __gong__toRawStringLiteral(parastichyncurves3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parastichyncurves3dshapeIdent, __gong__toRawStringLiteral(parastichyncurves3dshape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, plant2ddiagram := range __gong__sortStageSetInstances(stageSet.Stage.Plant2DDiagrams, stageSet.Stage.Plant2DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			plant2ddiagramIdent := "__models" + plant2ddiagram.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Plant2DDiagram{Name: %s}).Stage(stageSet.Stage)", plant2ddiagramIdent, __gong__toRawStringLiteral(plant2ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", plant2ddiagramIdent, __gong__toRawStringLiteral(plant2ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginX = %f", plant2ddiagramIdent, plant2ddiagram.OriginX))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginY = %f", plant2ddiagramIdent, plant2ddiagram.OriginY))
			values.WriteString(fmt.Sprintf("\n\t%s.Zoom = %f", plant2ddiagramIdent, plant2ddiagram.Zoom))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRhombusNodesExpanded = %t", plant2ddiagramIdent, plant2ddiagram.IsRhombusNodesExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsArcNodesExpanded = %t", plant2ddiagramIdent, plant2ddiagram.IsArcNodesExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenAxesShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenAxesShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenReferenceRhombus = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenReferenceRhombus))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPlantCircumferenceShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenPlantCircumferenceShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenGridPathShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenGridPathShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRhombusGridShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenRhombusGridShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenExplanationTextShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenExplanationTextShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedReferenceRhombus = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenRotatedReferenceRhombus))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedPlantCircumferenceShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenRotatedPlantCircumferenceShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedGridPathShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenRotatedGridPathShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedRhombusGridShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenRotatedRhombusGridShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenGrowthPathRhombusGridShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenGrowthPathRhombusGridShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenGrowthVectorShape = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenGrowthVectorShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPerpendicularVectorGrid = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenPerpendicularVectorGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenBaseVectorShapeGrid = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenBaseVectorShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenArcNormalVectorShapeGrid = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenArcNormalVectorShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStartArcShapeGrid = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenStartArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenMidArcVectorShapeGrid = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenMidArcVectorShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEndArcShapeGrid = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenEndArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenGrowthCurve2D = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenGrowthCurve2D))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfGrowthCurve2DByGrowthVector = %t", plant2ddiagramIdent, plant2ddiagram.IsHiddenStackOfGrowthCurve2DByGrowthVector))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", plant2ddiagramIdent, plant2ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", plant2ddiagramIdent, __gong__toRawStringLiteral(plant2ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", plant2ddiagramIdent, plant2ddiagram.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		for _, plant3ddiagram := range __gong__sortStageSetInstances(stageSet.Stage.Plant3DDiagrams, stageSet.Stage.Plant3DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			plant3ddiagramIdent := "__models" + plant3ddiagram.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Plant3DDiagram{Name: %s}).Stage(stageSet.Stage)", plant3ddiagramIdent, __gong__toRawStringLiteral(plant3ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", plant3ddiagramIdent, __gong__toRawStringLiteral(plant3ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStemCylinder3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenStemCylinder3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenParastichyNCurves3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenParastichyNCurves3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenParastichyMCurves3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenParastichyMCurves3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenCutLine3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenCutLine3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenCircumference3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenCircumference3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTiledFloor3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenTiledFloor3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenLeaves3DShape = %t", plant3ddiagramIdent, plant3ddiagram.IsHiddenLeaves3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", plant3ddiagramIdent, plant3ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", plant3ddiagramIdent, __gong__toRawStringLiteral(plant3ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", plant3ddiagramIdent, plant3ddiagram.IsExpanded))
			if plant3ddiagram.StemCylinder3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plant3ddiagram.StemCylinder3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StemCylinder3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
			if plant3ddiagram.ParastichyNCurves3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plant3ddiagram.ParastichyNCurves3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParastichyNCurves3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
			if plant3ddiagram.ParastichyMCurves3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plant3ddiagram.ParastichyMCurves3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParastichyMCurves3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
			if plant3ddiagram.CutLine3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plant3ddiagram.CutLine3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CutLine3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
			if plant3ddiagram.Circumference3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plant3ddiagram.Circumference3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Circumference3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
			if plant3ddiagram.Leaves3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plant3ddiagram.Leaves3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Leaves3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
			if plant3ddiagram.Rendered3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plant3ddiagram.Rendered3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rendered3DShape = %s", plant3ddiagramIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, plantabstract := range __gong__sortStageSetInstances(stageSet.Stage.PlantAbstracts, stageSet.Stage.PlantAbstract_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			plantabstractIdent := "__models" + plantabstract.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PlantAbstract{Name: %s}).Stage(stageSet.Stage)", plantabstractIdent, __gong__toRawStringLiteral(plantabstract.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", plantabstractIdent, __gong__toRawStringLiteral(plantabstract.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.N = %d", plantabstractIdent, plantabstract.N))
			values.WriteString(fmt.Sprintf("\n\t%s.M = %d", plantabstractIdent, plantabstract.M))
			values.WriteString(fmt.Sprintf("\n\t%s.StackHeight = %d", plantabstractIdent, plantabstract.StackHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.RhombusInsideAngle = %f", plantabstractIdent, plantabstract.RhombusInsideAngle))
			values.WriteString(fmt.Sprintf("\n\t%s.RhombusSideLength = %f", plantabstractIdent, plantabstract.RhombusSideLength))
			values.WriteString(fmt.Sprintf("\n\t%s.PlantType = %s", plantabstractIdent, __gong__toRawStringLiteral(string(plantabstract.PlantType))))
			values.WriteString(fmt.Sprintf("\n\t%s.CurrentView = %s", plantabstractIdent, __gong__toRawStringLiteral(string(plantabstract.CurrentView))))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", plantabstractIdent, __gong__toRawStringLiteral(plantabstract.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", plantabstractIdent, plantabstract.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelected = %t", plantabstractIdent, plantabstract.IsSelected))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPlant2DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsPlant2DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsPlant3DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsPlant3DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsVase2DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsVase2DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsTubeVase3DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsTubeVase3DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStool2DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsStool2DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStool3DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsStool3DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsClock2DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsClock2DDiagramsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsClock3DDiagramsNodeExpanded = %t", plantabstractIdent, plantabstract.IsClock3DDiagramsNodeExpanded))
			if plantabstract.TubeVaseAbstract != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.TubeVaseAbstract.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TubeVaseAbstract = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.StoolAbstract != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__stool" + plantabstract.StoolAbstract.GongGetIdentifier(stageSet.StoolStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StoolAbstract = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.ClockAbstract != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__clock" + plantabstract.ClockAbstract.GongGetIdentifier(stageSet.ClockStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClockAbstract = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.MusicAbstract != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__music" + plantabstract.MusicAbstract.GongGetIdentifier(stageSet.MusicStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MusicAbstract = %s", plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Plant2DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Plant2DDiagrams = append(%s.Plant2DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Plant3DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Plant3DDiagrams = append(%s.Plant3DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Vase2DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Vase2DDiagrams = append(%s.Vase2DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.TubeVase3DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TubeVase3DDiagrams = append(%s.TubeVase3DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Stool2DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stool2DDiagrams = append(%s.Stool2DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Stool3DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stool3DDiagrams = append(%s.Stool3DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Clock2DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Clock2DDiagrams = append(%s.Clock2DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
			for _, elem := range plantabstract.Clock3DDiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Clock3DDiagrams = append(%s.Clock3DDiagrams, %s)", plantabstractIdent, plantabstractIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, rendered3dshape := range __gong__sortStageSetInstances(stageSet.Stage.Rendered3DShapes, stageSet.Stage.Rendered3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rendered3dshapeIdent := "__models" + rendered3dshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Rendered3DShape{Name: %s}).Stage(stageSet.Stage)", rendered3dshapeIdent, __gong__toRawStringLiteral(rendered3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rendered3dshapeIdent, __gong__toRawStringLiteral(rendered3dshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ViewX = %f", rendered3dshapeIdent, rendered3dshape.ViewX))
			values.WriteString(fmt.Sprintf("\n\t%s.ViewY = %f", rendered3dshapeIdent, rendered3dshape.ViewY))
			values.WriteString(fmt.Sprintf("\n\t%s.ViewZ = %f", rendered3dshapeIdent, rendered3dshape.ViewZ))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetX = %f", rendered3dshapeIdent, rendered3dshape.TargetX))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetY = %f", rendered3dshapeIdent, rendered3dshape.TargetY))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetZ = %f", rendered3dshapeIdent, rendered3dshape.TargetZ))
			values.WriteString(fmt.Sprintf("\n\t%s.Fov = %f", rendered3dshapeIdent, rendered3dshape.Fov))
		}
	}
	if stageSet.Stage != nil {
		for _, sampledpoints3dshape := range __gong__sortStageSetInstances(stageSet.Stage.SampledPoints3DShapes, stageSet.Stage.SampledPoints3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sampledpoints3dshapeIdent := "__models" + sampledpoints3dshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SampledPoints3DShape{Name: %s}).Stage(stageSet.Stage)", sampledpoints3dshapeIdent, __gong__toRawStringLiteral(sampledpoints3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sampledpoints3dshapeIdent, __gong__toRawStringLiteral(sampledpoints3dshape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, stackofrotatedvasetrapezeringsshape := range __gong__sortStageSetInstances(stageSet.Stage.StackOfRotatedVaseTrapezeRingsShapes, stageSet.Stage.StackOfRotatedVaseTrapezeRingsShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stackofrotatedvasetrapezeringsshapeIdent := "__models" + stackofrotatedvasetrapezeringsshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StackOfRotatedVaseTrapezeRingsShape{Name: %s}).Stage(stageSet.Stage)", stackofrotatedvasetrapezeringsshapeIdent, __gong__toRawStringLiteral(stackofrotatedvasetrapezeringsshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stackofrotatedvasetrapezeringsshapeIdent, __gong__toRawStringLiteral(stackofrotatedvasetrapezeringsshape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, stackofvasetrapezeringsshape := range __gong__sortStageSetInstances(stageSet.Stage.StackOfVaseTrapezeRingsShapes, stageSet.Stage.StackOfVaseTrapezeRingsShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stackofvasetrapezeringsshapeIdent := "__models" + stackofvasetrapezeringsshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StackOfVaseTrapezeRingsShape{Name: %s}).Stage(stageSet.Stage)", stackofvasetrapezeringsshapeIdent, __gong__toRawStringLiteral(stackofvasetrapezeringsshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stackofvasetrapezeringsshapeIdent, __gong__toRawStringLiteral(stackofvasetrapezeringsshape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, stemcylinder3dshape := range __gong__sortStageSetInstances(stageSet.Stage.StemCylinder3DShapes, stageSet.Stage.StemCylinder3DShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stemcylinder3dshapeIdent := "__models" + stemcylinder3dshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StemCylinder3DShape{Name: %s}).Stage(stageSet.Stage)", stemcylinder3dshapeIdent, __gong__toRawStringLiteral(stemcylinder3dshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stemcylinder3dshapeIdent, __gong__toRawStringLiteral(stemcylinder3dshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transparency = %f", stemcylinder3dshapeIdent, stemcylinder3dshape.Transparency))
		}
	}
	if stageSet.Stage != nil {
		for _, stool2ddiagram := range __gong__sortStageSetInstances(stageSet.Stage.Stool2DDiagrams, stageSet.Stage.Stool2DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stool2ddiagramIdent := "__models" + stool2ddiagram.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Stool2DDiagram{Name: %s}).Stage(stageSet.Stage)", stool2ddiagramIdent, __gong__toRawStringLiteral(stool2ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stool2ddiagramIdent, __gong__toRawStringLiteral(stool2ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Zoom = %f", stool2ddiagramIdent, stool2ddiagram.Zoom))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenAxesShape = %t", stool2ddiagramIdent, stool2ddiagram.IsHiddenAxesShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", stool2ddiagramIdent, stool2ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", stool2ddiagramIdent, __gong__toRawStringLiteral(stool2ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", stool2ddiagramIdent, stool2ddiagram.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		for _, stool3ddiagram := range __gong__sortStageSetInstances(stageSet.Stage.Stool3DDiagrams, stageSet.Stage.Stool3DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stool3ddiagramIdent := "__models" + stool3ddiagram.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Stool3DDiagram{Name: %s}).Stage(stageSet.Stage)", stool3ddiagramIdent, __gong__toRawStringLiteral(stool3ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stool3ddiagramIdent, __gong__toRawStringLiteral(stool3ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSeatTopCurveShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenSeatTopCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedSeatTopCurveShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenRotatedSeatTopCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSeatBottomCurveShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenSeatBottomCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedSeatBottomCurveShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenRotatedSeatBottomCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTorus3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenTorus3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedTorusShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenRotatedTorusShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSampledPoints3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenSampledPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedSampledPoints3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenRotatedSampledPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEyeSampledPoints3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenEyeSampledPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEyeCornersSampledPoints3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenEyeCornersSampledPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEye3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenEye3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEyeSeatBottomCurveShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenEyeSeatBottomCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEyeStoolBottomCurveShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenEyeStoolBottomCurveShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSeat3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenSeat3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEyeVolume3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenEyeVolume3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSeatAndLegs3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenSeatAndLegs3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenRotatedSeatAndLegs3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenRotatedSeatAndLegs3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTiledFloor3DShape = %t", stool3ddiagramIdent, stool3ddiagram.IsHiddenTiledFloor3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", stool3ddiagramIdent, stool3ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", stool3ddiagramIdent, __gong__toRawStringLiteral(stool3ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", stool3ddiagramIdent, stool3ddiagram.IsExpanded))
			if stool3ddiagram.SampledPoints3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.SampledPoints3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SampledPoints3DShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.Rendered3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.Rendered3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rendered3DShape = %s", stool3ddiagramIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, topcurveplane1shape := range __gong__sortStageSetInstances(stageSet.Stage.TopCurvePlane1Shapes, stageSet.Stage.TopCurvePlane1Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			topcurveplane1shapeIdent := "__models" + topcurveplane1shape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TopCurvePlane1Shape{Name: %s}).Stage(stageSet.Stage)", topcurveplane1shapeIdent, __gong__toRawStringLiteral(topcurveplane1shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", topcurveplane1shapeIdent, __gong__toRawStringLiteral(topcurveplane1shape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, topcurveplane2shape := range __gong__sortStageSetInstances(stageSet.Stage.TopCurvePlane2Shapes, stageSet.Stage.TopCurvePlane2Shape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			topcurveplane2shapeIdent := "__models" + topcurveplane2shape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TopCurvePlane2Shape{Name: %s}).Stage(stageSet.Stage)", topcurveplane2shapeIdent, __gong__toRawStringLiteral(topcurveplane2shape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", topcurveplane2shapeIdent, __gong__toRawStringLiteral(topcurveplane2shape.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, tubevase3ddiagram := range __gong__sortStageSetInstances(stageSet.Stage.TubeVase3DDiagrams, stageSet.Stage.TubeVase3DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tubevase3ddiagramIdent := "__models" + tubevase3ddiagram.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TubeVase3DDiagram{Name: %s}).Stage(stageSet.Stage)", tubevase3ddiagramIdent, __gong__toRawStringLiteral(tubevase3ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tubevase3ddiagramIdent, __gong__toRawStringLiteral(tubevase3ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTorusStackShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenTorusStackShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenVerticalTorusStackShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenVerticalTorusStackShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPartiallyRotatedTorusShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenPartiallyRotatedTorusShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfPartiallyRotatedTorusShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenStackOfPartiallyRotatedTorusShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPointsAndLines3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenPointsAndLines3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenKeyHole3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenKeyHole3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenKey3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenKey3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenVolumeKey3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenVolumeKey3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTorusEdge3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenTorusEdge3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenSampledPoints3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenSampledPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenOriginalPoints3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenOriginalPoints3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenAngle0Shape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenAngle0Shape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTiledFloor3DShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenTiledFloor3DShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopCurvePlane1Shape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenTopCurvePlane1Shape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenBottomCurvePlane1Shape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenBottomCurvePlane1Shape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopCurvePlane2Shape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenTopCurvePlane2Shape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenBottomCurvePlane2Shape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenBottomCurvePlane2Shape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenVaseTrapezeRingShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenVaseTrapezeRingShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfVaseTrapezeRingsShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenStackOfVaseTrapezeRingsShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfRotatedVaseTrapezeRingsShape = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsHiddenStackOfRotatedVaseTrapezeRingsShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", tubevase3ddiagramIdent, __gong__toRawStringLiteral(tubevase3ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", tubevase3ddiagramIdent, tubevase3ddiagram.IsExpanded))
			if tubevase3ddiagram.Rendered3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.Rendered3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rendered3DShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.SampledPoints3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.SampledPoints3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SampledPoints3DShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.OriginalPoints3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.OriginalPoints3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OriginalPoints3DShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.Angle0Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.Angle0Shape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Angle0Shape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.TopCurvePlane1Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.TopCurvePlane1Shape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TopCurvePlane1Shape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.BottomCurvePlane1Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.BottomCurvePlane1Shape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BottomCurvePlane1Shape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.TopCurvePlane2Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.TopCurvePlane2Shape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TopCurvePlane2Shape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.BottomCurvePlane2Shape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.BottomCurvePlane2Shape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BottomCurvePlane2Shape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.VaseTrapezeRingShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.VaseTrapezeRingShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VaseTrapezeRingShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.StackOfVaseTrapezeRingsShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.StackOfVaseTrapezeRingsShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StackOfVaseTrapezeRingsShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.StackOfRotatedVaseTrapezeRingsShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.StackOfRotatedVaseTrapezeRingsShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StackOfRotatedVaseTrapezeRingsShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, tubevaseabstract := range __gong__sortStageSetInstances(stageSet.Stage.TubeVaseAbstracts, stageSet.Stage.TubeVaseAbstract_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tubevaseabstractIdent := "__models" + tubevaseabstract.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TubeVaseAbstract{Name: %s}).Stage(stageSet.Stage)", tubevaseabstractIdent, __gong__toRawStringLiteral(tubevaseabstract.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tubevaseabstractIdent, __gong__toRawStringLiteral(tubevaseabstract.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Z_Ribbon = %f", tubevaseabstractIdent, tubevaseabstract.Z_Ribbon))
			values.WriteString(fmt.Sprintf("\n\t%s.RibbonVerticalScale = %f", tubevaseabstractIdent, tubevaseabstract.RibbonVerticalScale))
			values.WriteString(fmt.Sprintf("\n\t%s.Plane1Height = %f", tubevaseabstractIdent, tubevaseabstract.Plane1Height))
			values.WriteString(fmt.Sprintf("\n\t%s.Plane2Height = %f", tubevaseabstractIdent, tubevaseabstract.Plane2Height))
			values.WriteString(fmt.Sprintf("\n\t%s.ProjectionAngle = %f", tubevaseabstractIdent, tubevaseabstract.ProjectionAngle))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeVerticalThickness = %f", tubevaseabstractIdent, tubevaseabstract.RelativeVerticalThickness))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeRadialThickness = %f", tubevaseabstractIdent, tubevaseabstract.RelativeRadialThickness))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeCuttedStackFloorHeight = %f", tubevaseabstractIdent, tubevaseabstract.RelativeCuttedStackFloorHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeRotatedTorusSeparation = %f", tubevaseabstractIdent, tubevaseabstract.RelativeRotatedTorusSeparation))
			values.WriteString(fmt.Sprintf("\n\t%s.RotationRatio = %f", tubevaseabstractIdent, tubevaseabstract.RotationRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.RadialRepetitions = %d", tubevaseabstractIdent, tubevaseabstract.RadialRepetitions))
			values.WriteString(fmt.Sprintf("\n\t%s.Transparency = %f", tubevaseabstractIdent, tubevaseabstract.Transparency))
			values.WriteString(fmt.Sprintf("\n\t%s.HasAlternatingRingColors = %t", tubevaseabstractIdent, tubevaseabstract.HasAlternatingRingColors))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeTrajectoryOffsetX = %f", tubevaseabstractIdent, tubevaseabstract.RelativeTrajectoryOffsetX))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeTrajectoryOffsetY = %f", tubevaseabstractIdent, tubevaseabstract.RelativeTrajectoryOffsetY))
			values.WriteString(fmt.Sprintf("\n\t%s.NbStepP1P2 = %d", tubevaseabstractIdent, tubevaseabstract.NbStepP1P2))
			values.WriteString(fmt.Sprintf("\n\t%s.ChosenStep = %d", tubevaseabstractIdent, tubevaseabstract.ChosenStep))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeHorizontalRingsHeight = %f", tubevaseabstractIdent, tubevaseabstract.RelativeHorizontalRingsHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.OffsetKeyX = %f", tubevaseabstractIdent, tubevaseabstract.OffsetKeyX))
			values.WriteString(fmt.Sprintf("\n\t%s.OffsetKeyY = %f", tubevaseabstractIdent, tubevaseabstract.OffsetKeyY))
			values.WriteString(fmt.Sprintf("\n\t%s.HeightKey = %f", tubevaseabstractIdent, tubevaseabstract.HeightKey))
			values.WriteString(fmt.Sprintf("\n\t%s.WidthKey = %f", tubevaseabstractIdent, tubevaseabstract.WidthKey))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeKeySize = %f", tubevaseabstractIdent, tubevaseabstract.RelativeKeySize))
			values.WriteString(fmt.Sprintf("\n\t%s.MovieNbFrames = %d", tubevaseabstractIdent, tubevaseabstract.MovieNbFrames))
		}
	}
	if stageSet.Stage != nil {
		for _, vase2ddiagram := range __gong__sortStageSetInstances(stageSet.Stage.Vase2DDiagrams, stageSet.Stage.Vase2DDiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vase2ddiagramIdent := "__models" + vase2ddiagram.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Vase2DDiagram{Name: %s}).Stage(stageSet.Stage)", vase2ddiagramIdent, __gong__toRawStringLiteral(vase2ddiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vase2ddiagramIdent, __gong__toRawStringLiteral(vase2ddiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Zoom = %f", vase2ddiagramIdent, vase2ddiagram.Zoom))
			values.WriteString(fmt.Sprintf("\n\t%s.IsVaseArcNodesExpanded = %t", vase2ddiagramIdent, vase2ddiagram.IsVaseArcNodesExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsVaseClampingNodesExpanded = %t", vase2ddiagramIdent, vase2ddiagram.IsVaseClampingNodesExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenAxesShape = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenAxesShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenBottomStartArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenBottomStartArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenBottomEndArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenBottomEndArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenBottomStackOfGrowthCurve = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenBottomStackOfGrowthCurve))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenShiftedLeftStackOfGrowthCurve = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenShiftedLeftStackOfGrowthCurve))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenShiftedLeftStackOfNormalVector = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenShiftedLeftStackOfNormalVector))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPerpendicularVectorGridHalfway = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenPerpendicularVectorGridHalfway))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopStartArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopStartArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenShiftedBottomTopStartArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenShiftedBottomTopStartArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopMidArcVectorShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopMidArcVectorShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStartHalfwayArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenStartHalfwayArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopStartHalfwayArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopStartHalfwayArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenEndHalfwayArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenEndHalfwayArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopEndHalfwayArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopEndHalfwayArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopEndArcShapeGrid = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopEndArcShapeGrid))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfGrowthCurve = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenStackOfGrowthCurve))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopStackOfGrowthCurve = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopStackOfGrowthCurve))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopGrowthCurve2D = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopGrowthCurve2D))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfGrowthCurve2D = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenStackOfGrowthCurve2D))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenTopStackOfGrowthCurve2D = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenTopStackOfGrowthCurve2D))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenShiftedRightGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenShiftedRightGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenShiftedLeftGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenShiftedLeftGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenStackOfGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenStackOfRotatedGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenStackOfRotatedGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPartiallyGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenPartiallyGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPartiallyGrowthCurve2DTrajectory = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenPartiallyGrowthCurve2DTrajectory))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2 = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenPxShape = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenPxShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenChosenP1P2PairShape = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenChosenP1P2PairShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHiddenKeyHoleShape = %t", vase2ddiagramIdent, vase2ddiagram.IsHiddenKeyHoleShape))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", vase2ddiagramIdent, vase2ddiagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", vase2ddiagramIdent, __gong__toRawStringLiteral(vase2ddiagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", vase2ddiagramIdent, vase2ddiagram.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		for _, vasetrapezeringshape := range __gong__sortStageSetInstances(stageSet.Stage.VaseTrapezeRingShapes, stageSet.Stage.VaseTrapezeRingShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			vasetrapezeringshapeIdent := "__models" + vasetrapezeringshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.VaseTrapezeRingShape{Name: %s}).Stage(stageSet.Stage)", vasetrapezeringshapeIdent, __gong__toRawStringLiteral(vasetrapezeringshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", vasetrapezeringshapeIdent, __gong__toRawStringLiteral(vasetrapezeringshape.Name)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/stool"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/clock"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
	_ *stool.Stage
	_ *music.Stage
	_ *clock.Stage
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
		case "github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/clock":
			aliasToCanonical[alias] = "clock"
		case "github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music":
			aliasToCanonical[alias] = "music"
		case "github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/stool":
			aliasToCanonical[alias] = "stool"
		case "github.com/fullstack-lang/gong/dsm/phylla/go/models":
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
			case "clock":
				switch typeName {
				case "ClockAbstract":
					identifierMap[ident.Name] = __gong__stageSetInit(new(clock.ClockAbstract), stageSet.ClockStage, ident.Name, instanceName, preserveOrder)
				}
			case "music":
				switch typeName {
				case "MusicAbstract":
					identifierMap[ident.Name] = __gong__stageSetInit(new(music.MusicAbstract), stageSet.MusicStage, ident.Name, instanceName, preserveOrder)
				}
			case "stool":
				switch typeName {
				case "StoolAbstract":
					identifierMap[ident.Name] = __gong__stageSetInit(new(stool.StoolAbstract), stageSet.StoolStage, ident.Name, instanceName, preserveOrder)
				}
			case "models":
				switch typeName {
				case "Angle0Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Angle0Shape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "BottomCurvePlane1Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BottomCurvePlane1Shape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "BottomCurvePlane2Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BottomCurvePlane2Shape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Circumference3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Circumference3DShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Clock2DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Clock2DDiagram), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Clock3DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Clock3DDiagram), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "CutLine3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CutLine3DShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Leaves3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Leaves3DShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "OriginalPoints3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OriginalPoints3DShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ParastichyMCurves3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParastichyMCurves3DShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ParastichyNCurves3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParastichyNCurves3DShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Plant2DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Plant2DDiagram), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Plant3DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Plant3DDiagram), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "PlantAbstract":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PlantAbstract), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Rendered3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Rendered3DShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SampledPoints3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SampledPoints3DShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StackOfRotatedVaseTrapezeRingsShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StackOfRotatedVaseTrapezeRingsShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StackOfVaseTrapezeRingsShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StackOfVaseTrapezeRingsShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "StemCylinder3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StemCylinder3DShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Stool2DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Stool2DDiagram), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Stool3DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Stool3DDiagram), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TopCurvePlane1Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TopCurvePlane1Shape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TopCurvePlane2Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TopCurvePlane2Shape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TubeVase3DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TubeVase3DDiagram), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "TubeVaseAbstract":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TubeVaseAbstract), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Vase2DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Vase2DDiagram), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "VaseTrapezeRingShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(VaseTrapezeRingShape), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *clock.ClockAbstract:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RadialRepetitions":
						inst.RadialRepetitions = GongExtractInt(rhs)
					case "Transparency":
						inst.Transparency = GongExtractFloat(rhs)
					case "RelativeTubeDiameter":
						inst.RelativeTubeDiameter = GongExtractFloat(rhs)
					case "RelativeHeight3DTorus":
						inst.RelativeHeight3DTorus = GongExtractFloat(rhs)
					case "ClockTorusVerticalScale":
						inst.ClockTorusVerticalScale = GongExtractFloat(rhs)
					case "RelativeHeight":
						inst.RelativeHeight = GongExtractFloat(rhs)
					case "ProjectionAngle":
						inst.ProjectionAngle = GongExtractFloat(rhs)
					}
				case *music.MusicAbstract:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "PitchHeight":
						inst.PitchHeight = GongExtractFloat(rhs)
					case "NbOfBeatsInTheme":
						inst.NbOfBeatsInTheme = GongExtractInt(rhs)
					case "BeatsPerSecond":
						inst.BeatsPerSecond = GongExtractFloat(rhs)
					case "FirstVoiceShiftX":
						inst.FirstVoiceShiftX = GongExtractFloat(rhs)
					case "FirstVoiceShiftY":
						inst.FirstVoiceShiftY = GongExtractFloat(rhs)
					case "PitchDifference":
						inst.PitchDifference = GongExtractInt(rhs)
					case "Level":
						inst.Level = GongExtractFloat(rhs)
					case "ActualBeatsTemporalShift":
						inst.ActualBeatsTemporalShift = GongExtractInt(rhs)
					case "IsMinor":
						inst.IsMinor = GongExtractBool(rhs)
					case "ThemeBinaryEncoding":
						inst.ThemeBinaryEncoding = GongExtractInt(rhs)
					case "BezierControlLengthRatio":
						inst.BezierControlLengthRatio = GongExtractFloat(rhs)
					case "NbPitchLines":
						inst.NbPitchLines = GongExtractInt(rhs)
					case "NbBeatLines":
						inst.NbBeatLines = GongExtractInt(rhs)
					case "OriginX":
						inst.OriginX = GongExtractFloat(rhs)
					case "OriginY":
						inst.OriginY = GongExtractFloat(rhs)
					case "ScoreScale":
						inst.ScoreScale = GongExtractFloat(rhs)
					case "ShowFirstVoice":
						inst.ShowFirstVoice = GongExtractBool(rhs)
					case "ShowFirstVoiceShiftRight":
						inst.ShowFirstVoiceShiftRight = GongExtractBool(rhs)
					case "ShowSecondVoice":
						inst.ShowSecondVoice = GongExtractBool(rhs)
					case "ShowSecondVoiceShiftRight":
						inst.ShowSecondVoiceShiftRight = GongExtractBool(rhs)
					case "ShowFirstVoiceNotes":
						inst.ShowFirstVoiceNotes = GongExtractBool(rhs)
					case "ShowFirstVoiceNotesShiftRight":
						inst.ShowFirstVoiceNotesShiftRight = GongExtractBool(rhs)
					case "ShowSecondVoiceNotes":
						inst.ShowSecondVoiceNotes = GongExtractBool(rhs)
					case "ShowSecondVoiceNotesShiftRight":
						inst.ShowSecondVoiceNotesShiftRight = GongExtractBool(rhs)
					case "IsComposerNodeExpanded":
						inst.IsComposerNodeExpanded = GongExtractBool(rhs)
					}
				case *stool.StoolAbstract:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RadialRepetitions":
						inst.RadialRepetitions = GongExtractInt(rhs)
					case "Transparency":
						inst.Transparency = GongExtractFloat(rhs)
					case "RelativeTubeDiameter":
						inst.RelativeTubeDiameter = GongExtractFloat(rhs)
					case "RelativeHeight3DTorus":
						inst.RelativeHeight3DTorus = GongExtractFloat(rhs)
					case "StoolTorusVerticalScale":
						inst.StoolTorusVerticalScale = GongExtractFloat(rhs)
					case "RelativeHeight":
						inst.RelativeHeight = GongExtractFloat(rhs)
					case "RelativeSeatThickness":
						inst.RelativeSeatThickness = GongExtractFloat(rhs)
					case "ProjectionAngle":
						inst.ProjectionAngle = GongExtractFloat(rhs)
					case "RelativeEyeSeparationCriteria":
						inst.RelativeEyeSeparationCriteria = GongExtractFloat(rhs)
					case "RelativeEyeCornerControlVectorStrength":
						inst.RelativeEyeCornerControlVectorStrength = GongExtractFloat(rhs)
					}
				case *Angle0Shape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *BottomCurvePlane1Shape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *BottomCurvePlane2Shape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Circumference3DShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Clock2DDiagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Zoom":
						inst.Zoom = GongExtractFloat(rhs)
					case "IsHiddenAxesShape":
						inst.IsHiddenAxesShape = GongExtractBool(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *Clock3DDiagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsHiddenClockTopCurveShape":
						inst.IsHiddenClockTopCurveShape = GongExtractBool(rhs)
					case "IsHiddenTorus3DShape":
						inst.IsHiddenTorus3DShape = GongExtractBool(rhs)
					case "IsHiddenSampledPoints3DShape":
						inst.IsHiddenSampledPoints3DShape = GongExtractBool(rhs)
					case "SampledPoints3DShape":
						__gong__assignPointer(&inst.SampledPoints3DShape, rhs, identifierMap)
					case "IsHiddenTiledFloor3DShape":
						inst.IsHiddenTiledFloor3DShape = GongExtractBool(rhs)
					case "Rendered3DShape":
						__gong__assignPointer(&inst.Rendered3DShape, rhs, identifierMap)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *CutLine3DShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Leaves3DShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Library:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Plants":
						__gong__assignSliceOfPointers(&inst.Plants, rhs, identifierMap)
					case "SubLibraries":
						__gong__assignSliceOfPointers(&inst.SubLibraries, rhs, identifierMap)
					case "NbPixPerCharacter":
						inst.NbPixPerCharacter = GongExtractFloat(rhs)
					case "LogoSVGFile":
						inst.LogoSVGFile = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsRootLibrary":
						inst.IsRootLibrary = GongExtractBool(rhs)
					}
				case *OriginalPoints3DShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *ParastichyMCurves3DShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *ParastichyNCurves3DShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Plant2DDiagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "OriginX":
						inst.OriginX = GongExtractFloat(rhs)
					case "OriginY":
						inst.OriginY = GongExtractFloat(rhs)
					case "Zoom":
						inst.Zoom = GongExtractFloat(rhs)
					case "IsRhombusNodesExpanded":
						inst.IsRhombusNodesExpanded = GongExtractBool(rhs)
					case "IsArcNodesExpanded":
						inst.IsArcNodesExpanded = GongExtractBool(rhs)
					case "IsHiddenAxesShape":
						inst.IsHiddenAxesShape = GongExtractBool(rhs)
					case "IsHiddenReferenceRhombus":
						inst.IsHiddenReferenceRhombus = GongExtractBool(rhs)
					case "IsHiddenPlantCircumferenceShape":
						inst.IsHiddenPlantCircumferenceShape = GongExtractBool(rhs)
					case "IsHiddenGridPathShape":
						inst.IsHiddenGridPathShape = GongExtractBool(rhs)
					case "IsHiddenRhombusGridShape":
						inst.IsHiddenRhombusGridShape = GongExtractBool(rhs)
					case "IsHiddenExplanationTextShape":
						inst.IsHiddenExplanationTextShape = GongExtractBool(rhs)
					case "IsHiddenRotatedReferenceRhombus":
						inst.IsHiddenRotatedReferenceRhombus = GongExtractBool(rhs)
					case "IsHiddenRotatedPlantCircumferenceShape":
						inst.IsHiddenRotatedPlantCircumferenceShape = GongExtractBool(rhs)
					case "IsHiddenRotatedGridPathShape":
						inst.IsHiddenRotatedGridPathShape = GongExtractBool(rhs)
					case "IsHiddenRotatedRhombusGridShape":
						inst.IsHiddenRotatedRhombusGridShape = GongExtractBool(rhs)
					case "IsHiddenGrowthPathRhombusGridShape":
						inst.IsHiddenGrowthPathRhombusGridShape = GongExtractBool(rhs)
					case "IsHiddenGrowthVectorShape":
						inst.IsHiddenGrowthVectorShape = GongExtractBool(rhs)
					case "IsHiddenPerpendicularVectorGrid":
						inst.IsHiddenPerpendicularVectorGrid = GongExtractBool(rhs)
					case "IsHiddenBaseVectorShapeGrid":
						inst.IsHiddenBaseVectorShapeGrid = GongExtractBool(rhs)
					case "IsHiddenArcNormalVectorShapeGrid":
						inst.IsHiddenArcNormalVectorShapeGrid = GongExtractBool(rhs)
					case "IsHiddenStartArcShapeGrid":
						inst.IsHiddenStartArcShapeGrid = GongExtractBool(rhs)
					case "IsHiddenMidArcVectorShapeGrid":
						inst.IsHiddenMidArcVectorShapeGrid = GongExtractBool(rhs)
					case "IsHiddenEndArcShapeGrid":
						inst.IsHiddenEndArcShapeGrid = GongExtractBool(rhs)
					case "IsHiddenGrowthCurve2D":
						inst.IsHiddenGrowthCurve2D = GongExtractBool(rhs)
					case "IsHiddenStackOfGrowthCurve2DByGrowthVector":
						inst.IsHiddenStackOfGrowthCurve2DByGrowthVector = GongExtractBool(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *Plant3DDiagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsHiddenStemCylinder3DShape":
						inst.IsHiddenStemCylinder3DShape = GongExtractBool(rhs)
					case "StemCylinder3DShape":
						__gong__assignPointer(&inst.StemCylinder3DShape, rhs, identifierMap)
					case "IsHiddenParastichyNCurves3DShape":
						inst.IsHiddenParastichyNCurves3DShape = GongExtractBool(rhs)
					case "ParastichyNCurves3DShape":
						__gong__assignPointer(&inst.ParastichyNCurves3DShape, rhs, identifierMap)
					case "IsHiddenParastichyMCurves3DShape":
						inst.IsHiddenParastichyMCurves3DShape = GongExtractBool(rhs)
					case "ParastichyMCurves3DShape":
						__gong__assignPointer(&inst.ParastichyMCurves3DShape, rhs, identifierMap)
					case "IsHiddenCutLine3DShape":
						inst.IsHiddenCutLine3DShape = GongExtractBool(rhs)
					case "CutLine3DShape":
						__gong__assignPointer(&inst.CutLine3DShape, rhs, identifierMap)
					case "IsHiddenCircumference3DShape":
						inst.IsHiddenCircumference3DShape = GongExtractBool(rhs)
					case "Circumference3DShape":
						__gong__assignPointer(&inst.Circumference3DShape, rhs, identifierMap)
					case "IsHiddenTiledFloor3DShape":
						inst.IsHiddenTiledFloor3DShape = GongExtractBool(rhs)
					case "IsHiddenLeaves3DShape":
						inst.IsHiddenLeaves3DShape = GongExtractBool(rhs)
					case "Leaves3DShape":
						__gong__assignPointer(&inst.Leaves3DShape, rhs, identifierMap)
					case "Rendered3DShape":
						__gong__assignPointer(&inst.Rendered3DShape, rhs, identifierMap)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *PlantAbstract:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "N":
						inst.N = GongExtractInt(rhs)
					case "M":
						inst.M = GongExtractInt(rhs)
					case "StackHeight":
						inst.StackHeight = GongExtractInt(rhs)
					case "RhombusInsideAngle":
						inst.RhombusInsideAngle = GongExtractFloat(rhs)
					case "RhombusSideLength":
						inst.RhombusSideLength = GongExtractFloat(rhs)
					case "PlantType":
						inst.PlantType = PlantType(GongExtractString(rhs))
					case "TubeVaseAbstract":
						__gong__assignPointer(&inst.TubeVaseAbstract, rhs, identifierMap)
					case "StoolAbstract":
						__gong__assignPointer(&inst.StoolAbstract, rhs, identifierMap)
					case "ClockAbstract":
						__gong__assignPointer(&inst.ClockAbstract, rhs, identifierMap)
					case "MusicAbstract":
						__gong__assignPointer(&inst.MusicAbstract, rhs, identifierMap)
					case "CurrentView":
						inst.CurrentView = ViewType(GongExtractString(rhs))
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsSelected":
						inst.IsSelected = GongExtractBool(rhs)
					case "IsPlant2DDiagramsNodeExpanded":
						inst.IsPlant2DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "Plant2DDiagrams":
						__gong__assignSliceOfPointers(&inst.Plant2DDiagrams, rhs, identifierMap)
					case "IsPlant3DDiagramsNodeExpanded":
						inst.IsPlant3DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "Plant3DDiagrams":
						__gong__assignSliceOfPointers(&inst.Plant3DDiagrams, rhs, identifierMap)
					case "IsVase2DDiagramsNodeExpanded":
						inst.IsVase2DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "Vase2DDiagrams":
						__gong__assignSliceOfPointers(&inst.Vase2DDiagrams, rhs, identifierMap)
					case "IsTubeVase3DDiagramsNodeExpanded":
						inst.IsTubeVase3DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "TubeVase3DDiagrams":
						__gong__assignSliceOfPointers(&inst.TubeVase3DDiagrams, rhs, identifierMap)
					case "IsStool2DDiagramsNodeExpanded":
						inst.IsStool2DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "Stool2DDiagrams":
						__gong__assignSliceOfPointers(&inst.Stool2DDiagrams, rhs, identifierMap)
					case "IsStool3DDiagramsNodeExpanded":
						inst.IsStool3DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "Stool3DDiagrams":
						__gong__assignSliceOfPointers(&inst.Stool3DDiagrams, rhs, identifierMap)
					case "IsClock2DDiagramsNodeExpanded":
						inst.IsClock2DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "Clock2DDiagrams":
						__gong__assignSliceOfPointers(&inst.Clock2DDiagrams, rhs, identifierMap)
					case "IsClock3DDiagramsNodeExpanded":
						inst.IsClock3DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "Clock3DDiagrams":
						__gong__assignSliceOfPointers(&inst.Clock3DDiagrams, rhs, identifierMap)
					}
				case *Rendered3DShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ViewX":
						inst.ViewX = GongExtractFloat(rhs)
					case "ViewY":
						inst.ViewY = GongExtractFloat(rhs)
					case "ViewZ":
						inst.ViewZ = GongExtractFloat(rhs)
					case "TargetX":
						inst.TargetX = GongExtractFloat(rhs)
					case "TargetY":
						inst.TargetY = GongExtractFloat(rhs)
					case "TargetZ":
						inst.TargetZ = GongExtractFloat(rhs)
					case "Fov":
						inst.Fov = GongExtractFloat(rhs)
					}
				case *SampledPoints3DShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *StackOfRotatedVaseTrapezeRingsShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *StackOfVaseTrapezeRingsShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *StemCylinder3DShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Transparency":
						inst.Transparency = GongExtractFloat(rhs)
					}
				case *Stool2DDiagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Zoom":
						inst.Zoom = GongExtractFloat(rhs)
					case "IsHiddenAxesShape":
						inst.IsHiddenAxesShape = GongExtractBool(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *Stool3DDiagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsHiddenSeatTopCurveShape":
						inst.IsHiddenSeatTopCurveShape = GongExtractBool(rhs)
					case "IsHiddenRotatedSeatTopCurveShape":
						inst.IsHiddenRotatedSeatTopCurveShape = GongExtractBool(rhs)
					case "IsHiddenSeatBottomCurveShape":
						inst.IsHiddenSeatBottomCurveShape = GongExtractBool(rhs)
					case "IsHiddenRotatedSeatBottomCurveShape":
						inst.IsHiddenRotatedSeatBottomCurveShape = GongExtractBool(rhs)
					case "IsHiddenTorus3DShape":
						inst.IsHiddenTorus3DShape = GongExtractBool(rhs)
					case "IsHiddenRotatedTorusShape":
						inst.IsHiddenRotatedTorusShape = GongExtractBool(rhs)
					case "IsHiddenSampledPoints3DShape":
						inst.IsHiddenSampledPoints3DShape = GongExtractBool(rhs)
					case "SampledPoints3DShape":
						__gong__assignPointer(&inst.SampledPoints3DShape, rhs, identifierMap)
					case "IsHiddenRotatedSampledPoints3DShape":
						inst.IsHiddenRotatedSampledPoints3DShape = GongExtractBool(rhs)
					case "IsHiddenEyeSampledPoints3DShape":
						inst.IsHiddenEyeSampledPoints3DShape = GongExtractBool(rhs)
					case "IsHiddenEyeCornersSampledPoints3DShape":
						inst.IsHiddenEyeCornersSampledPoints3DShape = GongExtractBool(rhs)
					case "IsHiddenEye3DShape":
						inst.IsHiddenEye3DShape = GongExtractBool(rhs)
					case "IsHiddenEyeSeatBottomCurveShape":
						inst.IsHiddenEyeSeatBottomCurveShape = GongExtractBool(rhs)
					case "IsHiddenEyeStoolBottomCurveShape":
						inst.IsHiddenEyeStoolBottomCurveShape = GongExtractBool(rhs)
					case "IsHiddenSeat3DShape":
						inst.IsHiddenSeat3DShape = GongExtractBool(rhs)
					case "IsHiddenEyeVolume3DShape":
						inst.IsHiddenEyeVolume3DShape = GongExtractBool(rhs)
					case "IsHiddenSeatAndLegs3DShape":
						inst.IsHiddenSeatAndLegs3DShape = GongExtractBool(rhs)
					case "IsHiddenRotatedSeatAndLegs3DShape":
						inst.IsHiddenRotatedSeatAndLegs3DShape = GongExtractBool(rhs)
					case "IsHiddenTiledFloor3DShape":
						inst.IsHiddenTiledFloor3DShape = GongExtractBool(rhs)
					case "Rendered3DShape":
						__gong__assignPointer(&inst.Rendered3DShape, rhs, identifierMap)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *TopCurvePlane1Shape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *TopCurvePlane2Shape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *TubeVase3DDiagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon":
						inst.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon = GongExtractBool(rhs)
					case "IsHiddenTorusStackShape":
						inst.IsHiddenTorusStackShape = GongExtractBool(rhs)
					case "IsHiddenVerticalTorusStackShape":
						inst.IsHiddenVerticalTorusStackShape = GongExtractBool(rhs)
					case "IsHiddenPartiallyRotatedTorusShape":
						inst.IsHiddenPartiallyRotatedTorusShape = GongExtractBool(rhs)
					case "IsHiddenStackOfPartiallyRotatedTorusShape":
						inst.IsHiddenStackOfPartiallyRotatedTorusShape = GongExtractBool(rhs)
					case "IsHiddenPointsAndLines3DShape":
						inst.IsHiddenPointsAndLines3DShape = GongExtractBool(rhs)
					case "IsHiddenKeyHole3DShape":
						inst.IsHiddenKeyHole3DShape = GongExtractBool(rhs)
					case "IsHiddenKey3DShape":
						inst.IsHiddenKey3DShape = GongExtractBool(rhs)
					case "IsHiddenVolumeKey3DShape":
						inst.IsHiddenVolumeKey3DShape = GongExtractBool(rhs)
					case "IsHiddenTorusEdge3DShape":
						inst.IsHiddenTorusEdge3DShape = GongExtractBool(rhs)
					case "IsHiddenSampledPoints3DShape":
						inst.IsHiddenSampledPoints3DShape = GongExtractBool(rhs)
					case "IsHiddenOriginalPoints3DShape":
						inst.IsHiddenOriginalPoints3DShape = GongExtractBool(rhs)
					case "IsHiddenAngle0Shape":
						inst.IsHiddenAngle0Shape = GongExtractBool(rhs)
					case "IsHiddenTiledFloor3DShape":
						inst.IsHiddenTiledFloor3DShape = GongExtractBool(rhs)
					case "IsHiddenTopCurvePlane1Shape":
						inst.IsHiddenTopCurvePlane1Shape = GongExtractBool(rhs)
					case "IsHiddenBottomCurvePlane1Shape":
						inst.IsHiddenBottomCurvePlane1Shape = GongExtractBool(rhs)
					case "IsHiddenTopCurvePlane2Shape":
						inst.IsHiddenTopCurvePlane2Shape = GongExtractBool(rhs)
					case "IsHiddenBottomCurvePlane2Shape":
						inst.IsHiddenBottomCurvePlane2Shape = GongExtractBool(rhs)
					case "IsHiddenVaseTrapezeRingShape":
						inst.IsHiddenVaseTrapezeRingShape = GongExtractBool(rhs)
					case "IsHiddenStackOfVaseTrapezeRingsShape":
						inst.IsHiddenStackOfVaseTrapezeRingsShape = GongExtractBool(rhs)
					case "IsHiddenStackOfRotatedVaseTrapezeRingsShape":
						inst.IsHiddenStackOfRotatedVaseTrapezeRingsShape = GongExtractBool(rhs)
					case "Rendered3DShape":
						__gong__assignPointer(&inst.Rendered3DShape, rhs, identifierMap)
					case "SampledPoints3DShape":
						__gong__assignPointer(&inst.SampledPoints3DShape, rhs, identifierMap)
					case "OriginalPoints3DShape":
						__gong__assignPointer(&inst.OriginalPoints3DShape, rhs, identifierMap)
					case "Angle0Shape":
						__gong__assignPointer(&inst.Angle0Shape, rhs, identifierMap)
					case "TopCurvePlane1Shape":
						__gong__assignPointer(&inst.TopCurvePlane1Shape, rhs, identifierMap)
					case "BottomCurvePlane1Shape":
						__gong__assignPointer(&inst.BottomCurvePlane1Shape, rhs, identifierMap)
					case "TopCurvePlane2Shape":
						__gong__assignPointer(&inst.TopCurvePlane2Shape, rhs, identifierMap)
					case "BottomCurvePlane2Shape":
						__gong__assignPointer(&inst.BottomCurvePlane2Shape, rhs, identifierMap)
					case "VaseTrapezeRingShape":
						__gong__assignPointer(&inst.VaseTrapezeRingShape, rhs, identifierMap)
					case "StackOfVaseTrapezeRingsShape":
						__gong__assignPointer(&inst.StackOfVaseTrapezeRingsShape, rhs, identifierMap)
					case "StackOfRotatedVaseTrapezeRingsShape":
						__gong__assignPointer(&inst.StackOfRotatedVaseTrapezeRingsShape, rhs, identifierMap)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *TubeVaseAbstract:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Z_Ribbon":
						inst.Z_Ribbon = GongExtractFloat(rhs)
					case "RibbonVerticalScale":
						inst.RibbonVerticalScale = GongExtractFloat(rhs)
					case "Plane1Height":
						inst.Plane1Height = GongExtractFloat(rhs)
					case "Plane2Height":
						inst.Plane2Height = GongExtractFloat(rhs)
					case "ProjectionAngle":
						inst.ProjectionAngle = GongExtractFloat(rhs)
					case "RelativeVerticalThickness":
						inst.RelativeVerticalThickness = GongExtractFloat(rhs)
					case "RelativeRadialThickness":
						inst.RelativeRadialThickness = GongExtractFloat(rhs)
					case "RelativeCuttedStackFloorHeight":
						inst.RelativeCuttedStackFloorHeight = GongExtractFloat(rhs)
					case "RelativeRotatedTorusSeparation":
						inst.RelativeRotatedTorusSeparation = GongExtractFloat(rhs)
					case "RotationRatio":
						inst.RotationRatio = GongExtractFloat(rhs)
					case "RadialRepetitions":
						inst.RadialRepetitions = GongExtractInt(rhs)
					case "Transparency":
						inst.Transparency = GongExtractFloat(rhs)
					case "HasAlternatingRingColors":
						inst.HasAlternatingRingColors = GongExtractBool(rhs)
					case "RelativeTrajectoryOffsetX":
						inst.RelativeTrajectoryOffsetX = GongExtractFloat(rhs)
					case "RelativeTrajectoryOffsetY":
						inst.RelativeTrajectoryOffsetY = GongExtractFloat(rhs)
					case "NbStepP1P2":
						inst.NbStepP1P2 = GongExtractInt(rhs)
					case "ChosenStep":
						inst.ChosenStep = GongExtractInt(rhs)
					case "RelativeHorizontalRingsHeight":
						inst.RelativeHorizontalRingsHeight = GongExtractFloat(rhs)
					case "OffsetKeyX":
						inst.OffsetKeyX = GongExtractFloat(rhs)
					case "OffsetKeyY":
						inst.OffsetKeyY = GongExtractFloat(rhs)
					case "HeightKey":
						inst.HeightKey = GongExtractFloat(rhs)
					case "WidthKey":
						inst.WidthKey = GongExtractFloat(rhs)
					case "RelativeKeySize":
						inst.RelativeKeySize = GongExtractFloat(rhs)
					case "MovieNbFrames":
						inst.MovieNbFrames = GongExtractInt(rhs)
					}
				case *Vase2DDiagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Zoom":
						inst.Zoom = GongExtractFloat(rhs)
					case "IsVaseArcNodesExpanded":
						inst.IsVaseArcNodesExpanded = GongExtractBool(rhs)
					case "IsVaseClampingNodesExpanded":
						inst.IsVaseClampingNodesExpanded = GongExtractBool(rhs)
					case "IsHiddenAxesShape":
						inst.IsHiddenAxesShape = GongExtractBool(rhs)
					case "IsHiddenBottomStartArcShapeGrid":
						inst.IsHiddenBottomStartArcShapeGrid = GongExtractBool(rhs)
					case "IsHiddenBottomEndArcShapeGrid":
						inst.IsHiddenBottomEndArcShapeGrid = GongExtractBool(rhs)
					case "IsHiddenBottomStackOfGrowthCurve":
						inst.IsHiddenBottomStackOfGrowthCurve = GongExtractBool(rhs)
					case "IsHiddenShiftedLeftStackOfGrowthCurve":
						inst.IsHiddenShiftedLeftStackOfGrowthCurve = GongExtractBool(rhs)
					case "IsHiddenShiftedLeftStackOfNormalVector":
						inst.IsHiddenShiftedLeftStackOfNormalVector = GongExtractBool(rhs)
					case "IsHiddenPerpendicularVectorGridHalfway":
						inst.IsHiddenPerpendicularVectorGridHalfway = GongExtractBool(rhs)
					case "IsHiddenTopStartArcShapeGrid":
						inst.IsHiddenTopStartArcShapeGrid = GongExtractBool(rhs)
					case "IsHiddenShiftedBottomTopStartArcShapeGrid":
						inst.IsHiddenShiftedBottomTopStartArcShapeGrid = GongExtractBool(rhs)
					case "IsHiddenTopMidArcVectorShapeGrid":
						inst.IsHiddenTopMidArcVectorShapeGrid = GongExtractBool(rhs)
					case "IsHiddenStartHalfwayArcShapeGrid":
						inst.IsHiddenStartHalfwayArcShapeGrid = GongExtractBool(rhs)
					case "IsHiddenTopStartHalfwayArcShapeGrid":
						inst.IsHiddenTopStartHalfwayArcShapeGrid = GongExtractBool(rhs)
					case "IsHiddenEndHalfwayArcShapeGrid":
						inst.IsHiddenEndHalfwayArcShapeGrid = GongExtractBool(rhs)
					case "IsHiddenTopEndHalfwayArcShapeGrid":
						inst.IsHiddenTopEndHalfwayArcShapeGrid = GongExtractBool(rhs)
					case "IsHiddenTopEndArcShapeGrid":
						inst.IsHiddenTopEndArcShapeGrid = GongExtractBool(rhs)
					case "IsHiddenStackOfGrowthCurve":
						inst.IsHiddenStackOfGrowthCurve = GongExtractBool(rhs)
					case "IsHiddenTopStackOfGrowthCurve":
						inst.IsHiddenTopStackOfGrowthCurve = GongExtractBool(rhs)
					case "IsHiddenTopGrowthCurve2D":
						inst.IsHiddenTopGrowthCurve2D = GongExtractBool(rhs)
					case "IsHiddenStackOfGrowthCurve2D":
						inst.IsHiddenStackOfGrowthCurve2D = GongExtractBool(rhs)
					case "IsHiddenTopStackOfGrowthCurve2D":
						inst.IsHiddenTopStackOfGrowthCurve2D = GongExtractBool(rhs)
					case "IsHiddenGrowthCurve2DRibbon":
						inst.IsHiddenGrowthCurve2DRibbon = GongExtractBool(rhs)
					case "IsHiddenShiftedRightGrowthCurve2DRibbon":
						inst.IsHiddenShiftedRightGrowthCurve2DRibbon = GongExtractBool(rhs)
					case "IsHiddenShiftedLeftGrowthCurve2DRibbon":
						inst.IsHiddenShiftedLeftGrowthCurve2DRibbon = GongExtractBool(rhs)
					case "IsHiddenStackOfGrowthCurve2DRibbon":
						inst.IsHiddenStackOfGrowthCurve2DRibbon = GongExtractBool(rhs)
					case "IsHiddenStackOfRotatedGrowthCurve2DRibbon":
						inst.IsHiddenStackOfRotatedGrowthCurve2DRibbon = GongExtractBool(rhs)
					case "IsHiddenPartiallyGrowthCurve2DRibbon":
						inst.IsHiddenPartiallyGrowthCurve2DRibbon = GongExtractBool(rhs)
					case "IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon":
						inst.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon = GongExtractBool(rhs)
					case "IsHiddenPartiallyGrowthCurve2DTrajectory":
						inst.IsHiddenPartiallyGrowthCurve2DTrajectory = GongExtractBool(rhs)
					case "IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2":
						inst.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2 = GongExtractBool(rhs)
					case "IsHiddenPxShape":
						inst.IsHiddenPxShape = GongExtractBool(rhs)
					case "IsHiddenChosenP1P2PairShape":
						inst.IsHiddenChosenP1P2PairShape = GongExtractBool(rhs)
					case "IsHiddenKeyHoleShape":
						inst.IsHiddenKeyHoleShape = GongExtractBool(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *VaseTrapezeRingShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
