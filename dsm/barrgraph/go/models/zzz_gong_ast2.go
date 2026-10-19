// generated code - do not edit
package models

import (
	"embed"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"slices"
	"strings"
	"time"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
)

var (
	_ = errors.New
	_ = parser.ParseFile
	_ = filepath.Base
	_ = slices.Index[[]int, int]
	_ = strings.Trim
	_ = time.Hour
)

// swagger:ignore
type GONG__ExpressionType = gong_runtime.GONG__ExpressionType

const (
	GONG__STRUCT_INSTANCE      = gong_runtime.GONG__STRUCT_INSTANCE
	GONG__FIELD_OR_CONST_VALUE = gong_runtime.GONG__FIELD_OR_CONST_VALUE
	GONG__FIELD_VALUE          = gong_runtime.GONG__FIELD_VALUE
	GONG__ENUM_CAST_INT        = gong_runtime.GONG__ENUM_CAST_INT
	GONG__ENUM_CAST_STRING     = gong_runtime.GONG__ENUM_CAST_STRING
	GONG__IDENTIFIER_CONST     = gong_runtime.GONG__IDENTIFIER_CONST
)

// ------------------------------------------------------------------------------------------------
// STATIC AST PARSING LOGIC
// ------------------------------------------------------------------------------------------------

// GongModelUnmarshaller abstracts the logic for setting fields on a staged instance
type GongModelUnmarshaller interface {
	// Initialize creates the struct, stages it, and returns the pointer as 'any'
	Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error)

	// UnmarshallField sets a field's value based on the AST expression
	UnmarshallField(stage *Stage, instance GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error
}

type ModelUnmarshaller = GongModelUnmarshaller

