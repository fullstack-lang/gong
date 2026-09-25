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
		for _, assplit := range __gong__sortStageSetInstances(stageSet.Stage.AsSplits, stageSet.Stage.AsSplit_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			assplitIdent := "__models" + assplit.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AsSplit{Name: %s}).Stage(stageSet.Stage)", assplitIdent, __gong__toRawStringLiteral(assplit.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", assplitIdent, __gong__toRawStringLiteral(assplit.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", assplitIdent, __gong__toRawStringLiteral(string(assplit.Direction))))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSizeInPixel = %t", assplitIdent, assplit.IsSizeInPixel))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCustomGutterSize = %t", assplitIdent, assplit.IsWithCustomGutterSize))
			values.WriteString(fmt.Sprintf("\n\t%s.GutterSize = %f", assplitIdent, assplit.GutterSize))
			for _, elem := range assplit.AsSplitAreas {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AsSplitAreas = append(%s.AsSplitAreas, %s)", assplitIdent, assplitIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, assplitarea := range __gong__sortStageSetInstances(stageSet.Stage.AsSplitAreas, stageSet.Stage.AsSplitArea_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			assplitareaIdent := "__models" + assplitarea.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AsSplitArea{Name: %s}).Stage(stageSet.Stage)", assplitareaIdent, __gong__toRawStringLiteral(assplitarea.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", assplitareaIdent, __gong__toRawStringLiteral(assplitarea.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowNameInHeader = %t", assplitareaIdent, assplitarea.ShowNameInHeader))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %f", assplitareaIdent, assplitarea.Size))
			values.WriteString(fmt.Sprintf("\n\t%s.IsAny = %t", assplitareaIdent, assplitarea.IsAny))
			values.WriteString(fmt.Sprintf("\n\t%s.HasDiv = %t", assplitareaIdent, assplitarea.HasDiv))
			values.WriteString(fmt.Sprintf("\n\t%s.DivStyle = %s", assplitareaIdent, __gong__toRawStringLiteral(assplitarea.DivStyle)))
			if assplitarea.AsSplit != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.AsSplit.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AsSplit = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Button != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.Button.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Button = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Cursor != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.Cursor.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Cursor = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Form != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.Form.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Form = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Load != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.Load.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Load = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Markdown != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.Markdown.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Markdown = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Slider != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.Slider.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slider = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Split != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.Split.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Split = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Svg != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.Svg.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Svg = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Table != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.Table.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Table = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Tone != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.Tone.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tone = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Tree != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.Tree.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tree = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Threejs != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.Threejs.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Threejs = %s", assplitareaIdent, targetIdent))
			}
			if assplitarea.Xlsx != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + assplitarea.Xlsx.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Xlsx = %s", assplitareaIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, button := range __gong__sortStageSetInstances(stageSet.Stage.Buttons, stageSet.Stage.Button_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			buttonIdent := "__models" + button.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Button{Name: %s}).Stage(stageSet.Stage)", buttonIdent, __gong__toRawStringLiteral(button.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", buttonIdent, __gong__toRawStringLiteral(button.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", buttonIdent, __gong__toRawStringLiteral(button.StackName)))
		}
	}
	if stageSet.Stage != nil {
		for _, cursor := range __gong__sortStageSetInstances(stageSet.Stage.Cursors, stageSet.Stage.Cursor_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			cursorIdent := "__models" + cursor.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Cursor{Name: %s}).Stage(stageSet.Stage)", cursorIdent, __gong__toRawStringLiteral(cursor.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cursorIdent, __gong__toRawStringLiteral(cursor.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", cursorIdent, __gong__toRawStringLiteral(cursor.StackName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Style = %s", cursorIdent, __gong__toRawStringLiteral(cursor.Style)))
		}
	}
	if stageSet.Stage != nil {
		for _, favicon := range __gong__sortStageSetInstances(stageSet.Stage.FavIcons, stageSet.Stage.FavIcon_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			faviconIdent := "__models" + favicon.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FavIcon{Name: %s}).Stage(stageSet.Stage)", faviconIdent, __gong__toRawStringLiteral(favicon.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", faviconIdent, __gong__toRawStringLiteral(favicon.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG = %s", faviconIdent, __gong__toRawStringLiteral(favicon.SVG)))
		}
	}
	if stageSet.Stage != nil {
		for _, form := range __gong__sortStageSetInstances(stageSet.Stage.Forms, stageSet.Stage.Form_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			formIdent := "__models" + form.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Form{Name: %s}).Stage(stageSet.Stage)", formIdent, __gong__toRawStringLiteral(form.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", formIdent, __gong__toRawStringLiteral(form.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", formIdent, __gong__toRawStringLiteral(form.StackName)))
		}
	}
	if stageSet.Stage != nil {
		for _, load := range __gong__sortStageSetInstances(stageSet.Stage.Loads, stageSet.Stage.Load_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			loadIdent := "__models" + load.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Load{Name: %s}).Stage(stageSet.Stage)", loadIdent, __gong__toRawStringLiteral(load.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", loadIdent, __gong__toRawStringLiteral(load.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", loadIdent, __gong__toRawStringLiteral(load.StackName)))
		}
	}
	if stageSet.Stage != nil {
		for _, logoontheleft := range __gong__sortStageSetInstances(stageSet.Stage.LogoOnTheLefts, stageSet.Stage.LogoOnTheLeft_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			logoontheleftIdent := "__models" + logoontheleft.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LogoOnTheLeft{Name: %s}).Stage(stageSet.Stage)", logoontheleftIdent, __gong__toRawStringLiteral(logoontheleft.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", logoontheleftIdent, __gong__toRawStringLiteral(logoontheleft.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %d", logoontheleftIdent, logoontheleft.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %d", logoontheleftIdent, logoontheleft.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG = %s", logoontheleftIdent, __gong__toRawStringLiteral(logoontheleft.SVG)))
		}
	}
	if stageSet.Stage != nil {
		for _, logoontheright := range __gong__sortStageSetInstances(stageSet.Stage.LogoOnTheRights, stageSet.Stage.LogoOnTheRight_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			logoontherightIdent := "__models" + logoontheright.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LogoOnTheRight{Name: %s}).Stage(stageSet.Stage)", logoontherightIdent, __gong__toRawStringLiteral(logoontheright.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", logoontherightIdent, __gong__toRawStringLiteral(logoontheright.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %d", logoontherightIdent, logoontheright.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %d", logoontherightIdent, logoontheright.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.SVG = %s", logoontherightIdent, __gong__toRawStringLiteral(logoontheright.SVG)))
		}
	}
	if stageSet.Stage != nil {
		for _, markdown := range __gong__sortStageSetInstances(stageSet.Stage.Markdowns, stageSet.Stage.Markdown_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			markdownIdent := "__models" + markdown.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Markdown{Name: %s}).Stage(stageSet.Stage)", markdownIdent, __gong__toRawStringLiteral(markdown.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", markdownIdent, __gong__toRawStringLiteral(markdown.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", markdownIdent, __gong__toRawStringLiteral(markdown.StackName)))
		}
	}
	if stageSet.Stage != nil {
		for _, slider := range __gong__sortStageSetInstances(stageSet.Stage.Sliders, stageSet.Stage.Slider_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sliderIdent := "__models" + slider.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Slider{Name: %s}).Stage(stageSet.Stage)", sliderIdent, __gong__toRawStringLiteral(slider.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sliderIdent, __gong__toRawStringLiteral(slider.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", sliderIdent, __gong__toRawStringLiteral(slider.StackName)))
		}
	}
	if stageSet.Stage != nil {
		for _, split := range __gong__sortStageSetInstances(stageSet.Stage.Splits, stageSet.Stage.Split_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			splitIdent := "__models" + split.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Split{Name: %s}).Stage(stageSet.Stage)", splitIdent, __gong__toRawStringLiteral(split.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", splitIdent, __gong__toRawStringLiteral(split.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", splitIdent, __gong__toRawStringLiteral(split.StackName)))
		}
	}
	if stageSet.Stage != nil {
		for _, svg := range __gong__sortStageSetInstances(stageSet.Stage.Svgs, stageSet.Stage.Svg_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			svgIdent := "__models" + svg.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Svg{Name: %s}).Stage(stageSet.Stage)", svgIdent, __gong__toRawStringLiteral(svg.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", svgIdent, __gong__toRawStringLiteral(svg.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", svgIdent, __gong__toRawStringLiteral(svg.StackName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Style = %s", svgIdent, __gong__toRawStringLiteral(svg.Style)))
		}
	}
	if stageSet.Stage != nil {
		for _, table := range __gong__sortStageSetInstances(stageSet.Stage.Tables, stageSet.Stage.Table_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tableIdent := "__models" + table.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Table{Name: %s}).Stage(stageSet.Stage)", tableIdent, __gong__toRawStringLiteral(table.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tableIdent, __gong__toRawStringLiteral(table.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", tableIdent, __gong__toRawStringLiteral(table.StackName)))
		}
	}
	if stageSet.Stage != nil {
		for _, threejs := range __gong__sortStageSetInstances(stageSet.Stage.Threejss, stageSet.Stage.Threejs_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			threejsIdent := "__models" + threejs.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Threejs{Name: %s}).Stage(stageSet.Stage)", threejsIdent, __gong__toRawStringLiteral(threejs.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", threejsIdent, __gong__toRawStringLiteral(threejs.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", threejsIdent, __gong__toRawStringLiteral(threejs.StackName)))
		}
	}
	if stageSet.Stage != nil {
		for _, title := range __gong__sortStageSetInstances(stageSet.Stage.Titles, stageSet.Stage.Title_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			titleIdent := "__models" + title.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Title{Name: %s}).Stage(stageSet.Stage)", titleIdent, __gong__toRawStringLiteral(title.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", titleIdent, __gong__toRawStringLiteral(title.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, tone := range __gong__sortStageSetInstances(stageSet.Stage.Tones, stageSet.Stage.Tone_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			toneIdent := "__models" + tone.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tone{Name: %s}).Stage(stageSet.Stage)", toneIdent, __gong__toRawStringLiteral(tone.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", toneIdent, __gong__toRawStringLiteral(tone.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", toneIdent, __gong__toRawStringLiteral(tone.StackName)))
		}
	}
	if stageSet.Stage != nil {
		for _, tree := range __gong__sortStageSetInstances(stageSet.Stage.Trees, stageSet.Stage.Tree_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			treeIdent := "__models" + tree.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Tree{Name: %s}).Stage(stageSet.Stage)", treeIdent, __gong__toRawStringLiteral(tree.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", treeIdent, __gong__toRawStringLiteral(tree.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", treeIdent, __gong__toRawStringLiteral(tree.StackName)))
		}
	}
	if stageSet.Stage != nil {
		for _, view := range __gong__sortStageSetInstances(stageSet.Stage.Views, stageSet.Stage.View_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			viewIdent := "__models" + view.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.View{Name: %s}).Stage(stageSet.Stage)", viewIdent, __gong__toRawStringLiteral(view.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", viewIdent, __gong__toRawStringLiteral(view.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowViewName = %t", viewIdent, view.ShowViewName))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelectedView = %t", viewIdent, view.IsSelectedView))
			values.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", viewIdent, __gong__toRawStringLiteral(string(view.Direction))))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSecondaryView = %t", viewIdent, view.IsSecondaryView))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSizeInPixel = %t", viewIdent, view.IsSizeInPixel))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithCustomGutterSize = %t", viewIdent, view.IsWithCustomGutterSize))
			values.WriteString(fmt.Sprintf("\n\t%s.GutterSize = %f", viewIdent, view.GutterSize))
			for _, elem := range view.RootAsSplitAreas {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootAsSplitAreas = append(%s.RootAsSplitAreas, %s)", viewIdent, viewIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, xlsx := range __gong__sortStageSetInstances(stageSet.Stage.Xlsxs, stageSet.Stage.Xlsx_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			xlsxIdent := "__models" + xlsx.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Xlsx{Name: %s}).Stage(stageSet.Stage)", xlsxIdent, __gong__toRawStringLiteral(xlsx.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", xlsxIdent, __gong__toRawStringLiteral(xlsx.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StackName = %s", xlsxIdent, __gong__toRawStringLiteral(xlsx.StackName)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/split/go/models"
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
		case "github.com/fullstack-lang/gong/lib/split/go/models":
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
				case "AsSplit":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AsSplit), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "AsSplitArea":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AsSplitArea), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Button":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Button), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Cursor":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Cursor), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "FavIcon":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FavIcon), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Form":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Form), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Load":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Load), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "LogoOnTheLeft":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LogoOnTheLeft), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "LogoOnTheRight":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LogoOnTheRight), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Markdown":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Markdown), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Slider":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Slider), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Split":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Split), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Svg":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Svg), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Table":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Table), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Threejs":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Threejs), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Title":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Title), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tone":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tone), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Tree":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Tree), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "View":
					identifierMap[ident.Name] = __gong__stageSetInit(new(View), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Xlsx":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Xlsx), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *AsSplit:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Direction":
						inst.Direction = Direction(GongExtractString(rhs))
					case "AsSplitAreas":
						__gong__assignSliceOfPointers(&inst.AsSplitAreas, rhs, identifierMap)
					case "IsSizeInPixel":
						inst.IsSizeInPixel = GongExtractBool(rhs)
					case "IsWithCustomGutterSize":
						inst.IsWithCustomGutterSize = GongExtractBool(rhs)
					case "GutterSize":
						inst.GutterSize = GongExtractFloat(rhs)
					}
				case *AsSplitArea:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowNameInHeader":
						inst.ShowNameInHeader = GongExtractBool(rhs)
					case "Size":
						inst.Size = GongExtractFloat(rhs)
					case "IsAny":
						inst.IsAny = GongExtractBool(rhs)
					case "AsSplit":
						__gong__assignPointer(&inst.AsSplit, rhs, identifierMap)
					case "Button":
						__gong__assignPointer(&inst.Button, rhs, identifierMap)
					case "Cursor":
						__gong__assignPointer(&inst.Cursor, rhs, identifierMap)
					case "Form":
						__gong__assignPointer(&inst.Form, rhs, identifierMap)
					case "Load":
						__gong__assignPointer(&inst.Load, rhs, identifierMap)
					case "Markdown":
						__gong__assignPointer(&inst.Markdown, rhs, identifierMap)
					case "Slider":
						__gong__assignPointer(&inst.Slider, rhs, identifierMap)
					case "Split":
						__gong__assignPointer(&inst.Split, rhs, identifierMap)
					case "Svg":
						__gong__assignPointer(&inst.Svg, rhs, identifierMap)
					case "Table":
						__gong__assignPointer(&inst.Table, rhs, identifierMap)
					case "Tone":
						__gong__assignPointer(&inst.Tone, rhs, identifierMap)
					case "Tree":
						__gong__assignPointer(&inst.Tree, rhs, identifierMap)
					case "Threejs":
						__gong__assignPointer(&inst.Threejs, rhs, identifierMap)
					case "Xlsx":
						__gong__assignPointer(&inst.Xlsx, rhs, identifierMap)
					case "HasDiv":
						inst.HasDiv = GongExtractBool(rhs)
					case "DivStyle":
						inst.DivStyle = GongExtractString(rhs)
					}
				case *Button:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Cursor:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					case "Style":
						inst.Style = GongExtractString(rhs)
					}
				case *FavIcon:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SVG":
						inst.SVG = GongExtractString(rhs)
					}
				case *Form:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Load:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *LogoOnTheLeft:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractInt(rhs)
					case "Height":
						inst.Height = GongExtractInt(rhs)
					case "SVG":
						inst.SVG = GongExtractString(rhs)
					}
				case *LogoOnTheRight:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Width":
						inst.Width = GongExtractInt(rhs)
					case "Height":
						inst.Height = GongExtractInt(rhs)
					case "SVG":
						inst.SVG = GongExtractString(rhs)
					}
				case *Markdown:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Slider:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Split:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Svg:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					case "Style":
						inst.Style = GongExtractString(rhs)
					}
				case *Table:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Threejs:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Title:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *Tone:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *Tree:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
					}
				case *View:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ShowViewName":
						inst.ShowViewName = GongExtractBool(rhs)
					case "RootAsSplitAreas":
						__gong__assignSliceOfPointers(&inst.RootAsSplitAreas, rhs, identifierMap)
					case "IsSelectedView":
						inst.IsSelectedView = GongExtractBool(rhs)
					case "Direction":
						inst.Direction = Direction(GongExtractString(rhs))
					case "IsSecondaryView":
						inst.IsSecondaryView = GongExtractBool(rhs)
					case "IsSizeInPixel":
						inst.IsSizeInPixel = GongExtractBool(rhs)
					case "IsWithCustomGutterSize":
						inst.IsWithCustomGutterSize = GongExtractBool(rhs)
					case "GutterSize":
						inst.GutterSize = GongExtractFloat(rhs)
					}
				case *Xlsx:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "StackName":
						inst.StackName = GongExtractString(rhs)
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
