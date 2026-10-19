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
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/stool"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/clock"
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

	if stage.ClockStage != nil {
		for _, instance := range stage.ClockStage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage.ClockStage)] = instance
			identifierMap["__clock"+instance.GongGetIdentifier(stage.ClockStage)] = instance
		}
	}
	if stage.MusicStage != nil {
		for _, instance := range stage.MusicStage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage.MusicStage)] = instance
			identifierMap["__music"+instance.GongGetIdentifier(stage.MusicStage)] = instance
		}
	}
	if stage.StoolStage != nil {
		for _, instance := range stage.StoolStage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage.StoolStage)] = instance
			identifierMap["__stool"+instance.GongGetIdentifier(stage.StoolStage)] = instance
		}
	}
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
					} else if pkgAlias == "" || pkgAlias == "models" || len(inFile.Imports) == 0 {
						pkgAlias = "models"
					}

					switch pkgAlias {
			case "clock":
				switch typeName {
				case "ClockAbstract":
					identifierMap[ident.Name] = __gong__stageSetInit(new(clock.ClockAbstract), stage.ClockStage, ident.Name, instanceName, preserveOrder)
				}
			case "music":
				switch typeName {
				case "MusicAbstract":
					identifierMap[ident.Name] = __gong__stageSetInit(new(music.MusicAbstract), stage.MusicStage, ident.Name, instanceName, preserveOrder)
				}
			case "stool":
				switch typeName {
				case "StoolAbstract":
					identifierMap[ident.Name] = __gong__stageSetInit(new(stool.StoolAbstract), stage.StoolStage, ident.Name, instanceName, preserveOrder)
				}
			case "models":
				switch typeName {
				case "Angle0Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Angle0Shape), stage, ident.Name, instanceName, preserveOrder)
				case "BottomCurvePlane1Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BottomCurvePlane1Shape), stage, ident.Name, instanceName, preserveOrder)
				case "BottomCurvePlane2Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BottomCurvePlane2Shape), stage, ident.Name, instanceName, preserveOrder)
				case "CarvedOutBottomCurvePlane1Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CarvedOutBottomCurvePlane1Shape), stage, ident.Name, instanceName, preserveOrder)
				case "CarvedOutTopCurvePlane1Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CarvedOutTopCurvePlane1Shape), stage, ident.Name, instanceName, preserveOrder)
				case "CarvedOutVaseTrapezeRingShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CarvedOutVaseTrapezeRingShape), stage, ident.Name, instanceName, preserveOrder)
				case "Circumference3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Circumference3DShape), stage, ident.Name, instanceName, preserveOrder)
				case "Clock2DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Clock2DDiagram), stage, ident.Name, instanceName, preserveOrder)
				case "Clock3DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Clock3DDiagram), stage, ident.Name, instanceName, preserveOrder)
				case "CutLine3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CutLine3DShape), stage, ident.Name, instanceName, preserveOrder)
				case "Leaves3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Leaves3DShape), stage, ident.Name, instanceName, preserveOrder)
				case "Library":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Library), stage, ident.Name, instanceName, preserveOrder)
				case "OriginalPoints3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(OriginalPoints3DShape), stage, ident.Name, instanceName, preserveOrder)
				case "ParastichyMCurves3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParastichyMCurves3DShape), stage, ident.Name, instanceName, preserveOrder)
				case "ParastichyNCurves3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ParastichyNCurves3DShape), stage, ident.Name, instanceName, preserveOrder)
				case "Plant2DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Plant2DDiagram), stage, ident.Name, instanceName, preserveOrder)
				case "Plant3DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Plant3DDiagram), stage, ident.Name, instanceName, preserveOrder)
				case "PlantAbstract":
					identifierMap[ident.Name] = __gong__stageSetInit(new(PlantAbstract), stage, ident.Name, instanceName, preserveOrder)
				case "Rendered3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Rendered3DShape), stage, ident.Name, instanceName, preserveOrder)
				case "SampledPoints3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(SampledPoints3DShape), stage, ident.Name, instanceName, preserveOrder)
				case "StackOfRotatedCarvedOutVaseTrapezeRingsShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StackOfRotatedCarvedOutVaseTrapezeRingsShape), stage, ident.Name, instanceName, preserveOrder)
				case "StemCylinder3DShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(StemCylinder3DShape), stage, ident.Name, instanceName, preserveOrder)
				case "Stool2DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Stool2DDiagram), stage, ident.Name, instanceName, preserveOrder)
				case "Stool3DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Stool3DDiagram), stage, ident.Name, instanceName, preserveOrder)
				case "TopCurvePlane1Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TopCurvePlane1Shape), stage, ident.Name, instanceName, preserveOrder)
				case "TopCurvePlane2Shape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TopCurvePlane2Shape), stage, ident.Name, instanceName, preserveOrder)
				case "TubeVase3DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TubeVase3DDiagram), stage, ident.Name, instanceName, preserveOrder)
				case "TubeVaseAbstract":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TubeVaseAbstract), stage, ident.Name, instanceName, preserveOrder)
				case "Vase2DDiagram":
					identifierMap[ident.Name] = __gong__stageSetInit(new(Vase2DDiagram), stage, ident.Name, instanceName, preserveOrder)
				case "VaseTrapezeBasePlateShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(VaseTrapezeBasePlateShape), stage, ident.Name, instanceName, preserveOrder)
				case "VaseTrapezeRingShape":
					identifierMap[ident.Name] = __gong__stageSetInit(new(VaseTrapezeRingShape), stage, ident.Name, instanceName, preserveOrder)
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
				case *CarvedOutBottomCurvePlane1Shape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *CarvedOutTopCurvePlane1Shape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *CarvedOutVaseTrapezeRingShape:
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
						GongUnmarshallEnum(&inst.PlantType, rhs)
					case "TubeVaseAbstract":
						__gong__assignPointer(&inst.TubeVaseAbstract, rhs, identifierMap)
					case "StoolAbstract":
						__gong__assignPointer(&inst.StoolAbstract, rhs, identifierMap)
					case "ClockAbstract":
						__gong__assignPointer(&inst.ClockAbstract, rhs, identifierMap)
					case "MusicAbstract":
						__gong__assignPointer(&inst.MusicAbstract, rhs, identifierMap)
					case "CurrentView":
						GongUnmarshallEnum(&inst.CurrentView, rhs)
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
				case *StackOfRotatedCarvedOutVaseTrapezeRingsShape:
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
					case "IsHiddenCarvedOutTopCurvePlane1Shape":
						inst.IsHiddenCarvedOutTopCurvePlane1Shape = GongExtractBool(rhs)
					case "IsHiddenCarvedOutBottomCurvePlane1Shape":
						inst.IsHiddenCarvedOutBottomCurvePlane1Shape = GongExtractBool(rhs)
					case "IsHiddenVaseTrapezeRingShape":
						inst.IsHiddenVaseTrapezeRingShape = GongExtractBool(rhs)
					case "IsHiddenCarvedOutVaseTrapezeRingShape":
						inst.IsHiddenCarvedOutVaseTrapezeRingShape = GongExtractBool(rhs)
					case "IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape":
						inst.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape = GongExtractBool(rhs)
					case "IsHiddenVaseTrapezeBasePlateShape":
						inst.IsHiddenVaseTrapezeBasePlateShape = GongExtractBool(rhs)
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
					case "CarvedOutTopCurvePlane1Shape":
						__gong__assignPointer(&inst.CarvedOutTopCurvePlane1Shape, rhs, identifierMap)
					case "CarvedOutBottomCurvePlane1Shape":
						__gong__assignPointer(&inst.CarvedOutBottomCurvePlane1Shape, rhs, identifierMap)
					case "VaseTrapezeRingShape":
						__gong__assignPointer(&inst.VaseTrapezeRingShape, rhs, identifierMap)
					case "CarvedOutVaseTrapezeRingShape":
						__gong__assignPointer(&inst.CarvedOutVaseTrapezeRingShape, rhs, identifierMap)
					case "StackOfRotatedCarvedOutVaseTrapezeRingsShape":
						__gong__assignPointer(&inst.StackOfRotatedCarvedOutVaseTrapezeRingsShape, rhs, identifierMap)
					case "VaseTrapezeBasePlateShape":
						__gong__assignPointer(&inst.VaseTrapezeBasePlateShape, rhs, identifierMap)
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
					case "TopPlaneHeight":
						inst.TopPlaneHeight = GongExtractFloat(rhs)
					case "BottomPlaneHeight":
						inst.BottomPlaneHeight = GongExtractFloat(rhs)
					case "BaseRingBottomPlaneHeight":
						inst.BaseRingBottomPlaneHeight = GongExtractFloat(rhs)
					case "ProjectionAngle":
						inst.ProjectionAngle = GongExtractFloat(rhs)
					case "BasePlateHeight":
						inst.BasePlateHeight = GongExtractFloat(rhs)
					case "CarvedOutTopRingsParameter":
						inst.CarvedOutTopRingsParameter = GongExtractFloat(rhs)
					case "BulbousStartTangentMagnitude":
						inst.BulbousStartTangentMagnitude = GongExtractFloat(rhs)
					case "BulbousEndAngle":
						inst.BulbousEndAngle = GongExtractFloat(rhs)
					case "BulbousEndTangentMagnitude":
						inst.BulbousEndTangentMagnitude = GongExtractFloat(rhs)
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
				case *VaseTrapezeBasePlateShape:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
		case *ast.ExprStmt:
			if call, ok := node.X.(*ast.CallExpr); ok {
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
					if sel.Sel.Name == "Unstage" {
						if ident, ok := sel.X.(*ast.Ident); ok {
							if instance, ok := identifierMap[ident.Name]; ok {
								if inst, ok := instance.(clock.GongstructIF); ok && stage.ClockStage != nil {
									inst.UnstageVoid(stage.ClockStage)
								}
								if inst, ok := instance.(music.GongstructIF); ok && stage.MusicStage != nil {
									inst.UnstageVoid(stage.MusicStage)
								}
								if inst, ok := instance.(stool.GongstructIF); ok && stage.StoolStage != nil {
									inst.UnstageVoid(stage.StoolStage)
								}
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
type Angle0ShapeUnmarshaller struct{}

func (u *Angle0ShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Angle0Shape), stage, identifier, instanceName, preserveOrder)
}

func (u *Angle0ShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Angle0Shape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ArcNormalVectorShapeUnmarshaller struct{}

func (u *ArcNormalVectorShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArcNormalVectorShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ArcNormalVectorShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArcNormalVectorShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type ArcNormalVectorShapeGridUnmarshaller struct{}

func (u *ArcNormalVectorShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ArcNormalVectorShapeGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *ArcNormalVectorShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ArcNormalVectorShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ArcNormalVectorShapes":
		GongUnmarshallSliceOfPointers(&instance.ArcNormalVectorShapes, valueExpr, identifierMap)
	}
	return nil
}

type AxesShapeUnmarshaller struct{}

func (u *AxesShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AxesShape), stage, identifier, instanceName, preserveOrder)
}

func (u *AxesShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AxesShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "LengthX":
		instance.LengthX = GongExtractFloat(valueExpr)
	case "LengthY":
		instance.LengthY = GongExtractFloat(valueExpr)
	case "IsWithHiddenHandle":
		instance.IsWithHiddenHandle = GongExtractBool(valueExpr)
	}
	return nil
}

type BaseVectorShapeUnmarshaller struct{}

func (u *BaseVectorShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(BaseVectorShape), stage, identifier, instanceName, preserveOrder)
}

func (u *BaseVectorShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BaseVectorShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type BaseVectorShapeGridUnmarshaller struct{}

func (u *BaseVectorShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(BaseVectorShapeGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *BaseVectorShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BaseVectorShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BaseVectorShapes":
		GongUnmarshallSliceOfPointers(&instance.BaseVectorShapes, valueExpr, identifierMap)
	}
	return nil
}

type BottomCurvePlane1ShapeUnmarshaller struct{}

func (u *BottomCurvePlane1ShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(BottomCurvePlane1Shape), stage, identifier, instanceName, preserveOrder)
}