// ParseAstFile Parse pathToFile and stages all instances declared in the file
func (stage *Stage) ParseAstFile(pathToFile string, preserveOrder bool) error {
	inFile, fset, err := gong_runtime.ParseAstFile(pathToFile)
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file
func (stage *Stage) ParseAstEmbeddedFile(directory embed.FS, pathToFile string) error {
	inFile, fset, err := gong_runtime.ParseAstEmbeddedFile(directory, pathToFile, stage.GetName())
	if err != nil {
		return err
	}
	return stage.ParseAstFileFromAst(inFile, fset, false)
}

// ParseAstString parses the Go source code from a string
func (stage *Stage) ParseAstString(blob string, preserveOrder bool) error {
	var inFile *ast.File
	var fset *token.FileSet
	var err error
	if strings.HasPrefix(strings.TrimSpace(blob), "package ") || strings.Contains(blob, "\npackage ") {
		fset = token.NewFileSet()
		inFile, err = parser.ParseFile(fset, "", blob, parser.ParseComments)
		if err != nil {
			return errors.New("Unable to parse " + err.Error())
		}
	} else {
		inFile, fset, err = gong_runtime.ParseAstString(blob)
		if err != nil {
			return err
		}
	}
	return stage.ParseAstFileFromAst(inFile, fset, preserveOrder)
}

// ParseAstFileFromAst traverses the AST and stages instances across stage and its sub-stages
func (stage *Stage) ParseAstFileFromAst(inFile *ast.File, fset *token.FileSet, preserveOrder bool) error {
	gong_runtime.CheckModuleVersion(stage.GetProbeIF(), inFile)

	identifierMap := make(map[string]any)

	if stage != nil {
		for _, instance := range stage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage)] = instance
			identifierMap["__models"+instance.GongGetIdentifier(stage)] = instance
		}
	}

	aliasToCanonical := make(map[string]string)
	for _, imp := range inFile.Imports {
		p := strings.Trim(imp.Path.Value, "\"`")
		alias := filepath.Base(p)
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		switch p {
		case "github.com/fullstack-lang/gong/dsm/barrgraph/go/models":
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
					} else if pkgAlias == "" || pkgAlias == "models" || len(inFile.Imports) == 0 {
						pkgAlias = "models"
					}

					switch pkgAlias {
			case "models":
				switch typeName {
				case "ArtefactType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArtefactType), stage, ident.Name, instanceName, preserveOrder)
				case "ArtefactTypeShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArtefactTypeShape), stage, ident.Name, instanceName, preserveOrder)
				case "Artist":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Artist), stage, ident.Name, instanceName, preserveOrder)
				case "ArtistShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ArtistShape), stage, ident.Name, instanceName, preserveOrder)
				case "ControlPointShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ControlPointShape), stage, ident.Name, instanceName, preserveOrder)
				case "Desk":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Desk), stage, ident.Name, instanceName, preserveOrder)
				case "Diagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Diagram), stage, ident.Name, instanceName, preserveOrder)
				case "Influence":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Influence), stage, ident.Name, instanceName, preserveOrder)
				case "InfluenceShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(InfluenceShape), stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stage, ident.Name, instanceName, preserveOrder)
				case "Movement":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Movement), stage, ident.Name, instanceName, preserveOrder)
				case "MovementShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(MovementShape), stage, ident.Name, instanceName, preserveOrder)
				case "Place":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Place), stage, ident.Name, instanceName, preserveOrder)
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
						__gong__assignPointer(&inst.ArtefactType, rhs, identifierMap)
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
						inst.DateOfDeath = GongExtractDate(rhs)
					case "Place":
						__gong__assignPointer(&inst.Place, rhs, identifierMap)
					}
				case *ArtistShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Artist":
						__gong__assignPointer(&inst.Artist, rhs, identifierMap)
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
						GongUnmarshallEnum(&inst.ImagePng_RectAnchorType, rhs)
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
						__gong__assignPointer(&inst.SelectedDiagram, rhs, identifierMap)
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
						__gong__assignSliceOfPointers(&inst.MovementShapes, rhs, identifierMap)
					case "ArtefactTypeShapes":
						__gong__assignSliceOfPointers(&inst.ArtefactTypeShapes, rhs, identifierMap)
					case "ArtistShapes":
						__gong__assignSliceOfPointers(&inst.ArtistShapes, rhs, identifierMap)
					case "InfluenceShapes":
						__gong__assignSliceOfPointers(&inst.InfluenceShapes, rhs, identifierMap)
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
						inst.StartDate = GongExtractDate(rhs)
					case "EndDate":
						inst.EndDate = GongExtractDate(rhs)
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
						GongUnmarshallEnum(&inst.MovementRectAnchorType, rhs)
					case "MovementTextAnchorType":
						GongUnmarshallEnum(&inst.MovementTextAnchorType, rhs)
					case "MovementDominantBaselineType":
						GongUnmarshallEnum(&inst.MovementDominantBaselineType, rhs)
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
						GongUnmarshallEnum(&inst.AbstractMovementRectAnchorType, rhs)
					case "AbstractMovementTextAnchorType":
						GongUnmarshallEnum(&inst.AbstractMovementTextAnchorType, rhs)
					case "AbstractDominantBaselineType":
						GongUnmarshallEnum(&inst.AbstractDominantBaselineType, rhs)
					case "MovementDateRectAnchorType":
						GongUnmarshallEnum(&inst.MovementDateRectAnchorType, rhs)
					case "MovementDateTextAnchorType":
						GongUnmarshallEnum(&inst.MovementDateTextAnchorType, rhs)
					case "MovementDateTextDominantBaselineType":
						GongUnmarshallEnum(&inst.MovementDateTextDominantBaselineType, rhs)
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
						GongUnmarshallEnum(&inst.MovementPlacesRectAnchorType, rhs)
					case "MovementPlacesTextAnchorType":
						GongUnmarshallEnum(&inst.MovementPlacesTextAnchorType, rhs)
					case "MovementPlacesDominantBaselineType":
						GongUnmarshallEnum(&inst.MovementPlacesDominantBaselineType, rhs)
					case "ArtefactTypeFontSize":
						inst.ArtefactTypeFontSize = GongExtractString(rhs)
					case "ArtefactTypeFontWeigth":
						inst.ArtefactTypeFontWeigth = GongExtractString(rhs)
					case "ArtefactTypeFontFamily":
						inst.ArtefactTypeFontFamily = GongExtractString(rhs)
					case "ArtefactTypeLetterSpacing":
						inst.ArtefactTypeLetterSpacing = GongExtractString(rhs)
					case "ArtefactTypeRectAnchorType":
						GongUnmarshallEnum(&inst.ArtefactTypeRectAnchorType, rhs)
					case "ArtefactDominantBaselineType":
						GongUnmarshallEnum(&inst.ArtefactDominantBaselineType, rhs)
					case "ArtefactTypeStrokeWidth":
						inst.ArtefactTypeStrokeWidth = GongExtractFloat(rhs)
					case "ArtistRectAnchorType":
						GongUnmarshallEnum(&inst.ArtistRectAnchorType, rhs)
					case "ArtistTextAnchorType":
						GongUnmarshallEnum(&inst.ArtistTextAnchorType, rhs)
					case "ArtistDominantBaselineType":
						GongUnmarshallEnum(&inst.ArtistDominantBaselineType, rhs)
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
						GongUnmarshallEnum(&inst.ArtistDateRectAnchorType, rhs)
					case "ArtistDateTextAnchorType":
						GongUnmarshallEnum(&inst.ArtistDateTextAnchorType, rhs)
					case "ArtistDateDominantBaselineType":
						GongUnmarshallEnum(&inst.ArtistDateDominantBaselineType, rhs)
					case "ArtistDateAndPlacesFontSize":
						inst.ArtistDateAndPlacesFontSize = GongExtractString(rhs)
					case "ArtistDateAndPlacesFontWeigth":
						inst.ArtistDateAndPlacesFontWeigth = GongExtractString(rhs)
					case "ArtistDateAndPlacesFontFamily":
						inst.ArtistDateAndPlacesFontFamily = GongExtractString(rhs)
					case "ArtistDateAndPlacesLetterSpacing":
						inst.ArtistDateAndPlacesLetterSpacing = GongExtractString(rhs)
					case "ArtistPlacesRectAnchorType":
						GongUnmarshallEnum(&inst.ArtistPlacesRectAnchorType, rhs)
					case "ArtistPlacesTextAnchorType":
						GongUnmarshallEnum(&inst.ArtistPlacesTextAnchorType, rhs)
					case "ArtistPlacesDominantBaselineType":
						GongUnmarshallEnum(&inst.ArtistPlacesDominantBaselineType, rhs)
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
						__gong__assignPointer(&inst.SourceMovement, rhs, identifierMap)
					case "SourceArtefactType":
						__gong__assignPointer(&inst.SourceArtefactType, rhs, identifierMap)
					case "SourceArtist":
						__gong__assignPointer(&inst.SourceArtist, rhs, identifierMap)
					case "TargetMovement":
						__gong__assignPointer(&inst.TargetMovement, rhs, identifierMap)
					case "TargetArtefactType":
						__gong__assignPointer(&inst.TargetArtefactType, rhs, identifierMap)
					case "TargetArtist":
						__gong__assignPointer(&inst.TargetArtist, rhs, identifierMap)
					case "IsHypothtical":
						inst.IsHypothtical = GongExtractBool(rhs)
					}
				case *InfluenceShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Influence":
						__gong__assignPointer(&inst.Influence, rhs, identifierMap)
					case "IsHidden":
						inst.IsHidden = GongExtractBool(rhs)
					case "ControlPointShapes":
						__gong__assignSliceOfPointers(&inst.ControlPointShapes, rhs, identifierMap)
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
						__gong__assignSliceOfPointers(&inst.SubLibraries, rhs, identifierMap)
					case "IsSubLibrariesNodeExpanded":
						inst.IsSubLibrariesNodeExpanded = GongExtractBool(rhs)
					case "SubLibrariesWhoseNodeIsExpanded":
						__gong__assignSliceOfPointers(&inst.SubLibrariesWhoseNodeIsExpanded, rhs, identifierMap)
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
						inst.Date = GongExtractDate(rhs)
					case "HideDate":
						inst.HideDate = GongExtractBool(rhs)
					case "Places":
						__gong__assignSliceOfPointers(&inst.Places, rhs, identifierMap)
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
						__gong__assignPointer(&inst.Movement, rhs, identifierMap)
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
		case *ast.ExprStmt:
			if call, ok := node.X.(*ast.CallExpr); ok {
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
					if sel.Sel.Name == "Unstage" {
						if ident, ok := sel.X.(*ast.Ident); ok {
							if instance, ok := identifierMap[ident.Name]; ok {
								if inst, ok := instance.(GongstructIF); ok {
									inst.UnstageVoid(stage)
								}
							}
						}
					}
					if sel.Sel.Name == "Commit" {
						if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "stage" {
							if stage.IsInDeltaMode() && stage.GetNavigationMode() != GongNavigationModeNavigating {
								stage.Commit()
							} else {
								stage.ComputeInstancesNb()
								stage.ComputeReferenceAndOrders()
								if stage.OnInitCommitCallback != nil {
									stage.OnInitCommitCallback.BeforeCommit(stage)
								}
								if stage.OnInitCommitFromBackCallback != nil {
									stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
								}
								stage.RunBeforeCommitHooks()
								stage.RunAfterCommitHooks()
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

// --- Generic Helpers for Unmarshallers (delegating to gong_runtime) ---

var GongExtractString = gong_runtime.ExtractString
var GongExtractInt = gong_runtime.ExtractInt
var GongExtractFloat = gong_runtime.ExtractFloat
var GongExtractBool = gong_runtime.ExtractBool
var GongExtractExpr = gong_runtime.ExtractExpr
var GongExtractDate = gong_runtime.ExtractDate

func __gong__extractMiddleUint(input string) (uint, error) {
	return gong_runtime.ExtractMiddleUint(input)
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
	if call, ok := rhs.(*ast.CallExpr); ok {
		funcName := ""
		var isSlices bool
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
			if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "slices" {
				isSlices = true
				funcName = sel.Sel.Name
			}
		} else if ident, ok := call.Fun.(*ast.Ident); ok {
			funcName = ident.Name
		}

		if isSlices {
			if funcName == "Delete" && len(call.Args) == 3 {
				start := GongExtractInt(call.Args[1])
				end := GongExtractInt(call.Args[2])
				if end <= len(*slice) {
					*slice = slices.Delete(*slice, start, end)
				}
			} else if funcName == "Insert" && len(call.Args) == 3 {
				index := GongExtractInt(call.Args[1])
				if index <= len(*slice) {
					if rIdent, ok := call.Args[2].(*ast.Ident); ok {
						if target, ok := identifierMap[rIdent.Name]; ok {
							if typedTarget, ok := target.(*T); ok {
								*slice = slices.Insert(*slice, index, typedTarget)
							}
						}
					}
				}
			}
		} else if funcName == "append" && len(call.Args) >= 2 {
			for i := 1; i < len(call.Args); i++ {
				if rIdent, ok := call.Args[i].(*ast.Ident); ok {
					if target, ok := identifierMap[rIdent.Name]; ok {
						if typedTarget, ok := target.(*T); ok {
							*slice = append(*slice, typedTarget)
						}
					}
				}
			}
		}
	}
}

// GongUnmarshallSliceOfPointers handles append, slices.Delete, and slices.Insert for slice fields
func GongUnmarshallSliceOfPointers[T GongstructPtr](
	slice *[]T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) (err error) {
	return gong_runtime.UnmarshallSliceOfPointers(slice, valueExpr, identifierMap)
}

// GongUnmarshallPointer handles assignment of a single pointer field
func GongUnmarshallPointer[T GongstructPtr](
	ptr *T,
	valueExpr ast.Expr,
	identifierMap map[string]GongstructIF) {
	gong_runtime.UnmarshallPointer(ptr, valueExpr, identifierMap)
}

// GongUnmarshallEnum handles assignment of enum fields (via SelectorExpr or String fallback)
func GongUnmarshallEnum[T interface{ FromCodeString(string) error }](
	ptr T,
	valueExpr ast.Expr) {
	gong_runtime.UnmarshallEnum(ptr, valueExpr)
}

// GongInitialize initializes a staged instance, sets its name, and stages it
func GongInitialize[P interface {
	GongstructIF
	StagePreserveOrder(stage *Stage, order uint)
}](instance P, stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	instance.SetName(instanceName)
	if !preserveOrder {
		instance.StageVoid(stage)
	} else {
		if newOrder, err := __gong__extractMiddleUint(identifier); err != nil {
			log.Println("UnmarshallGongstructStaging: Problem with parsing identifer", identifier)
			instance.StageVoid(stage)
		} else {
			instance.StagePreserveOrder(stage, newOrder)
		}
	}
	return instance, nil
}

// insertion point per named struct
type ArtefactTypeUnmarshaller struct{}

func (u *ArtefactTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArtefactType), stage, identifier, instanceName, preserveOrder)
}

func (u *ArtefactTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArtefactType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ArtefactTypeShapeUnmarshaller struct{}

func (u *ArtefactTypeShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArtefactTypeShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ArtefactTypeShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArtefactTypeShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ArtefactType":
		GongUnmarshallPointer(&instance.ArtefactType, valueExpr, identifierMap)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

type ArtistUnmarshaller struct{}

func (u *ArtistUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Artist), stage, identifier, instanceName, preserveOrder)
}

func (u *ArtistUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Artist)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsDead":
		instance.IsDead = GongExtractBool(valueExpr)
	case "DateOfDeath":
		instance.DateOfDeath = GongExtractDate(valueExpr)
	case "Place":
		GongUnmarshallPointer(&instance.Place, valueExpr, identifierMap)
	}
	return nil
}

