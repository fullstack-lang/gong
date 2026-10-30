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
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/xlink"
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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/xlink":
			aliasToCanonical[alias] = "xlink"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/gml":
			aliasToCanonical[alias] = "gml"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmabstractgmlobjecttypes":
			aliasToCanonical[alias] = "aixmabstractgmlobjecttypes"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmdatatypes":
			aliasToCanonical[alias] = "aixmdatatypes"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models/aixmfeatures":
			aliasToCanonical[alias] = "aixmfeatures"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/route/go/models":
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
				case "CodeATCReportingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeATCReportingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAerialRefuellingPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAerialRefuellingPointType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAerialRefuellingPrefixType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAerialRefuellingPrefixType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAerialRefuellingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAerialRefuellingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftCategoryType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftCategoryType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftEngineNumberType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftEngineNumberType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftEngineType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftEngineType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftICAOType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftICAOType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftWingspanClassType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftWingspanClassType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAltitudeAdjustmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAltitudeAdjustmentType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAltitudeUseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAltitudeUseType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAuthorityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAuthorityType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeCardinalDirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeCardinalDirectionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeCommunicationModeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeCommunicationModeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeComparisonType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeComparisonType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeCourseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeCourseType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeDayType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeDirectionTurnType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeDirectionTurnType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeDirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeDirectionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeEquipmentAntiCollisionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeEquipmentAntiCollisionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightOriginType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFlightOriginType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightPurposeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFlightPurposeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightRestrictionDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFlightRestrictionDesignatorType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightRestrictionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFlightRestrictionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightRuleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFlightRuleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFlightStatusType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFlightType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFlowConditionOperationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFlowConditionOperationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFreeFlightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFreeFlightType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLevelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLevelType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLocationQualifierType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLocationQualifierType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMeteoConditionsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMeteoConditionsType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryRoutePointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMilitaryRoutePointType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMilitaryStatusType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryTrainingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMilitaryTrainingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeNavigationEquipmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeNavigationEquipmentType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeNavigationSpecificationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeNavigationSpecificationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeNotePurposeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeNotePurposeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeObstacleAssessmentSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeObstacleAssessmentSurfaceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeObstructionIdSurfaceZoneType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeObstructionIdSurfaceZoneType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRNPType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRNPType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRVSMPointRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRVSMPointRoleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRVSMType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRVSMType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeReferenceRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeReferenceRoleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRouteAvailabilityType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteDesignatorLetterType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRouteDesignatorLetterType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteDesignatorPrefixType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRouteDesignatorPrefixType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteDesignatorSuffixType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRouteDesignatorSuffixType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteNavigationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRouteNavigationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteOriginType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRouteOriginType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteSegmentPathType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRouteSegmentPathType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRouteType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRuleProcedureTitleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRuleProcedureTitleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRuleProcedureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRuleProcedureType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeSideType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeSideType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeSpeedReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeSpeedReferenceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTACANChannelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTACANChannelType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTelecomNetworkType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTelecomNetworkType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventCombinationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTimeEventCombinationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTimeEventType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTimeReferenceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTransponderType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTransponderType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeUpperAlphaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeUpperAlphaType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeValueInterpretationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeValueInterpretationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalDatumType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeVerticalDatumType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeVerticalReferenceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeWakeTurbulenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeWakeTurbulenceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeYesNoType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeYesNoType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "DateMonthDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.DateMonthDayType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "NoNumberType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.NoNumberType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "NoSequenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.NoSequenceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextAddressType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextAddressType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextDesignatorType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextInstructionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextInstructionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextNameType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextNameType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextNoteType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextPhoneType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextPhoneType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextPropertyNameType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextPropertyNameType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TimeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TimeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValAngleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValAngleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValBearingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValBearingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceSignedType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDistanceSignedType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDistanceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceVerticalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDistanceVerticalType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDurationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDurationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValSlopeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValSlopeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValSpeedType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValSpeedType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValWeightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValWeightType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "XHTMLType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.XHTMLType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
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
				case "A_extension_13":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_13), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_14":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_14), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_15":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_15), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_16":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_16), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_17":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_17), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_18":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_18), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_19":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_19), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_2":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_2), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_20":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_20), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_21":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_21), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_22":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_22), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_23":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_23), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_24":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_24), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_25":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_25), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_26":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_26), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_27":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_27), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_28":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_28), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_29":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_29), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_3":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_3), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_30":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_30), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_31":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_31), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_32":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_32), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_33":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_33), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_34":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_34), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_35":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_35), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_36":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_36), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_37":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_37), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_38":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_38), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_39":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_39), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_4":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_4), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_40":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_40), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_41":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_41), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_42":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_42), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
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
				case "AbstractDirectFlightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractDirectFlightType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractPropertiesWithScheduleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractPropertiesWithScheduleType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractSegmentPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractSegmentPointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingAnchorPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AerialRefuellingAnchorPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingAnchorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AerialRefuellingAnchorType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AerialRefuellingPointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AerialRefuellingPointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AerialRefuellingPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AerialRefuellingTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AerialRefuellingTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingTrackPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AerialRefuellingTrackPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingTrackType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AerialRefuellingTrackType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AerialRefuellingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftCharacteristicPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftCharacteristicType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirspaceBorderCrossingPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirspaceBorderCrossingPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirspaceBorderCrossingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirspaceBorderCrossingTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirspaceBorderCrossingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirspaceBorderCrossingTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirspaceBorderCrossingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirspaceBorderCrossingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirspaceLayerPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirspaceLayerPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirspaceLayerType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirspaceLayerType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirspacePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirspacePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AltitudeAdjustmentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AltitudeAdjustmentPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AltitudeAdjustmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AltitudeAdjustmentType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AngleIndicationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AngleIndicationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AngleUsePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AngleUsePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AngleUseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AngleUseType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AuthorityForAerialRefuellingPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AuthorityForAerialRefuellingPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AuthorityForAerialRefuellingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AuthorityForAerialRefuellingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ChangeOverPointTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ChangeOverPointTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ChangeOverPointTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ChangeOverPointTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ChangeOverPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ChangeOverPointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ContactInformationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ContactInformationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ContactInformationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ContactInformationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CurvePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CurvePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CurveType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DMEPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DMEPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DesignatedPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DesignatedPointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DirectFlightClassType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DirectFlightClassType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DirectFlightPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DirectFlightPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DirectFlightSegmentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DirectFlightSegmentPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DirectFlightSegmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DirectFlightSegmentType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DistanceIndicationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DistanceIndicationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedCurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedCurveType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedPointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedSurfaceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "EnRouteSegmentPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.EnRouteSegmentPointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "EnRouteSegmentPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.EnRouteSegmentPointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightCharacteristicPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightCharacteristicPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightCharacteristicType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightCharacteristicType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightConditionCircumstancePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightConditionCircumstancePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightConditionCircumstanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightConditionCircumstanceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightConditionCombinationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightConditionCombinationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightConditionCombinationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightConditionCombinationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightConditionElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightConditionElementPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightConditionElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightConditionElementType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionLevelPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightRestrictionLevelPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionLevelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightRestrictionLevelType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionRoutePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightRestrictionRoutePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionRouteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightRestrictionRouteType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightRestrictionTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightRestrictionTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightRestrictionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightRestrictionType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightRoutingElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightRoutingElementPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightRoutingElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightRoutingElementType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LinguisticNotePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LinguisticNoteType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MeteorologyPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MeteorologyPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MeteorologyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MeteorologyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavaidPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavaidPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NotePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NoteType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ObstacleAssessmentAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ObstacleAssessmentAreaPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ObstacleAssessmentAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ObstacleAssessmentAreaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ObstaclePlacementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ObstaclePlacementPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ObstaclePlacementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ObstaclePlacementType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ObstructionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ObstructionPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ObstructionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ObstructionType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OnlineContactPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OnlineContactPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OnlineContactType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OnlineContactType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OrganisationAuthorityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OrganisationAuthorityPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PointReferencePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PointReferencePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PointReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PointReferenceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PostalAddressPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PostalAddressPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PostalAddressType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PostalAddressType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RadioFrequencyAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RadioFrequencyAreaPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RouteAvailabilityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RouteAvailabilityPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RouteAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RouteAvailabilityType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RouteDMETimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RouteDMETimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RouteDMETimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RouteDMETimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RouteDMEType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RouteDMEType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RoutePortionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RoutePortionPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RoutePortionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RoutePortionType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RoutePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RoutePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RouteSegmentTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RouteSegmentTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RouteSegmentTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RouteSegmentTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RouteSegmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RouteSegmentType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RouteTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RouteTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RouteTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RouteTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RouteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RouteType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RulesProceduresTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RulesProceduresTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RulesProceduresTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RulesProceduresTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RulesProceduresType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RulesProceduresType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayCentrelinePointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentArrivalPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardInstrumentArrivalPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentDeparturePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardInstrumentDeparturePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelColumnPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelColumnPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurfacePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurfacePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurfaceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TelephoneContactPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TelephoneContactPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TelephoneContactType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TelephoneContactType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TimesheetPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TimesheetPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TimesheetType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TimesheetType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "VerticalStructurePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.VerticalStructurePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
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
				case *aixmdatatypes.CodeATCReportingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAerialRefuellingPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAerialRefuellingPrefixType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAerialRefuellingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAircraftCategoryType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAircraftEngineNumberType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAircraftEngineType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAircraftICAOType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAircraftType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAircraftWingspanClassType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAltitudeAdjustmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAltitudeUseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAuthorityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeCardinalDirectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeCommunicationModeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeComparisonType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeCourseType:
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
				case *aixmdatatypes.CodeDirectionTurnType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeDirectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeEquipmentAntiCollisionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFlightOriginType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFlightPurposeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFlightRestrictionDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFlightRestrictionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFlightRuleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFlightStatusType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFlightType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFlowConditionOperationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFreeFlightType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeLevelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeLocationQualifierType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeMeteoConditionsType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeMilitaryRoutePointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeMilitaryStatusType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeMilitaryTrainingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeNavigationEquipmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeNavigationSpecificationType:
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
				case *aixmdatatypes.CodeObstacleAssessmentSurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeObstructionIdSurfaceZoneType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRNPType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRVSMPointRoleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRVSMType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeReferenceRoleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRouteAvailabilityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRouteDesignatorLetterType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRouteDesignatorPrefixType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRouteDesignatorSuffixType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRouteNavigationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRouteOriginType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRouteSegmentPathType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRouteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRuleProcedureTitleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRuleProcedureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeSideType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeSpeedReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeTACANChannelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeTelecomNetworkType:
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
				case *aixmdatatypes.CodeTransponderType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeUpperAlphaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeValueInterpretationType:
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
				case *aixmdatatypes.CodeVerticalReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeWakeTurbulenceType:
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
				case *aixmdatatypes.NoNumberType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *aixmdatatypes.NoSequenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *aixmdatatypes.TextAddressType:
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
				case *aixmdatatypes.TextInstructionType:
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
				case *aixmdatatypes.TextPhoneType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
				case *aixmdatatypes.ValAngleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.ValBearingType:
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
				case *aixmdatatypes.ValSlopeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.ValSpeedType:
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
				case *aixmdatatypes.ValWeightType:
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
				case *aixmdatatypes.XHTMLType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					}
				case *aixmfeatures.A_extension:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAerialRefuellingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAerialRefuellingExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAerialRefuellingAnchorExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAerialRefuellingAnchorExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_10:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPointReferenceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPointReferenceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_11:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractLinguisticNoteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractLinguisticNoteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_12:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractNoteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractNoteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_13:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractChangeOverPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractChangeOverPointExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_14:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRouteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRouteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_15:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractRouteAvailabilityExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRouteAvailabilityExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_16:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRouteDMEExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRouteDMEExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_17:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRoutePortionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRoutePortionExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_18:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRouteSegmentExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRouteSegmentExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_19:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAirspaceBorderCrossingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirspaceBorderCrossingExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_2:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSegmentPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSegmentPointExtension, rhs, identifierMap)
					case "AbstractAerialRefuellingPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAerialRefuellingPointExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_20:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractDirectFlightExtension":
						__gong__assignSliceOfPointers(&inst.AbstractDirectFlightExtension, rhs, identifierMap)
					case "AbstractDirectFlightClassExtension":
						__gong__assignSliceOfPointers(&inst.AbstractDirectFlightClassExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_21:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractDirectFlightExtension":
						__gong__assignSliceOfPointers(&inst.AbstractDirectFlightExtension, rhs, identifierMap)
					case "AbstractDirectFlightSegmentExtension":
						__gong__assignSliceOfPointers(&inst.AbstractDirectFlightSegmentExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_22:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFlightConditionCircumstanceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFlightConditionCircumstanceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_23:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractFlightConditionCombinationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFlightConditionCombinationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_24:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFlightConditionElementExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFlightConditionElementExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_25:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFlightRestrictionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFlightRestrictionExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_26:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFlightRestrictionLevelExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFlightRestrictionLevelExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_27:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFlightRestrictionRouteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFlightRestrictionRouteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_28:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFlightRoutingElementExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFlightRoutingElementExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_29:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRulesProceduresExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRulesProceduresExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_3:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAerialRefuellingTrackExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAerialRefuellingTrackExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_30:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractContactInformationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractContactInformationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_31:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractOnlineContactExtension":
						__gong__assignSliceOfPointers(&inst.AbstractOnlineContactExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_32:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractPostalAddressExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPostalAddressExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_33:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractTelephoneContactExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTelephoneContactExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_34:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAircraftCharacteristicExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAircraftCharacteristicExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_35:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFlightCharacteristicExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFlightCharacteristicExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_36:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAirspaceLayerExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirspaceLayerExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_37:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractMeteorologyExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMeteorologyExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_38:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTimesheetExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTimesheetExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_39:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAltitudeAdjustmentExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAltitudeAdjustmentExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_4:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAuthorityForAerialRefuellingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAuthorityForAerialRefuellingExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_40:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractObstacleAssessmentAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractObstacleAssessmentAreaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_41:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractObstaclePlacementExtension":
						__gong__assignSliceOfPointers(&inst.AbstractObstaclePlacementExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_42:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractObstructionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractObstructionExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_5:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractElevatedCurveExtension":
						__gong__assignSliceOfPointers(&inst.AbstractElevatedCurveExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_6:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractElevatedPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractElevatedPointExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_7:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractElevatedSurfaceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractElevatedSurfaceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_8:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAngleUseExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAngleUseExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_9:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSegmentPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSegmentPointExtension, rhs, identifierMap)
					case "AbstractEnRouteSegmentPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractEnRouteSegmentPointExtension, rhs, identifierMap)
					}
				case *aixmfeatures.AbstractDirectFlightType:
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
				case *aixmfeatures.AbstractSegmentPointType:
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
				case *aixmfeatures.AerialRefuellingAnchorPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AerialRefuellingAnchor":
						__gong__assignPointer(&inst.AerialRefuellingAnchor, rhs, identifierMap)
					}
				case *aixmfeatures.AerialRefuellingAnchorType:
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
					case "OutboundCourse":
						__gong__assignPointer(&inst.OutboundCourse, rhs, identifierMap)
					case "OutboundCourseType":
						__gong__assignPointer(&inst.OutboundCourseType, rhs, identifierMap)
					case "InboundCourse":
						__gong__assignPointer(&inst.InboundCourse, rhs, identifierMap)
					case "TurnDirection":
						__gong__assignPointer(&inst.TurnDirection, rhs, identifierMap)
					case "SpeedLimit":
						__gong__assignPointer(&inst.SpeedLimit, rhs, identifierMap)
					case "LegSeparation":
						__gong__assignPointer(&inst.LegSeparation, rhs, identifierMap)
					case "LegLength":
						__gong__assignPointer(&inst.LegLength, rhs, identifierMap)
					case "RefuellingBaseLevel":
						__gong__assignPointer(&inst.RefuellingBaseLevel, rhs, identifierMap)
					case "RefuellingBaseLevelReference":
						__gong__assignPointer(&inst.RefuellingBaseLevelReference, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "VerticalExtent":
						__gong__assignSliceOfPointers(&inst.VerticalExtent, rhs, identifierMap)
					case "Point":
						__gong__assignSliceOfPointers(&inst.Point, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AerialRefuellingPointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AerialRefuellingPoint":
						__gong__assignPointer(&inst.AerialRefuellingPoint, rhs, identifierMap)
					}
				case *aixmfeatures.AerialRefuellingPointType:
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
					case "ReportingATC":
						__gong__assignPointer(&inst.ReportingATC, rhs, identifierMap)
					case "FlyOver":
						__gong__assignPointer(&inst.FlyOver, rhs, identifierMap)
					case "Waypoint":
						__gong__assignPointer(&inst.Waypoint, rhs, identifierMap)
					case "RadarGuidance":
						__gong__assignPointer(&inst.RadarGuidance, rhs, identifierMap)
					case "FacilityMakeup":
						__gong__assignSliceOfPointers(&inst.FacilityMakeup, rhs, identifierMap)
					case "PointChoice_fixDesignatedPoint":
						__gong__assignPointer(&inst.PointChoice_fixDesignatedPoint, rhs, identifierMap)
					case "PointChoice_navaidSystem":
						__gong__assignPointer(&inst.PointChoice_navaidSystem, rhs, identifierMap)
					case "PointChoice_position":
						__gong__assignPointer(&inst.PointChoice_position, rhs, identifierMap)
					case "PointChoice_runwayPoint":
						__gong__assignPointer(&inst.PointChoice_runwayPoint, rhs, identifierMap)
					case "PointChoice_aimingPoint":
						__gong__assignPointer(&inst.PointChoice_aimingPoint, rhs, identifierMap)
					case "PointChoice_airportReferencePoint":
						__gong__assignPointer(&inst.PointChoice_airportReferencePoint, rhs, identifierMap)
					case "ExtendedServiceVolume":
						__gong__assignPointer(&inst.ExtendedServiceVolume, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Sequence":
						__gong__assignPointer(&inst.Sequence, rhs, identifierMap)
					case "UsageType":
						__gong__assignPointer(&inst.UsageType, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AerialRefuellingPropertyType:
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
				case *aixmfeatures.AerialRefuellingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AerialRefuellingTimeSlice":
						__gong__assignPointer(&inst.AerialRefuellingTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AerialRefuellingTimeSliceType:
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
					case "DesignatorPrefix":
						__gong__assignPointer(&inst.DesignatorPrefix, rhs, identifierMap)
					case "DesignatorNumber":
						__gong__assignPointer(&inst.DesignatorNumber, rhs, identifierMap)
					case "DesignatorSuffix":
						__gong__assignPointer(&inst.DesignatorSuffix, rhs, identifierMap)
					case "DesignatorDirection":
						__gong__assignPointer(&inst.DesignatorDirection, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "RadarBeaconSetting":
						__gong__assignPointer(&inst.RadarBeaconSetting, rhs, identifierMap)
					case "XbandRadarSetting":
						__gong__assignPointer(&inst.XbandRadarSetting, rhs, identifierMap)
					case "TankerChannel":
						__gong__assignPointer(&inst.TankerChannel, rhs, identifierMap)
					case "ReceiverChannel":
						__gong__assignPointer(&inst.ReceiverChannel, rhs, identifierMap)
					case "HelicopterRoute":
						__gong__assignPointer(&inst.HelicopterRoute, rhs, identifierMap)
					case "SpecialRefuelling":
						__gong__assignPointer(&inst.SpecialRefuelling, rhs, identifierMap)
					case "BidirectionalUse":
						__gong__assignPointer(&inst.BidirectionalUse, rhs, identifierMap)
					case "ReverseDirectionTurn":
						__gong__assignPointer(&inst.ReverseDirectionTurn, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "ProtectingAirspace":
						__gong__assignSliceOfPointers(&inst.ProtectingAirspace, rhs, identifierMap)
					case "Track":
						__gong__assignSliceOfPointers(&inst.Track, rhs, identifierMap)
					case "Anchor":
						__gong__assignSliceOfPointers(&inst.Anchor, rhs, identifierMap)
					case "OppositeTrack":
						__gong__assignPointer(&inst.OppositeTrack, rhs, identifierMap)
					case "ManagingOrganisation":
						__gong__assignSliceOfPointers(&inst.ManagingOrganisation, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AerialRefuellingTrackPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AerialRefuellingTrack":
						__gong__assignPointer(&inst.AerialRefuellingTrack, rhs, identifierMap)
					}
				case *aixmfeatures.AerialRefuellingTrackType:
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
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Point":
						__gong__assignSliceOfPointers(&inst.Point, rhs, identifierMap)
					case "VerticalExtent":
						__gong__assignSliceOfPointers(&inst.VerticalExtent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AerialRefuellingType:
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
				case *aixmfeatures.AircraftCharacteristicPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AircraftCharacteristic":
						__gong__assignPointer(&inst.AircraftCharacteristic, rhs, identifierMap)
					}
				case *aixmfeatures.AircraftCharacteristicType:
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
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Engine":
						__gong__assignPointer(&inst.Engine, rhs, identifierMap)
					case "NumberEngine":
						__gong__assignPointer(&inst.NumberEngine, rhs, identifierMap)
					case "TypeAircraftICAO":
						__gong__assignPointer(&inst.TypeAircraftICAO, rhs, identifierMap)
					case "AircraftLandingCategory":
						__gong__assignPointer(&inst.AircraftLandingCategory, rhs, identifierMap)
					case "WingSpan":
						__gong__assignPointer(&inst.WingSpan, rhs, identifierMap)
					case "WingSpanInterpretation":
						__gong__assignPointer(&inst.WingSpanInterpretation, rhs, identifierMap)
					case "ClassWingSpan":
						__gong__assignPointer(&inst.ClassWingSpan, rhs, identifierMap)
					case "Weight":
						__gong__assignPointer(&inst.Weight, rhs, identifierMap)
					case "WeightInterpretation":
						__gong__assignPointer(&inst.WeightInterpretation, rhs, identifierMap)
					case "Passengers":
						__gong__assignPointer(&inst.Passengers, rhs, identifierMap)
					case "PassengersInterpretation":
						__gong__assignPointer(&inst.PassengersInterpretation, rhs, identifierMap)
					case "Speed":
						__gong__assignPointer(&inst.Speed, rhs, identifierMap)
					case "SpeedInterpretation":
						__gong__assignPointer(&inst.SpeedInterpretation, rhs, identifierMap)
					case "WakeTurbulence":
						__gong__assignPointer(&inst.WakeTurbulence, rhs, identifierMap)
					case "NavigationEquipment":
						__gong__assignPointer(&inst.NavigationEquipment, rhs, identifierMap)
					case "NavigationSpecification":
						__gong__assignPointer(&inst.NavigationSpecification, rhs, identifierMap)
					case "VerticalSeparationCapability":
						__gong__assignPointer(&inst.VerticalSeparationCapability, rhs, identifierMap)
					case "AntiCollisionAndSeparationEquipment":
						__gong__assignPointer(&inst.AntiCollisionAndSeparationEquipment, rhs, identifierMap)
					case "CommunicationEquipment":
						__gong__assignPointer(&inst.CommunicationEquipment, rhs, identifierMap)
					case "SurveillanceEquipment":
						__gong__assignPointer(&inst.SurveillanceEquipment, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *aixmfeatures.AirspaceBorderCrossingPropertyType:
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
				case *aixmfeatures.AirspaceBorderCrossingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirspaceBorderCrossingTimeSlice":
						__gong__assignPointer(&inst.AirspaceBorderCrossingTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AirspaceBorderCrossingTimeSliceType:
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
					case "ExitedAirspace":
						__gong__assignPointer(&inst.ExitedAirspace, rhs, identifierMap)
					case "EnteredAirspace":
						__gong__assignPointer(&inst.EnteredAirspace, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AirspaceBorderCrossingType:
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
				case *aixmfeatures.AirspaceLayerPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirspaceLayer":
						__gong__assignPointer(&inst.AirspaceLayer, rhs, identifierMap)
					}
				case *aixmfeatures.AirspaceLayerType:
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
					case "UpperLimit":
						__gong__assignPointer(&inst.UpperLimit, rhs, identifierMap)
					case "UpperLimitReference":
						__gong__assignPointer(&inst.UpperLimitReference, rhs, identifierMap)
					case "LowerLimit":
						__gong__assignPointer(&inst.LowerLimit, rhs, identifierMap)
					case "LowerLimitReference":
						__gong__assignPointer(&inst.LowerLimitReference, rhs, identifierMap)
					case "AltitudeInterpretation":
						__gong__assignPointer(&inst.AltitudeInterpretation, rhs, identifierMap)
					case "DiscreteLevelSeries":
						__gong__assignPointer(&inst.DiscreteLevelSeries, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AirspacePropertyType:
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
				case *aixmfeatures.AltitudeAdjustmentPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AltitudeAdjustment":
						__gong__assignPointer(&inst.AltitudeAdjustment, rhs, identifierMap)
					}
				case *aixmfeatures.AltitudeAdjustmentType:
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
					case "AltitudeAdjustmentType":
						__gong__assignPointer(&inst.AltitudeAdjustmentType, rhs, identifierMap)
					case "PrimaryAlternateMinimum":
						__gong__assignPointer(&inst.PrimaryAlternateMinimum, rhs, identifierMap)
					case "AltitudeAdjustment":
						__gong__assignPointer(&inst.AltitudeAdjustment, rhs, identifierMap)
					case "LocalRemoteCode":
						__gong__assignPointer(&inst.LocalRemoteCode, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AngleIndicationPropertyType:
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
				case *aixmfeatures.AngleUsePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AngleUse":
						__gong__assignPointer(&inst.AngleUse, rhs, identifierMap)
					}
				case *aixmfeatures.AngleUseType:
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
					case "AlongCourseGuidance":
						__gong__assignPointer(&inst.AlongCourseGuidance, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheAngleIndication":
						__gong__assignPointer(&inst.TheAngleIndication, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AuthorityForAerialRefuellingPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AuthorityForAerialRefuelling":
						__gong__assignPointer(&inst.AuthorityForAerialRefuelling, rhs, identifierMap)
					}
				case *aixmfeatures.AuthorityForAerialRefuellingType:
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
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheOrganisationAuthority":
						__gong__assignPointer(&inst.TheOrganisationAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ChangeOverPointTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ChangeOverPointTimeSlice":
						__gong__assignPointer(&inst.ChangeOverPointTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.ChangeOverPointTimeSliceType:
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
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "Location_fixDesignatedPoint":
						__gong__assignPointer(&inst.Location_fixDesignatedPoint, rhs, identifierMap)
					case "Location_navaidSystem":
						__gong__assignPointer(&inst.Location_navaidSystem, rhs, identifierMap)
					case "Location_position":
						__gong__assignPointer(&inst.Location_position, rhs, identifierMap)
					case "Location_runwayPoint":
						__gong__assignPointer(&inst.Location_runwayPoint, rhs, identifierMap)
					case "Location_aimingPoint":
						__gong__assignPointer(&inst.Location_aimingPoint, rhs, identifierMap)
					case "Location_airportReferencePoint":
						__gong__assignPointer(&inst.Location_airportReferencePoint, rhs, identifierMap)
					case "ApplicableRoutePortion":
						__gong__assignPointer(&inst.ApplicableRoutePortion, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ChangeOverPointType:
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
				case *aixmfeatures.ContactInformationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ContactInformation":
						__gong__assignPointer(&inst.ContactInformation, rhs, identifierMap)
					}
				case *aixmfeatures.ContactInformationType:
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
					case "Title":
						__gong__assignPointer(&inst.Title, rhs, identifierMap)
					case "Address":
						__gong__assignSliceOfPointers(&inst.Address, rhs, identifierMap)
					case "NetworkNode":
						__gong__assignSliceOfPointers(&inst.NetworkNode, rhs, identifierMap)
					case "PhoneFax":
						__gong__assignSliceOfPointers(&inst.PhoneFax, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.CurvePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Curve":
						__gong__assignPointer(&inst.Curve, rhs, identifierMap)
					case "ElevatedCurve":
						__gong__assignPointer(&inst.ElevatedCurve, rhs, identifierMap)
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
				case *aixmfeatures.DMEPropertyType:
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
				case *aixmfeatures.DesignatedPointPropertyType:
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
				case *aixmfeatures.DirectFlightClassType:
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
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "ExceedLength":
						__gong__assignPointer(&inst.ExceedLength, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.DirectFlightPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "DirectFlightClass":
						__gong__assignPointer(&inst.DirectFlightClass, rhs, identifierMap)
					case "DirectFlightSegment":
						__gong__assignPointer(&inst.DirectFlightSegment, rhs, identifierMap)
					}
				case *aixmfeatures.DirectFlightSegmentPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "DirectFlightSegment":
						__gong__assignPointer(&inst.DirectFlightSegment, rhs, identifierMap)
					}
				case *aixmfeatures.DirectFlightSegmentType:
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
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "End_fixDesignatedPoint":
						__gong__assignPointer(&inst.End_fixDesignatedPoint, rhs, identifierMap)
					case "End_navaidSystem":
						__gong__assignPointer(&inst.End_navaidSystem, rhs, identifierMap)
					case "End_position":
						__gong__assignPointer(&inst.End_position, rhs, identifierMap)
					case "End_runwayPoint":
						__gong__assignPointer(&inst.End_runwayPoint, rhs, identifierMap)
					case "End_aimingPoint":
						__gong__assignPointer(&inst.End_aimingPoint, rhs, identifierMap)
					case "End_airportReferencePoint":
						__gong__assignPointer(&inst.End_airportReferencePoint, rhs, identifierMap)
					case "Start_fixDesignatedPoint":
						__gong__assignPointer(&inst.Start_fixDesignatedPoint, rhs, identifierMap)
					case "Start_navaidSystem":
						__gong__assignPointer(&inst.Start_navaidSystem, rhs, identifierMap)
					case "Start_position":
						__gong__assignPointer(&inst.Start_position, rhs, identifierMap)
					case "Start_runwayPoint":
						__gong__assignPointer(&inst.Start_runwayPoint, rhs, identifierMap)
					case "Start_aimingPoint":
						__gong__assignPointer(&inst.Start_aimingPoint, rhs, identifierMap)
					case "Start_airportReferencePoint":
						__gong__assignPointer(&inst.Start_airportReferencePoint, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.DistanceIndicationPropertyType:
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
				case *aixmfeatures.EnRouteSegmentPointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnRouteSegmentPoint":
						__gong__assignPointer(&inst.EnRouteSegmentPoint, rhs, identifierMap)
					}
				case *aixmfeatures.EnRouteSegmentPointType:
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
					case "ReportingATC":
						__gong__assignPointer(&inst.ReportingATC, rhs, identifierMap)
					case "FlyOver":
						__gong__assignPointer(&inst.FlyOver, rhs, identifierMap)
					case "Waypoint":
						__gong__assignPointer(&inst.Waypoint, rhs, identifierMap)
					case "RadarGuidance":
						__gong__assignPointer(&inst.RadarGuidance, rhs, identifierMap)
					case "FacilityMakeup":
						__gong__assignSliceOfPointers(&inst.FacilityMakeup, rhs, identifierMap)
					case "PointChoice_fixDesignatedPoint":
						__gong__assignPointer(&inst.PointChoice_fixDesignatedPoint, rhs, identifierMap)
					case "PointChoice_navaidSystem":
						__gong__assignPointer(&inst.PointChoice_navaidSystem, rhs, identifierMap)
					case "PointChoice_position":
						__gong__assignPointer(&inst.PointChoice_position, rhs, identifierMap)
					case "PointChoice_runwayPoint":
						__gong__assignPointer(&inst.PointChoice_runwayPoint, rhs, identifierMap)
					case "PointChoice_aimingPoint":
						__gong__assignPointer(&inst.PointChoice_aimingPoint, rhs, identifierMap)
					case "PointChoice_airportReferencePoint":
						__gong__assignPointer(&inst.PointChoice_airportReferencePoint, rhs, identifierMap)
					case "ExtendedServiceVolume":
						__gong__assignPointer(&inst.ExtendedServiceVolume, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "RoleFreeFlight":
						__gong__assignPointer(&inst.RoleFreeFlight, rhs, identifierMap)
					case "RoleRVSM":
						__gong__assignPointer(&inst.RoleRVSM, rhs, identifierMap)
					case "TurnRadius":
						__gong__assignPointer(&inst.TurnRadius, rhs, identifierMap)
					case "RoleMilitaryTraining":
						__gong__assignPointer(&inst.RoleMilitaryTraining, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.FlightCharacteristicPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightCharacteristic":
						__gong__assignPointer(&inst.FlightCharacteristic, rhs, identifierMap)
					}
				case *aixmfeatures.FlightCharacteristicType:
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
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Rule":
						__gong__assignPointer(&inst.Rule, rhs, identifierMap)
					case "Status":
						__gong__assignPointer(&inst.Status, rhs, identifierMap)
					case "Military":
						__gong__assignPointer(&inst.Military, rhs, identifierMap)
					case "Origin":
						__gong__assignPointer(&inst.Origin, rhs, identifierMap)
					case "Purpose":
						__gong__assignPointer(&inst.Purpose, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.FlightConditionCircumstancePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightConditionCircumstance":
						__gong__assignPointer(&inst.FlightConditionCircumstance, rhs, identifierMap)
					}
				case *aixmfeatures.FlightConditionCircumstanceType:
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
					case "ReferenceLocation":
						__gong__assignPointer(&inst.ReferenceLocation, rhs, identifierMap)
					case "RelationWithLocation":
						__gong__assignPointer(&inst.RelationWithLocation, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.FlightConditionCombinationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightConditionCombination":
						__gong__assignPointer(&inst.FlightConditionCombination, rhs, identifierMap)
					}
				case *aixmfeatures.FlightConditionCombinationType:
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
					case "LogicalOperator":
						__gong__assignPointer(&inst.LogicalOperator, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.FlightConditionElementPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightConditionElement":
						__gong__assignPointer(&inst.FlightConditionElement, rhs, identifierMap)
					}
				case *aixmfeatures.FlightConditionElementType:
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
					case "Index":
						__gong__assignPointer(&inst.Index, rhs, identifierMap)
					case "FlightCondition_weather":
						__gong__assignPointer(&inst.FlightCondition_weather, rhs, identifierMap)
					case "FlightCondition_aircraft":
						__gong__assignPointer(&inst.FlightCondition_aircraft, rhs, identifierMap)
					case "FlightCondition_flight":
						__gong__assignPointer(&inst.FlightCondition_flight, rhs, identifierMap)
					case "FlightCondition_directFlightCondition":
						__gong__assignPointer(&inst.FlightCondition_directFlightCondition, rhs, identifierMap)
					case "FlightCondition_borderCrossingCondition":
						__gong__assignPointer(&inst.FlightCondition_borderCrossingCondition, rhs, identifierMap)
					case "FlightCondition_operand":
						__gong__assignPointer(&inst.FlightCondition_operand, rhs, identifierMap)
					case "FlightCondition_routePortionCondition":
						__gong__assignPointer(&inst.FlightCondition_routePortionCondition, rhs, identifierMap)
					case "FlightCondition_standardInstrumentDepartureCondition":
						__gong__assignPointer(&inst.FlightCondition_standardInstrumentDepartureCondition, rhs, identifierMap)
					case "FlightCondition_standardInstrumentArrivalCondition":
						__gong__assignPointer(&inst.FlightCondition_standardInstrumentArrivalCondition, rhs, identifierMap)
					case "FlightCondition_organisationCondition":
						__gong__assignPointer(&inst.FlightCondition_organisationCondition, rhs, identifierMap)
					case "SignificantPointCondition_fixDesignatedPoint":
						__gong__assignPointer(&inst.SignificantPointCondition_fixDesignatedPoint, rhs, identifierMap)
					case "SignificantPointCondition_navaidSystem":
						__gong__assignPointer(&inst.SignificantPointCondition_navaidSystem, rhs, identifierMap)
					case "SignificantPointCondition_position":
						__gong__assignPointer(&inst.SignificantPointCondition_position, rhs, identifierMap)
					case "SignificantPointCondition_runwayPoint":
						__gong__assignPointer(&inst.SignificantPointCondition_runwayPoint, rhs, identifierMap)
					case "SignificantPointCondition_aimingPoint":
						__gong__assignPointer(&inst.SignificantPointCondition_aimingPoint, rhs, identifierMap)
					case "SignificantPointCondition_airportReferencePoint":
						__gong__assignPointer(&inst.SignificantPointCondition_airportReferencePoint, rhs, identifierMap)
					case "FlightCondition_airspaceCondition":
						__gong__assignPointer(&inst.FlightCondition_airspaceCondition, rhs, identifierMap)
					case "FlightCondition_airportHeliportCondition":
						__gong__assignPointer(&inst.FlightCondition_airportHeliportCondition, rhs, identifierMap)
					case "FlightCondition_aerialRefuellingCondition":
						__gong__assignPointer(&inst.FlightCondition_aerialRefuellingCondition, rhs, identifierMap)
					case "OperationalCondition":
						__gong__assignPointer(&inst.OperationalCondition, rhs, identifierMap)
					case "FlightLevel":
						__gong__assignSliceOfPointers(&inst.FlightLevel, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.FlightRestrictionLevelPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightRestrictionLevel":
						__gong__assignPointer(&inst.FlightRestrictionLevel, rhs, identifierMap)
					}
				case *aixmfeatures.FlightRestrictionLevelType:
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
					case "UpperLevel":
						__gong__assignPointer(&inst.UpperLevel, rhs, identifierMap)
					case "UpperLevelReference":
						__gong__assignPointer(&inst.UpperLevelReference, rhs, identifierMap)
					case "LowerLevel":
						__gong__assignPointer(&inst.LowerLevel, rhs, identifierMap)
					case "LowerLevelReference":
						__gong__assignPointer(&inst.LowerLevelReference, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.FlightRestrictionRoutePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightRestrictionRoute":
						__gong__assignPointer(&inst.FlightRestrictionRoute, rhs, identifierMap)
					}
				case *aixmfeatures.FlightRestrictionRouteType:
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
					case "PriorPermission":
						__gong__assignPointer(&inst.PriorPermission, rhs, identifierMap)
					case "RouteElement":
						__gong__assignSliceOfPointers(&inst.RouteElement, rhs, identifierMap)
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.FlightRestrictionTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "FlightRestrictionTimeSlice":
						__gong__assignPointer(&inst.FlightRestrictionTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.FlightRestrictionTimeSliceType:
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
					case "Designator":
						__gong__assignPointer(&inst.Designator, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Instruction":
						__gong__assignPointer(&inst.Instruction, rhs, identifierMap)
					case "Flight":
						__gong__assignPointer(&inst.Flight, rhs, identifierMap)
					case "RegulatedRoute":
						__gong__assignSliceOfPointers(&inst.RegulatedRoute, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.FlightRestrictionType:
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
				case *aixmfeatures.FlightRoutingElementPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FlightRoutingElement":
						__gong__assignPointer(&inst.FlightRoutingElement, rhs, identifierMap)
					}
				case *aixmfeatures.FlightRoutingElementType:
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
					case "OrderNumber":
						__gong__assignPointer(&inst.OrderNumber, rhs, identifierMap)
					case "Speed":
						__gong__assignPointer(&inst.Speed, rhs, identifierMap)
					case "SpeedReference":
						__gong__assignPointer(&inst.SpeedReference, rhs, identifierMap)
					case "SpeedCriteria":
						__gong__assignPointer(&inst.SpeedCriteria, rhs, identifierMap)
					case "FlightLevel":
						__gong__assignSliceOfPointers(&inst.FlightLevel, rhs, identifierMap)
					case "Element_directFlightElement":
						__gong__assignPointer(&inst.Element_directFlightElement, rhs, identifierMap)
					case "Element_routePortionElement":
						__gong__assignPointer(&inst.Element_routePortionElement, rhs, identifierMap)
					case "Element_standardInstrumentDepartureElement":
						__gong__assignPointer(&inst.Element_standardInstrumentDepartureElement, rhs, identifierMap)
					case "Element_standardInstrumentArrivalElement":
						__gong__assignPointer(&inst.Element_standardInstrumentArrivalElement, rhs, identifierMap)
					case "PointElement_fixDesignatedPoint":
						__gong__assignPointer(&inst.PointElement_fixDesignatedPoint, rhs, identifierMap)
					case "PointElement_navaidSystem":
						__gong__assignPointer(&inst.PointElement_navaidSystem, rhs, identifierMap)
					case "PointElement_position":
						__gong__assignPointer(&inst.PointElement_position, rhs, identifierMap)
					case "PointElement_runwayPoint":
						__gong__assignPointer(&inst.PointElement_runwayPoint, rhs, identifierMap)
					case "PointElement_aimingPoint":
						__gong__assignPointer(&inst.PointElement_aimingPoint, rhs, identifierMap)
					case "PointElement_airportReferencePoint":
						__gong__assignPointer(&inst.PointElement_airportReferencePoint, rhs, identifierMap)
					case "Element_airspaceElement":
						__gong__assignPointer(&inst.Element_airspaceElement, rhs, identifierMap)
					case "Element_airportHeliportElement":
						__gong__assignPointer(&inst.Element_airportHeliportElement, rhs, identifierMap)
					case "Element_aerialRefuellingElement":
						__gong__assignPointer(&inst.Element_aerialRefuellingElement, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
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
				case *aixmfeatures.MeteorologyPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Meteorology":
						__gong__assignPointer(&inst.Meteorology, rhs, identifierMap)
					}
				case *aixmfeatures.MeteorologyType:
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
					case "FlightConditions":
						__gong__assignPointer(&inst.FlightConditions, rhs, identifierMap)
					case "Visibility":
						__gong__assignPointer(&inst.Visibility, rhs, identifierMap)
					case "VisibilityInterpretation":
						__gong__assignPointer(&inst.VisibilityInterpretation, rhs, identifierMap)
					case "RunwayVisualRange":
						__gong__assignPointer(&inst.RunwayVisualRange, rhs, identifierMap)
					case "RunwayVisualRangeInterpretation":
						__gong__assignPointer(&inst.RunwayVisualRangeInterpretation, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.NavaidPropertyType:
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
				case *aixmfeatures.ObstacleAssessmentAreaPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ObstacleAssessmentArea":
						__gong__assignPointer(&inst.ObstacleAssessmentArea, rhs, identifierMap)
					}
				case *aixmfeatures.ObstacleAssessmentAreaType:
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
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "SectionNumber":
						__gong__assignPointer(&inst.SectionNumber, rhs, identifierMap)
					case "Slope":
						__gong__assignPointer(&inst.Slope, rhs, identifierMap)
					case "AssessedAltitude":
						__gong__assignPointer(&inst.AssessedAltitude, rhs, identifierMap)
					case "SlopeLowerAltitude":
						__gong__assignPointer(&inst.SlopeLowerAltitude, rhs, identifierMap)
					case "GradientLowHigh":
						__gong__assignPointer(&inst.GradientLowHigh, rhs, identifierMap)
					case "SurfaceZone":
						__gong__assignPointer(&inst.SurfaceZone, rhs, identifierMap)
					case "SafetyRegulation":
						__gong__assignPointer(&inst.SafetyRegulation, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "SignificantObstacle":
						__gong__assignSliceOfPointers(&inst.SignificantObstacle, rhs, identifierMap)
					case "Surface":
						__gong__assignPointer(&inst.Surface, rhs, identifierMap)
					case "StartingCurve":
						__gong__assignPointer(&inst.StartingCurve, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ObstaclePlacementPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ObstaclePlacement":
						__gong__assignPointer(&inst.ObstaclePlacement, rhs, identifierMap)
					}
				case *aixmfeatures.ObstaclePlacementType:
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
					case "ObstacleBearing":
						__gong__assignPointer(&inst.ObstacleBearing, rhs, identifierMap)
					case "ObstacleDistance":
						__gong__assignPointer(&inst.ObstacleDistance, rhs, identifierMap)
					case "PointType":
						__gong__assignPointer(&inst.PointType, rhs, identifierMap)
					case "ObstaclePlacement":
						__gong__assignPointer(&inst.ObstaclePlacement, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ObstructionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Obstruction":
						__gong__assignPointer(&inst.Obstruction, rhs, identifierMap)
					}
				case *aixmfeatures.ObstructionType:
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
					case "RequiredClearance":
						__gong__assignPointer(&inst.RequiredClearance, rhs, identifierMap)
					case "MinimumAltitude":
						__gong__assignPointer(&inst.MinimumAltitude, rhs, identifierMap)
					case "SurfacePenetration":
						__gong__assignPointer(&inst.SurfacePenetration, rhs, identifierMap)
					case "SlopePenetration":
						__gong__assignPointer(&inst.SlopePenetration, rhs, identifierMap)
					case "Controlling":
						__gong__assignPointer(&inst.Controlling, rhs, identifierMap)
					case "CloseIn":
						__gong__assignPointer(&inst.CloseIn, rhs, identifierMap)
					case "TheVerticalStructure":
						__gong__assignPointer(&inst.TheVerticalStructure, rhs, identifierMap)
					case "Adjustment":
						__gong__assignSliceOfPointers(&inst.Adjustment, rhs, identifierMap)
					case "ObstaclePlacement":
						__gong__assignSliceOfPointers(&inst.ObstaclePlacement, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.OnlineContactPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "OnlineContact":
						__gong__assignPointer(&inst.OnlineContact, rhs, identifierMap)
					}
				case *aixmfeatures.OnlineContactType:
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
					case "Network":
						__gong__assignPointer(&inst.Network, rhs, identifierMap)
					case "Linkage":
						__gong__assignPointer(&inst.Linkage, rhs, identifierMap)
					case "Protocol":
						__gong__assignPointer(&inst.Protocol, rhs, identifierMap)
					case "EMail":
						__gong__assignPointer(&inst.EMail, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *aixmfeatures.PointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Point":
						__gong__assignPointer(&inst.Point, rhs, identifierMap)
					case "ElevatedPoint":
						__gong__assignPointer(&inst.ElevatedPoint, rhs, identifierMap)
					}
				case *aixmfeatures.PointReferencePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "PointReference":
						__gong__assignPointer(&inst.PointReference, rhs, identifierMap)
					}
				case *aixmfeatures.PointReferenceType:
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
					case "Role":
						__gong__assignPointer(&inst.Role, rhs, identifierMap)
					case "PriorFixTolerance":
						__gong__assignPointer(&inst.PriorFixTolerance, rhs, identifierMap)
					case "PostFixTolerance":
						__gong__assignPointer(&inst.PostFixTolerance, rhs, identifierMap)
					case "Point":
						__gong__assignPointer(&inst.Point, rhs, identifierMap)
					case "FacilityAngle":
						__gong__assignSliceOfPointers(&inst.FacilityAngle, rhs, identifierMap)
					case "FacilityDistance":
						__gong__assignSliceOfPointers(&inst.FacilityDistance, rhs, identifierMap)
					case "FixToleranceArea":
						__gong__assignPointer(&inst.FixToleranceArea, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *aixmfeatures.PostalAddressPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "PostalAddress":
						__gong__assignPointer(&inst.PostalAddress, rhs, identifierMap)
					}
				case *aixmfeatures.PostalAddressType:
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
					case "DeliveryPoint":
						__gong__assignPointer(&inst.DeliveryPoint, rhs, identifierMap)
					case "City":
						__gong__assignPointer(&inst.City, rhs, identifierMap)
					case "AdministrativeArea":
						__gong__assignPointer(&inst.AdministrativeArea, rhs, identifierMap)
					case "PostalCode":
						__gong__assignPointer(&inst.PostalCode, rhs, identifierMap)
					case "Country":
						__gong__assignPointer(&inst.Country, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RadioFrequencyAreaPropertyType:
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
				case *aixmfeatures.RouteAvailabilityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RouteAvailability":
						__gong__assignPointer(&inst.RouteAvailability, rhs, identifierMap)
					}
				case *aixmfeatures.RouteAvailabilityType:
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
					case "Direction":
						__gong__assignPointer(&inst.Direction, rhs, identifierMap)
					case "CardinalDirection":
						__gong__assignPointer(&inst.CardinalDirection, rhs, identifierMap)
					case "Status":
						__gong__assignPointer(&inst.Status, rhs, identifierMap)
					case "Levels":
						__gong__assignSliceOfPointers(&inst.Levels, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RouteDMETimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RouteDMETimeSlice":
						__gong__assignPointer(&inst.RouteDMETimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RouteDMETimeSliceType:
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
					case "CriticalDME":
						__gong__assignPointer(&inst.CriticalDME, rhs, identifierMap)
					case "Satisfactory":
						__gong__assignPointer(&inst.Satisfactory, rhs, identifierMap)
					case "ReferencedDME":
						__gong__assignPointer(&inst.ReferencedDME, rhs, identifierMap)
					case "ApplicableRoutePortion":
						__gong__assignPointer(&inst.ApplicableRoutePortion, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RouteDMEType:
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
				case *aixmfeatures.RoutePortionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RoutePortion":
						__gong__assignPointer(&inst.RoutePortion, rhs, identifierMap)
					}
				case *aixmfeatures.RoutePortionType:
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
					case "Start_fixDesignatedPoint":
						__gong__assignPointer(&inst.Start_fixDesignatedPoint, rhs, identifierMap)
					case "Start_navaidSystem":
						__gong__assignPointer(&inst.Start_navaidSystem, rhs, identifierMap)
					case "Start_position":
						__gong__assignPointer(&inst.Start_position, rhs, identifierMap)
					case "Start_runwayPoint":
						__gong__assignPointer(&inst.Start_runwayPoint, rhs, identifierMap)
					case "Start_aimingPoint":
						__gong__assignPointer(&inst.Start_aimingPoint, rhs, identifierMap)
					case "Start_airportReferencePoint":
						__gong__assignPointer(&inst.Start_airportReferencePoint, rhs, identifierMap)
					case "IntermediatePoint_fixDesignatedPoint":
						__gong__assignPointer(&inst.IntermediatePoint_fixDesignatedPoint, rhs, identifierMap)
					case "IntermediatePoint_navaidSystem":
						__gong__assignPointer(&inst.IntermediatePoint_navaidSystem, rhs, identifierMap)
					case "IntermediatePoint_position":
						__gong__assignPointer(&inst.IntermediatePoint_position, rhs, identifierMap)
					case "IntermediatePoint_runwayPoint":
						__gong__assignPointer(&inst.IntermediatePoint_runwayPoint, rhs, identifierMap)
					case "IntermediatePoint_aimingPoint":
						__gong__assignPointer(&inst.IntermediatePoint_aimingPoint, rhs, identifierMap)
					case "IntermediatePoint_airportReferencePoint":
						__gong__assignPointer(&inst.IntermediatePoint_airportReferencePoint, rhs, identifierMap)
					case "ReferencedRoute":
						__gong__assignPointer(&inst.ReferencedRoute, rhs, identifierMap)
					case "End_fixDesignatedPoint":
						__gong__assignPointer(&inst.End_fixDesignatedPoint, rhs, identifierMap)
					case "End_navaidSystem":
						__gong__assignPointer(&inst.End_navaidSystem, rhs, identifierMap)
					case "End_position":
						__gong__assignPointer(&inst.End_position, rhs, identifierMap)
					case "End_runwayPoint":
						__gong__assignPointer(&inst.End_runwayPoint, rhs, identifierMap)
					case "End_aimingPoint":
						__gong__assignPointer(&inst.End_aimingPoint, rhs, identifierMap)
					case "End_airportReferencePoint":
						__gong__assignPointer(&inst.End_airportReferencePoint, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RoutePropertyType:
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
				case *aixmfeatures.RouteSegmentTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RouteSegmentTimeSlice":
						__gong__assignPointer(&inst.RouteSegmentTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RouteSegmentTimeSliceType:
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
					case "Level":
						__gong__assignPointer(&inst.Level, rhs, identifierMap)
					case "UpperLimit":
						__gong__assignPointer(&inst.UpperLimit, rhs, identifierMap)
					case "UpperLimitReference":
						__gong__assignPointer(&inst.UpperLimitReference, rhs, identifierMap)
					case "LowerLimit":
						__gong__assignPointer(&inst.LowerLimit, rhs, identifierMap)
					case "LowerLimitReference":
						__gong__assignPointer(&inst.LowerLimitReference, rhs, identifierMap)
					case "MinimumObstacleClearanceAltitude":
						__gong__assignPointer(&inst.MinimumObstacleClearanceAltitude, rhs, identifierMap)
					case "PathType":
						__gong__assignPointer(&inst.PathType, rhs, identifierMap)
					case "TrueTrack":
						__gong__assignPointer(&inst.TrueTrack, rhs, identifierMap)
					case "MagneticTrack":
						__gong__assignPointer(&inst.MagneticTrack, rhs, identifierMap)
					case "ReverseTrueTrack":
						__gong__assignPointer(&inst.ReverseTrueTrack, rhs, identifierMap)
					case "ReverseMagneticTrack":
						__gong__assignPointer(&inst.ReverseMagneticTrack, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "WidthLeft":
						__gong__assignPointer(&inst.WidthLeft, rhs, identifierMap)
					case "WidthRight":
						__gong__assignPointer(&inst.WidthRight, rhs, identifierMap)
					case "TurnDirection":
						__gong__assignPointer(&inst.TurnDirection, rhs, identifierMap)
					case "SignalGap":
						__gong__assignPointer(&inst.SignalGap, rhs, identifierMap)
					case "MinimumEnrouteAltitude":
						__gong__assignPointer(&inst.MinimumEnrouteAltitude, rhs, identifierMap)
					case "MinimumCrossingAtEnd":
						__gong__assignPointer(&inst.MinimumCrossingAtEnd, rhs, identifierMap)
					case "MinimumCrossingAtEndReference":
						__gong__assignPointer(&inst.MinimumCrossingAtEndReference, rhs, identifierMap)
					case "MaximumCrossingAtEnd":
						__gong__assignPointer(&inst.MaximumCrossingAtEnd, rhs, identifierMap)
					case "MaximumCrossingAtEndReference":
						__gong__assignPointer(&inst.MaximumCrossingAtEndReference, rhs, identifierMap)
					case "NavigationType":
						__gong__assignPointer(&inst.NavigationType, rhs, identifierMap)
					case "RequiredNavigationPerformance":
						__gong__assignPointer(&inst.RequiredNavigationPerformance, rhs, identifierMap)
					case "DesignatorSuffix":
						__gong__assignPointer(&inst.DesignatorSuffix, rhs, identifierMap)
					case "Start":
						__gong__assignPointer(&inst.Start, rhs, identifierMap)
					case "RouteFormed":
						__gong__assignPointer(&inst.RouteFormed, rhs, identifierMap)
					case "EvaluationArea":
						__gong__assignPointer(&inst.EvaluationArea, rhs, identifierMap)
					case "CurveExtent":
						__gong__assignPointer(&inst.CurveExtent, rhs, identifierMap)
					case "End":
						__gong__assignPointer(&inst.End, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RouteSegmentType:
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
				case *aixmfeatures.RouteTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RouteTimeSlice":
						__gong__assignPointer(&inst.RouteTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RouteTimeSliceType:
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
					case "DesignatorPrefix":
						__gong__assignPointer(&inst.DesignatorPrefix, rhs, identifierMap)
					case "DesignatorSecondLetter":
						__gong__assignPointer(&inst.DesignatorSecondLetter, rhs, identifierMap)
					case "DesignatorNumber":
						__gong__assignPointer(&inst.DesignatorNumber, rhs, identifierMap)
					case "MultipleIdentifier":
						__gong__assignPointer(&inst.MultipleIdentifier, rhs, identifierMap)
					case "LocationDesignator":
						__gong__assignPointer(&inst.LocationDesignator, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "FlightRule":
						__gong__assignPointer(&inst.FlightRule, rhs, identifierMap)
					case "InternationalUse":
						__gong__assignPointer(&inst.InternationalUse, rhs, identifierMap)
					case "MilitaryUse":
						__gong__assignPointer(&inst.MilitaryUse, rhs, identifierMap)
					case "MilitaryTrainingType":
						__gong__assignPointer(&inst.MilitaryTrainingType, rhs, identifierMap)
					case "UserOrganisation":
						__gong__assignPointer(&inst.UserOrganisation, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RouteType:
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
				case *aixmfeatures.RulesProceduresTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RulesProceduresTimeSlice":
						__gong__assignPointer(&inst.RulesProceduresTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RulesProceduresTimeSliceType:
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
					case "Category":
						__gong__assignPointer(&inst.Category, rhs, identifierMap)
					case "Title":
						__gong__assignPointer(&inst.Title, rhs, identifierMap)
					case "Content":
						__gong__assignPointer(&inst.Content, rhs, identifierMap)
					case "AffectedLocation":
						__gong__assignSliceOfPointers(&inst.AffectedLocation, rhs, identifierMap)
					case "AffectedArea":
						__gong__assignSliceOfPointers(&inst.AffectedArea, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RulesProceduresType:
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
				case *aixmfeatures.RunwayCentrelinePointPropertyType:
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
				case *aixmfeatures.StandardInstrumentArrivalPropertyType:
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
				case *aixmfeatures.StandardInstrumentDeparturePropertyType:
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
				case *aixmfeatures.StandardLevelColumnPropertyType:
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
				case *aixmfeatures.TelephoneContactPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "TelephoneContact":
						__gong__assignPointer(&inst.TelephoneContact, rhs, identifierMap)
					}
				case *aixmfeatures.TelephoneContactType:
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
					case "Voice":
						__gong__assignPointer(&inst.Voice, rhs, identifierMap)
					case "Facsimile":
						__gong__assignPointer(&inst.Facsimile, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *aixmfeatures.TouchDownLiftOffPropertyType:
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
					case "Route":
						__gong__assignPointer(&inst.Route, rhs, identifierMap)
					case "RouteSegment":
						__gong__assignPointer(&inst.RouteSegment, rhs, identifierMap)
					case "RouteDME":
						__gong__assignPointer(&inst.RouteDME, rhs, identifierMap)
					case "ChangeOverPoint":
						__gong__assignPointer(&inst.ChangeOverPoint, rhs, identifierMap)
					case "FlightRestriction":
						__gong__assignPointer(&inst.FlightRestriction, rhs, identifierMap)
					case "AirspaceBorderCrossing":
						__gong__assignPointer(&inst.AirspaceBorderCrossing, rhs, identifierMap)
					case "RulesProcedures":
						__gong__assignPointer(&inst.RulesProcedures, rhs, identifierMap)
					case "AerialRefuelling":
						__gong__assignPointer(&inst.AerialRefuelling, rhs, identifierMap)
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
