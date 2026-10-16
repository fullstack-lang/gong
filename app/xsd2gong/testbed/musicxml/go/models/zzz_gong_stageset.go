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
	"runtime/debug"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/score"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/barline"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
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
	ScoreStage *score.Stage
	NoteStage *note.Stage
	DirectionStage *direction.Stage
	BarlineStage *barline.Stage
	AttributesStage *attributes.Stage
	CommonStage *common.Stage
	XlinkStage *xlink.Stage
}


// Commit commits all stages in StageSet in dependency order
func (stageSet *StageSet) Commit() {
	if stageSet.XlinkStage != nil {
		stageSet.XlinkStage.Commit()
	}
	if stageSet.CommonStage != nil {
		stageSet.CommonStage.Commit()
	}
	if stageSet.AttributesStage != nil {
		stageSet.AttributesStage.Commit()
	}
	if stageSet.BarlineStage != nil {
		stageSet.BarlineStage.Commit()
	}
	if stageSet.DirectionStage != nil {
		stageSet.DirectionStage.Commit()
	}
	if stageSet.NoteStage != nil {
		stageSet.NoteStage.Commit()
	}
	if stageSet.ScoreStage != nil {
		stageSet.ScoreStage.Commit()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.Commit()
	}
}

// Checkout checkouts all stages in StageSet
func (stageSet *StageSet) Checkout() {
	if stageSet.XlinkStage != nil {
		stageSet.XlinkStage.Checkout()
	}
	if stageSet.CommonStage != nil {
		stageSet.CommonStage.Checkout()
	}
	if stageSet.AttributesStage != nil {
		stageSet.AttributesStage.Checkout()
	}
	if stageSet.BarlineStage != nil {
		stageSet.BarlineStage.Checkout()
	}
	if stageSet.DirectionStage != nil {
		stageSet.DirectionStage.Checkout()
	}
	if stageSet.NoteStage != nil {
		stageSet.NoteStage.Checkout()
	}
	if stageSet.ScoreStage != nil {
		stageSet.ScoreStage.Checkout()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.Checkout()
	}
}

// Reset resets all stages in StageSet
func (stageSet *StageSet) Reset() {
	if stageSet.XlinkStage != nil {
		stageSet.XlinkStage.Reset()
	}
	if stageSet.CommonStage != nil {
		stageSet.CommonStage.Reset()
	}
	if stageSet.AttributesStage != nil {
		stageSet.AttributesStage.Reset()
	}
	if stageSet.BarlineStage != nil {
		stageSet.BarlineStage.Reset()
	}
	if stageSet.DirectionStage != nil {
		stageSet.DirectionStage.Reset()
	}
	if stageSet.NoteStage != nil {
		stageSet.NoteStage.Reset()
	}
	if stageSet.ScoreStage != nil {
		stageSet.ScoreStage.Reset()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.Reset()
	}
}

// Clean cleans all stages in StageSet in dependency order
func (stageSet *StageSet) Clean() {
	if stageSet.XlinkStage != nil {
		stageSet.XlinkStage.Clean()
	}
	if stageSet.CommonStage != nil {
		stageSet.CommonStage.Clean()
	}
	if stageSet.AttributesStage != nil {
		stageSet.AttributesStage.Clean()
	}
	if stageSet.BarlineStage != nil {
		stageSet.BarlineStage.Clean()
	}
	if stageSet.DirectionStage != nil {
		stageSet.DirectionStage.Clean()
	}
	if stageSet.NoteStage != nil {
		stageSet.NoteStage.Clean()
	}
	if stageSet.ScoreStage != nil {
		stageSet.ScoreStage.Clean()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.Clean()
	}
}

// ComputeReverseMaps computes reverse maps on all stages in StageSet
func (stageSet *StageSet) ComputeReverseMaps() {
	if stageSet.XlinkStage != nil {
		stageSet.XlinkStage.ComputeReverseMaps()
	}
	if stageSet.CommonStage != nil {
		stageSet.CommonStage.ComputeReverseMaps()
	}
	if stageSet.AttributesStage != nil {
		stageSet.AttributesStage.ComputeReverseMaps()
	}
	if stageSet.BarlineStage != nil {
		stageSet.BarlineStage.ComputeReverseMaps()
	}
	if stageSet.DirectionStage != nil {
		stageSet.DirectionStage.ComputeReverseMaps()
	}
	if stageSet.NoteStage != nil {
		stageSet.NoteStage.ComputeReverseMaps()
	}
	if stageSet.ScoreStage != nil {
		stageSet.ScoreStage.ComputeReverseMaps()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReverseMaps()
	}
}

// ComputeInstancesNb computes instances nb on all stages in StageSet
func (stageSet *StageSet) ComputeInstancesNb() {
	if stageSet.XlinkStage != nil {
		stageSet.XlinkStage.ComputeInstancesNb()
	}
	if stageSet.CommonStage != nil {
		stageSet.CommonStage.ComputeInstancesNb()
	}
	if stageSet.AttributesStage != nil {
		stageSet.AttributesStage.ComputeInstancesNb()
	}
	if stageSet.BarlineStage != nil {
		stageSet.BarlineStage.ComputeInstancesNb()
	}
	if stageSet.DirectionStage != nil {
		stageSet.DirectionStage.ComputeInstancesNb()
	}
	if stageSet.NoteStage != nil {
		stageSet.NoteStage.ComputeInstancesNb()
	}
	if stageSet.ScoreStage != nil {
		stageSet.ScoreStage.ComputeInstancesNb()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeInstancesNb()
	}
}

// ComputeReferenceAndOrders computes reference and orders on all stages in StageSet
func (stageSet *StageSet) ComputeReferenceAndOrders() {
	if stageSet.XlinkStage != nil {
		stageSet.XlinkStage.ComputeReferenceAndOrders()
	}
	if stageSet.CommonStage != nil {
		stageSet.CommonStage.ComputeReferenceAndOrders()
	}
	if stageSet.AttributesStage != nil {
		stageSet.AttributesStage.ComputeReferenceAndOrders()
	}
	if stageSet.BarlineStage != nil {
		stageSet.BarlineStage.ComputeReferenceAndOrders()
	}
	if stageSet.DirectionStage != nil {
		stageSet.DirectionStage.ComputeReferenceAndOrders()
	}
	if stageSet.NoteStage != nil {
		stageSet.NoteStage.ComputeReferenceAndOrders()
	}
	if stageSet.ScoreStage != nil {
		stageSet.ScoreStage.ComputeReferenceAndOrders()
	}
	if stageSet.Stage != nil {
		stageSet.Stage.ComputeReferenceAndOrders()
	}
}

// NewStageSet creates a StageSet with all stages initialized
func NewStageSet(path string) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = NewStage(path)
	subPath_ScoreStage := "score"
	if path != "" {
	subPath_ScoreStage = path + "_score"
	}
	stageSet.ScoreStage = score.NewStage(subPath_ScoreStage)
	subPath_NoteStage := "note"
	if path != "" {
	subPath_NoteStage = path + "_note"
	}
	stageSet.NoteStage = note.NewStage(subPath_NoteStage)
	subPath_DirectionStage := "direction"
	if path != "" {
	subPath_DirectionStage = path + "_direction"
	}
	stageSet.DirectionStage = direction.NewStage(subPath_DirectionStage)
	subPath_BarlineStage := "barline"
	if path != "" {
	subPath_BarlineStage = path + "_barline"
	}
	stageSet.BarlineStage = barline.NewStage(subPath_BarlineStage)
	subPath_AttributesStage := "attributes"
	if path != "" {
	subPath_AttributesStage = path + "_attributes"
	}
	stageSet.AttributesStage = attributes.NewStage(subPath_AttributesStage)
	subPath_CommonStage := "common"
	if path != "" {
	subPath_CommonStage = path + "_common"
	}
	stageSet.CommonStage = common.NewStage(subPath_CommonStage)
	subPath_XlinkStage := "xlink"
	if path != "" {
	subPath_XlinkStage = path + "_xlink"
	}
	stageSet.XlinkStage = xlink.NewStage(subPath_XlinkStage)
	return stageSet
}