type ArtistShapeUnmarshaller struct{}

func (u *ArtistShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArtistShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ArtistShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArtistShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Artist":
		GongUnmarshallPointer(&instance.Artist, valueExpr, identifierMap)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	case "ImagePng_X":
		instance.ImagePng_X = GongExtractFloat(valueExpr)
	case "ImagePng_Y":
		instance.ImagePng_Y = GongExtractFloat(valueExpr)
	case "ImagePng_Width":
		instance.ImagePng_Width = GongExtractFloat(valueExpr)
	case "ImagePng_Height":
		instance.ImagePng_Height = GongExtractFloat(valueExpr)
	case "ImagePng_X_Offset":
		instance.ImagePng_X_Offset = GongExtractFloat(valueExpr)
	case "ImagePng_Y_Offset":
		instance.ImagePng_Y_Offset = GongExtractFloat(valueExpr)
	case "ImagePng_RectAnchorType":
		GongUnmarshallEnum(&instance.ImagePng_RectAnchorType, valueExpr)
	case "ImagePngBase64Content":
		instance.ImagePngBase64Content = GongExtractString(valueExpr)
	}
	return nil
}

type ControlPointShapeUnmarshaller struct{}

func (u *ControlPointShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ControlPointShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ControlPointShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ControlPointShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X_Relative":
		instance.X_Relative = GongExtractFloat(valueExpr)
	case "Y_Relative":
		instance.Y_Relative = GongExtractFloat(valueExpr)
	case "IsStartShapeTheClosestShape":
		instance.IsStartShapeTheClosestShape = GongExtractBool(valueExpr)
	}
	return nil
}