func (u *BottomCurvePlane1ShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BottomCurvePlane1Shape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type BottomCurvePlane2ShapeUnmarshaller struct{}

func (u *BottomCurvePlane2ShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(BottomCurvePlane2Shape), stage, identifier, instanceName, preserveOrder)
}

func (u *BottomCurvePlane2ShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BottomCurvePlane2Shape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type CarvedOutBottomCurvePlane1ShapeUnmarshaller struct{}

func (u *CarvedOutBottomCurvePlane1ShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CarvedOutBottomCurvePlane1Shape), stage, identifier, instanceName, preserveOrder)
}

func (u *CarvedOutBottomCurvePlane1ShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CarvedOutBottomCurvePlane1Shape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type CarvedOutTopCurvePlane1ShapeUnmarshaller struct{}

func (u *CarvedOutTopCurvePlane1ShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CarvedOutTopCurvePlane1Shape), stage, identifier, instanceName, preserveOrder)
}

func (u *CarvedOutTopCurvePlane1ShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CarvedOutTopCurvePlane1Shape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type CarvedOutVaseTrapezeRingShapeUnmarshaller struct{}

func (u *CarvedOutVaseTrapezeRingShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CarvedOutVaseTrapezeRingShape), stage, identifier, instanceName, preserveOrder)
}

func (u *CarvedOutVaseTrapezeRingShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CarvedOutVaseTrapezeRingShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ChosenP1P2PairShapeUnmarshaller struct{}

func (u *ChosenP1P2PairShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ChosenP1P2PairShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ChosenP1P2PairShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ChosenP1P2PairShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "P1X":
		instance.P1X = GongExtractFloat(valueExpr)
	case "P1Y":
		instance.P1Y = GongExtractFloat(valueExpr)
	case "P2X":
		instance.P2X = GongExtractFloat(valueExpr)
	case "P2Y":
		instance.P2Y = GongExtractFloat(valueExpr)
	case "PxX":
		instance.PxX = GongExtractFloat(valueExpr)
	case "PxY":
		instance.PxY = GongExtractFloat(valueExpr)
	case "DistanceP1Px":
		instance.DistanceP1Px = GongExtractFloat(valueExpr)
	case "DistanceP2Px":
		instance.DistanceP2Px = GongExtractFloat(valueExpr)
	case "DistanceSum":
		instance.DistanceSum = GongExtractFloat(valueExpr)
	}
	return nil
}

type CircleGridShapeUnmarshaller struct{}

func (u *CircleGridShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CircleGridShape), stage, identifier, instanceName, preserveOrder)
}

func (u *CircleGridShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CircleGridShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type Circumference3DShapeUnmarshaller struct{}

func (u *Circumference3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Circumference3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *Circumference3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Circumference3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type Clock2DDiagramUnmarshaller struct{}

func (u *Clock2DDiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Clock2DDiagram), stage, identifier, instanceName, preserveOrder)
}

func (u *Clock2DDiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Clock2DDiagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Zoom":
		instance.Zoom = GongExtractFloat(valueExpr)
	case "IsHiddenAxesShape":
		instance.IsHiddenAxesShape = GongExtractBool(valueExpr)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type Clock3DDiagramUnmarshaller struct{}

func (u *Clock3DDiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Clock3DDiagram), stage, identifier, instanceName, preserveOrder)
}

func (u *Clock3DDiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Clock3DDiagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsHiddenClockTopCurveShape":
		instance.IsHiddenClockTopCurveShape = GongExtractBool(valueExpr)
	case "ClockTopCurveShape":
		GongUnmarshallPointer(&instance.ClockTopCurveShape, valueExpr, identifierMap)
	case "IsHiddenTorus3DShape":
		instance.IsHiddenTorus3DShape = GongExtractBool(valueExpr)
	case "Torus3DShape":
		GongUnmarshallPointer(&instance.Torus3DShape, valueExpr, identifierMap)
	case "IsHiddenSampledPoints3DShape":
		instance.IsHiddenSampledPoints3DShape = GongExtractBool(valueExpr)
	case "SampledPoints3DShape":
		GongUnmarshallPointer(&instance.SampledPoints3DShape, valueExpr, identifierMap)
	case "IsHiddenTiledFloor3DShape":
		instance.IsHiddenTiledFloor3DShape = GongExtractBool(valueExpr)
	case "TiledFloor3DShape":
		GongUnmarshallPointer(&instance.TiledFloor3DShape, valueExpr, identifierMap)
	case "Rendered3DShape":
		GongUnmarshallPointer(&instance.Rendered3DShape, valueExpr, identifierMap)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type ClockTopCurveShapeUnmarshaller struct{}

func (u *ClockTopCurveShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ClockTopCurveShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ClockTopCurveShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ClockTopCurveShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type CutLine3DShapeUnmarshaller struct{}

func (u *CutLine3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CutLine3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *CutLine3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CutLine3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type EndArcShapeUnmarshaller struct{}

func (u *EndArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EndArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *EndArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EndArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type EndArcShapeGridUnmarshaller struct{}

func (u *EndArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EndArcShapeGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *EndArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EndArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EndArcShapes":
		GongUnmarshallSliceOfPointers(&instance.EndArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type EndHalfwayArcShapeUnmarshaller struct{}

func (u *EndHalfwayArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EndHalfwayArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *EndHalfwayArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EndHalfwayArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type EndHalfwayArcShapeGridUnmarshaller struct{}

func (u *EndHalfwayArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EndHalfwayArcShapeGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *EndHalfwayArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EndHalfwayArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "EndHalfwayArcShapes":
		GongUnmarshallSliceOfPointers(&instance.EndHalfwayArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type ExplanationTextShapeUnmarshaller struct{}

func (u *ExplanationTextShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ExplanationTextShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ExplanationTextShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ExplanationTextShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type Eye3DShapeUnmarshaller struct{}

func (u *Eye3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Eye3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *Eye3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Eye3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type EyeCornersSampledPoints3DShapeUnmarshaller struct{}

func (u *EyeCornersSampledPoints3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EyeCornersSampledPoints3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *EyeCornersSampledPoints3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EyeCornersSampledPoints3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type EyeSampledPoints3DShapeUnmarshaller struct{}

func (u *EyeSampledPoints3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EyeSampledPoints3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *EyeSampledPoints3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EyeSampledPoints3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type EyeSeatBottomCurveShapeUnmarshaller struct{}

func (u *EyeSeatBottomCurveShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EyeSeatBottomCurveShape), stage, identifier, instanceName, preserveOrder)
}

func (u *EyeSeatBottomCurveShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EyeSeatBottomCurveShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type EyeStoolBottomCurveShapeUnmarshaller struct{}

func (u *EyeStoolBottomCurveShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EyeStoolBottomCurveShape), stage, identifier, instanceName, preserveOrder)
}

func (u *EyeStoolBottomCurveShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EyeStoolBottomCurveShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type EyeVolume3DShapeUnmarshaller struct{}

func (u *EyeVolume3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(EyeVolume3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *EyeVolume3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*EyeVolume3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type GridPathShapeUnmarshaller struct{}

func (u *GridPathShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GridPathShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GridPathShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GridPathShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type GrowthCurve2DUnmarshaller struct{}

func (u *GrowthCurve2DUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GrowthCurve2D), stage, identifier, instanceName, preserveOrder)
}

func (u *GrowthCurve2DUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GrowthCurve2D)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.StartHalfwayArcShapeGrid, valueExpr, identifierMap)
	case "EndHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.EndHalfwayArcShapeGrid, valueExpr, identifierMap)
	}
	return nil
}

type GrowthCurve2DRibbonUnmarshaller struct{}

func (u *GrowthCurve2DRibbonUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GrowthCurve2DRibbon), stage, identifier, instanceName, preserveOrder)
}

func (u *GrowthCurve2DRibbonUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GrowthCurve2DRibbon)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GrowthCurve2DRibbonStartShapes":
		GongUnmarshallSliceOfPointers(&instance.GrowthCurve2DRibbonStartShapes, valueExpr, identifierMap)
	case "GrowthCurve2DRibbonEndShapes":
		GongUnmarshallSliceOfPointers(&instance.GrowthCurve2DRibbonEndShapes, valueExpr, identifierMap)
	}
	return nil
}

type GrowthCurve2DRibbonEndShapeUnmarshaller struct{}

func (u *GrowthCurve2DRibbonEndShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GrowthCurve2DRibbonEndShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GrowthCurve2DRibbonEndShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GrowthCurve2DRibbonEndShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type GrowthCurve2DRibbonStartShapeUnmarshaller struct{}

func (u *GrowthCurve2DRibbonStartShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GrowthCurve2DRibbonStartShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GrowthCurve2DRibbonStartShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GrowthCurve2DRibbonStartShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type GrowthCurveRhombusGridShapeUnmarshaller struct{}

func (u *GrowthCurveRhombusGridShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GrowthCurveRhombusGridShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GrowthCurveRhombusGridShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GrowthCurveRhombusGridShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "GrowthCurveRhombusShapes":
		GongUnmarshallSliceOfPointers(&instance.GrowthCurveRhombusShapes, valueExpr, identifierMap)
	}
	return nil
}

type GrowthCurveRhombusShapeUnmarshaller struct{}

func (u *GrowthCurveRhombusShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GrowthCurveRhombusShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GrowthCurveRhombusShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GrowthCurveRhombusShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type GrowthVectorShapeUnmarshaller struct{}

func (u *GrowthVectorShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(GrowthVectorShape), stage, identifier, instanceName, preserveOrder)
}

func (u *GrowthVectorShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*GrowthVectorShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type InitialRhombusGridShapeUnmarshaller struct{}

func (u *InitialRhombusGridShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InitialRhombusGridShape), stage, identifier, instanceName, preserveOrder)
}

func (u *InitialRhombusGridShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InitialRhombusGridShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "InitialRhombusShapes":
		GongUnmarshallSliceOfPointers(&instance.InitialRhombusShapes, valueExpr, identifierMap)
	}
	return nil
}

type InitialRhombusShapeUnmarshaller struct{}

func (u *InitialRhombusShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(InitialRhombusShape), stage, identifier, instanceName, preserveOrder)
}

func (u *InitialRhombusShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*InitialRhombusShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type Key3DShapeUnmarshaller struct{}

func (u *Key3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Key3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *Key3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Key3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type KeyHole3DShapeUnmarshaller struct{}

func (u *KeyHole3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(KeyHole3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *KeyHole3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*KeyHole3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type KeyHoleShapeUnmarshaller struct{}

func (u *KeyHoleShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(KeyHoleShape), stage, identifier, instanceName, preserveOrder)
}

func (u *KeyHoleShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*KeyHoleShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	case "Width":
		instance.Width = GongExtractFloat(valueExpr)
	case "Height":
		instance.Height = GongExtractFloat(valueExpr)
	}
	return nil
}

type Leaves3DShapeUnmarshaller struct{}

func (u *Leaves3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Leaves3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *Leaves3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Leaves3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
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
	case "Plants":
		GongUnmarshallSliceOfPointers(&instance.Plants, valueExpr, identifierMap)
	case "SubLibraries":
		GongUnmarshallSliceOfPointers(&instance.SubLibraries, valueExpr, identifierMap)
	case "NbPixPerCharacter":
		instance.NbPixPerCharacter = GongExtractFloat(valueExpr)
	case "LogoSVGFile":
		instance.LogoSVGFile = GongExtractString(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsRootLibrary":
		instance.IsRootLibrary = GongExtractBool(valueExpr)
	}
	return nil
}

type MidArcVectorShapeUnmarshaller struct{}

func (u *MidArcVectorShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MidArcVectorShape), stage, identifier, instanceName, preserveOrder)
}

func (u *MidArcVectorShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MidArcVectorShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type MidArcVectorShapeGridUnmarshaller struct{}

func (u *MidArcVectorShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(MidArcVectorShapeGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *MidArcVectorShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*MidArcVectorShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "MidArcVectorShapes":
		GongUnmarshallSliceOfPointers(&instance.MidArcVectorShapes, valueExpr, identifierMap)
	}
	return nil
}

type OriginalPoints3DShapeUnmarshaller struct{}

func (u *OriginalPoints3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(OriginalPoints3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *OriginalPoints3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*OriginalPoints3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ParastichyMCurves3DShapeUnmarshaller struct{}

func (u *ParastichyMCurves3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParastichyMCurves3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ParastichyMCurves3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParastichyMCurves3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ParastichyNCurves3DShapeUnmarshaller struct{}

func (u *ParastichyNCurves3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ParastichyNCurves3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ParastichyNCurves3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ParastichyNCurves3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type PartiallyGrowthCurve2DRibbonUnmarshaller struct{}

func (u *PartiallyGrowthCurve2DRibbonUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyGrowthCurve2DRibbon), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyGrowthCurve2DRibbonUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyGrowthCurve2DRibbon)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PartiallyGrowthCurve2DRibbonStartShapes":
		GongUnmarshallSliceOfPointers(&instance.PartiallyGrowthCurve2DRibbonStartShapes, valueExpr, identifierMap)
	case "PartiallyGrowthCurve2DRibbonEndShapes":
		GongUnmarshallSliceOfPointers(&instance.PartiallyGrowthCurve2DRibbonEndShapes, valueExpr, identifierMap)
	}
	return nil
}

type PartiallyGrowthCurve2DRibbonEndShapeUnmarshaller struct{}

func (u *PartiallyGrowthCurve2DRibbonEndShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyGrowthCurve2DRibbonEndShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyGrowthCurve2DRibbonEndShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyGrowthCurve2DRibbonEndShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type PartiallyGrowthCurve2DRibbonStartShapeUnmarshaller struct{}

func (u *PartiallyGrowthCurve2DRibbonStartShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyGrowthCurve2DRibbonStartShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyGrowthCurve2DRibbonStartShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyGrowthCurve2DRibbonStartShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type PartiallyGrowthCurve2DTrajectoryUnmarshaller struct{}

func (u *PartiallyGrowthCurve2DTrajectoryUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyGrowthCurve2DTrajectory), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyGrowthCurve2DTrajectoryUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyGrowthCurve2DTrajectory)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PartiallyGrowthCurve2DTrajectoryShapes":
		GongUnmarshallSliceOfPointers(&instance.PartiallyGrowthCurve2DTrajectoryShapes, valueExpr, identifierMap)
	}
	return nil
}

type PartiallyGrowthCurve2DTrajectoryP1CurveShapeUnmarshaller struct{}

func (u *PartiallyGrowthCurve2DTrajectoryP1CurveShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyGrowthCurve2DTrajectoryP1CurveShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyGrowthCurve2DTrajectoryP1CurveShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyGrowthCurve2DTrajectoryP1CurveShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type PartiallyGrowthCurve2DTrajectoryP1P2Unmarshaller struct{}

func (u *PartiallyGrowthCurve2DTrajectoryP1P2Unmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyGrowthCurve2DTrajectoryP1P2), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyGrowthCurve2DTrajectoryP1P2Unmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyGrowthCurve2DTrajectoryP1P2)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "P1PointShapes":
		GongUnmarshallSliceOfPointers(&instance.P1PointShapes, valueExpr, identifierMap)
	case "P2PointShapes":
		GongUnmarshallSliceOfPointers(&instance.P2PointShapes, valueExpr, identifierMap)
	case "P1CurveShapes":
		GongUnmarshallSliceOfPointers(&instance.P1CurveShapes, valueExpr, identifierMap)
	case "P2CurveShapes":
		GongUnmarshallSliceOfPointers(&instance.P2CurveShapes, valueExpr, identifierMap)
	case "P1P2PairLineShapes":
		GongUnmarshallSliceOfPointers(&instance.P1P2PairLineShapes, valueExpr, identifierMap)
	}
	return nil
}

type PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeUnmarshaller struct{}

func (u *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type PartiallyGrowthCurve2DTrajectoryP1PointShapeUnmarshaller struct{}

func (u *PartiallyGrowthCurve2DTrajectoryP1PointShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyGrowthCurve2DTrajectoryP1PointShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyGrowthCurve2DTrajectoryP1PointShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyGrowthCurve2DTrajectoryP1PointShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type PartiallyGrowthCurve2DTrajectoryP2CurveShapeUnmarshaller struct{}

func (u *PartiallyGrowthCurve2DTrajectoryP2CurveShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyGrowthCurve2DTrajectoryP2CurveShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyGrowthCurve2DTrajectoryP2CurveShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyGrowthCurve2DTrajectoryP2CurveShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type PartiallyGrowthCurve2DTrajectoryP2PointShapeUnmarshaller struct{}

func (u *PartiallyGrowthCurve2DTrajectoryP2PointShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyGrowthCurve2DTrajectoryP2PointShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyGrowthCurve2DTrajectoryP2PointShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyGrowthCurve2DTrajectoryP2PointShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type PartiallyGrowthCurve2DTrajectoryShapeUnmarshaller struct{}

func (u *PartiallyGrowthCurve2DTrajectoryShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyGrowthCurve2DTrajectoryShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyGrowthCurve2DTrajectoryShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyGrowthCurve2DTrajectoryShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type PartiallyRotatedSeatBottomCurveShapeUnmarshaller struct{}

func (u *PartiallyRotatedSeatBottomCurveShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyRotatedSeatBottomCurveShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyRotatedSeatBottomCurveShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyRotatedSeatBottomCurveShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type PartiallyRotatedSeatTopCurveShapeUnmarshaller struct{}

func (u *PartiallyRotatedSeatTopCurveShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyRotatedSeatTopCurveShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyRotatedSeatTopCurveShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyRotatedSeatTopCurveShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type PartiallyRotatedTorusShapeUnmarshaller struct{}

func (u *PartiallyRotatedTorusShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PartiallyRotatedTorusShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PartiallyRotatedTorusShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PartiallyRotatedTorusShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type PerpendicularVectorUnmarshaller struct{}

func (u *PerpendicularVectorUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PerpendicularVector), stage, identifier, instanceName, preserveOrder)
}

func (u *PerpendicularVectorUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PerpendicularVector)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type PerpendicularVectorGridUnmarshaller struct{}

func (u *PerpendicularVectorGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PerpendicularVectorGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *PerpendicularVectorGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PerpendicularVectorGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PerpendicularVectors":
		GongUnmarshallSliceOfPointers(&instance.PerpendicularVectors, valueExpr, identifierMap)
	}
	return nil
}

type PerpendicularVectorGridHalfwayUnmarshaller struct{}

func (u *PerpendicularVectorGridHalfwayUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PerpendicularVectorGridHalfway), stage, identifier, instanceName, preserveOrder)
}

func (u *PerpendicularVectorGridHalfwayUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PerpendicularVectorGridHalfway)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "PerpendicularVectorHalfways":
		GongUnmarshallSliceOfPointers(&instance.PerpendicularVectorHalfways, valueExpr, identifierMap)
	}
	return nil
}

type PerpendicularVectorHalfwayUnmarshaller struct{}

func (u *PerpendicularVectorHalfwayUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PerpendicularVectorHalfway), stage, identifier, instanceName, preserveOrder)
}

func (u *PerpendicularVectorHalfwayUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PerpendicularVectorHalfway)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type Plant2DDiagramUnmarshaller struct{}

func (u *Plant2DDiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Plant2DDiagram), stage, identifier, instanceName, preserveOrder)
}

