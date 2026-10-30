// generated code - do not edit
package aixmdatatypes

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
		case "github.com/fullstack-lang/gong/app/xsd/tests/aixm/go/models/aixmdatatypes":
			aliasToCanonical[alias] = "aixmdatatypes"
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
					} else if pkgAlias == "" || pkgAlias == "aixmdatatypes" || len(inFile.Imports) == 0 {
						pkgAlias = "aixmdatatypes"
					}

					switch pkgAlias {
			case "aixmdatatypes":
				switch typeName {
				case "CodeATCReportingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeATCReportingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAerialRefuellingPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAerialRefuellingPointType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAerialRefuellingPrefixType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAerialRefuellingPrefixType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAerialRefuellingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAerialRefuellingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftCategoryType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftCategoryType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftEngineNumberType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftEngineNumberType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftEngineType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftEngineType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftGroundServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftGroundServiceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftICAOType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftICAOType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftStandType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftStandType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAircraftWingspanClassType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAircraftWingspanClassType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAirportHeliportCollocationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAirportHeliportCollocationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAirportHeliportDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAirportHeliportDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAirportHeliportType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAirportHeliportType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAirportWarningType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAirportWarningType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAirspaceActivityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAirspaceActivityType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAirspaceAggregationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAirspaceAggregationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAirspaceClassificationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAirspaceClassificationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAirspaceDependencyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAirspaceDependencyType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAirspaceDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAirspaceDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAirspacePointPositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAirspacePointPositionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAirspacePointRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAirspacePointRoleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAirspaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAirspaceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAltitudeAdjustmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAltitudeAdjustmentType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAltitudeUseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAltitudeUseType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeApproachEquipmentAdditionalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeApproachEquipmentAdditionalType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeApproachGuidanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeApproachGuidanceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeApproachLightingICAOType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeApproachLightingICAOType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeApproachLightingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeApproachLightingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeApproachPrefixType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeApproachPrefixType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeApproachType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeApproachType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeApprovalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeApprovalType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeApronElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeApronElementType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeApronSectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeApronSectionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeArcDirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeArcDirectionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeArrestingGearEnergyAbsorbType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeArrestingGearEnergyAbsorbType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeArrestingGearEngageDeviceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeArrestingGearEngageDeviceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAuralMorseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAuralMorseType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAuthorityRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAuthorityRoleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAuthorityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAuthorityType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeAviationStandardsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeAviationStandardsType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeBearingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeBearingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeBuoyDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeBuoyDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeBuoyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeBuoyType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeCardinalDirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeCardinalDirectionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeCheckpointCategoryType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeCheckpointCategoryType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeColourType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeColourType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeCommunicationChannelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeCommunicationChannelType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeCommunicationDirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeCommunicationDirectionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeCommunicationModeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeCommunicationModeType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeComparisonType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeComparisonType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeContaminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeContaminationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeCourseQualityILSType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeCourseQualityILSType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeCourseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeCourseType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDMEChannelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDMEChannelType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDMEType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDMEType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDayType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDeclaredDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDeclaredDistanceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDesignStandardType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDesignStandardType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDesignatedPointDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDesignatedPointDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDesignatedPointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDesignatedPointType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDirectionReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDirectionReferenceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDirectionTurnType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDirectionTurnType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDirectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDirectionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeDistanceIndicationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeDistanceIndicationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeEmissionBandType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeEmissionBandType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeEquipmentAntiCollisionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeEquipmentAntiCollisionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeEquipmentUnavailableType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeEquipmentUnavailableType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFacilityRankingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFacilityRankingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFinalGuidanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFinalGuidanceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFireFightingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFireFightingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightDestinationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightDestinationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightOriginType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightOriginType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightPurposeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightPurposeType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightRestrictionDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightRestrictionDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightRestrictionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightRestrictionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightRuleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightRuleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightStatusType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlightType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFlowConditionOperationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFlowConditionOperationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFreeFlightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFreeFlightType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFrictionDeviceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFrictionDeviceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFrictionEstimateType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFrictionEstimateType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeFuelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeFuelType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeGeoBorderType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeGeoBorderType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeGradeSeparationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeGradeSeparationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeGroundLightingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeGroundLightingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeGuidanceLineType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeGuidanceLineType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeHeightReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeHeightReferenceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeHelicopterPerformanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeHelicopterPerformanceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeHoldingCategoryType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeHoldingCategoryType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeHoldingUsageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeHoldingUsageType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeHoldingUseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeHoldingUseType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeIATAType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeIATAType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeICAOCountryType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeICAOCountryType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeICAOType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeICAOType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeILSBackCourseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeILSBackCourseType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeIntegrityLevelILSType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeIntegrityLevelILSType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeIntensityStandByType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeIntensityStandByType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLanguageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLanguageType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLevelSeriesType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLevelSeriesType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLevelTableDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLevelTableDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLevelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLevelType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLightHoldingPositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLightHoldingPositionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLightIntensityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLightIntensityType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLightSourceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLightSourceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLightingJARType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLightingJARType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLoadingBridgeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLoadingBridgeType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLocationQualifierType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLocationQualifierType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeLogicalOperatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeLogicalOperatorType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMLSAzimuthType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMLSAzimuthType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMLSChannelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMLSChannelType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMarkerBeaconSignalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMarkerBeaconSignalType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMarkingConditionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMarkingConditionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMarkingStyleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMarkingStyleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMeteoConditionsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMeteoConditionsType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryOperationsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMilitaryOperationsType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryRoutePointType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMilitaryRoutePointType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryStatusType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMilitaryStatusType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMilitaryTrainingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMilitaryTrainingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMinimaFinalApproachPathType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMinimaFinalApproachPathType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMinimumAltitudeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMinimumAltitudeType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMinimumHeightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMinimumHeightType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeMissedApproachType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeMissedApproachType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNDBUsageType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNDBUsageType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNavaidDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNavaidDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNavaidPurposeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNavaidPurposeType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNavaidServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNavaidServiceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNavigationAreaRestrictionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNavigationAreaRestrictionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNavigationAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNavigationAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNavigationEquipmentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNavigationEquipmentType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNavigationSpecificationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNavigationSpecificationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNitrogenType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNitrogenType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNorthReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNorthReferenceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeNotePurposeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeNotePurposeType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeObstacleAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeObstacleAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeObstacleAssessmentSurfaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeObstacleAssessmentSurfaceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeObstructionIdSurfaceZoneType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeObstructionIdSurfaceZoneType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeOilType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeOilType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeOperationAirportHeliportType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeOperationAirportHeliportType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeOperationManoeuvringAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeOperationManoeuvringAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeOrganisationDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeOrganisationDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeOrganisationHierarchyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeOrganisationHierarchyType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeOrganisationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeOrganisationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeOxygenType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeOxygenType), stage, ident.Name, instanceName, preserveOrder)
				case "CodePARType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodePARType), stage, ident.Name, instanceName, preserveOrder)
				case "CodePCNMethodType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodePCNMethodType), stage, ident.Name, instanceName, preserveOrder)
				case "CodePCNPavementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodePCNPavementType), stage, ident.Name, instanceName, preserveOrder)
				case "CodePCNSubgradeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodePCNSubgradeType), stage, ident.Name, instanceName, preserveOrder)
				case "CodePCNTyrePressureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodePCNTyrePressureType), stage, ident.Name, instanceName, preserveOrder)
				case "CodePassengerServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodePassengerServiceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodePilotControlledLightingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodePilotControlledLightingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodePositionInILSType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodePositionInILSType), stage, ident.Name, instanceName, preserveOrder)
				case "CodePrimaryRadarType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodePrimaryRadarType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeProcedureAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeProcedureAvailabilityType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeProcedureCodingStandardType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeProcedureCodingStandardType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeProcedureDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeProcedureDistanceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeProcedureFixRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeProcedureFixRoleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeProcedurePhaseType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeProcedurePhaseType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeProtectAreaSectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeProtectAreaSectionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRNPType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRNPType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRVRReadingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRVRReadingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRVSMPointRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRVSMPointRoleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRVSMType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRVSMType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRadarServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRadarServiceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRadioEmissionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRadioEmissionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRadioFrequencyAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRadioFrequencyAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRadioSignalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRadioSignalType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeReferencePathIdentifierType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeReferencePathIdentifierType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeReferenceRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeReferenceRoleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeReflectorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeReflectorType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRelativePositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRelativePositionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRoadType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRoadType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteAvailabilityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteAvailabilityType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteDesignatorLetterType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteDesignatorLetterType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteDesignatorPrefixType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteDesignatorPrefixType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteDesignatorSuffixType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteDesignatorSuffixType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteIndicatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteIndicatorType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteNavigationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteNavigationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteOriginType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteOriginType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteSegmentPathType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteSegmentPathType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRouteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRouteType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRuleProcedureTitleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRuleProcedureTitleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRuleProcedureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRuleProcedureType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRunwayElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRunwayElementType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRunwayMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRunwayMarkingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRunwayPointRoleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRunwayPointRoleType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRunwayProtectionAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRunwayProtectionAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRunwaySectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRunwaySectionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeRunwayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeRunwayType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSafeAltitudeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSafeAltitudeType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSegmentPathType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSegmentPathType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSegmentTerminationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSegmentTerminationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeServiceATCType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeServiceATCType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeServiceATFMType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeServiceATFMType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeServiceGroundControlType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeServiceGroundControlType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeServiceInformationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeServiceInformationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeServiceSARType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeServiceSARType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSideType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSideType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSignalPerformanceILSType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSignalPerformanceILSType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSpecialDateType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSpecialDateType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSpecialNavigationChainDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSpecialNavigationChainDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSpecialNavigationStationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSpecialNavigationStationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSpecialNavigationSystemType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSpecialNavigationSystemType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSpeedReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSpeedReferenceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeStandbyPowerType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeStandbyPowerType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusAirportType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeStatusAirportType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusAirspaceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeStatusAirspaceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusConstructionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeStatusConstructionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusNavaidType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeStatusNavaidType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusOperationsType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeStatusOperationsType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeStatusServiceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeStatusServiceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSurfaceCompositionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSurfaceCompositionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSurfaceConditionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSurfaceConditionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSurfacePreparationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSurfacePreparationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeSystemActivationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeSystemActivationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTAAType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTAAType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTACANChannelType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTACANChannelType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTLOFSectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTLOFSectionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTaxiwayElementType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTaxiwayElementType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTaxiwaySectionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTaxiwaySectionType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTaxiwayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTaxiwayType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTelecomNetworkType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTelecomNetworkType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventCombinationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTimeEventCombinationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeEventType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTimeEventType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTimeReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTimeReferenceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTrajectoryType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTrajectoryType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeTransponderType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeTransponderType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeUUIDType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeUUIDType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeUnitDependencyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeUnitDependencyType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeUnitType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeUnitType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeUpperAlphaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeUpperAlphaType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeUsageLimitationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeUsageLimitationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVASISType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVASISType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVORType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVORType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeValueInterpretationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeValueInterpretationType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalDatumType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVerticalDatumType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalReferenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVerticalReferenceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalStructureMarkingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVerticalStructureMarkingType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalStructureMaterialType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVerticalStructureMaterialType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVerticalStructureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVerticalStructureType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeVisualDockingGuidanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeVisualDockingGuidanceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeWakeTurbulenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeWakeTurbulenceType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeWorkAreaType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeWorkAreaType), stage, ident.Name, instanceName, preserveOrder)
				case "CodeYesNoType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(CodeYesNoType), stage, ident.Name, instanceName, preserveOrder)
				case "DateMonthDayType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DateMonthDayType), stage, ident.Name, instanceName, preserveOrder)
				case "DateTimeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DateTimeType), stage, ident.Name, instanceName, preserveOrder)
				case "DateType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DateType), stage, ident.Name, instanceName, preserveOrder)
				case "DateYearType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(DateYearType), stage, ident.Name, instanceName, preserveOrder)
				case "NoNumberType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoNumberType), stage, ident.Name, instanceName, preserveOrder)
				case "NoSequenceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(NoSequenceType), stage, ident.Name, instanceName, preserveOrder)
				case "TextAddressType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextAddressType), stage, ident.Name, instanceName, preserveOrder)
				case "TextDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "TextInstructionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextInstructionType), stage, ident.Name, instanceName, preserveOrder)
				case "TextNameType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextNameType), stage, ident.Name, instanceName, preserveOrder)
				case "TextNoteType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextNoteType), stage, ident.Name, instanceName, preserveOrder)
				case "TextPhoneType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextPhoneType), stage, ident.Name, instanceName, preserveOrder)
				case "TextPropertyNameType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextPropertyNameType), stage, ident.Name, instanceName, preserveOrder)
				case "TextRemarkType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextRemarkType), stage, ident.Name, instanceName, preserveOrder)
				case "TextSIDSTARDesignatorType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TextSIDSTARDesignatorType), stage, ident.Name, instanceName, preserveOrder)
				case "TimeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(TimeType), stage, ident.Name, instanceName, preserveOrder)
				case "ValAlarmLimitType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValAlarmLimitType), stage, ident.Name, instanceName, preserveOrder)
				case "ValAngleType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValAngleType), stage, ident.Name, instanceName, preserveOrder)
				case "ValBearingType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValBearingType), stage, ident.Name, instanceName, preserveOrder)
				case "ValChannelNumberType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValChannelNumberType), stage, ident.Name, instanceName, preserveOrder)
				case "ValDepthType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValDepthType), stage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceSignedType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValDistanceSignedType), stage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValDistanceType), stage, ident.Name, instanceName, preserveOrder)
				case "ValDistanceVerticalType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValDistanceVerticalType), stage, ident.Name, instanceName, preserveOrder)
				case "ValDurationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValDurationType), stage, ident.Name, instanceName, preserveOrder)
				case "ValFLType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValFLType), stage, ident.Name, instanceName, preserveOrder)
				case "ValFrequencyType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValFrequencyType), stage, ident.Name, instanceName, preserveOrder)
				case "ValFrictionType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValFrictionType), stage, ident.Name, instanceName, preserveOrder)
				case "ValHexType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValHexType), stage, ident.Name, instanceName, preserveOrder)
				case "ValLCNType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValLCNType), stage, ident.Name, instanceName, preserveOrder)
				case "ValLightIntensityType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValLightIntensityType), stage, ident.Name, instanceName, preserveOrder)
				case "ValMagneticVariationChangeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValMagneticVariationChangeType), stage, ident.Name, instanceName, preserveOrder)
				case "ValMagneticVariationType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValMagneticVariationType), stage, ident.Name, instanceName, preserveOrder)
				case "ValPCNType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValPCNType), stage, ident.Name, instanceName, preserveOrder)
				case "ValPercentType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValPercentType), stage, ident.Name, instanceName, preserveOrder)
				case "ValPressureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValPressureType), stage, ident.Name, instanceName, preserveOrder)
				case "ValSlopeType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValSlopeType), stage, ident.Name, instanceName, preserveOrder)
				case "ValSpeedType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValSpeedType), stage, ident.Name, instanceName, preserveOrder)
				case "ValTemperatureType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValTemperatureType), stage, ident.Name, instanceName, preserveOrder)
				case "ValWeightType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(ValWeightType), stage, ident.Name, instanceName, preserveOrder)
				case "XHTMLType":
					identifierMap[ident.Name] = __gong__stageSetInit(new(XHTMLType), stage, ident.Name, instanceName, preserveOrder)
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
				case *CodeATCReportingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAerialRefuellingPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAerialRefuellingPrefixType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAerialRefuellingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftCategoryType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftEngineNumberType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftEngineType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftGroundServiceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftICAOType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftStandType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAircraftWingspanClassType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAirportHeliportCollocationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAirportHeliportDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAirportHeliportType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAirportWarningType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAirspaceActivityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAirspaceAggregationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAirspaceClassificationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAirspaceDependencyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAirspaceDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAirspacePointPositionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAirspacePointRoleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAirspaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAltitudeAdjustmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAltitudeUseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeApproachEquipmentAdditionalType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeApproachGuidanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeApproachLightingICAOType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeApproachLightingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeApproachPrefixType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeApproachType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeApprovalType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeApronElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeApronSectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeArcDirectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeArrestingGearEnergyAbsorbType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeArrestingGearEngageDeviceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAuralMorseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAuthorityRoleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAuthorityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeAviationStandardsType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeBearingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeBuoyDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeBuoyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeCardinalDirectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeCheckpointCategoryType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeColourType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeCommunicationChannelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeCommunicationDirectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeCommunicationModeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeComparisonType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeContaminationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeCourseQualityILSType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeCourseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDMEChannelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDMEType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDayType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDeclaredDistanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDesignStandardType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDesignatedPointDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDesignatedPointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDirectionReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDirectionTurnType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDirectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeDistanceIndicationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeEmissionBandType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeEquipmentAntiCollisionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeEquipmentUnavailableType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFacilityRankingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFinalGuidanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFireFightingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightDestinationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightOriginType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightPurposeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightRestrictionDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightRestrictionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightRuleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightStatusType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlightType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFlowConditionOperationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFreeFlightType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFrictionDeviceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFrictionEstimateType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeFuelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeGeoBorderType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeGradeSeparationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeGroundLightingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeGuidanceLineType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeHeightReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeHelicopterPerformanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeHoldingCategoryType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeHoldingUsageType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeHoldingUseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeIATAType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeICAOCountryType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeICAOType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeILSBackCourseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeIntegrityLevelILSType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeIntensityStandByType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLanguageType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLevelSeriesType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLevelTableDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLevelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLightHoldingPositionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLightIntensityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLightSourceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLightingJARType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLoadingBridgeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLocationQualifierType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeLogicalOperatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMLSAzimuthType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMLSChannelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMarkerBeaconSignalType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMarkingConditionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMarkingStyleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMeteoConditionsType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMilitaryOperationsType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMilitaryRoutePointType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMilitaryStatusType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMilitaryTrainingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMinimaFinalApproachPathType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMinimumAltitudeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMinimumHeightType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeMissedApproachType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNDBUsageType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNavaidDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNavaidPurposeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNavaidServiceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNavigationAreaRestrictionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNavigationAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNavigationEquipmentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNavigationSpecificationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNitrogenType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNorthReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeNotePurposeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeObstacleAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeObstacleAssessmentSurfaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeObstructionIdSurfaceZoneType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeOilType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeOperationAirportHeliportType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeOperationManoeuvringAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeOrganisationDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeOrganisationHierarchyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeOrganisationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeOxygenType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodePARType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodePCNMethodType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodePCNPavementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodePCNSubgradeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodePCNTyrePressureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodePassengerServiceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodePilotControlledLightingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodePositionInILSType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodePrimaryRadarType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeProcedureAvailabilityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeProcedureCodingStandardType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeProcedureDistanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeProcedureFixRoleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeProcedurePhaseType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeProtectAreaSectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRNPType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRVRReadingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRVSMPointRoleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRVSMType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRadarServiceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRadioEmissionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRadioFrequencyAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRadioSignalType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeReferencePathIdentifierType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeReferenceRoleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeReflectorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRelativePositionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRoadType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteAvailabilityType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteDesignatorLetterType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteDesignatorPrefixType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteDesignatorSuffixType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteIndicatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteNavigationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteOriginType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteSegmentPathType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRouteType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRuleProcedureTitleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRuleProcedureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRunwayElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRunwayMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRunwayPointRoleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRunwayProtectionAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRunwaySectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeRunwayType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSafeAltitudeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSegmentPathType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSegmentTerminationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeServiceATCType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeServiceATFMType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeServiceGroundControlType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeServiceInformationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeServiceSARType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSideType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSignalPerformanceILSType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSpecialDateType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSpecialNavigationChainDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSpecialNavigationStationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSpecialNavigationSystemType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSpeedReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeStandbyPowerType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeStatusAirportType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeStatusAirspaceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeStatusConstructionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeStatusNavaidType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeStatusOperationsType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeStatusServiceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSurfaceCompositionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSurfaceConditionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSurfacePreparationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeSystemActivationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTAAType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTACANChannelType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTLOFSectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTaxiwayElementType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTaxiwaySectionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTaxiwayType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTelecomNetworkType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTimeEventCombinationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTimeEventType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTimeReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTrajectoryType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeTransponderType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeUUIDType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeUnitDependencyType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeUnitType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeUpperAlphaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeUsageLimitationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVASISType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVORType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeValueInterpretationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVerticalDatumType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVerticalReferenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVerticalStructureMarkingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVerticalStructureMaterialType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVerticalStructureType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeVisualDockingGuidanceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeWakeTurbulenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeWorkAreaType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *CodeYesNoType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *DateMonthDayType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *DateTimeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *DateType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *DateYearType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *NoNumberType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *NoSequenceType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractInt(rhs)
					}
				case *TextAddressType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextInstructionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextNameType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextNoteType:
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
				case *TextPhoneType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextPropertyNameType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextRemarkType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TextSIDSTARDesignatorType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *TimeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValAlarmLimitType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValAngleType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValBearingType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValChannelNumberType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValDepthType:
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
				case *ValDistanceSignedType:
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
				case *ValDistanceType:
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
				case *ValDistanceVerticalType:
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
				case *ValDurationType:
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
				case *ValFLType:
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
				case *ValFrequencyType:
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
				case *ValFrictionType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValHexType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValLCNType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValLightIntensityType:
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
				case *ValMagneticVariationChangeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValMagneticVariationType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValPCNType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValPercentType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValPressureType:
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
				case *ValSlopeType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
					case "EnclosedText":
						inst.EnclosedText = GongExtractString(rhs)
					}
				case *ValSpeedType:
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
				case *ValTemperatureType:
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
				case *ValWeightType:
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
				case *XHTMLType:
					switch fieldName {
					case "Name":
						inst.Name = GongExtractString(rhs)
					case "NilReason":
						inst.NilReason = GongExtractString(rhs)
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
type CodeATCReportingTypeUnmarshaller struct{}

func (u *CodeATCReportingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeATCReportingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeATCReportingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeATCReportingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAerialRefuellingPointTypeUnmarshaller struct{}

func (u *CodeAerialRefuellingPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAerialRefuellingPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAerialRefuellingPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAerialRefuellingPointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAerialRefuellingPrefixTypeUnmarshaller struct{}

func (u *CodeAerialRefuellingPrefixTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAerialRefuellingPrefixType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAerialRefuellingPrefixTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAerialRefuellingPrefixType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAerialRefuellingTypeUnmarshaller struct{}

func (u *CodeAerialRefuellingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAerialRefuellingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAerialRefuellingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAerialRefuellingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftCategoryTypeUnmarshaller struct{}

func (u *CodeAircraftCategoryTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftCategoryType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftCategoryTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftCategoryType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftEngineNumberTypeUnmarshaller struct{}

func (u *CodeAircraftEngineNumberTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftEngineNumberType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftEngineNumberTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftEngineNumberType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftEngineTypeUnmarshaller struct{}

func (u *CodeAircraftEngineTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftEngineType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftEngineTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftEngineType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftGroundServiceTypeUnmarshaller struct{}

func (u *CodeAircraftGroundServiceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftGroundServiceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftGroundServiceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftGroundServiceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftICAOTypeUnmarshaller struct{}

func (u *CodeAircraftICAOTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftICAOType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftICAOTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftICAOType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftStandTypeUnmarshaller struct{}

func (u *CodeAircraftStandTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftStandType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftStandTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftStandType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftTypeUnmarshaller struct{}

func (u *CodeAircraftTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAircraftWingspanClassTypeUnmarshaller struct{}

func (u *CodeAircraftWingspanClassTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAircraftWingspanClassType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAircraftWingspanClassTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAircraftWingspanClassType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAirportHeliportCollocationTypeUnmarshaller struct{}

func (u *CodeAirportHeliportCollocationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAirportHeliportCollocationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAirportHeliportCollocationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAirportHeliportCollocationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAirportHeliportDesignatorTypeUnmarshaller struct{}

func (u *CodeAirportHeliportDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAirportHeliportDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAirportHeliportDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAirportHeliportDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAirportHeliportTypeUnmarshaller struct{}

func (u *CodeAirportHeliportTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAirportHeliportType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAirportHeliportTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAirportHeliportType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAirportWarningTypeUnmarshaller struct{}

func (u *CodeAirportWarningTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAirportWarningType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAirportWarningTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAirportWarningType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAirspaceActivityTypeUnmarshaller struct{}

func (u *CodeAirspaceActivityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAirspaceActivityType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAirspaceActivityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAirspaceActivityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAirspaceAggregationTypeUnmarshaller struct{}

func (u *CodeAirspaceAggregationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAirspaceAggregationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAirspaceAggregationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAirspaceAggregationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAirspaceClassificationTypeUnmarshaller struct{}

func (u *CodeAirspaceClassificationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAirspaceClassificationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAirspaceClassificationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAirspaceClassificationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAirspaceDependencyTypeUnmarshaller struct{}

func (u *CodeAirspaceDependencyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAirspaceDependencyType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAirspaceDependencyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAirspaceDependencyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAirspaceDesignatorTypeUnmarshaller struct{}

func (u *CodeAirspaceDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAirspaceDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAirspaceDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAirspaceDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAirspacePointPositionTypeUnmarshaller struct{}

func (u *CodeAirspacePointPositionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAirspacePointPositionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAirspacePointPositionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAirspacePointPositionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAirspacePointRoleTypeUnmarshaller struct{}

func (u *CodeAirspacePointRoleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAirspacePointRoleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAirspacePointRoleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAirspacePointRoleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAirspaceTypeUnmarshaller struct{}

func (u *CodeAirspaceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAirspaceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAirspaceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAirspaceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAltitudeAdjustmentTypeUnmarshaller struct{}

func (u *CodeAltitudeAdjustmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAltitudeAdjustmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAltitudeAdjustmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAltitudeAdjustmentType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAltitudeUseTypeUnmarshaller struct{}

func (u *CodeAltitudeUseTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAltitudeUseType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAltitudeUseTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAltitudeUseType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeApproachEquipmentAdditionalTypeUnmarshaller struct{}

func (u *CodeApproachEquipmentAdditionalTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeApproachEquipmentAdditionalType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeApproachEquipmentAdditionalTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeApproachEquipmentAdditionalType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeApproachGuidanceTypeUnmarshaller struct{}

func (u *CodeApproachGuidanceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeApproachGuidanceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeApproachGuidanceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeApproachGuidanceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeApproachLightingICAOTypeUnmarshaller struct{}

func (u *CodeApproachLightingICAOTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeApproachLightingICAOType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeApproachLightingICAOTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeApproachLightingICAOType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeApproachLightingTypeUnmarshaller struct{}

func (u *CodeApproachLightingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeApproachLightingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeApproachLightingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeApproachLightingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeApproachPrefixTypeUnmarshaller struct{}

func (u *CodeApproachPrefixTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeApproachPrefixType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeApproachPrefixTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeApproachPrefixType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeApproachTypeUnmarshaller struct{}

func (u *CodeApproachTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeApproachType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeApproachTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeApproachType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeApprovalTypeUnmarshaller struct{}

func (u *CodeApprovalTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeApprovalType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeApprovalTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeApprovalType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeApronElementTypeUnmarshaller struct{}

func (u *CodeApronElementTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeApronElementType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeApronElementTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeApronElementType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeApronSectionTypeUnmarshaller struct{}

func (u *CodeApronSectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeApronSectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeApronSectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeApronSectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeArcDirectionTypeUnmarshaller struct{}

func (u *CodeArcDirectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeArcDirectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeArcDirectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeArcDirectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeArrestingGearEnergyAbsorbTypeUnmarshaller struct{}

func (u *CodeArrestingGearEnergyAbsorbTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeArrestingGearEnergyAbsorbType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeArrestingGearEnergyAbsorbTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeArrestingGearEnergyAbsorbType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeArrestingGearEngageDeviceTypeUnmarshaller struct{}

func (u *CodeArrestingGearEngageDeviceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeArrestingGearEngageDeviceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeArrestingGearEngageDeviceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeArrestingGearEngageDeviceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAuralMorseTypeUnmarshaller struct{}

func (u *CodeAuralMorseTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAuralMorseType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAuralMorseTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAuralMorseType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAuthorityRoleTypeUnmarshaller struct{}

func (u *CodeAuthorityRoleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAuthorityRoleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAuthorityRoleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAuthorityRoleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAuthorityTypeUnmarshaller struct{}

func (u *CodeAuthorityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAuthorityType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAuthorityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAuthorityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeAviationStandardsTypeUnmarshaller struct{}

func (u *CodeAviationStandardsTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeAviationStandardsType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeAviationStandardsTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeAviationStandardsType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeBearingTypeUnmarshaller struct{}

func (u *CodeBearingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeBearingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeBearingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeBearingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeBuoyDesignatorTypeUnmarshaller struct{}

func (u *CodeBuoyDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeBuoyDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeBuoyDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeBuoyDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeBuoyTypeUnmarshaller struct{}

func (u *CodeBuoyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeBuoyType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeBuoyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeBuoyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeCardinalDirectionTypeUnmarshaller struct{}

func (u *CodeCardinalDirectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeCardinalDirectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeCardinalDirectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeCardinalDirectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeCheckpointCategoryTypeUnmarshaller struct{}

func (u *CodeCheckpointCategoryTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeCheckpointCategoryType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeCheckpointCategoryTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeCheckpointCategoryType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeColourTypeUnmarshaller struct{}

func (u *CodeColourTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeColourType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeColourTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeColourType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeCommunicationChannelTypeUnmarshaller struct{}

func (u *CodeCommunicationChannelTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeCommunicationChannelType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeCommunicationChannelTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeCommunicationChannelType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeCommunicationDirectionTypeUnmarshaller struct{}

func (u *CodeCommunicationDirectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeCommunicationDirectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeCommunicationDirectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeCommunicationDirectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeCommunicationModeTypeUnmarshaller struct{}

func (u *CodeCommunicationModeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeCommunicationModeType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeCommunicationModeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeCommunicationModeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeComparisonTypeUnmarshaller struct{}

func (u *CodeComparisonTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeComparisonType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeComparisonTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeComparisonType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeContaminationTypeUnmarshaller struct{}

func (u *CodeContaminationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeContaminationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeContaminationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeContaminationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeCourseQualityILSTypeUnmarshaller struct{}

func (u *CodeCourseQualityILSTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeCourseQualityILSType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeCourseQualityILSTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeCourseQualityILSType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeCourseTypeUnmarshaller struct{}

func (u *CodeCourseTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeCourseType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeCourseTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeCourseType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDMEChannelTypeUnmarshaller struct{}

func (u *CodeDMEChannelTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDMEChannelType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDMEChannelTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDMEChannelType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDMETypeUnmarshaller struct{}

func (u *CodeDMETypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDMEType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDMETypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDMEType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDayTypeUnmarshaller struct{}

func (u *CodeDayTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDayType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDayTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDayType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDeclaredDistanceTypeUnmarshaller struct{}

func (u *CodeDeclaredDistanceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDeclaredDistanceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDeclaredDistanceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDeclaredDistanceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDesignStandardTypeUnmarshaller struct{}

func (u *CodeDesignStandardTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDesignStandardType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDesignStandardTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDesignStandardType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDesignatedPointDesignatorTypeUnmarshaller struct{}

func (u *CodeDesignatedPointDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDesignatedPointDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDesignatedPointDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDesignatedPointDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDesignatedPointTypeUnmarshaller struct{}

func (u *CodeDesignatedPointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDesignatedPointType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDesignatedPointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDesignatedPointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDirectionReferenceTypeUnmarshaller struct{}

func (u *CodeDirectionReferenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDirectionReferenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDirectionReferenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDirectionReferenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDirectionTurnTypeUnmarshaller struct{}

func (u *CodeDirectionTurnTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDirectionTurnType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDirectionTurnTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDirectionTurnType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDirectionTypeUnmarshaller struct{}

func (u *CodeDirectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDirectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDirectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDirectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeDistanceIndicationTypeUnmarshaller struct{}

func (u *CodeDistanceIndicationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeDistanceIndicationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeDistanceIndicationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeDistanceIndicationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeEmissionBandTypeUnmarshaller struct{}

func (u *CodeEmissionBandTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeEmissionBandType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeEmissionBandTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeEmissionBandType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeEquipmentAntiCollisionTypeUnmarshaller struct{}

func (u *CodeEquipmentAntiCollisionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeEquipmentAntiCollisionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeEquipmentAntiCollisionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeEquipmentAntiCollisionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeEquipmentUnavailableTypeUnmarshaller struct{}

func (u *CodeEquipmentUnavailableTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeEquipmentUnavailableType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeEquipmentUnavailableTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeEquipmentUnavailableType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFacilityRankingTypeUnmarshaller struct{}

func (u *CodeFacilityRankingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFacilityRankingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFacilityRankingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFacilityRankingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFinalGuidanceTypeUnmarshaller struct{}

func (u *CodeFinalGuidanceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFinalGuidanceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFinalGuidanceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFinalGuidanceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFireFightingTypeUnmarshaller struct{}

func (u *CodeFireFightingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFireFightingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFireFightingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFireFightingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightDestinationTypeUnmarshaller struct{}

func (u *CodeFlightDestinationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightDestinationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightDestinationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightDestinationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightOriginTypeUnmarshaller struct{}

func (u *CodeFlightOriginTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightOriginType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightOriginTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightOriginType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightPurposeTypeUnmarshaller struct{}

func (u *CodeFlightPurposeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightPurposeType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightPurposeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightPurposeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightRestrictionDesignatorTypeUnmarshaller struct{}

func (u *CodeFlightRestrictionDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightRestrictionDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightRestrictionDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightRestrictionDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightRestrictionTypeUnmarshaller struct{}

func (u *CodeFlightRestrictionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightRestrictionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightRestrictionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightRestrictionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightRuleTypeUnmarshaller struct{}

func (u *CodeFlightRuleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightRuleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightRuleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightRuleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightStatusTypeUnmarshaller struct{}

func (u *CodeFlightStatusTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightStatusType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightStatusTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightStatusType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlightTypeUnmarshaller struct{}

func (u *CodeFlightTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlightType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlightTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlightType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFlowConditionOperationTypeUnmarshaller struct{}

func (u *CodeFlowConditionOperationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFlowConditionOperationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFlowConditionOperationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFlowConditionOperationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFreeFlightTypeUnmarshaller struct{}

func (u *CodeFreeFlightTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFreeFlightType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFreeFlightTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFreeFlightType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFrictionDeviceTypeUnmarshaller struct{}

func (u *CodeFrictionDeviceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFrictionDeviceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFrictionDeviceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFrictionDeviceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFrictionEstimateTypeUnmarshaller struct{}

func (u *CodeFrictionEstimateTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFrictionEstimateType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFrictionEstimateTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFrictionEstimateType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeFuelTypeUnmarshaller struct{}

func (u *CodeFuelTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeFuelType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeFuelTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeFuelType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeGeoBorderTypeUnmarshaller struct{}

func (u *CodeGeoBorderTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeGeoBorderType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeGeoBorderTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeGeoBorderType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeGradeSeparationTypeUnmarshaller struct{}

func (u *CodeGradeSeparationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeGradeSeparationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeGradeSeparationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeGradeSeparationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeGroundLightingTypeUnmarshaller struct{}

func (u *CodeGroundLightingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeGroundLightingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeGroundLightingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeGroundLightingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeGuidanceLineTypeUnmarshaller struct{}

func (u *CodeGuidanceLineTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeGuidanceLineType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeGuidanceLineTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeGuidanceLineType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeHeightReferenceTypeUnmarshaller struct{}

func (u *CodeHeightReferenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeHeightReferenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeHeightReferenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeHeightReferenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeHelicopterPerformanceTypeUnmarshaller struct{}

func (u *CodeHelicopterPerformanceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeHelicopterPerformanceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeHelicopterPerformanceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeHelicopterPerformanceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeHoldingCategoryTypeUnmarshaller struct{}

func (u *CodeHoldingCategoryTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeHoldingCategoryType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeHoldingCategoryTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeHoldingCategoryType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeHoldingUsageTypeUnmarshaller struct{}

func (u *CodeHoldingUsageTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeHoldingUsageType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeHoldingUsageTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeHoldingUsageType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeHoldingUseTypeUnmarshaller struct{}

func (u *CodeHoldingUseTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeHoldingUseType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeHoldingUseTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeHoldingUseType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeIATATypeUnmarshaller struct{}

func (u *CodeIATATypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeIATAType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeIATATypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeIATAType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeICAOCountryTypeUnmarshaller struct{}

func (u *CodeICAOCountryTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeICAOCountryType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeICAOCountryTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeICAOCountryType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeICAOTypeUnmarshaller struct{}

func (u *CodeICAOTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeICAOType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeICAOTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeICAOType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeILSBackCourseTypeUnmarshaller struct{}

func (u *CodeILSBackCourseTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeILSBackCourseType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeILSBackCourseTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeILSBackCourseType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeIntegrityLevelILSTypeUnmarshaller struct{}

func (u *CodeIntegrityLevelILSTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeIntegrityLevelILSType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeIntegrityLevelILSTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeIntegrityLevelILSType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeIntensityStandByTypeUnmarshaller struct{}

func (u *CodeIntensityStandByTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeIntensityStandByType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeIntensityStandByTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeIntensityStandByType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLanguageTypeUnmarshaller struct{}

func (u *CodeLanguageTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLanguageType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLanguageTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLanguageType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLevelSeriesTypeUnmarshaller struct{}

func (u *CodeLevelSeriesTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLevelSeriesType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLevelSeriesTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLevelSeriesType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLevelTableDesignatorTypeUnmarshaller struct{}

func (u *CodeLevelTableDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLevelTableDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLevelTableDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLevelTableDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLevelTypeUnmarshaller struct{}

func (u *CodeLevelTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLevelType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLevelTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLevelType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLightHoldingPositionTypeUnmarshaller struct{}

func (u *CodeLightHoldingPositionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLightHoldingPositionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLightHoldingPositionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLightHoldingPositionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLightIntensityTypeUnmarshaller struct{}

func (u *CodeLightIntensityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLightIntensityType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLightIntensityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLightIntensityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLightSourceTypeUnmarshaller struct{}

func (u *CodeLightSourceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLightSourceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLightSourceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLightSourceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLightingJARTypeUnmarshaller struct{}

func (u *CodeLightingJARTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLightingJARType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLightingJARTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLightingJARType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLoadingBridgeTypeUnmarshaller struct{}

func (u *CodeLoadingBridgeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLoadingBridgeType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLoadingBridgeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLoadingBridgeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLocationQualifierTypeUnmarshaller struct{}

func (u *CodeLocationQualifierTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLocationQualifierType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLocationQualifierTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLocationQualifierType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeLogicalOperatorTypeUnmarshaller struct{}

func (u *CodeLogicalOperatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeLogicalOperatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeLogicalOperatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeLogicalOperatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMLSAzimuthTypeUnmarshaller struct{}

func (u *CodeMLSAzimuthTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMLSAzimuthType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMLSAzimuthTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMLSAzimuthType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMLSChannelTypeUnmarshaller struct{}

func (u *CodeMLSChannelTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMLSChannelType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMLSChannelTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMLSChannelType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMarkerBeaconSignalTypeUnmarshaller struct{}

func (u *CodeMarkerBeaconSignalTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMarkerBeaconSignalType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMarkerBeaconSignalTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMarkerBeaconSignalType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMarkingConditionTypeUnmarshaller struct{}

func (u *CodeMarkingConditionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMarkingConditionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMarkingConditionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMarkingConditionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMarkingStyleTypeUnmarshaller struct{}

func (u *CodeMarkingStyleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMarkingStyleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMarkingStyleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMarkingStyleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMeteoConditionsTypeUnmarshaller struct{}

func (u *CodeMeteoConditionsTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMeteoConditionsType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMeteoConditionsTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMeteoConditionsType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMilitaryOperationsTypeUnmarshaller struct{}

func (u *CodeMilitaryOperationsTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMilitaryOperationsType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMilitaryOperationsTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMilitaryOperationsType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMilitaryRoutePointTypeUnmarshaller struct{}

func (u *CodeMilitaryRoutePointTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMilitaryRoutePointType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMilitaryRoutePointTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMilitaryRoutePointType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMilitaryStatusTypeUnmarshaller struct{}

func (u *CodeMilitaryStatusTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMilitaryStatusType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMilitaryStatusTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMilitaryStatusType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMilitaryTrainingTypeUnmarshaller struct{}

func (u *CodeMilitaryTrainingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMilitaryTrainingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMilitaryTrainingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMilitaryTrainingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMinimaFinalApproachPathTypeUnmarshaller struct{}

func (u *CodeMinimaFinalApproachPathTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMinimaFinalApproachPathType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMinimaFinalApproachPathTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMinimaFinalApproachPathType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMinimumAltitudeTypeUnmarshaller struct{}

func (u *CodeMinimumAltitudeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMinimumAltitudeType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMinimumAltitudeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMinimumAltitudeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMinimumHeightTypeUnmarshaller struct{}

func (u *CodeMinimumHeightTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMinimumHeightType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMinimumHeightTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMinimumHeightType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeMissedApproachTypeUnmarshaller struct{}

func (u *CodeMissedApproachTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeMissedApproachType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeMissedApproachTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeMissedApproachType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNDBUsageTypeUnmarshaller struct{}

func (u *CodeNDBUsageTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNDBUsageType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNDBUsageTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNDBUsageType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNavaidDesignatorTypeUnmarshaller struct{}

func (u *CodeNavaidDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNavaidDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNavaidDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNavaidDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNavaidPurposeTypeUnmarshaller struct{}

func (u *CodeNavaidPurposeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNavaidPurposeType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNavaidPurposeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNavaidPurposeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNavaidServiceTypeUnmarshaller struct{}

func (u *CodeNavaidServiceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNavaidServiceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNavaidServiceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNavaidServiceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNavigationAreaRestrictionTypeUnmarshaller struct{}

func (u *CodeNavigationAreaRestrictionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNavigationAreaRestrictionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNavigationAreaRestrictionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNavigationAreaRestrictionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNavigationAreaTypeUnmarshaller struct{}

func (u *CodeNavigationAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNavigationAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNavigationAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNavigationAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNavigationEquipmentTypeUnmarshaller struct{}

func (u *CodeNavigationEquipmentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNavigationEquipmentType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNavigationEquipmentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNavigationEquipmentType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNavigationSpecificationTypeUnmarshaller struct{}

func (u *CodeNavigationSpecificationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNavigationSpecificationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNavigationSpecificationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNavigationSpecificationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNitrogenTypeUnmarshaller struct{}

func (u *CodeNitrogenTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNitrogenType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNitrogenTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNitrogenType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNorthReferenceTypeUnmarshaller struct{}

func (u *CodeNorthReferenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNorthReferenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNorthReferenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNorthReferenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeNotePurposeTypeUnmarshaller struct{}

func (u *CodeNotePurposeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeNotePurposeType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeNotePurposeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeNotePurposeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeObstacleAreaTypeUnmarshaller struct{}

func (u *CodeObstacleAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeObstacleAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeObstacleAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeObstacleAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeObstacleAssessmentSurfaceTypeUnmarshaller struct{}

func (u *CodeObstacleAssessmentSurfaceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeObstacleAssessmentSurfaceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeObstacleAssessmentSurfaceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeObstacleAssessmentSurfaceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeObstructionIdSurfaceZoneTypeUnmarshaller struct{}

func (u *CodeObstructionIdSurfaceZoneTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeObstructionIdSurfaceZoneType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeObstructionIdSurfaceZoneTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeObstructionIdSurfaceZoneType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeOilTypeUnmarshaller struct{}

func (u *CodeOilTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeOilType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeOilTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeOilType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeOperationAirportHeliportTypeUnmarshaller struct{}

func (u *CodeOperationAirportHeliportTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeOperationAirportHeliportType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeOperationAirportHeliportTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeOperationAirportHeliportType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeOperationManoeuvringAreaTypeUnmarshaller struct{}

func (u *CodeOperationManoeuvringAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeOperationManoeuvringAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeOperationManoeuvringAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeOperationManoeuvringAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeOrganisationDesignatorTypeUnmarshaller struct{}

func (u *CodeOrganisationDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeOrganisationDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeOrganisationDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeOrganisationDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeOrganisationHierarchyTypeUnmarshaller struct{}

func (u *CodeOrganisationHierarchyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeOrganisationHierarchyType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeOrganisationHierarchyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeOrganisationHierarchyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeOrganisationTypeUnmarshaller struct{}

func (u *CodeOrganisationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeOrganisationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeOrganisationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeOrganisationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeOxygenTypeUnmarshaller struct{}

func (u *CodeOxygenTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeOxygenType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeOxygenTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeOxygenType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodePARTypeUnmarshaller struct{}

func (u *CodePARTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodePARType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodePARTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodePARType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodePCNMethodTypeUnmarshaller struct{}

func (u *CodePCNMethodTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodePCNMethodType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodePCNMethodTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodePCNMethodType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodePCNPavementTypeUnmarshaller struct{}

func (u *CodePCNPavementTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodePCNPavementType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodePCNPavementTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodePCNPavementType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodePCNSubgradeTypeUnmarshaller struct{}

func (u *CodePCNSubgradeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodePCNSubgradeType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodePCNSubgradeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodePCNSubgradeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodePCNTyrePressureTypeUnmarshaller struct{}

func (u *CodePCNTyrePressureTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodePCNTyrePressureType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodePCNTyrePressureTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodePCNTyrePressureType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodePassengerServiceTypeUnmarshaller struct{}

func (u *CodePassengerServiceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodePassengerServiceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodePassengerServiceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodePassengerServiceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodePilotControlledLightingTypeUnmarshaller struct{}

func (u *CodePilotControlledLightingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodePilotControlledLightingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodePilotControlledLightingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodePilotControlledLightingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodePositionInILSTypeUnmarshaller struct{}

func (u *CodePositionInILSTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodePositionInILSType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodePositionInILSTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodePositionInILSType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodePrimaryRadarTypeUnmarshaller struct{}

func (u *CodePrimaryRadarTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodePrimaryRadarType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodePrimaryRadarTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodePrimaryRadarType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeProcedureAvailabilityTypeUnmarshaller struct{}

func (u *CodeProcedureAvailabilityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeProcedureAvailabilityType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeProcedureAvailabilityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeProcedureAvailabilityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeProcedureCodingStandardTypeUnmarshaller struct{}

func (u *CodeProcedureCodingStandardTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeProcedureCodingStandardType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeProcedureCodingStandardTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeProcedureCodingStandardType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeProcedureDistanceTypeUnmarshaller struct{}

func (u *CodeProcedureDistanceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeProcedureDistanceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeProcedureDistanceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeProcedureDistanceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeProcedureFixRoleTypeUnmarshaller struct{}

func (u *CodeProcedureFixRoleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeProcedureFixRoleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeProcedureFixRoleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeProcedureFixRoleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeProcedurePhaseTypeUnmarshaller struct{}

func (u *CodeProcedurePhaseTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeProcedurePhaseType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeProcedurePhaseTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeProcedurePhaseType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeProtectAreaSectionTypeUnmarshaller struct{}

func (u *CodeProtectAreaSectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeProtectAreaSectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeProtectAreaSectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeProtectAreaSectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRNPTypeUnmarshaller struct{}

func (u *CodeRNPTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRNPType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRNPTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRNPType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRVRReadingTypeUnmarshaller struct{}

func (u *CodeRVRReadingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRVRReadingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRVRReadingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRVRReadingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRVSMPointRoleTypeUnmarshaller struct{}

func (u *CodeRVSMPointRoleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRVSMPointRoleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRVSMPointRoleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRVSMPointRoleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRVSMTypeUnmarshaller struct{}

func (u *CodeRVSMTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRVSMType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRVSMTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRVSMType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRadarServiceTypeUnmarshaller struct{}

func (u *CodeRadarServiceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRadarServiceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRadarServiceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRadarServiceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRadioEmissionTypeUnmarshaller struct{}

func (u *CodeRadioEmissionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRadioEmissionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRadioEmissionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRadioEmissionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRadioFrequencyAreaTypeUnmarshaller struct{}

func (u *CodeRadioFrequencyAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRadioFrequencyAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRadioFrequencyAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRadioFrequencyAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRadioSignalTypeUnmarshaller struct{}

func (u *CodeRadioSignalTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRadioSignalType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRadioSignalTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRadioSignalType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeReferencePathIdentifierTypeUnmarshaller struct{}

func (u *CodeReferencePathIdentifierTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeReferencePathIdentifierType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeReferencePathIdentifierTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeReferencePathIdentifierType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeReferenceRoleTypeUnmarshaller struct{}

func (u *CodeReferenceRoleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeReferenceRoleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeReferenceRoleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeReferenceRoleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeReflectorTypeUnmarshaller struct{}

func (u *CodeReflectorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeReflectorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeReflectorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeReflectorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRelativePositionTypeUnmarshaller struct{}

func (u *CodeRelativePositionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRelativePositionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRelativePositionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRelativePositionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRoadTypeUnmarshaller struct{}

func (u *CodeRoadTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRoadType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRoadTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRoadType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteAvailabilityTypeUnmarshaller struct{}

func (u *CodeRouteAvailabilityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteAvailabilityType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteAvailabilityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteAvailabilityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteDesignatorLetterTypeUnmarshaller struct{}

func (u *CodeRouteDesignatorLetterTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteDesignatorLetterType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteDesignatorLetterTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteDesignatorLetterType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteDesignatorPrefixTypeUnmarshaller struct{}

func (u *CodeRouteDesignatorPrefixTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteDesignatorPrefixType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteDesignatorPrefixTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteDesignatorPrefixType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteDesignatorSuffixTypeUnmarshaller struct{}

func (u *CodeRouteDesignatorSuffixTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteDesignatorSuffixType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteDesignatorSuffixTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteDesignatorSuffixType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteIndicatorTypeUnmarshaller struct{}

func (u *CodeRouteIndicatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteIndicatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteIndicatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteIndicatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteNavigationTypeUnmarshaller struct{}

func (u *CodeRouteNavigationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteNavigationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteNavigationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteNavigationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteOriginTypeUnmarshaller struct{}

func (u *CodeRouteOriginTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteOriginType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteOriginTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteOriginType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteSegmentPathTypeUnmarshaller struct{}

func (u *CodeRouteSegmentPathTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteSegmentPathType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteSegmentPathTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteSegmentPathType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRouteTypeUnmarshaller struct{}

func (u *CodeRouteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRouteType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRouteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRouteType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRuleProcedureTitleTypeUnmarshaller struct{}

func (u *CodeRuleProcedureTitleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRuleProcedureTitleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRuleProcedureTitleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRuleProcedureTitleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRuleProcedureTypeUnmarshaller struct{}

func (u *CodeRuleProcedureTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRuleProcedureType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRuleProcedureTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRuleProcedureType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRunwayElementTypeUnmarshaller struct{}

func (u *CodeRunwayElementTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRunwayElementType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRunwayElementTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRunwayElementType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRunwayMarkingTypeUnmarshaller struct{}

func (u *CodeRunwayMarkingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRunwayMarkingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRunwayMarkingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRunwayMarkingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRunwayPointRoleTypeUnmarshaller struct{}

func (u *CodeRunwayPointRoleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRunwayPointRoleType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRunwayPointRoleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRunwayPointRoleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRunwayProtectionAreaTypeUnmarshaller struct{}

func (u *CodeRunwayProtectionAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRunwayProtectionAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRunwayProtectionAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRunwayProtectionAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRunwaySectionTypeUnmarshaller struct{}

func (u *CodeRunwaySectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRunwaySectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRunwaySectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRunwaySectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeRunwayTypeUnmarshaller struct{}

func (u *CodeRunwayTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeRunwayType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeRunwayTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeRunwayType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSafeAltitudeTypeUnmarshaller struct{}

func (u *CodeSafeAltitudeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSafeAltitudeType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSafeAltitudeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSafeAltitudeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSegmentPathTypeUnmarshaller struct{}

func (u *CodeSegmentPathTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSegmentPathType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSegmentPathTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSegmentPathType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSegmentTerminationTypeUnmarshaller struct{}

func (u *CodeSegmentTerminationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSegmentTerminationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSegmentTerminationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSegmentTerminationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeServiceATCTypeUnmarshaller struct{}

func (u *CodeServiceATCTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeServiceATCType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeServiceATCTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeServiceATCType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeServiceATFMTypeUnmarshaller struct{}

func (u *CodeServiceATFMTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeServiceATFMType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeServiceATFMTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeServiceATFMType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeServiceGroundControlTypeUnmarshaller struct{}

func (u *CodeServiceGroundControlTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeServiceGroundControlType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeServiceGroundControlTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeServiceGroundControlType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeServiceInformationTypeUnmarshaller struct{}

func (u *CodeServiceInformationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeServiceInformationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeServiceInformationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeServiceInformationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeServiceSARTypeUnmarshaller struct{}

func (u *CodeServiceSARTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeServiceSARType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeServiceSARTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeServiceSARType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSideTypeUnmarshaller struct{}

func (u *CodeSideTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSideType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSideTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSideType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSignalPerformanceILSTypeUnmarshaller struct{}

func (u *CodeSignalPerformanceILSTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSignalPerformanceILSType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSignalPerformanceILSTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSignalPerformanceILSType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSpecialDateTypeUnmarshaller struct{}

func (u *CodeSpecialDateTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSpecialDateType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSpecialDateTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSpecialDateType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSpecialNavigationChainDesignatorTypeUnmarshaller struct{}

func (u *CodeSpecialNavigationChainDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSpecialNavigationChainDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSpecialNavigationChainDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSpecialNavigationChainDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSpecialNavigationStationTypeUnmarshaller struct{}

func (u *CodeSpecialNavigationStationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSpecialNavigationStationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSpecialNavigationStationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSpecialNavigationStationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSpecialNavigationSystemTypeUnmarshaller struct{}

func (u *CodeSpecialNavigationSystemTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSpecialNavigationSystemType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSpecialNavigationSystemTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSpecialNavigationSystemType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSpeedReferenceTypeUnmarshaller struct{}

func (u *CodeSpeedReferenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSpeedReferenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSpeedReferenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSpeedReferenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeStandbyPowerTypeUnmarshaller struct{}

func (u *CodeStandbyPowerTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeStandbyPowerType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeStandbyPowerTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeStandbyPowerType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeStatusAirportTypeUnmarshaller struct{}

func (u *CodeStatusAirportTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeStatusAirportType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeStatusAirportTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeStatusAirportType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeStatusAirspaceTypeUnmarshaller struct{}

func (u *CodeStatusAirspaceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeStatusAirspaceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeStatusAirspaceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeStatusAirspaceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeStatusConstructionTypeUnmarshaller struct{}

func (u *CodeStatusConstructionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeStatusConstructionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeStatusConstructionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeStatusConstructionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeStatusNavaidTypeUnmarshaller struct{}

func (u *CodeStatusNavaidTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeStatusNavaidType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeStatusNavaidTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeStatusNavaidType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeStatusOperationsTypeUnmarshaller struct{}

func (u *CodeStatusOperationsTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeStatusOperationsType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeStatusOperationsTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeStatusOperationsType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeStatusServiceTypeUnmarshaller struct{}

func (u *CodeStatusServiceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeStatusServiceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeStatusServiceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeStatusServiceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSurfaceCompositionTypeUnmarshaller struct{}

func (u *CodeSurfaceCompositionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSurfaceCompositionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSurfaceCompositionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSurfaceCompositionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSurfaceConditionTypeUnmarshaller struct{}

func (u *CodeSurfaceConditionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSurfaceConditionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSurfaceConditionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSurfaceConditionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSurfacePreparationTypeUnmarshaller struct{}

func (u *CodeSurfacePreparationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSurfacePreparationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSurfacePreparationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSurfacePreparationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeSystemActivationTypeUnmarshaller struct{}

func (u *CodeSystemActivationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeSystemActivationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeSystemActivationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeSystemActivationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTAATypeUnmarshaller struct{}

func (u *CodeTAATypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTAAType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTAATypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTAAType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTACANChannelTypeUnmarshaller struct{}

func (u *CodeTACANChannelTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTACANChannelType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTACANChannelTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTACANChannelType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTLOFSectionTypeUnmarshaller struct{}

func (u *CodeTLOFSectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTLOFSectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTLOFSectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTLOFSectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTaxiwayElementTypeUnmarshaller struct{}

func (u *CodeTaxiwayElementTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTaxiwayElementType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTaxiwayElementTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTaxiwayElementType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTaxiwaySectionTypeUnmarshaller struct{}

func (u *CodeTaxiwaySectionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTaxiwaySectionType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTaxiwaySectionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTaxiwaySectionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTaxiwayTypeUnmarshaller struct{}

func (u *CodeTaxiwayTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTaxiwayType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTaxiwayTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTaxiwayType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTelecomNetworkTypeUnmarshaller struct{}

func (u *CodeTelecomNetworkTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTelecomNetworkType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTelecomNetworkTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTelecomNetworkType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTimeEventCombinationTypeUnmarshaller struct{}

func (u *CodeTimeEventCombinationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTimeEventCombinationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTimeEventCombinationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTimeEventCombinationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTimeEventTypeUnmarshaller struct{}

func (u *CodeTimeEventTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTimeEventType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTimeEventTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTimeEventType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTimeReferenceTypeUnmarshaller struct{}

func (u *CodeTimeReferenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTimeReferenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTimeReferenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTimeReferenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTrajectoryTypeUnmarshaller struct{}

func (u *CodeTrajectoryTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTrajectoryType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTrajectoryTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTrajectoryType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeTransponderTypeUnmarshaller struct{}

func (u *CodeTransponderTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeTransponderType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeTransponderTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeTransponderType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeUUIDTypeUnmarshaller struct{}

func (u *CodeUUIDTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeUUIDType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeUUIDTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeUUIDType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeUnitDependencyTypeUnmarshaller struct{}

func (u *CodeUnitDependencyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeUnitDependencyType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeUnitDependencyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeUnitDependencyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeUnitTypeUnmarshaller struct{}

func (u *CodeUnitTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeUnitType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeUnitTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeUnitType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeUpperAlphaTypeUnmarshaller struct{}

func (u *CodeUpperAlphaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeUpperAlphaType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeUpperAlphaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeUpperAlphaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeUsageLimitationTypeUnmarshaller struct{}

func (u *CodeUsageLimitationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeUsageLimitationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeUsageLimitationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeUsageLimitationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVASISTypeUnmarshaller struct{}

func (u *CodeVASISTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVASISType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVASISTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVASISType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVORTypeUnmarshaller struct{}

func (u *CodeVORTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVORType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVORTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVORType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeValueInterpretationTypeUnmarshaller struct{}

func (u *CodeValueInterpretationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeValueInterpretationType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeValueInterpretationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeValueInterpretationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVerticalDatumTypeUnmarshaller struct{}

func (u *CodeVerticalDatumTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVerticalDatumType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVerticalDatumTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVerticalDatumType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVerticalReferenceTypeUnmarshaller struct{}

func (u *CodeVerticalReferenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVerticalReferenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVerticalReferenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVerticalReferenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVerticalStructureMarkingTypeUnmarshaller struct{}

func (u *CodeVerticalStructureMarkingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVerticalStructureMarkingType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVerticalStructureMarkingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVerticalStructureMarkingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVerticalStructureMaterialTypeUnmarshaller struct{}

func (u *CodeVerticalStructureMaterialTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVerticalStructureMaterialType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVerticalStructureMaterialTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVerticalStructureMaterialType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVerticalStructureTypeUnmarshaller struct{}

func (u *CodeVerticalStructureTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVerticalStructureType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVerticalStructureTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVerticalStructureType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeVisualDockingGuidanceTypeUnmarshaller struct{}

func (u *CodeVisualDockingGuidanceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeVisualDockingGuidanceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeVisualDockingGuidanceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeVisualDockingGuidanceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeWakeTurbulenceTypeUnmarshaller struct{}

func (u *CodeWakeTurbulenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeWakeTurbulenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeWakeTurbulenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeWakeTurbulenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeWorkAreaTypeUnmarshaller struct{}

func (u *CodeWorkAreaTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeWorkAreaType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeWorkAreaTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeWorkAreaType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type CodeYesNoTypeUnmarshaller struct{}

func (u *CodeYesNoTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(CodeYesNoType), stage, identifier, instanceName, preserveOrder)
}

func (u *CodeYesNoTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*CodeYesNoType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type DateMonthDayTypeUnmarshaller struct{}

func (u *DateMonthDayTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DateMonthDayType), stage, identifier, instanceName, preserveOrder)
}

func (u *DateMonthDayTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DateMonthDayType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type DateTimeTypeUnmarshaller struct{}

func (u *DateTimeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DateTimeType), stage, identifier, instanceName, preserveOrder)
}

func (u *DateTimeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DateTimeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type DateTypeUnmarshaller struct{}

func (u *DateTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DateType), stage, identifier, instanceName, preserveOrder)
}

func (u *DateTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DateType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type DateYearTypeUnmarshaller struct{}

func (u *DateYearTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(DateYearType), stage, identifier, instanceName, preserveOrder)
}

func (u *DateYearTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*DateYearType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type NoNumberTypeUnmarshaller struct{}

func (u *NoNumberTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoNumberType), stage, identifier, instanceName, preserveOrder)
}

func (u *NoNumberTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoNumberType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type NoSequenceTypeUnmarshaller struct{}

func (u *NoSequenceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(NoSequenceType), stage, identifier, instanceName, preserveOrder)
}

func (u *NoSequenceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*NoSequenceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type TextAddressTypeUnmarshaller struct{}

func (u *TextAddressTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextAddressType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextAddressTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextAddressType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextDesignatorTypeUnmarshaller struct{}

func (u *TextDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextInstructionTypeUnmarshaller struct{}

func (u *TextInstructionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextInstructionType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextInstructionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextInstructionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextNameTypeUnmarshaller struct{}

func (u *TextNameTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextNameType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextNameTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextNameType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextNoteTypeUnmarshaller struct{}

func (u *TextNoteTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextNoteType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextNoteTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextNoteType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "Lang":
		instance.Lang = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextPhoneTypeUnmarshaller struct{}

func (u *TextPhoneTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextPhoneType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextPhoneTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextPhoneType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextPropertyNameTypeUnmarshaller struct{}

func (u *TextPropertyNameTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextPropertyNameType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextPropertyNameTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextPropertyNameType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextRemarkTypeUnmarshaller struct{}

func (u *TextRemarkTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextRemarkType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextRemarkTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextRemarkType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TextSIDSTARDesignatorTypeUnmarshaller struct{}

func (u *TextSIDSTARDesignatorTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TextSIDSTARDesignatorType), stage, identifier, instanceName, preserveOrder)
}

func (u *TextSIDSTARDesignatorTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TextSIDSTARDesignatorType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type TimeTypeUnmarshaller struct{}

func (u *TimeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(TimeType), stage, identifier, instanceName, preserveOrder)
}

func (u *TimeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*TimeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValAlarmLimitTypeUnmarshaller struct{}

func (u *ValAlarmLimitTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValAlarmLimitType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValAlarmLimitTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValAlarmLimitType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValAngleTypeUnmarshaller struct{}

func (u *ValAngleTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValAngleType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValAngleTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValAngleType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValBearingTypeUnmarshaller struct{}

func (u *ValBearingTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValBearingType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValBearingTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValBearingType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValChannelNumberTypeUnmarshaller struct{}

func (u *ValChannelNumberTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValChannelNumberType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValChannelNumberTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValChannelNumberType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValDepthTypeUnmarshaller struct{}

func (u *ValDepthTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValDepthType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValDepthTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValDepthType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValDistanceSignedTypeUnmarshaller struct{}

func (u *ValDistanceSignedTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValDistanceSignedType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValDistanceSignedTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValDistanceSignedType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValDistanceTypeUnmarshaller struct{}

func (u *ValDistanceTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValDistanceType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValDistanceTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValDistanceType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValDistanceVerticalTypeUnmarshaller struct{}

func (u *ValDistanceVerticalTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValDistanceVerticalType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValDistanceVerticalTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValDistanceVerticalType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValDurationTypeUnmarshaller struct{}

func (u *ValDurationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValDurationType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValDurationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValDurationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValFLTypeUnmarshaller struct{}

func (u *ValFLTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValFLType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValFLTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValFLType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractInt(valueExpr)
	}
	return nil
}

type ValFrequencyTypeUnmarshaller struct{}

func (u *ValFrequencyTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValFrequencyType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValFrequencyTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValFrequencyType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValFrictionTypeUnmarshaller struct{}

func (u *ValFrictionTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValFrictionType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValFrictionTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValFrictionType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValHexTypeUnmarshaller struct{}

func (u *ValHexTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValHexType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValHexTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValHexType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValLCNTypeUnmarshaller struct{}

func (u *ValLCNTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValLCNType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValLCNTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValLCNType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValLightIntensityTypeUnmarshaller struct{}

func (u *ValLightIntensityTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValLightIntensityType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValLightIntensityTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValLightIntensityType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValMagneticVariationChangeTypeUnmarshaller struct{}

func (u *ValMagneticVariationChangeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValMagneticVariationChangeType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValMagneticVariationChangeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValMagneticVariationChangeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValMagneticVariationTypeUnmarshaller struct{}

func (u *ValMagneticVariationTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValMagneticVariationType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValMagneticVariationTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValMagneticVariationType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValPCNTypeUnmarshaller struct{}

func (u *ValPCNTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValPCNType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValPCNTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValPCNType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValPercentTypeUnmarshaller struct{}

func (u *ValPercentTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValPercentType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValPercentTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValPercentType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValPressureTypeUnmarshaller struct{}

func (u *ValPressureTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValPressureType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValPressureTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValPressureType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValSlopeTypeUnmarshaller struct{}

func (u *ValSlopeTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValSlopeType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValSlopeTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValSlopeType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValSpeedTypeUnmarshaller struct{}

func (u *ValSpeedTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValSpeedType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValSpeedTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValSpeedType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValTemperatureTypeUnmarshaller struct{}

func (u *ValTemperatureTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValTemperatureType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValTemperatureTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValTemperatureType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type ValWeightTypeUnmarshaller struct{}

func (u *ValWeightTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(ValWeightType), stage, identifier, instanceName, preserveOrder)
}

func (u *ValWeightTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*ValWeightType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "Uom":
		instance.Uom = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	case "EnclosedText":
		instance.EnclosedText = GongExtractString(valueExpr)
	}
	return nil
}

type XHTMLTypeUnmarshaller struct{}

func (u *XHTMLTypeUnmarshaller) Initialize(stage *Stage, identifier string, instanceName string, preserveOrder bool) (GongstructIF, error) {
	return GongInitialize(new(XHTMLType), stage, identifier, instanceName, preserveOrder)
}

func (u *XHTMLTypeUnmarshaller) UnmarshallField(stage *Stage, i GongstructIF, fieldName string, valueExpr ast.Expr, identifierMap map[string]GongstructIF) error {
	instance := i.(*XHTMLType)
	_ = instance
	switch fieldName {
	// insertion point per field
	case "Name":
		instance.Name = GongExtractString(valueExpr)
	case "NilReason":
		instance.NilReason = GongExtractString(valueExpr)
	}
	return nil
}