type DeskUnmarshaller struct{}

func (u *DeskUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Desk), stage, identifier, instanceName, preserveOrder)
}

func (u *DeskUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Desk)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "SelectedDiagram":
		GongUnmarshallPointer(&instance.SelectedDiagram, valueExpr, identifierMap)
	}
	return nil
}

type DiagramUnmarshaller struct{}

func (u *DiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Diagram), stage, identifier, instanceName, preserveOrder)
}

func (u *DiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Diagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "MovementShapes":
		GongUnmarshallSliceOfPointers(&instance.MovementShapes, valueExpr, identifierMap)
	case "ArtefactTypeShapes":
		GongUnmarshallSliceOfPointers(&instance.ArtefactTypeShapes, valueExpr, identifierMap)
	case "ArtistShapes":
		GongUnmarshallSliceOfPointers(&instance.ArtistShapes, valueExpr, identifierMap)
	case "InfluenceShapes":
		GongUnmarshallSliceOfPointers(&instance.InfluenceShapes, valueExpr, identifierMap)
	case "IsEditable":
		instance.IsEditable = GongExtractBool(valueExpr)
	case "IsNodeExpanded":
		instance.IsNodeExpanded = GongExtractBool(valueExpr)
	case "IsMovementCategoryNodeExpanded":
		instance.IsMovementCategoryNodeExpanded = GongExtractBool(valueExpr)
	case "IsArtefactTypeCategoryNodeExpanded":
		instance.IsArtefactTypeCategoryNodeExpanded = GongExtractBool(valueExpr)
	case "IsArtistCategoryNodeExpanded":
		instance.IsArtistCategoryNodeExpanded = GongExtractBool(valueExpr)
	case "IsInfluenceCategoryNodeExpanded":
		instance.IsInfluenceCategoryNodeExpanded = GongExtractBool(valueExpr)
	case "IsMovementCategoryHidden":
		instance.IsMovementCategoryHidden = GongExtractBool(valueExpr)
	case "IsArtefactTypeCategoryHidden":
		instance.IsArtefactTypeCategoryHidden = GongExtractBool(valueExpr)
	case "IsArtistCategoryHidden":
		instance.IsArtistCategoryHidden = GongExtractBool(valueExpr)
	case "IsInfluenceCategoryHidden":
		instance.IsInfluenceCategoryHidden = GongExtractBool(valueExpr)
	case "StartDate":
		instance.StartDate = GongExtractDate(valueExpr)
	case "EndDate":
		instance.EndDate = GongExtractDate(valueExpr)
	case "NbYearsForIntervals":
		instance.NbYearsForIntervals = GongExtractInt(valueExpr)
	case "XMargin":
		instance.XMargin = GongExtractFloat(valueExpr)
	case "YMargin":
		instance.YMargin = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "NextVerticalDateXMargin":
		instance.NextVerticalDateXMargin = GongExtractFloat(valueExpr)
	case "RedColorCode":
		instance.RedColorCode = GongExtractString(valueExpr)
	case "BackgroundGreyColorCode":
		instance.BackgroundGreyColorCode = GongExtractString(valueExpr)
	case "GrayColorCode":
		instance.GrayColorCode = GongExtractString(valueExpr)
	case "BottomBoxYOffset":
		instance.BottomBoxYOffset = GongExtractFloat(valueExpr)
	case "BottomBoxWidth":
		instance.BottomBoxWidth = GongExtractFloat(valueExpr)
	case "BottomBoxHeigth":
		instance.BottomBoxHeigth = GongExtractFloat(valueExpr)
	case "BottomBoxFontSize":
		instance.BottomBoxFontSize = GongExtractString(valueExpr)
	case "BottomBoxFontWeigth":
		instance.BottomBoxFontWeigth = GongExtractString(valueExpr)
	case "BottomBoxFontFamily":
		instance.BottomBoxFontFamily = GongExtractString(valueExpr)
	case "BottomBoxLetterSpacing":
		instance.BottomBoxLetterSpacing = GongExtractString(valueExpr)
	case "BottomBoxLetterColorCode":
		instance.BottomBoxLetterColorCode = GongExtractString(valueExpr)
	case "MovementRectAnchorType":
		GongUnmarshallEnum(&instance.MovementRectAnchorType, valueExpr)
	case "MovementTextAnchorType":
		GongUnmarshallEnum(&instance.MovementTextAnchorType, valueExpr)
	case "MovementDominantBaselineType":
		GongUnmarshallEnum(&instance.MovementDominantBaselineType, valueExpr)
	case "MovementFontSize":
		instance.MovementFontSize = GongExtractString(valueExpr)
	case "MajorMovementFontSize":
		instance.MajorMovementFontSize = GongExtractString(valueExpr)
	case "MinorMovementFontSize":
		instance.MinorMovementFontSize = GongExtractString(valueExpr)
	case "MovementFontWeigth":
		instance.MovementFontWeigth = GongExtractString(valueExpr)
	case "MovementFontFamily":
		instance.MovementFontFamily = GongExtractString(valueExpr)
	case "MovementLetterSpacing":
		instance.MovementLetterSpacing = GongExtractString(valueExpr)
	case "AbstractMovementFontSize":
		instance.AbstractMovementFontSize = GongExtractString(valueExpr)
	case "AbstractMovementRectAnchorType":
		GongUnmarshallEnum(&instance.AbstractMovementRectAnchorType, valueExpr)
	case "AbstractMovementTextAnchorType":
		GongUnmarshallEnum(&instance.AbstractMovementTextAnchorType, valueExpr)
	case "AbstractDominantBaselineType":
		GongUnmarshallEnum(&instance.AbstractDominantBaselineType, valueExpr)
	case "MovementDateRectAnchorType":
		GongUnmarshallEnum(&instance.MovementDateRectAnchorType, valueExpr)
	case "MovementDateTextAnchorType":
		GongUnmarshallEnum(&instance.MovementDateTextAnchorType, valueExpr)
	case "MovementDateTextDominantBaselineType":
		GongUnmarshallEnum(&instance.MovementDateTextDominantBaselineType, valueExpr)
	case "MovementDateAndPlacesFontSize":
		instance.MovementDateAndPlacesFontSize = GongExtractString(valueExpr)
	case "MovementDateAndPlacesFontWeigth":
		instance.MovementDateAndPlacesFontWeigth = GongExtractString(valueExpr)
	case "MovementDateAndPlacesFontFamily":
		instance.MovementDateAndPlacesFontFamily = GongExtractString(valueExpr)
	case "MovementDateAndPlacesLetterSpacing":
		instance.MovementDateAndPlacesLetterSpacing = GongExtractString(valueExpr)
	case "MovementBelowArcY_Offset":
		instance.MovementBelowArcY_Offset = GongExtractFloat(valueExpr)
	case "MovementBelowArcY_OffsetPerPlace":
		instance.MovementBelowArcY_OffsetPerPlace = GongExtractFloat(valueExpr)
	case "MovementPlacesRectAnchorType":
		GongUnmarshallEnum(&instance.MovementPlacesRectAnchorType, valueExpr)
	case "MovementPlacesTextAnchorType":
		GongUnmarshallEnum(&instance.MovementPlacesTextAnchorType, valueExpr)
	case "MovementPlacesDominantBaselineType":
		GongUnmarshallEnum(&instance.MovementPlacesDominantBaselineType, valueExpr)
	case "ArtefactTypeFontSize":
		instance.ArtefactTypeFontSize = GongExtractString(valueExpr)
	case "ArtefactTypeFontWeigth":
		instance.ArtefactTypeFontWeigth = GongExtractString(valueExpr)
	case "ArtefactTypeFontFamily":
		instance.ArtefactTypeFontFamily = GongExtractString(valueExpr)
	case "ArtefactTypeLetterSpacing":
		instance.ArtefactTypeLetterSpacing = GongExtractString(valueExpr)
	case "ArtefactTypeRectAnchorType":
		GongUnmarshallEnum(&instance.ArtefactTypeRectAnchorType, valueExpr)
	case "ArtefactDominantBaselineType":
		GongUnmarshallEnum(&instance.ArtefactDominantBaselineType, valueExpr)
	case "ArtefactTypeStrokeWidth":
		instance.ArtefactTypeStrokeWidth = GongExtractFloat(valueExpr)
	case "ArtistRectAnchorType":
		GongUnmarshallEnum(&instance.ArtistRectAnchorType, valueExpr)
	case "ArtistTextAnchorType":
		GongUnmarshallEnum(&instance.ArtistTextAnchorType, valueExpr)
	case "ArtistDominantBaselineType":
		GongUnmarshallEnum(&instance.ArtistDominantBaselineType, valueExpr)
	case "ArtistFontSize":
		instance.ArtistFontSize = GongExtractString(valueExpr)
	case "MajorArtistFontSize":
		instance.MajorArtistFontSize = GongExtractString(valueExpr)
	case "MinorArtistFontSize":
		instance.MinorArtistFontSize = GongExtractString(valueExpr)
	case "ArtistFontWeigth":
		instance.ArtistFontWeigth = GongExtractString(valueExpr)
	case "ArtistFontFamily":
		instance.ArtistFontFamily = GongExtractString(valueExpr)
	case "ArtistLetterSpacing":
		instance.ArtistLetterSpacing = GongExtractString(valueExpr)
	case "ArtistDateRectAnchorType":
		GongUnmarshallEnum(&instance.ArtistDateRectAnchorType, valueExpr)
	case "ArtistDateTextAnchorType":
		GongUnmarshallEnum(&instance.ArtistDateTextAnchorType, valueExpr)
	case "ArtistDateDominantBaselineType":
		GongUnmarshallEnum(&instance.ArtistDateDominantBaselineType, valueExpr)
	case "ArtistDateAndPlacesFontSize":
		instance.ArtistDateAndPlacesFontSize = GongExtractString(valueExpr)
	case "ArtistDateAndPlacesFontWeigth":
		instance.ArtistDateAndPlacesFontWeigth = GongExtractString(valueExpr)
	case "ArtistDateAndPlacesFontFamily":
		instance.ArtistDateAndPlacesFontFamily = GongExtractString(valueExpr)
	case "ArtistDateAndPlacesLetterSpacing":
		instance.ArtistDateAndPlacesLetterSpacing = GongExtractString(valueExpr)
	case "ArtistPlacesRectAnchorType":
		GongUnmarshallEnum(&instance.ArtistPlacesRectAnchorType, valueExpr)
	case "ArtistPlacesTextAnchorType":
		GongUnmarshallEnum(&instance.ArtistPlacesTextAnchorType, valueExpr)
	case "ArtistPlacesDominantBaselineType":
		GongUnmarshallEnum(&instance.ArtistPlacesDominantBaselineType, valueExpr)
	case "InfluenceArrowSize":
		instance.InfluenceArrowSize = GongExtractFloat(valueExpr)
	case "InfluenceArrowStartOffset":
		instance.InfluenceArrowStartOffset = GongExtractFloat(valueExpr)
	case "InfluenceArrowEndOffset":
		instance.InfluenceArrowEndOffset = GongExtractFloat(valueExpr)
	case "InfluenceCornerRadius":
		instance.InfluenceCornerRadius = GongExtractFloat(valueExpr)
	case "InfluenceDashedLinePattern":
		instance.InfluenceDashedLinePattern = GongExtractString(valueExpr)
	}
	return nil
}

