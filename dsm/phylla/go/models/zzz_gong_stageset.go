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
		angle0shapeOrdered := []*Angle0Shape{}
		for angle0shape := range stageSet.Stage.Angle0Shapes {
			angle0shapeOrdered = append(angle0shapeOrdered, angle0shape)
		}
		sort.Slice(angle0shapeOrdered, func(i, j int) bool {
			return stageSet.Stage.Angle0Shape_stagedOrder[angle0shapeOrdered[i]] < stageSet.Stage.Angle0Shape_stagedOrder[angle0shapeOrdered[j]]
		})
		for _, angle0shape := range angle0shapeOrdered {
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
		bottomcurveplane1shapeOrdered := []*BottomCurvePlane1Shape{}
		for bottomcurveplane1shape := range stageSet.Stage.BottomCurvePlane1Shapes {
			bottomcurveplane1shapeOrdered = append(bottomcurveplane1shapeOrdered, bottomcurveplane1shape)
		}
		sort.Slice(bottomcurveplane1shapeOrdered, func(i, j int) bool {
			return stageSet.Stage.BottomCurvePlane1Shape_stagedOrder[bottomcurveplane1shapeOrdered[i]] < stageSet.Stage.BottomCurvePlane1Shape_stagedOrder[bottomcurveplane1shapeOrdered[j]]
		})
		for _, bottomcurveplane1shape := range bottomcurveplane1shapeOrdered {
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
		bottomcurveplane2shapeOrdered := []*BottomCurvePlane2Shape{}
		for bottomcurveplane2shape := range stageSet.Stage.BottomCurvePlane2Shapes {
			bottomcurveplane2shapeOrdered = append(bottomcurveplane2shapeOrdered, bottomcurveplane2shape)
		}
		sort.Slice(bottomcurveplane2shapeOrdered, func(i, j int) bool {
			return stageSet.Stage.BottomCurvePlane2Shape_stagedOrder[bottomcurveplane2shapeOrdered[i]] < stageSet.Stage.BottomCurvePlane2Shape_stagedOrder[bottomcurveplane2shapeOrdered[j]]
		})
		for _, bottomcurveplane2shape := range bottomcurveplane2shapeOrdered {
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
		circumference3dshapeOrdered := []*Circumference3DShape{}
		for circumference3dshape := range stageSet.Stage.Circumference3DShapes {
			circumference3dshapeOrdered = append(circumference3dshapeOrdered, circumference3dshape)
		}
		sort.Slice(circumference3dshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.Circumference3DShape_stagedOrder[circumference3dshapeOrdered[i]] < stageSet.Stage.Circumference3DShape_stagedOrder[circumference3dshapeOrdered[j]]
		})
		for _, circumference3dshape := range circumference3dshapeOrdered {
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
		clock2ddiagramOrdered := []*Clock2DDiagram{}
		for clock2ddiagram := range stageSet.Stage.Clock2DDiagrams {
			clock2ddiagramOrdered = append(clock2ddiagramOrdered, clock2ddiagram)
		}
		sort.Slice(clock2ddiagramOrdered, func(i, j int) bool {
			return stageSet.Stage.Clock2DDiagram_stagedOrder[clock2ddiagramOrdered[i]] < stageSet.Stage.Clock2DDiagram_stagedOrder[clock2ddiagramOrdered[j]]
		})
		for _, clock2ddiagram := range clock2ddiagramOrdered {
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
		clock3ddiagramOrdered := []*Clock3DDiagram{}
		for clock3ddiagram := range stageSet.Stage.Clock3DDiagrams {
			clock3ddiagramOrdered = append(clock3ddiagramOrdered, clock3ddiagram)
		}
		sort.Slice(clock3ddiagramOrdered, func(i, j int) bool {
			return stageSet.Stage.Clock3DDiagram_stagedOrder[clock3ddiagramOrdered[i]] < stageSet.Stage.Clock3DDiagram_stagedOrder[clock3ddiagramOrdered[j]]
		})
		for _, clock3ddiagram := range clock3ddiagramOrdered {
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
			if clock3ddiagram.ClockTopCurveShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + clock3ddiagram.ClockTopCurveShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClockTopCurveShape = %s", clock3ddiagramIdent, targetIdent))
			}
			if clock3ddiagram.Torus3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + clock3ddiagram.Torus3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Torus3DShape = %s", clock3ddiagramIdent, targetIdent))
			}
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
			if clock3ddiagram.TiledFloor3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + clock3ddiagram.TiledFloor3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TiledFloor3DShape = %s", clock3ddiagramIdent, targetIdent))
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
		clockabstractOrdered := []*ClockAbstract{}
		for clockabstract := range stageSet.Stage.ClockAbstracts {
			clockabstractOrdered = append(clockabstractOrdered, clockabstract)
		}
		sort.Slice(clockabstractOrdered, func(i, j int) bool {
			return stageSet.Stage.ClockAbstract_stagedOrder[clockabstractOrdered[i]] < stageSet.Stage.ClockAbstract_stagedOrder[clockabstractOrdered[j]]
		})
		for _, clockabstract := range clockabstractOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			clockabstractIdent := "__models" + clockabstract.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ClockAbstract{Name: %s}).Stage(stageSet.Stage)", clockabstractIdent, __gong__toRawStringLiteral(clockabstract.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
	if stageSet.Stage != nil {
		cutline3dshapeOrdered := []*CutLine3DShape{}
		for cutline3dshape := range stageSet.Stage.CutLine3DShapes {
			cutline3dshapeOrdered = append(cutline3dshapeOrdered, cutline3dshape)
		}
		sort.Slice(cutline3dshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.CutLine3DShape_stagedOrder[cutline3dshapeOrdered[i]] < stageSet.Stage.CutLine3DShape_stagedOrder[cutline3dshapeOrdered[j]]
		})
		for _, cutline3dshape := range cutline3dshapeOrdered {
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
		leaves3dshapeOrdered := []*Leaves3DShape{}
		for leaves3dshape := range stageSet.Stage.Leaves3DShapes {
			leaves3dshapeOrdered = append(leaves3dshapeOrdered, leaves3dshape)
		}
		sort.Slice(leaves3dshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.Leaves3DShape_stagedOrder[leaves3dshapeOrdered[i]] < stageSet.Stage.Leaves3DShape_stagedOrder[leaves3dshapeOrdered[j]]
		})
		for _, leaves3dshape := range leaves3dshapeOrdered {
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
		libraryOrdered := []*Library{}
		for library := range stageSet.Stage.Librarys {
			libraryOrdered = append(libraryOrdered, library)
		}
		sort.Slice(libraryOrdered, func(i, j int) bool {
			return stageSet.Stage.Library_stagedOrder[libraryOrdered[i]] < stageSet.Stage.Library_stagedOrder[libraryOrdered[j]]
		})
		for _, library := range libraryOrdered {
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
		musicabstractOrdered := []*MusicAbstract{}
		for musicabstract := range stageSet.Stage.MusicAbstracts {
			musicabstractOrdered = append(musicabstractOrdered, musicabstract)
		}
		sort.Slice(musicabstractOrdered, func(i, j int) bool {
			return stageSet.Stage.MusicAbstract_stagedOrder[musicabstractOrdered[i]] < stageSet.Stage.MusicAbstract_stagedOrder[musicabstractOrdered[j]]
		})
		for _, musicabstract := range musicabstractOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			musicabstractIdent := "__models" + musicabstract.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MusicAbstract{Name: %s}).Stage(stageSet.Stage)", musicabstractIdent, __gong__toRawStringLiteral(musicabstract.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
	if stageSet.Stage != nil {
		originalpoints3dshapeOrdered := []*OriginalPoints3DShape{}
		for originalpoints3dshape := range stageSet.Stage.OriginalPoints3DShapes {
			originalpoints3dshapeOrdered = append(originalpoints3dshapeOrdered, originalpoints3dshape)
		}
		sort.Slice(originalpoints3dshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.OriginalPoints3DShape_stagedOrder[originalpoints3dshapeOrdered[i]] < stageSet.Stage.OriginalPoints3DShape_stagedOrder[originalpoints3dshapeOrdered[j]]
		})
		for _, originalpoints3dshape := range originalpoints3dshapeOrdered {
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
		parastichymcurves3dshapeOrdered := []*ParastichyMCurves3DShape{}
		for parastichymcurves3dshape := range stageSet.Stage.ParastichyMCurves3DShapes {
			parastichymcurves3dshapeOrdered = append(parastichymcurves3dshapeOrdered, parastichymcurves3dshape)
		}
		sort.Slice(parastichymcurves3dshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ParastichyMCurves3DShape_stagedOrder[parastichymcurves3dshapeOrdered[i]] < stageSet.Stage.ParastichyMCurves3DShape_stagedOrder[parastichymcurves3dshapeOrdered[j]]
		})
		for _, parastichymcurves3dshape := range parastichymcurves3dshapeOrdered {
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
		parastichyncurves3dshapeOrdered := []*ParastichyNCurves3DShape{}
		for parastichyncurves3dshape := range stageSet.Stage.ParastichyNCurves3DShapes {
			parastichyncurves3dshapeOrdered = append(parastichyncurves3dshapeOrdered, parastichyncurves3dshape)
		}
		sort.Slice(parastichyncurves3dshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ParastichyNCurves3DShape_stagedOrder[parastichyncurves3dshapeOrdered[i]] < stageSet.Stage.ParastichyNCurves3DShape_stagedOrder[parastichyncurves3dshapeOrdered[j]]
		})
		for _, parastichyncurves3dshape := range parastichyncurves3dshapeOrdered {
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
		plant2ddiagramOrdered := []*Plant2DDiagram{}
		for plant2ddiagram := range stageSet.Stage.Plant2DDiagrams {
			plant2ddiagramOrdered = append(plant2ddiagramOrdered, plant2ddiagram)
		}
		sort.Slice(plant2ddiagramOrdered, func(i, j int) bool {
			return stageSet.Stage.Plant2DDiagram_stagedOrder[plant2ddiagramOrdered[i]] < stageSet.Stage.Plant2DDiagram_stagedOrder[plant2ddiagramOrdered[j]]
		})
		for _, plant2ddiagram := range plant2ddiagramOrdered {
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
		plant3ddiagramOrdered := []*Plant3DDiagram{}
		for plant3ddiagram := range stageSet.Stage.Plant3DDiagrams {
			plant3ddiagramOrdered = append(plant3ddiagramOrdered, plant3ddiagram)
		}
		sort.Slice(plant3ddiagramOrdered, func(i, j int) bool {
			return stageSet.Stage.Plant3DDiagram_stagedOrder[plant3ddiagramOrdered[i]] < stageSet.Stage.Plant3DDiagram_stagedOrder[plant3ddiagramOrdered[j]]
		})
		for _, plant3ddiagram := range plant3ddiagramOrdered {
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
			if plant3ddiagram.TiledFloor3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plant3ddiagram.TiledFloor3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TiledFloor3DShape = %s", plant3ddiagramIdent, targetIdent))
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
		plantabstractOrdered := []*PlantAbstract{}
		for plantabstract := range stageSet.Stage.PlantAbstracts {
			plantabstractOrdered = append(plantabstractOrdered, plantabstract)
		}
		sort.Slice(plantabstractOrdered, func(i, j int) bool {
			return stageSet.Stage.PlantAbstract_stagedOrder[plantabstractOrdered[i]] < stageSet.Stage.PlantAbstract_stagedOrder[plantabstractOrdered[j]]
		})
		for _, plantabstract := range plantabstractOrdered {
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
				targetIdent := "__models" + plantabstract.StoolAbstract.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StoolAbstract = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.ClockAbstract != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.ClockAbstract.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClockAbstract = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.MusicAbstract != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.MusicAbstract.GongGetIdentifier(stageSet.Stage)
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
			if plantabstract.AxesShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.AxesShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AxesShape = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.RhombusStuff != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.RhombusStuff.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RhombusStuff = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.GrowthVectorShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.GrowthVectorShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GrowthVectorShape = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.PerpendicularVectorGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.PerpendicularVectorGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PerpendicularVectorGrid = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.BaseVectorShapeGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.BaseVectorShapeGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BaseVectorShapeGrid = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.ArcNormalVectorShapeGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.ArcNormalVectorShapeGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArcNormalVectorShapeGrid = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.StartArcShapeGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.StartArcShapeGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartArcShapeGrid = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.MidArcVectorShapeGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.MidArcVectorShapeGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MidArcVectorShapeGrid = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.EndArcShapeGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.EndArcShapeGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndArcShapeGrid = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.GrowthCurve2D != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.GrowthCurve2D.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GrowthCurve2D = %s", plantabstractIdent, targetIdent))
			}
			if plantabstract.StackOfGrowthCurve2DByGrowthVector != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + plantabstract.StackOfGrowthCurve2DByGrowthVector.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StackOfGrowthCurve2DByGrowthVector = %s", plantabstractIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		rendered3dshapeOrdered := []*Rendered3DShape{}
		for rendered3dshape := range stageSet.Stage.Rendered3DShapes {
			rendered3dshapeOrdered = append(rendered3dshapeOrdered, rendered3dshape)
		}
		sort.Slice(rendered3dshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.Rendered3DShape_stagedOrder[rendered3dshapeOrdered[i]] < stageSet.Stage.Rendered3DShape_stagedOrder[rendered3dshapeOrdered[j]]
		})
		for _, rendered3dshape := range rendered3dshapeOrdered {
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
		sampledpoints3dshapeOrdered := []*SampledPoints3DShape{}
		for sampledpoints3dshape := range stageSet.Stage.SampledPoints3DShapes {
			sampledpoints3dshapeOrdered = append(sampledpoints3dshapeOrdered, sampledpoints3dshape)
		}
		sort.Slice(sampledpoints3dshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.SampledPoints3DShape_stagedOrder[sampledpoints3dshapeOrdered[i]] < stageSet.Stage.SampledPoints3DShape_stagedOrder[sampledpoints3dshapeOrdered[j]]
		})
		for _, sampledpoints3dshape := range sampledpoints3dshapeOrdered {
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
		stemcylinder3dshapeOrdered := []*StemCylinder3DShape{}
		for stemcylinder3dshape := range stageSet.Stage.StemCylinder3DShapes {
			stemcylinder3dshapeOrdered = append(stemcylinder3dshapeOrdered, stemcylinder3dshape)
		}
		sort.Slice(stemcylinder3dshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.StemCylinder3DShape_stagedOrder[stemcylinder3dshapeOrdered[i]] < stageSet.Stage.StemCylinder3DShape_stagedOrder[stemcylinder3dshapeOrdered[j]]
		})
		for _, stemcylinder3dshape := range stemcylinder3dshapeOrdered {
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
		stool2ddiagramOrdered := []*Stool2DDiagram{}
		for stool2ddiagram := range stageSet.Stage.Stool2DDiagrams {
			stool2ddiagramOrdered = append(stool2ddiagramOrdered, stool2ddiagram)
		}
		sort.Slice(stool2ddiagramOrdered, func(i, j int) bool {
			return stageSet.Stage.Stool2DDiagram_stagedOrder[stool2ddiagramOrdered[i]] < stageSet.Stage.Stool2DDiagram_stagedOrder[stool2ddiagramOrdered[j]]
		})
		for _, stool2ddiagram := range stool2ddiagramOrdered {
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
		stool3ddiagramOrdered := []*Stool3DDiagram{}
		for stool3ddiagram := range stageSet.Stage.Stool3DDiagrams {
			stool3ddiagramOrdered = append(stool3ddiagramOrdered, stool3ddiagram)
		}
		sort.Slice(stool3ddiagramOrdered, func(i, j int) bool {
			return stageSet.Stage.Stool3DDiagram_stagedOrder[stool3ddiagramOrdered[i]] < stageSet.Stage.Stool3DDiagram_stagedOrder[stool3ddiagramOrdered[j]]
		})
		for _, stool3ddiagram := range stool3ddiagramOrdered {
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
			if stool3ddiagram.SeatTopCurveShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.SeatTopCurveShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SeatTopCurveShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.RotatedSeatTopCurveShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.RotatedSeatTopCurveShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RotatedSeatTopCurveShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.SeatBottomCurveShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.SeatBottomCurveShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SeatBottomCurveShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.RotatedSeatBottomCurveShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.RotatedSeatBottomCurveShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RotatedSeatBottomCurveShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.Torus3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.Torus3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Torus3DShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.RotatedTorusShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.RotatedTorusShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RotatedTorusShape = %s", stool3ddiagramIdent, targetIdent))
			}
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
			if stool3ddiagram.RotatedSampledPoints3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.RotatedSampledPoints3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RotatedSampledPoints3DShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.EyeSampledPoints3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.EyeSampledPoints3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EyeSampledPoints3DShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.EyeCornersSampledPoints3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.EyeCornersSampledPoints3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EyeCornersSampledPoints3DShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.Eye3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.Eye3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Eye3DShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.EyeSeatBottomCurveShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.EyeSeatBottomCurveShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EyeSeatBottomCurveShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.EyeStoolBottomCurveShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.EyeStoolBottomCurveShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EyeStoolBottomCurveShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.Seat3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.Seat3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Seat3DShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.EyeVolume3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.EyeVolume3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EyeVolume3DShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.SeatAndLegs3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.SeatAndLegs3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SeatAndLegs3DShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.RotatedSeatAndLegs3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.RotatedSeatAndLegs3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RotatedSeatAndLegs3DShape = %s", stool3ddiagramIdent, targetIdent))
			}
			if stool3ddiagram.TiledFloor3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + stool3ddiagram.TiledFloor3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TiledFloor3DShape = %s", stool3ddiagramIdent, targetIdent))
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
		stoolabstractOrdered := []*StoolAbstract{}
		for stoolabstract := range stageSet.Stage.StoolAbstracts {
			stoolabstractOrdered = append(stoolabstractOrdered, stoolabstract)
		}
		sort.Slice(stoolabstractOrdered, func(i, j int) bool {
			return stageSet.Stage.StoolAbstract_stagedOrder[stoolabstractOrdered[i]] < stageSet.Stage.StoolAbstract_stagedOrder[stoolabstractOrdered[j]]
		})
		for _, stoolabstract := range stoolabstractOrdered {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			stoolabstractIdent := "__models" + stoolabstract.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StoolAbstract{Name: %s}).Stage(stageSet.Stage)", stoolabstractIdent, __gong__toRawStringLiteral(stoolabstract.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
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
		topcurveplane1shapeOrdered := []*TopCurvePlane1Shape{}
		for topcurveplane1shape := range stageSet.Stage.TopCurvePlane1Shapes {
			topcurveplane1shapeOrdered = append(topcurveplane1shapeOrdered, topcurveplane1shape)
		}
		sort.Slice(topcurveplane1shapeOrdered, func(i, j int) bool {
			return stageSet.Stage.TopCurvePlane1Shape_stagedOrder[topcurveplane1shapeOrdered[i]] < stageSet.Stage.TopCurvePlane1Shape_stagedOrder[topcurveplane1shapeOrdered[j]]
		})
		for _, topcurveplane1shape := range topcurveplane1shapeOrdered {
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
		topcurveplane2shapeOrdered := []*TopCurvePlane2Shape{}
		for topcurveplane2shape := range stageSet.Stage.TopCurvePlane2Shapes {
			topcurveplane2shapeOrdered = append(topcurveplane2shapeOrdered, topcurveplane2shape)
		}
		sort.Slice(topcurveplane2shapeOrdered, func(i, j int) bool {
			return stageSet.Stage.TopCurvePlane2Shape_stagedOrder[topcurveplane2shapeOrdered[i]] < stageSet.Stage.TopCurvePlane2Shape_stagedOrder[topcurveplane2shapeOrdered[j]]
		})
		for _, topcurveplane2shape := range topcurveplane2shapeOrdered {
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
		tubevase3ddiagramOrdered := []*TubeVase3DDiagram{}
		for tubevase3ddiagram := range stageSet.Stage.TubeVase3DDiagrams {
			tubevase3ddiagramOrdered = append(tubevase3ddiagramOrdered, tubevase3ddiagram)
		}
		sort.Slice(tubevase3ddiagramOrdered, func(i, j int) bool {
			return stageSet.Stage.TubeVase3DDiagram_stagedOrder[tubevase3ddiagramOrdered[i]] < stageSet.Stage.TubeVase3DDiagram_stagedOrder[tubevase3ddiagramOrdered[j]]
		})
		for _, tubevase3ddiagram := range tubevase3ddiagramOrdered {
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
			if tubevase3ddiagram.TorusStackShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.TorusStackShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TorusStackShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.VerticalTorusStackShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.VerticalTorusStackShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalTorusStackShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.PartiallyRotatedTorusShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.PartiallyRotatedTorusShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartiallyRotatedTorusShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.StackOfPartiallyRotatedTorusShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.StackOfPartiallyRotatedTorusShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StackOfPartiallyRotatedTorusShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.PointsAndLines3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.PointsAndLines3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PointsAndLines3DShape = %s", tubevase3ddiagramIdent, targetIdent))
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
			if tubevase3ddiagram.KeyHole3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.KeyHole3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.KeyHole3DShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.Key3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.Key3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Key3DShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.VolumeKey3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.VolumeKey3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VolumeKey3DShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.TorusEdge3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.TorusEdge3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TorusEdge3DShape = %s", tubevase3ddiagramIdent, targetIdent))
			}
			if tubevase3ddiagram.TiledFloor3DShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevase3ddiagram.TiledFloor3DShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TiledFloor3DShape = %s", tubevase3ddiagramIdent, targetIdent))
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
		}
	}
	if stageSet.Stage != nil {
		tubevaseabstractOrdered := []*TubeVaseAbstract{}
		for tubevaseabstract := range stageSet.Stage.TubeVaseAbstracts {
			tubevaseabstractOrdered = append(tubevaseabstractOrdered, tubevaseabstract)
		}
		sort.Slice(tubevaseabstractOrdered, func(i, j int) bool {
			return stageSet.Stage.TubeVaseAbstract_stagedOrder[tubevaseabstractOrdered[i]] < stageSet.Stage.TubeVaseAbstract_stagedOrder[tubevaseabstractOrdered[j]]
		})
		for _, tubevaseabstract := range tubevaseabstractOrdered {
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
			if tubevaseabstract.PerpendicularVectorGridHalfway != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.PerpendicularVectorGridHalfway.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PerpendicularVectorGridHalfway = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.TopStartArcShapeGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.TopStartArcShapeGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TopStartArcShapeGrid = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.TopEndArcShapeGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.TopEndArcShapeGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TopEndArcShapeGrid = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.ShiftedBottomTopStartArcShapeGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.ShiftedBottomTopStartArcShapeGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ShiftedBottomTopStartArcShapeGrid = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.TopMidArcVectorShapeGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.TopMidArcVectorShapeGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TopMidArcVectorShapeGrid = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.StartHalfwayArcShapeGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.StartHalfwayArcShapeGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartHalfwayArcShapeGrid = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.TopStartHalfwayArcShapeGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.TopStartHalfwayArcShapeGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TopStartHalfwayArcShapeGrid = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.EndHalfwayArcShapeGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.EndHalfwayArcShapeGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndHalfwayArcShapeGrid = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.TopEndHalfwayArcShapeGrid != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.TopEndHalfwayArcShapeGrid.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TopEndHalfwayArcShapeGrid = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.StackOfRotatedGrowthCurve2D != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.StackOfRotatedGrowthCurve2D.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StackOfRotatedGrowthCurve2D = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.TopStackOfRotatedGrowthCurve2D != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.TopStackOfRotatedGrowthCurve2D.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TopStackOfRotatedGrowthCurve2D = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.TopGrowthCurve2D != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.TopGrowthCurve2D.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TopGrowthCurve2D = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.StackOfGrowthCurve2D != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.StackOfGrowthCurve2D.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StackOfGrowthCurve2D = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.TopStackOfGrowthCurve2D != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.TopStackOfGrowthCurve2D.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TopStackOfGrowthCurve2D = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.StackOfGrowthCurve2DRibbon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.StackOfGrowthCurve2DRibbon.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StackOfGrowthCurve2DRibbon = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.StackOfRotatedGrowthCurve2DRibbon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.StackOfRotatedGrowthCurve2DRibbon.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StackOfRotatedGrowthCurve2DRibbon = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.GrowthCurve2DRibbon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.GrowthCurve2DRibbon.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GrowthCurve2DRibbon = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.ShiftedRightGrowthCurve2DRibbon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.ShiftedRightGrowthCurve2DRibbon.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ShiftedRightGrowthCurve2DRibbon = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.ShiftedLeftGrowthCurve2DRibbon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.ShiftedLeftGrowthCurve2DRibbon.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ShiftedLeftGrowthCurve2DRibbon = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.PartiallyGrowthCurve2DRibbon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.PartiallyGrowthCurve2DRibbon.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartiallyGrowthCurve2DRibbon = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.ShiftedLeftPartiallyGrowthCurve2DRibbon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.ShiftedLeftPartiallyGrowthCurve2DRibbon.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ShiftedLeftPartiallyGrowthCurve2DRibbon = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.PartiallyGrowthCurve2DTrajectory != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.PartiallyGrowthCurve2DTrajectory.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartiallyGrowthCurve2DTrajectory = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.PartiallyGrowthCurve2DTrajectoryP1P2 != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.PartiallyGrowthCurve2DTrajectoryP1P2.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartiallyGrowthCurve2DTrajectoryP1P2 = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.PxShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.PxShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PxShape = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.ChosenP1P2PairShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.ChosenP1P2PairShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ChosenP1P2PairShape = %s", tubevaseabstractIdent, targetIdent))
			}
			if tubevaseabstract.KeyHoleShape != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + tubevaseabstract.KeyHoleShape.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.KeyHoleShape = %s", tubevaseabstractIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		vase2ddiagramOrdered := []*Vase2DDiagram{}
		for vase2ddiagram := range stageSet.Stage.Vase2DDiagrams {
			vase2ddiagramOrdered = append(vase2ddiagramOrdered, vase2ddiagram)
		}
		sort.Slice(vase2ddiagramOrdered, func(i, j int) bool {
			return stageSet.Stage.Vase2DDiagram_stagedOrder[vase2ddiagramOrdered[i]] < stageSet.Stage.Vase2DDiagram_stagedOrder[vase2ddiagramOrdered[j]]
		})
		for _, vase2ddiagram := range vase2ddiagramOrdered {
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


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
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
			case "models":
				switch typeName {
				case "Angle0Shape":
					if !preserveOrder {
						inst := (&Angle0Shape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Angle0Shape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "BottomCurvePlane1Shape":
					if !preserveOrder {
						inst := (&BottomCurvePlane1Shape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(BottomCurvePlane1Shape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "BottomCurvePlane2Shape":
					if !preserveOrder {
						inst := (&BottomCurvePlane2Shape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(BottomCurvePlane2Shape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Circumference3DShape":
					if !preserveOrder {
						inst := (&Circumference3DShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Circumference3DShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Clock2DDiagram":
					if !preserveOrder {
						inst := (&Clock2DDiagram{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Clock2DDiagram)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Clock3DDiagram":
					if !preserveOrder {
						inst := (&Clock3DDiagram{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Clock3DDiagram)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ClockAbstract":
					if !preserveOrder {
						inst := (&ClockAbstract{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ClockAbstract)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "CutLine3DShape":
					if !preserveOrder {
						inst := (&CutLine3DShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(CutLine3DShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Leaves3DShape":
					if !preserveOrder {
						inst := (&Leaves3DShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Leaves3DShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Library":
					if !preserveOrder {
						inst := (&Library{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Library)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "MusicAbstract":
					if !preserveOrder {
						inst := (&MusicAbstract{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(MusicAbstract)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "OriginalPoints3DShape":
					if !preserveOrder {
						inst := (&OriginalPoints3DShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(OriginalPoints3DShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ParastichyMCurves3DShape":
					if !preserveOrder {
						inst := (&ParastichyMCurves3DShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ParastichyMCurves3DShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ParastichyNCurves3DShape":
					if !preserveOrder {
						inst := (&ParastichyNCurves3DShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ParastichyNCurves3DShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Plant2DDiagram":
					if !preserveOrder {
						inst := (&Plant2DDiagram{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Plant2DDiagram)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Plant3DDiagram":
					if !preserveOrder {
						inst := (&Plant3DDiagram{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Plant3DDiagram)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "PlantAbstract":
					if !preserveOrder {
						inst := (&PlantAbstract{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(PlantAbstract)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Rendered3DShape":
					if !preserveOrder {
						inst := (&Rendered3DShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Rendered3DShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "SampledPoints3DShape":
					if !preserveOrder {
						inst := (&SampledPoints3DShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(SampledPoints3DShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StemCylinder3DShape":
					if !preserveOrder {
						inst := (&StemCylinder3DShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StemCylinder3DShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Stool2DDiagram":
					if !preserveOrder {
						inst := (&Stool2DDiagram{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Stool2DDiagram)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Stool3DDiagram":
					if !preserveOrder {
						inst := (&Stool3DDiagram{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Stool3DDiagram)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "StoolAbstract":
					if !preserveOrder {
						inst := (&StoolAbstract{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(StoolAbstract)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TopCurvePlane1Shape":
					if !preserveOrder {
						inst := (&TopCurvePlane1Shape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TopCurvePlane1Shape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TopCurvePlane2Shape":
					if !preserveOrder {
						inst := (&TopCurvePlane2Shape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TopCurvePlane2Shape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TubeVase3DDiagram":
					if !preserveOrder {
						inst := (&TubeVase3DDiagram{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TubeVase3DDiagram)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "TubeVaseAbstract":
					if !preserveOrder {
						inst := (&TubeVaseAbstract{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(TubeVaseAbstract)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Vase2DDiagram":
					if !preserveOrder {
						inst := (&Vase2DDiagram{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Vase2DDiagram)
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
					case "ClockTopCurveShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ClockTopCurveShape); ok {
									inst.ClockTopCurveShape = typedTarget
								}
							}
						}
					case "IsHiddenTorus3DShape":
						inst.IsHiddenTorus3DShape = GongExtractBool(rhs)
					case "Torus3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Torus3DShape); ok {
									inst.Torus3DShape = typedTarget
								}
							}
						}
					case "IsHiddenSampledPoints3DShape":
						inst.IsHiddenSampledPoints3DShape = GongExtractBool(rhs)
					case "SampledPoints3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SampledPoints3DShape); ok {
									inst.SampledPoints3DShape = typedTarget
								}
							}
						}
					case "IsHiddenTiledFloor3DShape":
						inst.IsHiddenTiledFloor3DShape = GongExtractBool(rhs)
					case "TiledFloor3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TiledFloor3DShape); ok {
									inst.TiledFloor3DShape = typedTarget
								}
							}
						}
					case "Rendered3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Rendered3DShape); ok {
									inst.Rendered3DShape = typedTarget
								}
							}
						}
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *ClockAbstract:
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*PlantAbstract); ok {
										inst.Plants = append(inst.Plants, typedTarget)
									}
								}
							}
						}
					case "SubLibraries":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Library); ok {
										inst.SubLibraries = append(inst.SubLibraries, typedTarget)
									}
								}
							}
						}
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
				case *MusicAbstract:
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StemCylinder3DShape); ok {
									inst.StemCylinder3DShape = typedTarget
								}
							}
						}
					case "IsHiddenParastichyNCurves3DShape":
						inst.IsHiddenParastichyNCurves3DShape = GongExtractBool(rhs)
					case "ParastichyNCurves3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ParastichyNCurves3DShape); ok {
									inst.ParastichyNCurves3DShape = typedTarget
								}
							}
						}
					case "IsHiddenParastichyMCurves3DShape":
						inst.IsHiddenParastichyMCurves3DShape = GongExtractBool(rhs)
					case "ParastichyMCurves3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ParastichyMCurves3DShape); ok {
									inst.ParastichyMCurves3DShape = typedTarget
								}
							}
						}
					case "IsHiddenCutLine3DShape":
						inst.IsHiddenCutLine3DShape = GongExtractBool(rhs)
					case "CutLine3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*CutLine3DShape); ok {
									inst.CutLine3DShape = typedTarget
								}
							}
						}
					case "IsHiddenCircumference3DShape":
						inst.IsHiddenCircumference3DShape = GongExtractBool(rhs)
					case "Circumference3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Circumference3DShape); ok {
									inst.Circumference3DShape = typedTarget
								}
							}
						}
					case "IsHiddenTiledFloor3DShape":
						inst.IsHiddenTiledFloor3DShape = GongExtractBool(rhs)
					case "TiledFloor3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TiledFloor3DShape); ok {
									inst.TiledFloor3DShape = typedTarget
								}
							}
						}
					case "IsHiddenLeaves3DShape":
						inst.IsHiddenLeaves3DShape = GongExtractBool(rhs)
					case "Leaves3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Leaves3DShape); ok {
									inst.Leaves3DShape = typedTarget
								}
							}
						}
					case "Rendered3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Rendered3DShape); ok {
									inst.Rendered3DShape = typedTarget
								}
							}
						}
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
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TubeVaseAbstract); ok {
									inst.TubeVaseAbstract = typedTarget
								}
							}
						}
					case "StoolAbstract":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StoolAbstract); ok {
									inst.StoolAbstract = typedTarget
								}
							}
						}
					case "ClockAbstract":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ClockAbstract); ok {
									inst.ClockAbstract = typedTarget
								}
							}
						}
					case "MusicAbstract":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*MusicAbstract); ok {
									inst.MusicAbstract = typedTarget
								}
							}
						}
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
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Plant2DDiagram); ok {
										inst.Plant2DDiagrams = append(inst.Plant2DDiagrams, typedTarget)
									}
								}
							}
						}
					case "IsPlant3DDiagramsNodeExpanded":
						inst.IsPlant3DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "Plant3DDiagrams":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Plant3DDiagram); ok {
										inst.Plant3DDiagrams = append(inst.Plant3DDiagrams, typedTarget)
									}
								}
							}
						}
					case "IsVase2DDiagramsNodeExpanded":
						inst.IsVase2DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "Vase2DDiagrams":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Vase2DDiagram); ok {
										inst.Vase2DDiagrams = append(inst.Vase2DDiagrams, typedTarget)
									}
								}
							}
						}
					case "IsTubeVase3DDiagramsNodeExpanded":
						inst.IsTubeVase3DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "TubeVase3DDiagrams":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*TubeVase3DDiagram); ok {
										inst.TubeVase3DDiagrams = append(inst.TubeVase3DDiagrams, typedTarget)
									}
								}
							}
						}
					case "IsStool2DDiagramsNodeExpanded":
						inst.IsStool2DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "Stool2DDiagrams":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Stool2DDiagram); ok {
										inst.Stool2DDiagrams = append(inst.Stool2DDiagrams, typedTarget)
									}
								}
							}
						}
					case "IsStool3DDiagramsNodeExpanded":
						inst.IsStool3DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "Stool3DDiagrams":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Stool3DDiagram); ok {
										inst.Stool3DDiagrams = append(inst.Stool3DDiagrams, typedTarget)
									}
								}
							}
						}
					case "IsClock2DDiagramsNodeExpanded":
						inst.IsClock2DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "Clock2DDiagrams":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Clock2DDiagram); ok {
										inst.Clock2DDiagrams = append(inst.Clock2DDiagrams, typedTarget)
									}
								}
							}
						}
					case "IsClock3DDiagramsNodeExpanded":
						inst.IsClock3DDiagramsNodeExpanded = GongExtractBool(rhs)
					case "Clock3DDiagrams":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
									if typedTarget, ok := target.(*Clock3DDiagram); ok {
										inst.Clock3DDiagrams = append(inst.Clock3DDiagrams, typedTarget)
									}
								}
							}
						}
					case "AxesShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*AxesShape); ok {
									inst.AxesShape = typedTarget
								}
							}
						}
					case "RhombusStuff":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*RhombusStuff); ok {
									inst.RhombusStuff = typedTarget
								}
							}
						}
					case "GrowthVectorShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*GrowthVectorShape); ok {
									inst.GrowthVectorShape = typedTarget
								}
							}
						}
					case "PerpendicularVectorGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PerpendicularVectorGrid); ok {
									inst.PerpendicularVectorGrid = typedTarget
								}
							}
						}
					case "BaseVectorShapeGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*BaseVectorShapeGrid); ok {
									inst.BaseVectorShapeGrid = typedTarget
								}
							}
						}
					case "ArcNormalVectorShapeGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ArcNormalVectorShapeGrid); ok {
									inst.ArcNormalVectorShapeGrid = typedTarget
								}
							}
						}
					case "StartArcShapeGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StartArcShapeGrid); ok {
									inst.StartArcShapeGrid = typedTarget
								}
							}
						}
					case "MidArcVectorShapeGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*MidArcVectorShapeGrid); ok {
									inst.MidArcVectorShapeGrid = typedTarget
								}
							}
						}
					case "EndArcShapeGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*EndArcShapeGrid); ok {
									inst.EndArcShapeGrid = typedTarget
								}
							}
						}
					case "GrowthCurve2D":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*GrowthCurve2D); ok {
									inst.GrowthCurve2D = typedTarget
								}
							}
						}
					case "StackOfGrowthCurve2DByGrowthVector":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StackOfGrowthCurve2DByGrowthVector); ok {
									inst.StackOfGrowthCurve2DByGrowthVector = typedTarget
								}
							}
						}
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
					case "SeatTopCurveShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SeatTopCurveShape); ok {
									inst.SeatTopCurveShape = typedTarget
								}
							}
						}
					case "IsHiddenRotatedSeatTopCurveShape":
						inst.IsHiddenRotatedSeatTopCurveShape = GongExtractBool(rhs)
					case "RotatedSeatTopCurveShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PartiallyRotatedSeatTopCurveShape); ok {
									inst.RotatedSeatTopCurveShape = typedTarget
								}
							}
						}
					case "IsHiddenSeatBottomCurveShape":
						inst.IsHiddenSeatBottomCurveShape = GongExtractBool(rhs)
					case "SeatBottomCurveShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SeatBottomCurveShape); ok {
									inst.SeatBottomCurveShape = typedTarget
								}
							}
						}
					case "IsHiddenRotatedSeatBottomCurveShape":
						inst.IsHiddenRotatedSeatBottomCurveShape = GongExtractBool(rhs)
					case "RotatedSeatBottomCurveShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PartiallyRotatedSeatBottomCurveShape); ok {
									inst.RotatedSeatBottomCurveShape = typedTarget
								}
							}
						}
					case "IsHiddenTorus3DShape":
						inst.IsHiddenTorus3DShape = GongExtractBool(rhs)
					case "Torus3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Torus3DShape); ok {
									inst.Torus3DShape = typedTarget
								}
							}
						}
					case "IsHiddenRotatedTorusShape":
						inst.IsHiddenRotatedTorusShape = GongExtractBool(rhs)
					case "RotatedTorusShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PartiallyRotatedTorusShape); ok {
									inst.RotatedTorusShape = typedTarget
								}
							}
						}
					case "IsHiddenSampledPoints3DShape":
						inst.IsHiddenSampledPoints3DShape = GongExtractBool(rhs)
					case "SampledPoints3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SampledPoints3DShape); ok {
									inst.SampledPoints3DShape = typedTarget
								}
							}
						}
					case "IsHiddenRotatedSampledPoints3DShape":
						inst.IsHiddenRotatedSampledPoints3DShape = GongExtractBool(rhs)
					case "RotatedSampledPoints3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*RotatedSampledPoints3DShape); ok {
									inst.RotatedSampledPoints3DShape = typedTarget
								}
							}
						}
					case "IsHiddenEyeSampledPoints3DShape":
						inst.IsHiddenEyeSampledPoints3DShape = GongExtractBool(rhs)
					case "EyeSampledPoints3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*EyeSampledPoints3DShape); ok {
									inst.EyeSampledPoints3DShape = typedTarget
								}
							}
						}
					case "IsHiddenEyeCornersSampledPoints3DShape":
						inst.IsHiddenEyeCornersSampledPoints3DShape = GongExtractBool(rhs)
					case "EyeCornersSampledPoints3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*EyeCornersSampledPoints3DShape); ok {
									inst.EyeCornersSampledPoints3DShape = typedTarget
								}
							}
						}
					case "IsHiddenEye3DShape":
						inst.IsHiddenEye3DShape = GongExtractBool(rhs)
					case "Eye3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Eye3DShape); ok {
									inst.Eye3DShape = typedTarget
								}
							}
						}
					case "IsHiddenEyeSeatBottomCurveShape":
						inst.IsHiddenEyeSeatBottomCurveShape = GongExtractBool(rhs)
					case "EyeSeatBottomCurveShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*EyeSeatBottomCurveShape); ok {
									inst.EyeSeatBottomCurveShape = typedTarget
								}
							}
						}
					case "IsHiddenEyeStoolBottomCurveShape":
						inst.IsHiddenEyeStoolBottomCurveShape = GongExtractBool(rhs)
					case "EyeStoolBottomCurveShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*EyeStoolBottomCurveShape); ok {
									inst.EyeStoolBottomCurveShape = typedTarget
								}
							}
						}
					case "IsHiddenSeat3DShape":
						inst.IsHiddenSeat3DShape = GongExtractBool(rhs)
					case "Seat3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Seat3DShape); ok {
									inst.Seat3DShape = typedTarget
								}
							}
						}
					case "IsHiddenEyeVolume3DShape":
						inst.IsHiddenEyeVolume3DShape = GongExtractBool(rhs)
					case "EyeVolume3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*EyeVolume3DShape); ok {
									inst.EyeVolume3DShape = typedTarget
								}
							}
						}
					case "IsHiddenSeatAndLegs3DShape":
						inst.IsHiddenSeatAndLegs3DShape = GongExtractBool(rhs)
					case "SeatAndLegs3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SeatAndLegs3DShape); ok {
									inst.SeatAndLegs3DShape = typedTarget
								}
							}
						}
					case "IsHiddenRotatedSeatAndLegs3DShape":
						inst.IsHiddenRotatedSeatAndLegs3DShape = GongExtractBool(rhs)
					case "RotatedSeatAndLegs3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*RotatedSeatAndLegs3DShape); ok {
									inst.RotatedSeatAndLegs3DShape = typedTarget
								}
							}
						}
					case "IsHiddenTiledFloor3DShape":
						inst.IsHiddenTiledFloor3DShape = GongExtractBool(rhs)
					case "TiledFloor3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TiledFloor3DShape); ok {
									inst.TiledFloor3DShape = typedTarget
								}
							}
						}
					case "Rendered3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Rendered3DShape); ok {
									inst.Rendered3DShape = typedTarget
								}
							}
						}
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *StoolAbstract:
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
					case "Rendered3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Rendered3DShape); ok {
									inst.Rendered3DShape = typedTarget
								}
							}
						}
					case "TorusStackShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TorusStackShape); ok {
									inst.TorusStackShape = typedTarget
								}
							}
						}
					case "VerticalTorusStackShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*VerticalTorusStackShape); ok {
									inst.VerticalTorusStackShape = typedTarget
								}
							}
						}
					case "PartiallyRotatedTorusShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PartiallyRotatedTorusShape); ok {
									inst.PartiallyRotatedTorusShape = typedTarget
								}
							}
						}
					case "StackOfPartiallyRotatedTorusShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StackOfPartiallyRotatedTorusShape); ok {
									inst.StackOfPartiallyRotatedTorusShape = typedTarget
								}
							}
						}
					case "PointsAndLines3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PointsAndLines3DShape); ok {
									inst.PointsAndLines3DShape = typedTarget
								}
							}
						}
					case "SampledPoints3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*SampledPoints3DShape); ok {
									inst.SampledPoints3DShape = typedTarget
								}
							}
						}
					case "OriginalPoints3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*OriginalPoints3DShape); ok {
									inst.OriginalPoints3DShape = typedTarget
								}
							}
						}
					case "Angle0Shape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Angle0Shape); ok {
									inst.Angle0Shape = typedTarget
								}
							}
						}
					case "KeyHole3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*KeyHole3DShape); ok {
									inst.KeyHole3DShape = typedTarget
								}
							}
						}
					case "Key3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Key3DShape); ok {
									inst.Key3DShape = typedTarget
								}
							}
						}
					case "VolumeKey3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*VolumeKey3DShape); ok {
									inst.VolumeKey3DShape = typedTarget
								}
							}
						}
					case "TorusEdge3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TorusEdge3DShape); ok {
									inst.TorusEdge3DShape = typedTarget
								}
							}
						}
					case "TiledFloor3DShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TiledFloor3DShape); ok {
									inst.TiledFloor3DShape = typedTarget
								}
							}
						}
					case "TopCurvePlane1Shape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TopCurvePlane1Shape); ok {
									inst.TopCurvePlane1Shape = typedTarget
								}
							}
						}
					case "BottomCurvePlane1Shape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*BottomCurvePlane1Shape); ok {
									inst.BottomCurvePlane1Shape = typedTarget
								}
							}
						}
					case "TopCurvePlane2Shape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TopCurvePlane2Shape); ok {
									inst.TopCurvePlane2Shape = typedTarget
								}
							}
						}
					case "BottomCurvePlane2Shape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*BottomCurvePlane2Shape); ok {
									inst.BottomCurvePlane2Shape = typedTarget
								}
							}
						}
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
					case "PerpendicularVectorGridHalfway":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PerpendicularVectorGridHalfway); ok {
									inst.PerpendicularVectorGridHalfway = typedTarget
								}
							}
						}
					case "TopStartArcShapeGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TopStartArcShapeGrid); ok {
									inst.TopStartArcShapeGrid = typedTarget
								}
							}
						}
					case "TopEndArcShapeGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TopEndArcShapeGrid); ok {
									inst.TopEndArcShapeGrid = typedTarget
								}
							}
						}
					case "ShiftedBottomTopStartArcShapeGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ShiftedBottomTopStartArcShapeGrid); ok {
									inst.ShiftedBottomTopStartArcShapeGrid = typedTarget
								}
							}
						}
					case "TopMidArcVectorShapeGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TopMidArcVectorShapeGrid); ok {
									inst.TopMidArcVectorShapeGrid = typedTarget
								}
							}
						}
					case "StartHalfwayArcShapeGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StartHalfwayArcShapeGrid); ok {
									inst.StartHalfwayArcShapeGrid = typedTarget
								}
							}
						}
					case "TopStartHalfwayArcShapeGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TopStartHalfwayArcShapeGrid); ok {
									inst.TopStartHalfwayArcShapeGrid = typedTarget
								}
							}
						}
					case "EndHalfwayArcShapeGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*EndHalfwayArcShapeGrid); ok {
									inst.EndHalfwayArcShapeGrid = typedTarget
								}
							}
						}
					case "TopEndHalfwayArcShapeGrid":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TopEndHalfwayArcShapeGrid); ok {
									inst.TopEndHalfwayArcShapeGrid = typedTarget
								}
							}
						}
					case "StackOfRotatedGrowthCurve2D":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StackOfRotatedGrowthCurve2D); ok {
									inst.StackOfRotatedGrowthCurve2D = typedTarget
								}
							}
						}
					case "TopStackOfRotatedGrowthCurve2D":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TopStackOfRotatedGrowthCurve2D); ok {
									inst.TopStackOfRotatedGrowthCurve2D = typedTarget
								}
							}
						}
					case "TopGrowthCurve2D":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TopGrowthCurve2D); ok {
									inst.TopGrowthCurve2D = typedTarget
								}
							}
						}
					case "StackOfGrowthCurve2D":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StackOfGrowthCurve2D); ok {
									inst.StackOfGrowthCurve2D = typedTarget
								}
							}
						}
					case "TopStackOfGrowthCurve2D":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*TopStackOfGrowthCurve2D); ok {
									inst.TopStackOfGrowthCurve2D = typedTarget
								}
							}
						}
					case "StackOfGrowthCurve2DRibbon":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StackOfGrowthCurve2DRibbon); ok {
									inst.StackOfGrowthCurve2DRibbon = typedTarget
								}
							}
						}
					case "StackOfRotatedGrowthCurve2DRibbon":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*StackOfRotatedGrowthCurve2DRibbon); ok {
									inst.StackOfRotatedGrowthCurve2DRibbon = typedTarget
								}
							}
						}
					case "GrowthCurve2DRibbon":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*GrowthCurve2DRibbon); ok {
									inst.GrowthCurve2DRibbon = typedTarget
								}
							}
						}
					case "ShiftedRightGrowthCurve2DRibbon":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ShiftedRightGrowthCurve2DRibbon); ok {
									inst.ShiftedRightGrowthCurve2DRibbon = typedTarget
								}
							}
						}
					case "ShiftedLeftGrowthCurve2DRibbon":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ShiftedLeftGrowthCurve2DRibbon); ok {
									inst.ShiftedLeftGrowthCurve2DRibbon = typedTarget
								}
							}
						}
					case "PartiallyGrowthCurve2DRibbon":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PartiallyGrowthCurve2DRibbon); ok {
									inst.PartiallyGrowthCurve2DRibbon = typedTarget
								}
							}
						}
					case "ShiftedLeftPartiallyGrowthCurve2DRibbon":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ShiftedLeftPartiallyGrowthCurve2DRibbon); ok {
									inst.ShiftedLeftPartiallyGrowthCurve2DRibbon = typedTarget
								}
							}
						}
					case "PartiallyGrowthCurve2DTrajectory":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PartiallyGrowthCurve2DTrajectory); ok {
									inst.PartiallyGrowthCurve2DTrajectory = typedTarget
								}
							}
						}
					case "PartiallyGrowthCurve2DTrajectoryP1P2":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PartiallyGrowthCurve2DTrajectoryP1P2); ok {
									inst.PartiallyGrowthCurve2DTrajectoryP1P2 = typedTarget
								}
							}
						}
					case "PxShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*PxShape); ok {
									inst.PxShape = typedTarget
								}
							}
						}
					case "ChosenP1P2PairShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ChosenP1P2PairShape); ok {
									inst.ChosenP1P2PairShape = typedTarget
								}
							}
						}
					case "KeyHoleShape":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*KeyHoleShape); ok {
									inst.KeyHoleShape = typedTarget
								}
							}
						}
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