func (u *Plant2DDiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Plant2DDiagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "OriginX":
		instance.OriginX = GongExtractFloat(valueExpr)
	case "OriginY":
		instance.OriginY = GongExtractFloat(valueExpr)
	case "Zoom":
		instance.Zoom = GongExtractFloat(valueExpr)
	case "IsRhombusNodesExpanded":
		instance.IsRhombusNodesExpanded = GongExtractBool(valueExpr)
	case "IsArcNodesExpanded":
		instance.IsArcNodesExpanded = GongExtractBool(valueExpr)
	case "IsHiddenAxesShape":
		instance.IsHiddenAxesShape = GongExtractBool(valueExpr)
	case "IsHiddenReferenceRhombus":
		instance.IsHiddenReferenceRhombus = GongExtractBool(valueExpr)
	case "IsHiddenPlantCircumferenceShape":
		instance.IsHiddenPlantCircumferenceShape = GongExtractBool(valueExpr)
	case "IsHiddenGridPathShape":
		instance.IsHiddenGridPathShape = GongExtractBool(valueExpr)
	case "IsHiddenRhombusGridShape":
		instance.IsHiddenRhombusGridShape = GongExtractBool(valueExpr)
	case "IsHiddenExplanationTextShape":
		instance.IsHiddenExplanationTextShape = GongExtractBool(valueExpr)
	case "IsHiddenRotatedReferenceRhombus":
		instance.IsHiddenRotatedReferenceRhombus = GongExtractBool(valueExpr)
	case "IsHiddenRotatedPlantCircumferenceShape":
		instance.IsHiddenRotatedPlantCircumferenceShape = GongExtractBool(valueExpr)
	case "IsHiddenRotatedGridPathShape":
		instance.IsHiddenRotatedGridPathShape = GongExtractBool(valueExpr)
	case "IsHiddenRotatedRhombusGridShape":
		instance.IsHiddenRotatedRhombusGridShape = GongExtractBool(valueExpr)
	case "IsHiddenGrowthPathRhombusGridShape":
		instance.IsHiddenGrowthPathRhombusGridShape = GongExtractBool(valueExpr)
	case "IsHiddenGrowthVectorShape":
		instance.IsHiddenGrowthVectorShape = GongExtractBool(valueExpr)
	case "IsHiddenPerpendicularVectorGrid":
		instance.IsHiddenPerpendicularVectorGrid = GongExtractBool(valueExpr)
	case "IsHiddenBaseVectorShapeGrid":
		instance.IsHiddenBaseVectorShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenArcNormalVectorShapeGrid":
		instance.IsHiddenArcNormalVectorShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenStartArcShapeGrid":
		instance.IsHiddenStartArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenMidArcVectorShapeGrid":
		instance.IsHiddenMidArcVectorShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenEndArcShapeGrid":
		instance.IsHiddenEndArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenGrowthCurve2D":
		instance.IsHiddenGrowthCurve2D = GongExtractBool(valueExpr)
	case "IsHiddenStackOfGrowthCurve2DByGrowthVector":
		instance.IsHiddenStackOfGrowthCurve2DByGrowthVector = GongExtractBool(valueExpr)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type Plant3DDiagramUnmarshaller struct{}

func (u *Plant3DDiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Plant3DDiagram), stage, identifier, instanceName, preserveOrder)
}

func (u *Plant3DDiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Plant3DDiagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsHiddenStemCylinder3DShape":
		instance.IsHiddenStemCylinder3DShape = GongExtractBool(valueExpr)
	case "StemCylinder3DShape":
		GongUnmarshallPointer(&instance.StemCylinder3DShape, valueExpr, identifierMap)
	case "IsHiddenParastichyNCurves3DShape":
		instance.IsHiddenParastichyNCurves3DShape = GongExtractBool(valueExpr)
	case "ParastichyNCurves3DShape":
		GongUnmarshallPointer(&instance.ParastichyNCurves3DShape, valueExpr, identifierMap)
	case "IsHiddenParastichyMCurves3DShape":
		instance.IsHiddenParastichyMCurves3DShape = GongExtractBool(valueExpr)
	case "ParastichyMCurves3DShape":
		GongUnmarshallPointer(&instance.ParastichyMCurves3DShape, valueExpr, identifierMap)
	case "IsHiddenCutLine3DShape":
		instance.IsHiddenCutLine3DShape = GongExtractBool(valueExpr)
	case "CutLine3DShape":
		GongUnmarshallPointer(&instance.CutLine3DShape, valueExpr, identifierMap)
	case "IsHiddenCircumference3DShape":
		instance.IsHiddenCircumference3DShape = GongExtractBool(valueExpr)
	case "Circumference3DShape":
		GongUnmarshallPointer(&instance.Circumference3DShape, valueExpr, identifierMap)
	case "IsHiddenTiledFloor3DShape":
		instance.IsHiddenTiledFloor3DShape = GongExtractBool(valueExpr)
	case "TiledFloor3DShape":
		GongUnmarshallPointer(&instance.TiledFloor3DShape, valueExpr, identifierMap)
	case "IsHiddenLeaves3DShape":
		instance.IsHiddenLeaves3DShape = GongExtractBool(valueExpr)
	case "Leaves3DShape":
		GongUnmarshallPointer(&instance.Leaves3DShape, valueExpr, identifierMap)
	case "Rendered3DShape":
		GongUnmarshallPointer(&instance.Rendered3DShape, valueExpr, identifierMap)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type PlantAbstractUnmarshaller struct{}

func (u *PlantAbstractUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PlantAbstract), stage, identifier, instanceName, preserveOrder)
}

