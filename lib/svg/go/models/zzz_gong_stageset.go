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
		for _, animate := range __gong__sortStageSetInstances(stageSet.Stage.Animates, stageSet.Stage.Animate_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			animateIdent := "__models" + animate.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Animate{Name: %s}).Stage(stageSet.Stage)", animateIdent, __gong__toRawStringLiteral(animate.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", animateIdent, __gong__toRawStringLiteral(animate.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.AttributeName = %s", animateIdent, __gong__toRawStringLiteral(animate.AttributeName)))
			values.WriteString(fmt.Sprintf("\n\t%s.Values = %s", animateIdent, __gong__toRawStringLiteral(animate.Values)))
			values.WriteString(fmt.Sprintf("\n\t%s.From = %s", animateIdent, __gong__toRawStringLiteral(animate.From)))
			values.WriteString(fmt.Sprintf("\n\t%s.To = %s", animateIdent, __gong__toRawStringLiteral(animate.To)))
			values.WriteString(fmt.Sprintf("\n\t%s.Dur = %s", animateIdent, __gong__toRawStringLiteral(animate.Dur)))
			values.WriteString(fmt.Sprintf("\n\t%s.RepeatCount = %s", animateIdent, __gong__toRawStringLiteral(animate.RepeatCount)))
		}
	}
	if stageSet.Stage != nil {
		for _, circle := range __gong__sortStageSetInstances(stageSet.Stage.Circles, stageSet.Stage.Circle_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			circleIdent := "__models" + circle.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Circle{Name: %s}).Stage(stageSet.Stage)", circleIdent, __gong__toRawStringLiteral(circle.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", circleIdent, __gong__toRawStringLiteral(circle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.CX = %f", circleIdent, circle.CX))
			values.WriteString(fmt.Sprintf("\n\t%s.CY = %f", circleIdent, circle.CY))
			values.WriteString(fmt.Sprintf("\n\t%s.Radius = %f", circleIdent, circle.Radius))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", circleIdent, __gong__toRawStringLiteral(circle.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", circleIdent, circle.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", circleIdent, __gong__toRawStringLiteral(circle.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", circleIdent, circle.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", circleIdent, circle.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", circleIdent, __gong__toRawStringLiteral(circle.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", circleIdent, __gong__toRawStringLiteral(circle.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", circleIdent, __gong__toRawStringLiteral(circle.Transform)))
			for _, elem := range circle.Animations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animations = append(%s.Animations, %s)", circleIdent, circleIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, condition := range __gong__sortStageSetInstances(stageSet.Stage.Conditions, stageSet.Stage.Condition_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			conditionIdent := "__models" + condition.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Condition{Name: %s}).Stage(stageSet.Stage)", conditionIdent, __gong__toRawStringLiteral(condition.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", conditionIdent, __gong__toRawStringLiteral(condition.Name)))
		}
	}
	if stageSet.Stage != nil {
		for _, controlpoint := range __gong__sortStageSetInstances(stageSet.Stage.ControlPoints, stageSet.Stage.ControlPoint_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			controlpointIdent := "__models" + controlpoint.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ControlPoint{Name: %s}).Stage(stageSet.Stage)", controlpointIdent, __gong__toRawStringLiteral(controlpoint.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", controlpointIdent, __gong__toRawStringLiteral(controlpoint.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Relative = %f", controlpointIdent, controlpoint.X_Relative))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Relative = %f", controlpointIdent, controlpoint.Y_Relative))
			if controlpoint.ClosestRect != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + controlpoint.ClosestRect.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ClosestRect = %s", controlpointIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, ellipse := range __gong__sortStageSetInstances(stageSet.Stage.Ellipses, stageSet.Stage.Ellipse_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			ellipseIdent := "__models" + ellipse.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Ellipse{Name: %s}).Stage(stageSet.Stage)", ellipseIdent, __gong__toRawStringLiteral(ellipse.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ellipseIdent, __gong__toRawStringLiteral(ellipse.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.CX = %f", ellipseIdent, ellipse.CX))
			values.WriteString(fmt.Sprintf("\n\t%s.CY = %f", ellipseIdent, ellipse.CY))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", ellipseIdent, ellipse.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.RY = %f", ellipseIdent, ellipse.RY))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", ellipseIdent, __gong__toRawStringLiteral(ellipse.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", ellipseIdent, ellipse.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", ellipseIdent, __gong__toRawStringLiteral(ellipse.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", ellipseIdent, ellipse.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", ellipseIdent, ellipse.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", ellipseIdent, __gong__toRawStringLiteral(ellipse.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", ellipseIdent, __gong__toRawStringLiteral(ellipse.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", ellipseIdent, __gong__toRawStringLiteral(ellipse.Transform)))
			for _, elem := range ellipse.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", ellipseIdent, ellipseIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, filetodownload := range __gong__sortStageSetInstances(stageSet.Stage.FileToDownloads, stageSet.Stage.FileToDownload_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			filetodownloadIdent := "__models" + filetodownload.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.FileToDownload{Name: %s}).Stage(stageSet.Stage)", filetodownloadIdent, __gong__toRawStringLiteral(filetodownload.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", filetodownloadIdent, __gong__toRawStringLiteral(filetodownload.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64EncodedContent = %s", filetodownloadIdent, __gong__toRawStringLiteral(filetodownload.Base64EncodedContent)))
		}
	}
	if stageSet.Stage != nil {
		for _, layer := range __gong__sortStageSetInstances(stageSet.Stage.Layers, stageSet.Stage.Layer_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			layerIdent := "__models" + layer.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Layer{Name: %s}).Stage(stageSet.Stage)", layerIdent, __gong__toRawStringLiteral(layer.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", layerIdent, __gong__toRawStringLiteral(layer.Name)))
			for _, elem := range layer.Rects {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rects = append(%s.Rects, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Texts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Texts = append(%s.Texts, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Circles {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Circles = append(%s.Circles, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Lines {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lines = append(%s.Lines, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Ellipses {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ellipses = append(%s.Ellipses, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Polylines {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Polylines = append(%s.Polylines, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Polygones {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Polygones = append(%s.Polygones, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Paths {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Paths = append(%s.Paths, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.Links {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Links = append(%s.Links, %s)", layerIdent, layerIdent, targetIdent))
			}
			for _, elem := range layer.RectLinkLinks {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RectLinkLinks = append(%s.RectLinkLinks, %s)", layerIdent, layerIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, line := range __gong__sortStageSetInstances(stageSet.Stage.Lines, stageSet.Stage.Line_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lineIdent := "__models" + line.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Line{Name: %s}).Stage(stageSet.Stage)", lineIdent, __gong__toRawStringLiteral(line.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lineIdent, __gong__toRawStringLiteral(line.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X1 = %f", lineIdent, line.X1))
			values.WriteString(fmt.Sprintf("\n\t%s.Y1 = %f", lineIdent, line.Y1))
			values.WriteString(fmt.Sprintf("\n\t%s.X2 = %f", lineIdent, line.X2))
			values.WriteString(fmt.Sprintf("\n\t%s.Y2 = %f", lineIdent, line.Y2))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", lineIdent, __gong__toRawStringLiteral(line.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", lineIdent, line.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", lineIdent, __gong__toRawStringLiteral(line.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", lineIdent, line.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", lineIdent, line.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", lineIdent, __gong__toRawStringLiteral(line.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", lineIdent, __gong__toRawStringLiteral(line.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", lineIdent, __gong__toRawStringLiteral(line.Transform)))
			values.WriteString(fmt.Sprintf("\n\t%s.MouseClickX = %f", lineIdent, line.MouseClickX))
			values.WriteString(fmt.Sprintf("\n\t%s.MouseClickY = %f", lineIdent, line.MouseClickY))
			for _, elem := range line.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", lineIdent, lineIdent, targetIdent))
			}
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
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", linkIdent, __gong__toRawStringLiteral(string(link.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.IsBezierCurve = %t", linkIdent, link.IsBezierCurve))
			values.WriteString(fmt.Sprintf("\n\t%s.StartAnchorType = %s", linkIdent, __gong__toRawStringLiteral(string(link.StartAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndAnchorType = %s", linkIdent, __gong__toRawStringLiteral(string(link.EndAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = %s", linkIdent, __gong__toRawStringLiteral(string(link.StartOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", linkIdent, link.StartRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = %s", linkIdent, __gong__toRawStringLiteral(string(link.EndOrientation))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", linkIdent, link.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", linkIdent, link.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerRadius = %f", linkIdent, link.CornerRadius))
			values.WriteString(fmt.Sprintf("\n\t%s.HasEndArrow = %t", linkIdent, link.HasEndArrow))
			values.WriteString(fmt.Sprintf("\n\t%s.EndArrowSize = %f", linkIdent, link.EndArrowSize))
			values.WriteString(fmt.Sprintf("\n\t%s.EndArrowOffset = %f", linkIdent, link.EndArrowOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.HasStartArrow = %t", linkIdent, link.HasStartArrow))
			values.WriteString(fmt.Sprintf("\n\t%s.StartArrowSize = %f", linkIdent, link.StartArrowSize))
			values.WriteString(fmt.Sprintf("\n\t%s.StartArrowOffset = %f", linkIdent, link.StartArrowOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", linkIdent, __gong__toRawStringLiteral(link.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", linkIdent, link.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", linkIdent, __gong__toRawStringLiteral(link.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", linkIdent, link.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", linkIdent, link.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", linkIdent, __gong__toRawStringLiteral(link.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", linkIdent, __gong__toRawStringLiteral(link.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", linkIdent, __gong__toRawStringLiteral(link.Transform)))
			values.WriteString(fmt.Sprintf("\n\t%s.MouseX = %f", linkIdent, link.MouseX))
			values.WriteString(fmt.Sprintf("\n\t%s.MouseY = %f", linkIdent, link.MouseY))
			values.WriteString(fmt.Sprintf("\n\t%s.MouseEventKey = %s", linkIdent, __gong__toRawStringLiteral(string(link.MouseEventKey))))
			if link.Start != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + link.Start.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", linkIdent, targetIdent))
			}
			if link.End != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + link.End.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End = %s", linkIdent, targetIdent))
			}
			for _, elem := range link.TextAtArrowStart {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TextAtArrowStart = append(%s.TextAtArrowStart, %s)", linkIdent, linkIdent, targetIdent))
			}
			for _, elem := range link.TextAtArrowEnd {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TextAtArrowEnd = append(%s.TextAtArrowEnd, %s)", linkIdent, linkIdent, targetIdent))
			}
			for _, elem := range link.TextAtCorner {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TextAtCorner = append(%s.TextAtCorner, %s)", linkIdent, linkIdent, targetIdent))
			}
			for _, elem := range link.PathAtArrowStart {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PathAtArrowStart = append(%s.PathAtArrowStart, %s)", linkIdent, linkIdent, targetIdent))
			}
			for _, elem := range link.PathAtArrowEnd {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PathAtArrowEnd = append(%s.PathAtArrowEnd, %s)", linkIdent, linkIdent, targetIdent))
			}
			for _, elem := range link.PathAtCorner {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PathAtCorner = append(%s.PathAtCorner, %s)", linkIdent, linkIdent, targetIdent))
			}
			for _, elem := range link.ControlPoints {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPoints = append(%s.ControlPoints, %s)", linkIdent, linkIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, linkanchoredpath := range __gong__sortStageSetInstances(stageSet.Stage.LinkAnchoredPaths, stageSet.Stage.LinkAnchoredPath_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linkanchoredpathIdent := "__models" + linkanchoredpath.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LinkAnchoredPath{Name: %s}).Stage(stageSet.Stage)", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Definition = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.Definition)))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", linkanchoredpathIdent, linkanchoredpath.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", linkanchoredpathIdent, linkanchoredpath.Y_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.ScalePropotionnally = %t", linkanchoredpathIdent, linkanchoredpath.ScalePropotionnally))
			values.WriteString(fmt.Sprintf("\n\t%s.AppliedScaling = %f", linkanchoredpathIdent, linkanchoredpath.AppliedScaling))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", linkanchoredpathIdent, linkanchoredpath.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", linkanchoredpathIdent, linkanchoredpath.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", linkanchoredpathIdent, linkanchoredpath.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", linkanchoredpathIdent, __gong__toRawStringLiteral(linkanchoredpath.Transform)))
		}
	}
	if stageSet.Stage != nil {
		for _, linkanchoredtext := range __gong__sortStageSetInstances(stageSet.Stage.LinkAnchoredTexts, stageSet.Stage.LinkAnchoredText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linkanchoredtextIdent := "__models" + linkanchoredtext.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LinkAnchoredText{Name: %s}).Stage(stageSet.Stage)", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.AutomaticLayout = %t", linkanchoredtextIdent, linkanchoredtext.AutomaticLayout))
			values.WriteString(fmt.Sprintf("\n\t%s.LinkAnchorType = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(string(linkanchoredtext.LinkAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", linkanchoredtextIdent, linkanchoredtext.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", linkanchoredtextIdent, linkanchoredtext.Y_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.FontStyle)))
			values.WriteString(fmt.Sprintf("\n\t%s.LetterSpacing = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.LetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.WhiteSpace = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(string(linkanchoredtext.WhiteSpace))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", linkanchoredtextIdent, linkanchoredtext.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", linkanchoredtextIdent, linkanchoredtext.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", linkanchoredtextIdent, linkanchoredtext.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", linkanchoredtextIdent, __gong__toRawStringLiteral(linkanchoredtext.Transform)))
			for _, elem := range linkanchoredtext.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", linkanchoredtextIdent, linkanchoredtextIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, path := range __gong__sortStageSetInstances(stageSet.Stage.Paths, stageSet.Stage.Path_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pathIdent := "__models" + path.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Path{Name: %s}).Stage(stageSet.Stage)", pathIdent, __gong__toRawStringLiteral(path.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pathIdent, __gong__toRawStringLiteral(path.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Definition = %s", pathIdent, __gong__toRawStringLiteral(path.Definition)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", pathIdent, __gong__toRawStringLiteral(path.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", pathIdent, path.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", pathIdent, __gong__toRawStringLiteral(path.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", pathIdent, path.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", pathIdent, path.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", pathIdent, __gong__toRawStringLiteral(path.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", pathIdent, __gong__toRawStringLiteral(path.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", pathIdent, __gong__toRawStringLiteral(path.Transform)))
			for _, elem := range path.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", pathIdent, pathIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, point := range __gong__sortStageSetInstances(stageSet.Stage.Points, stageSet.Stage.Point_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pointIdent := "__models" + point.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Point{Name: %s}).Stage(stageSet.Stage)", pointIdent, __gong__toRawStringLiteral(point.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointIdent, __gong__toRawStringLiteral(point.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", pointIdent, point.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", pointIdent, point.Y))
		}
	}
	if stageSet.Stage != nil {
		for _, polygone := range __gong__sortStageSetInstances(stageSet.Stage.Polygones, stageSet.Stage.Polygone_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			polygoneIdent := "__models" + polygone.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Polygone{Name: %s}).Stage(stageSet.Stage)", polygoneIdent, __gong__toRawStringLiteral(polygone.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Points = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.Points)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", polygoneIdent, polygone.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", polygoneIdent, polygone.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", polygoneIdent, polygone.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", polygoneIdent, __gong__toRawStringLiteral(polygone.Transform)))
			for _, elem := range polygone.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", polygoneIdent, polygoneIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, polyline := range __gong__sortStageSetInstances(stageSet.Stage.Polylines, stageSet.Stage.Polyline_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			polylineIdent := "__models" + polyline.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Polyline{Name: %s}).Stage(stageSet.Stage)", polylineIdent, __gong__toRawStringLiteral(polyline.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", polylineIdent, __gong__toRawStringLiteral(polyline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Points = %s", polylineIdent, __gong__toRawStringLiteral(polyline.Points)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", polylineIdent, __gong__toRawStringLiteral(polyline.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", polylineIdent, polyline.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", polylineIdent, __gong__toRawStringLiteral(polyline.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", polylineIdent, polyline.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", polylineIdent, polyline.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", polylineIdent, __gong__toRawStringLiteral(polyline.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", polylineIdent, __gong__toRawStringLiteral(polyline.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", polylineIdent, __gong__toRawStringLiteral(polyline.Transform)))
			for _, elem := range polyline.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", polylineIdent, polylineIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, rect := range __gong__sortStageSetInstances(stageSet.Stage.Rects, stageSet.Stage.Rect_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rectIdent := "__models" + rect.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Rect{Name: %s}).Stage(stageSet.Stage)", rectIdent, __gong__toRawStringLiteral(rect.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rectIdent, __gong__toRawStringLiteral(rect.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", rectIdent, rect.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", rectIdent, rect.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", rectIdent, rect.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", rectIdent, rect.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", rectIdent, rect.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", rectIdent, __gong__toRawStringLiteral(rect.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", rectIdent, rect.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", rectIdent, __gong__toRawStringLiteral(rect.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", rectIdent, rect.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", rectIdent, rect.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", rectIdent, __gong__toRawStringLiteral(rect.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", rectIdent, __gong__toRawStringLiteral(rect.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", rectIdent, __gong__toRawStringLiteral(rect.Transform)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelectable = %t", rectIdent, rect.IsSelectable))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelected = %t", rectIdent, rect.IsSelected))
			values.WriteString(fmt.Sprintf("\n\t%s.CanHaveLeftHandle = %t", rectIdent, rect.CanHaveLeftHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.HasLeftHandle = %t", rectIdent, rect.HasLeftHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.CanHaveRightHandle = %t", rectIdent, rect.CanHaveRightHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.HasRightHandle = %t", rectIdent, rect.HasRightHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.CanHaveTopHandle = %t", rectIdent, rect.CanHaveTopHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.HasTopHandle = %t", rectIdent, rect.HasTopHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.IsScalingProportionally = %t", rectIdent, rect.IsScalingProportionally))
			values.WriteString(fmt.Sprintf("\n\t%s.CanHaveBottomHandle = %t", rectIdent, rect.CanHaveBottomHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.HasBottomHandle = %t", rectIdent, rect.HasBottomHandle))
			values.WriteString(fmt.Sprintf("\n\t%s.CanMoveHorizontaly = %t", rectIdent, rect.CanMoveHorizontaly))
			values.WriteString(fmt.Sprintf("\n\t%s.CanMoveVerticaly = %t", rectIdent, rect.CanMoveVerticaly))
			values.WriteString(fmt.Sprintf("\n\t%s.ChangeColorWhenHovered = %t", rectIdent, rect.ChangeColorWhenHovered))
			values.WriteString(fmt.Sprintf("\n\t%s.ColorWhenHovered = %s", rectIdent, __gong__toRawStringLiteral(rect.ColorWhenHovered)))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginalColor = %s", rectIdent, __gong__toRawStringLiteral(rect.OriginalColor)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacityWhenHovered = %f", rectIdent, rect.FillOpacityWhenHovered))
			values.WriteString(fmt.Sprintf("\n\t%s.OriginalFillOpacity = %f", rectIdent, rect.OriginalFillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.HasToolTip = %t", rectIdent, rect.HasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipText = %s", rectIdent, __gong__toRawStringLiteral(rect.ToolTipText)))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipPosition = %s", rectIdent, __gong__toRawStringLiteral(string(rect.ToolTipPosition))))
			values.WriteString(fmt.Sprintf("\n\t%s.MouseX = %f", rectIdent, rect.MouseX))
			values.WriteString(fmt.Sprintf("\n\t%s.MouseY = %f", rectIdent, rect.MouseY))
			values.WriteString(fmt.Sprintf("\n\t%s.MouseEventKey = %s", rectIdent, __gong__toRawStringLiteral(string(rect.MouseEventKey))))
			values.WriteString(fmt.Sprintf("\n\t%s.URLPath = %s", rectIdent, __gong__toRawStringLiteral(rect.URLPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.URLTarget = %s", rectIdent, __gong__toRawStringLiteral(string(rect.URLTarget))))
			for _, elem := range rect.Peers {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Peers = append(%s.Peers, %s)", rectIdent, rectIdent, targetIdent))
			}
			if rect.EnclosingRect != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + rect.EnclosingRect.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingRect = %s", rectIdent, targetIdent))
			}
			for _, elem := range rect.Obstacles {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Obstacles = append(%s.Obstacles, %s)", rectIdent, rectIdent, targetIdent))
			}
			if rect.AnchoredTo != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + rect.AnchoredTo.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AnchoredTo = %s", rectIdent, targetIdent))
			}
			for _, elem := range rect.HoveringTrigger {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HoveringTrigger = append(%s.HoveringTrigger, %s)", rectIdent, rectIdent, targetIdent))
			}
			for _, elem := range rect.DisplayConditions {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DisplayConditions = append(%s.DisplayConditions, %s)", rectIdent, rectIdent, targetIdent))
			}
			for _, elem := range rect.Animations {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animations = append(%s.Animations, %s)", rectIdent, rectIdent, targetIdent))
			}
			for _, elem := range rect.RectAnchoredTexts {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RectAnchoredTexts = append(%s.RectAnchoredTexts, %s)", rectIdent, rectIdent, targetIdent))
			}
			for _, elem := range rect.RectAnchoredRects {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RectAnchoredRects = append(%s.RectAnchoredRects, %s)", rectIdent, rectIdent, targetIdent))
			}
			for _, elem := range rect.RectAnchoredPaths {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RectAnchoredPaths = append(%s.RectAnchoredPaths, %s)", rectIdent, rectIdent, targetIdent))
			}
			for _, elem := range rect.RectAnchoredPngImages {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RectAnchoredPngImages = append(%s.RectAnchoredPngImages, %s)", rectIdent, rectIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, rectanchoredpath := range __gong__sortStageSetInstances(stageSet.Stage.RectAnchoredPaths, stageSet.Stage.RectAnchoredPath_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rectanchoredpathIdent := "__models" + rectanchoredpath.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RectAnchoredPath{Name: %s}).Stage(stageSet.Stage)", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Definition = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.Definition)))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", rectanchoredpathIdent, rectanchoredpath.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", rectanchoredpathIdent, rectanchoredpath.Y_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(string(rectanchoredpath.RectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ScalePropotionnally = %t", rectanchoredpathIdent, rectanchoredpath.ScalePropotionnally))
			values.WriteString(fmt.Sprintf("\n\t%s.AppliedScaling = %f", rectanchoredpathIdent, rectanchoredpath.AppliedScaling))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", rectanchoredpathIdent, rectanchoredpath.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", rectanchoredpathIdent, rectanchoredpath.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", rectanchoredpathIdent, rectanchoredpath.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", rectanchoredpathIdent, __gong__toRawStringLiteral(rectanchoredpath.Transform)))
		}
	}
	if stageSet.Stage != nil {
		for _, rectanchoredpngimage := range __gong__sortStageSetInstances(stageSet.Stage.RectAnchoredPngImages, stageSet.Stage.RectAnchoredPngImage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rectanchoredpngimageIdent := "__models" + rectanchoredpngimage.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RectAnchoredPngImage{Name: %s}).Stage(stageSet.Stage)", rectanchoredpngimageIdent, __gong__toRawStringLiteral(rectanchoredpngimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rectanchoredpngimageIdent, __gong__toRawStringLiteral(rectanchoredpngimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", rectanchoredpngimageIdent, rectanchoredpngimage.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", rectanchoredpngimageIdent, rectanchoredpngimage.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", rectanchoredpngimageIdent, rectanchoredpngimage.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", rectanchoredpngimageIdent, rectanchoredpngimage.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", rectanchoredpngimageIdent, rectanchoredpngimage.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", rectanchoredpngimageIdent, rectanchoredpngimage.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", rectanchoredpngimageIdent, rectanchoredpngimage.Y_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = %s", rectanchoredpngimageIdent, __gong__toRawStringLiteral(string(rectanchoredpngimage.RectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", rectanchoredpngimageIdent, __gong__toRawStringLiteral(rectanchoredpngimage.Base64Content)))
		}
	}
	if stageSet.Stage != nil {
		for _, rectanchoredrect := range __gong__sortStageSetInstances(stageSet.Stage.RectAnchoredRects, stageSet.Stage.RectAnchoredRect_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rectanchoredrectIdent := "__models" + rectanchoredrect.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RectAnchoredRect{Name: %s}).Stage(stageSet.Stage)", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", rectanchoredrectIdent, rectanchoredrect.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", rectanchoredrectIdent, rectanchoredrect.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", rectanchoredrectIdent, rectanchoredrect.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", rectanchoredrectIdent, rectanchoredrect.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.RX = %f", rectanchoredrectIdent, rectanchoredrect.RX))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", rectanchoredrectIdent, rectanchoredrect.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", rectanchoredrectIdent, rectanchoredrect.Y_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(string(rectanchoredrect.RectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.WidthFollowRect = %t", rectanchoredrectIdent, rectanchoredrect.WidthFollowRect))
			values.WriteString(fmt.Sprintf("\n\t%s.HeightFollowRect = %t", rectanchoredrectIdent, rectanchoredrect.HeightFollowRect))
			values.WriteString(fmt.Sprintf("\n\t%s.HasToolTip = %t", rectanchoredrectIdent, rectanchoredrect.HasToolTip))
			values.WriteString(fmt.Sprintf("\n\t%s.ToolTipText = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.ToolTipText)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", rectanchoredrectIdent, rectanchoredrect.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", rectanchoredrectIdent, rectanchoredrect.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", rectanchoredrectIdent, rectanchoredrect.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", rectanchoredrectIdent, __gong__toRawStringLiteral(rectanchoredrect.Transform)))
		}
	}
	if stageSet.Stage != nil {
		for _, rectanchoredtext := range __gong__sortStageSetInstances(stageSet.Stage.RectAnchoredTexts, stageSet.Stage.RectAnchoredText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rectanchoredtextIdent := "__models" + rectanchoredtext.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RectAnchoredText{Name: %s}).Stage(stageSet.Stage)", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.FontStyle)))
			values.WriteString(fmt.Sprintf("\n\t%s.LetterSpacing = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.LetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.WhiteSpace = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(string(rectanchoredtext.WhiteSpace))))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Offset = %f", rectanchoredtextIdent, rectanchoredtext.X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Offset = %f", rectanchoredtextIdent, rectanchoredtext.Y_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.RectAnchorType = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(string(rectanchoredtext.RectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.TextAnchorType = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(string(rectanchoredtext.TextAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.DominantBaseline = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(string(rectanchoredtext.DominantBaseline))))
			values.WriteString(fmt.Sprintf("\n\t%s.WritingMode = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(string(rectanchoredtext.WritingMode))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", rectanchoredtextIdent, rectanchoredtext.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", rectanchoredtextIdent, rectanchoredtext.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", rectanchoredtextIdent, rectanchoredtext.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.Transform)))
			values.WriteString(fmt.Sprintf("\n\t%s.URLPath = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(rectanchoredtext.URLPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.URLTarget = %s", rectanchoredtextIdent, __gong__toRawStringLiteral(string(rectanchoredtext.URLTarget))))
			for _, elem := range rectanchoredtext.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", rectanchoredtextIdent, rectanchoredtextIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, rectlinklink := range __gong__sortStageSetInstances(stageSet.Stage.RectLinkLinks, stageSet.Stage.RectLinkLink_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			rectlinklinkIdent := "__models" + rectlinklink.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RectLinkLink{Name: %s}).Stage(stageSet.Stage)", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetAnchorPosition = %f", rectlinklinkIdent, rectlinklink.TargetAnchorPosition))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", rectlinklinkIdent, rectlinklink.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", rectlinklinkIdent, rectlinklink.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", rectlinklinkIdent, rectlinklink.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", rectlinklinkIdent, __gong__toRawStringLiteral(rectlinklink.Transform)))
			if rectlinklink.Start != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + rectlinklink.Start.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Start = %s", rectlinklinkIdent, targetIdent))
			}
			if rectlinklink.End != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + rectlinklink.End.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.End = %s", rectlinklinkIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, svg := range __gong__sortStageSetInstances(stageSet.Stage.SVGs, stageSet.Stage.SVG_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			svgIdent := "__models" + svg.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SVG{Name: %s}).Stage(stageSet.Stage)", svgIdent, __gong__toRawStringLiteral(svg.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", svgIdent, __gong__toRawStringLiteral(svg.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DrawingState = %s", svgIdent, __gong__toRawStringLiteral(string(svg.DrawingState))))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable = %t", svgIdent, svg.IsEditable))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSVGFrontEndFileGenerated = %t", svgIdent, svg.IsSVGFrontEndFileGenerated))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSVGBackEndFileGenerated = %t", svgIdent, svg.IsSVGBackEndFileGenerated))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultDirectoryForGeneratedImages = %s", svgIdent, __gong__toRawStringLiteral(svg.DefaultDirectoryForGeneratedImages)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsControlBannerHidden = %t", svgIdent, svg.IsControlBannerHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.PanX = %f", svgIdent, svg.PanX))
			values.WriteString(fmt.Sprintf("\n\t%s.PanY = %f", svgIdent, svg.PanY))
			values.WriteString(fmt.Sprintf("\n\t%s.Zoom = %f", svgIdent, svg.Zoom))
			values.WriteString(fmt.Sprintf("\n\t%s.OverrideWidth = %t", svgIdent, svg.OverrideWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.OverriddenWidth = %f", svgIdent, svg.OverriddenWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.OverrideHeight = %t", svgIdent, svg.OverrideHeight))
			values.WriteString(fmt.Sprintf("\n\t%s.OverriddenHeight = %f", svgIdent, svg.OverriddenHeight))
			for _, elem := range svg.Layers {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Layers = append(%s.Layers, %s)", svgIdent, svgIdent, targetIdent))
			}
			if svg.StartRect != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + svg.StartRect.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StartRect = %s", svgIdent, targetIdent))
			}
			if svg.EndRect != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + svg.EndRect.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndRect = %s", svgIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, svgtext := range __gong__sortStageSetInstances(stageSet.Stage.SvgTexts, stageSet.Stage.SvgText_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			svgtextIdent := "__models" + svgtext.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SvgText{Name: %s}).Stage(stageSet.Stage)", svgtextIdent, __gong__toRawStringLiteral(svgtext.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", svgtextIdent, __gong__toRawStringLiteral(svgtext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", svgtextIdent, __gong__toRawStringLiteral(svgtext.Text)))
		}
	}
	if stageSet.Stage != nil {
		for _, text := range __gong__sortStageSetInstances(stageSet.Stage.Texts, stageSet.Stage.Text_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textIdent := "__models" + text.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Text{Name: %s}).Stage(stageSet.Stage)", textIdent, __gong__toRawStringLiteral(text.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textIdent, __gong__toRawStringLiteral(text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", textIdent, text.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", textIdent, text.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", textIdent, __gong__toRawStringLiteral(text.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", textIdent, __gong__toRawStringLiteral(text.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.FillOpacity = %f", textIdent, text.FillOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.Stroke = %s", textIdent, __gong__toRawStringLiteral(text.Stroke)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeOpacity = %f", textIdent, text.StrokeOpacity))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeWidth = %f", textIdent, text.StrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArray = %s", textIdent, __gong__toRawStringLiteral(text.StrokeDashArray)))
			values.WriteString(fmt.Sprintf("\n\t%s.StrokeDashArrayWhenSelected = %s", textIdent, __gong__toRawStringLiteral(text.StrokeDashArrayWhenSelected)))
			values.WriteString(fmt.Sprintf("\n\t%s.Transform = %s", textIdent, __gong__toRawStringLiteral(text.Transform)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", textIdent, __gong__toRawStringLiteral(text.FontWeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", textIdent, __gong__toRawStringLiteral(text.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", textIdent, __gong__toRawStringLiteral(text.FontStyle)))
			values.WriteString(fmt.Sprintf("\n\t%s.LetterSpacing = %s", textIdent, __gong__toRawStringLiteral(text.LetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", textIdent, __gong__toRawStringLiteral(text.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.WhiteSpace = %s", textIdent, __gong__toRawStringLiteral(string(text.WhiteSpace))))
			for _, elem := range text.Animates {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__models" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Animates = append(%s.Animates, %s)", textIdent, textIdent, targetIdent))
			}
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
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
		case "github.com/fullstack-lang/gong/lib/svg/go/models":
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
				case "Animate":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Animate), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Circle":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Circle), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Condition":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Condition), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "ControlPoint":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ControlPoint), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Ellipse":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Ellipse), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "FileToDownload":
					identifierMap[ident.Name] = __gong__stageSetInit(new(FileToDownload), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Layer":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Layer), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Line":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Line), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Link":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Link), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "LinkAnchoredPath":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinkAnchoredPath), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "LinkAnchoredText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(LinkAnchoredText), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Path":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Path), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Point":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Point), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Polygone":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Polygone), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Polyline":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Polyline), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Rect":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Rect), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "RectAnchoredPath":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RectAnchoredPath), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "RectAnchoredPngImage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RectAnchoredPngImage), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "RectAnchoredRect":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RectAnchoredRect), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "RectAnchoredText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RectAnchoredText), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "RectLinkLink":
					identifierMap[ident.Name] = __gong__stageSetInit(new(RectLinkLink), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SVG":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SVG), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "SvgText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SvgText), stageSet.Stage, ident.Name, instanceName, preserveOrder)
				case "Text":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Text), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *Animate:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "AttributeName":
						inst.AttributeName = GongExtractString(rhs)
					case "Values":
						inst.Values = GongExtractString(rhs)
					case "From":
						inst.From = GongExtractString(rhs)
					case "To":
						inst.To = GongExtractString(rhs)
					case "Dur":
						inst.Dur = GongExtractString(rhs)
					case "RepeatCount":
						inst.RepeatCount = GongExtractString(rhs)
					}
				case *Circle:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "CX":
						inst.CX = GongExtractFloat(rhs)
					case "CY":
						inst.CY = GongExtractFloat(rhs)
					case "Radius":
						inst.Radius = GongExtractFloat(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animations":
						__gong__assignSliceOfPointers(&inst.Animations, rhs, identifierMap)
					}
				case *Condition:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *ControlPoint:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X_Relative":
						inst.X_Relative = GongExtractFloat(rhs)
					case "Y_Relative":
						inst.Y_Relative = GongExtractFloat(rhs)
					case "ClosestRect":
						__gong__assignPointer(&inst.ClosestRect, rhs, identifierMap)
					}
				case *Ellipse:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "CX":
						inst.CX = GongExtractFloat(rhs)
					case "CY":
						inst.CY = GongExtractFloat(rhs)
					case "RX":
						inst.RX = GongExtractFloat(rhs)
					case "RY":
						inst.RY = GongExtractFloat(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					}
				case *FileToDownload:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Base64EncodedContent":
						inst.Base64EncodedContent = GongExtractString(rhs)
					}
				case *Layer:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Rects":
						__gong__assignSliceOfPointers(&inst.Rects, rhs, identifierMap)
					case "Texts":
						__gong__assignSliceOfPointers(&inst.Texts, rhs, identifierMap)
					case "Circles":
						__gong__assignSliceOfPointers(&inst.Circles, rhs, identifierMap)
					case "Lines":
						__gong__assignSliceOfPointers(&inst.Lines, rhs, identifierMap)
					case "Ellipses":
						__gong__assignSliceOfPointers(&inst.Ellipses, rhs, identifierMap)
					case "Polylines":
						__gong__assignSliceOfPointers(&inst.Polylines, rhs, identifierMap)
					case "Polygones":
						__gong__assignSliceOfPointers(&inst.Polygones, rhs, identifierMap)
					case "Paths":
						__gong__assignSliceOfPointers(&inst.Paths, rhs, identifierMap)
					case "Links":
						__gong__assignSliceOfPointers(&inst.Links, rhs, identifierMap)
					case "RectLinkLinks":
						__gong__assignSliceOfPointers(&inst.RectLinkLinks, rhs, identifierMap)
					}
				case *Line:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X1":
						inst.X1 = GongExtractFloat(rhs)
					case "Y1":
						inst.Y1 = GongExtractFloat(rhs)
					case "X2":
						inst.X2 = GongExtractFloat(rhs)
					case "Y2":
						inst.Y2 = GongExtractFloat(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					case "MouseClickX":
						inst.MouseClickX = GongExtractFloat(rhs)
					case "MouseClickY":
						inst.MouseClickY = GongExtractFloat(rhs)
					}
				case *Link:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = LinkType(GongExtractString(rhs))
					case "IsBezierCurve":
						inst.IsBezierCurve = GongExtractBool(rhs)
					case "Start":
						__gong__assignPointer(&inst.Start, rhs, identifierMap)
					case "StartAnchorType":
						inst.StartAnchorType = AnchorType(GongExtractString(rhs))
					case "End":
						__gong__assignPointer(&inst.End, rhs, identifierMap)
					case "EndAnchorType":
						inst.EndAnchorType = AnchorType(GongExtractString(rhs))
					case "StartOrientation":
						inst.StartOrientation = OrientationType(GongExtractString(rhs))
					case "StartRatio":
						inst.StartRatio = GongExtractFloat(rhs)
					case "EndOrientation":
						inst.EndOrientation = OrientationType(GongExtractString(rhs))
					case "EndRatio":
						inst.EndRatio = GongExtractFloat(rhs)
					case "CornerOffsetRatio":
						inst.CornerOffsetRatio = GongExtractFloat(rhs)
					case "CornerRadius":
						inst.CornerRadius = GongExtractFloat(rhs)
					case "HasEndArrow":
						inst.HasEndArrow = GongExtractBool(rhs)
					case "EndArrowSize":
						inst.EndArrowSize = GongExtractFloat(rhs)
					case "EndArrowOffset":
						inst.EndArrowOffset = GongExtractFloat(rhs)
					case "HasStartArrow":
						inst.HasStartArrow = GongExtractBool(rhs)
					case "StartArrowSize":
						inst.StartArrowSize = GongExtractFloat(rhs)
					case "StartArrowOffset":
						inst.StartArrowOffset = GongExtractFloat(rhs)
					case "TextAtArrowStart":
						__gong__assignSliceOfPointers(&inst.TextAtArrowStart, rhs, identifierMap)
					case "TextAtArrowEnd":
						__gong__assignSliceOfPointers(&inst.TextAtArrowEnd, rhs, identifierMap)
					case "TextAtCorner":
						__gong__assignSliceOfPointers(&inst.TextAtCorner, rhs, identifierMap)
					case "PathAtArrowStart":
						__gong__assignSliceOfPointers(&inst.PathAtArrowStart, rhs, identifierMap)
					case "PathAtArrowEnd":
						__gong__assignSliceOfPointers(&inst.PathAtArrowEnd, rhs, identifierMap)
					case "PathAtCorner":
						__gong__assignSliceOfPointers(&inst.PathAtCorner, rhs, identifierMap)
					case "ControlPoints":
						__gong__assignSliceOfPointers(&inst.ControlPoints, rhs, identifierMap)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "MouseX":
						inst.MouseX = GongExtractFloat(rhs)
					case "MouseY":
						inst.MouseY = GongExtractFloat(rhs)
					case "MouseEventKey":
						inst.MouseEventKey = MouseEventKey(GongExtractString(rhs))
					}
				case *LinkAnchoredPath:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Definition":
						inst.Definition = GongExtractString(rhs)
					case "X_Offset":
						inst.X_Offset = GongExtractFloat(rhs)
					case "Y_Offset":
						inst.Y_Offset = GongExtractFloat(rhs)
					case "ScalePropotionnally":
						inst.ScalePropotionnally = GongExtractBool(rhs)
					case "AppliedScaling":
						inst.AppliedScaling = GongExtractFloat(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					}
				case *LinkAnchoredText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "AutomaticLayout":
						inst.AutomaticLayout = GongExtractBool(rhs)
					case "LinkAnchorType":
						inst.LinkAnchorType = LinkAnchorType(GongExtractString(rhs))
					case "X_Offset":
						inst.X_Offset = GongExtractFloat(rhs)
					case "Y_Offset":
						inst.Y_Offset = GongExtractFloat(rhs)
					case "FontWeight":
						inst.FontWeight = GongExtractString(rhs)
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = GongExtractString(rhs)
					case "LetterSpacing":
						inst.LetterSpacing = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "WhiteSpace":
						inst.WhiteSpace = WhiteSpaceEnum(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					}
				case *Path:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Definition":
						inst.Definition = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					}
				case *Point:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					}
				case *Polygone:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Points":
						inst.Points = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					}
				case *Polyline:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Points":
						inst.Points = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					}
				case *Rect:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "RX":
						inst.RX = GongExtractFloat(rhs)
					case "Peers":
						__gong__assignSliceOfPointers(&inst.Peers, rhs, identifierMap)
					case "EnclosingRect":
						__gong__assignPointer(&inst.EnclosingRect, rhs, identifierMap)
					case "Obstacles":
						__gong__assignSliceOfPointers(&inst.Obstacles, rhs, identifierMap)
					case "AnchoredTo":
						__gong__assignPointer(&inst.AnchoredTo, rhs, identifierMap)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "HoveringTrigger":
						__gong__assignSliceOfPointers(&inst.HoveringTrigger, rhs, identifierMap)
					case "DisplayConditions":
						__gong__assignSliceOfPointers(&inst.DisplayConditions, rhs, identifierMap)
					case "Animations":
						__gong__assignSliceOfPointers(&inst.Animations, rhs, identifierMap)
					case "IsSelectable":
						inst.IsSelectable = GongExtractBool(rhs)
					case "IsSelected":
						inst.IsSelected = GongExtractBool(rhs)
					case "CanHaveLeftHandle":
						inst.CanHaveLeftHandle = GongExtractBool(rhs)
					case "HasLeftHandle":
						inst.HasLeftHandle = GongExtractBool(rhs)
					case "CanHaveRightHandle":
						inst.CanHaveRightHandle = GongExtractBool(rhs)
					case "HasRightHandle":
						inst.HasRightHandle = GongExtractBool(rhs)
					case "CanHaveTopHandle":
						inst.CanHaveTopHandle = GongExtractBool(rhs)
					case "HasTopHandle":
						inst.HasTopHandle = GongExtractBool(rhs)
					case "IsScalingProportionally":
						inst.IsScalingProportionally = GongExtractBool(rhs)
					case "CanHaveBottomHandle":
						inst.CanHaveBottomHandle = GongExtractBool(rhs)
					case "HasBottomHandle":
						inst.HasBottomHandle = GongExtractBool(rhs)
					case "CanMoveHorizontaly":
						inst.CanMoveHorizontaly = GongExtractBool(rhs)
					case "CanMoveVerticaly":
						inst.CanMoveVerticaly = GongExtractBool(rhs)
					case "RectAnchoredTexts":
						__gong__assignSliceOfPointers(&inst.RectAnchoredTexts, rhs, identifierMap)
					case "RectAnchoredRects":
						__gong__assignSliceOfPointers(&inst.RectAnchoredRects, rhs, identifierMap)
					case "RectAnchoredPaths":
						__gong__assignSliceOfPointers(&inst.RectAnchoredPaths, rhs, identifierMap)
					case "RectAnchoredPngImages":
						__gong__assignSliceOfPointers(&inst.RectAnchoredPngImages, rhs, identifierMap)
					case "ChangeColorWhenHovered":
						inst.ChangeColorWhenHovered = GongExtractBool(rhs)
					case "ColorWhenHovered":
						inst.ColorWhenHovered = GongExtractString(rhs)
					case "OriginalColor":
						inst.OriginalColor = GongExtractString(rhs)
					case "FillOpacityWhenHovered":
						inst.FillOpacityWhenHovered = GongExtractFloat(rhs)
					case "OriginalFillOpacity":
						inst.OriginalFillOpacity = GongExtractFloat(rhs)
					case "HasToolTip":
						inst.HasToolTip = GongExtractBool(rhs)
					case "ToolTipText":
						inst.ToolTipText = GongExtractString(rhs)
					case "ToolTipPosition":
						inst.ToolTipPosition = ToolTipPositionEnum(GongExtractString(rhs))
					case "MouseX":
						inst.MouseX = GongExtractFloat(rhs)
					case "MouseY":
						inst.MouseY = GongExtractFloat(rhs)
					case "MouseEventKey":
						inst.MouseEventKey = MouseEventKey(GongExtractString(rhs))
					case "URLPath":
						inst.URLPath = GongExtractString(rhs)
					case "URLTarget":
						inst.URLTarget = LinkTargetType(GongExtractString(rhs))
					}
				case *RectAnchoredPath:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Definition":
						inst.Definition = GongExtractString(rhs)
					case "X_Offset":
						inst.X_Offset = GongExtractFloat(rhs)
					case "Y_Offset":
						inst.Y_Offset = GongExtractFloat(rhs)
					case "RectAnchorType":
						inst.RectAnchorType = RectAnchorType(GongExtractString(rhs))
					case "ScalePropotionnally":
						inst.ScalePropotionnally = GongExtractBool(rhs)
					case "AppliedScaling":
						inst.AppliedScaling = GongExtractFloat(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					}
				case *RectAnchoredPngImage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "RX":
						inst.RX = GongExtractFloat(rhs)
					case "X_Offset":
						inst.X_Offset = GongExtractFloat(rhs)
					case "Y_Offset":
						inst.Y_Offset = GongExtractFloat(rhs)
					case "RectAnchorType":
						inst.RectAnchorType = RectAnchorType(GongExtractString(rhs))
					case "Base64Content":
						inst.Base64Content = GongExtractString(rhs)
					}
				case *RectAnchoredRect:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "RX":
						inst.RX = GongExtractFloat(rhs)
					case "X_Offset":
						inst.X_Offset = GongExtractFloat(rhs)
					case "Y_Offset":
						inst.Y_Offset = GongExtractFloat(rhs)
					case "RectAnchorType":
						inst.RectAnchorType = RectAnchorType(GongExtractString(rhs))
					case "WidthFollowRect":
						inst.WidthFollowRect = GongExtractBool(rhs)
					case "HeightFollowRect":
						inst.HeightFollowRect = GongExtractBool(rhs)
					case "HasToolTip":
						inst.HasToolTip = GongExtractBool(rhs)
					case "ToolTipText":
						inst.ToolTipText = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					}
				case *RectAnchoredText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = GongExtractString(rhs)
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = GongExtractString(rhs)
					case "LetterSpacing":
						inst.LetterSpacing = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "WhiteSpace":
						inst.WhiteSpace = WhiteSpaceEnum(GongExtractString(rhs))
					case "X_Offset":
						inst.X_Offset = GongExtractFloat(rhs)
					case "Y_Offset":
						inst.Y_Offset = GongExtractFloat(rhs)
					case "RectAnchorType":
						inst.RectAnchorType = RectAnchorType(GongExtractString(rhs))
					case "TextAnchorType":
						inst.TextAnchorType = TextAnchorType(GongExtractString(rhs))
					case "DominantBaseline":
						inst.DominantBaseline = DominantBaselineType(GongExtractString(rhs))
					case "WritingMode":
						inst.WritingMode = WritingMode(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
					case "URLPath":
						inst.URLPath = GongExtractString(rhs)
					case "URLTarget":
						inst.URLTarget = LinkTargetType(GongExtractString(rhs))
					}
				case *RectLinkLink:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Start":
						__gong__assignPointer(&inst.Start, rhs, identifierMap)
					case "End":
						__gong__assignPointer(&inst.End, rhs, identifierMap)
					case "TargetAnchorPosition":
						inst.TargetAnchorPosition = GongExtractFloat(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					}
				case *SVG:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Layers":
						__gong__assignSliceOfPointers(&inst.Layers, rhs, identifierMap)
					case "DrawingState":
						inst.DrawingState = DrawingState(GongExtractString(rhs))
					case "StartRect":
						__gong__assignPointer(&inst.StartRect, rhs, identifierMap)
					case "EndRect":
						__gong__assignPointer(&inst.EndRect, rhs, identifierMap)
					case "IsEditable":
						inst.IsEditable = GongExtractBool(rhs)
					case "IsSVGFrontEndFileGenerated":
						inst.IsSVGFrontEndFileGenerated = GongExtractBool(rhs)
					case "IsSVGBackEndFileGenerated":
						inst.IsSVGBackEndFileGenerated = GongExtractBool(rhs)
					case "DefaultDirectoryForGeneratedImages":
						inst.DefaultDirectoryForGeneratedImages = GongExtractString(rhs)
					case "IsControlBannerHidden":
						inst.IsControlBannerHidden = GongExtractBool(rhs)
					case "PanX":
						inst.PanX = GongExtractFloat(rhs)
					case "PanY":
						inst.PanY = GongExtractFloat(rhs)
					case "Zoom":
						inst.Zoom = GongExtractFloat(rhs)
					case "OverrideWidth":
						inst.OverrideWidth = GongExtractBool(rhs)
					case "OverriddenWidth":
						inst.OverriddenWidth = GongExtractFloat(rhs)
					case "OverrideHeight":
						inst.OverrideHeight = GongExtractBool(rhs)
					case "OverriddenHeight":
						inst.OverriddenHeight = GongExtractFloat(rhs)
					}
				case *SvgText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Text":
						inst.Text = GongExtractString(rhs)
					}
				case *Text:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Content":
						inst.Content = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "FillOpacity":
						inst.FillOpacity = GongExtractFloat(rhs)
					case "Stroke":
						inst.Stroke = GongExtractString(rhs)
					case "StrokeOpacity":
						inst.StrokeOpacity = GongExtractFloat(rhs)
					case "StrokeWidth":
						inst.StrokeWidth = GongExtractFloat(rhs)
					case "StrokeDashArray":
						inst.StrokeDashArray = GongExtractString(rhs)
					case "StrokeDashArrayWhenSelected":
						inst.StrokeDashArrayWhenSelected = GongExtractString(rhs)
					case "Transform":
						inst.Transform = GongExtractString(rhs)
					case "FontWeight":
						inst.FontWeight = GongExtractString(rhs)
					case "FontSize":
						inst.FontSize = GongExtractString(rhs)
					case "FontStyle":
						inst.FontStyle = GongExtractString(rhs)
					case "LetterSpacing":
						inst.LetterSpacing = GongExtractString(rhs)
					case "FontFamily":
						inst.FontFamily = GongExtractString(rhs)
					case "WhiteSpace":
						inst.WhiteSpace = WhiteSpaceEnum(GongExtractString(rhs))
					case "Animates":
						__gong__assignSliceOfPointers(&inst.Animates, rhs, identifierMap)
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
