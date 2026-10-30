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
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/xlink"
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

	if stage.XlinkStage != nil {
		for _, instance := range stage.XlinkStage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage.XlinkStage)] = instance
		}
	}
	if stage.GmlStage != nil {
		for _, instance := range stage.GmlStage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage.GmlStage)] = instance
		}
	}
	if stage.AixmabstractgmlobjecttypesStage != nil {
		for _, instance := range stage.AixmabstractgmlobjecttypesStage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage.AixmabstractgmlobjecttypesStage)] = instance
		}
	}
	if stage.AixmdatatypesStage != nil {
		for _, instance := range stage.AixmdatatypesStage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage.AixmdatatypesStage)] = instance
		}
	}
	if stage.AixmfeaturesStage != nil {
		for _, instance := range stage.AixmfeaturesStage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage.AixmfeaturesStage)] = instance
		}
	}
	if stage != nil {
		for _, instance := range stage.GetInstances() {
			identifierMap[instance.GongGetIdentifier(stage)] = instance
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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/xlink":
			aliasToCanonical[alias] = "xlink"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/gml":
			aliasToCanonical[alias] = "gml"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmabstractgmlobjecttypes":
			aliasToCanonical[alias] = "aixmabstractgmlobjecttypes"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmdatatypes":
			aliasToCanonical[alias] = "aixmdatatypes"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures":
			aliasToCanonical[alias] = "aixmfeatures"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models":
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
			case "xlink":
				switch typeName {
				case "Xlink_SchemaElement":
					identifierMap[ident.Name] = __gong__stageSetInit(new(xlink.Xlink_SchemaElement), stage.XlinkStage, ident.Name, instanceName, preserveOrder)
				}
			case "gml":
				switch typeName {
				case "AbstractCurveSegmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractCurveSegmentType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractCurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractCurveType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractFeatureMemberType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractFeatureMemberType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractFeatureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractFeatureType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractGMLType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractGMLType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractGeometricPrimitiveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractGeometricPrimitiveType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractGeometryType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractGeometryType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractMetadataPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractMetadataPropertyType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractRingPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractRingPropertyType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractRingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractRingType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractSurfacePatchType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractSurfacePatchType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractSurfaceType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractTimeGeometricPrimitiveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractTimeGeometricPrimitiveType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractTimeObjectType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractTimeObjectType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractTimePrimitiveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractTimePrimitiveType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractTimeSliceType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AngleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AngleType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "ArcByCenterPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.ArcByCenterPointType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "ArcStringType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.ArcStringType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "ArcType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.ArcType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "BoundingShapeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.BoundingShapeType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "CircleByCenterPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.CircleByCenterPointType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "CircleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.CircleType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "CodeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.CodeType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "CodeWithAuthorityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.CodeWithAuthorityType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "CompositeCurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.CompositeCurveType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "CurvePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.CurvePropertyType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "CurveSegmentArrayPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.CurveSegmentArrayPropertyType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "CurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.CurveType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "DirectPositionListType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.DirectPositionListType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "DirectPositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.DirectPositionType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "DynamicFeatureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.DynamicFeatureType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "EnvelopeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.EnvelopeType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "GeodesicStringType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.GeodesicStringType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "GeodesicType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.GeodesicType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "LengthType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.LengthType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "LineStringSegmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.LineStringSegmentType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "MeasureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.MeasureType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "OrientableCurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.OrientableCurveType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "PointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.PointPropertyType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "PointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.PointType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "PolygonPatchType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.PolygonPatchType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "RingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.RingType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "SurfacePatchArrayPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.SurfacePatchArrayPropertyType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "SurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.SurfaceType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "TimeInstantType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.TimeInstantType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "TimePeriodType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.TimePeriodType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "TimePositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.TimePositionType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "TimePrimitivePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.TimePrimitivePropertyType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				}
			case "aixmabstractgmlobjecttypes":
				switch typeName {
				case "AbstractAIXMFeatureBaseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType), stage.AixmabstractgmlobjecttypesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractAIXMFeatureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmabstractgmlobjecttypes.AbstractAIXMFeatureType), stage.AixmabstractgmlobjecttypesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractAIXMMessageBaseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType), stage.AixmabstractgmlobjecttypesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractAIXMMessageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmabstractgmlobjecttypes.AbstractAIXMMessageType), stage.AixmabstractgmlobjecttypesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractAIXMObjectType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmabstractgmlobjecttypes.AbstractAIXMObjectType), stage.AixmabstractgmlobjecttypesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractAIXMPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmabstractgmlobjecttypes.AbstractAIXMPropertyType), stage.AixmabstractgmlobjecttypesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractAIXMTimeSliceBaseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType), stage.AixmabstractgmlobjecttypesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractAIXMTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType), stage.AixmabstractgmlobjecttypesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractExtensionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmabstractgmlobjecttypes.AbstractExtensionType), stage.AixmabstractgmlobjecttypesStage, ident.Name, instanceName, preserveOrder)
				case "FeatureMetadataPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmabstractgmlobjecttypes.FeatureMetadataPropertyType), stage.AixmabstractgmlobjecttypesStage, ident.Name, instanceName, preserveOrder)
				case "FeatureTimeSliceMetadataPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType), stage.AixmabstractgmlobjecttypesStage, ident.Name, instanceName, preserveOrder)
				case "MessageMetadataPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmabstractgmlobjecttypes.MessageMetadataPropertyType), stage.AixmabstractgmlobjecttypesStage, ident.Name, instanceName, preserveOrder)
				}
			case "aixmdatatypes":
				switch typeName {
				case "CodeColourType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeColourType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeDayType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeGroundLightingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeGroundLightingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLightIntensityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLightIntensityType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLightSourceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLightSourceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeNotePurposeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeNotePurposeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeObstacleAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeObstacleAreaType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeObstacleAssessmentSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeObstacleAssessmentSurfaceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusConstructionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeStatusConstructionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusOperationsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeStatusOperationsType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventCombinationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTimeEventCombinationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTimeEventType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTimeReferenceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalDatumType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeVerticalDatumType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalStructureMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeVerticalStructureMarkingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalStructureMaterialType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeVerticalStructureMaterialType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalStructureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeVerticalStructureType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeYesNoType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeYesNoType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "DateMonthDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.DateMonthDayType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextDesignatorType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextNameType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextNameType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextNoteType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextPropertyNameType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextPropertyNameType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TimeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TimeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceSignedType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDistanceSignedType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDistanceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceVerticalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDistanceVerticalType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDurationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDurationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValLightIntensityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValLightIntensityType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				}
			case "aixmfeatures":
				switch typeName {
				case "A_extension":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_1":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_1), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_10":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_10), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_11":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_11), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_12":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_12), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_2":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_2), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_3":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_3), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_4":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_4), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_5":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_5), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_6":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_6), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_7":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_7), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_8":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_8), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_9":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_9), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractPropertiesWithScheduleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractPropertiesWithScheduleType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AeronauticalGroundLightTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AeronauticalGroundLightTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AeronauticalGroundLightTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AeronauticalGroundLightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AeronauticalGroundLightType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CurveType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedCurvePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedCurvePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedCurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedCurveType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedPointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedPointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedSurfacePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedSurfacePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedSurfaceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GroundLightSystemPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GroundLightSystemPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LightElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LightElementPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LightElementStatusPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LightElementStatusPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LightElementStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LightElementStatusType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LightElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LightElementType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LinguisticNotePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LinguisticNoteType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MarkerBeaconPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MarkerBeaconPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavaidEquipmentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavaidEquipmentPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NotePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NoteType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ObstacleAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ObstacleAreaTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ObstacleAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ObstacleAreaTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ObstacleAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ObstacleAreaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OrganisationAuthorityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OrganisationAuthorityPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PassengerServicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PassengerServicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayDirectionPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ServicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ServicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SpecialNavigationStationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SpecialNavigationStationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurfacePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurfacePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurfaceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TimesheetPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TimesheetPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TimesheetType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TimesheetType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "UnitPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.UnitPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "VerticalStructureLightingStatusPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.VerticalStructureLightingStatusPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "VerticalStructureLightingStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.VerticalStructureLightingStatusType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "VerticalStructurePartPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.VerticalStructurePartPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "VerticalStructurePartType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.VerticalStructurePartType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "VerticalStructurePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.VerticalStructurePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "VerticalStructureTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.VerticalStructureTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "VerticalStructureTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.VerticalStructureTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "VerticalStructureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.VerticalStructureType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				}
			case "models":
				switch typeName {
				case "AIXMBasicMessage":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AIXMBasicMessage), stage, ident.Name, instanceName, preserveOrder)
				case "AIXMBasicMessageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(AIXMBasicMessageType), stage, ident.Name, instanceName, preserveOrder)
				case "BasicMessageMemberAIXMPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(BasicMessageMemberAIXMPropertyType), stage, ident.Name, instanceName, preserveOrder)
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
				case *xlink.Xlink_SchemaElement:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *gml.AbstractCurveSegmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					}
				case *gml.AbstractCurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					}
				case *gml.AbstractFeatureMemberType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *gml.AbstractFeatureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BoundedBy":
						__gong__assignPointer(&inst.BoundedBy, rhs, identifierMap)
					}
				case *gml.AbstractGMLType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					}
				case *gml.AbstractGeometricPrimitiveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					}
				case *gml.AbstractGeometryType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					}
				case *gml.AbstractMetadataPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *gml.AbstractRingPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Ring":
						__gong__assignPointer(&inst.Ring, rhs, identifierMap)
					}
				case *gml.AbstractRingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *gml.AbstractSurfacePatchType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					}
				case *gml.AbstractSurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					}
				case *gml.AbstractTimeGeometricPrimitiveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Frame":
						inst.Frame = GongExtractString(rhs)
					}
				case *gml.AbstractTimeObjectType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					}
				case *gml.AbstractTimePrimitiveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					}
				case *gml.AbstractTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "ValidTime":
						__gong__assignPointer(&inst.ValidTime, rhs, identifierMap)
					}
				case *gml.AngleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *gml.ArcByCenterPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "NumArc":
						inst.NumArc = GongExtractInt(rhs)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignPointer(&inst.PointProperty, rhs, identifierMap)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					case "Radius":
						__gong__assignPointer(&inst.Radius, rhs, identifierMap)
					case "StartAngle":
						__gong__assignPointer(&inst.StartAngle, rhs, identifierMap)
					case "EndAngle":
						__gong__assignPointer(&inst.EndAngle, rhs, identifierMap)
					}
				case *gml.ArcStringType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "NumArc":
						inst.NumArc = GongExtractInt(rhs)
					case "Pos":
						__gong__assignSliceOfPointers(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignSliceOfPointers(&inst.PointProperty, rhs, identifierMap)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					}
				case *gml.ArcType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "NumArc":
						inst.NumArc = GongExtractInt(rhs)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignPointer(&inst.PointProperty, rhs, identifierMap)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					}
				case *gml.BoundingShapeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Envelope":
						__gong__assignPointer(&inst.Envelope, rhs, identifierMap)
					}
				case *gml.CircleByCenterPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "NumArc":
						inst.NumArc = GongExtractInt(rhs)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignPointer(&inst.PointProperty, rhs, identifierMap)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					case "Radius":
						__gong__assignPointer(&inst.Radius, rhs, identifierMap)
					}
				case *gml.CircleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "NumArc":
						inst.NumArc = GongExtractInt(rhs)
					case "Pos":
						__gong__assignSliceOfPointers(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignSliceOfPointers(&inst.PointProperty, rhs, identifierMap)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					}
				case *gml.CodeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "CodeSpace":
						inst.CodeSpace = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *gml.CodeWithAuthorityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "CodeSpace":
						inst.CodeSpace = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *gml.CompositeCurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "AggregationType":
						inst.AggregationType = GongExtractString(rhs)
					case "CurveMember":
						__gong__assignSliceOfPointers(&inst.CurveMember, rhs, identifierMap)
					}
				case *gml.CurvePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					case "Curve":
						__gong__assignPointer(&inst.Curve, rhs, identifierMap)
					case "OrientableCurve":
						__gong__assignPointer(&inst.OrientableCurve, rhs, identifierMap)
					case "CompositeCurve":
						__gong__assignPointer(&inst.CompositeCurve, rhs, identifierMap)
					}
				case *gml.CurveSegmentArrayPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "LineStringSegment":
						__gong__assignSliceOfPointers(&inst.LineStringSegment, rhs, identifierMap)
					case "ArcString":
						__gong__assignSliceOfPointers(&inst.ArcString, rhs, identifierMap)
					case "Arc":
						__gong__assignSliceOfPointers(&inst.Arc, rhs, identifierMap)
					case "Circle":
						__gong__assignSliceOfPointers(&inst.Circle, rhs, identifierMap)
					case "ArcByCenterPoint":
						__gong__assignSliceOfPointers(&inst.ArcByCenterPoint, rhs, identifierMap)
					case "CircleByCenterPoint":
						__gong__assignSliceOfPointers(&inst.CircleByCenterPoint, rhs, identifierMap)
					case "GeodesicString":
						__gong__assignSliceOfPointers(&inst.GeodesicString, rhs, identifierMap)
					case "Geodesic":
						__gong__assignSliceOfPointers(&inst.Geodesic, rhs, identifierMap)
					}
				case *gml.CurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Segments":
						__gong__assignPointer(&inst.Segments, rhs, identifierMap)
					}
				case *gml.DirectPositionListType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *gml.DirectPositionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *gml.DynamicFeatureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BoundedBy":
						__gong__assignPointer(&inst.BoundedBy, rhs, identifierMap)
					}
				case *gml.EnvelopeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "LowerCorner":
						__gong__assignPointer(&inst.LowerCorner, rhs, identifierMap)
					case "UpperCorner":
						__gong__assignPointer(&inst.UpperCorner, rhs, identifierMap)
					}
				case *gml.GeodesicStringType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignPointer(&inst.PointProperty, rhs, identifierMap)
					}
				case *gml.GeodesicType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignPointer(&inst.PointProperty, rhs, identifierMap)
					}
				case *gml.LengthType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *gml.LineStringSegmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NumDerivativesAtStart":
						inst.NumDerivativesAtStart = GongExtractInt(rhs)
					case "NumDerivativesAtEnd":
						inst.NumDerivativesAtEnd = GongExtractInt(rhs)
					case "NumDerivativeInterior":
						inst.NumDerivativeInterior = GongExtractInt(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "Pos":
						__gong__assignSliceOfPointers(&inst.Pos, rhs, identifierMap)
					case "PointProperty":
						__gong__assignSliceOfPointers(&inst.PointProperty, rhs, identifierMap)
					case "PosList":
						__gong__assignPointer(&inst.PosList, rhs, identifierMap)
					}
				case *gml.MeasureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractFloat(rhs)
					}
				case *gml.OrientableCurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Orientation":
						GongUnmarshallEnum(&inst.Orientation, rhs)
					case "BaseCurve":
						__gong__assignPointer(&inst.BaseCurve, rhs, identifierMap)
					}
				case *gml.PointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					case "Point":
						__gong__assignPointer(&inst.Point, rhs, identifierMap)
					}
				case *gml.PointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					}
				case *gml.PolygonPatchType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Interpolation":
						GongUnmarshallEnum(&inst.Interpolation, rhs)
					case "Exterior":
						__gong__assignPointer(&inst.Exterior, rhs, identifierMap)
					}
				case *gml.RingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "AggregationType":
						inst.AggregationType = GongExtractString(rhs)
					case "CurveMember":
						__gong__assignSliceOfPointers(&inst.CurveMember, rhs, identifierMap)
					}
				case *gml.SurfacePatchArrayPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "PolygonPatch":
						__gong__assignSliceOfPointers(&inst.PolygonPatch, rhs, identifierMap)
					}
				case *gml.SurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Patches":
						__gong__assignPointer(&inst.Patches, rhs, identifierMap)
					}
				case *gml.TimeInstantType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Frame":
						inst.Frame = GongExtractString(rhs)
					case "TimePosition":
						__gong__assignPointer(&inst.TimePosition, rhs, identifierMap)
					}
				case *gml.TimePeriodType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Frame":
						inst.Frame = GongExtractString(rhs)
					case "BeginPosition":
						__gong__assignPointer(&inst.BeginPosition, rhs, identifierMap)
					case "EndPosition":
						__gong__assignPointer(&inst.EndPosition, rhs, identifierMap)
					}
				case *gml.TimePositionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Frame":
						inst.Frame = GongExtractString(rhs)
					case "IndeterminatePosition":
						GongUnmarshallEnum(&inst.IndeterminatePosition, rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *gml.TimePrimitivePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					case "TimeInstant":
						__gong__assignPointer(&inst.TimeInstant, rhs, identifierMap)
					case "TimePeriod":
						__gong__assignPointer(&inst.TimePeriod, rhs, identifierMap)
					}
				case *aixmabstractgmlobjecttypes.AbstractAIXMFeatureBaseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BoundedBy":
						__gong__assignPointer(&inst.BoundedBy, rhs, identifierMap)
					}
				case *aixmabstractgmlobjecttypes.AbstractAIXMFeatureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BoundedBy":
						__gong__assignPointer(&inst.BoundedBy, rhs, identifierMap)
					case "FeatureMetadata":
						__gong__assignPointer(&inst.FeatureMetadata, rhs, identifierMap)
					}
				case *aixmabstractgmlobjecttypes.AbstractAIXMMessageBaseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BoundedBy":
						__gong__assignPointer(&inst.BoundedBy, rhs, identifierMap)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "MessageMetadata":
						__gong__assignPointer(&inst.MessageMetadata, rhs, identifierMap)
					}
				case *aixmabstractgmlobjecttypes.AbstractAIXMMessageType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BoundedBy":
						__gong__assignPointer(&inst.BoundedBy, rhs, identifierMap)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "MessageMetadata":
						__gong__assignPointer(&inst.MessageMetadata, rhs, identifierMap)
					case "AggregationType":
						inst.AggregationType = GongExtractString(rhs)
					}
				case *aixmabstractgmlobjecttypes.AbstractAIXMObjectType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					}
				case *aixmabstractgmlobjecttypes.AbstractAIXMPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					}
				case *aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceBaseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "ValidTime":
						__gong__assignPointer(&inst.ValidTime, rhs, identifierMap)
					}
				case *aixmabstractgmlobjecttypes.AbstractAIXMTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "ValidTime":
						__gong__assignPointer(&inst.ValidTime, rhs, identifierMap)
					case "Interpretation":
						GongUnmarshallEnum(&inst.Interpretation, rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "TimeSliceMetadata":
						__gong__assignPointer(&inst.TimeSliceMetadata, rhs, identifierMap)
					case "FeatureLifetime":
						__gong__assignPointer(&inst.FeatureLifetime, rhs, identifierMap)
					}
				case *aixmabstractgmlobjecttypes.AbstractExtensionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					}
				case *aixmabstractgmlobjecttypes.FeatureMetadataPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *aixmabstractgmlobjecttypes.FeatureTimeSliceMetadataPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *aixmabstractgmlobjecttypes.MessageMetadataPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					}
				case *aixmdatatypes.CodeColourType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeDayType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeGroundLightingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeLightIntensityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeLightSourceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeNotePurposeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeObstacleAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeObstacleAssessmentSurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeStatusConstructionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeStatusOperationsType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeTimeEventCombinationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeTimeEventType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeTimeReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeVerticalDatumType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeVerticalStructureMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeVerticalStructureMaterialType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeVerticalStructureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeYesNoType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.DateMonthDayType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.TextDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.TextNameType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.TextNoteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Lang":
						inst.Lang = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.TextPropertyNameType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.TimeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.ValDistanceSignedType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.ValDistanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.ValDistanceVerticalType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.ValDurationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.ValLightIntensityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmfeatures.A_extension:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractElevatedCurveExtension":
						__gong__assignSliceOfPointers(&inst.AbstractElevatedCurveExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractElevatedPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractElevatedPointExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_10:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractLightElementExtension":
						__gong__assignSliceOfPointers(&inst.AbstractLightElementExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_11:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractLightElementStatusExtension":
						__gong__assignSliceOfPointers(&inst.AbstractLightElementStatusExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_12:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTimesheetExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTimesheetExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_2:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractElevatedSurfaceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractElevatedSurfaceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_3:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAeronauticalGroundLightExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAeronauticalGroundLightExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_4:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractLinguisticNoteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractLinguisticNoteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_5:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractNoteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractNoteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_6:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractObstacleAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractObstacleAreaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_7:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractVerticalStructureExtension":
						__gong__assignSliceOfPointers(&inst.AbstractVerticalStructureExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_8:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractVerticalStructureLightingStatusExtension":
						__gong__assignSliceOfPointers(&inst.AbstractVerticalStructureLightingStatusExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_9:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractVerticalStructurePartExtension":
						__gong__assignSliceOfPointers(&inst.AbstractVerticalStructurePartExtension, rhs, identifierMap)
					}
				case *aixmfeatures.AbstractPropertiesWithScheduleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					}
				case *aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AeronauticalGroundLightTimeSlice":
						__gong__assignPointer(&inst.AeronauticalGroundLightTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AeronauticalGroundLightTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "ValidTime":
						__gong__assignPointer(&inst.ValidTime, rhs, identifierMap)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "TimeSliceMetadata":
						__gong__assignPointer(&inst.TimeSliceMetadata, rhs, identifierMap)
					case "FeatureLifetime":
						__gong__assignPointer(&inst.FeatureLifetime, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Colour":
						__gong__assignPointer(&inst.Colour, rhs, identifierMap)
					case "Flashing":
						__gong__assignPointer(&inst.Flashing, rhs, identifierMap)
					case "StructureBeacon":
						__gong__assignPointer(&inst.StructureBeacon, rhs, identifierMap)
					case "AerodromeBeacon":
						__gong__assignPointer(&inst.AerodromeBeacon, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AeronauticalGroundLightType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BoundedBy":
						__gong__assignPointer(&inst.BoundedBy, rhs, identifierMap)
					case "FeatureMetadata":
						__gong__assignPointer(&inst.FeatureMetadata, rhs, identifierMap)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHeliportPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *aixmfeatures.CurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Segments":
						__gong__assignPointer(&inst.Segments, rhs, identifierMap)
					case "HorizontalAccuracy":
						__gong__assignPointer(&inst.HorizontalAccuracy, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					}
				case *aixmfeatures.ElevatedCurvePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ElevatedCurve":
						__gong__assignPointer(&inst.ElevatedCurve, rhs, identifierMap)
					}
				case *aixmfeatures.ElevatedCurveType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Segments":
						__gong__assignPointer(&inst.Segments, rhs, identifierMap)
					case "HorizontalAccuracy":
						__gong__assignPointer(&inst.HorizontalAccuracy, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Elevation":
						__gong__assignPointer(&inst.Elevation, rhs, identifierMap)
					case "GeoidUndulation":
						__gong__assignPointer(&inst.GeoidUndulation, rhs, identifierMap)
					case "VerticalDatum":
						__gong__assignPointer(&inst.VerticalDatum, rhs, identifierMap)
					case "VerticalAccuracy":
						__gong__assignPointer(&inst.VerticalAccuracy, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ElevatedPointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ElevatedPoint":
						__gong__assignPointer(&inst.ElevatedPoint, rhs, identifierMap)
					}
				case *aixmfeatures.ElevatedPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					case "HorizontalAccuracy":
						__gong__assignPointer(&inst.HorizontalAccuracy, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Elevation":
						__gong__assignPointer(&inst.Elevation, rhs, identifierMap)
					case "GeoidUndulation":
						__gong__assignPointer(&inst.GeoidUndulation, rhs, identifierMap)
					case "VerticalDatum":
						__gong__assignPointer(&inst.VerticalDatum, rhs, identifierMap)
					case "VerticalAccuracy":
						__gong__assignPointer(&inst.VerticalAccuracy, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ElevatedSurfacePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ElevatedSurface":
						__gong__assignPointer(&inst.ElevatedSurface, rhs, identifierMap)
					}
				case *aixmfeatures.ElevatedSurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Patches":
						__gong__assignPointer(&inst.Patches, rhs, identifierMap)
					case "HorizontalAccuracy":
						__gong__assignPointer(&inst.HorizontalAccuracy, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Elevation":
						__gong__assignPointer(&inst.Elevation, rhs, identifierMap)
					case "GeoidUndulation":
						__gong__assignPointer(&inst.GeoidUndulation, rhs, identifierMap)
					case "VerticalDatum":
						__gong__assignPointer(&inst.VerticalDatum, rhs, identifierMap)
					case "VerticalAccuracy":
						__gong__assignPointer(&inst.VerticalAccuracy, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.GroundLightSystemPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *aixmfeatures.LightElementPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "LightElement":
						__gong__assignPointer(&inst.LightElement, rhs, identifierMap)
					}
				case *aixmfeatures.LightElementStatusPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "LightElementStatus":
						__gong__assignPointer(&inst.LightElementStatus, rhs, identifierMap)
					}
				case *aixmfeatures.LightElementStatusType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Status":
						__gong__assignPointer(&inst.Status, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.LightElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Colour":
						__gong__assignPointer(&inst.Colour, rhs, identifierMap)
					case "IntensityLevel":
						__gong__assignPointer(&inst.IntensityLevel, rhs, identifierMap)
					case "Intensity":
						__gong__assignPointer(&inst.Intensity, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.LinguisticNotePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "LinguisticNote":
						__gong__assignPointer(&inst.LinguisticNote, rhs, identifierMap)
					}
				case *aixmfeatures.LinguisticNoteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.MarkerBeaconPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *aixmfeatures.NavaidEquipmentPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *aixmfeatures.NotePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Note":
						__gong__assignPointer(&inst.Note, rhs, identifierMap)
					}
				case *aixmfeatures.NoteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "PropertyName":
						__gong__assignPointer(&inst.PropertyName, rhs, identifierMap)
					case "Purpose":
						__gong__assignPointer(&inst.Purpose, rhs, identifierMap)
					case "TranslatedNote":
						__gong__assignSliceOfPointers(&inst.TranslatedNote, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ObstacleAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ObstacleAreaTimeSlice":
						__gong__assignPointer(&inst.ObstacleAreaTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.ObstacleAreaTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "ValidTime":
						__gong__assignPointer(&inst.ValidTime, rhs, identifierMap)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "TimeSliceMetadata":
						__gong__assignPointer(&inst.TimeSliceMetadata, rhs, identifierMap)
					case "FeatureLifetime":
						__gong__assignPointer(&inst.FeatureLifetime, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "ObstructionIdSurfaceCondition":
						__gong__assignPointer(&inst.ObstructionIdSurfaceCondition, rhs, identifierMap)
					case "Reference_ownerOrganisation":
						__gong__assignPointer(&inst.Reference_ownerOrganisation, rhs, identifierMap)
					case "Reference_ownerRunway":
						__gong__assignPointer(&inst.Reference_ownerRunway, rhs, identifierMap)
					case "Reference_ownerAirport":
						__gong__assignPointer(&inst.Reference_ownerAirport, rhs, identifierMap)
					case "SurfaceExtent":
						__gong__assignPointer(&inst.SurfaceExtent, rhs, identifierMap)
					case "Obstacle":
						__gong__assignSliceOfPointers(&inst.Obstacle, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ObstacleAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BoundedBy":
						__gong__assignPointer(&inst.BoundedBy, rhs, identifierMap)
					case "FeatureMetadata":
						__gong__assignPointer(&inst.FeatureMetadata, rhs, identifierMap)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.OrganisationAuthorityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *aixmfeatures.PassengerServicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *aixmfeatures.PointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Pos":
						__gong__assignPointer(&inst.Pos, rhs, identifierMap)
					case "HorizontalAccuracy":
						__gong__assignPointer(&inst.HorizontalAccuracy, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayDirectionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *aixmfeatures.ServicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *aixmfeatures.SpecialNavigationStationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *aixmfeatures.SurfacePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Surface":
						__gong__assignPointer(&inst.Surface, rhs, identifierMap)
					case "ElevatedSurface":
						__gong__assignPointer(&inst.ElevatedSurface, rhs, identifierMap)
					}
				case *aixmfeatures.SurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SrsName":
						inst.SrsName = GongExtractString(rhs)
					case "Patches":
						__gong__assignPointer(&inst.Patches, rhs, identifierMap)
					case "HorizontalAccuracy":
						__gong__assignPointer(&inst.HorizontalAccuracy, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					}
				case *aixmfeatures.TimesheetPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Timesheet":
						__gong__assignPointer(&inst.Timesheet, rhs, identifierMap)
					}
				case *aixmfeatures.TimesheetType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "TimeReference":
						__gong__assignPointer(&inst.TimeReference, rhs, identifierMap)
					case "StartDate":
						__gong__assignPointer(&inst.StartDate, rhs, identifierMap)
					case "EndDate":
						__gong__assignPointer(&inst.EndDate, rhs, identifierMap)
					case "Day":
						__gong__assignPointer(&inst.Day, rhs, identifierMap)
					case "DayTil":
						__gong__assignPointer(&inst.DayTil, rhs, identifierMap)
					case "StartTime":
						__gong__assignPointer(&inst.StartTime, rhs, identifierMap)
					case "StartEvent":
						__gong__assignPointer(&inst.StartEvent, rhs, identifierMap)
					case "StartTimeRelativeEvent":
						__gong__assignPointer(&inst.StartTimeRelativeEvent, rhs, identifierMap)
					case "StartEventInterpretation":
						__gong__assignPointer(&inst.StartEventInterpretation, rhs, identifierMap)
					case "EndTime":
						__gong__assignPointer(&inst.EndTime, rhs, identifierMap)
					case "EndEvent":
						__gong__assignPointer(&inst.EndEvent, rhs, identifierMap)
					case "EndTimeRelativeEvent":
						__gong__assignPointer(&inst.EndTimeRelativeEvent, rhs, identifierMap)
					case "EndEventInterpretation":
						__gong__assignPointer(&inst.EndEventInterpretation, rhs, identifierMap)
					case "DaylightSavingAdjust":
						__gong__assignPointer(&inst.DaylightSavingAdjust, rhs, identifierMap)
					case "Excluded":
						__gong__assignPointer(&inst.Excluded, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.UnitPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *aixmfeatures.VerticalStructureLightingStatusPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "VerticalStructureLightingStatus":
						__gong__assignPointer(&inst.VerticalStructureLightingStatus, rhs, identifierMap)
					}
				case *aixmfeatures.VerticalStructureLightingStatusType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "Status":
						__gong__assignPointer(&inst.Status, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.VerticalStructurePartPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "VerticalStructurePart":
						__gong__assignPointer(&inst.VerticalStructurePart, rhs, identifierMap)
					}
				case *aixmfeatures.VerticalStructurePartType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "TimeInterval":
						__gong__assignSliceOfPointers(&inst.TimeInterval, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SpecialDateAuthority":
						__gong__assignSliceOfPointers(&inst.SpecialDateAuthority, rhs, identifierMap)
					case "VerticalExtent":
						__gong__assignPointer(&inst.VerticalExtent, rhs, identifierMap)
					case "VerticalExtentAccuracy":
						__gong__assignPointer(&inst.VerticalExtentAccuracy, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "ConstructionStatus":
						__gong__assignPointer(&inst.ConstructionStatus, rhs, identifierMap)
					case "MarkingPattern":
						__gong__assignPointer(&inst.MarkingPattern, rhs, identifierMap)
					case "MarkingFirstColour":
						__gong__assignPointer(&inst.MarkingFirstColour, rhs, identifierMap)
					case "MarkingSecondColour":
						__gong__assignPointer(&inst.MarkingSecondColour, rhs, identifierMap)
					case "Mobile":
						__gong__assignPointer(&inst.Mobile, rhs, identifierMap)
					case "Frangible":
						__gong__assignPointer(&inst.Frangible, rhs, identifierMap)
					case "VisibleMaterial":
						__gong__assignPointer(&inst.VisibleMaterial, rhs, identifierMap)
					case "Designator":
						__gong__assignPointer(&inst.Designator, rhs, identifierMap)
					case "HorizontalProjection_surfaceExtent":
						__gong__assignPointer(&inst.HorizontalProjection_surfaceExtent, rhs, identifierMap)
					case "HorizontalProjection_linearExtent":
						__gong__assignPointer(&inst.HorizontalProjection_linearExtent, rhs, identifierMap)
					case "HorizontalProjection_location":
						__gong__assignPointer(&inst.HorizontalProjection_location, rhs, identifierMap)
					case "Lighting":
						__gong__assignSliceOfPointers(&inst.Lighting, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.VerticalStructurePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					}
				case *aixmfeatures.VerticalStructureTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "VerticalStructureTimeSlice":
						__gong__assignPointer(&inst.VerticalStructureTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.VerticalStructureTimeSliceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "ValidTime":
						__gong__assignPointer(&inst.ValidTime, rhs, identifierMap)
					case "Interpretation":
						inst.Interpretation = GongExtractString(rhs)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "CorrectionNumber":
						inst.CorrectionNumber = GongExtractInt(rhs)
					case "TimeSliceMetadata":
						__gong__assignPointer(&inst.TimeSliceMetadata, rhs, identifierMap)
					case "FeatureLifetime":
						__gong__assignPointer(&inst.FeatureLifetime, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Lighted":
						__gong__assignPointer(&inst.Lighted, rhs, identifierMap)
					case "MarkingICAOStandard":
						__gong__assignPointer(&inst.MarkingICAOStandard, rhs, identifierMap)
					case "Group":
						__gong__assignPointer(&inst.Group, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Width":
						__gong__assignPointer(&inst.Width, rhs, identifierMap)
					case "Radius":
						__gong__assignPointer(&inst.Radius, rhs, identifierMap)
					case "LightingICAOStandard":
						__gong__assignPointer(&inst.LightingICAOStandard, rhs, identifierMap)
					case "SynchronisedLighting":
						__gong__assignPointer(&inst.SynchronisedLighting, rhs, identifierMap)
					case "Marker":
						__gong__assignPointer(&inst.Marker, rhs, identifierMap)
					case "Part":
						__gong__assignSliceOfPointers(&inst.Part, rhs, identifierMap)
					case "HostedPassengerService":
						__gong__assignSliceOfPointers(&inst.HostedPassengerService, rhs, identifierMap)
					case "SupportedGroundLight":
						__gong__assignSliceOfPointers(&inst.SupportedGroundLight, rhs, identifierMap)
					case "HostedNavaidEquipment":
						__gong__assignSliceOfPointers(&inst.HostedNavaidEquipment, rhs, identifierMap)
					case "HostedSpecialNavStation":
						__gong__assignSliceOfPointers(&inst.HostedSpecialNavStation, rhs, identifierMap)
					case "HostedUnit":
						__gong__assignSliceOfPointers(&inst.HostedUnit, rhs, identifierMap)
					case "HostedOrganisation":
						__gong__assignSliceOfPointers(&inst.HostedOrganisation, rhs, identifierMap)
					case "SupportedService":
						__gong__assignSliceOfPointers(&inst.SupportedService, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "LightingAvailability":
						__gong__assignSliceOfPointers(&inst.LightingAvailability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.VerticalStructureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BoundedBy":
						__gong__assignPointer(&inst.BoundedBy, rhs, identifierMap)
					case "FeatureMetadata":
						__gong__assignPointer(&inst.FeatureMetadata, rhs, identifierMap)
					case "TimeSlice":
						__gong__assignSliceOfPointers(&inst.TimeSlice, rhs, identifierMap)
					}
				case *AIXMBasicMessage:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BoundedBy":
						__gong__assignPointer(&inst.BoundedBy, rhs, identifierMap)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "MessageMetadata":
						__gong__assignPointer(&inst.MessageMetadata, rhs, identifierMap)
					case "AggregationType":
						inst.AggregationType = GongExtractString(rhs)
					case "HasMember":
						__gong__assignSliceOfPointers(&inst.HasMember, rhs, identifierMap)
					}
				case *AIXMBasicMessageType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BoundedBy":
						__gong__assignPointer(&inst.BoundedBy, rhs, identifierMap)
					case "SequenceNumber":
						inst.SequenceNumber = GongExtractInt(rhs)
					case "MessageMetadata":
						__gong__assignPointer(&inst.MessageMetadata, rhs, identifierMap)
					case "AggregationType":
						inst.AggregationType = GongExtractString(rhs)
					case "HasMember":
						__gong__assignSliceOfPointers(&inst.HasMember, rhs, identifierMap)
					}
				case *BasicMessageMemberAIXMPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Type_XSD":
						inst.Type_XSD = GongExtractString(rhs)
					case "Href":
						inst.Href = GongExtractString(rhs)
					case "Role":
						inst.Role = GongExtractString(rhs)
					case "Arcrole":
						inst.Arcrole = GongExtractString(rhs)
					case "Title":
						inst.Title = GongExtractString(rhs)
					case "Show":
						inst.Show = GongExtractString(rhs)
					case "Actuate":
						inst.Actuate = GongExtractString(rhs)
					case "VerticalStructure":
						__gong__assignPointer(&inst.VerticalStructure, rhs, identifierMap)
					case "ObstacleArea":
						__gong__assignPointer(&inst.ObstacleArea, rhs, identifierMap)
					case "AeronauticalGroundLight":
						__gong__assignPointer(&inst.AeronauticalGroundLight, rhs, identifierMap)
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
								if inst, ok := instance.(xlink.GongstructIF); ok && stage.XlinkStage != nil {
									inst.UnstageVoid(stage.XlinkStage)
								}
								if inst, ok := instance.(gml.GongstructIF); ok && stage.GmlStage != nil {
									inst.UnstageVoid(stage.GmlStage)
								}
								if inst, ok := instance.(aixmabstractgmlobjecttypes.GongstructIF); ok && stage.AixmabstractgmlobjecttypesStage != nil {
									inst.UnstageVoid(stage.AixmabstractgmlobjecttypesStage)
								}
								if inst, ok := instance.(aixmdatatypes.GongstructIF); ok && stage.AixmdatatypesStage != nil {
									inst.UnstageVoid(stage.AixmdatatypesStage)
								}
								if inst, ok := instance.(aixmfeatures.GongstructIF); ok && stage.AixmfeaturesStage != nil {
									inst.UnstageVoid(stage.AixmfeaturesStage)
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
type AIXMBasicMessageUnmarshaller struct{}

func (u *AIXMBasicMessageUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AIXMBasicMessage), stage, identifier, instanceName, preserveOrder)
}

func (u *AIXMBasicMessageUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AIXMBasicMessage)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "AggregationType":
		instance.AggregationType = GongExtractString(valueExpr)
	case "HasMember":
		GongUnmarshallSliceOfPointers(&instance.HasMember, valueExpr, identifierMap)
	}
	return nil
}

type AIXMBasicMessageTypeUnmarshaller struct{}

func (u *AIXMBasicMessageTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(AIXMBasicMessageType), stage, identifier, instanceName, preserveOrder)
}

func (u *AIXMBasicMessageTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*AIXMBasicMessageType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Id":
		instance.Id = GongExtractString(valueExpr)
	case "Description":
		instance.Description = GongExtractString(valueExpr)
	case "SequenceNumber":
		instance.SequenceNumber = GongExtractInt(valueExpr)
	case "AggregationType":
		instance.AggregationType = GongExtractString(valueExpr)
	case "HasMember":
		GongUnmarshallSliceOfPointers(&instance.HasMember, valueExpr, identifierMap)
	}
	return nil
}

type BasicMessageMemberAIXMPropertyTypeUnmarshaller struct{}

func (u *BasicMessageMemberAIXMPropertyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(BasicMessageMemberAIXMPropertyType), stage, identifier, instanceName, preserveOrder)
}

func (u *BasicMessageMemberAIXMPropertyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*BasicMessageMemberAIXMPropertyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Type_XSD":
		instance.Type_XSD = GongExtractString(valueExpr)
	case "Href":
		instance.Href = GongExtractString(valueExpr)
	case "Role":
		instance.Role = GongExtractString(valueExpr)
	case "Arcrole":
		instance.Arcrole = GongExtractString(valueExpr)
	case "Title":
		instance.Title = GongExtractString(valueExpr)
	case "Show":
		instance.Show = GongExtractString(valueExpr)
	case "Actuate":
		instance.Actuate = GongExtractString(valueExpr)
	}
	return nil
}
