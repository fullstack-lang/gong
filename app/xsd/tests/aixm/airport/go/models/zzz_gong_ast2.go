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
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/xlink"
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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/xlink":
			aliasToCanonical[alias] = "xlink"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/gml":
			aliasToCanonical[alias] = "gml"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmabstractgmlobjecttypes":
			aliasToCanonical[alias] = "aixmabstractgmlobjecttypes"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmdatatypes":
			aliasToCanonical[alias] = "aixmdatatypes"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models/aixmfeatures":
			aliasToCanonical[alias] = "aixmfeatures"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airport/go/models":
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
				case "CodeAircraftCategoryType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftCategoryType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftEngineNumberType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftEngineNumberType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftEngineType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftEngineType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftICAOType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftICAOType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftStandType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftStandType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftWingspanClassType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftWingspanClassType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAirportHeliportCollocationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAirportHeliportCollocationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAirportHeliportDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAirportHeliportDesignatorType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAirportHeliportType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAirportHeliportType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAirportWarningType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAirportWarningType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeApproachGuidanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeApproachGuidanceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeApproachLightingICAOType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeApproachLightingICAOType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeApproachLightingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeApproachLightingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeApronElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeApronElementType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeApronSectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeApronSectionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeArrestingGearEnergyAbsorbType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeArrestingGearEnergyAbsorbType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeArrestingGearEngageDeviceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeArrestingGearEngageDeviceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAuthorityRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAuthorityRoleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeBuoyDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeBuoyDesignatorType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeBuoyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeBuoyType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeColourType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeColourType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeCommunicationModeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeCommunicationModeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeContaminationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeDayType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeDeclaredDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeDeclaredDistanceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
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
				case "CodeFlightRuleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFlightRuleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFlightStatusType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFlightType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFrictionDeviceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFrictionDeviceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFrictionEstimateType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFrictionEstimateType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeGradeSeparationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeGradeSeparationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeGuidanceLineType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeGuidanceLineType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeHelicopterPerformanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeHelicopterPerformanceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeHoldingCategoryType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeHoldingCategoryType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeIATAType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeIATAType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeICAOType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeICAOType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLightHoldingPositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLightHoldingPositionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLightIntensityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLightIntensityType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLightSourceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLightSourceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLightingJARType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLightingJARType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLoadingBridgeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLoadingBridgeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLogicalOperatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLogicalOperatorType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMarkingConditionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMarkingConditionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMarkingStyleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMarkingStyleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMeteoConditionsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMeteoConditionsType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryOperationsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMilitaryOperationsType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMilitaryStatusType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeNavigationEquipmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeNavigationEquipmentType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeNavigationSpecificationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeNavigationSpecificationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeNotePurposeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeNotePurposeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeOperationAirportHeliportType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeOperationAirportHeliportType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeOperationManoeuvringAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeOperationManoeuvringAreaType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodePCNMethodType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodePCNMethodType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodePCNPavementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodePCNPavementType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodePCNSubgradeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodePCNSubgradeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodePCNTyrePressureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodePCNTyrePressureType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeProtectAreaSectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeProtectAreaSectionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRVRReadingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRVRReadingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRVSMType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRVSMType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRoadType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRoadType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRunwayElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRunwayElementType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRunwayMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRunwayMarkingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRunwayPointRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRunwayPointRoleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRunwayProtectionAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRunwayProtectionAreaType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRunwaySectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRunwaySectionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRunwayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRunwayType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeSideType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeSideType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusAirportType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeStatusAirportType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusOperationsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeStatusOperationsType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeSurfaceCompositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeSurfaceCompositionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeSurfaceConditionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeSurfaceConditionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeSurfacePreparationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeSurfacePreparationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTLOFSectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTLOFSectionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTaxiwayElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTaxiwayElementType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTaxiwaySectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTaxiwaySectionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTaxiwayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTaxiwayType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
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
				case "CodeUsageLimitationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeUsageLimitationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeVASISType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeVASISType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeValueInterpretationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeValueInterpretationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalDatumType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeVerticalDatumType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeVisualDockingGuidanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeVisualDockingGuidanceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeWakeTurbulenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeWakeTurbulenceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeWorkAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeWorkAreaType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeYesNoType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeYesNoType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "DateMonthDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.DateMonthDayType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "DateTimeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.DateTimeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "DateType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.DateType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "DateYearType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.DateYearType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
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
				case "ValDepthType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDepthType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceSignedType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDistanceSignedType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDistanceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceVerticalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDistanceVerticalType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDurationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDurationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValFLType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValFLType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValFrictionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValFrictionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValLCNType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValLCNType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValLightIntensityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValLightIntensityType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValMagneticVariationChangeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValMagneticVariationChangeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValMagneticVariationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValMagneticVariationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValPCNType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValPCNType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValPercentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValPercentType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValPressureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValPressureType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValSlopeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValSlopeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValSpeedType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValSpeedType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValTemperatureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValTemperatureType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValWeightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValWeightType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
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
				case "A_extension_43":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_43), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_44":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_44), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_45":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_45), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_46":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_46), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_47":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_47), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_48":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_48), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_49":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_49), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_5":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_5), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_50":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_50), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_51":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_51), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_52":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_52), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_53":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_53), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_54":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_54), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_55":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_55), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_56":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_56), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_57":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_57), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_58":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_58), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_59":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_59), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_6":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_6), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_60":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_60), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_61":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_61), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_62":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_62), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_63":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_63), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_64":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_64), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_65":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_65), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_66":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_66), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_67":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_67), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_68":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_68), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_69":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_69), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_7":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_7), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_70":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_70), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_71":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_71), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_72":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_72), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_73":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_73), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_74":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_74), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_75":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_75), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_76":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_76), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_77":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_77), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_78":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_78), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_79":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_79), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_8":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_8), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_80":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_80), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_81":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_81), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_82":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_82), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_83":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_83), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_84":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_84), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_85":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_85), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_86":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_86), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_87":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_87), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_88":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_88), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_89":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_89), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_9":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_9), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractAirportHeliportProtectionAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractAirportHeliportProtectionAreaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractGroundLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractGroundLightSystemType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractMarkingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractPropertiesWithScheduleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractPropertiesWithScheduleType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractSurfaceContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractSurfaceContaminationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractUsageConditionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractUsageConditionType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftCharacteristicPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftCharacteristicType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftStandContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftStandContaminationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftStandContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftStandContaminationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftStandPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftStandPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftStandTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftStandTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftStandTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftStandTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftStandType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftStandType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportAvailabilityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportAvailabilityPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportAvailabilityType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportCollocationTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportCollocationTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportCollocationTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportCollocationTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportCollocationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportCollocationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportContaminationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportContaminationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportProtectionAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportProtectionAreaPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportResponsibilityOrganisationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportResponsibilityOrganisationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportResponsibilityOrganisationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportResponsibilityOrganisationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportUsagePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportUsagePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportUsageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportUsageType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHotSpotTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHotSpotTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHotSpotTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHotSpotTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHotSpotType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHotSpotType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportProtectionAreaMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportProtectionAreaMarkingTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportProtectionAreaMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportProtectionAreaMarkingTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportProtectionAreaMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportProtectionAreaMarkingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportSuppliesServicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportSuppliesServicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourcePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AltimeterSourcePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourceStatusPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AltimeterSourceStatusPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourceStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AltimeterSourceStatusType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AltimeterSourceTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AltimeterSourceTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AltimeterSourceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApproachLightingSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApproachLightingSystemTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApproachLightingSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApproachLightingSystemTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApproachLightingSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApproachLightingSystemType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronAreaAvailabilityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronAreaAvailabilityPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronAreaAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronAreaAvailabilityType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronAreaUsagePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronAreaUsagePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronAreaUsageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronAreaUsageType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronContaminationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronContaminationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronElementPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronElementTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronElementTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronElementTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronElementTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronElementType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronLightSystemTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronLightSystemTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronLightSystemType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronMarkingTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronMarkingTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronMarkingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApronType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApronType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ArrestingGearTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ArrestingGearTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ArrestingGearTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ArrestingGearTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ArrestingGearType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ArrestingGearType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CityPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CityType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ConditionCombinationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ConditionCombinationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ConditionCombinationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ConditionCombinationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ContactInformationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ContactInformationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ContactInformationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ContactInformationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CurveType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DeicingAreaMarkingTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DeicingAreaMarkingTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DeicingAreaMarkingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DeicingAreaPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DeicingAreaTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DeicingAreaTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DeicingAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DeicingAreaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
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
				case "FlightCharacteristicPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightCharacteristicPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FlightCharacteristicType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FlightCharacteristicType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FloatingDockSitePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FloatingDockSitePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FloatingDockSiteTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FloatingDockSiteTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FloatingDockSiteTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FloatingDockSiteTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FloatingDockSiteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FloatingDockSiteType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GroundLightingAvailabilityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GroundLightingAvailabilityPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GroundLightingAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GroundLightingAvailabilityType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GuidanceLineLightSystemTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GuidanceLineLightSystemTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GuidanceLineLightSystemType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GuidanceLineMarkingTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GuidanceLineMarkingTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GuidanceLineMarkingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLinePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GuidanceLinePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GuidanceLineTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GuidanceLineTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GuidanceLineType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GuidanceLineType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
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
				case "ManoeuvringAreaAvailabilityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ManoeuvringAreaAvailabilityPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ManoeuvringAreaAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ManoeuvringAreaAvailabilityType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ManoeuvringAreaUsagePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ManoeuvringAreaUsagePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ManoeuvringAreaUsageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ManoeuvringAreaUsageType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MarkingBuoyTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MarkingBuoyTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MarkingBuoyTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MarkingBuoyTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MarkingBuoyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MarkingBuoyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MarkingElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MarkingElementPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MarkingElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MarkingElementType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MeteorologyPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MeteorologyPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MeteorologyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MeteorologyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavaidEquipmentDistancePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavaidEquipmentDistancePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavaidEquipmentDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavaidEquipmentDistanceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavaidEquipmentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavaidEquipmentPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NonMovementAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NonMovementAreaTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NonMovementAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NonMovementAreaTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NonMovementAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NonMovementAreaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NotePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NoteType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OnlineContactPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OnlineContactPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OnlineContactType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OnlineContactType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OrganisationAuthorityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OrganisationAuthorityPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PassengerLoadingBridgeTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PassengerLoadingBridgeTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PassengerLoadingBridgeTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PassengerLoadingBridgeTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PassengerLoadingBridgeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PassengerLoadingBridgeType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PostalAddressPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PostalAddressPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PostalAddressType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PostalAddressType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RidgePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RidgePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RidgeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RidgeType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RoadTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RoadTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RoadTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RoadTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RoadType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RoadType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayBlastPadTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayBlastPadTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayBlastPadTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayBlastPadTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayBlastPadType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayBlastPadType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayCentrelinePointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayCentrelinePointTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayCentrelinePointTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayCentrelinePointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayContaminationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayContaminationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayDeclaredDistancePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayDeclaredDistancePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayDeclaredDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayDeclaredDistanceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayDeclaredDistanceValuePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayDeclaredDistanceValuePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayDeclaredDistanceValueType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayDeclaredDistanceValueType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayDirectionLightSystemTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayDirectionLightSystemTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayDirectionLightSystemType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayDirectionPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayDirectionTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayDirectionTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayDirectionType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayElementPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayElementTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayElementTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayElementTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayElementTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayElementType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayMarkingTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayMarkingTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayMarkingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayProtectAreaLightSystemTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayProtectAreaLightSystemTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayProtectAreaLightSystemType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayProtectAreaPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayProtectAreaTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayProtectAreaTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayProtectAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayProtectAreaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwaySectionContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwaySectionContaminationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwaySectionContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwaySectionContaminationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayVisualRangeTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayVisualRangeTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayVisualRangeTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayVisualRangeTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayVisualRangeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayVisualRangeType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneLandingAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SeaplaneLandingAreaPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneLandingAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SeaplaneLandingAreaTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneLandingAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SeaplaneLandingAreaTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneLandingAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SeaplaneLandingAreaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneRampSitePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SeaplaneRampSitePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneRampSiteTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SeaplaneRampSiteTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneRampSiteTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SeaplaneRampSiteTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SeaplaneRampSiteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SeaplaneRampSiteType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandMarkingTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandMarkingTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandMarkingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurfaceCharacteristicsPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurfaceCharacteristicsPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurfaceCharacteristicsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurfaceCharacteristicsType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurfaceContaminationLayerPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurfaceContaminationLayerPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurfaceContaminationLayerType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurfaceContaminationLayerType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurfaceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurveyControlPointTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurveyControlPointTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurveyControlPointTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurveyControlPointTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurveyControlPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurveyControlPointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiHoldingPositionLightSystemTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiHoldingPositionLightSystemTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiHoldingPositionLightSystemType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiHoldingPositionMarkingTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiHoldingPositionMarkingTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiHoldingPositionMarkingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiHoldingPositionPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiHoldingPositionTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiHoldingPositionTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiHoldingPositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiHoldingPositionType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayContaminationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayContaminationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayElementPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayElementPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayElementTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayElementTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayElementTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayElementTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayElementType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayLightSystemTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayLightSystemTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayLightSystemType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayMarkingTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayMarkingTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayMarkingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TaxiwayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TaxiwayType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TelephoneContactPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TelephoneContactPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TelephoneContactType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TelephoneContactType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TimesheetPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TimesheetPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TimesheetType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TimesheetType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffContaminationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffContaminationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffContaminationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffLightSystemTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffLightSystemTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffLightSystemTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffLightSystemTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffLightSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffLightSystemType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffMarkingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffMarkingTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffMarkingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffMarkingTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffMarkingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffSafeAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffSafeAreaTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffSafeAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffSafeAreaTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffSafeAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffSafeAreaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "VisualGlideSlopeIndicatorTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.VisualGlideSlopeIndicatorTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "VisualGlideSlopeIndicatorTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.VisualGlideSlopeIndicatorTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "VisualGlideSlopeIndicatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.VisualGlideSlopeIndicatorType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "WorkAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.WorkAreaTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "WorkAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.WorkAreaTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "WorkAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.WorkAreaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "WorkareaActivityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.WorkareaActivityPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "WorkareaActivityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.WorkareaActivityType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
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
				case *aixmdatatypes.CodeAircraftStandType:
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
				case *aixmdatatypes.CodeAirportHeliportCollocationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAirportHeliportDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAirportHeliportType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAirportWarningType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeApproachGuidanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeApproachLightingICAOType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeApproachLightingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeApronElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeApronSectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeArrestingGearEnergyAbsorbType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeArrestingGearEngageDeviceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAuthorityRoleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeBuoyDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeBuoyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
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
				case *aixmdatatypes.CodeCommunicationModeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeContaminationType:
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
				case *aixmdatatypes.CodeDeclaredDistanceType:
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
				case *aixmdatatypes.CodeFrictionDeviceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFrictionEstimateType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeGradeSeparationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeGuidanceLineType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeHelicopterPerformanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeHoldingCategoryType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeIATAType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeICAOType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeLightHoldingPositionType:
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
				case *aixmdatatypes.CodeLightingJARType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeLoadingBridgeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeLogicalOperatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeMarkingConditionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeMarkingStyleType:
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
				case *aixmdatatypes.CodeMilitaryOperationsType:
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
				case *aixmdatatypes.CodeOperationAirportHeliportType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeOperationManoeuvringAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodePCNMethodType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodePCNPavementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodePCNSubgradeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodePCNTyrePressureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeProtectAreaSectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRVRReadingType:
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
				case *aixmdatatypes.CodeRoadType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRunwayElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRunwayMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRunwayPointRoleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRunwayProtectionAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRunwaySectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRunwayType:
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
				case *aixmdatatypes.CodeStatusAirportType:
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
				case *aixmdatatypes.CodeSurfaceCompositionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeSurfaceConditionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeSurfacePreparationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeTLOFSectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeTaxiwayElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeTaxiwaySectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeTaxiwayType:
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
				case *aixmdatatypes.CodeUsageLimitationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeVASISType:
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
				case *aixmdatatypes.CodeVisualDockingGuidanceType:
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
				case *aixmdatatypes.CodeWorkAreaType:
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
				case *aixmdatatypes.DateTimeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.DateType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.DateYearType:
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
				case *aixmdatatypes.ValDepthType:
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
				case *aixmdatatypes.ValFLType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Uom":
						inst.Uom = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *aixmdatatypes.ValFrictionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.ValLCNType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
				case *aixmdatatypes.ValMagneticVariationChangeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.ValMagneticVariationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.ValPCNType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.ValPercentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.ValPressureType:
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
				case *aixmdatatypes.ValTemperatureType:
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
				case *aixmfeatures.A_extension:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAirportHeliportExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportHeliportExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractAirportHeliportAvailabilityExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportHeliportAvailabilityExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_10:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractNonMovementAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractNonMovementAreaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_11:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSurfaceCharacteristicsExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSurfaceCharacteristicsExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_12:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSurveyControlPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSurveyControlPointExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_13:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractWorkAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractWorkAreaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_14:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractWorkareaActivityExtension":
						__gong__assignSliceOfPointers(&inst.AbstractWorkareaActivityExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_15:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAircraftStandExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAircraftStandExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_16:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractApronExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApronExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_17:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractApronAreaAvailabilityExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApronAreaAvailabilityExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_18:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractUsageConditionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractUsageConditionExtension, rhs, identifierMap)
					case "AbstractApronAreaUsageExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApronAreaUsageExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_19:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractApronElementExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApronElementExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_2:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAirportHeliportCollocationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportHeliportCollocationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_20:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractDeicingAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractDeicingAreaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_21:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPassengerLoadingBridgeExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPassengerLoadingBridgeExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_22:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRoadExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRoadExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_23:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTouchDownLiftOffExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTouchDownLiftOffExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_24:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTouchDownLiftOffSafeAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTouchDownLiftOffSafeAreaExtension, rhs, identifierMap)
					case "AbstractAirportHeliportProtectionAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportHeliportProtectionAreaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_25:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractApproachLightingSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApproachLightingSystemExtension, rhs, identifierMap)
					case "AbstractGroundLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGroundLightSystemExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_26:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractApronLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApronLightSystemExtension, rhs, identifierMap)
					case "AbstractGroundLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGroundLightSystemExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_27:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractGroundLightingAvailabilityExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGroundLightingAvailabilityExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_28:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractGuidanceLineLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGuidanceLineLightSystemExtension, rhs, identifierMap)
					case "AbstractGroundLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGroundLightSystemExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_29:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRunwayDirectionLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwayDirectionLightSystemExtension, rhs, identifierMap)
					case "AbstractGroundLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGroundLightSystemExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_3:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractAirportHeliportResponsibilityOrganisationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportHeliportResponsibilityOrganisationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_30:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRunwayProtectAreaLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwayProtectAreaLightSystemExtension, rhs, identifierMap)
					case "AbstractGroundLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGroundLightSystemExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_31:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTaxiHoldingPositionLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTaxiHoldingPositionLightSystemExtension, rhs, identifierMap)
					case "AbstractGroundLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGroundLightSystemExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_32:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTaxiwayLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTaxiwayLightSystemExtension, rhs, identifierMap)
					case "AbstractGroundLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGroundLightSystemExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_33:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTouchDownLiftOffLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTouchDownLiftOffLightSystemExtension, rhs, identifierMap)
					case "AbstractGroundLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGroundLightSystemExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_34:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAirportProtectionAreaMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportProtectionAreaMarkingExtension, rhs, identifierMap)
					case "AbstractMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMarkingExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_35:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractApronMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApronMarkingExtension, rhs, identifierMap)
					case "AbstractMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMarkingExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_36:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractDeicingAreaMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractDeicingAreaMarkingExtension, rhs, identifierMap)
					case "AbstractMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMarkingExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_37:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractGuidanceLineMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGuidanceLineMarkingExtension, rhs, identifierMap)
					case "AbstractMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMarkingExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_38:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractMarkingElementExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMarkingElementExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_39:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRunwayMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwayMarkingExtension, rhs, identifierMap)
					case "AbstractMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMarkingExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_4:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractUsageConditionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractUsageConditionExtension, rhs, identifierMap)
					case "AbstractAirportHeliportUsageExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportHeliportUsageExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_40:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractStandMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractStandMarkingExtension, rhs, identifierMap)
					case "AbstractMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMarkingExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_41:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTaxiHoldingPositionMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTaxiHoldingPositionMarkingExtension, rhs, identifierMap)
					case "AbstractMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMarkingExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_42:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTaxiwayMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTaxiwayMarkingExtension, rhs, identifierMap)
					case "AbstractMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMarkingExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_43:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTouchDownLiftOffMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTouchDownLiftOffMarkingExtension, rhs, identifierMap)
					case "AbstractMarkingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMarkingExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_44:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractArrestingGearExtension":
						__gong__assignSliceOfPointers(&inst.AbstractArrestingGearExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_45:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractManoeuvringAreaAvailabilityExtension":
						__gong__assignSliceOfPointers(&inst.AbstractManoeuvringAreaAvailabilityExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_46:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractUsageConditionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractUsageConditionExtension, rhs, identifierMap)
					case "AbstractManoeuvringAreaUsageExtension":
						__gong__assignSliceOfPointers(&inst.AbstractManoeuvringAreaUsageExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_47:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractNavaidEquipmentDistanceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractNavaidEquipmentDistanceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_48:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRunwayExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwayExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_49:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRunwayBlastPadExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwayBlastPadExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_5:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAirportHotSpotExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportHotSpotExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_50:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRunwayCentrelinePointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwayCentrelinePointExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_51:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRunwayDeclaredDistanceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwayDeclaredDistanceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_52:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractRunwayDeclaredDistanceValueExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwayDeclaredDistanceValueExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_53:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRunwayDirectionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwayDirectionExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_54:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRunwayElementExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwayElementExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_55:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRunwayProtectAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwayProtectAreaExtension, rhs, identifierMap)
					case "AbstractAirportHeliportProtectionAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportHeliportProtectionAreaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_56:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRunwayVisualRangeExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwayVisualRangeExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_57:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractVisualGlideSlopeIndicatorExtension":
						__gong__assignSliceOfPointers(&inst.AbstractVisualGlideSlopeIndicatorExtension, rhs, identifierMap)
					case "AbstractGroundLightSystemExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGroundLightSystemExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_58:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFloatingDockSiteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFloatingDockSiteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_59:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractMarkingBuoyExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMarkingBuoyExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_6:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAltimeterSourceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAltimeterSourceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_60:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSeaplaneLandingAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSeaplaneLandingAreaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_61:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSeaplaneRampSiteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSeaplaneRampSiteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_62:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSurfaceContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSurfaceContaminationExtension, rhs, identifierMap)
					case "AbstractAircraftStandContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAircraftStandContaminationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_63:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSurfaceContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSurfaceContaminationExtension, rhs, identifierMap)
					case "AbstractAirportHeliportContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportHeliportContaminationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_64:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSurfaceContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSurfaceContaminationExtension, rhs, identifierMap)
					case "AbstractApronContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApronContaminationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_65:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRidgeExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRidgeExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_66:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSurfaceContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSurfaceContaminationExtension, rhs, identifierMap)
					case "AbstractRunwayContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwayContaminationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_67:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSurfaceContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSurfaceContaminationExtension, rhs, identifierMap)
					case "AbstractRunwaySectionContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRunwaySectionContaminationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_68:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSurfaceContaminationLayerExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSurfaceContaminationLayerExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_69:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSurfaceContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSurfaceContaminationExtension, rhs, identifierMap)
					case "AbstractTaxiwayContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTaxiwayContaminationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_7:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractAltimeterSourceStatusExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAltimeterSourceStatusExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_70:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSurfaceContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSurfaceContaminationExtension, rhs, identifierMap)
					case "AbstractTouchDownLiftOffContaminationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTouchDownLiftOffContaminationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_71:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractGuidanceLineExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGuidanceLineExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_72:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTaxiHoldingPositionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTaxiHoldingPositionExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_73:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTaxiwayExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTaxiwayExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_74:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTaxiwayElementExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTaxiwayElementExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_75:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractElevatedCurveExtension":
						__gong__assignSliceOfPointers(&inst.AbstractElevatedCurveExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_76:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractElevatedPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractElevatedPointExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_77:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractElevatedSurfaceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractElevatedSurfaceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_78:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractLinguisticNoteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractLinguisticNoteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_79:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractNoteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractNoteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_8:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractCityExtension":
						__gong__assignSliceOfPointers(&inst.AbstractCityExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_80:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractContactInformationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractContactInformationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_81:
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
				case *aixmfeatures.A_extension_82:
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
				case *aixmfeatures.A_extension_83:
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
				case *aixmfeatures.A_extension_84:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAircraftCharacteristicExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAircraftCharacteristicExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_85:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFlightCharacteristicExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFlightCharacteristicExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_86:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractLightElementExtension":
						__gong__assignSliceOfPointers(&inst.AbstractLightElementExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_87:
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
				case *aixmfeatures.A_extension_88:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractMeteorologyExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMeteorologyExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_89:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTimesheetExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTimesheetExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_9:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractConditionCombinationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractConditionCombinationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.AbstractAirportHeliportProtectionAreaType:
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
				case *aixmfeatures.AbstractGroundLightSystemType:
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
				case *aixmfeatures.AbstractMarkingType:
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
				case *aixmfeatures.AbstractSurfaceContaminationType:
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
				case *aixmfeatures.AbstractUsageConditionType:
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
				case *aixmfeatures.AircraftStandContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AircraftStandContamination":
						__gong__assignPointer(&inst.AircraftStandContamination, rhs, identifierMap)
					}
				case *aixmfeatures.AircraftStandContaminationType:
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
					case "ObservationTime":
						__gong__assignPointer(&inst.ObservationTime, rhs, identifierMap)
					case "Depth":
						__gong__assignPointer(&inst.Depth, rhs, identifierMap)
					case "FrictionCoefficient":
						__gong__assignPointer(&inst.FrictionCoefficient, rhs, identifierMap)
					case "FrictionEstimation":
						__gong__assignPointer(&inst.FrictionEstimation, rhs, identifierMap)
					case "FrictionDevice":
						__gong__assignPointer(&inst.FrictionDevice, rhs, identifierMap)
					case "ObscuredLights":
						__gong__assignPointer(&inst.ObscuredLights, rhs, identifierMap)
					case "FurtherClearanceTime":
						__gong__assignPointer(&inst.FurtherClearanceTime, rhs, identifierMap)
					case "FurtherTotalClearance":
						__gong__assignPointer(&inst.FurtherTotalClearance, rhs, identifierMap)
					case "NextObservationTime":
						__gong__assignPointer(&inst.NextObservationTime, rhs, identifierMap)
					case "Proportion":
						__gong__assignPointer(&inst.Proportion, rhs, identifierMap)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AircraftStandPropertyType:
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
				case *aixmfeatures.AircraftStandTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AircraftStandTimeSlice":
						__gong__assignPointer(&inst.AircraftStandTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AircraftStandTimeSliceType:
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
					case "VisualDockingSystem":
						__gong__assignPointer(&inst.VisualDockingSystem, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "ApronLocation":
						__gong__assignPointer(&inst.ApronLocation, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Contaminant":
						__gong__assignSliceOfPointers(&inst.Contaminant, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AircraftStandType:
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
				case *aixmfeatures.AirportHeliportAvailabilityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirportHeliportAvailability":
						__gong__assignPointer(&inst.AirportHeliportAvailability, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHeliportAvailabilityType:
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
					case "OperationalStatus":
						__gong__assignPointer(&inst.OperationalStatus, rhs, identifierMap)
					case "Warning":
						__gong__assignPointer(&inst.Warning, rhs, identifierMap)
					case "Usage":
						__gong__assignSliceOfPointers(&inst.Usage, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHeliportCollocationTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirportHeliportCollocationTimeSlice":
						__gong__assignPointer(&inst.AirportHeliportCollocationTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHeliportCollocationTimeSliceType:
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
					case "HostAirport":
						__gong__assignPointer(&inst.HostAirport, rhs, identifierMap)
					case "DependentAirport":
						__gong__assignPointer(&inst.DependentAirport, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHeliportCollocationType:
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
				case *aixmfeatures.AirportHeliportContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirportHeliportContamination":
						__gong__assignPointer(&inst.AirportHeliportContamination, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHeliportContaminationType:
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
					case "ObservationTime":
						__gong__assignPointer(&inst.ObservationTime, rhs, identifierMap)
					case "Depth":
						__gong__assignPointer(&inst.Depth, rhs, identifierMap)
					case "FrictionCoefficient":
						__gong__assignPointer(&inst.FrictionCoefficient, rhs, identifierMap)
					case "FrictionEstimation":
						__gong__assignPointer(&inst.FrictionEstimation, rhs, identifierMap)
					case "FrictionDevice":
						__gong__assignPointer(&inst.FrictionDevice, rhs, identifierMap)
					case "ObscuredLights":
						__gong__assignPointer(&inst.ObscuredLights, rhs, identifierMap)
					case "FurtherClearanceTime":
						__gong__assignPointer(&inst.FurtherClearanceTime, rhs, identifierMap)
					case "FurtherTotalClearance":
						__gong__assignPointer(&inst.FurtherTotalClearance, rhs, identifierMap)
					case "NextObservationTime":
						__gong__assignPointer(&inst.NextObservationTime, rhs, identifierMap)
					case "Proportion":
						__gong__assignPointer(&inst.Proportion, rhs, identifierMap)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
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
				case *aixmfeatures.AirportHeliportProtectionAreaPropertyType:
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
				case *aixmfeatures.AirportHeliportResponsibilityOrganisationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirportHeliportResponsibilityOrganisation":
						__gong__assignPointer(&inst.AirportHeliportResponsibilityOrganisation, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHeliportResponsibilityOrganisationType:
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
					case "Role":
						__gong__assignPointer(&inst.Role, rhs, identifierMap)
					case "TheOrganisationAuthority":
						__gong__assignPointer(&inst.TheOrganisationAuthority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHeliportTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirportHeliportTimeSlice":
						__gong__assignPointer(&inst.AirportHeliportTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHeliportTimeSliceType:
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
					case "LocationIndicatorICAO":
						__gong__assignPointer(&inst.LocationIndicatorICAO, rhs, identifierMap)
					case "DesignatorIATA":
						__gong__assignPointer(&inst.DesignatorIATA, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "CertifiedICAO":
						__gong__assignPointer(&inst.CertifiedICAO, rhs, identifierMap)
					case "PrivateUse":
						__gong__assignPointer(&inst.PrivateUse, rhs, identifierMap)
					case "ControlType":
						__gong__assignPointer(&inst.ControlType, rhs, identifierMap)
					case "FieldElevation":
						__gong__assignPointer(&inst.FieldElevation, rhs, identifierMap)
					case "FieldElevationAccuracy":
						__gong__assignPointer(&inst.FieldElevationAccuracy, rhs, identifierMap)
					case "VerticalDatum":
						__gong__assignPointer(&inst.VerticalDatum, rhs, identifierMap)
					case "MagneticVariation":
						__gong__assignPointer(&inst.MagneticVariation, rhs, identifierMap)
					case "MagneticVariationAccuracy":
						__gong__assignPointer(&inst.MagneticVariationAccuracy, rhs, identifierMap)
					case "DateMagneticVariation":
						__gong__assignPointer(&inst.DateMagneticVariation, rhs, identifierMap)
					case "MagneticVariationChange":
						__gong__assignPointer(&inst.MagneticVariationChange, rhs, identifierMap)
					case "ReferenceTemperature":
						__gong__assignPointer(&inst.ReferenceTemperature, rhs, identifierMap)
					case "AltimeterCheckLocation":
						__gong__assignPointer(&inst.AltimeterCheckLocation, rhs, identifierMap)
					case "SecondaryPowerSupply":
						__gong__assignPointer(&inst.SecondaryPowerSupply, rhs, identifierMap)
					case "WindDirectionIndicator":
						__gong__assignPointer(&inst.WindDirectionIndicator, rhs, identifierMap)
					case "LandingDirectionIndicator":
						__gong__assignPointer(&inst.LandingDirectionIndicator, rhs, identifierMap)
					case "TransitionAltitude":
						__gong__assignPointer(&inst.TransitionAltitude, rhs, identifierMap)
					case "TransitionLevel":
						__gong__assignPointer(&inst.TransitionLevel, rhs, identifierMap)
					case "LowestTemperature":
						__gong__assignPointer(&inst.LowestTemperature, rhs, identifierMap)
					case "Abandoned":
						__gong__assignPointer(&inst.Abandoned, rhs, identifierMap)
					case "CertificationDate":
						__gong__assignPointer(&inst.CertificationDate, rhs, identifierMap)
					case "CertificationExpirationDate":
						__gong__assignPointer(&inst.CertificationExpirationDate, rhs, identifierMap)
					case "Contaminant":
						__gong__assignSliceOfPointers(&inst.Contaminant, rhs, identifierMap)
					case "ServedCity":
						__gong__assignSliceOfPointers(&inst.ServedCity, rhs, identifierMap)
					case "ResponsibleOrganisation":
						__gong__assignPointer(&inst.ResponsibleOrganisation, rhs, identifierMap)
					case "ARP":
						__gong__assignPointer(&inst.ARP, rhs, identifierMap)
					case "AviationBoundary":
						__gong__assignPointer(&inst.AviationBoundary, rhs, identifierMap)
					case "AltimeterSource":
						__gong__assignSliceOfPointers(&inst.AltimeterSource, rhs, identifierMap)
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHeliportType:
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
				case *aixmfeatures.AirportHeliportUsagePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AirportHeliportUsage":
						__gong__assignPointer(&inst.AirportHeliportUsage, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHeliportUsageType:
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
					case "PriorPermission":
						__gong__assignPointer(&inst.PriorPermission, rhs, identifierMap)
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "Selection":
						__gong__assignPointer(&inst.Selection, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Operation":
						__gong__assignPointer(&inst.Operation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHotSpotTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirportHotSpotTimeSlice":
						__gong__assignPointer(&inst.AirportHotSpotTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHotSpotTimeSliceType:
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
					case "Instruction":
						__gong__assignPointer(&inst.Instruction, rhs, identifierMap)
					case "Area":
						__gong__assignPointer(&inst.Area, rhs, identifierMap)
					case "AffectedAirport":
						__gong__assignPointer(&inst.AffectedAirport, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AirportHotSpotType:
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
				case *aixmfeatures.AirportProtectionAreaMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirportProtectionAreaMarkingTimeSlice":
						__gong__assignPointer(&inst.AirportProtectionAreaMarkingTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AirportProtectionAreaMarkingTimeSliceType:
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
					case "MarkingICAOStandard":
						__gong__assignPointer(&inst.MarkingICAOStandard, rhs, identifierMap)
					case "Condition":
						__gong__assignPointer(&inst.Condition, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkingLocation":
						__gong__assignPointer(&inst.MarkingLocation, rhs, identifierMap)
					case "MarkedProtectionArea":
						__gong__assignPointer(&inst.MarkedProtectionArea, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AirportProtectionAreaMarkingType:
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
				case *aixmfeatures.AirportSuppliesServicePropertyType:
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
				case *aixmfeatures.AltimeterSourcePropertyType:
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
				case *aixmfeatures.AltimeterSourceStatusPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AltimeterSourceStatus":
						__gong__assignPointer(&inst.AltimeterSourceStatus, rhs, identifierMap)
					}
				case *aixmfeatures.AltimeterSourceStatusType:
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
					case "OperationalStatus":
						__gong__assignPointer(&inst.OperationalStatus, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AltimeterSourceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AltimeterSourceTimeSlice":
						__gong__assignPointer(&inst.AltimeterSourceTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AltimeterSourceTimeSliceType:
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
					case "IsRemote":
						__gong__assignPointer(&inst.IsRemote, rhs, identifierMap)
					case "IsPrimary":
						__gong__assignPointer(&inst.IsPrimary, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AltimeterSourceType:
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
				case *aixmfeatures.ApproachLightingSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ApproachLightingSystemTimeSlice":
						__gong__assignPointer(&inst.ApproachLightingSystemTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.ApproachLightingSystemTimeSliceType:
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
					case "EmergencyLighting":
						__gong__assignPointer(&inst.EmergencyLighting, rhs, identifierMap)
					case "IntensityLevel":
						__gong__assignPointer(&inst.IntensityLevel, rhs, identifierMap)
					case "Colour":
						__gong__assignPointer(&inst.Colour, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "ClassICAO":
						__gong__assignPointer(&inst.ClassICAO, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "SequencedFlashing":
						__gong__assignPointer(&inst.SequencedFlashing, rhs, identifierMap)
					case "AlignmentIndicator":
						__gong__assignPointer(&inst.AlignmentIndicator, rhs, identifierMap)
					case "ServedRunwayDirection":
						__gong__assignPointer(&inst.ServedRunwayDirection, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ApproachLightingSystemType:
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
				case *aixmfeatures.ApronAreaAvailabilityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApronAreaAvailability":
						__gong__assignPointer(&inst.ApronAreaAvailability, rhs, identifierMap)
					}
				case *aixmfeatures.ApronAreaAvailabilityType:
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
					case "OperationalStatus":
						__gong__assignPointer(&inst.OperationalStatus, rhs, identifierMap)
					case "Warning":
						__gong__assignPointer(&inst.Warning, rhs, identifierMap)
					case "Usage":
						__gong__assignSliceOfPointers(&inst.Usage, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ApronAreaUsagePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApronAreaUsage":
						__gong__assignPointer(&inst.ApronAreaUsage, rhs, identifierMap)
					}
				case *aixmfeatures.ApronAreaUsageType:
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
					case "PriorPermission":
						__gong__assignPointer(&inst.PriorPermission, rhs, identifierMap)
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "Selection":
						__gong__assignPointer(&inst.Selection, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ApronContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApronContamination":
						__gong__assignPointer(&inst.ApronContamination, rhs, identifierMap)
					}
				case *aixmfeatures.ApronContaminationType:
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
					case "ObservationTime":
						__gong__assignPointer(&inst.ObservationTime, rhs, identifierMap)
					case "Depth":
						__gong__assignPointer(&inst.Depth, rhs, identifierMap)
					case "FrictionCoefficient":
						__gong__assignPointer(&inst.FrictionCoefficient, rhs, identifierMap)
					case "FrictionEstimation":
						__gong__assignPointer(&inst.FrictionEstimation, rhs, identifierMap)
					case "FrictionDevice":
						__gong__assignPointer(&inst.FrictionDevice, rhs, identifierMap)
					case "ObscuredLights":
						__gong__assignPointer(&inst.ObscuredLights, rhs, identifierMap)
					case "FurtherClearanceTime":
						__gong__assignPointer(&inst.FurtherClearanceTime, rhs, identifierMap)
					case "FurtherTotalClearance":
						__gong__assignPointer(&inst.FurtherTotalClearance, rhs, identifierMap)
					case "NextObservationTime":
						__gong__assignPointer(&inst.NextObservationTime, rhs, identifierMap)
					case "Proportion":
						__gong__assignPointer(&inst.Proportion, rhs, identifierMap)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ApronElementPropertyType:
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
				case *aixmfeatures.ApronElementTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ApronElementTimeSlice":
						__gong__assignPointer(&inst.ApronElementTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.ApronElementTimeSliceType:
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
					case "JetwayAvailability":
						__gong__assignPointer(&inst.JetwayAvailability, rhs, identifierMap)
					case "TowingAvailability":
						__gong__assignPointer(&inst.TowingAvailability, rhs, identifierMap)
					case "DockingAvailability":
						__gong__assignPointer(&inst.DockingAvailability, rhs, identifierMap)
					case "GroundPowerAvailability":
						__gong__assignPointer(&inst.GroundPowerAvailability, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Width":
						__gong__assignPointer(&inst.Width, rhs, identifierMap)
					case "AssociatedApron":
						__gong__assignPointer(&inst.AssociatedApron, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "SupplyService":
						__gong__assignSliceOfPointers(&inst.SupplyService, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ApronElementType:
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
				case *aixmfeatures.ApronLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ApronLightSystemTimeSlice":
						__gong__assignPointer(&inst.ApronLightSystemTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.ApronLightSystemTimeSliceType:
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
					case "EmergencyLighting":
						__gong__assignPointer(&inst.EmergencyLighting, rhs, identifierMap)
					case "IntensityLevel":
						__gong__assignPointer(&inst.IntensityLevel, rhs, identifierMap)
					case "Colour":
						__gong__assignPointer(&inst.Colour, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Position":
						__gong__assignPointer(&inst.Position, rhs, identifierMap)
					case "LightedApron":
						__gong__assignPointer(&inst.LightedApron, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ApronLightSystemType:
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
				case *aixmfeatures.ApronMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ApronMarkingTimeSlice":
						__gong__assignPointer(&inst.ApronMarkingTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.ApronMarkingTimeSliceType:
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
					case "MarkingICAOStandard":
						__gong__assignPointer(&inst.MarkingICAOStandard, rhs, identifierMap)
					case "Condition":
						__gong__assignPointer(&inst.Condition, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkingLocation":
						__gong__assignPointer(&inst.MarkingLocation, rhs, identifierMap)
					case "MarkedApron":
						__gong__assignPointer(&inst.MarkedApron, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ApronMarkingType:
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
				case *aixmfeatures.ApronPropertyType:
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
				case *aixmfeatures.ApronTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ApronTimeSlice":
						__gong__assignPointer(&inst.ApronTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.ApronTimeSliceType:
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
					case "Abandoned":
						__gong__assignPointer(&inst.Abandoned, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "Contaminant":
						__gong__assignSliceOfPointers(&inst.Contaminant, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ApronType:
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
				case *aixmfeatures.ArrestingGearTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ArrestingGearTimeSlice":
						__gong__assignPointer(&inst.ArrestingGearTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.ArrestingGearTimeSliceType:
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
					case "Status":
						__gong__assignPointer(&inst.Status, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Width":
						__gong__assignPointer(&inst.Width, rhs, identifierMap)
					case "EngageDevice":
						__gong__assignPointer(&inst.EngageDevice, rhs, identifierMap)
					case "AbsorbType":
						__gong__assignPointer(&inst.AbsorbType, rhs, identifierMap)
					case "Bidirectional":
						__gong__assignPointer(&inst.Bidirectional, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "RunwayDirection":
						__gong__assignSliceOfPointers(&inst.RunwayDirection, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Extent_surfaceExtent":
						__gong__assignPointer(&inst.Extent_surfaceExtent, rhs, identifierMap)
					case "Extent_curveExtent":
						__gong__assignPointer(&inst.Extent_curveExtent, rhs, identifierMap)
					case "Extent_pointExtent":
						__gong__assignPointer(&inst.Extent_pointExtent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ArrestingGearType:
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
				case *aixmfeatures.CityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "City":
						__gong__assignPointer(&inst.City, rhs, identifierMap)
					}
				case *aixmfeatures.CityType:
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
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ConditionCombinationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ConditionCombination":
						__gong__assignPointer(&inst.ConditionCombination, rhs, identifierMap)
					}
				case *aixmfeatures.ConditionCombinationType:
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
					case "Weather":
						__gong__assignSliceOfPointers(&inst.Weather, rhs, identifierMap)
					case "Aircraft":
						__gong__assignSliceOfPointers(&inst.Aircraft, rhs, identifierMap)
					case "Flight":
						__gong__assignSliceOfPointers(&inst.Flight, rhs, identifierMap)
					case "SubCondition":
						__gong__assignSliceOfPointers(&inst.SubCondition, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *aixmfeatures.DeicingAreaMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "DeicingAreaMarkingTimeSlice":
						__gong__assignPointer(&inst.DeicingAreaMarkingTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.DeicingAreaMarkingTimeSliceType:
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
					case "MarkingICAOStandard":
						__gong__assignPointer(&inst.MarkingICAOStandard, rhs, identifierMap)
					case "Condition":
						__gong__assignPointer(&inst.Condition, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkedDeicingArea":
						__gong__assignPointer(&inst.MarkedDeicingArea, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.DeicingAreaMarkingType:
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
				case *aixmfeatures.DeicingAreaPropertyType:
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
				case *aixmfeatures.DeicingAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "DeicingAreaTimeSlice":
						__gong__assignPointer(&inst.DeicingAreaTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.DeicingAreaTimeSliceType:
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
					case "AssociatedApron":
						__gong__assignPointer(&inst.AssociatedApron, rhs, identifierMap)
					case "TaxiwayLocation":
						__gong__assignPointer(&inst.TaxiwayLocation, rhs, identifierMap)
					case "StandLocation":
						__gong__assignPointer(&inst.StandLocation, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.DeicingAreaType:
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
				case *aixmfeatures.FloatingDockSitePropertyType:
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
				case *aixmfeatures.FloatingDockSiteTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "FloatingDockSiteTimeSlice":
						__gong__assignPointer(&inst.FloatingDockSiteTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.FloatingDockSiteTimeSliceType:
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
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.FloatingDockSiteType:
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
				case *aixmfeatures.GroundLightingAvailabilityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "GroundLightingAvailability":
						__gong__assignPointer(&inst.GroundLightingAvailability, rhs, identifierMap)
					}
				case *aixmfeatures.GroundLightingAvailabilityType:
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
					case "OperationalStatus":
						__gong__assignPointer(&inst.OperationalStatus, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.GuidanceLineLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "GuidanceLineLightSystemTimeSlice":
						__gong__assignPointer(&inst.GuidanceLineLightSystemTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.GuidanceLineLightSystemTimeSliceType:
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
					case "EmergencyLighting":
						__gong__assignPointer(&inst.EmergencyLighting, rhs, identifierMap)
					case "IntensityLevel":
						__gong__assignPointer(&inst.IntensityLevel, rhs, identifierMap)
					case "Colour":
						__gong__assignPointer(&inst.Colour, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "LightedGuidanceLine":
						__gong__assignPointer(&inst.LightedGuidanceLine, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.GuidanceLineLightSystemType:
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
				case *aixmfeatures.GuidanceLineMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "GuidanceLineMarkingTimeSlice":
						__gong__assignPointer(&inst.GuidanceLineMarkingTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.GuidanceLineMarkingTimeSliceType:
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
					case "MarkingICAOStandard":
						__gong__assignPointer(&inst.MarkingICAOStandard, rhs, identifierMap)
					case "Condition":
						__gong__assignPointer(&inst.Condition, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkedGuidanceLine":
						__gong__assignPointer(&inst.MarkedGuidanceLine, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.GuidanceLineMarkingType:
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
				case *aixmfeatures.GuidanceLinePropertyType:
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
				case *aixmfeatures.GuidanceLineTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "GuidanceLineTimeSlice":
						__gong__assignPointer(&inst.GuidanceLineTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.GuidanceLineTimeSliceType:
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
					case "MaxSpeed":
						__gong__assignPointer(&inst.MaxSpeed, rhs, identifierMap)
					case "UsageDirection":
						__gong__assignPointer(&inst.UsageDirection, rhs, identifierMap)
					case "ConnectedTouchDownLiftOff":
						__gong__assignSliceOfPointers(&inst.ConnectedTouchDownLiftOff, rhs, identifierMap)
					case "ConnectedRunwayCentrelinePoint":
						__gong__assignSliceOfPointers(&inst.ConnectedRunwayCentrelinePoint, rhs, identifierMap)
					case "ConnectedApron":
						__gong__assignSliceOfPointers(&inst.ConnectedApron, rhs, identifierMap)
					case "ConnectedStand":
						__gong__assignSliceOfPointers(&inst.ConnectedStand, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "ConnectedTaxiway":
						__gong__assignSliceOfPointers(&inst.ConnectedTaxiway, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.GuidanceLineType:
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
				case *aixmfeatures.ManoeuvringAreaAvailabilityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ManoeuvringAreaAvailability":
						__gong__assignPointer(&inst.ManoeuvringAreaAvailability, rhs, identifierMap)
					}
				case *aixmfeatures.ManoeuvringAreaAvailabilityType:
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
					case "OperationalStatus":
						__gong__assignPointer(&inst.OperationalStatus, rhs, identifierMap)
					case "Warning":
						__gong__assignPointer(&inst.Warning, rhs, identifierMap)
					case "Usage":
						__gong__assignSliceOfPointers(&inst.Usage, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ManoeuvringAreaUsagePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ManoeuvringAreaUsage":
						__gong__assignPointer(&inst.ManoeuvringAreaUsage, rhs, identifierMap)
					}
				case *aixmfeatures.ManoeuvringAreaUsageType:
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
					case "PriorPermission":
						__gong__assignPointer(&inst.PriorPermission, rhs, identifierMap)
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "Selection":
						__gong__assignPointer(&inst.Selection, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Operation":
						__gong__assignPointer(&inst.Operation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.MarkingBuoyTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "MarkingBuoyTimeSlice":
						__gong__assignPointer(&inst.MarkingBuoyTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.MarkingBuoyTimeSliceType:
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
					case "Colour":
						__gong__assignPointer(&inst.Colour, rhs, identifierMap)
					case "TheSeaplaneLandingArea":
						__gong__assignPointer(&inst.TheSeaplaneLandingArea, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.MarkingBuoyType:
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
				case *aixmfeatures.MarkingElementPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "MarkingElement":
						__gong__assignPointer(&inst.MarkingElement, rhs, identifierMap)
					}
				case *aixmfeatures.MarkingElementType:
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
					case "Style":
						__gong__assignPointer(&inst.Style, rhs, identifierMap)
					case "Extent_curveExtent":
						__gong__assignPointer(&inst.Extent_curveExtent, rhs, identifierMap)
					case "Extent_surfaceExtent":
						__gong__assignPointer(&inst.Extent_surfaceExtent, rhs, identifierMap)
					case "Extent_location":
						__gong__assignPointer(&inst.Extent_location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
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
				case *aixmfeatures.NavaidEquipmentDistancePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "NavaidEquipmentDistance":
						__gong__assignPointer(&inst.NavaidEquipmentDistance, rhs, identifierMap)
					}
				case *aixmfeatures.NavaidEquipmentDistanceType:
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
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "DistanceAccuracy":
						__gong__assignPointer(&inst.DistanceAccuracy, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheNavaidEquipment":
						__gong__assignPointer(&inst.TheNavaidEquipment, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *aixmfeatures.NonMovementAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NonMovementAreaTimeSlice":
						__gong__assignPointer(&inst.NonMovementAreaTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.NonMovementAreaTimeSliceType:
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
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.NonMovementAreaType:
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
				case *aixmfeatures.PassengerLoadingBridgeTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "PassengerLoadingBridgeTimeSlice":
						__gong__assignPointer(&inst.PassengerLoadingBridgeTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.PassengerLoadingBridgeTimeSliceType:
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
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "AssociatedStand":
						__gong__assignSliceOfPointers(&inst.AssociatedStand, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.PassengerLoadingBridgeType:
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
				case *aixmfeatures.RidgePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Ridge":
						__gong__assignPointer(&inst.Ridge, rhs, identifierMap)
					}
				case *aixmfeatures.RidgeType:
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
					case "Side":
						__gong__assignPointer(&inst.Side, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "Depth":
						__gong__assignPointer(&inst.Depth, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RoadTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RoadTimeSlice":
						__gong__assignPointer(&inst.RoadTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RoadTimeSliceType:
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
					case "Status":
						__gong__assignPointer(&inst.Status, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Abandoned":
						__gong__assignPointer(&inst.Abandoned, rhs, identifierMap)
					case "AssociatedAirport":
						__gong__assignPointer(&inst.AssociatedAirport, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AccessibleStand":
						__gong__assignSliceOfPointers(&inst.AccessibleStand, rhs, identifierMap)
					case "SurfaceExtent":
						__gong__assignPointer(&inst.SurfaceExtent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RoadType:
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
				case *aixmfeatures.RunwayBlastPadTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayBlastPadTimeSlice":
						__gong__assignPointer(&inst.RunwayBlastPadTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayBlastPadTimeSliceType:
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
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Status":
						__gong__assignPointer(&inst.Status, rhs, identifierMap)
					case "UsedRunwayDirection":
						__gong__assignPointer(&inst.UsedRunwayDirection, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayBlastPadType:
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
				case *aixmfeatures.RunwayCentrelinePointTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayCentrelinePointTimeSlice":
						__gong__assignPointer(&inst.RunwayCentrelinePointTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayCentrelinePointTimeSliceType:
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
					case "Role":
						__gong__assignPointer(&inst.Role, rhs, identifierMap)
					case "Designator":
						__gong__assignPointer(&inst.Designator, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "OnRunway":
						__gong__assignPointer(&inst.OnRunway, rhs, identifierMap)
					case "AssociatedDeclaredDistance":
						__gong__assignSliceOfPointers(&inst.AssociatedDeclaredDistance, rhs, identifierMap)
					case "NavaidEquipment":
						__gong__assignSliceOfPointers(&inst.NavaidEquipment, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayCentrelinePointType:
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
				case *aixmfeatures.RunwayContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RunwayContamination":
						__gong__assignPointer(&inst.RunwayContamination, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayContaminationType:
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
					case "ObservationTime":
						__gong__assignPointer(&inst.ObservationTime, rhs, identifierMap)
					case "Depth":
						__gong__assignPointer(&inst.Depth, rhs, identifierMap)
					case "FrictionCoefficient":
						__gong__assignPointer(&inst.FrictionCoefficient, rhs, identifierMap)
					case "FrictionEstimation":
						__gong__assignPointer(&inst.FrictionEstimation, rhs, identifierMap)
					case "FrictionDevice":
						__gong__assignPointer(&inst.FrictionDevice, rhs, identifierMap)
					case "ObscuredLights":
						__gong__assignPointer(&inst.ObscuredLights, rhs, identifierMap)
					case "FurtherClearanceTime":
						__gong__assignPointer(&inst.FurtherClearanceTime, rhs, identifierMap)
					case "FurtherTotalClearance":
						__gong__assignPointer(&inst.FurtherTotalClearance, rhs, identifierMap)
					case "NextObservationTime":
						__gong__assignPointer(&inst.NextObservationTime, rhs, identifierMap)
					case "Proportion":
						__gong__assignPointer(&inst.Proportion, rhs, identifierMap)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "ClearedLength":
						__gong__assignPointer(&inst.ClearedLength, rhs, identifierMap)
					case "ClearedWidth":
						__gong__assignPointer(&inst.ClearedWidth, rhs, identifierMap)
					case "ClearedSide":
						__gong__assignPointer(&inst.ClearedSide, rhs, identifierMap)
					case "FurtherClearanceLength":
						__gong__assignPointer(&inst.FurtherClearanceLength, rhs, identifierMap)
					case "FurtherClearanceWidth":
						__gong__assignPointer(&inst.FurtherClearanceWidth, rhs, identifierMap)
					case "ObscuredLightsSide":
						__gong__assignPointer(&inst.ObscuredLightsSide, rhs, identifierMap)
					case "ClearedLengthBegin":
						__gong__assignPointer(&inst.ClearedLengthBegin, rhs, identifierMap)
					case "TaxiwayAvailable":
						__gong__assignPointer(&inst.TaxiwayAvailable, rhs, identifierMap)
					case "ApronAvailable":
						__gong__assignPointer(&inst.ApronAvailable, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayDeclaredDistancePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RunwayDeclaredDistance":
						__gong__assignPointer(&inst.RunwayDeclaredDistance, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayDeclaredDistanceType:
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
					case "DeclaredValue":
						__gong__assignSliceOfPointers(&inst.DeclaredValue, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayDeclaredDistanceValuePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RunwayDeclaredDistanceValue":
						__gong__assignPointer(&inst.RunwayDeclaredDistanceValue, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayDeclaredDistanceValueType:
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
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "DistanceAccuracy":
						__gong__assignPointer(&inst.DistanceAccuracy, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayDirectionLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayDirectionLightSystemTimeSlice":
						__gong__assignPointer(&inst.RunwayDirectionLightSystemTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayDirectionLightSystemTimeSliceType:
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
					case "EmergencyLighting":
						__gong__assignPointer(&inst.EmergencyLighting, rhs, identifierMap)
					case "IntensityLevel":
						__gong__assignPointer(&inst.IntensityLevel, rhs, identifierMap)
					case "Colour":
						__gong__assignPointer(&inst.Colour, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Position":
						__gong__assignPointer(&inst.Position, rhs, identifierMap)
					case "AssociatedRunwayDirection":
						__gong__assignPointer(&inst.AssociatedRunwayDirection, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayDirectionLightSystemType:
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
				case *aixmfeatures.RunwayDirectionTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayDirectionTimeSlice":
						__gong__assignPointer(&inst.RunwayDirectionTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayDirectionTimeSliceType:
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
					case "TrueBearing":
						__gong__assignPointer(&inst.TrueBearing, rhs, identifierMap)
					case "TrueBearingAccuracy":
						__gong__assignPointer(&inst.TrueBearingAccuracy, rhs, identifierMap)
					case "MagneticBearing":
						__gong__assignPointer(&inst.MagneticBearing, rhs, identifierMap)
					case "PatternVFR":
						__gong__assignPointer(&inst.PatternVFR, rhs, identifierMap)
					case "SlopeTDZ":
						__gong__assignPointer(&inst.SlopeTDZ, rhs, identifierMap)
					case "ElevationTDZ":
						__gong__assignPointer(&inst.ElevationTDZ, rhs, identifierMap)
					case "ElevationTDZAccuracy":
						__gong__assignPointer(&inst.ElevationTDZAccuracy, rhs, identifierMap)
					case "ApproachMarkingType":
						__gong__assignPointer(&inst.ApproachMarkingType, rhs, identifierMap)
					case "ApproachMarkingCondition":
						__gong__assignPointer(&inst.ApproachMarkingCondition, rhs, identifierMap)
					case "ClassLightingJAR":
						__gong__assignPointer(&inst.ClassLightingJAR, rhs, identifierMap)
					case "PrecisionApproachGuidance":
						__gong__assignPointer(&inst.PrecisionApproachGuidance, rhs, identifierMap)
					case "UsedRunway":
						__gong__assignPointer(&inst.UsedRunway, rhs, identifierMap)
					case "StartingElement":
						__gong__assignPointer(&inst.StartingElement, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayDirectionType:
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
				case *aixmfeatures.RunwayElementPropertyType:
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
				case *aixmfeatures.RunwayElementTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayElementTimeSlice":
						__gong__assignPointer(&inst.RunwayElementTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayElementTimeSliceType:
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
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Width":
						__gong__assignPointer(&inst.Width, rhs, identifierMap)
					case "GradeSeparation":
						__gong__assignPointer(&inst.GradeSeparation, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AssociatedRunway":
						__gong__assignSliceOfPointers(&inst.AssociatedRunway, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayElementType:
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
				case *aixmfeatures.RunwayMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayMarkingTimeSlice":
						__gong__assignPointer(&inst.RunwayMarkingTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayMarkingTimeSliceType:
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
					case "MarkingICAOStandard":
						__gong__assignPointer(&inst.MarkingICAOStandard, rhs, identifierMap)
					case "Condition":
						__gong__assignPointer(&inst.Condition, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkingLocation":
						__gong__assignPointer(&inst.MarkingLocation, rhs, identifierMap)
					case "MarkedRunway":
						__gong__assignPointer(&inst.MarkedRunway, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayMarkingType:
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
				case *aixmfeatures.RunwayPropertyType:
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
				case *aixmfeatures.RunwayProtectAreaLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayProtectAreaLightSystemTimeSlice":
						__gong__assignPointer(&inst.RunwayProtectAreaLightSystemTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayProtectAreaLightSystemTimeSliceType:
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
					case "EmergencyLighting":
						__gong__assignPointer(&inst.EmergencyLighting, rhs, identifierMap)
					case "IntensityLevel":
						__gong__assignPointer(&inst.IntensityLevel, rhs, identifierMap)
					case "Colour":
						__gong__assignPointer(&inst.Colour, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Position":
						__gong__assignPointer(&inst.Position, rhs, identifierMap)
					case "LightedArea":
						__gong__assignPointer(&inst.LightedArea, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayProtectAreaLightSystemType:
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
				case *aixmfeatures.RunwayProtectAreaPropertyType:
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
				case *aixmfeatures.RunwayProtectAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayProtectAreaTimeSlice":
						__gong__assignPointer(&inst.RunwayProtectAreaTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayProtectAreaTimeSliceType:
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
					case "Width":
						__gong__assignPointer(&inst.Width, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Lighting":
						__gong__assignPointer(&inst.Lighting, rhs, identifierMap)
					case "ObstacleFree":
						__gong__assignPointer(&inst.ObstacleFree, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Status":
						__gong__assignPointer(&inst.Status, rhs, identifierMap)
					case "ProtectedRunwayDirection":
						__gong__assignPointer(&inst.ProtectedRunwayDirection, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayProtectAreaType:
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
				case *aixmfeatures.RunwaySectionContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RunwaySectionContamination":
						__gong__assignPointer(&inst.RunwaySectionContamination, rhs, identifierMap)
					}
				case *aixmfeatures.RunwaySectionContaminationType:
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
					case "ObservationTime":
						__gong__assignPointer(&inst.ObservationTime, rhs, identifierMap)
					case "Depth":
						__gong__assignPointer(&inst.Depth, rhs, identifierMap)
					case "FrictionCoefficient":
						__gong__assignPointer(&inst.FrictionCoefficient, rhs, identifierMap)
					case "FrictionEstimation":
						__gong__assignPointer(&inst.FrictionEstimation, rhs, identifierMap)
					case "FrictionDevice":
						__gong__assignPointer(&inst.FrictionDevice, rhs, identifierMap)
					case "ObscuredLights":
						__gong__assignPointer(&inst.ObscuredLights, rhs, identifierMap)
					case "FurtherClearanceTime":
						__gong__assignPointer(&inst.FurtherClearanceTime, rhs, identifierMap)
					case "FurtherTotalClearance":
						__gong__assignPointer(&inst.FurtherTotalClearance, rhs, identifierMap)
					case "NextObservationTime":
						__gong__assignPointer(&inst.NextObservationTime, rhs, identifierMap)
					case "Proportion":
						__gong__assignPointer(&inst.Proportion, rhs, identifierMap)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Section":
						__gong__assignPointer(&inst.Section, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayTimeSlice":
						__gong__assignPointer(&inst.RunwayTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayTimeSliceType:
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
					case "NominalLength":
						__gong__assignPointer(&inst.NominalLength, rhs, identifierMap)
					case "LengthAccuracy":
						__gong__assignPointer(&inst.LengthAccuracy, rhs, identifierMap)
					case "NominalWidth":
						__gong__assignPointer(&inst.NominalWidth, rhs, identifierMap)
					case "WidthAccuracy":
						__gong__assignPointer(&inst.WidthAccuracy, rhs, identifierMap)
					case "WidthShoulder":
						__gong__assignPointer(&inst.WidthShoulder, rhs, identifierMap)
					case "LengthStrip":
						__gong__assignPointer(&inst.LengthStrip, rhs, identifierMap)
					case "WidthStrip":
						__gong__assignPointer(&inst.WidthStrip, rhs, identifierMap)
					case "LengthOffset":
						__gong__assignPointer(&inst.LengthOffset, rhs, identifierMap)
					case "WidthOffset":
						__gong__assignPointer(&inst.WidthOffset, rhs, identifierMap)
					case "Abandoned":
						__gong__assignPointer(&inst.Abandoned, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "OverallContaminant":
						__gong__assignSliceOfPointers(&inst.OverallContaminant, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "AreaContaminant":
						__gong__assignSliceOfPointers(&inst.AreaContaminant, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayType:
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
				case *aixmfeatures.RunwayVisualRangeTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RunwayVisualRangeTimeSlice":
						__gong__assignPointer(&inst.RunwayVisualRangeTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayVisualRangeTimeSliceType:
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
					case "ReadingPosition":
						__gong__assignPointer(&inst.ReadingPosition, rhs, identifierMap)
					case "AssociatedRunwayDirection":
						__gong__assignSliceOfPointers(&inst.AssociatedRunwayDirection, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RunwayVisualRangeType:
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
				case *aixmfeatures.SeaplaneLandingAreaPropertyType:
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
				case *aixmfeatures.SeaplaneLandingAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SeaplaneLandingAreaTimeSlice":
						__gong__assignPointer(&inst.SeaplaneLandingAreaTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.SeaplaneLandingAreaTimeSliceType:
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
					case "RampSite":
						__gong__assignSliceOfPointers(&inst.RampSite, rhs, identifierMap)
					case "DockSite":
						__gong__assignSliceOfPointers(&inst.DockSite, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.SeaplaneLandingAreaType:
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
				case *aixmfeatures.SeaplaneRampSitePropertyType:
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
				case *aixmfeatures.SeaplaneRampSiteTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SeaplaneRampSiteTimeSlice":
						__gong__assignPointer(&inst.SeaplaneRampSiteTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.SeaplaneRampSiteTimeSliceType:
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
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Centreline":
						__gong__assignPointer(&inst.Centreline, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.SeaplaneRampSiteType:
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
				case *aixmfeatures.StandMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "StandMarkingTimeSlice":
						__gong__assignPointer(&inst.StandMarkingTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.StandMarkingTimeSliceType:
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
					case "MarkingICAOStandard":
						__gong__assignPointer(&inst.MarkingICAOStandard, rhs, identifierMap)
					case "Condition":
						__gong__assignPointer(&inst.Condition, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkedStand":
						__gong__assignPointer(&inst.MarkedStand, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.StandMarkingType:
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
				case *aixmfeatures.SurfaceCharacteristicsPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "SurfaceCharacteristics":
						__gong__assignPointer(&inst.SurfaceCharacteristics, rhs, identifierMap)
					}
				case *aixmfeatures.SurfaceCharacteristicsType:
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
					case "Composition":
						__gong__assignPointer(&inst.Composition, rhs, identifierMap)
					case "Preparation":
						__gong__assignPointer(&inst.Preparation, rhs, identifierMap)
					case "SurfaceCondition":
						__gong__assignPointer(&inst.SurfaceCondition, rhs, identifierMap)
					case "ClassPCN":
						__gong__assignPointer(&inst.ClassPCN, rhs, identifierMap)
					case "PavementTypePCN":
						__gong__assignPointer(&inst.PavementTypePCN, rhs, identifierMap)
					case "PavementSubgradePCN":
						__gong__assignPointer(&inst.PavementSubgradePCN, rhs, identifierMap)
					case "MaxTyrePressurePCN":
						__gong__assignPointer(&inst.MaxTyrePressurePCN, rhs, identifierMap)
					case "EvaluationMethodPCN":
						__gong__assignPointer(&inst.EvaluationMethodPCN, rhs, identifierMap)
					case "ClassLCN":
						__gong__assignPointer(&inst.ClassLCN, rhs, identifierMap)
					case "WeightSIWL":
						__gong__assignPointer(&inst.WeightSIWL, rhs, identifierMap)
					case "TyrePressureSIWL":
						__gong__assignPointer(&inst.TyrePressureSIWL, rhs, identifierMap)
					case "WeightAUW":
						__gong__assignPointer(&inst.WeightAUW, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.SurfaceContaminationLayerPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "SurfaceContaminationLayer":
						__gong__assignPointer(&inst.SurfaceContaminationLayer, rhs, identifierMap)
					}
				case *aixmfeatures.SurfaceContaminationLayerType:
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
					case "LayerOrder":
						__gong__assignPointer(&inst.LayerOrder, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Extent":
						__gong__assignSliceOfPointers(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *aixmfeatures.SurveyControlPointTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SurveyControlPointTimeSlice":
						__gong__assignPointer(&inst.SurveyControlPointTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.SurveyControlPointTimeSliceType:
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
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.SurveyControlPointType:
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
				case *aixmfeatures.TaxiHoldingPositionLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiHoldingPositionLightSystemTimeSlice":
						__gong__assignPointer(&inst.TaxiHoldingPositionLightSystemTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiHoldingPositionLightSystemTimeSliceType:
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
					case "EmergencyLighting":
						__gong__assignPointer(&inst.EmergencyLighting, rhs, identifierMap)
					case "IntensityLevel":
						__gong__assignPointer(&inst.IntensityLevel, rhs, identifierMap)
					case "Colour":
						__gong__assignPointer(&inst.Colour, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "TaxiHolding":
						__gong__assignPointer(&inst.TaxiHolding, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiHoldingPositionLightSystemType:
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
				case *aixmfeatures.TaxiHoldingPositionMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiHoldingPositionMarkingTimeSlice":
						__gong__assignPointer(&inst.TaxiHoldingPositionMarkingTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiHoldingPositionMarkingTimeSliceType:
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
					case "MarkingICAOStandard":
						__gong__assignPointer(&inst.MarkingICAOStandard, rhs, identifierMap)
					case "Condition":
						__gong__assignPointer(&inst.Condition, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkedTaxiHold":
						__gong__assignPointer(&inst.MarkedTaxiHold, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiHoldingPositionMarkingType:
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
				case *aixmfeatures.TaxiHoldingPositionPropertyType:
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
				case *aixmfeatures.TaxiHoldingPositionTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiHoldingPositionTimeSlice":
						__gong__assignPointer(&inst.TaxiHoldingPositionTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiHoldingPositionTimeSliceType:
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
					case "LandingCategory":
						__gong__assignPointer(&inst.LandingCategory, rhs, identifierMap)
					case "Status":
						__gong__assignPointer(&inst.Status, rhs, identifierMap)
					case "AssociatedGuidanceLine":
						__gong__assignPointer(&inst.AssociatedGuidanceLine, rhs, identifierMap)
					case "ProtectedRunway":
						__gong__assignSliceOfPointers(&inst.ProtectedRunway, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiHoldingPositionType:
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
				case *aixmfeatures.TaxiwayContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "TaxiwayContamination":
						__gong__assignPointer(&inst.TaxiwayContamination, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiwayContaminationType:
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
					case "ObservationTime":
						__gong__assignPointer(&inst.ObservationTime, rhs, identifierMap)
					case "Depth":
						__gong__assignPointer(&inst.Depth, rhs, identifierMap)
					case "FrictionCoefficient":
						__gong__assignPointer(&inst.FrictionCoefficient, rhs, identifierMap)
					case "FrictionEstimation":
						__gong__assignPointer(&inst.FrictionEstimation, rhs, identifierMap)
					case "FrictionDevice":
						__gong__assignPointer(&inst.FrictionDevice, rhs, identifierMap)
					case "ObscuredLights":
						__gong__assignPointer(&inst.ObscuredLights, rhs, identifierMap)
					case "FurtherClearanceTime":
						__gong__assignPointer(&inst.FurtherClearanceTime, rhs, identifierMap)
					case "FurtherTotalClearance":
						__gong__assignPointer(&inst.FurtherTotalClearance, rhs, identifierMap)
					case "NextObservationTime":
						__gong__assignPointer(&inst.NextObservationTime, rhs, identifierMap)
					case "Proportion":
						__gong__assignPointer(&inst.Proportion, rhs, identifierMap)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "ClearedWidth":
						__gong__assignPointer(&inst.ClearedWidth, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiwayElementPropertyType:
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
				case *aixmfeatures.TaxiwayElementTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiwayElementTimeSlice":
						__gong__assignPointer(&inst.TaxiwayElementTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiwayElementTimeSliceType:
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
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Width":
						__gong__assignPointer(&inst.Width, rhs, identifierMap)
					case "GradeSeparation":
						__gong__assignPointer(&inst.GradeSeparation, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AssociatedTaxiway":
						__gong__assignPointer(&inst.AssociatedTaxiway, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiwayElementType:
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
				case *aixmfeatures.TaxiwayLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiwayLightSystemTimeSlice":
						__gong__assignPointer(&inst.TaxiwayLightSystemTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiwayLightSystemTimeSliceType:
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
					case "EmergencyLighting":
						__gong__assignPointer(&inst.EmergencyLighting, rhs, identifierMap)
					case "IntensityLevel":
						__gong__assignPointer(&inst.IntensityLevel, rhs, identifierMap)
					case "Colour":
						__gong__assignPointer(&inst.Colour, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Position":
						__gong__assignPointer(&inst.Position, rhs, identifierMap)
					case "LightedTaxiway":
						__gong__assignPointer(&inst.LightedTaxiway, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiwayLightSystemType:
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
				case *aixmfeatures.TaxiwayMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiwayMarkingTimeSlice":
						__gong__assignPointer(&inst.TaxiwayMarkingTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiwayMarkingTimeSliceType:
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
					case "MarkingICAOStandard":
						__gong__assignPointer(&inst.MarkingICAOStandard, rhs, identifierMap)
					case "Condition":
						__gong__assignPointer(&inst.Condition, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkingLocation":
						__gong__assignPointer(&inst.MarkingLocation, rhs, identifierMap)
					case "MarkedTaxiway":
						__gong__assignPointer(&inst.MarkedTaxiway, rhs, identifierMap)
					case "MarkedElement":
						__gong__assignPointer(&inst.MarkedElement, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiwayMarkingType:
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
				case *aixmfeatures.TaxiwayPropertyType:
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
				case *aixmfeatures.TaxiwayTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TaxiwayTimeSlice":
						__gong__assignPointer(&inst.TaxiwayTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiwayTimeSliceType:
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
					case "Width":
						__gong__assignPointer(&inst.Width, rhs, identifierMap)
					case "WidthShoulder":
						__gong__assignPointer(&inst.WidthShoulder, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Abandoned":
						__gong__assignPointer(&inst.Abandoned, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "Contaminant":
						__gong__assignSliceOfPointers(&inst.Contaminant, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TaxiwayType:
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
				case *aixmfeatures.TouchDownLiftOffContaminationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "TouchDownLiftOffContamination":
						__gong__assignPointer(&inst.TouchDownLiftOffContamination, rhs, identifierMap)
					}
				case *aixmfeatures.TouchDownLiftOffContaminationType:
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
					case "ObservationTime":
						__gong__assignPointer(&inst.ObservationTime, rhs, identifierMap)
					case "Depth":
						__gong__assignPointer(&inst.Depth, rhs, identifierMap)
					case "FrictionCoefficient":
						__gong__assignPointer(&inst.FrictionCoefficient, rhs, identifierMap)
					case "FrictionEstimation":
						__gong__assignPointer(&inst.FrictionEstimation, rhs, identifierMap)
					case "FrictionDevice":
						__gong__assignPointer(&inst.FrictionDevice, rhs, identifierMap)
					case "ObscuredLights":
						__gong__assignPointer(&inst.ObscuredLights, rhs, identifierMap)
					case "FurtherClearanceTime":
						__gong__assignPointer(&inst.FurtherClearanceTime, rhs, identifierMap)
					case "FurtherTotalClearance":
						__gong__assignPointer(&inst.FurtherTotalClearance, rhs, identifierMap)
					case "NextObservationTime":
						__gong__assignPointer(&inst.NextObservationTime, rhs, identifierMap)
					case "Proportion":
						__gong__assignPointer(&inst.Proportion, rhs, identifierMap)
					case "CriticalRidge":
						__gong__assignSliceOfPointers(&inst.CriticalRidge, rhs, identifierMap)
					case "Layer":
						__gong__assignSliceOfPointers(&inst.Layer, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TouchDownLiftOffLightSystemTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TouchDownLiftOffLightSystemTimeSlice":
						__gong__assignPointer(&inst.TouchDownLiftOffLightSystemTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.TouchDownLiftOffLightSystemTimeSliceType:
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
					case "EmergencyLighting":
						__gong__assignPointer(&inst.EmergencyLighting, rhs, identifierMap)
					case "IntensityLevel":
						__gong__assignPointer(&inst.IntensityLevel, rhs, identifierMap)
					case "Colour":
						__gong__assignPointer(&inst.Colour, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Position":
						__gong__assignPointer(&inst.Position, rhs, identifierMap)
					case "LightedTouchDownLiftOff":
						__gong__assignPointer(&inst.LightedTouchDownLiftOff, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TouchDownLiftOffLightSystemType:
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
				case *aixmfeatures.TouchDownLiftOffMarkingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TouchDownLiftOffMarkingTimeSlice":
						__gong__assignPointer(&inst.TouchDownLiftOffMarkingTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.TouchDownLiftOffMarkingTimeSliceType:
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
					case "MarkingICAOStandard":
						__gong__assignPointer(&inst.MarkingICAOStandard, rhs, identifierMap)
					case "Condition":
						__gong__assignPointer(&inst.Condition, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "MarkingLocation":
						__gong__assignPointer(&inst.MarkingLocation, rhs, identifierMap)
					case "MarkedTouchDownLiftOff":
						__gong__assignPointer(&inst.MarkedTouchDownLiftOff, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TouchDownLiftOffMarkingType:
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
				case *aixmfeatures.TouchDownLiftOffSafeAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TouchDownLiftOffSafeAreaTimeSlice":
						__gong__assignPointer(&inst.TouchDownLiftOffSafeAreaTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.TouchDownLiftOffSafeAreaTimeSliceType:
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
					case "Width":
						__gong__assignPointer(&inst.Width, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Lighting":
						__gong__assignPointer(&inst.Lighting, rhs, identifierMap)
					case "ObstacleFree":
						__gong__assignPointer(&inst.ObstacleFree, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "ProtectedTouchDownLiftOff":
						__gong__assignPointer(&inst.ProtectedTouchDownLiftOff, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TouchDownLiftOffSafeAreaType:
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
				case *aixmfeatures.TouchDownLiftOffTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TouchDownLiftOffTimeSlice":
						__gong__assignPointer(&inst.TouchDownLiftOffTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.TouchDownLiftOffTimeSliceType:
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
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Width":
						__gong__assignPointer(&inst.Width, rhs, identifierMap)
					case "Slope":
						__gong__assignPointer(&inst.Slope, rhs, identifierMap)
					case "HelicopterClass":
						__gong__assignPointer(&inst.HelicopterClass, rhs, identifierMap)
					case "Abandoned":
						__gong__assignPointer(&inst.Abandoned, rhs, identifierMap)
					case "AimingPoint":
						__gong__assignPointer(&inst.AimingPoint, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "SurfaceProperties":
						__gong__assignPointer(&inst.SurfaceProperties, rhs, identifierMap)
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "ApproachTakeOffArea":
						__gong__assignPointer(&inst.ApproachTakeOffArea, rhs, identifierMap)
					case "Contaminant":
						__gong__assignSliceOfPointers(&inst.Contaminant, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TouchDownLiftOffType:
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
				case *aixmfeatures.VisualGlideSlopeIndicatorTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "VisualGlideSlopeIndicatorTimeSlice":
						__gong__assignPointer(&inst.VisualGlideSlopeIndicatorTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.VisualGlideSlopeIndicatorTimeSliceType:
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
					case "EmergencyLighting":
						__gong__assignPointer(&inst.EmergencyLighting, rhs, identifierMap)
					case "IntensityLevel":
						__gong__assignPointer(&inst.IntensityLevel, rhs, identifierMap)
					case "Colour":
						__gong__assignPointer(&inst.Colour, rhs, identifierMap)
					case "Element":
						__gong__assignSliceOfPointers(&inst.Element, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Position":
						__gong__assignPointer(&inst.Position, rhs, identifierMap)
					case "NumberBox":
						__gong__assignPointer(&inst.NumberBox, rhs, identifierMap)
					case "Portable":
						__gong__assignPointer(&inst.Portable, rhs, identifierMap)
					case "SlopeAngle":
						__gong__assignPointer(&inst.SlopeAngle, rhs, identifierMap)
					case "MinimumEyeHeightOverThreshold":
						__gong__assignPointer(&inst.MinimumEyeHeightOverThreshold, rhs, identifierMap)
					case "RunwayDirection":
						__gong__assignPointer(&inst.RunwayDirection, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.VisualGlideSlopeIndicatorType:
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
				case *aixmfeatures.WorkAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "WorkAreaTimeSlice":
						__gong__assignPointer(&inst.WorkAreaTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.WorkAreaTimeSliceType:
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
					case "PlannedOperational":
						__gong__assignPointer(&inst.PlannedOperational, rhs, identifierMap)
					case "AssociatedAirportHeliport":
						__gong__assignPointer(&inst.AssociatedAirportHeliport, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Activation":
						__gong__assignSliceOfPointers(&inst.Activation, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.WorkAreaType:
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
				case *aixmfeatures.WorkareaActivityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "WorkareaActivity":
						__gong__assignPointer(&inst.WorkareaActivity, rhs, identifierMap)
					}
				case *aixmfeatures.WorkareaActivityType:
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
					case "IsActive":
						__gong__assignPointer(&inst.IsActive, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
					case "AirportHeliport":
						__gong__assignPointer(&inst.AirportHeliport, rhs, identifierMap)
					case "AirportHeliportCollocation":
						__gong__assignPointer(&inst.AirportHeliportCollocation, rhs, identifierMap)
					case "AirportHotSpot":
						__gong__assignPointer(&inst.AirportHotSpot, rhs, identifierMap)
					case "AltimeterSource":
						__gong__assignPointer(&inst.AltimeterSource, rhs, identifierMap)
					case "NonMovementArea":
						__gong__assignPointer(&inst.NonMovementArea, rhs, identifierMap)
					case "SurveyControlPoint":
						__gong__assignPointer(&inst.SurveyControlPoint, rhs, identifierMap)
					case "WorkArea":
						__gong__assignPointer(&inst.WorkArea, rhs, identifierMap)
					case "Apron":
						__gong__assignPointer(&inst.Apron, rhs, identifierMap)
					case "ApronElement":
						__gong__assignPointer(&inst.ApronElement, rhs, identifierMap)
					case "AircraftStand":
						__gong__assignPointer(&inst.AircraftStand, rhs, identifierMap)
					case "DeicingArea":
						__gong__assignPointer(&inst.DeicingArea, rhs, identifierMap)
					case "PassengerLoadingBridge":
						__gong__assignPointer(&inst.PassengerLoadingBridge, rhs, identifierMap)
					case "Road":
						__gong__assignPointer(&inst.Road, rhs, identifierMap)
					case "TouchDownLiftOff":
						__gong__assignPointer(&inst.TouchDownLiftOff, rhs, identifierMap)
					case "TouchDownLiftOffSafeArea":
						__gong__assignPointer(&inst.TouchDownLiftOffSafeArea, rhs, identifierMap)
					case "Runway":
						__gong__assignPointer(&inst.Runway, rhs, identifierMap)
					case "RunwayElement":
						__gong__assignPointer(&inst.RunwayElement, rhs, identifierMap)
					case "RunwayDirection":
						__gong__assignPointer(&inst.RunwayDirection, rhs, identifierMap)
					case "RunwayCentrelinePoint":
						__gong__assignPointer(&inst.RunwayCentrelinePoint, rhs, identifierMap)
					case "RunwayBlastPad":
						__gong__assignPointer(&inst.RunwayBlastPad, rhs, identifierMap)
					case "RunwayProtectArea":
						__gong__assignPointer(&inst.RunwayProtectArea, rhs, identifierMap)
					case "RunwayVisualRange":
						__gong__assignPointer(&inst.RunwayVisualRange, rhs, identifierMap)
					case "VisualGlideSlopeIndicator":
						__gong__assignPointer(&inst.VisualGlideSlopeIndicator, rhs, identifierMap)
					case "ArrestingGear":
						__gong__assignPointer(&inst.ArrestingGear, rhs, identifierMap)
					case "Taxiway":
						__gong__assignPointer(&inst.Taxiway, rhs, identifierMap)
					case "TaxiwayElement":
						__gong__assignPointer(&inst.TaxiwayElement, rhs, identifierMap)
					case "TaxiHoldingPosition":
						__gong__assignPointer(&inst.TaxiHoldingPosition, rhs, identifierMap)
					case "GuidanceLine":
						__gong__assignPointer(&inst.GuidanceLine, rhs, identifierMap)
					case "SeaplaneLandingArea":
						__gong__assignPointer(&inst.SeaplaneLandingArea, rhs, identifierMap)
					case "SeaplaneRampSite":
						__gong__assignPointer(&inst.SeaplaneRampSite, rhs, identifierMap)
					case "FloatingDockSite":
						__gong__assignPointer(&inst.FloatingDockSite, rhs, identifierMap)
					case "MarkingBuoy":
						__gong__assignPointer(&inst.MarkingBuoy, rhs, identifierMap)
					case "ApproachLightingSystem":
						__gong__assignPointer(&inst.ApproachLightingSystem, rhs, identifierMap)
					case "RunwayDirectionLightSystem":
						__gong__assignPointer(&inst.RunwayDirectionLightSystem, rhs, identifierMap)
					case "RunwayProtectAreaLightSystem":
						__gong__assignPointer(&inst.RunwayProtectAreaLightSystem, rhs, identifierMap)
					case "TaxiwayLightSystem":
						__gong__assignPointer(&inst.TaxiwayLightSystem, rhs, identifierMap)
					case "TaxiHoldingPositionLightSystem":
						__gong__assignPointer(&inst.TaxiHoldingPositionLightSystem, rhs, identifierMap)
					case "GuidanceLineLightSystem":
						__gong__assignPointer(&inst.GuidanceLineLightSystem, rhs, identifierMap)
					case "ApronLightSystem":
						__gong__assignPointer(&inst.ApronLightSystem, rhs, identifierMap)
					case "TouchDownLiftOffLightSystem":
						__gong__assignPointer(&inst.TouchDownLiftOffLightSystem, rhs, identifierMap)
					case "RunwayMarking":
						__gong__assignPointer(&inst.RunwayMarking, rhs, identifierMap)
					case "TaxiwayMarking":
						__gong__assignPointer(&inst.TaxiwayMarking, rhs, identifierMap)
					case "TaxiHoldingPositionMarking":
						__gong__assignPointer(&inst.TaxiHoldingPositionMarking, rhs, identifierMap)
					case "StandMarking":
						__gong__assignPointer(&inst.StandMarking, rhs, identifierMap)
					case "ApronMarking":
						__gong__assignPointer(&inst.ApronMarking, rhs, identifierMap)
					case "DeicingAreaMarking":
						__gong__assignPointer(&inst.DeicingAreaMarking, rhs, identifierMap)
					case "GuidanceLineMarking":
						__gong__assignPointer(&inst.GuidanceLineMarking, rhs, identifierMap)
					case "TouchDownLiftOffMarking":
						__gong__assignPointer(&inst.TouchDownLiftOffMarking, rhs, identifierMap)
					case "AirportProtectionAreaMarking":
						__gong__assignPointer(&inst.AirportProtectionAreaMarking, rhs, identifierMap)
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