func (u *PlantAbstractUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PlantAbstract)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "N":
		instance.N = GongExtractInt(valueExpr)
	case "M":
		instance.M = GongExtractInt(valueExpr)
	case "StackHeight":
		instance.StackHeight = GongExtractInt(valueExpr)
	case "RhombusInsideAngle":
		instance.RhombusInsideAngle = GongExtractFloat(valueExpr)
	case "RhombusSideLength":
		instance.RhombusSideLength = GongExtractFloat(valueExpr)
	case "PlantType":
		GongUnmarshallEnum(&instance.PlantType, valueExpr)
	case "TubeVaseAbstract":
		GongUnmarshallPointer(&instance.TubeVaseAbstract, valueExpr, identifierMap)
	case "CurrentView":
		GongUnmarshallEnum(&instance.CurrentView, valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	case "IsSelected":
		instance.IsSelected = GongExtractBool(valueExpr)
	case "IsPlant2DDiagramsNodeExpanded":
		instance.IsPlant2DDiagramsNodeExpanded = GongExtractBool(valueExpr)
	case "Plant2DDiagrams":
		GongUnmarshallSliceOfPointers(&instance.Plant2DDiagrams, valueExpr, identifierMap)
	case "IsPlant3DDiagramsNodeExpanded":
		instance.IsPlant3DDiagramsNodeExpanded = GongExtractBool(valueExpr)
	case "Plant3DDiagrams":
		GongUnmarshallSliceOfPointers(&instance.Plant3DDiagrams, valueExpr, identifierMap)
	case "IsVase2DDiagramsNodeExpanded":
		instance.IsVase2DDiagramsNodeExpanded = GongExtractBool(valueExpr)
	case "Vase2DDiagrams":
		GongUnmarshallSliceOfPointers(&instance.Vase2DDiagrams, valueExpr, identifierMap)
	case "IsTubeVase3DDiagramsNodeExpanded":
		instance.IsTubeVase3DDiagramsNodeExpanded = GongExtractBool(valueExpr)
	case "TubeVase3DDiagrams":
		GongUnmarshallSliceOfPointers(&instance.TubeVase3DDiagrams, valueExpr, identifierMap)
	case "IsStool2DDiagramsNodeExpanded":
		instance.IsStool2DDiagramsNodeExpanded = GongExtractBool(valueExpr)
	case "Stool2DDiagrams":
		GongUnmarshallSliceOfPointers(&instance.Stool2DDiagrams, valueExpr, identifierMap)
	case "IsStool3DDiagramsNodeExpanded":
		instance.IsStool3DDiagramsNodeExpanded = GongExtractBool(valueExpr)
	case "Stool3DDiagrams":
		GongUnmarshallSliceOfPointers(&instance.Stool3DDiagrams, valueExpr, identifierMap)
	case "IsClock2DDiagramsNodeExpanded":
		instance.IsClock2DDiagramsNodeExpanded = GongExtractBool(valueExpr)
	case "Clock2DDiagrams":
		GongUnmarshallSliceOfPointers(&instance.Clock2DDiagrams, valueExpr, identifierMap)
	case "IsClock3DDiagramsNodeExpanded":
		instance.IsClock3DDiagramsNodeExpanded = GongExtractBool(valueExpr)
	case "Clock3DDiagrams":
		GongUnmarshallSliceOfPointers(&instance.Clock3DDiagrams, valueExpr, identifierMap)
	case "AxesShape":
		GongUnmarshallPointer(&instance.AxesShape, valueExpr, identifierMap)
	case "RhombusStuff":
		GongUnmarshallPointer(&instance.RhombusStuff, valueExpr, identifierMap)
	case "GrowthVectorShape":
		GongUnmarshallPointer(&instance.GrowthVectorShape, valueExpr, identifierMap)
	case "PerpendicularVectorGrid":
		GongUnmarshallPointer(&instance.PerpendicularVectorGrid, valueExpr, identifierMap)
	case "BaseVectorShapeGrid":
		GongUnmarshallPointer(&instance.BaseVectorShapeGrid, valueExpr, identifierMap)
	case "ArcNormalVectorShapeGrid":
		GongUnmarshallPointer(&instance.ArcNormalVectorShapeGrid, valueExpr, identifierMap)
	case "StartArcShapeGrid":
		GongUnmarshallPointer(&instance.StartArcShapeGrid, valueExpr, identifierMap)
	case "MidArcVectorShapeGrid":
		GongUnmarshallPointer(&instance.MidArcVectorShapeGrid, valueExpr, identifierMap)
	case "EndArcShapeGrid":
		GongUnmarshallPointer(&instance.EndArcShapeGrid, valueExpr, identifierMap)
	case "GrowthCurve2D":
		GongUnmarshallPointer(&instance.GrowthCurve2D, valueExpr, identifierMap)
	case "StackOfGrowthCurve2DByGrowthVector":
		GongUnmarshallPointer(&instance.StackOfGrowthCurve2DByGrowthVector, valueExpr, identifierMap)
	}
	return nil
}

type PlantCircumferenceShapeUnmarshaller struct{}

func (u *PlantCircumferenceShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PlantCircumferenceShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PlantCircumferenceShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PlantCircumferenceShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "AngleDegree":
		instance.AngleDegree = GongExtractFloat(valueExpr)
	case "Length":
		instance.Length = GongExtractFloat(valueExpr)
	}
	return nil
}

type PointsAndLines3DShapeUnmarshaller struct{}

func (u *PointsAndLines3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PointsAndLines3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PointsAndLines3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PointsAndLines3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type PxShapeUnmarshaller struct{}

func (u *PxShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(PxShape), stage, identifier, instanceName, preserveOrder)
}

func (u *PxShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*PxShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type Rendered3DShapeUnmarshaller struct{}

func (u *Rendered3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Rendered3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *Rendered3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Rendered3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ViewX":
		instance.ViewX = GongExtractFloat(valueExpr)
	case "ViewY":
		instance.ViewY = GongExtractFloat(valueExpr)
	case "ViewZ":
		instance.ViewZ = GongExtractFloat(valueExpr)
	case "TargetX":
		instance.TargetX = GongExtractFloat(valueExpr)
	case "TargetY":
		instance.TargetY = GongExtractFloat(valueExpr)
	case "TargetZ":
		instance.TargetZ = GongExtractFloat(valueExpr)
	case "Fov":
		instance.Fov = GongExtractFloat(valueExpr)
	}
	return nil
}

type RhombusShapeUnmarshaller struct{}

func (u *RhombusShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RhombusShape), stage, identifier, instanceName, preserveOrder)
}

func (u *RhombusShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RhombusShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type RhombusStuffUnmarshaller struct{}

func (u *RhombusStuffUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RhombusStuff), stage, identifier, instanceName, preserveOrder)
}

func (u *RhombusStuffUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RhombusStuff)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ReferenceRhombus":
		GongUnmarshallPointer(&instance.ReferenceRhombus, valueExpr, identifierMap)
	case "PlantCircumferenceShape":
		GongUnmarshallPointer(&instance.PlantCircumferenceShape, valueExpr, identifierMap)
	case "GridPathShape":
		GongUnmarshallPointer(&instance.GridPathShape, valueExpr, identifierMap)
	case "InitialRhombusGridShape":
		GongUnmarshallPointer(&instance.InitialRhombusGridShape, valueExpr, identifierMap)
	case "ExplanationTextShape":
		GongUnmarshallPointer(&instance.ExplanationTextShape, valueExpr, identifierMap)
	case "RotatedReferenceRhombus":
		GongUnmarshallPointer(&instance.RotatedReferenceRhombus, valueExpr, identifierMap)
	case "RotatedPlantCircumferenceShape":
		GongUnmarshallPointer(&instance.RotatedPlantCircumferenceShape, valueExpr, identifierMap)
	case "RotatedGridPathShape":
		GongUnmarshallPointer(&instance.RotatedGridPathShape, valueExpr, identifierMap)
	case "RotatedRhombusGridShape2":
		GongUnmarshallPointer(&instance.RotatedRhombusGridShape2, valueExpr, identifierMap)
	case "GrowthCurveRhombusGridShape":
		GongUnmarshallPointer(&instance.GrowthCurveRhombusGridShape, valueExpr, identifierMap)
	}
	return nil
}

type RotatedRhombusGridShapeUnmarshaller struct{}

func (u *RotatedRhombusGridShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RotatedRhombusGridShape), stage, identifier, instanceName, preserveOrder)
}

func (u *RotatedRhombusGridShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RotatedRhombusGridShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "RotatedRhombusShapes":
		GongUnmarshallSliceOfPointers(&instance.RotatedRhombusShapes, valueExpr, identifierMap)
	}
	return nil
}

type RotatedRhombusShapeUnmarshaller struct{}

func (u *RotatedRhombusShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RotatedRhombusShape), stage, identifier, instanceName, preserveOrder)
}

func (u *RotatedRhombusShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RotatedRhombusShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "X":
		instance.X = GongExtractFloat(valueExpr)
	case "Y":
		instance.Y = GongExtractFloat(valueExpr)
	}
	return nil
}

type RotatedSampledPoints3DShapeUnmarshaller struct{}

