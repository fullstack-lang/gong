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

	if stageSet.Stage != nil {
		artefacttypeOrdered := []*ArtefactType{}
		for artefacttype := range stageSet.Stage.ArtefactTypes {
			artefacttypeOrdered = append(artefacttypeOrdered, artefacttype)
		}
		sort.Slice(artefacttypeOrdered, func(i, j int) bool {
			return stageSet.Stage.ArtefactType_stagedOrder[artefacttypeOrdered[i]] < stageSet.Stage.ArtefactType_stagedOrder[artefacttypeOrdered[j]]
		})
		for _, artefacttype := range artefacttypeOrdered {
			artefacttypeIdent := "__stage_0" + artefacttype.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ArtefactType{Name: %s}).Stage(stageSet.Stage)", artefacttypeIdent, __gong__toRawStringLiteral(artefacttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", artefacttypeIdent, __gong__toRawStringLiteral(artefacttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", artefacttypeIdent, __gong__toRawStringLiteral(artefacttype.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", artefacttypeIdent, artefacttype.IsExpanded))
		}
	}
	if stageSet.Stage != nil {
		artefacttypeshapeOrdered := []*ArtefactTypeShape{}
		for artefacttypeshape := range stageSet.Stage.ArtefactTypeShapes {
			artefacttypeshapeOrdered = append(artefacttypeshapeOrdered, artefacttypeshape)
		}
		sort.Slice(artefacttypeshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ArtefactTypeShape_stagedOrder[artefacttypeshapeOrdered[i]] < stageSet.Stage.ArtefactTypeShape_stagedOrder[artefacttypeshapeOrdered[j]]
		})
		for _, artefacttypeshape := range artefacttypeshapeOrdered {
			artefacttypeshapeIdent := "__stage_0" + artefacttypeshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ArtefactTypeShape{Name: %s}).Stage(stageSet.Stage)", artefacttypeshapeIdent, __gong__toRawStringLiteral(artefacttypeshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", artefacttypeshapeIdent, __gong__toRawStringLiteral(artefacttypeshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", artefacttypeshapeIdent, artefacttypeshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", artefacttypeshapeIdent, artefacttypeshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", artefacttypeshapeIdent, artefacttypeshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", artefacttypeshapeIdent, artefacttypeshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", artefacttypeshapeIdent, artefacttypeshape.IsHidden))
			if artefacttypeshape.ArtefactType != nil {
				targetIdent := "__stage_0" + artefacttypeshape.ArtefactType.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArtefactType = %s", artefacttypeshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		artistOrdered := []*Artist{}
		for artist := range stageSet.Stage.Artists {
			artistOrdered = append(artistOrdered, artist)
		}
		sort.Slice(artistOrdered, func(i, j int) bool {
			return stageSet.Stage.Artist_stagedOrder[artistOrdered[i]] < stageSet.Stage.Artist_stagedOrder[artistOrdered[j]]
		})
		for _, artist := range artistOrdered {
			artistIdent := "__stage_0" + artist.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Artist{Name: %s}).Stage(stageSet.Stage)", artistIdent, __gong__toRawStringLiteral(artist.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", artistIdent, __gong__toRawStringLiteral(artist.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", artistIdent, __gong__toRawStringLiteral(artist.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", artistIdent, artist.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDead = %t", artistIdent, artist.IsDead))
			values.WriteString(fmt.Sprintf("\n\t%s.DateOfDeath, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", artistIdent, artist.DateOfDeath.String()))
			if artist.Place != nil {
				targetIdent := "__stage_0" + artist.Place.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Place = %s", artistIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		artistshapeOrdered := []*ArtistShape{}
		for artistshape := range stageSet.Stage.ArtistShapes {
			artistshapeOrdered = append(artistshapeOrdered, artistshape)
		}
		sort.Slice(artistshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ArtistShape_stagedOrder[artistshapeOrdered[i]] < stageSet.Stage.ArtistShape_stagedOrder[artistshapeOrdered[j]]
		})
		for _, artistshape := range artistshapeOrdered {
			artistshapeIdent := "__stage_0" + artistshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ArtistShape{Name: %s}).Stage(stageSet.Stage)", artistshapeIdent, __gong__toRawStringLiteral(artistshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", artistshapeIdent, __gong__toRawStringLiteral(artistshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", artistshapeIdent, artistshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", artistshapeIdent, artistshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", artistshapeIdent, artistshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", artistshapeIdent, artistshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", artistshapeIdent, artistshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_X = %f", artistshapeIdent, artistshape.ImagePng_X))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_Y = %f", artistshapeIdent, artistshape.ImagePng_Y))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_Width = %f", artistshapeIdent, artistshape.ImagePng_Width))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_Height = %f", artistshapeIdent, artistshape.ImagePng_Height))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_X_Offset = %f", artistshapeIdent, artistshape.ImagePng_X_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_Y_Offset = %f", artistshapeIdent, artistshape.ImagePng_Y_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePng_RectAnchorType = %s", artistshapeIdent, __gong__toRawStringLiteral(string(artistshape.ImagePng_RectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ImagePngBase64Content = %s", artistshapeIdent, __gong__toRawStringLiteral(artistshape.ImagePngBase64Content)))
			if artistshape.Artist != nil {
				targetIdent := "__stage_0" + artistshape.Artist.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Artist = %s", artistshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		controlpointshapeOrdered := []*ControlPointShape{}
		for controlpointshape := range stageSet.Stage.ControlPointShapes {
			controlpointshapeOrdered = append(controlpointshapeOrdered, controlpointshape)
		}
		sort.Slice(controlpointshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.ControlPointShape_stagedOrder[controlpointshapeOrdered[i]] < stageSet.Stage.ControlPointShape_stagedOrder[controlpointshapeOrdered[j]]
		})
		for _, controlpointshape := range controlpointshapeOrdered {
			controlpointshapeIdent := "__stage_0" + controlpointshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.ControlPointShape{Name: %s}).Stage(stageSet.Stage)", controlpointshapeIdent, __gong__toRawStringLiteral(controlpointshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", controlpointshapeIdent, __gong__toRawStringLiteral(controlpointshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X_Relative = %f", controlpointshapeIdent, controlpointshape.X_Relative))
			values.WriteString(fmt.Sprintf("\n\t%s.Y_Relative = %f", controlpointshapeIdent, controlpointshape.Y_Relative))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStartShapeTheClosestShape = %t", controlpointshapeIdent, controlpointshape.IsStartShapeTheClosestShape))
		}
	}
	if stageSet.Stage != nil {
		deskOrdered := []*Desk{}
		for desk := range stageSet.Stage.Desks {
			deskOrdered = append(deskOrdered, desk)
		}
		sort.Slice(deskOrdered, func(i, j int) bool {
			return stageSet.Stage.Desk_stagedOrder[deskOrdered[i]] < stageSet.Stage.Desk_stagedOrder[deskOrdered[j]]
		})
		for _, desk := range deskOrdered {
			deskIdent := "__stage_0" + desk.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Desk{Name: %s}).Stage(stageSet.Stage)", deskIdent, __gong__toRawStringLiteral(desk.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", deskIdent, __gong__toRawStringLiteral(desk.Name)))
			if desk.SelectedDiagram != nil {
				targetIdent := "__stage_0" + desk.SelectedDiagram.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SelectedDiagram = %s", deskIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		diagramOrdered := []*Diagram{}
		for diagram := range stageSet.Stage.Diagrams {
			diagramOrdered = append(diagramOrdered, diagram)
		}
		sort.Slice(diagramOrdered, func(i, j int) bool {
			return stageSet.Stage.Diagram_stagedOrder[diagramOrdered[i]] < stageSet.Stage.Diagram_stagedOrder[diagramOrdered[j]]
		})
		for _, diagram := range diagramOrdered {
			diagramIdent := "__stage_0" + diagram.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Diagram{Name: %s}).Stage(stageSet.Stage)", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagramIdent, __gong__toRawStringLiteral(diagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", diagramIdent, diagram.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsChecked = %t", diagramIdent, diagram.IsChecked))
			values.WriteString(fmt.Sprintf("\n\t%s.IsEditable = %t", diagramIdent, diagram.IsEditable))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNodeExpanded = %t", diagramIdent, diagram.IsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMovementCategoryNodeExpanded = %t", diagramIdent, diagram.IsMovementCategoryNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsArtefactTypeCategoryNodeExpanded = %t", diagramIdent, diagram.IsArtefactTypeCategoryNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsArtistCategoryNodeExpanded = %t", diagramIdent, diagram.IsArtistCategoryNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInfluenceCategoryNodeExpanded = %t", diagramIdent, diagram.IsInfluenceCategoryNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMovementCategoryHidden = %t", diagramIdent, diagram.IsMovementCategoryHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsArtefactTypeCategoryHidden = %t", diagramIdent, diagram.IsArtefactTypeCategoryHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsArtistCategoryHidden = %t", diagramIdent, diagram.IsArtistCategoryHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInfluenceCategoryHidden = %t", diagramIdent, diagram.IsInfluenceCategoryHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.StartDate, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.StartDate.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.EndDate, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", diagramIdent, diagram.EndDate.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.NbYearsForIntervals = %d", diagramIdent, diagram.NbYearsForIntervals))
			values.WriteString(fmt.Sprintf("\n\t%s.XMargin = %f", diagramIdent, diagram.XMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.YMargin = %f", diagramIdent, diagram.YMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", diagramIdent, diagram.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.NextVerticalDateXMargin = %f", diagramIdent, diagram.NextVerticalDateXMargin))
			values.WriteString(fmt.Sprintf("\n\t%s.RedColorCode = %s", diagramIdent, __gong__toRawStringLiteral(diagram.RedColorCode)))
			values.WriteString(fmt.Sprintf("\n\t%s.BackgroundGreyColorCode = %s", diagramIdent, __gong__toRawStringLiteral(diagram.BackgroundGreyColorCode)))
			values.WriteString(fmt.Sprintf("\n\t%s.GrayColorCode = %s", diagramIdent, __gong__toRawStringLiteral(diagram.GrayColorCode)))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxYOffset = %f", diagramIdent, diagram.BottomBoxYOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxWidth = %f", diagramIdent, diagram.BottomBoxWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxHeigth = %f", diagramIdent, diagram.BottomBoxHeigth))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.BottomBoxFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxFontWeigth = %s", diagramIdent, __gong__toRawStringLiteral(diagram.BottomBoxFontWeigth)))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxFontFamily = %s", diagramIdent, __gong__toRawStringLiteral(diagram.BottomBoxFontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxLetterSpacing = %s", diagramIdent, __gong__toRawStringLiteral(diagram.BottomBoxLetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.BottomBoxLetterColorCode = %s", diagramIdent, __gong__toRawStringLiteral(diagram.BottomBoxLetterColorCode)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementRectAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.MovementRectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementTextAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.MovementTextAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementDominantBaselineType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.MovementDominantBaselineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.MajorMovementFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MajorMovementFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.MinorMovementFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MinorMovementFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementFontWeigth = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementFontWeigth)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementFontFamily = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementFontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementLetterSpacing = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementLetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.AbstractMovementFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.AbstractMovementFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.AbstractMovementRectAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.AbstractMovementRectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.AbstractMovementTextAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.AbstractMovementTextAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.AbstractDominantBaselineType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.AbstractDominantBaselineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementDateRectAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.MovementDateRectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementDateTextAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.MovementDateTextAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementDateTextDominantBaselineType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.MovementDateTextDominantBaselineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementDateAndPlacesFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementDateAndPlacesFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementDateAndPlacesFontWeigth = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementDateAndPlacesFontWeigth)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementDateAndPlacesFontFamily = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementDateAndPlacesFontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementDateAndPlacesLetterSpacing = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MovementDateAndPlacesLetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementBelowArcY_Offset = %f", diagramIdent, diagram.MovementBelowArcY_Offset))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementBelowArcY_OffsetPerPlace = %f", diagramIdent, diagram.MovementBelowArcY_OffsetPerPlace))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementPlacesRectAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.MovementPlacesRectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementPlacesTextAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.MovementPlacesTextAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.MovementPlacesDominantBaselineType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.MovementPlacesDominantBaselineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtefactTypeFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeFontWeigth = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtefactTypeFontWeigth)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeFontFamily = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtefactTypeFontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeLetterSpacing = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtefactTypeLetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeRectAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.ArtefactTypeRectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtefactDominantBaselineType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.ArtefactDominantBaselineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeStrokeWidth = %f", diagramIdent, diagram.ArtefactTypeStrokeWidth))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistRectAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.ArtistRectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistTextAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.ArtistTextAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistDominantBaselineType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.ArtistDominantBaselineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.MajorArtistFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MajorArtistFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.MinorArtistFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.MinorArtistFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistFontWeigth = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistFontWeigth)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistFontFamily = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistFontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistLetterSpacing = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistLetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateRectAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.ArtistDateRectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateTextAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.ArtistDateTextAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateDominantBaselineType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.ArtistDateDominantBaselineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateAndPlacesFontSize = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistDateAndPlacesFontSize)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateAndPlacesFontWeigth = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistDateAndPlacesFontWeigth)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateAndPlacesFontFamily = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistDateAndPlacesFontFamily)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistDateAndPlacesLetterSpacing = %s", diagramIdent, __gong__toRawStringLiteral(diagram.ArtistDateAndPlacesLetterSpacing)))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistPlacesRectAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.ArtistPlacesRectAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistPlacesTextAnchorType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.ArtistPlacesTextAnchorType))))
			values.WriteString(fmt.Sprintf("\n\t%s.ArtistPlacesDominantBaselineType = %s", diagramIdent, __gong__toRawStringLiteral(string(diagram.ArtistPlacesDominantBaselineType))))
			values.WriteString(fmt.Sprintf("\n\t%s.InfluenceArrowSize = %f", diagramIdent, diagram.InfluenceArrowSize))
			values.WriteString(fmt.Sprintf("\n\t%s.InfluenceArrowStartOffset = %f", diagramIdent, diagram.InfluenceArrowStartOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.InfluenceArrowEndOffset = %f", diagramIdent, diagram.InfluenceArrowEndOffset))
			values.WriteString(fmt.Sprintf("\n\t%s.InfluenceCornerRadius = %f", diagramIdent, diagram.InfluenceCornerRadius))
			values.WriteString(fmt.Sprintf("\n\t%s.InfluenceDashedLinePattern = %s", diagramIdent, __gong__toRawStringLiteral(diagram.InfluenceDashedLinePattern)))
			for _, elem := range diagram.MovementShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.MovementShapes = append(%s.MovementShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ArtefactTypeShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArtefactTypeShapes = append(%s.ArtefactTypeShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.ArtistShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ArtistShapes = append(%s.ArtistShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
			for _, elem := range diagram.InfluenceShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.InfluenceShapes = append(%s.InfluenceShapes, %s)", diagramIdent, diagramIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		influenceOrdered := []*Influence{}
		for influence := range stageSet.Stage.Influences {
			influenceOrdered = append(influenceOrdered, influence)
		}
		sort.Slice(influenceOrdered, func(i, j int) bool {
			return stageSet.Stage.Influence_stagedOrder[influenceOrdered[i]] < stageSet.Stage.Influence_stagedOrder[influenceOrdered[j]]
		})
		for _, influence := range influenceOrdered {
			influenceIdent := "__stage_0" + influence.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Influence{Name: %s}).Stage(stageSet.Stage)", influenceIdent, __gong__toRawStringLiteral(influence.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", influenceIdent, __gong__toRawStringLiteral(influence.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", influenceIdent, __gong__toRawStringLiteral(influence.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", influenceIdent, influence.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHypothtical = %t", influenceIdent, influence.IsHypothtical))
			if influence.SourceMovement != nil {
				targetIdent := "__stage_0" + influence.SourceMovement.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SourceMovement = %s", influenceIdent, targetIdent))
			}
			if influence.SourceArtefactType != nil {
				targetIdent := "__stage_0" + influence.SourceArtefactType.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SourceArtefactType = %s", influenceIdent, targetIdent))
			}
			if influence.SourceArtist != nil {
				targetIdent := "__stage_0" + influence.SourceArtist.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SourceArtist = %s", influenceIdent, targetIdent))
			}
			if influence.TargetMovement != nil {
				targetIdent := "__stage_0" + influence.TargetMovement.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TargetMovement = %s", influenceIdent, targetIdent))
			}
			if influence.TargetArtefactType != nil {
				targetIdent := "__stage_0" + influence.TargetArtefactType.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TargetArtefactType = %s", influenceIdent, targetIdent))
			}
			if influence.TargetArtist != nil {
				targetIdent := "__stage_0" + influence.TargetArtist.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TargetArtist = %s", influenceIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		influenceshapeOrdered := []*InfluenceShape{}
		for influenceshape := range stageSet.Stage.InfluenceShapes {
			influenceshapeOrdered = append(influenceshapeOrdered, influenceshape)
		}
		sort.Slice(influenceshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.InfluenceShape_stagedOrder[influenceshapeOrdered[i]] < stageSet.Stage.InfluenceShape_stagedOrder[influenceshapeOrdered[j]]
		})
		for _, influenceshape := range influenceshapeOrdered {
			influenceshapeIdent := "__stage_0" + influenceshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.InfluenceShape{Name: %s}).Stage(stageSet.Stage)", influenceshapeIdent, __gong__toRawStringLiteral(influenceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", influenceshapeIdent, __gong__toRawStringLiteral(influenceshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", influenceshapeIdent, influenceshape.IsHidden))
			if influenceshape.Influence != nil {
				targetIdent := "__stage_0" + influenceshape.Influence.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Influence = %s", influenceshapeIdent, targetIdent))
			}
			for _, elem := range influenceshape.ControlPointShapes {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ControlPointShapes = append(%s.ControlPointShapes, %s)", influenceshapeIdent, influenceshapeIdent, targetIdent))
			}
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
			libraryIdent := "__stage_0" + library.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Library{Name: %s}).Stage(stageSet.Stage)", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", libraryIdent, __gong__toRawStringLiteral(library.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", libraryIdent, __gong__toRawStringLiteral(library.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", libraryIdent, __gong__toRawStringLiteral(library.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", libraryIdent, library.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsRootLibrary = %t", libraryIdent, library.IsRootLibrary))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSubLibrariesNodeExpanded = %t", libraryIdent, library.IsSubLibrariesNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NbPixPerCharacter = %f", libraryIdent, library.NbPixPerCharacter))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", libraryIdent, __gong__toRawStringLiteral(library.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpandedTmp = %t", libraryIdent, library.IsExpandedTmp))
			for _, elem := range library.SubLibraries {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibraries = append(%s.SubLibraries, %s)", libraryIdent, libraryIdent, targetIdent))
			}
			for _, elem := range library.SubLibrariesWhoseNodeIsExpanded {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubLibrariesWhoseNodeIsExpanded = append(%s.SubLibrariesWhoseNodeIsExpanded, %s)", libraryIdent, libraryIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		movementOrdered := []*Movement{}
		for movement := range stageSet.Stage.Movements {
			movementOrdered = append(movementOrdered, movement)
		}
		sort.Slice(movementOrdered, func(i, j int) bool {
			return stageSet.Stage.Movement_stagedOrder[movementOrdered[i]] < stageSet.Stage.Movement_stagedOrder[movementOrdered[j]]
		})
		for _, movement := range movementOrdered {
			movementIdent := "__stage_0" + movement.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Movement{Name: %s}).Stage(stageSet.Stage)", movementIdent, __gong__toRawStringLiteral(movement.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", movementIdent, __gong__toRawStringLiteral(movement.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ComputedPrefix = %s", movementIdent, __gong__toRawStringLiteral(movement.ComputedPrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", movementIdent, movement.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.Date, _ = time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", \"%s\")", movementIdent, movement.Date.String()))
			values.WriteString(fmt.Sprintf("\n\t%s.HideDate = %t", movementIdent, movement.HideDate))
			values.WriteString(fmt.Sprintf("\n\t%s.HasTaxonomicFilter = %t", movementIdent, movement.HasTaxonomicFilter))
			values.WriteString(fmt.Sprintf("\n\t%s.TaxonomicFilter = %s", movementIdent, __gong__toRawStringLiteral(movement.TaxonomicFilter)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsFeatured = %t", movementIdent, movement.IsFeatured))
			values.WriteString(fmt.Sprintf("\n\t%s.FeaturePrefix = %s", movementIdent, __gong__toRawStringLiteral(movement.FeaturePrefix)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMajor = %t", movementIdent, movement.IsMajor))
			values.WriteString(fmt.Sprintf("\n\t%s.IsMinor = %t", movementIdent, movement.IsMinor))
			values.WriteString(fmt.Sprintf("\n\t%s.AdditionnalName = %s", movementIdent, __gong__toRawStringLiteral(movement.AdditionnalName)))
			for _, elem := range movement.Places {
				targetIdent := "__stage_0" + elem.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Places = append(%s.Places, %s)", movementIdent, movementIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		movementshapeOrdered := []*MovementShape{}
		for movementshape := range stageSet.Stage.MovementShapes {
			movementshapeOrdered = append(movementshapeOrdered, movementshape)
		}
		sort.Slice(movementshapeOrdered, func(i, j int) bool {
			return stageSet.Stage.MovementShape_stagedOrder[movementshapeOrdered[i]] < stageSet.Stage.MovementShape_stagedOrder[movementshapeOrdered[j]]
		})
		for _, movementshape := range movementshapeOrdered {
			movementshapeIdent := "__stage_0" + movementshape.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.MovementShape{Name: %s}).Stage(stageSet.Stage)", movementshapeIdent, __gong__toRawStringLiteral(movementshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", movementshapeIdent, __gong__toRawStringLiteral(movementshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", movementshapeIdent, movementshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", movementshapeIdent, movementshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", movementshapeIdent, movementshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", movementshapeIdent, movementshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", movementshapeIdent, movementshape.IsHidden))
			if movementshape.Movement != nil {
				targetIdent := "__stage_0" + movementshape.Movement.GongGetIdentifier(stageSet.Stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Movement = %s", movementshapeIdent, targetIdent))
			}
		}
	}
	if stageSet.Stage != nil {
		placeOrdered := []*Place{}
		for place := range stageSet.Stage.Places {
			placeOrdered = append(placeOrdered, place)
		}
		sort.Slice(placeOrdered, func(i, j int) bool {
			return stageSet.Stage.Place_stagedOrder[placeOrdered[i]] < stageSet.Stage.Place_stagedOrder[placeOrdered[j]]
		})
		for _, place := range placeOrdered {
			placeIdent := "__stage_0" + place.GongGetIdentifier(stageSet.Stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&__stage_0__.Place{Name: %s}).Stage(stageSet.Stage)", placeIdent, __gong__toRawStringLiteral(place.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", placeIdent, __gong__toRawStringLiteral(place.Name)))
		}
	}


	res = fmt.Sprintf(`// file generated by gong
package %s

import (
	"slices"
	"time"

	__stage_0__ "github.com/fullstack-lang/gong/dsm/barrgraph/go/models"
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *__stage_0__.Stage
)

// function will stage objects across all coordinated stages
func _(stageSet *__stage_0__.StageSet) {

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

					switch pkgAlias {
			case "__stage_0__":
				switch typeName {
				case "ArtefactType":
					if !preserveOrder {
						inst := (&ArtefactType{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ArtefactType)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ArtefactTypeShape":
					if !preserveOrder {
						inst := (&ArtefactTypeShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ArtefactTypeShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Artist":
					if !preserveOrder {
						inst := (&Artist{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Artist)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ArtistShape":
					if !preserveOrder {
						inst := (&ArtistShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ArtistShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "ControlPointShape":
					if !preserveOrder {
						inst := (&ControlPointShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(ControlPointShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Desk":
					if !preserveOrder {
						inst := (&Desk{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Desk)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Diagram":
					if !preserveOrder {
						inst := (&Diagram{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Diagram)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Influence":
					if !preserveOrder {
						inst := (&Influence{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Influence)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "InfluenceShape":
					if !preserveOrder {
						inst := (&InfluenceShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(InfluenceShape)
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
				case "Movement":
					if !preserveOrder {
						inst := (&Movement{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Movement)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "MovementShape":
					if !preserveOrder {
						inst := (&MovementShape{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(MovementShape)
						inst.Name = instanceName
						order, _ := __gong__extractMiddleUint(ident.Name)
						inst.StagePreserveOrder(stageSet.Stage, uint(order))
						identifierMap[ident.Name] = inst
					}
				case "Place":
					if !preserveOrder {
						inst := (&Place{Name: instanceName}).Stage(stageSet.Stage)
						identifierMap[ident.Name] = inst
					} else {
						inst := new(Place)
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
				case *ArtefactType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					}
				case *ArtefactTypeShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ArtefactType":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ArtefactType); ok {
									inst.ArtefactType = typedTarget
								}
							}
						}
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					}
				case *Artist:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsDead":
						inst.IsDead = GongExtractBool(rhs)
					case "DateOfDeath":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.DateOfDeath, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "Place":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Place); ok {
									inst.Place = typedTarget
								}
							}
						}
					}
				case *ArtistShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Artist":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Artist); ok {
									inst.Artist = typedTarget
								}
							}
						}
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					case "ImagePng_X":
						inst.ImagePng_X = GongExtractFloat(rhs)
					case "ImagePng_Y":
						inst.ImagePng_Y = GongExtractFloat(rhs)
					case "ImagePng_Width":
						inst.ImagePng_Width = GongExtractFloat(rhs)
					case "ImagePng_Height":
						inst.ImagePng_Height = GongExtractFloat(rhs)
					case "ImagePng_X_Offset":
						inst.ImagePng_X_Offset = GongExtractFloat(rhs)
					case "ImagePng_Y_Offset":
						inst.ImagePng_Y_Offset = GongExtractFloat(rhs)
					case "ImagePng_RectAnchorType":
						inst.ImagePng_RectAnchorType = RectAnchorType(GongExtractString(rhs))
					case "ImagePngBase64Content":
						inst.ImagePngBase64Content = GongExtractString(rhs)
					}
				case *ControlPointShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "X_Relative":
						inst.X_Relative = GongExtractFloat(rhs)
					case "Y_Relative":
						inst.Y_Relative = GongExtractFloat(rhs)
					case "IsStartShapeTheClosestShape":
						inst.IsStartShapeTheClosestShape = GongExtractBool(rhs)
					}
				case *Desk:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SelectedDiagram":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Diagram); ok {
									inst.SelectedDiagram = typedTarget
								}
							}
						}
					}
				case *Diagram:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsChecked":
						inst.IsChecked = GongExtractBool(rhs)
					case "MovementShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*MovementShape); ok {
										inst.MovementShapes = append(inst.MovementShapes, typedTarget)
									}
								}
							}
						}
					case "ArtefactTypeShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ArtefactTypeShape); ok {
										inst.ArtefactTypeShapes = append(inst.ArtefactTypeShapes, typedTarget)
									}
								}
							}
						}
					case "ArtistShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ArtistShape); ok {
										inst.ArtistShapes = append(inst.ArtistShapes, typedTarget)
									}
								}
							}
						}
					case "InfluenceShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*InfluenceShape); ok {
										inst.InfluenceShapes = append(inst.InfluenceShapes, typedTarget)
									}
								}
							}
						}
					case "IsEditable":
						inst.IsEditable = GongExtractBool(rhs)
					case "IsNodeExpanded":
						inst.IsNodeExpanded = GongExtractBool(rhs)
					case "IsMovementCategoryNodeExpanded":
						inst.IsMovementCategoryNodeExpanded = GongExtractBool(rhs)
					case "IsArtefactTypeCategoryNodeExpanded":
						inst.IsArtefactTypeCategoryNodeExpanded = GongExtractBool(rhs)
					case "IsArtistCategoryNodeExpanded":
						inst.IsArtistCategoryNodeExpanded = GongExtractBool(rhs)
					case "IsInfluenceCategoryNodeExpanded":
						inst.IsInfluenceCategoryNodeExpanded = GongExtractBool(rhs)
					case "IsMovementCategoryHidden":
						inst.IsMovementCategoryHidden = GongExtractBool(rhs)
					case "IsArtefactTypeCategoryHidden":
						inst.IsArtefactTypeCategoryHidden = GongExtractBool(rhs)
					case "IsArtistCategoryHidden":
						inst.IsArtistCategoryHidden = GongExtractBool(rhs)
					case "IsInfluenceCategoryHidden":
						inst.IsInfluenceCategoryHidden = GongExtractBool(rhs)
					case "StartDate":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.StartDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "EndDate":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.EndDate, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "NbYearsForIntervals":
						inst.NbYearsForIntervals = GongExtractInt(rhs)
					case "XMargin":
						inst.XMargin = GongExtractFloat(rhs)
					case "YMargin":
						inst.YMargin = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "NextVerticalDateXMargin":
						inst.NextVerticalDateXMargin = GongExtractFloat(rhs)
					case "RedColorCode":
						inst.RedColorCode = GongExtractString(rhs)
					case "BackgroundGreyColorCode":
						inst.BackgroundGreyColorCode = GongExtractString(rhs)
					case "GrayColorCode":
						inst.GrayColorCode = GongExtractString(rhs)
					case "BottomBoxYOffset":
						inst.BottomBoxYOffset = GongExtractFloat(rhs)
					case "BottomBoxWidth":
						inst.BottomBoxWidth = GongExtractFloat(rhs)
					case "BottomBoxHeigth":
						inst.BottomBoxHeigth = GongExtractFloat(rhs)
					case "BottomBoxFontSize":
						inst.BottomBoxFontSize = GongExtractString(rhs)
					case "BottomBoxFontWeigth":
						inst.BottomBoxFontWeigth = GongExtractString(rhs)
					case "BottomBoxFontFamily":
						inst.BottomBoxFontFamily = GongExtractString(rhs)
					case "BottomBoxLetterSpacing":
						inst.BottomBoxLetterSpacing = GongExtractString(rhs)
					case "BottomBoxLetterColorCode":
						inst.BottomBoxLetterColorCode = GongExtractString(rhs)
					case "MovementRectAnchorType":
						inst.MovementRectAnchorType = RectAnchorType(GongExtractString(rhs))
					case "MovementTextAnchorType":
						inst.MovementTextAnchorType = TextAnchorType(GongExtractString(rhs))
					case "MovementDominantBaselineType":
						inst.MovementDominantBaselineType = DominantBaselineType(GongExtractString(rhs))
					case "MovementFontSize":
						inst.MovementFontSize = GongExtractString(rhs)
					case "MajorMovementFontSize":
						inst.MajorMovementFontSize = GongExtractString(rhs)
					case "MinorMovementFontSize":
						inst.MinorMovementFontSize = GongExtractString(rhs)
					case "MovementFontWeigth":
						inst.MovementFontWeigth = GongExtractString(rhs)
					case "MovementFontFamily":
						inst.MovementFontFamily = GongExtractString(rhs)
					case "MovementLetterSpacing":
						inst.MovementLetterSpacing = GongExtractString(rhs)
					case "AbstractMovementFontSize":
						inst.AbstractMovementFontSize = GongExtractString(rhs)
					case "AbstractMovementRectAnchorType":
						inst.AbstractMovementRectAnchorType = RectAnchorType(GongExtractString(rhs))
					case "AbstractMovementTextAnchorType":
						inst.AbstractMovementTextAnchorType = TextAnchorType(GongExtractString(rhs))
					case "AbstractDominantBaselineType":
						inst.AbstractDominantBaselineType = DominantBaselineType(GongExtractString(rhs))
					case "MovementDateRectAnchorType":
						inst.MovementDateRectAnchorType = RectAnchorType(GongExtractString(rhs))
					case "MovementDateTextAnchorType":
						inst.MovementDateTextAnchorType = TextAnchorType(GongExtractString(rhs))
					case "MovementDateTextDominantBaselineType":
						inst.MovementDateTextDominantBaselineType = DominantBaselineType(GongExtractString(rhs))
					case "MovementDateAndPlacesFontSize":
						inst.MovementDateAndPlacesFontSize = GongExtractString(rhs)
					case "MovementDateAndPlacesFontWeigth":
						inst.MovementDateAndPlacesFontWeigth = GongExtractString(rhs)
					case "MovementDateAndPlacesFontFamily":
						inst.MovementDateAndPlacesFontFamily = GongExtractString(rhs)
					case "MovementDateAndPlacesLetterSpacing":
						inst.MovementDateAndPlacesLetterSpacing = GongExtractString(rhs)
					case "MovementBelowArcY_Offset":
						inst.MovementBelowArcY_Offset = GongExtractFloat(rhs)
					case "MovementBelowArcY_OffsetPerPlace":
						inst.MovementBelowArcY_OffsetPerPlace = GongExtractFloat(rhs)
					case "MovementPlacesRectAnchorType":
						inst.MovementPlacesRectAnchorType = RectAnchorType(GongExtractString(rhs))
					case "MovementPlacesTextAnchorType":
						inst.MovementPlacesTextAnchorType = TextAnchorType(GongExtractString(rhs))
					case "MovementPlacesDominantBaselineType":
						inst.MovementPlacesDominantBaselineType = DominantBaselineType(GongExtractString(rhs))
					case "ArtefactTypeFontSize":
						inst.ArtefactTypeFontSize = GongExtractString(rhs)
					case "ArtefactTypeFontWeigth":
						inst.ArtefactTypeFontWeigth = GongExtractString(rhs)
					case "ArtefactTypeFontFamily":
						inst.ArtefactTypeFontFamily = GongExtractString(rhs)
					case "ArtefactTypeLetterSpacing":
						inst.ArtefactTypeLetterSpacing = GongExtractString(rhs)
					case "ArtefactTypeRectAnchorType":
						inst.ArtefactTypeRectAnchorType = RectAnchorType(GongExtractString(rhs))
					case "ArtefactDominantBaselineType":
						inst.ArtefactDominantBaselineType = DominantBaselineType(GongExtractString(rhs))
					case "ArtefactTypeStrokeWidth":
						inst.ArtefactTypeStrokeWidth = GongExtractFloat(rhs)
					case "ArtistRectAnchorType":
						inst.ArtistRectAnchorType = RectAnchorType(GongExtractString(rhs))
					case "ArtistTextAnchorType":
						inst.ArtistTextAnchorType = TextAnchorType(GongExtractString(rhs))
					case "ArtistDominantBaselineType":
						inst.ArtistDominantBaselineType = DominantBaselineType(GongExtractString(rhs))
					case "ArtistFontSize":
						inst.ArtistFontSize = GongExtractString(rhs)
					case "MajorArtistFontSize":
						inst.MajorArtistFontSize = GongExtractString(rhs)
					case "MinorArtistFontSize":
						inst.MinorArtistFontSize = GongExtractString(rhs)
					case "ArtistFontWeigth":
						inst.ArtistFontWeigth = GongExtractString(rhs)
					case "ArtistFontFamily":
						inst.ArtistFontFamily = GongExtractString(rhs)
					case "ArtistLetterSpacing":
						inst.ArtistLetterSpacing = GongExtractString(rhs)
					case "ArtistDateRectAnchorType":
						inst.ArtistDateRectAnchorType = RectAnchorType(GongExtractString(rhs))
					case "ArtistDateTextAnchorType":
						inst.ArtistDateTextAnchorType = TextAnchorType(GongExtractString(rhs))
					case "ArtistDateDominantBaselineType":
						inst.ArtistDateDominantBaselineType = DominantBaselineType(GongExtractString(rhs))
					case "ArtistDateAndPlacesFontSize":
						inst.ArtistDateAndPlacesFontSize = GongExtractString(rhs)
					case "ArtistDateAndPlacesFontWeigth":
						inst.ArtistDateAndPlacesFontWeigth = GongExtractString(rhs)
					case "ArtistDateAndPlacesFontFamily":
						inst.ArtistDateAndPlacesFontFamily = GongExtractString(rhs)
					case "ArtistDateAndPlacesLetterSpacing":
						inst.ArtistDateAndPlacesLetterSpacing = GongExtractString(rhs)
					case "ArtistPlacesRectAnchorType":
						inst.ArtistPlacesRectAnchorType = RectAnchorType(GongExtractString(rhs))
					case "ArtistPlacesTextAnchorType":
						inst.ArtistPlacesTextAnchorType = TextAnchorType(GongExtractString(rhs))
					case "ArtistPlacesDominantBaselineType":
						inst.ArtistPlacesDominantBaselineType = DominantBaselineType(GongExtractString(rhs))
					case "InfluenceArrowSize":
						inst.InfluenceArrowSize = GongExtractFloat(rhs)
					case "InfluenceArrowStartOffset":
						inst.InfluenceArrowStartOffset = GongExtractFloat(rhs)
					case "InfluenceArrowEndOffset":
						inst.InfluenceArrowEndOffset = GongExtractFloat(rhs)
					case "InfluenceCornerRadius":
						inst.InfluenceCornerRadius = GongExtractFloat(rhs)
					case "InfluenceDashedLinePattern":
						inst.InfluenceDashedLinePattern = GongExtractString(rhs)
					}
				case *Influence:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "SourceMovement":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Movement); ok {
									inst.SourceMovement = typedTarget
								}
							}
						}
					case "SourceArtefactType":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ArtefactType); ok {
									inst.SourceArtefactType = typedTarget
								}
							}
						}
					case "SourceArtist":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Artist); ok {
									inst.SourceArtist = typedTarget
								}
							}
						}
					case "TargetMovement":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Movement); ok {
									inst.TargetMovement = typedTarget
								}
							}
						}
					case "TargetArtefactType":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ArtefactType); ok {
									inst.TargetArtefactType = typedTarget
								}
							}
						}
					case "TargetArtist":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Artist); ok {
									inst.TargetArtist = typedTarget
								}
							}
						}
					case "IsHypothtical":
						inst.IsHypothtical = GongExtractBool(rhs)
					}
				case *InfluenceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Influence":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Influence); ok {
									inst.Influence = typedTarget
								}
							}
						}
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					case "ControlPointShapes":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*ControlPointShape); ok {
										inst.ControlPointShapes = append(inst.ControlPointShapes, typedTarget)
									}
								}
							}
						}
					}
				case *Library:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "IsRootLibrary":
						inst.IsRootLibrary = GongExtractBool(rhs)
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
					case "IsSubLibrariesNodeExpanded":
						inst.IsSubLibrariesNodeExpanded = GongExtractBool(rhs)
					case "SubLibrariesWhoseNodeIsExpanded":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Library); ok {
										inst.SubLibrariesWhoseNodeIsExpanded = append(inst.SubLibrariesWhoseNodeIsExpanded, typedTarget)
									}
								}
							}
						}
					case "NbPixPerCharacter":
						inst.NbPixPerCharacter = GongExtractFloat(rhs)
					case "LogoSVGFile":
						inst.LogoSVGFile = GongExtractString(rhs)
					case "IsExpandedTmp":
						inst.IsExpandedTmp = GongExtractBool(rhs)
					}
				case *Movement:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "ComputedPrefix":
						inst.ComputedPrefix = GongExtractString(rhs)
					case "IsExpanded":
						inst.IsExpanded = GongExtractBool(rhs)
					case "Date":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if bl, ok := call.Args[1].(*ast.BasicLit); ok {
								inst.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Trim(bl.Value, "\"`"))
							}
						}
					case "HideDate":
						inst.HideDate = GongExtractBool(rhs)
					case "Places":
						if call, ok := rhs.(*ast.CallExpr); ok && len(call.Args) == 2 {
							if rIdent, ok := call.Args[1].(*ast.Ident); ok {
								if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Place); ok {
										inst.Places = append(inst.Places, typedTarget)
									}
								}
							}
						}
					case "HasTaxonomicFilter":
						inst.HasTaxonomicFilter = GongExtractBool(rhs)
					case "TaxonomicFilter":
						inst.TaxonomicFilter = GongExtractString(rhs)
					case "IsFeatured":
						inst.IsFeatured = GongExtractBool(rhs)
					case "FeaturePrefix":
						inst.FeaturePrefix = GongExtractString(rhs)
					case "IsMajor":
						inst.IsMajor = GongExtractBool(rhs)
					case "IsMinor":
						inst.IsMinor = GongExtractBool(rhs)
					case "AdditionnalName":
						inst.AdditionnalName = GongExtractString(rhs)
					}
				case *MovementShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Movement":
						if rIdent, ok := rhs.(*ast.Ident); ok {
							if target, ok := identifierMap[rIdent.Name]; ok {
								if typedTarget, ok := target.(*Movement); ok {
									inst.Movement = typedTarget
								}
							}
						}
					case "X":
						inst.X = GongExtractFloat(rhs)
					case "Y":
						inst.Y = GongExtractFloat(rhs)
					case "Width":
						inst.Width = GongExtractFloat(rhs)
					case "Height":
						inst.Height = GongExtractFloat(rhs)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					}
				case *Place:
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