type InfluenceUnmarshaller struct{}

func (u *InfluenceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Influence), stage, identifier, instanceName, preserveOrder)
}

func (u *InfluenceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Influence)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "SourceMovement":
		GongUnmarshallPointer(&instance.SourceMovement, valueExpr, identifierMap)
	case "SourceArtefactType":
		GongUnmarshallPointer(&instance.SourceArtefactType, valueExpr, identifierMap)
	case "SourceArtist":
		GongUnmarshallPointer(&instance.SourceArtist, valueExpr, identifierMap)
	case "TargetMovement":
		GongUnmarshallPointer(&instance.TargetMovement, valueExpr, identifierMap)
	case "TargetArtefactType":
		GongUnmarshallPointer(&instance.TargetArtefactType, valueExpr, identifierMap)
	case "TargetArtist":
		GongUnmarshallPointer(&instance.TargetArtist, valueExpr, identifierMap)
	case "IsHypothtical":
		instance.IsHypothtical = GongExtractBool(valueExpr)
	}
	return nil
}

type InfluenceShapeUnmarshaller struct{}

func (u *InfluenceShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InfluenceShape), stage, identifier, instanceName, preserveOrder)
}

func (u *InfluenceShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InfluenceShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Influence":
		GongUnmarshallPointer(&instance.Influence, valueExpr, identifierMap)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	case "ControlPointShapes":
		GongUnmarshallSliceOfPointers(&instance.ControlPointShapes, valueExpr, identifierMap)
	}
	return nil
}