func (u *RotatedSampledPoints3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RotatedSampledPoints3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *RotatedSampledPoints3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RotatedSampledPoints3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type RotatedSeatAndLegs3DShapeUnmarshaller struct{}

func (u *RotatedSeatAndLegs3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(RotatedSeatAndLegs3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *RotatedSeatAndLegs3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*RotatedSeatAndLegs3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type SampledPoints3DShapeUnmarshaller struct{}

func (u *SampledPoints3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SampledPoints3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *SampledPoints3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SampledPoints3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type Seat3DShapeUnmarshaller struct{}

func (u *Seat3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Seat3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *Seat3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Seat3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type SeatAndLegs3DShapeUnmarshaller struct{}

func (u *SeatAndLegs3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SeatAndLegs3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *SeatAndLegs3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SeatAndLegs3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type SeatBottomCurveShapeUnmarshaller struct{}

func (u *SeatBottomCurveShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SeatBottomCurveShape), stage, identifier, instanceName, preserveOrder)
}

func (u *SeatBottomCurveShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SeatBottomCurveShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type SeatTopCurveShapeUnmarshaller struct{}

func (u *SeatTopCurveShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(SeatTopCurveShape), stage, identifier, instanceName, preserveOrder)
}

func (u *SeatTopCurveShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*SeatTopCurveShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type ShiftedBottomTopStartArcShapeUnmarshaller struct{}

func (u *ShiftedBottomTopStartArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedBottomTopStartArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedBottomTopStartArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedBottomTopStartArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type ShiftedBottomTopStartArcShapeGridUnmarshaller struct{}

func (u *ShiftedBottomTopStartArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedBottomTopStartArcShapeGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedBottomTopStartArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedBottomTopStartArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShiftedBottomTopStartArcShapes":
		GongUnmarshallSliceOfPointers(&instance.ShiftedBottomTopStartArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type ShiftedLeftGrowthCurve2DRibbonUnmarshaller struct{}

func (u *ShiftedLeftGrowthCurve2DRibbonUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedLeftGrowthCurve2DRibbon), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedLeftGrowthCurve2DRibbonUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftGrowthCurve2DRibbon)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShiftedLeftGrowthCurve2DRibbonStartShapes":
		GongUnmarshallSliceOfPointers(&instance.ShiftedLeftGrowthCurve2DRibbonStartShapes, valueExpr, identifierMap)
	case "ShiftedLeftGrowthCurve2DRibbonEndShapes":
		GongUnmarshallSliceOfPointers(&instance.ShiftedLeftGrowthCurve2DRibbonEndShapes, valueExpr, identifierMap)
	}
	return nil
}

type ShiftedLeftGrowthCurve2DRibbonEndShapeUnmarshaller struct{}

func (u *ShiftedLeftGrowthCurve2DRibbonEndShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedLeftGrowthCurve2DRibbonEndShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedLeftGrowthCurve2DRibbonEndShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftGrowthCurve2DRibbonEndShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type ShiftedLeftGrowthCurve2DRibbonStartShapeUnmarshaller struct{}

func (u *ShiftedLeftGrowthCurve2DRibbonStartShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedLeftGrowthCurve2DRibbonStartShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedLeftGrowthCurve2DRibbonStartShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftGrowthCurve2DRibbonStartShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type ShiftedLeftPartiallyGrowthCurve2DRibbonUnmarshaller struct{}

func (u *ShiftedLeftPartiallyGrowthCurve2DRibbonUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedLeftPartiallyGrowthCurve2DRibbon), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedLeftPartiallyGrowthCurve2DRibbonUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftPartiallyGrowthCurve2DRibbon)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes":
		GongUnmarshallSliceOfPointers(&instance.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes, valueExpr, identifierMap)
	case "ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes":
		GongUnmarshallSliceOfPointers(&instance.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes, valueExpr, identifierMap)
	}
	return nil
}

type ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeUnmarshaller struct{}

func (u *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeUnmarshaller struct{}

func (u *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type ShiftedLeftStackGrowthCurveEndArcShapeUnmarshaller struct{}

func (u *ShiftedLeftStackGrowthCurveEndArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedLeftStackGrowthCurveEndArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedLeftStackGrowthCurveEndArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftStackGrowthCurveEndArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type ShiftedLeftStackGrowthCurveStartArcShapeUnmarshaller struct{}

func (u *ShiftedLeftStackGrowthCurveStartArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedLeftStackGrowthCurveStartArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedLeftStackGrowthCurveStartArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftStackGrowthCurveStartArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type ShiftedLeftStackNormalVectorUnmarshaller struct{}

func (u *ShiftedLeftStackNormalVectorUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedLeftStackNormalVector), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedLeftStackNormalVectorUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftStackNormalVector)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type ShiftedLeftStackOfGrowthCurveUnmarshaller struct{}

func (u *ShiftedLeftStackOfGrowthCurveUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedLeftStackOfGrowthCurve), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedLeftStackOfGrowthCurveUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftStackOfGrowthCurve)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShiftedLeftStackGrowthCurveStartArcShapes":
		GongUnmarshallSliceOfPointers(&instance.ShiftedLeftStackGrowthCurveStartArcShapes, valueExpr, identifierMap)
	case "ShiftedLeftStackGrowthCurveEndArcShapes":
		GongUnmarshallSliceOfPointers(&instance.ShiftedLeftStackGrowthCurveEndArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type ShiftedLeftStackOfNormalVectorUnmarshaller struct{}

func (u *ShiftedLeftStackOfNormalVectorUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedLeftStackOfNormalVector), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedLeftStackOfNormalVectorUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedLeftStackOfNormalVector)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShiftedLeftStackNormalVectors":
		GongUnmarshallSliceOfPointers(&instance.ShiftedLeftStackNormalVectors, valueExpr, identifierMap)
	}
	return nil
}

type ShiftedRightGrowthCurve2DRibbonUnmarshaller struct{}

func (u *ShiftedRightGrowthCurve2DRibbonUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedRightGrowthCurve2DRibbon), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedRightGrowthCurve2DRibbonUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedRightGrowthCurve2DRibbon)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "ShiftedRightGrowthCurve2DRibbonStartShapes":
		GongUnmarshallSliceOfPointers(&instance.ShiftedRightGrowthCurve2DRibbonStartShapes, valueExpr, identifierMap)
	case "ShiftedRightGrowthCurve2DRibbonEndShapes":
		GongUnmarshallSliceOfPointers(&instance.ShiftedRightGrowthCurve2DRibbonEndShapes, valueExpr, identifierMap)
	}
	return nil
}

type ShiftedRightGrowthCurve2DRibbonEndShapeUnmarshaller struct{}

func (u *ShiftedRightGrowthCurve2DRibbonEndShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedRightGrowthCurve2DRibbonEndShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedRightGrowthCurve2DRibbonEndShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedRightGrowthCurve2DRibbonEndShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type ShiftedRightGrowthCurve2DRibbonStartShapeUnmarshaller struct{}

func (u *ShiftedRightGrowthCurve2DRibbonStartShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ShiftedRightGrowthCurve2DRibbonStartShape), stage, identifier, instanceName, preserveOrder)
}

func (u *ShiftedRightGrowthCurve2DRibbonStartShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ShiftedRightGrowthCurve2DRibbonStartShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type StackGrowthCurve2DEndHalfwayArcShapeUnmarshaller struct{}

func (u *StackGrowthCurve2DEndHalfwayArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackGrowthCurve2DEndHalfwayArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StackGrowthCurve2DEndHalfwayArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackGrowthCurve2DEndHalfwayArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type StackGrowthCurve2DRibbonEndShapeUnmarshaller struct{}

func (u *StackGrowthCurve2DRibbonEndShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackGrowthCurve2DRibbonEndShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StackGrowthCurve2DRibbonEndShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackGrowthCurve2DRibbonEndShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type StackGrowthCurve2DRibbonStartShapeUnmarshaller struct{}

func (u *StackGrowthCurve2DRibbonStartShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackGrowthCurve2DRibbonStartShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StackGrowthCurve2DRibbonStartShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackGrowthCurve2DRibbonStartShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type StackGrowthCurve2DStartHalfwayArcShapeUnmarshaller struct{}

func (u *StackGrowthCurve2DStartHalfwayArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackGrowthCurve2DStartHalfwayArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StackGrowthCurve2DStartHalfwayArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackGrowthCurve2DStartHalfwayArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type StackOfGrowthCurve2DUnmarshaller struct{}

func (u *StackOfGrowthCurve2DUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackOfGrowthCurve2D), stage, identifier, instanceName, preserveOrder)
}

func (u *StackOfGrowthCurve2DUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackOfGrowthCurve2D)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackGrowthCurve2DStartHalfwayArcShapes":
		GongUnmarshallSliceOfPointers(&instance.StackGrowthCurve2DStartHalfwayArcShapes, valueExpr, identifierMap)
	case "StackGrowthCurve2DEndHalfwayArcShapes":
		GongUnmarshallSliceOfPointers(&instance.StackGrowthCurve2DEndHalfwayArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type StackOfGrowthCurve2DByGrowthVectorUnmarshaller struct{}

func (u *StackOfGrowthCurve2DByGrowthVectorUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackOfGrowthCurve2DByGrowthVector), stage, identifier, instanceName, preserveOrder)
}

func (u *StackOfGrowthCurve2DByGrowthVectorUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackOfGrowthCurve2DByGrowthVector)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type StackOfGrowthCurve2DRibbonUnmarshaller struct{}

func (u *StackOfGrowthCurve2DRibbonUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackOfGrowthCurve2DRibbon), stage, identifier, instanceName, preserveOrder)
}

func (u *StackOfGrowthCurve2DRibbonUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackOfGrowthCurve2DRibbon)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackGrowthCurve2DRibbonStartShapes":
		GongUnmarshallSliceOfPointers(&instance.StackGrowthCurve2DRibbonStartShapes, valueExpr, identifierMap)
	case "StackGrowthCurve2DRibbonEndShapes":
		GongUnmarshallSliceOfPointers(&instance.StackGrowthCurve2DRibbonEndShapes, valueExpr, identifierMap)
	}
	return nil
}

type StackOfPartiallyRotatedTorusShapeUnmarshaller struct{}

func (u *StackOfPartiallyRotatedTorusShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackOfPartiallyRotatedTorusShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StackOfPartiallyRotatedTorusShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackOfPartiallyRotatedTorusShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type StackOfRotatedCarvedOutVaseTrapezeRingsShapeUnmarshaller struct{}

func (u *StackOfRotatedCarvedOutVaseTrapezeRingsShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackOfRotatedCarvedOutVaseTrapezeRingsShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StackOfRotatedCarvedOutVaseTrapezeRingsShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackOfRotatedCarvedOutVaseTrapezeRingsShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type StackOfRotatedGrowthCurve2DUnmarshaller struct{}

func (u *StackOfRotatedGrowthCurve2DUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackOfRotatedGrowthCurve2D), stage, identifier, instanceName, preserveOrder)
}

func (u *StackOfRotatedGrowthCurve2DUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackOfRotatedGrowthCurve2D)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackRotatedGrowthCurve2DStartArcShapes":
		GongUnmarshallSliceOfPointers(&instance.StackRotatedGrowthCurve2DStartArcShapes, valueExpr, identifierMap)
	case "StackRotatedGrowthCurve2DEndArcShapes":
		GongUnmarshallSliceOfPointers(&instance.StackRotatedGrowthCurve2DEndArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type StackOfRotatedGrowthCurve2DRibbonUnmarshaller struct{}

func (u *StackOfRotatedGrowthCurve2DRibbonUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackOfRotatedGrowthCurve2DRibbon), stage, identifier, instanceName, preserveOrder)
}

func (u *StackOfRotatedGrowthCurve2DRibbonUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackOfRotatedGrowthCurve2DRibbon)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StackRotatedGrowthCurve2DRibbonStartShapes":
		GongUnmarshallSliceOfPointers(&instance.StackRotatedGrowthCurve2DRibbonStartShapes, valueExpr, identifierMap)
	case "StackRotatedGrowthCurve2DRibbonEndShapes":
		GongUnmarshallSliceOfPointers(&instance.StackRotatedGrowthCurve2DRibbonEndShapes, valueExpr, identifierMap)
	}
	return nil
}

type StackRotatedGrowthCurve2DEndArcShapeUnmarshaller struct{}