// NewStageSetFromStage creates a StageSet using an existing root stage
func NewStageSetFromStage(stage *Stage) (stageSet *StageSet) {
	stageSet = new(StageSet)
	stageSet.Stage = stage
	subPath_ScoreStage := "score"
	if stage != nil && stage.GetName() != "" {
	subPath_ScoreStage = stage.GetName() + "_score"
	}
	stageSet.ScoreStage = score.NewStage(subPath_ScoreStage)
	subPath_NoteStage := "note"
	if stage != nil && stage.GetName() != "" {
	subPath_NoteStage = stage.GetName() + "_note"
	}
	stageSet.NoteStage = note.NewStage(subPath_NoteStage)
	subPath_DirectionStage := "direction"
	if stage != nil && stage.GetName() != "" {
	subPath_DirectionStage = stage.GetName() + "_direction"
	}
	stageSet.DirectionStage = direction.NewStage(subPath_DirectionStage)
	subPath_BarlineStage := "barline"
	if stage != nil && stage.GetName() != "" {
	subPath_BarlineStage = stage.GetName() + "_barline"
	}
	stageSet.BarlineStage = barline.NewStage(subPath_BarlineStage)
	subPath_AttributesStage := "attributes"
	if stage != nil && stage.GetName() != "" {
	subPath_AttributesStage = stage.GetName() + "_attributes"
	}
	stageSet.AttributesStage = attributes.NewStage(subPath_AttributesStage)
	subPath_CommonStage := "common"
	if stage != nil && stage.GetName() != "" {
	subPath_CommonStage = stage.GetName() + "_common"
	}
	stageSet.CommonStage = common.NewStage(subPath_CommonStage)
	subPath_XlinkStage := "xlink"
	if stage != nil && stage.GetName() != "" {
	subPath_XlinkStage = stage.GetName() + "_xlink"
	}
	stageSet.XlinkStage = xlink.NewStage(subPath_XlinkStage)
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

	if stageSet.XlinkStage != nil {
		for _, arctype := range __gong__sortStageSetInstances(stageSet.XlinkStage.ArcTypes, stageSet.XlinkStage.ArcType_stagedOrder) {
			if lastStageDecl != "XlinkStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "XlinkStage"
			}
			arctypeIdent := "__xlink" + arctype.GongGetIdentifier(stageSet.XlinkStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.ArcType{Name: %s}).Stage(stageSet.XlinkStage)", arctypeIdent, __gong__toRawStringLiteral(arctype.Name)))
			if lastStageVal != "XlinkStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "XlinkStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", arctypeIdent, __gong__toRawStringLiteral(string(arctype.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", arctypeIdent, __gong__toRawStringLiteral(string(arctype.Show))))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", arctypeIdent, __gong__toRawStringLiteral(string(arctype.Actuate))))
			values.WriteString(fmt.Sprintf("\n\t%s.From = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.From)))
			values.WriteString(fmt.Sprintf("\n\t%s.To = %s", arctypeIdent, __gong__toRawStringLiteral(arctype.To)))
			for _, elem := range arctype.Title_1 {
				if lastStagePtr != "XlinkStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "XlinkStage"
				}
				targetIdent := "__xlink" + elem.GongGetIdentifier(stageSet.XlinkStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Title_1 = append(%s.Title_1, %s)", arctypeIdent, arctypeIdent, targetIdent))
			}
		}
	}
	if stageSet.XlinkStage != nil {
		for _, extended := range __gong__sortStageSetInstances(stageSet.XlinkStage.Extendeds, stageSet.XlinkStage.Extended_stagedOrder) {
			if lastStageDecl != "XlinkStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "XlinkStage"
			}
			extendedIdent := "__xlink" + extended.GongGetIdentifier(stageSet.XlinkStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.Extended{Name: %s}).Stage(stageSet.XlinkStage)", extendedIdent, __gong__toRawStringLiteral(extended.Name)))
			if lastStageVal != "XlinkStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "XlinkStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", extendedIdent, __gong__toRawStringLiteral(extended.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", extendedIdent, __gong__toRawStringLiteral(string(extended.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", extendedIdent, __gong__toRawStringLiteral(extended.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", extendedIdent, __gong__toRawStringLiteral(extended.Title)))
			for _, elem := range extended.Title_1 {
				if lastStagePtr != "XlinkStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "XlinkStage"
				}
				targetIdent := "__xlink" + elem.GongGetIdentifier(stageSet.XlinkStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Title_1 = append(%s.Title_1, %s)", extendedIdent, extendedIdent, targetIdent))
			}
			for _, elem := range extended.Resource {
				if lastStagePtr != "XlinkStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "XlinkStage"
				}
				targetIdent := "__xlink" + elem.GongGetIdentifier(stageSet.XlinkStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Resource = append(%s.Resource, %s)", extendedIdent, extendedIdent, targetIdent))
			}
			for _, elem := range extended.Locator {
				if lastStagePtr != "XlinkStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "XlinkStage"
				}
				targetIdent := "__xlink" + elem.GongGetIdentifier(stageSet.XlinkStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Locator = append(%s.Locator, %s)", extendedIdent, extendedIdent, targetIdent))
			}
			for _, elem := range extended.Arc {
				if lastStagePtr != "XlinkStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "XlinkStage"
				}
				targetIdent := "__xlink" + elem.GongGetIdentifier(stageSet.XlinkStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arc = append(%s.Arc, %s)", extendedIdent, extendedIdent, targetIdent))
			}
		}
	}
	if stageSet.XlinkStage != nil {
		for _, locatortype := range __gong__sortStageSetInstances(stageSet.XlinkStage.LocatorTypes, stageSet.XlinkStage.LocatorType_stagedOrder) {
			if lastStageDecl != "XlinkStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "XlinkStage"
			}
			locatortypeIdent := "__xlink" + locatortype.GongGetIdentifier(stageSet.XlinkStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.LocatorType{Name: %s}).Stage(stageSet.XlinkStage)", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Name)))
			if lastStageVal != "XlinkStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "XlinkStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", locatortypeIdent, __gong__toRawStringLiteral(string(locatortype.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", locatortypeIdent, __gong__toRawStringLiteral(locatortype.Label)))
			for _, elem := range locatortype.Title_1 {
				if lastStagePtr != "XlinkStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "XlinkStage"
				}
				targetIdent := "__xlink" + elem.GongGetIdentifier(stageSet.XlinkStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Title_1 = append(%s.Title_1, %s)", locatortypeIdent, locatortypeIdent, targetIdent))
			}
		}
	}
	if stageSet.XlinkStage != nil {
		for _, resourcetype := range __gong__sortStageSetInstances(stageSet.XlinkStage.ResourceTypes, stageSet.XlinkStage.ResourceType_stagedOrder) {
			if lastStageDecl != "XlinkStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "XlinkStage"
			}
			resourcetypeIdent := "__xlink" + resourcetype.GongGetIdentifier(stageSet.XlinkStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.ResourceType{Name: %s}).Stage(stageSet.XlinkStage)", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Name)))
			if lastStageVal != "XlinkStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "XlinkStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", resourcetypeIdent, __gong__toRawStringLiteral(string(resourcetype.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Label = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Label)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", resourcetypeIdent, __gong__toRawStringLiteral(resourcetype.Value)))
		}
	}
	if stageSet.XlinkStage != nil {
		for _, simple := range __gong__sortStageSetInstances(stageSet.XlinkStage.Simples, stageSet.XlinkStage.Simple_stagedOrder) {
			if lastStageDecl != "XlinkStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "XlinkStage"
			}
			simpleIdent := "__xlink" + simple.GongGetIdentifier(stageSet.XlinkStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.Simple{Name: %s}).Stage(stageSet.XlinkStage)", simpleIdent, __gong__toRawStringLiteral(simple.Name)))
			if lastStageVal != "XlinkStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "XlinkStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", simpleIdent, __gong__toRawStringLiteral(simple.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", simpleIdent, __gong__toRawStringLiteral(string(simple.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", simpleIdent, __gong__toRawStringLiteral(simple.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", simpleIdent, __gong__toRawStringLiteral(simple.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", simpleIdent, __gong__toRawStringLiteral(simple.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", simpleIdent, __gong__toRawStringLiteral(simple.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", simpleIdent, __gong__toRawStringLiteral(string(simple.Show))))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", simpleIdent, __gong__toRawStringLiteral(string(simple.Actuate))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", simpleIdent, __gong__toRawStringLiteral(simple.Value)))
		}
	}
	if stageSet.XlinkStage != nil {
		for _, titleelttype := range __gong__sortStageSetInstances(stageSet.XlinkStage.TitleEltTypes, stageSet.XlinkStage.TitleEltType_stagedOrder) {
			if lastStageDecl != "XlinkStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "XlinkStage"
			}
			titleelttypeIdent := "__xlink" + titleelttype.GongGetIdentifier(stageSet.XlinkStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&xlink.TitleEltType{Name: %s}).Stage(stageSet.XlinkStage)", titleelttypeIdent, __gong__toRawStringLiteral(titleelttype.Name)))
			if lastStageVal != "XlinkStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "XlinkStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", titleelttypeIdent, __gong__toRawStringLiteral(titleelttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", titleelttypeIdent, __gong__toRawStringLiteral(string(titleelttype.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", titleelttypeIdent, __gong__toRawStringLiteral(titleelttype.Lang)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", titleelttypeIdent, __gong__toRawStringLiteral(titleelttype.Value)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, accidentaltext := range __gong__sortStageSetInstances(stageSet.CommonStage.AccidentalTexts, stageSet.CommonStage.AccidentalText_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			accidentaltextIdent := "__common" + accidentaltext.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.AccidentalText{Name: %s}).Stage(stageSet.CommonStage)", accidentaltextIdent, __gong__toRawStringLiteral(accidentaltext.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, appearance := range __gong__sortStageSetInstances(stageSet.CommonStage.Appearances, stageSet.CommonStage.Appearance_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			appearanceIdent := "__common" + appearance.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Appearance{Name: %s}).Stage(stageSet.CommonStage)", appearanceIdent, __gong__toRawStringLiteral(appearance.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", appearanceIdent, __gong__toRawStringLiteral(appearance.Name)))
			for _, elem := range appearance.LineWidth {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LineWidth = append(%s.LineWidth, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.NoteSize {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteSize = append(%s.NoteSize, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.Distance {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Distance = append(%s.Distance, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.Glyph {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Glyph = append(%s.Glyph, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
			for _, elem := range appearance.OtherAppearance {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherAppearance = append(%s.OtherAppearance, %s)", appearanceIdent, appearanceIdent, targetIdent))
			}
		}
	}
	if stageSet.CommonStage != nil {
		for _, bookmark := range __gong__sortStageSetInstances(stageSet.CommonStage.Bookmarks, stageSet.CommonStage.Bookmark_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			bookmarkIdent := "__common" + bookmark.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Bookmark{Name: %s}).Stage(stageSet.CommonStage)", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameAttr = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.NameAttr)))
			values.WriteString(fmt.Sprintf("\n\t%s.Element = %s", bookmarkIdent, __gong__toRawStringLiteral(bookmark.Element)))
			values.WriteString(fmt.Sprintf("\n\t%s.Position = %d", bookmarkIdent, bookmark.Position))
		}
	}
	if stageSet.CommonStage != nil {
		for _, coda := range __gong__sortStageSetInstances(stageSet.CommonStage.Codas, stageSet.CommonStage.Coda_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			codaIdent := "__common" + coda.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Coda{Name: %s}).Stage(stageSet.CommonStage)", codaIdent, __gong__toRawStringLiteral(coda.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, distance := range __gong__sortStageSetInstances(stageSet.CommonStage.Distances, stageSet.CommonStage.Distance_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			distanceIdent := "__common" + distance.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Distance{Name: %s}).Stage(stageSet.CommonStage)", distanceIdent, __gong__toRawStringLiteral(distance.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", distanceIdent, __gong__toRawStringLiteral(distance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", distanceIdent, __gong__toRawStringLiteral(distance.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", distanceIdent, __gong__toRawStringLiteral(distance.Value)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, dynamics := range __gong__sortStageSetInstances(stageSet.CommonStage.Dynamicss, stageSet.CommonStage.Dynamics_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			dynamicsIdent := "__common" + dynamics.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Dynamics{Name: %s}).Stage(stageSet.CommonStage)", dynamicsIdent, __gong__toRawStringLiteral(dynamics.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.P = append(%s.P, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Pp {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pp = append(%s.Pp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Ppp {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ppp = append(%s.Ppp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Pppp {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pppp = append(%s.Pppp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Ppppp {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ppppp = append(%s.Ppppp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Pppppp {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pppppp = append(%s.Pppppp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.F {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.F = append(%s.F, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Ff {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ff = append(%s.Ff, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Fff {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fff = append(%s.Fff, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Ffff {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ffff = append(%s.Ffff, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Fffff {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fffff = append(%s.Fffff, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Ffffff {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ffffff = append(%s.Ffffff, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Mp {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mp = append(%s.Mp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Mf {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mf = append(%s.Mf, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Sf {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sf = append(%s.Sf, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Sfp {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sfp = append(%s.Sfp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Sfpp {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sfpp = append(%s.Sfpp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Fp {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fp = append(%s.Fp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Rf {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rf = append(%s.Rf, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Rfz {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rfz = append(%s.Rfz, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Sfz {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sfz = append(%s.Sfz, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Sffz {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sffz = append(%s.Sffz, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Fz {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fz = append(%s.Fz, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.N {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.N = append(%s.N, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Pf {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pf = append(%s.Pf, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.Sfzp {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sfzp = append(%s.Sfzp, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
			for _, elem := range dynamics.OtherDynamics {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherDynamics = append(%s.OtherDynamics, %s)", dynamicsIdent, dynamicsIdent, targetIdent))
			}
		}
	}
	if stageSet.CommonStage != nil {
		for _, empty := range __gong__sortStageSetInstances(stageSet.CommonStage.Emptys, stageSet.CommonStage.Empty_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			emptyIdent := "__common" + empty.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Empty{Name: %s}).Stage(stageSet.CommonStage)", emptyIdent, __gong__toRawStringLiteral(empty.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", emptyIdent, __gong__toRawStringLiteral(empty.Name)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, emptyplacement := range __gong__sortStageSetInstances(stageSet.CommonStage.EmptyPlacements, stageSet.CommonStage.EmptyPlacement_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			emptyplacementIdent := "__common" + emptyplacement.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyPlacement{Name: %s}).Stage(stageSet.CommonStage)", emptyplacementIdent, __gong__toRawStringLiteral(emptyplacement.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, emptyplacementsmufl := range __gong__sortStageSetInstances(stageSet.CommonStage.EmptyPlacementSmufls, stageSet.CommonStage.EmptyPlacementSmufl_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			emptyplacementsmuflIdent := "__common" + emptyplacementsmufl.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyPlacementSmufl{Name: %s}).Stage(stageSet.CommonStage)", emptyplacementsmuflIdent, __gong__toRawStringLiteral(emptyplacementsmufl.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, emptyprintobjectstylealign := range __gong__sortStageSetInstances(stageSet.CommonStage.EmptyPrintObjectStyleAligns, stageSet.CommonStage.EmptyPrintObjectStyleAlign_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			emptyprintobjectstylealignIdent := "__common" + emptyprintobjectstylealign.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyPrintObjectStyleAlign{Name: %s}).Stage(stageSet.CommonStage)", emptyprintobjectstylealignIdent, __gong__toRawStringLiteral(emptyprintobjectstylealign.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, emptyprintstyle := range __gong__sortStageSetInstances(stageSet.CommonStage.EmptyPrintStyles, stageSet.CommonStage.EmptyPrintStyle_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			emptyprintstyleIdent := "__common" + emptyprintstyle.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyPrintStyle{Name: %s}).Stage(stageSet.CommonStage)", emptyprintstyleIdent, __gong__toRawStringLiteral(emptyprintstyle.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, emptyprintstylealign := range __gong__sortStageSetInstances(stageSet.CommonStage.EmptyPrintStyleAligns, stageSet.CommonStage.EmptyPrintStyleAlign_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			emptyprintstylealignIdent := "__common" + emptyprintstylealign.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyPrintStyleAlign{Name: %s}).Stage(stageSet.CommonStage)", emptyprintstylealignIdent, __gong__toRawStringLiteral(emptyprintstylealign.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, emptyprintstylealignid := range __gong__sortStageSetInstances(stageSet.CommonStage.EmptyPrintStyleAlignIds, stageSet.CommonStage.EmptyPrintStyleAlignId_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			emptyprintstylealignidIdent := "__common" + emptyprintstylealignid.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyPrintStyleAlignId{Name: %s}).Stage(stageSet.CommonStage)", emptyprintstylealignidIdent, __gong__toRawStringLiteral(emptyprintstylealignid.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, emptytrillsound := range __gong__sortStageSetInstances(stageSet.CommonStage.EmptyTrillSounds, stageSet.CommonStage.EmptyTrillSound_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			emptytrillsoundIdent := "__common" + emptytrillsound.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.EmptyTrillSound{Name: %s}).Stage(stageSet.CommonStage)", emptytrillsoundIdent, __gong__toRawStringLiteral(emptytrillsound.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, encoding := range __gong__sortStageSetInstances(stageSet.CommonStage.Encodings, stageSet.CommonStage.Encoding_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			encodingIdent := "__common" + encoding.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Encoding{Name: %s}).Stage(stageSet.CommonStage)", encodingIdent, __gong__toRawStringLiteral(encoding.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", encodingIdent, __gong__toRawStringLiteral(encoding.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.EncodingDate = %s", encodingIdent, __gong__toRawStringLiteral(encoding.EncodingDate)))
			values.WriteString(fmt.Sprintf("\n\t%s.Software = %s", encodingIdent, __gong__toRawStringLiteral(encoding.Software)))
			values.WriteString(fmt.Sprintf("\n\t%s.EncodingDescription = %s", encodingIdent, __gong__toRawStringLiteral(encoding.EncodingDescription)))
			for _, elem := range encoding.Encoder {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Encoder = append(%s.Encoder, %s)", encodingIdent, encodingIdent, targetIdent))
			}
			for _, elem := range encoding.Supports {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Supports = append(%s.Supports, %s)", encodingIdent, encodingIdent, targetIdent))
			}
		}
	}
	if stageSet.CommonStage != nil {
		for _, fermata := range __gong__sortStageSetInstances(stageSet.CommonStage.Fermatas, stageSet.CommonStage.Fermata_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			fermataIdent := "__common" + fermata.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Fermata{Name: %s}).Stage(stageSet.CommonStage)", fermataIdent, __gong__toRawStringLiteral(fermata.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, fingering := range __gong__sortStageSetInstances(stageSet.CommonStage.Fingerings, stageSet.CommonStage.Fingering_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			fingeringIdent := "__common" + fingering.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Fingering{Name: %s}).Stage(stageSet.CommonStage)", fingeringIdent, __gong__toRawStringLiteral(fingering.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, formattedsymbol := range __gong__sortStageSetInstances(stageSet.CommonStage.FormattedSymbols, stageSet.CommonStage.FormattedSymbol_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			formattedsymbolIdent := "__common" + formattedsymbol.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.FormattedSymbol{Name: %s}).Stage(stageSet.CommonStage)", formattedsymbolIdent, __gong__toRawStringLiteral(formattedsymbol.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, formattedsymbolid := range __gong__sortStageSetInstances(stageSet.CommonStage.FormattedSymbolIds, stageSet.CommonStage.FormattedSymbolId_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			formattedsymbolidIdent := "__common" + formattedsymbolid.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.FormattedSymbolId{Name: %s}).Stage(stageSet.CommonStage)", formattedsymbolidIdent, __gong__toRawStringLiteral(formattedsymbolid.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, formattedtext := range __gong__sortStageSetInstances(stageSet.CommonStage.FormattedTexts, stageSet.CommonStage.FormattedText_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			formattedtextIdent := "__common" + formattedtext.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.FormattedText{Name: %s}).Stage(stageSet.CommonStage)", formattedtextIdent, __gong__toRawStringLiteral(formattedtext.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, formattedtextid := range __gong__sortStageSetInstances(stageSet.CommonStage.FormattedTextIds, stageSet.CommonStage.FormattedTextId_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			formattedtextidIdent := "__common" + formattedtextid.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.FormattedTextId{Name: %s}).Stage(stageSet.CommonStage)", formattedtextidIdent, __gong__toRawStringLiteral(formattedtextid.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, fret := range __gong__sortStageSetInstances(stageSet.CommonStage.Frets, stageSet.CommonStage.Fret_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			fretIdent := "__common" + fret.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Fret{Name: %s}).Stage(stageSet.CommonStage)", fretIdent, __gong__toRawStringLiteral(fret.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, glyph := range __gong__sortStageSetInstances(stageSet.CommonStage.Glyphs, stageSet.CommonStage.Glyph_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			glyphIdent := "__common" + glyph.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Glyph{Name: %s}).Stage(stageSet.CommonStage)", glyphIdent, __gong__toRawStringLiteral(glyph.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glyphIdent, __gong__toRawStringLiteral(glyph.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", glyphIdent, __gong__toRawStringLiteral(glyph.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", glyphIdent, __gong__toRawStringLiteral(glyph.Value)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, horizontalturn := range __gong__sortStageSetInstances(stageSet.CommonStage.HorizontalTurns, stageSet.CommonStage.HorizontalTurn_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			horizontalturnIdent := "__common" + horizontalturn.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.HorizontalTurn{Name: %s}).Stage(stageSet.CommonStage)", horizontalturnIdent, __gong__toRawStringLiteral(horizontalturn.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, identification := range __gong__sortStageSetInstances(stageSet.CommonStage.Identifications, stageSet.CommonStage.Identification_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			identificationIdent := "__common" + identification.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Identification{Name: %s}).Stage(stageSet.CommonStage)", identificationIdent, __gong__toRawStringLiteral(identification.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", identificationIdent, __gong__toRawStringLiteral(identification.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Source = %s", identificationIdent, __gong__toRawStringLiteral(identification.Source)))
			for _, elem := range identification.Creator {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Creator = append(%s.Creator, %s)", identificationIdent, identificationIdent, targetIdent))
			}
			for _, elem := range identification.Rights {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rights = append(%s.Rights, %s)", identificationIdent, identificationIdent, targetIdent))
			}
			if identification.Encoding != nil {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + identification.Encoding.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Encoding = %s", identificationIdent, targetIdent))
			}
			for _, elem := range identification.Relation {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Relation = append(%s.Relation, %s)", identificationIdent, identificationIdent, targetIdent))
			}
			if identification.Miscellaneous != nil {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + identification.Miscellaneous.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Miscellaneous = %s", identificationIdent, targetIdent))
			}
		}
	}
	if stageSet.CommonStage != nil {
		for _, level := range __gong__sortStageSetInstances(stageSet.CommonStage.Levels, stageSet.CommonStage.Level_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			levelIdent := "__common" + level.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Level{Name: %s}).Stage(stageSet.CommonStage)", levelIdent, __gong__toRawStringLiteral(level.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, linewidth := range __gong__sortStageSetInstances(stageSet.CommonStage.LineWidths, stageSet.CommonStage.LineWidth_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			linewidthIdent := "__common" + linewidth.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.LineWidth{Name: %s}).Stage(stageSet.CommonStage)", linewidthIdent, __gong__toRawStringLiteral(linewidth.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linewidthIdent, __gong__toRawStringLiteral(linewidth.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", linewidthIdent, __gong__toRawStringLiteral(linewidth.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", linewidthIdent, __gong__toRawStringLiteral(linewidth.Value)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, link := range __gong__sortStageSetInstances(stageSet.CommonStage.Links, stageSet.CommonStage.Link_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			linkIdent := "__common" + link.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Link{Name: %s}).Stage(stageSet.CommonStage)", linkIdent, __gong__toRawStringLiteral(link.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, measurelayout := range __gong__sortStageSetInstances(stageSet.CommonStage.MeasureLayouts, stageSet.CommonStage.MeasureLayout_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			measurelayoutIdent := "__common" + measurelayout.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.MeasureLayout{Name: %s}).Stage(stageSet.CommonStage)", measurelayoutIdent, __gong__toRawStringLiteral(measurelayout.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measurelayoutIdent, __gong__toRawStringLiteral(measurelayout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MeasureDistance = %s", measurelayoutIdent, __gong__toRawStringLiteral(measurelayout.MeasureDistance)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, mididevice := range __gong__sortStageSetInstances(stageSet.CommonStage.MidiDevices, stageSet.CommonStage.MidiDevice_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			midideviceIdent := "__common" + mididevice.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.MidiDevice{Name: %s}).Stage(stageSet.CommonStage)", midideviceIdent, __gong__toRawStringLiteral(mididevice.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", midideviceIdent, __gong__toRawStringLiteral(mididevice.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Port = %d", midideviceIdent, mididevice.Port))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", midideviceIdent, __gong__toRawStringLiteral(mididevice.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", midideviceIdent, __gong__toRawStringLiteral(mididevice.Value)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, midiinstrument := range __gong__sortStageSetInstances(stageSet.CommonStage.MidiInstruments, stageSet.CommonStage.MidiInstrument_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			midiinstrumentIdent := "__common" + midiinstrument.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.MidiInstrument{Name: %s}).Stage(stageSet.CommonStage)", midiinstrumentIdent, __gong__toRawStringLiteral(midiinstrument.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, miscellaneous := range __gong__sortStageSetInstances(stageSet.CommonStage.Miscellaneouss, stageSet.CommonStage.Miscellaneous_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			miscellaneousIdent := "__common" + miscellaneous.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Miscellaneous{Name: %s}).Stage(stageSet.CommonStage)", miscellaneousIdent, __gong__toRawStringLiteral(miscellaneous.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", miscellaneousIdent, __gong__toRawStringLiteral(miscellaneous.Name)))
			for _, elem := range miscellaneous.MiscellaneousField {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MiscellaneousField = append(%s.MiscellaneousField, %s)", miscellaneousIdent, miscellaneousIdent, targetIdent))
			}
		}
	}
	if stageSet.CommonStage != nil {
		for _, miscellaneousfield := range __gong__sortStageSetInstances(stageSet.CommonStage.MiscellaneousFields, stageSet.CommonStage.MiscellaneousField_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			miscellaneousfieldIdent := "__common" + miscellaneousfield.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.MiscellaneousField{Name: %s}).Stage(stageSet.CommonStage)", miscellaneousfieldIdent, __gong__toRawStringLiteral(miscellaneousfield.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", miscellaneousfieldIdent, __gong__toRawStringLiteral(miscellaneousfield.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameAttr = %s", miscellaneousfieldIdent, __gong__toRawStringLiteral(miscellaneousfield.NameAttr)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", miscellaneousfieldIdent, __gong__toRawStringLiteral(miscellaneousfield.Value)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, namedisplay := range __gong__sortStageSetInstances(stageSet.CommonStage.NameDisplays, stageSet.CommonStage.NameDisplay_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			namedisplayIdent := "__common" + namedisplay.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.NameDisplay{Name: %s}).Stage(stageSet.CommonStage)", namedisplayIdent, __gong__toRawStringLiteral(namedisplay.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", namedisplayIdent, __gong__toRawStringLiteral(namedisplay.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", namedisplayIdent, __gong__toRawStringLiteral(string(namedisplay.PrintObject))))
			for _, elem := range namedisplay.DisplayText {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DisplayText = append(%s.DisplayText, %s)", namedisplayIdent, namedisplayIdent, targetIdent))
			}
			for _, elem := range namedisplay.AccidentalText {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AccidentalText = append(%s.AccidentalText, %s)", namedisplayIdent, namedisplayIdent, targetIdent))
			}
		}
	}
	if stageSet.CommonStage != nil {
		for _, notesize := range __gong__sortStageSetInstances(stageSet.CommonStage.NoteSizes, stageSet.CommonStage.NoteSize_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			notesizeIdent := "__common" + notesize.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.NoteSize{Name: %s}).Stage(stageSet.CommonStage)", notesizeIdent, __gong__toRawStringLiteral(notesize.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notesizeIdent, __gong__toRawStringLiteral(notesize.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", notesizeIdent, __gong__toRawStringLiteral(string(notesize.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", notesizeIdent, __gong__toRawStringLiteral(notesize.Value)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, otherappearance := range __gong__sortStageSetInstances(stageSet.CommonStage.OtherAppearances, stageSet.CommonStage.OtherAppearance_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			otherappearanceIdent := "__common" + otherappearance.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.OtherAppearance{Name: %s}).Stage(stageSet.CommonStage)", otherappearanceIdent, __gong__toRawStringLiteral(otherappearance.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", otherappearanceIdent, __gong__toRawStringLiteral(otherappearance.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", otherappearanceIdent, __gong__toRawStringLiteral(otherappearance.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", otherappearanceIdent, __gong__toRawStringLiteral(otherappearance.Value)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, otherplay := range __gong__sortStageSetInstances(stageSet.CommonStage.OtherPlays, stageSet.CommonStage.OtherPlay_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			otherplayIdent := "__common" + otherplay.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.OtherPlay{Name: %s}).Stage(stageSet.CommonStage)", otherplayIdent, __gong__toRawStringLiteral(otherplay.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", otherplayIdent, __gong__toRawStringLiteral(otherplay.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", otherplayIdent, __gong__toRawStringLiteral(otherplay.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", otherplayIdent, __gong__toRawStringLiteral(otherplay.Value)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, othertext := range __gong__sortStageSetInstances(stageSet.CommonStage.OtherTexts, stageSet.CommonStage.OtherText_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			othertextIdent := "__common" + othertext.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.OtherText{Name: %s}).Stage(stageSet.CommonStage)", othertextIdent, __gong__toRawStringLiteral(othertext.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", othertextIdent, __gong__toRawStringLiteral(othertext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", othertextIdent, __gong__toRawStringLiteral(othertext.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", othertextIdent, __gong__toRawStringLiteral(othertext.Value)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, pagelayout := range __gong__sortStageSetInstances(stageSet.CommonStage.PageLayouts, stageSet.CommonStage.PageLayout_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			pagelayoutIdent := "__common" + pagelayout.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.PageLayout{Name: %s}).Stage(stageSet.CommonStage)", pagelayoutIdent, __gong__toRawStringLiteral(pagelayout.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pagelayoutIdent, __gong__toRawStringLiteral(pagelayout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PageHeight = %s", pagelayoutIdent, __gong__toRawStringLiteral(pagelayout.PageHeight)))
			values.WriteString(fmt.Sprintf("\n\t%s.PageWidth = %s", pagelayoutIdent, __gong__toRawStringLiteral(pagelayout.PageWidth)))
			for _, elem := range pagelayout.PageMargins {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PageMargins = append(%s.PageMargins, %s)", pagelayoutIdent, pagelayoutIdent, targetIdent))
			}
		}
	}
	if stageSet.CommonStage != nil {
		for _, pagemargins := range __gong__sortStageSetInstances(stageSet.CommonStage.PageMarginss, stageSet.CommonStage.PageMargins_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			pagemarginsIdent := "__common" + pagemargins.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.PageMargins{Name: %s}).Stage(stageSet.CommonStage)", pagemarginsIdent, __gong__toRawStringLiteral(pagemargins.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pagemarginsIdent, __gong__toRawStringLiteral(pagemargins.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", pagemarginsIdent, __gong__toRawStringLiteral(string(pagemargins.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.LeftMargin = %s", pagemarginsIdent, __gong__toRawStringLiteral(pagemargins.LeftMargin)))
			values.WriteString(fmt.Sprintf("\n\t%s.RightMargin = %s", pagemarginsIdent, __gong__toRawStringLiteral(pagemargins.RightMargin)))
			values.WriteString(fmt.Sprintf("\n\t%s.TopMargin = %s", pagemarginsIdent, __gong__toRawStringLiteral(pagemargins.TopMargin)))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomMargin = %s", pagemarginsIdent, __gong__toRawStringLiteral(pagemargins.BottomMargin)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, play := range __gong__sortStageSetInstances(stageSet.CommonStage.Plays, stageSet.CommonStage.Play_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			playIdent := "__common" + play.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Play{Name: %s}).Stage(stageSet.CommonStage)", playIdent, __gong__toRawStringLiteral(play.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", playIdent, __gong__toRawStringLiteral(play.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", playIdent, __gong__toRawStringLiteral(play.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ipa = %s", playIdent, __gong__toRawStringLiteral(play.Ipa)))
			values.WriteString(fmt.Sprintf("\n\t%s.Mute = %s", playIdent, __gong__toRawStringLiteral(string(play.Mute))))
			values.WriteString(fmt.Sprintf("\n\t%s.SemiPitched = %s", playIdent, __gong__toRawStringLiteral(string(play.SemiPitched))))
			for _, elem := range play.OtherPlay {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherPlay = append(%s.OtherPlay, %s)", playIdent, playIdent, targetIdent))
			}
		}
	}
	if stageSet.CommonStage != nil {
		for _, scaling := range __gong__sortStageSetInstances(stageSet.CommonStage.Scalings, stageSet.CommonStage.Scaling_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			scalingIdent := "__common" + scaling.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Scaling{Name: %s}).Stage(stageSet.CommonStage)", scalingIdent, __gong__toRawStringLiteral(scaling.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scalingIdent, __gong__toRawStringLiteral(scaling.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Millimeters = %s", scalingIdent, __gong__toRawStringLiteral(scaling.Millimeters)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tenths = %s", scalingIdent, __gong__toRawStringLiteral(scaling.Tenths)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, segno := range __gong__sortStageSetInstances(stageSet.CommonStage.Segnos, stageSet.CommonStage.Segno_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			segnoIdent := "__common" + segno.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Segno{Name: %s}).Stage(stageSet.CommonStage)", segnoIdent, __gong__toRawStringLiteral(segno.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, stafflayout := range __gong__sortStageSetInstances(stageSet.CommonStage.StaffLayouts, stageSet.CommonStage.StaffLayout_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			stafflayoutIdent := "__common" + stafflayout.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.StaffLayout{Name: %s}).Stage(stageSet.CommonStage)", stafflayoutIdent, __gong__toRawStringLiteral(stafflayout.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stafflayoutIdent, __gong__toRawStringLiteral(stafflayout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", stafflayoutIdent, stafflayout.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.StaffDistance = %s", stafflayoutIdent, __gong__toRawStringLiteral(stafflayout.StaffDistance)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, stringtype := range __gong__sortStageSetInstances(stageSet.CommonStage.StringTypes, stageSet.CommonStage.StringType_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			stringtypeIdent := "__common" + stringtype.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.StringType{Name: %s}).Stage(stageSet.CommonStage)", stringtypeIdent, __gong__toRawStringLiteral(stringtype.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.CommonStage != nil {
		for _, supports := range __gong__sortStageSetInstances(stageSet.CommonStage.Supportss, stageSet.CommonStage.Supports_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			supportsIdent := "__common" + supports.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.Supports{Name: %s}).Stage(stageSet.CommonStage)", supportsIdent, __gong__toRawStringLiteral(supports.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", supportsIdent, __gong__toRawStringLiteral(supports.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", supportsIdent, __gong__toRawStringLiteral(string(supports.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Element = %s", supportsIdent, __gong__toRawStringLiteral(supports.Element)))
			values.WriteString(fmt.Sprintf("\n\t%s.Attribute = %s", supportsIdent, __gong__toRawStringLiteral(supports.Attribute)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", supportsIdent, __gong__toRawStringLiteral(supports.Value)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, systemdividers := range __gong__sortStageSetInstances(stageSet.CommonStage.SystemDividerss, stageSet.CommonStage.SystemDividers_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			systemdividersIdent := "__common" + systemdividers.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.SystemDividers{Name: %s}).Stage(stageSet.CommonStage)", systemdividersIdent, __gong__toRawStringLiteral(systemdividers.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", systemdividersIdent, __gong__toRawStringLiteral(systemdividers.Name)))
			if systemdividers.LeftDivider != nil {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + systemdividers.LeftDivider.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LeftDivider = %s", systemdividersIdent, targetIdent))
			}
			if systemdividers.RightDivider != nil {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + systemdividers.RightDivider.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RightDivider = %s", systemdividersIdent, targetIdent))
			}
		}
	}
	if stageSet.CommonStage != nil {
		for _, systemlayout := range __gong__sortStageSetInstances(stageSet.CommonStage.SystemLayouts, stageSet.CommonStage.SystemLayout_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			systemlayoutIdent := "__common" + systemlayout.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.SystemLayout{Name: %s}).Stage(stageSet.CommonStage)", systemlayoutIdent, __gong__toRawStringLiteral(systemlayout.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", systemlayoutIdent, __gong__toRawStringLiteral(systemlayout.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SystemDistance = %s", systemlayoutIdent, __gong__toRawStringLiteral(systemlayout.SystemDistance)))
			values.WriteString(fmt.Sprintf("\n\t%s.TopSystemDistance = %s", systemlayoutIdent, __gong__toRawStringLiteral(systemlayout.TopSystemDistance)))
			if systemlayout.SystemMargins != nil {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + systemlayout.SystemMargins.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemMargins = %s", systemlayoutIdent, targetIdent))
			}
			if systemlayout.SystemDividers != nil {
				if lastStagePtr != "CommonStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "CommonStage"
				}
				targetIdent := "__common" + systemlayout.SystemDividers.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemDividers = %s", systemlayoutIdent, targetIdent))
			}
		}
	}
	if stageSet.CommonStage != nil {
		for _, systemmargins := range __gong__sortStageSetInstances(stageSet.CommonStage.SystemMarginss, stageSet.CommonStage.SystemMargins_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			systemmarginsIdent := "__common" + systemmargins.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.SystemMargins{Name: %s}).Stage(stageSet.CommonStage)", systemmarginsIdent, __gong__toRawStringLiteral(systemmargins.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", systemmarginsIdent, __gong__toRawStringLiteral(systemmargins.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.LeftMargin = %s", systemmarginsIdent, __gong__toRawStringLiteral(systemmargins.LeftMargin)))
			values.WriteString(fmt.Sprintf("\n\t%s.RightMargin = %s", systemmarginsIdent, __gong__toRawStringLiteral(systemmargins.RightMargin)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, typedtext := range __gong__sortStageSetInstances(stageSet.CommonStage.TypedTexts, stageSet.CommonStage.TypedText_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			typedtextIdent := "__common" + typedtext.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.TypedText{Name: %s}).Stage(stageSet.CommonStage)", typedtextIdent, __gong__toRawStringLiteral(typedtext.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", typedtextIdent, __gong__toRawStringLiteral(typedtext.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", typedtextIdent, __gong__toRawStringLiteral(typedtext.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", typedtextIdent, __gong__toRawStringLiteral(typedtext.Value)))
		}
	}
	if stageSet.CommonStage != nil {
		for _, wavyline := range __gong__sortStageSetInstances(stageSet.CommonStage.WavyLines, stageSet.CommonStage.WavyLine_stagedOrder) {
			if lastStageDecl != "CommonStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "CommonStage"
			}
			wavylineIdent := "__common" + wavyline.GongGetIdentifier(stageSet.CommonStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&common.WavyLine{Name: %s}).Stage(stageSet.CommonStage)", wavylineIdent, __gong__toRawStringLiteral(wavyline.Name)))
			if lastStageVal != "CommonStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "CommonStage"
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
	if stageSet.AttributesStage != nil {
		for _, attributes := range __gong__sortStageSetInstances(stageSet.AttributesStage.Attributess, stageSet.AttributesStage.Attributes_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			attributesIdent := "__attributes" + attributes.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Attributes{Name: %s}).Stage(stageSet.AttributesStage)", attributesIdent, __gong__toRawStringLiteral(attributes.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attributesIdent, __gong__toRawStringLiteral(attributes.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Divisions = %s", attributesIdent, __gong__toRawStringLiteral(attributes.Divisions)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staves = %d", attributesIdent, attributes.Staves))
			values.WriteString(fmt.Sprintf("\n\t%s.Instruments = %d", attributesIdent, attributes.Instruments))
			if attributes.Footnote != nil {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__common" + attributes.Footnote.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", attributesIdent, targetIdent))
			}
			if attributes.Level != nil {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__common" + attributes.Level.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Key {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Key = append(%s.Key, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Time {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Time = append(%s.Time, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			if attributes.PartSymbol != nil {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + attributes.PartSymbol.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartSymbol = %s", attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Clef {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Clef = append(%s.Clef, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.StaffDetails {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StaffDetails = append(%s.StaffDetails, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Transpose {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Transpose = append(%s.Transpose, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.ForPart {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ForPart = append(%s.ForPart, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.Directive {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Directive = append(%s.Directive, %s)", attributesIdent, attributesIdent, targetIdent))
			}
			for _, elem := range attributes.MeasureStyle {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MeasureStyle = append(%s.MeasureStyle, %s)", attributesIdent, attributesIdent, targetIdent))
			}
		}
	}
	if stageSet.AttributesStage != nil {
		for _, attributesdirective := range __gong__sortStageSetInstances(stageSet.AttributesStage.AttributesDirectives, stageSet.AttributesStage.AttributesDirective_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			attributesdirectiveIdent := "__attributes" + attributesdirective.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.AttributesDirective{Name: %s}).Stage(stageSet.AttributesStage)", attributesdirectiveIdent, __gong__toRawStringLiteral(attributesdirective.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
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
	if stageSet.AttributesStage != nil {
		for _, beatrepeat := range __gong__sortStageSetInstances(stageSet.AttributesStage.BeatRepeats, stageSet.AttributesStage.BeatRepeat_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			beatrepeatIdent := "__attributes" + beatrepeat.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.BeatRepeat{Name: %s}).Stage(stageSet.AttributesStage)", beatrepeatIdent, __gong__toRawStringLiteral(beatrepeat.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", beatrepeatIdent, __gong__toRawStringLiteral(beatrepeat.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", beatrepeatIdent, __gong__toRawStringLiteral(string(beatrepeat.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Slashes = %d", beatrepeatIdent, beatrepeat.Slashes))
			values.WriteString(fmt.Sprintf("\n\t%s.UseDots = %s", beatrepeatIdent, __gong__toRawStringLiteral(string(beatrepeat.UseDots))))
			values.WriteString(fmt.Sprintf("\n\t%s.SlashType = %s", beatrepeatIdent, __gong__toRawStringLiteral(string(beatrepeat.SlashType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ExceptVoice = %s", beatrepeatIdent, __gong__toRawStringLiteral(beatrepeat.ExceptVoice)))
			for _, elem := range beatrepeat.SlashDot {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SlashDot = append(%s.SlashDot, %s)", beatrepeatIdent, beatrepeatIdent, targetIdent))
			}
		}
	}
	if stageSet.AttributesStage != nil {
		for _, cancel := range __gong__sortStageSetInstances(stageSet.AttributesStage.Cancels, stageSet.AttributesStage.Cancel_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			cancelIdent := "__attributes" + cancel.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Cancel{Name: %s}).Stage(stageSet.AttributesStage)", cancelIdent, __gong__toRawStringLiteral(cancel.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", cancelIdent, __gong__toRawStringLiteral(cancel.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", cancelIdent, __gong__toRawStringLiteral(string(cancel.Location))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", cancelIdent, cancel.Value))
		}
	}
	if stageSet.AttributesStage != nil {
		for _, clef := range __gong__sortStageSetInstances(stageSet.AttributesStage.Clefs, stageSet.AttributesStage.Clef_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			clefIdent := "__attributes" + clef.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Clef{Name: %s}).Stage(stageSet.AttributesStage)", clefIdent, __gong__toRawStringLiteral(clef.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
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
	if stageSet.AttributesStage != nil {
		for _, double := range __gong__sortStageSetInstances(stageSet.AttributesStage.Doubles, stageSet.AttributesStage.Double_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			doubleIdent := "__attributes" + double.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Double{Name: %s}).Stage(stageSet.AttributesStage)", doubleIdent, __gong__toRawStringLiteral(double.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", doubleIdent, __gong__toRawStringLiteral(double.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Above = %s", doubleIdent, __gong__toRawStringLiteral(string(double.Above))))
		}
	}
	if stageSet.AttributesStage != nil {
		for _, forpart := range __gong__sortStageSetInstances(stageSet.AttributesStage.ForParts, stageSet.AttributesStage.ForPart_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			forpartIdent := "__attributes" + forpart.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.ForPart{Name: %s}).Stage(stageSet.AttributesStage)", forpartIdent, __gong__toRawStringLiteral(forpart.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", forpartIdent, __gong__toRawStringLiteral(forpart.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", forpartIdent, forpart.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", forpartIdent, __gong__toRawStringLiteral(forpart.Id)))
			if forpart.PartClef != nil {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + forpart.PartClef.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartClef = %s", forpartIdent, targetIdent))
			}
			if forpart.PartTranspose != nil {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + forpart.PartTranspose.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartTranspose = %s", forpartIdent, targetIdent))
			}
		}
	}
	if stageSet.AttributesStage != nil {
		for _, interchangeable := range __gong__sortStageSetInstances(stageSet.AttributesStage.Interchangeables, stageSet.AttributesStage.Interchangeable_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			interchangeableIdent := "__attributes" + interchangeable.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Interchangeable{Name: %s}).Stage(stageSet.AttributesStage)", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Symbol = %s", interchangeableIdent, __gong__toRawStringLiteral(string(interchangeable.Symbol))))
			values.WriteString(fmt.Sprintf("\n\t%s.Separator = %s", interchangeableIdent, __gong__toRawStringLiteral(string(interchangeable.Separator))))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeRelation = %s", interchangeableIdent, __gong__toRawStringLiteral(string(interchangeable.TimeRelation))))
			values.WriteString(fmt.Sprintf("\n\t%s.Beats = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.Beats)))
			values.WriteString(fmt.Sprintf("\n\t%s.BeatType = %s", interchangeableIdent, __gong__toRawStringLiteral(interchangeable.BeatType)))
		}
	}
	if stageSet.AttributesStage != nil {
		for _, key := range __gong__sortStageSetInstances(stageSet.AttributesStage.Keys, stageSet.AttributesStage.Key_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			keyIdent := "__attributes" + key.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Key{Name: %s}).Stage(stageSet.AttributesStage)", keyIdent, __gong__toRawStringLiteral(key.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
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
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + key.Cancel.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Cancel = %s", keyIdent, targetIdent))
			}
			for _, elem := range key.KeyAccidental {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.KeyAccidental = append(%s.KeyAccidental, %s)", keyIdent, keyIdent, targetIdent))
			}
			for _, elem := range key.KeyOctave {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.KeyOctave = append(%s.KeyOctave, %s)", keyIdent, keyIdent, targetIdent))
			}
		}
	}
	if stageSet.AttributesStage != nil {
		for _, keyaccidental := range __gong__sortStageSetInstances(stageSet.AttributesStage.KeyAccidentals, stageSet.AttributesStage.KeyAccidental_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			keyaccidentalIdent := "__attributes" + keyaccidental.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.KeyAccidental{Name: %s}).Stage(stageSet.AttributesStage)", keyaccidentalIdent, __gong__toRawStringLiteral(keyaccidental.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", keyaccidentalIdent, __gong__toRawStringLiteral(keyaccidental.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", keyaccidentalIdent, __gong__toRawStringLiteral(keyaccidental.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", keyaccidentalIdent, __gong__toRawStringLiteral(string(keyaccidental.Value))))
		}
	}
	if stageSet.AttributesStage != nil {
		for _, keyoctave := range __gong__sortStageSetInstances(stageSet.AttributesStage.KeyOctaves, stageSet.AttributesStage.KeyOctave_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			keyoctaveIdent := "__attributes" + keyoctave.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.KeyOctave{Name: %s}).Stage(stageSet.AttributesStage)", keyoctaveIdent, __gong__toRawStringLiteral(keyoctave.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", keyoctaveIdent, __gong__toRawStringLiteral(keyoctave.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", keyoctaveIdent, keyoctave.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Cancel = %s", keyoctaveIdent, __gong__toRawStringLiteral(string(keyoctave.Cancel))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", keyoctaveIdent, keyoctave.Value))
		}
	}
	if stageSet.AttributesStage != nil {
		for _, linedetail := range __gong__sortStageSetInstances(stageSet.AttributesStage.LineDetails, stageSet.AttributesStage.LineDetail_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			linedetailIdent := "__attributes" + linedetail.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.LineDetail{Name: %s}).Stage(stageSet.AttributesStage)", linedetailIdent, __gong__toRawStringLiteral(linedetail.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linedetailIdent, __gong__toRawStringLiteral(linedetail.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", linedetailIdent, linedetail.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", linedetailIdent, __gong__toRawStringLiteral(linedetail.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", linedetailIdent, __gong__toRawStringLiteral(linedetail.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.LineType = %s", linedetailIdent, __gong__toRawStringLiteral(string(linedetail.LineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", linedetailIdent, __gong__toRawStringLiteral(string(linedetail.PrintObject))))
		}
	}
	if stageSet.AttributesStage != nil {
		for _, measurerepeat := range __gong__sortStageSetInstances(stageSet.AttributesStage.MeasureRepeats, stageSet.AttributesStage.MeasureRepeat_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			measurerepeatIdent := "__attributes" + measurerepeat.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.MeasureRepeat{Name: %s}).Stage(stageSet.AttributesStage)", measurerepeatIdent, __gong__toRawStringLiteral(measurerepeat.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measurerepeatIdent, __gong__toRawStringLiteral(measurerepeat.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", measurerepeatIdent, __gong__toRawStringLiteral(string(measurerepeat.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Slashes = %d", measurerepeatIdent, measurerepeat.Slashes))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", measurerepeatIdent, __gong__toRawStringLiteral(measurerepeat.Value)))
		}
	}
	if stageSet.AttributesStage != nil {
		for _, measurestyle := range __gong__sortStageSetInstances(stageSet.AttributesStage.MeasureStyles, stageSet.AttributesStage.MeasureStyle_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			measurestyleIdent := "__attributes" + measurestyle.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.MeasureStyle{Name: %s}).Stage(stageSet.AttributesStage)", measurestyleIdent, __gong__toRawStringLiteral(measurestyle.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
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
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + measurestyle.MultipleRest.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MultipleRest = %s", measurestyleIdent, targetIdent))
			}
			if measurestyle.MeasureRepeat != nil {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + measurestyle.MeasureRepeat.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MeasureRepeat = %s", measurestyleIdent, targetIdent))
			}
			if measurestyle.BeatRepeat != nil {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + measurestyle.BeatRepeat.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BeatRepeat = %s", measurestyleIdent, targetIdent))
			}
			if measurestyle.Slash != nil {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + measurestyle.Slash.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slash = %s", measurestyleIdent, targetIdent))
			}
		}
	}
	if stageSet.AttributesStage != nil {
		for _, multiplerest := range __gong__sortStageSetInstances(stageSet.AttributesStage.MultipleRests, stageSet.AttributesStage.MultipleRest_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			multiplerestIdent := "__attributes" + multiplerest.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.MultipleRest{Name: %s}).Stage(stageSet.AttributesStage)", multiplerestIdent, __gong__toRawStringLiteral(multiplerest.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", multiplerestIdent, __gong__toRawStringLiteral(multiplerest.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.UseSymbols = %s", multiplerestIdent, __gong__toRawStringLiteral(string(multiplerest.UseSymbols))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", multiplerestIdent, multiplerest.Value))
		}
	}
	if stageSet.AttributesStage != nil {
		for _, partclef := range __gong__sortStageSetInstances(stageSet.AttributesStage.PartClefs, stageSet.AttributesStage.PartClef_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			partclefIdent := "__attributes" + partclef.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.PartClef{Name: %s}).Stage(stageSet.AttributesStage)", partclefIdent, __gong__toRawStringLiteral(partclef.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partclefIdent, __gong__toRawStringLiteral(partclef.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sign = %s", partclefIdent, __gong__toRawStringLiteral(string(partclef.Sign))))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", partclefIdent, partclef.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.ClefOctaveChange = %d", partclefIdent, partclef.ClefOctaveChange))
		}
	}
	if stageSet.AttributesStage != nil {
		for _, partsymbol := range __gong__sortStageSetInstances(stageSet.AttributesStage.PartSymbols, stageSet.AttributesStage.PartSymbol_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			partsymbolIdent := "__attributes" + partsymbol.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.PartSymbol{Name: %s}).Stage(stageSet.AttributesStage)", partsymbolIdent, __gong__toRawStringLiteral(partsymbol.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
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
	if stageSet.AttributesStage != nil {
		for _, parttranspose := range __gong__sortStageSetInstances(stageSet.AttributesStage.PartTransposes, stageSet.AttributesStage.PartTranspose_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			parttransposeIdent := "__attributes" + parttranspose.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.PartTranspose{Name: %s}).Stage(stageSet.AttributesStage)", parttransposeIdent, __gong__toRawStringLiteral(parttranspose.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", parttransposeIdent, __gong__toRawStringLiteral(parttranspose.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Diatonic = %d", parttransposeIdent, parttranspose.Diatonic))
			values.WriteString(fmt.Sprintf("\n\t%s.Chromatic = %s", parttransposeIdent, __gong__toRawStringLiteral(parttranspose.Chromatic)))
			values.WriteString(fmt.Sprintf("\n\t%s.OctaveChange = %d", parttransposeIdent, parttranspose.OctaveChange))
			if parttranspose.Double != nil {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + parttranspose.Double.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Double = %s", parttransposeIdent, targetIdent))
			}
		}
	}
	if stageSet.AttributesStage != nil {
		for _, slash := range __gong__sortStageSetInstances(stageSet.AttributesStage.Slashs, stageSet.AttributesStage.Slash_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			slashIdent := "__attributes" + slash.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Slash{Name: %s}).Stage(stageSet.AttributesStage)", slashIdent, __gong__toRawStringLiteral(slash.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", slashIdent, __gong__toRawStringLiteral(slash.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", slashIdent, __gong__toRawStringLiteral(string(slash.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.UseDots = %s", slashIdent, __gong__toRawStringLiteral(string(slash.UseDots))))
			values.WriteString(fmt.Sprintf("\n\t%s.UseStems = %s", slashIdent, __gong__toRawStringLiteral(string(slash.UseStems))))
			values.WriteString(fmt.Sprintf("\n\t%s.SlashType = %s", slashIdent, __gong__toRawStringLiteral(string(slash.SlashType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ExceptVoice = %s", slashIdent, __gong__toRawStringLiteral(slash.ExceptVoice)))
			for _, elem := range slash.SlashDot {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SlashDot = append(%s.SlashDot, %s)", slashIdent, slashIdent, targetIdent))
			}
		}
	}
	if stageSet.AttributesStage != nil {
		for _, staffdetails := range __gong__sortStageSetInstances(stageSet.AttributesStage.StaffDetailss, stageSet.AttributesStage.StaffDetails_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			staffdetailsIdent := "__attributes" + staffdetails.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.StaffDetails{Name: %s}).Stage(stageSet.AttributesStage)", staffdetailsIdent, __gong__toRawStringLiteral(staffdetails.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
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
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LineDetail = append(%s.LineDetail, %s)", staffdetailsIdent, staffdetailsIdent, targetIdent))
			}
			for _, elem := range staffdetails.StaffTuning {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StaffTuning = append(%s.StaffTuning, %s)", staffdetailsIdent, staffdetailsIdent, targetIdent))
			}
			if staffdetails.StaffSize != nil {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + staffdetails.StaffSize.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StaffSize = %s", staffdetailsIdent, targetIdent))
			}
		}
	}
	if stageSet.AttributesStage != nil {
		for _, staffsize := range __gong__sortStageSetInstances(stageSet.AttributesStage.StaffSizes, stageSet.AttributesStage.StaffSize_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			staffsizeIdent := "__attributes" + staffsize.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.StaffSize{Name: %s}).Stage(stageSet.AttributesStage)", staffsizeIdent, __gong__toRawStringLiteral(staffsize.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staffsizeIdent, __gong__toRawStringLiteral(staffsize.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Scaling = %s", staffsizeIdent, __gong__toRawStringLiteral(staffsize.Scaling)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", staffsizeIdent, __gong__toRawStringLiteral(staffsize.Value)))
		}
	}
	if stageSet.AttributesStage != nil {
		for _, stafftuning := range __gong__sortStageSetInstances(stageSet.AttributesStage.StaffTunings, stageSet.AttributesStage.StaffTuning_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			stafftuningIdent := "__attributes" + stafftuning.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.StaffTuning{Name: %s}).Stage(stageSet.AttributesStage)", stafftuningIdent, __gong__toRawStringLiteral(stafftuning.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stafftuningIdent, __gong__toRawStringLiteral(stafftuning.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Line = %d", stafftuningIdent, stafftuning.Line))
			values.WriteString(fmt.Sprintf("\n\t%s.TuningStep = %s", stafftuningIdent, __gong__toRawStringLiteral(string(stafftuning.TuningStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.TuningAlter = %s", stafftuningIdent, __gong__toRawStringLiteral(stafftuning.TuningAlter)))
			values.WriteString(fmt.Sprintf("\n\t%s.TuningOctave = %d", stafftuningIdent, stafftuning.TuningOctave))
		}
	}
	if stageSet.AttributesStage != nil {
		for _, time := range __gong__sortStageSetInstances(stageSet.AttributesStage.Times, stageSet.AttributesStage.Time_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			timeIdent := "__attributes" + time.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Time{Name: %s}).Stage(stageSet.AttributesStage)", timeIdent, __gong__toRawStringLiteral(time.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
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
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + time.Interchangeable.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Interchangeable = %s", timeIdent, targetIdent))
			}
		}
	}
	if stageSet.AttributesStage != nil {
		for _, transpose := range __gong__sortStageSetInstances(stageSet.AttributesStage.Transposes, stageSet.AttributesStage.Transpose_stagedOrder) {
			if lastStageDecl != "AttributesStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "AttributesStage"
			}
			transposeIdent := "__attributes" + transpose.GongGetIdentifier(stageSet.AttributesStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&attributes.Transpose{Name: %s}).Stage(stageSet.AttributesStage)", transposeIdent, __gong__toRawStringLiteral(transpose.Name)))
			if lastStageVal != "AttributesStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "AttributesStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", transposeIdent, __gong__toRawStringLiteral(transpose.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", transposeIdent, transpose.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", transposeIdent, __gong__toRawStringLiteral(transpose.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Diatonic = %d", transposeIdent, transpose.Diatonic))
			values.WriteString(fmt.Sprintf("\n\t%s.Chromatic = %s", transposeIdent, __gong__toRawStringLiteral(transpose.Chromatic)))
			values.WriteString(fmt.Sprintf("\n\t%s.OctaveChange = %d", transposeIdent, transpose.OctaveChange))
			if transpose.Double != nil {
				if lastStagePtr != "AttributesStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "AttributesStage"
				}
				targetIdent := "__attributes" + transpose.Double.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Double = %s", transposeIdent, targetIdent))
			}
		}
	}
	if stageSet.BarlineStage != nil {
		for _, barstylecolor := range __gong__sortStageSetInstances(stageSet.BarlineStage.BarStyleColors, stageSet.BarlineStage.BarStyleColor_stagedOrder) {
			if lastStageDecl != "BarlineStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "BarlineStage"
			}
			barstylecolorIdent := "__barline" + barstylecolor.GongGetIdentifier(stageSet.BarlineStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&barline.BarStyleColor{Name: %s}).Stage(stageSet.BarlineStage)", barstylecolorIdent, __gong__toRawStringLiteral(barstylecolor.Name)))
			if lastStageVal != "BarlineStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "BarlineStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", barstylecolorIdent, __gong__toRawStringLiteral(barstylecolor.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", barstylecolorIdent, __gong__toRawStringLiteral(barstylecolor.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", barstylecolorIdent, __gong__toRawStringLiteral(string(barstylecolor.Value))))
		}
	}
	if stageSet.BarlineStage != nil {
		for _, barline := range __gong__sortStageSetInstances(stageSet.BarlineStage.Barlines, stageSet.BarlineStage.Barline_stagedOrder) {
			if lastStageDecl != "BarlineStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "BarlineStage"
			}
			barlineIdent := "__barline" + barline.GongGetIdentifier(stageSet.BarlineStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&barline.Barline{Name: %s}).Stage(stageSet.BarlineStage)", barlineIdent, __gong__toRawStringLiteral(barline.Name)))
			if lastStageVal != "BarlineStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "BarlineStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", barlineIdent, __gong__toRawStringLiteral(barline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", barlineIdent, __gong__toRawStringLiteral(string(barline.Location))))
			values.WriteString(fmt.Sprintf("\n\t%s.Segno = %s", barlineIdent, __gong__toRawStringLiteral(barline.Segno)))
			values.WriteString(fmt.Sprintf("\n\t%s.Coda = %s", barlineIdent, __gong__toRawStringLiteral(barline.Coda)))
			values.WriteString(fmt.Sprintf("\n\t%s.Divisions = %s", barlineIdent, __gong__toRawStringLiteral(barline.Divisions)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", barlineIdent, __gong__toRawStringLiteral(barline.Id)))
			if barline.BarStyle != nil {
				if lastStagePtr != "BarlineStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "BarlineStage"
				}
				targetIdent := "__barline" + barline.BarStyle.GongGetIdentifier(stageSet.BarlineStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BarStyle = %s", barlineIdent, targetIdent))
			}
			if barline.Footnote != nil {
				if lastStagePtr != "BarlineStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "BarlineStage"
				}
				targetIdent := "__common" + barline.Footnote.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", barlineIdent, targetIdent))
			}
			if barline.Level != nil {
				if lastStagePtr != "BarlineStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "BarlineStage"
				}
				targetIdent := "__common" + barline.Level.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", barlineIdent, targetIdent))
			}
			if barline.WavyLine != nil {
				if lastStagePtr != "BarlineStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "BarlineStage"
				}
				targetIdent := "__common" + barline.WavyLine.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WavyLine = %s", barlineIdent, targetIdent))
			}
			if barline.Segno_1 != nil {
				if lastStagePtr != "BarlineStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "BarlineStage"
				}
				targetIdent := "__common" + barline.Segno_1.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Segno_1 = %s", barlineIdent, targetIdent))
			}
			if barline.Coda_1 != nil {
				if lastStagePtr != "BarlineStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "BarlineStage"
				}
				targetIdent := "__common" + barline.Coda_1.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Coda_1 = %s", barlineIdent, targetIdent))
			}
			for _, elem := range barline.Fermata {
				if lastStagePtr != "BarlineStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "BarlineStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fermata = append(%s.Fermata, %s)", barlineIdent, barlineIdent, targetIdent))
			}
			if barline.Ending != nil {
				if lastStagePtr != "BarlineStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "BarlineStage"
				}
				targetIdent := "__barline" + barline.Ending.GongGetIdentifier(stageSet.BarlineStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ending = %s", barlineIdent, targetIdent))
			}
			if barline.Repeat != nil {
				if lastStagePtr != "BarlineStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "BarlineStage"
				}
				targetIdent := "__barline" + barline.Repeat.GongGetIdentifier(stageSet.BarlineStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Repeat = %s", barlineIdent, targetIdent))
			}
		}
	}
	if stageSet.BarlineStage != nil {
		for _, ending := range __gong__sortStageSetInstances(stageSet.BarlineStage.Endings, stageSet.BarlineStage.Ending_stagedOrder) {
			if lastStageDecl != "BarlineStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "BarlineStage"
			}
			endingIdent := "__barline" + ending.GongGetIdentifier(stageSet.BarlineStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&barline.Ending{Name: %s}).Stage(stageSet.BarlineStage)", endingIdent, __gong__toRawStringLiteral(ending.Name)))
			if lastStageVal != "BarlineStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "BarlineStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", endingIdent, __gong__toRawStringLiteral(ending.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", endingIdent, __gong__toRawStringLiteral(ending.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", endingIdent, __gong__toRawStringLiteral(string(ending.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.EndLength = %s", endingIdent, __gong__toRawStringLiteral(ending.EndLength)))
			values.WriteString(fmt.Sprintf("\n\t%s.TextX = %s", endingIdent, __gong__toRawStringLiteral(ending.TextX)))
			values.WriteString(fmt.Sprintf("\n\t%s.TextY = %s", endingIdent, __gong__toRawStringLiteral(ending.TextY)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", endingIdent, __gong__toRawStringLiteral(string(ending.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultX = %s", endingIdent, __gong__toRawStringLiteral(ending.DefaultX)))
			values.WriteString(fmt.Sprintf("\n\t%s.DefaultY = %s", endingIdent, __gong__toRawStringLiteral(ending.DefaultY)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeX = %s", endingIdent, __gong__toRawStringLiteral(ending.RelativeX)))
			values.WriteString(fmt.Sprintf("\n\t%s.RelativeY = %s", endingIdent, __gong__toRawStringLiteral(ending.RelativeY)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", endingIdent, __gong__toRawStringLiteral(ending.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", endingIdent, __gong__toRawStringLiteral(string(ending.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", endingIdent, __gong__toRawStringLiteral(ending.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", endingIdent, __gong__toRawStringLiteral(string(ending.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", endingIdent, __gong__toRawStringLiteral(ending.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.System = %s", endingIdent, __gong__toRawStringLiteral(string(ending.System))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", endingIdent, __gong__toRawStringLiteral(ending.Value)))
		}
	}
	if stageSet.BarlineStage != nil {
		for _, repeat := range __gong__sortStageSetInstances(stageSet.BarlineStage.Repeats, stageSet.BarlineStage.Repeat_stagedOrder) {
			if lastStageDecl != "BarlineStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "BarlineStage"
			}
			repeatIdent := "__barline" + repeat.GongGetIdentifier(stageSet.BarlineStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&barline.Repeat{Name: %s}).Stage(stageSet.BarlineStage)", repeatIdent, __gong__toRawStringLiteral(repeat.Name)))
			if lastStageVal != "BarlineStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "BarlineStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", repeatIdent, __gong__toRawStringLiteral(repeat.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", repeatIdent, __gong__toRawStringLiteral(string(repeat.Direction))))
			values.WriteString(fmt.Sprintf("\n\t%s.Times = %d", repeatIdent, repeat.Times))
			values.WriteString(fmt.Sprintf("\n\t%s.AfterJump = %s", repeatIdent, __gong__toRawStringLiteral(string(repeat.AfterJump))))
			values.WriteString(fmt.Sprintf("\n\t%s.Winged = %s", repeatIdent, __gong__toRawStringLiteral(string(repeat.Winged))))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, accord := range __gong__sortStageSetInstances(stageSet.DirectionStage.Accords, stageSet.DirectionStage.Accord_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			accordIdent := "__direction" + accord.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Accord{Name: %s}).Stage(stageSet.DirectionStage)", accordIdent, __gong__toRawStringLiteral(accord.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", accordIdent, __gong__toRawStringLiteral(accord.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.String = %d", accordIdent, accord.String))
			values.WriteString(fmt.Sprintf("\n\t%s.TuningStep = %s", accordIdent, __gong__toRawStringLiteral(string(accord.TuningStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.TuningAlter = %s", accordIdent, __gong__toRawStringLiteral(accord.TuningAlter)))
			values.WriteString(fmt.Sprintf("\n\t%s.TuningOctave = %d", accordIdent, accord.TuningOctave))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, accordionregistration := range __gong__sortStageSetInstances(stageSet.DirectionStage.AccordionRegistrations, stageSet.DirectionStage.AccordionRegistration_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			accordionregistrationIdent := "__direction" + accordionregistration.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.AccordionRegistration{Name: %s}).Stage(stageSet.DirectionStage)", accordionregistrationIdent, __gong__toRawStringLiteral(accordionregistration.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
			if accordionregistration.AccordionHigh != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + accordionregistration.AccordionHigh.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AccordionHigh = %s", accordionregistrationIdent, targetIdent))
			}
			if accordionregistration.AccordionLow != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + accordionregistration.AccordionLow.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AccordionLow = %s", accordionregistrationIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, barre := range __gong__sortStageSetInstances(stageSet.DirectionStage.Barres, stageSet.DirectionStage.Barre_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			barreIdent := "__direction" + barre.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Barre{Name: %s}).Stage(stageSet.DirectionStage)", barreIdent, __gong__toRawStringLiteral(barre.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", barreIdent, __gong__toRawStringLiteral(barre.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", barreIdent, __gong__toRawStringLiteral(string(barre.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", barreIdent, __gong__toRawStringLiteral(barre.Color)))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, bass := range __gong__sortStageSetInstances(stageSet.DirectionStage.Basss, stageSet.DirectionStage.Bass_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			bassIdent := "__direction" + bass.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Bass{Name: %s}).Stage(stageSet.DirectionStage)", bassIdent, __gong__toRawStringLiteral(bass.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bassIdent, __gong__toRawStringLiteral(bass.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arrangement = %s", bassIdent, __gong__toRawStringLiteral(string(bass.Arrangement))))
			if bass.BassSeparator != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + bass.BassSeparator.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BassSeparator = %s", bassIdent, targetIdent))
			}
			if bass.BassStep != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + bass.BassStep.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BassStep = %s", bassIdent, targetIdent))
			}
			if bass.BassAlter != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + bass.BassAlter.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BassAlter = %s", bassIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, bassstep := range __gong__sortStageSetInstances(stageSet.DirectionStage.BassSteps, stageSet.DirectionStage.BassStep_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			bassstepIdent := "__direction" + bassstep.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.BassStep{Name: %s}).Stage(stageSet.DirectionStage)", bassstepIdent, __gong__toRawStringLiteral(bassstep.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, beatunittied := range __gong__sortStageSetInstances(stageSet.DirectionStage.BeatUnitTieds, stageSet.DirectionStage.BeatUnitTied_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			beatunittiedIdent := "__direction" + beatunittied.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.BeatUnitTied{Name: %s}).Stage(stageSet.DirectionStage)", beatunittiedIdent, __gong__toRawStringLiteral(beatunittied.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", beatunittiedIdent, __gong__toRawStringLiteral(beatunittied.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.BeatUnit = %s", beatunittiedIdent, __gong__toRawStringLiteral(string(beatunittied.BeatUnit))))
			for _, elem := range beatunittied.BeatUnitDot {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BeatUnitDot = append(%s.BeatUnitDot, %s)", beatunittiedIdent, beatunittiedIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, beater := range __gong__sortStageSetInstances(stageSet.DirectionStage.Beaters, stageSet.DirectionStage.Beater_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			beaterIdent := "__direction" + beater.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Beater{Name: %s}).Stage(stageSet.DirectionStage)", beaterIdent, __gong__toRawStringLiteral(beater.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", beaterIdent, __gong__toRawStringLiteral(beater.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tip = %s", beaterIdent, __gong__toRawStringLiteral(string(beater.Tip))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", beaterIdent, __gong__toRawStringLiteral(string(beater.Value))))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, bracket := range __gong__sortStageSetInstances(stageSet.DirectionStage.Brackets, stageSet.DirectionStage.Bracket_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			bracketIdent := "__direction" + bracket.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Bracket{Name: %s}).Stage(stageSet.DirectionStage)", bracketIdent, __gong__toRawStringLiteral(bracket.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, dashes := range __gong__sortStageSetInstances(stageSet.DirectionStage.Dashess, stageSet.DirectionStage.Dashes_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			dashesIdent := "__direction" + dashes.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Dashes{Name: %s}).Stage(stageSet.DirectionStage)", dashesIdent, __gong__toRawStringLiteral(dashes.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, degree := range __gong__sortStageSetInstances(stageSet.DirectionStage.Degrees, stageSet.DirectionStage.Degree_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			degreeIdent := "__direction" + degree.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Degree{Name: %s}).Stage(stageSet.DirectionStage)", degreeIdent, __gong__toRawStringLiteral(degree.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", degreeIdent, __gong__toRawStringLiteral(degree.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", degreeIdent, __gong__toRawStringLiteral(string(degree.PrintObject))))
			if degree.DegreeValue != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + degree.DegreeValue.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DegreeValue = %s", degreeIdent, targetIdent))
			}
			if degree.DegreeAlter != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + degree.DegreeAlter.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DegreeAlter = %s", degreeIdent, targetIdent))
			}
			if degree.DegreeType != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + degree.DegreeType.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DegreeType = %s", degreeIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, degreealter := range __gong__sortStageSetInstances(stageSet.DirectionStage.DegreeAlters, stageSet.DirectionStage.DegreeAlter_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			degreealterIdent := "__direction" + degreealter.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.DegreeAlter{Name: %s}).Stage(stageSet.DirectionStage)", degreealterIdent, __gong__toRawStringLiteral(degreealter.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, degreetype := range __gong__sortStageSetInstances(stageSet.DirectionStage.DegreeTypes, stageSet.DirectionStage.DegreeType_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			degreetypeIdent := "__direction" + degreetype.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.DegreeType{Name: %s}).Stage(stageSet.DirectionStage)", degreetypeIdent, __gong__toRawStringLiteral(degreetype.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, degreevalue := range __gong__sortStageSetInstances(stageSet.DirectionStage.DegreeValues, stageSet.DirectionStage.DegreeValue_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			degreevalueIdent := "__direction" + degreevalue.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.DegreeValue{Name: %s}).Stage(stageSet.DirectionStage)", degreevalueIdent, __gong__toRawStringLiteral(degreevalue.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, direction := range __gong__sortStageSetInstances(stageSet.DirectionStage.Directions, stageSet.DirectionStage.Direction_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			directionIdent := "__direction" + direction.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Direction{Name: %s}).Stage(stageSet.DirectionStage)", directionIdent, __gong__toRawStringLiteral(direction.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directionIdent, __gong__toRawStringLiteral(direction.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Placement = %s", directionIdent, __gong__toRawStringLiteral(string(direction.Placement))))
			values.WriteString(fmt.Sprintf("\n\t%s.Directive = %s", directionIdent, __gong__toRawStringLiteral(string(direction.Directive))))
			values.WriteString(fmt.Sprintf("\n\t%s.System = %s", directionIdent, __gong__toRawStringLiteral(string(direction.System))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directionIdent, __gong__toRawStringLiteral(direction.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Voice = %s", directionIdent, __gong__toRawStringLiteral(direction.Voice)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff = %d", directionIdent, direction.Staff))
			for _, elem := range direction.DirectionType {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DirectionType = append(%s.DirectionType, %s)", directionIdent, directionIdent, targetIdent))
			}
			if direction.Offset != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + direction.Offset.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", directionIdent, targetIdent))
			}
			if direction.Footnote != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + direction.Footnote.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", directionIdent, targetIdent))
			}
			if direction.Level != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + direction.Level.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", directionIdent, targetIdent))
			}
			if direction.Sound != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + direction.Sound.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sound = %s", directionIdent, targetIdent))
			}
			if direction.Listening != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + direction.Listening.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listening = %s", directionIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, directiontype := range __gong__sortStageSetInstances(stageSet.DirectionStage.DirectionTypes, stageSet.DirectionStage.DirectionType_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			directiontypeIdent := "__direction" + directiontype.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.DirectionType{Name: %s}).Stage(stageSet.DirectionStage)", directiontypeIdent, __gong__toRawStringLiteral(directiontype.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", directiontypeIdent, __gong__toRawStringLiteral(directiontype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", directiontypeIdent, __gong__toRawStringLiteral(directiontype.Id)))
			for _, elem := range directiontype.Rehearsal {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rehearsal = append(%s.Rehearsal, %s)", directiontypeIdent, directiontypeIdent, targetIdent))
			}
			for _, elem := range directiontype.Segno {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Segno = append(%s.Segno, %s)", directiontypeIdent, directiontypeIdent, targetIdent))
			}
			for _, elem := range directiontype.Coda {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Coda = append(%s.Coda, %s)", directiontypeIdent, directiontypeIdent, targetIdent))
			}
			for _, elem := range directiontype.Words {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Words = append(%s.Words, %s)", directiontypeIdent, directiontypeIdent, targetIdent))
			}
			for _, elem := range directiontype.Symbol {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Symbol = append(%s.Symbol, %s)", directiontypeIdent, directiontypeIdent, targetIdent))
			}
			if directiontype.Wedge != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.Wedge.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wedge = %s", directiontypeIdent, targetIdent))
			}
			for _, elem := range directiontype.Dynamics {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dynamics = append(%s.Dynamics, %s)", directiontypeIdent, directiontypeIdent, targetIdent))
			}
			if directiontype.Dashes != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.Dashes.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dashes = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.Bracket != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.Bracket.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bracket = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.Pedal != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.Pedal.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pedal = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.Metronome != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.Metronome.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metronome = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.OctaveShift != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.OctaveShift.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OctaveShift = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.HarpPedals != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.HarpPedals.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HarpPedals = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.Damp != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + directiontype.Damp.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Damp = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.DampAll != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + directiontype.DampAll.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DampAll = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.Eyeglasses != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + directiontype.Eyeglasses.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Eyeglasses = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.StringMute != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.StringMute.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StringMute = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.Scordatura != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.Scordatura.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Scordatura = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.Image != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.Image.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Image = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.PrincipalVoice != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.PrincipalVoice.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PrincipalVoice = %s", directiontypeIdent, targetIdent))
			}
			for _, elem := range directiontype.Percussion {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Percussion = append(%s.Percussion, %s)", directiontypeIdent, directiontypeIdent, targetIdent))
			}
			if directiontype.AccordionRegistration != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.AccordionRegistration.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AccordionRegistration = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.StaffDivide != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.StaffDivide.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StaffDivide = %s", directiontypeIdent, targetIdent))
			}
			if directiontype.OtherDirection != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + directiontype.OtherDirection.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherDirection = %s", directiontypeIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, effect := range __gong__sortStageSetInstances(stageSet.DirectionStage.Effects, stageSet.DirectionStage.Effect_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			effectIdent := "__direction" + effect.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Effect{Name: %s}).Stage(stageSet.DirectionStage)", effectIdent, __gong__toRawStringLiteral(effect.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", effectIdent, __gong__toRawStringLiteral(effect.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", effectIdent, __gong__toRawStringLiteral(effect.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", effectIdent, __gong__toRawStringLiteral(string(effect.Value))))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, feature := range __gong__sortStageSetInstances(stageSet.DirectionStage.Features, stageSet.DirectionStage.Feature_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			featureIdent := "__direction" + feature.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Feature{Name: %s}).Stage(stageSet.DirectionStage)", featureIdent, __gong__toRawStringLiteral(feature.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", featureIdent, __gong__toRawStringLiteral(feature.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", featureIdent, __gong__toRawStringLiteral(feature.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", featureIdent, __gong__toRawStringLiteral(feature.Value)))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, firstfret := range __gong__sortStageSetInstances(stageSet.DirectionStage.FirstFrets, stageSet.DirectionStage.FirstFret_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			firstfretIdent := "__direction" + firstfret.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.FirstFret{Name: %s}).Stage(stageSet.DirectionStage)", firstfretIdent, __gong__toRawStringLiteral(firstfret.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", firstfretIdent, __gong__toRawStringLiteral(firstfret.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", firstfretIdent, __gong__toRawStringLiteral(firstfret.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", firstfretIdent, __gong__toRawStringLiteral(string(firstfret.Location))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %d", firstfretIdent, firstfret.Value))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, frame := range __gong__sortStageSetInstances(stageSet.DirectionStage.Frames, stageSet.DirectionStage.Frame_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			frameIdent := "__direction" + frame.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Frame{Name: %s}).Stage(stageSet.DirectionStage)", frameIdent, __gong__toRawStringLiteral(frame.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + frame.FirstFret.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FirstFret = %s", frameIdent, targetIdent))
			}
			for _, elem := range frame.FrameNote {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FrameNote = append(%s.FrameNote, %s)", frameIdent, frameIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, framenote := range __gong__sortStageSetInstances(stageSet.DirectionStage.FrameNotes, stageSet.DirectionStage.FrameNote_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			framenoteIdent := "__direction" + framenote.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.FrameNote{Name: %s}).Stage(stageSet.DirectionStage)", framenoteIdent, __gong__toRawStringLiteral(framenote.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", framenoteIdent, __gong__toRawStringLiteral(framenote.Name)))
			if framenote.String != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + framenote.String.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.String = %s", framenoteIdent, targetIdent))
			}
			if framenote.Fret != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + framenote.Fret.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fret = %s", framenoteIdent, targetIdent))
			}
			if framenote.Fingering != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + framenote.Fingering.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fingering = %s", framenoteIdent, targetIdent))
			}
			if framenote.Barre != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + framenote.Barre.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Barre = %s", framenoteIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, glass := range __gong__sortStageSetInstances(stageSet.DirectionStage.Glasss, stageSet.DirectionStage.Glass_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			glassIdent := "__direction" + glass.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Glass{Name: %s}).Stage(stageSet.DirectionStage)", glassIdent, __gong__toRawStringLiteral(glass.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", glassIdent, __gong__toRawStringLiteral(glass.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", glassIdent, __gong__toRawStringLiteral(glass.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", glassIdent, __gong__toRawStringLiteral(string(glass.Value))))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, grouping := range __gong__sortStageSetInstances(stageSet.DirectionStage.Groupings, stageSet.DirectionStage.Grouping_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			groupingIdent := "__direction" + grouping.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Grouping{Name: %s}).Stage(stageSet.DirectionStage)", groupingIdent, __gong__toRawStringLiteral(grouping.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", groupingIdent, __gong__toRawStringLiteral(string(grouping.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.MemberOf = %s", groupingIdent, __gong__toRawStringLiteral(grouping.MemberOf)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", groupingIdent, __gong__toRawStringLiteral(grouping.Id)))
			for _, elem := range grouping.Feature {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Feature = append(%s.Feature, %s)", groupingIdent, groupingIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, harmony := range __gong__sortStageSetInstances(stageSet.DirectionStage.Harmonys, stageSet.DirectionStage.Harmony_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			harmonyIdent := "__direction" + harmony.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Harmony{Name: %s}).Stage(stageSet.DirectionStage)", harmonyIdent, __gong__toRawStringLiteral(harmony.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Root = append(%s.Root, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Numeral {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Numeral = append(%s.Numeral, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Function {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Function = append(%s.Function, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Kind {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Kind = append(%s.Kind, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Inversion {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Inversion = append(%s.Inversion, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Bass {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bass = append(%s.Bass, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			for _, elem := range harmony.Degree {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Degree = append(%s.Degree, %s)", harmonyIdent, harmonyIdent, targetIdent))
			}
			if harmony.Frame != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + harmony.Frame.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Frame = %s", harmonyIdent, targetIdent))
			}
			if harmony.Offset != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + harmony.Offset.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", harmonyIdent, targetIdent))
			}
			if harmony.Footnote != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + harmony.Footnote.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", harmonyIdent, targetIdent))
			}
			if harmony.Level != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + harmony.Level.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", harmonyIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, harmonyalter := range __gong__sortStageSetInstances(stageSet.DirectionStage.HarmonyAlters, stageSet.DirectionStage.HarmonyAlter_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			harmonyalterIdent := "__direction" + harmonyalter.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.HarmonyAlter{Name: %s}).Stage(stageSet.DirectionStage)", harmonyalterIdent, __gong__toRawStringLiteral(harmonyalter.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, harppedals := range __gong__sortStageSetInstances(stageSet.DirectionStage.HarpPedalss, stageSet.DirectionStage.HarpPedals_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			harppedalsIdent := "__direction" + harppedals.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.HarpPedals{Name: %s}).Stage(stageSet.DirectionStage)", harppedalsIdent, __gong__toRawStringLiteral(harppedals.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PedalTuning = append(%s.PedalTuning, %s)", harppedalsIdent, harppedalsIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, image := range __gong__sortStageSetInstances(stageSet.DirectionStage.Images, stageSet.DirectionStage.Image_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			imageIdent := "__direction" + image.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Image{Name: %s}).Stage(stageSet.DirectionStage)", imageIdent, __gong__toRawStringLiteral(image.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, instrumentchange := range __gong__sortStageSetInstances(stageSet.DirectionStage.InstrumentChanges, stageSet.DirectionStage.InstrumentChange_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			instrumentchangeIdent := "__direction" + instrumentchange.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.InstrumentChange{Name: %s}).Stage(stageSet.DirectionStage)", instrumentchangeIdent, __gong__toRawStringLiteral(instrumentchange.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrumentchangeIdent, __gong__toRawStringLiteral(instrumentchange.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", instrumentchangeIdent, __gong__toRawStringLiteral(instrumentchange.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.InstrumentSound = %s", instrumentchangeIdent, __gong__toRawStringLiteral(instrumentchange.InstrumentSound)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ensemble = %s", instrumentchangeIdent, __gong__toRawStringLiteral(instrumentchange.Ensemble)))
			if instrumentchange.Solo != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + instrumentchange.Solo.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Solo = %s", instrumentchangeIdent, targetIdent))
			}
			if instrumentchange.VirtualInstrument != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + instrumentchange.VirtualInstrument.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VirtualInstrument = %s", instrumentchangeIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, inversion := range __gong__sortStageSetInstances(stageSet.DirectionStage.Inversions, stageSet.DirectionStage.Inversion_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			inversionIdent := "__direction" + inversion.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Inversion{Name: %s}).Stage(stageSet.DirectionStage)", inversionIdent, __gong__toRawStringLiteral(inversion.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, kind := range __gong__sortStageSetInstances(stageSet.DirectionStage.Kinds, stageSet.DirectionStage.Kind_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			kindIdent := "__direction" + kind.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Kind{Name: %s}).Stage(stageSet.DirectionStage)", kindIdent, __gong__toRawStringLiteral(kind.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, listening := range __gong__sortStageSetInstances(stageSet.DirectionStage.Listenings, stageSet.DirectionStage.Listening_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			listeningIdent := "__direction" + listening.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Listening{Name: %s}).Stage(stageSet.DirectionStage)", listeningIdent, __gong__toRawStringLiteral(listening.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", listeningIdent, __gong__toRawStringLiteral(listening.Name)))
			for _, elem := range listening.Sync {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sync = append(%s.Sync, %s)", listeningIdent, listeningIdent, targetIdent))
			}
			for _, elem := range listening.OtherListening {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherListening = append(%s.OtherListening, %s)", listeningIdent, listeningIdent, targetIdent))
			}
			if listening.Offset != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + listening.Offset.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", listeningIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, measurenumbering := range __gong__sortStageSetInstances(stageSet.DirectionStage.MeasureNumberings, stageSet.DirectionStage.MeasureNumbering_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			measurenumberingIdent := "__direction" + measurenumbering.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.MeasureNumbering{Name: %s}).Stage(stageSet.DirectionStage)", measurenumberingIdent, __gong__toRawStringLiteral(measurenumbering.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, membrane := range __gong__sortStageSetInstances(stageSet.DirectionStage.Membranes, stageSet.DirectionStage.Membrane_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			membraneIdent := "__direction" + membrane.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Membrane{Name: %s}).Stage(stageSet.DirectionStage)", membraneIdent, __gong__toRawStringLiteral(membrane.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", membraneIdent, __gong__toRawStringLiteral(membrane.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", membraneIdent, __gong__toRawStringLiteral(membrane.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", membraneIdent, __gong__toRawStringLiteral(string(membrane.Value))))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, metal := range __gong__sortStageSetInstances(stageSet.DirectionStage.Metals, stageSet.DirectionStage.Metal_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			metalIdent := "__direction" + metal.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Metal{Name: %s}).Stage(stageSet.DirectionStage)", metalIdent, __gong__toRawStringLiteral(metal.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metalIdent, __gong__toRawStringLiteral(metal.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", metalIdent, __gong__toRawStringLiteral(metal.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", metalIdent, __gong__toRawStringLiteral(string(metal.Value))))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, metronome := range __gong__sortStageSetInstances(stageSet.DirectionStage.Metronomes, stageSet.DirectionStage.Metronome_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			metronomeIdent := "__direction" + metronome.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Metronome{Name: %s}).Stage(stageSet.DirectionStage)", metronomeIdent, __gong__toRawStringLiteral(metronome.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
			for _, elem := range metronome.BeatUnitDot {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BeatUnitDot = append(%s.BeatUnitDot, %s)", metronomeIdent, metronomeIdent, targetIdent))
			}
			for _, elem := range metronome.BeatUnitTied {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BeatUnitTied = append(%s.BeatUnitTied, %s)", metronomeIdent, metronomeIdent, targetIdent))
			}
			if metronome.PerMinute != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + metronome.PerMinute.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PerMinute = %s", metronomeIdent, targetIdent))
			}
			if metronome.MetronomeArrows != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + metronome.MetronomeArrows.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MetronomeArrows = %s", metronomeIdent, targetIdent))
			}
			for _, elem := range metronome.MetronomeNote {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MetronomeNote = append(%s.MetronomeNote, %s)", metronomeIdent, metronomeIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, metronomebeam := range __gong__sortStageSetInstances(stageSet.DirectionStage.MetronomeBeams, stageSet.DirectionStage.MetronomeBeam_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			metronomebeamIdent := "__direction" + metronomebeam.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.MetronomeBeam{Name: %s}).Stage(stageSet.DirectionStage)", metronomebeamIdent, __gong__toRawStringLiteral(metronomebeam.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronomebeamIdent, __gong__toRawStringLiteral(metronomebeam.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %d", metronomebeamIdent, metronomebeam.Number))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", metronomebeamIdent, __gong__toRawStringLiteral(string(metronomebeam.Value))))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, metronomenote := range __gong__sortStageSetInstances(stageSet.DirectionStage.MetronomeNotes, stageSet.DirectionStage.MetronomeNote_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			metronomenoteIdent := "__direction" + metronomenote.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.MetronomeNote{Name: %s}).Stage(stageSet.DirectionStage)", metronomenoteIdent, __gong__toRawStringLiteral(metronomenote.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronomenoteIdent, __gong__toRawStringLiteral(metronomenote.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MetronomeType = %s", metronomenoteIdent, __gong__toRawStringLiteral(string(metronomenote.MetronomeType))))
			for _, elem := range metronomenote.MetronomeDot {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MetronomeDot = append(%s.MetronomeDot, %s)", metronomenoteIdent, metronomenoteIdent, targetIdent))
			}
			for _, elem := range metronomenote.MetronomeBeam {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MetronomeBeam = append(%s.MetronomeBeam, %s)", metronomenoteIdent, metronomenoteIdent, targetIdent))
			}
			if metronomenote.MetronomeTied != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + metronomenote.MetronomeTied.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MetronomeTied = %s", metronomenoteIdent, targetIdent))
			}
			if metronomenote.MetronomeTuplet != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + metronomenote.MetronomeTuplet.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MetronomeTuplet = %s", metronomenoteIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, metronometied := range __gong__sortStageSetInstances(stageSet.DirectionStage.MetronomeTieds, stageSet.DirectionStage.MetronomeTied_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			metronometiedIdent := "__direction" + metronometied.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.MetronomeTied{Name: %s}).Stage(stageSet.DirectionStage)", metronometiedIdent, __gong__toRawStringLiteral(metronometied.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronometiedIdent, __gong__toRawStringLiteral(metronometied.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", metronometiedIdent, __gong__toRawStringLiteral(string(metronometied.Type))))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, metronometuplet := range __gong__sortStageSetInstances(stageSet.DirectionStage.MetronomeTuplets, stageSet.DirectionStage.MetronomeTuplet_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			metronometupletIdent := "__direction" + metronometuplet.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.MetronomeTuplet{Name: %s}).Stage(stageSet.DirectionStage)", metronometupletIdent, __gong__toRawStringLiteral(metronometuplet.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", metronometupletIdent, __gong__toRawStringLiteral(metronometuplet.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", metronometupletIdent, __gong__toRawStringLiteral(string(metronometuplet.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Bracket = %s", metronometupletIdent, __gong__toRawStringLiteral(string(metronometuplet.Bracket))))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowNumber = %s", metronometupletIdent, __gong__toRawStringLiteral(string(metronometuplet.ShowNumber))))
			values.WriteString(fmt.Sprintf("\n\t%s.ActualNotes = %d", metronometupletIdent, metronometuplet.ActualNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.NormalNotes = %d", metronometupletIdent, metronometuplet.NormalNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.NormalType = %s", metronometupletIdent, __gong__toRawStringLiteral(string(metronometuplet.NormalType))))
			for _, elem := range metronometuplet.NormalDot {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NormalDot = append(%s.NormalDot, %s)", metronometupletIdent, metronometupletIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, numeral := range __gong__sortStageSetInstances(stageSet.DirectionStage.Numerals, stageSet.DirectionStage.Numeral_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			numeralIdent := "__direction" + numeral.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Numeral{Name: %s}).Stage(stageSet.DirectionStage)", numeralIdent, __gong__toRawStringLiteral(numeral.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", numeralIdent, __gong__toRawStringLiteral(numeral.Name)))
			if numeral.NumeralRoot != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + numeral.NumeralRoot.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NumeralRoot = %s", numeralIdent, targetIdent))
			}
			if numeral.NumeralAlter != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + numeral.NumeralAlter.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NumeralAlter = %s", numeralIdent, targetIdent))
			}
			if numeral.NumeralKey != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + numeral.NumeralKey.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NumeralKey = %s", numeralIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, numeralkey := range __gong__sortStageSetInstances(stageSet.DirectionStage.NumeralKeys, stageSet.DirectionStage.NumeralKey_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			numeralkeyIdent := "__direction" + numeralkey.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.NumeralKey{Name: %s}).Stage(stageSet.DirectionStage)", numeralkeyIdent, __gong__toRawStringLiteral(numeralkey.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", numeralkeyIdent, __gong__toRawStringLiteral(numeralkey.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", numeralkeyIdent, __gong__toRawStringLiteral(string(numeralkey.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.NumeralFifths = %d", numeralkeyIdent, numeralkey.NumeralFifths))
			values.WriteString(fmt.Sprintf("\n\t%s.NumeralMode = %s", numeralkeyIdent, __gong__toRawStringLiteral(string(numeralkey.NumeralMode))))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, numeralroot := range __gong__sortStageSetInstances(stageSet.DirectionStage.NumeralRoots, stageSet.DirectionStage.NumeralRoot_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			numeralrootIdent := "__direction" + numeralroot.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.NumeralRoot{Name: %s}).Stage(stageSet.DirectionStage)", numeralrootIdent, __gong__toRawStringLiteral(numeralroot.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, octaveshift := range __gong__sortStageSetInstances(stageSet.DirectionStage.OctaveShifts, stageSet.DirectionStage.OctaveShift_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			octaveshiftIdent := "__direction" + octaveshift.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.OctaveShift{Name: %s}).Stage(stageSet.DirectionStage)", octaveshiftIdent, __gong__toRawStringLiteral(octaveshift.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, offset := range __gong__sortStageSetInstances(stageSet.DirectionStage.Offsets, stageSet.DirectionStage.Offset_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			offsetIdent := "__direction" + offset.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Offset{Name: %s}).Stage(stageSet.DirectionStage)", offsetIdent, __gong__toRawStringLiteral(offset.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", offsetIdent, __gong__toRawStringLiteral(offset.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Sound = %s", offsetIdent, __gong__toRawStringLiteral(string(offset.Sound))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", offsetIdent, __gong__toRawStringLiteral(offset.Value)))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, otherdirection := range __gong__sortStageSetInstances(stageSet.DirectionStage.OtherDirections, stageSet.DirectionStage.OtherDirection_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			otherdirectionIdent := "__direction" + otherdirection.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.OtherDirection{Name: %s}).Stage(stageSet.DirectionStage)", otherdirectionIdent, __gong__toRawStringLiteral(otherdirection.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, otherlistening := range __gong__sortStageSetInstances(stageSet.DirectionStage.OtherListenings, stageSet.DirectionStage.OtherListening_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			otherlisteningIdent := "__direction" + otherlistening.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.OtherListening{Name: %s}).Stage(stageSet.DirectionStage)", otherlisteningIdent, __gong__toRawStringLiteral(otherlistening.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", otherlisteningIdent, __gong__toRawStringLiteral(otherlistening.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", otherlisteningIdent, __gong__toRawStringLiteral(otherlistening.Type)))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", otherlisteningIdent, __gong__toRawStringLiteral(otherlistening.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeOnly = %s", otherlisteningIdent, __gong__toRawStringLiteral(otherlistening.TimeOnly)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", otherlisteningIdent, __gong__toRawStringLiteral(otherlistening.Value)))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, pedal := range __gong__sortStageSetInstances(stageSet.DirectionStage.Pedals, stageSet.DirectionStage.Pedal_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			pedalIdent := "__direction" + pedal.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Pedal{Name: %s}).Stage(stageSet.DirectionStage)", pedalIdent, __gong__toRawStringLiteral(pedal.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, pedaltuning := range __gong__sortStageSetInstances(stageSet.DirectionStage.PedalTunings, stageSet.DirectionStage.PedalTuning_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			pedaltuningIdent := "__direction" + pedaltuning.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.PedalTuning{Name: %s}).Stage(stageSet.DirectionStage)", pedaltuningIdent, __gong__toRawStringLiteral(pedaltuning.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pedaltuningIdent, __gong__toRawStringLiteral(pedaltuning.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PedalStep = %s", pedaltuningIdent, __gong__toRawStringLiteral(string(pedaltuning.PedalStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.PedalAlter = %s", pedaltuningIdent, __gong__toRawStringLiteral(pedaltuning.PedalAlter)))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, perminute := range __gong__sortStageSetInstances(stageSet.DirectionStage.PerMinutes, stageSet.DirectionStage.PerMinute_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			perminuteIdent := "__direction" + perminute.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.PerMinute{Name: %s}).Stage(stageSet.DirectionStage)", perminuteIdent, __gong__toRawStringLiteral(perminute.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", perminuteIdent, __gong__toRawStringLiteral(perminute.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", perminuteIdent, __gong__toRawStringLiteral(perminute.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", perminuteIdent, __gong__toRawStringLiteral(string(perminute.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", perminuteIdent, __gong__toRawStringLiteral(perminute.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", perminuteIdent, __gong__toRawStringLiteral(string(perminute.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", perminuteIdent, __gong__toRawStringLiteral(perminute.Value)))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, percussion := range __gong__sortStageSetInstances(stageSet.DirectionStage.Percussions, stageSet.DirectionStage.Percussion_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			percussionIdent := "__direction" + percussion.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Percussion{Name: %s}).Stage(stageSet.DirectionStage)", percussionIdent, __gong__toRawStringLiteral(percussion.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + percussion.Glass.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Glass = %s", percussionIdent, targetIdent))
			}
			if percussion.Metal != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + percussion.Metal.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Metal = %s", percussionIdent, targetIdent))
			}
			if percussion.Wood != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + percussion.Wood.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wood = %s", percussionIdent, targetIdent))
			}
			if percussion.Pitched != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + percussion.Pitched.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pitched = %s", percussionIdent, targetIdent))
			}
			if percussion.Membrane != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + percussion.Membrane.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Membrane = %s", percussionIdent, targetIdent))
			}
			if percussion.Effect != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + percussion.Effect.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Effect = %s", percussionIdent, targetIdent))
			}
			if percussion.Timpani != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + percussion.Timpani.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Timpani = %s", percussionIdent, targetIdent))
			}
			if percussion.Beater != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + percussion.Beater.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Beater = %s", percussionIdent, targetIdent))
			}
			if percussion.Stick != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + percussion.Stick.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stick = %s", percussionIdent, targetIdent))
			}
			if percussion.OtherPercussion != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + percussion.OtherPercussion.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherPercussion = %s", percussionIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, pitched := range __gong__sortStageSetInstances(stageSet.DirectionStage.Pitcheds, stageSet.DirectionStage.Pitched_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			pitchedIdent := "__direction" + pitched.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Pitched{Name: %s}).Stage(stageSet.DirectionStage)", pitchedIdent, __gong__toRawStringLiteral(pitched.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pitchedIdent, __gong__toRawStringLiteral(pitched.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", pitchedIdent, __gong__toRawStringLiteral(pitched.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", pitchedIdent, __gong__toRawStringLiteral(string(pitched.Value))))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, principalvoice := range __gong__sortStageSetInstances(stageSet.DirectionStage.PrincipalVoices, stageSet.DirectionStage.PrincipalVoice_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			principalvoiceIdent := "__direction" + principalvoice.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.PrincipalVoice{Name: %s}).Stage(stageSet.DirectionStage)", principalvoiceIdent, __gong__toRawStringLiteral(principalvoice.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, print := range __gong__sortStageSetInstances(stageSet.DirectionStage.Prints, stageSet.DirectionStage.Print_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			printIdent := "__direction" + print.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Print{Name: %s}).Stage(stageSet.DirectionStage)", printIdent, __gong__toRawStringLiteral(print.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", printIdent, __gong__toRawStringLiteral(print.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StaffSpacing = %s", printIdent, __gong__toRawStringLiteral(print.StaffSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.NewSystem = %s", printIdent, __gong__toRawStringLiteral(string(print.NewSystem))))
			values.WriteString(fmt.Sprintf("\n\t%s.NewPage = %s", printIdent, __gong__toRawStringLiteral(string(print.NewPage))))
			values.WriteString(fmt.Sprintf("\n\t%s.BlankPage = %d", printIdent, print.BlankPage))
			values.WriteString(fmt.Sprintf("\n\t%s.PageNumber = %s", printIdent, __gong__toRawStringLiteral(print.PageNumber)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", printIdent, __gong__toRawStringLiteral(print.Id)))
			if print.PageLayout != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + print.PageLayout.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PageLayout = %s", printIdent, targetIdent))
			}
			if print.SystemLayout != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + print.SystemLayout.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemLayout = %s", printIdent, targetIdent))
			}
			for _, elem := range print.StaffLayout {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StaffLayout = append(%s.StaffLayout, %s)", printIdent, printIdent, targetIdent))
			}
			if print.MeasureLayout != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + print.MeasureLayout.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MeasureLayout = %s", printIdent, targetIdent))
			}
			if print.MeasureNumbering != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + print.MeasureNumbering.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MeasureNumbering = %s", printIdent, targetIdent))
			}
			if print.PartNameDisplay != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + print.PartNameDisplay.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartNameDisplay = %s", printIdent, targetIdent))
			}
			if print.PartAbbreviationDisplay != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + print.PartAbbreviationDisplay.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartAbbreviationDisplay = %s", printIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, root := range __gong__sortStageSetInstances(stageSet.DirectionStage.Roots, stageSet.DirectionStage.Root_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			rootIdent := "__direction" + root.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Root{Name: %s}).Stage(stageSet.DirectionStage)", rootIdent, __gong__toRawStringLiteral(root.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", rootIdent, __gong__toRawStringLiteral(root.Name)))
			if root.RootStep != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + root.RootStep.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootStep = %s", rootIdent, targetIdent))
			}
			if root.RootAlter != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + root.RootAlter.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RootAlter = %s", rootIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, rootstep := range __gong__sortStageSetInstances(stageSet.DirectionStage.RootSteps, stageSet.DirectionStage.RootStep_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			rootstepIdent := "__direction" + rootstep.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.RootStep{Name: %s}).Stage(stageSet.DirectionStage)", rootstepIdent, __gong__toRawStringLiteral(rootstep.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, scordatura := range __gong__sortStageSetInstances(stageSet.DirectionStage.Scordaturas, stageSet.DirectionStage.Scordatura_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			scordaturaIdent := "__direction" + scordatura.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Scordatura{Name: %s}).Stage(stageSet.DirectionStage)", scordaturaIdent, __gong__toRawStringLiteral(scordatura.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scordaturaIdent, __gong__toRawStringLiteral(scordatura.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", scordaturaIdent, __gong__toRawStringLiteral(scordatura.Id)))
			for _, elem := range scordatura.Accord {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accord = append(%s.Accord, %s)", scordaturaIdent, scordaturaIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, sound := range __gong__sortStageSetInstances(stageSet.DirectionStage.Sounds, stageSet.DirectionStage.Sound_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			soundIdent := "__direction" + sound.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Sound{Name: %s}).Stage(stageSet.DirectionStage)", soundIdent, __gong__toRawStringLiteral(sound.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InstrumentChange = append(%s.InstrumentChange, %s)", soundIdent, soundIdent, targetIdent))
			}
			for _, elem := range sound.MidiDevice {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MidiDevice = append(%s.MidiDevice, %s)", soundIdent, soundIdent, targetIdent))
			}
			for _, elem := range sound.MidiInstrument {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MidiInstrument = append(%s.MidiInstrument, %s)", soundIdent, soundIdent, targetIdent))
			}
			for _, elem := range sound.Play {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Play = append(%s.Play, %s)", soundIdent, soundIdent, targetIdent))
			}
			if sound.Swing != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + sound.Swing.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Swing = %s", soundIdent, targetIdent))
			}
			if sound.Offset != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__direction" + sound.Offset.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", soundIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, staffdivide := range __gong__sortStageSetInstances(stageSet.DirectionStage.StaffDivides, stageSet.DirectionStage.StaffDivide_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			staffdivideIdent := "__direction" + staffdivide.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.StaffDivide{Name: %s}).Stage(stageSet.DirectionStage)", staffdivideIdent, __gong__toRawStringLiteral(staffdivide.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, stick := range __gong__sortStageSetInstances(stageSet.DirectionStage.Sticks, stageSet.DirectionStage.Stick_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			stickIdent := "__direction" + stick.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Stick{Name: %s}).Stage(stageSet.DirectionStage)", stickIdent, __gong__toRawStringLiteral(stick.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", stickIdent, __gong__toRawStringLiteral(stick.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Tip = %s", stickIdent, __gong__toRawStringLiteral(string(stick.Tip))))
			values.WriteString(fmt.Sprintf("\n\t%s.Parentheses = %s", stickIdent, __gong__toRawStringLiteral(string(stick.Parentheses))))
			values.WriteString(fmt.Sprintf("\n\t%s.DashedCircle = %s", stickIdent, __gong__toRawStringLiteral(string(stick.DashedCircle))))
			values.WriteString(fmt.Sprintf("\n\t%s.StickType = %s", stickIdent, __gong__toRawStringLiteral(string(stick.StickType))))
			values.WriteString(fmt.Sprintf("\n\t%s.StickMaterial = %s", stickIdent, __gong__toRawStringLiteral(string(stick.StickMaterial))))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, stringmute := range __gong__sortStageSetInstances(stageSet.DirectionStage.StringMutes, stageSet.DirectionStage.StringMute_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			stringmuteIdent := "__direction" + stringmute.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.StringMute{Name: %s}).Stage(stageSet.DirectionStage)", stringmuteIdent, __gong__toRawStringLiteral(stringmute.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, styletext := range __gong__sortStageSetInstances(stageSet.DirectionStage.StyleTexts, stageSet.DirectionStage.StyleText_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			styletextIdent := "__direction" + styletext.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.StyleText{Name: %s}).Stage(stageSet.DirectionStage)", styletextIdent, __gong__toRawStringLiteral(styletext.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, swing := range __gong__sortStageSetInstances(stageSet.DirectionStage.Swings, stageSet.DirectionStage.Swing_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			swingIdent := "__direction" + swing.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Swing{Name: %s}).Stage(stageSet.DirectionStage)", swingIdent, __gong__toRawStringLiteral(swing.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", swingIdent, __gong__toRawStringLiteral(swing.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.First = %d", swingIdent, swing.First))
			values.WriteString(fmt.Sprintf("\n\t%s.Second = %d", swingIdent, swing.Second))
			values.WriteString(fmt.Sprintf("\n\t%s.SwingType = %s", swingIdent, __gong__toRawStringLiteral(string(swing.SwingType))))
			values.WriteString(fmt.Sprintf("\n\t%s.SwingStyle = %s", swingIdent, __gong__toRawStringLiteral(swing.SwingStyle)))
			if swing.Straight != nil {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + swing.Straight.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Straight = %s", swingIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, sync := range __gong__sortStageSetInstances(stageSet.DirectionStage.Syncs, stageSet.DirectionStage.Sync_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			syncIdent := "__direction" + sync.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Sync{Name: %s}).Stage(stageSet.DirectionStage)", syncIdent, __gong__toRawStringLiteral(sync.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", syncIdent, __gong__toRawStringLiteral(sync.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", syncIdent, __gong__toRawStringLiteral(string(sync.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Latency = %d", syncIdent, sync.Latency))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", syncIdent, __gong__toRawStringLiteral(sync.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeOnly = %s", syncIdent, __gong__toRawStringLiteral(sync.TimeOnly)))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, timemodification := range __gong__sortStageSetInstances(stageSet.DirectionStage.TimeModifications, stageSet.DirectionStage.TimeModification_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			timemodificationIdent := "__direction" + timemodification.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.TimeModification{Name: %s}).Stage(stageSet.DirectionStage)", timemodificationIdent, __gong__toRawStringLiteral(timemodification.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timemodificationIdent, __gong__toRawStringLiteral(timemodification.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ActualNotes = %d", timemodificationIdent, timemodification.ActualNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.NormalNotes = %d", timemodificationIdent, timemodification.NormalNotes))
			values.WriteString(fmt.Sprintf("\n\t%s.NormalType = %s", timemodificationIdent, __gong__toRawStringLiteral(string(timemodification.NormalType))))
			for _, elem := range timemodification.NormalDot {
				if lastStagePtr != "DirectionStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "DirectionStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NormalDot = append(%s.NormalDot, %s)", timemodificationIdent, timemodificationIdent, targetIdent))
			}
		}
	}
	if stageSet.DirectionStage != nil {
		for _, timpani := range __gong__sortStageSetInstances(stageSet.DirectionStage.Timpanis, stageSet.DirectionStage.Timpani_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			timpaniIdent := "__direction" + timpani.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Timpani{Name: %s}).Stage(stageSet.DirectionStage)", timpaniIdent, __gong__toRawStringLiteral(timpani.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timpaniIdent, __gong__toRawStringLiteral(timpani.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", timpaniIdent, __gong__toRawStringLiteral(timpani.Smufl)))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, virtualinstrument := range __gong__sortStageSetInstances(stageSet.DirectionStage.VirtualInstruments, stageSet.DirectionStage.VirtualInstrument_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			virtualinstrumentIdent := "__direction" + virtualinstrument.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.VirtualInstrument{Name: %s}).Stage(stageSet.DirectionStage)", virtualinstrumentIdent, __gong__toRawStringLiteral(virtualinstrument.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", virtualinstrumentIdent, __gong__toRawStringLiteral(virtualinstrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.VirtualLibrary = %s", virtualinstrumentIdent, __gong__toRawStringLiteral(virtualinstrument.VirtualLibrary)))
			values.WriteString(fmt.Sprintf("\n\t%s.VirtualName = %s", virtualinstrumentIdent, __gong__toRawStringLiteral(virtualinstrument.VirtualName)))
		}
	}
	if stageSet.DirectionStage != nil {
		for _, wedge := range __gong__sortStageSetInstances(stageSet.DirectionStage.Wedges, stageSet.DirectionStage.Wedge_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			wedgeIdent := "__direction" + wedge.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Wedge{Name: %s}).Stage(stageSet.DirectionStage)", wedgeIdent, __gong__toRawStringLiteral(wedge.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
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
	if stageSet.DirectionStage != nil {
		for _, wood := range __gong__sortStageSetInstances(stageSet.DirectionStage.Woods, stageSet.DirectionStage.Wood_stagedOrder) {
			if lastStageDecl != "DirectionStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "DirectionStage"
			}
			woodIdent := "__direction" + wood.GongGetIdentifier(stageSet.DirectionStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&direction.Wood{Name: %s}).Stage(stageSet.DirectionStage)", woodIdent, __gong__toRawStringLiteral(wood.Name)))
			if lastStageVal != "DirectionStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "DirectionStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", woodIdent, __gong__toRawStringLiteral(wood.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Smufl = %s", woodIdent, __gong__toRawStringLiteral(wood.Smufl)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", woodIdent, __gong__toRawStringLiteral(string(wood.Value))))
		}
	}
	if stageSet.NoteStage != nil {
		for _, accidental := range __gong__sortStageSetInstances(stageSet.NoteStage.Accidentals, stageSet.NoteStage.Accidental_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			accidentalIdent := "__note" + accidental.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Accidental{Name: %s}).Stage(stageSet.NoteStage)", accidentalIdent, __gong__toRawStringLiteral(accidental.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, accidentalmark := range __gong__sortStageSetInstances(stageSet.NoteStage.AccidentalMarks, stageSet.NoteStage.AccidentalMark_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			accidentalmarkIdent := "__note" + accidentalmark.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.AccidentalMark{Name: %s}).Stage(stageSet.NoteStage)", accidentalmarkIdent, __gong__toRawStringLiteral(accidentalmark.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, arpeggiate := range __gong__sortStageSetInstances(stageSet.NoteStage.Arpeggiates, stageSet.NoteStage.Arpeggiate_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			arpeggiateIdent := "__note" + arpeggiate.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Arpeggiate{Name: %s}).Stage(stageSet.NoteStage)", arpeggiateIdent, __gong__toRawStringLiteral(arpeggiate.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, arrow := range __gong__sortStageSetInstances(stageSet.NoteStage.Arrows, stageSet.NoteStage.Arrow_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			arrowIdent := "__note" + arrow.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Arrow{Name: %s}).Stage(stageSet.NoteStage)", arrowIdent, __gong__toRawStringLiteral(arrow.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
			if arrow.Arrowhead != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + arrow.Arrowhead.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arrowhead = %s", arrowIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, articulations := range __gong__sortStageSetInstances(stageSet.NoteStage.Articulationss, stageSet.NoteStage.Articulations_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			articulationsIdent := "__note" + articulations.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Articulations{Name: %s}).Stage(stageSet.NoteStage)", articulationsIdent, __gong__toRawStringLiteral(articulations.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", articulationsIdent, __gong__toRawStringLiteral(articulations.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", articulationsIdent, __gong__toRawStringLiteral(articulations.Id)))
			for _, elem := range articulations.Accent {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accent = append(%s.Accent, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.StrongAccent {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StrongAccent = append(%s.StrongAccent, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Staccato {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staccato = append(%s.Staccato, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Tenuto {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tenuto = append(%s.Tenuto, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.DetachedLegato {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DetachedLegato = append(%s.DetachedLegato, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Staccatissimo {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Staccatissimo = append(%s.Staccatissimo, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Spiccato {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Spiccato = append(%s.Spiccato, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Scoop {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Scoop = append(%s.Scoop, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Plop {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Plop = append(%s.Plop, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Doit {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Doit = append(%s.Doit, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Falloff {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Falloff = append(%s.Falloff, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.BreathMark {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BreathMark = append(%s.BreathMark, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Caesura {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Caesura = append(%s.Caesura, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Stress {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stress = append(%s.Stress, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.Unstress {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Unstress = append(%s.Unstress, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.SoftAccent {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SoftAccent = append(%s.SoftAccent, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
			for _, elem := range articulations.OtherArticulation {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherArticulation = append(%s.OtherArticulation, %s)", articulationsIdent, articulationsIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, assess := range __gong__sortStageSetInstances(stageSet.NoteStage.Assesss, stageSet.NoteStage.Assess_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			assessIdent := "__note" + assess.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Assess{Name: %s}).Stage(stageSet.NoteStage)", assessIdent, __gong__toRawStringLiteral(assess.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", assessIdent, __gong__toRawStringLiteral(assess.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", assessIdent, __gong__toRawStringLiteral(string(assess.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", assessIdent, __gong__toRawStringLiteral(assess.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeOnly = %s", assessIdent, __gong__toRawStringLiteral(assess.TimeOnly)))
		}
	}
	if stageSet.NoteStage != nil {
		for _, backup := range __gong__sortStageSetInstances(stageSet.NoteStage.Backups, stageSet.NoteStage.Backup_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			backupIdent := "__note" + backup.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Backup{Name: %s}).Stage(stageSet.NoteStage)", backupIdent, __gong__toRawStringLiteral(backup.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", backupIdent, __gong__toRawStringLiteral(backup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = %s", backupIdent, __gong__toRawStringLiteral(backup.Duration)))
			if backup.Footnote != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + backup.Footnote.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", backupIdent, targetIdent))
			}
			if backup.Level != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + backup.Level.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", backupIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, beam := range __gong__sortStageSetInstances(stageSet.NoteStage.Beams, stageSet.NoteStage.Beam_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			beamIdent := "__note" + beam.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Beam{Name: %s}).Stage(stageSet.NoteStage)", beamIdent, __gong__toRawStringLiteral(beam.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, bend := range __gong__sortStageSetInstances(stageSet.NoteStage.Bends, stageSet.NoteStage.Bend_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			bendIdent := "__note" + bend.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Bend{Name: %s}).Stage(stageSet.NoteStage)", bendIdent, __gong__toRawStringLiteral(bend.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
			if bend.PreBend != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + bend.PreBend.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PreBend = %s", bendIdent, targetIdent))
			}
			if bend.Release != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + bend.Release.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Release = %s", bendIdent, targetIdent))
			}
			if bend.WithBar != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + bend.WithBar.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WithBar = %s", bendIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, breathmark := range __gong__sortStageSetInstances(stageSet.NoteStage.BreathMarks, stageSet.NoteStage.BreathMark_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			breathmarkIdent := "__note" + breathmark.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.BreathMark{Name: %s}).Stage(stageSet.NoteStage)", breathmarkIdent, __gong__toRawStringLiteral(breathmark.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, caesura := range __gong__sortStageSetInstances(stageSet.NoteStage.Caesuras, stageSet.NoteStage.Caesura_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			caesuraIdent := "__note" + caesura.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Caesura{Name: %s}).Stage(stageSet.NoteStage)", caesuraIdent, __gong__toRawStringLiteral(caesura.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, elision := range __gong__sortStageSetInstances(stageSet.NoteStage.Elisions, stageSet.NoteStage.Elision_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			elisionIdent := "__note" + elision.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Elision{Name: %s}).Stage(stageSet.NoteStage)", elisionIdent, __gong__toRawStringLiteral(elision.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, emptyline := range __gong__sortStageSetInstances(stageSet.NoteStage.EmptyLines, stageSet.NoteStage.EmptyLine_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			emptylineIdent := "__note" + emptyline.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.EmptyLine{Name: %s}).Stage(stageSet.NoteStage)", emptylineIdent, __gong__toRawStringLiteral(emptyline.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, extend := range __gong__sortStageSetInstances(stageSet.NoteStage.Extends, stageSet.NoteStage.Extend_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			extendIdent := "__note" + extend.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Extend{Name: %s}).Stage(stageSet.NoteStage)", extendIdent, __gong__toRawStringLiteral(extend.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, figure := range __gong__sortStageSetInstances(stageSet.NoteStage.Figures, stageSet.NoteStage.Figure_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			figureIdent := "__note" + figure.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Figure{Name: %s}).Stage(stageSet.NoteStage)", figureIdent, __gong__toRawStringLiteral(figure.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", figureIdent, __gong__toRawStringLiteral(figure.Name)))
			if figure.Prefix != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__direction" + figure.Prefix.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Prefix = %s", figureIdent, targetIdent))
			}
			if figure.FigureNumber != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__direction" + figure.FigureNumber.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FigureNumber = %s", figureIdent, targetIdent))
			}
			if figure.Suffix != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__direction" + figure.Suffix.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Suffix = %s", figureIdent, targetIdent))
			}
			if figure.Extend != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + figure.Extend.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extend = %s", figureIdent, targetIdent))
			}
			if figure.Footnote != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + figure.Footnote.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", figureIdent, targetIdent))
			}
			if figure.Level != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + figure.Level.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", figureIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, figuredbass := range __gong__sortStageSetInstances(stageSet.NoteStage.FiguredBasss, stageSet.NoteStage.FiguredBass_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			figuredbassIdent := "__note" + figuredbass.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.FiguredBass{Name: %s}).Stage(stageSet.NoteStage)", figuredbassIdent, __gong__toRawStringLiteral(figuredbass.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Figure = append(%s.Figure, %s)", figuredbassIdent, figuredbassIdent, targetIdent))
			}
			if figuredbass.Footnote != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + figuredbass.Footnote.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", figuredbassIdent, targetIdent))
			}
			if figuredbass.Level != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + figuredbass.Level.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", figuredbassIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, forward := range __gong__sortStageSetInstances(stageSet.NoteStage.Forwards, stageSet.NoteStage.Forward_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			forwardIdent := "__note" + forward.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Forward{Name: %s}).Stage(stageSet.NoteStage)", forwardIdent, __gong__toRawStringLiteral(forward.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", forwardIdent, __gong__toRawStringLiteral(forward.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Duration = %s", forwardIdent, __gong__toRawStringLiteral(forward.Duration)))
			values.WriteString(fmt.Sprintf("\n\t%s.Voice = %s", forwardIdent, __gong__toRawStringLiteral(forward.Voice)))
			values.WriteString(fmt.Sprintf("\n\t%s.Staff = %d", forwardIdent, forward.Staff))
			if forward.Footnote != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + forward.Footnote.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", forwardIdent, targetIdent))
			}
			if forward.Level != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + forward.Level.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", forwardIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, glissando := range __gong__sortStageSetInstances(stageSet.NoteStage.Glissandos, stageSet.NoteStage.Glissando_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			glissandoIdent := "__note" + glissando.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Glissando{Name: %s}).Stage(stageSet.NoteStage)", glissandoIdent, __gong__toRawStringLiteral(glissando.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, grace := range __gong__sortStageSetInstances(stageSet.NoteStage.Graces, stageSet.NoteStage.Grace_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			graceIdent := "__note" + grace.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Grace{Name: %s}).Stage(stageSet.NoteStage)", graceIdent, __gong__toRawStringLiteral(grace.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", graceIdent, __gong__toRawStringLiteral(grace.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.StealTimePrevious = %s", graceIdent, __gong__toRawStringLiteral(grace.StealTimePrevious)))
			values.WriteString(fmt.Sprintf("\n\t%s.StealTimeFollowing = %s", graceIdent, __gong__toRawStringLiteral(grace.StealTimeFollowing)))
			values.WriteString(fmt.Sprintf("\n\t%s.MakeTime = %s", graceIdent, __gong__toRawStringLiteral(grace.MakeTime)))
			values.WriteString(fmt.Sprintf("\n\t%s.Slash = %s", graceIdent, __gong__toRawStringLiteral(string(grace.Slash))))
		}
	}
	if stageSet.NoteStage != nil {
		for _, hammeronpulloff := range __gong__sortStageSetInstances(stageSet.NoteStage.HammerOnPullOffs, stageSet.NoteStage.HammerOnPullOff_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			hammeronpulloffIdent := "__note" + hammeronpulloff.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.HammerOnPullOff{Name: %s}).Stage(stageSet.NoteStage)", hammeronpulloffIdent, __gong__toRawStringLiteral(hammeronpulloff.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, handbell := range __gong__sortStageSetInstances(stageSet.NoteStage.Handbells, stageSet.NoteStage.Handbell_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			handbellIdent := "__note" + handbell.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Handbell{Name: %s}).Stage(stageSet.NoteStage)", handbellIdent, __gong__toRawStringLiteral(handbell.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, harmonclosed := range __gong__sortStageSetInstances(stageSet.NoteStage.HarmonCloseds, stageSet.NoteStage.HarmonClosed_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			harmonclosedIdent := "__note" + harmonclosed.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.HarmonClosed{Name: %s}).Stage(stageSet.NoteStage)", harmonclosedIdent, __gong__toRawStringLiteral(harmonclosed.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", harmonclosedIdent, __gong__toRawStringLiteral(harmonclosed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", harmonclosedIdent, __gong__toRawStringLiteral(string(harmonclosed.Location))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", harmonclosedIdent, __gong__toRawStringLiteral(string(harmonclosed.Value))))
		}
	}
	if stageSet.NoteStage != nil {
		for _, harmonmute := range __gong__sortStageSetInstances(stageSet.NoteStage.HarmonMutes, stageSet.NoteStage.HarmonMute_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			harmonmuteIdent := "__note" + harmonmute.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.HarmonMute{Name: %s}).Stage(stageSet.NoteStage)", harmonmuteIdent, __gong__toRawStringLiteral(harmonmute.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + harmonmute.HarmonClosed.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HarmonClosed = %s", harmonmuteIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, harmonic := range __gong__sortStageSetInstances(stageSet.NoteStage.Harmonics, stageSet.NoteStage.Harmonic_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			harmonicIdent := "__note" + harmonic.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Harmonic{Name: %s}).Stage(stageSet.NoteStage)", harmonicIdent, __gong__toRawStringLiteral(harmonic.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
			if harmonic.Natural != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + harmonic.Natural.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Natural = %s", harmonicIdent, targetIdent))
			}
			if harmonic.Artificial != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + harmonic.Artificial.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Artificial = %s", harmonicIdent, targetIdent))
			}
			if harmonic.BasePitch != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + harmonic.BasePitch.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BasePitch = %s", harmonicIdent, targetIdent))
			}
			if harmonic.TouchingPitch != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + harmonic.TouchingPitch.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TouchingPitch = %s", harmonicIdent, targetIdent))
			}
			if harmonic.SoundingPitch != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + harmonic.SoundingPitch.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SoundingPitch = %s", harmonicIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, heeltoe := range __gong__sortStageSetInstances(stageSet.NoteStage.HeelToes, stageSet.NoteStage.HeelToe_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			heeltoeIdent := "__note" + heeltoe.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.HeelToe{Name: %s}).Stage(stageSet.NoteStage)", heeltoeIdent, __gong__toRawStringLiteral(heeltoe.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, hole := range __gong__sortStageSetInstances(stageSet.NoteStage.Holes, stageSet.NoteStage.Hole_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			holeIdent := "__note" + hole.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Hole{Name: %s}).Stage(stageSet.NoteStage)", holeIdent, __gong__toRawStringLiteral(hole.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + hole.HoleClosed.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HoleClosed = %s", holeIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, holeclosed := range __gong__sortStageSetInstances(stageSet.NoteStage.HoleCloseds, stageSet.NoteStage.HoleClosed_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			holeclosedIdent := "__note" + holeclosed.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.HoleClosed{Name: %s}).Stage(stageSet.NoteStage)", holeclosedIdent, __gong__toRawStringLiteral(holeclosed.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", holeclosedIdent, __gong__toRawStringLiteral(holeclosed.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Location = %s", holeclosedIdent, __gong__toRawStringLiteral(string(holeclosed.Location))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", holeclosedIdent, __gong__toRawStringLiteral(string(holeclosed.Value))))
		}
	}
	if stageSet.NoteStage != nil {
		for _, instrument := range __gong__sortStageSetInstances(stageSet.NoteStage.Instruments, stageSet.NoteStage.Instrument_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			instrumentIdent := "__note" + instrument.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Instrument{Name: %s}).Stage(stageSet.NoteStage)", instrumentIdent, __gong__toRawStringLiteral(instrument.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrumentIdent, __gong__toRawStringLiteral(instrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", instrumentIdent, __gong__toRawStringLiteral(instrument.Id)))
		}
	}
	if stageSet.NoteStage != nil {
		for _, listen := range __gong__sortStageSetInstances(stageSet.NoteStage.Listens, stageSet.NoteStage.Listen_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			listenIdent := "__note" + listen.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Listen{Name: %s}).Stage(stageSet.NoteStage)", listenIdent, __gong__toRawStringLiteral(listen.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", listenIdent, __gong__toRawStringLiteral(listen.Name)))
			for _, elem := range listen.Assess {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Assess = append(%s.Assess, %s)", listenIdent, listenIdent, targetIdent))
			}
			for _, elem := range listen.Wait {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Wait = append(%s.Wait, %s)", listenIdent, listenIdent, targetIdent))
			}
			for _, elem := range listen.OtherListen {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherListen = append(%s.OtherListen, %s)", listenIdent, listenIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, lyric := range __gong__sortStageSetInstances(stageSet.NoteStage.Lyrics, stageSet.NoteStage.Lyric_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			lyricIdent := "__note" + lyric.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Lyric{Name: %s}).Stage(stageSet.NoteStage)", lyricIdent, __gong__toRawStringLiteral(lyric.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + lyric.Text.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Text = %s", lyricIdent, targetIdent))
			}
			for _, elem := range lyric.Elision {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Elision = append(%s.Elision, %s)", lyricIdent, lyricIdent, targetIdent))
			}
			if lyric.Extend != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + lyric.Extend.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extend = %s", lyricIdent, targetIdent))
			}
			if lyric.Laughing != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + lyric.Laughing.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Laughing = %s", lyricIdent, targetIdent))
			}
			if lyric.Humming != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + lyric.Humming.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Humming = %s", lyricIdent, targetIdent))
			}
			if lyric.EndLine != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + lyric.EndLine.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndLine = %s", lyricIdent, targetIdent))
			}
			if lyric.EndParagraph != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + lyric.EndParagraph.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EndParagraph = %s", lyricIdent, targetIdent))
			}
			if lyric.Footnote != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + lyric.Footnote.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", lyricIdent, targetIdent))
			}
			if lyric.Level != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + lyric.Level.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", lyricIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, mordent := range __gong__sortStageSetInstances(stageSet.NoteStage.Mordents, stageSet.NoteStage.Mordent_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			mordentIdent := "__note" + mordent.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Mordent{Name: %s}).Stage(stageSet.NoteStage)", mordentIdent, __gong__toRawStringLiteral(mordent.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, nonarpeggiate := range __gong__sortStageSetInstances(stageSet.NoteStage.NonArpeggiates, stageSet.NoteStage.NonArpeggiate_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			nonarpeggiateIdent := "__note" + nonarpeggiate.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.NonArpeggiate{Name: %s}).Stage(stageSet.NoteStage)", nonarpeggiateIdent, __gong__toRawStringLiteral(nonarpeggiate.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, notations := range __gong__sortStageSetInstances(stageSet.NoteStage.Notationss, stageSet.NoteStage.Notations_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			notationsIdent := "__note" + notations.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Notations{Name: %s}).Stage(stageSet.NoteStage)", notationsIdent, __gong__toRawStringLiteral(notations.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notationsIdent, __gong__toRawStringLiteral(notations.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.PrintObject = %s", notationsIdent, __gong__toRawStringLiteral(string(notations.PrintObject))))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", notationsIdent, __gong__toRawStringLiteral(notations.Id)))
			if notations.Footnote != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + notations.Footnote.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", notationsIdent, targetIdent))
			}
			if notations.Level != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + notations.Level.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", notationsIdent, targetIdent))
			}
			for _, elem := range notations.Tied {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tied = append(%s.Tied, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Slur {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slur = append(%s.Slur, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Tuplet {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tuplet = append(%s.Tuplet, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Glissando {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Glissando = append(%s.Glissando, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Slide {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Slide = append(%s.Slide, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Ornaments {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Ornaments = append(%s.Ornaments, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Technical {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Technical = append(%s.Technical, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Articulations {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Articulations = append(%s.Articulations, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Dynamics {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dynamics = append(%s.Dynamics, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Fermata {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fermata = append(%s.Fermata, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.Arpeggiate {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arpeggiate = append(%s.Arpeggiate, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.NonArpeggiate {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NonArpeggiate = append(%s.NonArpeggiate, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.AccidentalMark {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AccidentalMark = append(%s.AccidentalMark, %s)", notationsIdent, notationsIdent, targetIdent))
			}
			for _, elem := range notations.OtherNotation {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherNotation = append(%s.OtherNotation, %s)", notationsIdent, notationsIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, note := range __gong__sortStageSetInstances(stageSet.NoteStage.Notes, stageSet.NoteStage.Note_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			noteIdent := "__note" + note.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Note{Name: %s}).Stage(stageSet.NoteStage)", noteIdent, __gong__toRawStringLiteral(note.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + note.Grace.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Grace = %s", noteIdent, targetIdent))
			}
			if note.Chord != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + note.Chord.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Chord = %s", noteIdent, targetIdent))
			}
			if note.Pitch != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + note.Pitch.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pitch = %s", noteIdent, targetIdent))
			}
			if note.Unpitched != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + note.Unpitched.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Unpitched = %s", noteIdent, targetIdent))
			}
			if note.Rest != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + note.Rest.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Rest = %s", noteIdent, targetIdent))
			}
			for _, elem := range note.Tie {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tie = append(%s.Tie, %s)", noteIdent, noteIdent, targetIdent))
			}
			if note.Cue != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + note.Cue.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Cue = %s", noteIdent, targetIdent))
			}
			for _, elem := range note.Instrument {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Instrument = append(%s.Instrument, %s)", noteIdent, noteIdent, targetIdent))
			}
			if note.Footnote != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + note.Footnote.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", noteIdent, targetIdent))
			}
			if note.Level != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + note.Level.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", noteIdent, targetIdent))
			}
			if note.Type != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + note.Type.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Type = %s", noteIdent, targetIdent))
			}
			for _, elem := range note.Dot {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Dot = append(%s.Dot, %s)", noteIdent, noteIdent, targetIdent))
			}
			if note.Accidental != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + note.Accidental.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Accidental = %s", noteIdent, targetIdent))
			}
			if note.TimeModification != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__direction" + note.TimeModification.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeModification = %s", noteIdent, targetIdent))
			}
			if note.Stem != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + note.Stem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stem = %s", noteIdent, targetIdent))
			}
			if note.Notehead != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + note.Notehead.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notehead = %s", noteIdent, targetIdent))
			}
			if note.NoteheadText != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + note.NoteheadText.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.NoteheadText = %s", noteIdent, targetIdent))
			}
			for _, elem := range note.Beam {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Beam = append(%s.Beam, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Notations {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Notations = append(%s.Notations, %s)", noteIdent, noteIdent, targetIdent))
			}
			for _, elem := range note.Lyric {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lyric = append(%s.Lyric, %s)", noteIdent, noteIdent, targetIdent))
			}
			if note.Play != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + note.Play.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Play = %s", noteIdent, targetIdent))
			}
			if note.Listen != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + note.Listen.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listen = %s", noteIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, notetype := range __gong__sortStageSetInstances(stageSet.NoteStage.NoteTypes, stageSet.NoteStage.NoteType_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			notetypeIdent := "__note" + notetype.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.NoteType{Name: %s}).Stage(stageSet.NoteStage)", notetypeIdent, __gong__toRawStringLiteral(notetype.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notetypeIdent, __gong__toRawStringLiteral(notetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Size = %s", notetypeIdent, __gong__toRawStringLiteral(string(notetype.Size))))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", notetypeIdent, __gong__toRawStringLiteral(string(notetype.Value))))
		}
	}
	if stageSet.NoteStage != nil {
		for _, notehead := range __gong__sortStageSetInstances(stageSet.NoteStage.Noteheads, stageSet.NoteStage.Notehead_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			noteheadIdent := "__note" + notehead.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Notehead{Name: %s}).Stage(stageSet.NoteStage)", noteheadIdent, __gong__toRawStringLiteral(notehead.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, noteheadtext := range __gong__sortStageSetInstances(stageSet.NoteStage.NoteheadTexts, stageSet.NoteStage.NoteheadText_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			noteheadtextIdent := "__note" + noteheadtext.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.NoteheadText{Name: %s}).Stage(stageSet.NoteStage)", noteheadtextIdent, __gong__toRawStringLiteral(noteheadtext.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", noteheadtextIdent, __gong__toRawStringLiteral(noteheadtext.Name)))
			for _, elem := range noteheadtext.DisplayText {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DisplayText = append(%s.DisplayText, %s)", noteheadtextIdent, noteheadtextIdent, targetIdent))
			}
			for _, elem := range noteheadtext.AccidentalText {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AccidentalText = append(%s.AccidentalText, %s)", noteheadtextIdent, noteheadtextIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, ornaments := range __gong__sortStageSetInstances(stageSet.NoteStage.Ornamentss, stageSet.NoteStage.Ornaments_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			ornamentsIdent := "__note" + ornaments.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Ornaments{Name: %s}).Stage(stageSet.NoteStage)", ornamentsIdent, __gong__toRawStringLiteral(ornaments.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", ornamentsIdent, __gong__toRawStringLiteral(ornaments.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", ornamentsIdent, __gong__toRawStringLiteral(ornaments.Id)))
			for _, elem := range ornaments.TrillMark {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TrillMark = append(%s.TrillMark, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Turn {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Turn = append(%s.Turn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.DelayedTurn {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DelayedTurn = append(%s.DelayedTurn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.InvertedTurn {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InvertedTurn = append(%s.InvertedTurn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.DelayedInvertedTurn {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DelayedInvertedTurn = append(%s.DelayedInvertedTurn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.VerticalTurn {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalTurn = append(%s.VerticalTurn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.InvertedVerticalTurn {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InvertedVerticalTurn = append(%s.InvertedVerticalTurn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Shake {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Shake = append(%s.Shake, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.WavyLine {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WavyLine = append(%s.WavyLine, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Mordent {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Mordent = append(%s.Mordent, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.InvertedMordent {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InvertedMordent = append(%s.InvertedMordent, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Schleifer {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Schleifer = append(%s.Schleifer, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Tremolo {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tremolo = append(%s.Tremolo, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.Haydn {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Haydn = append(%s.Haydn, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.OtherOrnament {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherOrnament = append(%s.OtherOrnament, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
			for _, elem := range ornaments.AccidentalMark {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AccidentalMark = append(%s.AccidentalMark, %s)", ornamentsIdent, ornamentsIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, othernotation := range __gong__sortStageSetInstances(stageSet.NoteStage.OtherNotations, stageSet.NoteStage.OtherNotation_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			othernotationIdent := "__note" + othernotation.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.OtherNotation{Name: %s}).Stage(stageSet.NoteStage)", othernotationIdent, __gong__toRawStringLiteral(othernotation.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, otherplacementtext := range __gong__sortStageSetInstances(stageSet.NoteStage.OtherPlacementTexts, stageSet.NoteStage.OtherPlacementText_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			otherplacementtextIdent := "__note" + otherplacementtext.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.OtherPlacementText{Name: %s}).Stage(stageSet.NoteStage)", otherplacementtextIdent, __gong__toRawStringLiteral(otherplacementtext.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, pitch := range __gong__sortStageSetInstances(stageSet.NoteStage.Pitchs, stageSet.NoteStage.Pitch_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			pitchIdent := "__note" + pitch.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Pitch{Name: %s}).Stage(stageSet.NoteStage)", pitchIdent, __gong__toRawStringLiteral(pitch.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pitchIdent, __gong__toRawStringLiteral(pitch.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Step = %s", pitchIdent, __gong__toRawStringLiteral(string(pitch.Step))))
			values.WriteString(fmt.Sprintf("\n\t%s.Alter = %s", pitchIdent, __gong__toRawStringLiteral(pitch.Alter)))
			values.WriteString(fmt.Sprintf("\n\t%s.Octave = %d", pitchIdent, pitch.Octave))
		}
	}
	if stageSet.NoteStage != nil {
		for _, placementtext := range __gong__sortStageSetInstances(stageSet.NoteStage.PlacementTexts, stageSet.NoteStage.PlacementText_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			placementtextIdent := "__note" + placementtext.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.PlacementText{Name: %s}).Stage(stageSet.NoteStage)", placementtextIdent, __gong__toRawStringLiteral(placementtext.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, release := range __gong__sortStageSetInstances(stageSet.NoteStage.Releases, stageSet.NoteStage.Release_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			releaseIdent := "__note" + release.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Release{Name: %s}).Stage(stageSet.NoteStage)", releaseIdent, __gong__toRawStringLiteral(release.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", releaseIdent, __gong__toRawStringLiteral(release.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Offset = %s", releaseIdent, __gong__toRawStringLiteral(release.Offset)))
		}
	}
	if stageSet.NoteStage != nil {
		for _, rest := range __gong__sortStageSetInstances(stageSet.NoteStage.Rests, stageSet.NoteStage.Rest_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			restIdent := "__note" + rest.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Rest{Name: %s}).Stage(stageSet.NoteStage)", restIdent, __gong__toRawStringLiteral(rest.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", restIdent, __gong__toRawStringLiteral(rest.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Measure = %s", restIdent, __gong__toRawStringLiteral(string(rest.Measure))))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayStep = %s", restIdent, __gong__toRawStringLiteral(string(rest.DisplayStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayOctave = %d", restIdent, rest.DisplayOctave))
		}
	}
	if stageSet.NoteStage != nil {
		for _, slide := range __gong__sortStageSetInstances(stageSet.NoteStage.Slides, stageSet.NoteStage.Slide_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			slideIdent := "__note" + slide.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Slide{Name: %s}).Stage(stageSet.NoteStage)", slideIdent, __gong__toRawStringLiteral(slide.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, slur := range __gong__sortStageSetInstances(stageSet.NoteStage.Slurs, stageSet.NoteStage.Slur_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			slurIdent := "__note" + slur.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Slur{Name: %s}).Stage(stageSet.NoteStage)", slurIdent, __gong__toRawStringLiteral(slur.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, stem := range __gong__sortStageSetInstances(stageSet.NoteStage.Stems, stageSet.NoteStage.Stem_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			stemIdent := "__note" + stem.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Stem{Name: %s}).Stage(stageSet.NoteStage)", stemIdent, __gong__toRawStringLiteral(stem.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, strongaccent := range __gong__sortStageSetInstances(stageSet.NoteStage.StrongAccents, stageSet.NoteStage.StrongAccent_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			strongaccentIdent := "__note" + strongaccent.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.StrongAccent{Name: %s}).Stage(stageSet.NoteStage)", strongaccentIdent, __gong__toRawStringLiteral(strongaccent.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, tap := range __gong__sortStageSetInstances(stageSet.NoteStage.Taps, stageSet.NoteStage.Tap_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			tapIdent := "__note" + tap.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Tap{Name: %s}).Stage(stageSet.NoteStage)", tapIdent, __gong__toRawStringLiteral(tap.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, technical := range __gong__sortStageSetInstances(stageSet.NoteStage.Technicals, stageSet.NoteStage.Technical_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			technicalIdent := "__note" + technical.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Technical{Name: %s}).Stage(stageSet.NoteStage)", technicalIdent, __gong__toRawStringLiteral(technical.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", technicalIdent, __gong__toRawStringLiteral(technical.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", technicalIdent, __gong__toRawStringLiteral(technical.Id)))
			for _, elem := range technical.UpBow {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.UpBow = append(%s.UpBow, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.DownBow {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DownBow = append(%s.DownBow, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Harmonic {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmonic = append(%s.Harmonic, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.OpenString {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OpenString = append(%s.OpenString, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.ThumbPosition {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ThumbPosition = append(%s.ThumbPosition, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Fingering {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fingering = append(%s.Fingering, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Pluck {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pluck = append(%s.Pluck, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.DoubleTongue {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DoubleTongue = append(%s.DoubleTongue, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.TripleTongue {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TripleTongue = append(%s.TripleTongue, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Stopped {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Stopped = append(%s.Stopped, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.SnapPizzicato {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SnapPizzicato = append(%s.SnapPizzicato, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Fret {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fret = append(%s.Fret, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.String {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.String = append(%s.String, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.HammerOn {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HammerOn = append(%s.HammerOn, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.PullOff {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PullOff = append(%s.PullOff, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Bend {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bend = append(%s.Bend, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Tap {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tap = append(%s.Tap, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Heel {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Heel = append(%s.Heel, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Toe {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Toe = append(%s.Toe, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Fingernails {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Fingernails = append(%s.Fingernails, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Hole {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Hole = append(%s.Hole, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Arrow {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Arrow = append(%s.Arrow, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Handbell {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Handbell = append(%s.Handbell, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.BrassBend {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.BrassBend = append(%s.BrassBend, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Flip {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Flip = append(%s.Flip, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Smear {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Smear = append(%s.Smear, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Open {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Open = append(%s.Open, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.HalfMuted {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HalfMuted = append(%s.HalfMuted, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.HarmonMute {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HarmonMute = append(%s.HarmonMute, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.Golpe {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Golpe = append(%s.Golpe, %s)", technicalIdent, technicalIdent, targetIdent))
			}
			for _, elem := range technical.OtherTechnical {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OtherTechnical = append(%s.OtherTechnical, %s)", technicalIdent, technicalIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, textelementdata := range __gong__sortStageSetInstances(stageSet.NoteStage.TextElementDatas, stageSet.NoteStage.TextElementData_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			textelementdataIdent := "__note" + textelementdata.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.TextElementData{Name: %s}).Stage(stageSet.NoteStage)", textelementdataIdent, __gong__toRawStringLiteral(textelementdata.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, tie := range __gong__sortStageSetInstances(stageSet.NoteStage.Ties, stageSet.NoteStage.Tie_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			tieIdent := "__note" + tie.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Tie{Name: %s}).Stage(stageSet.NoteStage)", tieIdent, __gong__toRawStringLiteral(tie.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tieIdent, __gong__toRawStringLiteral(tie.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", tieIdent, __gong__toRawStringLiteral(string(tie.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeOnly = %s", tieIdent, __gong__toRawStringLiteral(tie.TimeOnly)))
		}
	}
	if stageSet.NoteStage != nil {
		for _, tied := range __gong__sortStageSetInstances(stageSet.NoteStage.Tieds, stageSet.NoteStage.Tied_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			tiedIdent := "__note" + tied.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Tied{Name: %s}).Stage(stageSet.NoteStage)", tiedIdent, __gong__toRawStringLiteral(tied.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, tremolo := range __gong__sortStageSetInstances(stageSet.NoteStage.Tremolos, stageSet.NoteStage.Tremolo_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			tremoloIdent := "__note" + tremolo.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Tremolo{Name: %s}).Stage(stageSet.NoteStage)", tremoloIdent, __gong__toRawStringLiteral(tremolo.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, tuplet := range __gong__sortStageSetInstances(stageSet.NoteStage.Tuplets, stageSet.NoteStage.Tuplet_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			tupletIdent := "__note" + tuplet.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Tuplet{Name: %s}).Stage(stageSet.NoteStage)", tupletIdent, __gong__toRawStringLiteral(tuplet.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + tuplet.TupletActual.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TupletActual = %s", tupletIdent, targetIdent))
			}
			if tuplet.TupletNormal != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + tuplet.TupletNormal.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TupletNormal = %s", tupletIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, tupletdot := range __gong__sortStageSetInstances(stageSet.NoteStage.TupletDots, stageSet.NoteStage.TupletDot_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			tupletdotIdent := "__note" + tupletdot.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.TupletDot{Name: %s}).Stage(stageSet.NoteStage)", tupletdotIdent, __gong__toRawStringLiteral(tupletdot.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tupletdotIdent, __gong__toRawStringLiteral(tupletdot.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", tupletdotIdent, __gong__toRawStringLiteral(tupletdot.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", tupletdotIdent, __gong__toRawStringLiteral(string(tupletdot.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", tupletdotIdent, __gong__toRawStringLiteral(tupletdot.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", tupletdotIdent, __gong__toRawStringLiteral(string(tupletdot.FontWeight))))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", tupletdotIdent, __gong__toRawStringLiteral(tupletdot.Color)))
		}
	}
	if stageSet.NoteStage != nil {
		for _, tupletnumber := range __gong__sortStageSetInstances(stageSet.NoteStage.TupletNumbers, stageSet.NoteStage.TupletNumber_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			tupletnumberIdent := "__note" + tupletnumber.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.TupletNumber{Name: %s}).Stage(stageSet.NoteStage)", tupletnumberIdent, __gong__toRawStringLiteral(tupletnumber.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, tupletportion := range __gong__sortStageSetInstances(stageSet.NoteStage.TupletPortions, stageSet.NoteStage.TupletPortion_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			tupletportionIdent := "__note" + tupletportion.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.TupletPortion{Name: %s}).Stage(stageSet.NoteStage)", tupletportionIdent, __gong__toRawStringLiteral(tupletportion.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tupletportionIdent, __gong__toRawStringLiteral(tupletportion.Name)))
			if tupletportion.TupletNumber != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + tupletportion.TupletNumber.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TupletNumber = %s", tupletportionIdent, targetIdent))
			}
			if tupletportion.TupletType != nil {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + tupletportion.TupletType.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TupletType = %s", tupletportionIdent, targetIdent))
			}
			for _, elem := range tupletportion.TupletDot {
				if lastStagePtr != "NoteStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "NoteStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TupletDot = append(%s.TupletDot, %s)", tupletportionIdent, tupletportionIdent, targetIdent))
			}
		}
	}
	if stageSet.NoteStage != nil {
		for _, tuplettype := range __gong__sortStageSetInstances(stageSet.NoteStage.TupletTypes, stageSet.NoteStage.TupletType_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			tuplettypeIdent := "__note" + tuplettype.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.TupletType{Name: %s}).Stage(stageSet.NoteStage)", tuplettypeIdent, __gong__toRawStringLiteral(tuplettype.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
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
	if stageSet.NoteStage != nil {
		for _, unpitched := range __gong__sortStageSetInstances(stageSet.NoteStage.Unpitcheds, stageSet.NoteStage.Unpitched_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			unpitchedIdent := "__note" + unpitched.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Unpitched{Name: %s}).Stage(stageSet.NoteStage)", unpitchedIdent, __gong__toRawStringLiteral(unpitched.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unpitchedIdent, __gong__toRawStringLiteral(unpitched.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayStep = %s", unpitchedIdent, __gong__toRawStringLiteral(string(unpitched.DisplayStep))))
			values.WriteString(fmt.Sprintf("\n\t%s.DisplayOctave = %d", unpitchedIdent, unpitched.DisplayOctave))
		}
	}
	if stageSet.NoteStage != nil {
		for _, wait := range __gong__sortStageSetInstances(stageSet.NoteStage.Waits, stageSet.NoteStage.Wait_stagedOrder) {
			if lastStageDecl != "NoteStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "NoteStage"
			}
			waitIdent := "__note" + wait.GongGetIdentifier(stageSet.NoteStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&note.Wait{Name: %s}).Stage(stageSet.NoteStage)", waitIdent, __gong__toRawStringLiteral(wait.Name)))
			if lastStageVal != "NoteStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "NoteStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", waitIdent, __gong__toRawStringLiteral(wait.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Player = %s", waitIdent, __gong__toRawStringLiteral(wait.Player)))
			values.WriteString(fmt.Sprintf("\n\t%s.TimeOnly = %s", waitIdent, __gong__toRawStringLiteral(wait.TimeOnly)))
		}
	}
	if stageSet.ScoreStage != nil {
		for _, credit := range __gong__sortStageSetInstances(stageSet.ScoreStage.Credits, stageSet.ScoreStage.Credit_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			creditIdent := "__score" + credit.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Credit{Name: %s}).Stage(stageSet.ScoreStage)", creditIdent, __gong__toRawStringLiteral(credit.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", creditIdent, __gong__toRawStringLiteral(credit.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Page = %d", creditIdent, credit.Page))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", creditIdent, __gong__toRawStringLiteral(credit.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.CreditType = %s", creditIdent, __gong__toRawStringLiteral(credit.CreditType)))
			for _, elem := range credit.Link {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Link = append(%s.Link, %s)", creditIdent, creditIdent, targetIdent))
			}
			for _, elem := range credit.Bookmark {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bookmark = append(%s.Bookmark, %s)", creditIdent, creditIdent, targetIdent))
			}
			if credit.CreditImage != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + credit.CreditImage.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CreditImage = %s", creditIdent, targetIdent))
			}
			for _, elem := range credit.CreditWords {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CreditWords = append(%s.CreditWords, %s)", creditIdent, creditIdent, targetIdent))
			}
			for _, elem := range credit.CreditSymbol {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CreditSymbol = append(%s.CreditSymbol, %s)", creditIdent, creditIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, defaults := range __gong__sortStageSetInstances(stageSet.ScoreStage.Defaultss, stageSet.ScoreStage.Defaults_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			defaultsIdent := "__score" + defaults.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Defaults{Name: %s}).Stage(stageSet.ScoreStage)", defaultsIdent, __gong__toRawStringLiteral(defaults.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", defaultsIdent, __gong__toRawStringLiteral(defaults.Name)))
			if defaults.Scaling != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + defaults.Scaling.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Scaling = %s", defaultsIdent, targetIdent))
			}
			if defaults.ConcertScore != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + defaults.ConcertScore.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ConcertScore = %s", defaultsIdent, targetIdent))
			}
			if defaults.PageLayout != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + defaults.PageLayout.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PageLayout = %s", defaultsIdent, targetIdent))
			}
			if defaults.SystemLayout != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + defaults.SystemLayout.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SystemLayout = %s", defaultsIdent, targetIdent))
			}
			for _, elem := range defaults.StaffLayout {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StaffLayout = append(%s.StaffLayout, %s)", defaultsIdent, defaultsIdent, targetIdent))
			}
			if defaults.Appearance != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + defaults.Appearance.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Appearance = %s", defaultsIdent, targetIdent))
			}
			if defaults.MusicFont != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + defaults.MusicFont.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MusicFont = %s", defaultsIdent, targetIdent))
			}
			if defaults.WordFont != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + defaults.WordFont.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.WordFont = %s", defaultsIdent, targetIdent))
			}
			for _, elem := range defaults.LyricFont {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LyricFont = append(%s.LyricFont, %s)", defaultsIdent, defaultsIdent, targetIdent))
			}
			for _, elem := range defaults.LyricLanguage {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LyricLanguage = append(%s.LyricLanguage, %s)", defaultsIdent, defaultsIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, emptyfont := range __gong__sortStageSetInstances(stageSet.ScoreStage.EmptyFonts, stageSet.ScoreStage.EmptyFont_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			emptyfontIdent := "__score" + emptyfont.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.EmptyFont{Name: %s}).Stage(stageSet.ScoreStage)", emptyfontIdent, __gong__toRawStringLiteral(emptyfont.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", emptyfontIdent, __gong__toRawStringLiteral(emptyfont.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontFamily = %s", emptyfontIdent, __gong__toRawStringLiteral(emptyfont.FontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontStyle = %s", emptyfontIdent, __gong__toRawStringLiteral(string(emptyfont.FontStyle))))
			values.WriteString(fmt.Sprintf("\n\t%s.FontSize = %s", emptyfontIdent, __gong__toRawStringLiteral(emptyfont.FontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.FontWeight = %s", emptyfontIdent, __gong__toRawStringLiteral(string(emptyfont.FontWeight))))
		}
	}
	if stageSet.ScoreStage != nil {
		for _, groupbarline := range __gong__sortStageSetInstances(stageSet.ScoreStage.GroupBarlines, stageSet.ScoreStage.GroupBarline_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			groupbarlineIdent := "__score" + groupbarline.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.GroupBarline{Name: %s}).Stage(stageSet.ScoreStage)", groupbarlineIdent, __gong__toRawStringLiteral(groupbarline.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groupbarlineIdent, __gong__toRawStringLiteral(groupbarline.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Color = %s", groupbarlineIdent, __gong__toRawStringLiteral(groupbarline.Color)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %s", groupbarlineIdent, __gong__toRawStringLiteral(string(groupbarline.Value))))
		}
	}
	if stageSet.ScoreStage != nil {
		for _, groupname := range __gong__sortStageSetInstances(stageSet.ScoreStage.GroupNames, stageSet.ScoreStage.GroupName_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			groupnameIdent := "__score" + groupname.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.GroupName{Name: %s}).Stage(stageSet.ScoreStage)", groupnameIdent, __gong__toRawStringLiteral(groupname.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
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
	if stageSet.ScoreStage != nil {
		for _, groupsymbol := range __gong__sortStageSetInstances(stageSet.ScoreStage.GroupSymbols, stageSet.ScoreStage.GroupSymbol_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			groupsymbolIdent := "__score" + groupsymbol.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.GroupSymbol{Name: %s}).Stage(stageSet.ScoreStage)", groupsymbolIdent, __gong__toRawStringLiteral(groupsymbol.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
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
	if stageSet.ScoreStage != nil {
		for _, instrumentlink := range __gong__sortStageSetInstances(stageSet.ScoreStage.InstrumentLinks, stageSet.ScoreStage.InstrumentLink_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			instrumentlinkIdent := "__score" + instrumentlink.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.InstrumentLink{Name: %s}).Stage(stageSet.ScoreStage)", instrumentlinkIdent, __gong__toRawStringLiteral(instrumentlink.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", instrumentlinkIdent, __gong__toRawStringLiteral(instrumentlink.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", instrumentlinkIdent, __gong__toRawStringLiteral(instrumentlink.Id)))
			if instrumentlink.RefScoreInstrument != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + instrumentlink.RefScoreInstrument.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RefScoreInstrument = %s", instrumentlinkIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, lyricfont := range __gong__sortStageSetInstances(stageSet.ScoreStage.LyricFonts, stageSet.ScoreStage.LyricFont_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			lyricfontIdent := "__score" + lyricfont.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.LyricFont{Name: %s}).Stage(stageSet.ScoreStage)", lyricfontIdent, __gong__toRawStringLiteral(lyricfont.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
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
	if stageSet.ScoreStage != nil {
		for _, lyriclanguage := range __gong__sortStageSetInstances(stageSet.ScoreStage.LyricLanguages, stageSet.ScoreStage.LyricLanguage_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			lyriclanguageIdent := "__score" + lyriclanguage.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.LyricLanguage{Name: %s}).Stage(stageSet.ScoreStage)", lyriclanguageIdent, __gong__toRawStringLiteral(lyriclanguage.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lyriclanguageIdent, __gong__toRawStringLiteral(lyriclanguage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", lyriclanguageIdent, __gong__toRawStringLiteral(lyriclanguage.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.NameAttr = %s", lyriclanguageIdent, __gong__toRawStringLiteral(lyriclanguage.NameAttr)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", lyriclanguageIdent, __gong__toRawStringLiteral(lyriclanguage.Lang)))
		}
	}
	if stageSet.ScoreStage != nil {
		for _, measure := range __gong__sortStageSetInstances(stageSet.ScoreStage.Measures, stageSet.ScoreStage.Measure_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			measureIdent := "__score" + measure.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Measure{Name: %s}).Stage(stageSet.ScoreStage)", measureIdent, __gong__toRawStringLiteral(measure.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measureIdent, __gong__toRawStringLiteral(measure.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", measureIdent, __gong__toRawStringLiteral(measure.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", measureIdent, __gong__toRawStringLiteral(measure.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Implicit = %s", measureIdent, __gong__toRawStringLiteral(string(measure.Implicit))))
			values.WriteString(fmt.Sprintf("\n\t%s.NonControlling = %s", measureIdent, __gong__toRawStringLiteral(string(measure.NonControlling))))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", measureIdent, __gong__toRawStringLiteral(measure.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", measureIdent, __gong__toRawStringLiteral(measure.Id)))
			for _, elem := range measure.MusicData {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MusicData = append(%s.MusicData, %s)", measureIdent, measureIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, measuremusicdatachoice := range __gong__sortStageSetInstances(stageSet.ScoreStage.MeasureMusicDataChoices, stageSet.ScoreStage.MeasureMusicDataChoice_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			measuremusicdatachoiceIdent := "__score" + measuremusicdatachoice.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.MeasureMusicDataChoice{Name: %s}).Stage(stageSet.ScoreStage)", measuremusicdatachoiceIdent, __gong__toRawStringLiteral(measuremusicdatachoice.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", measuremusicdatachoiceIdent, __gong__toRawStringLiteral(measuremusicdatachoice.Name)))
			if measuremusicdatachoice.Note != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__note" + measuremusicdatachoice.Note.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", measuremusicdatachoiceIdent, targetIdent))
			}
			if measuremusicdatachoice.Backup != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__note" + measuremusicdatachoice.Backup.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Backup = %s", measuremusicdatachoiceIdent, targetIdent))
			}
			if measuremusicdatachoice.Forward != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__note" + measuremusicdatachoice.Forward.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Forward = %s", measuremusicdatachoiceIdent, targetIdent))
			}
			if measuremusicdatachoice.Direction != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + measuremusicdatachoice.Direction.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Direction = %s", measuremusicdatachoiceIdent, targetIdent))
			}
			if measuremusicdatachoice.Attributes != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__attributes" + measuremusicdatachoice.Attributes.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = %s", measuremusicdatachoiceIdent, targetIdent))
			}
			if measuremusicdatachoice.Harmony != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + measuremusicdatachoice.Harmony.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmony = %s", measuremusicdatachoiceIdent, targetIdent))
			}
			if measuremusicdatachoice.FiguredBass != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__note" + measuremusicdatachoice.FiguredBass.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FiguredBass = %s", measuremusicdatachoiceIdent, targetIdent))
			}
			if measuremusicdatachoice.Print != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + measuremusicdatachoice.Print.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Print = %s", measuremusicdatachoiceIdent, targetIdent))
			}
			if measuremusicdatachoice.Sound != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + measuremusicdatachoice.Sound.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sound = %s", measuremusicdatachoiceIdent, targetIdent))
			}
			if measuremusicdatachoice.Listening != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + measuremusicdatachoice.Listening.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listening = %s", measuremusicdatachoiceIdent, targetIdent))
			}
			if measuremusicdatachoice.Barline != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__barline" + measuremusicdatachoice.Barline.GongGetIdentifier(stageSet.BarlineStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Barline = %s", measuremusicdatachoiceIdent, targetIdent))
			}
			if measuremusicdatachoice.Grouping != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + measuremusicdatachoice.Grouping.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Grouping = %s", measuremusicdatachoiceIdent, targetIdent))
			}
			if measuremusicdatachoice.Link != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + measuremusicdatachoice.Link.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Link = %s", measuremusicdatachoiceIdent, targetIdent))
			}
			if measuremusicdatachoice.Bookmark != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + measuremusicdatachoice.Bookmark.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bookmark = %s", measuremusicdatachoiceIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, opus := range __gong__sortStageSetInstances(stageSet.ScoreStage.Opuss, stageSet.ScoreStage.Opus_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			opusIdent := "__score" + opus.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Opus{Name: %s}).Stage(stageSet.ScoreStage)", opusIdent, __gong__toRawStringLiteral(opus.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
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
	if stageSet.ScoreStage != nil {
		for _, part := range __gong__sortStageSetInstances(stageSet.ScoreStage.Parts, stageSet.ScoreStage.Part_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			partIdent := "__score" + part.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Part{Name: %s}).Stage(stageSet.ScoreStage)", partIdent, __gong__toRawStringLiteral(part.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partIdent, __gong__toRawStringLiteral(part.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", partIdent, __gong__toRawStringLiteral(part.Id)))
			if part.RefScorePart != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + part.RefScorePart.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RefScorePart = %s", partIdent, targetIdent))
			}
			for _, elem := range part.Measure {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure = append(%s.Measure, %s)", partIdent, partIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, partgroup := range __gong__sortStageSetInstances(stageSet.ScoreStage.PartGroups, stageSet.ScoreStage.PartGroup_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			partgroupIdent := "__score" + partgroup.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.PartGroup{Name: %s}).Stage(stageSet.ScoreStage)", partgroupIdent, __gong__toRawStringLiteral(partgroup.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partgroupIdent, __gong__toRawStringLiteral(partgroup.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type = %s", partgroupIdent, __gong__toRawStringLiteral(string(partgroup.Type))))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", partgroupIdent, __gong__toRawStringLiteral(partgroup.Number)))
			if partgroup.GroupName != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + partgroup.GroupName.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupName = %s", partgroupIdent, targetIdent))
			}
			if partgroup.GroupNameDisplay != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + partgroup.GroupNameDisplay.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupNameDisplay = %s", partgroupIdent, targetIdent))
			}
			if partgroup.GroupAbbreviation != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + partgroup.GroupAbbreviation.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupAbbreviation = %s", partgroupIdent, targetIdent))
			}
			if partgroup.GroupAbbreviationDisplay != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + partgroup.GroupAbbreviationDisplay.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupAbbreviationDisplay = %s", partgroupIdent, targetIdent))
			}
			if partgroup.GroupSymbol != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + partgroup.GroupSymbol.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupSymbol = %s", partgroupIdent, targetIdent))
			}
			if partgroup.GroupBarline != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + partgroup.GroupBarline.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupBarline = %s", partgroupIdent, targetIdent))
			}
			if partgroup.GroupTime != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + partgroup.GroupTime.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GroupTime = %s", partgroupIdent, targetIdent))
			}
			if partgroup.Footnote != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + partgroup.Footnote.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Footnote = %s", partgroupIdent, targetIdent))
			}
			if partgroup.Level != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + partgroup.Level.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Level = %s", partgroupIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, partlink := range __gong__sortStageSetInstances(stageSet.ScoreStage.PartLinks, stageSet.ScoreStage.PartLink_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			partlinkIdent := "__score" + partlink.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.PartLink{Name: %s}).Stage(stageSet.ScoreStage)", partlinkIdent, __gong__toRawStringLiteral(partlink.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
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
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InstrumentLink = append(%s.InstrumentLink, %s)", partlinkIdent, partlinkIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, partlist := range __gong__sortStageSetInstances(stageSet.ScoreStage.PartLists, stageSet.ScoreStage.PartList_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			partlistIdent := "__score" + partlist.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.PartList{Name: %s}).Stage(stageSet.ScoreStage)", partlistIdent, __gong__toRawStringLiteral(partlist.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", partlistIdent, __gong__toRawStringLiteral(partlist.Name)))
			for _, elem := range partlist.PartGroup {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartGroup = append(%s.PartGroup, %s)", partlistIdent, partlistIdent, targetIdent))
			}
			for _, elem := range partlist.ScorePart {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ScorePart = append(%s.ScorePart, %s)", partlistIdent, partlistIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, partname := range __gong__sortStageSetInstances(stageSet.ScoreStage.PartNames, stageSet.ScoreStage.PartName_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			partnameIdent := "__score" + partname.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.PartName{Name: %s}).Stage(stageSet.ScoreStage)", partnameIdent, __gong__toRawStringLiteral(partname.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
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
	if stageSet.ScoreStage != nil {
		for _, player := range __gong__sortStageSetInstances(stageSet.ScoreStage.Players, stageSet.ScoreStage.Player_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			playerIdent := "__score" + player.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Player{Name: %s}).Stage(stageSet.ScoreStage)", playerIdent, __gong__toRawStringLiteral(player.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", playerIdent, __gong__toRawStringLiteral(player.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", playerIdent, __gong__toRawStringLiteral(player.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.PlayerName = %s", playerIdent, __gong__toRawStringLiteral(player.PlayerName)))
		}
	}
	if stageSet.ScoreStage != nil {
		for _, scoreinstrument := range __gong__sortStageSetInstances(stageSet.ScoreStage.ScoreInstruments, stageSet.ScoreStage.ScoreInstrument_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			scoreinstrumentIdent := "__score" + scoreinstrument.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.ScoreInstrument{Name: %s}).Stage(stageSet.ScoreStage)", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.InstrumentName = %s", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.InstrumentName)))
			values.WriteString(fmt.Sprintf("\n\t%s.InstrumentAbbreviation = %s", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.InstrumentAbbreviation)))
			values.WriteString(fmt.Sprintf("\n\t%s.InstrumentSound = %s", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.InstrumentSound)))
			values.WriteString(fmt.Sprintf("\n\t%s.Ensemble = %s", scoreinstrumentIdent, __gong__toRawStringLiteral(scoreinstrument.Ensemble)))
			if scoreinstrument.Solo != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + scoreinstrument.Solo.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Solo = %s", scoreinstrumentIdent, targetIdent))
			}
			if scoreinstrument.VirtualInstrument != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + scoreinstrument.VirtualInstrument.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VirtualInstrument = %s", scoreinstrumentIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, scorepart := range __gong__sortStageSetInstances(stageSet.ScoreStage.ScoreParts, stageSet.ScoreStage.ScorePart_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			scorepartIdent := "__score" + scorepart.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.ScorePart{Name: %s}).Stage(stageSet.ScoreStage)", scorepartIdent, __gong__toRawStringLiteral(scorepart.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scorepartIdent, __gong__toRawStringLiteral(scorepart.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", scorepartIdent, __gong__toRawStringLiteral(scorepart.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Group = %s", scorepartIdent, __gong__toRawStringLiteral(scorepart.Group)))
			if scorepart.Identification != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + scorepart.Identification.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identification = %s", scorepartIdent, targetIdent))
			}
			for _, elem := range scorepart.PartLink {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartLink = append(%s.PartLink, %s)", scorepartIdent, scorepartIdent, targetIdent))
			}
			if scorepart.PartName != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + scorepart.PartName.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartName = %s", scorepartIdent, targetIdent))
			}
			if scorepart.PartNameDisplay != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + scorepart.PartNameDisplay.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartNameDisplay = %s", scorepartIdent, targetIdent))
			}
			if scorepart.PartAbbreviation != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + scorepart.PartAbbreviation.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartAbbreviation = %s", scorepartIdent, targetIdent))
			}
			if scorepart.PartAbbreviationDisplay != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + scorepart.PartAbbreviationDisplay.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartAbbreviationDisplay = %s", scorepartIdent, targetIdent))
			}
			for _, elem := range scorepart.ScoreInstrument {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ScoreInstrument = append(%s.ScoreInstrument, %s)", scorepartIdent, scorepartIdent, targetIdent))
			}
			for _, elem := range scorepart.Player {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Player = append(%s.Player, %s)", scorepartIdent, scorepartIdent, targetIdent))
			}
			for _, elem := range scorepart.MidiDevice {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MidiDevice = append(%s.MidiDevice, %s)", scorepartIdent, scorepartIdent, targetIdent))
			}
			for _, elem := range scorepart.MidiInstrument {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MidiInstrument = append(%s.MidiInstrument, %s)", scorepartIdent, scorepartIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, scorepartwise := range __gong__sortStageSetInstances(stageSet.ScoreStage.ScorePartwises, stageSet.ScoreStage.ScorePartwise_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			scorepartwiseIdent := "__score" + scorepartwise.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.ScorePartwise{Name: %s}).Stage(stageSet.ScoreStage)", scorepartwiseIdent, __gong__toRawStringLiteral(scorepartwise.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scorepartwiseIdent, __gong__toRawStringLiteral(scorepartwise.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Version = %s", scorepartwiseIdent, __gong__toRawStringLiteral(scorepartwise.Version)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementNumber = %s", scorepartwiseIdent, __gong__toRawStringLiteral(scorepartwise.MovementNumber)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementTitle = %s", scorepartwiseIdent, __gong__toRawStringLiteral(scorepartwise.MovementTitle)))
			if scorepartwise.Work != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + scorepartwise.Work.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Work = %s", scorepartwiseIdent, targetIdent))
			}
			if scorepartwise.Identification != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + scorepartwise.Identification.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identification = %s", scorepartwiseIdent, targetIdent))
			}
			if scorepartwise.Defaults != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + scorepartwise.Defaults.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Defaults = %s", scorepartwiseIdent, targetIdent))
			}
			for _, elem := range scorepartwise.Credit {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit = append(%s.Credit, %s)", scorepartwiseIdent, scorepartwiseIdent, targetIdent))
			}
			if scorepartwise.PartList != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + scorepartwise.PartList.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartList = %s", scorepartwiseIdent, targetIdent))
			}
			for _, elem := range scorepartwise.Part {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = append(%s.Part, %s)", scorepartwiseIdent, scorepartwiseIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, scoretimewise := range __gong__sortStageSetInstances(stageSet.ScoreStage.ScoreTimewises, stageSet.ScoreStage.ScoreTimewise_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			scoretimewiseIdent := "__score" + scoretimewise.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.ScoreTimewise{Name: %s}).Stage(stageSet.ScoreStage)", scoretimewiseIdent, __gong__toRawStringLiteral(scoretimewise.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scoretimewiseIdent, __gong__toRawStringLiteral(scoretimewise.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Version = %s", scoretimewiseIdent, __gong__toRawStringLiteral(scoretimewise.Version)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementNumber = %s", scoretimewiseIdent, __gong__toRawStringLiteral(scoretimewise.MovementNumber)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementTitle = %s", scoretimewiseIdent, __gong__toRawStringLiteral(scoretimewise.MovementTitle)))
			if scoretimewise.Work != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + scoretimewise.Work.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Work = %s", scoretimewiseIdent, targetIdent))
			}
			if scoretimewise.Identification != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + scoretimewise.Identification.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Identification = %s", scoretimewiseIdent, targetIdent))
			}
			if scoretimewise.Defaults != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + scoretimewise.Defaults.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Defaults = %s", scoretimewiseIdent, targetIdent))
			}
			for _, elem := range scoretimewise.Credit {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Credit = append(%s.Credit, %s)", scoretimewiseIdent, scoretimewiseIdent, targetIdent))
			}
			if scoretimewise.PartList != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + scoretimewise.PartList.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PartList = %s", scoretimewiseIdent, targetIdent))
			}
			for _, elem := range scoretimewise.Measure {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Measure = append(%s.Measure, %s)", scoretimewiseIdent, scoretimewiseIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, scoretimewisemeasure := range __gong__sortStageSetInstances(stageSet.ScoreStage.ScoreTimewiseMeasures, stageSet.ScoreStage.ScoreTimewiseMeasure_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			scoretimewisemeasureIdent := "__score" + scoretimewisemeasure.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.ScoreTimewiseMeasure{Name: %s}).Stage(stageSet.ScoreStage)", scoretimewisemeasureIdent, __gong__toRawStringLiteral(scoretimewisemeasure.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(scoretimewisemeasure.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Number = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(scoretimewisemeasure.Number)))
			values.WriteString(fmt.Sprintf("\n\t%s.Text = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(scoretimewisemeasure.Text)))
			values.WriteString(fmt.Sprintf("\n\t%s.Implicit = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(string(scoretimewisemeasure.Implicit))))
			values.WriteString(fmt.Sprintf("\n\t%s.NonControlling = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(string(scoretimewisemeasure.NonControlling))))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(scoretimewisemeasure.Width)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", scoretimewisemeasureIdent, __gong__toRawStringLiteral(scoretimewisemeasure.Id)))
			for _, elem := range scoretimewisemeasure.Part {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + elem.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = append(%s.Part, %s)", scoretimewisemeasureIdent, scoretimewisemeasureIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, scoretimewisemeasurepart := range __gong__sortStageSetInstances(stageSet.ScoreStage.ScoreTimewiseMeasureParts, stageSet.ScoreStage.ScoreTimewiseMeasurePart_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			scoretimewisemeasurepartIdent := "__score" + scoretimewisemeasurepart.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.ScoreTimewiseMeasurePart{Name: %s}).Stage(stageSet.ScoreStage)", scoretimewisemeasurepartIdent, __gong__toRawStringLiteral(scoretimewisemeasurepart.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", scoretimewisemeasurepartIdent, __gong__toRawStringLiteral(scoretimewisemeasurepart.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", scoretimewisemeasurepartIdent, __gong__toRawStringLiteral(scoretimewisemeasurepart.Id)))
			if scoretimewisemeasurepart.RefScorePart != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + scoretimewisemeasurepart.RefScorePart.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RefScorePart = %s", scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.Note {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = append(%s.Note, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.Backup {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Backup = append(%s.Backup, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.Forward {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Forward = append(%s.Forward, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.Direction {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Direction = append(%s.Direction, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.Attributes {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__attributes" + elem.GongGetIdentifier(stageSet.AttributesStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Attributes = append(%s.Attributes, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.Harmony {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Harmony = append(%s.Harmony, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.FiguredBass {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__note" + elem.GongGetIdentifier(stageSet.NoteStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.FiguredBass = append(%s.FiguredBass, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.Print {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Print = append(%s.Print, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.Sound {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sound = append(%s.Sound, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.Listening {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Listening = append(%s.Listening, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.Barline {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__barline" + elem.GongGetIdentifier(stageSet.BarlineStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Barline = append(%s.Barline, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.Grouping {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__direction" + elem.GongGetIdentifier(stageSet.DirectionStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Grouping = append(%s.Grouping, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.Link {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Link = append(%s.Link, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
			for _, elem := range scoretimewisemeasurepart.Bookmark {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__common" + elem.GongGetIdentifier(stageSet.CommonStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Bookmark = append(%s.Bookmark, %s)", scoretimewisemeasurepartIdent, scoretimewisemeasurepartIdent, targetIdent))
			}
		}
	}
	if stageSet.ScoreStage != nil {
		for _, work := range __gong__sortStageSetInstances(stageSet.ScoreStage.Works, stageSet.ScoreStage.Work_stagedOrder) {
			if lastStageDecl != "ScoreStage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "ScoreStage"
			}
			workIdent := "__score" + work.GongGetIdentifier(stageSet.ScoreStage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&score.Work{Name: %s}).Stage(stageSet.ScoreStage)", workIdent, __gong__toRawStringLiteral(work.Name)))
			if lastStageVal != "ScoreStage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "ScoreStage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", workIdent, __gong__toRawStringLiteral(work.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.WorkNumber = %s", workIdent, __gong__toRawStringLiteral(work.WorkNumber)))
			values.WriteString(fmt.Sprintf("\n\t%s.WorkTitle = %s", workIdent, __gong__toRawStringLiteral(work.WorkTitle)))
			if work.Opus != nil {
				if lastStagePtr != "ScoreStage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "ScoreStage"
				}
				targetIdent := "__score" + work.Opus.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Opus = %s", workIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		for _, musicxmldocument := range __gong__sortStageSetInstances(stageSet.Stage.MusicXMLDocuments, stageSet.Stage.MusicXMLDocument_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			musicxmldocumentIdent := "__models" + musicxmldocument.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.MusicXMLDocument{Name: %s}).Stage(stageSet.Stage)", musicxmldocumentIdent, __gong__toRawStringLiteral(musicxmldocument.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", musicxmldocumentIdent, __gong__toRawStringLiteral(musicxmldocument.Name)))
			if musicxmldocument.ScorePartwise != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + musicxmldocument.ScorePartwise.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ScorePartwise = %s", musicxmldocumentIdent, targetIdent))
			}
			if musicxmldocument.ScoreTimewise != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := "__score" + musicxmldocument.ScoreTimewise.GongGetIdentifier(stageSet.ScoreStage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ScoreTimewise = %s", musicxmldocumentIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/score"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/barline"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common"
	"github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/xlink"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
	_ *score.Stage
	_ *note.Stage
	_ *direction.Stage
	_ *barline.Stage
	_ *attributes.Stage
	_ *common.Stage
	_ *xlink.Stage
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
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/xlink":
			aliasToCanonical[alias] = "xlink"
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/common":
			aliasToCanonical[alias] = "common"
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/attributes":
			aliasToCanonical[alias] = "attributes"
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/barline":
			aliasToCanonical[alias] = "barline"
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/direction":
			aliasToCanonical[alias] = "direction"
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/note":
			aliasToCanonical[alias] = "note"
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models/score":
			aliasToCanonical[alias] = "score"
		case "github.com/fullstack-lang/gong/app/xsd2gong/testbed/musicxml/go/models":
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
			case "xlink":
				switch typeName {
				case "ArcType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xlink.ArcType), stageSet.XlinkStage, ident.Name, instanceName, preserveOrder)
				case "Extended":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xlink.Extended), stageSet.XlinkStage, ident.Name, instanceName, preserveOrder)
				case "LocatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xlink.LocatorType), stageSet.XlinkStage, ident.Name, instanceName, preserveOrder)
				case "ResourceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xlink.ResourceType), stageSet.XlinkStage, ident.Name, instanceName, preserveOrder)
				case "Simple":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xlink.Simple), stageSet.XlinkStage, ident.Name, instanceName, preserveOrder)
				case "TitleEltType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xlink.TitleEltType), stageSet.XlinkStage, ident.Name, instanceName, preserveOrder)
				}
			case "common":
				switch typeName {
				case "AccidentalText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.AccidentalText), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Appearance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Appearance), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Bookmark":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Bookmark), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Coda":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Coda), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Distance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Distance), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Dynamics":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Dynamics), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Empty":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Empty), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "EmptyPlacement":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.EmptyPlacement), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "EmptyPlacementSmufl":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.EmptyPlacementSmufl), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "EmptyPrintObjectStyleAlign":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.EmptyPrintObjectStyleAlign), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "EmptyPrintStyle":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.EmptyPrintStyle), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "EmptyPrintStyleAlign":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.EmptyPrintStyleAlign), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "EmptyPrintStyleAlignId":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.EmptyPrintStyleAlignId), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "EmptyTrillSound":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.EmptyTrillSound), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Encoding":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Encoding), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Fermata":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Fermata), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Fingering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Fingering), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "FormattedSymbol":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.FormattedSymbol), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "FormattedSymbolId":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.FormattedSymbolId), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "FormattedText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.FormattedText), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "FormattedTextId":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.FormattedTextId), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Fret":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Fret), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Glyph":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Glyph), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "HorizontalTurn":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.HorizontalTurn), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Identification":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Identification), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Level":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Level), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "LineWidth":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.LineWidth), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Link":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Link), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "MeasureLayout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.MeasureLayout), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "MidiDevice":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.MidiDevice), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "MidiInstrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.MidiInstrument), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Miscellaneous":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Miscellaneous), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "MiscellaneousField":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.MiscellaneousField), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "NameDisplay":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.NameDisplay), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "NoteSize":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.NoteSize), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "OtherAppearance":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.OtherAppearance), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "OtherPlay":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.OtherPlay), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "OtherText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.OtherText), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "PageLayout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.PageLayout), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "PageMargins":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.PageMargins), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Play":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Play), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Scaling":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Scaling), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Segno":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Segno), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "StaffLayout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.StaffLayout), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "StringType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.StringType), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "Supports":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.Supports), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "SystemDividers":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.SystemDividers), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "SystemLayout":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.SystemLayout), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "SystemMargins":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.SystemMargins), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "TypedText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.TypedText), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				case "WavyLine":
					identifierMap[ident.Name] = __gong__stageSetInit(new(common.WavyLine), stageSet.CommonStage, ident.Name, instanceName, preserveOrder)
				}
			case "attributes":
				switch typeName {
				case "Attributes":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.Attributes), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "AttributesDirective":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.AttributesDirective), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "BeatRepeat":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.BeatRepeat), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "Cancel":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.Cancel), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "Clef":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.Clef), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "Double":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.Double), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "ForPart":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.ForPart), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "Interchangeable":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.Interchangeable), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "Key":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.Key), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "KeyAccidental":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.KeyAccidental), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "KeyOctave":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.KeyOctave), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "LineDetail":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.LineDetail), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "MeasureRepeat":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.MeasureRepeat), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "MeasureStyle":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.MeasureStyle), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "MultipleRest":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.MultipleRest), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "PartClef":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.PartClef), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "PartSymbol":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.PartSymbol), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "PartTranspose":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.PartTranspose), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "Slash":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.Slash), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "StaffDetails":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.StaffDetails), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "StaffSize":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.StaffSize), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "StaffTuning":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.StaffTuning), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "Time":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.Time), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				case "Transpose":
					identifierMap[ident.Name] = __gong__stageSetInit(new(attributes.Transpose), stageSet.AttributesStage, ident.Name, instanceName, preserveOrder)
				}
			case "barline":
				switch typeName {
				case "BarStyleColor":
					identifierMap[ident.Name] = __gong__stageSetInit(new(barline.BarStyleColor), stageSet.BarlineStage, ident.Name, instanceName, preserveOrder)
				case "Barline":
					identifierMap[ident.Name] = __gong__stageSetInit(new(barline.Barline), stageSet.BarlineStage, ident.Name, instanceName, preserveOrder)
				case "Ending":
					identifierMap[ident.Name] = __gong__stageSetInit(new(barline.Ending), stageSet.BarlineStage, ident.Name, instanceName, preserveOrder)
				case "Repeat":
					identifierMap[ident.Name] = __gong__stageSetInit(new(barline.Repeat), stageSet.BarlineStage, ident.Name, instanceName, preserveOrder)
				}
			case "direction":
				switch typeName {
				case "Accord":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Accord), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "AccordionRegistration":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.AccordionRegistration), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Barre":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Barre), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Bass":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Bass), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "BassStep":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.BassStep), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "BeatUnitTied":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.BeatUnitTied), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Beater":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Beater), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Bracket":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Bracket), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Dashes":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Dashes), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Degree":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Degree), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "DegreeAlter":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.DegreeAlter), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "DegreeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.DegreeType), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "DegreeValue":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.DegreeValue), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Direction":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Direction), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "DirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.DirectionType), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Effect":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Effect), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Feature":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Feature), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "FirstFret":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.FirstFret), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Frame":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Frame), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "FrameNote":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.FrameNote), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Glass":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Glass), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Grouping":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Grouping), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Harmony":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Harmony), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "HarmonyAlter":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.HarmonyAlter), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "HarpPedals":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.HarpPedals), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Image":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Image), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "InstrumentChange":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.InstrumentChange), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Inversion":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Inversion), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Kind":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Kind), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Listening":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Listening), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "MeasureNumbering":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.MeasureNumbering), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Membrane":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Membrane), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Metal":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Metal), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Metronome":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Metronome), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "MetronomeBeam":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.MetronomeBeam), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "MetronomeNote":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.MetronomeNote), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "MetronomeTied":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.MetronomeTied), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "MetronomeTuplet":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.MetronomeTuplet), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Numeral":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Numeral), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "NumeralKey":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.NumeralKey), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "NumeralRoot":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.NumeralRoot), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "OctaveShift":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.OctaveShift), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Offset":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Offset), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "OtherDirection":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.OtherDirection), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "OtherListening":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.OtherListening), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Pedal":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Pedal), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "PedalTuning":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.PedalTuning), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "PerMinute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.PerMinute), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Percussion":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Percussion), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Pitched":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Pitched), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "PrincipalVoice":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.PrincipalVoice), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Print":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Print), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Root":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Root), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "RootStep":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.RootStep), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Scordatura":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Scordatura), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Sound":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Sound), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "StaffDivide":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.StaffDivide), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Stick":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Stick), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "StringMute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.StringMute), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "StyleText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.StyleText), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Swing":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Swing), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Sync":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Sync), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "TimeModification":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.TimeModification), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Timpani":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Timpani), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "VirtualInstrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.VirtualInstrument), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Wedge":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Wedge), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				case "Wood":
					identifierMap[ident.Name] = __gong__stageSetInit(new(direction.Wood), stageSet.DirectionStage, ident.Name, instanceName, preserveOrder)
				}
			case "note":
				switch typeName {
				case "Accidental":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Accidental), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "AccidentalMark":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.AccidentalMark), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Arpeggiate":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Arpeggiate), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Arrow":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Arrow), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Articulations":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Articulations), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Assess":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Assess), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Backup":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Backup), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Beam":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Beam), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Bend":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Bend), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "BreathMark":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.BreathMark), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Caesura":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Caesura), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Elision":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Elision), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "EmptyLine":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.EmptyLine), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Extend":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Extend), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Figure":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Figure), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "FiguredBass":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.FiguredBass), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Forward":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Forward), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Glissando":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Glissando), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Grace":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Grace), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "HammerOnPullOff":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.HammerOnPullOff), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Handbell":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Handbell), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "HarmonClosed":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.HarmonClosed), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "HarmonMute":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.HarmonMute), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Harmonic":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Harmonic), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "HeelToe":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.HeelToe), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Hole":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Hole), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "HoleClosed":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.HoleClosed), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Instrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Instrument), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Listen":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Listen), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Lyric":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Lyric), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Mordent":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Mordent), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "NonArpeggiate":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.NonArpeggiate), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Notations":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Notations), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Note":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Note), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "NoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.NoteType), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Notehead":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Notehead), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "NoteheadText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.NoteheadText), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Ornaments":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Ornaments), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "OtherNotation":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.OtherNotation), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "OtherPlacementText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.OtherPlacementText), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Pitch":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Pitch), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "PlacementText":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.PlacementText), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Release":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Release), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Rest":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Rest), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Slide":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Slide), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Slur":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Slur), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Stem":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Stem), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "StrongAccent":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.StrongAccent), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Tap":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Tap), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Technical":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Technical), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "TextElementData":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.TextElementData), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Tie":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Tie), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Tied":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Tied), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Tremolo":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Tremolo), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Tuplet":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Tuplet), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "TupletDot":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.TupletDot), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "TupletNumber":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.TupletNumber), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "TupletPortion":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.TupletPortion), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "TupletType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.TupletType), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Unpitched":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Unpitched), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				case "Wait":
					identifierMap[ident.Name] = __gong__stageSetInit(new(note.Wait), stageSet.NoteStage, ident.Name, instanceName, preserveOrder)
				}
			case "score":
				switch typeName {
				case "Credit":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.Credit), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "Defaults":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.Defaults), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "EmptyFont":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.EmptyFont), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "GroupBarline":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.GroupBarline), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "GroupName":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.GroupName), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "GroupSymbol":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.GroupSymbol), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "InstrumentLink":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.InstrumentLink), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "LyricFont":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.LyricFont), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "LyricLanguage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.LyricLanguage), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "Measure":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.Measure), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "MeasureMusicDataChoice":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.MeasureMusicDataChoice), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "Opus":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.Opus), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "Part":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.Part), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "PartGroup":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.PartGroup), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "PartLink":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.PartLink), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "PartList":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.PartList), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "PartName":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.PartName), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "Player":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.Player), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "ScoreInstrument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.ScoreInstrument), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "ScorePart":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.ScorePart), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "ScorePartwise":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.ScorePartwise), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "ScoreTimewise":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.ScoreTimewise), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "ScoreTimewiseMeasure":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.ScoreTimewiseMeasure), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "ScoreTimewiseMeasurePart":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.ScoreTimewiseMeasurePart), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				case "Work":
					identifierMap[ident.Name] = __gong__stageSetInit(new(score.Work), stageSet.ScoreStage, ident.Name, instanceName, preserveOrder)
				}
			case "models":
				switch typeName {
				case "MusicXMLDocument":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MusicXMLDocument), stageSet.Stage, ident.Name, instanceName, preserveOrder)
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
				case *xlink.ArcType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = xlink.EnumTypeType(GongExtractString(rhs))
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = xlink.EnumShowType(GongExtractString(rhs))
					case "Actuate":
						inst.Actuate = xlink.EnumActuateType(GongExtractString(rhs))
					case "From":
						inst.From = GongExtractString(rhs)
					case "To":
						inst.To = GongExtractString(rhs)
					case "Title_1":
						__gong__assignSliceOfPointers(&inst.Title_1, rhs, identifierMap)
					}
				case *xlink.Extended:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = xlink.EnumTypeType(GongExtractString(rhs))
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Title_1":
						__gong__assignSliceOfPointers(&inst.Title_1, rhs, identifierMap)
					case "Resource":
						__gong__assignSliceOfPointers(&inst.Resource, rhs, identifierMap)
					case "Locator":
						__gong__assignSliceOfPointers(&inst.Locator, rhs, identifierMap)
					case "Arc":
						__gong__assignSliceOfPointers(&inst.Arc, rhs, identifierMap)
					}
				case *xlink.LocatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = xlink.EnumTypeType(GongExtractString(rhs))
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "Title_1":
						__gong__assignSliceOfPointers(&inst.Title_1, rhs, identifierMap)
					}
				case *xlink.ResourceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = xlink.EnumTypeType(GongExtractString(rhs))
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Label":
						inst.Label = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *xlink.Simple:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = xlink.EnumTypeType(GongExtractString(rhs))
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = xlink.EnumShowType(GongExtractString(rhs))
					case "Actuate":
						inst.Actuate = xlink.EnumActuateType(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *xlink.TitleEltType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = xlink.EnumTypeType(GongExtractString(rhs))
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.AccidentalText:
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
						inst.Justify = common.EnumLeftCenterRight(GongExtractString(rhs))
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
						inst.Dir = common.EnumTextDirection(GongExtractString(rhs))
					case "Enclosure":
						inst.Enclosure = common.EnumEnclosureShape(GongExtractString(rhs))
					case "Value":
						inst.Value = common.EnumAccidentalValue(GongExtractString(rhs))
					}
				case *common.Appearance:
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
				case *common.Bookmark:
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
				case *common.Coda:
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
				case *common.Distance:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.Dynamics:
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
					case "Placement":
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Underline":
						inst.Underline = GongExtractInt(rhs)
					case "Overline":
						inst.Overline = GongExtractInt(rhs)
					case "LineThrough":
						inst.LineThrough = GongExtractInt(rhs)
					case "Enclosure":
						inst.Enclosure = common.EnumEnclosureShape(GongExtractString(rhs))
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
				case *common.Empty:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *common.EmptyPlacement:
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
					}
				case *common.EmptyPlacementSmufl:
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
					}
				case *common.EmptyPrintObjectStyleAlign:
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
					}
				case *common.EmptyPrintStyle:
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
					}
				case *common.EmptyPrintStyleAlign:
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
					}
				case *common.EmptyPrintStyleAlignId:
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
					}
				case *common.EmptyTrillSound:
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
					}
				case *common.Encoding:
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
				case *common.Fermata:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumUprightInverted(GongExtractString(rhs))
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
						inst.Value = common.EnumFermataShape(GongExtractString(rhs))
					}
				case *common.Fingering:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Substitution":
						inst.Substitution = common.EnumYesNo(GongExtractString(rhs))
					case "Alternate":
						inst.Alternate = common.EnumYesNo(GongExtractString(rhs))
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
				case *common.FormattedSymbol:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
						inst.Dir = common.EnumTextDirection(GongExtractString(rhs))
					case "Enclosure":
						inst.Enclosure = common.EnumEnclosureShape(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.FormattedSymbolId:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
						inst.Dir = common.EnumTextDirection(GongExtractString(rhs))
					case "Enclosure":
						inst.Enclosure = common.EnumEnclosureShape(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.FormattedText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Space":
						inst.Space = GongExtractString(rhs)
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
						inst.Dir = common.EnumTextDirection(GongExtractString(rhs))
					case "Enclosure":
						inst.Enclosure = common.EnumEnclosureShape(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.FormattedTextId:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "Space":
						inst.Space = GongExtractString(rhs)
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
						inst.Dir = common.EnumTextDirection(GongExtractString(rhs))
					case "Enclosure":
						inst.Enclosure = common.EnumEnclosureShape(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.Fret:
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
				case *common.Glyph:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.HorizontalTurn:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Slash":
						inst.Slash = common.EnumYesNo(GongExtractString(rhs))
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
					}
				case *common.Identification:
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
				case *common.Level:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Reference":
						inst.Reference = common.EnumYesNo(GongExtractString(rhs))
					case "Type":
						inst.Type = common.EnumStartStopSingle(GongExtractString(rhs))
					case "Parentheses":
						inst.Parentheses = common.EnumYesNo(GongExtractString(rhs))
					case "Bracket":
						inst.Bracket = common.EnumYesNo(GongExtractString(rhs))
					case "Size":
						inst.Size = common.EnumSymbolSize(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.LineWidth:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.Link:
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
				case *common.MeasureLayout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MeasureDistance":
						inst.MeasureDistance = GongExtractString(rhs)
					}
				case *common.MidiDevice:
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
				case *common.MidiInstrument:
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
				case *common.Miscellaneous:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MiscellaneousField":
						__gong__assignSliceOfPointers(&inst.MiscellaneousField, rhs, identifierMap)
					}
				case *common.MiscellaneousField:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NameAttr":
						inst.NameAttr = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.NameDisplay:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "DisplayText":
						__gong__assignSliceOfPointers(&inst.DisplayText, rhs, identifierMap)
					case "AccidentalText":
						__gong__assignSliceOfPointers(&inst.AccidentalText, rhs, identifierMap)
					}
				case *common.NoteSize:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumNoteSizeType(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.OtherAppearance:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.OtherPlay:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.OtherText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.PageLayout:
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
				case *common.PageMargins:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumMarginType(GongExtractString(rhs))
					case "LeftMargin":
						inst.LeftMargin = GongExtractString(rhs)
					case "RightMargin":
						inst.RightMargin = GongExtractString(rhs)
					case "TopMargin":
						inst.TopMargin = GongExtractString(rhs)
					case "BottomMargin":
						inst.BottomMargin = GongExtractString(rhs)
					}
				case *common.Play:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Ipa":
						inst.Ipa = GongExtractString(rhs)
					case "Mute":
						inst.Mute = common.EnumMute(GongExtractString(rhs))
					case "SemiPitched":
						inst.SemiPitched = common.EnumSemiPitched(GongExtractString(rhs))
					case "OtherPlay":
						__gong__assignSliceOfPointers(&inst.OtherPlay, rhs, identifierMap)
					}
				case *common.Scaling:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Millimeters":
						inst.Millimeters = GongExtractString(rhs)
					case "Tenths":
						inst.Tenths = GongExtractString(rhs)
					}
				case *common.Segno:
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
				case *common.StaffLayout:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "StaffDistance":
						inst.StaffDistance = GongExtractString(rhs)
					}
				case *common.StringType:
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
						inst.Value = GongExtractInt(rhs)
					}
				case *common.Supports:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumYesNo(GongExtractString(rhs))
					case "Element":
						inst.Element = GongExtractString(rhs)
					case "Attribute":
						inst.Attribute = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.SystemDividers:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "LeftDivider":
						__gong__assignPointer(&inst.LeftDivider, rhs, identifierMap)
					case "RightDivider":
						__gong__assignPointer(&inst.RightDivider, rhs, identifierMap)
					}
				case *common.SystemLayout:
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
				case *common.SystemMargins:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "LeftMargin":
						inst.LeftMargin = GongExtractString(rhs)
					case "RightMargin":
						inst.RightMargin = GongExtractString(rhs)
					}
				case *common.TypedText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *common.WavyLine:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStopContinue(GongExtractString(rhs))
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
						inst.Placement = common.EnumAboveBelow(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
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
					}
				case *attributes.Attributes:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
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
				case *attributes.AttributesDirective:
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
				case *attributes.BeatRepeat:
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
					case "SlashDot":
						__gong__assignSliceOfPointers(&inst.SlashDot, rhs, identifierMap)
					case "ExceptVoice":
						inst.ExceptVoice = GongExtractString(rhs)
					}
				case *attributes.Cancel:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						inst.Location = attributes.EnumCancelLocation(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractInt(rhs)
					}
				case *attributes.Clef:
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
						inst.Sign = attributes.EnumClefSign(GongExtractString(rhs))
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "ClefOctaveChange":
						inst.ClefOctaveChange = GongExtractInt(rhs)
					}
				case *attributes.Double:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Above":
						inst.Above = common.EnumYesNo(GongExtractString(rhs))
					}
				case *attributes.ForPart:
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
				case *attributes.Interchangeable:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Symbol":
						inst.Symbol = attributes.EnumTimeSymbol(GongExtractString(rhs))
					case "Separator":
						inst.Separator = attributes.EnumTimeSeparator(GongExtractString(rhs))
					case "TimeRelation":
						inst.TimeRelation = attributes.EnumTimeRelation(GongExtractString(rhs))
					case "Beats":
						inst.Beats = GongExtractString(rhs)
					case "BeatType":
						inst.BeatType = GongExtractString(rhs)
					}
				case *attributes.Key:
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
						inst.KeyStep = attributes.EnumStep(GongExtractString(rhs))
					case "KeyAlter":
						inst.KeyAlter = GongExtractString(rhs)
					case "KeyAccidental":
						__gong__assignSliceOfPointers(&inst.KeyAccidental, rhs, identifierMap)
					case "KeyOctave":
						__gong__assignSliceOfPointers(&inst.KeyOctave, rhs, identifierMap)
					}
				case *attributes.KeyAccidental:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = common.EnumAccidentalValue(GongExtractString(rhs))
					}
				case *attributes.KeyOctave:
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
				case *attributes.LineDetail:
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
				case *attributes.MeasureRepeat:
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
				case *attributes.MeasureStyle:
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
				case *attributes.MultipleRest:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "UseSymbols":
						inst.UseSymbols = common.EnumYesNo(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractInt(rhs)
					}
				case *attributes.PartClef:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Sign":
						inst.Sign = attributes.EnumClefSign(GongExtractString(rhs))
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "ClefOctaveChange":
						inst.ClefOctaveChange = GongExtractInt(rhs)
					}
				case *attributes.PartSymbol:
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
						inst.Value = attributes.EnumGroupSymbolValue(GongExtractString(rhs))
					}
				case *attributes.PartTranspose:
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
				case *attributes.Slash:
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
					case "SlashDot":
						__gong__assignSliceOfPointers(&inst.SlashDot, rhs, identifierMap)
					case "ExceptVoice":
						inst.ExceptVoice = GongExtractString(rhs)
					}
				case *attributes.StaffDetails:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "ShowFrets":
						inst.ShowFrets = attributes.EnumShowFrets(GongExtractString(rhs))
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "PrintSpacing":
						inst.PrintSpacing = common.EnumYesNo(GongExtractString(rhs))
					case "StaffType":
						inst.StaffType = attributes.EnumStaffType(GongExtractString(rhs))
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
				case *attributes.StaffSize:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Scaling":
						inst.Scaling = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *attributes.StaffTuning:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Line":
						inst.Line = GongExtractInt(rhs)
					case "TuningStep":
						inst.TuningStep = attributes.EnumStep(GongExtractString(rhs))
					case "TuningAlter":
						inst.TuningAlter = GongExtractString(rhs)
					case "TuningOctave":
						inst.TuningOctave = GongExtractInt(rhs)
					}
				case *attributes.Time:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Symbol":
						inst.Symbol = attributes.EnumTimeSymbol(GongExtractString(rhs))
					case "Separator":
						inst.Separator = attributes.EnumTimeSeparator(GongExtractString(rhs))
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
				case *attributes.Transpose:
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
				case *barline.BarStyleColor:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Value":
						inst.Value = barline.EnumBarStyle(GongExtractString(rhs))
					}
				case *barline.Barline:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						inst.Location = barline.EnumRightLeftMiddle(GongExtractString(rhs))
					case "Segno":
						inst.Segno = GongExtractString(rhs)
					case "Coda":
						inst.Coda = GongExtractString(rhs)
					case "Divisions":
						inst.Divisions = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "BarStyle":
						__gong__assignPointer(&inst.BarStyle, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "WavyLine":
						__gong__assignPointer(&inst.WavyLine, rhs, identifierMap)
					case "Segno_1":
						__gong__assignPointer(&inst.Segno_1, rhs, identifierMap)
					case "Coda_1":
						__gong__assignPointer(&inst.Coda_1, rhs, identifierMap)
					case "Fermata":
						__gong__assignSliceOfPointers(&inst.Fermata, rhs, identifierMap)
					case "Ending":
						__gong__assignPointer(&inst.Ending, rhs, identifierMap)
					case "Repeat":
						__gong__assignPointer(&inst.Repeat, rhs, identifierMap)
					}
				case *barline.Ending:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "Type":
						inst.Type = barline.EnumStartStopDiscontinue(GongExtractString(rhs))
					case "EndLength":
						inst.EndLength = GongExtractString(rhs)
					case "TextX":
						inst.TextX = GongExtractString(rhs)
					case "TextY":
						inst.TextY = GongExtractString(rhs)
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
					case "System":
						inst.System = barline.EnumSystemRelation(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *barline.Repeat:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Direction":
						inst.Direction = barline.EnumBackwardForward(GongExtractString(rhs))
					case "Times":
						inst.Times = GongExtractInt(rhs)
					case "AfterJump":
						inst.AfterJump = common.EnumYesNo(GongExtractString(rhs))
					case "Winged":
						inst.Winged = barline.EnumWinged(GongExtractString(rhs))
					}
				case *direction.Accord:
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
				case *direction.AccordionRegistration:
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
					case "AccordionHigh":
						__gong__assignPointer(&inst.AccordionHigh, rhs, identifierMap)
					case "AccordionMiddle":
						inst.AccordionMiddle = GongExtractInt(rhs)
					case "AccordionLow":
						__gong__assignPointer(&inst.AccordionLow, rhs, identifierMap)
					}
				case *direction.Barre:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					}
				case *direction.Bass:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Arrangement":
						inst.Arrangement = direction.EnumHarmonyArrangement(GongExtractString(rhs))
					case "BassSeparator":
						__gong__assignPointer(&inst.BassSeparator, rhs, identifierMap)
					case "BassStep":
						__gong__assignPointer(&inst.BassStep, rhs, identifierMap)
					case "BassAlter":
						__gong__assignPointer(&inst.BassAlter, rhs, identifierMap)
					}
				case *direction.BassStep:
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
				case *direction.BeatUnitTied:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "BeatUnit":
						inst.BeatUnit = common.EnumNoteTypeValue(GongExtractString(rhs))
					case "BeatUnitDot":
						__gong__assignSliceOfPointers(&inst.BeatUnitDot, rhs, identifierMap)
					}
				case *direction.Beater:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tip":
						inst.Tip = direction.EnumTipDirection(GongExtractString(rhs))
					case "Value":
						inst.Value = direction.EnumBeaterValue(GongExtractString(rhs))
					}
				case *direction.Bracket:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStopContinue(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "LineEnd":
						inst.LineEnd = direction.EnumLineEnd(GongExtractString(rhs))
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
				case *direction.Dashes:
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
				case *direction.Degree:
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
				case *direction.DegreeAlter:
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
				case *direction.DegreeType:
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
						inst.Value = direction.EnumDegreeTypeValue(GongExtractString(rhs))
					}
				case *direction.DegreeValue:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Symbol":
						inst.Symbol = direction.EnumDegreeSymbolValue(GongExtractString(rhs))
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
				case *direction.Direction:
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
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					case "Sound":
						__gong__assignPointer(&inst.Sound, rhs, identifierMap)
					case "Listening":
						__gong__assignPointer(&inst.Listening, rhs, identifierMap)
					}
				case *direction.DirectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Rehearsal":
						__gong__assignSliceOfPointers(&inst.Rehearsal, rhs, identifierMap)
					case "Segno":
						__gong__assignSliceOfPointers(&inst.Segno, rhs, identifierMap)
					case "Coda":
						__gong__assignSliceOfPointers(&inst.Coda, rhs, identifierMap)
					case "Words":
						__gong__assignSliceOfPointers(&inst.Words, rhs, identifierMap)
					case "Symbol":
						__gong__assignSliceOfPointers(&inst.Symbol, rhs, identifierMap)
					case "Wedge":
						__gong__assignPointer(&inst.Wedge, rhs, identifierMap)
					case "Dynamics":
						__gong__assignSliceOfPointers(&inst.Dynamics, rhs, identifierMap)
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
					case "Damp":
						__gong__assignPointer(&inst.Damp, rhs, identifierMap)
					case "DampAll":
						__gong__assignPointer(&inst.DampAll, rhs, identifierMap)
					case "Eyeglasses":
						__gong__assignPointer(&inst.Eyeglasses, rhs, identifierMap)
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
				case *direction.Effect:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = direction.EnumEffectValue(GongExtractString(rhs))
					}
				case *direction.Feature:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = GongExtractString(rhs)
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *direction.FirstFret:
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
				case *direction.Frame:
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
				case *direction.FrameNote:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "String":
						__gong__assignPointer(&inst.String, rhs, identifierMap)
					case "Fret":
						__gong__assignPointer(&inst.Fret, rhs, identifierMap)
					case "Fingering":
						__gong__assignPointer(&inst.Fingering, rhs, identifierMap)
					case "Barre":
						__gong__assignPointer(&inst.Barre, rhs, identifierMap)
					}
				case *direction.Glass:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = direction.EnumGlassValue(GongExtractString(rhs))
					}
				case *direction.Grouping:
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
				case *direction.Harmony:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = direction.EnumHarmonyType(GongExtractString(rhs))
					case "PrintFrame":
						inst.PrintFrame = common.EnumYesNo(GongExtractString(rhs))
					case "Arrangement":
						inst.Arrangement = direction.EnumHarmonyArrangement(GongExtractString(rhs))
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
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					}
				case *direction.HarmonyAlter:
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
				case *direction.HarpPedals:
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
				case *direction.Image:
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
				case *direction.InstrumentChange:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "InstrumentSound":
						inst.InstrumentSound = GongExtractString(rhs)
					case "Solo":
						__gong__assignPointer(&inst.Solo, rhs, identifierMap)
					case "Ensemble":
						inst.Ensemble = GongExtractString(rhs)
					case "VirtualInstrument":
						__gong__assignPointer(&inst.VirtualInstrument, rhs, identifierMap)
					}
				case *direction.Inversion:
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
				case *direction.Kind:
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
						inst.Value = direction.EnumKindValue(GongExtractString(rhs))
					}
				case *direction.Listening:
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
				case *direction.MeasureNumbering:
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
						inst.Value = direction.EnumMeasureNumberingValue(GongExtractString(rhs))
					}
				case *direction.Membrane:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = direction.EnumMembraneValue(GongExtractString(rhs))
					}
				case *direction.Metal:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = direction.EnumMetalValue(GongExtractString(rhs))
					}
				case *direction.Metronome:
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
					case "BeatUnitDot":
						__gong__assignSliceOfPointers(&inst.BeatUnitDot, rhs, identifierMap)
					case "BeatUnitTied":
						__gong__assignSliceOfPointers(&inst.BeatUnitTied, rhs, identifierMap)
					case "PerMinute":
						__gong__assignPointer(&inst.PerMinute, rhs, identifierMap)
					case "MetronomeArrows":
						__gong__assignPointer(&inst.MetronomeArrows, rhs, identifierMap)
					case "MetronomeNote":
						__gong__assignSliceOfPointers(&inst.MetronomeNote, rhs, identifierMap)
					case "MetronomeRelation":
						inst.MetronomeRelation = GongExtractString(rhs)
					}
				case *direction.MetronomeBeam:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Value":
						inst.Value = direction.EnumBeamValue(GongExtractString(rhs))
					}
				case *direction.MetronomeNote:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "MetronomeType":
						inst.MetronomeType = common.EnumNoteTypeValue(GongExtractString(rhs))
					case "MetronomeDot":
						__gong__assignSliceOfPointers(&inst.MetronomeDot, rhs, identifierMap)
					case "MetronomeBeam":
						__gong__assignSliceOfPointers(&inst.MetronomeBeam, rhs, identifierMap)
					case "MetronomeTied":
						__gong__assignPointer(&inst.MetronomeTied, rhs, identifierMap)
					case "MetronomeTuplet":
						__gong__assignPointer(&inst.MetronomeTuplet, rhs, identifierMap)
					}
				case *direction.MetronomeTied:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					}
				case *direction.MetronomeTuplet:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Bracket":
						inst.Bracket = common.EnumYesNo(GongExtractString(rhs))
					case "ShowNumber":
						inst.ShowNumber = direction.EnumShowTuplet(GongExtractString(rhs))
					case "ActualNotes":
						inst.ActualNotes = GongExtractInt(rhs)
					case "NormalNotes":
						inst.NormalNotes = GongExtractInt(rhs)
					case "NormalType":
						inst.NormalType = common.EnumNoteTypeValue(GongExtractString(rhs))
					case "NormalDot":
						__gong__assignSliceOfPointers(&inst.NormalDot, rhs, identifierMap)
					}
				case *direction.Numeral:
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
				case *direction.NumeralKey:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "NumeralFifths":
						inst.NumeralFifths = GongExtractInt(rhs)
					case "NumeralMode":
						inst.NumeralMode = direction.EnumNumeralMode(GongExtractString(rhs))
					}
				case *direction.NumeralRoot:
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
				case *direction.OctaveShift:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = direction.EnumUpDownStopContinue(GongExtractString(rhs))
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
				case *direction.Offset:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Sound":
						inst.Sound = common.EnumYesNo(GongExtractString(rhs))
					case "Value":
						inst.Value = GongExtractString(rhs)
					}
				case *direction.OtherDirection:
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
				case *direction.OtherListening:
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
				case *direction.Pedal:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = direction.EnumPedalType(GongExtractString(rhs))
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
				case *direction.PedalTuning:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PedalStep":
						inst.PedalStep = attributes.EnumStep(GongExtractString(rhs))
					case "PedalAlter":
						inst.PedalAlter = GongExtractString(rhs)
					}
				case *direction.PerMinute:
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
				case *direction.Percussion:
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
						inst.StickLocation = direction.EnumStickLocation(GongExtractString(rhs))
					case "OtherPercussion":
						__gong__assignPointer(&inst.OtherPercussion, rhs, identifierMap)
					}
				case *direction.Pitched:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = direction.EnumPitchedValue(GongExtractString(rhs))
					}
				case *direction.PrincipalVoice:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Symbol":
						inst.Symbol = direction.EnumPrincipalVoiceSymbol(GongExtractString(rhs))
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
				case *direction.Print:
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
					case "PageLayout":
						__gong__assignPointer(&inst.PageLayout, rhs, identifierMap)
					case "SystemLayout":
						__gong__assignPointer(&inst.SystemLayout, rhs, identifierMap)
					case "StaffLayout":
						__gong__assignSliceOfPointers(&inst.StaffLayout, rhs, identifierMap)
					case "MeasureLayout":
						__gong__assignPointer(&inst.MeasureLayout, rhs, identifierMap)
					case "MeasureNumbering":
						__gong__assignPointer(&inst.MeasureNumbering, rhs, identifierMap)
					case "PartNameDisplay":
						__gong__assignPointer(&inst.PartNameDisplay, rhs, identifierMap)
					case "PartAbbreviationDisplay":
						__gong__assignPointer(&inst.PartAbbreviationDisplay, rhs, identifierMap)
					}
				case *direction.Root:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RootStep":
						__gong__assignPointer(&inst.RootStep, rhs, identifierMap)
					case "RootAlter":
						__gong__assignPointer(&inst.RootAlter, rhs, identifierMap)
					}
				case *direction.RootStep:
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
				case *direction.Scordatura:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Accord":
						__gong__assignSliceOfPointers(&inst.Accord, rhs, identifierMap)
					}
				case *direction.Sound:
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
					case "MidiDevice":
						__gong__assignSliceOfPointers(&inst.MidiDevice, rhs, identifierMap)
					case "MidiInstrument":
						__gong__assignSliceOfPointers(&inst.MidiInstrument, rhs, identifierMap)
					case "Play":
						__gong__assignSliceOfPointers(&inst.Play, rhs, identifierMap)
					case "Swing":
						__gong__assignPointer(&inst.Swing, rhs, identifierMap)
					case "Offset":
						__gong__assignPointer(&inst.Offset, rhs, identifierMap)
					}
				case *direction.StaffDivide:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = direction.EnumStaffDivideSymbol(GongExtractString(rhs))
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
				case *direction.Stick:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Tip":
						inst.Tip = direction.EnumTipDirection(GongExtractString(rhs))
					case "Parentheses":
						inst.Parentheses = common.EnumYesNo(GongExtractString(rhs))
					case "DashedCircle":
						inst.DashedCircle = common.EnumYesNo(GongExtractString(rhs))
					case "StickType":
						inst.StickType = direction.EnumStickType(GongExtractString(rhs))
					case "StickMaterial":
						inst.StickMaterial = direction.EnumStickMaterial(GongExtractString(rhs))
					}
				case *direction.StringMute:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = direction.EnumOnOff(GongExtractString(rhs))
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
				case *direction.StyleText:
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
				case *direction.Swing:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Straight":
						__gong__assignPointer(&inst.Straight, rhs, identifierMap)
					case "First":
						inst.First = GongExtractInt(rhs)
					case "Second":
						inst.Second = GongExtractInt(rhs)
					case "SwingType":
						inst.SwingType = common.EnumSwingTypeValue(GongExtractString(rhs))
					case "SwingStyle":
						inst.SwingStyle = GongExtractString(rhs)
					}
				case *direction.Sync:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = direction.EnumSyncType(GongExtractString(rhs))
					case "Latency":
						inst.Latency = GongExtractInt(rhs)
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "TimeOnly":
						inst.TimeOnly = GongExtractString(rhs)
					}
				case *direction.TimeModification:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ActualNotes":
						inst.ActualNotes = GongExtractInt(rhs)
					case "NormalNotes":
						inst.NormalNotes = GongExtractInt(rhs)
					case "NormalType":
						inst.NormalType = common.EnumNoteTypeValue(GongExtractString(rhs))
					case "NormalDot":
						__gong__assignSliceOfPointers(&inst.NormalDot, rhs, identifierMap)
					}
				case *direction.Timpani:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					}
				case *direction.VirtualInstrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "VirtualLibrary":
						inst.VirtualLibrary = GongExtractString(rhs)
					case "VirtualName":
						inst.VirtualName = GongExtractString(rhs)
					}
				case *direction.Wedge:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = direction.EnumWedgeType(GongExtractString(rhs))
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
				case *direction.Wood:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Smufl":
						inst.Smufl = GongExtractString(rhs)
					case "Value":
						inst.Value = direction.EnumWoodValue(GongExtractString(rhs))
					}
				case *note.Accidental:
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
				case *note.AccidentalMark:
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
				case *note.Arpeggiate:
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
				case *note.Arrow:
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
						inst.ArrowDirection = note.EnumArrowDirection(GongExtractString(rhs))
					case "ArrowStyle":
						inst.ArrowStyle = note.EnumArrowStyle(GongExtractString(rhs))
					case "Arrowhead":
						__gong__assignPointer(&inst.Arrowhead, rhs, identifierMap)
					case "CircularArrow":
						inst.CircularArrow = note.EnumCircularArrow(GongExtractString(rhs))
					}
				case *note.Articulations:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Accent":
						__gong__assignSliceOfPointers(&inst.Accent, rhs, identifierMap)
					case "StrongAccent":
						__gong__assignSliceOfPointers(&inst.StrongAccent, rhs, identifierMap)
					case "Staccato":
						__gong__assignSliceOfPointers(&inst.Staccato, rhs, identifierMap)
					case "Tenuto":
						__gong__assignSliceOfPointers(&inst.Tenuto, rhs, identifierMap)
					case "DetachedLegato":
						__gong__assignSliceOfPointers(&inst.DetachedLegato, rhs, identifierMap)
					case "Staccatissimo":
						__gong__assignSliceOfPointers(&inst.Staccatissimo, rhs, identifierMap)
					case "Spiccato":
						__gong__assignSliceOfPointers(&inst.Spiccato, rhs, identifierMap)
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
					case "Stress":
						__gong__assignSliceOfPointers(&inst.Stress, rhs, identifierMap)
					case "Unstress":
						__gong__assignSliceOfPointers(&inst.Unstress, rhs, identifierMap)
					case "SoftAccent":
						__gong__assignSliceOfPointers(&inst.SoftAccent, rhs, identifierMap)
					case "OtherArticulation":
						__gong__assignSliceOfPointers(&inst.OtherArticulation, rhs, identifierMap)
					}
				case *note.Assess:
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
				case *note.Backup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *note.Beam:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Number":
						inst.Number = GongExtractInt(rhs)
					case "Repeater":
						inst.Repeater = common.EnumYesNo(GongExtractString(rhs))
					case "Fan":
						inst.Fan = note.EnumFan(GongExtractString(rhs))
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Value":
						inst.Value = direction.EnumBeamValue(GongExtractString(rhs))
					}
				case *note.Bend:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Shape":
						inst.Shape = note.EnumBendShape(GongExtractString(rhs))
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
					case "PreBend":
						__gong__assignPointer(&inst.PreBend, rhs, identifierMap)
					case "Release":
						__gong__assignPointer(&inst.Release, rhs, identifierMap)
					case "WithBar":
						__gong__assignPointer(&inst.WithBar, rhs, identifierMap)
					}
				case *note.BreathMark:
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
						inst.Value = note.EnumBreathMarkValue(GongExtractString(rhs))
					}
				case *note.Caesura:
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
						inst.Value = note.EnumCaesuraValue(GongExtractString(rhs))
					}
				case *note.Elision:
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
				case *note.EmptyLine:
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
				case *note.Extend:
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
				case *note.Figure:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Prefix":
						__gong__assignPointer(&inst.Prefix, rhs, identifierMap)
					case "FigureNumber":
						__gong__assignPointer(&inst.FigureNumber, rhs, identifierMap)
					case "Suffix":
						__gong__assignPointer(&inst.Suffix, rhs, identifierMap)
					case "Extend":
						__gong__assignPointer(&inst.Extend, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *note.FiguredBass:
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
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *note.Forward:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Staff":
						inst.Staff = GongExtractInt(rhs)
					}
				case *note.Glissando:
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
				case *note.Grace:
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
				case *note.HammerOnPullOff:
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
				case *note.Handbell:
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
						inst.Value = note.EnumHandbellValue(GongExtractString(rhs))
					}
				case *note.HarmonClosed:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						inst.Location = note.EnumHarmonClosedLocation(GongExtractString(rhs))
					case "Value":
						inst.Value = note.EnumHarmonClosedValue(GongExtractString(rhs))
					}
				case *note.HarmonMute:
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
				case *note.Harmonic:
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
					case "Natural":
						__gong__assignPointer(&inst.Natural, rhs, identifierMap)
					case "Artificial":
						__gong__assignPointer(&inst.Artificial, rhs, identifierMap)
					case "BasePitch":
						__gong__assignPointer(&inst.BasePitch, rhs, identifierMap)
					case "TouchingPitch":
						__gong__assignPointer(&inst.TouchingPitch, rhs, identifierMap)
					case "SoundingPitch":
						__gong__assignPointer(&inst.SoundingPitch, rhs, identifierMap)
					}
				case *note.HeelToe:
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
				case *note.Hole:
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
				case *note.HoleClosed:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Location":
						inst.Location = note.EnumHoleClosedLocation(GongExtractString(rhs))
					case "Value":
						inst.Value = note.EnumHoleClosedValue(GongExtractString(rhs))
					}
				case *note.Instrument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *note.Listen:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Assess":
						__gong__assignSliceOfPointers(&inst.Assess, rhs, identifierMap)
					case "Wait":
						__gong__assignSliceOfPointers(&inst.Wait, rhs, identifierMap)
					case "OtherListen":
						__gong__assignSliceOfPointers(&inst.OtherListen, rhs, identifierMap)
					}
				case *note.Lyric:
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
						inst.Syllabic = note.EnumSyllabic(GongExtractString(rhs))
					case "Text":
						__gong__assignPointer(&inst.Text, rhs, identifierMap)
					case "Elision":
						__gong__assignSliceOfPointers(&inst.Elision, rhs, identifierMap)
					case "Extend":
						__gong__assignPointer(&inst.Extend, rhs, identifierMap)
					case "Laughing":
						__gong__assignPointer(&inst.Laughing, rhs, identifierMap)
					case "Humming":
						__gong__assignPointer(&inst.Humming, rhs, identifierMap)
					case "EndLine":
						__gong__assignPointer(&inst.EndLine, rhs, identifierMap)
					case "EndParagraph":
						__gong__assignPointer(&inst.EndParagraph, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *note.Mordent:
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
				case *note.NonArpeggiate:
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
				case *note.Notations:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PrintObject":
						inst.PrintObject = common.EnumYesNo(GongExtractString(rhs))
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
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
					case "Dynamics":
						__gong__assignSliceOfPointers(&inst.Dynamics, rhs, identifierMap)
					case "Fermata":
						__gong__assignSliceOfPointers(&inst.Fermata, rhs, identifierMap)
					case "Arpeggiate":
						__gong__assignSliceOfPointers(&inst.Arpeggiate, rhs, identifierMap)
					case "NonArpeggiate":
						__gong__assignSliceOfPointers(&inst.NonArpeggiate, rhs, identifierMap)
					case "AccidentalMark":
						__gong__assignSliceOfPointers(&inst.AccidentalMark, rhs, identifierMap)
					case "OtherNotation":
						__gong__assignSliceOfPointers(&inst.OtherNotation, rhs, identifierMap)
					}
				case *note.Note:
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
					case "Chord":
						__gong__assignPointer(&inst.Chord, rhs, identifierMap)
					case "Pitch":
						__gong__assignPointer(&inst.Pitch, rhs, identifierMap)
					case "Unpitched":
						__gong__assignPointer(&inst.Unpitched, rhs, identifierMap)
					case "Rest":
						__gong__assignPointer(&inst.Rest, rhs, identifierMap)
					case "Tie":
						__gong__assignSliceOfPointers(&inst.Tie, rhs, identifierMap)
					case "Cue":
						__gong__assignPointer(&inst.Cue, rhs, identifierMap)
					case "Duration":
						inst.Duration = GongExtractString(rhs)
					case "Instrument":
						__gong__assignSliceOfPointers(&inst.Instrument, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "Voice":
						inst.Voice = GongExtractString(rhs)
					case "Type":
						__gong__assignPointer(&inst.Type, rhs, identifierMap)
					case "Dot":
						__gong__assignSliceOfPointers(&inst.Dot, rhs, identifierMap)
					case "Accidental":
						__gong__assignPointer(&inst.Accidental, rhs, identifierMap)
					case "TimeModification":
						__gong__assignPointer(&inst.TimeModification, rhs, identifierMap)
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
					case "Play":
						__gong__assignPointer(&inst.Play, rhs, identifierMap)
					case "Listen":
						__gong__assignPointer(&inst.Listen, rhs, identifierMap)
					}
				case *note.NoteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Size":
						inst.Size = common.EnumSymbolSize(GongExtractString(rhs))
					case "Value":
						inst.Value = common.EnumNoteTypeValue(GongExtractString(rhs))
					}
				case *note.Notehead:
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
						inst.Value = note.EnumNoteheadValue(GongExtractString(rhs))
					}
				case *note.NoteheadText:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DisplayText":
						__gong__assignSliceOfPointers(&inst.DisplayText, rhs, identifierMap)
					case "AccidentalText":
						__gong__assignSliceOfPointers(&inst.AccidentalText, rhs, identifierMap)
					}
				case *note.Ornaments:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "TrillMark":
						__gong__assignSliceOfPointers(&inst.TrillMark, rhs, identifierMap)
					case "Turn":
						__gong__assignSliceOfPointers(&inst.Turn, rhs, identifierMap)
					case "DelayedTurn":
						__gong__assignSliceOfPointers(&inst.DelayedTurn, rhs, identifierMap)
					case "InvertedTurn":
						__gong__assignSliceOfPointers(&inst.InvertedTurn, rhs, identifierMap)
					case "DelayedInvertedTurn":
						__gong__assignSliceOfPointers(&inst.DelayedInvertedTurn, rhs, identifierMap)
					case "VerticalTurn":
						__gong__assignSliceOfPointers(&inst.VerticalTurn, rhs, identifierMap)
					case "InvertedVerticalTurn":
						__gong__assignSliceOfPointers(&inst.InvertedVerticalTurn, rhs, identifierMap)
					case "Shake":
						__gong__assignSliceOfPointers(&inst.Shake, rhs, identifierMap)
					case "WavyLine":
						__gong__assignSliceOfPointers(&inst.WavyLine, rhs, identifierMap)
					case "Mordent":
						__gong__assignSliceOfPointers(&inst.Mordent, rhs, identifierMap)
					case "InvertedMordent":
						__gong__assignSliceOfPointers(&inst.InvertedMordent, rhs, identifierMap)
					case "Schleifer":
						__gong__assignSliceOfPointers(&inst.Schleifer, rhs, identifierMap)
					case "Tremolo":
						__gong__assignSliceOfPointers(&inst.Tremolo, rhs, identifierMap)
					case "Haydn":
						__gong__assignSliceOfPointers(&inst.Haydn, rhs, identifierMap)
					case "OtherOrnament":
						__gong__assignSliceOfPointers(&inst.OtherOrnament, rhs, identifierMap)
					case "AccidentalMark":
						__gong__assignSliceOfPointers(&inst.AccidentalMark, rhs, identifierMap)
					}
				case *note.OtherNotation:
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
				case *note.OtherPlacementText:
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
				case *note.Pitch:
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
				case *note.PlacementText:
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
				case *note.Release:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Offset":
						inst.Offset = GongExtractString(rhs)
					}
				case *note.Rest:
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
				case *note.Slide:
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
				case *note.Slur:
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
				case *note.Stem:
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
						inst.Value = note.EnumStemValue(GongExtractString(rhs))
					}
				case *note.StrongAccent:
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
				case *note.Tap:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Hand":
						inst.Hand = note.EnumTapHand(GongExtractString(rhs))
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
				case *note.Technical:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "UpBow":
						__gong__assignSliceOfPointers(&inst.UpBow, rhs, identifierMap)
					case "DownBow":
						__gong__assignSliceOfPointers(&inst.DownBow, rhs, identifierMap)
					case "Harmonic":
						__gong__assignSliceOfPointers(&inst.Harmonic, rhs, identifierMap)
					case "OpenString":
						__gong__assignSliceOfPointers(&inst.OpenString, rhs, identifierMap)
					case "ThumbPosition":
						__gong__assignSliceOfPointers(&inst.ThumbPosition, rhs, identifierMap)
					case "Fingering":
						__gong__assignSliceOfPointers(&inst.Fingering, rhs, identifierMap)
					case "Pluck":
						__gong__assignSliceOfPointers(&inst.Pluck, rhs, identifierMap)
					case "DoubleTongue":
						__gong__assignSliceOfPointers(&inst.DoubleTongue, rhs, identifierMap)
					case "TripleTongue":
						__gong__assignSliceOfPointers(&inst.TripleTongue, rhs, identifierMap)
					case "Stopped":
						__gong__assignSliceOfPointers(&inst.Stopped, rhs, identifierMap)
					case "SnapPizzicato":
						__gong__assignSliceOfPointers(&inst.SnapPizzicato, rhs, identifierMap)
					case "Fret":
						__gong__assignSliceOfPointers(&inst.Fret, rhs, identifierMap)
					case "String":
						__gong__assignSliceOfPointers(&inst.String, rhs, identifierMap)
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
					case "Fingernails":
						__gong__assignSliceOfPointers(&inst.Fingernails, rhs, identifierMap)
					case "Hole":
						__gong__assignSliceOfPointers(&inst.Hole, rhs, identifierMap)
					case "Arrow":
						__gong__assignSliceOfPointers(&inst.Arrow, rhs, identifierMap)
					case "Handbell":
						__gong__assignSliceOfPointers(&inst.Handbell, rhs, identifierMap)
					case "BrassBend":
						__gong__assignSliceOfPointers(&inst.BrassBend, rhs, identifierMap)
					case "Flip":
						__gong__assignSliceOfPointers(&inst.Flip, rhs, identifierMap)
					case "Smear":
						__gong__assignSliceOfPointers(&inst.Smear, rhs, identifierMap)
					case "Open":
						__gong__assignSliceOfPointers(&inst.Open, rhs, identifierMap)
					case "HalfMuted":
						__gong__assignSliceOfPointers(&inst.HalfMuted, rhs, identifierMap)
					case "HarmonMute":
						__gong__assignSliceOfPointers(&inst.HarmonMute, rhs, identifierMap)
					case "Golpe":
						__gong__assignSliceOfPointers(&inst.Golpe, rhs, identifierMap)
					case "OtherTechnical":
						__gong__assignSliceOfPointers(&inst.OtherTechnical, rhs, identifierMap)
					}
				case *note.TextElementData:
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
				case *note.Tie:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "TimeOnly":
						inst.TimeOnly = GongExtractString(rhs)
					}
				case *note.Tied:
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
				case *note.Tremolo:
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
				case *note.Tuplet:
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
				case *note.TupletDot:
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
				case *note.TupletNumber:
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
				case *note.TupletPortion:
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
				case *note.TupletType:
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
				case *note.Unpitched:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "DisplayStep":
						inst.DisplayStep = attributes.EnumStep(GongExtractString(rhs))
					case "DisplayOctave":
						inst.DisplayOctave = GongExtractInt(rhs)
					}
				case *note.Wait:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Player":
						inst.Player = GongExtractString(rhs)
					case "TimeOnly":
						inst.TimeOnly = GongExtractString(rhs)
					}
				case *score.Credit:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Page":
						inst.Page = GongExtractInt(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "CreditType":
						inst.CreditType = GongExtractString(rhs)
					case "Link":
						__gong__assignSliceOfPointers(&inst.Link, rhs, identifierMap)
					case "Bookmark":
						__gong__assignSliceOfPointers(&inst.Bookmark, rhs, identifierMap)
					case "CreditImage":
						__gong__assignPointer(&inst.CreditImage, rhs, identifierMap)
					case "CreditWords":
						__gong__assignSliceOfPointers(&inst.CreditWords, rhs, identifierMap)
					case "CreditSymbol":
						__gong__assignSliceOfPointers(&inst.CreditSymbol, rhs, identifierMap)
					}
				case *score.Defaults:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Scaling":
						__gong__assignPointer(&inst.Scaling, rhs, identifierMap)
					case "ConcertScore":
						__gong__assignPointer(&inst.ConcertScore, rhs, identifierMap)
					case "PageLayout":
						__gong__assignPointer(&inst.PageLayout, rhs, identifierMap)
					case "SystemLayout":
						__gong__assignPointer(&inst.SystemLayout, rhs, identifierMap)
					case "StaffLayout":
						__gong__assignSliceOfPointers(&inst.StaffLayout, rhs, identifierMap)
					case "Appearance":
						__gong__assignPointer(&inst.Appearance, rhs, identifierMap)
					case "MusicFont":
						__gong__assignPointer(&inst.MusicFont, rhs, identifierMap)
					case "WordFont":
						__gong__assignPointer(&inst.WordFont, rhs, identifierMap)
					case "LyricFont":
						__gong__assignSliceOfPointers(&inst.LyricFont, rhs, identifierMap)
					case "LyricLanguage":
						__gong__assignSliceOfPointers(&inst.LyricLanguage, rhs, identifierMap)
					}
				case *score.EmptyFont:
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
				case *score.GroupBarline:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Color":
						inst.Color = GongExtractString(rhs)
					case "Value":
						inst.Value = score.EnumGroupBarlineValue(GongExtractString(rhs))
					}
				case *score.GroupName:
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
				case *score.GroupSymbol:
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
				case *score.InstrumentLink:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RefScoreInstrument":
						__gong__assignPointer(&inst.RefScoreInstrument, rhs, identifierMap)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *score.LyricFont:
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
				case *score.LyricLanguage:
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
				case *score.Measure:
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
				case *score.MeasureMusicDataChoice:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Backup":
						__gong__assignPointer(&inst.Backup, rhs, identifierMap)
					case "Forward":
						__gong__assignPointer(&inst.Forward, rhs, identifierMap)
					case "Direction":
						__gong__assignPointer(&inst.Direction, rhs, identifierMap)
					case "Attributes":
						__gong__assignPointer(&inst.Attributes, rhs, identifierMap)
					case "Harmony":
						__gong__assignPointer(&inst.Harmony, rhs, identifierMap)
					case "FiguredBass":
						__gong__assignPointer(&inst.FiguredBass, rhs, identifierMap)
					case "Print":
						__gong__assignPointer(&inst.Print, rhs, identifierMap)
					case "Sound":
						__gong__assignPointer(&inst.Sound, rhs, identifierMap)
					case "Listening":
						__gong__assignPointer(&inst.Listening, rhs, identifierMap)
					case "Barline":
						__gong__assignPointer(&inst.Barline, rhs, identifierMap)
					case "Grouping":
						__gong__assignPointer(&inst.Grouping, rhs, identifierMap)
					case "Link":
						__gong__assignPointer(&inst.Link, rhs, identifierMap)
					case "Bookmark":
						__gong__assignPointer(&inst.Bookmark, rhs, identifierMap)
					}
				case *score.Opus:
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
				case *score.Part:
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
				case *score.PartGroup:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Type":
						inst.Type = common.EnumStartStop(GongExtractString(rhs))
					case "Number":
						inst.Number = GongExtractString(rhs)
					case "GroupName":
						__gong__assignPointer(&inst.GroupName, rhs, identifierMap)
					case "GroupNameDisplay":
						__gong__assignPointer(&inst.GroupNameDisplay, rhs, identifierMap)
					case "GroupAbbreviation":
						__gong__assignPointer(&inst.GroupAbbreviation, rhs, identifierMap)
					case "GroupAbbreviationDisplay":
						__gong__assignPointer(&inst.GroupAbbreviationDisplay, rhs, identifierMap)
					case "GroupSymbol":
						__gong__assignPointer(&inst.GroupSymbol, rhs, identifierMap)
					case "GroupBarline":
						__gong__assignPointer(&inst.GroupBarline, rhs, identifierMap)
					case "GroupTime":
						__gong__assignPointer(&inst.GroupTime, rhs, identifierMap)
					case "Footnote":
						__gong__assignPointer(&inst.Footnote, rhs, identifierMap)
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					}
				case *score.PartLink:
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
				case *score.PartList:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PartGroup":
						__gong__assignSliceOfPointers(&inst.PartGroup, rhs, identifierMap)
					case "ScorePart":
						__gong__assignSliceOfPointers(&inst.ScorePart, rhs, identifierMap)
					}
				case *score.PartName:
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
				case *score.Player:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "PlayerName":
						inst.PlayerName = GongExtractString(rhs)
					}
				case *score.ScoreInstrument:
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
					case "Solo":
						__gong__assignPointer(&inst.Solo, rhs, identifierMap)
					case "Ensemble":
						inst.Ensemble = GongExtractString(rhs)
					case "VirtualInstrument":
						__gong__assignPointer(&inst.VirtualInstrument, rhs, identifierMap)
					}
				case *score.ScorePart:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Identification":
						__gong__assignPointer(&inst.Identification, rhs, identifierMap)
					case "PartLink":
						__gong__assignSliceOfPointers(&inst.PartLink, rhs, identifierMap)
					case "PartName":
						__gong__assignPointer(&inst.PartName, rhs, identifierMap)
					case "PartNameDisplay":
						__gong__assignPointer(&inst.PartNameDisplay, rhs, identifierMap)
					case "PartAbbreviation":
						__gong__assignPointer(&inst.PartAbbreviation, rhs, identifierMap)
					case "PartAbbreviationDisplay":
						__gong__assignPointer(&inst.PartAbbreviationDisplay, rhs, identifierMap)
					case "Group":
						inst.Group = GongExtractString(rhs)
					case "ScoreInstrument":
						__gong__assignSliceOfPointers(&inst.ScoreInstrument, rhs, identifierMap)
					case "Player":
						__gong__assignSliceOfPointers(&inst.Player, rhs, identifierMap)
					case "MidiDevice":
						__gong__assignSliceOfPointers(&inst.MidiDevice, rhs, identifierMap)
					case "MidiInstrument":
						__gong__assignSliceOfPointers(&inst.MidiInstrument, rhs, identifierMap)
					}
				case *score.ScorePartwise:
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
					case "Identification":
						__gong__assignPointer(&inst.Identification, rhs, identifierMap)
					case "Defaults":
						__gong__assignPointer(&inst.Defaults, rhs, identifierMap)
					case "Credit":
						__gong__assignSliceOfPointers(&inst.Credit, rhs, identifierMap)
					case "PartList":
						__gong__assignPointer(&inst.PartList, rhs, identifierMap)
					case "Part":
						__gong__assignSliceOfPointers(&inst.Part, rhs, identifierMap)
					}
				case *score.ScoreTimewise:
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
					case "Identification":
						__gong__assignPointer(&inst.Identification, rhs, identifierMap)
					case "Defaults":
						__gong__assignPointer(&inst.Defaults, rhs, identifierMap)
					case "Credit":
						__gong__assignSliceOfPointers(&inst.Credit, rhs, identifierMap)
					case "PartList":
						__gong__assignPointer(&inst.PartList, rhs, identifierMap)
					case "Measure":
						__gong__assignSliceOfPointers(&inst.Measure, rhs, identifierMap)
					}
				case *score.ScoreTimewiseMeasure:
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
				case *score.ScoreTimewiseMeasurePart:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "RefScorePart":
						__gong__assignPointer(&inst.RefScorePart, rhs, identifierMap)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Note":
						__gong__assignSliceOfPointers(&inst.Note, rhs, identifierMap)
					case "Backup":
						__gong__assignSliceOfPointers(&inst.Backup, rhs, identifierMap)
					case "Forward":
						__gong__assignSliceOfPointers(&inst.Forward, rhs, identifierMap)
					case "Direction":
						__gong__assignSliceOfPointers(&inst.Direction, rhs, identifierMap)
					case "Attributes":
						__gong__assignSliceOfPointers(&inst.Attributes, rhs, identifierMap)
					case "Harmony":
						__gong__assignSliceOfPointers(&inst.Harmony, rhs, identifierMap)
					case "FiguredBass":
						__gong__assignSliceOfPointers(&inst.FiguredBass, rhs, identifierMap)
					case "Print":
						__gong__assignSliceOfPointers(&inst.Print, rhs, identifierMap)
					case "Sound":
						__gong__assignSliceOfPointers(&inst.Sound, rhs, identifierMap)
					case "Listening":
						__gong__assignSliceOfPointers(&inst.Listening, rhs, identifierMap)
					case "Barline":
						__gong__assignSliceOfPointers(&inst.Barline, rhs, identifierMap)
					case "Grouping":
						__gong__assignSliceOfPointers(&inst.Grouping, rhs, identifierMap)
					case "Link":
						__gong__assignSliceOfPointers(&inst.Link, rhs, identifierMap)
					case "Bookmark":
						__gong__assignSliceOfPointers(&inst.Bookmark, rhs, identifierMap)
					}
				case *score.Work:
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
				case *MusicXMLDocument:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ScorePartwise":
						__gong__assignPointer(&inst.ScorePartwise, rhs, identifierMap)
					case "ScoreTimewise":
						__gong__assignPointer(&inst.ScoreTimewise, rhs, identifierMap)
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