type LibraryUnmarshaller struct{}

func (u *LibraryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Library), stage, identifier, instanceName, preserveOrder)
}

func (u *LibraryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Library)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsRootLibrary":
		instance.IsRootLibrary = GongExtractBool(valueExpr)
	case "SubLibraries":
		GongUnmarshallSliceOfPointers(&instance.SubLibraries, valueExpr, identifierMap)
	case "IsSubLibrariesNodeExpanded":
		instance.IsSubLibrariesNodeExpanded = GongExtractBool(valueExpr)
	case "SubLibrariesWhoseNodeIsExpanded":
		GongUnmarshallSliceOfPointers(&instance.SubLibrariesWhoseNodeIsExpanded, valueExpr, identifierMap)
	case "NbPixPerCharacter":
		instance.NbPixPerCharacter = GongExtractFloat(valueExpr)
	case "LogoSVGFile":
		instance.LogoSVGFile = GongExtractString(valueExpr)
	case "IsExpandedTmp":
		instance.IsExpandedTmp = GongExtractBool(valueExpr)
	}
	return nil
}

type MovementUnmarshaller struct{}

func (u *MovementUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Movement), stage, identifier, instanceName, preserveOrder)
}

func (u *MovementUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Movement)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "Date":
		instance.Date = GongExtractDate(valueExpr)
	case "HideDate":
		instance.HideDate = GongExtractBool(valueExpr)
	case "Places":
		GongUnmarshallSliceOfPointers(&instance.Places, valueExpr, identifierMap)
	case "HasTaxonomicFilter":
		instance.HasTaxonomicFilter = GongExtractBool(valueExpr)
	case "TaxonomicFilter":
		instance.TaxonomicFilter = GongExtractString(valueExpr)
	case "IsFeatured":
		instance.IsFeatured = GongExtractBool(valueExpr)
	case "FeaturePrefix":
		instance.FeaturePrefix = GongExtractString(valueExpr)
	case "IsMajor":
		instance.IsMajor = GongExtractBool(valueExpr)
	case "IsMinor":
		instance.IsMinor = GongExtractBool(valueExpr)
	case "AdditionnalName":
		instance.AdditionnalName = GongExtractString(valueExpr)
	}
	return nil
}

type MovementShapeUnmarshaller struct{}

func (u *MovementShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MovementShape), stage, identifier, instanceName, preserveOrder)
}

func (u *MovementShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MovementShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Movement":
		GongUnmarshallPointer(&instance.Movement, valueExpr, identifierMap)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	case "IsHidden":
		instance.IsHidden = GongExtractBool(valueExpr)
	}
	return nil
}

type PlaceUnmarshaller struct{}

func (u *PlaceUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Place), stage, identifier, instanceName, preserveOrder)
}

func (u *PlaceUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Place)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}