func (u *StackRotatedGrowthCurve2DEndArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackRotatedGrowthCurve2DEndArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StackRotatedGrowthCurve2DEndArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackRotatedGrowthCurve2DEndArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type StackRotatedGrowthCurve2DRibbonEndShapeUnmarshaller struct{}

func (u *StackRotatedGrowthCurve2DRibbonEndShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackRotatedGrowthCurve2DRibbonEndShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StackRotatedGrowthCurve2DRibbonEndShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackRotatedGrowthCurve2DRibbonEndShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type StackRotatedGrowthCurve2DRibbonStartShapeUnmarshaller struct{}

func (u *StackRotatedGrowthCurve2DRibbonStartShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackRotatedGrowthCurve2DRibbonStartShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StackRotatedGrowthCurve2DRibbonStartShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackRotatedGrowthCurve2DRibbonStartShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "BottomStartX":
		instance.BottomStartX = GongExtractFloat(valueExpr)
	case "BottomStartY":
		instance.BottomStartY = GongExtractFloat(valueExpr)
	case "BottomEndX":
		instance.BottomEndX = GongExtractFloat(valueExpr)
	case "BottomEndY":
		instance.BottomEndY = GongExtractFloat(valueExpr)
	case "BottomRadiusX":
		instance.BottomRadiusX = GongExtractFloat(valueExpr)
	case "BottomRadiusY":
		instance.BottomRadiusY = GongExtractFloat(valueExpr)
	case "BottomXAxisRotation":
		instance.BottomXAxisRotation = GongExtractFloat(valueExpr)
	case "BottomLargeArcFlag":
		instance.BottomLargeArcFlag = GongExtractBool(valueExpr)
	case "BottomSweepFlag":
		instance.BottomSweepFlag = GongExtractBool(valueExpr)
	case "TopStartX":
		instance.TopStartX = GongExtractFloat(valueExpr)
	case "TopStartY":
		instance.TopStartY = GongExtractFloat(valueExpr)
	case "TopEndX":
		instance.TopEndX = GongExtractFloat(valueExpr)
	case "TopEndY":
		instance.TopEndY = GongExtractFloat(valueExpr)
	case "TopRadiusX":
		instance.TopRadiusX = GongExtractFloat(valueExpr)
	case "TopRadiusY":
		instance.TopRadiusY = GongExtractFloat(valueExpr)
	case "TopXAxisRotation":
		instance.TopXAxisRotation = GongExtractFloat(valueExpr)
	case "TopLargeArcFlag":
		instance.TopLargeArcFlag = GongExtractBool(valueExpr)
	case "TopSweepFlag":
		instance.TopSweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type StackRotatedGrowthCurve2DStartArcShapeUnmarshaller struct{}

func (u *StackRotatedGrowthCurve2DStartArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StackRotatedGrowthCurve2DStartArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StackRotatedGrowthCurve2DStartArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StackRotatedGrowthCurve2DStartArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type StartArcShapeUnmarshaller struct{}

func (u *StartArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StartArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StartArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StartArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type StartArcShapeGridUnmarshaller struct{}

func (u *StartArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StartArcShapeGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *StartArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StartArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartArcShapes":
		GongUnmarshallSliceOfPointers(&instance.StartArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type StartHalfwayArcShapeUnmarshaller struct{}

func (u *StartHalfwayArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StartHalfwayArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StartHalfwayArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StartHalfwayArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type StartHalfwayArcShapeGridUnmarshaller struct{}

func (u *StartHalfwayArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StartHalfwayArcShapeGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *StartHalfwayArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StartHalfwayArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartHalfwayArcShapes":
		GongUnmarshallSliceOfPointers(&instance.StartHalfwayArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type StemCylinder3DShapeUnmarshaller struct{}

func (u *StemCylinder3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(StemCylinder3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *StemCylinder3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*StemCylinder3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Transparency":
		instance.Transparency = GongExtractFloat(valueExpr)
	}
	return nil
}

type Stool2DDiagramUnmarshaller struct{}

func (u *Stool2DDiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Stool2DDiagram), stage, identifier, instanceName, preserveOrder)
}

func (u *Stool2DDiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Stool2DDiagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Zoom":
		instance.Zoom = GongExtractFloat(valueExpr)
	case "IsHiddenAxesShape":
		instance.IsHiddenAxesShape = GongExtractBool(valueExpr)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type Stool3DDiagramUnmarshaller struct{}

func (u *Stool3DDiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Stool3DDiagram), stage, identifier, instanceName, preserveOrder)
}

func (u *Stool3DDiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Stool3DDiagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsHiddenSeatTopCurveShape":
		instance.IsHiddenSeatTopCurveShape = GongExtractBool(valueExpr)
	case "SeatTopCurveShape":
		GongUnmarshallPointer(&instance.SeatTopCurveShape, valueExpr, identifierMap)
	case "IsHiddenRotatedSeatTopCurveShape":
		instance.IsHiddenRotatedSeatTopCurveShape = GongExtractBool(valueExpr)
	case "RotatedSeatTopCurveShape":
		GongUnmarshallPointer(&instance.RotatedSeatTopCurveShape, valueExpr, identifierMap)
	case "IsHiddenSeatBottomCurveShape":
		instance.IsHiddenSeatBottomCurveShape = GongExtractBool(valueExpr)
	case "SeatBottomCurveShape":
		GongUnmarshallPointer(&instance.SeatBottomCurveShape, valueExpr, identifierMap)
	case "IsHiddenRotatedSeatBottomCurveShape":
		instance.IsHiddenRotatedSeatBottomCurveShape = GongExtractBool(valueExpr)
	case "RotatedSeatBottomCurveShape":
		GongUnmarshallPointer(&instance.RotatedSeatBottomCurveShape, valueExpr, identifierMap)
	case "IsHiddenTorus3DShape":
		instance.IsHiddenTorus3DShape = GongExtractBool(valueExpr)
	case "Torus3DShape":
		GongUnmarshallPointer(&instance.Torus3DShape, valueExpr, identifierMap)
	case "IsHiddenRotatedTorusShape":
		instance.IsHiddenRotatedTorusShape = GongExtractBool(valueExpr)
	case "RotatedTorusShape":
		GongUnmarshallPointer(&instance.RotatedTorusShape, valueExpr, identifierMap)
	case "IsHiddenSampledPoints3DShape":
		instance.IsHiddenSampledPoints3DShape = GongExtractBool(valueExpr)
	case "SampledPoints3DShape":
		GongUnmarshallPointer(&instance.SampledPoints3DShape, valueExpr, identifierMap)
	case "IsHiddenRotatedSampledPoints3DShape":
		instance.IsHiddenRotatedSampledPoints3DShape = GongExtractBool(valueExpr)
	case "RotatedSampledPoints3DShape":
		GongUnmarshallPointer(&instance.RotatedSampledPoints3DShape, valueExpr, identifierMap)
	case "IsHiddenEyeSampledPoints3DShape":
		instance.IsHiddenEyeSampledPoints3DShape = GongExtractBool(valueExpr)
	case "EyeSampledPoints3DShape":
		GongUnmarshallPointer(&instance.EyeSampledPoints3DShape, valueExpr, identifierMap)
	case "IsHiddenEyeCornersSampledPoints3DShape":
		instance.IsHiddenEyeCornersSampledPoints3DShape = GongExtractBool(valueExpr)
	case "EyeCornersSampledPoints3DShape":
		GongUnmarshallPointer(&instance.EyeCornersSampledPoints3DShape, valueExpr, identifierMap)
	case "IsHiddenEye3DShape":
		instance.IsHiddenEye3DShape = GongExtractBool(valueExpr)
	case "Eye3DShape":
		GongUnmarshallPointer(&instance.Eye3DShape, valueExpr, identifierMap)
	case "IsHiddenEyeSeatBottomCurveShape":
		instance.IsHiddenEyeSeatBottomCurveShape = GongExtractBool(valueExpr)
	case "EyeSeatBottomCurveShape":
		GongUnmarshallPointer(&instance.EyeSeatBottomCurveShape, valueExpr, identifierMap)
	case "IsHiddenEyeStoolBottomCurveShape":
		instance.IsHiddenEyeStoolBottomCurveShape = GongExtractBool(valueExpr)
	case "EyeStoolBottomCurveShape":
		GongUnmarshallPointer(&instance.EyeStoolBottomCurveShape, valueExpr, identifierMap)
	case "IsHiddenSeat3DShape":
		instance.IsHiddenSeat3DShape = GongExtractBool(valueExpr)
	case "Seat3DShape":
		GongUnmarshallPointer(&instance.Seat3DShape, valueExpr, identifierMap)
	case "IsHiddenEyeVolume3DShape":
		instance.IsHiddenEyeVolume3DShape = GongExtractBool(valueExpr)
	case "EyeVolume3DShape":
		GongUnmarshallPointer(&instance.EyeVolume3DShape, valueExpr, identifierMap)
	case "IsHiddenSeatAndLegs3DShape":
		instance.IsHiddenSeatAndLegs3DShape = GongExtractBool(valueExpr)
	case "SeatAndLegs3DShape":
		GongUnmarshallPointer(&instance.SeatAndLegs3DShape, valueExpr, identifierMap)
	case "IsHiddenRotatedSeatAndLegs3DShape":
		instance.IsHiddenRotatedSeatAndLegs3DShape = GongExtractBool(valueExpr)
	case "RotatedSeatAndLegs3DShape":
		GongUnmarshallPointer(&instance.RotatedSeatAndLegs3DShape, valueExpr, identifierMap)
	case "IsHiddenTiledFloor3DShape":
		instance.IsHiddenTiledFloor3DShape = GongExtractBool(valueExpr)
	case "TiledFloor3DShape":
		GongUnmarshallPointer(&instance.TiledFloor3DShape, valueExpr, identifierMap)
	case "Rendered3DShape":
		GongUnmarshallPointer(&instance.Rendered3DShape, valueExpr, identifierMap)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type TiledFloor3DShapeUnmarshaller struct{}

func (u *TiledFloor3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TiledFloor3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TiledFloor3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TiledFloor3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type TopCurvePlane1ShapeUnmarshaller struct{}

func (u *TopCurvePlane1ShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopCurvePlane1Shape), stage, identifier, instanceName, preserveOrder)
}

func (u *TopCurvePlane1ShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopCurvePlane1Shape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type TopCurvePlane2ShapeUnmarshaller struct{}

func (u *TopCurvePlane2ShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopCurvePlane2Shape), stage, identifier, instanceName, preserveOrder)
}

func (u *TopCurvePlane2ShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopCurvePlane2Shape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type TopEndArcShapeUnmarshaller struct{}

func (u *TopEndArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopEndArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TopEndArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopEndArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type TopEndArcShapeGridUnmarshaller struct{}

func (u *TopEndArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopEndArcShapeGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *TopEndArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopEndArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopEndArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopEndArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type TopEndHalfwayArcShapeUnmarshaller struct{}

func (u *TopEndHalfwayArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopEndHalfwayArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TopEndHalfwayArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopEndHalfwayArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type TopEndHalfwayArcShapeGridUnmarshaller struct{}

func (u *TopEndHalfwayArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopEndHalfwayArcShapeGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *TopEndHalfwayArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopEndHalfwayArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopEndHalfwayArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopEndHalfwayArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type TopGrowthCurve2DUnmarshaller struct{}

func (u *TopGrowthCurve2DUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopGrowthCurve2D), stage, identifier, instanceName, preserveOrder)
}

func (u *TopGrowthCurve2DUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopGrowthCurve2D)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopStartHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.TopStartHalfwayArcShapeGrid, valueExpr, identifierMap)
	case "TopEndHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.TopEndHalfwayArcShapeGrid, valueExpr, identifierMap)
	}
	return nil
}

type TopMidArcVectorShapeUnmarshaller struct{}

func (u *TopMidArcVectorShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopMidArcVectorShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TopMidArcVectorShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopMidArcVectorShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	}
	return nil
}

type TopMidArcVectorShapeGridUnmarshaller struct{}

func (u *TopMidArcVectorShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopMidArcVectorShapeGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *TopMidArcVectorShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopMidArcVectorShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopMidArcVectorShapes":
		GongUnmarshallSliceOfPointers(&instance.TopMidArcVectorShapes, valueExpr, identifierMap)
	}
	return nil
}

type TopStackGrowthCurve2DEndHalfwayArcShapeUnmarshaller struct{}

func (u *TopStackGrowthCurve2DEndHalfwayArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopStackGrowthCurve2DEndHalfwayArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TopStackGrowthCurve2DEndHalfwayArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStackGrowthCurve2DEndHalfwayArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type TopStackGrowthCurve2DStartHalfwayArcShapeUnmarshaller struct{}

func (u *TopStackGrowthCurve2DStartHalfwayArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopStackGrowthCurve2DStartHalfwayArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TopStackGrowthCurve2DStartHalfwayArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStackGrowthCurve2DStartHalfwayArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type TopStackOfGrowthCurve2DUnmarshaller struct{}

func (u *TopStackOfGrowthCurve2DUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopStackOfGrowthCurve2D), stage, identifier, instanceName, preserveOrder)
}

func (u *TopStackOfGrowthCurve2DUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStackOfGrowthCurve2D)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopStackGrowthCurve2DStartHalfwayArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopStackGrowthCurve2DStartHalfwayArcShapes, valueExpr, identifierMap)
	case "TopStackGrowthCurve2DEndHalfwayArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopStackGrowthCurve2DEndHalfwayArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type TopStackOfRotatedGrowthCurve2DUnmarshaller struct{}

func (u *TopStackOfRotatedGrowthCurve2DUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopStackOfRotatedGrowthCurve2D), stage, identifier, instanceName, preserveOrder)
}

func (u *TopStackOfRotatedGrowthCurve2DUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStackOfRotatedGrowthCurve2D)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopStackOfRotatedGrowthCurve2DStartArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopStackOfRotatedGrowthCurve2DStartArcShapes, valueExpr, identifierMap)
	case "TopStackOfRotatedGrowthCurve2DEndArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopStackOfRotatedGrowthCurve2DEndArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type TopStackOfRotatedGrowthCurve2DEndArcShapeUnmarshaller struct{}

