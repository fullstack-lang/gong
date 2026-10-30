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
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmfeatures"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmabstractgmlobjecttypes"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/xlink"
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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/xlink":
			aliasToCanonical[alias] = "xlink"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/gml":
			aliasToCanonical[alias] = "gml"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmabstractgmlobjecttypes":
			aliasToCanonical[alias] = "aixmabstractgmlobjecttypes"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmdatatypes":
			aliasToCanonical[alias] = "aixmdatatypes"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models/aixmfeatures":
			aliasToCanonical[alias] = "aixmfeatures"
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/service/go/models":
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
				case "AbstractTimeGeometricPrimitiveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractTimeGeometricPrimitiveType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractTimeObjectType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractTimeObjectType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractTimePrimitiveType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractTimePrimitiveType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "AbstractTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.AbstractTimeSliceType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "BoundingShapeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.BoundingShapeType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "CodeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.CodeType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "CodeWithAuthorityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.CodeWithAuthorityType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "DirectPositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.DirectPositionType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "DynamicFeatureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.DynamicFeatureType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "EnvelopeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.EnvelopeType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
				case "PointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(gml.PointType), stage.GmlStage, ident.Name, instanceName, preserveOrder)
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
				case "CodeAircraftGroundServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAircraftGroundServiceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeAviationStandardsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeAviationStandardsType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeCommunicationChannelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeCommunicationChannelType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeCommunicationDirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeCommunicationDirectionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeCommunicationModeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeCommunicationModeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeDayType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFacilityRankingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFacilityRankingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFireFightingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFireFightingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightDestinationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFlightDestinationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightRuleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFlightRuleType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeFuelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeFuelType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeIntensityStandByType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeIntensityStandByType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLanguageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLanguageType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLevelSeriesType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLevelSeriesType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLevelTableDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLevelTableDesignatorType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeLightIntensityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeLightIntensityType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryOperationsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeMilitaryOperationsType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeNitrogenType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeNitrogenType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeNorthReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeNorthReferenceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeNotePurposeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeNotePurposeType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeOilType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeOilType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeOrganisationDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeOrganisationDesignatorType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeOrganisationHierarchyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeOrganisationHierarchyType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeOrganisationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeOrganisationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeOxygenType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeOxygenType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodePassengerServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodePassengerServiceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodePilotControlledLightingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodePilotControlledLightingType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRVSMType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRVSMType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeRadioEmissionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeRadioEmissionType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeServiceATCType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeServiceATCType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeServiceATFMType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeServiceATFMType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeServiceGroundControlType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeServiceGroundControlType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeServiceInformationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeServiceInformationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeServiceSARType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeServiceSARType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeSpecialDateType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeSpecialDateType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusOperationsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeStatusOperationsType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeStatusServiceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeSystemActivationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeSystemActivationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTelecomNetworkType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTelecomNetworkType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventCombinationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTimeEventCombinationType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTimeEventType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeTimeReferenceType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeUnitDependencyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeUnitDependencyType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeUnitType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeUnitType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalDatumType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeVerticalDatumType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "CodeYesNoType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.CodeYesNoType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "DateMonthDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.DateMonthDayType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "DateYearType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.DateYearType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
				case "NoNumberType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.NoNumberType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
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
				case "ValFrequencyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmdatatypes.ValFrequencyType), stage.AixmdatatypesStage, ident.Name, instanceName, preserveOrder)
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
				case "AbstractAirportGroundServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractAirportGroundServiceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractPropertiesWithScheduleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractPropertiesWithScheduleType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractServiceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AbstractTrafficSeparationServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AbstractTrafficSeparationServiceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AerialRefuellingPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AerialRefuellingPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirTrafficControlServiceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirTrafficControlServiceTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirTrafficControlServiceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirTrafficControlServiceTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirTrafficControlServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirTrafficControlServiceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirTrafficManagementServiceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirTrafficManagementServiceTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirTrafficManagementServiceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirTrafficManagementServiceTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirTrafficManagementServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirTrafficManagementServiceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftGroundServiceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftGroundServiceTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftGroundServiceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftGroundServiceTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AircraftGroundServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AircraftGroundServiceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportClearanceServiceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportClearanceServiceTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportClearanceServiceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportClearanceServiceTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportClearanceServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportClearanceServiceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportHeliportPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportHeliportPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportSuppliesServiceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportSuppliesServiceTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportSuppliesServiceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportSuppliesServiceTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirportSuppliesServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirportSuppliesServiceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "AirspacePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.AirspacePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CallsignDetailPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CallsignDetailPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "CallsignDetailType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.CallsignDetailType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ContactInformationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ContactInformationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ContactInformationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ContactInformationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DesignatedPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DesignatedPointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "DirectionFinderPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.DirectionFinderPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedPointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedPointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ElevatedPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ElevatedPointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FireFightingServiceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FireFightingServiceTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FireFightingServiceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FireFightingServiceTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FireFightingServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FireFightingServiceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FuelPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FuelPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "FuelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.FuelType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GroundLightSystemPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GroundLightSystemPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GroundTrafficControlServiceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GroundTrafficControlServiceTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GroundTrafficControlServiceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GroundTrafficControlServiceTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "GroundTrafficControlServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.GroundTrafficControlServiceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "HoldingPatternPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.HoldingPatternPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "InformationServiceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.InformationServiceTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "InformationServiceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.InformationServiceTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "InformationServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.InformationServiceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LightActivationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LightActivationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LightActivationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LightActivationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LinguisticNotePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "LinguisticNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.LinguisticNoteType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NavaidPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NavaidPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NitrogenPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NitrogenPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NitrogenType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NitrogenType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NotePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NotePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "NoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.NoteType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OilPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OilPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OilType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OilType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OnlineContactPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OnlineContactPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OnlineContactType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OnlineContactType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OrganisationAuthorityAssociationPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OrganisationAuthorityAssociationPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OrganisationAuthorityAssociationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OrganisationAuthorityAssociationType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OrganisationAuthorityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OrganisationAuthorityPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OrganisationAuthorityTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OrganisationAuthorityTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OrganisationAuthorityTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OrganisationAuthorityTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OrganisationAuthorityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OrganisationAuthorityType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OxygenPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OxygenPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "OxygenType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.OxygenType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PassengerServiceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PassengerServiceTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PassengerServiceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PassengerServiceTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PassengerServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PassengerServiceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PilotControlledLightingTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PilotControlledLightingTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PilotControlledLightingTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PilotControlledLightingTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PilotControlledLightingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PilotControlledLightingType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PointType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PostalAddressPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PostalAddressPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "PostalAddressType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.PostalAddressType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ProcedurePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ProcedurePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RadioCommunicationChannelPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RadioCommunicationChannelPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RadioCommunicationChannelTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RadioCommunicationChannelTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RadioCommunicationChannelTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RadioCommunicationChannelTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RadioCommunicationChannelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RadioCommunicationChannelType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RadioCommunicationOperationalStatusPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RadioCommunicationOperationalStatusPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RadioCommunicationOperationalStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RadioCommunicationOperationalStatusType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RoutePortionPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RoutePortionPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RoutePortionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RoutePortionType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RoutePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RoutePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "RunwayCentrelinePointPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.RunwayCentrelinePointPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SearchRescueServiceTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SearchRescueServiceTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SearchRescueServiceTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SearchRescueServiceTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SearchRescueServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SearchRescueServiceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ServiceOperationalStatusPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ServiceOperationalStatusPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "ServiceOperationalStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.ServiceOperationalStatusType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SpecialDateTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SpecialDateTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SpecialDateTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SpecialDateTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "SpecialDateType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.SpecialDateType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelColumnPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelColumnPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelColumnTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelColumnTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelColumnTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelColumnTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelColumnType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelColumnType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelSectorTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelSectorTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelSectorTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelSectorTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelSectorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelSectorType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelTablePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelTablePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelTableTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelTableTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelTableTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelTableTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelTableType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelTableType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "StandardLevelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.StandardLevelType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
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
				case "UnitAvailabilityPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.UnitAvailabilityPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "UnitAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.UnitAvailabilityType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "UnitDependencyPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.UnitDependencyPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "UnitDependencyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.UnitDependencyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "UnitPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.UnitPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "UnitTimeSlicePropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.UnitTimeSlicePropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "UnitTimeSliceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.UnitTimeSliceType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "UnitType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.UnitType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
				case "VORPropertyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(aixmfeatures.VORPropertyType), stage.AixmfeaturesStage, ident.Name, instanceName, preserveOrder)
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
				case *gml.BoundingShapeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Envelope":
						__gong__assignPointer(&inst.Envelope, rhs, identifierMap)
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
				case *aixmdatatypes.CodeAircraftGroundServiceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeAviationStandardsType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeCommunicationChannelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeCommunicationDirectionType:
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
				case *aixmdatatypes.CodeDayType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFacilityRankingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFireFightingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeFlightDestinationType:
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
				case *aixmdatatypes.CodeFuelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeIntensityStandByType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeLanguageType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeLevelSeriesType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeLevelTableDesignatorType:
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
				case *aixmdatatypes.CodeMilitaryOperationsType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeNitrogenType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeNorthReferenceType:
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
				case *aixmdatatypes.CodeOilType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeOrganisationDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeOrganisationHierarchyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeOrganisationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeOxygenType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodePassengerServiceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodePilotControlledLightingType:
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
				case *aixmdatatypes.CodeRadioEmissionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeServiceATCType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeServiceATFMType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeServiceGroundControlType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeServiceInformationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeServiceSARType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeSpecialDateType:
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
				case *aixmdatatypes.CodeStatusServiceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeSystemActivationType:
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
				case *aixmdatatypes.CodeUnitDependencyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *aixmdatatypes.CodeUnitType:
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
				case *aixmdatatypes.ValFrequencyType:
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
					case "AbstractLightActivationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractLightActivationExtension, rhs, identifierMap)
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
					case "AbstractAircraftGroundServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAircraftGroundServiceExtension, rhs, identifierMap)
					case "AbstractAirportGroundServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportGroundServiceExtension, rhs, identifierMap)
					case "AbstractServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractServiceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_11:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAirportClearanceServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportClearanceServiceExtension, rhs, identifierMap)
					case "AbstractAirportGroundServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportGroundServiceExtension, rhs, identifierMap)
					case "AbstractServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractServiceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_12:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAirportSuppliesServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportSuppliesServiceExtension, rhs, identifierMap)
					case "AbstractAirportGroundServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportGroundServiceExtension, rhs, identifierMap)
					case "AbstractServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractServiceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_13:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAirTrafficControlServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirTrafficControlServiceExtension, rhs, identifierMap)
					case "AbstractTrafficSeparationServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTrafficSeparationServiceExtension, rhs, identifierMap)
					case "AbstractServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractServiceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_14:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractAirTrafficManagementServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirTrafficManagementServiceExtension, rhs, identifierMap)
					case "AbstractServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractServiceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_15:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractCallsignDetailExtension":
						__gong__assignSliceOfPointers(&inst.AbstractCallsignDetailExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_16:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFireFightingServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFireFightingServiceExtension, rhs, identifierMap)
					case "AbstractAirportGroundServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportGroundServiceExtension, rhs, identifierMap)
					case "AbstractServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractServiceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_17:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractFuelExtension":
						__gong__assignSliceOfPointers(&inst.AbstractFuelExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_18:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractGroundTrafficControlServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractGroundTrafficControlServiceExtension, rhs, identifierMap)
					case "AbstractTrafficSeparationServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTrafficSeparationServiceExtension, rhs, identifierMap)
					case "AbstractServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractServiceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_19:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractInformationServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractInformationServiceExtension, rhs, identifierMap)
					case "AbstractServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractServiceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_2:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractLinguisticNoteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractLinguisticNoteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_20:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractNitrogenExtension":
						__gong__assignSliceOfPointers(&inst.AbstractNitrogenExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_21:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractOilExtension":
						__gong__assignSliceOfPointers(&inst.AbstractOilExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_22:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractOxygenExtension":
						__gong__assignSliceOfPointers(&inst.AbstractOxygenExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_23:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPassengerServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPassengerServiceExtension, rhs, identifierMap)
					case "AbstractAirportGroundServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractAirportGroundServiceExtension, rhs, identifierMap)
					case "AbstractServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractServiceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_24:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPilotControlledLightingExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPilotControlledLightingExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_25:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRadioCommunicationChannelExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRadioCommunicationChannelExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_26:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractRadioCommunicationOperationalStatusExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRadioCommunicationOperationalStatusExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_27:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSearchRescueServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSearchRescueServiceExtension, rhs, identifierMap)
					case "AbstractServiceExtension":
						__gong__assignSliceOfPointers(&inst.AbstractServiceExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_28:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractServiceOperationalStatusExtension":
						__gong__assignSliceOfPointers(&inst.AbstractServiceOperationalStatusExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_29:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractContactInformationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractContactInformationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_3:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractNoteExtension":
						__gong__assignSliceOfPointers(&inst.AbstractNoteExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_30:
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
				case *aixmfeatures.A_extension_31:
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
				case *aixmfeatures.A_extension_32:
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
				case *aixmfeatures.A_extension_33:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractSpecialDateExtension":
						__gong__assignSliceOfPointers(&inst.AbstractSpecialDateExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_34:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractTimesheetExtension":
						__gong__assignSliceOfPointers(&inst.AbstractTimesheetExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_35:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractStandardLevelExtension":
						__gong__assignSliceOfPointers(&inst.AbstractStandardLevelExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_36:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractStandardLevelColumnExtension":
						__gong__assignSliceOfPointers(&inst.AbstractStandardLevelColumnExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_37:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractStandardLevelSectorExtension":
						__gong__assignSliceOfPointers(&inst.AbstractStandardLevelSectorExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_38:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractStandardLevelTableExtension":
						__gong__assignSliceOfPointers(&inst.AbstractStandardLevelTableExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_4:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractOrganisationAuthorityExtension":
						__gong__assignSliceOfPointers(&inst.AbstractOrganisationAuthorityExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_5:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractOrganisationAuthorityAssociationExtension":
						__gong__assignSliceOfPointers(&inst.AbstractOrganisationAuthorityAssociationExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_6:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractUnitExtension":
						__gong__assignSliceOfPointers(&inst.AbstractUnitExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_7:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractPropertiesWithScheduleExtension":
						__gong__assignSliceOfPointers(&inst.AbstractPropertiesWithScheduleExtension, rhs, identifierMap)
					case "AbstractUnitAvailabilityExtension":
						__gong__assignSliceOfPointers(&inst.AbstractUnitAvailabilityExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_8:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractUnitDependencyExtension":
						__gong__assignSliceOfPointers(&inst.AbstractUnitDependencyExtension, rhs, identifierMap)
					}
				case *aixmfeatures.A_extension_9:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AbstractRoutePortionExtension":
						__gong__assignSliceOfPointers(&inst.AbstractRoutePortionExtension, rhs, identifierMap)
					}
				case *aixmfeatures.AbstractAirportGroundServiceType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.AbstractServiceType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.AbstractTrafficSeparationServiceType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.AirTrafficControlServiceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirTrafficControlServiceTimeSlice":
						__gong__assignPointer(&inst.AirTrafficControlServiceTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AirTrafficControlServiceTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "FlightOperations":
						__gong__assignPointer(&inst.FlightOperations, rhs, identifierMap)
					case "Rank":
						__gong__assignPointer(&inst.Rank, rhs, identifierMap)
					case "CompliantICAO":
						__gong__assignPointer(&inst.CompliantICAO, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "ServiceProvider":
						__gong__assignPointer(&inst.ServiceProvider, rhs, identifierMap)
					case "Call_sign":
						__gong__assignSliceOfPointers(&inst.Call_sign, rhs, identifierMap)
					case "RadioCommunication":
						__gong__assignSliceOfPointers(&inst.RadioCommunication, rhs, identifierMap)
					case "GroundCommunication":
						__gong__assignSliceOfPointers(&inst.GroundCommunication, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "RadarAssisted":
						__gong__assignPointer(&inst.RadarAssisted, rhs, identifierMap)
					case "DataLinkEnabled":
						__gong__assignPointer(&inst.DataLinkEnabled, rhs, identifierMap)
					case "DataLinkChannel":
						__gong__assignPointer(&inst.DataLinkChannel, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "ClientAirport":
						__gong__assignSliceOfPointers(&inst.ClientAirport, rhs, identifierMap)
					case "ClientAirspace":
						__gong__assignSliceOfPointers(&inst.ClientAirspace, rhs, identifierMap)
					case "ClientRoute":
						__gong__assignSliceOfPointers(&inst.ClientRoute, rhs, identifierMap)
					case "ClientProcedure":
						__gong__assignSliceOfPointers(&inst.ClientProcedure, rhs, identifierMap)
					case "ClientHolding":
						__gong__assignSliceOfPointers(&inst.ClientHolding, rhs, identifierMap)
					case "ClientAerialRefuelling":
						__gong__assignSliceOfPointers(&inst.ClientAerialRefuelling, rhs, identifierMap)
					case "AircraftLocator":
						__gong__assignPointer(&inst.AircraftLocator, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AirTrafficControlServiceType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.AirTrafficManagementServiceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirTrafficManagementServiceTimeSlice":
						__gong__assignPointer(&inst.AirTrafficManagementServiceTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AirTrafficManagementServiceTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "FlightOperations":
						__gong__assignPointer(&inst.FlightOperations, rhs, identifierMap)
					case "Rank":
						__gong__assignPointer(&inst.Rank, rhs, identifierMap)
					case "CompliantICAO":
						__gong__assignPointer(&inst.CompliantICAO, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "ServiceProvider":
						__gong__assignPointer(&inst.ServiceProvider, rhs, identifierMap)
					case "Call_sign":
						__gong__assignSliceOfPointers(&inst.Call_sign, rhs, identifierMap)
					case "RadioCommunication":
						__gong__assignSliceOfPointers(&inst.RadioCommunication, rhs, identifierMap)
					case "GroundCommunication":
						__gong__assignSliceOfPointers(&inst.GroundCommunication, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "ClientAirspace":
						__gong__assignSliceOfPointers(&inst.ClientAirspace, rhs, identifierMap)
					case "ClientAerialRefuelling":
						__gong__assignSliceOfPointers(&inst.ClientAerialRefuelling, rhs, identifierMap)
					case "ClientRoute":
						__gong__assignSliceOfPointers(&inst.ClientRoute, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AirTrafficManagementServiceType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.AircraftGroundServiceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AircraftGroundServiceTimeSlice":
						__gong__assignPointer(&inst.AircraftGroundServiceTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AircraftGroundServiceTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "FlightOperations":
						__gong__assignPointer(&inst.FlightOperations, rhs, identifierMap)
					case "Rank":
						__gong__assignPointer(&inst.Rank, rhs, identifierMap)
					case "CompliantICAO":
						__gong__assignPointer(&inst.CompliantICAO, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "ServiceProvider":
						__gong__assignPointer(&inst.ServiceProvider, rhs, identifierMap)
					case "Call_sign":
						__gong__assignSliceOfPointers(&inst.Call_sign, rhs, identifierMap)
					case "RadioCommunication":
						__gong__assignSliceOfPointers(&inst.RadioCommunication, rhs, identifierMap)
					case "GroundCommunication":
						__gong__assignSliceOfPointers(&inst.GroundCommunication, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignSliceOfPointers(&inst.AirportHeliport, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AircraftGroundServiceType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.AirportClearanceServiceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirportClearanceServiceTimeSlice":
						__gong__assignPointer(&inst.AirportClearanceServiceTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AirportClearanceServiceTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "FlightOperations":
						__gong__assignPointer(&inst.FlightOperations, rhs, identifierMap)
					case "Rank":
						__gong__assignPointer(&inst.Rank, rhs, identifierMap)
					case "CompliantICAO":
						__gong__assignPointer(&inst.CompliantICAO, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "ServiceProvider":
						__gong__assignPointer(&inst.ServiceProvider, rhs, identifierMap)
					case "Call_sign":
						__gong__assignSliceOfPointers(&inst.Call_sign, rhs, identifierMap)
					case "RadioCommunication":
						__gong__assignSliceOfPointers(&inst.RadioCommunication, rhs, identifierMap)
					case "GroundCommunication":
						__gong__assignSliceOfPointers(&inst.GroundCommunication, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignSliceOfPointers(&inst.AirportHeliport, rhs, identifierMap)
					case "SnowPlan":
						__gong__assignPointer(&inst.SnowPlan, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AirportClearanceServiceType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.AirportSuppliesServiceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "AirportSuppliesServiceTimeSlice":
						__gong__assignPointer(&inst.AirportSuppliesServiceTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.AirportSuppliesServiceTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "FlightOperations":
						__gong__assignPointer(&inst.FlightOperations, rhs, identifierMap)
					case "Rank":
						__gong__assignPointer(&inst.Rank, rhs, identifierMap)
					case "CompliantICAO":
						__gong__assignPointer(&inst.CompliantICAO, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "ServiceProvider":
						__gong__assignPointer(&inst.ServiceProvider, rhs, identifierMap)
					case "Call_sign":
						__gong__assignSliceOfPointers(&inst.Call_sign, rhs, identifierMap)
					case "RadioCommunication":
						__gong__assignSliceOfPointers(&inst.RadioCommunication, rhs, identifierMap)
					case "GroundCommunication":
						__gong__assignSliceOfPointers(&inst.GroundCommunication, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignSliceOfPointers(&inst.AirportHeliport, rhs, identifierMap)
					case "FuelSupply":
						__gong__assignSliceOfPointers(&inst.FuelSupply, rhs, identifierMap)
					case "OilSupply":
						__gong__assignSliceOfPointers(&inst.OilSupply, rhs, identifierMap)
					case "NitrogenSupply":
						__gong__assignSliceOfPointers(&inst.NitrogenSupply, rhs, identifierMap)
					case "OxygenSupply":
						__gong__assignSliceOfPointers(&inst.OxygenSupply, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.AirportSuppliesServiceType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.CallsignDetailPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "CallsignDetail":
						__gong__assignPointer(&inst.CallsignDetail, rhs, identifierMap)
					}
				case *aixmfeatures.CallsignDetailType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "CallSign":
						__gong__assignPointer(&inst.CallSign, rhs, identifierMap)
					case "Language":
						__gong__assignPointer(&inst.Language, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
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
				case *aixmfeatures.DirectionFinderPropertyType:
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
				case *aixmfeatures.FireFightingServiceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "FireFightingServiceTimeSlice":
						__gong__assignPointer(&inst.FireFightingServiceTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.FireFightingServiceTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "FlightOperations":
						__gong__assignPointer(&inst.FlightOperations, rhs, identifierMap)
					case "Rank":
						__gong__assignPointer(&inst.Rank, rhs, identifierMap)
					case "CompliantICAO":
						__gong__assignPointer(&inst.CompliantICAO, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "ServiceProvider":
						__gong__assignPointer(&inst.ServiceProvider, rhs, identifierMap)
					case "Call_sign":
						__gong__assignSliceOfPointers(&inst.Call_sign, rhs, identifierMap)
					case "RadioCommunication":
						__gong__assignSliceOfPointers(&inst.RadioCommunication, rhs, identifierMap)
					case "GroundCommunication":
						__gong__assignSliceOfPointers(&inst.GroundCommunication, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignSliceOfPointers(&inst.AirportHeliport, rhs, identifierMap)
					case "Category":
						__gong__assignPointer(&inst.Category, rhs, identifierMap)
					case "Standard":
						__gong__assignPointer(&inst.Standard, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.FireFightingServiceType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.FuelPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Fuel":
						__gong__assignPointer(&inst.Fuel, rhs, identifierMap)
					}
				case *aixmfeatures.FuelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Category":
						__gong__assignPointer(&inst.Category, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
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
				case *aixmfeatures.GroundTrafficControlServiceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "GroundTrafficControlServiceTimeSlice":
						__gong__assignPointer(&inst.GroundTrafficControlServiceTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.GroundTrafficControlServiceTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "FlightOperations":
						__gong__assignPointer(&inst.FlightOperations, rhs, identifierMap)
					case "Rank":
						__gong__assignPointer(&inst.Rank, rhs, identifierMap)
					case "CompliantICAO":
						__gong__assignPointer(&inst.CompliantICAO, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "ServiceProvider":
						__gong__assignPointer(&inst.ServiceProvider, rhs, identifierMap)
					case "Call_sign":
						__gong__assignSliceOfPointers(&inst.Call_sign, rhs, identifierMap)
					case "RadioCommunication":
						__gong__assignSliceOfPointers(&inst.RadioCommunication, rhs, identifierMap)
					case "GroundCommunication":
						__gong__assignSliceOfPointers(&inst.GroundCommunication, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "RadarAssisted":
						__gong__assignPointer(&inst.RadarAssisted, rhs, identifierMap)
					case "DataLinkEnabled":
						__gong__assignPointer(&inst.DataLinkEnabled, rhs, identifierMap)
					case "DataLinkChannel":
						__gong__assignPointer(&inst.DataLinkChannel, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "ClientAirport":
						__gong__assignPointer(&inst.ClientAirport, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.GroundTrafficControlServiceType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.InformationServiceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "InformationServiceTimeSlice":
						__gong__assignPointer(&inst.InformationServiceTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.InformationServiceTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "FlightOperations":
						__gong__assignPointer(&inst.FlightOperations, rhs, identifierMap)
					case "Rank":
						__gong__assignPointer(&inst.Rank, rhs, identifierMap)
					case "CompliantICAO":
						__gong__assignPointer(&inst.CompliantICAO, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "ServiceProvider":
						__gong__assignPointer(&inst.ServiceProvider, rhs, identifierMap)
					case "Call_sign":
						__gong__assignSliceOfPointers(&inst.Call_sign, rhs, identifierMap)
					case "RadioCommunication":
						__gong__assignSliceOfPointers(&inst.RadioCommunication, rhs, identifierMap)
					case "GroundCommunication":
						__gong__assignSliceOfPointers(&inst.GroundCommunication, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Voice":
						__gong__assignPointer(&inst.Voice, rhs, identifierMap)
					case "DataLink":
						__gong__assignPointer(&inst.DataLink, rhs, identifierMap)
					case "Recorded":
						__gong__assignPointer(&inst.Recorded, rhs, identifierMap)
					case "NavaidBroadcast":
						__gong__assignSliceOfPointers(&inst.NavaidBroadcast, rhs, identifierMap)
					case "ClientAirspace":
						__gong__assignSliceOfPointers(&inst.ClientAirspace, rhs, identifierMap)
					case "ClientAirport":
						__gong__assignSliceOfPointers(&inst.ClientAirport, rhs, identifierMap)
					case "ClientRoute":
						__gong__assignSliceOfPointers(&inst.ClientRoute, rhs, identifierMap)
					case "ClientProcedure":
						__gong__assignSliceOfPointers(&inst.ClientProcedure, rhs, identifierMap)
					case "ClientHolding":
						__gong__assignSliceOfPointers(&inst.ClientHolding, rhs, identifierMap)
					case "ClientAerialRefuelling":
						__gong__assignSliceOfPointers(&inst.ClientAerialRefuelling, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.InformationServiceType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.LightActivationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "LightActivation":
						__gong__assignPointer(&inst.LightActivation, rhs, identifierMap)
					}
				case *aixmfeatures.LightActivationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Clicks":
						__gong__assignPointer(&inst.Clicks, rhs, identifierMap)
					case "IntensityLevel":
						__gong__assignPointer(&inst.IntensityLevel, rhs, identifierMap)
					case "Activation":
						__gong__assignPointer(&inst.Activation, rhs, identifierMap)
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
				case *aixmfeatures.NitrogenPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Nitrogen":
						__gong__assignPointer(&inst.Nitrogen, rhs, identifierMap)
					}
				case *aixmfeatures.NitrogenType:
					switch fieldName {
					case "Name":
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
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *aixmfeatures.OilPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Oil":
						__gong__assignPointer(&inst.Oil, rhs, identifierMap)
					}
				case *aixmfeatures.OilType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "Category":
						__gong__assignPointer(&inst.Category, rhs, identifierMap)
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
				case *aixmfeatures.OrganisationAuthorityAssociationPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "OrganisationAuthorityAssociation":
						__gong__assignPointer(&inst.OrganisationAuthorityAssociation, rhs, identifierMap)
					}
				case *aixmfeatures.OrganisationAuthorityAssociationType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.OrganisationAuthorityTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "OrganisationAuthorityTimeSlice":
						__gong__assignPointer(&inst.OrganisationAuthorityTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.OrganisationAuthorityTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "Military":
						__gong__assignPointer(&inst.Military, rhs, identifierMap)
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "RelatedOrganisationAuthority":
						__gong__assignSliceOfPointers(&inst.RelatedOrganisationAuthority, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.OrganisationAuthorityType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.OxygenPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "Oxygen":
						__gong__assignPointer(&inst.Oxygen, rhs, identifierMap)
					}
				case *aixmfeatures.OxygenType:
					switch fieldName {
					case "Name":
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
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.PassengerServiceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "PassengerServiceTimeSlice":
						__gong__assignPointer(&inst.PassengerServiceTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.PassengerServiceTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "FlightOperations":
						__gong__assignPointer(&inst.FlightOperations, rhs, identifierMap)
					case "Rank":
						__gong__assignPointer(&inst.Rank, rhs, identifierMap)
					case "CompliantICAO":
						__gong__assignPointer(&inst.CompliantICAO, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "ServiceProvider":
						__gong__assignPointer(&inst.ServiceProvider, rhs, identifierMap)
					case "Call_sign":
						__gong__assignSliceOfPointers(&inst.Call_sign, rhs, identifierMap)
					case "RadioCommunication":
						__gong__assignSliceOfPointers(&inst.RadioCommunication, rhs, identifierMap)
					case "GroundCommunication":
						__gong__assignSliceOfPointers(&inst.GroundCommunication, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "AirportHeliport":
						__gong__assignSliceOfPointers(&inst.AirportHeliport, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.PassengerServiceType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.PilotControlledLightingTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "PilotControlledLightingTimeSlice":
						__gong__assignPointer(&inst.PilotControlledLightingTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.PilotControlledLightingTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "Duration":
						__gong__assignPointer(&inst.Duration, rhs, identifierMap)
					case "IntensitySteps":
						__gong__assignPointer(&inst.IntensitySteps, rhs, identifierMap)
					case "StandByIntensity":
						__gong__assignPointer(&inst.StandByIntensity, rhs, identifierMap)
					case "RadioFrequency":
						__gong__assignPointer(&inst.RadioFrequency, rhs, identifierMap)
					case "ActivationInstruction":
						__gong__assignPointer(&inst.ActivationInstruction, rhs, identifierMap)
					case "ControlledLightIntensity":
						__gong__assignSliceOfPointers(&inst.ControlledLightIntensity, rhs, identifierMap)
					case "ActivatedGroundLighting":
						__gong__assignSliceOfPointers(&inst.ActivatedGroundLighting, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.PilotControlledLightingType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.RadioCommunicationChannelPropertyType:
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
				case *aixmfeatures.RadioCommunicationChannelTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "RadioCommunicationChannelTimeSlice":
						__gong__assignPointer(&inst.RadioCommunicationChannelTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.RadioCommunicationChannelTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "Mode":
						__gong__assignPointer(&inst.Mode, rhs, identifierMap)
					case "Rank":
						__gong__assignPointer(&inst.Rank, rhs, identifierMap)
					case "FrequencyTransmission":
						__gong__assignPointer(&inst.FrequencyTransmission, rhs, identifierMap)
					case "FrequencyReception":
						__gong__assignPointer(&inst.FrequencyReception, rhs, identifierMap)
					case "Channel":
						__gong__assignPointer(&inst.Channel, rhs, identifierMap)
					case "Logon":
						__gong__assignPointer(&inst.Logon, rhs, identifierMap)
					case "EmissionType":
						__gong__assignPointer(&inst.EmissionType, rhs, identifierMap)
					case "SelectiveCall":
						__gong__assignPointer(&inst.SelectiveCall, rhs, identifierMap)
					case "FlightChecked":
						__gong__assignPointer(&inst.FlightChecked, rhs, identifierMap)
					case "TrafficDirection":
						__gong__assignPointer(&inst.TrafficDirection, rhs, identifierMap)
					case "Location":
						__gong__assignSliceOfPointers(&inst.Location, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.RadioCommunicationChannelType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.RadioCommunicationOperationalStatusPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "RadioCommunicationOperationalStatus":
						__gong__assignPointer(&inst.RadioCommunicationOperationalStatus, rhs, identifierMap)
					}
				case *aixmfeatures.RadioCommunicationOperationalStatusType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.SearchRescueServiceTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SearchRescueServiceTimeSlice":
						__gong__assignPointer(&inst.SearchRescueServiceTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.SearchRescueServiceTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "FlightOperations":
						__gong__assignPointer(&inst.FlightOperations, rhs, identifierMap)
					case "Rank":
						__gong__assignPointer(&inst.Rank, rhs, identifierMap)
					case "CompliantICAO":
						__gong__assignPointer(&inst.CompliantICAO, rhs, identifierMap)
					case "Location":
						__gong__assignPointer(&inst.Location, rhs, identifierMap)
					case "ServiceProvider":
						__gong__assignPointer(&inst.ServiceProvider, rhs, identifierMap)
					case "Call_sign":
						__gong__assignSliceOfPointers(&inst.Call_sign, rhs, identifierMap)
					case "RadioCommunication":
						__gong__assignSliceOfPointers(&inst.RadioCommunication, rhs, identifierMap)
					case "GroundCommunication":
						__gong__assignSliceOfPointers(&inst.GroundCommunication, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Type_XSD":
						__gong__assignPointer(&inst.Type_XSD, rhs, identifierMap)
					case "ClientAirspace":
						__gong__assignSliceOfPointers(&inst.ClientAirspace, rhs, identifierMap)
					case "ClientRoute":
						__gong__assignSliceOfPointers(&inst.ClientRoute, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.SearchRescueServiceType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.ServiceOperationalStatusPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "ServiceOperationalStatus":
						__gong__assignPointer(&inst.ServiceOperationalStatus, rhs, identifierMap)
					}
				case *aixmfeatures.ServiceOperationalStatusType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.SpecialDateTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "SpecialDateTimeSlice":
						__gong__assignPointer(&inst.SpecialDateTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.SpecialDateTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "DateDay":
						__gong__assignPointer(&inst.DateDay, rhs, identifierMap)
					case "DateYear":
						__gong__assignPointer(&inst.DateYear, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Authority":
						__gong__assignPointer(&inst.Authority, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.SpecialDateType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.StandardLevelColumnTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "StandardLevelColumnTimeSlice":
						__gong__assignPointer(&inst.StandardLevelColumnTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.StandardLevelColumnTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "Series":
						__gong__assignPointer(&inst.Series, rhs, identifierMap)
					case "UnitOfMeasurement":
						inst.UnitOfMeasurement = GongExtractString(rhs)
					case "Separation":
						__gong__assignPointer(&inst.Separation, rhs, identifierMap)
					case "Level":
						__gong__assignSliceOfPointers(&inst.Level, rhs, identifierMap)
					case "LevelTable":
						__gong__assignPointer(&inst.LevelTable, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.StandardLevelColumnType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.StandardLevelPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "StandardLevel":
						__gong__assignPointer(&inst.StandardLevel, rhs, identifierMap)
					}
				case *aixmfeatures.StandardLevelSectorTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "StandardLevelSectorTimeSlice":
						__gong__assignPointer(&inst.StandardLevelSectorTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.StandardLevelSectorTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "FlightRule":
						__gong__assignPointer(&inst.FlightRule, rhs, identifierMap)
					case "FromTrack":
						__gong__assignPointer(&inst.FromTrack, rhs, identifierMap)
					case "ToTrack":
						__gong__assignPointer(&inst.ToTrack, rhs, identifierMap)
					case "AngleType":
						__gong__assignPointer(&inst.AngleType, rhs, identifierMap)
					case "ApplicableAirspace":
						__gong__assignSliceOfPointers(&inst.ApplicableAirspace, rhs, identifierMap)
					case "ApplicableLevelColumn":
						__gong__assignPointer(&inst.ApplicableLevelColumn, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.StandardLevelSectorType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.StandardLevelTablePropertyType:
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
				case *aixmfeatures.StandardLevelTableTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "StandardLevelTableTimeSlice":
						__gong__assignPointer(&inst.StandardLevelTableTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.StandardLevelTableTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "StandardICAO":
						__gong__assignPointer(&inst.StandardICAO, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.StandardLevelTableType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.StandardLevelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Id":
						inst.Id = GongExtractString(rhs)
					case "Description":
						inst.Description = GongExtractString(rhs)
					case "Identifier":
						__gong__assignPointer(&inst.Identifier, rhs, identifierMap)
					case "NameXSD":
						__gong__assignSliceOfPointers(&inst.NameXSD, rhs, identifierMap)
					case "VerticalDistance":
						__gong__assignPointer(&inst.VerticalDistance, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
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
				case *aixmfeatures.UnitAvailabilityPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "UnitAvailability":
						__gong__assignPointer(&inst.UnitAvailability, rhs, identifierMap)
					}
				case *aixmfeatures.UnitAvailabilityType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.UnitDependencyPropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "UnitDependency":
						__gong__assignPointer(&inst.UnitDependency, rhs, identifierMap)
					}
				case *aixmfeatures.UnitDependencyType:
					switch fieldName {
					case "Name":
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
					case "TheUnit":
						__gong__assignPointer(&inst.TheUnit, rhs, identifierMap)
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
				case *aixmfeatures.UnitTimeSlicePropertyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "Owns":
						inst.Owns = GongExtractBool(rhs)
					case "UnitTimeSlice":
						__gong__assignPointer(&inst.UnitTimeSlice, rhs, identifierMap)
					}
				case *aixmfeatures.UnitTimeSliceType:
					switch fieldName {
					case "Name":
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
					case "CompliantICAO":
						__gong__assignPointer(&inst.CompliantICAO, rhs, identifierMap)
					case "Designator":
						__gong__assignPointer(&inst.Designator, rhs, identifierMap)
					case "Military":
						__gong__assignPointer(&inst.Military, rhs, identifierMap)
					case "Position":
						__gong__assignPointer(&inst.Position, rhs, identifierMap)
					case "AirportLocation":
						__gong__assignPointer(&inst.AirportLocation, rhs, identifierMap)
					case "OwnerOrganisation":
						__gong__assignPointer(&inst.OwnerOrganisation, rhs, identifierMap)
					case "Contact":
						__gong__assignSliceOfPointers(&inst.Contact, rhs, identifierMap)
					case "RelatedUnit":
						__gong__assignSliceOfPointers(&inst.RelatedUnit, rhs, identifierMap)
					case "Availability":
						__gong__assignSliceOfPointers(&inst.Availability, rhs, identifierMap)
					case "Annotation":
						__gong__assignSliceOfPointers(&inst.Annotation, rhs, identifierMap)
					case "Extension":
						__gong__assignSliceOfPointers(&inst.Extension, rhs, identifierMap)
					}
				case *aixmfeatures.UnitType:
					switch fieldName {
					case "Name":
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
				case *aixmfeatures.VORPropertyType:
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
					case "OrganisationAuthority":
						__gong__assignPointer(&inst.OrganisationAuthority, rhs, identifierMap)
					case "Unit":
						__gong__assignPointer(&inst.Unit, rhs, identifierMap)
					case "AirTrafficControlService":
						__gong__assignPointer(&inst.AirTrafficControlService, rhs, identifierMap)
					case "AirTrafficManagementService":
						__gong__assignPointer(&inst.AirTrafficManagementService, rhs, identifierMap)
					case "GroundTrafficControlService":
						__gong__assignPointer(&inst.GroundTrafficControlService, rhs, identifierMap)
					case "InformationService":
						__gong__assignPointer(&inst.InformationService, rhs, identifierMap)
					case "SearchRescueService":
						__gong__assignPointer(&inst.SearchRescueService, rhs, identifierMap)
					case "AircraftGroundService":
						__gong__assignPointer(&inst.AircraftGroundService, rhs, identifierMap)
					case "AirportClearanceService":
						__gong__assignPointer(&inst.AirportClearanceService, rhs, identifierMap)
					case "AirportSuppliesService":
						__gong__assignPointer(&inst.AirportSuppliesService, rhs, identifierMap)
					case "FireFightingService":
						__gong__assignPointer(&inst.FireFightingService, rhs, identifierMap)
					case "PassengerService":
						__gong__assignPointer(&inst.PassengerService, rhs, identifierMap)
					case "PilotControlledLighting":
						__gong__assignPointer(&inst.PilotControlledLighting, rhs, identifierMap)
					case "RadioCommunicationChannel":
						__gong__assignPointer(&inst.RadioCommunicationChannel, rhs, identifierMap)
					case "SpecialDate":
						__gong__assignPointer(&inst.SpecialDate, rhs, identifierMap)
					case "StandardLevelTable":
						__gong__assignPointer(&inst.StandardLevelTable, rhs, identifierMap)
					case "StandardLevelColumn":
						__gong__assignPointer(&inst.StandardLevelColumn, rhs, identifierMap)
					case "StandardLevelSector":
						__gong__assignPointer(&inst.StandardLevelSector, rhs, identifierMap)
					case "StandardLevel":
						__gong__assignPointer(&inst.StandardLevel, rhs, identifierMap)
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
