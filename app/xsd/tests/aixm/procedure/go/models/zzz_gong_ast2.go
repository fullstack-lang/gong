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
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/xlink"
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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/xlink":
			aliasToCanonical[alias] = "xlink"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/gml":
			aliasToCanonical[alias] = "gml"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmabstractgmlobjecttypes":
			aliasToCanonical[alias] = "aixmabstractgmlobjecttypes"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmdatatypes":
			aliasToCanonical[alias] = "aixmdatatypes"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models/aixmfeatures":
			aliasToCanonical[alias] = "aixmfeatures"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/procedure/go/models":
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
				case "CodeApproachEquipmentAdditionalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeApproachEquipmentAdditionalType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeApproachGuidanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeApproachGuidanceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeApproachPrefixType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeApproachPrefixType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeApproachType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeApproachType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeApprovalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeApprovalType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeArcDirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeArcDirectionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeBearingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeBearingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeCommunicationModeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeCommunicationModeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeCourseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeCourseType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeDayType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeDesignStandardType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeDesignStandardType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeDesignatedPointDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeDesignatedPointDesignatorType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeDirectionReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeDirectionReferenceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeDirectionTurnType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeDirectionTurnType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeEquipmentAntiCollisionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeEquipmentAntiCollisionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeEquipmentUnavailableType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeEquipmentUnavailableType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFinalGuidanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFinalGuidanceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFreeFlightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFreeFlightType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeHeightReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeHeightReferenceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeHoldingUsageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeHoldingUsageType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeHoldingUseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeHoldingUseType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeICAOCountryType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeICAOCountryType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryRoutePointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMilitaryRoutePointType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMinimaFinalApproachPathType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMinimaFinalApproachPathType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMinimumAltitudeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMinimumAltitudeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMinimumHeightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMinimumHeightType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMissedApproachType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMissedApproachType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeNavigationAreaRestrictionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeNavigationAreaRestrictionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeNavigationAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeNavigationAreaType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
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
				case "CodeProcedureAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeProcedureAvailabilityType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeProcedureCodingStandardType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeProcedureCodingStandardType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeProcedureDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeProcedureDistanceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeProcedureFixRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeProcedureFixRoleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeProcedurePhaseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeProcedurePhaseType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRNPType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRNPType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRVSMPointRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRVSMPointRoleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRVSMType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRVSMType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeReferencePathIdentifierType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeReferencePathIdentifierType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeReferenceRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeReferenceRoleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRelativePositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRelativePositionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteIndicatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRouteIndicatorType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeSafeAltitudeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeSafeAltitudeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeSegmentPathType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeSegmentPathType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeSegmentTerminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeSegmentTerminationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeSideType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeSideType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeSpeedReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeSpeedReferenceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTAAType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTAAType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventCombinationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTimeEventCombinationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTimeEventType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTimeReferenceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTrajectoryType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTrajectoryType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
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
				case "TextInstructionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextInstructionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextNameType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextNameType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextNoteType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextPropertyNameType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextPropertyNameType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TextSIDSTARDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TextSIDSTARDesignatorType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "TimeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.TimeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValAlarmLimitType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValAlarmLimitType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValAngleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValAngleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValBearingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValBearingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValChannelNumberType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValChannelNumberType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceSignedType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDistanceSignedType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDistanceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceVerticalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDistanceVerticalType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValDurationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValDurationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "ValHexType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValHexType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
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
				case "A_extension_6":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_6), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_7":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_7), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_8":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_8), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "A_extension_9":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.A_extension_9), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractApproachLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractApproachLegType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractProcedureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractProcedureType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractPropertiesWithScheduleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractPropertiesWithScheduleType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractSegmentLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractSegmentLegType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractSegmentPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractSegmentPointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AerialRefuellingPointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftCharacteristicPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftCharacteristicType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftCharacteristicType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AltimeterSourcePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AltimeterSourcePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
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
				case "ApproachAltitudeTablePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApproachAltitudeTablePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApproachAltitudeTableType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApproachAltitudeTableType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApproachConditionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApproachConditionPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApproachConditionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApproachConditionType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApproachDistanceTablePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApproachDistanceTablePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApproachDistanceTableType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApproachDistanceTableType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApproachTimingTablePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApproachTimingTablePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ApproachTimingTableType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ApproachTimingTableType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ArrivalFeederLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ArrivalFeederLegTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ArrivalFeederLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ArrivalFeederLegTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ArrivalFeederLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ArrivalFeederLegType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ArrivalLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ArrivalLegTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ArrivalLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ArrivalLegTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ArrivalLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ArrivalLegType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CircleSectorPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CircleSectorPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CircleSectorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CircleSectorType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CirclingAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CirclingAreaTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CirclingAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CirclingAreaTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CirclingAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CirclingAreaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CirclingRestrictionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CirclingRestrictionPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CirclingRestrictionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CirclingRestrictionType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CurvePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CurvePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CurveType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DMEPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DMEPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DepartureArrivalConditionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DepartureArrivalConditionPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DepartureArrivalConditionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DepartureArrivalConditionType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DepartureLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DepartureLegTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DepartureLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DepartureLegTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DepartureLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DepartureLegType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DesignatedPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DesignatedPointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DistanceIndicationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DistanceIndicationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedCurveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedCurveType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedPointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedSurfaceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "EnRouteSegmentPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.EnRouteSegmentPointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "EquipmentUnavailableAdjustmentColumnPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.EquipmentUnavailableAdjustmentColumnPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "EquipmentUnavailableAdjustmentColumnType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.EquipmentUnavailableAdjustmentColumnType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "EquipmentUnavailableAdjustmentPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.EquipmentUnavailableAdjustmentPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "EquipmentUnavailableAdjustmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.EquipmentUnavailableAdjustmentType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FASDataBlockPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FASDataBlockPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FASDataBlockType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FASDataBlockType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FinalLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FinalLegTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FinalLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FinalLegTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FinalLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FinalLegType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FinalProfilePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FinalProfilePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FinalProfileType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FinalProfileType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingAssessmentTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingAssessmentTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingAssessmentTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingAssessmentTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingAssessmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingAssessmentType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternDistancePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingPatternDistancePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingPatternDistanceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternDurationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingPatternDurationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternDurationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingPatternDurationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingPatternPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingPatternTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingPatternTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingPatternType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingUsePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingUsePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingUseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingUseType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "InitialLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.InitialLegTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "InitialLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.InitialLegTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "InitialLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.InitialLegType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "InstrumentApproachProcedurePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.InstrumentApproachProcedurePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "InstrumentApproachProcedureTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.InstrumentApproachProcedureTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "InstrumentApproachProcedureTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.InstrumentApproachProcedureTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "InstrumentApproachProcedureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.InstrumentApproachProcedureType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "IntermediateLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.IntermediateLegTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "IntermediateLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.IntermediateLegTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "IntermediateLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.IntermediateLegType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LandingTakeoffAreaCollectionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LandingTakeoffAreaCollectionPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LandingTakeoffAreaCollectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LandingTakeoffAreaCollectionType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LinguisticNotePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LinguisticNoteType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MinimaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MinimaPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MinimaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MinimaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MissedApproachGroupPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MissedApproachGroupPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MissedApproachGroupType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MissedApproachGroupType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MissedApproachLegTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MissedApproachLegTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MissedApproachLegTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MissedApproachLegTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "MissedApproachLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.MissedApproachLegType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavaidPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavaidPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaRestrictionTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavigationAreaRestrictionTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaRestrictionTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavigationAreaRestrictionTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaRestrictionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavigationAreaRestrictionType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaSectorPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavigationAreaSectorPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaSectorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavigationAreaSectorType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavigationAreaTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavigationAreaTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavigationAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavigationAreaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
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
				case "ProcedureAvailabilityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ProcedureAvailabilityPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ProcedureAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ProcedureAvailabilityType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ProcedureDMETimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ProcedureDMETimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ProcedureDMETimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ProcedureDMETimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ProcedureDMEType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ProcedureDMEType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ProcedurePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ProcedurePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ProcedureTransitionLegPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ProcedureTransitionLegPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ProcedureTransitionLegType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ProcedureTransitionLegType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ProcedureTransitionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ProcedureTransitionPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ProcedureTransitionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ProcedureTransitionType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RadarSystemPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RadarSystemPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RadioFrequencyAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RadioFrequencyAreaPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayCentrelinePointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayDirectionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayDirectionPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SafeAltitudeAreaPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SafeAltitudeAreaPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SafeAltitudeAreaSectorPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SafeAltitudeAreaSectorPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SafeAltitudeAreaSectorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SafeAltitudeAreaSectorType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SafeAltitudeAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SafeAltitudeAreaTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SafeAltitudeAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SafeAltitudeAreaTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SafeAltitudeAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SafeAltitudeAreaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SectorDesignPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SectorDesignPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SectorDesignType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SectorDesignType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SegmentLegPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SegmentLegPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SegmentPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SegmentPointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SpecialNavigationSystemPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SpecialNavigationSystemPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentArrivalPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardInstrumentArrivalPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentArrivalTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardInstrumentArrivalTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentArrivalTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardInstrumentArrivalTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentArrivalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardInstrumentArrivalType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentDeparturePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardInstrumentDeparturePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentDepartureTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardInstrumentDepartureTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentDepartureTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardInstrumentDepartureTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardInstrumentDepartureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardInstrumentDepartureType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurfacePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurfacePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SurfaceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TerminalArrivalAreaSectorPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TerminalArrivalAreaSectorPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TerminalArrivalAreaSectorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TerminalArrivalAreaSectorType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TerminalArrivalAreaTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TerminalArrivalAreaTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TerminalArrivalAreaTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TerminalArrivalAreaTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TerminalArrivalAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TerminalArrivalAreaType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TerminalSegmentPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TerminalSegmentPointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TerminalSegmentPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TerminalSegmentPointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TimesheetPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TimesheetPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TimesheetType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TimesheetType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "TouchDownLiftOffPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.TouchDownLiftOffPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "UnplannedHoldingPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.UnplannedHoldingPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "UnplannedHoldingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.UnplannedHoldingTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "UnplannedHoldingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.UnplannedHoldingTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "UnplannedHoldingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.UnplannedHoldingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
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
				case *aixmdatatypes.CodeApproachEquipmentAdditionalType:
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
				case *aixmdatatypes.CodeApproachPrefixType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeApproachType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeApprovalType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeArcDirectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeBearingType:
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
				case *aixmdatatypes.CodeDesignStandardType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeDesignatedPointDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeDirectionReferenceType:
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
				case *aixmdatatypes.CodeEquipmentAntiCollisionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeEquipmentUnavailableType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFinalGuidanceType:
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
				case *aixmdatatypes.CodeHeightReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeHoldingUsageType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeHoldingUseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeICAOCountryType:
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
				case *aixmdatatypes.CodeMinimaFinalApproachPathType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeMinimumAltitudeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeMinimumHeightType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeMissedApproachType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeNavigationAreaRestrictionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeNavigationAreaType:
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
				case *aixmdatatypes.CodeProcedureAvailabilityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeProcedureCodingStandardType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeProcedureDistanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeProcedureFixRoleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeProcedurePhaseType:
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
				case *aixmdatatypes.CodeReferencePathIdentifierType:
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
				case *aixmdatatypes.CodeRelativePositionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeRouteIndicatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeSafeAltitudeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeSegmentPathType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeSegmentTerminationType:
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
				case *aixmdatatypes.CodeTAAType:
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
				case *aixmdatatypes.CodeTrajectoryType:
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
				case *aixmdatatypes.TextPropertyNameType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.TextSIDSTARDesignatorType:
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
				case *aixmdatatypes.ValAlarmLimitType:
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
				case *aixmdatatypes.ValChannelNumberType:
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
				case *aixmdatatypes.ValHexType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
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
					case "AbstractSegmentPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSegmentPointExtension, rhs, identifierMap)
					case "AbstractAerialRefuellingPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAerialRefuellingPointExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_1:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractElevatedCurveExtension":
						__gong__assignSliceOfPointers(&inst.AbstractElevatedCurveExtension, rhs, identifierMap)
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
					case "AbstractSegmentPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSegmentPointExtension, rhs, identifierMap)
					case "AbstractTerminalSegmentPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTerminalSegmentPointExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_12:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractLinguisticNoteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractLinguisticNoteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_13:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractNoteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractNoteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_14:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractApproachAltitudeTableExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApproachAltitudeTableExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_15:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractApproachConditionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApproachConditionExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_16:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractApproachDistanceTableExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApproachDistanceTableExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_17:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractApproachTimingTableExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApproachTimingTableExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_18:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFinalProfileExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFinalProfileExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_19:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractInstrumentApproachProcedureExtension":
						__gong__assignSliceOfPointers(&inst.AbstractInstrumentApproachProcedureExtension, rhs, identifierMap)
					case "AbstractProcedureExtension":
						__gong__assignSliceOfPointers(&inst.AbstractProcedureExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_2:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractElevatedPointExtension":
						__gong__assignSliceOfPointers(&inst.AbstractElevatedPointExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_20:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractMissedApproachGroupExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMissedApproachGroupExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_21:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTerminalArrivalAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTerminalArrivalAreaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_22:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTerminalArrivalAreaSectorExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTerminalArrivalAreaSectorExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_23:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFASDataBlockExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFASDataBlockExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_24:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractCirclingAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractCirclingAreaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_25:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractCirclingRestrictionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractCirclingRestrictionExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_26:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractEquipmentUnavailableAdjustmentExtension":
						__gong__assignSliceOfPointers(&inst.AbstractEquipmentUnavailableAdjustmentExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_27:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractEquipmentUnavailableAdjustmentColumnExtension":
						__gong__assignSliceOfPointers(&inst.AbstractEquipmentUnavailableAdjustmentColumnExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_28:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractMinimaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMinimaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_29:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractStandardInstrumentArrivalExtension":
						__gong__assignSliceOfPointers(&inst.AbstractStandardInstrumentArrivalExtension, rhs, identifierMap)
					case "AbstractProcedureExtension":
						__gong__assignSliceOfPointers(&inst.AbstractProcedureExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_3:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractElevatedSurfaceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractElevatedSurfaceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_30:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractDepartureArrivalConditionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractDepartureArrivalConditionExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_31:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractNavigationAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractNavigationAreaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_32:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractNavigationAreaSectorExtension":
						__gong__assignSliceOfPointers(&inst.AbstractNavigationAreaSectorExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_33:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSectorDesignExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSectorDesignExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_34:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractStandardInstrumentDepartureExtension":
						__gong__assignSliceOfPointers(&inst.AbstractStandardInstrumentDepartureExtension, rhs, identifierMap)
					case "AbstractProcedureExtension":
						__gong__assignSliceOfPointers(&inst.AbstractProcedureExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_35:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSafeAltitudeAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSafeAltitudeAreaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_36:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSafeAltitudeAreaSectorExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSafeAltitudeAreaSectorExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_37:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractArrivalFeederLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractArrivalFeederLegExtension, rhs, identifierMap)
					case "AbstractApproachLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApproachLegExtension, rhs, identifierMap)
					case "AbstractSegmentLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSegmentLegExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_38:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractArrivalLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractArrivalLegExtension, rhs, identifierMap)
					case "AbstractSegmentLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSegmentLegExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_39:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractDepartureLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractDepartureLegExtension, rhs, identifierMap)
					case "AbstractSegmentLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSegmentLegExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_4:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractHoldingPatternExtension":
						__gong__assignSliceOfPointers(&inst.AbstractHoldingPatternExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_40:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFinalLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFinalLegExtension, rhs, identifierMap)
					case "AbstractApproachLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApproachLegExtension, rhs, identifierMap)
					case "AbstractSegmentLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSegmentLegExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_41:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractHoldingUseExtension":
						__gong__assignSliceOfPointers(&inst.AbstractHoldingUseExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_42:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractInitialLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractInitialLegExtension, rhs, identifierMap)
					case "AbstractApproachLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApproachLegExtension, rhs, identifierMap)
					case "AbstractSegmentLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSegmentLegExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_43:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractIntermediateLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractIntermediateLegExtension, rhs, identifierMap)
					case "AbstractApproachLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApproachLegExtension, rhs, identifierMap)
					case "AbstractSegmentLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSegmentLegExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_44:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractLandingTakeoffAreaCollectionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractLandingTakeoffAreaCollectionExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_45:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractMissedApproachLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractMissedApproachLegExtension, rhs, identifierMap)
					case "AbstractApproachLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractApproachLegExtension, rhs, identifierMap)
					case "AbstractSegmentLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSegmentLegExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_46:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractNavigationAreaRestrictionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractNavigationAreaRestrictionExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_47:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractProcedureDMEExtension":
						__gong__assignSliceOfPointers(&inst.AbstractProcedureDMEExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_48:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractProcedureTransitionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractProcedureTransitionExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_49:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractProcedureTransitionLegExtension":
						__gong__assignSliceOfPointers(&inst.AbstractProcedureTransitionLegExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_5:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractHoldingPatternDistanceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractHoldingPatternDistanceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_50:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractProcedureAvailabilityExtension":
						__gong__assignSliceOfPointers(&inst.AbstractProcedureAvailabilityExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_51:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAircraftCharacteristicExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAircraftCharacteristicExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_52:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractCircleSectorExtension":
						__gong__assignSliceOfPointers(&inst.AbstractCircleSectorExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_53:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTimesheetExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTimesheetExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_54:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAltitudeAdjustmentExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAltitudeAdjustmentExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_55:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractHoldingAssessmentExtension":
						__gong__assignSliceOfPointers(&inst.AbstractHoldingAssessmentExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_56:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractObstacleAssessmentAreaExtension":
						__gong__assignSliceOfPointers(&inst.AbstractObstacleAssessmentAreaExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_57:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractObstaclePlacementExtension":
						__gong__assignSliceOfPointers(&inst.AbstractObstaclePlacementExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_58:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractObstructionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractObstructionExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_6:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractHoldingPatternDurationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractHoldingPatternDurationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_7:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractUnplannedHoldingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractUnplannedHoldingExtension, rhs, identifierMap)
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
				case *aixmfeatures.AbstractApproachLegType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.AbstractProcedureType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.AbstractSegmentLegType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.ApproachAltitudeTablePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApproachAltitudeTable":
						__gong__assignPointer(&inst.ApproachAltitudeTable, rhs, identifierMap)
					}
				case *aixmfeatures.ApproachAltitudeTableType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "MeasurementPoint":
						__gong__assignPointer(&inst.MeasurementPoint, rhs, identifierMap)
					case "Altitude":
						__gong__assignPointer(&inst.Altitude, rhs, identifierMap)
					case "AltitudeReference":
						__gong__assignPointer(&inst.AltitudeReference, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ApproachConditionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApproachCondition":
						__gong__assignPointer(&inst.ApproachCondition, rhs, identifierMap)
					}
				case *aixmfeatures.ApproachConditionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "FinalApproachPath":
						__gong__assignPointer(&inst.FinalApproachPath, rhs, identifierMap)
					case "RequiredNavigationPerformance":
						__gong__assignPointer(&inst.RequiredNavigationPerformance, rhs, identifierMap)
					case "ClimbGradient":
						__gong__assignPointer(&inst.ClimbGradient, rhs, identifierMap)
					case "MinimumSet":
						__gong__assignPointer(&inst.MinimumSet, rhs, identifierMap)
					case "CirclingRestriction":
						__gong__assignSliceOfPointers(&inst.CirclingRestriction, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "LandingArea":
						__gong__assignSliceOfPointers(&inst.LandingArea, rhs, identifierMap)
					case "Altimeter":
						__gong__assignPointer(&inst.Altimeter, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ApproachDistanceTablePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApproachDistanceTable":
						__gong__assignPointer(&inst.ApproachDistanceTable, rhs, identifierMap)
					}
				case *aixmfeatures.ApproachDistanceTableType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "StartingMeasurementPoint":
						__gong__assignPointer(&inst.StartingMeasurementPoint, rhs, identifierMap)
					case "ValueHAT":
						__gong__assignPointer(&inst.ValueHAT, rhs, identifierMap)
					case "EndingMeasurementPoint":
						__gong__assignPointer(&inst.EndingMeasurementPoint, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ApproachTimingTablePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ApproachTimingTable":
						__gong__assignPointer(&inst.ApproachTimingTable, rhs, identifierMap)
					}
				case *aixmfeatures.ApproachTimingTableType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "StartingMeasurementPoint":
						__gong__assignPointer(&inst.StartingMeasurementPoint, rhs, identifierMap)
					case "EndingMeasurementPoint":
						__gong__assignPointer(&inst.EndingMeasurementPoint, rhs, identifierMap)
					case "Time":
						__gong__assignPointer(&inst.Time, rhs, identifierMap)
					case "Speed":
						__gong__assignPointer(&inst.Speed, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ArrivalFeederLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ArrivalFeederLegTimeSlice":
						__gong__assignPointer(&inst.ArrivalFeederLegTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.ArrivalFeederLegTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "EndConditionDesignator":
						__gong__assignPointer(&inst.EndConditionDesignator, rhs, identifierMap)
					case "LegPath":
						__gong__assignPointer(&inst.LegPath, rhs, identifierMap)
					case "LegTypeARINC":
						__gong__assignPointer(&inst.LegTypeARINC, rhs, identifierMap)
					case "Course":
						__gong__assignPointer(&inst.Course, rhs, identifierMap)
					case "CourseType":
						__gong__assignPointer(&inst.CourseType, rhs, identifierMap)
					case "CourseDirection":
						__gong__assignPointer(&inst.CourseDirection, rhs, identifierMap)
					case "TurnDirection":
						__gong__assignPointer(&inst.TurnDirection, rhs, identifierMap)
					case "SpeedLimit":
						__gong__assignPointer(&inst.SpeedLimit, rhs, identifierMap)
					case "SpeedReference":
						__gong__assignPointer(&inst.SpeedReference, rhs, identifierMap)
					case "SpeedInterpretation":
						__gong__assignPointer(&inst.SpeedInterpretation, rhs, identifierMap)
					case "BankAngle":
						__gong__assignPointer(&inst.BankAngle, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Duration":
						__gong__assignPointer(&inst.Duration, rhs, identifierMap)
					case "ProcedureTurnRequired":
						__gong__assignPointer(&inst.ProcedureTurnRequired, rhs, identifierMap)
					case "UpperLimitAltitude":
						__gong__assignPointer(&inst.UpperLimitAltitude, rhs, identifierMap)
					case "UpperLimitReference":
						__gong__assignPointer(&inst.UpperLimitReference, rhs, identifierMap)
					case "LowerLimitAltitude":
						__gong__assignPointer(&inst.LowerLimitAltitude, rhs, identifierMap)
					case "LowerLimitReference":
						__gong__assignPointer(&inst.LowerLimitReference, rhs, identifierMap)
					case "AltitudeInterpretation":
						__gong__assignPointer(&inst.AltitudeInterpretation, rhs, identifierMap)
					case "AltitudeOverrideATC":
						__gong__assignPointer(&inst.AltitudeOverrideATC, rhs, identifierMap)
					case "AltitudeOverrideReference":
						__gong__assignPointer(&inst.AltitudeOverrideReference, rhs, identifierMap)
					case "VerticalAngle":
						__gong__assignPointer(&inst.VerticalAngle, rhs, identifierMap)
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Approach":
						__gong__assignPointer(&inst.Approach, rhs, identifierMap)
					case "RequiredNavigationPerformance":
						__gong__assignPointer(&inst.RequiredNavigationPerformance, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ArrivalFeederLegType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.ArrivalLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ArrivalLegTimeSlice":
						__gong__assignPointer(&inst.ArrivalLegTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.ArrivalLegTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "EndConditionDesignator":
						__gong__assignPointer(&inst.EndConditionDesignator, rhs, identifierMap)
					case "LegPath":
						__gong__assignPointer(&inst.LegPath, rhs, identifierMap)
					case "LegTypeARINC":
						__gong__assignPointer(&inst.LegTypeARINC, rhs, identifierMap)
					case "Course":
						__gong__assignPointer(&inst.Course, rhs, identifierMap)
					case "CourseType":
						__gong__assignPointer(&inst.CourseType, rhs, identifierMap)
					case "CourseDirection":
						__gong__assignPointer(&inst.CourseDirection, rhs, identifierMap)
					case "TurnDirection":
						__gong__assignPointer(&inst.TurnDirection, rhs, identifierMap)
					case "SpeedLimit":
						__gong__assignPointer(&inst.SpeedLimit, rhs, identifierMap)
					case "SpeedReference":
						__gong__assignPointer(&inst.SpeedReference, rhs, identifierMap)
					case "SpeedInterpretation":
						__gong__assignPointer(&inst.SpeedInterpretation, rhs, identifierMap)
					case "BankAngle":
						__gong__assignPointer(&inst.BankAngle, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Duration":
						__gong__assignPointer(&inst.Duration, rhs, identifierMap)
					case "ProcedureTurnRequired":
						__gong__assignPointer(&inst.ProcedureTurnRequired, rhs, identifierMap)
					case "UpperLimitAltitude":
						__gong__assignPointer(&inst.UpperLimitAltitude, rhs, identifierMap)
					case "UpperLimitReference":
						__gong__assignPointer(&inst.UpperLimitReference, rhs, identifierMap)
					case "LowerLimitAltitude":
						__gong__assignPointer(&inst.LowerLimitAltitude, rhs, identifierMap)
					case "LowerLimitReference":
						__gong__assignPointer(&inst.LowerLimitReference, rhs, identifierMap)
					case "AltitudeInterpretation":
						__gong__assignPointer(&inst.AltitudeInterpretation, rhs, identifierMap)
					case "AltitudeOverrideATC":
						__gong__assignPointer(&inst.AltitudeOverrideATC, rhs, identifierMap)
					case "AltitudeOverrideReference":
						__gong__assignPointer(&inst.AltitudeOverrideReference, rhs, identifierMap)
					case "VerticalAngle":
						__gong__assignPointer(&inst.VerticalAngle, rhs, identifierMap)
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "RequiredNavigationPerformance":
						__gong__assignPointer(&inst.RequiredNavigationPerformance, rhs, identifierMap)
					case "Arrival":
						__gong__assignPointer(&inst.Arrival, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ArrivalLegType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.CircleSectorPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "CircleSector":
						__gong__assignPointer(&inst.CircleSector, rhs, identifierMap)
					}
				case *aixmfeatures.CircleSectorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "ArcDirection":
						__gong__assignPointer(&inst.ArcDirection, rhs, identifierMap)
					case "FromAngle":
						__gong__assignPointer(&inst.FromAngle, rhs, identifierMap)
					case "ToAngle":
						__gong__assignPointer(&inst.ToAngle, rhs, identifierMap)
					case "AngleType":
						__gong__assignPointer(&inst.AngleType, rhs, identifierMap)
					case "AngleDirectionReference":
						__gong__assignPointer(&inst.AngleDirectionReference, rhs, identifierMap)
					case "InnerDistance":
						__gong__assignPointer(&inst.InnerDistance, rhs, identifierMap)
					case "OuterDistance":
						__gong__assignPointer(&inst.OuterDistance, rhs, identifierMap)
					case "UpperLimit":
						__gong__assignPointer(&inst.UpperLimit, rhs, identifierMap)
					case "UpperLimitReference":
						__gong__assignPointer(&inst.UpperLimitReference, rhs, identifierMap)
					case "LowerLimit":
						__gong__assignPointer(&inst.LowerLimit, rhs, identifierMap)
					case "LowerLimitReference":
						__gong__assignPointer(&inst.LowerLimitReference, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.CirclingAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "CirclingAreaTimeSlice":
						__gong__assignPointer(&inst.CirclingAreaTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.CirclingAreaTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "Approach":
						__gong__assignPointer(&inst.Approach, rhs, identifierMap)
					case "Condition":
						__gong__assignSliceOfPointers(&inst.Condition, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignPointer(&inst.AircraftCategory, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.CirclingAreaType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.CirclingRestrictionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "CirclingRestriction":
						__gong__assignPointer(&inst.CirclingRestriction, rhs, identifierMap)
					}
				case *aixmfeatures.CirclingRestrictionType:
					switch fieldName {
					case "Name":
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
					case "SectorDescription":
						__gong__assignPointer(&inst.SectorDescription, rhs, identifierMap)
					case "RestrictionArea":
						__gong__assignPointer(&inst.RestrictionArea, rhs, identifierMap)
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
				case *aixmfeatures.DepartureArrivalConditionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "DepartureArrivalCondition":
						__gong__assignPointer(&inst.DepartureArrivalCondition, rhs, identifierMap)
					}
				case *aixmfeatures.DepartureArrivalConditionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
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
					case "EngineType":
						__gong__assignPointer(&inst.EngineType, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.DepartureLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "DepartureLegTimeSlice":
						__gong__assignPointer(&inst.DepartureLegTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.DepartureLegTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "EndConditionDesignator":
						__gong__assignPointer(&inst.EndConditionDesignator, rhs, identifierMap)
					case "LegPath":
						__gong__assignPointer(&inst.LegPath, rhs, identifierMap)
					case "LegTypeARINC":
						__gong__assignPointer(&inst.LegTypeARINC, rhs, identifierMap)
					case "Course":
						__gong__assignPointer(&inst.Course, rhs, identifierMap)
					case "CourseType":
						__gong__assignPointer(&inst.CourseType, rhs, identifierMap)
					case "CourseDirection":
						__gong__assignPointer(&inst.CourseDirection, rhs, identifierMap)
					case "TurnDirection":
						__gong__assignPointer(&inst.TurnDirection, rhs, identifierMap)
					case "SpeedLimit":
						__gong__assignPointer(&inst.SpeedLimit, rhs, identifierMap)
					case "SpeedReference":
						__gong__assignPointer(&inst.SpeedReference, rhs, identifierMap)
					case "SpeedInterpretation":
						__gong__assignPointer(&inst.SpeedInterpretation, rhs, identifierMap)
					case "BankAngle":
						__gong__assignPointer(&inst.BankAngle, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Duration":
						__gong__assignPointer(&inst.Duration, rhs, identifierMap)
					case "ProcedureTurnRequired":
						__gong__assignPointer(&inst.ProcedureTurnRequired, rhs, identifierMap)
					case "UpperLimitAltitude":
						__gong__assignPointer(&inst.UpperLimitAltitude, rhs, identifierMap)
					case "UpperLimitReference":
						__gong__assignPointer(&inst.UpperLimitReference, rhs, identifierMap)
					case "LowerLimitAltitude":
						__gong__assignPointer(&inst.LowerLimitAltitude, rhs, identifierMap)
					case "LowerLimitReference":
						__gong__assignPointer(&inst.LowerLimitReference, rhs, identifierMap)
					case "AltitudeInterpretation":
						__gong__assignPointer(&inst.AltitudeInterpretation, rhs, identifierMap)
					case "AltitudeOverrideATC":
						__gong__assignPointer(&inst.AltitudeOverrideATC, rhs, identifierMap)
					case "AltitudeOverrideReference":
						__gong__assignPointer(&inst.AltitudeOverrideReference, rhs, identifierMap)
					case "VerticalAngle":
						__gong__assignPointer(&inst.VerticalAngle, rhs, identifierMap)
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "RequiredNavigationPerformance":
						__gong__assignPointer(&inst.RequiredNavigationPerformance, rhs, identifierMap)
					case "MinimumObstacleClearanceAltitude":
						__gong__assignPointer(&inst.MinimumObstacleClearanceAltitude, rhs, identifierMap)
					case "Departure":
						__gong__assignPointer(&inst.Departure, rhs, identifierMap)
					case "Condition":
						__gong__assignSliceOfPointers(&inst.Condition, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.DepartureLegType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.EquipmentUnavailableAdjustmentColumnPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EquipmentUnavailableAdjustmentColumn":
						__gong__assignPointer(&inst.EquipmentUnavailableAdjustmentColumn, rhs, identifierMap)
					}
				case *aixmfeatures.EquipmentUnavailableAdjustmentColumnType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "GuidanceEquipment":
						__gong__assignPointer(&inst.GuidanceEquipment, rhs, identifierMap)
					case "LandingSystemLights":
						__gong__assignPointer(&inst.LandingSystemLights, rhs, identifierMap)
					case "EquipmentRVR":
						__gong__assignPointer(&inst.EquipmentRVR, rhs, identifierMap)
					case "VisibilityAdjustment":
						__gong__assignPointer(&inst.VisibilityAdjustment, rhs, identifierMap)
					case "ApproachLightingInoperative":
						__gong__assignPointer(&inst.ApproachLightingInoperative, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.EquipmentUnavailableAdjustmentPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EquipmentUnavailableAdjustment":
						__gong__assignPointer(&inst.EquipmentUnavailableAdjustment, rhs, identifierMap)
					}
				case *aixmfeatures.EquipmentUnavailableAdjustmentType:
					switch fieldName {
					case "Name":
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
					case "ApproachLightingInoperative":
						__gong__assignPointer(&inst.ApproachLightingInoperative, rhs, identifierMap)
					case "AdjustmentINOPCol":
						__gong__assignSliceOfPointers(&inst.AdjustmentINOPCol, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.FASDataBlockPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FASDataBlock":
						__gong__assignPointer(&inst.FASDataBlock, rhs, identifierMap)
					}
				case *aixmfeatures.FASDataBlockType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "HorizontalAlarmLimit":
						__gong__assignPointer(&inst.HorizontalAlarmLimit, rhs, identifierMap)
					case "VerticalAlarmLimit":
						__gong__assignPointer(&inst.VerticalAlarmLimit, rhs, identifierMap)
					case "ThresholdCourseWidth":
						__gong__assignPointer(&inst.ThresholdCourseWidth, rhs, identifierMap)
					case "LengthOffset":
						__gong__assignPointer(&inst.LengthOffset, rhs, identifierMap)
					case "CRCRemainder":
						__gong__assignPointer(&inst.CRCRemainder, rhs, identifierMap)
					case "OperationType":
						__gong__assignPointer(&inst.OperationType, rhs, identifierMap)
					case "ServiceProviderSBAS":
						__gong__assignPointer(&inst.ServiceProviderSBAS, rhs, identifierMap)
					case "ApproachPerformanceDesignator":
						__gong__assignPointer(&inst.ApproachPerformanceDesignator, rhs, identifierMap)
					case "RouteIndicator":
						__gong__assignPointer(&inst.RouteIndicator, rhs, identifierMap)
					case "ReferencePathDataSelector":
						__gong__assignPointer(&inst.ReferencePathDataSelector, rhs, identifierMap)
					case "ReferencePathIdentifier":
						__gong__assignPointer(&inst.ReferencePathIdentifier, rhs, identifierMap)
					case "CodeICAO":
						__gong__assignPointer(&inst.CodeICAO, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.FinalLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "FinalLegTimeSlice":
						__gong__assignPointer(&inst.FinalLegTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.FinalLegTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "EndConditionDesignator":
						__gong__assignPointer(&inst.EndConditionDesignator, rhs, identifierMap)
					case "LegPath":
						__gong__assignPointer(&inst.LegPath, rhs, identifierMap)
					case "LegTypeARINC":
						__gong__assignPointer(&inst.LegTypeARINC, rhs, identifierMap)
					case "Course":
						__gong__assignPointer(&inst.Course, rhs, identifierMap)
					case "CourseType":
						__gong__assignPointer(&inst.CourseType, rhs, identifierMap)
					case "CourseDirection":
						__gong__assignPointer(&inst.CourseDirection, rhs, identifierMap)
					case "TurnDirection":
						__gong__assignPointer(&inst.TurnDirection, rhs, identifierMap)
					case "SpeedLimit":
						__gong__assignPointer(&inst.SpeedLimit, rhs, identifierMap)
					case "SpeedReference":
						__gong__assignPointer(&inst.SpeedReference, rhs, identifierMap)
					case "SpeedInterpretation":
						__gong__assignPointer(&inst.SpeedInterpretation, rhs, identifierMap)
					case "BankAngle":
						__gong__assignPointer(&inst.BankAngle, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Duration":
						__gong__assignPointer(&inst.Duration, rhs, identifierMap)
					case "ProcedureTurnRequired":
						__gong__assignPointer(&inst.ProcedureTurnRequired, rhs, identifierMap)
					case "UpperLimitAltitude":
						__gong__assignPointer(&inst.UpperLimitAltitude, rhs, identifierMap)
					case "UpperLimitReference":
						__gong__assignPointer(&inst.UpperLimitReference, rhs, identifierMap)
					case "LowerLimitAltitude":
						__gong__assignPointer(&inst.LowerLimitAltitude, rhs, identifierMap)
					case "LowerLimitReference":
						__gong__assignPointer(&inst.LowerLimitReference, rhs, identifierMap)
					case "AltitudeInterpretation":
						__gong__assignPointer(&inst.AltitudeInterpretation, rhs, identifierMap)
					case "AltitudeOverrideATC":
						__gong__assignPointer(&inst.AltitudeOverrideATC, rhs, identifierMap)
					case "AltitudeOverrideReference":
						__gong__assignPointer(&inst.AltitudeOverrideReference, rhs, identifierMap)
					case "VerticalAngle":
						__gong__assignPointer(&inst.VerticalAngle, rhs, identifierMap)
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Approach":
						__gong__assignPointer(&inst.Approach, rhs, identifierMap)
					case "GuidanceSystem":
						__gong__assignPointer(&inst.GuidanceSystem, rhs, identifierMap)
					case "LandingSystemCategory":
						__gong__assignPointer(&inst.LandingSystemCategory, rhs, identifierMap)
					case "MinimumBaroVnavTemperature":
						__gong__assignPointer(&inst.MinimumBaroVnavTemperature, rhs, identifierMap)
					case "RnpDMEAuthorized":
						__gong__assignPointer(&inst.RnpDMEAuthorized, rhs, identifierMap)
					case "CourseOffsetAngle":
						__gong__assignPointer(&inst.CourseOffsetAngle, rhs, identifierMap)
					case "CourseOffsetSide":
						__gong__assignPointer(&inst.CourseOffsetSide, rhs, identifierMap)
					case "CourseCentrelineDistance":
						__gong__assignPointer(&inst.CourseCentrelineDistance, rhs, identifierMap)
					case "CourseOffsetDistance":
						__gong__assignPointer(&inst.CourseOffsetDistance, rhs, identifierMap)
					case "CourseCentrelineIntersect":
						__gong__assignPointer(&inst.CourseCentrelineIntersect, rhs, identifierMap)
					case "Condition":
						__gong__assignSliceOfPointers(&inst.Condition, rhs, identifierMap)
					case "FinalPathAlignmentPoint_fixDesignatedPoint":
						__gong__assignPointer(&inst.FinalPathAlignmentPoint_fixDesignatedPoint, rhs, identifierMap)
					case "FinalPathAlignmentPoint_navaidSystem":
						__gong__assignPointer(&inst.FinalPathAlignmentPoint_navaidSystem, rhs, identifierMap)
					case "FinalPathAlignmentPoint_position":
						__gong__assignPointer(&inst.FinalPathAlignmentPoint_position, rhs, identifierMap)
					case "FinalPathAlignmentPoint_runwayPoint":
						__gong__assignPointer(&inst.FinalPathAlignmentPoint_runwayPoint, rhs, identifierMap)
					case "FinalPathAlignmentPoint_aimingPoint":
						__gong__assignPointer(&inst.FinalPathAlignmentPoint_aimingPoint, rhs, identifierMap)
					case "FinalPathAlignmentPoint_airportReferencePoint":
						__gong__assignPointer(&inst.FinalPathAlignmentPoint_airportReferencePoint, rhs, identifierMap)
					case "VisualDescentPoint":
						__gong__assignPointer(&inst.VisualDescentPoint, rhs, identifierMap)
					case "FASData":
						__gong__assignPointer(&inst.FASData, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.FinalLegType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.FinalProfilePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "FinalProfile":
						__gong__assignPointer(&inst.FinalProfile, rhs, identifierMap)
					}
				case *aixmfeatures.FinalProfileType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Altitude":
						__gong__assignSliceOfPointers(&inst.Altitude, rhs, identifierMap)
					case "Distance":
						__gong__assignSliceOfPointers(&inst.Distance, rhs, identifierMap)
					case "Timing":
						__gong__assignSliceOfPointers(&inst.Timing, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.HoldingAssessmentTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "HoldingAssessmentTimeSlice":
						__gong__assignPointer(&inst.HoldingAssessmentTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.HoldingAssessmentTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "UpperLimit":
						__gong__assignPointer(&inst.UpperLimit, rhs, identifierMap)
					case "UpperLimitReference":
						__gong__assignPointer(&inst.UpperLimitReference, rhs, identifierMap)
					case "LowerLimit":
						__gong__assignPointer(&inst.LowerLimit, rhs, identifierMap)
					case "LowerLimitReference":
						__gong__assignPointer(&inst.LowerLimitReference, rhs, identifierMap)
					case "SpeedLimit":
						__gong__assignPointer(&inst.SpeedLimit, rhs, identifierMap)
					case "PatternTemplate":
						__gong__assignPointer(&inst.PatternTemplate, rhs, identifierMap)
					case "TurbulentAir":
						__gong__assignPointer(&inst.TurbulentAir, rhs, identifierMap)
					case "LegLengthToward":
						__gong__assignPointer(&inst.LegLengthToward, rhs, identifierMap)
					case "LegLengthAway":
						__gong__assignPointer(&inst.LegLengthAway, rhs, identifierMap)
					case "HoldingPoint":
						__gong__assignPointer(&inst.HoldingPoint, rhs, identifierMap)
					case "UnplannedHolding":
						__gong__assignPointer(&inst.UnplannedHolding, rhs, identifierMap)
					case "AssessedHoldingPattern":
						__gong__assignPointer(&inst.AssessedHoldingPattern, rhs, identifierMap)
					case "ObstacleAssessment":
						__gong__assignSliceOfPointers(&inst.ObstacleAssessment, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.HoldingAssessmentType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.HoldingPatternDistancePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "HoldingPatternDistance":
						__gong__assignPointer(&inst.HoldingPatternDistance, rhs, identifierMap)
					}
				case *aixmfeatures.HoldingPatternDistanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.HoldingPatternDurationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "HoldingPatternDuration":
						__gong__assignPointer(&inst.HoldingPatternDuration, rhs, identifierMap)
					}
				case *aixmfeatures.HoldingPatternDurationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Duration":
						__gong__assignPointer(&inst.Duration, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.HoldingPatternPropertyType:
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
				case *aixmfeatures.HoldingPatternTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "HoldingPatternTimeSlice":
						__gong__assignPointer(&inst.HoldingPatternTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.HoldingPatternTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "OutboundCourse":
						__gong__assignPointer(&inst.OutboundCourse, rhs, identifierMap)
					case "OutboundCourseType":
						__gong__assignPointer(&inst.OutboundCourseType, rhs, identifierMap)
					case "InboundCourse":
						__gong__assignPointer(&inst.InboundCourse, rhs, identifierMap)
					case "TurnDirection":
						__gong__assignPointer(&inst.TurnDirection, rhs, identifierMap)
					case "UpperLimit":
						__gong__assignPointer(&inst.UpperLimit, rhs, identifierMap)
					case "UpperLimitReference":
						__gong__assignPointer(&inst.UpperLimitReference, rhs, identifierMap)
					case "LowerLimit":
						__gong__assignPointer(&inst.LowerLimit, rhs, identifierMap)
					case "LowerLimitReference":
						__gong__assignPointer(&inst.LowerLimitReference, rhs, identifierMap)
					case "SpeedLimit":
						__gong__assignPointer(&inst.SpeedLimit, rhs, identifierMap)
					case "Instruction":
						__gong__assignPointer(&inst.Instruction, rhs, identifierMap)
					case "NonStandardHolding":
						__gong__assignPointer(&inst.NonStandardHolding, rhs, identifierMap)
					case "OutboundLegSpan_endPoint":
						__gong__assignPointer(&inst.OutboundLegSpan_endPoint, rhs, identifierMap)
					case "OutboundLegSpan_endDistance":
						__gong__assignPointer(&inst.OutboundLegSpan_endDistance, rhs, identifierMap)
					case "OutboundLegSpan_endTime":
						__gong__assignPointer(&inst.OutboundLegSpan_endTime, rhs, identifierMap)
					case "HoldingPoint":
						__gong__assignPointer(&inst.HoldingPoint, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.HoldingPatternType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.HoldingUsePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "HoldingUse":
						__gong__assignPointer(&inst.HoldingUse, rhs, identifierMap)
					}
				case *aixmfeatures.HoldingUseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "HoldingUse":
						__gong__assignPointer(&inst.HoldingUse, rhs, identifierMap)
					case "Instruction":
						__gong__assignPointer(&inst.Instruction, rhs, identifierMap)
					case "InstructedAltitude":
						__gong__assignPointer(&inst.InstructedAltitude, rhs, identifierMap)
					case "InstructionAltitudeReference":
						__gong__assignPointer(&inst.InstructionAltitudeReference, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheHoldingPattern":
						__gong__assignPointer(&inst.TheHoldingPattern, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.InitialLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "InitialLegTimeSlice":
						__gong__assignPointer(&inst.InitialLegTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.InitialLegTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "EndConditionDesignator":
						__gong__assignPointer(&inst.EndConditionDesignator, rhs, identifierMap)
					case "LegPath":
						__gong__assignPointer(&inst.LegPath, rhs, identifierMap)
					case "LegTypeARINC":
						__gong__assignPointer(&inst.LegTypeARINC, rhs, identifierMap)
					case "Course":
						__gong__assignPointer(&inst.Course, rhs, identifierMap)
					case "CourseType":
						__gong__assignPointer(&inst.CourseType, rhs, identifierMap)
					case "CourseDirection":
						__gong__assignPointer(&inst.CourseDirection, rhs, identifierMap)
					case "TurnDirection":
						__gong__assignPointer(&inst.TurnDirection, rhs, identifierMap)
					case "SpeedLimit":
						__gong__assignPointer(&inst.SpeedLimit, rhs, identifierMap)
					case "SpeedReference":
						__gong__assignPointer(&inst.SpeedReference, rhs, identifierMap)
					case "SpeedInterpretation":
						__gong__assignPointer(&inst.SpeedInterpretation, rhs, identifierMap)
					case "BankAngle":
						__gong__assignPointer(&inst.BankAngle, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Duration":
						__gong__assignPointer(&inst.Duration, rhs, identifierMap)
					case "ProcedureTurnRequired":
						__gong__assignPointer(&inst.ProcedureTurnRequired, rhs, identifierMap)
					case "UpperLimitAltitude":
						__gong__assignPointer(&inst.UpperLimitAltitude, rhs, identifierMap)
					case "UpperLimitReference":
						__gong__assignPointer(&inst.UpperLimitReference, rhs, identifierMap)
					case "LowerLimitAltitude":
						__gong__assignPointer(&inst.LowerLimitAltitude, rhs, identifierMap)
					case "LowerLimitReference":
						__gong__assignPointer(&inst.LowerLimitReference, rhs, identifierMap)
					case "AltitudeInterpretation":
						__gong__assignPointer(&inst.AltitudeInterpretation, rhs, identifierMap)
					case "AltitudeOverrideATC":
						__gong__assignPointer(&inst.AltitudeOverrideATC, rhs, identifierMap)
					case "AltitudeOverrideReference":
						__gong__assignPointer(&inst.AltitudeOverrideReference, rhs, identifierMap)
					case "VerticalAngle":
						__gong__assignPointer(&inst.VerticalAngle, rhs, identifierMap)
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Approach":
						__gong__assignPointer(&inst.Approach, rhs, identifierMap)
					case "RequiredNavigationPerformance":
						__gong__assignPointer(&inst.RequiredNavigationPerformance, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.InitialLegType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.InstrumentApproachProcedurePropertyType:
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
				case *aixmfeatures.InstrumentApproachProcedureTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "InstrumentApproachProcedureTimeSlice":
						__gong__assignPointer(&inst.InstrumentApproachProcedureTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.InstrumentApproachProcedureTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "CommunicationFailureInstruction":
						__gong__assignPointer(&inst.CommunicationFailureInstruction, rhs, identifierMap)
					case "Instruction":
						__gong__assignPointer(&inst.Instruction, rhs, identifierMap)
					case "DesignCriteria":
						__gong__assignPointer(&inst.DesignCriteria, rhs, identifierMap)
					case "CodingStandard":
						__gong__assignPointer(&inst.CodingStandard, rhs, identifierMap)
					case "FlightChecked":
						__gong__assignPointer(&inst.FlightChecked, rhs, identifierMap)
					case "RNAV":
						__gong__assignPointer(&inst.RNAV, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignSliceOfPointers(&inst.AirportHeliport, rhs, identifierMap)
					case "AircraftCharacteristic":
						__gong__assignSliceOfPointers(&inst.AircraftCharacteristic, rhs, identifierMap)
					case "FlightTransition":
						__gong__assignSliceOfPointers(&inst.FlightTransition, rhs, identifierMap)
					case "GuidanceFacility_specialNavigationSystem":
						__gong__assignPointer(&inst.GuidanceFacility_specialNavigationSystem, rhs, identifierMap)
					case "GuidanceFacility_radar":
						__gong__assignPointer(&inst.GuidanceFacility_radar, rhs, identifierMap)
					case "GuidanceFacility_navaid":
						__gong__assignPointer(&inst.GuidanceFacility_navaid, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SafeAltitude":
						__gong__assignPointer(&inst.SafeAltitude, rhs, identifierMap)
					case "ApproachPrefix":
						__gong__assignPointer(&inst.ApproachPrefix, rhs, identifierMap)
					case "ApproachType":
						__gong__assignPointer(&inst.ApproachType, rhs, identifierMap)
					case "MultipleIdentification":
						__gong__assignPointer(&inst.MultipleIdentification, rhs, identifierMap)
					case "CopterTrack":
						__gong__assignPointer(&inst.CopterTrack, rhs, identifierMap)
					case "CirclingIdentification":
						__gong__assignPointer(&inst.CirclingIdentification, rhs, identifierMap)
					case "CourseReversalInstruction":
						__gong__assignPointer(&inst.CourseReversalInstruction, rhs, identifierMap)
					case "AdditionalEquipment":
						__gong__assignPointer(&inst.AdditionalEquipment, rhs, identifierMap)
					case "ChannelGNSS":
						__gong__assignPointer(&inst.ChannelGNSS, rhs, identifierMap)
					case "WAASReliable":
						__gong__assignPointer(&inst.WAASReliable, rhs, identifierMap)
					case "Landing":
						__gong__assignPointer(&inst.Landing, rhs, identifierMap)
					case "MissedInstruction":
						__gong__assignSliceOfPointers(&inst.MissedInstruction, rhs, identifierMap)
					case "FinalProfile":
						__gong__assignPointer(&inst.FinalProfile, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.InstrumentApproachProcedureType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.IntermediateLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "IntermediateLegTimeSlice":
						__gong__assignPointer(&inst.IntermediateLegTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.IntermediateLegTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "EndConditionDesignator":
						__gong__assignPointer(&inst.EndConditionDesignator, rhs, identifierMap)
					case "LegPath":
						__gong__assignPointer(&inst.LegPath, rhs, identifierMap)
					case "LegTypeARINC":
						__gong__assignPointer(&inst.LegTypeARINC, rhs, identifierMap)
					case "Course":
						__gong__assignPointer(&inst.Course, rhs, identifierMap)
					case "CourseType":
						__gong__assignPointer(&inst.CourseType, rhs, identifierMap)
					case "CourseDirection":
						__gong__assignPointer(&inst.CourseDirection, rhs, identifierMap)
					case "TurnDirection":
						__gong__assignPointer(&inst.TurnDirection, rhs, identifierMap)
					case "SpeedLimit":
						__gong__assignPointer(&inst.SpeedLimit, rhs, identifierMap)
					case "SpeedReference":
						__gong__assignPointer(&inst.SpeedReference, rhs, identifierMap)
					case "SpeedInterpretation":
						__gong__assignPointer(&inst.SpeedInterpretation, rhs, identifierMap)
					case "BankAngle":
						__gong__assignPointer(&inst.BankAngle, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Duration":
						__gong__assignPointer(&inst.Duration, rhs, identifierMap)
					case "ProcedureTurnRequired":
						__gong__assignPointer(&inst.ProcedureTurnRequired, rhs, identifierMap)
					case "UpperLimitAltitude":
						__gong__assignPointer(&inst.UpperLimitAltitude, rhs, identifierMap)
					case "UpperLimitReference":
						__gong__assignPointer(&inst.UpperLimitReference, rhs, identifierMap)
					case "LowerLimitAltitude":
						__gong__assignPointer(&inst.LowerLimitAltitude, rhs, identifierMap)
					case "LowerLimitReference":
						__gong__assignPointer(&inst.LowerLimitReference, rhs, identifierMap)
					case "AltitudeInterpretation":
						__gong__assignPointer(&inst.AltitudeInterpretation, rhs, identifierMap)
					case "AltitudeOverrideATC":
						__gong__assignPointer(&inst.AltitudeOverrideATC, rhs, identifierMap)
					case "AltitudeOverrideReference":
						__gong__assignPointer(&inst.AltitudeOverrideReference, rhs, identifierMap)
					case "VerticalAngle":
						__gong__assignPointer(&inst.VerticalAngle, rhs, identifierMap)
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Approach":
						__gong__assignPointer(&inst.Approach, rhs, identifierMap)
					case "RequiredNavigationPerformance":
						__gong__assignPointer(&inst.RequiredNavigationPerformance, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.IntermediateLegType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.LandingTakeoffAreaCollectionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "LandingTakeoffAreaCollection":
						__gong__assignPointer(&inst.LandingTakeoffAreaCollection, rhs, identifierMap)
					}
				case *aixmfeatures.LandingTakeoffAreaCollectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Runway":
						__gong__assignSliceOfPointers(&inst.Runway, rhs, identifierMap)
					case "TLOF":
						__gong__assignSliceOfPointers(&inst.TLOF, rhs, identifierMap)
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
				case *aixmfeatures.MinimaPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Minima":
						__gong__assignPointer(&inst.Minima, rhs, identifierMap)
					}
				case *aixmfeatures.MinimaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Altitude":
						__gong__assignPointer(&inst.Altitude, rhs, identifierMap)
					case "AltitudeCode":
						__gong__assignPointer(&inst.AltitudeCode, rhs, identifierMap)
					case "AltitudeReference":
						__gong__assignPointer(&inst.AltitudeReference, rhs, identifierMap)
					case "Height":
						__gong__assignPointer(&inst.Height, rhs, identifierMap)
					case "MilitaryHeight":
						__gong__assignPointer(&inst.MilitaryHeight, rhs, identifierMap)
					case "RadioHeight":
						__gong__assignPointer(&inst.RadioHeight, rhs, identifierMap)
					case "HeightCode":
						__gong__assignPointer(&inst.HeightCode, rhs, identifierMap)
					case "HeightReference":
						__gong__assignPointer(&inst.HeightReference, rhs, identifierMap)
					case "Visibility":
						__gong__assignPointer(&inst.Visibility, rhs, identifierMap)
					case "MilitaryVisibility":
						__gong__assignPointer(&inst.MilitaryVisibility, rhs, identifierMap)
					case "MandatoryRVR":
						__gong__assignPointer(&inst.MandatoryRVR, rhs, identifierMap)
					case "RemoteAltimeterMinima":
						__gong__assignPointer(&inst.RemoteAltimeterMinima, rhs, identifierMap)
					case "AdjustmentINOP":
						__gong__assignSliceOfPointers(&inst.AdjustmentINOP, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.MissedApproachGroupPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "MissedApproachGroup":
						__gong__assignPointer(&inst.MissedApproachGroup, rhs, identifierMap)
					}
				case *aixmfeatures.MissedApproachGroupType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Instruction":
						__gong__assignPointer(&inst.Instruction, rhs, identifierMap)
					case "AlternateClimbInstruction":
						__gong__assignPointer(&inst.AlternateClimbInstruction, rhs, identifierMap)
					case "AlternateClimbAltitude":
						__gong__assignPointer(&inst.AlternateClimbAltitude, rhs, identifierMap)
					case "Altimeter":
						__gong__assignSliceOfPointers(&inst.Altimeter, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.MissedApproachLegTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "MissedApproachLegTimeSlice":
						__gong__assignPointer(&inst.MissedApproachLegTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.MissedApproachLegTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "EndConditionDesignator":
						__gong__assignPointer(&inst.EndConditionDesignator, rhs, identifierMap)
					case "LegPath":
						__gong__assignPointer(&inst.LegPath, rhs, identifierMap)
					case "LegTypeARINC":
						__gong__assignPointer(&inst.LegTypeARINC, rhs, identifierMap)
					case "Course":
						__gong__assignPointer(&inst.Course, rhs, identifierMap)
					case "CourseType":
						__gong__assignPointer(&inst.CourseType, rhs, identifierMap)
					case "CourseDirection":
						__gong__assignPointer(&inst.CourseDirection, rhs, identifierMap)
					case "TurnDirection":
						__gong__assignPointer(&inst.TurnDirection, rhs, identifierMap)
					case "SpeedLimit":
						__gong__assignPointer(&inst.SpeedLimit, rhs, identifierMap)
					case "SpeedReference":
						__gong__assignPointer(&inst.SpeedReference, rhs, identifierMap)
					case "SpeedInterpretation":
						__gong__assignPointer(&inst.SpeedInterpretation, rhs, identifierMap)
					case "BankAngle":
						__gong__assignPointer(&inst.BankAngle, rhs, identifierMap)
					case "Length":
						__gong__assignPointer(&inst.Length, rhs, identifierMap)
					case "Duration":
						__gong__assignPointer(&inst.Duration, rhs, identifierMap)
					case "ProcedureTurnRequired":
						__gong__assignPointer(&inst.ProcedureTurnRequired, rhs, identifierMap)
					case "UpperLimitAltitude":
						__gong__assignPointer(&inst.UpperLimitAltitude, rhs, identifierMap)
					case "UpperLimitReference":
						__gong__assignPointer(&inst.UpperLimitReference, rhs, identifierMap)
					case "LowerLimitAltitude":
						__gong__assignPointer(&inst.LowerLimitAltitude, rhs, identifierMap)
					case "LowerLimitReference":
						__gong__assignPointer(&inst.LowerLimitReference, rhs, identifierMap)
					case "AltitudeInterpretation":
						__gong__assignPointer(&inst.AltitudeInterpretation, rhs, identifierMap)
					case "AltitudeOverrideATC":
						__gong__assignPointer(&inst.AltitudeOverrideATC, rhs, identifierMap)
					case "AltitudeOverrideReference":
						__gong__assignPointer(&inst.AltitudeOverrideReference, rhs, identifierMap)
					case "VerticalAngle":
						__gong__assignPointer(&inst.VerticalAngle, rhs, identifierMap)
					case "StartPoint":
						__gong__assignPointer(&inst.StartPoint, rhs, identifierMap)
					case "EndPoint":
						__gong__assignPointer(&inst.EndPoint, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "ArcCentre":
						__gong__assignPointer(&inst.ArcCentre, rhs, identifierMap)
					case "Angle":
						__gong__assignPointer(&inst.Angle, rhs, identifierMap)
					case "Distance":
						__gong__assignPointer(&inst.Distance, rhs, identifierMap)
					case "AircraftCategory":
						__gong__assignSliceOfPointers(&inst.AircraftCategory, rhs, identifierMap)
					case "Holding":
						__gong__assignPointer(&inst.Holding, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignSliceOfPointers(&inst.DesignSurface, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Approach":
						__gong__assignPointer(&inst.Approach, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "ThresholdAfterMAPT":
						__gong__assignPointer(&inst.ThresholdAfterMAPT, rhs, identifierMap)
					case "HeightMAPT":
						__gong__assignPointer(&inst.HeightMAPT, rhs, identifierMap)
					case "RequiredNavigationPerformance":
						__gong__assignPointer(&inst.RequiredNavigationPerformance, rhs, identifierMap)
					case "Condition":
						__gong__assignSliceOfPointers(&inst.Condition, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.MissedApproachLegType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.NavigationAreaRestrictionTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NavigationAreaRestrictionTimeSlice":
						__gong__assignPointer(&inst.NavigationAreaRestrictionTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.NavigationAreaRestrictionTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "Procedure":
						__gong__assignSliceOfPointers(&inst.Procedure, rhs, identifierMap)
					case "DesignSurface":
						__gong__assignPointer(&inst.DesignSurface, rhs, identifierMap)
					case "SectorDefinition":
						__gong__assignPointer(&inst.SectorDefinition, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.NavigationAreaRestrictionType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.NavigationAreaSectorPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "NavigationAreaSector":
						__gong__assignPointer(&inst.NavigationAreaSector, rhs, identifierMap)
					}
				case *aixmfeatures.NavigationAreaSectorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SectorDefinition":
						__gong__assignPointer(&inst.SectorDefinition, rhs, identifierMap)
					case "SignificantObstacle":
						__gong__assignSliceOfPointers(&inst.SignificantObstacle, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "SectorCriteria":
						__gong__assignSliceOfPointers(&inst.SectorCriteria, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.NavigationAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "NavigationAreaTimeSlice":
						__gong__assignPointer(&inst.NavigationAreaTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.NavigationAreaTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "NavigationAreaType":
						__gong__assignPointer(&inst.NavigationAreaType, rhs, identifierMap)
					case "MinimumCeiling":
						__gong__assignPointer(&inst.MinimumCeiling, rhs, identifierMap)
					case "MinimumVisibility":
						__gong__assignPointer(&inst.MinimumVisibility, rhs, identifierMap)
					case "Departure":
						__gong__assignPointer(&inst.Departure, rhs, identifierMap)
					case "Sector":
						__gong__assignSliceOfPointers(&inst.Sector, rhs, identifierMap)
					case "CentrePoint_fixDesignatedPoint":
						__gong__assignPointer(&inst.CentrePoint_fixDesignatedPoint, rhs, identifierMap)
					case "CentrePoint_navaidSystem":
						__gong__assignPointer(&inst.CentrePoint_navaidSystem, rhs, identifierMap)
					case "CentrePoint_position":
						__gong__assignPointer(&inst.CentrePoint_position, rhs, identifierMap)
					case "CentrePoint_runwayPoint":
						__gong__assignPointer(&inst.CentrePoint_runwayPoint, rhs, identifierMap)
					case "CentrePoint_aimingPoint":
						__gong__assignPointer(&inst.CentrePoint_aimingPoint, rhs, identifierMap)
					case "CentrePoint_airportReferencePoint":
						__gong__assignPointer(&inst.CentrePoint_airportReferencePoint, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.NavigationAreaType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.ProcedureAvailabilityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ProcedureAvailability":
						__gong__assignPointer(&inst.ProcedureAvailability, rhs, identifierMap)
					}
				case *aixmfeatures.ProcedureAvailabilityType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.ProcedureDMETimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "ProcedureDMETimeSlice":
						__gong__assignPointer(&inst.ProcedureDMETimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.ProcedureDMETimeSliceType:
					switch fieldName {
					case "Name":
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
					case "DME":
						__gong__assignPointer(&inst.DME, rhs, identifierMap)
					case "SegmentLeg":
						__gong__assignPointer(&inst.SegmentLeg, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ProcedureDMEType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.ProcedurePropertyType:
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
				case *aixmfeatures.ProcedureTransitionLegPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ProcedureTransitionLeg":
						__gong__assignPointer(&inst.ProcedureTransitionLeg, rhs, identifierMap)
					}
				case *aixmfeatures.ProcedureTransitionLegType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "SeqNumberARINC":
						__gong__assignPointer(&inst.SeqNumberARINC, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "TheSegmentLeg":
						__gong__assignPointer(&inst.TheSegmentLeg, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.ProcedureTransitionPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ProcedureTransition":
						__gong__assignPointer(&inst.ProcedureTransition, rhs, identifierMap)
					}
				case *aixmfeatures.ProcedureTransitionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "TransitionId":
						__gong__assignPointer(&inst.TransitionId, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Instruction":
						__gong__assignPointer(&inst.Instruction, rhs, identifierMap)
					case "VectorHeading":
						__gong__assignPointer(&inst.VectorHeading, rhs, identifierMap)
					case "DepartureRunwayTransition":
						__gong__assignPointer(&inst.DepartureRunwayTransition, rhs, identifierMap)
					case "Trajectory":
						__gong__assignPointer(&inst.Trajectory, rhs, identifierMap)
					case "TransitionLeg":
						__gong__assignSliceOfPointers(&inst.TransitionLeg, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RadarSystemPropertyType:
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
				case *aixmfeatures.SafeAltitudeAreaPropertyType:
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
				case *aixmfeatures.SafeAltitudeAreaSectorPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "SafeAltitudeAreaSector":
						__gong__assignPointer(&inst.SafeAltitudeAreaSector, rhs, identifierMap)
					}
				case *aixmfeatures.SafeAltitudeAreaSectorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "BufferWidth":
						__gong__assignPointer(&inst.BufferWidth, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "SignificantObstacle":
						__gong__assignSliceOfPointers(&inst.SignificantObstacle, rhs, identifierMap)
					case "SectorDefinition":
						__gong__assignPointer(&inst.SectorDefinition, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.SafeAltitudeAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SafeAltitudeAreaTimeSlice":
						__gong__assignPointer(&inst.SafeAltitudeAreaTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.SafeAltitudeAreaTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "SafeAreaType":
						__gong__assignPointer(&inst.SafeAreaType, rhs, identifierMap)
					case "CentrePoint_fixDesignatedPoint":
						__gong__assignPointer(&inst.CentrePoint_fixDesignatedPoint, rhs, identifierMap)
					case "CentrePoint_navaidSystem":
						__gong__assignPointer(&inst.CentrePoint_navaidSystem, rhs, identifierMap)
					case "CentrePoint_position":
						__gong__assignPointer(&inst.CentrePoint_position, rhs, identifierMap)
					case "CentrePoint_runwayPoint":
						__gong__assignPointer(&inst.CentrePoint_runwayPoint, rhs, identifierMap)
					case "CentrePoint_aimingPoint":
						__gong__assignPointer(&inst.CentrePoint_aimingPoint, rhs, identifierMap)
					case "CentrePoint_airportReferencePoint":
						__gong__assignPointer(&inst.CentrePoint_airportReferencePoint, rhs, identifierMap)
					case "Sector":
						__gong__assignSliceOfPointers(&inst.Sector, rhs, identifierMap)
					case "Location":
						__gong__assignSliceOfPointers(&inst.Location, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.SafeAltitudeAreaType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.SectorDesignPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "SectorDesign":
						__gong__assignPointer(&inst.SectorDesign, rhs, identifierMap)
					}
				case *aixmfeatures.SectorDesignType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "TurnDirection":
						__gong__assignPointer(&inst.TurnDirection, rhs, identifierMap)
					case "DesignGradient":
						__gong__assignPointer(&inst.DesignGradient, rhs, identifierMap)
					case "TerminationAltitude":
						__gong__assignPointer(&inst.TerminationAltitude, rhs, identifierMap)
					case "TurnPermitted":
						__gong__assignPointer(&inst.TurnPermitted, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.SegmentLegPropertyType:
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
				case *aixmfeatures.SegmentPointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "AerialRefuellingPoint":
						__gong__assignPointer(&inst.AerialRefuellingPoint, rhs, identifierMap)
					case "EnRouteSegmentPoint":
						__gong__assignPointer(&inst.EnRouteSegmentPoint, rhs, identifierMap)
					case "TerminalSegmentPoint":
						__gong__assignPointer(&inst.TerminalSegmentPoint, rhs, identifierMap)
					}
				case *aixmfeatures.SpecialNavigationSystemPropertyType:
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
				case *aixmfeatures.StandardInstrumentArrivalTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "StandardInstrumentArrivalTimeSlice":
						__gong__assignPointer(&inst.StandardInstrumentArrivalTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.StandardInstrumentArrivalTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "CommunicationFailureInstruction":
						__gong__assignPointer(&inst.CommunicationFailureInstruction, rhs, identifierMap)
					case "Instruction":
						__gong__assignPointer(&inst.Instruction, rhs, identifierMap)
					case "DesignCriteria":
						__gong__assignPointer(&inst.DesignCriteria, rhs, identifierMap)
					case "CodingStandard":
						__gong__assignPointer(&inst.CodingStandard, rhs, identifierMap)
					case "FlightChecked":
						__gong__assignPointer(&inst.FlightChecked, rhs, identifierMap)
					case "RNAV":
						__gong__assignPointer(&inst.RNAV, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignSliceOfPointers(&inst.AirportHeliport, rhs, identifierMap)
					case "AircraftCharacteristic":
						__gong__assignSliceOfPointers(&inst.AircraftCharacteristic, rhs, identifierMap)
					case "FlightTransition":
						__gong__assignSliceOfPointers(&inst.FlightTransition, rhs, identifierMap)
					case "GuidanceFacility_specialNavigationSystem":
						__gong__assignPointer(&inst.GuidanceFacility_specialNavigationSystem, rhs, identifierMap)
					case "GuidanceFacility_radar":
						__gong__assignPointer(&inst.GuidanceFacility_radar, rhs, identifierMap)
					case "GuidanceFacility_navaid":
						__gong__assignPointer(&inst.GuidanceFacility_navaid, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SafeAltitude":
						__gong__assignPointer(&inst.SafeAltitude, rhs, identifierMap)
					case "Designator":
						__gong__assignPointer(&inst.Designator, rhs, identifierMap)
					case "Arrival":
						__gong__assignPointer(&inst.Arrival, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.StandardInstrumentArrivalType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.StandardInstrumentDepartureTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "StandardInstrumentDepartureTimeSlice":
						__gong__assignPointer(&inst.StandardInstrumentDepartureTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.StandardInstrumentDepartureTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "CommunicationFailureInstruction":
						__gong__assignPointer(&inst.CommunicationFailureInstruction, rhs, identifierMap)
					case "Instruction":
						__gong__assignPointer(&inst.Instruction, rhs, identifierMap)
					case "DesignCriteria":
						__gong__assignPointer(&inst.DesignCriteria, rhs, identifierMap)
					case "CodingStandard":
						__gong__assignPointer(&inst.CodingStandard, rhs, identifierMap)
					case "FlightChecked":
						__gong__assignPointer(&inst.FlightChecked, rhs, identifierMap)
					case "RNAV":
						__gong__assignPointer(&inst.RNAV, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignSliceOfPointers(&inst.AirportHeliport, rhs, identifierMap)
					case "AircraftCharacteristic":
						__gong__assignSliceOfPointers(&inst.AircraftCharacteristic, rhs, identifierMap)
					case "FlightTransition":
						__gong__assignSliceOfPointers(&inst.FlightTransition, rhs, identifierMap)
					case "GuidanceFacility_specialNavigationSystem":
						__gong__assignPointer(&inst.GuidanceFacility_specialNavigationSystem, rhs, identifierMap)
					case "GuidanceFacility_radar":
						__gong__assignPointer(&inst.GuidanceFacility_radar, rhs, identifierMap)
					case "GuidanceFacility_navaid":
						__gong__assignPointer(&inst.GuidanceFacility_navaid, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "SafeAltitude":
						__gong__assignPointer(&inst.SafeAltitude, rhs, identifierMap)
					case "Designator":
						__gong__assignPointer(&inst.Designator, rhs, identifierMap)
					case "ContingencyRoute":
						__gong__assignPointer(&inst.ContingencyRoute, rhs, identifierMap)
					case "Takeoff":
						__gong__assignPointer(&inst.Takeoff, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.StandardInstrumentDepartureType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.TerminalArrivalAreaSectorPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "TerminalArrivalAreaSector":
						__gong__assignPointer(&inst.TerminalArrivalAreaSector, rhs, identifierMap)
					}
				case *aixmfeatures.TerminalArrivalAreaSectorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "FlyByCode":
						__gong__assignPointer(&inst.FlyByCode, rhs, identifierMap)
					case "ProcedureTurnRequired":
						__gong__assignPointer(&inst.ProcedureTurnRequired, rhs, identifierMap)
					case "AltitudeDescription":
						__gong__assignPointer(&inst.AltitudeDescription, rhs, identifierMap)
					case "SectorDefinition":
						__gong__assignPointer(&inst.SectorDefinition, rhs, identifierMap)
					case "Extent":
						__gong__assignPointer(&inst.Extent, rhs, identifierMap)
					case "SignificantObstacle":
						__gong__assignSliceOfPointers(&inst.SignificantObstacle, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TerminalArrivalAreaTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "TerminalArrivalAreaTimeSlice":
						__gong__assignPointer(&inst.TerminalArrivalAreaTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.TerminalArrivalAreaTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "ArrivalAreaType":
						__gong__assignPointer(&inst.ArrivalAreaType, rhs, identifierMap)
					case "OuterBufferWidth":
						__gong__assignPointer(&inst.OuterBufferWidth, rhs, identifierMap)
					case "LateralBufferWidth":
						__gong__assignPointer(&inst.LateralBufferWidth, rhs, identifierMap)
					case "IF_fixDesignatedPoint":
						__gong__assignPointer(&inst.IF_fixDesignatedPoint, rhs, identifierMap)
					case "IF_navaidSystem":
						__gong__assignPointer(&inst.IF_navaidSystem, rhs, identifierMap)
					case "IF_position":
						__gong__assignPointer(&inst.IF_position, rhs, identifierMap)
					case "IF_runwayPoint":
						__gong__assignPointer(&inst.IF_runwayPoint, rhs, identifierMap)
					case "IF_aimingPoint":
						__gong__assignPointer(&inst.IF_aimingPoint, rhs, identifierMap)
					case "IF_airportReferencePoint":
						__gong__assignPointer(&inst.IF_airportReferencePoint, rhs, identifierMap)
					case "IAF_fixDesignatedPoint":
						__gong__assignPointer(&inst.IAF_fixDesignatedPoint, rhs, identifierMap)
					case "IAF_navaidSystem":
						__gong__assignPointer(&inst.IAF_navaidSystem, rhs, identifierMap)
					case "IAF_position":
						__gong__assignPointer(&inst.IAF_position, rhs, identifierMap)
					case "IAF_runwayPoint":
						__gong__assignPointer(&inst.IAF_runwayPoint, rhs, identifierMap)
					case "IAF_aimingPoint":
						__gong__assignPointer(&inst.IAF_aimingPoint, rhs, identifierMap)
					case "IAF_airportReferencePoint":
						__gong__assignPointer(&inst.IAF_airportReferencePoint, rhs, identifierMap)
					case "Buffer":
						__gong__assignPointer(&inst.Buffer, rhs, identifierMap)
					case "Sector":
						__gong__assignSliceOfPointers(&inst.Sector, rhs, identifierMap)
					case "ApproachRNAV":
						__gong__assignPointer(&inst.ApproachRNAV, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.TerminalArrivalAreaType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.TerminalSegmentPointPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "TerminalSegmentPoint":
						__gong__assignPointer(&inst.TerminalSegmentPoint, rhs, identifierMap)
					}
				case *aixmfeatures.TerminalSegmentPointType:
					switch fieldName {
					case "Name":
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
					case "Role":
						__gong__assignPointer(&inst.Role, rhs, identifierMap)
					case "LeadRadial":
						__gong__assignPointer(&inst.LeadRadial, rhs, identifierMap)
					case "LeadDME":
						__gong__assignPointer(&inst.LeadDME, rhs, identifierMap)
					case "IndicatorFACF":
						__gong__assignPointer(&inst.IndicatorFACF, rhs, identifierMap)
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
				case *aixmfeatures.UnplannedHoldingPropertyType:
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
				case *aixmfeatures.UnplannedHoldingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "UnplannedHoldingTimeSlice":
						__gong__assignPointer(&inst.UnplannedHoldingTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.UnplannedHoldingTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "UnplannedHolding":
						__gong__assignPointer(&inst.UnplannedHolding, rhs, identifierMap)
					case "AuthorizedAltitude":
						__gong__assignPointer(&inst.AuthorizedAltitude, rhs, identifierMap)
					case "AltitudeReference":
						__gong__assignPointer(&inst.AltitudeReference, rhs, identifierMap)
					case "ControlledAirspace":
						__gong__assignPointer(&inst.ControlledAirspace, rhs, identifierMap)
					case "HoldingPoint":
						__gong__assignPointer(&inst.HoldingPoint, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.UnplannedHoldingType:
					switch fieldName {
					case "Name":
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
					case "StandardInstrumentDeparture":
						__gong__assignPointer(&inst.StandardInstrumentDeparture, rhs, identifierMap)
					case "StandardInstrumentArrival":
						__gong__assignPointer(&inst.StandardInstrumentArrival, rhs, identifierMap)
					case "InstrumentApproachProcedure":
						__gong__assignPointer(&inst.InstrumentApproachProcedure, rhs, identifierMap)
					case "NavigationArea":
						__gong__assignPointer(&inst.NavigationArea, rhs, identifierMap)
					case "TerminalArrivalArea":
						__gong__assignPointer(&inst.TerminalArrivalArea, rhs, identifierMap)
					case "CirclingArea":
						__gong__assignPointer(&inst.CirclingArea, rhs, identifierMap)
					case "DepartureLeg":
						__gong__assignPointer(&inst.DepartureLeg, rhs, identifierMap)
					case "ArrivalLeg":
						__gong__assignPointer(&inst.ArrivalLeg, rhs, identifierMap)
					case "ArrivalFeederLeg":
						__gong__assignPointer(&inst.ArrivalFeederLeg, rhs, identifierMap)
					case "InitialLeg":
						__gong__assignPointer(&inst.InitialLeg, rhs, identifierMap)
					case "IntermediateLeg":
						__gong__assignPointer(&inst.IntermediateLeg, rhs, identifierMap)
					case "FinalLeg":
						__gong__assignPointer(&inst.FinalLeg, rhs, identifierMap)
					case "MissedApproachLeg":
						__gong__assignPointer(&inst.MissedApproachLeg, rhs, identifierMap)
					case "ProcedureDME":
						__gong__assignPointer(&inst.ProcedureDME, rhs, identifierMap)
					case "NavigationAreaRestriction":
						__gong__assignPointer(&inst.NavigationAreaRestriction, rhs, identifierMap)
					case "SafeAltitudeArea":
						__gong__assignPointer(&inst.SafeAltitudeArea, rhs, identifierMap)
					case "HoldingPattern":
						__gong__assignPointer(&inst.HoldingPattern, rhs, identifierMap)
					case "UnplannedHolding":
						__gong__assignPointer(&inst.UnplannedHolding, rhs, identifierMap)
					case "HoldingAssessment":
						__gong__assignPointer(&inst.HoldingAssessment, rhs, identifierMap)
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