func (u *TopStackOfRotatedGrowthCurve2DEndArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopStackOfRotatedGrowthCurve2DEndArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TopStackOfRotatedGrowthCurve2DEndArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStackOfRotatedGrowthCurve2DEndArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type TopStackOfRotatedGrowthCurve2DStartArcShapeUnmarshaller struct{}

func (u *TopStackOfRotatedGrowthCurve2DStartArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopStackOfRotatedGrowthCurve2DStartArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TopStackOfRotatedGrowthCurve2DStartArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStackOfRotatedGrowthCurve2DStartArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type TopStartArcShapeUnmarshaller struct{}

func (u *TopStartArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopStartArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TopStartArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStartArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	}
	return nil
}

type TopStartArcShapeGridUnmarshaller struct{}

func (u *TopStartArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopStartArcShapeGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *TopStartArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStartArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopStartArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopStartArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type TopStartHalfwayArcShapeUnmarshaller struct{}

func (u *TopStartHalfwayArcShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopStartHalfwayArcShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TopStartHalfwayArcShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStartHalfwayArcShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "StartX":
		instance.StartX = GongExtractFloat(valueExpr)
	case "StartY":
		instance.StartY = GongExtractFloat(valueExpr)
	case "EndX":
		instance.EndX = GongExtractFloat(valueExpr)
	case "EndY":
		instance.EndY = GongExtractFloat(valueExpr)
	case "RadiusX":
		instance.RadiusX = GongExtractFloat(valueExpr)
	case "RadiusY":
		instance.RadiusY = GongExtractFloat(valueExpr)
	case "XAxisRotation":
		instance.XAxisRotation = GongExtractFloat(valueExpr)
	case "LargeArcFlag":
		instance.LargeArcFlag = GongExtractBool(valueExpr)
	case "SweepFlag":
		instance.SweepFlag = GongExtractBool(valueExpr)
	}
	return nil
}

type TopStartHalfwayArcShapeGridUnmarshaller struct{}

func (u *TopStartHalfwayArcShapeGridUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TopStartHalfwayArcShapeGrid), stage, identifier, instanceName, preserveOrder)
}

func (u *TopStartHalfwayArcShapeGridUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TopStartHalfwayArcShapeGrid)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "TopStartHalfwayArcShapes":
		GongUnmarshallSliceOfPointers(&instance.TopStartHalfwayArcShapes, valueExpr, identifierMap)
	}
	return nil
}

type Torus3DShapeUnmarshaller struct{}

func (u *Torus3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Torus3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *Torus3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Torus3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type TorusEdge3DShapeUnmarshaller struct{}

func (u *TorusEdge3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TorusEdge3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TorusEdge3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TorusEdge3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type TorusStackShapeUnmarshaller struct{}

func (u *TorusStackShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TorusStackShape), stage, identifier, instanceName, preserveOrder)
}

func (u *TorusStackShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TorusStackShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type TubeVase3DDiagramUnmarshaller struct{}

func (u *TubeVase3DDiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TubeVase3DDiagram), stage, identifier, instanceName, preserveOrder)
}

func (u *TubeVase3DDiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TubeVase3DDiagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon":
		instance.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon = GongExtractBool(valueExpr)
	case "IsHiddenTorusStackShape":
		instance.IsHiddenTorusStackShape = GongExtractBool(valueExpr)
	case "IsHiddenVerticalTorusStackShape":
		instance.IsHiddenVerticalTorusStackShape = GongExtractBool(valueExpr)
	case "IsHiddenPartiallyRotatedTorusShape":
		instance.IsHiddenPartiallyRotatedTorusShape = GongExtractBool(valueExpr)
	case "IsHiddenStackOfPartiallyRotatedTorusShape":
		instance.IsHiddenStackOfPartiallyRotatedTorusShape = GongExtractBool(valueExpr)
	case "IsHiddenPointsAndLines3DShape":
		instance.IsHiddenPointsAndLines3DShape = GongExtractBool(valueExpr)
	case "IsHiddenKeyHole3DShape":
		instance.IsHiddenKeyHole3DShape = GongExtractBool(valueExpr)
	case "IsHiddenKey3DShape":
		instance.IsHiddenKey3DShape = GongExtractBool(valueExpr)
	case "IsHiddenVolumeKey3DShape":
		instance.IsHiddenVolumeKey3DShape = GongExtractBool(valueExpr)
	case "IsHiddenTorusEdge3DShape":
		instance.IsHiddenTorusEdge3DShape = GongExtractBool(valueExpr)
	case "IsHiddenSampledPoints3DShape":
		instance.IsHiddenSampledPoints3DShape = GongExtractBool(valueExpr)
	case "IsHiddenOriginalPoints3DShape":
		instance.IsHiddenOriginalPoints3DShape = GongExtractBool(valueExpr)
	case "IsHiddenAngle0Shape":
		instance.IsHiddenAngle0Shape = GongExtractBool(valueExpr)
	case "IsHiddenTiledFloor3DShape":
		instance.IsHiddenTiledFloor3DShape = GongExtractBool(valueExpr)
	case "IsHiddenTopCurvePlane1Shape":
		instance.IsHiddenTopCurvePlane1Shape = GongExtractBool(valueExpr)
	case "IsHiddenBottomCurvePlane1Shape":
		instance.IsHiddenBottomCurvePlane1Shape = GongExtractBool(valueExpr)
	case "IsHiddenTopCurvePlane2Shape":
		instance.IsHiddenTopCurvePlane2Shape = GongExtractBool(valueExpr)
	case "IsHiddenBottomCurvePlane2Shape":
		instance.IsHiddenBottomCurvePlane2Shape = GongExtractBool(valueExpr)
	case "IsHiddenCarvedOutTopCurvePlane1Shape":
		instance.IsHiddenCarvedOutTopCurvePlane1Shape = GongExtractBool(valueExpr)
	case "IsHiddenCarvedOutBottomCurvePlane1Shape":
		instance.IsHiddenCarvedOutBottomCurvePlane1Shape = GongExtractBool(valueExpr)
	case "IsHiddenVaseTrapezeRingShape":
		instance.IsHiddenVaseTrapezeRingShape = GongExtractBool(valueExpr)
	case "IsHiddenCarvedOutVaseTrapezeRingShape":
		instance.IsHiddenCarvedOutVaseTrapezeRingShape = GongExtractBool(valueExpr)
	case "IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape":
		instance.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape = GongExtractBool(valueExpr)
	case "IsHiddenVaseTrapezeBasePlateShape":
		instance.IsHiddenVaseTrapezeBasePlateShape = GongExtractBool(valueExpr)
	case "Rendered3DShape":
		GongUnmarshallPointer(&instance.Rendered3DShape, valueExpr, identifierMap)
	case "TorusStackShape":
		GongUnmarshallPointer(&instance.TorusStackShape, valueExpr, identifierMap)
	case "VerticalTorusStackShape":
		GongUnmarshallPointer(&instance.VerticalTorusStackShape, valueExpr, identifierMap)
	case "PartiallyRotatedTorusShape":
		GongUnmarshallPointer(&instance.PartiallyRotatedTorusShape, valueExpr, identifierMap)
	case "StackOfPartiallyRotatedTorusShape":
		GongUnmarshallPointer(&instance.StackOfPartiallyRotatedTorusShape, valueExpr, identifierMap)
	case "PointsAndLines3DShape":
		GongUnmarshallPointer(&instance.PointsAndLines3DShape, valueExpr, identifierMap)
	case "SampledPoints3DShape":
		GongUnmarshallPointer(&instance.SampledPoints3DShape, valueExpr, identifierMap)
	case "OriginalPoints3DShape":
		GongUnmarshallPointer(&instance.OriginalPoints3DShape, valueExpr, identifierMap)
	case "Angle0Shape":
		GongUnmarshallPointer(&instance.Angle0Shape, valueExpr, identifierMap)
	case "KeyHole3DShape":
		GongUnmarshallPointer(&instance.KeyHole3DShape, valueExpr, identifierMap)
	case "Key3DShape":
		GongUnmarshallPointer(&instance.Key3DShape, valueExpr, identifierMap)
	case "VolumeKey3DShape":
		GongUnmarshallPointer(&instance.VolumeKey3DShape, valueExpr, identifierMap)
	case "TorusEdge3DShape":
		GongUnmarshallPointer(&instance.TorusEdge3DShape, valueExpr, identifierMap)
	case "TiledFloor3DShape":
		GongUnmarshallPointer(&instance.TiledFloor3DShape, valueExpr, identifierMap)
	case "TopCurvePlane1Shape":
		GongUnmarshallPointer(&instance.TopCurvePlane1Shape, valueExpr, identifierMap)
	case "BottomCurvePlane1Shape":
		GongUnmarshallPointer(&instance.BottomCurvePlane1Shape, valueExpr, identifierMap)
	case "TopCurvePlane2Shape":
		GongUnmarshallPointer(&instance.TopCurvePlane2Shape, valueExpr, identifierMap)
	case "BottomCurvePlane2Shape":
		GongUnmarshallPointer(&instance.BottomCurvePlane2Shape, valueExpr, identifierMap)
	case "CarvedOutTopCurvePlane1Shape":
		GongUnmarshallPointer(&instance.CarvedOutTopCurvePlane1Shape, valueExpr, identifierMap)
	case "CarvedOutBottomCurvePlane1Shape":
		GongUnmarshallPointer(&instance.CarvedOutBottomCurvePlane1Shape, valueExpr, identifierMap)
	case "VaseTrapezeRingShape":
		GongUnmarshallPointer(&instance.VaseTrapezeRingShape, valueExpr, identifierMap)
	case "CarvedOutVaseTrapezeRingShape":
		GongUnmarshallPointer(&instance.CarvedOutVaseTrapezeRingShape, valueExpr, identifierMap)
	case "StackOfRotatedCarvedOutVaseTrapezeRingsShape":
		GongUnmarshallPointer(&instance.StackOfRotatedCarvedOutVaseTrapezeRingsShape, valueExpr, identifierMap)
	case "VaseTrapezeBasePlateShape":
		GongUnmarshallPointer(&instance.VaseTrapezeBasePlateShape, valueExpr, identifierMap)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type TubeVaseAbstractUnmarshaller struct{}

func (u *TubeVaseAbstractUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TubeVaseAbstract), stage, identifier, instanceName, preserveOrder)
}

func (u *TubeVaseAbstractUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TubeVaseAbstract)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Z_Ribbon":
		instance.Z_Ribbon = GongExtractFloat(valueExpr)
	case "RibbonVerticalScale":
		instance.RibbonVerticalScale = GongExtractFloat(valueExpr)
	case "TopPlaneHeight":
		instance.TopPlaneHeight = GongExtractFloat(valueExpr)
	case "BottomPlaneHeight":
		instance.BottomPlaneHeight = GongExtractFloat(valueExpr)
	case "BaseRingBottomPlaneHeight":
		instance.BaseRingBottomPlaneHeight = GongExtractFloat(valueExpr)
	case "ProjectionAngle":
		instance.ProjectionAngle = GongExtractFloat(valueExpr)
	case "BasePlateHeight":
		instance.BasePlateHeight = GongExtractFloat(valueExpr)
	case "CarvedOutTopRingsParameter":
		instance.CarvedOutTopRingsParameter = GongExtractFloat(valueExpr)
	case "BulbousStartTangentMagnitude":
		instance.BulbousStartTangentMagnitude = GongExtractFloat(valueExpr)
	case "BulbousEndAngle":
		instance.BulbousEndAngle = GongExtractFloat(valueExpr)
	case "BulbousEndTangentMagnitude":
		instance.BulbousEndTangentMagnitude = GongExtractFloat(valueExpr)
	case "RelativeVerticalThickness":
		instance.RelativeVerticalThickness = GongExtractFloat(valueExpr)
	case "RelativeRadialThickness":
		instance.RelativeRadialThickness = GongExtractFloat(valueExpr)
	case "RelativeCuttedStackFloorHeight":
		instance.RelativeCuttedStackFloorHeight = GongExtractFloat(valueExpr)
	case "RelativeRotatedTorusSeparation":
		instance.RelativeRotatedTorusSeparation = GongExtractFloat(valueExpr)
	case "RotationRatio":
		instance.RotationRatio = GongExtractFloat(valueExpr)
	case "RadialRepetitions":
		instance.RadialRepetitions = GongExtractInt(valueExpr)
	case "Transparency":
		instance.Transparency = GongExtractFloat(valueExpr)
	case "HasAlternatingRingColors":
		instance.HasAlternatingRingColors = GongExtractBool(valueExpr)
	case "RelativeTrajectoryOffsetX":
		instance.RelativeTrajectoryOffsetX = GongExtractFloat(valueExpr)
	case "RelativeTrajectoryOffsetY":
		instance.RelativeTrajectoryOffsetY = GongExtractFloat(valueExpr)
	case "NbStepP1P2":
		instance.NbStepP1P2 = GongExtractInt(valueExpr)
	case "ChosenStep":
		instance.ChosenStep = GongExtractInt(valueExpr)
	case "RelativeHorizontalRingsHeight":
		instance.RelativeHorizontalRingsHeight = GongExtractFloat(valueExpr)
	case "OffsetKeyX":
		instance.OffsetKeyX = GongExtractFloat(valueExpr)
	case "OffsetKeyY":
		instance.OffsetKeyY = GongExtractFloat(valueExpr)
	case "HeightKey":
		instance.HeightKey = GongExtractFloat(valueExpr)
	case "WidthKey":
		instance.WidthKey = GongExtractFloat(valueExpr)
	case "RelativeKeySize":
		instance.RelativeKeySize = GongExtractFloat(valueExpr)
	case "MovieNbFrames":
		instance.MovieNbFrames = GongExtractInt(valueExpr)
	case "PerpendicularVectorGridHalfway":
		GongUnmarshallPointer(&instance.PerpendicularVectorGridHalfway, valueExpr, identifierMap)
	case "TopStartArcShapeGrid":
		GongUnmarshallPointer(&instance.TopStartArcShapeGrid, valueExpr, identifierMap)
	case "TopEndArcShapeGrid":
		GongUnmarshallPointer(&instance.TopEndArcShapeGrid, valueExpr, identifierMap)
	case "ShiftedBottomTopStartArcShapeGrid":
		GongUnmarshallPointer(&instance.ShiftedBottomTopStartArcShapeGrid, valueExpr, identifierMap)
	case "TopMidArcVectorShapeGrid":
		GongUnmarshallPointer(&instance.TopMidArcVectorShapeGrid, valueExpr, identifierMap)
	case "StartHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.StartHalfwayArcShapeGrid, valueExpr, identifierMap)
	case "TopStartHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.TopStartHalfwayArcShapeGrid, valueExpr, identifierMap)
	case "EndHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.EndHalfwayArcShapeGrid, valueExpr, identifierMap)
	case "TopEndHalfwayArcShapeGrid":
		GongUnmarshallPointer(&instance.TopEndHalfwayArcShapeGrid, valueExpr, identifierMap)
	case "StackOfRotatedGrowthCurve2D":
		GongUnmarshallPointer(&instance.StackOfRotatedGrowthCurve2D, valueExpr, identifierMap)
	case "TopStackOfRotatedGrowthCurve2D":
		GongUnmarshallPointer(&instance.TopStackOfRotatedGrowthCurve2D, valueExpr, identifierMap)
	case "TopGrowthCurve2D":
		GongUnmarshallPointer(&instance.TopGrowthCurve2D, valueExpr, identifierMap)
	case "StackOfGrowthCurve2D":
		GongUnmarshallPointer(&instance.StackOfGrowthCurve2D, valueExpr, identifierMap)
	case "TopStackOfGrowthCurve2D":
		GongUnmarshallPointer(&instance.TopStackOfGrowthCurve2D, valueExpr, identifierMap)
	case "StackOfGrowthCurve2DRibbon":
		GongUnmarshallPointer(&instance.StackOfGrowthCurve2DRibbon, valueExpr, identifierMap)
	case "StackOfRotatedGrowthCurve2DRibbon":
		GongUnmarshallPointer(&instance.StackOfRotatedGrowthCurve2DRibbon, valueExpr, identifierMap)
	case "GrowthCurve2DRibbon":
		GongUnmarshallPointer(&instance.GrowthCurve2DRibbon, valueExpr, identifierMap)
	case "ShiftedRightGrowthCurve2DRibbon":
		GongUnmarshallPointer(&instance.ShiftedRightGrowthCurve2DRibbon, valueExpr, identifierMap)
	case "ShiftedLeftGrowthCurve2DRibbon":
		GongUnmarshallPointer(&instance.ShiftedLeftGrowthCurve2DRibbon, valueExpr, identifierMap)
	case "PartiallyGrowthCurve2DRibbon":
		GongUnmarshallPointer(&instance.PartiallyGrowthCurve2DRibbon, valueExpr, identifierMap)
	case "ShiftedLeftPartiallyGrowthCurve2DRibbon":
		GongUnmarshallPointer(&instance.ShiftedLeftPartiallyGrowthCurve2DRibbon, valueExpr, identifierMap)
	case "PartiallyGrowthCurve2DTrajectory":
		GongUnmarshallPointer(&instance.PartiallyGrowthCurve2DTrajectory, valueExpr, identifierMap)
	case "PartiallyGrowthCurve2DTrajectoryP1P2":
		GongUnmarshallPointer(&instance.PartiallyGrowthCurve2DTrajectoryP1P2, valueExpr, identifierMap)
	case "PxShape":
		GongUnmarshallPointer(&instance.PxShape, valueExpr, identifierMap)
	case "ChosenP1P2PairShape":
		GongUnmarshallPointer(&instance.ChosenP1P2PairShape, valueExpr, identifierMap)
	case "KeyHoleShape":
		GongUnmarshallPointer(&instance.KeyHoleShape, valueExpr, identifierMap)
	}
	return nil
}

type Vase2DDiagramUnmarshaller struct{}

func (u *Vase2DDiagramUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(Vase2DDiagram), stage, identifier, instanceName, preserveOrder)
}

func (u *Vase2DDiagramUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*Vase2DDiagram)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Zoom":
		instance.Zoom = GongExtractFloat(valueExpr)
	case "IsVaseArcNodesExpanded":
		instance.IsVaseArcNodesExpanded = GongExtractBool(valueExpr)
	case "IsVaseClampingNodesExpanded":
		instance.IsVaseClampingNodesExpanded = GongExtractBool(valueExpr)
	case "IsHiddenAxesShape":
		instance.IsHiddenAxesShape = GongExtractBool(valueExpr)
	case "IsHiddenBottomStartArcShapeGrid":
		instance.IsHiddenBottomStartArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenBottomEndArcShapeGrid":
		instance.IsHiddenBottomEndArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenBottomStackOfGrowthCurve":
		instance.IsHiddenBottomStackOfGrowthCurve = GongExtractBool(valueExpr)
	case "IsHiddenShiftedLeftStackOfGrowthCurve":
		instance.IsHiddenShiftedLeftStackOfGrowthCurve = GongExtractBool(valueExpr)
	case "IsHiddenShiftedLeftStackOfNormalVector":
		instance.IsHiddenShiftedLeftStackOfNormalVector = GongExtractBool(valueExpr)
	case "IsHiddenPerpendicularVectorGridHalfway":
		instance.IsHiddenPerpendicularVectorGridHalfway = GongExtractBool(valueExpr)
	case "IsHiddenTopStartArcShapeGrid":
		instance.IsHiddenTopStartArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenShiftedBottomTopStartArcShapeGrid":
		instance.IsHiddenShiftedBottomTopStartArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenTopMidArcVectorShapeGrid":
		instance.IsHiddenTopMidArcVectorShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenStartHalfwayArcShapeGrid":
		instance.IsHiddenStartHalfwayArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenTopStartHalfwayArcShapeGrid":
		instance.IsHiddenTopStartHalfwayArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenEndHalfwayArcShapeGrid":
		instance.IsHiddenEndHalfwayArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenTopEndHalfwayArcShapeGrid":
		instance.IsHiddenTopEndHalfwayArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenTopEndArcShapeGrid":
		instance.IsHiddenTopEndArcShapeGrid = GongExtractBool(valueExpr)
	case "IsHiddenStackOfGrowthCurve":
		instance.IsHiddenStackOfGrowthCurve = GongExtractBool(valueExpr)
	case "IsHiddenTopStackOfGrowthCurve":
		instance.IsHiddenTopStackOfGrowthCurve = GongExtractBool(valueExpr)
	case "IsHiddenTopGrowthCurve2D":
		instance.IsHiddenTopGrowthCurve2D = GongExtractBool(valueExpr)
	case "IsHiddenStackOfGrowthCurve2D":
		instance.IsHiddenStackOfGrowthCurve2D = GongExtractBool(valueExpr)
	case "IsHiddenTopStackOfGrowthCurve2D":
		instance.IsHiddenTopStackOfGrowthCurve2D = GongExtractBool(valueExpr)
	case "IsHiddenGrowthCurve2DRibbon":
		instance.IsHiddenGrowthCurve2DRibbon = GongExtractBool(valueExpr)
	case "IsHiddenShiftedRightGrowthCurve2DRibbon":
		instance.IsHiddenShiftedRightGrowthCurve2DRibbon = GongExtractBool(valueExpr)
	case "IsHiddenShiftedLeftGrowthCurve2DRibbon":
		instance.IsHiddenShiftedLeftGrowthCurve2DRibbon = GongExtractBool(valueExpr)
	case "IsHiddenStackOfGrowthCurve2DRibbon":
		instance.IsHiddenStackOfGrowthCurve2DRibbon = GongExtractBool(valueExpr)
	case "IsHiddenStackOfRotatedGrowthCurve2DRibbon":
		instance.IsHiddenStackOfRotatedGrowthCurve2DRibbon = GongExtractBool(valueExpr)
	case "IsHiddenPartiallyGrowthCurve2DRibbon":
		instance.IsHiddenPartiallyGrowthCurve2DRibbon = GongExtractBool(valueExpr)
	case "IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon":
		instance.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon = GongExtractBool(valueExpr)
	case "IsHiddenPartiallyGrowthCurve2DTrajectory":
		instance.IsHiddenPartiallyGrowthCurve2DTrajectory = GongExtractBool(valueExpr)
	case "IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2":
		instance.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2 = GongExtractBool(valueExpr)
	case "IsHiddenPxShape":
		instance.IsHiddenPxShape = GongExtractBool(valueExpr)
	case "IsHiddenChosenP1P2PairShape":
		instance.IsHiddenChosenP1P2PairShape = GongExtractBool(valueExpr)
	case "IsHiddenKeyHoleShape":
		instance.IsHiddenKeyHoleShape = GongExtractBool(valueExpr)
	case "IsChecked":
		instance.IsChecked = GongExtractBool(valueExpr)
	case "ComputedPrefix":
		instance.ComputedPrefix = GongExtractString(valueExpr)
	case "IsExpanded":
		instance.IsExpanded = GongExtractBool(valueExpr)
	}
	return nil
}

type VaseTrapezeBasePlateShapeUnmarshaller struct{}

func (u *VaseTrapezeBasePlateShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(VaseTrapezeBasePlateShape), stage, identifier, instanceName, preserveOrder)
}

func (u *VaseTrapezeBasePlateShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VaseTrapezeBasePlateShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type VaseTrapezeRingShapeUnmarshaller struct{}

func (u *VaseTrapezeRingShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(VaseTrapezeRingShape), stage, identifier, instanceName, preserveOrder)
}

func (u *VaseTrapezeRingShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VaseTrapezeRingShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type VerticalTorusStackShapeUnmarshaller struct{}

func (u *VerticalTorusStackShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(VerticalTorusStackShape), stage, identifier, instanceName, preserveOrder)
}

func (u *VerticalTorusStackShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VerticalTorusStackShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}

type VolumeKey3DShapeUnmarshaller struct{}

func (u *VolumeKey3DShapeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(VolumeKey3DShape), stage, identifier, instanceName, preserveOrder)
}

func (u *VolumeKey3DShapeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*VolumeKey3DShape)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	}
	return nil
}
